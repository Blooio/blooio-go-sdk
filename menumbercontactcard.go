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
	"github.com/Blooio/blooio-go-sdk/packages/param"
	"github.com/Blooio/blooio-go-sdk/packages/respjson"
)

// Manage and share your iMessage contact card (Name & Photo)
//
// MeNumberContactCardService contains methods and other services that help with
// interacting with the blooio API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMeNumberContactCardService] method instead.
type MeNumberContactCardService struct {
	options []option.RequestOption
}

// NewMeNumberContactCardService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMeNumberContactCardService(opts ...option.RequestOption) (r MeNumberContactCardService) {
	r = MeNumberContactCardService{}
	r.options = opts
	return
}

// ⚠️ **COMING SOON** - This endpoint is temporarily disabled while we stabilize
// this feature.
//
// Get the personal contact card (Name & Photo) for the specified phone number.
// This is the identity that gets shared with contacts in iMessage.
func (r *MeNumberContactCardService) Get(ctx context.Context, number string, opts ...option.RequestOption) (res *MeNumberContactCardGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if number == "" {
		err = errors.New("missing required number parameter")
		return nil, err
	}
	path := fmt.Sprintf("me/numbers/%s/contact-card", url.PathEscape(number))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update the personal contact card (Name & Photo) for the specified phone number.
// All fields are optional — only provided fields are updated.
func (r *MeNumberContactCardService) Update(ctx context.Context, number string, body MeNumberContactCardUpdateParams, opts ...option.RequestOption) (res *MeNumberContactCardUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if number == "" {
		err = errors.New("missing required number parameter")
		return nil, err
	}
	path := fmt.Sprintf("me/numbers/%s/contact-card", url.PathEscape(number))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

type MeNumberContactCardGetResponse struct {
	// Base64-encoded JPEG/PNG image
	Avatar       string `json:"avatar" api:"nullable"`
	FirstName    string `json:"first_name" api:"nullable"`
	HasWallpaper bool   `json:"has_wallpaper"`
	LastName     string `json:"last_name" api:"nullable"`
	// Display name
	Name        string                                `json:"name" api:"nullable"`
	PhoneNumber string                                `json:"phone_number"`
	Sharing     MeNumberContactCardGetResponseSharing `json:"sharing"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Avatar       respjson.Field
		FirstName    respjson.Field
		HasWallpaper respjson.Field
		LastName     respjson.Field
		Name         respjson.Field
		PhoneNumber  respjson.Field
		Sharing      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeNumberContactCardGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MeNumberContactCardGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeNumberContactCardGetResponseSharing struct {
	// 0 = Contacts Only, 1 = Always Ask
	Audience int64 `json:"audience"`
	// Whether Name & Photo sharing is enabled
	Enabled bool `json:"enabled"`
	// 0 = First & Last, 1 = First Only
	NameFormat int64 `json:"name_format"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Audience    respjson.Field
		Enabled     respjson.Field
		NameFormat  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeNumberContactCardGetResponseSharing) RawJSON() string { return r.JSON.raw }
func (r *MeNumberContactCardGetResponseSharing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeNumberContactCardUpdateResponse struct {
	FirstName   string `json:"first_name" api:"nullable"`
	LastName    string `json:"last_name" api:"nullable"`
	PhoneNumber string `json:"phone_number"`
	Success     bool   `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstName   respjson.Field
		LastName    respjson.Field
		PhoneNumber respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeNumberContactCardUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *MeNumberContactCardUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeNumberContactCardUpdateParams struct {
	// Profile photo as base64-encoded JPEG/PNG
	Avatar param.Opt[string] `json:"avatar,omitzero"`
	// First name
	FirstName param.Opt[string] `json:"first_name,omitzero"`
	// Last name
	LastName param.Opt[string]                      `json:"last_name,omitzero"`
	Sharing  MeNumberContactCardUpdateParamsSharing `json:"sharing,omitzero"`
	paramObj
}

func (r MeNumberContactCardUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow MeNumberContactCardUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeNumberContactCardUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeNumberContactCardUpdateParamsSharing struct {
	// 0 = Contacts Only, 1 = Always Ask
	Audience param.Opt[int64] `json:"audience,omitzero"`
	// Enable/disable Name & Photo sharing
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// 0 = First & Last, 1 = First Only
	NameFormat param.Opt[int64] `json:"name_format,omitzero"`
	paramObj
}

func (r MeNumberContactCardUpdateParamsSharing) MarshalJSON() (data []byte, err error) {
	type shadow MeNumberContactCardUpdateParamsSharing
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeNumberContactCardUpdateParamsSharing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
