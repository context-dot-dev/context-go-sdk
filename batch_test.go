// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contextdev_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/context-dot-dev/context-go-sdk/v2"
	"github.com/context-dot-dev/context-go-sdk/v2/internal/testutil"
	"github.com/context-dot-dev/context-go-sdk/v2/option"
)

func TestBatchGet(t *testing.T) {
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
	_, err := client.Batch.Get(context.TODO(), "batch_9f2c8a")
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBatchListWithOptionalParams(t *testing.T) {
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
	_, err := client.Batch.List(context.TODO(), contextdev.BatchListParams{
		Cursor:     contextdev.String("cursor"),
		Limit:      contextdev.Int(1),
		Q:          contextdev.String("batch_1a2b"),
		SearchType: contextdev.BatchListParamsSearchTypeExact,
		Status:     contextdev.BatchListParamsStatusQueued,
		Tags:       contextdev.String("docs,competitor"),
	})
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBatchDelete(t *testing.T) {
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
	_, err := client.Batch.Delete(context.TODO(), "batch_9f2c8a")
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBatchCancel(t *testing.T) {
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
	_, err := client.Batch.Cancel(context.TODO(), "batch_9f2c8a")
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBatchGetResultsWithOptionalParams(t *testing.T) {
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
	_, err := client.Batch.GetResults(
		context.TODO(),
		"batch_9f2c8a",
		contextdev.BatchGetResultsParams{
			Cursor: contextdev.String("cursor"),
			Limit:  contextdev.Int(1),
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

func TestBatchSubmitWithOptionalParams(t *testing.T) {
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
	_, err := client.Batch.Submit(context.TODO(), contextdev.BatchSubmitParams{
		Input: contextdev.BatchSubmitParamsInputUnion{
			OfScrape: &contextdev.BatchSubmitParamsInputScrape{
				Data: contextdev.BatchSubmitParamsInputScrapeDataUnion{
					OfMarkdown: &contextdev.BatchSubmitParamsInputScrapeDataMarkdown{
						URLs: []contextdev.BatchSubmitParamsInputScrapeDataMarkdownURL{{
							URL:    "https://example.com/products/anvil",
							ItemID: contextdev.String("sku-1"),
							Meta: map[string]any{
								"category": "bar",
							},
						}, {
							URL:    "https://example.com/products/hammer",
							ItemID: contextdev.String("sku-2"),
							Meta: map[string]any{
								"foo": "bar",
							},
						}},
						Options: contextdev.BatchSubmitParamsInputScrapeDataMarkdownOptions{
							Country:          "de",
							ExcludeSelectors: []string{"x"},
							IncludeHTML:      contextdev.Bool(true),
							IncludeImages:    contextdev.Bool(true),
							IncludeLinks:     contextdev.Bool(true),
							IncludeSelectors: []string{"x"},
							MaxAgeMs:         contextdev.Int(0),
							Pdf: contextdev.BatchSubmitParamsInputScrapeDataMarkdownOptionsPdf{
								End:         contextdev.Int(1),
								Ocr:         contextdev.Bool(true),
								ShouldParse: contextdev.Bool(true),
								Start:       contextdev.Int(1),
							},
							SettleAnimations:    contextdev.Bool(true),
							ShortenBase64Images: contextdev.Bool(true),
							UseMainContentOnly:  contextdev.Bool(true),
							WaitForMs:           contextdev.Int(0),
						},
					},
				},
			},
		},
		Tags:           []string{"docs", "competitor"},
		WebhookURL:     contextdev.String("webhookUrl"),
		IdempotencyKey: contextdev.String("Idempotency-Key"),
	})
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
