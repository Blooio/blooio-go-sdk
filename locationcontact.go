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
	"github.com/Blooio/blooio-go-sdk/internal/requestconfig"
	"github.com/Blooio/blooio-go-sdk/option"
	"github.com/Blooio/blooio-go-sdk/packages/respjson"
)

// FindMy contact location tracking
//
// LocationContactService contains methods and other services that help with
// interacting with the blooio API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLocationContactService] method instead.
type LocationContactService struct {
	options []option.RequestOption
}

// NewLocationContactService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLocationContactService(opts ...option.RequestOption) (r LocationContactService) {
	r = LocationContactService{}
	r.options = opts
	return
}

// Returns the cached location for a specific contact identified by phone number or
// email.
func (r *LocationContactService) Get(ctx context.Context, handle string, opts ...option.RequestOption) (res *ContactLocation, err error) {
	opts = slices.Concat(r.options, opts)
	if handle == "" {
		err = errors.New("missing required handle parameter")
		return nil, err
	}
	path := fmt.Sprintf("location/contacts/%s", url.PathEscape(handle))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns cached FindMy contact locations available through your blooio account.
// Each entry includes the contact's handle (phone/email), coordinates, and last
// update time.
func (r *LocationContactService) List(ctx context.Context, opts ...option.RequestOption) (res *LocationContactListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "location/contacts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Triggers a refresh of cached FindMy contact locations. Updated results may take
// 15-20 seconds to appear.
func (r *LocationContactService) Refresh(ctx context.Context, opts ...option.RequestOption) (res *LocationContactRefreshResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "location/contacts/refresh"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type ContactLocation struct {
	// GPS coordinates [latitude, longitude]
	Coordinates []float64 `json:"coordinates"`
	// Contact's phone number or email
	Handle string `json:"handle"`
	// Timestamp of last location update (epoch ms)
	LastUpdated int64 `json:"last_updated"`
	// Location status (e.g., 'live', 'shallow', 'legacy')
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Coordinates respjson.Field
		Handle      respjson.Field
		LastUpdated respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactLocation) RawJSON() string { return r.JSON.raw }
func (r *ContactLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LocationContactListResponse struct {
	Friends []ContactLocation `json:"friends"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Friends     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LocationContactListResponse) RawJSON() string { return r.JSON.raw }
func (r *LocationContactListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LocationContactRefreshResponse struct {
	Friends []ContactLocation `json:"friends"`
	Success bool              `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Friends     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LocationContactRefreshResponse) RawJSON() string { return r.JSON.raw }
func (r *LocationContactRefreshResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
