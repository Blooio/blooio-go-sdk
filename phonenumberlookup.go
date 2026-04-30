// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blooio

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/blooio-go/internal/apijson"
	"github.com/stainless-sdks/blooio-go/internal/apiquery"
	"github.com/stainless-sdks/blooio-go/internal/requestconfig"
	"github.com/stainless-sdks/blooio-go/option"
	"github.com/stainless-sdks/blooio-go/packages/param"
	"github.com/stainless-sdks/blooio-go/packages/respjson"
)

// Phone number validation, formatting, and NANPA geocoding. Requires an Enterprise
// plan (Dedicated Enterprise).
//
// PhoneNumberLookupService contains methods and other services that help with
// interacting with the blooio API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPhoneNumberLookupService] method instead.
type PhoneNumberLookupService struct {
	options []option.RequestOption
}

// NewPhoneNumberLookupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPhoneNumberLookupService(opts ...option.RequestOption) (r PhoneNumberLookupService) {
	r = PhoneNumberLookupService{}
	r.options = opts
	return
}

// Same as the GET endpoint, but accepts the phone number in the request body.
// Useful when the number contains characters that are difficult to URL-encode.
//
// **Requires an Enterprise plan** (Dedicated Enterprise).
func (r *PhoneNumberLookupService) New(ctx context.Context, body PhoneNumberLookupNewParams, opts ...option.RequestOption) (res *PhoneNumberLookupResult, err error) {
	opts = slices.Concat(r.options, opts)
	path := "phone-numbers/lookup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns detailed information about a phone number including validation,
// formatting (E.164, national, international), number type, and NANPA geocoding
// (city, state/province) for North American numbers. The geocoding data is sourced
// from different database with 240,000+ NPA-NXX entries.
//
// **Requires an Enterprise plan** (Dedicated Enterprise). Returns 403 if your
// organization does not have an active enterprise subscription.
func (r *PhoneNumberLookupService) Get(ctx context.Context, query PhoneNumberLookupGetParams, opts ...option.RequestOption) (res *PhoneNumberLookupResult, err error) {
	opts = slices.Concat(r.options, opts)
	path := "phone-numbers/lookup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type PhoneNumberLookupResult struct {
	// NPA area code (first 3 digits of national number, only for NANP numbers)
	AreaCode string `json:"area_code"`
	// General region for the area code (most common city, only for NANP numbers)
	AreaCodeRegion string `json:"area_code_region"`
	// ISO 3166-1 alpha-2 country code
	Country string `json:"country" api:"nullable"`
	// Country calling code without +
	CountryCallingCode string `json:"country_calling_code"`
	// E.164 formatted number
	E164 string `json:"e164"`
	// NXX exchange code (digits 4-6 of national number, only for NANP numbers)
	Exchange string `json:"exchange"`
	// The original input string
	Input string `json:"input"`
	// International formatted number
	International string `json:"international"`
	// NANPA geocoding location (only for North American numbers with country code 1)
	Location PhoneNumberLookupResultLocation `json:"location" api:"nullable"`
	// National formatted number
	National string `json:"national"`
	// National number without country code
	NationalNumber string `json:"national_number"`
	// Whether the phone number is a possible number (less strict than valid)
	Possible bool `json:"possible"`
	// Number type detected by libphonenumber
	//
	// Any of "FIXED_LINE", "MOBILE", "FIXED_LINE_OR_MOBILE", "TOLL_FREE",
	// "PREMIUM_RATE", "SHARED_COST", "VOIP", "PERSONAL_NUMBER", "PAGER", "UAN",
	// "VOICEMAIL".
	Type PhoneNumberLookupResultType `json:"type" api:"nullable"`
	// Whether the phone number is valid
	Valid bool `json:"valid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AreaCode           respjson.Field
		AreaCodeRegion     respjson.Field
		Country            respjson.Field
		CountryCallingCode respjson.Field
		E164               respjson.Field
		Exchange           respjson.Field
		Input              respjson.Field
		International      respjson.Field
		Location           respjson.Field
		National           respjson.Field
		NationalNumber     respjson.Field
		Possible           respjson.Field
		Type               respjson.Field
		Valid              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberLookupResult) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberLookupResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// NANPA geocoding location (only for North American numbers with country code 1)
type PhoneNumberLookupResultLocation struct {
	// City name
	City string `json:"city" api:"nullable"`
	// State/province abbreviation
	Region string `json:"region" api:"nullable"`
	// Full state/province name
	RegionName string `json:"region_name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Region      respjson.Field
		RegionName  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberLookupResultLocation) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberLookupResultLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Number type detected by libphonenumber
type PhoneNumberLookupResultType string

const (
	PhoneNumberLookupResultTypeFixedLine         PhoneNumberLookupResultType = "FIXED_LINE"
	PhoneNumberLookupResultTypeMobile            PhoneNumberLookupResultType = "MOBILE"
	PhoneNumberLookupResultTypeFixedLineOrMobile PhoneNumberLookupResultType = "FIXED_LINE_OR_MOBILE"
	PhoneNumberLookupResultTypeTollFree          PhoneNumberLookupResultType = "TOLL_FREE"
	PhoneNumberLookupResultTypePremiumRate       PhoneNumberLookupResultType = "PREMIUM_RATE"
	PhoneNumberLookupResultTypeSharedCost        PhoneNumberLookupResultType = "SHARED_COST"
	PhoneNumberLookupResultTypeVoip              PhoneNumberLookupResultType = "VOIP"
	PhoneNumberLookupResultTypePersonalNumber    PhoneNumberLookupResultType = "PERSONAL_NUMBER"
	PhoneNumberLookupResultTypePager             PhoneNumberLookupResultType = "PAGER"
	PhoneNumberLookupResultTypeUan               PhoneNumberLookupResultType = "UAN"
	PhoneNumberLookupResultTypeVoicemail         PhoneNumberLookupResultType = "VOICEMAIL"
)

type PhoneNumberLookupNewParams struct {
	// Phone number to look up
	Number string `json:"number" api:"required"`
	paramObj
}

func (r PhoneNumberLookupNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PhoneNumberLookupNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PhoneNumberLookupNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberLookupGetParams struct {
	// Phone number to look up. Can be E.164 format (+12125551234), national format
	// (2125551234), or with formatting ((212) 555-1234).
	Number string `query:"number" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumberLookupGetParams]'s query parameters as
// `url.Values`.
func (r PhoneNumberLookupGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
