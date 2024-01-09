/*
 * Copyright (c) 2022-2024. Veteran Software
 *
 *  Discord API Wrapper - A custom wrapper for the Discord REST API developed for a proprietary project.
 *
 *  This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public
 *  License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License along with this program.
 *  If not, see <http://www.gnu.org/licenses/>.
 */

package api

import (
	"fmt"
	"net/http"

	log "github.com/veteran-software/nowlive-logging"
)

// An Interaction is the message that your application receives when a user uses an ApplicationCommand or a Message Component.
//
// For Slash Commands, it includes the values that the user submitted.
//
// For User Commands and Message Commands, it includes the resolved User or Message on which the action was taken.
//
// For Message Components it includes identifying information about the component that was used.
//
// It will also include some metadata about how the interaction was triggered: the `guild_id`, `channel_id`, `member` and other fields.
type Interaction struct {
	ID             Snowflake              `json:"id"`                        // ID of the interaction
	ApplicationID  Snowflake              `json:"application_id"`            // ID of the application this interaction is for
	Type           InteractionType        `json:"type"`                      // Type of interaction
	Data           ApplicationCommandData `json:"data,omitempty"`            // Interaction data payload
	GuildID        Snowflake              `json:"guild_id,omitempty"`        // Guild that the interaction was sent from
	Channel        Channel                `json:"channel,omitempty"`         // Channel that the interaction was sent from
	ChannelID      Snowflake              `json:"channel_id,omitempty"`      // Channel that the interaction was sent from
	Member         GuildMember            `json:"member,omitempty"`          // GuildMember data for the invoking user, including permissions
	User           *User                  `json:"user,omitempty"`            // User object for the invoking user, if invoked in a DM
	Token          string                 `json:"token"`                     // Continuation token for responding to the interaction
	Version        int                    `json:"version"`                   // Read-only property, always `1`
	Message        *Message               `json:"message,omitempty"`         // For components, the message they were attached to
	AppPermissions string                 `json:"app_permissions,omitempty"` // Bitwise set of permissions the app or bot has within the channel the interaction was sent from
	Locale         string                 `json:"locale,omitempty"`          // Selected language of the invoking user
	GuildLocale    string                 `json:"guild_locale,omitempty"`    // Guild's preferred locale, if invoked in a Guild
}

// InteractionType - The type of Interaction
type InteractionType int

//goland:noinspection GoUnusedConst
const (
	InteractionTypePing                           InteractionType = iota + 1 // PING
	InteractionTypeApplicationCommand                                        // APPLICATION_COMMAND
	InteractionTypeMessageComponent                                          // MESSAGE_COMPONENT
	InteractionTypeApplicationCommandAutocomplete                            // APPLICATION_COMMAND_AUTOCOMPLETE
	InteractionTypeModalSubmit                                               // MODAL_SUBMIT
)

// ApplicationCommandData
//
// While the data field is guaranteed to be present for all interaction types besides InteractionTypePing, its structure will vary.
type ApplicationCommandData struct {
	ID       Snowflake                                  `json:"id,omitempty"`        // the ID of the invoked command
	Name     string                                     `json:"name,omitempty"`      // the name of the invoked command
	Type     ApplicationCommandType                     `json:"type,omitempty"`      // the type of the invoked command
	Resolved ResolvedData                               `json:"resolved,omitempty"`  // converted users + roles + channels
	Options  []*ApplicationCommandInteractionDataOption `json:"options,omitempty"`   // the params + values from the user
	GuildID  Snowflake                                  `json:"guild_id,omitempty"`  // the id of the guild the command is registered to
	TargetID Snowflake                                  `json:"target_id,omitempty"` // id the of user or message targeted by a user or message command
}

type MessageComponentData struct {
	CustomID      string        `json:"custom_id"`      // the custom_id of the component
	ComponentType ComponentType `json:"component_type"` // the type of the component
	Values        []string      `json:"values"`         // values the user selected in a select menu component
}

type ModalSubmitData struct {
	CustomID   string       `json:"custom_id"`  // the custom_id of the modal
	Components []*Component `json:"components"` // the values submitted by the user
}

// ResolvedData - Descriptive data about the Interaction
//
// If data for a GuildMember is included, data for its corresponding User will also be included.
type ResolvedData struct {
	Users       map[Snowflake]User        `json:"users,omitempty"`       // the IDs and DiscordUser objects
	Members     map[Snowflake]GuildMember `json:"members,omitempty"`     // the IDs and partial GuildMember objects
	Roles       map[Snowflake]Role        `json:"roles,omitempty"`       // the IDs and GuildRole objects
	Channels    map[Snowflake]Channel     `json:"channels,omitempty"`    // the IDs and partial GuildChannel objects
	Messages    map[Snowflake]Message     `json:"messages,omitempty"`    // the ids and partial Message objects
	Attachments map[Snowflake]Attachment  `json:"attachments,omitempty"` // the ids and attachment objects
}

