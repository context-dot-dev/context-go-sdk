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

// Search the full text of papers from arXiv and Europe PMC by topic or question.
// Each result includes the matching passage and a direct link to the full text.
// Page through up to 250 matches, ranked by relevance or newest first.
func (r *ScholarService) Search(ctx context.Context, body ScholarSearchParams, opts ...option.RequestOption) (res *ScholarSearchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "scholar/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ScholarSearchResponse struct {
	// Papers on this page, in the requested order.
	Data []ScholarSearchResponseData `json:"data" api:"required"`
	// Whether more results are available after this page.
	HasMore bool `json:"has_more" api:"required"`
	// Summary of this page.
	Meta ScholarSearchResponseMeta `json:"meta" api:"required"`
	// Pass as `cursor` to fetch the next page. `null` on the last page.
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
	// Unique ID of the paper.
	ID string `json:"id" api:"required"`
	// Format of the full text at `full_text_url`. `jats` is JATS XML.
	//
	// Any of "pdf", "html", "jats".
	FullTextFormat string `json:"full_text_format" api:"required"`
	// Direct URL of the full text.
	FullTextURL string `json:"full_text_url" api:"required" format:"uri"`
	// When the paper was published. `null` when unknown.
	PublishedAt time.Time `json:"published_at" api:"required" format:"date-time"`
	// Passage that matches the query, or the opening passage when the query is only a
	// paper ID.
	Snippet string `json:"snippet" api:"required"`
	// Title of the paper.
	Title string `json:"title" api:"required"`
	// URL of the paper's landing page.
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

// Summary of this page.
type ScholarSearchResponseMeta struct {
	// Number of papers on this page.
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
	// Credits your organization has left after this request.
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
	// Topic or question to search for. Append one `arxiv:`, `pmcid:`, `pmid:`, or
	// `doi:` ID to search within that paper, such as `protein folding pmid:34265844`,
	// or send the ID on its own.
	Query string `json:"query" api:"required"`
	// Cursor from the previous page's `next_cursor`. Omit for the first page.
	Cursor param.Opt[string] `json:"cursor,omitzero"`
	// Maximum number of papers to return per page.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// Filter by publication date.
	FilterBy ScholarSearchParamsFilterBy `json:"filterBy,omitzero"`
	// `relevance` puts the closest matches first; `date_desc` puts the newest papers
	// first, with undated papers last. Both orders apply to the top 250 matches.
	//
	// Any of "relevance", "date_desc".
	Sort ScholarSearchParamsSort `json:"sort,omitzero"`
	// Labels for this request. Filter usage and request logs by them in the dashboard
	// or with `GET /org/logs`.
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

// `relevance` puts the closest matches first; `date_desc` puts the newest papers
// first, with undated papers last. Both orders apply to the top 250 matches.
type ScholarSearchParamsSort string

const (
	ScholarSearchParamsSortRelevance ScholarSearchParamsSort = "relevance"
	ScholarSearchParamsSortDateDesc  ScholarSearchParamsSort = "date_desc"
)
