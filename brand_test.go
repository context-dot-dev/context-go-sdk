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

func TestBrandGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Brand.Get(context.TODO(), contextdev.BrandGetParams{
		OfByDomain: &contextdev.BrandGetParamsBodyByDomain{
			Domain:        "stripe.com",
			ForceLanguage: "afrikaans",
			MaxAgeMs:      contextdev.Int(0),
			MaxSpeed:      contextdev.Bool(true),
			Tags:          []string{"production", "team-alpha"},
			TimeoutMs:     contextdev.Int(1000),
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

func TestBrandGetSimplifiedWithOptionalParams(t *testing.T) {
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
	_, err := client.Brand.GetSimplified(context.TODO(), contextdev.BrandGetSimplifiedParams{
		Domain:    "xxx",
		MaxAgeMs:  contextdev.Int(0),
		Tags:      []string{"production", "team-alpha"},
		Theme:     contextdev.BrandGetSimplifiedParamsThemeLight,
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

func TestBrandSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Brand.Search(context.TODO(), contextdev.BrandSearchParams{
		Query: "x",
		Tags:  []string{"production", "team-alpha"},
	})
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