// ApplicationCommandInteractionDataOption - All options have names, and an option can either be a parameter and input value--in which case value will be set--or it can denote a subcommand or group--in which case it will contain a top-level key and another array of options.
//
//	value and options are mutually exclusive.
type ApplicationCommandInteractionDataOption struct {
	Name    string                                     `json:"name"`              // the name of the parameter
	Type    ApplicationCommandOptionType               `json:"type"`              // value of application command option type
	Value   any                                        `json:"value,omitempty"`   // the value of the pair
	Options []*ApplicationCommandInteractionDataOption `json:"options,omitempty"` // present if this option is a group or subcommand
	Focused bool                                       `json:"focused,omitempty"` // true if this option is the currently focused option for autocomplete
}

// MessageInteraction - This is sent on the message object when the message is a response to an Interaction.
//
// This means responses to Message Components do not include this property, instead including a MessageReference object as components always exist on preexisting messages.
type MessageInteraction struct {
	ID     Snowflake       `json:"id"`               // id of the Interaction
	Type   InteractionType `json:"type"`             // the type of Interaction
	Name   string          `json:"name"`             // the name of the ApplicationCommand
	User   User            `json:"user"`             // the user who invoked the interaction
	Member GuildMember     `json:"member,omitempty"` // the Member who invoked the interaction in the Guild
}

// InteractionResponseMessages - After receiving an interaction, you must respond to acknowledge it.
//
// You can choose to respond with a message immediately using type 4, or you can choose to send a deferred response with type 5.
//
// If choosing a deferred response, the user will see a loading state for the interaction, and you'll have up to 15 minutes to edit the original deferred response using Edit Original Interaction Response.
//
// Interaction responses can also be public—everyone can see it—or "ephemeral"—only the invoking user can see it.
//
// That is determined by setting flags to 64 on the InteractionCallbackDataMessages.
type InteractionResponseMessages struct {
	Type InteractionCallbackType          `json:"type"`           // the type of response
	Data *InteractionCallbackDataMessages `json:"data,omitempty"` // an optional response message
}

// InteractionResponseAutocomplete - After receiving an interaction, you must respond to acknowledge it.
//
// You can choose to respond with a message immediately using type 4, or you can choose to send a deferred response with type 5.
//
// If choosing a deferred response, the user will see a loading state for the interaction, and you'll have up to 15 minutes to edit the original deferred response using Edit Original Interaction Response.
//
// Interaction responses can also be public—everyone can see it—or "ephemeral"—only the invoking user can see it.
//
// That is determined by setting flags to 64 on the InteractionCallbackDataMessages.
type InteractionResponseAutocomplete struct {
	Type InteractionCallbackType              `json:"type"`           // the type of response
	Data *InteractionCallbackDataAutocomplete `json:"data,omitempty"` // options for the autocomplete result
}

// InteractionResponseModal - After receiving an interaction, you must respond to acknowledge it.
//
// You can choose to respond with a message immediately using type 4, or you can choose to send a deferred response with type 5.
//
// If choosing a deferred response, the user will see a loading state for the interaction, and you'll have up to 15 minutes to edit the original deferred response using Edit Original Interaction Response.
//
// Interaction responses can also be public—everyone can see it—or "ephemeral"—only the invoking user can see it.
//
// That is determined by setting flags to 64 on the InteractionCallbackDataMessages.
type InteractionResponseModal struct {
	CallbackType InteractionCallbackType       `json:"type"`           // the type of response
	Data         *InteractionCallbackDataModal `json:"data,omitempty"` // the information submitted through the modal
}

// InteractionCallbackType - The type of callback to an interaction with respond
type InteractionCallbackType int

//goland:noinspection GoUnusedConst
const (
	Pong                             InteractionCallbackType = iota + 1 // ACK a Ping
	ChannelMessageWithSource         InteractionCallbackType = iota + 3 // respond to an interaction with a message
	DeferredChannelMessageWithSource                                    // ACK an interaction and edit a response later, the user sees a loading state
	DeferredUpdateMessage                                               // for components, ACK an interaction and edit the original message later; the user does not see a loading state; edit the message using EditOriginalInteractionResponse
	UpdateMessage                                                       // for components, edit the message the component was attached to
	AutocompleteResult                                                  // respond to an autocomplete interaction with suggested choices
	Modal                                                               // respond to an interaction with a popup modal ** Not available for MODAL_SUBMIT and PING interactions.
)

// InteractionCallbackDataMessages - Not all message fields are currently supported by Discord
//
// Data payload for InteractionResponseMessages
type InteractionCallbackDataMessages struct {
	TTS             bool             `json:"tts,omitempty"`         // is the response TTS
	Content         string           `json:"content,omitempty"`     // message content
	Embeds          []*Embed         `json:"embeds,omitempty"`      // supports up to 10 embeds
	AllowedMentions *AllowedMentions `json:"allowed_mentions"`      // AllowedMentionType object
	Flags           MessageFlags     `json:"flags,omitempty"`       // set to 64 to make your response Ephemeral
	Components      []*Component     `json:"components,omitempty"`  // message components
	Attachments     []*Attachment    `json:"attachments,omitempty"` // attachment objects with filename and description
}

