// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contextdev_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/context-dot-dev/context-go-sdk"
	"github.com/context-dot-dev/context-go-sdk/internal/testutil"
	"github.com/context-dot-dev/context-go-sdk/option"
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
		Domain:        "domain",
		ForceLanguage: contextdev.BrandGetParamsForceLanguageAfrikaans,
		MaxAgeMs:      contextdev.Int(86400000),
		MaxSpeed:      contextdev.Bool(true),
		TimeoutMs:     contextdev.Int(1000),
	})
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBrandIdentifyFromTransactionWithOptionalParams(t *testing.T) {
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
	_, err := client.Brand.IdentifyFromTransaction(context.TODO(), contextdev.BrandIdentifyFromTransactionParams{
		TransactionInfo:    "transaction_info",
		City:               contextdev.String("city"),
		CountryGl:          contextdev.BrandIdentifyFromTransactionParamsCountryGlAd,
		ForceLanguage:      contextdev.BrandIdentifyFromTransactionParamsForceLanguageAfrikaans,
		HighConfidenceOnly: contextdev.Bool(true),
		MaxSpeed:           contextdev.Bool(true),
		Mcc:                contextdev.String("mcc"),
		Phone:              contextdev.Float(0),
		TimeoutMs:          contextdev.Int(1000),
	})
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBrandGetByEmailWithOptionalParams(t *testing.T) {
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
	_, err := client.Brand.GetByEmail(context.TODO(), contextdev.BrandGetByEmailParams{
		Email:         "dev@stainless.com",
		ForceLanguage: contextdev.BrandGetByEmailParamsForceLanguageAfrikaans,
		MaxAgeMs:      contextdev.Int(86400000),
		MaxSpeed:      contextdev.Bool(true),
		TimeoutMs:     contextdev.Int(1000),
	})
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBrandGetByIsinWithOptionalParams(t *testing.T) {
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
	_, err := client.Brand.GetByIsin(context.TODO(), contextdev.BrandGetByIsinParams{
		Isin:          "SE60513A9993",
		ForceLanguage: contextdev.BrandGetByIsinParamsForceLanguageAfrikaans,
		MaxAgeMs:      contextdev.Int(86400000),
		MaxSpeed:      contextdev.Bool(true),
		TimeoutMs:     contextdev.Int(1000),
	})
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBrandGetByNameWithOptionalParams(t *testing.T) {
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
	_, err := client.Brand.GetByName(context.TODO(), contextdev.BrandGetByNameParams{
		Name:          "xxx",
		CountryGl:     contextdev.BrandGetByNameParamsCountryGlAd,
		ForceLanguage: contextdev.BrandGetByNameParamsForceLanguageAfrikaans,
		MaxAgeMs:      contextdev.Int(86400000),
		MaxSpeed:      contextdev.Bool(true),
		TimeoutMs:     contextdev.Int(1000),
	})
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBrandGetByTickerWithOptionalParams(t *testing.T) {
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
	_, err := client.Brand.GetByTicker(context.TODO(), contextdev.BrandGetByTickerParams{
		Ticker:         "ticker",
		ForceLanguage:  contextdev.BrandGetByTickerParamsForceLanguageAfrikaans,
		MaxAgeMs:       contextdev.Int(86400000),
		MaxSpeed:       contextdev.Bool(true),
		TickerExchange: contextdev.BrandGetByTickerParamsTickerExchangeAmex,
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
		Domain:    "domain",
		MaxAgeMs:  contextdev.Int(86400000),
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
