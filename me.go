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

// Authentication and account information
//
// MeService contains methods and other services that help with interacting with
// the blooio API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMeService] method instead.
type MeService struct {
	options []option.RequestOption
	// Manage phone numbers linked to your account
	Numbers MeNumberService
}

// NewMeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMeService(opts ...option.RequestOption) (r MeService) {
	r = MeService{}
	r.options = opts
	r.Numbers = NewMeNumberService(opts...)
	return
}

// Returns details about the authenticated API key or dashboard user, including
// organization info, devices, and usage statistics.
func (r *MeService) Get(ctx context.Context, opts ...option.RequestOption) (res *MeGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Response depends on auth_type. For 'api_key': includes full API key details. For
// 'dashboard': includes user_id and organization info only.
type MeGetResponse struct {
	// The API key (only for api_key auth)
	APIKey string `json:"api_key"`
	// Type of authentication used
	//
	// Any of "api_key", "dashboard".
	AuthType MeGetResponseAuthType `json:"auth_type"`
	// List of devices associated with this API key (only for api_key auth)
	Devices []MeGetResponseDevice `json:"devices"`
	// Integration details if the API key is associated with an integration (only for
	// api_key auth)
	IntegrationDetails any `json:"integration_details" api:"nullable"`
	// API key metadata (only for api_key auth)
	Metadata     any                       `json:"metadata"`
	Organization MeGetResponseOrganization `json:"organization"`
	// Organization ID (only for api_key auth)
	OrganizationID string `json:"organization_id"`
	// Usage statistics (only for api_key auth)
	Usage MeGetResponseUsage `json:"usage"`
	// User ID (only for dashboard auth)
	UserID string `json:"user_id" api:"nullable"`
	// Whether the API key is valid (only for api_key auth)
	Valid bool `json:"valid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKey             respjson.Field
		AuthType           respjson.Field
		Devices            respjson.Field
		IntegrationDetails respjson.Field
		Metadata           respjson.Field
		Organization       respjson.Field
		OrganizationID     respjson.Field
		Usage              respjson.Field
		UserID             respjson.Field
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

// Type of authentication used
type MeGetResponseAuthType string

const (
	MeGetResponseAuthTypeAPIKey    MeGetResponseAuthType = "api_key"
	MeGetResponseAuthTypeDashboard MeGetResponseAuthType = "dashboard"
)

type MeGetResponseDevice struct {
	IsActive   bool  `json:"is_active"`
	LastActive int64 `json:"last_active" api:"nullable"`
	// Phone number assigned to this device (E.164 format)
	PhoneNumber string `json:"phone_number" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsActive    respjson.Field
		LastActive  respjson.Field
		PhoneNumber respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeGetResponseDevice) RawJSON() string { return r.JSON.raw }
func (r *MeGetResponseDevice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeGetResponseOrganization struct {
	CountryCode    string `json:"country_code" api:"nullable"`
	CreatedAt      int64  `json:"created_at"`
	Name           string `json:"name"`
	OrganizationID string `json:"organization_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CountryCode    respjson.Field
		CreatedAt      respjson.Field
		Name           respjson.Field
		OrganizationID respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeGetResponseOrganization) RawJSON() string { return r.JSON.raw }
func (r *MeGetResponseOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Usage statistics (only for api_key auth)
type MeGetResponseUsage struct {
	InboundMessages  int64 `json:"inbound_messages"`
	LastMessageSent  int64 `json:"last_message_sent" api:"nullable"`
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
