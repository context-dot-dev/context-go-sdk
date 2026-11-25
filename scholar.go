// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contextdev

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/context-dot-dev/context-go-sdk/v2/internal/apijson"
	"github.com/context-dot-dev/context-go-sdk/v2/internal/requestconfig"
	"github.com/context-dot-dev/context-go-sdk/v2/option"
	"github.com/context-dot-dev/context-go-sdk/v2/packages/param"
	"github.com/context-dot-dev/context-go-sdk/v2/packages/respjson"
)

// Search full-text academic papers from arXiv and Europe PMC.
//
// ScholarService contains methods and other services that help with interacting
// with the context.dev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScholarService] method instead.
type ScholarService struct {
	options []option.RequestOption
}

// NewScholarService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewScholarService(opts ...option.RequestOption) (r ScholarService) {
	r = ScholarService{}
	r.options = opts
	return
}

// Search full-text papers from arXiv and Europe PMC. Returns up to 250 relevant
// papers with matching passages and full-text links.
func (r *ScholarService) Search(ctx context.Context, body ScholarSearchParams, opts ...option.RequestOption) (res *ScholarSearchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "scholar/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ScholarSearchResponse struct {
	// Matching papers in the requested order.
	Data []ScholarSearchResponseData `json:"data" api:"required"`
	// Whether another page is available.
	HasMore bool `json:"has_more" api:"required"`
	// Page summary.
	Meta ScholarSearchResponseMeta `json:"meta" api:"required"`
	// Cursor for the next page, or null on the last page.
	NextCursor string `json:"next_cursor" api:"required"`
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Credits this request used and your remaining balance.
	KeyMetadata ScholarSearchResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		Meta        respjson.Field
		NextCursor  respjson.Field
		RequestID   respjson.Field
		KeyMetadata respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScholarSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *ScholarSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScholarSearchResponseData struct {
	// Stable paper identifier.
	ID string `json:"id" api:"required"`
	// Format of the linked full text. jats is XML.
	//
	// Any of "pdf", "html", "jats".
	FullTextFormat string `json:"full_text_format" api:"required"`
	// Direct link to the full text.
	FullTextURL string `json:"full_text_url" api:"required" format:"uri"`
	// Publication time in ISO 8601, or null when unknown.
	PublishedAt time.Time `json:"published_at" api:"required" format:"date-time"`
	// Matching passage, or opening passage for an ID-only query.
	Snippet string `json:"snippet" api:"required"`
	// Paper title.
	Title string `json:"title" api:"required"`
	// Paper landing page.
	URL string `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		FullTextFormat respjson.Field
		FullTextURL    respjson.Field
		PublishedAt    respjson.Field
		Snippet        respjson.Field
		Title          respjson.Field
		URL            respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScholarSearchResponseData) RawJSON() string { return r.JSON.raw }
func (r *ScholarSearchResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Page summary.
type ScholarSearchResponseMeta struct {
	// Papers in this page.
	Count int64 `json:"count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScholarSearchResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *ScholarSearchResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credits this request used and your remaining balance.
type ScholarSearchResponseKeyMetadata struct {
	// Credits charged for this request.
	CreditsConsumed int64 `json:"credits_consumed" api:"required"`
	// Credits remaining for your organization.
	CreditsRemaining int64 `json:"credits_remaining" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditsConsumed  respjson.Field
		CreditsRemaining respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScholarSearchResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *ScholarSearchResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScholarSearchParams struct {
	// A topic or question. Append one arxiv:, pmcid:, pmid:, or doi: ID to search
	// within an indexed paper.
	Query string `json:"query" api:"required"`
	// Use next_cursor to fetch the next page.
	Cursor param.Opt[string] `json:"cursor,omitzero"`
	// Papers per page.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// Filter by publication date.
	FilterBy ScholarSearchParamsFilterBy `json:"filterBy,omitzero"`
	// Order the top 250 matches. date_desc puts newest first, with undated papers
	// last.
	//
	// Any of "relevance", "date_desc".
	Sort ScholarSearchParamsSort `json:"sort,omitzero"`
	// Labels for filtering usage in the dashboard.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r ScholarSearchParams) MarshalJSON() (data []byte, err error) {
	type shadow ScholarSearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScholarSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter by publication date.
type ScholarSearchParamsFilterBy struct {
	// Publication date range. Undated papers are excluded.
	Date ScholarSearchParamsFilterByDate `json:"date,omitzero"`
	paramObj
}

func (r ScholarSearchParamsFilterBy) MarshalJSON() (data []byte, err error) {
	type shadow ScholarSearchParamsFilterBy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScholarSearchParamsFilterBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Publication date range. Undated papers are excluded.
type ScholarSearchParamsFilterByDate struct {
	// Inclusive start date, in Unix milliseconds.
	From param.Opt[int64] `json:"from,omitzero"`
	// Inclusive end date, in Unix milliseconds.
	To param.Opt[int64] `json:"to,omitzero"`
	paramObj
}

func (r ScholarSearchParamsFilterByDate) MarshalJSON() (data []byte, err error) {
	type shadow ScholarSearchParamsFilterByDate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScholarSearchParamsFilterByDate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Order the top 250 matches. date_desc puts newest first, with undated papers
// last.
type ScholarSearchParamsSort string

const (
	ScholarSearchParamsSortRelevance ScholarSearchParamsSort = "relevance"
	ScholarSearchParamsSortDateDesc  ScholarSearchParamsSort = "date_desc"
)
