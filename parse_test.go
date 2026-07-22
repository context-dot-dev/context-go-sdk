// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contextdev_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/context-dot-dev/context-go-sdk/v2"
	"github.com/context-dot-dev/context-go-sdk/v2/internal/testutil"
	"github.com/context-dot-dev/context-go-sdk/v2/option"
)

func TestParseHandleWithOptionalParams(t *testing.T) {
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
	_, err := client.Parse.Handle(
		context.TODO(),
		io.Reader(bytes.NewBuffer([]byte("Example data"))),
		contextdev.ParseHandleParams{
			Client:    contextdev.String("x"),
			Extension: contextdev.ParseHandleParamsExtensionTxt,
			IncludeImages: contextdev.ParseHandleParamsIncludeImagesUnion{
				OfParseHandlesIncludeImagesString: contextdev.String("true"),
			},
			IncludeLinks: contextdev.ParseHandleParamsIncludeLinksUnion{
				OfParseHandlesIncludeLinksString: contextdev.String("true"),
			},
			Ocr: contextdev.ParseHandleParamsOcrUnion{
				OfParseHandlesOcrString: contextdev.String("true"),
			},
			Pdf: contextdev.ParseHandleParamsPdf{
				End:   contextdev.Int(1),
				Start: contextdev.Int(1),
			},
			ShortenBase64Images: contextdev.ParseHandleParamsShortenBase64ImagesUnion{
				OfParseHandlesShortenBase64ImagesString: contextdev.String("true"),
			},
			Tags: []string{"production", "team-alpha"},
			UseMainContentOnly: contextdev.ParseHandleParamsUseMainContentOnlyUnion{
				OfParseHandlesUseMainContentOnlyString: contextdev.String("true"),
			},
			Zdr: contextdev.ParseHandleParamsZdrEnabled,
		},
	)
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
