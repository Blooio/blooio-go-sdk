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
	"github.com/Blooio/blooio-go-sdk/packages/respjson"
)

// Control typing indicators for conversations
//
// ChatTypingService contains methods and other services that help with interacting
// with the blooio API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChatTypingService] method instead.
type ChatTypingService struct {
	options []option.RequestOption
}

// NewChatTypingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewChatTypingService(opts ...option.RequestOption) (r ChatTypingService) {
	r = ChatTypingService{}
	r.options = opts
	return
}

// Start the typing indicator for a chat. The indicator shows the recipient that
// you are typing.
func (r *ChatTypingService) Start(ctx context.Context, chatID string, opts ...option.RequestOption) (res *TypingResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if chatID == "" {
		err = errors.New("missing required chatId parameter")
		return nil, err
	}
	path := fmt.Sprintf("chats/%s/typing", url.PathEscape(chatID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Stop the typing indicator for a chat.
func (r *ChatTypingService) Stop(ctx context.Context, chatID string, opts ...option.RequestOption) (res *TypingResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if chatID == "" {
		err = errors.New("missing required chatId parameter")
		return nil, err
	}
	path := fmt.Sprintf("chats/%s/typing", url.PathEscape(chatID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type TypingResponse struct {
	// Chat identifier
	ChatID string `json:"chat_id"`
	// Timestamp when typing started (only for start)
	StartedAt int64 `json:"started_at"`
	// Timestamp when typing stopped (only for stop)
	StoppedAt int64 `json:"stopped_at"`
	// Whether typing indicator is active
	Typing bool `json:"typing"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChatID      respjson.Field
		StartedAt   respjson.Field
		StoppedAt   respjson.Field
		Typing      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TypingResponse) RawJSON() string { return r.JSON.raw }
func (r *TypingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
