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

// Manage contacts (phone numbers and emails)
//
// ContactTagService contains methods and other services that help with interacting
// with the blooio API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContactTagService] method instead.
type ContactTagService struct {
	options []option.RequestOption
}

// NewContactTagService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewContactTagService(opts ...option.RequestOption) (r ContactTagService) {
	r = ContactTagService{}
	r.options = opts
	return
}

// List all tags assigned to a contact.
func (r *ContactTagService) List(ctx context.Context, contactID string, opts ...option.RequestOption) (res *ContactTagListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if contactID == "" {
		err = errors.New("missing required contactId parameter")
		return nil, err
	}
	path := fmt.Sprintf("contacts/%s/tags", url.PathEscape(contactID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Add one or more tags to a contact. If a tag already exists on the contact, it is
// re-activated (idempotent). Tags are free-form strings.
func (r *ContactTagService) Add(ctx context.Context, contactID string, body ContactTagAddParams, opts ...option.RequestOption) (res *ContactTagAddResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if contactID == "" {
		err = errors.New("missing required contactId parameter")
		return nil, err
	}
	path := fmt.Sprintf("contacts/%s/tags", url.PathEscape(contactID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Remove a specific tag from a contact. The tag is soft-deleted and can be
// re-added later.
func (r *ContactTagService) Remove(ctx context.Context, tag string, body ContactTagRemoveParams, opts ...option.RequestOption) (res *DeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.ContactID == "" {
		err = errors.New("missing required contactId parameter")
		return nil, err
	}
	if tag == "" {
		err = errors.New("missing required tag parameter")
		return nil, err
	}
	path := fmt.Sprintf("contacts/%s/tags/%s", url.PathEscape(body.ContactID), url.PathEscape(tag))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type ContactTagListResponse struct {
	Tags []ContactTagListResponseTag `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Tags        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactTagListResponse) RawJSON() string { return r.JSON.raw }
func (r *ContactTagListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactTagListResponseTag struct {
	// Timestamp when the tag was added (ms since epoch)
	CreatedAt int64 `json:"created_at"`
	// The tag value
	Tag string `json:"tag"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Tag         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactTagListResponseTag) RawJSON() string { return r.JSON.raw }
func (r *ContactTagListResponseTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactTagAddResponse struct {
	Success bool `json:"success"`
	// Tags that were added
	TagsAdded []string `json:"tags_added"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		TagsAdded   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactTagAddResponse) RawJSON() string { return r.JSON.raw }
func (r *ContactTagAddResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactTagAddParams struct {
	// Tags to add
	Tags []string `json:"tags,omitzero" api:"required"`
	paramObj
}

func (r ContactTagAddParams) MarshalJSON() (data []byte, err error) {
	type shadow ContactTagAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContactTagAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactTagRemoveParams struct {
	ContactID string `path:"contactId" api:"required" json:"-"`
	paramObj
}
