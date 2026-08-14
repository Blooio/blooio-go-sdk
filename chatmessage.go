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

// ChatMessageService contains methods and other services that help with
// interacting with the blooio API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChatMessageService] method instead.
type ChatMessageService struct {
	options []option.RequestOption
}

// NewChatMessageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewChatMessageService(opts ...option.RequestOption) (r ChatMessageService) {
	r = ChatMessageService{}
	r.options = opts
	return
}

// Get details for a specific message.
func (r *ChatMessageService) Get(ctx context.Context, messageID string, query ChatMessageGetParams, opts ...option.RequestOption) (res *ChatMessageGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if query.ChatID == "" {
		err = errors.New("missing required chatId parameter")
		return nil, err
	}
	if messageID == "" {
		err = errors.New("missing required messageId parameter")
		return nil, err
	}
	path := fmt.Sprintf("chats/%s/messages/%s", url.PathEscape(query.ChatID), url.PathEscape(messageID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List all messages in a conversation with optional filtering.
func (r *ChatMessageService) List(ctx context.Context, chatID string, query ChatMessageListParams, opts ...option.RequestOption) (res *ChatMessageListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if chatID == "" {
		err = errors.New("missing required chatId parameter")
		return nil, err
	}
	path := fmt.Sprintf("chats/%s/messages", url.PathEscape(chatID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get delivery status for a specific message.
func (r *ChatMessageService) GetStatus(ctx context.Context, messageID string, query ChatMessageGetStatusParams, opts ...option.RequestOption) (res *ChatMessageGetStatusResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if query.ChatID == "" {
		err = errors.New("missing required chatId parameter")
		return nil, err
	}
	if messageID == "" {
		err = errors.New("missing required messageId parameter")
		return nil, err
	}
	path := fmt.Sprintf("chats/%s/messages/%s/status", url.PathEscape(query.ChatID), url.PathEscape(messageID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Add or remove a reaction to a message. Supports classic iMessage tapbacks (love,
// like, dislike, laugh, emphasize, question) and emoji reactions (e.g. +😂, -😂).
//
// The messageId can be an explicit message ID (e.g., msg_xxx) or a relative index
// (-1 for last message, -2 for second-to-last, etc.). When using relative indices,
// you can optionally filter by message direction (inbound/outbound only).
func (r *ChatMessageService) React(ctx context.Context, messageID string, params ChatMessageReactParams, opts ...option.RequestOption) (res *ChatMessageReactResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if params.ChatID == "" {
		err = errors.New("missing required chatId parameter")
		return nil, err
	}
	if messageID == "" {
		err = errors.New("missing required messageId parameter")
		return nil, err
	}
	path := fmt.Sprintf("chats/%s/messages/%s/reactions", url.PathEscape(params.ChatID), url.PathEscape(messageID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Send a message to a chat. The chatId can be: (1) E.164 phone number, (2) email
// address, (3) group ID (grp_xxxx), or (4) comma-separated list of phone/email for
// multi-recipient chats. For multi-recipient, an unnamed group is automatically
// created or reused if the exact participant combination already exists. For
// explicit groups, the group must be linked to an existing iMessage chat.
//
// **iMessage send-with-effect:** set the optional `effect` field to attach an
// Apple expressive send (slam, loud, gentle, invisible-ink) or screen effect
// (echo, spotlight, balloons, confetti, love, lasers, fireworks, celebration).
// Effects are an iMessage-only feature — when the recipient is on SMS/RCS the
// message is delivered without the animation. Effects are not supported in
// multipart (`parts`) mode.
//
// **Threaded replies (iMessage inline reply):** set the optional `reply_to` field
// to send the outgoing message as a reply to a specific earlier message. Two
// shapes are accepted: `{ "message_id": "msg_…" }` references a Blooio-minted
// message in the same chat (most common — the message*id returned by an earlier
// send or surfaced on a `message.received` webhook), or
// `{ "guid": "…", "part_index": 0 }` references the raw iMessage GUID for the rare
// case where the parent wasn't recorded by Blooio. The reply must target the same
// chat and the same from-number as the new send, and the parent must be no older
// than 30 days (the iMessage on-device retention horizon). Reply support is
// iMessage-only and is rejected on Twilio, dashboard-Twilio, and hybrid send
// paths; it's also rejected on multi-message fan-outs (`text` array or per-part
// URL-balloon batch). See the `400` responses for the full set of
// `reply_target*\*` error codes.
func (r *ChatMessageService) Send(ctx context.Context, chatID string, params ChatMessageSendParams, opts ...option.RequestOption) (res *ChatMessageSendResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if chatID == "" {
		err = errors.New("missing required chatId parameter")
		return nil, err
	}
	path := fmt.Sprintf("chats/%s/messages", url.PathEscape(chatID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Rich-link-preview overrides for URL messages (iMessage URL balloon). All fields
// are optional. Only applies when the message text (or the concatenated part text)
// is exactly a single http(s) URL. If omitted but the text is a URL, Blooio
// auto-fetches the page's Open Graph metadata to generate a preview. If the image
// download fails, the send still succeeds — Blooio silently falls back to the
// auto-generated preview.
type LinkPreviewParam struct {
	// HTTPS URL to an image (png, jpg, webp, gif). Blooio downloads the image
	// server-side and attaches it as the rich-link hero. Max 16 MB. If the download
	// fails or returns a non-image MIME, the send falls back to auto-fetched OG
	// metadata.
	ImageURL param.Opt[string] `json:"image_url,omitzero" format:"uri"`
	// Bold title line rendered in the iMessage bubble. Overrides the page's
	// `<meta property="og:title">`.
	Title param.Opt[string] `json:"title,omitzero"`
	paramObj
}

func (r LinkPreviewParam) MarshalJSON() (data []byte, err error) {
	type shadow LinkPreviewParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkPreviewParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Reaction struct {
	// Whether the reaction is currently active (true) or was removed (false)
	IsAdded bool `json:"is_added"`
	// The reaction value. Classic tapbacks: love, like, dislike, laugh, emphasize,
	// question. Emoji reactions: the emoji character (e.g. 😂, 👍).
	Reaction string `json:"reaction"`
	// Phone number or email of who sent the reaction. Null when the reaction was sent
	// by you (outbound).
	Sender string `json:"sender" api:"nullable"`
	// Timestamp when the reaction was sent (ms)
	TimeSent int64 `json:"time_sent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsAdded     respjson.Field
		Reaction    respjson.Field
		Sender      respjson.Field
		TimeSent    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Reaction) RawJSON() string { return r.JSON.raw }
func (r *Reaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatMessageGetResponse struct {
	Attachments []any                         `json:"attachments"`
	ChatID      string                        `json:"chat_id"`
	Contact     ChatMessageGetResponseContact `json:"contact" api:"nullable"`
	// Any of "inbound", "outbound".
	Direction ChatMessageGetResponseDirection `json:"direction"`
	Error     string                          `json:"error" api:"nullable"`
	// Organization phone number (from-number) used for this message
	InternalID string `json:"internal_id" api:"nullable"`
	MessageID  string `json:"message_id"`
	// Transport used to carry the message; never null. `pending` = accepted and
	// dispatched, wire service not resolved yet (settles within seconds of send);
	// `imessage` = delivered over iMessage (blue bubble); `rcs` = delivered over RCS;
	// `sms` = fell back to SMS/MMS (green bubble); `unknown` = accepted by the carrier
	// but the wire service could not be resolved before the tracking window closed
	// (see `error`).
	//
	// Any of "pending", "unknown", "imessage", "sms", "rcs".
	Protocol ChatMessageGetResponseProtocol `json:"protocol"`
	// Reactions on this message (tapbacks and emoji reactions)
	Reactions []Reaction `json:"reactions"`
	// Inline-reply parent reference. Identical shape on `message.received` webhooks
	// and on every GET endpoint that returns a single message or a list of messages.
	ReplyTo ChatMessageGetResponseReplyTo `json:"reply_to" api:"nullable"`
	// Sender's phone number or email for inbound group messages. Null for outbound
	// messages and 1-1 chats.
	Sender string `json:"sender" api:"nullable"`
	// Delivery lifecycle state. `pending` = persisted and being prepared for dispatch;
	// `queued` = accepted and waiting to be handed to Apple/the carrier; `sent` =
	// handed off to Apple/the carrier (protocol resolution happens around here);
	// `delivered` = a delivery receipt was received; `failed` = could not be delivered
	// (see `error`); `cancellation_requested` = a cancel was requested for a
	// still-queued message (best-effort); `cancelled` = cancelled before dispatch.
	// Inbound messages are surfaced via webhooks with `received`; read receipts arrive
	// as a `read` event.
	//
	// Any of "pending", "queued", "sent", "delivered", "failed",
	// "cancellation_requested", "cancelled".
	Status        ChatMessageGetResponseStatus `json:"status" api:"nullable"`
	Text          string                       `json:"text" api:"nullable"`
	TimeDelivered int64                        `json:"time_delivered" api:"nullable"`
	TimeSent      int64                        `json:"time_sent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attachments   respjson.Field
		ChatID        respjson.Field
		Contact       respjson.Field
		Direction     respjson.Field
		Error         respjson.Field
		InternalID    respjson.Field
		MessageID     respjson.Field
		Protocol      respjson.Field
		Reactions     respjson.Field
		ReplyTo       respjson.Field
		Sender        respjson.Field
		Status        respjson.Field
		Text          respjson.Field
		TimeDelivered respjson.Field
		TimeSent      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatMessageGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatMessageGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatMessageGetResponseContact struct {
	ContactID string `json:"contact_id"`
	// The contact's phone number or email
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
func (r ChatMessageGetResponseContact) RawJSON() string { return r.JSON.raw }
func (r *ChatMessageGetResponseContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatMessageGetResponseDirection string

const (
	ChatMessageGetResponseDirectionInbound  ChatMessageGetResponseDirection = "inbound"
	ChatMessageGetResponseDirectionOutbound ChatMessageGetResponseDirection = "outbound"
)

// Transport used to carry the message; never null. `pending` = accepted and
// dispatched, wire service not resolved yet (settles within seconds of send);
// `imessage` = delivered over iMessage (blue bubble); `rcs` = delivered over RCS;
// `sms` = fell back to SMS/MMS (green bubble); `unknown` = accepted by the carrier
// but the wire service could not be resolved before the tracking window closed
// (see `error`).
type ChatMessageGetResponseProtocol string

const (
	ChatMessageGetResponseProtocolPending  ChatMessageGetResponseProtocol = "pending"
	ChatMessageGetResponseProtocolUnknown  ChatMessageGetResponseProtocol = "unknown"
	ChatMessageGetResponseProtocolImessage ChatMessageGetResponseProtocol = "imessage"
	ChatMessageGetResponseProtocolSMS      ChatMessageGetResponseProtocol = "sms"
	ChatMessageGetResponseProtocolRcs      ChatMessageGetResponseProtocol = "rcs"
)

// Inline-reply parent reference. Identical shape on `message.received` webhooks
// and on every GET endpoint that returns a single message or a list of messages.
type ChatMessageGetResponseReplyTo struct {
	// The raw iMessage GUID of the parent. Always populated on real inline replies;
	// the on-device record-of-truth identifier that survives even when `message_id`
	// cannot be resolved.
	Guid string `json:"guid" api:"required"`
	// The Blooio `message_id` of the parent message. NULL when the parent isn't in our
	// `messages` table (e.g., the original was sent from outside Blooio's pipeline).
	MessageID string `json:"message_id" api:"required"`
	// Which part of the parent was replied to. 0 for the common single-part case.
	PartIndex int64 `json:"part_index" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Guid        respjson.Field
		MessageID   respjson.Field
		PartIndex   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatMessageGetResponseReplyTo) RawJSON() string { return r.JSON.raw }
func (r *ChatMessageGetResponseReplyTo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Delivery lifecycle state. `pending` = persisted and being prepared for dispatch;
// `queued` = accepted and waiting to be handed to Apple/the carrier; `sent` =
// handed off to Apple/the carrier (protocol resolution happens around here);
// `delivered` = a delivery receipt was received; `failed` = could not be delivered
// (see `error`); `cancellation_requested` = a cancel was requested for a
// still-queued message (best-effort); `cancelled` = cancelled before dispatch.
// Inbound messages are surfaced via webhooks with `received`; read receipts arrive
// as a `read` event.
type ChatMessageGetResponseStatus string

const (
	ChatMessageGetResponseStatusPending               ChatMessageGetResponseStatus = "pending"
	ChatMessageGetResponseStatusQueued                ChatMessageGetResponseStatus = "queued"
	ChatMessageGetResponseStatusSent                  ChatMessageGetResponseStatus = "sent"
	ChatMessageGetResponseStatusDelivered             ChatMessageGetResponseStatus = "delivered"
	ChatMessageGetResponseStatusFailed                ChatMessageGetResponseStatus = "failed"
	ChatMessageGetResponseStatusCancellationRequested ChatMessageGetResponseStatus = "cancellation_requested"
	ChatMessageGetResponseStatusCancelled             ChatMessageGetResponseStatus = "cancelled"
)

type ChatMessageListResponse struct {
	ChatID     string                           `json:"chat_id"`
	Messages   []ChatMessageListResponseMessage `json:"messages"`
	Pagination Pagination                       `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChatID      respjson.Field
		Messages    respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatMessageListResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatMessageListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatMessageListResponseMessage struct {
	Attachments []any `json:"attachments"`
	// Any of "inbound", "outbound".
	Direction string `json:"direction"`
	Error     string `json:"error" api:"nullable"`
	// Phone number or email of the contact, or group ID for group messages
	ExternalID string `json:"external_id"`
	// Organization phone number (from-number) used for this message
	InternalID string `json:"internal_id" api:"nullable"`
	MessageID  string `json:"message_id"`
	// Transport used to carry the message; never null. `pending` = accepted and
	// dispatched, wire service not resolved yet (settles within seconds of send);
	// `imessage` = delivered over iMessage (blue bubble); `rcs` = delivered over RCS;
	// `sms` = fell back to SMS/MMS (green bubble); `unknown` = accepted by the carrier
	// but the wire service could not be resolved before the tracking window closed
	// (see `error`).
	//
	// Any of "pending", "unknown", "imessage", "sms", "rcs".
	Protocol string `json:"protocol"`
	// Reactions on this message (tapbacks and emoji reactions)
	Reactions []Reaction `json:"reactions"`
	// Inline-reply parent reference. Identical shape on `message.received` webhooks
	// and on every GET endpoint that returns a single message or a list of messages.
	ReplyTo ChatMessageListResponseMessageReplyTo `json:"reply_to" api:"nullable"`
	// Sender's phone number or email for inbound group messages. Null for outbound
	// messages and 1-1 chats.
	Sender string `json:"sender" api:"nullable"`
	// Delivery lifecycle state. `pending` = persisted and being prepared for dispatch;
	// `queued` = accepted and waiting to be handed to Apple/the carrier; `sent` =
	// handed off to Apple/the carrier (protocol resolution happens around here);
	// `delivered` = a delivery receipt was received; `failed` = could not be delivered
	// (see `error`); `cancellation_requested` = a cancel was requested for a
	// still-queued message (best-effort); `cancelled` = cancelled before dispatch.
	// Inbound messages are surfaced via webhooks with `received`; read receipts arrive
	// as a `read` event.
	//
	// Any of "pending", "queued", "sent", "delivered", "failed",
	// "cancellation_requested", "cancelled".
	Status        string `json:"status" api:"nullable"`
	Text          string `json:"text" api:"nullable"`
	TimeDelivered int64  `json:"time_delivered" api:"nullable"`
	TimeSent      int64  `json:"time_sent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attachments   respjson.Field
		Direction     respjson.Field
		Error         respjson.Field
		ExternalID    respjson.Field
		InternalID    respjson.Field
		MessageID     respjson.Field
		Protocol      respjson.Field
		Reactions     respjson.Field
		ReplyTo       respjson.Field
		Sender        respjson.Field
		Status        respjson.Field
		Text          respjson.Field
		TimeDelivered respjson.Field
		TimeSent      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatMessageListResponseMessage) RawJSON() string { return r.JSON.raw }
func (r *ChatMessageListResponseMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inline-reply parent reference. Identical shape on `message.received` webhooks
// and on every GET endpoint that returns a single message or a list of messages.
type ChatMessageListResponseMessageReplyTo struct {
	// The raw iMessage GUID of the parent. Always populated on real inline replies;
	// the on-device record-of-truth identifier that survives even when `message_id`
	// cannot be resolved.
	Guid string `json:"guid" api:"required"`
	// The Blooio `message_id` of the parent message. NULL when the parent isn't in our
	// `messages` table (e.g., the original was sent from outside Blooio's pipeline).
	MessageID string `json:"message_id" api:"required"`
	// Which part of the parent was replied to. 0 for the common single-part case.
	PartIndex int64 `json:"part_index" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Guid        respjson.Field
		MessageID   respjson.Field
		PartIndex   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatMessageListResponseMessageReplyTo) RawJSON() string { return r.JSON.raw }
func (r *ChatMessageListResponseMessageReplyTo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatMessageGetStatusResponse struct {
	ChatID string `json:"chat_id"`
	// Any of "inbound", "outbound".
	Direction ChatMessageGetStatusResponseDirection `json:"direction"`
	Error     string                                `json:"error" api:"nullable"`
	MessageID string                                `json:"message_id"`
	// Transport used to carry the message; never null. `pending` = accepted and
	// dispatched, wire service not resolved yet (settles within seconds of send);
	// `imessage` = delivered over iMessage (blue bubble); `rcs` = delivered over RCS;
	// `sms` = fell back to SMS/MMS (green bubble); `unknown` = accepted by the carrier
	// but the wire service could not be resolved before the tracking window closed
	// (see `error`).
	//
	// Any of "pending", "unknown", "imessage", "sms", "rcs".
	Protocol ChatMessageGetStatusResponseProtocol `json:"protocol"`
	// Delivery lifecycle state. `pending` = persisted and being prepared for dispatch;
	// `queued` = accepted and waiting to be handed to Apple/the carrier; `sent` =
	// handed off to Apple/the carrier (protocol resolution happens around here);
	// `delivered` = a delivery receipt was received; `failed` = could not be delivered
	// (see `error`); `cancellation_requested` = a cancel was requested for a
	// still-queued message (best-effort); `cancelled` = cancelled before dispatch.
	// Inbound messages are surfaced via webhooks with `received`; read receipts arrive
	// as a `read` event.
	//
	// Any of "pending", "queued", "sent", "delivered", "failed",
	// "cancellation_requested", "cancelled".
	Status        ChatMessageGetStatusResponseStatus `json:"status" api:"nullable"`
	TimeDelivered int64                              `json:"time_delivered" api:"nullable"`
	TimeSent      int64                              `json:"time_sent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChatID        respjson.Field
		Direction     respjson.Field
		Error         respjson.Field
		MessageID     respjson.Field
		Protocol      respjson.Field
		Status        respjson.Field
		TimeDelivered respjson.Field
		TimeSent      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatMessageGetStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatMessageGetStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatMessageGetStatusResponseDirection string

const (
	ChatMessageGetStatusResponseDirectionInbound  ChatMessageGetStatusResponseDirection = "inbound"
	ChatMessageGetStatusResponseDirectionOutbound ChatMessageGetStatusResponseDirection = "outbound"
)

// Transport used to carry the message; never null. `pending` = accepted and
// dispatched, wire service not resolved yet (settles within seconds of send);
// `imessage` = delivered over iMessage (blue bubble); `rcs` = delivered over RCS;
// `sms` = fell back to SMS/MMS (green bubble); `unknown` = accepted by the carrier
// but the wire service could not be resolved before the tracking window closed
// (see `error`).
type ChatMessageGetStatusResponseProtocol string

const (
	ChatMessageGetStatusResponseProtocolPending  ChatMessageGetStatusResponseProtocol = "pending"
	ChatMessageGetStatusResponseProtocolUnknown  ChatMessageGetStatusResponseProtocol = "unknown"
	ChatMessageGetStatusResponseProtocolImessage ChatMessageGetStatusResponseProtocol = "imessage"
	ChatMessageGetStatusResponseProtocolSMS      ChatMessageGetStatusResponseProtocol = "sms"
	ChatMessageGetStatusResponseProtocolRcs      ChatMessageGetStatusResponseProtocol = "rcs"
)

// Delivery lifecycle state. `pending` = persisted and being prepared for dispatch;
// `queued` = accepted and waiting to be handed to Apple/the carrier; `sent` =
// handed off to Apple/the carrier (protocol resolution happens around here);
// `delivered` = a delivery receipt was received; `failed` = could not be delivered
// (see `error`); `cancellation_requested` = a cancel was requested for a
// still-queued message (best-effort); `cancelled` = cancelled before dispatch.
// Inbound messages are surfaced via webhooks with `received`; read receipts arrive
// as a `read` event.
type ChatMessageGetStatusResponseStatus string

const (
	ChatMessageGetStatusResponseStatusPending               ChatMessageGetStatusResponseStatus = "pending"
	ChatMessageGetStatusResponseStatusQueued                ChatMessageGetStatusResponseStatus = "queued"
	ChatMessageGetStatusResponseStatusSent                  ChatMessageGetStatusResponseStatus = "sent"
	ChatMessageGetStatusResponseStatusDelivered             ChatMessageGetStatusResponseStatus = "delivered"
	ChatMessageGetStatusResponseStatusFailed                ChatMessageGetStatusResponseStatus = "failed"
	ChatMessageGetStatusResponseStatusCancellationRequested ChatMessageGetStatusResponseStatus = "cancellation_requested"
	ChatMessageGetStatusResponseStatusCancelled             ChatMessageGetStatusResponseStatus = "cancelled"
)

type ChatMessageReactResponse struct {
	// The action that was performed
	//
	// Any of "add", "remove".
	Action ChatMessageReactResponseAction `json:"action"`
	// The ID of the message that was reacted to
	MessageID string `json:"message_id"`
	// The reaction that was added or removed. For classic tapbacks: love, like,
	// dislike, laugh, emphasize, question. For emoji reactions: the emoji character
	// (e.g. 😂, 👍, 🔥).
	Reaction string `json:"reaction"`
	// Whether the reaction was sent successfully
	Success bool `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		MessageID   respjson.Field
		Reaction    respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatMessageReactResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatMessageReactResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The action that was performed
type ChatMessageReactResponseAction string

const (
	ChatMessageReactResponseActionAdd    ChatMessageReactResponseAction = "add"
	ChatMessageReactResponseActionRemove ChatMessageReactResponseAction = "remove"
)

// Response after sending a message
type ChatMessageSendResponse struct {
	// Number of messages sent. Only present in URL-balloon batch mode.
	Count int64 `json:"count"`
	// True if a new unnamed group was created for this multi-recipient message
	GroupCreated bool `json:"group_created"`
	// Group ID when sending to multi-recipient (new or existing)
	GroupID string `json:"group_id"`
	// ID of the sent message (single-message sends)
	MessageID string `json:"message_id"`
	// IDs of sent messages. Present when `text` is an array or when `parts` uses
	// per-part `link_preview` (URL-balloon batch mode).
	MessageIDs []string `json:"message_ids"`
	// Present (and `true`) only when `reply_to.guid` was supplied without a
	// `message_id` and the GUID didn't map to any Blooio-minted row. The send still
	// proceeds and the device may still thread it; this flag signals that Blooio
	// couldn't link the new message to a known parent.
	ParentUnresolved bool `json:"parent_unresolved"`
	// List of participants (present for multi-recipient)
	Participants []string `json:"participants"`
	// Initial status of the message(s). `queued` = accepted for delivery (the normal
	// 202 result); `failed` = rejected before dispatch. Subsequent transitions (`sent`
	// → `delivered`, or `failed`) are reported via the status endpoint and
	// `message.status` webhooks.
	//
	// Any of "queued", "failed".
	Status ChatMessageSendResponseStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count            respjson.Field
		GroupCreated     respjson.Field
		GroupID          respjson.Field
		MessageID        respjson.Field
		MessageIDs       respjson.Field
		ParentUnresolved respjson.Field
		Participants     respjson.Field
		Status           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatMessageSendResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatMessageSendResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Initial status of the message(s). `queued` = accepted for delivery (the normal
// 202 result); `failed` = rejected before dispatch. Subsequent transitions (`sent`
// → `delivered`, or `failed`) are reported via the status endpoint and
// `message.status` webhooks.
type ChatMessageSendResponseStatus string

const (
	ChatMessageSendResponseStatusQueued ChatMessageSendResponseStatus = "queued"
	ChatMessageSendResponseStatusFailed ChatMessageSendResponseStatus = "failed"
)

type ChatMessageGetParams struct {
	ChatID string `path:"chatId" api:"required" json:"-"`
	paramObj
}

type ChatMessageListParams struct {
	// Maximum number of items to return in a single response. Must be between 1 and
	// 200; defaults to 50. Use together with `offset` to page through large result
	// sets.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip before returning results. Combine with `limit` for
	// page-based pagination (e.g. `offset=50&limit=50` returns the second page).
	// Defaults to 0.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Only messages sent after this timestamp (ms)
	Since param.Opt[int64] `query:"since,omitzero" json:"-"`
	// Only messages sent before this timestamp (ms)
	Until param.Opt[int64] `query:"until,omitzero" json:"-"`
	// Filter by message direction
	//
	// Any of "inbound", "outbound".
	Direction ChatMessageListParamsDirection `query:"direction,omitzero" json:"-"`
	// Sort order by time
	//
	// Any of "asc", "desc".
	Sort ChatMessageListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ChatMessageListParams]'s query parameters as `url.Values`.
func (r ChatMessageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by message direction
type ChatMessageListParamsDirection string

const (
	ChatMessageListParamsDirectionInbound  ChatMessageListParamsDirection = "inbound"
	ChatMessageListParamsDirectionOutbound ChatMessageListParamsDirection = "outbound"
)

// Sort order by time
type ChatMessageListParamsSort string

const (
	ChatMessageListParamsSortAsc  ChatMessageListParamsSort = "asc"
	ChatMessageListParamsSortDesc ChatMessageListParamsSort = "desc"
)

type ChatMessageGetStatusParams struct {
	ChatID string `path:"chatId" api:"required" json:"-"`
	paramObj
}

type ChatMessageReactParams struct {
	ChatID string `path:"chatId" api:"required" json:"-"`
	// The reaction to add or remove. Must be prefixed with `+` to add or `-` to
	// remove.
	//
	// **Classic tapbacks:** `+love`, `-love`, `+like`, `-like`, `+dislike`,
	// `-dislike`, `+laugh`, `-laugh`, `+emphasize`, `-emphasize`, `+question`,
	// `-question`
	//
	// **Emoji reactions:** Any emoji prefixed with `+` or `-` (e.g. `+😂`, `-😂`,
	// `+👍`, `-🔥`).
	Reaction string `json:"reaction" api:"required"`
	// Filter by message direction (only used when messageId is a relative index like
	// -1, -2)
	//
	// Any of "inbound", "outbound".
	Direction ChatMessageReactParamsDirection `json:"direction,omitzero"`
	paramObj
}

func (r ChatMessageReactParams) MarshalJSON() (data []byte, err error) {
	type shadow ChatMessageReactParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatMessageReactParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter by message direction (only used when messageId is a relative index like
// -1, -2)
type ChatMessageReactParamsDirection string

const (
	ChatMessageReactParamsDirectionInbound  ChatMessageReactParamsDirection = "inbound"
	ChatMessageReactParamsDirectionOutbound ChatMessageReactParamsDirection = "outbound"
)

type ChatMessageSendParams struct {
	// E.164 phone number to send from. For Twilio API keys, this is optional — if
	// omitted, the first assigned Twilio number is auto-selected. For Blooio
	// (iMessage) API keys, this selects a specific number from your pool. Must be a
	// number assigned to your API key.
	FromNumber param.Opt[string] `json:"from_number,omitzero"`
	// If true, the contact card (Name & Photo) will be shared with this message. The
	// contact card is piggybacked onto the outgoing message. Defaults to false. ⚠️
	// Only available on **Dedicated Commercial** and **Dedicated Enterprise** plans —
	// other plans receive a `403`.
	ShareContact param.Opt[bool] `json:"share_contact,omitzero"`
	// Whether to show typing indicator before sending. Defaults to org preference.
	UseTypingIndicator param.Opt[bool]   `json:"use_typing_indicator,omitzero"`
	IdempotencyKey     param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	// Optional. Attach an iMessage send-with-effect to the outgoing message.
	//
	// **Bubble effects** (apply to a single text bubble):
	//
	// - `slam` — Slam
	// - `loud` — Loud
	// - `gentle` — Gentle
	// - `invisible-ink` — Invisible Ink
	//
	// **Screen effects** (full-screen animation in the recipient's chat):
	//
	// - `echo` — Echo
	// - `spotlight` — Spotlight
	// - `balloons` — Balloons
	// - `confetti` — Confetti
	// - `love` — Love (heart)
	// - `lasers` — Lasers
	// - `fireworks` — Fireworks
	// - `celebration` — Celebration (sparkles)
	//
	// Values are case-insensitive and accept either dashes or spaces
	// (`"Invisible Ink"` and `"invisible-ink"` both work). Pass `"none"` or omit the
	// field to send without an effect.
	//
	// **Limitations:**
	//
	//   - iMessage-only — when the chat is delivered as SMS or RCS the message is sent
	//     without an animation.
	//   - Not supported alongside the `parts` array (multipart bubbles cannot carry an
	//     effect). Use the top-level `text` field instead.
	//   - When `text` is an array, every message in the array is sent with the same
	//     effect.
	//
	// Any of "slam", "loud", "gentle", "invisible-ink", "echo", "spotlight",
	// "balloons", "confetti", "love", "lasers", "fireworks", "celebration", "none".
	Effect ChatMessageSendParamsEffect `json:"effect,omitzero"`
	// Inline-reply target on `POST /chats/{chatId}/messages`. Pass either `message_id`
	// (preferred — references a Blooio-minted message) or `guid` (raw iMessage GUID,
	// useful for replying to messages received before the row was minted in Blooio).
	// The new send is dispatched to Lava with the resolved `selectedMessageGuid` +
	// `partIndex`, which iMessage renders as an inline reply on the recipient's
	// device.
	ReplyTo ChatMessageSendParamsReplyTo `json:"reply_to,omitzero"`
	// Array of attachment URLs or objects with url/name.
	//
	// **Voice memos:** a single audio file (`.mp3`, `.m4a`, `.wav`, `.aac`, `.opus`,
	// `.ogg`) is automatically sent as a voice memo (the native waveform/scrubber
	// bubble), not a plain audio-file attachment — no extra field is needed. A voice
	// memo is a standalone bubble, so it cannot be combined with `text` or any other
	// attachment; send the voice memo and the text as two separate messages.
	Attachments []ChatMessageSendParamsAttachmentUnion `json:"attachments,omitzero"`
	// Rich-link-preview overrides for URL messages (iMessage URL balloon). All fields
	// are optional. Only applies when the message text (or the concatenated part text)
	// is exactly a single http(s) URL. If omitted but the text is a URL, Blooio
	// auto-fetches the page's Open Graph metadata to generate a preview. If the image
	// download fails, the send still succeeds — Blooio silently falls back to the
	// auto-generated preview.
	LinkPreview LinkPreviewParam `json:"link_preview,omitzero"`
	// Ordered array of message parts. Two modes:
	//
	//  1. **Multipart mode** — parts sent as a single unified iMessage bubble (mix of
	//     text and attachment parts). This is the default.
	//  2. **URL-balloon batch mode** — triggered when any part has a `link_preview`
	//     object. Each part becomes its own rich-link-preview iMessage; parts are sent
	//     sequentially in array order. In batch mode every part must be text-only with
	//     `text` being a single http(s) URL. Response contains `message_ids[]` +
	//     `count` instead of `message_id`.
	Parts []ChatMessageSendParamsPart `json:"parts,omitzero"`
	// Message text. Can be a single string or array of strings (each becomes a
	// separate message)
	Text ChatMessageSendParamsTextUnion `json:"text,omitzero"`
	paramObj
}

func (r ChatMessageSendParams) MarshalJSON() (data []byte, err error) {
	type shadow ChatMessageSendParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatMessageSendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatMessageSendParamsAttachmentUnion struct {
	OfString                                        param.Opt[string]                                   `json:",omitzero,inline"`
	OfChatMessageSendsAttachmentUnionObjectVariant1 *ChatMessageSendParamsAttachmentUnionObjectVariant1 `json:",omitzero,inline"`
	paramUnion
}

func (u ChatMessageSendParamsAttachmentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfChatMessageSendsAttachmentUnionObjectVariant1)
}
func (u *ChatMessageSendParamsAttachmentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property URL is required.
type ChatMessageSendParamsAttachmentUnionObjectVariant1 struct {
	URL  string            `json:"url" api:"required"`
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r ChatMessageSendParamsAttachmentUnionObjectVariant1) MarshalJSON() (data []byte, err error) {
	type shadow ChatMessageSendParamsAttachmentUnionObjectVariant1
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatMessageSendParamsAttachmentUnionObjectVariant1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional. Attach an iMessage send-with-effect to the outgoing message.
//
// **Bubble effects** (apply to a single text bubble):
//
// - `slam` — Slam
// - `loud` — Loud
// - `gentle` — Gentle
// - `invisible-ink` — Invisible Ink
//
// **Screen effects** (full-screen animation in the recipient's chat):
//
// - `echo` — Echo
// - `spotlight` — Spotlight
// - `balloons` — Balloons
// - `confetti` — Confetti
// - `love` — Love (heart)
// - `lasers` — Lasers
// - `fireworks` — Fireworks
// - `celebration` — Celebration (sparkles)
//
// Values are case-insensitive and accept either dashes or spaces
// (`"Invisible Ink"` and `"invisible-ink"` both work). Pass `"none"` or omit the
// field to send without an effect.
//
// **Limitations:**
//
//   - iMessage-only — when the chat is delivered as SMS or RCS the message is sent
//     without an animation.
//   - Not supported alongside the `parts` array (multipart bubbles cannot carry an
//     effect). Use the top-level `text` field instead.
//   - When `text` is an array, every message in the array is sent with the same
//     effect.
type ChatMessageSendParamsEffect string

const (
	ChatMessageSendParamsEffectSlam         ChatMessageSendParamsEffect = "slam"
	ChatMessageSendParamsEffectLoud         ChatMessageSendParamsEffect = "loud"
	ChatMessageSendParamsEffectGentle       ChatMessageSendParamsEffect = "gentle"
	ChatMessageSendParamsEffectInvisibleInk ChatMessageSendParamsEffect = "invisible-ink"
	ChatMessageSendParamsEffectEcho         ChatMessageSendParamsEffect = "echo"
	ChatMessageSendParamsEffectSpotlight    ChatMessageSendParamsEffect = "spotlight"
	ChatMessageSendParamsEffectBalloons     ChatMessageSendParamsEffect = "balloons"
	ChatMessageSendParamsEffectConfetti     ChatMessageSendParamsEffect = "confetti"
	ChatMessageSendParamsEffectLove         ChatMessageSendParamsEffect = "love"
	ChatMessageSendParamsEffectLasers       ChatMessageSendParamsEffect = "lasers"
	ChatMessageSendParamsEffectFireworks    ChatMessageSendParamsEffect = "fireworks"
	ChatMessageSendParamsEffectCelebration  ChatMessageSendParamsEffect = "celebration"
	ChatMessageSendParamsEffectNone         ChatMessageSendParamsEffect = "none"
)

type ChatMessageSendParamsPart struct {
	// Participant phone number or email to @-mention. Only valid with 'text'. The
	// entire text of the part is rendered as the mention.
	Mention param.Opt[string] `json:"mention,omitzero"`
	// Filename for the attachment. Only valid with 'url'.
	Name param.Opt[string] `json:"name,omitzero"`
	// Text content for this part. Mutually exclusive with 'url'.
	Text param.Opt[string] `json:"text,omitzero"`
	// URL to an attachment for this part. Mutually exclusive with 'text'.
	URL param.Opt[string] `json:"url,omitzero"`
	// Rich-link-preview overrides for URL messages (iMessage URL balloon). All fields
	// are optional. Only applies when the message text (or the concatenated part text)
	// is exactly a single http(s) URL. If omitted but the text is a URL, Blooio
	// auto-fetches the page's Open Graph metadata to generate a preview. If the image
	// download fails, the send still succeeds — Blooio silently falls back to the
	// auto-generated preview.
	LinkPreview LinkPreviewParam `json:"link_preview,omitzero"`
	paramObj
}

func (r ChatMessageSendParamsPart) MarshalJSON() (data []byte, err error) {
	type shadow ChatMessageSendParamsPart
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatMessageSendParamsPart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inline-reply target on `POST /chats/{chatId}/messages`. Pass either `message_id`
// (preferred — references a Blooio-minted message) or `guid` (raw iMessage GUID,
// useful for replying to messages received before the row was minted in Blooio).
// The new send is dispatched to Lava with the resolved `selectedMessageGuid` +
// `partIndex`, which iMessage renders as an inline reply on the recipient's
// device.
type ChatMessageSendParamsReplyTo struct {
	// Raw iMessage GUID of the parent. When supplied without a `message_id`, Blooio
	// attempts to look up the parent via `provider_message_guid`; if the parent isn't
	// in our table the send still proceeds (Lava will thread on the device when
	// possible) and the response carries `parent_unresolved: true`.
	Guid param.Opt[string] `json:"guid,omitzero"`
	// Blooio `message_id` of the parent. Must belong to the same chat, same
	// from-number, and be no older than 30 days. Returns 404 `reply_target_not_found`
	// if unknown.
	MessageID param.Opt[string] `json:"message_id,omitzero"`
	// Which part of the parent to reply to. Defaults to 0 (covers the 99% case of
	// replying to a single-part text message).
	PartIndex param.Opt[int64] `json:"part_index,omitzero"`
	paramObj
}

func (r ChatMessageSendParamsReplyTo) MarshalJSON() (data []byte, err error) {
	type shadow ChatMessageSendParamsReplyTo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatMessageSendParamsReplyTo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatMessageSendParamsTextUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ChatMessageSendParamsTextUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ChatMessageSendParamsTextUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}
