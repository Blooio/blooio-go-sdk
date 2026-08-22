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
)

// Initiate FaceTime calls
//
// FacetimeService contains methods and other services that help with interacting
// with the blooio API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFacetimeService] method instead.
type FacetimeService struct {
	options []option.RequestOption
}

// NewFacetimeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFacetimeService(opts ...option.RequestOption) (r FacetimeService) {
	r = FacetimeService{}
	r.options = opts
	return
}

// **Coming Soon** -- This endpoint is temporarily disabled while we stabilize the
// FaceTime call flow.
//
// Initiates a FaceTime call to the specified phone number or email address.
// Returns a shareable FaceTime link that anyone can use to join the call. The call
// will ring the contact and auto-admit the first person who joins via the link.
func (r *FacetimeService) InitiateCall(ctx context.Context, body FacetimeInitiateCallParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "facetime/calls"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

type FacetimeInitiateCallParams struct {
	// Phone number (E.164) or email address to call
	Handle string `json:"handle" api:"required"`
	paramObj
}

func (r FacetimeInitiateCallParams) MarshalJSON() (data []byte, err error) {
	type shadow FacetimeInitiateCallParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FacetimeInitiateCallParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
