// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contextdev_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/context.dev-go"
	"github.com/stainless-sdks/context.dev-go/internal/testutil"
	"github.com/stainless-sdks/context.dev-go/option"
)

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
		Domain:    contextdev.String("domain"),
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
		DirectURL: contextdev.String("https://example.com"),
		Domain:    contextdev.String("domain"),
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
		DirectURL:         contextdev.String("https://example.com"),
		Domain:            contextdev.String("domain"),
		FullScreenshot:    contextdev.WebScreenshotParamsFullScreenshotTrue,
		HandleCookiePopup: contextdev.WebScreenshotParamsHandleCookiePopupTrue,
		MaxAgeMs:          contextdev.Int(0),
		Page:              contextdev.WebScreenshotParamsPageLogin,
		TimeoutMs:         contextdev.Int(1000),
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
		FollowSubdomains: contextdev.Bool(true),
		IncludeFrames:    contextdev.Bool(true),
		IncludeImages:    contextdev.Bool(true),
		IncludeLinks:     contextdev.Bool(true),
		MaxAgeMs:         contextdev.Int(0),
		MaxDepth:         contextdev.Int(0),
		MaxPages:         contextdev.Int(1),
		Pdf: contextdev.WebWebCrawlMdParamsPdf{
			End:         contextdev.Int(1),
			ShouldParse: contextdev.Bool(true),
			Start:       contextdev.Int(1),
		},
		ShortenBase64Images: contextdev.Bool(true),
		StopAfterMs:         contextdev.Int(10000),
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
		URL:           "https://example.com",
		IncludeFrames: contextdev.Bool(true),
		MaxAgeMs:      contextdev.Int(0),
		Pdf: contextdev.WebWebScrapeHTMLParamsPdf{
			End:         contextdev.Int(1),
			ShouldParse: contextdev.Bool(true),
			Start:       contextdev.Int(1),
		},
		TimeoutMs: contextdev.Int(1000),
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
		Enrichment: contextdev.WebWebScrapeImagesParamsEnrichment{
			Classification: contextdev.Bool(true),
			HostedURL:      contextdev.Bool(true),
			MaxTimePerMs:   contextdev.Int(1),
			Resolution:     contextdev.Bool(true),
		},
		MaxAgeMs:  contextdev.Int(0),
		TimeoutMs: contextdev.Int(1000),
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
		URL:           "https://example.com",
		IncludeFrames: contextdev.Bool(true),
		IncludeImages: contextdev.Bool(true),
		IncludeLinks:  contextdev.Bool(true),
		MaxAgeMs:      contextdev.Int(0),
		Pdf: contextdev.WebWebScrapeMdParamsPdf{
			End:         contextdev.Int(1),
			ShouldParse: contextdev.Bool(true),
			Start:       contextdev.Int(1),
		},
		ShortenBase64Images: contextdev.Bool(true),
		TimeoutMs:           contextdev.Int(1000),
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
		Domain:    "domain",
		MaxLinks:  contextdev.Int(1),
		TimeoutMs: contextdev.Int(1000),
		URLRegex:  contextdev.String("^https?://[^/]+/blog/"),
	})
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
