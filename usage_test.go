// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contextdev_test

import (
	"context"
	"os"
	"testing"

	"github.com/context-dot-dev/context-go-sdk"
	"github.com/context-dot-dev/context-go-sdk/internal/testutil"
	"github.com/context-dot-dev/context-go-sdk/option"
)

func TestUsage(t *testing.T) {
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
	response, err := client.Web.Extract(context.TODO(), contextdev.WebExtractParams{
		Schema: map[string]any{
			"type":                 "bar",
			"properties":           "bar",
			"required":             "bar",
			"additionalProperties": "bar",
		},
		URL: "https://example.com",
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response.Data)
}
