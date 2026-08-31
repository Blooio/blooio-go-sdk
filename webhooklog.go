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

// View and replay webhook deliveries
//
// WebhookLogService contains methods and other services that help with interacting
// with the blooio API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookLogService] method instead.
type WebhookLogService struct {
	options []option.RequestOption
}

// NewWebhookLogService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWebhookLogService(opts ...option.RequestOption) (r WebhookLogService) {
	r = WebhookLogService{}
	r.options = opts
	return
}

// List delivery logs for a specific webhook.
func (r *WebhookLogService) List(ctx context.Context, webhookID string, query WebhookLogListParams, opts ...option.RequestOption) (res *WebhookLogListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks/%s/logs", url.PathEscape(webhookID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Re-send a webhook event to the configured URL.
func (r *WebhookLogService) Replay(ctx context.Context, eventID string, body WebhookLogReplayParams, opts ...option.RequestOption) (res *WebhookLogReplayResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.WebhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return nil, err
	}
	if eventID == "" {
		err = errors.New("missing required eventId parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks/%s/logs/%s/replay", url.PathEscape(body.WebhookID), url.PathEscape(eventID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type WebhookLogListResponse struct {
	Logs       []WebhookLogListResponseLog      `json:"logs"`
	Pagination WebhookLogListResponsePagination `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Logs        respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookLogListResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookLogListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookLogListResponseLog struct {
	AttemptedTime int64 `json:"attempted_time"`
	// Webhook event payload. Structure varies by event type. All message events
	// include group information when applicable.
	EventBody WebhookLogListResponseLogEventBody `json:"event_body"`
	EventID   string                             `json:"event_id"`
	// Additional metadata about the webhook delivery
	Metadata WebhookLogListResponseLogMetadata `json:"metadata"`
	// Response body from the webhook endpoint (if JSON)
	ResponseJson       any   `json:"response_json" api:"nullable"`
	ResponseReceivedAt int64 `json:"response_received_at" api:"nullable"`
	// HTTP status code received from the webhook endpoint
	ResponseStatus int64 `json:"response_status" api:"nullable"`
	// Any of "api", "integration", "org".
	Scope      string `json:"scope"`
	WebhookURL string `json:"webhook_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttemptedTime      respjson.Field
		EventBody          respjson.Field
		EventID            respjson.Field
		Metadata           respjson.Field
		ResponseJson       respjson.Field
		ResponseReceivedAt respjson.Field
		ResponseStatus     respjson.Field
		Scope              respjson.Field
		WebhookURL         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookLogListResponseLog) RawJSON() string { return r.JSON.raw }
func (r *WebhookLogListResponseLog) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Webhook event payload. Structure varies by event type. All message events
// include group information when applicable.
type WebhookLogListResponseLogEventBody struct {
	// Array of attachment objects
	Attachments []WebhookLogListResponseLogEventBodyAttachment `json:"attachments" api:"nullable"`
	// The device's own identifier for the group conversation this message arrived in
	// (only on `message.received` when is_group=true). Two group chats can hold the
	// same members and are then indistinguishable by `group_id` and `participants`
	// alone; `chat_guid` is what tells them apart. Matches the `chat_guid` on GET
	// /groups/{groupId}.
	ChatGuid string `json:"chat_guid" api:"nullable"`
	// The name the device reports for the conversation (only on `message.received`
	// when is_group=true). May differ from `group_name`, or be present when
	// `group_name` is null.
	ChatName string `json:"chat_name" api:"nullable"`
	// Timestamp when message was delivered (for message.delivered events)
	DeliveredAt int64 `json:"delivered_at" api:"nullable"`
	// Error code (for message.failed events)
	ErrorCode string `json:"error_code" api:"nullable"`
	// Error description (for message.failed events)
	ErrorMessage string `json:"error_message" api:"nullable"`
	// Event type (e.g., message.received, message.sent, message.delivered,
	// message.failed, message.read)
	Event string `json:"event"`
	// Recipient identifier (phone number, email, or group ID)
	ExternalID string `json:"external_id"`
	// Group ID (only present when is_group=true)
	GroupID string `json:"group_id" api:"nullable"`
	// Group display name (only present when is_group=true)
	GroupName string `json:"group_name" api:"nullable"`
	// Phone number that sent/received the message
	InternalID string `json:"internal_id" api:"nullable"`
	// Whether this message is from/to a group chat. Always present.
	IsGroup bool `json:"is_group"`
	// Unique message identifier
	MessageID string `json:"message_id"`
	// Array of group participants (only present when is_group=true). One entry per
	// person: a participant appears once even if Blooio holds more than one identity
	// for their number.
	Participants []WebhookLogListResponseLogEventBodyParticipant `json:"participants" api:"nullable"`
	// Transport used to carry the message; never null. `pending` = accepted and
	// dispatched, wire service not resolved yet (settles within seconds of send);
	// `imessage` = delivered over iMessage (blue bubble); `rcs` = delivered over RCS;
	// `sms` = fell back to SMS/MMS (green bubble); `unknown` = accepted by the carrier
	// but the wire service could not be resolved before the tracking window closed
	// (see `error`).
	//
	// Any of "pending", "unknown", "imessage", "sms", "rcs".
	Protocol string `json:"protocol"`
	// Timestamp when message was read (for message.read events)
	ReadAt int64 `json:"read_at" api:"nullable"`
	// Sender identifier (for inbound messages)
	Sender string `json:"sender" api:"nullable"`
	// Timestamp when message was sent (for message.sent events)
	SentAt int64 `json:"sent_at" api:"nullable"`
	// Message status carried by the event. `queued` / `pending` = accepted, not yet
	// handed off; `sent` = handed to Apple/the carrier; `delivered` = a delivery
	// receipt was received; `read` = a read receipt was received (iMessage, when the
	// recipient has read receipts on); `failed` = delivery failed (see `error_code` /
	// `error_message`); `received` = an inbound message arrived.
	//
	// Any of "queued", "pending", "sent", "delivered", "failed", "read", "received".
	Status string `json:"status"`
	// Message text content
	Text string `json:"text" api:"nullable"`
	// Event timestamp in milliseconds
	Timestamp int64 `json:"timestamp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attachments  respjson.Field
		ChatGuid     respjson.Field
		ChatName     respjson.Field
		DeliveredAt  respjson.Field
		ErrorCode    respjson.Field
		ErrorMessage respjson.Field
		Event        respjson.Field
		ExternalID   respjson.Field
		GroupID      respjson.Field
		GroupName    respjson.Field
		InternalID   respjson.Field
		IsGroup      respjson.Field
		MessageID    respjson.Field
		Participants respjson.Field
		Protocol     respjson.Field
		ReadAt       respjson.Field
		Sender       respjson.Field
		SentAt       respjson.Field
		Status       respjson.Field
		Text         respjson.Field
		Timestamp    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookLogListResponseLogEventBody) RawJSON() string { return r.JSON.raw }
func (r *WebhookLogListResponseLogEventBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookLogListResponseLogEventBodyAttachment struct {
	Name string `json:"name" api:"nullable"`
	URL  string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookLogListResponseLogEventBodyAttachment) RawJSON() string { return r.JSON.raw }
func (r *WebhookLogListResponseLogEventBodyAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookLogListResponseLogEventBodyParticipant struct {
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
func (r WebhookLogListResponseLogEventBodyParticipant) RawJSON() string { return r.JSON.raw }
func (r *WebhookLogListResponseLogEventBodyParticipant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional metadata about the webhook delivery
type WebhookLogListResponseLogMetadata struct {
	DurationMs      int64  `json:"duration_ms"`
	EventName       string `json:"event_name"`
	IsReplay        bool   `json:"is_replay"`
	MessageID       string `json:"message_id"`
	OrganizationID  string `json:"organization_id"`
	OriginalEventID string `json:"original_event_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationMs      respjson.Field
		EventName       respjson.Field
		IsReplay        respjson.Field
		MessageID       respjson.Field
		OrganizationID  respjson.Field
		OriginalEventID respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookLogListResponseLogMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebhookLogListResponseLogMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookLogListResponsePagination struct {
	// Whether there are more logs to fetch
	HasMore bool  `json:"has_more"`
	Limit   int64 `json:"limit"`
	Offset  int64 `json:"offset"`
	// Number of logs returned in this response
	Returned int64 `json:"returned"`
	// Total number of matching logs
	Total int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		Limit       respjson.Field
		Offset      respjson.Field
		Returned    respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookLogListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *WebhookLogListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookLogReplayResponse struct {
	// Time taken for the replay request in milliseconds
	DurationMs int64 `json:"duration_ms"`
	// The original event ID that was replayed
	OriginalEventID string `json:"original_event_id"`
	// New event ID for this replay attempt
	ReplayEventID string `json:"replay_event_id"`
	// Response details from the replay attempt
	ResponseData WebhookLogReplayResponseResponseData `json:"response_data"`
	// HTTP status code from replay attempt
	ResponseStatus int64 `json:"response_status"`
	// Whether the replay received a 2xx response
	Success    bool   `json:"success"`
	WebhookID  string `json:"webhook_id"`
	WebhookURL string `json:"webhook_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationMs      respjson.Field
		OriginalEventID respjson.Field
		ReplayEventID   respjson.Field
		ResponseData    respjson.Field
		ResponseStatus  respjson.Field
		Success         respjson.Field
		WebhookID       respjson.Field
		WebhookURL      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookLogReplayResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookLogReplayResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response details from the replay attempt
type WebhookLogReplayResponseResponseData struct {
	// Response body (if parseable)
	Body        any    `json:"body"`
	ContentType string `json:"contentType"`
	Duration    int64  `json:"duration"`
	Error       string `json:"error" api:"nullable"`
	ErrorType   string `json:"errorType" api:"nullable"`
	Headers     any    `json:"headers"`
	Size        int64  `json:"size"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		ContentType respjson.Field
		Duration    respjson.Field
		Error       respjson.Field
		ErrorType   respjson.Field
		Headers     respjson.Field
		Size        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookLogReplayResponseResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookLogReplayResponseResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookLogListParams struct {
	// Maximum number of items to return in a single response. Must be between 1 and
	// 200; defaults to 50. Use together with `offset` to page through large result
	// sets.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Maximum HTTP status code
	MaxStatus param.Opt[int64] `query:"max_status,omitzero" json:"-"`
	// Minimum HTTP status code
	MinStatus param.Opt[int64] `query:"min_status,omitzero" json:"-"`
	// Number of items to skip before returning results. Combine with `limit` for
	// page-based pagination (e.g. `offset=50&limit=50` returns the second page).
	// Defaults to 0.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Filter by exact HTTP status code
	Status param.Opt[int64] `query:"status,omitzero" json:"-"`
	// Sort order by attempted time
	//
	// Any of "asc", "desc".
	Sort WebhookLogListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookLogListParams]'s query parameters as `url.Values`.
func (r WebhookLogListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order by attempted time
type WebhookLogListParamsSort string

const (
	WebhookLogListParamsSortAsc  WebhookLogListParamsSort = "asc"
	WebhookLogListParamsSortDesc WebhookLogListParamsSort = "desc"
)

type WebhookLogReplayParams struct {
	WebhookID string `path:"webhookId" api:"required" json:"-"`
	paramObj
}
