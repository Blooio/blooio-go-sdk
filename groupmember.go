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

// Manage group membership
//
// GroupMemberService contains methods and other services that help with
// interacting with the blooio API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGroupMemberService] method instead.
type GroupMemberService struct {
	options []option.RequestOption
}

// NewGroupMemberService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGroupMemberService(opts ...option.RequestOption) (r GroupMemberService) {
	r = GroupMemberService{}
	r.options = opts
	return
}

// List all members of a group.
func (r *GroupMemberService) List(ctx context.Context, groupID string, query GroupMemberListParams, opts ...option.RequestOption) (res *GroupMemberListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if groupID == "" {
		err = errors.New("missing required groupId parameter")
		return nil, err
	}
	path := fmt.Sprintf("groups/%s/members", url.PathEscape(groupID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// ⚠️ **COMING SOON** - This endpoint is temporarily disabled while we stabilize
// this feature.
//
// Add an existing contact to a group. If the group is linked to an existing
// iMessage chat, also adds the participant to that chat.
func (r *GroupMemberService) Add(ctx context.Context, groupID string, body GroupMemberAddParams, opts ...option.RequestOption) (res *GroupMemberAddResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if groupID == "" {
		err = errors.New("missing required groupId parameter")
		return nil, err
	}
	path := fmt.Sprintf("groups/%s/members", url.PathEscape(groupID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// ⚠️ **COMING SOON** - This endpoint is temporarily disabled while we stabilize
// this feature.
//
// Remove a contact from a group. If the group is linked to an existing iMessage
// chat, also removes the participant from that chat. If the contact being removed
// is the organization's own phone number, leaves the group chat instead.
func (r *GroupMemberService) Remove(ctx context.Context, contactID string, body GroupMemberRemoveParams, opts ...option.RequestOption) (res *GroupMemberRemoveResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.GroupID == "" {
		err = errors.New("missing required groupId parameter")
		return nil, err
	}
	if contactID == "" {
		err = errors.New("missing required contactId parameter")
		return nil, err
	}
	path := fmt.Sprintf("groups/%s/members/%s", url.PathEscape(body.GroupID), url.PathEscape(contactID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type GroupMember struct {
	// Contact identifier (phone or email)
	ID         string `json:"id"`
	AddedAt    int64  `json:"added_at"`
	ContactID  string `json:"contact_id"`
	Identifier string `json:"identifier"`
	Name       string `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AddedAt     respjson.Field
		ContactID   respjson.Field
		Identifier  respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GroupMember) RawJSON() string { return r.JSON.raw }
func (r *GroupMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GroupMemberListResponse struct {
	// The group ID
	GroupID string `json:"group_id"`
	// The group name
	GroupName string `json:"group_name" api:"nullable"`
	// URL of the group icon/photo
	IconURL    string        `json:"icon_url" api:"nullable"`
	Members    []GroupMember `json:"members"`
	Pagination Pagination    `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GroupID     respjson.Field
		GroupName   respjson.Field
		IconURL     respjson.Field
		Members     respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GroupMemberListResponse) RawJSON() string { return r.JSON.raw }
func (r *GroupMemberListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GroupMemberAddResponse struct {
	Member  GroupMember `json:"member"`
	Message string      `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Member      respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GroupMemberAddResponse) RawJSON() string { return r.JSON.raw }
func (r *GroupMemberAddResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GroupMemberRemoveResponse struct {
	RemovedAt int64 `json:"removed_at"`
	Success   bool  `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RemovedAt   respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GroupMemberRemoveResponse) RawJSON() string { return r.JSON.raw }
func (r *GroupMemberRemoveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GroupMemberListParams struct {
	// Maximum number of items to return in a single response. Must be between 1 and
	// 200; defaults to 50. Use together with `offset` to page through large result
	// sets.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip before returning results. Combine with `limit` for
	// page-based pagination (e.g. `offset=50&limit=50` returns the second page).
	// Defaults to 0.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GroupMemberListParams]'s query parameters as `url.Values`.
func (r GroupMemberListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GroupMemberAddParams struct {
	// Contact identifier (phone number or email)
	ContactID string `json:"contact_id" api:"required"`
	paramObj
}

func (r GroupMemberAddParams) MarshalJSON() (data []byte, err error) {
	type shadow GroupMemberAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GroupMemberAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GroupMemberRemoveParams struct {
	GroupID string `path:"groupId" api:"required" json:"-"`
	paramObj
}
