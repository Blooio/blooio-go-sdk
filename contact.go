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

// Manage contacts (phone numbers and emails)
//
// ContactService contains methods and other services that help with interacting
// with the blooio API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContactService] method instead.
type ContactService struct {
	options []option.RequestOption
	// Manage contacts (phone numbers and emails)
	Tags ContactTagService
}

// NewContactService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewContactService(opts ...option.RequestOption) (r ContactService) {
	r = ContactService{}
	r.options = opts
	r.Tags = NewContactTagService(opts...)
	return
}

// Create a new contact with a phone number (E.164 format) or email address.
func (r *ContactService) New(ctx context.Context, body ContactNewParams, opts ...option.RequestOption) (res *Contact, err error) {
	opts = slices.Concat(r.options, opts)
	path := "contacts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get details for a specific contact by phone number or email.
func (r *ContactService) Get(ctx context.Context, contactID string, opts ...option.RequestOption) (res *Contact, err error) {
	opts = slices.Concat(r.options, opts)
	if contactID == "" {
		err = errors.New("missing required contactId parameter")
		return nil, err
	}
	path := fmt.Sprintf("contacts/%s", url.PathEscape(contactID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a contact's name.
func (r *ContactService) Update(ctx context.Context, contactID string, body ContactUpdateParams, opts ...option.RequestOption) (res *Contact, err error) {
	opts = slices.Concat(r.options, opts)
	if contactID == "" {
		err = errors.New("missing required contactId parameter")
		return nil, err
	}
	path := fmt.Sprintf("contacts/%s", url.PathEscape(contactID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List all contacts for the organization with optional search and pagination.
func (r *ContactService) List(ctx context.Context, query ContactListParams, opts ...option.RequestOption) (res *ContactListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "contacts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Soft-delete a contact.
func (r *ContactService) Delete(ctx context.Context, contactID string, opts ...option.RequestOption) (res *DeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if contactID == "" {
		err = errors.New("missing required contactId parameter")
		return nil, err
	}
	path := fmt.Sprintf("contacts/%s", url.PathEscape(contactID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Check if a contact supports iMessage and/or SMS.
func (r *ContactService) CheckCapabilities(ctx context.Context, contactID string, opts ...option.RequestOption) (res *ContactCheckCapabilitiesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if contactID == "" {
		err = errors.New("missing required contactId parameter")
		return nil, err
	}
	path := fmt.Sprintf("contacts/%s/capabilities", url.PathEscape(contactID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type Contact struct {
	// Contact identifier (phone or email)
	ID string `json:"id"`
	// Internal contact ID
	ContactID string `json:"contact_id"`
	CreatedAt int64  `json:"created_at"`
	// Phone number (E.164) or email
	Identifier      string   `json:"identifier"`
	LastMessageTime int64    `json:"last_message_time" api:"nullable"`
	Name            string   `json:"name" api:"nullable"`
	Tags            []string `json:"tags"`
	// Any of "phone", "email".
	Type ContactType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ContactID       respjson.Field
		CreatedAt       respjson.Field
		Identifier      respjson.Field
		LastMessageTime respjson.Field
		Name            respjson.Field
		Tags            respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Contact) RawJSON() string { return r.JSON.raw }
func (r *Contact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactType string

const (
	ContactTypePhone ContactType = "phone"
	ContactTypeEmail ContactType = "email"
)

type DeleteResponse struct {
	DeletedAt int64 `json:"deleted_at"`
	Success   bool  `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeletedAt   respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *DeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Pagination struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
	Total  int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Offset      respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Pagination) RawJSON() string { return r.JSON.raw }
func (r *Pagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactListResponse struct {
	Contacts   []Contact  `json:"contacts"`
	Pagination Pagination `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contacts    respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactListResponse) RawJSON() string { return r.JSON.raw }
func (r *ContactListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactCheckCapabilitiesResponse struct {
	Capabilities ContactCheckCapabilitiesResponseCapabilities `json:"capabilities"`
	// Normalized contact identifier
	Contact string `json:"contact"`
	// Timestamp when capabilities were checked
	LastChecked int64 `json:"last_checked"`
	// Any of "phone", "email".
	Type ContactCheckCapabilitiesResponseType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Capabilities respjson.Field
		Contact      respjson.Field
		LastChecked  respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactCheckCapabilitiesResponse) RawJSON() string { return r.JSON.raw }
func (r *ContactCheckCapabilitiesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactCheckCapabilitiesResponseCapabilities struct {
	// Whether FaceTime is available
	Facetime bool `json:"facetime"`
	// Whether iMessage is available
	Imessage bool `json:"imessage"`
	// Whether SMS is available (phone only)
	SMS bool `json:"sms"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Facetime    respjson.Field
		Imessage    respjson.Field
		SMS         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactCheckCapabilitiesResponseCapabilities) RawJSON() string { return r.JSON.raw }
func (r *ContactCheckCapabilitiesResponseCapabilities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactCheckCapabilitiesResponseType string

const (
	ContactCheckCapabilitiesResponseTypePhone ContactCheckCapabilitiesResponseType = "phone"
	ContactCheckCapabilitiesResponseTypeEmail ContactCheckCapabilitiesResponseType = "email"
)

type ContactNewParams struct {
	// Phone number (E.164 format, e.g., +15551234567) or email address
	Identifier string `json:"identifier" api:"required"`
	// Display name for the contact
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r ContactNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ContactNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContactNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactUpdateParams struct {
	// New display name (null to clear)
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r ContactUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ContactUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContactUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactListParams struct {
	// Maximum number of items to return in a single response. Must be between 1 and
	// 200; defaults to 50. Use together with `offset` to page through large result
	// sets.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip before returning results. Combine with `limit` for
	// page-based pagination (e.g. `offset=50&limit=50` returns the second page).
	// Defaults to 0.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Search query (matches identifier or name)
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Sort order
	//
	// Any of "recent", "oldest", "name_asc", "name_desc".
	Sort ContactListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ContactListParams]'s query parameters as `url.Values`.
func (r ContactListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order
type ContactListParamsSort string

const (
	ContactListParamsSortRecent   ContactListParamsSort = "recent"
	ContactListParamsSortOldest   ContactListParamsSort = "oldest"
	ContactListParamsSortNameAsc  ContactListParamsSort = "name_asc"
	ContactListParamsSortNameDesc ContactListParamsSort = "name_desc"
)
