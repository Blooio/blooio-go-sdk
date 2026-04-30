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

// Manage contact groups
//
// GroupService contains methods and other services that help with interacting with
// the blooio API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGroupService] method instead.
type GroupService struct {
	options []option.RequestOption
	// Manage group membership
	Members GroupMemberService
	// Manage contact groups
	Icon GroupIconService
}

// NewGroupService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewGroupService(opts ...option.RequestOption) (r GroupService) {
	r = GroupService{}
	r.options = opts
	r.Members = NewGroupMemberService(opts...)
	r.Icon = NewGroupIconService(opts...)
	return
}

// Create a new group. There are two modes:
//
// **1. Link to existing iMessage chat:** Provide `chat_guid` to join an existing
// group chat that was created outside the API. The `members` list records who is
// in the group but does NOT add them to the linked iMessage chat. Multiple groups
// can have the same participants if they have different `chat_guid`s.
//
// **2. Create new group:** Omit `chat_guid` to create a new group. When you send
// the first message, a new iMessage chat will be created. Note: iMessage only
// allows one chat per unique participant set when created via API.
func (r *GroupService) New(ctx context.Context, body GroupNewParams, opts ...option.RequestOption) (res *GroupNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get details for a specific group.
func (r *GroupService) Get(ctx context.Context, groupID string, opts ...option.RequestOption) (res *Group, err error) {
	opts = slices.Concat(r.options, opts)
	if groupID == "" {
		err = errors.New("missing required groupId parameter")
		return nil, err
	}
	path := fmt.Sprintf("groups/%s", url.PathEscape(groupID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a group's name. If the group has a linked `chat_guid`, the display name
// will also be updated in the linked iMessage chat. Note: iMessage only allows one
// chat per unique participant set, so renaming simply changes the display name on
// the existing chat thread.
func (r *GroupService) Update(ctx context.Context, groupID string, body GroupUpdateParams, opts ...option.RequestOption) (res *GroupUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if groupID == "" {
		err = errors.New("missing required groupId parameter")
		return nil, err
	}
	path := fmt.Sprintf("groups/%s", url.PathEscape(groupID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List all groups for the organization with optional search and pagination.
func (r *GroupService) List(ctx context.Context, query GroupListParams, opts ...option.RequestOption) (res *GroupListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Soft-delete a group. Members are automatically removed. If the group is linked
// to an existing iMessage chat, the number also leaves that chat.
func (r *GroupService) Delete(ctx context.Context, groupID string, opts ...option.RequestOption) (res *GroupDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if groupID == "" {
		err = errors.New("missing required groupId parameter")
		return nil, err
	}
	path := fmt.Sprintf("groups/%s", url.PathEscape(groupID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type Group struct {
	// BlueBubbles chat GUID if linked to a device group chat
	ChatGuid  string `json:"chat_guid" api:"nullable"`
	CreatedAt int64  `json:"created_at"`
	GroupID   string `json:"group_id"`
	// URL of the group icon/photo
	IconURL string `json:"icon_url" api:"nullable"`
	// Direction of the most recent message
	//
	// Any of "inbound", "outbound".
	LastMessageDirection GroupLastMessageDirection `json:"last_message_direction" api:"nullable"`
	// Text of the most recent message in the group
	LastMessageText string `json:"last_message_text" api:"nullable"`
	// Timestamp of the most recent message
	LastMessageTime int64 `json:"last_message_time" api:"nullable"`
	MemberCount     int64 `json:"member_count"`
	// Total number of messages in this group
	MessageCount int64 `json:"message_count"`
	// Group name. Null for unnamed groups.
	Name string `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChatGuid             respjson.Field
		CreatedAt            respjson.Field
		GroupID              respjson.Field
		IconURL              respjson.Field
		LastMessageDirection respjson.Field
		LastMessageText      respjson.Field
		LastMessageTime      respjson.Field
		MemberCount          respjson.Field
		MessageCount         respjson.Field
		Name                 respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Group) RawJSON() string { return r.JSON.raw }
func (r *Group) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Direction of the most recent message
type GroupLastMessageDirection string

const (
	GroupLastMessageDirectionInbound  GroupLastMessageDirection = "inbound"
	GroupLastMessageDirectionOutbound GroupLastMessageDirection = "outbound"
)

type GroupNewResponse struct {
	// List of member identifiers that were added to the group
	AddedMembers []string `json:"added_members"`
	// List of contacts that were auto-created
	CreatedContacts []string `json:"created_contacts"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddedMembers    respjson.Field
		CreatedContacts respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
	Group
}

// Returns the unmodified JSON received from the API
func (r GroupNewResponse) RawJSON() string { return r.JSON.raw }
func (r *GroupNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GroupUpdateResponse struct {
	// Result of syncing the operation to a linked iMessage chat
	DeviceSync GroupUpdateResponseDeviceSync `json:"device_sync"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeviceSync  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Group
}

// Returns the unmodified JSON received from the API
func (r GroupUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *GroupUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of syncing the operation to a linked iMessage chat
type GroupUpdateResponseDeviceSync struct {
	// The action that was performed for the linked chat
	//
	// Any of "add_participant", "remove_participant", "leave".
	Action string `json:"action"`
	// The linked iMessage chat GUID
	ChatGuid string `json:"chat_guid"`
	// Error message if sync failed
	Error string `json:"error" api:"nullable"`
	// Whether the sync was successful
	Synced bool `json:"synced"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		ChatGuid    respjson.Field
		Error       respjson.Field
		Synced      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GroupUpdateResponseDeviceSync) RawJSON() string { return r.JSON.raw }
func (r *GroupUpdateResponseDeviceSync) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GroupListResponse struct {
	Groups     []Group    `json:"groups"`
	Pagination Pagination `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Groups      respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GroupListResponse) RawJSON() string { return r.JSON.raw }
func (r *GroupListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GroupDeleteResponse struct {
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
func (r GroupDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *GroupDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GroupNewParams struct {
	// Group name (max 255 characters)
	Name string `json:"name" api:"required"`
	// BlueBubbles chat GUID to link this group to an existing iMessage chat. Use this
	// to join groups created elsewhere. You can get this from the BlueBubbles API or
	// from inbound message webhooks.
	ChatGuid param.Opt[string] `json:"chat_guid,omitzero"`
	// Phone numbers or emails of contacts in the group. When linking via chat_guid,
	// this is for record-keeping only (members are not added to the linked iMessage
	// chat).
	Members []string `json:"members,omitzero"`
	paramObj
}

func (r GroupNewParams) MarshalJSON() (data []byte, err error) {
	type shadow GroupNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GroupNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GroupUpdateParams struct {
	// New group name
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r GroupUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow GroupUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GroupUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GroupListParams struct {
	// Maximum number of items to return (1-200)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Search query (matches group name)
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Sort order
	//
	// Any of "recent", "oldest", "name_asc", "name_desc".
	Sort GroupListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GroupListParams]'s query parameters as `url.Values`.
func (r GroupListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order
type GroupListParamsSort string

const (
	GroupListParamsSortRecent   GroupListParamsSort = "recent"
	GroupListParamsSortOldest   GroupListParamsSort = "oldest"
	GroupListParamsSortNameAsc  GroupListParamsSort = "name_asc"
	GroupListParamsSortNameDesc GroupListParamsSort = "name_desc"
)
