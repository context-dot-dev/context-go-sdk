// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contextdev

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/context-dot-dev/context-go-sdk/v2/internal/apijson"
	"github.com/context-dot-dev/context-go-sdk/v2/internal/apiquery"
	"github.com/context-dot-dev/context-go-sdk/v2/internal/requestconfig"
	"github.com/context-dot-dev/context-go-sdk/v2/option"
	"github.com/context-dot-dev/context-go-sdk/v2/packages/param"
	"github.com/context-dot-dev/context-go-sdk/v2/packages/respjson"
	"github.com/context-dot-dev/context-go-sdk/v2/shared/constant"
)

// BatchService contains methods and other services that help with interacting with
// the context.dev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBatchService] method instead.
type BatchService struct {
	options []option.RequestOption
}

// NewBatchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBatchService(opts ...option.RequestOption) (r BatchService) {
	r = BatchService{}
	r.options = opts
	return
}

// Check progress and get download links when the batch finishes. Also returns the
// rejected-URL list from submission. The webhook signing secret is not repeated
// here — it is returned once, by the submit response.
func (r *BatchService) Get(ctx context.Context, batchID string, opts ...option.RequestOption) (res *BatchGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if batchID == "" {
		err = errors.New("missing required batch_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("batch/%s", url.PathEscape(batchID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List your batches from newest to oldest. Filter by status or continue with a
// cursor.
func (r *BatchService) List(ctx context.Context, query BatchListParams, opts ...option.RequestOption) (res *BatchListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "batch/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Stop a batch from starting new pages. In-progress pages finish, and unused
// credits are refunded.
func (r *BatchService) Cancel(ctx context.Context, batchID string, opts ...option.RequestOption) (res *BatchCancelResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if batchID == "" {
		err = errors.New("missing required batch_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("batch/%s/cancel", url.PathEscape(batchID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Page through the result records of a finished batch as JSON, in the same order
// as the downloadable result files. Use this instead of downloading and parsing
// the NDJSON files yourself.
func (r *BatchService) GetResults(ctx context.Context, batchID string, query BatchGetResultsParams, opts ...option.RequestOption) (res *BatchGetResultsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if batchID == "" {
		err = errors.New("missing required batch_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("batch/%s/results", url.PathEscape(batchID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve and normalize a person profile from identifiers.
func (r *BatchService) Submit(ctx context.Context, body BatchSubmitParams, opts ...option.RequestOption) (res *BatchSubmitResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "people/retrieve"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Page failures sharing one error code.
type PageErrorCount struct {
	// Error code for these failures.
	Code string `json:"code" api:"required"`
	// Pages that failed with this code.
	Count int64 `json:"count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Count       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PageErrorCount) RawJSON() string { return r.JSON.raw }
func (r *PageErrorCount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A failure of the batch as a whole, distinct from the per-page failures in
// `page_errors`.
type Failure struct {
	// Why the batch itself stopped.
	Code string `json:"code" api:"required"`
	// Human-readable explanation.
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Failure) RawJSON() string { return r.JSON.raw }
func (r *Failure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The crawl controls as submitted, so the limits requested can be compared against
// what the crawl reached.
type CrawlControls struct {
	// Whether links to subdomains were followed. Always false for a sitemap crawl.
	FollowSubdomains bool `json:"follow_subdomains" api:"required"`
	// Link depth limit. Always 0 for a sitemap crawl, which never follows links off
	// its URLs; null when a `start_url` crawl set no limit.
	MaxDepth int64 `json:"max_depth" api:"required"`
	// The `maxUrls` submitted with the crawl. A sitemap crawl scrapes only the URLs
	// its sitemap actually lists, up to this many, so `input.reserved` is often lower.
	MaxPages int64 `json:"max_pages" api:"required"`
	// Where the crawl started.
	Source CrawlControlsSourceUnion `json:"source" api:"required"`
	// RE2 pattern URLs had to match to be crawled. Null when the crawl set none.
	URLPattern string `json:"url_pattern" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FollowSubdomains respjson.Field
		MaxDepth         respjson.Field
		MaxPages         respjson.Field
		Source           respjson.Field
		URLPattern       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CrawlControls) RawJSON() string { return r.JSON.raw }
func (r *CrawlControls) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CrawlControlsSourceUnion contains all possible properties and values from
// [CrawlControlsSourceObject], [CrawlControlsSourceObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CrawlControlsSourceUnion struct {
	Type string `json:"type"`
	// This field is from variant [CrawlControlsSourceObject].
	URL string `json:"url"`
	// This field is from variant [CrawlControlsSourceObject2].
	Domain string `json:"domain"`
	JSON   struct {
		Type   respjson.Field
		URL    respjson.Field
		Domain respjson.Field
		raw    string
	} `json:"-"`
}

func (u CrawlControlsSourceUnion) AsCrawlControlsSourceObject() (v CrawlControlsSourceObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CrawlControlsSourceUnion) AsCrawlControlsSourceObject2() (v CrawlControlsSourceObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CrawlControlsSourceUnion) RawJSON() string { return u.JSON.raw }

func (r *CrawlControlsSourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlControlsSourceObject struct {
	// Any of "start_url".
	Type string `json:"type" api:"required"`
	// Page the crawl started from.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CrawlControlsSourceObject) RawJSON() string { return r.JSON.raw }
func (r *CrawlControlsSourceObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlControlsSourceObject2 struct {
	// Domain whose sitemap supplied the pages.
	Domain string `json:"domain" api:"required"`
	// Any of "sitemap".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CrawlControlsSourceObject2) RawJSON() string { return r.JSON.raw }
func (r *CrawlControlsSourceObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What submission took in, and what it charged for.
type Intake struct {
	// URLs dropped before reserving because another entry resolved to the same page.
	// Non-zero for sitemap crawls too, whose sitemaps routinely list a page more than
	// once.
	Duplicates int64 `json:"duplicates" api:"required"`
	// URLs from your list rejected as unusable; the same ones are itemised in
	// `invalid_urls` at submission. Null for a crawl — a crawl that resolves no usable
	// page is rejected outright with a 400 rather than accepted with an empty list.
	Invalid int64 `json:"invalid" api:"required"`
	// Pages credits were reserved for. Everything else — progress, the refund, the
	// completion percentage — is measured against this.
	Reserved int64 `json:"reserved" api:"required"`
	// Whether `reserved` is an upper bound the batch may finish under. True only for a
	// crawl that follows links, whose reachable page count is unknowable until it
	// runs. False for a scrape and for a sitemap crawl, where `reserved` is an exact
	// page count.
	ReservedIsCeiling bool `json:"reserved_is_ceiling" api:"required"`
	// URLs in the list you sent, before validation and de-duplication. Null for a
	// crawl, which is given a source rather than a list.
	Submitted int64 `json:"submitted" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Duplicates        respjson.Field
		Invalid           respjson.Field
		Reserved          respjson.Field
		ReservedIsCeiling respjson.Field
		Submitted         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Intake) RawJSON() string { return r.JSON.raw }
func (r *Intake) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchGetResponse struct {
	// Batch ID used to retrieve or cancel the job.
	ID string `json:"id" api:"required"`
	// The crawl controls as submitted, so the limits requested can be compared against
	// what the crawl reached.
	Crawl CrawlControls `json:"crawl" api:"required"`
	// What this batch has done to your credit balance.
	Credits BatchGetResponseCredits `json:"credits" api:"required"`
	// A failure of the batch as a whole, distinct from the per-page failures in
	// `page_errors`.
	Failure Failure `json:"failure" api:"required"`
	// What each page is returned as. Matches `input.data.format` on the submit
	// request.
	//
	// Any of "markdown", "html".
	Format BatchGetResponseFormat `json:"format" api:"required"`
	// What submission took in, and what it charged for.
	Input Intake `json:"input" api:"required"`
	// Rejected URLs, up to 100. These are not charged.
	InvalidURLs []BatchGetResponseInvalidURL `json:"invalid_urls" api:"required"`
	// How pages were selected. Matches `input.mode` on the submit request.
	//
	// Any of "scrape", "crawl".
	Mode BatchGetResponseMode `json:"mode" api:"required"`
	// Individual page failures grouped by error code, sorted by count. Unrelated to
	// `failure`, which is the batch itself failing.
	PageErrors []PageErrorCount `json:"page_errors" api:"required"`
	// Pages attempted so far. Use `status` to check completion.
	Progress BatchGetResponseProgress `json:"progress" api:"required"`
	// Download links, available once the batch reaches a final status and null before
	// then. GET /batch/{batch_id}/results serves the same records as paginated JSON.
	Results BatchGetResponseResults `json:"results" api:"required"`
	// Current state. `completed`, `cancelled`, and `failed` are final.
	//
	// Any of "queued", "running", "cancelling", "completed", "cancelled", "failed".
	Status BatchGetResponseStatus `json:"status" api:"required"`
	// Tags stored on the batch at submission.
	Tags   []string               `json:"tags" api:"required"`
	Timing BatchGetResponseTiming `json:"timing" api:"required"`
	// API key usage for this request.
	KeyMetadata BatchGetResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Crawl       respjson.Field
		Credits     respjson.Field
		Failure     respjson.Field
		Format      respjson.Field
		Input       respjson.Field
		InvalidURLs respjson.Field
		Mode        respjson.Field
		PageErrors  respjson.Field
		Progress    respjson.Field
		Results     respjson.Field
		Status      respjson.Field
		Tags        respjson.Field
		Timing      respjson.Field
		KeyMetadata respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What this batch has done to your credit balance.
type BatchGetResponseCredits struct {
	// `reserved` minus `refunded` — what the batch has cost so far. Equal to
	// `reserved` until the batch settles.
	Net int64 `json:"net" api:"required"`
	// Credits returned for pages that did not succeed. Stays 0 until the batch reaches
	// a final status, then settles in one movement.
	Refunded int64 `json:"refunded" api:"required"`
	// Credits debited from your balance the moment the batch was accepted. This is a
	// charge, not a forecast — the whole amount leaves the balance up front.
	Reserved int64 `json:"reserved" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Net         respjson.Field
		Refunded    respjson.Field
		Reserved    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResponseCredits) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResponseCredits) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What each page is returned as. Matches `input.data.format` on the submit
// request.
type BatchGetResponseFormat string

const (
	BatchGetResponseFormatMarkdown BatchGetResponseFormat = "markdown"
	BatchGetResponseFormatHTML     BatchGetResponseFormat = "html"
)

type BatchGetResponseInvalidURL struct {
	// Why it was rejected.
	Reason string `json:"reason" api:"required"`
	// Rejected URL.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reason      respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResponseInvalidURL) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResponseInvalidURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How pages were selected. Matches `input.mode` on the submit request.
type BatchGetResponseMode string

const (
	BatchGetResponseModeScrape BatchGetResponseMode = "scrape"
	BatchGetResponseModeCrawl  BatchGetResponseMode = "crawl"
)

// Pages attempted so far. Use `status` to check completion.
type BatchGetResponseProgress struct {
	// Pages that could not be scraped.
	Failed int64 `json:"failed" api:"required"`
	// Reserved pages not yet attempted. A cancelled batch keeps reporting the URLs it
	// never reached; a crawl whose `input.reserved_is_ceiling` is true reports 0 once
	// final, because its unspent budget was never real pages.
	Pending int64 `json:"pending" api:"required"`
	// Pages scraped successfully.
	Succeeded int64 `json:"succeeded" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Failed      respjson.Field
		Pending     respjson.Field
		Succeeded   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResponseProgress) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResponseProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Download links, available once the batch reaches a final status and null before
// then. GET /batch/{batch_id}/results serves the same records as paginated JSON.
type BatchGetResponseResults struct {
	// When the download URLs expire.
	ExpiresAt string `json:"expires_at" api:"required"`
	// Result files. Order is not guaranteed.
	Files []BatchGetResponseResultsFile `json:"files" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		Files       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResponseResults) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResponseResults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchGetResponseResultsFile struct {
	// Compressed file size in bytes.
	Bytes int64 `json:"bytes" api:"required"`
	// Results in this file.
	Items int64 `json:"items" api:"required"`
	// Temporary URL for a gzipped NDJSON file.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bytes       respjson.Field
		Items       respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResponseResultsFile) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResponseResultsFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current state. `completed`, `cancelled`, and `failed` are final.
type BatchGetResponseStatus string

const (
	BatchGetResponseStatusQueued     BatchGetResponseStatus = "queued"
	BatchGetResponseStatusRunning    BatchGetResponseStatus = "running"
	BatchGetResponseStatusCancelling BatchGetResponseStatus = "cancelling"
	BatchGetResponseStatusCompleted  BatchGetResponseStatus = "completed"
	BatchGetResponseStatusCancelled  BatchGetResponseStatus = "cancelled"
	BatchGetResponseStatusFailed     BatchGetResponseStatus = "failed"
)

type BatchGetResponseTiming struct {
	// When processing finished. Null while active.
	CompletedAt string `json:"completed_at" api:"required"`
	// When the batch was created.
	CreatedAt string `json:"created_at" api:"required"`
	// When processing started. Null while queued.
	StartedAt string `json:"started_at" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		CreatedAt   respjson.Field
		StartedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResponseTiming) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResponseTiming) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// API key usage for this request.
type BatchGetResponseKeyMetadata struct {
	// The number of credits consumed by this request.
	CreditsConsumed int64 `json:"credits_consumed" api:"required"`
	// The number of credits remaining for your organization after this request.
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
func (r BatchGetResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchListResponse struct {
	// Batches on this page.
	Data []BatchListResponseData `json:"data"`
	// Whether another page is available.
	HasMore bool `json:"has_more"`
	// Metadata about the API key used for the request. Included in every response
	// whenever a valid API key is provided, even when the response status is not 200.
	KeyMetadata BatchListResponseKeyMetadata `json:"key_metadata"`
	// Cursor for the next page.
	NextCursor string `json:"next_cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		KeyMetadata respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchListResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An asynchronous web scraping job.
type BatchListResponseData struct {
	// Batch ID used to retrieve or cancel the job.
	ID string `json:"id" api:"required"`
	// The crawl controls as submitted, so the limits requested can be compared against
	// what the crawl reached.
	Crawl CrawlControls `json:"crawl" api:"required"`
	// What this batch has done to your credit balance.
	Credits BatchListResponseDataCredits `json:"credits" api:"required"`
	// A failure of the batch as a whole, distinct from the per-page failures in
	// `page_errors`.
	Failure Failure `json:"failure" api:"required"`
	// What each page is returned as. Matches `input.data.format` on the submit
	// request.
	//
	// Any of "markdown", "html".
	Format string `json:"format" api:"required"`
	// What submission took in, and what it charged for.
	Input Intake `json:"input" api:"required"`
	// How pages were selected. Matches `input.mode` on the submit request.
	//
	// Any of "scrape", "crawl".
	Mode string `json:"mode" api:"required"`
	// Individual page failures grouped by error code, sorted by count. Unrelated to
	// `failure`, which is the batch itself failing.
	PageErrors []PageErrorCount `json:"page_errors" api:"required"`
	// Pages attempted so far. Use `status` to check completion.
	Progress BatchListResponseDataProgress `json:"progress" api:"required"`
	// Download links, available once the batch reaches a final status and null before
	// then. GET /batch/{batch_id}/results serves the same records as paginated JSON.
	Results BatchListResponseDataResults `json:"results" api:"required"`
	// Current state. `completed`, `cancelled`, and `failed` are final.
	//
	// Any of "queued", "running", "cancelling", "completed", "cancelled", "failed".
	Status string `json:"status" api:"required"`
	// Tags stored on the batch at submission.
	Tags   []string                    `json:"tags" api:"required"`
	Timing BatchListResponseDataTiming `json:"timing" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Crawl       respjson.Field
		Credits     respjson.Field
		Failure     respjson.Field
		Format      respjson.Field
		Input       respjson.Field
		Mode        respjson.Field
		PageErrors  respjson.Field
		Progress    respjson.Field
		Results     respjson.Field
		Status      respjson.Field
		Tags        respjson.Field
		Timing      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchListResponseData) RawJSON() string { return r.JSON.raw }
func (r *BatchListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What this batch has done to your credit balance.
type BatchListResponseDataCredits struct {
	// `reserved` minus `refunded` — what the batch has cost so far. Equal to
	// `reserved` until the batch settles.
	Net int64 `json:"net" api:"required"`
	// Credits returned for pages that did not succeed. Stays 0 until the batch reaches
	// a final status, then settles in one movement.
	Refunded int64 `json:"refunded" api:"required"`
	// Credits debited from your balance the moment the batch was accepted. This is a
	// charge, not a forecast — the whole amount leaves the balance up front.
	Reserved int64 `json:"reserved" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Net         respjson.Field
		Refunded    respjson.Field
		Reserved    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchListResponseDataCredits) RawJSON() string { return r.JSON.raw }
func (r *BatchListResponseDataCredits) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pages attempted so far. Use `status` to check completion.
type BatchListResponseDataProgress struct {
	// Pages that could not be scraped.
	Failed int64 `json:"failed" api:"required"`
	// Reserved pages not yet attempted. A cancelled batch keeps reporting the URLs it
	// never reached; a crawl whose `input.reserved_is_ceiling` is true reports 0 once
	// final, because its unspent budget was never real pages.
	Pending int64 `json:"pending" api:"required"`
	// Pages scraped successfully.
	Succeeded int64 `json:"succeeded" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Failed      respjson.Field
		Pending     respjson.Field
		Succeeded   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchListResponseDataProgress) RawJSON() string { return r.JSON.raw }
func (r *BatchListResponseDataProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Download links, available once the batch reaches a final status and null before
// then. GET /batch/{batch_id}/results serves the same records as paginated JSON.
type BatchListResponseDataResults struct {
	// When the download URLs expire.
	ExpiresAt string `json:"expires_at" api:"required"`
	// Result files. Order is not guaranteed.
	Files []BatchListResponseDataResultsFile `json:"files" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		Files       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchListResponseDataResults) RawJSON() string { return r.JSON.raw }
func (r *BatchListResponseDataResults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchListResponseDataResultsFile struct {
	// Compressed file size in bytes.
	Bytes int64 `json:"bytes" api:"required"`
	// Results in this file.
	Items int64 `json:"items" api:"required"`
	// Temporary URL for a gzipped NDJSON file.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bytes       respjson.Field
		Items       respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchListResponseDataResultsFile) RawJSON() string { return r.JSON.raw }
func (r *BatchListResponseDataResultsFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchListResponseDataTiming struct {
	// When processing finished. Null while active.
	CompletedAt string `json:"completed_at" api:"required"`
	// When the batch was created.
	CreatedAt string `json:"created_at" api:"required"`
	// When processing started. Null while queued.
	StartedAt string `json:"started_at" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		CreatedAt   respjson.Field
		StartedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchListResponseDataTiming) RawJSON() string { return r.JSON.raw }
func (r *BatchListResponseDataTiming) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the API key used for the request. Included in every response
// whenever a valid API key is provided, even when the response status is not 200.
type BatchListResponseKeyMetadata struct {
	// The number of credits consumed by this request.
	CreditsConsumed int64 `json:"credits_consumed" api:"required"`
	// The number of credits remaining for your organization after this request.
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
func (r BatchListResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *BatchListResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchCancelResponse struct {
	// Batch ID used to retrieve or cancel the job.
	ID string `json:"id" api:"required"`
	// The crawl controls as submitted, so the limits requested can be compared against
	// what the crawl reached.
	Crawl CrawlControls `json:"crawl" api:"required"`
	// What this batch has done to your credit balance.
	Credits BatchCancelResponseCredits `json:"credits" api:"required"`
	// A failure of the batch as a whole, distinct from the per-page failures in
	// `page_errors`.
	Failure Failure `json:"failure" api:"required"`
	// What each page is returned as. Matches `input.data.format` on the submit
	// request.
	//
	// Any of "markdown", "html".
	Format BatchCancelResponseFormat `json:"format" api:"required"`
	// What submission took in, and what it charged for.
	Input Intake `json:"input" api:"required"`
	// How pages were selected. Matches `input.mode` on the submit request.
	//
	// Any of "scrape", "crawl".
	Mode BatchCancelResponseMode `json:"mode" api:"required"`
	// Individual page failures grouped by error code, sorted by count. Unrelated to
	// `failure`, which is the batch itself failing.
	PageErrors []PageErrorCount `json:"page_errors" api:"required"`
	// Pages attempted so far. Use `status` to check completion.
	Progress BatchCancelResponseProgress `json:"progress" api:"required"`
	// Download links, available once the batch reaches a final status and null before
	// then. GET /batch/{batch_id}/results serves the same records as paginated JSON.
	Results BatchCancelResponseResults `json:"results" api:"required"`
	// Current state. `completed`, `cancelled`, and `failed` are final.
	//
	// Any of "queued", "running", "cancelling", "completed", "cancelled", "failed".
	Status BatchCancelResponseStatus `json:"status" api:"required"`
	// Tags stored on the batch at submission.
	Tags   []string                  `json:"tags" api:"required"`
	Timing BatchCancelResponseTiming `json:"timing" api:"required"`
	// API key usage for this request.
	KeyMetadata BatchCancelResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Crawl       respjson.Field
		Credits     respjson.Field
		Failure     respjson.Field
		Format      respjson.Field
		Input       respjson.Field
		Mode        respjson.Field
		PageErrors  respjson.Field
		Progress    respjson.Field
		Results     respjson.Field
		Status      respjson.Field
		Tags        respjson.Field
		Timing      respjson.Field
		KeyMetadata respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchCancelResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchCancelResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What this batch has done to your credit balance.
type BatchCancelResponseCredits struct {
	// `reserved` minus `refunded` — what the batch has cost so far. Equal to
	// `reserved` until the batch settles.
	Net int64 `json:"net" api:"required"`
	// Credits returned for pages that did not succeed. Stays 0 until the batch reaches
	// a final status, then settles in one movement.
	Refunded int64 `json:"refunded" api:"required"`
	// Credits debited from your balance the moment the batch was accepted. This is a
	// charge, not a forecast — the whole amount leaves the balance up front.
	Reserved int64 `json:"reserved" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Net         respjson.Field
		Refunded    respjson.Field
		Reserved    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchCancelResponseCredits) RawJSON() string { return r.JSON.raw }
func (r *BatchCancelResponseCredits) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What each page is returned as. Matches `input.data.format` on the submit
// request.
type BatchCancelResponseFormat string

const (
	BatchCancelResponseFormatMarkdown BatchCancelResponseFormat = "markdown"
	BatchCancelResponseFormatHTML     BatchCancelResponseFormat = "html"
)

// How pages were selected. Matches `input.mode` on the submit request.
type BatchCancelResponseMode string

const (
	BatchCancelResponseModeScrape BatchCancelResponseMode = "scrape"
	BatchCancelResponseModeCrawl  BatchCancelResponseMode = "crawl"
)

// Pages attempted so far. Use `status` to check completion.
type BatchCancelResponseProgress struct {
	// Pages that could not be scraped.
	Failed int64 `json:"failed" api:"required"`
	// Reserved pages not yet attempted. A cancelled batch keeps reporting the URLs it
	// never reached; a crawl whose `input.reserved_is_ceiling` is true reports 0 once
	// final, because its unspent budget was never real pages.
	Pending int64 `json:"pending" api:"required"`
	// Pages scraped successfully.
	Succeeded int64 `json:"succeeded" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Failed      respjson.Field
		Pending     respjson.Field
		Succeeded   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchCancelResponseProgress) RawJSON() string { return r.JSON.raw }
func (r *BatchCancelResponseProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Download links, available once the batch reaches a final status and null before
// then. GET /batch/{batch_id}/results serves the same records as paginated JSON.
type BatchCancelResponseResults struct {
	// When the download URLs expire.
	ExpiresAt string `json:"expires_at" api:"required"`
	// Result files. Order is not guaranteed.
	Files []BatchCancelResponseResultsFile `json:"files" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		Files       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchCancelResponseResults) RawJSON() string { return r.JSON.raw }
func (r *BatchCancelResponseResults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchCancelResponseResultsFile struct {
	// Compressed file size in bytes.
	Bytes int64 `json:"bytes" api:"required"`
	// Results in this file.
	Items int64 `json:"items" api:"required"`
	// Temporary URL for a gzipped NDJSON file.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bytes       respjson.Field
		Items       respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchCancelResponseResultsFile) RawJSON() string { return r.JSON.raw }
func (r *BatchCancelResponseResultsFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current state. `completed`, `cancelled`, and `failed` are final.
type BatchCancelResponseStatus string

const (
	BatchCancelResponseStatusQueued     BatchCancelResponseStatus = "queued"
	BatchCancelResponseStatusRunning    BatchCancelResponseStatus = "running"
	BatchCancelResponseStatusCancelling BatchCancelResponseStatus = "cancelling"
	BatchCancelResponseStatusCompleted  BatchCancelResponseStatus = "completed"
	BatchCancelResponseStatusCancelled  BatchCancelResponseStatus = "cancelled"
	BatchCancelResponseStatusFailed     BatchCancelResponseStatus = "failed"
)

type BatchCancelResponseTiming struct {
	// When processing finished. Null while active.
	CompletedAt string `json:"completed_at" api:"required"`
	// When the batch was created.
	CreatedAt string `json:"created_at" api:"required"`
	// When processing started. Null while queued.
	StartedAt string `json:"started_at" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		CreatedAt   respjson.Field
		StartedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchCancelResponseTiming) RawJSON() string { return r.JSON.raw }
func (r *BatchCancelResponseTiming) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// API key usage for this request.
type BatchCancelResponseKeyMetadata struct {
	// The number of credits consumed by this request.
	CreditsConsumed int64 `json:"credits_consumed" api:"required"`
	// The number of credits remaining for your organization after this request.
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
func (r BatchCancelResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *BatchCancelResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchGetResultsResponse struct {
	// Result records on this page.
	Data []BatchGetResultsResponseDataUnion `json:"data"`
	// Whether another page is available.
	HasMore bool `json:"has_more"`
	// Metadata about the API key used for the request. Included in every response
	// whenever a valid API key is provided, even when the response status is not 200.
	KeyMetadata BatchGetResultsResponseKeyMetadata `json:"key_metadata"`
	// Cursor for the next page.
	NextCursor string `json:"next_cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		KeyMetadata respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchGetResultsResponseDataUnion contains all possible properties and values
// from [BatchGetResultsResponseDataOk], [BatchGetResultsResponseDataError].
//
// Use the [BatchGetResultsResponseDataUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type BatchGetResultsResponseDataUnion struct {
	// This field is from variant [BatchGetResultsResponseDataOk].
	FinalURL string `json:"final_url"`
	// This field is from variant [BatchGetResultsResponseDataOk].
	HTTPStatus int64 `json:"http_status"`
	// This field is from variant [BatchGetResultsResponseDataOk].
	Metadata BatchGetResultsResponseDataOkMetadata `json:"metadata"`
	// Any of "ok", "error".
	Status string `json:"status"`
	URL    string `json:"url"`
	// This field is from variant [BatchGetResultsResponseDataOk].
	HTML   string `json:"html"`
	ItemID string `json:"itemId"`
	// This field is from variant [BatchGetResultsResponseDataOk].
	Markdown string `json:"markdown"`
	Meta     any    `json:"meta"`
	// This field is from variant [BatchGetResultsResponseDataError].
	ErrorCode string `json:"error_code"`
	// This field is from variant [BatchGetResultsResponseDataError].
	Message string `json:"message"`
	JSON    struct {
		FinalURL   respjson.Field
		HTTPStatus respjson.Field
		Metadata   respjson.Field
		Status     respjson.Field
		URL        respjson.Field
		HTML       respjson.Field
		ItemID     respjson.Field
		Markdown   respjson.Field
		Meta       respjson.Field
		ErrorCode  respjson.Field
		Message    respjson.Field
		raw        string
	} `json:"-"`
}

// anyBatchGetResultsResponseData is implemented by each variant of
// [BatchGetResultsResponseDataUnion] to add type safety for the return type of
// [BatchGetResultsResponseDataUnion.AsAny]
type anyBatchGetResultsResponseData interface {
	implBatchGetResultsResponseDataUnion()
}

func (BatchGetResultsResponseDataOk) implBatchGetResultsResponseDataUnion()    {}
func (BatchGetResultsResponseDataError) implBatchGetResultsResponseDataUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := BatchGetResultsResponseDataUnion.AsAny().(type) {
//	case contextdev.BatchGetResultsResponseDataOk:
//	case contextdev.BatchGetResultsResponseDataError:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u BatchGetResultsResponseDataUnion) AsAny() anyBatchGetResultsResponseData {
	switch u.Status {
	case "ok":
		return u.AsOk()
	case "error":
		return u.AsError()
	}
	return nil
}

func (u BatchGetResultsResponseDataUnion) AsOk() (v BatchGetResultsResponseDataOk) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BatchGetResultsResponseDataUnion) AsError() (v BatchGetResultsResponseDataError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BatchGetResultsResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *BatchGetResultsResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A page the batch fetched successfully.
type BatchGetResultsResponseDataOk struct {
	// URL the content was read from, after redirects.
	FinalURL string `json:"final_url" api:"required"`
	// HTTP status of the final response, when known.
	HTTPStatus int64 `json:"http_status" api:"required"`
	// Metadata extracted from the scraped page HTML.
	Metadata BatchGetResultsResponseDataOkMetadata `json:"metadata" api:"required"`
	// The page was scraped.
	Status constant.Ok `json:"status" default:"ok"`
	// URL as submitted, or as discovered by the crawl.
	URL string `json:"url" api:"required"`
	// Raw page HTML. Present on html batches.
	HTML string `json:"html"`
	// Caller-supplied identifier echoed from submission.
	ItemID string `json:"itemId"`
	// Page content as Markdown. Present on markdown batches.
	Markdown string `json:"markdown"`
	// Caller-supplied metadata echoed from submission.
	Meta map[string]any `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FinalURL    respjson.Field
		HTTPStatus  respjson.Field
		Metadata    respjson.Field
		Status      respjson.Field
		URL         respjson.Field
		HTML        respjson.Field
		ItemID      respjson.Field
		Markdown    respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataOk) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataOk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata extracted from the scraped page HTML.
type BatchGetResultsResponseDataOkMetadata struct {
	// Final URL scraped after redirects or scraper fallback, when known. Falls back to
	// sourceUrl when unavailable.
	FinalURL string `json:"finalUrl" api:"required"`
	// Original URL requested by the caller.
	SourceURL string `json:"sourceUrl" api:"required"`
	// Additional non-social meta tags not promoted to top-level metadata fields.
	AdditionalMeta map[string]BatchGetResultsResponseDataOkMetadataAdditionalMetaUnion `json:"additionalMeta"`
	// Resolved alternate links from link rel=alternate tags.
	Alternates []BatchGetResultsResponseDataOkMetadataAlternate `json:"alternates"`
	// Author metadata, when present.
	Author string `json:"author"`
	// Resolved canonical URL, when present.
	CanonicalURL string `json:"canonicalUrl"`
	// Best description extracted from standard, Open Graph, or Twitter metadata.
	Description string `json:"description"`
	// Resolved favicon URL, when present.
	Favicon string `json:"favicon"`
	// Primary resolved preview image from Open Graph, Twitter, or image metadata.
	Image string `json:"image"`
	// JSON-LD structured data blocks parsed from the page.
	JsonLd []map[string]any `json:"jsonLd"`
	// Keywords extracted from the page's keywords meta tag.
	Keywords []string `json:"keywords"`
	// Language extracted from html lang or language meta tags.
	Language string `json:"language"`
	// Modified timestamp/date from page metadata, when present.
	ModifiedTime string `json:"modifiedTime"`
	// Open Graph metadata with the og: prefix removed and keys camel-cased.
	OpenGraph map[string]BatchGetResultsResponseDataOkMetadataOpenGraphUnion `json:"openGraph"`
	// Published timestamp/date from page metadata, when present.
	PublishedTime string `json:"publishedTime"`
	// Robots meta directive, when present.
	Robots string `json:"robots"`
	// Site or application name from page metadata.
	SiteName string `json:"siteName"`
	// Best title extracted from the page.
	Title string `json:"title"`
	// Twitter card metadata with the twitter: prefix removed and keys camel-cased.
	Twitter map[string]BatchGetResultsResponseDataOkMetadataTwitterUnion `json:"twitter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FinalURL       respjson.Field
		SourceURL      respjson.Field
		AdditionalMeta respjson.Field
		Alternates     respjson.Field
		Author         respjson.Field
		CanonicalURL   respjson.Field
		Description    respjson.Field
		Favicon        respjson.Field
		Image          respjson.Field
		JsonLd         respjson.Field
		Keywords       respjson.Field
		Language       respjson.Field
		ModifiedTime   respjson.Field
		OpenGraph      respjson.Field
		PublishedTime  respjson.Field
		Robots         respjson.Field
		SiteName       respjson.Field
		Title          respjson.Field
		Twitter        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataOkMetadata) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataOkMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchGetResultsResponseDataOkMetadataAdditionalMetaUnion contains all possible
// properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type BatchGetResultsResponseDataOkMetadataAdditionalMetaUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u BatchGetResultsResponseDataOkMetadataAdditionalMetaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BatchGetResultsResponseDataOkMetadataAdditionalMetaUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BatchGetResultsResponseDataOkMetadataAdditionalMetaUnion) RawJSON() string { return u.JSON.raw }

func (r *BatchGetResultsResponseDataOkMetadataAdditionalMetaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchGetResultsResponseDataOkMetadataAlternate struct {
	// Resolved alternate URL.
	Href string `json:"href" api:"required"`
	// Language or locale for the alternate URL, when present.
	Hreflang string `json:"hreflang"`
	// Alternate resource title, when present.
	Title string `json:"title"`
	// Alternate resource MIME type, when present.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Href        respjson.Field
		Hreflang    respjson.Field
		Title       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataOkMetadataAlternate) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataOkMetadataAlternate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchGetResultsResponseDataOkMetadataOpenGraphUnion contains all possible
// properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type BatchGetResultsResponseDataOkMetadataOpenGraphUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u BatchGetResultsResponseDataOkMetadataOpenGraphUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BatchGetResultsResponseDataOkMetadataOpenGraphUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BatchGetResultsResponseDataOkMetadataOpenGraphUnion) RawJSON() string { return u.JSON.raw }

func (r *BatchGetResultsResponseDataOkMetadataOpenGraphUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchGetResultsResponseDataOkMetadataTwitterUnion contains all possible
// properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type BatchGetResultsResponseDataOkMetadataTwitterUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u BatchGetResultsResponseDataOkMetadataTwitterUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BatchGetResultsResponseDataOkMetadataTwitterUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BatchGetResultsResponseDataOkMetadataTwitterUnion) RawJSON() string { return u.JSON.raw }

func (r *BatchGetResultsResponseDataOkMetadataTwitterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A page the batch could not fetch.
type BatchGetResultsResponseDataError struct {
	// Why the page failed.
	ErrorCode string `json:"error_code" api:"required"`
	// Human-readable failure detail.
	Message string `json:"message" api:"required"`
	// The page could not be scraped.
	Status constant.Error `json:"status" default:"error"`
	// URL as submitted, or as discovered by the crawl.
	URL string `json:"url" api:"required"`
	// Caller-supplied identifier echoed from submission.
	ItemID string `json:"itemId"`
	// Caller-supplied metadata echoed from submission.
	Meta map[string]any `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ErrorCode   respjson.Field
		Message     respjson.Field
		Status      respjson.Field
		URL         respjson.Field
		ItemID      respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataError) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the API key used for the request. Included in every response
// whenever a valid API key is provided, even when the response status is not 200.
type BatchGetResultsResponseKeyMetadata struct {
	// The number of credits consumed by this request.
	CreditsConsumed int64 `json:"credits_consumed" api:"required"`
	// The number of credits remaining for your organization after this request.
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
func (r BatchGetResultsResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchSubmitResponse struct {
	// HTTP status code.
	//
	// Any of 200.
	Code int64 `json:"code" api:"required"`
	// Additional response details.
	Metadata BatchSubmitResponseMetadata `json:"metadata" api:"required"`
	// Retrieved person profile.
	Person BatchSubmitResponsePerson `json:"person" api:"required"`
	// Response status.
	//
	// Any of "ok".
	Status BatchSubmitResponseStatus `json:"status" api:"required"`
	// Metadata about the API key used for the request. Included in every response
	// whenever a valid API key is provided, even when the response status is not 200.
	KeyMetadata BatchSubmitResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Metadata    respjson.Field
		Person      respjson.Field
		Status      respjson.Field
		KeyMetadata respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchSubmitResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchSubmitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional response details.
type BatchSubmitResponseMetadata struct {
	// Identifiers returned for the person.
	Identifiers BatchSubmitResponseMetadataIdentifiers `json:"identifiers" api:"required"`
	// Source categories checked.
	//
	// Any of "linkedin", "cv", "manual", "github", "other".
	SourcesAttempted []string `json:"sourcesAttempted" api:"required"`
	// Source categories with data.
	//
	// Any of "linkedin", "cv", "manual", "github", "other".
	SourcesSucceeded []string `json:"sourcesSucceeded" api:"required"`
	// URLs reviewed for this profile.
	URLsAnalyzed []string `json:"urlsAnalyzed" api:"required" format:"uri"`
	// Personal website URL, when found.
	PersonalWebsiteURL string `json:"personalWebsiteUrl" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Identifiers        respjson.Field
		SourcesAttempted   respjson.Field
		SourcesSucceeded   respjson.Field
		URLsAnalyzed       respjson.Field
		PersonalWebsiteURL respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchSubmitResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *BatchSubmitResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifiers returned for the person.
type BatchSubmitResponseMetadataIdentifiers struct {
	// LinkedIn profile URL.
	LinkedinURL string `json:"linkedinUrl" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LinkedinURL respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchSubmitResponseMetadataIdentifiers) RawJSON() string { return r.JSON.raw }
func (r *BatchSubmitResponseMetadataIdentifiers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Retrieved person profile.
type BatchSubmitResponsePerson struct {
	// Education history.
	Education []BatchSubmitResponsePersonEducation `json:"education" api:"required"`
	// Work history.
	Experience []BatchSubmitResponsePersonExperience `json:"experience" api:"required"`
	// Core profile details.
	Profile BatchSubmitResponsePersonProfile `json:"profile" api:"required"`
	// Listed skills.
	Skills []BatchSubmitResponsePersonSkill `json:"skills" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Education   respjson.Field
		Experience  respjson.Field
		Profile     respjson.Field
		Skills      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchSubmitResponsePerson) RawJSON() string { return r.JSON.raw }
func (r *BatchSubmitResponsePerson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchSubmitResponsePersonEducation struct {
	// School or institution name.
	Institution BatchSubmitResponsePersonEducationInstitution `json:"institution" api:"required"`
	// Education dates.
	Dates BatchSubmitResponsePersonEducationDates `json:"dates"`
	// Additional education details.
	Description string `json:"description"`
	// Area of study.
	FieldOfStudy string `json:"fieldOfStudy"`
	// Degree, certificate, or credential.
	Qualification string `json:"qualification"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Institution   respjson.Field
		Dates         respjson.Field
		Description   respjson.Field
		FieldOfStudy  respjson.Field
		Qualification respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchSubmitResponsePersonEducation) RawJSON() string { return r.JSON.raw }
func (r *BatchSubmitResponsePersonEducation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// School or institution name.
type BatchSubmitResponsePersonEducationInstitution struct {
	// Display name.
	Display string `json:"display" api:"required"`
	// Standardized name, when available.
	Normalized string `json:"normalized"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Display     respjson.Field
		Normalized  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchSubmitResponsePersonEducationInstitution) RawJSON() string { return r.JSON.raw }
func (r *BatchSubmitResponsePersonEducationInstitution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Education dates.
type BatchSubmitResponsePersonEducationDates struct {
	// End date, when known.
	EndDate BatchSubmitResponsePersonEducationDatesEndDate `json:"endDate"`
	// Whether the entry is current.
	IsCurrent bool `json:"isCurrent"`
	// Start date, when known.
	StartDate BatchSubmitResponsePersonEducationDatesStartDate `json:"startDate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndDate     respjson.Field
		IsCurrent   respjson.Field
		StartDate   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchSubmitResponsePersonEducationDates) RawJSON() string { return r.JSON.raw }
func (r *BatchSubmitResponsePersonEducationDates) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// End date, when known.
type BatchSubmitResponsePersonEducationDatesEndDate struct {
	// Year value.
	Year int64 `json:"year" api:"required"`
	// Day value, when known.
	Day int64 `json:"day"`
	// Month value, when known.
	Month int64 `json:"month"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Year        respjson.Field
		Day         respjson.Field
		Month       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchSubmitResponsePersonEducationDatesEndDate) RawJSON() string { return r.JSON.raw }
func (r *BatchSubmitResponsePersonEducationDatesEndDate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Start date, when known.
type BatchSubmitResponsePersonEducationDatesStartDate struct {
	// Year value.
	Year int64 `json:"year" api:"required"`
	// Day value, when known.
	Day int64 `json:"day"`
	// Month value, when known.
	Month int64 `json:"month"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Year        respjson.Field
		Day         respjson.Field
		Month       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchSubmitResponsePersonEducationDatesStartDate) RawJSON() string { return r.JSON.raw }
func (r *BatchSubmitResponsePersonEducationDatesStartDate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchSubmitResponsePersonExperience struct {
	// Company or organization name.
	Company BatchSubmitResponsePersonExperienceCompany `json:"company" api:"required"`
	// Role or job title.
	Title string `json:"title" api:"required"`
	// Role dates.
	Dates BatchSubmitResponsePersonExperienceDates `json:"dates"`
	// Role description.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Company     respjson.Field
		Title       respjson.Field
		Dates       respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchSubmitResponsePersonExperience) RawJSON() string { return r.JSON.raw }
func (r *BatchSubmitResponsePersonExperience) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Company or organization name.
type BatchSubmitResponsePersonExperienceCompany struct {
	// Display name.
	Display string `json:"display" api:"required"`
	// Standardized name, when available.
	Normalized string `json:"normalized"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Display     respjson.Field
		Normalized  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchSubmitResponsePersonExperienceCompany) RawJSON() string { return r.JSON.raw }
func (r *BatchSubmitResponsePersonExperienceCompany) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Role dates.
type BatchSubmitResponsePersonExperienceDates struct {
	// End date, when known.
	EndDate BatchSubmitResponsePersonExperienceDatesEndDate `json:"endDate"`
	// Whether the entry is current.
	IsCurrent bool `json:"isCurrent"`
	// Start date, when known.
	StartDate BatchSubmitResponsePersonExperienceDatesStartDate `json:"startDate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndDate     respjson.Field
		IsCurrent   respjson.Field
		StartDate   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchSubmitResponsePersonExperienceDates) RawJSON() string { return r.JSON.raw }
func (r *BatchSubmitResponsePersonExperienceDates) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// End date, when known.
type BatchSubmitResponsePersonExperienceDatesEndDate struct {
	// Year value.
	Year int64 `json:"year" api:"required"`
	// Day value, when known.
	Day int64 `json:"day"`
	// Month value, when known.
	Month int64 `json:"month"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Year        respjson.Field
		Day         respjson.Field
		Month       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchSubmitResponsePersonExperienceDatesEndDate) RawJSON() string { return r.JSON.raw }
func (r *BatchSubmitResponsePersonExperienceDatesEndDate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Start date, when known.
type BatchSubmitResponsePersonExperienceDatesStartDate struct {
	// Year value.
	Year int64 `json:"year" api:"required"`
	// Day value, when known.
	Day int64 `json:"day"`
	// Month value, when known.
	Month int64 `json:"month"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Year        respjson.Field
		Day         respjson.Field
		Month       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchSubmitResponsePersonExperienceDatesStartDate) RawJSON() string { return r.JSON.raw }
func (r *BatchSubmitResponsePersonExperienceDatesStartDate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Core profile details.
type BatchSubmitResponsePersonProfile struct {
	// Person's full name.
	FullName string `json:"fullName"`
	// Short professional headline.
	Headline string `json:"headline"`
	// Person's listed location.
	Location string `json:"location"`
	// Profile image URL.
	ProfilePictureURL string `json:"profilePictureUrl" format:"uri"`
	// Brief profile summary.
	Summary string `json:"summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FullName          respjson.Field
		Headline          respjson.Field
		Location          respjson.Field
		ProfilePictureURL respjson.Field
		Summary           respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchSubmitResponsePersonProfile) RawJSON() string { return r.JSON.raw }
func (r *BatchSubmitResponsePersonProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchSubmitResponsePersonSkill struct {
	// Skill name.
	Name string `json:"name" api:"required"`
	// Standardized skill name, when available.
	Normalized string `json:"normalized"`
	// Skill proficiency, when available.
	Proficiency string `json:"proficiency"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Normalized  respjson.Field
		Proficiency respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchSubmitResponsePersonSkill) RawJSON() string { return r.JSON.raw }
func (r *BatchSubmitResponsePersonSkill) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response status.
type BatchSubmitResponseStatus string

const (
	BatchSubmitResponseStatusOk BatchSubmitResponseStatus = "ok"
)

// Metadata about the API key used for the request. Included in every response
// whenever a valid API key is provided, even when the response status is not 200.
type BatchSubmitResponseKeyMetadata struct {
	// The number of credits consumed by this request.
	CreditsConsumed int64 `json:"credits_consumed" api:"required"`
	// The number of credits remaining for your organization after this request.
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
func (r BatchSubmitResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *BatchSubmitResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchListParams struct {
	// Cursor from the previous page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Batches per page. Defaults to 25.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Free-text search term, matched against the batch id, crawl source (start URL or
	// sitemap domain), and tags.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Comma-separated list of tags to filter by (matches batches having any of them).
	Tags param.Opt[string] `query:"tags,omitzero" json:"-"`
	// `prefix` for as-you-type prefix matching (default), `exact` for full-token
	// matching.
	//
	// Any of "exact", "prefix".
	SearchType BatchListParamsSearchType `query:"search_type,omitzero" json:"-"`
	// Filter by status.
	//
	// Any of "queued", "running", "cancelling", "completed", "cancelled", "failed".
	Status BatchListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BatchListParams]'s query parameters as `url.Values`.
func (r BatchListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// `prefix` for as-you-type prefix matching (default), `exact` for full-token
// matching.
type BatchListParamsSearchType string

const (
	BatchListParamsSearchTypeExact  BatchListParamsSearchType = "exact"
	BatchListParamsSearchTypePrefix BatchListParamsSearchType = "prefix"
)

// Filter by status.
type BatchListParamsStatus string

const (
	BatchListParamsStatusQueued     BatchListParamsStatus = "queued"
	BatchListParamsStatusRunning    BatchListParamsStatus = "running"
	BatchListParamsStatusCancelling BatchListParamsStatus = "cancelling"
	BatchListParamsStatusCompleted  BatchListParamsStatus = "completed"
	BatchListParamsStatusCancelled  BatchListParamsStatus = "cancelled"
	BatchListParamsStatusFailed     BatchListParamsStatus = "failed"
)

type BatchGetResultsParams struct {
	// next_cursor from the previous page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Records per page. Defaults to 25. A page can close early so its payload stays
	// under ~8 MB; rely on next_cursor rather than counting records.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BatchGetResultsParams]'s query parameters as `url.Values`.
func (r BatchGetResultsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BatchSubmitParams struct {
	// Known identifiers for the person. At least one identifier is required.
	Identifiers BatchSubmitParamsIdentifiers `json:"identifiers,omitzero" api:"required"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `json:"timeoutMS,omitzero"`
	// Optional tags for tracking usage. Up to 20 tags, each 1 to 50 characters.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r BatchSubmitParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Known identifiers for the person. At least one identifier is required.
type BatchSubmitParamsIdentifiers struct {
	// LinkedIn profile URL, e.g. https://www.linkedin.com/in/yahia-bakour/.
	LinkedinURL param.Opt[string] `json:"linkedinUrl,omitzero" format:"uri"`
	paramObj
}

func (r BatchSubmitParamsIdentifiers) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsIdentifiers
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsIdentifiers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
