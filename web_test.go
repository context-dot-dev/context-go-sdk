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
		URL:              "https://example.com",
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
		TimeoutMs:        contextdev.Int(1000),
		WaitForMs:        contextdev.Int(0),
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
		TimeoutMs:      contextdev.Int(1000),
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
		TimeoutMs: contextdev.Int(1000),
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
		TimeoutMs:   contextdev.Int(1000),
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
		ColorScheme:    contextdev.WebScreenshotParamsColorSchemeLight,
		Country:        contextdev.WebScreenshotParamsCountryDe,
		DirectURL:      contextdev.String("https://example.com"),
		Domain:         contextdev.String("xxx"),
		FullScreenshot: contextdev.WebScreenshotParamsFullScreenshotTrue,
		HandleCookiePopup: contextdev.WebScreenshotParamsHandleCookiePopupUnion{
			OfWebScreenshotsHandleCookiePopupString: contextdev.String("true"),
		},
		MaxAgeMs:     contextdev.Int(0),
		Page:         contextdev.WebScreenshotParamsPageLogin,
		ScrollOffset: contextdev.Int(0),
		Tags:         []string{"production", "team-alpha"},
		TimeoutMs:    contextdev.Int(1),
		Viewport: contextdev.WebScreenshotParamsViewport{
			Height: contextdev.Int(240),
			Width:  contextdev.Int(240),
		},
		WaitForMs: contextdev.Int(0),
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
			TimeoutMs:           contextdev.Int(1000),
			UseMainContentOnly:  contextdev.Bool(true),
			WaitForMs:           contextdev.Int(0),
		},
		NumResults:  contextdev.Int(10),
		QueryFanout: contextdev.Bool(true),
		Tags:        []string{"production", "team-alpha"},
		TimeoutMs:   contextdev.Int(1000),
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
		TimeoutMs:           contextdev.Int(1000),
		URLRegex:            contextdev.String("^https?://[^/]+/blog/"),
		UseMainContentOnly:  contextdev.Bool(true),
		WaitForMs:           contextdev.Int(0),
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
		URL:              "https://example.com",
		Country:          contextdev.WebWebScrapeHTMLParamsCountryDe,
		ExcludeSelectors: []string{"x"},
		Headers: map[string]string{
			"foo": "J!",
		},
		IncludeFrames: contextdev.WebWebScrapeHTMLParamsIncludeFramesUnion{
			OfWebWebScrapeHTMLsIncludeFramesString: contextdev.String("true"),
		},
		IncludeSelectors: []string{"x"},
		MaxAgeMs:         contextdev.Int(0),
		Pdf: contextdev.WebWebScrapeHTMLParamsPdf{
			End: contextdev.Int(1),
			Ocr: contextdev.WebWebScrapeHTMLParamsPdfOcrUnion{
				OfWebWebScrapeHTMLsPdfOcrString: contextdev.String("true"),
			},
			ShouldParse: contextdev.WebWebScrapeHTMLParamsPdfShouldParseUnion{
				OfWebWebScrapeHTMLsPdfShouldParseString: contextdev.String("true"),
			},
			Start: contextdev.Int(1),
		},
		SettleAnimations: contextdev.WebWebScrapeHTMLParamsSettleAnimationsUnion{
			OfWebWebScrapeHTMLsSettleAnimationsString: contextdev.String("true"),
		},
		Tags:      []string{"production", "team-alpha"},
		TimeoutMs: contextdev.Int(1),
		UseMainContentOnly: contextdev.WebWebScrapeHTMLParamsUseMainContentOnlyUnion{
			OfWebWebScrapeHTMLsUseMainContentOnlyString: contextdev.String("true"),
		},
		WaitForMs: contextdev.Int(0),
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
		Dedupe: contextdev.WebWebScrapeImagesParamsDedupeUnion{
			OfWebWebScrapeImagessDedupeString: contextdev.String("true"),
		},
		Enrichment: contextdev.WebWebScrapeImagesParamsEnrichment{
			Classification: contextdev.WebWebScrapeImagesParamsEnrichmentClassificationUnion{
				OfWebWebScrapeImagessEnrichmentClassificationString: contextdev.String("true"),
			},
			HostedURL: contextdev.WebWebScrapeImagesParamsEnrichmentHostedURLUnion{
				OfWebWebScrapeImagessEnrichmentHostedURLString: contextdev.String("true"),
			},
			MaxTimePerMs: contextdev.Int(1),
			Resolution: contextdev.WebWebScrapeImagesParamsEnrichmentResolutionUnion{
				OfWebWebScrapeImagessEnrichmentResolutionString: contextdev.String("true"),
			},
		},
		Headers: map[string]string{
			"foo": "J!",
		},
		MaxAgeMs:  contextdev.Int(0),
		Tags:      []string{"production", "team-alpha"},
		TimeoutMs: contextdev.Int(1),
		WaitForMs: contextdev.Int(0),
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
		URL:              "https://example.com",
		Country:          contextdev.WebWebScrapeMdParamsCountryDe,
		ExcludeSelectors: []string{"x"},
		Headers: map[string]string{
			"foo": "J!",
		},
		IncludeFrames: contextdev.WebWebScrapeMdParamsIncludeFramesUnion{
			OfWebWebScrapeMdsIncludeFramesString: contextdev.String("true"),
		},
		IncludeImages: contextdev.WebWebScrapeMdParamsIncludeImagesUnion{
			OfWebWebScrapeMdsIncludeImagesString: contextdev.String("true"),
		},
		IncludeLinks: contextdev.WebWebScrapeMdParamsIncludeLinksUnion{
			OfWebWebScrapeMdsIncludeLinksString: contextdev.String("true"),
		},
		IncludeSelectors: []string{"x"},
		MaxAgeMs:         contextdev.Int(0),
		Pdf: contextdev.WebWebScrapeMdParamsPdf{
			End: contextdev.Int(1),
			Ocr: contextdev.WebWebScrapeMdParamsPdfOcrUnion{
				OfWebWebScrapeMdsPdfOcrString: contextdev.String("true"),
			},
			ShouldParse: contextdev.WebWebScrapeMdParamsPdfShouldParseUnion{
				OfWebWebScrapeMdsPdfShouldParseString: contextdev.String("true"),
			},
			Start: contextdev.Int(1),
		},
		SettleAnimations: contextdev.WebWebScrapeMdParamsSettleAnimationsUnion{
			OfWebWebScrapeMdsSettleAnimationsString: contextdev.String("true"),
		},
		ShortenBase64Images: contextdev.WebWebScrapeMdParamsShortenBase64ImagesUnion{
			OfWebWebScrapeMdsShortenBase64ImagesString: contextdev.String("true"),
		},
		Tags:      []string{"production", "team-alpha"},
		TimeoutMs: contextdev.Int(1),
		UseMainContentOnly: contextdev.WebWebScrapeMdParamsUseMainContentOnlyUnion{
			OfWebWebScrapeMdsUseMainContentOnlyString: contextdev.String("true"),
		},
		WaitForMs: contextdev.Int(0),
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
		MaxLinks:   contextdev.Int(1),
		SitemapURL: contextdev.String("https://example.com"),
		Tags:       []string{"production", "team-alpha"},
		TimeoutMs:  contextdev.Int(1),
		URLRegex:   contextdev.String("^https?://[^/]+/blog/"),
	})
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
