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
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/veteran-software/discord-api-wrapper/v10/logging"
)

// Emoji - Routes for controlling emojis do not follow the normal rate limit conventions.
//
// These routes are specifically limited on a per-guild basis to prevent abuse.
//
// This means that the quota returned by our APIs may be inaccurate, and you may encounter 429s.
type Emoji struct {
	ID            *Snowflake `json:"id"`                       // ID - emoji id
	Name          string     `json:"name"`                     // Name - emoji name
	Roles         []Role     `json:"roles,omitempty"`          // Roles - roles allowed to use this emoji
	User          *User      `json:"user,omitempty"`           // User - user that created this emoji
	RequireColons bool       `json:"require_colons,omitempty"` // RequireColons - whether this emoji must be wrapped in colons
	Managed       bool       `json:"managed,omitempty"`        // Managed - whether this emoji is managed
	Animated      bool       `json:"animated,omitempty"`       // Animated - whether this emoji is animated
	Available     bool       `json:"available,omitempty"`      // Available - whether this emoji can be used, may be false due to loss of Server Boosts
}

// ListGuildEmojis - Returns a list of emoji objects for the given guild.
func (g *Guild) ListGuildEmojis() ([]Emoji, error) {
	u, err := url.Parse(fmt.Sprintf(listGuildEmojis, api, g.ID.String()))
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}

	resp, err := Rest.Request(http.MethodGet, u.String(), nil, nil)
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var emojis []Emoji
	err = json.NewDecoder(resp.Body).Decode(&emojis)
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}

	return emojis, nil
}

// GetGuildEmoji - Returns an emoji object for the given guild and emoji IDs.
func (g *Guild) GetGuildEmoji(emoji Emoji) (*Emoji, error) {
	u, err := url.Parse(fmt.Sprintf(getGuildEmoji, api, g.ID.String(), emoji.ID.String()))
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}

	resp, err := Rest.Request(http.MethodGet, u.String(), nil, nil)
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var emojis *Emoji
	err = json.NewDecoder(resp.Body).Decode(&emojis)
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}

	return emojis, nil
}

// CreateGuildEmoji - Create a new emoji for the guild.
//
// Requires the ManageEmojisAndStickers permission.
//
// Returns the new Emoji object on success.
//
// Fires a Guild Emojis Update Gateway event.
//
// Emojis and animated emojis have a maximum file size of 256kb.
//
// Attempting to upload an emoji larger than this limit will fail and return "400 Bad Request" and an error message, but not a JSON status code.
//
// This endpoint supports the "X-Audit-Log-Reason" header.
func (g *Guild) CreateGuildEmoji(payload CreateEmojiJSON, reason *string) (*Emoji, error) {
	u, err := url.Parse(fmt.Sprintf(createGuildEmoji, api, g.ID.String()))
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}

	resp, err := Rest.Request(http.MethodPost, u.String(), payload, reason)
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var emojis *Emoji
	err = json.NewDecoder(resp.Body).Decode(&emojis)
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}

	return emojis, nil
}

// CreateEmojiJSON - Parameters to pass in the JSON payload
//
// TODO: Validate the base64.Encoding
type CreateEmojiJSON struct {
	Name  string          `json:"name"`  // Name - name of the emoji
	Image base64.Encoding `json:"image"` // Image - the 128x128 emoji image
	Roles []Snowflake     `json:"roles"` // Roles - roles allowed to use this emoji
}

// ModifyGuildEmoji - Modify the given emoji.
//
// Requires the ManageEmojisAndStickers permission.
//
// Returns the updated emoji object on success.
//
// Fires a Guild Emojis Update Gateway event.
//
// All JSON parameters to this endpoint are optional.
//
// This endpoint supports the X-Audit-Log-Reason header.
func (g *Guild) ModifyGuildEmoji(emoji Emoji, payload ModifyGuildEmojiJSON, reason *string) (*Emoji, error) {
	u, err := url.Parse(fmt.Sprintf(modifyGuildEmoji, api, g.ID.String(), emoji.ID.String()))
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}

	resp, err := Rest.Request(http.MethodPatch, u.String(), payload, reason)
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var e *Emoji
	err = json.NewDecoder(resp.Body).Decode(&e)
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}

	return e, nil
}

// ModifyGuildEmojiJSON - Parameters to pass in the JSON payload
type ModifyGuildEmojiJSON struct {
	Name  string       `json:"name,omitempty"`  // Name - name of the emoji
	Roles []*Snowflake `json:"roles,omitempty"` // Roles - roles allowed to use this emoji
}

// DeleteGuildEmoji - Delete the given emoji.
//
// Requires the ManageEmojisAndStickers permission.
//
// Returns "204 No Content" on success.
//
// Fires a GuildEmojisUpdate Gateway event.
//
// This endpoint supports the "X-Audit-Log-Reason" header.
func (g *Guild) DeleteGuildEmoji(emoji Emoji, reason *string) error {
	u, err := url.Parse(fmt.Sprintf(deleteGuildEmoji, api, g.ID.String(), emoji.ID.String()))
	if err != nil {
		logging.Errorln(err)
		return err
	}

	resp, err := Rest.Request(http.MethodDelete, u.String(), nil, reason)
	if err != nil {
		logging.Errorln(err)
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	return nil
}
