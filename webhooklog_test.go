// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blooio_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Blooio/blooio-go-sdk"
	"github.com/Blooio/blooio-go-sdk/internal/testutil"
	"github.com/Blooio/blooio-go-sdk/option"
)

func TestWebhookLogListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := blooio.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Webhooks.Logs.List(
		context.TODO(),
		"wh_abc123def456",
		blooio.WebhookLogListParams{
			Limit:     blooio.Int(1),
			MaxStatus: blooio.Int(0),
			MinStatus: blooio.Int(0),
			Offset:    blooio.Int(0),
			Sort:      blooio.WebhookLogListParamsSortAsc,
			Status:    blooio.Int(0),
		},
	)
	if err != nil {
		var apierr *blooio.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebhookLogReplay(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := blooio.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Webhooks.Logs.Replay(
		context.TODO(),
		"eventId",
		blooio.WebhookLogReplayParams{
			WebhookID: "wh_abc123def456",
		},
	)
	if err != nil {
		var apierr *blooio.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
