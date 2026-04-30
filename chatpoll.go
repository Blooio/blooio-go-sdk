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

// Send native iMessage polls and retrieve poll results with vote counts. Poll
// events are delivered via separate webhook event types (poll.received,
// poll.created, poll.voted) and require webhook_type 'poll' or 'all'.
//
// ChatPollService contains methods and other services that help with interacting
// with the blooio API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChatPollService] method instead.
type ChatPollService struct {
	options []option.RequestOption
}

// NewChatPollService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewChatPollService(opts ...option.RequestOption) (r ChatPollService) {
	r = ChatPollService{}
	r.options = opts
	return
}

// Retrieve a poll's definition and aggregated vote counts. The pollId is the
// poll_id returned in the poll.received or poll.created webhook event.
func (r *ChatPollService) GetResults(ctx context.Context, pollID string, query ChatPollGetResultsParams, opts ...option.RequestOption) (res *ChatPollGetResultsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if query.ChatID == "" {
		err = errors.New("missing required chatId parameter")
		return nil, err
	}
	if pollID == "" {
		err = errors.New("missing required pollId parameter")
		return nil, err
	}
	path := fmt.Sprintf("chats/%s/polls/%s", url.PathEscape(query.ChatID), url.PathEscape(pollID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Send a native iMessage poll to a chat. The poll appears as an interactive ballot
// that recipients can vote on.
func (r *ChatPollService) Send(ctx context.Context, chatID string, body ChatPollSendParams, opts ...option.RequestOption) (res *ChatPollSendResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if chatID == "" {
		err = errors.New("missing required chatId parameter")
		return nil, err
	}
	path := fmt.Sprintf("chats/%s/polls", url.PathEscape(chatID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ChatPollGetResultsResponse struct {
	ChatID     string                             `json:"chat_id"`
	Options    []ChatPollGetResultsResponseOption `json:"options"`
	PollID     string                             `json:"poll_id"`
	Title      string                             `json:"title"`
	TotalVotes int64                              `json:"total_votes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChatID      respjson.Field
		Options     respjson.Field
		PollID      respjson.Field
		Title       respjson.Field
		TotalVotes  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatPollGetResultsResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatPollGetResultsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatPollGetResultsResponseOption struct {
	Text  string `json:"text"`
	Votes int64  `json:"votes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Votes       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatPollGetResultsResponseOption) RawJSON() string { return r.JSON.raw }
func (r *ChatPollGetResultsResponseOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatPollSendResponse struct {
	ChatID string                   `json:"chat_id"`
	Poll   ChatPollSendResponsePoll `json:"poll"`
	// Unique identifier for the poll
	PollID string  `json:"poll_id"`
	SentAt float64 `json:"sent_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChatID      respjson.Field
		Poll        respjson.Field
		PollID      respjson.Field
		SentAt      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatPollSendResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatPollSendResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatPollSendResponsePoll struct {
	Options []string `json:"options"`
	Title   string   `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Options     respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatPollSendResponsePoll) RawJSON() string { return r.JSON.raw }
func (r *ChatPollSendResponsePoll) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatPollGetResultsParams struct {
	ChatID string `path:"chatId" api:"required" json:"-"`
	paramObj
}

type ChatPollSendParams struct {
	// Array of 2-10 option strings for the poll
	Options []string `json:"options,omitzero" api:"required"`
	// Poll question or title (optional)
	Title param.Opt[string] `json:"title,omitzero"`
	paramObj
}

func (r ChatPollSendParams) MarshalJSON() (data []byte, err error) {
	type shadow ChatPollSendParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatPollSendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
