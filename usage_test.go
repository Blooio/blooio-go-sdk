// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blooio_test

import (
	"context"
	"os"
	"testing"

	"github.com/Blooio/blooio-go-sdk"
	"github.com/Blooio/blooio-go-sdk/internal/testutil"
	"github.com/Blooio/blooio-go-sdk/option"
)

func TestUsage(t *testing.T) {
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
	t.Skip("Prism tests are disabled")
	me, err := client.Me.Get(context.TODO())
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", me.Valid)
}
