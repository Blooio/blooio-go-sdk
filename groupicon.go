// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blooio

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"

	"github.com/Blooio/blooio-go-sdk/internal/apiform"
	"github.com/Blooio/blooio-go-sdk/internal/apijson"
	"github.com/Blooio/blooio-go-sdk/internal/requestconfig"
	"github.com/Blooio/blooio-go-sdk/option"
	"github.com/Blooio/blooio-go-sdk/packages/respjson"
)

// Manage contact groups
//
// GroupIconService contains methods and other services that help with interacting
// with the blooio API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGroupIconService] method instead.
type GroupIconService struct {
	options []option.RequestOption
}

// NewGroupIconService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGroupIconService(opts ...option.RequestOption) (r GroupIconService) {
	r = GroupIconService{}
	r.options = opts
	return
}

// Remove the group icon/photo. Requires the group to have a linked chat_guid.
//
// The icon is removed from both Blooio storage and the linked iMessage chat before
// the request returns.
func (r *GroupIconService) Remove(ctx context.Context, groupID string, opts ...option.RequestOption) (res *GroupIcon, err error) {
	opts = slices.Concat(r.options, opts)
	if groupID == "" {
		err = errors.New("missing required groupId parameter")
		return nil, err
	}
	path := fmt.Sprintf("groups/%s/icon", url.PathEscape(groupID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Set the group icon/photo. Requires the group to have a linked chat_guid. Uses
// multipart/form-data.
//
// The uploaded image is stored in Blooio storage and synced to the linked iMessage
// chat before the request returns.
func (r *GroupIconService) Set(ctx context.Context, groupID string, body GroupIconSetParams, opts ...option.RequestOption) (res *GroupIcon, err error) {
	opts = slices.Concat(r.options, opts)
	if groupID == "" {
		err = errors.New("missing required groupId parameter")
		return nil, err
	}
	path := fmt.Sprintf("groups/%s/icon", url.PathEscape(groupID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Response for group icon operations
type GroupIcon struct {
	// The BlueBubbles chat GUID
	ChatGuid string `json:"chat_guid"`
	// Linked chat sync status
	DeviceSync GroupIconDeviceSync `json:"device_sync"`
	GroupID    string              `json:"group_id"`
	// URL of the uploaded icon (only present on set)
	IconURL string `json:"icon_url"`
	Message string `json:"message"`
	Success bool   `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChatGuid    respjson.Field
		DeviceSync  respjson.Field
		GroupID     respjson.Field
		IconURL     respjson.Field
		Message     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GroupIcon) RawJSON() string { return r.JSON.raw }
func (r *GroupIcon) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Linked chat sync status
type GroupIconDeviceSync struct {
	ChatGuid string `json:"chat_guid"`
	// Status message about linked chat sync
	Message string `json:"message"`
	// Whether the icon change was synced to the linked iMessage chat. This will be
	// true on successful set/remove operations.
	Synced bool `json:"synced"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChatGuid    respjson.Field
		Message     respjson.Field
		Synced      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GroupIconDeviceSync) RawJSON() string { return r.JSON.raw }
func (r *GroupIconDeviceSync) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GroupIconSetParams struct {
	// The icon image file to set as the group photo
	Icon io.Reader `json:"icon,omitzero" api:"required" format:"binary"`
	paramObj
}

func (r GroupIconSetParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}
