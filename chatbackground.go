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

	"github.com/stainless-sdks/blooio-go/internal/apiform"
	"github.com/stainless-sdks/blooio-go/internal/apijson"
	"github.com/stainless-sdks/blooio-go/internal/requestconfig"
	"github.com/stainless-sdks/blooio-go/option"
	"github.com/stainless-sdks/blooio-go/packages/respjson"
)

// Set, get, and remove conversation backgrounds
//
// ChatBackgroundService contains methods and other services that help with
// interacting with the blooio API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChatBackgroundService] method instead.
type ChatBackgroundService struct {
	options []option.RequestOption
}

// NewChatBackgroundService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewChatBackgroundService(opts ...option.RequestOption) (r ChatBackgroundService) {
	r = ChatBackgroundService{}
	r.options = opts
	return
}

// Get the current background image metadata for a conversation. Works for both
// 1-on-1 and group chats.
func (r *ChatBackgroundService) Get(ctx context.Context, chatID string, opts ...option.RequestOption) (res *ChatBackgroundResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if chatID == "" {
		err = errors.New("missing required chatId parameter")
		return nil, err
	}
	path := fmt.Sprintf("chats/%s/background", url.PathEscape(chatID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Remove the background image from a conversation, reverting to the default
// appearance.
func (r *ChatBackgroundService) Remove(ctx context.Context, chatID string, opts ...option.RequestOption) (res *ChatBackgroundResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if chatID == "" {
		err = errors.New("missing required chatId parameter")
		return nil, err
	}
	path := fmt.Sprintf("chats/%s/background", url.PathEscape(chatID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Set or update the background image for a conversation. Works for both 1-on-1 and
// group chats.
//
// The uploaded image is converted into a PosterKit-compatible archive and applied
// to the iMessage conversation on the linked device. Supported formats: JPEG, PNG,
// GIF, WebP, HEIC/HEIF. Maximum file size: 10 MB.
func (r *ChatBackgroundService) Set(ctx context.Context, chatID string, body ChatBackgroundSetParams, opts ...option.RequestOption) (res *ChatBackgroundResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if chatID == "" {
		err = errors.New("missing required chatId parameter")
		return nil, err
	}
	path := fmt.Sprintf("chats/%s/background", url.PathEscape(chatID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Response for chat background operations
type ChatBackgroundResponse struct {
	// Unique identifier for the current background, or null if none
	BackgroundID string `json:"background_id" api:"nullable"`
	// Version number of the background (for cache invalidation)
	BackgroundVersion int64 `json:"background_version" api:"nullable"`
	// Whether the background was changed by this operation (only present on PUT)
	Changed bool `json:"changed"`
	// Normalized chat identifier (phone number, email, or group ID)
	ChatID string `json:"chat_id"`
	// Whether the chat currently has a background set
	HasBackground bool `json:"has_background"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BackgroundID      respjson.Field
		BackgroundVersion respjson.Field
		Changed           respjson.Field
		ChatID            respjson.Field
		HasBackground     respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatBackgroundResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatBackgroundResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatBackgroundSetParams struct {
	// The image file to set as the chat background
	Background io.Reader `json:"background,omitzero" api:"required" format:"binary"`
	paramObj
}

func (r ChatBackgroundSetParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
