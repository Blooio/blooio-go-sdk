// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blooio

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/Blooio/blooio-go-sdk/internal/apijson"
	"github.com/Blooio/blooio-go-sdk/internal/requestconfig"
	"github.com/Blooio/blooio-go-sdk/option"
	"github.com/Blooio/blooio-go-sdk/packages/respjson"
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
	// Plan type the underlying allocation runs on. Sourced directly from
	// `allocation_pool.type` — the enum mirrors the DB `CHECK` constraint (see
	// migration 2026-05-09-inbound-plan.sql), so any value here is also a valid type
	// stored in the database. `inbound` numbers are reply-only — outbound to a
	// recipient (a contact for 1:1 chats, the group for group chats) requires that
	// recipient to have messaged the number first (otherwise the send returns
	// `403 inbound_only_no_prior_inbound`). `null` indicates the underlying allocation
	// predates the type column or is unattributed; clients should treat `null` the
	// same as `dedicated` for routing decisions.
	//
	// Any of "shared", "dedicated", "inbound", "trial", "2fa".
	PlanKind string `json:"plan_kind" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsActive    respjson.Field
		LastActive  respjson.Field
		PhoneNumber respjson.Field
		PlanKind    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeNumberListResponseNumber) RawJSON() string { return r.JSON.raw }
func (r *MeNumberListResponseNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
