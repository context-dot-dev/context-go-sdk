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

// Get batch progress and result download links. Result files are deleted 180 days
// after the batch finishes.
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

// List your batches, newest first, with optional filters.
func (r *BatchService) List(ctx context.Context, query BatchListParams, opts ...option.RequestOption) (res *BatchListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "batch/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Permanently delete a finished batch and its results. Its webhook deliveries can
// no longer be retried.
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

// Stop a batch from starting new pages. Pages already in progress finish before
// the batch becomes cancelled.
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

// Page through a finished batch’s results as JSON. Each record contains result
// with the same outputs and per-format errors as POST /web/scrape. Results remain
// available for 180 days.
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

// Scrape up to 100,000 URLs, or crawl a site, asynchronously using the same
// formats and scrape parameters as POST /web/scrape. Choose Markdown, HTML,
// screenshots, images, bytes, parsed fields, products, highlights, and JSON. Poll
// the batch ID or receive a webhook when it finishes.
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

// Crawl settings as submitted.
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

// What the submission accepted.
type Intake struct {
	// URLs dropped before reserving because another entry resolved to the same page.
	// Non-zero for sitemap crawls too, whose sitemaps routinely list a page more than
	// once.
	Duplicates int64 `json:"duplicates" api:"required"`
	// Rejected input URLs; `null` for a crawl.
	Invalid int64 `json:"invalid" api:"required"`
	// Pages accepted; progress counts toward this total.
	Reserved int64 `json:"reserved" api:"required"`
	// True when `reserved` is a crawl ceiling; false when it is an exact URL count.
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
	// Batch ID.
	ID string `json:"id" api:"required"`
	// Crawl settings as submitted.
	Crawl CrawlControls `json:"crawl" api:"required"`
	// Batch credit usage and settlement.
	Credits BatchGetResponseCredits `json:"credits" api:"required"`
	// A failure of the batch as a whole, distinct from the per-page failures in
	// `page_errors`.
	Failure Failure `json:"failure" api:"required"`
	// Outputs requested for each page. Matches input.data.formats on submission.
	Formats BatchGetResponseFormats `json:"formats" api:"required"`
	// What the submission accepted.
	Input Intake `json:"input" api:"required"`
	// Rejected URLs (first 100).
	InvalidURLs []BatchGetResponseInvalidURL `json:"invalid_urls" api:"required"`
	// `scrape` (URL list) or `crawl`.
	//
	// Any of "scrape", "crawl".
	Mode BatchGetResponseMode `json:"mode" api:"required"`
	// Individual page failures grouped by error code, sorted by count. Unrelated to
	// `failure`, which is the batch itself failing.
	PageErrors []PageErrorCount `json:"page_errors" api:"required"`
	// Pages attempted so far. Use `status` to check completion.
	Progress BatchGetResponseProgress `json:"progress" api:"required"`
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Result download links; null until the batch finishes. Files are deleted 180 days
	// after the batch finishes.
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
	// Batch completion delivery ID, when available.
	WebhookDeliveryID string `json:"webhook_delivery_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Crawl             respjson.Field
		Credits           respjson.Field
		Failure           respjson.Field
		Formats           respjson.Field
		Input             respjson.Field
		InvalidURLs       respjson.Field
		Mode              respjson.Field
		PageErrors        respjson.Field
		Progress          respjson.Field
		RequestID         respjson.Field
		Results           respjson.Field
		Status            respjson.Field
		Tags              respjson.Field
		Timing            respjson.Field
		KeyMetadata       respjson.Field
		WebhookDeliveryID respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Batch credit usage and settlement.
type BatchGetResponseCredits struct {
	// `reserved` minus `refunded` plus `ocr_charged`.
	Net int64 `json:"net" api:"required"`
	// OCR usage charged when the batch settles.
	OcrCharged int64 `json:"ocr_charged" api:"required"`
	// Credits returned for unsuccessful pages when the batch settles.
	Refunded int64 `json:"refunded" api:"required"`
	// Credits held when the batch was accepted.
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

// Outputs requested for each page. Matches input.data.formats on submission.
type BatchGetResponseFormats struct {
	// The original HTTP response body.
	Bytes bool `json:"bytes"`
	// Markdown excerpts relevant to `highlightsParams.query`.
	Highlights bool `json:"highlights"`
	// Rendered HTML.
	HTML bool `json:"html"`
	// Images found on the page.
	Images bool `json:"images"`
	// An object matching `jsonParams.schema`, extracted from the page.
	Json bool `json:"json"`
	// Page content as Markdown.
	Markdown bool `json:"markdown"`
	// Fields extracted with custom rules or a prebuilt parser, returned as `parsed`.
	Parse bool `json:"parse"`
	// Product details such as name, price, and availability.
	Product bool `json:"product"`
	// A screenshot of the page.
	Screenshot bool `json:"screenshot"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bytes       respjson.Field
		Highlights  respjson.Field
		HTML        respjson.Field
		Images      respjson.Field
		Json        respjson.Field
		Markdown    respjson.Field
		Parse       respjson.Field
		Product     respjson.Field
		Screenshot  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResponseFormats) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResponseFormats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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

// `scrape` (URL list) or `crawl`.
type BatchGetResponseMode string

const (
	BatchGetResponseModeScrape BatchGetResponseMode = "scrape"
	BatchGetResponseModeCrawl  BatchGetResponseMode = "crawl"
)

// Pages attempted so far. Use `status` to check completion.
type BatchGetResponseProgress struct {
	// Pages that could not be scraped.
	Failed int64 `json:"failed" api:"required"`
	// Accepted pages not yet attempted. Unused crawl capacity is excluded after
	// completion.
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

// Result download links; null until the batch finishes. Files are deleted 180 days
// after the batch finishes.
type BatchGetResponseResults struct {
	// When these links expire (24 hours after this response).
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
func (r BatchGetResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchListResponse struct {
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Batches on this page.
	Data []BatchListResponseData `json:"data"`
	// Whether another page is available.
	HasMore bool `json:"has_more"`
	// Credits this request used and your remaining balance.
	KeyMetadata BatchListResponseKeyMetadata `json:"key_metadata"`
	// Cursor for the next page.
	NextCursor string `json:"next_cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestID   respjson.Field
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
	// Batch ID.
	ID string `json:"id" api:"required"`
	// Crawl settings as submitted.
	Crawl CrawlControls `json:"crawl" api:"required"`
	// Batch credit usage and settlement.
	Credits BatchListResponseDataCredits `json:"credits" api:"required"`
	// A failure of the batch as a whole, distinct from the per-page failures in
	// `page_errors`.
	Failure Failure `json:"failure" api:"required"`
	// Outputs requested for each page. Matches input.data.formats on submission.
	Formats BatchListResponseDataFormats `json:"formats" api:"required"`
	// What the submission accepted.
	Input Intake `json:"input" api:"required"`
	// `scrape` (URL list) or `crawl`.
	//
	// Any of "scrape", "crawl".
	Mode string `json:"mode" api:"required"`
	// Individual page failures grouped by error code, sorted by count. Unrelated to
	// `failure`, which is the batch itself failing.
	PageErrors []PageErrorCount `json:"page_errors" api:"required"`
	// Pages attempted so far. Use `status` to check completion.
	Progress BatchListResponseDataProgress `json:"progress" api:"required"`
	// Result download links; null until the batch finishes. Files are deleted 180 days
	// after the batch finishes.
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
		Formats     respjson.Field
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

// Batch credit usage and settlement.
type BatchListResponseDataCredits struct {
	// `reserved` minus `refunded` plus `ocr_charged`.
	Net int64 `json:"net" api:"required"`
	// OCR usage charged when the batch settles.
	OcrCharged int64 `json:"ocr_charged" api:"required"`
	// Credits returned for unsuccessful pages when the batch settles.
	Refunded int64 `json:"refunded" api:"required"`
	// Credits held when the batch was accepted.
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

// Outputs requested for each page. Matches input.data.formats on submission.
type BatchListResponseDataFormats struct {
	// The original HTTP response body.
	Bytes bool `json:"bytes"`
	// Markdown excerpts relevant to `highlightsParams.query`.
	Highlights bool `json:"highlights"`
	// Rendered HTML.
	HTML bool `json:"html"`
	// Images found on the page.
	Images bool `json:"images"`
	// An object matching `jsonParams.schema`, extracted from the page.
	Json bool `json:"json"`
	// Page content as Markdown.
	Markdown bool `json:"markdown"`
	// Fields extracted with custom rules or a prebuilt parser, returned as `parsed`.
	Parse bool `json:"parse"`
	// Product details such as name, price, and availability.
	Product bool `json:"product"`
	// A screenshot of the page.
	Screenshot bool `json:"screenshot"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bytes       respjson.Field
		Highlights  respjson.Field
		HTML        respjson.Field
		Images      respjson.Field
		Json        respjson.Field
		Markdown    respjson.Field
		Parse       respjson.Field
		Product     respjson.Field
		Screenshot  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchListResponseDataFormats) RawJSON() string { return r.JSON.raw }
func (r *BatchListResponseDataFormats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pages attempted so far. Use `status` to check completion.
type BatchListResponseDataProgress struct {
	// Pages that could not be scraped.
	Failed int64 `json:"failed" api:"required"`
	// Accepted pages not yet attempted. Unused crawl capacity is excluded after
	// completion.
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

// Result download links; null until the batch finishes. Files are deleted 180 days
// after the batch finishes.
type BatchListResponseDataResults struct {
	// When these links expire (24 hours after this response).
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

// Credits this request used and your remaining balance.
type BatchListResponseKeyMetadata struct {
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
func (r BatchListResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *BatchListResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchDeleteResponse struct {
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// ID of the deleted batch.
	ID string `json:"id"`
	// Always true on success.
	Deleted bool `json:"deleted"`
	// Credits this request used and your remaining balance.
	KeyMetadata BatchDeleteResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestID   respjson.Field
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

// Credits this request used and your remaining balance.
type BatchDeleteResponseKeyMetadata struct {
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
func (r BatchDeleteResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *BatchDeleteResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchCancelResponse struct {
	// Batch ID.
	ID string `json:"id" api:"required"`
	// Crawl settings as submitted.
	Crawl CrawlControls `json:"crawl" api:"required"`
	// What this batch cost so far.
	Credits BatchCancelResponseCredits `json:"credits" api:"required"`
	// Outputs requested for each page. Matches input.data.formats on submission.
	Formats BatchCancelResponseFormats `json:"formats" api:"required"`
	// What the submission accepted.
	Input Intake `json:"input" api:"required"`
	// How pages were selected.
	//
	// Any of "scrape", "crawl".
	Mode BatchCancelResponseMode `json:"mode" api:"required"`
	// Page failures so far, grouped by error code and sorted by count.
	PageErrors []PageErrorCount `json:"page_errors" api:"required"`
	// How far the batch got before cancellation.
	Progress BatchCancelResponseProgress `json:"progress" api:"required"`
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Always `cancelling`. Work already in flight finishes; the batch reaches
	// `cancelled` shortly after.
	//
	// Any of "cancelling".
	Status BatchCancelResponseStatus `json:"status" api:"required"`
	// Tags stored on the batch at submission.
	Tags []string `json:"tags" api:"required"`
	// Batch timestamps.
	Timing BatchCancelResponseTiming `json:"timing" api:"required"`
	// API key usage for this request.
	KeyMetadata BatchCancelResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Crawl       respjson.Field
		Credits     respjson.Field
		Formats     respjson.Field
		Input       respjson.Field
		Mode        respjson.Field
		PageErrors  respjson.Field
		Progress    respjson.Field
		RequestID   respjson.Field
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
	// Credits held at submission; unused credits are refunded when the batch settles.
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

// Outputs requested for each page. Matches input.data.formats on submission.
type BatchCancelResponseFormats struct {
	// The original HTTP response body.
	Bytes bool `json:"bytes"`
	// Markdown excerpts relevant to `highlightsParams.query`.
	Highlights bool `json:"highlights"`
	// Rendered HTML.
	HTML bool `json:"html"`
	// Images found on the page.
	Images bool `json:"images"`
	// An object matching `jsonParams.schema`, extracted from the page.
	Json bool `json:"json"`
	// Page content as Markdown.
	Markdown bool `json:"markdown"`
	// Fields extracted with custom rules or a prebuilt parser, returned as `parsed`.
	Parse bool `json:"parse"`
	// Product details such as name, price, and availability.
	Product bool `json:"product"`
	// A screenshot of the page.
	Screenshot bool `json:"screenshot"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bytes       respjson.Field
		Highlights  respjson.Field
		HTML        respjson.Field
		Images      respjson.Field
		Json        respjson.Field
		Markdown    respjson.Field
		Parse       respjson.Field
		Product     respjson.Field
		Screenshot  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchCancelResponseFormats) RawJSON() string { return r.JSON.raw }
func (r *BatchCancelResponseFormats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
	// Pages that will be skipped.
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

// Batch timestamps.
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
func (r BatchCancelResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *BatchCancelResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchGetResultsResponse struct {
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Result records on this page.
	Data []BatchGetResultsResponseDataUnion `json:"data"`
	// Whether another page is available.
	HasMore bool `json:"has_more"`
	// Credits this request used and your remaining balance.
	KeyMetadata BatchGetResultsResponseKeyMetadata `json:"key_metadata"`
	// Cursor for the next page.
	NextCursor string `json:"next_cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestID   respjson.Field
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
	CreditsUsed int64 `json:"credits_used"`
	// This field is a union of [BatchGetResultsResponseDataOkResult],
	// [BatchGetResultsResponseDataErrorResult]
	Result BatchGetResultsResponseDataUnionResult `json:"result"`
	// Any of "ok", "error".
	Status   string `json:"status"`
	URL      string `json:"url"`
	ItemID   string `json:"itemId"`
	Meta     any    `json:"meta"`
	OcrPages int64  `json:"ocr_pages"`
	// This field is from variant [BatchGetResultsResponseDataError].
	ErrorCode string `json:"error_code"`
	// This field is from variant [BatchGetResultsResponseDataError].
	Message string `json:"message"`
	JSON    struct {
		CreditsUsed respjson.Field
		Result      respjson.Field
		Status      respjson.Field
		URL         respjson.Field
		ItemID      respjson.Field
		Meta        respjson.Field
		OcrPages    respjson.Field
		ErrorCode   respjson.Field
		Message     respjson.Field
		raw         string
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

// BatchGetResultsResponseDataUnionResult is an implicit subunion of
// [BatchGetResultsResponseDataUnion]. BatchGetResultsResponseDataUnionResult
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [BatchGetResultsResponseDataUnion].
type BatchGetResultsResponseDataUnionResult struct {
	// This field is a union of [BatchGetResultsResponseDataOkResultBytes],
	// [BatchGetResultsResponseDataErrorResultBytes]
	Bytes BatchGetResultsResponseDataUnionResultBytes `json:"bytes"`
	// This field is a union of [BatchGetResultsResponseDataOkResultCacheMetadata],
	// [BatchGetResultsResponseDataErrorResultCacheMetadata]
	CacheMetadata BatchGetResultsResponseDataUnionResultCacheMetadata `json:"cache_metadata"`
	// This field is a union of [BatchGetResultsResponseDataOkResultHighlights],
	// [BatchGetResultsResponseDataErrorResultHighlights]
	Highlights BatchGetResultsResponseDataUnionResultHighlights `json:"highlights"`
	// This field is a union of [BatchGetResultsResponseDataOkResultHTML],
	// [BatchGetResultsResponseDataErrorResultHTML]
	HTML BatchGetResultsResponseDataUnionResultHTML `json:"html"`
	// This field is a union of [BatchGetResultsResponseDataOkResultImages],
	// [BatchGetResultsResponseDataErrorResultImages]
	Images BatchGetResultsResponseDataUnionResultImages `json:"images"`
	// This field is a union of [BatchGetResultsResponseDataOkResultJson],
	// [BatchGetResultsResponseDataErrorResultJson]
	Json BatchGetResultsResponseDataUnionResultJson `json:"json"`
	// This field is a union of [BatchGetResultsResponseDataOkResultMarkdown],
	// [BatchGetResultsResponseDataErrorResultMarkdown]
	Markdown BatchGetResultsResponseDataUnionResultMarkdown `json:"markdown"`
	// This field is a union of [BatchGetResultsResponseDataOkResultMetadata],
	// [BatchGetResultsResponseDataErrorResultMetadata]
	Metadata BatchGetResultsResponseDataUnionResultMetadata `json:"metadata"`
	// This field is a union of [BatchGetResultsResponseDataOkResultParsed],
	// [BatchGetResultsResponseDataErrorResultParsed]
	Parsed BatchGetResultsResponseDataUnionResultParsed `json:"parsed"`
	// This field is a union of [BatchGetResultsResponseDataOkResultProduct],
	// [BatchGetResultsResponseDataErrorResultProduct]
	Product BatchGetResultsResponseDataUnionResultProduct `json:"product"`
	// This field is a union of [BatchGetResultsResponseDataOkResultProfile],
	// [BatchGetResultsResponseDataErrorResultProfile]
	Profile   BatchGetResultsResponseDataUnionResultProfile `json:"profile"`
	RequestID string                                        `json:"request_id"`
	// This field is a union of [BatchGetResultsResponseDataOkResultScreenshot],
	// [BatchGetResultsResponseDataErrorResultScreenshot]
	Screenshot BatchGetResultsResponseDataUnionResultScreenshot `json:"screenshot"`
	URL        string                                           `json:"url"`
	IsPartial  bool                                             `json:"isPartial"`
	// This field is a union of [BatchGetResultsResponseDataOkResultKeyMetadata],
	// [BatchGetResultsResponseDataErrorResultKeyMetadata]
	KeyMetadata BatchGetResultsResponseDataUnionResultKeyMetadata `json:"key_metadata"`
	JSON        struct {
		Bytes         respjson.Field
		CacheMetadata respjson.Field
		Highlights    respjson.Field
		HTML          respjson.Field
		Images        respjson.Field
		Json          respjson.Field
		Markdown      respjson.Field
		Metadata      respjson.Field
		Parsed        respjson.Field
		Product       respjson.Field
		Profile       respjson.Field
		RequestID     respjson.Field
		Screenshot    respjson.Field
		URL           respjson.Field
		IsPartial     respjson.Field
		KeyMetadata   respjson.Field
		raw           string
	} `json:"-"`
}

func (r *BatchGetResultsResponseDataUnionResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchGetResultsResponseDataUnionResultBytes is an implicit subunion of
// [BatchGetResultsResponseDataUnion]. BatchGetResultsResponseDataUnionResultBytes
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [BatchGetResultsResponseDataUnion].
type BatchGetResultsResponseDataUnionResultBytes struct {
	// This field is a union of [BatchGetResultsResponseDataOkResultBytesData],
	// [BatchGetResultsResponseDataErrorResultBytesData]
	Data      BatchGetResultsResponseDataUnionResultBytesData `json:"data"`
	Requested bool                                            `json:"requested"`
	Success   bool                                            `json:"success"`
	ErrorCode string                                          `json:"error_code"`
	Message   string                                          `json:"message"`
	JSON      struct {
		Data      respjson.Field
		Requested respjson.Field
		Success   respjson.Field
		ErrorCode respjson.Field
		Message   respjson.Field
		raw       string
	} `json:"-"`
}

func (r *BatchGetResultsResponseDataUnionResultBytes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchGetResultsResponseDataUnionResultBytesData is an implicit subunion of
// [BatchGetResultsResponseDataUnion].
// BatchGetResultsResponseDataUnionResultBytesData provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [BatchGetResultsResponseDataUnion].
type BatchGetResultsResponseDataUnionResultBytesData struct {
	Base64      string `json:"base64"`
	ContentType string `json:"contentType"`
	JSON        struct {
		Base64      respjson.Field
		ContentType respjson.Field
		raw         string
	} `json:"-"`
}

func (r *BatchGetResultsResponseDataUnionResultBytesData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchGetResultsResponseDataUnionResultCacheMetadata is an implicit subunion of
// [BatchGetResultsResponseDataUnion].
// BatchGetResultsResponseDataUnionResultCacheMetadata provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [BatchGetResultsResponseDataUnion].
type BatchGetResultsResponseDataUnionResultCacheMetadata struct {
	AgeMs  int64  `json:"age_ms"`
	Status string `json:"status"`
	JSON   struct {
		AgeMs  respjson.Field
		Status respjson.Field
		raw    string
	} `json:"-"`
}

func (r *BatchGetResultsResponseDataUnionResultCacheMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchGetResultsResponseDataUnionResultHighlights is an implicit subunion of
// [BatchGetResultsResponseDataUnion].
// BatchGetResultsResponseDataUnionResultHighlights provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [BatchGetResultsResponseDataUnion].
type BatchGetResultsResponseDataUnionResultHighlights struct {
	Data      []string `json:"data"`
	Requested bool     `json:"requested"`
	Success   bool     `json:"success"`
	ErrorCode string   `json:"error_code"`
	Message   string   `json:"message"`
	JSON      struct {
		Data      respjson.Field
		Requested respjson.Field
		Success   respjson.Field
		ErrorCode respjson.Field
		Message   respjson.Field
		raw       string
	} `json:"-"`
}

func (r *BatchGetResultsResponseDataUnionResultHighlights) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchGetResultsResponseDataUnionResultHTML is an implicit subunion of
// [BatchGetResultsResponseDataUnion]. BatchGetResultsResponseDataUnionResultHTML
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [BatchGetResultsResponseDataUnion].
type BatchGetResultsResponseDataUnionResultHTML struct {
	Data      string `json:"data"`
	Requested bool   `json:"requested"`
	Success   bool   `json:"success"`
	ErrorCode string `json:"error_code"`
	Message   string `json:"message"`
	JSON      struct {
		Data      respjson.Field
		Requested respjson.Field
		Success   respjson.Field
		ErrorCode respjson.Field
		Message   respjson.Field
		raw       string
	} `json:"-"`
}

func (r *BatchGetResultsResponseDataUnionResultHTML) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchGetResultsResponseDataUnionResultImages is an implicit subunion of
// [BatchGetResultsResponseDataUnion]. BatchGetResultsResponseDataUnionResultImages
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [BatchGetResultsResponseDataUnion].
type BatchGetResultsResponseDataUnionResultImages struct {
	// This field is a union of [[]BatchGetResultsResponseDataOkResultImagesData],
	// [[]BatchGetResultsResponseDataErrorResultImagesData]
	Data      BatchGetResultsResponseDataUnionResultImagesData `json:"data"`
	Requested bool                                             `json:"requested"`
	Success   bool                                             `json:"success"`
	ErrorCode string                                           `json:"error_code"`
	Message   string                                           `json:"message"`
	JSON      struct {
		Data      respjson.Field
		Requested respjson.Field
		Success   respjson.Field
		ErrorCode respjson.Field
		Message   respjson.Field
		raw       string
	} `json:"-"`
}

func (r *BatchGetResultsResponseDataUnionResultImages) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchGetResultsResponseDataUnionResultImagesData is an implicit subunion of
// [BatchGetResultsResponseDataUnion].
// BatchGetResultsResponseDataUnionResultImagesData provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [BatchGetResultsResponseDataUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBatchGetResultsResponseDataOkResultImagesData
// OfBatchGetResultsResponseDataErrorResultImagesData]
type BatchGetResultsResponseDataUnionResultImagesData struct {
	// This field will be present if the value is a
	// [[]BatchGetResultsResponseDataOkResultImagesData] instead of an object.
	OfBatchGetResultsResponseDataOkResultImagesData []BatchGetResultsResponseDataOkResultImagesData `json:",inline"`
	// This field will be present if the value is a
	// [[]BatchGetResultsResponseDataErrorResultImagesData] instead of an object.
	OfBatchGetResultsResponseDataErrorResultImagesData []BatchGetResultsResponseDataErrorResultImagesData `json:",inline"`
	JSON                                               struct {
		OfBatchGetResultsResponseDataOkResultImagesData    respjson.Field
		OfBatchGetResultsResponseDataErrorResultImagesData respjson.Field
		raw                                                string
	} `json:"-"`
}

func (r *BatchGetResultsResponseDataUnionResultImagesData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchGetResultsResponseDataUnionResultJson is an implicit subunion of
// [BatchGetResultsResponseDataUnion]. BatchGetResultsResponseDataUnionResultJson
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [BatchGetResultsResponseDataUnion].
type BatchGetResultsResponseDataUnionResultJson struct {
	Data        any    `json:"data"`
	Requested   bool   `json:"requested"`
	Success     bool   `json:"success"`
	ErrorCode   string `json:"error_code"`
	IsTruncated bool   `json:"isTruncated"`
	Message     string `json:"message"`
	JSON        struct {
		Data        respjson.Field
		Requested   respjson.Field
		Success     respjson.Field
		ErrorCode   respjson.Field
		IsTruncated respjson.Field
		Message     respjson.Field
		raw         string
	} `json:"-"`
}

func (r *BatchGetResultsResponseDataUnionResultJson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchGetResultsResponseDataUnionResultMarkdown is an implicit subunion of
// [BatchGetResultsResponseDataUnion].
// BatchGetResultsResponseDataUnionResultMarkdown provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [BatchGetResultsResponseDataUnion].
type BatchGetResultsResponseDataUnionResultMarkdown struct {
	Data      string `json:"data"`
	Requested bool   `json:"requested"`
	Success   bool   `json:"success"`
	ErrorCode string `json:"error_code"`
	Message   string `json:"message"`
	JSON      struct {
		Data      respjson.Field
		Requested respjson.Field
		Success   respjson.Field
		ErrorCode respjson.Field
		Message   respjson.Field
		raw       string
	} `json:"-"`
}

func (r *BatchGetResultsResponseDataUnionResultMarkdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchGetResultsResponseDataUnionResultMetadata is an implicit subunion of
// [BatchGetResultsResponseDataUnion].
// BatchGetResultsResponseDataUnionResultMetadata provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [BatchGetResultsResponseDataUnion].
type BatchGetResultsResponseDataUnionResultMetadata struct {
	// This field is a union of
	// [map[string]BatchGetResultsResponseDataOkResultMetadataAdditionalMetaUnion],
	// [map[string]BatchGetResultsResponseDataErrorResultMetadataAdditionalMetaUnion]
	AdditionalMeta BatchGetResultsResponseDataUnionResultMetadataAdditionalMeta `json:"additionalMeta"`
	// This field is a union of
	// [[]BatchGetResultsResponseDataOkResultMetadataAlternate],
	// [[]BatchGetResultsResponseDataErrorResultMetadataAlternate]
	Alternates   BatchGetResultsResponseDataUnionResultMetadataAlternates `json:"alternates"`
	Author       string                                                   `json:"author"`
	CanonicalURL string                                                   `json:"canonicalUrl"`
	Description  string                                                   `json:"description"`
	Favicon      string                                                   `json:"favicon"`
	// This field is a union of [[]BatchGetResultsResponseDataOkResultMetadataHeading],
	// [[]BatchGetResultsResponseDataErrorResultMetadataHeading]
	Headings     BatchGetResultsResponseDataUnionResultMetadataHeadings `json:"headings"`
	Image        string                                                 `json:"image"`
	JsonLd       []map[string]any                                       `json:"jsonLd"`
	Keywords     []string                                               `json:"keywords"`
	Language     string                                                 `json:"language"`
	ModifiedTime string                                                 `json:"modifiedTime"`
	// This field is a union of
	// [map[string]BatchGetResultsResponseDataOkResultMetadataOpenGraphUnion],
	// [map[string]BatchGetResultsResponseDataErrorResultMetadataOpenGraphUnion]
	OpenGraph     BatchGetResultsResponseDataUnionResultMetadataOpenGraph `json:"openGraph"`
	PublishedTime string                                                  `json:"publishedTime"`
	Robots        string                                                  `json:"robots"`
	SiteName      string                                                  `json:"siteName"`
	Title         string                                                  `json:"title"`
	// This field is a union of
	// [map[string]BatchGetResultsResponseDataOkResultMetadataTwitterUnion],
	// [map[string]BatchGetResultsResponseDataErrorResultMetadataTwitterUnion]
	Twitter BatchGetResultsResponseDataUnionResultMetadataTwitter `json:"twitter"`
	JSON    struct {
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
		raw            string
	} `json:"-"`
}

func (r *BatchGetResultsResponseDataUnionResultMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchGetResultsResponseDataUnionResultMetadataAdditionalMeta is an implicit
// subunion of [BatchGetResultsResponseDataUnion].
// BatchGetResultsResponseDataUnionResultMetadataAdditionalMeta provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [BatchGetResultsResponseDataUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type BatchGetResultsResponseDataUnionResultMetadataAdditionalMeta struct {
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

func (r *BatchGetResultsResponseDataUnionResultMetadataAdditionalMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchGetResultsResponseDataUnionResultMetadataAlternates is an implicit subunion
// of [BatchGetResultsResponseDataUnion].
// BatchGetResultsResponseDataUnionResultMetadataAlternates provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [BatchGetResultsResponseDataUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBatchGetResultsResponseDataOkResultMetadataAlternates
// OfBatchGetResultsResponseDataErrorResultMetadataAlternates]
type BatchGetResultsResponseDataUnionResultMetadataAlternates struct {
	// This field will be present if the value is a
	// [[]BatchGetResultsResponseDataOkResultMetadataAlternate] instead of an object.
	OfBatchGetResultsResponseDataOkResultMetadataAlternates []BatchGetResultsResponseDataOkResultMetadataAlternate `json:",inline"`
	// This field will be present if the value is a
	// [[]BatchGetResultsResponseDataErrorResultMetadataAlternate] instead of an
	// object.
	OfBatchGetResultsResponseDataErrorResultMetadataAlternates []BatchGetResultsResponseDataErrorResultMetadataAlternate `json:",inline"`
	JSON                                                       struct {
		OfBatchGetResultsResponseDataOkResultMetadataAlternates    respjson.Field
		OfBatchGetResultsResponseDataErrorResultMetadataAlternates respjson.Field
		raw                                                        string
	} `json:"-"`
}

func (r *BatchGetResultsResponseDataUnionResultMetadataAlternates) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchGetResultsResponseDataUnionResultMetadataHeadings is an implicit subunion
// of [BatchGetResultsResponseDataUnion].
// BatchGetResultsResponseDataUnionResultMetadataHeadings provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [BatchGetResultsResponseDataUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBatchGetResultsResponseDataOkResultMetadataHeadings
// OfBatchGetResultsResponseDataErrorResultMetadataHeadings]
type BatchGetResultsResponseDataUnionResultMetadataHeadings struct {
	// This field will be present if the value is a
	// [[]BatchGetResultsResponseDataOkResultMetadataHeading] instead of an object.
	OfBatchGetResultsResponseDataOkResultMetadataHeadings []BatchGetResultsResponseDataOkResultMetadataHeading `json:",inline"`
	// This field will be present if the value is a
	// [[]BatchGetResultsResponseDataErrorResultMetadataHeading] instead of an object.
	OfBatchGetResultsResponseDataErrorResultMetadataHeadings []BatchGetResultsResponseDataErrorResultMetadataHeading `json:",inline"`
	JSON                                                     struct {
		OfBatchGetResultsResponseDataOkResultMetadataHeadings    respjson.Field
		OfBatchGetResultsResponseDataErrorResultMetadataHeadings respjson.Field
		raw                                                      string
	} `json:"-"`
}

func (r *BatchGetResultsResponseDataUnionResultMetadataHeadings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchGetResultsResponseDataUnionResultMetadataOpenGraph is an implicit subunion
// of [BatchGetResultsResponseDataUnion].
// BatchGetResultsResponseDataUnionResultMetadataOpenGraph provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [BatchGetResultsResponseDataUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type BatchGetResultsResponseDataUnionResultMetadataOpenGraph struct {
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

func (r *BatchGetResultsResponseDataUnionResultMetadataOpenGraph) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchGetResultsResponseDataUnionResultMetadataTwitter is an implicit subunion of
// [BatchGetResultsResponseDataUnion].
// BatchGetResultsResponseDataUnionResultMetadataTwitter provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [BatchGetResultsResponseDataUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type BatchGetResultsResponseDataUnionResultMetadataTwitter struct {
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

func (r *BatchGetResultsResponseDataUnionResultMetadataTwitter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchGetResultsResponseDataUnionResultParsed is an implicit subunion of
// [BatchGetResultsResponseDataUnion]. BatchGetResultsResponseDataUnionResultParsed
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [BatchGetResultsResponseDataUnion].
type BatchGetResultsResponseDataUnionResultParsed struct {
	Data      any    `json:"data"`
	Requested bool   `json:"requested"`
	Success   bool   `json:"success"`
	Type      string `json:"type"`
	ErrorCode string `json:"error_code"`
	Message   string `json:"message"`
	JSON      struct {
		Data      respjson.Field
		Requested respjson.Field
		Success   respjson.Field
		Type      respjson.Field
		ErrorCode respjson.Field
		Message   respjson.Field
		raw       string
	} `json:"-"`
}

func (r *BatchGetResultsResponseDataUnionResultParsed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchGetResultsResponseDataUnionResultProduct is an implicit subunion of
// [BatchGetResultsResponseDataUnion].
// BatchGetResultsResponseDataUnionResultProduct provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [BatchGetResultsResponseDataUnion].
type BatchGetResultsResponseDataUnionResultProduct struct {
	// This field is a union of [BatchGetResultsResponseDataOkResultProductData],
	// [BatchGetResultsResponseDataErrorResultProductData]
	Data      BatchGetResultsResponseDataUnionResultProductData `json:"data"`
	Requested bool                                              `json:"requested"`
	Success   bool                                              `json:"success"`
	ErrorCode string                                            `json:"error_code"`
	Message   string                                            `json:"message"`
	JSON      struct {
		Data      respjson.Field
		Requested respjson.Field
		Success   respjson.Field
		ErrorCode respjson.Field
		Message   respjson.Field
		raw       string
	} `json:"-"`
}

func (r *BatchGetResultsResponseDataUnionResultProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchGetResultsResponseDataUnionResultProductData is an implicit subunion of
// [BatchGetResultsResponseDataUnion].
// BatchGetResultsResponseDataUnionResultProductData provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [BatchGetResultsResponseDataUnion].
type BatchGetResultsResponseDataUnionResultProductData struct {
	IsProductPage bool `json:"isProductPage"`
	// This field is a union of
	// [BatchGetResultsResponseDataOkResultProductDataProduct],
	// [BatchGetResultsResponseDataErrorResultProductDataProduct]
	Product BatchGetResultsResponseDataUnionResultProductDataProduct `json:"product"`
	JSON    struct {
		IsProductPage respjson.Field
		Product       respjson.Field
		raw           string
	} `json:"-"`
}

func (r *BatchGetResultsResponseDataUnionResultProductData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchGetResultsResponseDataUnionResultProductDataProduct is an implicit subunion
// of [BatchGetResultsResponseDataUnion].
// BatchGetResultsResponseDataUnionResultProductDataProduct provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [BatchGetResultsResponseDataUnion].
type BatchGetResultsResponseDataUnionResultProductDataProduct struct {
	Availability   string   `json:"availability"`
	Brand          string   `json:"brand"`
	Category       string   `json:"category"`
	Currency       string   `json:"currency"`
	Description    string   `json:"description"`
	Dimensions     []string `json:"dimensions"`
	Features       []string `json:"features"`
	Images         []string `json:"images"`
	ImageURL       string   `json:"imageUrl"`
	Name           string   `json:"name"`
	Price          float64  `json:"price"`
	RegularPrice   float64  `json:"regularPrice"`
	SKU            string   `json:"sku"`
	Tags           []string `json:"tags"`
	TargetAudience []string `json:"targetAudience"`
	// This field is a union of
	// [[]BatchGetResultsResponseDataOkResultProductDataProductVariant],
	// [[]BatchGetResultsResponseDataErrorResultProductDataProductVariant]
	Variants BatchGetResultsResponseDataUnionResultProductDataProductVariants `json:"variants"`
	JSON     struct {
		Availability   respjson.Field
		Brand          respjson.Field
		Category       respjson.Field
		Currency       respjson.Field
		Description    respjson.Field
		Dimensions     respjson.Field
		Features       respjson.Field
		Images         respjson.Field
		ImageURL       respjson.Field
		Name           respjson.Field
		Price          respjson.Field
		RegularPrice   respjson.Field
		SKU            respjson.Field
		Tags           respjson.Field
		TargetAudience respjson.Field
		Variants       respjson.Field
		raw            string
	} `json:"-"`
}

func (r *BatchGetResultsResponseDataUnionResultProductDataProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchGetResultsResponseDataUnionResultProductDataProductVariants is an implicit
// subunion of [BatchGetResultsResponseDataUnion].
// BatchGetResultsResponseDataUnionResultProductDataProductVariants provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [BatchGetResultsResponseDataUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBatchGetResultsResponseDataOkResultProductDataProductVariants
// OfBatchGetResultsResponseDataErrorResultProductDataProductVariants]
type BatchGetResultsResponseDataUnionResultProductDataProductVariants struct {
	// This field will be present if the value is a
	// [[]BatchGetResultsResponseDataOkResultProductDataProductVariant] instead of an
	// object.
	OfBatchGetResultsResponseDataOkResultProductDataProductVariants []BatchGetResultsResponseDataOkResultProductDataProductVariant `json:",inline"`
	// This field will be present if the value is a
	// [[]BatchGetResultsResponseDataErrorResultProductDataProductVariant] instead of
	// an object.
	OfBatchGetResultsResponseDataErrorResultProductDataProductVariants []BatchGetResultsResponseDataErrorResultProductDataProductVariant `json:",inline"`
	JSON                                                               struct {
		OfBatchGetResultsResponseDataOkResultProductDataProductVariants    respjson.Field
		OfBatchGetResultsResponseDataErrorResultProductDataProductVariants respjson.Field
		raw                                                                string
	} `json:"-"`
}

func (r *BatchGetResultsResponseDataUnionResultProductDataProductVariants) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchGetResultsResponseDataUnionResultProfile is an implicit subunion of
// [BatchGetResultsResponseDataUnion].
// BatchGetResultsResponseDataUnionResultProfile provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [BatchGetResultsResponseDataUnion].
type BatchGetResultsResponseDataUnionResultProfile struct {
	ID   string `json:"id"`
	JSON struct {
		ID  respjson.Field
		raw string
	} `json:"-"`
}

func (r *BatchGetResultsResponseDataUnionResultProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchGetResultsResponseDataUnionResultScreenshot is an implicit subunion of
// [BatchGetResultsResponseDataUnion].
// BatchGetResultsResponseDataUnionResultScreenshot provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [BatchGetResultsResponseDataUnion].
type BatchGetResultsResponseDataUnionResultScreenshot struct {
	Data      string `json:"data"`
	Requested bool   `json:"requested"`
	Success   bool   `json:"success"`
	ErrorCode string `json:"error_code"`
	Message   string `json:"message"`
	JSON      struct {
		Data      respjson.Field
		Requested respjson.Field
		Success   respjson.Field
		ErrorCode respjson.Field
		Message   respjson.Field
		raw       string
	} `json:"-"`
}

func (r *BatchGetResultsResponseDataUnionResultScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchGetResultsResponseDataUnionResultKeyMetadata is an implicit subunion of
// [BatchGetResultsResponseDataUnion].
// BatchGetResultsResponseDataUnionResultKeyMetadata provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [BatchGetResultsResponseDataUnion].
type BatchGetResultsResponseDataUnionResultKeyMetadata struct {
	CreditsConsumed  int64 `json:"credits_consumed"`
	CreditsRemaining int64 `json:"credits_remaining"`
	JSON             struct {
		CreditsConsumed  respjson.Field
		CreditsRemaining respjson.Field
		raw              string
	} `json:"-"`
}

func (r *BatchGetResultsResponseDataUnionResultKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// At least one requested output succeeded. Check result.isPartial and each output
// success for incomplete results.
type BatchGetResultsResponseDataOk struct {
	// Credits charged for this page, including OCR. Cached outputs use the same
	// pricing as POST /web/scrape.
	CreditsUsed int64 `json:"credits_used" api:"required"`
	// The same response as POST /web/scrape, including per-format success, failure,
	// and partial results.
	Result BatchGetResultsResponseDataOkResult `json:"result" api:"required"`
	Status constant.Ok                         `json:"status" default:"ok"`
	// URL as submitted, or as discovered by the crawl.
	URL string `json:"url" api:"required"`
	// Your `itemId` from submission.
	ItemID string `json:"itemId"`
	// Caller-supplied metadata echoed from submission.
	Meta map[string]any `json:"meta"`
	// PDF pages recovered by OCR. Omitted when OCR did not run.
	OcrPages int64 `json:"ocr_pages"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditsUsed respjson.Field
		Result      respjson.Field
		Status      respjson.Field
		URL         respjson.Field
		ItemID      respjson.Field
		Meta        respjson.Field
		OcrPages    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataOk) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataOk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The same response as POST /web/scrape, including per-format success, failure,
// and partial results.
type BatchGetResultsResponseDataOkResult struct {
	// The original HTTP response body, unchanged by waits, actions, and filters.
	Bytes BatchGetResultsResponseDataOkResultBytes `json:"bytes" api:"required"`
	// Whether this response came from cache.
	CacheMetadata BatchGetResultsResponseDataOkResultCacheMetadata `json:"cache_metadata" api:"required"`
	// Relevant Markdown excerpts in page order. `[Heading]` adds context; `…` marks
	// omitted text.
	Highlights BatchGetResultsResponseDataOkResultHighlights `json:"highlights" api:"required"`
	// Rendered HTML after content filters.
	HTML BatchGetResultsResponseDataOkResultHTML `json:"html" api:"required"`
	// Images after content filters. `[]` when none are found.
	Images BatchGetResultsResponseDataOkResultImages `json:"images" api:"required"`
	// Object matching `jsonParams.schema`.
	Json BatchGetResultsResponseDataOkResultJson `json:"json" api:"required"`
	// Markdown after content filters.
	Markdown BatchGetResultsResponseDataOkResultMarkdown `json:"markdown" api:"required"`
	// Page metadata. Fields are omitted when not found.
	Metadata BatchGetResultsResponseDataOkResultMetadata `json:"metadata" api:"required"`
	// Custom rules return fields after content filters; unmatched fields are null ([]
	// for lists). The google-maps parser returns business records in data.results,
	// google-maps-reviews returns data.reviews, and the google-play parser returns one
	// app record as data, google-play-reviews returns data.reviews, app-store and
	// app-store-reviews do the same for the App Store, g2-reviews returns data.reviews
	// for a G2 product, emails, phone-numbers, and links return data.emails,
	// data.phone_numbers, and data.links from the rendered page; the store, Maps, and
	// G2 parsers read their data directly, independently of DOM filters and other
	// outputs.
	Parsed BatchGetResultsResponseDataOkResultParsed `json:"parsed" api:"required"`
	// Product details found on the page.
	Product BatchGetResultsResponseDataOkResultProduct `json:"product" api:"required"`
	// Authenticated browser profile used for this scrape, or null when no profile was
	// requested. New and existing profile requests both return the profile ID.
	Profile BatchGetResultsResponseDataOkResultProfile `json:"profile" api:"required"`
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Screenshot as a base64 image data URL.
	Screenshot BatchGetResultsResponseDataOkResultScreenshot `json:"screenshot" api:"required"`
	// Final URL after redirects and browser actions.
	URL string `json:"url" api:"required" format:"uri"`
	// True when at least one requested output succeeds but the response has failed or
	// incomplete outputs. Absent when all requested outputs fail.
	//
	// Any of true.
	IsPartial bool `json:"isPartial"`
	// Credits this request used and your remaining balance.
	KeyMetadata BatchGetResultsResponseDataOkResultKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bytes         respjson.Field
		CacheMetadata respjson.Field
		Highlights    respjson.Field
		HTML          respjson.Field
		Images        respjson.Field
		Json          respjson.Field
		Markdown      respjson.Field
		Metadata      respjson.Field
		Parsed        respjson.Field
		Product       respjson.Field
		Profile       respjson.Field
		RequestID     respjson.Field
		Screenshot    respjson.Field
		URL           respjson.Field
		IsPartial     respjson.Field
		KeyMetadata   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataOkResult) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataOkResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The original HTTP response body, unchanged by waits, actions, and filters.
type BatchGetResultsResponseDataOkResultBytes struct {
	Data      BatchGetResultsResponseDataOkResultBytesData `json:"data" api:"required"`
	Requested bool                                         `json:"requested" api:"required"`
	// `true` if returned, `false` if it failed, `null` if not requested.
	Success bool `json:"success" api:"required"`
	// Why the output failed. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Explanation of the failure and possible next steps.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		Success     respjson.Field
		ErrorCode   respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataOkResultBytes) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataOkResultBytes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchGetResultsResponseDataOkResultBytesData struct {
	// Body as base64, after HTTP decompression. Up to 50 MiB decoded.
	Base64      string `json:"base64" api:"required"`
	ContentType string `json:"contentType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Base64      respjson.Field
		ContentType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataOkResultBytesData) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataOkResultBytesData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether this response came from cache.
type BatchGetResultsResponseDataOkResultCacheMetadata struct {
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
func (r BatchGetResultsResponseDataOkResultCacheMetadata) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataOkResultCacheMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Relevant Markdown excerpts in page order. `[Heading]` adds context; `…` marks
// omitted text.
type BatchGetResultsResponseDataOkResultHighlights struct {
	Data      []string `json:"data" api:"required"`
	Requested bool     `json:"requested" api:"required"`
	// `true` if returned, `false` if it failed, `null` if not requested.
	Success bool `json:"success" api:"required"`
	// Why the output failed. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Explanation of the failure and possible next steps.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		Success     respjson.Field
		ErrorCode   respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataOkResultHighlights) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataOkResultHighlights) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rendered HTML after content filters.
type BatchGetResultsResponseDataOkResultHTML struct {
	Data      string `json:"data" api:"required"`
	Requested bool   `json:"requested" api:"required"`
	// `true` if returned, `false` if it failed, `null` if not requested.
	Success bool `json:"success" api:"required"`
	// Why the output failed. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Explanation of the failure and possible next steps.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		Success     respjson.Field
		ErrorCode   respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataOkResultHTML) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataOkResultHTML) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Images after content filters. `[]` when none are found.
type BatchGetResultsResponseDataOkResultImages struct {
	Data      []BatchGetResultsResponseDataOkResultImagesData `json:"data" api:"required"`
	Requested bool                                            `json:"requested" api:"required"`
	// `true` if returned, `false` if it failed, `null` if not requested.
	Success bool `json:"success" api:"required"`
	// Why the output failed. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Explanation of the failure and possible next steps.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		Success     respjson.Field
		ErrorCode   respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataOkResultImages) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataOkResultImages) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchGetResultsResponseDataOkResultImagesData struct {
	// Alt text, if present.
	Alt string `json:"alt" api:"required"`
	// Image URL, or a data URI for inline images.
	URL string `json:"url" api:"required" format:"uri"`
	// Any of "photography", "illustration", "logo", "wordmark", "icon", "pattern",
	// "graphic", "other".
	Classification string `json:"classification"`
	// Hosted image URL, valid for 24 hours after capture. Requires `file` enrichment
	// and ZDR disabled.
	FileURL string `json:"fileUrl" format:"uri"`
	Height  int64  `json:"height"`
	Width   int64  `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Alt            respjson.Field
		URL            respjson.Field
		Classification respjson.Field
		FileURL        respjson.Field
		Height         respjson.Field
		Width          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataOkResultImagesData) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataOkResultImagesData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Object matching `jsonParams.schema`.
type BatchGetResultsResponseDataOkResultJson struct {
	Data      map[string]any `json:"data" api:"required"`
	Requested bool           `json:"requested" api:"required"`
	// `true` if returned, `false` if it failed, `null` if not requested.
	Success bool `json:"success" api:"required"`
	// Why the output failed. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// True when the page was too long to read in full, so values found only in the
	// unread parts may be missing.
	//
	// Any of true.
	IsTruncated bool `json:"isTruncated"`
	// Explanation of the failure and possible next steps.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		Success     respjson.Field
		ErrorCode   respjson.Field
		IsTruncated respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataOkResultJson) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataOkResultJson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Markdown after content filters.
type BatchGetResultsResponseDataOkResultMarkdown struct {
	Data      string `json:"data" api:"required"`
	Requested bool   `json:"requested" api:"required"`
	// `true` if returned, `false` if it failed, `null` if not requested.
	Success bool `json:"success" api:"required"`
	// Why the output failed. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Explanation of the failure and possible next steps.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		Success     respjson.Field
		ErrorCode   respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataOkResultMarkdown) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataOkResultMarkdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Page metadata. Fields are omitted when not found.
type BatchGetResultsResponseDataOkResultMetadata struct {
	// Additional non-social meta tags not promoted to top-level metadata fields.
	AdditionalMeta map[string]BatchGetResultsResponseDataOkResultMetadataAdditionalMetaUnion `json:"additionalMeta"`
	// Resolved alternate links from link rel=alternate tags.
	Alternates []BatchGetResultsResponseDataOkResultMetadataAlternate `json:"alternates"`
	// Author metadata, when present.
	Author string `json:"author"`
	// Resolved canonical URL, when present.
	CanonicalURL string `json:"canonicalUrl"`
	// Best description extracted from standard, Open Graph, or Twitter metadata.
	Description string `json:"description"`
	// Resolved favicon URL, when present.
	Favicon string `json:"favicon"`
	// Up to 500 h1–h6 headings in document order, before content filtering.
	Headings []BatchGetResultsResponseDataOkResultMetadataHeading `json:"headings"`
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
	OpenGraph map[string]BatchGetResultsResponseDataOkResultMetadataOpenGraphUnion `json:"openGraph"`
	// Published timestamp/date from page metadata, when present.
	PublishedTime string `json:"publishedTime"`
	// Robots meta directive, when present.
	Robots string `json:"robots"`
	// Site or application name from page metadata.
	SiteName string `json:"siteName"`
	// Best title extracted from the page.
	Title string `json:"title"`
	// Twitter card metadata with the twitter: prefix removed and keys camel-cased.
	Twitter map[string]BatchGetResultsResponseDataOkResultMetadataTwitterUnion `json:"twitter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
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
func (r BatchGetResultsResponseDataOkResultMetadata) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataOkResultMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchGetResultsResponseDataOkResultMetadataAdditionalMetaUnion contains all
// possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type BatchGetResultsResponseDataOkResultMetadataAdditionalMetaUnion struct {
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

func (u BatchGetResultsResponseDataOkResultMetadataAdditionalMetaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BatchGetResultsResponseDataOkResultMetadataAdditionalMetaUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BatchGetResultsResponseDataOkResultMetadataAdditionalMetaUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *BatchGetResultsResponseDataOkResultMetadataAdditionalMetaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchGetResultsResponseDataOkResultMetadataAlternate struct {
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
func (r BatchGetResultsResponseDataOkResultMetadataAlternate) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataOkResultMetadataAlternate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchGetResultsResponseDataOkResultMetadataHeading struct {
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
func (r BatchGetResultsResponseDataOkResultMetadataHeading) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataOkResultMetadataHeading) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchGetResultsResponseDataOkResultMetadataOpenGraphUnion contains all possible
// properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type BatchGetResultsResponseDataOkResultMetadataOpenGraphUnion struct {
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

func (u BatchGetResultsResponseDataOkResultMetadataOpenGraphUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BatchGetResultsResponseDataOkResultMetadataOpenGraphUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BatchGetResultsResponseDataOkResultMetadataOpenGraphUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *BatchGetResultsResponseDataOkResultMetadataOpenGraphUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchGetResultsResponseDataOkResultMetadataTwitterUnion contains all possible
// properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type BatchGetResultsResponseDataOkResultMetadataTwitterUnion struct {
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

func (u BatchGetResultsResponseDataOkResultMetadataTwitterUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BatchGetResultsResponseDataOkResultMetadataTwitterUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BatchGetResultsResponseDataOkResultMetadataTwitterUnion) RawJSON() string { return u.JSON.raw }

func (r *BatchGetResultsResponseDataOkResultMetadataTwitterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom rules return fields after content filters; unmatched fields are null ([]
// for lists). The google-maps parser returns business records in data.results,
// google-maps-reviews returns data.reviews, and the google-play parser returns one
// app record as data, google-play-reviews returns data.reviews, app-store and
// app-store-reviews do the same for the App Store, g2-reviews returns data.reviews
// for a G2 product, emails, phone-numbers, and links return data.emails,
// data.phone_numbers, and data.links from the rendered page; the store, Maps, and
// G2 parsers read their data directly, independently of DOM filters and other
// outputs.
type BatchGetResultsResponseDataOkResultParsed struct {
	Data      map[string]any `json:"data" api:"required"`
	Requested bool           `json:"requested" api:"required"`
	// `true` if returned, `false` if it failed, `null` if not requested.
	Success bool `json:"success" api:"required"`
	// The selected prebuilt parser name, or custom when type is omitted, custom rules
	// are used, or parsing is not requested.
	//
	// Any of "custom", "google-maps", "google-maps-reviews", "google-play",
	// "google-play-reviews", "app-store", "app-store-reviews", "g2-reviews", "emails",
	// "phone-numbers", "links".
	Type string `json:"type" api:"required"`
	// Why the output failed. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Explanation of the failure and possible next steps.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		Success     respjson.Field
		Type        respjson.Field
		ErrorCode   respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataOkResultParsed) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataOkResultParsed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Product details found on the page.
type BatchGetResultsResponseDataOkResultProduct struct {
	Data      BatchGetResultsResponseDataOkResultProductData `json:"data" api:"required"`
	Requested bool                                           `json:"requested" api:"required"`
	// `true` if returned, `false` if it failed, `null` if not requested.
	Success bool `json:"success" api:"required"`
	// Why the output failed. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Explanation of the failure and possible next steps.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		Success     respjson.Field
		ErrorCode   respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataOkResultProduct) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataOkResultProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchGetResultsResponseDataOkResultProductData struct {
	// Whether the page is a product detail page.
	IsProductPage bool `json:"isProductPage" api:"required"`
	// The extracted product, or null when the page is not a product detail page.
	Product BatchGetResultsResponseDataOkResultProductDataProduct `json:"product" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsProductPage respjson.Field
		Product       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataOkResultProductData) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataOkResultProductData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The extracted product, or null when the page is not a product detail page.
type BatchGetResultsResponseDataOkResultProductDataProduct struct {
	// Stock or ordering availability.
	//
	// Any of "in_stock", "out_of_stock", "limited_availability", "preorder",
	// "backorder", "made_to_order", "discontinued".
	Availability string `json:"availability" api:"required"`
	// Brand or vendor.
	Brand string `json:"brand" api:"required"`
	// Product category.
	Category string `json:"category" api:"required"`
	// ISO 4217 currency code.
	Currency string `json:"currency" api:"required"`
	// Product description.
	Description string `json:"description" api:"required"`
	// Product dimensions as shown on the page.
	Dimensions []string `json:"dimensions" api:"required"`
	// Key features and specifications.
	Features []string `json:"features" api:"required"`
	// Product image URLs, main image first.
	Images []string `json:"images" api:"required" format:"uri"`
	// Main product image URL.
	ImageURL string `json:"imageUrl" api:"required" format:"uri"`
	// Product name.
	Name string `json:"name" api:"required"`
	// Current price.
	Price float64 `json:"price" api:"required"`
	// List price before any discount.
	RegularPrice float64 `json:"regularPrice" api:"required"`
	// Product identifier such as a SKU or model number.
	SKU string `json:"sku" api:"required"`
	// Product tags.
	Tags []string `json:"tags" api:"required"`
	// Intended audience.
	TargetAudience []string `json:"targetAudience" api:"required"`
	// Product variations, such as different colors or sizes, with their attributes and
	// images. Empty if none are found. May not include every variation offered by the
	// store.
	Variants []BatchGetResultsResponseDataOkResultProductDataProductVariant `json:"variants" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Availability   respjson.Field
		Brand          respjson.Field
		Category       respjson.Field
		Currency       respjson.Field
		Description    respjson.Field
		Dimensions     respjson.Field
		Features       respjson.Field
		Images         respjson.Field
		ImageURL       respjson.Field
		Name           respjson.Field
		Price          respjson.Field
		RegularPrice   respjson.Field
		SKU            respjson.Field
		Tags           respjson.Field
		TargetAudience respjson.Field
		Variants       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataOkResultProductDataProduct) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataOkResultProductDataProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchGetResultsResponseDataOkResultProductDataProductVariant struct {
	// Explicit variant attributes such as color, size, material, pattern and
	// properties declared by page.
	Attributes map[string]string `json:"attributes" api:"required"`
	// Original source image URLs explicitly attached to this variant.
	Images []string `json:"images" api:"required"`
	SKU    string   `json:"sku" api:"required"`
	// Variant or offer URL when provided by the source. May be shared by variants.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attributes  respjson.Field
		Images      respjson.Field
		SKU         respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataOkResultProductDataProductVariant) RawJSON() string {
	return r.JSON.raw
}
func (r *BatchGetResultsResponseDataOkResultProductDataProductVariant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Authenticated browser profile used for this scrape, or null when no profile was
// requested. New and existing profile requests both return the profile ID.
type BatchGetResultsResponseDataOkResultProfile struct {
	// Save this ID after creating a profile, then reuse it as
	// sharedParams.profile.profileId with type "existing".
	ID string `json:"id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataOkResultProfile) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataOkResultProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot as a base64 image data URL.
type BatchGetResultsResponseDataOkResultScreenshot struct {
	Data      string `json:"data" api:"required" format:"uri"`
	Requested bool   `json:"requested" api:"required"`
	// `true` if returned, `false` if it failed, `null` if not requested.
	Success bool `json:"success" api:"required"`
	// Why the output failed. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Explanation of the failure and possible next steps.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		Success     respjson.Field
		ErrorCode   respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataOkResultScreenshot) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataOkResultScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credits this request used and your remaining balance.
type BatchGetResultsResponseDataOkResultKeyMetadata struct {
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
func (r BatchGetResultsResponseDataOkResultKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataOkResultKeyMetadata) UnmarshalJSON(data []byte) error {
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
	// Credits charged for this page, including OCR. Cached outputs use the same
	// pricing as POST /web/scrape.
	CreditsUsed int64 `json:"credits_used"`
	// Your `itemId` from submission.
	ItemID string `json:"itemId"`
	// Caller-supplied metadata echoed from submission.
	Meta map[string]any `json:"meta"`
	// PDF pages recovered by OCR. Omitted when OCR did not run.
	OcrPages int64 `json:"ocr_pages"`
	// Per-format failures when scraping returned a structured response. Omitted for
	// errors that prevent scraping.
	Result BatchGetResultsResponseDataErrorResult `json:"result"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ErrorCode   respjson.Field
		Message     respjson.Field
		Status      respjson.Field
		URL         respjson.Field
		CreditsUsed respjson.Field
		ItemID      respjson.Field
		Meta        respjson.Field
		OcrPages    respjson.Field
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataError) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-format failures when scraping returned a structured response. Omitted for
// errors that prevent scraping.
type BatchGetResultsResponseDataErrorResult struct {
	// The original HTTP response body, unchanged by waits, actions, and filters.
	Bytes BatchGetResultsResponseDataErrorResultBytes `json:"bytes" api:"required"`
	// Whether this response came from cache.
	CacheMetadata BatchGetResultsResponseDataErrorResultCacheMetadata `json:"cache_metadata" api:"required"`
	// Relevant Markdown excerpts in page order. `[Heading]` adds context; `…` marks
	// omitted text.
	Highlights BatchGetResultsResponseDataErrorResultHighlights `json:"highlights" api:"required"`
	// Rendered HTML after content filters.
	HTML BatchGetResultsResponseDataErrorResultHTML `json:"html" api:"required"`
	// Images after content filters. `[]` when none are found.
	Images BatchGetResultsResponseDataErrorResultImages `json:"images" api:"required"`
	// Object matching `jsonParams.schema`.
	Json BatchGetResultsResponseDataErrorResultJson `json:"json" api:"required"`
	// Markdown after content filters.
	Markdown BatchGetResultsResponseDataErrorResultMarkdown `json:"markdown" api:"required"`
	// Page metadata. Fields are omitted when not found.
	Metadata BatchGetResultsResponseDataErrorResultMetadata `json:"metadata" api:"required"`
	// Custom rules return fields after content filters; unmatched fields are null ([]
	// for lists). The google-maps parser returns business records in data.results,
	// google-maps-reviews returns data.reviews, and the google-play parser returns one
	// app record as data, google-play-reviews returns data.reviews, app-store and
	// app-store-reviews do the same for the App Store, g2-reviews returns data.reviews
	// for a G2 product, emails, phone-numbers, and links return data.emails,
	// data.phone_numbers, and data.links from the rendered page; the store, Maps, and
	// G2 parsers read their data directly, independently of DOM filters and other
	// outputs.
	Parsed BatchGetResultsResponseDataErrorResultParsed `json:"parsed" api:"required"`
	// Product details found on the page.
	Product BatchGetResultsResponseDataErrorResultProduct `json:"product" api:"required"`
	// Authenticated browser profile used for this scrape, or null when no profile was
	// requested. New and existing profile requests both return the profile ID.
	Profile BatchGetResultsResponseDataErrorResultProfile `json:"profile" api:"required"`
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Screenshot as a base64 image data URL.
	Screenshot BatchGetResultsResponseDataErrorResultScreenshot `json:"screenshot" api:"required"`
	// Final URL after redirects and browser actions.
	URL string `json:"url" api:"required" format:"uri"`
	// True when at least one requested output succeeds but the response has failed or
	// incomplete outputs. Absent when all requested outputs fail.
	//
	// Any of true.
	IsPartial bool `json:"isPartial"`
	// Credits this request used and your remaining balance.
	KeyMetadata BatchGetResultsResponseDataErrorResultKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bytes         respjson.Field
		CacheMetadata respjson.Field
		Highlights    respjson.Field
		HTML          respjson.Field
		Images        respjson.Field
		Json          respjson.Field
		Markdown      respjson.Field
		Metadata      respjson.Field
		Parsed        respjson.Field
		Product       respjson.Field
		Profile       respjson.Field
		RequestID     respjson.Field
		Screenshot    respjson.Field
		URL           respjson.Field
		IsPartial     respjson.Field
		KeyMetadata   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataErrorResult) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataErrorResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The original HTTP response body, unchanged by waits, actions, and filters.
type BatchGetResultsResponseDataErrorResultBytes struct {
	Data      BatchGetResultsResponseDataErrorResultBytesData `json:"data" api:"required"`
	Requested bool                                            `json:"requested" api:"required"`
	// `true` if returned, `false` if it failed, `null` if not requested.
	Success bool `json:"success" api:"required"`
	// Why the output failed. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Explanation of the failure and possible next steps.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		Success     respjson.Field
		ErrorCode   respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataErrorResultBytes) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataErrorResultBytes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchGetResultsResponseDataErrorResultBytesData struct {
	// Body as base64, after HTTP decompression. Up to 50 MiB decoded.
	Base64      string `json:"base64" api:"required"`
	ContentType string `json:"contentType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Base64      respjson.Field
		ContentType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataErrorResultBytesData) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataErrorResultBytesData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether this response came from cache.
type BatchGetResultsResponseDataErrorResultCacheMetadata struct {
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
func (r BatchGetResultsResponseDataErrorResultCacheMetadata) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataErrorResultCacheMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Relevant Markdown excerpts in page order. `[Heading]` adds context; `…` marks
// omitted text.
type BatchGetResultsResponseDataErrorResultHighlights struct {
	Data      []string `json:"data" api:"required"`
	Requested bool     `json:"requested" api:"required"`
	// `true` if returned, `false` if it failed, `null` if not requested.
	Success bool `json:"success" api:"required"`
	// Why the output failed. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Explanation of the failure and possible next steps.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		Success     respjson.Field
		ErrorCode   respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataErrorResultHighlights) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataErrorResultHighlights) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rendered HTML after content filters.
type BatchGetResultsResponseDataErrorResultHTML struct {
	Data      string `json:"data" api:"required"`
	Requested bool   `json:"requested" api:"required"`
	// `true` if returned, `false` if it failed, `null` if not requested.
	Success bool `json:"success" api:"required"`
	// Why the output failed. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Explanation of the failure and possible next steps.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		Success     respjson.Field
		ErrorCode   respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataErrorResultHTML) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataErrorResultHTML) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Images after content filters. `[]` when none are found.
type BatchGetResultsResponseDataErrorResultImages struct {
	Data      []BatchGetResultsResponseDataErrorResultImagesData `json:"data" api:"required"`
	Requested bool                                               `json:"requested" api:"required"`
	// `true` if returned, `false` if it failed, `null` if not requested.
	Success bool `json:"success" api:"required"`
	// Why the output failed. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Explanation of the failure and possible next steps.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		Success     respjson.Field
		ErrorCode   respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataErrorResultImages) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataErrorResultImages) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchGetResultsResponseDataErrorResultImagesData struct {
	// Alt text, if present.
	Alt string `json:"alt" api:"required"`
	// Image URL, or a data URI for inline images.
	URL string `json:"url" api:"required" format:"uri"`
	// Any of "photography", "illustration", "logo", "wordmark", "icon", "pattern",
	// "graphic", "other".
	Classification string `json:"classification"`
	// Hosted image URL, valid for 24 hours after capture. Requires `file` enrichment
	// and ZDR disabled.
	FileURL string `json:"fileUrl" format:"uri"`
	Height  int64  `json:"height"`
	Width   int64  `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Alt            respjson.Field
		URL            respjson.Field
		Classification respjson.Field
		FileURL        respjson.Field
		Height         respjson.Field
		Width          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataErrorResultImagesData) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataErrorResultImagesData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Object matching `jsonParams.schema`.
type BatchGetResultsResponseDataErrorResultJson struct {
	Data      map[string]any `json:"data" api:"required"`
	Requested bool           `json:"requested" api:"required"`
	// `true` if returned, `false` if it failed, `null` if not requested.
	Success bool `json:"success" api:"required"`
	// Why the output failed. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// True when the page was too long to read in full, so values found only in the
	// unread parts may be missing.
	//
	// Any of true.
	IsTruncated bool `json:"isTruncated"`
	// Explanation of the failure and possible next steps.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		Success     respjson.Field
		ErrorCode   respjson.Field
		IsTruncated respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataErrorResultJson) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataErrorResultJson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Markdown after content filters.
type BatchGetResultsResponseDataErrorResultMarkdown struct {
	Data      string `json:"data" api:"required"`
	Requested bool   `json:"requested" api:"required"`
	// `true` if returned, `false` if it failed, `null` if not requested.
	Success bool `json:"success" api:"required"`
	// Why the output failed. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Explanation of the failure and possible next steps.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		Success     respjson.Field
		ErrorCode   respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataErrorResultMarkdown) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataErrorResultMarkdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Page metadata. Fields are omitted when not found.
type BatchGetResultsResponseDataErrorResultMetadata struct {
	// Additional non-social meta tags not promoted to top-level metadata fields.
	AdditionalMeta map[string]BatchGetResultsResponseDataErrorResultMetadataAdditionalMetaUnion `json:"additionalMeta"`
	// Resolved alternate links from link rel=alternate tags.
	Alternates []BatchGetResultsResponseDataErrorResultMetadataAlternate `json:"alternates"`
	// Author metadata, when present.
	Author string `json:"author"`
	// Resolved canonical URL, when present.
	CanonicalURL string `json:"canonicalUrl"`
	// Best description extracted from standard, Open Graph, or Twitter metadata.
	Description string `json:"description"`
	// Resolved favicon URL, when present.
	Favicon string `json:"favicon"`
	// Up to 500 h1–h6 headings in document order, before content filtering.
	Headings []BatchGetResultsResponseDataErrorResultMetadataHeading `json:"headings"`
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
	OpenGraph map[string]BatchGetResultsResponseDataErrorResultMetadataOpenGraphUnion `json:"openGraph"`
	// Published timestamp/date from page metadata, when present.
	PublishedTime string `json:"publishedTime"`
	// Robots meta directive, when present.
	Robots string `json:"robots"`
	// Site or application name from page metadata.
	SiteName string `json:"siteName"`
	// Best title extracted from the page.
	Title string `json:"title"`
	// Twitter card metadata with the twitter: prefix removed and keys camel-cased.
	Twitter map[string]BatchGetResultsResponseDataErrorResultMetadataTwitterUnion `json:"twitter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
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
func (r BatchGetResultsResponseDataErrorResultMetadata) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataErrorResultMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchGetResultsResponseDataErrorResultMetadataAdditionalMetaUnion contains all
// possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type BatchGetResultsResponseDataErrorResultMetadataAdditionalMetaUnion struct {
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

func (u BatchGetResultsResponseDataErrorResultMetadataAdditionalMetaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BatchGetResultsResponseDataErrorResultMetadataAdditionalMetaUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BatchGetResultsResponseDataErrorResultMetadataAdditionalMetaUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *BatchGetResultsResponseDataErrorResultMetadataAdditionalMetaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchGetResultsResponseDataErrorResultMetadataAlternate struct {
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
func (r BatchGetResultsResponseDataErrorResultMetadataAlternate) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataErrorResultMetadataAlternate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchGetResultsResponseDataErrorResultMetadataHeading struct {
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
func (r BatchGetResultsResponseDataErrorResultMetadataHeading) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataErrorResultMetadataHeading) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchGetResultsResponseDataErrorResultMetadataOpenGraphUnion contains all
// possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type BatchGetResultsResponseDataErrorResultMetadataOpenGraphUnion struct {
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

func (u BatchGetResultsResponseDataErrorResultMetadataOpenGraphUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BatchGetResultsResponseDataErrorResultMetadataOpenGraphUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BatchGetResultsResponseDataErrorResultMetadataOpenGraphUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *BatchGetResultsResponseDataErrorResultMetadataOpenGraphUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BatchGetResultsResponseDataErrorResultMetadataTwitterUnion contains all possible
// properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type BatchGetResultsResponseDataErrorResultMetadataTwitterUnion struct {
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

func (u BatchGetResultsResponseDataErrorResultMetadataTwitterUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BatchGetResultsResponseDataErrorResultMetadataTwitterUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BatchGetResultsResponseDataErrorResultMetadataTwitterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *BatchGetResultsResponseDataErrorResultMetadataTwitterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom rules return fields after content filters; unmatched fields are null ([]
// for lists). The google-maps parser returns business records in data.results,
// google-maps-reviews returns data.reviews, and the google-play parser returns one
// app record as data, google-play-reviews returns data.reviews, app-store and
// app-store-reviews do the same for the App Store, g2-reviews returns data.reviews
// for a G2 product, emails, phone-numbers, and links return data.emails,
// data.phone_numbers, and data.links from the rendered page; the store, Maps, and
// G2 parsers read their data directly, independently of DOM filters and other
// outputs.
type BatchGetResultsResponseDataErrorResultParsed struct {
	Data      map[string]any `json:"data" api:"required"`
	Requested bool           `json:"requested" api:"required"`
	// `true` if returned, `false` if it failed, `null` if not requested.
	Success bool `json:"success" api:"required"`
	// The selected prebuilt parser name, or custom when type is omitted, custom rules
	// are used, or parsing is not requested.
	//
	// Any of "custom", "google-maps", "google-maps-reviews", "google-play",
	// "google-play-reviews", "app-store", "app-store-reviews", "g2-reviews", "emails",
	// "phone-numbers", "links".
	Type string `json:"type" api:"required"`
	// Why the output failed. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Explanation of the failure and possible next steps.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		Success     respjson.Field
		Type        respjson.Field
		ErrorCode   respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataErrorResultParsed) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataErrorResultParsed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Product details found on the page.
type BatchGetResultsResponseDataErrorResultProduct struct {
	Data      BatchGetResultsResponseDataErrorResultProductData `json:"data" api:"required"`
	Requested bool                                              `json:"requested" api:"required"`
	// `true` if returned, `false` if it failed, `null` if not requested.
	Success bool `json:"success" api:"required"`
	// Why the output failed. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Explanation of the failure and possible next steps.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		Success     respjson.Field
		ErrorCode   respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataErrorResultProduct) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataErrorResultProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchGetResultsResponseDataErrorResultProductData struct {
	// Whether the page is a product detail page.
	IsProductPage bool `json:"isProductPage" api:"required"`
	// The extracted product, or null when the page is not a product detail page.
	Product BatchGetResultsResponseDataErrorResultProductDataProduct `json:"product" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsProductPage respjson.Field
		Product       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataErrorResultProductData) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataErrorResultProductData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The extracted product, or null when the page is not a product detail page.
type BatchGetResultsResponseDataErrorResultProductDataProduct struct {
	// Stock or ordering availability.
	//
	// Any of "in_stock", "out_of_stock", "limited_availability", "preorder",
	// "backorder", "made_to_order", "discontinued".
	Availability string `json:"availability" api:"required"`
	// Brand or vendor.
	Brand string `json:"brand" api:"required"`
	// Product category.
	Category string `json:"category" api:"required"`
	// ISO 4217 currency code.
	Currency string `json:"currency" api:"required"`
	// Product description.
	Description string `json:"description" api:"required"`
	// Product dimensions as shown on the page.
	Dimensions []string `json:"dimensions" api:"required"`
	// Key features and specifications.
	Features []string `json:"features" api:"required"`
	// Product image URLs, main image first.
	Images []string `json:"images" api:"required" format:"uri"`
	// Main product image URL.
	ImageURL string `json:"imageUrl" api:"required" format:"uri"`
	// Product name.
	Name string `json:"name" api:"required"`
	// Current price.
	Price float64 `json:"price" api:"required"`
	// List price before any discount.
	RegularPrice float64 `json:"regularPrice" api:"required"`
	// Product identifier such as a SKU or model number.
	SKU string `json:"sku" api:"required"`
	// Product tags.
	Tags []string `json:"tags" api:"required"`
	// Intended audience.
	TargetAudience []string `json:"targetAudience" api:"required"`
	// Product variations, such as different colors or sizes, with their attributes and
	// images. Empty if none are found. May not include every variation offered by the
	// store.
	Variants []BatchGetResultsResponseDataErrorResultProductDataProductVariant `json:"variants" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Availability   respjson.Field
		Brand          respjson.Field
		Category       respjson.Field
		Currency       respjson.Field
		Description    respjson.Field
		Dimensions     respjson.Field
		Features       respjson.Field
		Images         respjson.Field
		ImageURL       respjson.Field
		Name           respjson.Field
		Price          respjson.Field
		RegularPrice   respjson.Field
		SKU            respjson.Field
		Tags           respjson.Field
		TargetAudience respjson.Field
		Variants       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataErrorResultProductDataProduct) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataErrorResultProductDataProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchGetResultsResponseDataErrorResultProductDataProductVariant struct {
	// Explicit variant attributes such as color, size, material, pattern and
	// properties declared by page.
	Attributes map[string]string `json:"attributes" api:"required"`
	// Original source image URLs explicitly attached to this variant.
	Images []string `json:"images" api:"required"`
	SKU    string   `json:"sku" api:"required"`
	// Variant or offer URL when provided by the source. May be shared by variants.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attributes  respjson.Field
		Images      respjson.Field
		SKU         respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataErrorResultProductDataProductVariant) RawJSON() string {
	return r.JSON.raw
}
func (r *BatchGetResultsResponseDataErrorResultProductDataProductVariant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Authenticated browser profile used for this scrape, or null when no profile was
// requested. New and existing profile requests both return the profile ID.
type BatchGetResultsResponseDataErrorResultProfile struct {
	// Save this ID after creating a profile, then reuse it as
	// sharedParams.profile.profileId with type "existing".
	ID string `json:"id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataErrorResultProfile) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataErrorResultProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot as a base64 image data URL.
type BatchGetResultsResponseDataErrorResultScreenshot struct {
	Data      string `json:"data" api:"required" format:"uri"`
	Requested bool   `json:"requested" api:"required"`
	// `true` if returned, `false` if it failed, `null` if not requested.
	Success bool `json:"success" api:"required"`
	// Why the output failed. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Explanation of the failure and possible next steps.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		Success     respjson.Field
		ErrorCode   respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataErrorResultScreenshot) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataErrorResultScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credits this request used and your remaining balance.
type BatchGetResultsResponseDataErrorResultKeyMetadata struct {
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
func (r BatchGetResultsResponseDataErrorResultKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataErrorResultKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credits this request used and your remaining balance.
type BatchGetResultsResponseKeyMetadata struct {
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
func (r BatchGetResultsResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchSubmitResponse struct {
	// Batch ID. Poll GET /batch/{batch_id} with it.
	ID string `json:"id" api:"required"`
	// Whether this response came from cache.
	CacheMetadata BatchSubmitResponseCacheMetadata `json:"cache_metadata" api:"required"`
	// Crawl settings as submitted.
	Crawl CrawlControls `json:"crawl" api:"required"`
	// When the batch was created.
	CreatedAt string `json:"created_at" api:"required"`
	// What accepting this batch cost.
	Credits BatchSubmitResponseCredits `json:"credits" api:"required"`
	// Outputs requested for each page. Matches input.data.formats on submission.
	Formats BatchSubmitResponseFormats `json:"formats" api:"required"`
	// What the submission accepted.
	Input Intake `json:"input" api:"required"`
	// Rejected URLs (first 100).
	InvalidURLs []BatchSubmitResponseInvalidURL `json:"invalid_urls" api:"required"`
	// How pages will be selected.
	//
	// Any of "scrape", "crawl".
	Mode BatchSubmitResponseMode `json:"mode" api:"required"`
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Always `queued`. An accepted batch has not started yet.
	//
	// Any of "queued".
	Status BatchSubmitResponseStatus `json:"status" api:"required"`
	// Tags stored on the batch.
	Tags []string `json:"tags" api:"required"`
	// API key usage for this request.
	KeyMetadata BatchSubmitResponseKeyMetadata `json:"key_metadata"`
	// Secret for verifying `X-Context-Signature`. Only submit returns it, so store it.
	WebhookSecret string `json:"webhook_secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CacheMetadata respjson.Field
		Crawl         respjson.Field
		CreatedAt     respjson.Field
		Credits       respjson.Field
		Formats       respjson.Field
		Input         respjson.Field
		InvalidURLs   respjson.Field
		Mode          respjson.Field
		RequestID     respjson.Field
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

// Whether this response came from cache.
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
	// Credits held at submission.
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

// Outputs requested for each page. Matches input.data.formats on submission.
type BatchSubmitResponseFormats struct {
	// The original HTTP response body.
	Bytes bool `json:"bytes"`
	// Markdown excerpts relevant to `highlightsParams.query`.
	Highlights bool `json:"highlights"`
	// Rendered HTML.
	HTML bool `json:"html"`
	// Images found on the page.
	Images bool `json:"images"`
	// An object matching `jsonParams.schema`, extracted from the page.
	Json bool `json:"json"`
	// Page content as Markdown.
	Markdown bool `json:"markdown"`
	// Fields extracted with custom rules or a prebuilt parser, returned as `parsed`.
	Parse bool `json:"parse"`
	// Product details such as name, price, and availability.
	Product bool `json:"product"`
	// A screenshot of the page.
	Screenshot bool `json:"screenshot"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bytes       respjson.Field
		Highlights  respjson.Field
		HTML        respjson.Field
		Images      respjson.Field
		Json        respjson.Field
		Markdown    respjson.Field
		Parse       respjson.Field
		Product     respjson.Field
		Screenshot  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchSubmitResponseFormats) RawJSON() string { return r.JSON.raw }
func (r *BatchSubmitResponseFormats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
	// Include batches submitted with Scrape Anything formats. Defaults to false.
	ShouldIncludeUnified param.Opt[bool] `query:"shouldIncludeUnified,omitzero" json:"-"`
	// `prefix` for as-you-type prefix matching (default), `exact` for full-token
	// matching.
	//
	// Any of "exact", "prefix".
	SearchType BatchListParamsSearchType `query:"search_type,omitzero" json:"-"`
	// Filter by status.
	//
	// Any of "queued", "running", "cancelling", "completed", "cancelled", "failed".
	Status BatchListParamsStatus `query:"status,omitzero" json:"-"`
	// Tags to filter by (matches batches having any of them). Pass repeated `tags`
	// params or one comma-separated list, e.g. `tags=docs,competitor`.
	Tags BatchListParamsTagsUnion `query:"tags,omitzero" json:"-"`
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchListParamsTagsUnion struct {
	OfString      param.Opt[string] `query:",omitzero,inline"`
	OfStringArray []string          `query:",omitzero,inline"`
	paramUnion
}

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
	// Legacy URL notified when the batch finishes. Preserves one best-effort attempt.
	// Cannot be combined with webhook.
	WebhookURL param.Opt[string] `json:"webhookUrl,omitzero" format:"uri"`
	// Unique key per submission. Retrying with the same key and body returns the
	// original batch; a different body returns `409`.
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	// Tags stored on the batch. Filter the batch list by them later.
	Tags []string `json:"tags,omitzero"`
	// Where to send the batch's final-status event. Omit `retry` for one attempt; `{}`
	// uses the default retry schedule.
	Webhook BatchSubmitParamsWebhook `json:"webhook,omitzero"`
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

// Scrape a list of up to 100,000 URLs.
//
// The properties Data, Mode are required.
type BatchSubmitParamsInputScrape struct {
	// Use the same formats and parameters as POST /web/scrape, replacing url with
	// urls. Options apply to every page.
	Data BatchSubmitParamsInputScrapeData `json:"data,omitzero" api:"required"`
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

// Use the same formats and parameters as POST /web/scrape, replacing url with
// urls. Options apply to every page.
//
// The properties Formats, URLs are required.
type BatchSubmitParamsInputScrapeData struct {
	// Outputs to return. Set at least one to `true`.
	Formats BatchSubmitParamsInputScrapeDataFormats `json:"formats,omitzero" api:"required"`
	// Pages to scrape. Maximum 100000.
	URLs []BatchSubmitParamsInputScrapeDataURL `json:"urls,omitzero" api:"required"`
	// Maximum age of a cached output, in milliseconds. `0` fetches fresh. Defaults to
	// 3 days (259200000 ms). Maximum: 1 year (31536000000 ms).
	MaxAgeMs param.Opt[int64] `json:"maxAgeMs,omitzero"`
	// Requires `formats.highlights: true`; required when it is set.
	HighlightsParams BatchSubmitParamsInputScrapeDataHighlightsParams `json:"highlightsParams,omitzero"`
	// Image options. Requires formats.images: true.
	ImageParams BatchSubmitParamsInputScrapeDataImageParams `json:"imageParams,omitzero"`
	// Requires `formats.json: true`; required when it is set.
	JsonParams BatchSubmitParamsInputScrapeDataJsonParams `json:"jsonParams,omitzero"`
	// Markdown options. Requires `formats.markdown`.
	MarkdownParams BatchSubmitParamsInputScrapeDataMarkdownParams `json:"markdownParams,omitzero"`
	// Requires `formats.parse: true`; required when it is set.
	ParseParams BatchSubmitParamsInputScrapeDataParseParamsUnion `json:"parseParams,omitzero"`
	// Product options. Requires formats.product: true.
	ProductParams BatchSubmitParamsInputScrapeDataProductParams `json:"productParams,omitzero"`
	// Screenshot options. Requires formats.screenshot: true.
	ScreenshotParams BatchSubmitParamsInputScrapeDataScreenshotParams `json:"screenshotParams,omitzero"`
	// Browser and content settings shared by all outputs.
	SharedParams BatchSubmitParamsInputScrapeDataSharedParams `json:"sharedParams,omitzero"`
	// Labels for tracking request usage. Not retained when zdr is enabled.
	Tags []string `json:"tags,omitzero"`
	// Deadline for each page, using the same behavior and defaults as POST
	// /web/scrape.
	TimeoutOpts BatchSubmitParamsInputScrapeDataTimeoutOpts `json:"timeoutOpts,omitzero"`
	// Batches persist their results. Only disabled is supported; enabled returns
	// ZDR_NOT_SUPPORTED.
	//
	// Any of "disabled".
	Zdr string `json:"zdr,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeData) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputScrapeData](
		"zdr", "disabled",
	)
}

// Outputs to return. Set at least one to `true`.
type BatchSubmitParamsInputScrapeDataFormats struct {
	// The original HTTP response body.
	Bytes param.Opt[bool] `json:"bytes,omitzero"`
	// Markdown excerpts relevant to `highlightsParams.query`.
	Highlights param.Opt[bool] `json:"highlights,omitzero"`
	// Rendered HTML.
	HTML param.Opt[bool] `json:"html,omitzero"`
	// Images found on the page.
	Images param.Opt[bool] `json:"images,omitzero"`
	// An object matching `jsonParams.schema`, extracted from the page.
	Json param.Opt[bool] `json:"json,omitzero"`
	// Page content as Markdown.
	Markdown param.Opt[bool] `json:"markdown,omitzero"`
	// Fields extracted with custom rules or a prebuilt parser, returned as `parsed`.
	Parse param.Opt[bool] `json:"parse,omitzero"`
	// Product details such as name, price, and availability.
	Product param.Opt[bool] `json:"product,omitzero"`
	// A screenshot of the page.
	Screenshot param.Opt[bool] `json:"screenshot,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataFormats) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataFormats
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataFormats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A page to scrape, with optional data for matching results.
//
// The property URL is required.
type BatchSubmitParamsInputScrapeDataURL struct {
	// Page URL to scrape.
	URL string `json:"url" api:"required"`
	// Your ID for this page, returned with its result. The same URL can use different
	// IDs.
	ItemID param.Opt[string] `json:"itemId,omitzero"`
	// Custom JSON returned unchanged with this page result.
	Meta map[string]any `json:"meta,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataURL) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Requires `formats.highlights: true`; required when it is set.
//
// The property Query is required.
type BatchSubmitParamsInputScrapeDataHighlightsParams struct {
	// The question or topic to find passages for.
	Query string `json:"query" api:"required"`
	// Maximum combined length of returned passages.
	MaxCharacters param.Opt[int64] `json:"maxCharacters,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataHighlightsParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataHighlightsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataHighlightsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image options. Requires formats.images: true.
type BatchSubmitParamsInputScrapeDataImageParams struct {
	// Set `visual` to drop visual duplicates, keeping the largest copy.
	//
	// Any of "none", "visual".
	Dedupe string `json:"dedupe,omitzero"`
	// Extra data per image: `dimensions`, `classification`, or a hosted `file` URL.
	//
	// Any of "dimensions", "classification", "file".
	Enrich []string `json:"enrich,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataImageParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataImageParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataImageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputScrapeDataImageParams](
		"dedupe", "none", "visual",
	)
}

// Requires `formats.json: true`; required when it is set.
//
// The property Schema is required.
type BatchSubmitParamsInputScrapeDataJsonParams struct {
	// JSON Schema (not an example object) for a top-level object, up to 50 KB. Use
	// optional or nullable fields for missing facts.
	Schema map[string]any `json:"schema,omitzero" api:"required"`
	// Extra guidance, such as which facts to prefer or how to read a field.
	Instructions param.Opt[string] `json:"instructions,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataJsonParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataJsonParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataJsonParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Markdown options. Requires `formats.markdown`.
type BatchSubmitParamsInputScrapeDataMarkdownParams struct {
	// Include images in the Markdown using image syntax with URLs and alt text.
	IncludeImages param.Opt[bool] `json:"includeImages,omitzero"`
	// Keep link URLs in the Markdown. Set false to return link text without URLs.
	IncludeLinks param.Opt[bool] `json:"includeLinks,omitzero"`
	// How base64 images appear: `placeholder` (default) or `preserve`. Requires
	// `includeImages`.
	//
	// Any of "placeholder", "preserve".
	InlineImages string `json:"inlineImages,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataMarkdownParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataMarkdownParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataMarkdownParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputScrapeDataMarkdownParams](
		"inlineImages", "placeholder", "preserve",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputScrapeDataParseParamsUnion struct {
	OfCustomParser   *BatchSubmitParamsInputScrapeDataParseParamsCustomParser   `json:",omitzero,inline"`
	OfPrebuiltParser *BatchSubmitParamsInputScrapeDataParseParamsPrebuiltParser `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputScrapeDataParseParamsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCustomParser, u.OfPrebuiltParser)
}
func (u *BatchSubmitParamsInputScrapeDataParseParamsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property Rules is required.
type BatchSubmitParamsInputScrapeDataParseParamsCustomParser struct {
	// Field names mapped to CSS selectors (`h1`, `a@href`) or rule objects. Max 100
	// fields, 5 levels.
	Rules map[string]BatchSubmitParamsInputScrapeDataParseParamsCustomParserRuleUnion `json:"rules,omitzero" api:"required"`
	// Extract fields using custom CSS rules. This is the default when type is omitted.
	//
	// Any of "custom".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataParseParamsCustomParser) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataParseParamsCustomParser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataParseParamsCustomParser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputScrapeDataParseParamsCustomParser](
		"type", "custom",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputScrapeDataParseParamsCustomParserRuleUnion struct {
	OfString                                                       param.Opt[string]                                                  `json:",omitzero,inline"`
	OfBatchSubmitsInputScrapeDataParseParamsCustomParserRuleObject *BatchSubmitParamsInputScrapeDataParseParamsCustomParserRuleObject `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputScrapeDataParseParamsCustomParserRuleUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfBatchSubmitsInputScrapeDataParseParamsCustomParserRuleObject)
}
func (u *BatchSubmitParamsInputScrapeDataParseParamsCustomParserRuleUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property Selector is required.
type BatchSubmitParamsInputScrapeDataParseParamsCustomParserRuleObject struct {
	// CSS selector to match within the current page or parent rule.
	Selector string `json:"selector" api:"required"`
	// Return text, HTML, an attribute such as `@href`, or nested field rules. Defaults
	// to text.
	Output string `json:"output,omitzero"`
	// Return the first match with `item` or all matches with `list`.
	//
	// Any of "item", "list".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataParseParamsCustomParserRuleObject) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataParseParamsCustomParserRuleObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataParseParamsCustomParserRuleObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputScrapeDataParseParamsCustomParserRuleObject](
		"type", "item", "list",
	)
}

// The properties Parser, Type are required.
type BatchSubmitParamsInputScrapeDataParseParamsPrebuiltParser struct {
	// google-maps parses Google Maps business data into data.results. Accepts Maps
	// place, search, and internal data URLs. Uses a US exit and English labels
	// regardless of hl or Accept-Language. google-maps-reviews returns the reviews of
	// a place in data.reviews, in the same shape as the Serper reviews API (rating,
	// date, isoDate, snippet, translatedSnippet, likes, user, media, response, link,
	// id), plus language, guided ratings in details, and user.isLocalGuide; data also
	// carries the place featureId, cid, and sort. sort and limit pick the order and
	// count, and oldest is exact. It accepts Maps place, search, and ?cid= URLs, plus
	// search.google.com/local/reviews URLs, so the google_maps_url and
	// google_reviews_url fields of google-maps results work as input. An unknown place
	// returns no reviews. google-play parses a Play Store app into data: details,
	// developer contact, ratings, price, permissions, and data safety.
	// google-play-reviews returns the 200 newest reviews of the app, with developer
	// replies, in data.reviews. Both Play parsers accept app details URLs
	// (play.google.com/store/apps/details?id=…); hl and gl set the language and
	// storefront, defaulting to en and us. app-store parses an App Store app into
	// data: details, developer, ratings, price, version, and screenshots.
	// app-store-reviews returns up to the 500 newest reviews in data.reviews. Both App
	// Store parsers accept app URLs (apps.apple.com/us/app/…/id…); the country in the
	// path sets the storefront, defaulting to us. g2-reviews returns the 25 most
	// recently published reviews of a G2 product in data.reviews (rating, title,
	// likes, dislikes, problems_solved, reviewer role, company size, and industry),
	// plus data.product. It accepts G2 product URLs (g2.com/products/…), follows
	// renamed products, and keeps star-rating filters from the URL
	// (filters[nps_score][]=1 to 5); other filters, sort orders, and pages are
	// ignored. emails accepts any page and returns its email addresses as data.emails
	// ([{ email }]), read from mailto links, Cloudflare-protected addresses, and
	// visible text. phone-numbers accepts any page and returns its valid phone numbers
	// in E.164 as data.phone_numbers ([{ phone }]), read from tel links and visible
	// text; numbers without a country code are read for the country of the site TLD,
	// defaulting to US. links accepts any page and returns its unique http(s) links as
	// data.links ([{ url, text }]), with absolute URLs; text falls back to aria-label
	// or image alt and is null when a link has none. The emails parser reads the
	// rendered page and respects content filters; the other prebuilt parsers read
	// their data independently of DOM filters.
	//
	// Any of "google-maps", "google-maps-reviews", "google-play",
	// "google-play-reviews", "app-store", "app-store-reviews", "g2-reviews", "emails",
	// "phone-numbers", "links".
	Parser string `json:"parser,omitzero" api:"required"`
	// google-maps-reviews only. Maximum number of reviews to return. Defaults to 100;
	// maximum 1000.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// google-maps-reviews only. Order of data.reviews. newest, relevant, highest, and
	// lowest follow Google; oldest returns the earliest reviews first, which Google
	// itself does not offer.
	//
	// Any of "newest", "oldest", "relevant", "highest", "lowest".
	Sort string `json:"sort,omitzero"`
	// Use a prebuilt parser instead of custom rules.
	//
	// This field can be elided, and will marshal its zero value as "prebuilt".
	Type constant.Prebuilt `json:"type" default:"prebuilt"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataParseParamsPrebuiltParser) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataParseParamsPrebuiltParser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataParseParamsPrebuiltParser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputScrapeDataParseParamsPrebuiltParser](
		"parser", "google-maps", "google-maps-reviews", "google-play", "google-play-reviews", "app-store", "app-store-reviews", "g2-reviews", "emails", "phone-numbers", "links",
	)
	apijson.RegisterFieldValidator[BatchSubmitParamsInputScrapeDataParseParamsPrebuiltParser](
		"sort", "newest", "oldest", "relevant", "highest", "lowest",
	)
}

// Product options. Requires formats.product: true.
type BatchSubmitParamsInputScrapeDataProductParams struct {
	// Drop visually duplicate product images, keeping the largest copy.
	DedupeImages param.Opt[bool] `json:"dedupeImages,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataProductParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataProductParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataProductParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot options. Requires formats.screenshot: true.
type BatchSubmitParamsInputScrapeDataScreenshotParams struct {
	// What to capture: `viewport`, `fullPage`, one element, or a rectangle. Max 40
	// megapixels.
	Area BatchSubmitParamsInputScrapeDataScreenshotParamsAreaUnion `json:"area,omitzero"`
	// Image format for the screenshot.
	//
	// Any of "png", "jpeg", "webp".
	Format string `json:"format,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataScreenshotParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataScreenshotParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataScreenshotParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputScrapeDataScreenshotParams](
		"format", "png", "jpeg", "webp",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputScrapeDataScreenshotParamsAreaUnion struct {
	// Check if union is this variant with !param.IsOmitted(union.OfPage)
	OfPage      param.Opt[string]                                              `json:",omitzero,inline"`
	OfElement   *BatchSubmitParamsInputScrapeDataScreenshotParamsAreaElement   `json:",omitzero,inline"`
	OfRectangle *BatchSubmitParamsInputScrapeDataScreenshotParamsAreaRectangle `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputScrapeDataScreenshotParamsAreaUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPage, u.OfElement, u.OfRectangle)
}
func (u *BatchSubmitParamsInputScrapeDataScreenshotParamsAreaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type BatchSubmitParamsInputScrapeDataScreenshotParamsAreaPage string

const (
	BatchSubmitParamsInputScrapeDataScreenshotParamsAreaPageViewport BatchSubmitParamsInputScrapeDataScreenshotParamsAreaPage = "viewport"
	BatchSubmitParamsInputScrapeDataScreenshotParamsAreaPageFullPage BatchSubmitParamsInputScrapeDataScreenshotParamsAreaPage = "fullPage"
)

// The property Selector is required.
type BatchSubmitParamsInputScrapeDataScreenshotParamsAreaElement struct {
	// CSS selector matching exactly one visible element.
	Selector string `json:"selector" api:"required"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataScreenshotParamsAreaElement) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataScreenshotParamsAreaElement
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataScreenshotParamsAreaElement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pixels from the document origin.
//
// The properties Height, Width, X, Y are required.
type BatchSubmitParamsInputScrapeDataScreenshotParamsAreaRectangle struct {
	// Height of the capture in pixels.
	Height int64 `json:"height" api:"required"`
	// Width of the capture in pixels.
	Width int64 `json:"width" api:"required"`
	// Left edge of the capture, in pixels from the document origin.
	X int64 `json:"x" api:"required"`
	// Top edge of the capture, in pixels from the document origin.
	Y int64 `json:"y" api:"required"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataScreenshotParamsAreaRectangle) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataScreenshotParamsAreaRectangle
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataScreenshotParamsAreaRectangle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Browser and content settings shared by all outputs.
type BatchSubmitParamsInputScrapeDataSharedParams struct {
	// Proxy country as a two-letter code, such as `US`. Case-insensitive.
	Country param.Opt[string] `json:"country,omitzero"`
	// Accept cookie banners before actions and capture.
	DismissCookies param.Opt[bool] `json:"dismissCookies,omitzero"`
	// Close other popups before actions and capture.
	DismissPopups param.Opt[bool] `json:"dismissPopups,omitzero"`
	// Include iframe content in HTML and text outputs. Screenshots always show visible
	// frames.
	IncludeFrames param.Opt[bool] `json:"includeFrames,omitzero"`
	// Keep only the main content. Doesn't affect `screenshot`, `bytes`, or `product`.
	MainContentOnly param.Opt[bool] `json:"mainContentOnly,omitzero"`
	// Wait for CSS animations to finish before capture. Defaults to `true` when
	// `screenshot` is requested.
	SettleAnimations param.Opt[bool] `json:"settleAnimations,omitzero"`
	// Browser steps run in order before capture. Requires a paid plan. Skips the
	// cache.
	Actions []BatchSubmitParamsInputScrapeDataSharedParamsActionUnion `json:"actions,omitzero"`
	// Remove elements matching these CSS selectors. Overrides `includeSelectors`.
	ExcludeSelectors []string `json:"excludeSelectors,omitzero"`
	// HTTP headers to send to the target site. Requests with headers skip the cache.
	Headers map[string]string `json:"headers,omitzero"`
	// Keep only elements matching these CSS selectors.
	IncludeSelectors []string `json:"includeSelectors,omitzero"`
	// Document parsing options.
	Parsers BatchSubmitParamsInputScrapeDataSharedParamsParsers `json:"parsers,omitzero"`
	// Create or reuse an authenticated browser profile. Requires a paid plan and an
	// HTTPS target. Profiles persist session state, bypass shared caches, and cannot
	// be combined with ZDR, custom headers, raw bytes, or prebuilt parsers.
	Profile BatchSubmitParamsInputScrapeDataSharedParamsProfileUnion `json:"profile,omitzero"`
	// Emulate a light or dark color scheme.
	//
	// Any of "light", "dark".
	Theme string `json:"theme,omitzero"`
	// Browser size in pixels. Omit for 1920 × 1080. When provided, missing dimensions
	// default to 1440 × 900.
	Viewport BatchSubmitParamsInputScrapeDataSharedParamsViewport `json:"viewport,omitzero"`
	// Milliseconds, or a CSS selector to wait for, after actions. Defaults to 500
	// (2000 with frames or XML).
	WaitFor BatchSubmitParamsInputScrapeDataSharedParamsWaitForUnion `json:"waitFor,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataSharedParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataSharedParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataSharedParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputScrapeDataSharedParams](
		"theme", "light", "dark",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputScrapeDataSharedParamsActionUnion struct {
	OfPerform *BatchSubmitParamsInputScrapeDataSharedParamsActionPerform `json:",omitzero,inline"`
	OfScroll  *BatchSubmitParamsInputScrapeDataSharedParamsActionScroll  `json:",omitzero,inline"`
	OfWait    *BatchSubmitParamsInputScrapeDataSharedParamsActionWait    `json:",omitzero,inline"`
	OfWaitFor *BatchSubmitParamsInputScrapeDataSharedParamsActionWaitFor `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputScrapeDataSharedParamsActionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPerform, u.OfScroll, u.OfWait, u.OfWaitFor)
}
func (u *BatchSubmitParamsInputScrapeDataSharedParamsActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[BatchSubmitParamsInputScrapeDataSharedParamsActionUnion](
		"type",
		apijson.Discriminator[BatchSubmitParamsInputScrapeDataSharedParamsActionPerform]("perform"),
		apijson.Discriminator[BatchSubmitParamsInputScrapeDataSharedParamsActionScroll]("scroll"),
		apijson.Discriminator[BatchSubmitParamsInputScrapeDataSharedParamsActionWait]("wait"),
		apijson.Discriminator[BatchSubmitParamsInputScrapeDataSharedParamsActionWaitFor]("waitFor"),
	)
}

// The properties Action, Type are required.
type BatchSubmitParamsInputScrapeDataSharedParamsActionPerform struct {
	// One browser instruction, such as clicking a button or entering text.
	Action string `json:"action" api:"required"`
	// Use `perform` for a plain-language browser instruction.
	//
	// This field can be elided, and will marshal its zero value as "perform".
	Type constant.Perform `json:"type" default:"perform"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataSharedParamsActionPerform) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataSharedParamsActionPerform
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataSharedParamsActionPerform) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type BatchSubmitParamsInputScrapeDataSharedParamsActionScroll struct {
	// Maximum number of scroll steps for this action.
	MaxScrolls param.Opt[int64] `json:"maxScrolls,omitzero"`
	// Scroll this container. Omit to scroll the page.
	Selector param.Opt[string] `json:"selector,omitzero"`
	// Distance per scroll: pixels, one `viewport`, or `max` to reach the end.
	Amount BatchSubmitParamsInputScrapeDataSharedParamsActionScrollAmountUnion `json:"amount,omitzero"`
	// Direction to scroll.
	//
	// Any of "down", "up", "left", "right".
	Direction string `json:"direction,omitzero"`
	// Use `scroll` to move through the page or a container.
	//
	// This field can be elided, and will marshal its zero value as "scroll".
	Type constant.Scroll `json:"type" default:"scroll"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataSharedParamsActionScroll) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataSharedParamsActionScroll
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataSharedParamsActionScroll) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputScrapeDataSharedParamsActionScroll](
		"direction", "down", "up", "left", "right",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputScrapeDataSharedParamsActionScrollAmountUnion struct {
	OfInt param.Opt[int64] `json:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfBatchSubmitsInputScrapeDataSharedParamsActionScrollAmountString)
	OfBatchSubmitsInputScrapeDataSharedParamsActionScrollAmountString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputScrapeDataSharedParamsActionScrollAmountUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfBatchSubmitsInputScrapeDataSharedParamsActionScrollAmountString)
}
func (u *BatchSubmitParamsInputScrapeDataSharedParamsActionScrollAmountUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type BatchSubmitParamsInputScrapeDataSharedParamsActionScrollAmountString string

const (
	BatchSubmitParamsInputScrapeDataSharedParamsActionScrollAmountStringViewport BatchSubmitParamsInputScrapeDataSharedParamsActionScrollAmountString = "viewport"
	BatchSubmitParamsInputScrapeDataSharedParamsActionScrollAmountStringMax      BatchSubmitParamsInputScrapeDataSharedParamsActionScrollAmountString = "max"
)

// The properties Milliseconds, Type are required.
type BatchSubmitParamsInputScrapeDataSharedParamsActionWait struct {
	// Time to pause in milliseconds before the next action.
	Milliseconds int64 `json:"milliseconds" api:"required"`
	// Use `wait` to pause for a fixed duration.
	//
	// This field can be elided, and will marshal its zero value as "wait".
	Type constant.Wait `json:"type" default:"wait"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataSharedParamsActionWait) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataSharedParamsActionWait
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataSharedParamsActionWait) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Selector, Type are required.
type BatchSubmitParamsInputScrapeDataSharedParamsActionWaitFor struct {
	// CSS selector to wait for before continuing.
	Selector string `json:"selector" api:"required"`
	// Use `waitFor` to wait for a matching element.
	//
	// This field can be elided, and will marshal its zero value as "waitFor".
	Type constant.WaitFor `json:"type" default:"waitFor"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataSharedParamsActionWaitFor) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataSharedParamsActionWaitFor
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataSharedParamsActionWaitFor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Document parsing options.
type BatchSubmitParamsInputScrapeDataSharedParamsParsers struct {
	// PDF page range and OCR.
	Pdf BatchSubmitParamsInputScrapeDataSharedParamsParsersPdf `json:"pdf,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataSharedParamsParsers) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataSharedParamsParsers
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataSharedParamsParsers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PDF page range and OCR.
type BatchSubmitParamsInputScrapeDataSharedParamsParsersPdf struct {
	// Last page to parse. Must be at least `startPage`.
	EndPage param.Opt[int64] `json:"endPage,omitzero"`
	// First page to parse, starting at 1.
	StartPage param.Opt[int64] `json:"startPage,omitzero"`
	// Set `auto` to read scanned pages with OCR.
	//
	// Any of "off", "auto".
	Ocr string `json:"ocr,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataSharedParamsParsersPdf) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataSharedParamsParsersPdf
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataSharedParamsParsersPdf) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputScrapeDataSharedParamsParsersPdf](
		"ocr", "off", "auto",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputScrapeDataSharedParamsProfileUnion struct {
	OfExisting *BatchSubmitParamsInputScrapeDataSharedParamsProfileExisting `json:",omitzero,inline"`
	OfNew      *BatchSubmitParamsInputScrapeDataSharedParamsProfileNew      `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputScrapeDataSharedParamsProfileUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExisting, u.OfNew)
}
func (u *BatchSubmitParamsInputScrapeDataSharedParamsProfileUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[BatchSubmitParamsInputScrapeDataSharedParamsProfileUnion](
		"type",
		apijson.Discriminator[BatchSubmitParamsInputScrapeDataSharedParamsProfileExisting]("existing"),
		apijson.Discriminator[BatchSubmitParamsInputScrapeDataSharedParamsProfileNew]("new"),
	)
}

// The properties ProfileID, Type are required.
type BatchSubmitParamsInputScrapeDataSharedParamsProfileExisting struct {
	// The ID returned in profile.id by an earlier scrape.
	ProfileID string `json:"profileId" api:"required"`
	// Reuse a saved browser profile.
	//
	// This field can be elided, and will marshal its zero value as "existing".
	Type constant.Existing `json:"type" default:"existing"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataSharedParamsProfileExisting) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataSharedParamsProfileExisting
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataSharedParamsProfileExisting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Identifier, Password, Type are required.
type BatchSubmitParamsInputScrapeDataSharedParamsProfileNew struct {
	// Username or email for the target website.
	Identifier string `json:"identifier" api:"required"`
	// Password used to sign in. It is never stored in the profile.
	Password string `json:"password" api:"required"`
	// Optional label shown when listing profiles with `GET /org/browser/profiles`.
	Description param.Opt[string] `json:"description,omitzero"`
	// Sign in and create a browser profile.
	//
	// This field can be elided, and will marshal its zero value as "new".
	Type constant.New `json:"type" default:"new"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataSharedParamsProfileNew) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataSharedParamsProfileNew
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataSharedParamsProfileNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Browser size in pixels. Omit for 1920 × 1080. When provided, missing dimensions
// default to 1440 × 900.
type BatchSubmitParamsInputScrapeDataSharedParamsViewport struct {
	// Browser viewport height in pixels.
	Height param.Opt[int64] `json:"height,omitzero"`
	// Browser viewport width in pixels.
	Width param.Opt[int64] `json:"width,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataSharedParamsViewport) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataSharedParamsViewport
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataSharedParamsViewport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputScrapeDataSharedParamsWaitForUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputScrapeDataSharedParamsWaitForUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *BatchSubmitParamsInputScrapeDataSharedParamsWaitForUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Deadline for each page, using the same behavior and defaults as POST
// /web/scrape.
//
// The property Milliseconds is required.
type BatchSubmitParamsInputScrapeDataTimeoutOpts struct {
	// Deadline in milliseconds.
	Milliseconds int64 `json:"milliseconds" api:"required"`
	// "fail" returns 408 at the deadline. "return-partial" returns available results;
	// inspect the response’s partial flag. "return-partial" requires at least 5000 ms.
	//
	// Any of "fail", "return-partial".
	Behavior string `json:"behavior,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataTimeoutOpts) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataTimeoutOpts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataTimeoutOpts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputScrapeDataTimeoutOpts](
		"behavior", "fail", "return-partial",
	)
}

// Crawl pages starting from a URL or from a domain's sitemap.
//
// The properties Data, Mode are required.
type BatchSubmitParamsInputCrawl struct {
	// Use the same formats and parameters as POST /web/scrape, replacing url with
	// source. Options apply to every page.
	Data BatchSubmitParamsInputCrawlData `json:"data,omitzero" api:"required"`
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

// Use the same formats and parameters as POST /web/scrape, replacing url with
// source. Options apply to every page.
//
// The properties Formats, Source are required.
type BatchSubmitParamsInputCrawlData struct {
	// Outputs to return. Set at least one to `true`.
	Formats BatchSubmitParamsInputCrawlDataFormats `json:"formats,omitzero" api:"required"`
	// How to find pages to crawl.
	Source BatchSubmitParamsInputCrawlDataSourceUnion `json:"source,omitzero" api:"required"`
	// Maximum age of a cached output, in milliseconds. `0` fetches fresh. Defaults to
	// 3 days (259200000 ms). Maximum: 1 year (31536000000 ms).
	MaxAgeMs param.Opt[int64] `json:"maxAgeMs,omitzero"`
	// Requires `formats.highlights: true`; required when it is set.
	HighlightsParams BatchSubmitParamsInputCrawlDataHighlightsParams `json:"highlightsParams,omitzero"`
	// Image options. Requires formats.images: true.
	ImageParams BatchSubmitParamsInputCrawlDataImageParams `json:"imageParams,omitzero"`
	// Requires `formats.json: true`; required when it is set.
	JsonParams BatchSubmitParamsInputCrawlDataJsonParams `json:"jsonParams,omitzero"`
	// Markdown options. Requires `formats.markdown`.
	MarkdownParams BatchSubmitParamsInputCrawlDataMarkdownParams `json:"markdownParams,omitzero"`
	// Requires `formats.parse: true`; required when it is set.
	ParseParams BatchSubmitParamsInputCrawlDataParseParamsUnion `json:"parseParams,omitzero"`
	// Product options. Requires formats.product: true.
	ProductParams BatchSubmitParamsInputCrawlDataProductParams `json:"productParams,omitzero"`
	// Screenshot options. Requires formats.screenshot: true.
	ScreenshotParams BatchSubmitParamsInputCrawlDataScreenshotParams `json:"screenshotParams,omitzero"`
	// Browser and content settings shared by all outputs.
	SharedParams BatchSubmitParamsInputCrawlDataSharedParams `json:"sharedParams,omitzero"`
	// Labels for tracking request usage. Not retained when zdr is enabled.
	Tags []string `json:"tags,omitzero"`
	// Deadline for each page, using the same behavior and defaults as POST
	// /web/scrape.
	TimeoutOpts BatchSubmitParamsInputCrawlDataTimeoutOpts `json:"timeoutOpts,omitzero"`
	// Batches persist their results. Only disabled is supported; enabled returns
	// ZDR_NOT_SUPPORTED.
	//
	// Any of "disabled".
	Zdr string `json:"zdr,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlData) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputCrawlData](
		"zdr", "disabled",
	)
}

// Outputs to return. Set at least one to `true`.
type BatchSubmitParamsInputCrawlDataFormats struct {
	// The original HTTP response body.
	Bytes param.Opt[bool] `json:"bytes,omitzero"`
	// Markdown excerpts relevant to `highlightsParams.query`.
	Highlights param.Opt[bool] `json:"highlights,omitzero"`
	// Rendered HTML.
	HTML param.Opt[bool] `json:"html,omitzero"`
	// Images found on the page.
	Images param.Opt[bool] `json:"images,omitzero"`
	// An object matching `jsonParams.schema`, extracted from the page.
	Json param.Opt[bool] `json:"json,omitzero"`
	// Page content as Markdown.
	Markdown param.Opt[bool] `json:"markdown,omitzero"`
	// Fields extracted with custom rules or a prebuilt parser, returned as `parsed`.
	Parse param.Opt[bool] `json:"parse,omitzero"`
	// Product details such as name, price, and availability.
	Product param.Opt[bool] `json:"product,omitzero"`
	// A screenshot of the page.
	Screenshot param.Opt[bool] `json:"screenshot,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataFormats) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataFormats
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataFormats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataSourceUnion struct {
	OfStartURL *BatchSubmitParamsInputCrawlDataSourceStartURL `json:",omitzero,inline"`
	OfSitemap  *BatchSubmitParamsInputCrawlDataSourceSitemap  `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataSourceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfStartURL, u.OfSitemap)
}
func (u *BatchSubmitParamsInputCrawlDataSourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[BatchSubmitParamsInputCrawlDataSourceUnion](
		"type",
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataSourceStartURL]("start_url"),
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataSourceSitemap]("sitemap"),
	)
}

// Discover pages by following links from one URL.
//
// The properties Type, URL are required.
type BatchSubmitParamsInputCrawlDataSourceStartURL struct {
	// Page where crawling begins. A URL without a scheme is read as https://.
	URL string `json:"url" api:"required"`
	// Limits and filters for page discovery.
	Controls BatchSubmitParamsInputCrawlDataSourceStartURLControls `json:"controls,omitzero"`
	// Start from one page.
	//
	// This field can be elided, and will marshal its zero value as "start_url".
	Type constant.StartURL `json:"type" default:"start_url"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataSourceStartURL) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataSourceStartURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataSourceStartURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Limits and filters for page discovery.
type BatchSubmitParamsInputCrawlDataSourceStartURLControls struct {
	// Follow links to subdomains.
	FollowSubdomains param.Opt[bool] `json:"followSubdomains,omitzero"`
	// Maximum link depth. Source pages are depth 0. No limit when omitted.
	MaxDepth param.Opt[int64] `json:"maxDepth,omitzero"`
	// Maximum pages to fetch. Unused reserved credits are refunded. Maximum 100000.
	MaxURLs param.Opt[int64] `json:"maxUrls,omitzero"`
	// RE2 pattern for URLs to include. The `start_url` itself is always included.
	Regex param.Opt[string] `json:"regex,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataSourceStartURLControls) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataSourceStartURLControls
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataSourceStartURLControls) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Scrape the pages listed in a domain's sitemap. Links on those pages are not
// followed.
//
// The properties Domain, Type are required.
type BatchSubmitParamsInputCrawlDataSourceSitemap struct {
	// Domain whose sitemap lists the pages to scrape. A full URL is reduced to its
	// domain.
	Domain string `json:"domain" api:"required"`
	// Limits and filters for the sitemap URLs. A sitemap batch scrapes exactly those
	// URLs and never follows links off them, so there is no crawl depth here.
	Controls BatchSubmitParamsInputCrawlDataSourceSitemapControls `json:"controls,omitzero"`
	// Scrape the URLs in the domain's sitemap.
	//
	// This field can be elided, and will marshal its zero value as "sitemap".
	Type constant.Sitemap `json:"type" default:"sitemap"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataSourceSitemap) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataSourceSitemap
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataSourceSitemap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Limits and filters for the sitemap URLs. A sitemap batch scrapes exactly those
// URLs and never follows links off them, so there is no crawl depth here.
type BatchSubmitParamsInputCrawlDataSourceSitemapControls struct {
	// Maximum pages to fetch. Unused reserved credits are refunded. Maximum 100000.
	MaxURLs param.Opt[int64] `json:"maxUrls,omitzero"`
	// RE2 pattern; only sitemap URLs matching it are scraped.
	Regex param.Opt[string] `json:"regex,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataSourceSitemapControls) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataSourceSitemapControls
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataSourceSitemapControls) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Requires `formats.highlights: true`; required when it is set.
//
// The property Query is required.
type BatchSubmitParamsInputCrawlDataHighlightsParams struct {
	// The question or topic to find passages for.
	Query string `json:"query" api:"required"`
	// Maximum combined length of returned passages.
	MaxCharacters param.Opt[int64] `json:"maxCharacters,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataHighlightsParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataHighlightsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataHighlightsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image options. Requires formats.images: true.
type BatchSubmitParamsInputCrawlDataImageParams struct {
	// Set `visual` to drop visual duplicates, keeping the largest copy.
	//
	// Any of "none", "visual".
	Dedupe string `json:"dedupe,omitzero"`
	// Extra data per image: `dimensions`, `classification`, or a hosted `file` URL.
	//
	// Any of "dimensions", "classification", "file".
	Enrich []string `json:"enrich,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataImageParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataImageParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataImageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputCrawlDataImageParams](
		"dedupe", "none", "visual",
	)
}

// Requires `formats.json: true`; required when it is set.
//
// The property Schema is required.
type BatchSubmitParamsInputCrawlDataJsonParams struct {
	// JSON Schema (not an example object) for a top-level object, up to 50 KB. Use
	// optional or nullable fields for missing facts.
	Schema map[string]any `json:"schema,omitzero" api:"required"`
	// Extra guidance, such as which facts to prefer or how to read a field.
	Instructions param.Opt[string] `json:"instructions,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataJsonParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataJsonParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataJsonParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Markdown options. Requires `formats.markdown`.
type BatchSubmitParamsInputCrawlDataMarkdownParams struct {
	// Include images in the Markdown using image syntax with URLs and alt text.
	IncludeImages param.Opt[bool] `json:"includeImages,omitzero"`
	// Keep link URLs in the Markdown. Set false to return link text without URLs.
	IncludeLinks param.Opt[bool] `json:"includeLinks,omitzero"`
	// How base64 images appear: `placeholder` (default) or `preserve`. Requires
	// `includeImages`.
	//
	// Any of "placeholder", "preserve".
	InlineImages string `json:"inlineImages,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataMarkdownParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataMarkdownParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataMarkdownParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputCrawlDataMarkdownParams](
		"inlineImages", "placeholder", "preserve",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataParseParamsUnion struct {
	OfCustomParser   *BatchSubmitParamsInputCrawlDataParseParamsCustomParser   `json:",omitzero,inline"`
	OfPrebuiltParser *BatchSubmitParamsInputCrawlDataParseParamsPrebuiltParser `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataParseParamsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCustomParser, u.OfPrebuiltParser)
}
func (u *BatchSubmitParamsInputCrawlDataParseParamsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property Rules is required.
type BatchSubmitParamsInputCrawlDataParseParamsCustomParser struct {
	// Field names mapped to CSS selectors (`h1`, `a@href`) or rule objects. Max 100
	// fields, 5 levels.
	Rules map[string]BatchSubmitParamsInputCrawlDataParseParamsCustomParserRuleUnion `json:"rules,omitzero" api:"required"`
	// Extract fields using custom CSS rules. This is the default when type is omitted.
	//
	// Any of "custom".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataParseParamsCustomParser) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataParseParamsCustomParser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataParseParamsCustomParser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputCrawlDataParseParamsCustomParser](
		"type", "custom",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataParseParamsCustomParserRuleUnion struct {
	OfString                                                      param.Opt[string]                                                 `json:",omitzero,inline"`
	OfBatchSubmitsInputCrawlDataParseParamsCustomParserRuleObject *BatchSubmitParamsInputCrawlDataParseParamsCustomParserRuleObject `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataParseParamsCustomParserRuleUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfBatchSubmitsInputCrawlDataParseParamsCustomParserRuleObject)
}
func (u *BatchSubmitParamsInputCrawlDataParseParamsCustomParserRuleUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property Selector is required.
type BatchSubmitParamsInputCrawlDataParseParamsCustomParserRuleObject struct {
	// CSS selector to match within the current page or parent rule.
	Selector string `json:"selector" api:"required"`
	// Return text, HTML, an attribute such as `@href`, or nested field rules. Defaults
	// to text.
	Output string `json:"output,omitzero"`
	// Return the first match with `item` or all matches with `list`.
	//
	// Any of "item", "list".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataParseParamsCustomParserRuleObject) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataParseParamsCustomParserRuleObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataParseParamsCustomParserRuleObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputCrawlDataParseParamsCustomParserRuleObject](
		"type", "item", "list",
	)
}

// The properties Parser, Type are required.
type BatchSubmitParamsInputCrawlDataParseParamsPrebuiltParser struct {
	// google-maps parses Google Maps business data into data.results. Accepts Maps
	// place, search, and internal data URLs. Uses a US exit and English labels
	// regardless of hl or Accept-Language. google-maps-reviews returns the reviews of
	// a place in data.reviews, in the same shape as the Serper reviews API (rating,
	// date, isoDate, snippet, translatedSnippet, likes, user, media, response, link,
	// id), plus language, guided ratings in details, and user.isLocalGuide; data also
	// carries the place featureId, cid, and sort. sort and limit pick the order and
	// count, and oldest is exact. It accepts Maps place, search, and ?cid= URLs, plus
	// search.google.com/local/reviews URLs, so the google_maps_url and
	// google_reviews_url fields of google-maps results work as input. An unknown place
	// returns no reviews. google-play parses a Play Store app into data: details,
	// developer contact, ratings, price, permissions, and data safety.
	// google-play-reviews returns the 200 newest reviews of the app, with developer
	// replies, in data.reviews. Both Play parsers accept app details URLs
	// (play.google.com/store/apps/details?id=…); hl and gl set the language and
	// storefront, defaulting to en and us. app-store parses an App Store app into
	// data: details, developer, ratings, price, version, and screenshots.
	// app-store-reviews returns up to the 500 newest reviews in data.reviews. Both App
	// Store parsers accept app URLs (apps.apple.com/us/app/…/id…); the country in the
	// path sets the storefront, defaulting to us. g2-reviews returns the 25 most
	// recently published reviews of a G2 product in data.reviews (rating, title,
	// likes, dislikes, problems_solved, reviewer role, company size, and industry),
	// plus data.product. It accepts G2 product URLs (g2.com/products/…), follows
	// renamed products, and keeps star-rating filters from the URL
	// (filters[nps_score][]=1 to 5); other filters, sort orders, and pages are
	// ignored. emails accepts any page and returns its email addresses as data.emails
	// ([{ email }]), read from mailto links, Cloudflare-protected addresses, and
	// visible text. phone-numbers accepts any page and returns its valid phone numbers
	// in E.164 as data.phone_numbers ([{ phone }]), read from tel links and visible
	// text; numbers without a country code are read for the country of the site TLD,
	// defaulting to US. links accepts any page and returns its unique http(s) links as
	// data.links ([{ url, text }]), with absolute URLs; text falls back to aria-label
	// or image alt and is null when a link has none. The emails parser reads the
	// rendered page and respects content filters; the other prebuilt parsers read
	// their data independently of DOM filters.
	//
	// Any of "google-maps", "google-maps-reviews", "google-play",
	// "google-play-reviews", "app-store", "app-store-reviews", "g2-reviews", "emails",
	// "phone-numbers", "links".
	Parser string `json:"parser,omitzero" api:"required"`
	// google-maps-reviews only. Maximum number of reviews to return. Defaults to 100;
	// maximum 1000.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// google-maps-reviews only. Order of data.reviews. newest, relevant, highest, and
	// lowest follow Google; oldest returns the earliest reviews first, which Google
	// itself does not offer.
	//
	// Any of "newest", "oldest", "relevant", "highest", "lowest".
	Sort string `json:"sort,omitzero"`
	// Use a prebuilt parser instead of custom rules.
	//
	// This field can be elided, and will marshal its zero value as "prebuilt".
	Type constant.Prebuilt `json:"type" default:"prebuilt"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataParseParamsPrebuiltParser) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataParseParamsPrebuiltParser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataParseParamsPrebuiltParser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputCrawlDataParseParamsPrebuiltParser](
		"parser", "google-maps", "google-maps-reviews", "google-play", "google-play-reviews", "app-store", "app-store-reviews", "g2-reviews", "emails", "phone-numbers", "links",
	)
	apijson.RegisterFieldValidator[BatchSubmitParamsInputCrawlDataParseParamsPrebuiltParser](
		"sort", "newest", "oldest", "relevant", "highest", "lowest",
	)
}

// Product options. Requires formats.product: true.
type BatchSubmitParamsInputCrawlDataProductParams struct {
	// Drop visually duplicate product images, keeping the largest copy.
	DedupeImages param.Opt[bool] `json:"dedupeImages,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataProductParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataProductParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataProductParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot options. Requires formats.screenshot: true.
type BatchSubmitParamsInputCrawlDataScreenshotParams struct {
	// What to capture: `viewport`, `fullPage`, one element, or a rectangle. Max 40
	// megapixels.
	Area BatchSubmitParamsInputCrawlDataScreenshotParamsAreaUnion `json:"area,omitzero"`
	// Image format for the screenshot.
	//
	// Any of "png", "jpeg", "webp".
	Format string `json:"format,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataScreenshotParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataScreenshotParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataScreenshotParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputCrawlDataScreenshotParams](
		"format", "png", "jpeg", "webp",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataScreenshotParamsAreaUnion struct {
	// Check if union is this variant with !param.IsOmitted(union.OfPage)
	OfPage      param.Opt[string]                                             `json:",omitzero,inline"`
	OfElement   *BatchSubmitParamsInputCrawlDataScreenshotParamsAreaElement   `json:",omitzero,inline"`
	OfRectangle *BatchSubmitParamsInputCrawlDataScreenshotParamsAreaRectangle `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataScreenshotParamsAreaUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPage, u.OfElement, u.OfRectangle)
}
func (u *BatchSubmitParamsInputCrawlDataScreenshotParamsAreaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type BatchSubmitParamsInputCrawlDataScreenshotParamsAreaPage string

const (
	BatchSubmitParamsInputCrawlDataScreenshotParamsAreaPageViewport BatchSubmitParamsInputCrawlDataScreenshotParamsAreaPage = "viewport"
	BatchSubmitParamsInputCrawlDataScreenshotParamsAreaPageFullPage BatchSubmitParamsInputCrawlDataScreenshotParamsAreaPage = "fullPage"
)

// The property Selector is required.
type BatchSubmitParamsInputCrawlDataScreenshotParamsAreaElement struct {
	// CSS selector matching exactly one visible element.
	Selector string `json:"selector" api:"required"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataScreenshotParamsAreaElement) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataScreenshotParamsAreaElement
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataScreenshotParamsAreaElement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pixels from the document origin.
//
// The properties Height, Width, X, Y are required.
type BatchSubmitParamsInputCrawlDataScreenshotParamsAreaRectangle struct {
	// Height of the capture in pixels.
	Height int64 `json:"height" api:"required"`
	// Width of the capture in pixels.
	Width int64 `json:"width" api:"required"`
	// Left edge of the capture, in pixels from the document origin.
	X int64 `json:"x" api:"required"`
	// Top edge of the capture, in pixels from the document origin.
	Y int64 `json:"y" api:"required"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataScreenshotParamsAreaRectangle) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataScreenshotParamsAreaRectangle
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataScreenshotParamsAreaRectangle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Browser and content settings shared by all outputs.
type BatchSubmitParamsInputCrawlDataSharedParams struct {
	// Proxy country as a two-letter code, such as `US`. Case-insensitive.
	Country param.Opt[string] `json:"country,omitzero"`
	// Accept cookie banners before actions and capture.
	DismissCookies param.Opt[bool] `json:"dismissCookies,omitzero"`
	// Close other popups before actions and capture.
	DismissPopups param.Opt[bool] `json:"dismissPopups,omitzero"`
	// Include iframe content in HTML and text outputs. Screenshots always show visible
	// frames.
	IncludeFrames param.Opt[bool] `json:"includeFrames,omitzero"`
	// Keep only the main content. Doesn't affect `screenshot`, `bytes`, or `product`.
	MainContentOnly param.Opt[bool] `json:"mainContentOnly,omitzero"`
	// Wait for CSS animations to finish before capture. Defaults to `true` when
	// `screenshot` is requested.
	SettleAnimations param.Opt[bool] `json:"settleAnimations,omitzero"`
	// Browser steps run in order before capture. Requires a paid plan. Skips the
	// cache.
	Actions []BatchSubmitParamsInputCrawlDataSharedParamsActionUnion `json:"actions,omitzero"`
	// Remove elements matching these CSS selectors. Overrides `includeSelectors`.
	ExcludeSelectors []string `json:"excludeSelectors,omitzero"`
	// HTTP headers to send to the target site. Requests with headers skip the cache.
	Headers map[string]string `json:"headers,omitzero"`
	// Keep only elements matching these CSS selectors.
	IncludeSelectors []string `json:"includeSelectors,omitzero"`
	// Document parsing options.
	Parsers BatchSubmitParamsInputCrawlDataSharedParamsParsers `json:"parsers,omitzero"`
	// Create or reuse an authenticated browser profile. Requires a paid plan and an
	// HTTPS target. Profiles persist session state, bypass shared caches, and cannot
	// be combined with ZDR, custom headers, raw bytes, or prebuilt parsers.
	Profile BatchSubmitParamsInputCrawlDataSharedParamsProfileUnion `json:"profile,omitzero"`
	// Emulate a light or dark color scheme.
	//
	// Any of "light", "dark".
	Theme string `json:"theme,omitzero"`
	// Browser size in pixels. Omit for 1920 × 1080. When provided, missing dimensions
	// default to 1440 × 900.
	Viewport BatchSubmitParamsInputCrawlDataSharedParamsViewport `json:"viewport,omitzero"`
	// Milliseconds, or a CSS selector to wait for, after actions. Defaults to 500
	// (2000 with frames or XML).
	WaitFor BatchSubmitParamsInputCrawlDataSharedParamsWaitForUnion `json:"waitFor,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataSharedParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataSharedParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataSharedParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputCrawlDataSharedParams](
		"theme", "light", "dark",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataSharedParamsActionUnion struct {
	OfPerform *BatchSubmitParamsInputCrawlDataSharedParamsActionPerform `json:",omitzero,inline"`
	OfScroll  *BatchSubmitParamsInputCrawlDataSharedParamsActionScroll  `json:",omitzero,inline"`
	OfWait    *BatchSubmitParamsInputCrawlDataSharedParamsActionWait    `json:",omitzero,inline"`
	OfWaitFor *BatchSubmitParamsInputCrawlDataSharedParamsActionWaitFor `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataSharedParamsActionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPerform, u.OfScroll, u.OfWait, u.OfWaitFor)
}
func (u *BatchSubmitParamsInputCrawlDataSharedParamsActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[BatchSubmitParamsInputCrawlDataSharedParamsActionUnion](
		"type",
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataSharedParamsActionPerform]("perform"),
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataSharedParamsActionScroll]("scroll"),
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataSharedParamsActionWait]("wait"),
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataSharedParamsActionWaitFor]("waitFor"),
	)
}

// The properties Action, Type are required.
type BatchSubmitParamsInputCrawlDataSharedParamsActionPerform struct {
	// One browser instruction, such as clicking a button or entering text.
	Action string `json:"action" api:"required"`
	// Use `perform` for a plain-language browser instruction.
	//
	// This field can be elided, and will marshal its zero value as "perform".
	Type constant.Perform `json:"type" default:"perform"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataSharedParamsActionPerform) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataSharedParamsActionPerform
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataSharedParamsActionPerform) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type BatchSubmitParamsInputCrawlDataSharedParamsActionScroll struct {
	// Maximum number of scroll steps for this action.
	MaxScrolls param.Opt[int64] `json:"maxScrolls,omitzero"`
	// Scroll this container. Omit to scroll the page.
	Selector param.Opt[string] `json:"selector,omitzero"`
	// Distance per scroll: pixels, one `viewport`, or `max` to reach the end.
	Amount BatchSubmitParamsInputCrawlDataSharedParamsActionScrollAmountUnion `json:"amount,omitzero"`
	// Direction to scroll.
	//
	// Any of "down", "up", "left", "right".
	Direction string `json:"direction,omitzero"`
	// Use `scroll` to move through the page or a container.
	//
	// This field can be elided, and will marshal its zero value as "scroll".
	Type constant.Scroll `json:"type" default:"scroll"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataSharedParamsActionScroll) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataSharedParamsActionScroll
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataSharedParamsActionScroll) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputCrawlDataSharedParamsActionScroll](
		"direction", "down", "up", "left", "right",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataSharedParamsActionScrollAmountUnion struct {
	OfInt param.Opt[int64] `json:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfBatchSubmitsInputCrawlDataSharedParamsActionScrollAmountString)
	OfBatchSubmitsInputCrawlDataSharedParamsActionScrollAmountString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataSharedParamsActionScrollAmountUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfBatchSubmitsInputCrawlDataSharedParamsActionScrollAmountString)
}
func (u *BatchSubmitParamsInputCrawlDataSharedParamsActionScrollAmountUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type BatchSubmitParamsInputCrawlDataSharedParamsActionScrollAmountString string

const (
	BatchSubmitParamsInputCrawlDataSharedParamsActionScrollAmountStringViewport BatchSubmitParamsInputCrawlDataSharedParamsActionScrollAmountString = "viewport"
	BatchSubmitParamsInputCrawlDataSharedParamsActionScrollAmountStringMax      BatchSubmitParamsInputCrawlDataSharedParamsActionScrollAmountString = "max"
)

// The properties Milliseconds, Type are required.
type BatchSubmitParamsInputCrawlDataSharedParamsActionWait struct {
	// Time to pause in milliseconds before the next action.
	Milliseconds int64 `json:"milliseconds" api:"required"`
	// Use `wait` to pause for a fixed duration.
	//
	// This field can be elided, and will marshal its zero value as "wait".
	Type constant.Wait `json:"type" default:"wait"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataSharedParamsActionWait) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataSharedParamsActionWait
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataSharedParamsActionWait) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Selector, Type are required.
type BatchSubmitParamsInputCrawlDataSharedParamsActionWaitFor struct {
	// CSS selector to wait for before continuing.
	Selector string `json:"selector" api:"required"`
	// Use `waitFor` to wait for a matching element.
	//
	// This field can be elided, and will marshal its zero value as "waitFor".
	Type constant.WaitFor `json:"type" default:"waitFor"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataSharedParamsActionWaitFor) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataSharedParamsActionWaitFor
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataSharedParamsActionWaitFor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Document parsing options.
type BatchSubmitParamsInputCrawlDataSharedParamsParsers struct {
	// PDF page range and OCR.
	Pdf BatchSubmitParamsInputCrawlDataSharedParamsParsersPdf `json:"pdf,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataSharedParamsParsers) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataSharedParamsParsers
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataSharedParamsParsers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PDF page range and OCR.
type BatchSubmitParamsInputCrawlDataSharedParamsParsersPdf struct {
	// Last page to parse. Must be at least `startPage`.
	EndPage param.Opt[int64] `json:"endPage,omitzero"`
	// First page to parse, starting at 1.
	StartPage param.Opt[int64] `json:"startPage,omitzero"`
	// Set `auto` to read scanned pages with OCR.
	//
	// Any of "off", "auto".
	Ocr string `json:"ocr,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataSharedParamsParsersPdf) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataSharedParamsParsersPdf
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataSharedParamsParsersPdf) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputCrawlDataSharedParamsParsersPdf](
		"ocr", "off", "auto",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataSharedParamsProfileUnion struct {
	OfExisting *BatchSubmitParamsInputCrawlDataSharedParamsProfileExisting `json:",omitzero,inline"`
	OfNew      *BatchSubmitParamsInputCrawlDataSharedParamsProfileNew      `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataSharedParamsProfileUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExisting, u.OfNew)
}
func (u *BatchSubmitParamsInputCrawlDataSharedParamsProfileUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[BatchSubmitParamsInputCrawlDataSharedParamsProfileUnion](
		"type",
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataSharedParamsProfileExisting]("existing"),
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataSharedParamsProfileNew]("new"),
	)
}

// The properties ProfileID, Type are required.
type BatchSubmitParamsInputCrawlDataSharedParamsProfileExisting struct {
	// The ID returned in profile.id by an earlier scrape.
	ProfileID string `json:"profileId" api:"required"`
	// Reuse a saved browser profile.
	//
	// This field can be elided, and will marshal its zero value as "existing".
	Type constant.Existing `json:"type" default:"existing"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataSharedParamsProfileExisting) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataSharedParamsProfileExisting
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataSharedParamsProfileExisting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Identifier, Password, Type are required.
type BatchSubmitParamsInputCrawlDataSharedParamsProfileNew struct {
	// Username or email for the target website.
	Identifier string `json:"identifier" api:"required"`
	// Password used to sign in. It is never stored in the profile.
	Password string `json:"password" api:"required"`
	// Optional label shown when listing profiles with `GET /org/browser/profiles`.
	Description param.Opt[string] `json:"description,omitzero"`
	// Sign in and create a browser profile.
	//
	// This field can be elided, and will marshal its zero value as "new".
	Type constant.New `json:"type" default:"new"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataSharedParamsProfileNew) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataSharedParamsProfileNew
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataSharedParamsProfileNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Browser size in pixels. Omit for 1920 × 1080. When provided, missing dimensions
// default to 1440 × 900.
type BatchSubmitParamsInputCrawlDataSharedParamsViewport struct {
	// Browser viewport height in pixels.
	Height param.Opt[int64] `json:"height,omitzero"`
	// Browser viewport width in pixels.
	Width param.Opt[int64] `json:"width,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataSharedParamsViewport) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataSharedParamsViewport
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataSharedParamsViewport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataSharedParamsWaitForUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataSharedParamsWaitForUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *BatchSubmitParamsInputCrawlDataSharedParamsWaitForUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Deadline for each page, using the same behavior and defaults as POST
// /web/scrape.
//
// The property Milliseconds is required.
type BatchSubmitParamsInputCrawlDataTimeoutOpts struct {
	// Deadline in milliseconds.
	Milliseconds int64 `json:"milliseconds" api:"required"`
	// "fail" returns 408 at the deadline. "return-partial" returns available results;
	// inspect the response’s partial flag. "return-partial" requires at least 5000 ms.
	//
	// Any of "fail", "return-partial".
	Behavior string `json:"behavior,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataTimeoutOpts) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataTimeoutOpts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataTimeoutOpts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputCrawlDataTimeoutOpts](
		"behavior", "fail", "return-partial",
	)
}

// Where to send the batch's final-status event. Omit `retry` for one attempt; `{}`
// uses the default retry schedule.
//
// The property URL is required.
type BatchSubmitParamsWebhook struct {
	// Public HTTP(S) URL that receives batch completion, failure, or cancellation
	// events.
	URL string `json:"url" api:"required" format:"uri"`
	// Webhook retry settings. Use {} for the default schedule.
	Retry RetryConfigParam `json:"retry,omitzero"`
	paramObj
}

func (r BatchSubmitParamsWebhook) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsWebhook
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
