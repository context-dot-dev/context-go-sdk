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

// Scrape many pages or crawl a site asynchronously.
//
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

// Check progress, and get download links once the batch finishes.
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

// Permanently delete a finished batch and its stored results. Active batches must
// settle first.
func (r *BatchService) Delete(ctx context.Context, batchID string, opts ...option.RequestOption) (res *BatchDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if batchID == "" {
		err = errors.New("missing required batch_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("batch/%s", url.PathEscape(batchID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
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

// Page through a finished batch's results as JSON instead of downloading the
// NDJSON files.
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

// Scrape 25K URLs or crawl large websites asynchronously.
func (r *BatchService) Submit(ctx context.Context, params BatchSubmitParams, opts ...option.RequestOption) (res *BatchSubmitResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "batch/submit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
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
// [CrawlControlsSourceStartURL], [CrawlControlsSourceSitemap].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CrawlControlsSourceUnion struct {
	Type string `json:"type"`
	// This field is from variant [CrawlControlsSourceStartURL].
	URL string `json:"url"`
	// This field is from variant [CrawlControlsSourceSitemap].
	Domain string `json:"domain"`
	JSON   struct {
		Type   respjson.Field
		URL    respjson.Field
		Domain respjson.Field
		raw    string
	} `json:"-"`
}

func (u CrawlControlsSourceUnion) AsStartURL() (v CrawlControlsSourceStartURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CrawlControlsSourceUnion) AsSitemap() (v CrawlControlsSourceSitemap) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CrawlControlsSourceUnion) RawJSON() string { return u.JSON.raw }

func (r *CrawlControlsSourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The crawl discovered pages by following links from one URL.
type CrawlControlsSourceStartURL struct {
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
func (r CrawlControlsSourceStartURL) RawJSON() string { return r.JSON.raw }
func (r *CrawlControlsSourceStartURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The crawl scraped the pages listed in the domain's sitemap.
type CrawlControlsSourceSitemap struct {
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
func (r CrawlControlsSourceSitemap) RawJSON() string { return r.JSON.raw }
func (r *CrawlControlsSourceSitemap) UnmarshalJSON(data []byte) error {
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
	// `reserved` minus `refunded` plus `ocr_charged` — what the batch has cost so far.
	// Equal to `reserved` until the batch settles.
	Net int64 `json:"net" api:"required"`
	// Credits charged for PDF pages recovered by OCR (pdf.ocr=true), 1 per recovered
	// page, on top of `reserved`. Stays 0 until the batch settles.
	OcrCharged int64 `json:"ocr_charged" api:"required"`
	// Credits returned for pages that did not succeed. Stays 0 until the batch reaches
	// a final status, then settles in one movement.
	Refunded int64 `json:"refunded" api:"required"`
	// Credits debited from your balance the moment the batch was accepted. This is a
	// charge, not a forecast — the whole amount leaves the balance up front.
	Reserved int64 `json:"reserved" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Net         respjson.Field
		OcrCharged  respjson.Field
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
	// `reserved` minus `refunded` plus `ocr_charged` — what the batch has cost so far.
	// Equal to `reserved` until the batch settles.
	Net int64 `json:"net" api:"required"`
	// Credits charged for PDF pages recovered by OCR (pdf.ocr=true), 1 per recovered
	// page, on top of `reserved`. Stays 0 until the batch settles.
	OcrCharged int64 `json:"ocr_charged" api:"required"`
	// Credits returned for pages that did not succeed. Stays 0 until the batch reaches
	// a final status, then settles in one movement.
	Refunded int64 `json:"refunded" api:"required"`
	// Credits debited from your balance the moment the batch was accepted. This is a
	// charge, not a forecast — the whole amount leaves the balance up front.
	Reserved int64 `json:"reserved" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Net         respjson.Field
		OcrCharged  respjson.Field
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

type BatchDeleteResponse struct {
	// ID of the deleted batch.
	ID string `json:"id"`
	// Always true on success.
	Deleted bool `json:"deleted"`
	// Metadata about the API key used for the request. Included in every response
	// whenever a valid API key is provided, even when the response status is not 200.
	KeyMetadata BatchDeleteResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Deleted     respjson.Field
		KeyMetadata respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the API key used for the request. Included in every response
// whenever a valid API key is provided, even when the response status is not 200.
type BatchDeleteResponseKeyMetadata struct {
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
func (r BatchDeleteResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *BatchDeleteResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchCancelResponse struct {
	// Batch ID.
	ID string `json:"id" api:"required"`
	// The crawl controls as submitted, so the limits requested can be compared against
	// what the crawl reached.
	Crawl CrawlControls `json:"crawl" api:"required"`
	// What this batch cost so far.
	Credits BatchCancelResponseCredits `json:"credits" api:"required"`
	// What each page is returned as.
	//
	// Any of "markdown", "html".
	Format BatchCancelResponseFormat `json:"format" api:"required"`
	// What submission took in, and what it charged for.
	Input Intake `json:"input" api:"required"`
	// How pages were selected.
	//
	// Any of "scrape", "crawl".
	Mode BatchCancelResponseMode `json:"mode" api:"required"`
	// Page failures so far, grouped by error code and sorted by count.
	PageErrors []PageErrorCount `json:"page_errors" api:"required"`
	// How far the batch got before cancellation.
	Progress BatchCancelResponseProgress `json:"progress" api:"required"`
	// Always `cancelling`. Work already in flight finishes; the batch reaches
	// `cancelled` shortly after.
	//
	// Any of "cancelling".
	Status BatchCancelResponseStatus `json:"status" api:"required"`
	// Tags stored on the batch at submission.
	Tags []string `json:"tags" api:"required"`
	// There is no finish time yet — the batch is still winding down.
	Timing BatchCancelResponseTiming `json:"timing" api:"required"`
	// API key usage for this request.
	KeyMetadata BatchCancelResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Crawl       respjson.Field
		Credits     respjson.Field
		Format      respjson.Field
		Input       respjson.Field
		Mode        respjson.Field
		PageErrors  respjson.Field
		Progress    respjson.Field
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

// What this batch cost so far.
type BatchCancelResponseCredits struct {
	// Credits debited at submission. The unspent remainder is refunded once the batch
	// settles — read `credits.refunded` from GET /batch/{batch_id} then.
	Reserved int64 `json:"reserved" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
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

// What each page is returned as.
type BatchCancelResponseFormat string

const (
	BatchCancelResponseFormatMarkdown BatchCancelResponseFormat = "markdown"
	BatchCancelResponseFormatHTML     BatchCancelResponseFormat = "html"
)

// How pages were selected.
type BatchCancelResponseMode string

const (
	BatchCancelResponseModeScrape BatchCancelResponseMode = "scrape"
	BatchCancelResponseModeCrawl  BatchCancelResponseMode = "crawl"
)

// How far the batch got before cancellation.
type BatchCancelResponseProgress struct {
	// Pages that could not be scraped before the request landed.
	Failed int64 `json:"failed" api:"required"`
	// Reserved pages that will now be skipped, and refunded when the batch settles.
	Pending int64 `json:"pending" api:"required"`
	// Pages scraped successfully before the request landed.
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

// Always `cancelling`. Work already in flight finishes; the batch reaches
// `cancelled` shortly after.
type BatchCancelResponseStatus string

const (
	BatchCancelResponseStatusCancelling BatchCancelResponseStatus = "cancelling"
)

// There is no finish time yet — the batch is still winding down.
type BatchCancelResponseTiming struct {
	// When the batch was created.
	CreatedAt string `json:"created_at" api:"required"`
	// When processing started. Null if it was cancelled while still queued.
	StartedAt string `json:"started_at" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
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
	CacheMetadata BatchGetResultsResponseDataOkCacheMetadata `json:"cache_metadata"`
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
	// This field is from variant [BatchGetResultsResponseDataOk].
	OcrPages int64 `json:"ocr_pages"`
	// This field is from variant [BatchGetResultsResponseDataError].
	ErrorCode string `json:"error_code"`
	// This field is from variant [BatchGetResultsResponseDataError].
	Message string `json:"message"`
	JSON    struct {
		CacheMetadata respjson.Field
		FinalURL      respjson.Field
		HTTPStatus    respjson.Field
		Metadata      respjson.Field
		Status        respjson.Field
		URL           respjson.Field
		HTML          respjson.Field
		ItemID        respjson.Field
		Markdown      respjson.Field
		Meta          respjson.Field
		OcrPages      respjson.Field
		ErrorCode     respjson.Field
		Message       respjson.Field
		raw           string
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
	// Cache outcome for this response. Composite responses are hits only when every
	// cache-controlled fetch contributing to the output was a hit; age_ms is the
	// oldest contributing hit.
	CacheMetadata BatchGetResultsResponseDataOkCacheMetadata `json:"cache_metadata" api:"required"`
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
	// Page HTML. Present on html batches, and on markdown batches submitted with
	// `options.includeHTML`.
	HTML string `json:"html"`
	// Caller-supplied identifier echoed from submission.
	ItemID string `json:"itemId"`
	// Page content as Markdown. Present on markdown batches.
	Markdown string `json:"markdown"`
	// Caller-supplied metadata echoed from submission.
	Meta map[string]any `json:"meta"`
	// PDF pages of this document recovered by OCR (pdf.ocr=true). Each recovered page
	// bills 1 credit on top of the page base credit; absent when no OCR ran.
	OcrPages int64 `json:"ocr_pages"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CacheMetadata respjson.Field
		FinalURL      respjson.Field
		HTTPStatus    respjson.Field
		Metadata      respjson.Field
		Status        respjson.Field
		URL           respjson.Field
		HTML          respjson.Field
		ItemID        respjson.Field
		Markdown      respjson.Field
		Meta          respjson.Field
		OcrPages      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataOk) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataOk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache outcome for this response. Composite responses are hits only when every
// cache-controlled fetch contributing to the output was a hit; age_ms is the
// oldest contributing hit.
type BatchGetResultsResponseDataOkCacheMetadata struct {
	// Age of the cached data in milliseconds. Zero for miss and zdr responses.
	AgeMs int64 `json:"age_ms" api:"required"`
	// Whether the response was served from cache, required fresh work, or honored
	// zero-data-retention cache bypass.
	//
	// Any of "hit", "miss", "zdr".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgeMs       respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataOkCacheMetadata) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataOkCacheMetadata) UnmarshalJSON(data []byte) error {
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
	// Page headings (h1–h6) in document order, extracted from the unfiltered document.
	// Capped at the first 500 headings. Omitted when the page has none.
	Headings []BatchGetResultsResponseDataOkMetadataHeading `json:"headings"`
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
		Headings       respjson.Field
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

type BatchGetResultsResponseDataOkMetadataHeading struct {
	// Heading level, 1–6 (from h1–h6).
	Level int64 `json:"level" api:"required"`
	// Heading text with whitespace collapsed, truncated to 1000 characters.
	Text string `json:"text" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Level       respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataOkMetadataHeading) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataOkMetadataHeading) UnmarshalJSON(data []byte) error {
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
	// Batch ID. Poll GET /batch/{batch_id} with it.
	ID string `json:"id" api:"required"`
	// Cache outcome for this response. Composite responses are hits only when every
	// cache-controlled fetch contributing to the output was a hit; age_ms is the
	// oldest contributing hit.
	CacheMetadata BatchSubmitResponseCacheMetadata `json:"cache_metadata" api:"required"`
	// The crawl controls as submitted, so the limits requested can be compared against
	// what the crawl reached.
	Crawl CrawlControls `json:"crawl" api:"required"`
	// When the batch was created.
	CreatedAt string `json:"created_at" api:"required"`
	// What accepting this batch cost.
	Credits BatchSubmitResponseCredits `json:"credits" api:"required"`
	// What each page will be returned as.
	//
	// Any of "markdown", "html".
	Format BatchSubmitResponseFormat `json:"format" api:"required"`
	// What submission took in, and what it charged for.
	Input Intake `json:"input" api:"required"`
	// Rejected URLs, up to 100. These are not charged.
	InvalidURLs []BatchSubmitResponseInvalidURL `json:"invalid_urls" api:"required"`
	// How pages will be selected.
	//
	// Any of "scrape", "crawl".
	Mode BatchSubmitResponseMode `json:"mode" api:"required"`
	// Always `queued`. An accepted batch has not started yet.
	//
	// Any of "queued".
	Status BatchSubmitResponseStatus `json:"status" api:"required"`
	// Tags stored on the batch.
	Tags []string `json:"tags" api:"required"`
	// API key usage for this request.
	KeyMetadata BatchSubmitResponseKeyMetadata `json:"key_metadata"`
	// Signing secret for the completion webhook, returned only here and never again.
	// Store it now; it is not repeated by GET /batch/{batch_id}.
	WebhookSecret string `json:"webhook_secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CacheMetadata respjson.Field
		Crawl         respjson.Field
		CreatedAt     respjson.Field
		Credits       respjson.Field
		Format        respjson.Field
		Input         respjson.Field
		InvalidURLs   respjson.Field
		Mode          respjson.Field
		Status        respjson.Field
		Tags          respjson.Field
		KeyMetadata   respjson.Field
		WebhookSecret respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchSubmitResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchSubmitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache outcome for this response. Composite responses are hits only when every
// cache-controlled fetch contributing to the output was a hit; age_ms is the
// oldest contributing hit.
type BatchSubmitResponseCacheMetadata struct {
	// Age of the cached data in milliseconds. Zero for miss and zdr responses.
	AgeMs int64 `json:"age_ms" api:"required"`
	// Whether the response was served from cache, required fresh work, or honored
	// zero-data-retention cache bypass.
	//
	// Any of "hit", "miss", "zdr".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgeMs       respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchSubmitResponseCacheMetadata) RawJSON() string { return r.JSON.raw }
func (r *BatchSubmitResponseCacheMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What accepting this batch cost.
type BatchSubmitResponseCredits struct {
	// Credits just debited from your balance. Whatever the batch does not spend is
	// refunded when it settles.
	Reserved int64 `json:"reserved" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reserved    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchSubmitResponseCredits) RawJSON() string { return r.JSON.raw }
func (r *BatchSubmitResponseCredits) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What each page will be returned as.
type BatchSubmitResponseFormat string

const (
	BatchSubmitResponseFormatMarkdown BatchSubmitResponseFormat = "markdown"
	BatchSubmitResponseFormatHTML     BatchSubmitResponseFormat = "html"
)

type BatchSubmitResponseInvalidURL struct {
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
func (r BatchSubmitResponseInvalidURL) RawJSON() string { return r.JSON.raw }
func (r *BatchSubmitResponseInvalidURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How pages will be selected.
type BatchSubmitResponseMode string

const (
	BatchSubmitResponseModeScrape BatchSubmitResponseMode = "scrape"
	BatchSubmitResponseModeCrawl  BatchSubmitResponseMode = "crawl"
)

// Always `queued`. An accepted batch has not started yet.
type BatchSubmitResponseStatus string

const (
	BatchSubmitResponseStatusQueued BatchSubmitResponseStatus = "queued"
)

// API key usage for this request.
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
	// Choose a URL list or a site crawl.
	Input BatchSubmitParamsInputUnion `json:"input,omitzero" api:"required"`
	// URL notified when the batch finishes.
	WebhookURL param.Opt[string] `json:"webhookUrl,omitzero"`
	// Any string unique to this submission. Retries with the same key return the
	// original batch.
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	// Tags stored on the batch. Filter the batch list by them later.
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputUnion struct {
	OfScrape *BatchSubmitParamsInputScrape `json:",omitzero,inline"`
	OfCrawl  *BatchSubmitParamsInputCrawl  `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfScrape, u.OfCrawl)
}
func (u *BatchSubmitParamsInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[BatchSubmitParamsInputUnion](
		"mode",
		apijson.Discriminator[BatchSubmitParamsInputScrape]("scrape"),
		apijson.Discriminator[BatchSubmitParamsInputCrawl]("crawl"),
	)
}

// Scrape up to 25K URLs in one batch.
//
// The properties Data, Mode are required.
type BatchSubmitParamsInputScrape struct {
	// Pages to scrape and their output format.
	Data BatchSubmitParamsInputScrapeDataUnion `json:"data,omitzero" api:"required"`
	// Scrape the pages in `data.urls`.
	//
	// This field can be elided, and will marshal its zero value as "scrape".
	Mode constant.Scrape `json:"mode" default:"scrape"`
	paramObj
}

func (r BatchSubmitParamsInputScrape) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrape
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrape) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputScrapeDataUnion struct {
	OfMarkdown *BatchSubmitParamsInputScrapeDataMarkdown `json:",omitzero,inline"`
	OfHTML     *BatchSubmitParamsInputScrapeDataHTML     `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputScrapeDataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfMarkdown, u.OfHTML)
}
func (u *BatchSubmitParamsInputScrapeDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[BatchSubmitParamsInputScrapeDataUnion](
		"format",
		apijson.Discriminator[BatchSubmitParamsInputScrapeDataMarkdown]("markdown"),
		apijson.Discriminator[BatchSubmitParamsInputScrapeDataHTML]("html"),
	)
}

// Scrape the listed pages as Markdown.
//
// The properties Format, URLs are required.
type BatchSubmitParamsInputScrapeDataMarkdown struct {
	// Pages to scrape. Maximum 25000.
	URLs []BatchSubmitParamsInputScrapeDataMarkdownURL `json:"urls,omitzero" api:"required"`
	// Options for Markdown output.
	Options BatchSubmitParamsInputScrapeDataMarkdownOptions `json:"options,omitzero"`
	// Return page content as Markdown.
	//
	// This field can be elided, and will marshal its zero value as "markdown".
	Format constant.Markdown `json:"format" default:"markdown"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataMarkdown) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataMarkdown
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataMarkdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A page to scrape, with optional data for matching results.
//
// The property URL is required.
type BatchSubmitParamsInputScrapeDataMarkdownURL struct {
	// Page URL to scrape.
	URL string `json:"url" api:"required"`
	// Your ID for this page, returned with its result. The same URL can use different
	// IDs.
	ItemID param.Opt[string] `json:"itemId,omitzero"`
	// Custom JSON returned unchanged with this page result.
	Meta map[string]any `json:"meta,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataMarkdownURL) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataMarkdownURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataMarkdownURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Options for Markdown output.
type BatchSubmitParamsInputScrapeDataMarkdownOptions struct {
	// Return a cached result if a prior scrape for the same parameters exists and is
	// younger than this many milliseconds. Defaults to 1 day (86400000 ms) when
	// omitted. Max is 30 days (2592000000 ms). Set to 0 to always scrape fresh.
	MaxAgeMs param.Opt[int64] `json:"maxAgeMs,omitzero"`
	// Also include each page's HTML in its result record, as an `html` field alongside
	// the Markdown.
	IncludeHTML param.Opt[bool] `json:"includeHTML,omitzero"`
	// Include image references in the Markdown.
	IncludeImages param.Opt[bool] `json:"includeImages,omitzero"`
	// Include links in the Markdown.
	IncludeLinks param.Opt[bool] `json:"includeLinks,omitzero"`
	// Wait briefly for CSS and transition animations to settle before extraction, on
	// pages that render in a browser.
	SettleAnimations param.Opt[bool] `json:"settleAnimations,omitzero"`
	// Shorten inline base64 image data.
	ShortenBase64Images param.Opt[bool] `json:"shortenBase64Images,omitzero"`
	// Return the main content without navigation or footers.
	UseMainContentOnly param.Opt[bool] `json:"useMainContentOnly,omitzero"`
	// How long to wait after initial page load, in milliseconds. `0` waits 500 ms.
	WaitForMs param.Opt[int64] `json:"waitForMs,omitzero"`
	// Remove elements matching these CSS selectors. Applied after `includeSelectors`,
	// so an element matching both is removed.
	ExcludeSelectors []string `json:"excludeSelectors,omitzero"`
	// Keep only the subtrees matching these CSS selectors. Filtered pages are always
	// fetched fresh, ignoring `maxAgeMs`.
	IncludeSelectors []string `json:"includeSelectors,omitzero"`
	// Fetch the target page through a residential proxy in this country (ISO 3166-1
	// alpha-2).
	//
	// Any of "ad", "ae", "af", "ag", "ai", "al", "am", "ao", "ar", "at", "au", "aw",
	// "az", "ba", "bb", "bd", "be", "bf", "bg", "bh", "bi", "bj", "bm", "bn", "bo",
	// "bq", "br", "bs", "bw", "by", "bz", "ca", "cd", "cf", "cg", "ch", "ci", "cl",
	// "cm", "cn", "co", "cr", "cv", "cw", "cy", "cz", "de", "dj", "dk", "dm", "do",
	// "dz", "ec", "ee", "eg", "es", "et", "fi", "fj", "fr", "ga", "gb", "gd", "ge",
	// "gf", "gg", "gh", "gm", "gn", "gp", "gq", "gr", "gt", "gu", "gw", "gy", "hk",
	// "hn", "hr", "ht", "hu", "id", "ie", "il", "im", "in", "iq", "ir", "is", "it",
	// "je", "jm", "jo", "jp", "ke", "kg", "kh", "kn", "kr", "kw", "ky", "kz", "la",
	// "lb", "lc", "lk", "lr", "ls", "lt", "lu", "lv", "ly", "ma", "mc", "md", "me",
	// "mf", "mg", "mk", "ml", "mm", "mn", "mo", "mq", "mr", "mt", "mu", "mv", "mw",
	// "mx", "my", "mz", "na", "nc", "ne", "ng", "ni", "nl", "no", "np", "nz", "om",
	// "pa", "pe", "pf", "pg", "ph", "pk", "pl", "pr", "ps", "pt", "py", "qa", "re",
	// "ro", "rs", "ru", "rw", "sa", "sc", "sd", "se", "sg", "si", "sk", "sl", "sm",
	// "sn", "so", "sr", "ss", "st", "sv", "sx", "sy", "sz", "tc", "td", "tg", "th",
	// "tj", "tl", "tm", "tn", "tr", "tt", "tw", "tz", "ua", "ug", "us", "uy", "uz",
	// "vc", "ve", "vg", "vi", "vn", "ye", "yt", "za", "zm", "zw".
	Country string `json:"country,omitzero"`
	// PDF parsing controls. Use start/end to limit text extraction and embedded-image
	// detection/OCR to an inclusive 1-based page range.
	Pdf BatchSubmitParamsInputScrapeDataMarkdownOptionsPdf `json:"pdf,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataMarkdownOptions) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataMarkdownOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataMarkdownOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputScrapeDataMarkdownOptions](
		"country", "ad", "ae", "af", "ag", "ai", "al", "am", "ao", "ar", "at", "au", "aw", "az", "ba", "bb", "bd", "be", "bf", "bg", "bh", "bi", "bj", "bm", "bn", "bo", "bq", "br", "bs", "bw", "by", "bz", "ca", "cd", "cf", "cg", "ch", "ci", "cl", "cm", "cn", "co", "cr", "cv", "cw", "cy", "cz", "de", "dj", "dk", "dm", "do", "dz", "ec", "ee", "eg", "es", "et", "fi", "fj", "fr", "ga", "gb", "gd", "ge", "gf", "gg", "gh", "gm", "gn", "gp", "gq", "gr", "gt", "gu", "gw", "gy", "hk", "hn", "hr", "ht", "hu", "id", "ie", "il", "im", "in", "iq", "ir", "is", "it", "je", "jm", "jo", "jp", "ke", "kg", "kh", "kn", "kr", "kw", "ky", "kz", "la", "lb", "lc", "lk", "lr", "ls", "lt", "lu", "lv", "ly", "ma", "mc", "md", "me", "mf", "mg", "mk", "ml", "mm", "mn", "mo", "mq", "mr", "mt", "mu", "mv", "mw", "mx", "my", "mz", "na", "nc", "ne", "ng", "ni", "nl", "no", "np", "nz", "om", "pa", "pe", "pf", "pg", "ph", "pk", "pl", "pr", "ps", "pt", "py", "qa", "re", "ro", "rs", "ru", "rw", "sa", "sc", "sd", "se", "sg", "si", "sk", "sl", "sm", "sn", "so", "sr", "ss", "st", "sv", "sx", "sy", "sz", "tc", "td", "tg", "th", "tj", "tl", "tm", "tn", "tr", "tt", "tw", "tz", "ua", "ug", "us", "uy", "uz", "vc", "ve", "vg", "vi", "vn", "ye", "yt", "za", "zm", "zw",
	)
}

// PDF parsing controls. Use start/end to limit text extraction and embedded-image
// detection/OCR to an inclusive 1-based page range.
type BatchSubmitParamsInputScrapeDataMarkdownOptionsPdf struct {
	// Last 1-based PDF page to parse. When omitted, parsing ends at the last page.
	// Must be greater than or equal to start when both are provided.
	End param.Opt[int64] `json:"end,omitzero"`
	// When true, OCR the selected PDF pages that have no usable text layer (scans),
	// replacing each recovered page's text with the OCR result while pages with a real
	// text layer keep it. Billed at 1 credit per page OCR actually recovered, on top
	// of the base request cost. When false, no OCR runs.
	Ocr param.Opt[bool] `json:"ocr,omitzero"`
	// When true, PDF URLs are fetched and parsed. When false, PDF URLs are skipped and
	// a 400 PDF_SKIPPED is returned.
	ShouldParse param.Opt[bool] `json:"shouldParse,omitzero"`
	// First 1-based PDF page to parse. When omitted, parsing starts at the first page.
	Start param.Opt[int64] `json:"start,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataMarkdownOptionsPdf) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataMarkdownOptionsPdf
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataMarkdownOptionsPdf) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Scrape the listed pages as HTML.
//
// The properties Format, URLs are required.
type BatchSubmitParamsInputScrapeDataHTML struct {
	// Pages to scrape. Maximum 25000.
	URLs []BatchSubmitParamsInputScrapeDataHTMLURL `json:"urls,omitzero" api:"required"`
	// Options for HTML output.
	Options BatchSubmitParamsInputScrapeDataHTMLOptions `json:"options,omitzero"`
	// Return page content as HTML.
	//
	// This field can be elided, and will marshal its zero value as "html".
	Format constant.HTML `json:"format" default:"html"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataHTML) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataHTML
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataHTML) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A page to scrape, with optional data for matching results.
//
// The property URL is required.
type BatchSubmitParamsInputScrapeDataHTMLURL struct {
	// Page URL to scrape.
	URL string `json:"url" api:"required"`
	// Your ID for this page, returned with its result. The same URL can use different
	// IDs.
	ItemID param.Opt[string] `json:"itemId,omitzero"`
	// Custom JSON returned unchanged with this page result.
	Meta map[string]any `json:"meta,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataHTMLURL) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataHTMLURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataHTMLURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Options for HTML output.
type BatchSubmitParamsInputScrapeDataHTMLOptions struct {
	// Return a cached result if a prior scrape for the same parameters exists and is
	// younger than this many milliseconds. Defaults to 1 day (86400000 ms) when
	// omitted. Max is 30 days (2592000000 ms). Set to 0 to always scrape fresh.
	MaxAgeMs param.Opt[int64] `json:"maxAgeMs,omitzero"`
	// Wait briefly for CSS and transition animations to settle before extraction, on
	// pages that render in a browser.
	SettleAnimations param.Opt[bool] `json:"settleAnimations,omitzero"`
	// Return the main content without navigation or footers.
	UseMainContentOnly param.Opt[bool] `json:"useMainContentOnly,omitzero"`
	// How long to wait after initial page load, in milliseconds. `0` waits 500 ms.
	WaitForMs param.Opt[int64] `json:"waitForMs,omitzero"`
	// Remove elements matching these CSS selectors. Applied after `includeSelectors`,
	// so an element matching both is removed.
	ExcludeSelectors []string `json:"excludeSelectors,omitzero"`
	// Keep only the subtrees matching these CSS selectors. Filtered pages are always
	// fetched fresh, ignoring `maxAgeMs`.
	IncludeSelectors []string `json:"includeSelectors,omitzero"`
	// Fetch the target page through a residential proxy in this country (ISO 3166-1
	// alpha-2).
	//
	// Any of "ad", "ae", "af", "ag", "ai", "al", "am", "ao", "ar", "at", "au", "aw",
	// "az", "ba", "bb", "bd", "be", "bf", "bg", "bh", "bi", "bj", "bm", "bn", "bo",
	// "bq", "br", "bs", "bw", "by", "bz", "ca", "cd", "cf", "cg", "ch", "ci", "cl",
	// "cm", "cn", "co", "cr", "cv", "cw", "cy", "cz", "de", "dj", "dk", "dm", "do",
	// "dz", "ec", "ee", "eg", "es", "et", "fi", "fj", "fr", "ga", "gb", "gd", "ge",
	// "gf", "gg", "gh", "gm", "gn", "gp", "gq", "gr", "gt", "gu", "gw", "gy", "hk",
	// "hn", "hr", "ht", "hu", "id", "ie", "il", "im", "in", "iq", "ir", "is", "it",
	// "je", "jm", "jo", "jp", "ke", "kg", "kh", "kn", "kr", "kw", "ky", "kz", "la",
	// "lb", "lc", "lk", "lr", "ls", "lt", "lu", "lv", "ly", "ma", "mc", "md", "me",
	// "mf", "mg", "mk", "ml", "mm", "mn", "mo", "mq", "mr", "mt", "mu", "mv", "mw",
	// "mx", "my", "mz", "na", "nc", "ne", "ng", "ni", "nl", "no", "np", "nz", "om",
	// "pa", "pe", "pf", "pg", "ph", "pk", "pl", "pr", "ps", "pt", "py", "qa", "re",
	// "ro", "rs", "ru", "rw", "sa", "sc", "sd", "se", "sg", "si", "sk", "sl", "sm",
	// "sn", "so", "sr", "ss", "st", "sv", "sx", "sy", "sz", "tc", "td", "tg", "th",
	// "tj", "tl", "tm", "tn", "tr", "tt", "tw", "tz", "ua", "ug", "us", "uy", "uz",
	// "vc", "ve", "vg", "vi", "vn", "ye", "yt", "za", "zm", "zw".
	Country string `json:"country,omitzero"`
	// PDF parsing controls. Use start/end to limit text extraction and embedded-image
	// detection/OCR to an inclusive 1-based page range.
	Pdf BatchSubmitParamsInputScrapeDataHTMLOptionsPdf `json:"pdf,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataHTMLOptions) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataHTMLOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataHTMLOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputScrapeDataHTMLOptions](
		"country", "ad", "ae", "af", "ag", "ai", "al", "am", "ao", "ar", "at", "au", "aw", "az", "ba", "bb", "bd", "be", "bf", "bg", "bh", "bi", "bj", "bm", "bn", "bo", "bq", "br", "bs", "bw", "by", "bz", "ca", "cd", "cf", "cg", "ch", "ci", "cl", "cm", "cn", "co", "cr", "cv", "cw", "cy", "cz", "de", "dj", "dk", "dm", "do", "dz", "ec", "ee", "eg", "es", "et", "fi", "fj", "fr", "ga", "gb", "gd", "ge", "gf", "gg", "gh", "gm", "gn", "gp", "gq", "gr", "gt", "gu", "gw", "gy", "hk", "hn", "hr", "ht", "hu", "id", "ie", "il", "im", "in", "iq", "ir", "is", "it", "je", "jm", "jo", "jp", "ke", "kg", "kh", "kn", "kr", "kw", "ky", "kz", "la", "lb", "lc", "lk", "lr", "ls", "lt", "lu", "lv", "ly", "ma", "mc", "md", "me", "mf", "mg", "mk", "ml", "mm", "mn", "mo", "mq", "mr", "mt", "mu", "mv", "mw", "mx", "my", "mz", "na", "nc", "ne", "ng", "ni", "nl", "no", "np", "nz", "om", "pa", "pe", "pf", "pg", "ph", "pk", "pl", "pr", "ps", "pt", "py", "qa", "re", "ro", "rs", "ru", "rw", "sa", "sc", "sd", "se", "sg", "si", "sk", "sl", "sm", "sn", "so", "sr", "ss", "st", "sv", "sx", "sy", "sz", "tc", "td", "tg", "th", "tj", "tl", "tm", "tn", "tr", "tt", "tw", "tz", "ua", "ug", "us", "uy", "uz", "vc", "ve", "vg", "vi", "vn", "ye", "yt", "za", "zm", "zw",
	)
}

// PDF parsing controls. Use start/end to limit text extraction and embedded-image
// detection/OCR to an inclusive 1-based page range.
type BatchSubmitParamsInputScrapeDataHTMLOptionsPdf struct {
	// Last 1-based PDF page to parse. When omitted, parsing ends at the last page.
	// Must be greater than or equal to start when both are provided.
	End param.Opt[int64] `json:"end,omitzero"`
	// When true, OCR the selected PDF pages that have no usable text layer (scans),
	// replacing each recovered page's text with the OCR result while pages with a real
	// text layer keep it. Billed at 1 credit per page OCR actually recovered, on top
	// of the base request cost. When false, no OCR runs.
	Ocr param.Opt[bool] `json:"ocr,omitzero"`
	// When true, PDF URLs are fetched and parsed. When false, PDF URLs are skipped and
	// a 400 PDF_SKIPPED is returned.
	ShouldParse param.Opt[bool] `json:"shouldParse,omitzero"`
	// First 1-based PDF page to parse. When omitted, parsing starts at the first page.
	Start param.Opt[int64] `json:"start,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataHTMLOptionsPdf) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataHTMLOptionsPdf
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataHTMLOptionsPdf) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Crawl pages starting from a URL or from a domain's sitemap.
//
// The properties Data, Mode are required.
type BatchSubmitParamsInputCrawl struct {
	// Crawl source and output format.
	Data BatchSubmitParamsInputCrawlDataUnion `json:"data,omitzero" api:"required"`
	// Discover and scrape pages from `data.source`.
	//
	// This field can be elided, and will marshal its zero value as "crawl".
	Mode constant.Crawl `json:"mode" default:"crawl"`
	paramObj
}

func (r BatchSubmitParamsInputCrawl) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawl
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawl) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataUnion struct {
	OfMarkdown *BatchSubmitParamsInputCrawlDataMarkdown `json:",omitzero,inline"`
	OfHTML     *BatchSubmitParamsInputCrawlDataHTML     `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfMarkdown, u.OfHTML)
}
func (u *BatchSubmitParamsInputCrawlDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[BatchSubmitParamsInputCrawlDataUnion](
		"format",
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataMarkdown]("markdown"),
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataHTML]("html"),
	)
}

// Crawl pages and return Markdown.
//
// The properties Format, Source are required.
type BatchSubmitParamsInputCrawlDataMarkdown struct {
	// How to find pages to crawl.
	Source BatchSubmitParamsInputCrawlDataMarkdownSourceUnion `json:"source,omitzero" api:"required"`
	// Options for Markdown output.
	Options BatchSubmitParamsInputCrawlDataMarkdownOptions `json:"options,omitzero"`
	// Return page content as Markdown.
	//
	// This field can be elided, and will marshal its zero value as "markdown".
	Format constant.Markdown `json:"format" default:"markdown"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataMarkdown) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataMarkdown
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataMarkdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataMarkdownSourceUnion struct {
	OfStartURL *BatchSubmitParamsInputCrawlDataMarkdownSourceStartURL `json:",omitzero,inline"`
	OfSitemap  *BatchSubmitParamsInputCrawlDataMarkdownSourceSitemap  `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataMarkdownSourceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfStartURL, u.OfSitemap)
}
func (u *BatchSubmitParamsInputCrawlDataMarkdownSourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[BatchSubmitParamsInputCrawlDataMarkdownSourceUnion](
		"type",
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataMarkdownSourceStartURL]("start_url"),
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataMarkdownSourceSitemap]("sitemap"),
	)
}

// Discover pages by following links from one URL.
//
// The properties Type, URL are required.
type BatchSubmitParamsInputCrawlDataMarkdownSourceStartURL struct {
	// Page where crawling begins. A URL without a scheme is read as https://.
	URL string `json:"url" api:"required"`
	// Limits and filters for page discovery.
	Controls BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControls `json:"controls,omitzero"`
	// Start from one page.
	//
	// This field can be elided, and will marshal its zero value as "start_url".
	Type constant.StartURL `json:"type" default:"start_url"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataMarkdownSourceStartURL) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataMarkdownSourceStartURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataMarkdownSourceStartURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Limits and filters for page discovery.
type BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControls struct {
	// Follow links to subdomains.
	FollowSubdomains param.Opt[bool] `json:"followSubdomains,omitzero"`
	// Maximum link depth. Source pages are depth 0. No limit when omitted.
	MaxDepth param.Opt[int64] `json:"maxDepth,omitzero"`
	// Maximum pages to fetch. Unused reserved credits are refunded. Maximum 25000.
	MaxURLs param.Opt[int64] `json:"maxUrls,omitzero"`
	// RE2 pattern for URLs to include. The `start_url` itself is always included.
	Regex param.Opt[string] `json:"regex,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControls) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControls
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControls) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Scrape the pages listed in a domain's sitemap. Links on those pages are not
// followed.
//
// The properties Domain, Type are required.
type BatchSubmitParamsInputCrawlDataMarkdownSourceSitemap struct {
	// Domain whose sitemap lists the pages to scrape. A full URL is reduced to its
	// domain.
	Domain string `json:"domain" api:"required"`
	// Limits and filters for the sitemap URLs. A sitemap batch scrapes exactly those
	// URLs and never follows links off them, so there is no crawl depth here.
	Controls BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControls `json:"controls,omitzero"`
	// Scrape the URLs in the domain's sitemap.
	//
	// This field can be elided, and will marshal its zero value as "sitemap".
	Type constant.Sitemap `json:"type" default:"sitemap"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataMarkdownSourceSitemap) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataMarkdownSourceSitemap
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataMarkdownSourceSitemap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Limits and filters for the sitemap URLs. A sitemap batch scrapes exactly those
// URLs and never follows links off them, so there is no crawl depth here.
type BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControls struct {
	// Maximum pages to fetch. Unused reserved credits are refunded. Maximum 25000.
	MaxURLs param.Opt[int64] `json:"maxUrls,omitzero"`
	// RE2 pattern; only sitemap URLs matching it are scraped.
	Regex param.Opt[string] `json:"regex,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControls) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControls
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControls) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Options for Markdown output.
type BatchSubmitParamsInputCrawlDataMarkdownOptions struct {
	// Return a cached result if a prior scrape for the same parameters exists and is
	// younger than this many milliseconds. Defaults to 1 day (86400000 ms) when
	// omitted. Max is 30 days (2592000000 ms). Set to 0 to always scrape fresh.
	MaxAgeMs param.Opt[int64] `json:"maxAgeMs,omitzero"`
	// Also include each page's HTML in its result record, as an `html` field alongside
	// the Markdown.
	IncludeHTML param.Opt[bool] `json:"includeHTML,omitzero"`
	// Include image references in the Markdown.
	IncludeImages param.Opt[bool] `json:"includeImages,omitzero"`
	// Include links in the Markdown.
	IncludeLinks param.Opt[bool] `json:"includeLinks,omitzero"`
	// Wait briefly for CSS and transition animations to settle before extraction, on
	// pages that render in a browser.
	SettleAnimations param.Opt[bool] `json:"settleAnimations,omitzero"`
	// Shorten inline base64 image data.
	ShortenBase64Images param.Opt[bool] `json:"shortenBase64Images,omitzero"`
	// Return the main content without navigation or footers.
	UseMainContentOnly param.Opt[bool] `json:"useMainContentOnly,omitzero"`
	// How long to wait after initial page load, in milliseconds. `0` waits 500 ms.
	WaitForMs param.Opt[int64] `json:"waitForMs,omitzero"`
	// Remove elements matching these CSS selectors. Applied after `includeSelectors`,
	// so an element matching both is removed.
	ExcludeSelectors []string `json:"excludeSelectors,omitzero"`
	// Keep only the subtrees matching these CSS selectors. Filtered pages are always
	// fetched fresh, ignoring `maxAgeMs`.
	IncludeSelectors []string `json:"includeSelectors,omitzero"`
	// Fetch the target page through a residential proxy in this country (ISO 3166-1
	// alpha-2).
	//
	// Any of "ad", "ae", "af", "ag", "ai", "al", "am", "ao", "ar", "at", "au", "aw",
	// "az", "ba", "bb", "bd", "be", "bf", "bg", "bh", "bi", "bj", "bm", "bn", "bo",
	// "bq", "br", "bs", "bw", "by", "bz", "ca", "cd", "cf", "cg", "ch", "ci", "cl",
	// "cm", "cn", "co", "cr", "cv", "cw", "cy", "cz", "de", "dj", "dk", "dm", "do",
	// "dz", "ec", "ee", "eg", "es", "et", "fi", "fj", "fr", "ga", "gb", "gd", "ge",
	// "gf", "gg", "gh", "gm", "gn", "gp", "gq", "gr", "gt", "gu", "gw", "gy", "hk",
	// "hn", "hr", "ht", "hu", "id", "ie", "il", "im", "in", "iq", "ir", "is", "it",
	// "je", "jm", "jo", "jp", "ke", "kg", "kh", "kn", "kr", "kw", "ky", "kz", "la",
	// "lb", "lc", "lk", "lr", "ls", "lt", "lu", "lv", "ly", "ma", "mc", "md", "me",
	// "mf", "mg", "mk", "ml", "mm", "mn", "mo", "mq", "mr", "mt", "mu", "mv", "mw",
	// "mx", "my", "mz", "na", "nc", "ne", "ng", "ni", "nl", "no", "np", "nz", "om",
	// "pa", "pe", "pf", "pg", "ph", "pk", "pl", "pr", "ps", "pt", "py", "qa", "re",
	// "ro", "rs", "ru", "rw", "sa", "sc", "sd", "se", "sg", "si", "sk", "sl", "sm",
	// "sn", "so", "sr", "ss", "st", "sv", "sx", "sy", "sz", "tc", "td", "tg", "th",
	// "tj", "tl", "tm", "tn", "tr", "tt", "tw", "tz", "ua", "ug", "us", "uy", "uz",
	// "vc", "ve", "vg", "vi", "vn", "ye", "yt", "za", "zm", "zw".
	Country string `json:"country,omitzero"`
	// PDF parsing controls. Use start/end to limit text extraction and embedded-image
	// detection/OCR to an inclusive 1-based page range.
	Pdf BatchSubmitParamsInputCrawlDataMarkdownOptionsPdf `json:"pdf,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataMarkdownOptions) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataMarkdownOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataMarkdownOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputCrawlDataMarkdownOptions](
		"country", "ad", "ae", "af", "ag", "ai", "al", "am", "ao", "ar", "at", "au", "aw", "az", "ba", "bb", "bd", "be", "bf", "bg", "bh", "bi", "bj", "bm", "bn", "bo", "bq", "br", "bs", "bw", "by", "bz", "ca", "cd", "cf", "cg", "ch", "ci", "cl", "cm", "cn", "co", "cr", "cv", "cw", "cy", "cz", "de", "dj", "dk", "dm", "do", "dz", "ec", "ee", "eg", "es", "et", "fi", "fj", "fr", "ga", "gb", "gd", "ge", "gf", "gg", "gh", "gm", "gn", "gp", "gq", "gr", "gt", "gu", "gw", "gy", "hk", "hn", "hr", "ht", "hu", "id", "ie", "il", "im", "in", "iq", "ir", "is", "it", "je", "jm", "jo", "jp", "ke", "kg", "kh", "kn", "kr", "kw", "ky", "kz", "la", "lb", "lc", "lk", "lr", "ls", "lt", "lu", "lv", "ly", "ma", "mc", "md", "me", "mf", "mg", "mk", "ml", "mm", "mn", "mo", "mq", "mr", "mt", "mu", "mv", "mw", "mx", "my", "mz", "na", "nc", "ne", "ng", "ni", "nl", "no", "np", "nz", "om", "pa", "pe", "pf", "pg", "ph", "pk", "pl", "pr", "ps", "pt", "py", "qa", "re", "ro", "rs", "ru", "rw", "sa", "sc", "sd", "se", "sg", "si", "sk", "sl", "sm", "sn", "so", "sr", "ss", "st", "sv", "sx", "sy", "sz", "tc", "td", "tg", "th", "tj", "tl", "tm", "tn", "tr", "tt", "tw", "tz", "ua", "ug", "us", "uy", "uz", "vc", "ve", "vg", "vi", "vn", "ye", "yt", "za", "zm", "zw",
	)
}

// PDF parsing controls. Use start/end to limit text extraction and embedded-image
// detection/OCR to an inclusive 1-based page range.
type BatchSubmitParamsInputCrawlDataMarkdownOptionsPdf struct {
	// Last 1-based PDF page to parse. When omitted, parsing ends at the last page.
	// Must be greater than or equal to start when both are provided.
	End param.Opt[int64] `json:"end,omitzero"`
	// When true, OCR the selected PDF pages that have no usable text layer (scans),
	// replacing each recovered page's text with the OCR result while pages with a real
	// text layer keep it. Billed at 1 credit per page OCR actually recovered, on top
	// of the base request cost. When false, no OCR runs.
	Ocr param.Opt[bool] `json:"ocr,omitzero"`
	// When true, PDF URLs are fetched and parsed. When false, PDF URLs are skipped and
	// a 400 PDF_SKIPPED is returned.
	ShouldParse param.Opt[bool] `json:"shouldParse,omitzero"`
	// First 1-based PDF page to parse. When omitted, parsing starts at the first page.
	Start param.Opt[int64] `json:"start,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataMarkdownOptionsPdf) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataMarkdownOptionsPdf
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataMarkdownOptionsPdf) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Crawl pages and return HTML.
//
// The properties Format, Source are required.
type BatchSubmitParamsInputCrawlDataHTML struct {
	// How to find pages to crawl.
	Source BatchSubmitParamsInputCrawlDataHTMLSourceUnion `json:"source,omitzero" api:"required"`
	// Options for HTML output.
	Options BatchSubmitParamsInputCrawlDataHTMLOptions `json:"options,omitzero"`
	// Return page content as HTML.
	//
	// This field can be elided, and will marshal its zero value as "html".
	Format constant.HTML `json:"format" default:"html"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataHTML) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataHTML
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataHTML) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataHTMLSourceUnion struct {
	OfStartURL *BatchSubmitParamsInputCrawlDataHTMLSourceStartURL `json:",omitzero,inline"`
	OfSitemap  *BatchSubmitParamsInputCrawlDataHTMLSourceSitemap  `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataHTMLSourceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfStartURL, u.OfSitemap)
}
func (u *BatchSubmitParamsInputCrawlDataHTMLSourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[BatchSubmitParamsInputCrawlDataHTMLSourceUnion](
		"type",
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataHTMLSourceStartURL]("start_url"),
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataHTMLSourceSitemap]("sitemap"),
	)
}

// Discover pages by following links from one URL.
//
// The properties Type, URL are required.
type BatchSubmitParamsInputCrawlDataHTMLSourceStartURL struct {
	// Page where crawling begins. A URL without a scheme is read as https://.
	URL string `json:"url" api:"required"`
	// Limits and filters for page discovery.
	Controls BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControls `json:"controls,omitzero"`
	// Start from one page.
	//
	// This field can be elided, and will marshal its zero value as "start_url".
	Type constant.StartURL `json:"type" default:"start_url"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataHTMLSourceStartURL) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataHTMLSourceStartURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataHTMLSourceStartURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Limits and filters for page discovery.
type BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControls struct {
	// Follow links to subdomains.
	FollowSubdomains param.Opt[bool] `json:"followSubdomains,omitzero"`
	// Maximum link depth. Source pages are depth 0. No limit when omitted.
	MaxDepth param.Opt[int64] `json:"maxDepth,omitzero"`
	// Maximum pages to fetch. Unused reserved credits are refunded. Maximum 25000.
	MaxURLs param.Opt[int64] `json:"maxUrls,omitzero"`
	// RE2 pattern for URLs to include. The `start_url` itself is always included.
	Regex param.Opt[string] `json:"regex,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControls) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControls
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControls) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Scrape the pages listed in a domain's sitemap. Links on those pages are not
// followed.
//
// The properties Domain, Type are required.
type BatchSubmitParamsInputCrawlDataHTMLSourceSitemap struct {
	// Domain whose sitemap lists the pages to scrape. A full URL is reduced to its
	// domain.
	Domain string `json:"domain" api:"required"`
	// Limits and filters for the sitemap URLs. A sitemap batch scrapes exactly those
	// URLs and never follows links off them, so there is no crawl depth here.
	Controls BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControls `json:"controls,omitzero"`
	// Scrape the URLs in the domain's sitemap.
	//
	// This field can be elided, and will marshal its zero value as "sitemap".
	Type constant.Sitemap `json:"type" default:"sitemap"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataHTMLSourceSitemap) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataHTMLSourceSitemap
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataHTMLSourceSitemap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Limits and filters for the sitemap URLs. A sitemap batch scrapes exactly those
// URLs and never follows links off them, so there is no crawl depth here.
type BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControls struct {
	// Maximum pages to fetch. Unused reserved credits are refunded. Maximum 25000.
	MaxURLs param.Opt[int64] `json:"maxUrls,omitzero"`
	// RE2 pattern; only sitemap URLs matching it are scraped.
	Regex param.Opt[string] `json:"regex,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControls) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControls
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControls) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Options for HTML output.
type BatchSubmitParamsInputCrawlDataHTMLOptions struct {
	// Return a cached result if a prior scrape for the same parameters exists and is
	// younger than this many milliseconds. Defaults to 1 day (86400000 ms) when
	// omitted. Max is 30 days (2592000000 ms). Set to 0 to always scrape fresh.
	MaxAgeMs param.Opt[int64] `json:"maxAgeMs,omitzero"`
	// Wait briefly for CSS and transition animations to settle before extraction, on
	// pages that render in a browser.
	SettleAnimations param.Opt[bool] `json:"settleAnimations,omitzero"`
	// Return the main content without navigation or footers.
	UseMainContentOnly param.Opt[bool] `json:"useMainContentOnly,omitzero"`
	// How long to wait after initial page load, in milliseconds. `0` waits 500 ms.
	WaitForMs param.Opt[int64] `json:"waitForMs,omitzero"`
	// Remove elements matching these CSS selectors. Applied after `includeSelectors`,
	// so an element matching both is removed.
	ExcludeSelectors []string `json:"excludeSelectors,omitzero"`
	// Keep only the subtrees matching these CSS selectors. Filtered pages are always
	// fetched fresh, ignoring `maxAgeMs`.
	IncludeSelectors []string `json:"includeSelectors,omitzero"`
	// Fetch the target page through a residential proxy in this country (ISO 3166-1
	// alpha-2).
	//
	// Any of "ad", "ae", "af", "ag", "ai", "al", "am", "ao", "ar", "at", "au", "aw",
	// "az", "ba", "bb", "bd", "be", "bf", "bg", "bh", "bi", "bj", "bm", "bn", "bo",
	// "bq", "br", "bs", "bw", "by", "bz", "ca", "cd", "cf", "cg", "ch", "ci", "cl",
	// "cm", "cn", "co", "cr", "cv", "cw", "cy", "cz", "de", "dj", "dk", "dm", "do",
	// "dz", "ec", "ee", "eg", "es", "et", "fi", "fj", "fr", "ga", "gb", "gd", "ge",
	// "gf", "gg", "gh", "gm", "gn", "gp", "gq", "gr", "gt", "gu", "gw", "gy", "hk",
	// "hn", "hr", "ht", "hu", "id", "ie", "il", "im", "in", "iq", "ir", "is", "it",
	// "je", "jm", "jo", "jp", "ke", "kg", "kh", "kn", "kr", "kw", "ky", "kz", "la",
	// "lb", "lc", "lk", "lr", "ls", "lt", "lu", "lv", "ly", "ma", "mc", "md", "me",
	// "mf", "mg", "mk", "ml", "mm", "mn", "mo", "mq", "mr", "mt", "mu", "mv", "mw",
	// "mx", "my", "mz", "na", "nc", "ne", "ng", "ni", "nl", "no", "np", "nz", "om",
	// "pa", "pe", "pf", "pg", "ph", "pk", "pl", "pr", "ps", "pt", "py", "qa", "re",
	// "ro", "rs", "ru", "rw", "sa", "sc", "sd", "se", "sg", "si", "sk", "sl", "sm",
	// "sn", "so", "sr", "ss", "st", "sv", "sx", "sy", "sz", "tc", "td", "tg", "th",
	// "tj", "tl", "tm", "tn", "tr", "tt", "tw", "tz", "ua", "ug", "us", "uy", "uz",
	// "vc", "ve", "vg", "vi", "vn", "ye", "yt", "za", "zm", "zw".
	Country string `json:"country,omitzero"`
	// PDF parsing controls. Use start/end to limit text extraction and embedded-image
	// detection/OCR to an inclusive 1-based page range.
	Pdf BatchSubmitParamsInputCrawlDataHTMLOptionsPdf `json:"pdf,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataHTMLOptions) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataHTMLOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataHTMLOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputCrawlDataHTMLOptions](
		"country", "ad", "ae", "af", "ag", "ai", "al", "am", "ao", "ar", "at", "au", "aw", "az", "ba", "bb", "bd", "be", "bf", "bg", "bh", "bi", "bj", "bm", "bn", "bo", "bq", "br", "bs", "bw", "by", "bz", "ca", "cd", "cf", "cg", "ch", "ci", "cl", "cm", "cn", "co", "cr", "cv", "cw", "cy", "cz", "de", "dj", "dk", "dm", "do", "dz", "ec", "ee", "eg", "es", "et", "fi", "fj", "fr", "ga", "gb", "gd", "ge", "gf", "gg", "gh", "gm", "gn", "gp", "gq", "gr", "gt", "gu", "gw", "gy", "hk", "hn", "hr", "ht", "hu", "id", "ie", "il", "im", "in", "iq", "ir", "is", "it", "je", "jm", "jo", "jp", "ke", "kg", "kh", "kn", "kr", "kw", "ky", "kz", "la", "lb", "lc", "lk", "lr", "ls", "lt", "lu", "lv", "ly", "ma", "mc", "md", "me", "mf", "mg", "mk", "ml", "mm", "mn", "mo", "mq", "mr", "mt", "mu", "mv", "mw", "mx", "my", "mz", "na", "nc", "ne", "ng", "ni", "nl", "no", "np", "nz", "om", "pa", "pe", "pf", "pg", "ph", "pk", "pl", "pr", "ps", "pt", "py", "qa", "re", "ro", "rs", "ru", "rw", "sa", "sc", "sd", "se", "sg", "si", "sk", "sl", "sm", "sn", "so", "sr", "ss", "st", "sv", "sx", "sy", "sz", "tc", "td", "tg", "th", "tj", "tl", "tm", "tn", "tr", "tt", "tw", "tz", "ua", "ug", "us", "uy", "uz", "vc", "ve", "vg", "vi", "vn", "ye", "yt", "za", "zm", "zw",
	)
}

// PDF parsing controls. Use start/end to limit text extraction and embedded-image
// detection/OCR to an inclusive 1-based page range.
type BatchSubmitParamsInputCrawlDataHTMLOptionsPdf struct {
	// Last 1-based PDF page to parse. When omitted, parsing ends at the last page.
	// Must be greater than or equal to start when both are provided.
	End param.Opt[int64] `json:"end,omitzero"`
	// When true, OCR the selected PDF pages that have no usable text layer (scans),
	// replacing each recovered page's text with the OCR result while pages with a real
	// text layer keep it. Billed at 1 credit per page OCR actually recovered, on top
	// of the base request cost. When false, no OCR runs.
	Ocr param.Opt[bool] `json:"ocr,omitzero"`
	// When true, PDF URLs are fetched and parsed. When false, PDF URLs are skipped and
	// a 400 PDF_SKIPPED is returned.
	ShouldParse param.Opt[bool] `json:"shouldParse,omitzero"`
	// First 1-based PDF page to parse. When omitted, parsing starts at the first page.
	Start param.Opt[int64] `json:"start,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataHTMLOptionsPdf) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataHTMLOptionsPdf
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataHTMLOptionsPdf) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