// InteractionCallbackDataAutocomplete - Data payload for InteractionResponseAutocomplete
type InteractionCallbackDataAutocomplete struct {
	Choices []*ApplicationCommandOptionChoice `json:"choices"` // autocomplete choices (max of 25 choices)
}

// InteractionCallbackDataModal - Data payload for InteractionResponseModal
type InteractionCallbackDataModal struct {
	CustomID   string       `json:"custom_id"`  // a developer-defined identifier for the component, max 100 characters
	Title      string       `json:"title"`      // the title of the popup modal, max 45 characters
	Components []*Component `json:"components"` // between 1 and 5 (inclusive) components that make up the modal
}

// BuildResponse
// Deprecated: helper method for building a basic message response
func (i *Interaction) BuildResponse(embeds []*Embed) *InteractionResponseMessages {
	ir := &InteractionResponseMessages{
		Data: &InteractionCallbackDataMessages{},
	}

	for _, embed := range embeds {
		ir.Data.Embeds = append(ir.Data.Embeds, embed)
	}

	if i.Type == InteractionTypeApplicationCommand {
		ir.Type = ChannelMessageWithSource

	} else {
		ir.Type = UpdateMessage
	}

	return ir
}

// TODO: Find the best way to handle these; overloads?

// CreateInteractionResponse
//
// Create a response to an Interaction from the gateway. Body is an InteractionResponse. Returns 204 No Content.
//
// This endpoint also supports file attachments similar to the webhook endpoints.
// Refer to Uploading Files for details on uploading files and `multipart/form-data` requests.
func (i *Interaction) CreateInteractionResponse(payload any) error {
	// verify that we only accept the payload that we want
	// maybe future language version will make this easier/cleaner
	switch payload.(type) {
	case **InteractionResponseMessages:
	case **InteractionResponseAutocomplete:
	case **InteractionResponseModal:
	default:
		return nil
	}

	u := parseRoute(fmt.Sprintf(createInteractionResponse, api, i.ID.String(), i.Token))

	_, err := firePostRequest(u, payload, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return err
	}

	return nil
}

// GetOriginalInteractionResponse Returns the initial Interaction response.
func (i *Interaction) GetOriginalInteractionResponse() (method string, route string) {
	return http.MethodGet, fmt.Sprintf(getOriginalInteractionResponse, api, i.ApplicationID.String(), i.Token)
}

// EditOriginalInteractionResponse Edits the initial Interaction response.
func (i *Interaction) EditOriginalInteractionResponse() (method string, route string) {
	return http.MethodPatch, fmt.Sprintf(editOriginalInteractionResponse, api, i.ApplicationID.String(), i.Token)
}

// DeleteOriginalInteractionResponse Deletes the initial Interaction response. Returns 204 on success.
func (i *Interaction) DeleteOriginalInteractionResponse() (method string, route string) {
	return http.MethodDelete, fmt.Sprintf(deleteOriginalInteractionResponse, api, i.ApplicationID.String(), i.Token)
}

// CreateFollowupMessage - Create a followup message for an Interaction.
//
// Functions the same as Execute Webhook, but wait is always true, and flags can be set to 64 in the body to send an ephemeral message.
//
// The thread_id, avatar_url, and username parameters are not supported when using this endpoint for interaction followups.
func (i *Interaction) CreateFollowupMessage() (method string, route string) {
	return http.MethodPost, fmt.Sprintf(createFollowupMessage, api, i.ApplicationID.String(), i.Token)
}

// GetFollowupMessage - Returns a followup message for an Interaction.
//
// Functions the same as Get Webhook Message.
//
//	Does not support ephemeral followups.
func (i *Interaction) GetFollowupMessage() (method string, route string) {
	return http.MethodGet, fmt.Sprintf(getFollowupMessage, api, i.ApplicationID.String(), i.Token, i.Message.ID)
}

// EditFollowupMessage - Edits a followup message for an Interaction.
//
// Functions the same as Edit Webhook Message.
//
//	Does not support ephemeral followups.
func (i *Interaction) EditFollowupMessage() (method string, route string) {
	return http.MethodPatch, fmt.Sprintf(editFollowupMessage, api, i.ApplicationID, i.Token, i.Message.ID)
}

// DeleteFollowupMessage - Deletes a followup message for an Interaction.
//
// Returns 204 No Content on success.
//
//	Does not support ephemeral followups.
func (i *Interaction) DeleteFollowupMessage() (method string, route string) {
	return http.MethodDelete, fmt.Sprintf(deleteFollowupMessage, api, i.ApplicationID, i.Token, i.Message.ID)
}
