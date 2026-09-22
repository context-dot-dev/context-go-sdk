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

func TestWebAnswersWithOptionalParams(t *testing.T) {
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
	_, err := client.Web.Answers(context.TODO(), contextdev.WebAnswersParams{
		Task: "Find the pricing page URL and plan names for context.dev.",
		JsonFormat: map[string]any{
			"pricing_page_url": "bar",
			"plans":            "bar",
		},
		Mode: contextdev.WebAnswersParamsModeFast,
		Tags: []string{"production", "team-alpha"},
		TimeoutOpts: contextdev.WebAnswersParamsTimeoutOpts{
			Milliseconds: 1000,
			Behavior:     "fail",
		},
		Zdr: contextdev.WebAnswersParamsZdrEnabled,
	})
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebExtractCompetitorsWithOptionalParams(t *testing.T) {
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
	_, err := client.Web.ExtractCompetitors(context.TODO(), contextdev.WebExtractCompetitorsParams{
		Domain:         "xxx",
		NumCompetitors: contextdev.Int(1),
		Tags:           []string{"production", "team-alpha"},
		TimeoutOpts: contextdev.WebExtractCompetitorsParamsTimeoutOpts{
			Milliseconds: 1000,
			Behavior:     "fail",
		},
		Zdr: contextdev.WebExtractCompetitorsParamsZdrEnabled,
	})
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebExtractStyleguideWithOptionalParams(t *testing.T) {
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
	_, err := client.Web.ExtractStyleguide(context.TODO(), contextdev.WebExtractStyleguideParams{
		ColorScheme: contextdev.WebExtractStyleguideParamsColorSchemeLight,
		DirectURL:   contextdev.String("https://example.com"),
		Domain:      contextdev.String("xxx"),
		MaxAgeMs:    contextdev.Int(0),
		Tags:        []string{"production", "team-alpha"},
		TimeoutOpts: contextdev.WebExtractStyleguideParamsTimeoutOpts{
			Milliseconds: 1,
			Behavior:     "fail",
		},
		Zdr: contextdev.WebExtractStyleguideParamsZdrEnabled,
	})
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebMapURLsWithOptionalParams(t *testing.T) {
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
	_, err := client.Web.MapURLs(context.TODO(), contextdev.WebMapURLsParams{
		Domain: "xxx",
		Headers: map[string]string{
			"foo": "J!",
		},
		IncludeSubdomains: contextdev.Bool(true),
		MaxLinks:          contextdev.Int(1),
		Search:            contextdev.String("help center and troubleshooting articles"),
		SitemapURL:        contextdev.String("https://example.com"),
		Tags:              []string{"production", "team-alpha"},
		TimeoutOpts: contextdev.WebMapURLsParamsTimeoutOpts{
			Milliseconds: 1,
			Behavior:     "fail",
		},
		URLRegex: contextdev.String("^https?://[^/]+/blog/"),
		Zdr:      contextdev.WebMapURLsParamsZdrEnabled,
	})
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebScrapeWithOptionalParams(t *testing.T) {
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
	_, err := client.Web.Scrape(context.TODO(), contextdev.WebScrapeParams{
		Formats: contextdev.WebScrapeParamsFormats{
			Bytes:      contextdev.Bool(true),
			HTML:       contextdev.Bool(true),
			Images:     contextdev.Bool(true),
			Markdown:   contextdev.Bool(true),
			Parse:      contextdev.Bool(true),
			Screenshot: contextdev.Bool(true),
		},
		URL: "https://example.com",
		ImageParams: contextdev.WebScrapeParamsImageParams{
			Dedupe: "none",
			Enrich: []string{"dimensions"},
		},
		MarkdownParams: contextdev.WebScrapeParamsMarkdownParams{
			IncludeImages: contextdev.Bool(true),
			IncludeLinks:  contextdev.Bool(true),
			InlineImages:  "placeholder",
		},
		MaxAgeMs: contextdev.Int(0),
		ParseParams: contextdev.WebScrapeParamsParseParams{
			Rules: map[string]contextdev.WebScrapeParamsParseParamsRuleUnion{
				"title": {
					OfString: contextdev.String("h1"),
				},
				"links": {
					OfWebScrapesParseParamsRuleObject: &contextdev.WebScrapeParamsParseParamsRuleObject{
						Selector: "a",
						Output:   "text",
						Type:     "list",
					},
				},
			},
		},
		ScreenshotParams: contextdev.WebScrapeParamsScreenshotParams{
			Area: contextdev.WebScrapeParamsScreenshotParamsAreaUnion{
				OfPage: contextdev.String("viewport"),
			},
			Format: "png",
		},
		SharedParams: contextdev.WebScrapeParamsSharedParams{
			Actions: []contextdev.WebScrapeParamsSharedParamsActionUnion{{
				OfPerform: &contextdev.WebScrapeParamsSharedParamsActionPerform{
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
			Parsers: contextdev.WebScrapeParamsSharedParamsParsers{
				Pdf: contextdev.WebScrapeParamsSharedParamsParsersPdf{
					EndPage:   contextdev.Int(1),
					Ocr:       "off",
					StartPage: contextdev.Int(1),
				},
			},
			SettleAnimations: contextdev.Bool(true),
			Theme:            "light",
			Viewport: contextdev.WebScrapeParamsSharedParamsViewport{
				Height: contextdev.Int(240),
				Width:  contextdev.Int(240),
			},
			WaitFor: contextdev.WebScrapeParamsSharedParamsWaitForUnion{
				OfInt: contextdev.Int(500),
			},
		},
		Tags: []string{"production", "team-alpha"},
		TimeoutOpts: contextdev.WebScrapeParamsTimeoutOpts{
			Milliseconds: 1,
			Behavior:     "fail",
		},
		Zdr: contextdev.WebScrapeParamsZdrEnabled,
	})
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebScreenshotWithOptionalParams(t *testing.T) {
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
	_, err := client.Web.Screenshot(context.TODO(), contextdev.WebScreenshotParams{
		ClearPopups:       contextdev.Bool(true),
		ColorScheme:       contextdev.WebScreenshotParamsColorSchemeLight,
		Country:           contextdev.WebScreenshotParamsCountryDe,
		DirectURL:         contextdev.String("https://example.com"),
		Domain:            contextdev.String("xxx"),
		FullScreenshot:    contextdev.WebScreenshotParamsFullScreenshotTrue,
		HandleCookiePopup: contextdev.Bool(true),
		Headers: map[string]string{
			"foo": "J!",
		},
		MaxAgeMs:     contextdev.Int(0),
		Page:         contextdev.WebScreenshotParamsPageLogin,
		ScrollOffset: contextdev.Int(0),
		Tags:         []string{"production", "team-alpha"},
		TimeoutOpts: contextdev.WebScreenshotParamsTimeoutOpts{
			Milliseconds: 1,
			Behavior:     "fail",
		},
		Viewport: contextdev.WebScreenshotParamsViewport{
			Height: contextdev.Int(240),
			Width:  contextdev.Int(240),
		},
		WaitForMs: contextdev.Int(0),
		Zdr:       contextdev.WebScreenshotParamsZdrEnabled,
	})
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Web.Search(context.TODO(), contextdev.WebSearchParams{
		Query:          "x",
		Country:        contextdev.WebSearchParamsCountryAf,
		ExcludeDomains: []string{"string"},
		Freshness:      contextdev.WebSearchParamsFreshnessLast24Hours,
		IncludeDomains: []string{"string"},
		MarkdownOptions: contextdev.WebSearchParamsMarkdownOptions{
			Enabled:       contextdev.Bool(true),
			IncludeFrames: contextdev.Bool(true),
			IncludeImages: contextdev.Bool(true),
			IncludeLinks:  contextdev.Bool(true),
			MaxAgeMs:      contextdev.Int(0),
			Pdf: contextdev.WebSearchParamsMarkdownOptionsPdf{
				End:         contextdev.Int(1),
				ShouldParse: contextdev.Bool(true),
				Start:       contextdev.Int(1),
			},
			ShortenBase64Images: contextdev.Bool(true),
			TimeoutOpts: contextdev.WebSearchParamsMarkdownOptionsTimeoutOpts{
				Milliseconds: 1,
				Behavior:     "fail",
			},
			UseMainContentOnly: contextdev.Bool(true),
			WaitForMs:          contextdev.Int(0),
		},
		NumResults:  contextdev.Int(10),
		QueryFanout: contextdev.Bool(true),
		Tags:        []string{"production", "team-alpha"},
		TimeoutOpts: contextdev.WebSearchParamsTimeoutOpts{
			Milliseconds: 1000,
			Behavior:     "fail",
		},
		Zdr: contextdev.WebSearchParamsZdrEnabled,
	})
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebWebCrawlMdWithOptionalParams(t *testing.T) {
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
	_, err := client.Web.WebCrawlMd(context.TODO(), contextdev.WebWebCrawlMdParams{
		URL:              "https://example.com",
		Country:          contextdev.WebWebCrawlMdParamsCountryDe,
		ExcludeSelectors: []string{"string"},
		FollowSubdomains: contextdev.Bool(true),
		IncludeFrames:    contextdev.Bool(true),
		IncludeImages:    contextdev.Bool(true),
		IncludeLinks:     contextdev.Bool(true),
		IncludeSelectors: []string{"string"},
		MaxAgeMs:         contextdev.Int(0),
		MaxDepth:         contextdev.Int(0),
		MaxPages:         contextdev.Int(1),
		Pdf: contextdev.WebWebCrawlMdParamsPdf{
			End:         contextdev.Int(1),
			Ocr:         contextdev.Bool(true),
			ShouldParse: contextdev.Bool(true),
			Start:       contextdev.Int(1),
		},
		SettleAnimations:    contextdev.Bool(true),
		ShortenBase64Images: contextdev.Bool(true),
		StopAfterMs:         contextdev.Int(10000),
		Tags:                []string{"production", "team-alpha"},
		TimeoutOpts: contextdev.WebWebCrawlMdParamsTimeoutOpts{
			Milliseconds: 1000,
			Behavior:     "fail",
		},
		URLRegex:           contextdev.String("^https?://[^/]+/blog/"),
		UseMainContentOnly: contextdev.Bool(true),
		WaitForMs:          contextdev.Int(0),
		Zdr:                contextdev.WebWebCrawlMdParamsZdrEnabled,
	})
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
