// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blooio

import (
	"context"
	"net/http"
	"slices"

	"github.com/Blooio/blooio-go-sdk/internal/apijson"
	"github.com/Blooio/blooio-go-sdk/internal/requestconfig"
	"github.com/Blooio/blooio-go-sdk/option"
	"github.com/Blooio/blooio-go-sdk/packages/param"
	"github.com/Blooio/blooio-go-sdk/packages/respjson"
)

// Phone number validation, formatting, and NANPA geocoding. Requires an Enterprise
// plan (Dedicated Enterprise).
//
// PhoneNumberService contains methods and other services that help with
// interacting with the blooio API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPhoneNumberService] method instead.
type PhoneNumberService struct {
	options []option.RequestOption
	// Phone number validation, formatting, and NANPA geocoding. Requires an Enterprise
	// plan (Dedicated Enterprise).
	Lookup PhoneNumberLookupService
}

// NewPhoneNumberService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPhoneNumberService(opts ...option.RequestOption) (r PhoneNumberService) {
	r = PhoneNumberService{}
	r.options = opts
	r.Lookup = NewPhoneNumberLookupService(opts...)
	return
}

// Look up multiple phone numbers in a single request. Returns the same detailed
// information as the single lookup endpoint for each number. Maximum 100 numbers
// per request.
//
// **Requires an Enterprise plan** (Dedicated Enterprise).
func (r *PhoneNumberService) BatchNew(ctx context.Context, body PhoneNumberBatchNewParams, opts ...option.RequestOption) (res *PhoneNumberBatchNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "phone-numbers/batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type PhoneNumberBatchNewResponse struct {
	Results []PhoneNumberLookupResult `json:"results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberBatchNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberBatchNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberBatchNewParams struct {
	// Array of phone numbers to look up
	Numbers []string `json:"numbers,omitzero" api:"required"`
	paramObj
}

func (r PhoneNumberBatchNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PhoneNumberBatchNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PhoneNumberBatchNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
