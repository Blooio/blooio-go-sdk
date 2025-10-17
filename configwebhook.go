// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blooio

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/blooio-go/internal/apijson"
	"github.com/stainless-sdks/blooio-go/internal/requestconfig"
	"github.com/stainless-sdks/blooio-go/option"
	"github.com/stainless-sdks/blooio-go/packages/param"
	"github.com/stainless-sdks/blooio-go/packages/respjson"
)

// ConfigWebhookService contains methods and other services that help with
// interacting with the blooio API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConfigWebhookService] method instead.
type ConfigWebhookService struct {
	Options []option.RequestOption
}

// NewConfigWebhookService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConfigWebhookService(opts ...option.RequestOption) (r ConfigWebhookService) {
	r = ConfigWebhookService{}
	r.Options = opts
	return
}

// Get the current webhook URL
func (r *ConfigWebhookService) Get(ctx context.Context, opts ...option.RequestOption) (res *ConfigWebhookGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/api/config/webhook"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Configure the webhook URL that will receive all message events. Once configured,
// your endpoint will receive POST requests with the following event types:
//
//   - `message.received` - When you receive an inbound message
//   - `message.sent` - When your outbound message is sent to the carrier
//   - `message.delivered` - When your outbound message is delivered to the recipient
//   - `message.failed` - When your outbound message fails to deliver
//   - `message.read` - When your outbound message is read by the recipient (iMessage
//     only, recipient must have read receipts enabled)
//
// See the webhook event schemas for detailed payload formats.
func (r *ConfigWebhookService) Update(ctx context.Context, body ConfigWebhookUpdateParams, opts ...option.RequestOption) (res *ConfigWebhookUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/api/config/webhook"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ConfigWebhookGetResponse struct {
	// Unix timestamp (ms) when the webhook URL was last updated.
	UpdatedAt int64 `json:"updated_at"`
	// The current webhook URL or null if not set.
	WebhookURL string `json:"webhook_url,nullable" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		UpdatedAt   respjson.Field
		WebhookURL  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigWebhookGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ConfigWebhookGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigWebhookUpdateResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigWebhookUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ConfigWebhookUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigWebhookUpdateParams struct {
	WebhookURL string `json:"webhook_url,required" format:"uri"`
	paramObj
}

func (r ConfigWebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ConfigWebhookUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConfigWebhookUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
