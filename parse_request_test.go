package contextdev_test

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/context-dot-dev/context-go-sdk/v2"
	"github.com/context-dot-dev/context-go-sdk/v2/option"
)

func TestParseHandleRequest(t *testing.T) {
	const document = "<p>Example document</p>"
	tests := []struct {
		name   string
		params contextdev.ParseHandleParams
		opts   []option.RequestOption
		query  url.Values
	}{
		{
			name:  "no options",
			query: url.Values{},
		},
		{
			name: "document options",
			params: contextdev.ParseHandleParams{
				Extension:          contextdev.ParseHandleParamsExtensionHTML,
				IncludeLinks:       contextdev.Bool(true),
				UseMainContentOnly: contextdev.Bool(false),
				Tags:               []string{"preview", "team & tools"},
			},
			query: url.Values{
				"extension":          {"html"},
				"includeLinks":       {"true"},
				"useMainContentOnly": {"false"},
				"tags":               {"preview,team & tools"},
			},
		},
		{
			name: "custom request query takes precedence",
			params: contextdev.ParseHandleParams{
				Client: contextdev.String("default"),
				Pdf: contextdev.ParseHandleParamsPdf{
					Start: contextdev.Int(2),
					End:   contextdev.Int(5),
				},
			},
			opts: []option.RequestOption{option.WithQuery("client", "integration")},
			query: url.Values{
				"client":     {"integration"},
				"pdf[start]": {"2"},
				"pdf[end]":   {"5"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := contextdev.NewClient(
				option.WithAPIKey("offline-test"),
				option.WithBaseURL("https://context.invalid/v1/"),
				option.WithHTTPClient(&http.Client{Transport: &closureTransport{
					fn: func(req *http.Request) (*http.Response, error) {
						if req.Method != http.MethodPost || req.URL.Path != "/v1/parse" {
							t.Errorf("got %s %s, want POST /v1/parse", req.Method, req.URL.Path)
						}
						if got := req.URL.Query(); !reflect.DeepEqual(got, tt.query) {
							t.Errorf("query = %v, want %v", got, tt.query)
						}
						if got := req.Header.Get("Content-Type"); got != "application/octet-stream" {
							t.Errorf("content type = %q, want application/octet-stream", got)
						}
						body, err := io.ReadAll(req.Body)
						if err != nil {
							return nil, err
						}
						if got := string(body); got != document {
							t.Errorf("body = %q, want %q", got, document)
						}
						return &http.Response{
							StatusCode: http.StatusOK,
							Header:     http.Header{"Content-Type": {"application/json"}},
							Body:       io.NopCloser(strings.NewReader(`{"markdown":"Example document","success":true,"type":"html"}`)),
						}, nil
					},
				}}),
			)

			result, err := client.Parse.Handle(context.Background(), strings.NewReader(document), tt.params, tt.opts...)
			if err != nil {
				t.Fatal(err)
			}
			if result.Markdown != "Example document" {
				t.Errorf("markdown = %q, want Example document", result.Markdown)
			}
		})
	}
}
