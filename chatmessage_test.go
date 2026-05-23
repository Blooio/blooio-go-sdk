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

func TestChatMessageGet(t *testing.T) {
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
	_, err := client.Chats.Messages.Get(
		context.TODO(),
		"msg_abc123def456",
		blooio.ChatMessageGetParams{
			ChatID: "chatId",
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

func TestChatMessageListWithOptionalParams(t *testing.T) {
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
	_, err := client.Chats.Messages.List(
		context.TODO(),
		"chatId",
		blooio.ChatMessageListParams{
			Direction: blooio.ChatMessageListParamsDirectionInbound,
			Limit:     blooio.Int(1),
			Offset:    blooio.Int(0),
			Since:     blooio.Int(0),
			Sort:      blooio.ChatMessageListParamsSortAsc,
			Until:     blooio.Int(0),
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

func TestChatMessageGetStatus(t *testing.T) {
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
	_, err := client.Chats.Messages.GetStatus(
		context.TODO(),
		"msg_abc123def456",
		blooio.ChatMessageGetStatusParams{
			ChatID: "chatId",
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

func TestChatMessageReactWithOptionalParams(t *testing.T) {
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
	_, err := client.Chats.Messages.React(
		context.TODO(),
		"messageId",
		blooio.ChatMessageReactParams{
			ChatID:    "chatId",
			Reaction:  "+love",
			Direction: blooio.ChatMessageReactParamsDirectionInbound,
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

func TestChatMessageSendWithOptionalParams(t *testing.T) {
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
	_, err := client.Chats.Messages.Send(
		context.TODO(),
		"chatId",
		blooio.ChatMessageSendParams{
			Attachments: []blooio.ChatMessageSendParamsAttachmentUnion{{
				OfString: blooio.String("string"),
			}},
			Effect:     blooio.ChatMessageSendParamsEffectSlam,
			FromNumber: blooio.String("from_number"),
			LinkPreview: blooio.LinkPreviewParam{
				ImageURL: blooio.String("https://example.com"),
				Title:    blooio.String("title"),
			},
			Parts: []blooio.ChatMessageSendParamsPart{{
				LinkPreview: blooio.LinkPreviewParam{
					ImageURL: blooio.String("https://example.com"),
					Title:    blooio.String("title"),
				},
				Mention: blooio.String("mention"),
				Name:    blooio.String("name"),
				Text:    blooio.String("text"),
				URL:     blooio.String("url"),
			}},
			ShareContact: blooio.Bool(true),
			Text: blooio.ChatMessageSendParamsTextUnion{
				OfString: blooio.String("string"),
			},
			UseTypingIndicator: blooio.Bool(true),
			IdempotencyKey:     blooio.String("Idempotency-Key"),
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
