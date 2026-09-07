// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contextdev_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/context-dot-dev/context-go-sdk/v2"
	"github.com/context-dot-dev/context-go-sdk/v2/internal/testutil"
	"github.com/context-dot-dev/context-go-sdk/v2/option"
)

func TestWebhookDeliveryGetWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := contextdev.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Webhooks.Deliveries.Get(
		context.TODO(),
		"whd_210b9798eb53baa4e69d31c1071cf03d",
		contextdev.WebhookDeliveryGetParams{
			Tags: []string{"production", "team-alpha"},
		},
	)
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebhookDeliveryListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := contextdev.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Webhooks.Deliveries.List(context.TODO(), contextdev.WebhookDeliveryListParams{
		OfBatch: &contextdev.WebhookDeliveryListParamsBodyBatch{
			BatchID:      contextdev.String("batch_id"),
			CreatedAfter: contextdev.Time(time.Now()),
			Cursor:       contextdev.String("whd_210b9798eb53baa4e69d31c1071cf03d"),
			Limit:        contextdev.Int(25),
			Status:       "failed",
			Tags:         []string{"production", "team-alpha"},
		},
	})
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebhookDeliveryListAttemptsWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := contextdev.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Webhooks.Deliveries.ListAttempts(
		context.TODO(),
		"whd_210b9798eb53baa4e69d31c1071cf03d",
		contextdev.WebhookDeliveryListAttemptsParams{
			Cursor: contextdev.String("wha_210b9798eb53baa4e69d31c1071cf03d"),
			Limit:  contextdev.Int(1),
			Tags:   []string{"production", "team-alpha"},
		},
	)
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebhookDeliveryRetryWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := contextdev.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Webhooks.Deliveries.Retry(
		context.TODO(),
		"whd_210b9798eb53baa4e69d31c1071cf03d",
		contextdev.WebhookDeliveryRetryParams{
			Force:          contextdev.Bool(true),
			Tags:           []string{"production", "team-alpha"},
			IdempotencyKey: contextdev.String("Idempotency-Key"),
		},
	)
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
