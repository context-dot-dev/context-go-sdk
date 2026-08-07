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

func TestPersonEnrichWithOptionalParams(t *testing.T) {
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
	_, err := client.People.Enrich(context.TODO(), contextdev.PersonEnrichParams{
		Company: contextdev.PersonEnrichParamsCompany{
			Domain: contextdev.String("analyticalengines.example"),
			Name:   contextdev.String("Analytical Engines"),
		},
		Education: []contextdev.PersonEnrichParamsEducation{{
			Degree:         contextdev.String("x"),
			FieldOfStudy:   contextdev.String("x"),
			GraduationYear: contextdev.Int(1900),
			Institution: contextdev.PersonEnrichParamsEducationInstitution{
				Domain: contextdev.String("x"),
				Name:   contextdev.String("x"),
			},
		}},
		Email: contextdev.String("dev@stainless.com"),
		Location: contextdev.PersonEnrichParamsLocation{
			City:    contextdev.String("x"),
			Country: contextdev.String("x"),
			Region:  contextdev.String("x"),
		},
		Name: contextdev.PersonEnrichParamsName{
			First: contextdev.String("Ada"),
			Last:  contextdev.String("Lovelace"),
		},
		SocialURLs: []string{"https://www.linkedin.com/in/ada-lovelace/"},
		Tags:       []string{"production", "team-alpha"},
		TimeoutMs:  contextdev.Int(1000),
	})
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
