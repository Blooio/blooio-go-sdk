// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blooio

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/Blooio/blooio-go-sdk/internal/apijson"
	"github.com/Blooio/blooio-go-sdk/internal/apiquery"
	"github.com/Blooio/blooio-go-sdk/internal/requestconfig"
	"github.com/Blooio/blooio-go-sdk/option"
	"github.com/Blooio/blooio-go-sdk/packages/param"
	"github.com/Blooio/blooio-go-sdk/packages/respjson"
)

// ChatService contains methods and other services that help with interacting with
// the blooio API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChatService] method instead.
type ChatService struct {
	options  []option.RequestOption
	Messages ChatMessageService
	// Send native iMessage polls and retrieve poll results with vote counts. Poll
	// events are delivered via separate webhook event types (poll.received,
	// poll.created, poll.voted) and require webhook_type 'poll' or 'all'.
	Polls ChatPollService
	// Control typing indicators for conversations
	Typing ChatTypingService
	// Set, get, and remove conversation backgrounds
	Background ChatBackgroundService
}

// NewChatService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewChatService(opts ...option.RequestOption) (r ChatService) {
	r = ChatService{}
	r.options = opts
	r.Messages = NewChatMessageService(opts...)
	r.Polls = NewChatPollService(opts...)
	r.Typing = NewChatTypingService(opts...)
	r.Background = NewChatBackgroundService(opts...)
	return
}

