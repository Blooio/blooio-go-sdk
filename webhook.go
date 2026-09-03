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

// Manage webhook subscriptions
//
// WebhookService contains methods and other services that help with interacting
// with the blooio API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	options []option.RequestOption
	// Manage webhook subscriptions
	Secret WebhookSecretService
	// View and replay webhook deliveries
	Logs WebhookLogService
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r WebhookService) {
	r = WebhookService{}
	r.options = opts
	r.Secret = NewWebhookSecretService(opts...)
	r.Logs = NewWebhookLogService(opts...)
	return
}

// Registration through this endpoint is closed and returns 410. Use POST
// /v4/webhooks to create new subscriptions. Existing webhooks keep working and can
// still be listed, updated, and deleted here. Re-posting the URL of a webhook that
// already exists still returns 200 with that webhook, so idempotent provisioning
// scripts continue to work unchanged.
//
// Deprecated: deprecated
func (r *WebhookService) New(ctx context.Context, body WebhookNewParams, opts ...option.RequestOption) (res *WebhookNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get details for a specific webhook.
func (r *WebhookService) Get(ctx context.Context, webhookID string, opts ...option.RequestOption) (res *Webhook, err error) {
	opts = slices.Concat(r.options, opts)
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks/%s", url.PathEscape(webhookID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a webhook's configuration.
func (r *WebhookService) Update(ctx context.Context, webhookID string, body WebhookUpdateParams, opts ...option.RequestOption) (res *Webhook, err error) {
	opts = slices.Concat(r.options, opts)
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks/%s", url.PathEscape(webhookID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List all webhooks for the organization.
func (r *WebhookService) List(ctx context.Context, opts ...option.RequestOption) (res *WebhookListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Permanently delete a webhook.
func (r *WebhookService) Delete(ctx context.Context, webhookID string, opts ...option.RequestOption) (res *WebhookDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks/%s", url.PathEscape(webhookID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type Webhook struct {
	// Name of the API key (if scope is api_key)
	APIKeyName   string `json:"api_key_name" api:"nullable"`
	CreatedAt    int64  `json:"created_at"`
	DeprecatedAt int64  `json:"deprecated_at" api:"nullable"`
	FailureCount int64  `json:"failure_count"`
	// Name of the integration (if scope is integration)
	IntegrationName string `json:"integration_name" api:"nullable"`
	// Whether the webhook is active (not deprecated)
	IsActive      bool  `json:"is_active"`
	LastTriggered int64 `json:"last_triggered" api:"nullable"`
	// Any of "api_key", "organization", "integration".
	Scope WebhookScope `json:"scope"`
	// -1 means no expiration
	ValidUntil int64  `json:"valid_until"`
	WebhookID  string `json:"webhook_id"`
	// Any of "message", "status", "all".
	WebhookType WebhookWebhookType `json:"webhook_type"`
	WebhookURL  string             `json:"webhook_url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKeyName      respjson.Field
		CreatedAt       respjson.Field
		DeprecatedAt    respjson.Field
		FailureCount    respjson.Field
		IntegrationName respjson.Field
		IsActive        respjson.Field
		LastTriggered   respjson.Field
		Scope           respjson.Field
		ValidUntil      respjson.Field
		WebhookID       respjson.Field
		WebhookType     respjson.Field
		WebhookURL      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Webhook) RawJSON() string { return r.JSON.raw }
func (r *Webhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookScope string

const (
	WebhookScopeAPIKey       WebhookScope = "api_key"
	WebhookScopeOrganization WebhookScope = "organization"
	WebhookScopeIntegration  WebhookScope = "integration"
)

type WebhookWebhookType string

const (
	WebhookWebhookTypeMessage WebhookWebhookType = "message"
	WebhookWebhookTypeStatus  WebhookWebhookType = "status"
	WebhookWebhookTypeAll     WebhookWebhookType = "all"
)

type WebhookNewResponse struct {
	Message string `json:"message"`
	// Any of "api_key", "organization".
	Scope      WebhookNewResponseScope `json:"scope"`
	WebhookID  string                  `json:"webhook_id"`
	WebhookURL string                  `json:"webhook_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Scope       respjson.Field
		WebhookID   respjson.Field
		WebhookURL  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookNewResponseScope string

const (
	WebhookNewResponseScopeAPIKey       WebhookNewResponseScope = "api_key"
	WebhookNewResponseScopeOrganization WebhookNewResponseScope = "organization"
)

type WebhookListResponse struct {
	Webhooks []Webhook `json:"webhooks"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Webhooks    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookListResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeleteResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookNewParams struct {
	// URL of an existing webhook, for the idempotent 200 response. A URL that does not
	// already exist returns 410.
	WebhookURL string `json:"webhook_url" api:"required" format:"uri"`
	// Ignored. Retained so existing request bodies stay valid.
	ValidUntil param.Opt[int64] `json:"valid_until,omitzero"`
	paramObj
}

func (r WebhookNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookUpdateParams struct {
	// Set to true to deprecate, false to undeprecate
	Deprecate param.Opt[bool] `json:"deprecate,omitzero"`
	// Expiration timestamp. Use -1 or null for no expiration.
	ValidUntil param.Opt[int64] `json:"valid_until,omitzero"`
	// Type of events to receive
	//
	// Any of "message", "status", "all".
	WebhookType WebhookUpdateParamsWebhookType `json:"webhook_type,omitzero"`
	paramObj
}

func (r WebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of events to receive
type WebhookUpdateParamsWebhookType string

const (
	WebhookUpdateParamsWebhookTypeMessage WebhookUpdateParamsWebhookType = "message"
	WebhookUpdateParamsWebhookTypeStatus  WebhookUpdateParamsWebhookType = "status"
	WebhookUpdateParamsWebhookTypeAll     WebhookUpdateParamsWebhookType = "all"
)
