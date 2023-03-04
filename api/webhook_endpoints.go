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
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/vincent-petithory/dataurl"
)

// CreateWebhook - Create a new webhook.
//
// Requires the ManageWebhooks permission.
//
// Returns a Webhook object on success.
//
// Webhook names follow our naming restrictions that can be found in our Usernames and Nicknames documentation, with the following additional stipulations:
//
//   - Webhook names cannot be: 'clyde'
//
// This endpoint supports the "X-Audit-Log-Reason" header.
func (c *Channel) CreateWebhook(name string, avatar *dataurl.DataURL, reason *string) (*Webhook, error) {
	if len(name) < 1 || len(name) > 80 || strings.Contains(strings.ToLower(name), "clyde") {
		return nil, errors.New("webhook length is incorrect or the name contains \"Clyde\"")
	}

	params := struct {
		Name   string `json:"name"`
		Avatar string `json:"avatar"`
	}{
		Name:   name,
		Avatar: avatar.String(),
	}

	// TODO: Check for ManageWebhooks permission

	u := parseRoute(fmt.Sprintf(createWebhook, api, c.ID.String()))

	var webhook *Webhook
	err := json.Unmarshal(firePostRequest(u, params, reason), &webhook)

	return webhook, err
}

// GetChannelWebhooks - Returns a list of channel webhook objects. Requires the ManageWebhooks permission.
func (c *Channel) GetChannelWebhooks() ([]*Webhook, error) {
	// TODO: Check for ManageWebhooks permission

	u := parseRoute(fmt.Sprintf(getChannelWebhooks, api, c.ID.String()))

	var webhooks []*Webhook
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &webhooks)

	return webhooks, err
}

// GetGuildWebhooks - Returns a list of guild webhook objects. Requires the ManageWebhooks permission.
func (g *Guild) GetGuildWebhooks() ([]*Webhook, error) {
	// TODO: Check for ManageWebhooks permission

	u := parseRoute(fmt.Sprintf(getGuildWebhooks, api, g.ID.String()))

	var webhooks []*Webhook
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &webhooks)

	return webhooks, err
}

// GetWebhook - Returns the new webhook object for the given id.
func (w *Webhook) GetWebhook() (*Webhook, error) {
	u := parseRoute(fmt.Sprintf(getWebhook, api, w.ID.String()))

	var webhook *Webhook
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &webhook)

	return webhook, err
}

// GetWebhookWithToken - Same as above, except this call does not require authentication and returns no user in the webhook object.
func (w *Webhook) GetWebhookWithToken() (*Webhook, error) {
	u := parseRoute(fmt.Sprintf(getWebhookWithToken, api, w.ID.String(), w.Token))

	var webhook *Webhook
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &webhook)

	return webhook, err
}

// ModifyWebhook - Modify a webhook. Requires the ManageWebhooks permission. Returns the updated Webhook object on success. Fires a Webhooks Update Gateway event.
//
// # All parameters to this endpoint are optional
//
// This endpoint supports the "X-Audit-Log-Reason" header.
func (w *Webhook) ModifyWebhook(name *string, avatar *dataurl.DataURL, channelId *Snowflake, reason *string) (
	*Webhook,
	error,
) {
	params := struct {
		Name      string    `json:"name,omitempty"`
		Avatar    string    `json:"avatar,omitempty"`
		ChannelId Snowflake `json:"channel_id,omitempty"`
	}{}

	if name != nil {
		params.Name = *name
	}
	if avatar != nil {
		params.Avatar = avatar.String()
	}
	if channelId != nil {
		params.ChannelId = *channelId
	}

	// TODO: Check for ManageWebhooks permission

	u := parseRoute(fmt.Sprintf(modifyWebhook, api, w.ID.String()))

	var webhook *Webhook
	err := json.Unmarshal(firePatchRequest(u, params, reason), &webhook)

	return webhook, err
}

// ModifyWebhookWithToken - Same as above, except this call does not require authentication, does not accept a channel_id parameter in the body, and does not return a user in the webhook object.
func (w *Webhook) ModifyWebhookWithToken(name *string, avatar *dataurl.DataURL, reason *string) (*Webhook, error) {
	params := struct {
		Name   string `json:"name,omitempty"`
		Avatar string `json:"avatar,omitempty"`
	}{}

	if name != nil {
		params.Name = *name
	}
	if avatar != nil {
		params.Avatar = avatar.String()
	}

	u := parseRoute(fmt.Sprintf(modifyWebhookWithToken, api, w.ID.String(), w.Token))

	var webhook *Webhook
	err := json.Unmarshal(firePatchRequest(u, params, reason), &webhook)

	return webhook, err
}

// DeleteWebhook - Delete a webhook permanently. Requires the ManageWebhooks permission. Returns a 204 No Content response on success.
//
// This endpoint supports the "X-Audit-Log-Reason" header.
func (w *Webhook) DeleteWebhook(reason *string) error {
	u := parseRoute(fmt.Sprintf(deleteWebhook, api, w.ID.String()))

	return fireDeleteRequest(u, reason)
}

// DeleteWebhookWithToken - Same as above, except this call does not require authentication.
func (w *Webhook) DeleteWebhookWithToken(reason *string) error {
	u := parseRoute(fmt.Sprintf(deleteWebhookWithToken, api, w.ID.String(), w.Token))

	return fireDeleteRequest(u, reason)
}