// Get details for a specific conversation.
func (r *ChatService) Get(ctx context.Context, chatID string, opts ...option.RequestOption) (res *ChatGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if chatID == "" {
		err = errors.New("missing required chatId parameter")
		return nil, err
	}
	path := fmt.Sprintf("chats/%s", url.PathEscape(chatID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List all unique conversations for the organization, sorted by most recent
// message.
func (r *ChatService) List(ctx context.Context, query ChatListParams, opts ...option.RequestOption) (res *ChatListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "chats"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Mark all messages in a chat as read. This sends a read receipt to the sender.
func (r *ChatService) MarkAsRead(ctx context.Context, chatID string, opts ...option.RequestOption) (res *ChatMarkAsReadResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if chatID == "" {
		err = errors.New("missing required chatId parameter")
		return nil, err
	}
	path := fmt.Sprintf("chats/%s/read", url.PathEscape(chatID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Stage the contact card (Name & Photo) for sharing in a chat. The contact card
// will be piggybacked onto the next outgoing message (text or attachment) sent to
// this chat. This is idempotent — calling it multiple times is harmless.
func (r *ChatService) ShareContactCard(ctx context.Context, chatID string, opts ...option.RequestOption) (res *ChatShareContactCardResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if chatID == "" {
		err = errors.New("missing required chatId parameter")
		return nil, err
	}
	path := fmt.Sprintf("chats/%s/contact-card", url.PathEscape(chatID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type LastMessage struct {
	// Any of "inbound", "outbound".
	Direction LastMessageDirection `json:"direction"`
	MessageID string               `json:"message_id"`
	Text      string               `json:"text" api:"nullable"`
	TimeSent  int64                `json:"time_sent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Direction   respjson.Field
		MessageID   respjson.Field
		Text        respjson.Field
		TimeSent    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LastMessage) RawJSON() string { return r.JSON.raw }
func (r *LastMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LastMessageDirection string

const (
	LastMessageDirectionInbound  LastMessageDirection = "inbound"
	LastMessageDirectionOutbound LastMessageDirection = "outbound"
)

type ChatGetResponse struct {
	// Chat identifier (phone number, email, or group ID)
	ID string `json:"id"`
	// Identifier for the active chat background
	BackgroundID string `json:"background_id" api:"nullable"`
	// Public URL of the chat background image (if one has been set via the API)
	BackgroundURL string `json:"background_url" api:"nullable" format:"uri"`
	// Contact info (only for non-group chats)
	Contact          ChatGetResponseContact `json:"contact" api:"nullable"`
	FirstMessageTime int64                  `json:"first_message_time"`
	// Group ID (only for group chats)
	GroupID string `json:"group_id" api:"nullable"`
	// Group name (only for group chats)
	GroupName    string `json:"group_name" api:"nullable"`
	InboundCount int64  `json:"inbound_count"`
	// Whether this is a group chat
	IsGroup          bool        `json:"is_group"`
	LastInboundTime  int64       `json:"last_inbound_time" api:"nullable"`
	LastMessage      LastMessage `json:"last_message"`
	LastMessageTime  int64       `json:"last_message_time"`
	LastOutboundTime int64       `json:"last_outbound_time" api:"nullable"`
	// Number of members (only for group chats)
	MemberCount   int64 `json:"member_count"`
	MessageCount  int64 `json:"message_count"`
	OutboundCount int64 `json:"outbound_count"`
	// Any of "phone", "email", "group".
	Type ChatGetResponseType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		BackgroundID     respjson.Field
		BackgroundURL    respjson.Field
		Contact          respjson.Field
		FirstMessageTime respjson.Field
		GroupID          respjson.Field
		GroupName        respjson.Field
		InboundCount     respjson.Field
		IsGroup          respjson.Field
		LastInboundTime  respjson.Field
		LastMessage      respjson.Field
		LastMessageTime  respjson.Field
		LastOutboundTime respjson.Field
		MemberCount      respjson.Field
		MessageCount     respjson.Field
		OutboundCount    respjson.Field
		Type             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact info (only for non-group chats)
type ChatGetResponseContact struct {
	ContactID  string `json:"contact_id"`
	Identifier string `json:"identifier"`
	Name       string `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContactID   respjson.Field
		Identifier  respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatGetResponseContact) RawJSON() string { return r.JSON.raw }
func (r *ChatGetResponseContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatGetResponseType string

const (
	ChatGetResponseTypePhone ChatGetResponseType = "phone"
	ChatGetResponseTypeEmail ChatGetResponseType = "email"
	ChatGetResponseTypeGroup ChatGetResponseType = "group"
)

type ChatListResponse struct {
	Chats      []ChatListResponseChat `json:"chats"`
	Pagination Pagination             `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Chats       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatListResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatListResponseChat struct {
	// Chat identifier (phone number, email, or group ID)
	ID string `json:"id"`
	// Identifier for the active chat background
	BackgroundID string `json:"background_id" api:"nullable"`
	// Public URL of the chat background image (if one has been set via the API)
	BackgroundURL string `json:"background_url" api:"nullable" format:"uri"`
	// Contact info (only for non-group chats)
	Contact ChatListResponseChatContact `json:"contact" api:"nullable"`
	// Group ID (only for group chats)
	GroupID string `json:"group_id" api:"nullable"`
	// Group name (only for group chats)
	GroupName    string `json:"group_name" api:"nullable"`
	InboundCount int64  `json:"inbound_count"`
	// Whether this is a group chat
	IsGroup          bool        `json:"is_group"`
	LastInboundTime  int64       `json:"last_inbound_time" api:"nullable"`
	LastMessage      LastMessage `json:"last_message"`
	LastMessageTime  int64       `json:"last_message_time"`
	LastOutboundTime int64       `json:"last_outbound_time" api:"nullable"`
	// Number of members (only for group chats)
	MemberCount   int64 `json:"member_count"`
	MessageCount  int64 `json:"message_count"`
	OutboundCount int64 `json:"outbound_count"`
	// Any of "phone", "email", "group".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		BackgroundID     respjson.Field
		BackgroundURL    respjson.Field
		Contact          respjson.Field
		GroupID          respjson.Field
		GroupName        respjson.Field
		InboundCount     respjson.Field
		IsGroup          respjson.Field
		LastInboundTime  respjson.Field
		LastMessage      respjson.Field
		LastMessageTime  respjson.Field
		LastOutboundTime respjson.Field
		MemberCount      respjson.Field
		MessageCount     respjson.Field
		OutboundCount    respjson.Field
		Type             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatListResponseChat) RawJSON() string { return r.JSON.raw }
func (r *ChatListResponseChat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact info (only for non-group chats)
type ChatListResponseChatContact struct {
	ContactID  string `json:"contact_id"`
	Identifier string `json:"identifier"`
	Name       string `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContactID   respjson.Field
		Identifier  respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatListResponseChatContact) RawJSON() string { return r.JSON.raw }
func (r *ChatListResponseChatContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatMarkAsReadResponse struct {
	// Chat identifier
	ChatID string `json:"chat_id"`
	// Timestamp when marked as read
	MarkedAt int64 `json:"marked_at"`
	// Read status
	//
	// Any of "read".
	Status ChatMarkAsReadResponseStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChatID      respjson.Field
		MarkedAt    respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatMarkAsReadResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatMarkAsReadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Read status
type ChatMarkAsReadResponseStatus string

const (
	ChatMarkAsReadResponseStatusRead ChatMarkAsReadResponseStatus = "read"
)

type ChatShareContactCardResponse struct {
	// Normalized chat identifier
	ChatID  string `json:"chat_id"`
	Message string `json:"message"`
	Success bool   `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChatID      respjson.Field
		Message     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatShareContactCardResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatShareContactCardResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatListParams struct {
	// Maximum number of items to return (1-200)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Search query (matches phone/email or contact name)
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Sort order
	//
	// Any of "recent", "oldest".
	Sort ChatListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ChatListParams]'s query parameters as `url.Values`.
func (r ChatListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order
type ChatListParamsSort string

const (
	ChatListParamsSortRecent ChatListParamsSort = "recent"
	ChatListParamsSortOldest ChatListParamsSort = "oldest"
)
