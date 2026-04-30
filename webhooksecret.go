// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blooio

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/blooio-go/internal/apijson"
	"github.com/stainless-sdks/blooio-go/internal/requestconfig"
	"github.com/stainless-sdks/blooio-go/option"
	"github.com/stainless-sdks/blooio-go/packages/respjson"
)

// Manage webhook subscriptions
//
// WebhookSecretService contains methods and other services that help with
// interacting with the blooio API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookSecretService] method instead.
type WebhookSecretService struct {
	options []option.RequestOption
}

// NewWebhookSecretService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWebhookSecretService(opts ...option.RequestOption) (r WebhookSecretService) {
	r = WebhookSecretService{}
	r.options = opts
	return
}

// Generate a new signing secret for the webhook. The new secret is returned only
// once in this response - store it securely. The old secret becomes invalid
// immediately.
func (r *WebhookSecretService) Rotate(ctx context.Context, webhookID string, opts ...option.RequestOption) (res *WebhookSecretRotateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks/%s/secret/rotate", url.PathEscape(webhookID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type WebhookSecretRotateResponse struct {
	// Timestamp when the secret was rotated
	RotatedAt int64 `json:"rotated_at"`
	// Identifier of who rotated the secret
	RotatedBy string `json:"rotated_by"`
	// Total number of times this secret has been rotated
	RotationCount int64 `json:"rotation_count"`
	// The new signing secret. Store this securely - it will not be shown again.
	SigningSecret string `json:"signing_secret"`
	WebhookID     string `json:"webhook_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RotatedAt     respjson.Field
		RotatedBy     respjson.Field
		RotationCount respjson.Field
		SigningSecret respjson.Field
		WebhookID     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookSecretRotateResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookSecretRotateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
