// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blooio

import (
	"context"
	"net/http"
	"slices"

	"github.com/Blooio/blooio-go-sdk/internal/apijson"
	"github.com/Blooio/blooio-go-sdk/internal/requestconfig"
	"github.com/Blooio/blooio-go-sdk/option"
	"github.com/Blooio/blooio-go-sdk/packages/respjson"
)

// MeService contains methods and other services that help with interacting with
// the blooio API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMeService] method instead.
type MeService struct {
	Options []option.RequestOption
}

// NewMeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMeService(opts ...option.RequestOption) (r MeService) {
	r = MeService{}
	r.Options = opts
	return
}

// Returns information about the authenticated API key including plan, devices,
// usage statistics, and integration details.
func (r *MeService) Get(ctx context.Context, opts ...option.RequestOption) (res *MeGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/api/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type MeGetResponse struct {
	// The API key used for authentication.
	APIKey string `json:"api_key"`
	// List of devices associated with this API key.
	Devices []MeGetResponseDevice `json:"devices"`
	// Integration-specific details (GHL or API integration).
	IntegrationDetails MeGetResponseIntegrationDetails `json:"integration_details,nullable"`
	// Custom metadata associated with the API key.
	Metadata any `json:"metadata"`
	// The plan associated with this API key.
	Plan string `json:"plan"`
	// Usage statistics for this API key.
	Usage MeGetResponseUsage `json:"usage"`
	// Whether the API key is valid.
	Valid bool `json:"valid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKey             respjson.Field
		Devices            respjson.Field
		IntegrationDetails respjson.Field
		Metadata           respjson.Field
		Plan               respjson.Field
		Usage              respjson.Field
		Valid              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MeGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeGetResponseDevice struct {
	// Hashed device identifier.
	DeviceHash string `json:"device_hash"`
	// Whether the device is currently active.
	IsActive bool `json:"is_active"`
	// Unix timestamp (ms) of last device activity.
	LastActive int64 `json:"last_active,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeviceHash  respjson.Field
		IsActive    respjson.Field
		LastActive  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeGetResponseDevice) RawJSON() string { return r.JSON.raw }
func (r *MeGetResponseDevice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Integration-specific details (GHL or API integration).
type MeGetResponseIntegrationDetails struct {
	// Webhook URL for API integrations.
	CustomerWebhookURL string `json:"customer_webhook_url" format:"uri"`
	// Integration-specific metadata.
	Metadata any `json:"metadata"`
	// Name of the integration (GHL only).
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomerWebhookURL respjson.Field
		Metadata           respjson.Field
		Name               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeGetResponseIntegrationDetails) RawJSON() string { return r.JSON.raw }
func (r *MeGetResponseIntegrationDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Usage statistics for this API key.
type MeGetResponseUsage struct {
	// Total number of inbound messages.
	InboundMessages int64 `json:"inbound_messages"`
	// Unix timestamp (ms) of the last message sent.
	LastMessageSent int64 `json:"last_message_sent,nullable"`
	// Total number of outbound messages.
	OutboundMessages int64 `json:"outbound_messages"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InboundMessages  respjson.Field
		LastMessageSent  respjson.Field
		OutboundMessages respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeGetResponseUsage) RawJSON() string { return r.JSON.raw }
func (r *MeGetResponseUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
