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

func TestWebExtractWithOptionalParams(t *testing.T) {
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
	_, err := client.Web.Extract(context.TODO(), contextdev.WebExtractParams{
		Schema: map[string]any{
			"type":                 "bar",
			"properties":           "bar",
			"required":             "bar",
			"additionalProperties": "bar",
		},
		URL: "https://example.com",
		Actions: []contextdev.WebExtractParamsActionUnion{{
			OfWait: &contextdev.WebExtractParamsActionWait{
				TimeMs: 0,
			},
		}},
		FactCheck:        contextdev.Bool(true),
		FollowSubdomains: contextdev.Bool(true),
		IncludeFrames:    contextdev.Bool(true),
		Instructions:     contextdev.String("instructions"),
		MaxAgeMs:         contextdev.Int(0),
		MaxDepth:         contextdev.Int(0),
		MaxPages:         contextdev.Int(1),
		Pdf: contextdev.WebExtractParamsPdf{
			End:         contextdev.Int(1),
			ShouldParse: contextdev.Bool(true),
			Start:       contextdev.Int(1),
		},
		SettleAnimations: contextdev.Bool(true),
		StopAfterMs:      contextdev.Int(10000),
		Tags:             []string{"production", "team-alpha"},
		TimeoutOpts: contextdev.WebExtractParamsTimeoutOpts{
			Milliseconds: 1000,
			Behavior:     "fail",
		},
		WaitForMs: contextdev.Int(0),
		Zdr:       contextdev.WebExtractParamsZdrEnabled,
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

func TestWebExtractFontsWithOptionalParams(t *testing.T) {
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
	_, err := client.Web.ExtractFonts(context.TODO(), contextdev.WebExtractFontsParams{
		DirectURL: contextdev.String("https://example.com"),
		Domain:    contextdev.String("xxx"),
		MaxAgeMs:  contextdev.Int(0),
		Tags:      []string{"production", "team-alpha"},
		TimeoutOpts: contextdev.WebExtractFontsParamsTimeoutOpts{
			Milliseconds: 1,
			Behavior:     "fail",
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
		MaxAgeMs:          contextdev.Int(0),
		Page:              contextdev.WebScreenshotParamsPageLogin,
		ScrollOffset:      contextdev.Int(0),
		Tags:              []string{"production", "team-alpha"},
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

func TestWebWebScrapeBytesWithOptionalParams(t *testing.T) {
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
	_, err := client.Web.WebScrapeBytes(context.TODO(), contextdev.WebWebScrapeBytesParams{
		URL:     "https://example.com",
		Country: contextdev.WebWebScrapeBytesParamsCountryDe,
		Headers: map[string]string{
			"foo": "J!",
		},
		Tags: []string{"production", "team-alpha"},
		TimeoutOpts: contextdev.WebWebScrapeBytesParamsTimeoutOpts{
			Milliseconds: 1,
			Behavior:     "fail",
		},
		Zdr: contextdev.WebWebScrapeBytesParamsZdrEnabled,
	})
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebWebScrapeHTMLWithOptionalParams(t *testing.T) {
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
	_, err := client.Web.WebScrapeHTML(context.TODO(), contextdev.WebWebScrapeHTMLParams{
		URL: "https://example.com",
		Actions: []contextdev.WebWebScrapeHTMLParamsActionUnion{{
			OfWait: &contextdev.WebWebScrapeHTMLParamsActionWait{
				TimeMs: 0,
			},
		}},
		Country:          contextdev.WebWebScrapeHTMLParamsCountryDe,
		ExcludeSelectors: []string{"x"},
		ExtractRules: map[string]contextdev.WebWebScrapeHTMLParamsExtractRuleUnion{
			"foo": {
				OfString: contextdev.String("x"),
			},
		},
		Headers: map[string]string{
			"foo": "J!",
		},
		IncludeFrames:    contextdev.Bool(true),
		IncludeSelectors: []string{"x"},
		MaxAgeMs:         contextdev.Int(0),
		Pdf: contextdev.WebWebScrapeHTMLParamsPdf{
			End:         contextdev.Int(1),
			Ocr:         contextdev.Bool(true),
			ShouldParse: contextdev.Bool(true),
			Start:       contextdev.Int(1),
		},
		SettleAnimations: contextdev.Bool(true),
		Tags:             []string{"production", "team-alpha"},
		TimeoutOpts: contextdev.WebWebScrapeHTMLParamsTimeoutOpts{
			Milliseconds: 1,
			Behavior:     "fail",
		},
		UseMainContentOnly: contextdev.Bool(true),
		WaitForMs:          contextdev.Int(0),
		Zdr:                contextdev.WebWebScrapeHTMLParamsZdrEnabled,
	})
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebWebScrapeImagesWithOptionalParams(t *testing.T) {
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
	_, err := client.Web.WebScrapeImages(context.TODO(), contextdev.WebWebScrapeImagesParams{
		URL: "https://example.com",
		Actions: []contextdev.WebWebScrapeImagesParamsActionUnion{{
			OfWait: &contextdev.WebWebScrapeImagesParamsActionWait{
				TimeMs: 0,
			},
		}},
		Dedupe: contextdev.Bool(true),
		Enrichment: contextdev.WebWebScrapeImagesParamsEnrichment{
			Classification: contextdev.Bool(true),
			HostedURL:      contextdev.Bool(true),
			MaxTimePerMs:   contextdev.Int(1),
			Resolution:     contextdev.Bool(true),
		},
		Headers: map[string]string{
			"foo": "J!",
		},
		MaxAgeMs: contextdev.Int(0),
		Tags:     []string{"production", "team-alpha"},
		TimeoutOpts: contextdev.WebWebScrapeImagesParamsTimeoutOpts{
			Milliseconds: 1,
			Behavior:     "fail",
		},
		WaitForMs: contextdev.Int(0),
		Zdr:       contextdev.WebWebScrapeImagesParamsZdrEnabled,
	})
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebWebScrapeMdWithOptionalParams(t *testing.T) {
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
	_, err := client.Web.WebScrapeMd(context.TODO(), contextdev.WebWebScrapeMdParams{
		URL: "https://example.com",
		Actions: []contextdev.WebWebScrapeMdParamsActionUnion{{
			OfWait: &contextdev.WebWebScrapeMdParamsActionWait{
				TimeMs: 0,
			},
		}},
		Country:          contextdev.WebWebScrapeMdParamsCountryDe,
		ExcludeSelectors: []string{"x"},
		Headers: map[string]string{
			"foo": "J!",
		},
		IncludeFrames:    contextdev.Bool(true),
		IncludeHTML:      contextdev.Bool(true),
		IncludeImages:    contextdev.Bool(true),
		IncludeLinks:     contextdev.Bool(true),
		IncludeSelectors: []string{"x"},
		MaxAgeMs:         contextdev.Int(0),
		Pdf: contextdev.WebWebScrapeMdParamsPdf{
			End:         contextdev.Int(1),
			Ocr:         contextdev.Bool(true),
			ShouldParse: contextdev.Bool(true),
			Start:       contextdev.Int(1),
		},
		SettleAnimations:    contextdev.Bool(true),
		ShortenBase64Images: contextdev.Bool(true),
		Tags:                []string{"production", "team-alpha"},
		TimeoutOpts: contextdev.WebWebScrapeMdParamsTimeoutOpts{
			Milliseconds: 1,
			Behavior:     "fail",
		},
		UseMainContentOnly: contextdev.Bool(true),
		WaitForMs:          contextdev.Int(0),
		Zdr:                contextdev.WebWebScrapeMdParamsZdrEnabled,
	})
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebWebScrapeScreenshotWithOptionalParams(t *testing.T) {
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
	_, err := client.Web.WebScrapeScreenshot(context.TODO(), contextdev.WebWebScrapeScreenshotParams{
		URL:               "https://example.com",
		ClearPopups:       contextdev.Bool(true),
		ColorScheme:       contextdev.WebWebScrapeScreenshotParamsColorSchemeLight,
		Country:           contextdev.WebWebScrapeScreenshotParamsCountryDe,
		FullScreenshot:    contextdev.WebWebScrapeScreenshotParamsFullScreenshotTrue,
		HandleCookiePopup: contextdev.Bool(true),
		MaxAgeMs:          contextdev.Int(0),
		ScrollOffset:      contextdev.Int(0),
		Tags:              []string{"production", "team-alpha"},
		TimeoutOpts: contextdev.WebWebScrapeScreenshotParamsTimeoutOpts{
			Milliseconds: 1,
			Behavior:     "fail",
		},
		Viewport: contextdev.WebWebScrapeScreenshotParamsViewport{
			Height: contextdev.Int(240),
			Width:  contextdev.Int(240),
		},
		WaitForMs: contextdev.Int(0),
		Zdr:       contextdev.WebWebScrapeScreenshotParamsZdrEnabled,
	})
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebWebScrapeSitemapWithOptionalParams(t *testing.T) {
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
	_, err := client.Web.WebScrapeSitemap(context.TODO(), contextdev.WebWebScrapeSitemapParams{
		Domain: "xxx",
		Headers: map[string]string{
			"foo": "J!",
		},
		IncludeSubdomains: contextdev.Bool(true),
		MaxLinks:          contextdev.Int(1),
		Search:            contextdev.String("help center and troubleshooting articles"),
		SitemapURL:        contextdev.String("https://example.com"),
		Tags:              []string{"production", "team-alpha"},
		TimeoutOpts: contextdev.WebWebScrapeSitemapParamsTimeoutOpts{
			Milliseconds: 1,
			Behavior:     "fail",
		},
		URLRegex: contextdev.String("^https?://[^/]+/blog/"),
		Zdr:      contextdev.WebWebScrapeSitemapParamsZdrEnabled,
	})
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
