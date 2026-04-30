// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blooio

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/blooio-go/internal/apijson"
	"github.com/stainless-sdks/blooio-go/internal/requestconfig"
	"github.com/stainless-sdks/blooio-go/option"
	"github.com/stainless-sdks/blooio-go/packages/respjson"
)

// Manage phone numbers linked to your account
//
// MeNumberService contains methods and other services that help with interacting
// with the blooio API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMeNumberService] method instead.
type MeNumberService struct {
	options []option.RequestOption
	// Manage and share your iMessage contact card (Name & Photo)
	ContactCard MeNumberContactCardService
}

// NewMeNumberService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMeNumberService(opts ...option.RequestOption) (r MeNumberService) {
	r = MeNumberService{}
	r.options = opts
	r.ContactCard = NewMeNumberContactCardService(opts...)
	return
}

// List all phone numbers bound to this API key with their availability status. Use
// the returned phone numbers as the `:number` path parameter for other
// `/me/numbers/` endpoints.
func (r *MeNumberService) List(ctx context.Context, opts ...option.RequestOption) (res *MeNumberListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "me/numbers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type MeNumberListResponse struct {
	Numbers []MeNumberListResponseNumber `json:"numbers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Numbers     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeNumberListResponse) RawJSON() string { return r.JSON.raw }
func (r *MeNumberListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeNumberListResponseNumber struct {
	IsActive    bool      `json:"is_active"`
	LastActive  time.Time `json:"last_active" api:"nullable" format:"date-time"`
	PhoneNumber string    `json:"phone_number"`
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
func (r MeNumberListResponseNumber) RawJSON() string { return r.JSON.raw }
func (r *MeNumberListResponseNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
