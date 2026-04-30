// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blooio_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/blooio-go"
	"github.com/stainless-sdks/blooio-go/internal/testutil"
	"github.com/stainless-sdks/blooio-go/option"
)

func TestMeNumberContactCardGet(t *testing.T) {
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
	_, err := client.Me.Numbers.ContactCard.Get(context.TODO(), "number")
	if err != nil {
		var apierr *blooio.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMeNumberContactCardUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Me.Numbers.ContactCard.Update(
		context.TODO(),
		"number",
		blooio.MeNumberContactCardUpdateParams{
			Avatar:    blooio.String("avatar"),
			FirstName: blooio.String("first_name"),
			LastName:  blooio.String("last_name"),
			Sharing: blooio.MeNumberContactCardUpdateParamsSharing{
				Audience:   blooio.Int(0),
				Enabled:    blooio.Bool(true),
				NameFormat: blooio.Int(0),
			},
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
