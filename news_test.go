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

func TestNewsSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.News.Search(context.TODO(), contextdev.NewsSearchParams{
		SearchBy: contextdev.NewsSearchParamsSearchBy{
			Entity: contextdev.NewsSearchParamsSearchByEntityUnion{
				OfName: &contextdev.NewsSearchParamsSearchByEntityName{
					Name: "xx",
				},
			},
			Type: "entity",
		},
		Cursor: contextdev.String("cursor"),
		FilterBy: contextdev.NewsSearchParamsFilterBy{
			ArticleLanguage: []string{"ar"},
			ArticleType:     []string{"editorial"},
			Date: contextdev.NewsSearchParamsFilterByDate{
				From: contextdev.Int(0),
				To:   contextdev.Int(0),
			},
			SourceCountry: []string{"ae"},
			SourceDomain:  []string{"x"},
		},
		Limit: contextdev.Int(1),
		SortBy: contextdev.NewsSearchParamsSortBy{
			Type: "relevance",
		},
		Tags: []string{"production", "team-alpha"},
	})
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
