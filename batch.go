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
// rejected-URL list and webhook signing secret from submission, so nothing is lost
// if the submit response was dropped.
func (r *BatchService) Get(ctx context.Context, batchID string, query BatchGetParams, opts ...option.RequestOption) (res *BatchGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if batchID == "" {
		err = errors.New("missing required batch_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("batch/%s", url.PathEscape(batchID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
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
func (r *BatchService) Cancel(ctx context.Context, batchID string, body BatchCancelParams, opts ...option.RequestOption) (res *BatchCancelResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if batchID == "" {
		err = errors.New("missing required batch_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("batch/%s/cancel", url.PathEscape(batchID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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

type BatchGetResponse struct {
	// Batch ID used to retrieve or cancel the job.
	ID string `json:"id" api:"required"`
	// Reserved and used credits.
	Credits BatchGetResponseCredits `json:"credits" api:"required"`
	// Batch-level error. Null unless `status` is `failed`.
	Error BatchGetResponseError `json:"error" api:"required"`
	// Page failures grouped by error code.
	Errors []BatchGetResponseError `json:"errors" api:"required"`
	// Submission counts.
	Input BatchGetResponseInput `json:"input" api:"required"`
	// Rejected URLs, up to 100. These are not charged.
	InvalidURLs []BatchGetResponseInvalidURL `json:"invalid_urls" api:"required"`
	// How pages are selected.
	//
	// Any of "scrape", "crawl".
	Mode BatchGetResponseMode `json:"mode" api:"required"`
	// Current processing counts. Use `status` to check completion.
	Progress BatchGetResponseProgress `json:"progress" api:"required"`
	// Download links available when the batch finishes. GET /batch/{batch_id}/results
	// serves the same records as paginated JSON.
	Results BatchGetResponseResults `json:"results" api:"required"`
	// Current state. `completed`, `cancelled`, and `failed` are final.
	//
	// Any of "queued", "running", "cancelling", "completed", "cancelled", "failed".
	Status BatchGetResponseStatus `json:"status" api:"required"`
	Timing BatchGetResponseTiming `json:"timing" api:"required"`
	// Output format.
	//
	// Any of "markdown", "html".
	Type BatchGetResponseType `json:"type" api:"required"`
	// API key usage for this request.
	KeyMetadata BatchGetResponseKeyMetadata `json:"key_metadata"`
	// Webhook signing secret. Also returned by GET /batch/{batch_id}.
	WebhookSecret string `json:"webhook_secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Credits       respjson.Field
		Error         respjson.Field
		Errors        respjson.Field
		Input         respjson.Field
		InvalidURLs   respjson.Field
		Mode          respjson.Field
		Progress      respjson.Field
		Results       respjson.Field
		Status        respjson.Field
		Timing        respjson.Field
		Type          respjson.Field
		KeyMetadata   respjson.Field
		WebhookSecret respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reserved and used credits.
type BatchGetResponseCredits struct {
	// Credits used by successful pages.
	Charged int64 `json:"charged" api:"required"`
	// Credits reserved when the batch was accepted.
	Estimated int64 `json:"estimated" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Charged     respjson.Field
		Estimated   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResponseCredits) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResponseCredits) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Batch-level error. Null unless `status` is `failed`.
type BatchGetResponseError struct {
	// Batch error code.
	Code string `json:"code" api:"required"`
	// Batch error message.
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
func (r BatchGetResponseError) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Submission counts.
type BatchGetResponseInput struct {
	// Pages accepted, or the crawl page limit. Credits are reserved for this count.
	Accepted int64 `json:"accepted" api:"required"`
	// Duplicate URL and `itemId` pairs skipped. Always 0 for crawls.
	Duplicates int64 `json:"duplicates" api:"required"`
	// Pages rejected during validation.
	Invalid int64 `json:"invalid" api:"required"`
	// Pages submitted before validation. For a crawl, the page limit.
	Submitted int64 `json:"submitted" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accepted    respjson.Field
		Duplicates  respjson.Field
		Invalid     respjson.Field
		Submitted   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResponseInput) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResponseInput) UnmarshalJSON(data []byte) error {
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

// How pages are selected.
type BatchGetResponseMode string

const (
	BatchGetResponseModeScrape BatchGetResponseMode = "scrape"
	BatchGetResponseModeCrawl  BatchGetResponseMode = "crawl"
)

// Current processing counts. Use `status` to check completion.
type BatchGetResponseProgress struct {
	// Pages that could not be scraped.
	Failed int64 `json:"failed" api:"required"`
	// Accepted pages not yet attempted. Always 0 once the batch completes; a crawl can
	// finish under its page limit when the site has no more reachable pages.
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

// Download links available when the batch finishes. GET /batch/{batch_id}/results
// serves the same records as paginated JSON.
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

// Output format.
type BatchGetResponseType string

const (
	BatchGetResponseTypeMarkdown BatchGetResponseType = "markdown"
	BatchGetResponseTypeHTML     BatchGetResponseType = "html"
)

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
	// Reserved and used credits.
	Credits BatchListResponseDataCredits `json:"credits" api:"required"`
	// Batch-level error. Null unless `status` is `failed`.
	Error BatchListResponseDataError `json:"error" api:"required"`
	// Page failures grouped by error code.
	Errors []BatchListResponseDataError `json:"errors" api:"required"`
	// Submission counts.
	Input BatchListResponseDataInput `json:"input" api:"required"`
	// How pages are selected.
	//
	// Any of "scrape", "crawl".
	Mode string `json:"mode" api:"required"`
	// Current processing counts. Use `status` to check completion.
	Progress BatchListResponseDataProgress `json:"progress" api:"required"`
	// Download links available when the batch finishes. GET /batch/{batch_id}/results
	// serves the same records as paginated JSON.
	Results BatchListResponseDataResults `json:"results" api:"required"`
	// Current state. `completed`, `cancelled`, and `failed` are final.
	//
	// Any of "queued", "running", "cancelling", "completed", "cancelled", "failed".
	Status string                      `json:"status" api:"required"`
	Timing BatchListResponseDataTiming `json:"timing" api:"required"`
	// Output format.
	//
	// Any of "markdown", "html".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Credits     respjson.Field
		Error       respjson.Field
		Errors      respjson.Field
		Input       respjson.Field
		Mode        respjson.Field
		Progress    respjson.Field
		Results     respjson.Field
		Status      respjson.Field
		Timing      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchListResponseData) RawJSON() string { return r.JSON.raw }
func (r *BatchListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reserved and used credits.
type BatchListResponseDataCredits struct {
	// Credits used by successful pages.
	Charged int64 `json:"charged" api:"required"`
	// Credits reserved when the batch was accepted.
	Estimated int64 `json:"estimated" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Charged     respjson.Field
		Estimated   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchListResponseDataCredits) RawJSON() string { return r.JSON.raw }
func (r *BatchListResponseDataCredits) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Batch-level error. Null unless `status` is `failed`.
type BatchListResponseDataError struct {
	// Batch error code.
	Code string `json:"code" api:"required"`
	// Batch error message.
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
func (r BatchListResponseDataError) RawJSON() string { return r.JSON.raw }
func (r *BatchListResponseDataError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Submission counts.
type BatchListResponseDataInput struct {
	// Pages accepted, or the crawl page limit. Credits are reserved for this count.
	Accepted int64 `json:"accepted" api:"required"`
	// Duplicate URL and `itemId` pairs skipped. Always 0 for crawls.
	Duplicates int64 `json:"duplicates" api:"required"`
	// Pages rejected during validation.
	Invalid int64 `json:"invalid" api:"required"`
	// Pages submitted before validation. For a crawl, the page limit.
	Submitted int64 `json:"submitted" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accepted    respjson.Field
		Duplicates  respjson.Field
		Invalid     respjson.Field
		Submitted   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchListResponseDataInput) RawJSON() string { return r.JSON.raw }
func (r *BatchListResponseDataInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current processing counts. Use `status` to check completion.
type BatchListResponseDataProgress struct {
	// Pages that could not be scraped.
	Failed int64 `json:"failed" api:"required"`
	// Accepted pages not yet attempted. Always 0 once the batch completes; a crawl can
	// finish under its page limit when the site has no more reachable pages.
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

// Download links available when the batch finishes. GET /batch/{batch_id}/results
// serves the same records as paginated JSON.
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
	// Reserved and used credits.
	Credits BatchCancelResponseCredits `json:"credits" api:"required"`
	// Batch-level error. Null unless `status` is `failed`.
	Error BatchCancelResponseError `json:"error" api:"required"`
	// Page failures grouped by error code.
	Errors []BatchCancelResponseError `json:"errors" api:"required"`
	// Submission counts.
	Input BatchCancelResponseInput `json:"input" api:"required"`
	// How pages are selected.
	//
	// Any of "scrape", "crawl".
	Mode BatchCancelResponseMode `json:"mode" api:"required"`
	// Current processing counts. Use `status` to check completion.
	Progress BatchCancelResponseProgress `json:"progress" api:"required"`
	// Download links available when the batch finishes. GET /batch/{batch_id}/results
	// serves the same records as paginated JSON.
	Results BatchCancelResponseResults `json:"results" api:"required"`
	// Current state. `completed`, `cancelled`, and `failed` are final.
	//
	// Any of "queued", "running", "cancelling", "completed", "cancelled", "failed".
	Status BatchCancelResponseStatus `json:"status" api:"required"`
	Timing BatchCancelResponseTiming `json:"timing" api:"required"`
	// Output format.
	//
	// Any of "markdown", "html".
	Type BatchCancelResponseType `json:"type" api:"required"`
	// API key usage for this request.
	KeyMetadata BatchCancelResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Credits     respjson.Field
		Error       respjson.Field
		Errors      respjson.Field
		Input       respjson.Field
		Mode        respjson.Field
		Progress    respjson.Field
		Results     respjson.Field
		Status      respjson.Field
		Timing      respjson.Field
		Type        respjson.Field
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

// Reserved and used credits.
type BatchCancelResponseCredits struct {
	// Credits used by successful pages.
	Charged int64 `json:"charged" api:"required"`
	// Credits reserved when the batch was accepted.
	Estimated int64 `json:"estimated" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Charged     respjson.Field
		Estimated   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchCancelResponseCredits) RawJSON() string { return r.JSON.raw }
func (r *BatchCancelResponseCredits) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Batch-level error. Null unless `status` is `failed`.
type BatchCancelResponseError struct {
	// Batch error code.
	Code string `json:"code" api:"required"`
	// Batch error message.
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
func (r BatchCancelResponseError) RawJSON() string { return r.JSON.raw }
func (r *BatchCancelResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Submission counts.
type BatchCancelResponseInput struct {
	// Pages accepted, or the crawl page limit. Credits are reserved for this count.
	Accepted int64 `json:"accepted" api:"required"`
	// Duplicate URL and `itemId` pairs skipped. Always 0 for crawls.
	Duplicates int64 `json:"duplicates" api:"required"`
	// Pages rejected during validation.
	Invalid int64 `json:"invalid" api:"required"`
	// Pages submitted before validation. For a crawl, the page limit.
	Submitted int64 `json:"submitted" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accepted    respjson.Field
		Duplicates  respjson.Field
		Invalid     respjson.Field
		Submitted   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchCancelResponseInput) RawJSON() string { return r.JSON.raw }
func (r *BatchCancelResponseInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How pages are selected.
type BatchCancelResponseMode string

const (
	BatchCancelResponseModeScrape BatchCancelResponseMode = "scrape"
	BatchCancelResponseModeCrawl  BatchCancelResponseMode = "crawl"
)

// Current processing counts. Use `status` to check completion.
type BatchCancelResponseProgress struct {
	// Pages that could not be scraped.
	Failed int64 `json:"failed" api:"required"`
	// Accepted pages not yet attempted. Always 0 once the batch completes; a crawl can
	// finish under its page limit when the site has no more reachable pages.
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

// Download links available when the batch finishes. GET /batch/{batch_id}/results
// serves the same records as paginated JSON.
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

// Output format.
type BatchCancelResponseType string

const (
	BatchCancelResponseTypeMarkdown BatchCancelResponseType = "markdown"
	BatchCancelResponseTypeHTML     BatchCancelResponseType = "html"
)

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

type BatchGetParams struct {
	// Optional comma-separated caller-defined tags for tracking this request. Tags are
	// recorded on the request's usage log and can be used to filter usage on the
	// dashboard usage page. Up to 20 tags, each 1-50 characters.
	Tags []string `query:"tags,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BatchGetParams]'s query parameters as `url.Values`.
func (r BatchGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BatchListParams struct {
	// Cursor from the previous page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Batches per page. Defaults to 25.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by status.
	//
	// Any of "queued", "running", "cancelling", "completed", "cancelled", "failed".
	Status BatchListParamsStatus `query:"status,omitzero" json:"-"`
	// Optional comma-separated caller-defined tags for tracking this request. Tags are
	// recorded on the request's usage log and can be used to filter usage on the
	// dashboard usage page. Up to 20 tags, each 1-50 characters.
	Tags []string `query:"tags,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BatchListParams]'s query parameters as `url.Values`.
func (r BatchListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

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

type BatchCancelParams struct {
	// Optional comma-separated caller-defined tags for tracking this request. Tags are
	// recorded on the request's usage log and can be used to filter usage on the
	// dashboard usage page. Up to 20 tags, each 1-50 characters.
	Tags []string `query:"tags,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BatchCancelParams]'s query parameters as `url.Values`.
func (r BatchCancelParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BatchGetResultsParams struct {
	// next_cursor from the previous page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Records per page. Defaults to 25. A page can close early so its payload stays
	// under ~8 MB; rely on next_cursor rather than counting records.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Optional comma-separated caller-defined tags for tracking this request. Tags are
	// recorded on the request's usage log and can be used to filter usage on the
	// dashboard usage page. Up to 20 tags, each 1-50 characters.
	Tags []string `query:"tags,omitzero" json:"-"`
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