// ExecuteWebhook - Refer to Uploading Files for details on attachments and multipart/form-data requests.
//
// Note that when sending a message, you must provide a value for at least one of content, embeds, or file.
//
// wait is required; threadID is optional; pass nil if not needed
func (w *Webhook) ExecuteWebhook(wait bool, threadID *Snowflake, params *ExecuteWebhookJSON) (*Message, error) {
	u := parseRoute(fmt.Sprintf(executeWebhook, api, w.ID, w.Token))

	q := u.Query()
	q.Set("wait", strconv.FormatBool(wait))
	if threadID != nil {
		q.Set("thread_id", threadID.String())
	}
	if len(q) > 0 {
		u.RawQuery = q.Encode()
	}

	var message *Message
	err := json.Unmarshal(firePostRequest(u, params, nil), &message)

	return message, err
}

// ExecuteWebhookJSON - JSON payload structure
type ExecuteWebhookJSON struct {
	Content         string           `json:"content"`                    // the message contents (up to 2000 characters); Required - one of content, file, embeds
	Username        string           `json:"username,omitempty"`         // override the default username of the webhook; Required - false
	AvatarURL       string           `json:"avatar_url,omitempty"`       // override the default avatar of the webhook; Required - false
	Tts             bool             `json:"tts,omitempty"`              // true if this is a TTS message; Required - false
	Embeds          []*Embed         `json:"embeds"`                     // embedded rich content; Required - one of content, file, embeds
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"` // allowed mentions for the message; Required - false
	Components      []*Component     `json:"components,omitempty"`       // the components to include with the message - Required - false
	PayloadJson     string           `json:"payload_json"`               // JSON encoded body of non-file params; Required - "multipart/form-data" only
	Attachments     []*Attachment    `json:"attachments,omitempty"`      // Attachment objects with filename and description; Required - false
	Flags           MessageFlags     `json:"flags,omitempty"`            // MessageFlags combined as a bitfield (only SuppressEmbeds can be set)
	ThreadName      string           `json:"thread_name"`                // name of thread to create (requires the webhook channel to be a forum channel)
}

// GetWebhookMessage - Returns a previously-sent webhook message from the same token. Returns a message object on success.
//
// threadID is optional; pass nil if not needed
func (w *Webhook) GetWebhookMessage(msgID *Snowflake, threadID *Snowflake) (*Message, error) {
	u := parseRoute(fmt.Sprintf(getWebhookMessage, api, w.ID.String(), w.Token, msgID.String()))

	q := u.Query()
	if threadID != nil {
		q.Set("thread_id", threadID.String())
	}
	if len(q) > 0 {
		u.RawQuery = q.Encode()
	}

	var message *Message
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &message)

	return message, err
}

// EditWebhookMessage - Edits a previously-sent webhook message from the same token. Returns a Message object on success.
//
// When the content field is edited, the mentions array in the message object will be reconstructed from scratch based on the new content.
// The AllowedMentions field of the edit request controls how this happens.
// If there is no explicit AllowedMentions in the edit request, the content will be parsed with default allowances, that is, without regard to whether or not an AllowedMentions was present in the request that originally created the Message.
//
// Refer to Uploading Files for details on attachments and `multipart/form-data requests`.
// Any provided files will be appended to the message.
// To remove or replace files you will have to supply the "attachments" field which specifies the files to retain on the message after edit.
//
// Starting with API v10, the attachments array must contain all attachments that should be present after edit, including retained and new attachments provided in the request body.
//
// All JSON parameters to this endpoint are optional and nullable.
//
// threadID is optional; pass nil if not needed
func (w *Webhook) EditWebhookMessage(msgID *Snowflake, threadID *Snowflake, payload *EditWebhookMessageJSON) (
	*Message,
	error,
) {
	u := parseRoute(fmt.Sprintf(editWebhookMessage, api, w.ID.String(), w.Token, msgID.String()))

	q := u.Query()
	if threadID != nil {
		q.Set("thread_id", threadID.String())
	}
	if len(q) > 0 {
		u.RawQuery = q.Encode()
	}

	var message *Message
	err := json.Unmarshal(firePatchRequest(u, payload, nil), &message)

	return message, err
}

// EditWebhookMessageJSON - All parameters to this endpoint are optional and nullable.
type EditWebhookMessageJSON struct {
	Content         *string          `json:"content,omitempty"`          // the message contents (up to 2000 characters)
	Embeds          []*Embed         `json:"embeds,omitempty"`           // embedded rich content
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"` // allowed mentions for the message
	Components      []*Component     `json:"components,omitempty"`       // the components to include with the message
	PayloadJson     *string          `json:"payload_json,omitempty"`     // JSON encoded body of non-file params (multipart/form-data only)
	Attachments     []*Attachment    `json:"attachments,omitempty"`      // attached files to keep and possible descriptions for new files
}

// DeleteWebhookMessage - Deletes a message that was created by the webhook. Returns a 204 No Content response on success.
//
// threadID is optional; pass nil if not needed
func (w *Webhook) DeleteWebhookMessage(msgID *Snowflake, threadID *Snowflake) error {
	u := parseRoute(fmt.Sprintf(deleteWebhookMessage, api, w.ID.String(), w.Token, msgID.String()))

	q := u.Query()
	if threadID != nil {
		q.Set("thread_id", threadID.String())
	}
	if len(q) > 0 {
		u.RawQuery = q.Encode()
	}

	return fireDeleteRequest(u, nil)
}
