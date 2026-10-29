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
		Cursor:               contextdev.String("cursor"),
		Limit:                contextdev.Int(1),
		Q:                    contextdev.String("batch_1a2b"),
		SearchType:           contextdev.BatchListParamsSearchTypeExact,
		ShouldIncludeUnified: contextdev.Bool(true),
		Status:               contextdev.BatchListParamsStatusQueued,
		Tags: contextdev.BatchListParamsTagsUnion{
			OfString: contextdev.String("docs,competitor"),
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
			Cursor: contextdev.String("321669910225:155771193"),
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
					OfScrapeAnything: &contextdev.BatchSubmitParamsInputScrapeDataScrapeAnything{
						Formats: contextdev.BatchSubmitParamsInputScrapeDataScrapeAnythingFormats{
							Bytes:      contextdev.Bool(true),
							Highlights: contextdev.Bool(true),
							HTML:       contextdev.Bool(true),
							Images:     contextdev.Bool(true),
							Json:       contextdev.Bool(true),
							Markdown:   contextdev.Bool(true),
							Parse:      contextdev.Bool(true),
							Product:    contextdev.Bool(true),
							Screenshot: contextdev.Bool(true),
						},
						URLs: []contextdev.BatchSubmitParamsInputScrapeDataScrapeAnythingURL{{
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
						HighlightsParams: contextdev.BatchSubmitParamsInputScrapeDataScrapeAnythingHighlightsParams{
							Query:         "x",
							MaxCharacters: contextdev.Int(100),
						},
						ImageParams: contextdev.BatchSubmitParamsInputScrapeDataScrapeAnythingImageParams{
							Dedupe: "none",
							Enrich: []string{"dimensions"},
						},
						JsonParams: contextdev.BatchSubmitParamsInputScrapeDataScrapeAnythingJsonParams{
							Schema: map[string]any{
								"type":                 "bar",
								"properties":           "bar",
								"required":             "bar",
								"additionalProperties": "bar",
							},
							Instructions: contextdev.String("instructions"),
						},
						MarkdownParams: contextdev.BatchSubmitParamsInputScrapeDataScrapeAnythingMarkdownParams{
							IncludeImages:    contextdev.Bool(true),
							IncludeLinks:     contextdev.Bool(true),
							InlineImages:     "placeholder",
							SourcePreference: []string{"html-conversion"},
						},
						MaxAgeMs: contextdev.Int(0),
						ParseParams: contextdev.BatchSubmitParamsInputScrapeDataScrapeAnythingParseParamsUnion{
							OfCustomParser: &contextdev.BatchSubmitParamsInputScrapeDataScrapeAnythingParseParamsCustomParser{
								Rules: map[string]contextdev.BatchSubmitParamsInputScrapeDataScrapeAnythingParseParamsCustomParserRuleUnion{
									"title": {
										OfString: contextdev.String("h1"),
									},
									"links": {
										OfBatchSubmitsInputScrapeDataScrapeAnythingParseParamsCustomParserRuleObject: &contextdev.BatchSubmitParamsInputScrapeDataScrapeAnythingParseParamsCustomParserRuleObject{
											Selector: "a",
											Output:   "text",
											Type:     "list",
										},
									},
								},
								Type: "custom",
							},
						},
						ProductParams: contextdev.BatchSubmitParamsInputScrapeDataScrapeAnythingProductParams{
							DedupeImages:  contextdev.Bool(true),
							UseAIFallback: contextdev.Bool(true),
						},
						ScreenshotParams: contextdev.BatchSubmitParamsInputScrapeDataScrapeAnythingScreenshotParams{
							Area: contextdev.BatchSubmitParamsInputScrapeDataScrapeAnythingScreenshotParamsAreaUnion{
								OfPage: contextdev.String("viewport"),
							},
							Format: "png",
						},
						SharedParams: contextdev.BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParams{
							Actions: []contextdev.BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionUnion{{
								OfPerform: &contextdev.BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionPerform{
									Action: "Click the product details tab",
								},
							}},
							Country:          contextdev.String("US"),
							DismissCookies:   contextdev.Bool(true),
							DismissPopups:    contextdev.Bool(true),
							ExcludeSelectors: []string{"P"},
							Headers: map[string]string{
								"Accept-Language": "en-US",
							},
							IncludeFrames:    contextdev.Bool(true),
							IncludeSelectors: []string{"P"},
							MainContentOnly:  contextdev.Bool(true),
							Parsers: contextdev.BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsParsers{
								Pdf: contextdev.BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsParsersPdf{
									EndPage:              contextdev.Int(1),
									IncludeNumberedPages: contextdev.Bool(true),
									Ocr:                  "off",
									StartPage:            contextdev.Int(1),
								},
							},
							Profile: contextdev.BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsProfileUnion{
								OfExisting: &contextdev.BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsProfileExisting{
									ProfileID: "prof_210b9798eb53baa4e69d31c1071cf03d",
								},
							},
							SettleAnimations: contextdev.Bool(true),
							Theme:            "light",
							Viewport: contextdev.BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsViewport{
								Height: contextdev.Int(240),
								Width:  contextdev.Int(240),
							},
							WaitFor: contextdev.BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsWaitForUnion{
								OfInt: contextdev.Int(500),
							},
						},
						Tags: []string{"production", "team-alpha"},
						TimeoutOpts: contextdev.BatchSubmitParamsInputScrapeDataScrapeAnythingTimeoutOpts{
							Milliseconds: 1,
							Behavior:     "fail",
						},
						Zdr: "disabled",
					},
				},
			},
		},
		Policies: []string{"pol_3f9c2a7e5b8d4c1f9a6e0b2d7c4f8a1e"},
		Tags:     []string{"docs", "competitor"},
		Webhook: contextdev.BatchSubmitParamsWebhook{
			URL: "https://example.com",
			Retry: contextdev.RetryConfigParam{
				DelaysSeconds: []int64{10, 60, 300, 1800, 7200, 21600, 57600},
			},
		},
		WebhookURL:     contextdev.String("https://example.com"),
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
