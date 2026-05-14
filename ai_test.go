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

func TestAIAIQueryWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.AIQuery(context.TODO(), contextdev.AIAIQueryParams{
		DataToExtract: []contextdev.AIAIQueryParamsDataToExtract{{
			DatapointDescription: "datapoint_description",
			DatapointExample:     "datapoint_example",
			DatapointName:        "datapoint_name",
			DatapointType:        "text",
			DatapointListType:    "string",
			DatapointObjectSchema: map[string]string{
				"testimonial_text":   "string",
				"testimonial_author": "string",
			},
		}},
		Domain: "domain",
		SpecificPages: contextdev.AIAIQueryParamsSpecificPages{
			AboutUs:            contextdev.Bool(true),
			Blog:               contextdev.Bool(true),
			Careers:            contextdev.Bool(true),
			ContactUs:          contextdev.Bool(true),
			Faq:                contextdev.Bool(true),
			HomePage:           contextdev.Bool(true),
			Pricing:            contextdev.Bool(true),
			PrivacyPolicy:      contextdev.Bool(true),
			TermsAndConditions: contextdev.Bool(true),
		},
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

func TestAIExtractProductWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.ExtractProduct(context.TODO(), contextdev.AIExtractProductParams{
		URL:       "https://example.com",
		MaxAgeMs:  contextdev.Int(0),
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

func TestAIExtractProductsWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.ExtractProducts(context.TODO(), contextdev.AIExtractProductsParams{
		OfByDomain: &contextdev.AIExtractProductsParamsBodyByDomain{
			Domain:      "domain",
			MaxAgeMs:    contextdev.Int(0),
			MaxProducts: contextdev.Int(1),
			TimeoutMs:   contextdev.Int(1000),
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
