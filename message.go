// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blooio

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/blooio-go/internal/apijson"
	"github.com/stainless-sdks/blooio-go/internal/requestconfig"
	"github.com/stainless-sdks/blooio-go/option"
	"github.com/stainless-sdks/blooio-go/packages/param"
	"github.com/stainless-sdks/blooio-go/packages/respjson"
)

// MessageService contains methods and other services that help with interacting
// with the blooio API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessageService] method instead.
type MessageService struct {
	Options []option.RequestOption
}

// NewMessageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMessageService(opts ...option.RequestOption) (r MessageService) {
	r = MessageService{}
	r.Options = opts
	return
}

// Retrieve full message metadata including direction, protocol, text length,
// attachments count, timestamps, current status, and original metadata.
func (r *MessageService) Get(ctx context.Context, messageID string, opts ...option.RequestOption) (res *MessageGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if messageID == "" {
		err = errors.New("missing required messageId parameter")
		return
	}
	path := fmt.Sprintf("v1/api/messages/%s", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Request cancellation for a queued or scheduled message. If the message is
// already in a non-cancellable state (e.g., sent), the response indicates the
// current status. Note: response may include deprecated `current_status`; use
// `status` instead.
func (r *MessageService) Cancel(ctx context.Context, messageID string, opts ...option.RequestOption) (res *MessageCancelResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if messageID == "" {
		err = errors.New("missing required messageId parameter")
		return
	}
	path := fmt.Sprintf("v1/api/messages/%s", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get current delivery status only
func (r *MessageService) GetStatus(ctx context.Context, messageID string, opts ...option.RequestOption) (res *MessageGetStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if messageID == "" {
		err = errors.New("missing required messageId parameter")
		return
	}
	path := fmt.Sprintf("v1/api/messages/%s/status", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Queues an outbound iMessage/SMS to the specified phone number.
//
// Supports optional media attachments via public URLs, and arbitrary metadata for
// correlation.
//
// For safe retries, provide an `Idempotency-Key` header. If the same key is used
// twice, the original message will be returned with 200; otherwise a new message
// is queued with 202.
func (r *MessageService) Send(ctx context.Context, params MessageSendParams, opts ...option.RequestOption) (res *MessageSendResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%s", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/api/messages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type MessageGetResponse struct {
	APIKey           string `json:"api_key"`
	AttachmentsCount int64  `json:"attachments_count"`
	// Any of "outbound", "inbound".
	Direction MessageGetResponseDirection `json:"direction"`
	// Recipient phone number.
	ExternalID string `json:"external_id"`
	MessageID  string `json:"message_id"`
	// Original metadata provided plus system generated fields.
	Metadata any `json:"metadata"`
	// The protocol used to send the message (e.g., imessage, rcs, sms).
	Protocol string `json:"protocol"`
	// Current delivery status.
	//
	// Any of "pending", "queued", "sent", "delivered", "failed", "cancelled",
	// "cancellation_requested".
	Status     MessageGetResponseStatus `json:"status"`
	TextLength int64                    `json:"text_length"`
	// Unix timestamp (ms) when the message was queued/sent.
	TimeSent int64 `json:"time_sent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKey           respjson.Field
		AttachmentsCount respjson.Field
		Direction        respjson.Field
		ExternalID       respjson.Field
		MessageID        respjson.Field
		Metadata         respjson.Field
		Protocol         respjson.Field
		Status           respjson.Field
		TextLength       respjson.Field
		TimeSent         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageGetResponseDirection string

const (
	MessageGetResponseDirectionOutbound MessageGetResponseDirection = "outbound"
	MessageGetResponseDirectionInbound  MessageGetResponseDirection = "inbound"
)

// Current delivery status.
type MessageGetResponseStatus string

const (
	MessageGetResponseStatusPending               MessageGetResponseStatus = "pending"
	MessageGetResponseStatusQueued                MessageGetResponseStatus = "queued"
	MessageGetResponseStatusSent                  MessageGetResponseStatus = "sent"
	MessageGetResponseStatusDelivered             MessageGetResponseStatus = "delivered"
	MessageGetResponseStatusFailed                MessageGetResponseStatus = "failed"
	MessageGetResponseStatusCancelled             MessageGetResponseStatus = "cancelled"
	MessageGetResponseStatusCancellationRequested MessageGetResponseStatus = "cancellation_requested"
)

type MessageCancelResponse struct {
	// True if cancellation was successful, false otherwise.
	Cancelled bool `json:"cancelled"`
	// The current status if cancellation failed (deprecated, use 'status' instead).
	CurrentStatus string `json:"current_status,nullable"`
	// Human-readable message about the cancellation result.
	Message   string `json:"message"`
	MessageID string `json:"message_id"`
	// The current status of the message.
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cancelled     respjson.Field
		CurrentStatus respjson.Field
		Message       respjson.Field
		MessageID     respjson.Field
		Status        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageCancelResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageCancelResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageGetStatusResponse struct {
	MessageID string `json:"message_id"`
	// Any of "pending", "queued", "sent", "delivered", "failed", "cancelled",
	// "cancellation_requested".
	Status MessageGetStatusResponseStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MessageID   respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageGetStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageGetStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageGetStatusResponseStatus string

const (
	MessageGetStatusResponseStatusPending               MessageGetStatusResponseStatus = "pending"
	MessageGetStatusResponseStatusQueued                MessageGetStatusResponseStatus = "queued"
	MessageGetStatusResponseStatusSent                  MessageGetStatusResponseStatus = "sent"
	MessageGetStatusResponseStatusDelivered             MessageGetStatusResponseStatus = "delivered"
	MessageGetStatusResponseStatusFailed                MessageGetStatusResponseStatus = "failed"
	MessageGetStatusResponseStatusCancelled             MessageGetStatusResponseStatus = "cancelled"
	MessageGetStatusResponseStatusCancellationRequested MessageGetStatusResponseStatus = "cancellation_requested"
)

type MessageSendResponse struct {
	MessageID string `json:"message_id"`
	Status    string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MessageID   respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageSendResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendParams struct {
	// Recipient phone number in E.164 format (e.g., +15551234567)
	To string `json:"to,required"`
	// Text body of the message.
	Text           param.Opt[string] `json:"text,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	// Array of publicly accessible URLs to media attachments.
	Attachments []string `json:"attachments,omitzero" format:"uri"`
	// Arbitrary key/value pairs to associate with the message.
	Metadata any `json:"metadata,omitzero"`
	paramObj
}

func (r MessageSendParams) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
