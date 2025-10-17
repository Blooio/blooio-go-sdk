// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blooio

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/blooio-go/internal/apijson"
	"github.com/stainless-sdks/blooio-go/internal/requestconfig"
	"github.com/stainless-sdks/blooio-go/option"
	"github.com/stainless-sdks/blooio-go/packages/respjson"
)

// ContactService contains methods and other services that help with interacting
// with the blooio API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContactService] method instead.
type ContactService struct {
	Options []option.RequestOption
}

// NewContactService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewContactService(opts ...option.RequestOption) (r ContactService) {
	r = ContactService{}
	r.Options = opts
	return
}

// Check if a phone number or email address supports iMessage, SMS, RCS, and other
// messaging capabilities.
func (r *ContactService) CheckCapabilities(ctx context.Context, contact string, opts ...option.RequestOption) (res *ContactCheckCapabilitiesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if contact == "" {
		err = errors.New("missing required contact parameter")
		return
	}
	path := fmt.Sprintf("v1/api/contacts/%s/capabilities", contact)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ContactCheckCapabilitiesResponse struct {
	// Messaging capabilities for this contact.
	Capabilities ContactCheckCapabilitiesResponseCapabilities `json:"capabilities"`
	// The contact identifier (phone number or email).
	Contact string `json:"contact"`
	// ISO 8601 timestamp of when capabilities were last checked.
	LastChecked time.Time `json:"lastChecked" format:"date-time"`
	// Type of contact identifier.
	//
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

// Messaging capabilities for this contact.
type ContactCheckCapabilitiesResponseCapabilities struct {
	// Whether this contact supports iMessage.
	Imessage bool `json:"imessage"`
	// Whether this contact supports SMS (always true for phone numbers, false for
	// emails).
	SMS bool `json:"sms"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
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

// Type of contact identifier.
type ContactCheckCapabilitiesResponseType string

const (
	ContactCheckCapabilitiesResponseTypePhone ContactCheckCapabilitiesResponseType = "phone"
	ContactCheckCapabilitiesResponseTypeEmail ContactCheckCapabilitiesResponseType = "email"
)
