/*
 * Copyright (c) 2022. Veteran Software
 *
 * Discord API Wrapper - A custom wrapper for the Discord REST API developed for a proprietary project.
 *
 * This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public
 * License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later
 * version.
 *
 * This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied
 * warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License along with this program.
 * If not, see <http://www.gnu.org/licenses/>.
 */

/*
 * Copyright (c) 2022. Veteran Software
 *
 * Discord API Wrapper - A custom wrapper for the Discord REST API developed for a proprietary project.
 *
 * This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public
 * License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later
 * version.
 *
 * This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied
 * warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License along with this program.
 * If not, see <http://www.gnu.org/licenses/>.
 */

package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/veteran-software/discord-api-wrapper/logging"
	"github.com/veteran-software/discord-api-wrapper/utilities"

	"github.com/gojek/heimdall/v7"
	"github.com/gojek/heimdall/v7/httpclient"
)

var (
	initialTimeout        = 500 * time.Millisecond
	maxTimeout            = 25 * time.Second
	exponentFactor        = 2.0
	maximumJitterInterval = 2 * time.Millisecond
	retryCount            = 2

	backoff = heimdall.NewExponentialBackoff(initialTimeout, maxTimeout, exponentFactor, maximumJitterInterval)

	// Create a new retry mechanism with the backoff
	retrier = heimdall.NewRetrier(backoff)

	httpClient = httpclient.NewClient(
		httpclient.WithRetrier(retrier),
		httpclient.WithRetryCount(retryCount),
	)
)

var (
	Rest *RateLimiter
)

func init() {
	Rest = NewRatelimiter()
}

func (r *RateLimiter) Request(method, route string, data interface{}, reason *string) (*http.Response, error) {
	return r.requestWithBucketID(method, route, data, strings.SplitN(route, "?", 2)[0], reason)
}

func (r *RateLimiter) requestWithBucketID(method, route string, data interface{}, bucketID string, reason *string) (*http.Response, error) {
	return r.request(method, route, "application/json", data, bucketID, 0, reason)
}

func (r *RateLimiter) request(method, route, contentType string, b interface{}, bucketID string, sequence int, reason *string) (*http.Response, error) {
	if bucketID == "" {
		bucketID = strings.SplitN(route, "?", 2)[0]
	}

	return r.requestWithLockedBucket(method, route, contentType, b, r.lockBucket(bucketID), sequence, reason)
}

func (r *RateLimiter) requestWithLockedBucket(method, route, contentType string, b interface{}, bucket *bucket, sequence int, reason *string) (*http.Response, error) {
	var buffer bytes.Buffer
	if b != nil {
		encoder := json.NewEncoder(&buffer)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(&b)
		if err != nil {
			_ = bucket.release(nil)
			return nil, err
		}
	}

	req, err := http.NewRequest(method, route, bytes.NewReader(buffer.Bytes()))
	if err != nil {
		_ = bucket.release(nil)
		return nil, err
	}

	req.Header.Set(http.CanonicalHeaderKey("Authorization"), fmt.Sprintf("Bot %s", utilities.Token))

	if b != nil {
		req.Header.Set(http.CanonicalHeaderKey("Content-Type"), contentType)
	}

	req.Header.Set(http.CanonicalHeaderKey("User-Agent"), UserAgent)

	ctx, cancel := context.WithDeadline(req.Context(), time.Now().Add(30*time.Second))
	go func(ctx context.Context) {
		defer cancel()

		select {
		case <-ctx.Done():
			switch ctx.Err() {
			case context.DeadlineExceeded:
				logging.Traceln(logging.LogPrefixDiscord, "context timeout exceeded")
			case context.Canceled:
				logging.Traceln(logging.LogPrefixDiscord, "context cancelled; process complete")
			}
		}
	}(ctx)

	resp, err := httpClient.Do(req.WithContext(ctx))
	logging.Infoln(resp.Status)

	if err != nil {
		_ = bucket.release(nil)
		if errors.Is(err, context.Canceled) {
			logging.Warnln("Context cancelled. Deadline was 12 seconds.")
			logging.Warnln("\tRequest was : ", method, " : ", route)
		} else if errors.Is(err, context.DeadlineExceeded) {
			logging.Warnln("Context deadline exceeded.")
			logging.Warnln("\tRequest was : ", method, " : ", route)
		}
		return nil, err
	}

	err = bucket.release(resp.Header)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusCreated:
	case http.StatusNoContent:
	case http.StatusBadGateway:
	case http.StatusTooManyRequests:
		logging.Warnln("Rate Limited!")
		logging.Infoln(route)
		logging.Infoln(resp.Status)

		var rlr rateLimitResponse
		err = json.NewDecoder(resp.Body).Decode(&rlr)
		if err != nil {
			return nil, err
		}

		time.Sleep(rlr.RetryAfter)

		resp, err = r.requestWithLockedBucket(method, route, contentType, b, r.lockBucketObject(bucket), sequence, reason)
	}

	return resp, nil
}
