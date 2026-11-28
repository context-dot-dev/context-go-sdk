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

// Scrape up to 100,000 URLs or crawl a site asynchronously.
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

// Retrieve a batch's status, progress, and credit usage, plus links to download
// its results once it finishes. Result files are deleted 180 days after the batch
// finishes.
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

// List your batches, newest first. Filter by status or tags, or search by batch
// ID, crawl source, or tag.
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

// Stop a batch from starting new pages. Pages already in progress finish, then the
// batch becomes `cancelled` and unused credits are refunded.
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

// Page through a finished batch's results as JSON. Each record's `result` has the
// same outputs and per-format errors as `POST /web/scrape`. Results are kept for
// 180 days after the batch finishes.
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

// Scrape up to 100,000 URLs, or crawl a site, asynchronously with every Scrape
// Anything format and option from `POST /web/scrape`. Poll `GET /batch/{batch_id}`
// or pass `webhook` to be notified when the batch finishes. URLs and credits are
// checked after the batch is accepted, so a batch that fails those checks ends as
// `failed` with a `failure` code instead of returning an error.
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
	// Number of pages that failed with this code.
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

// Why the batch as a whole failed. Individual page failures are counted in
// `page_errors`.
type Failure struct {
	// Machine-readable failure code, such as `stalled` or `USAGE_EXCEEDED`.
	Code string `json:"code" api:"required"`
	// Human-readable explanation of the failure.
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
	// Matchers for the pages the crawl skipped. `null` when it skipped none.
	Exclude []CrawlControlsExcludeUnion `json:"exclude" api:"required"`
	// Regex that excluded matching URLs from the crawl. `null` when the crawl set
	// none.
	ExcludeURLPattern string `json:"exclude_url_pattern" api:"required"`
	// Matchers for the links the crawl followed. `null` when it followed every link.
	Follow []CrawlControlsFollowUnion `json:"follow" api:"required"`
	// Whether links to subdomains were followed. Always `false` for a sitemap crawl.
	FollowSubdomains bool `json:"follow_subdomains" api:"required"`
	// Languages whose pages were followed. `null` when all languages were.
	Languages CrawlControlsLanguagesUnion `json:"languages" api:"required"`
	// Maximum link depth. Always `0` for a sitemap crawl; `null` when a `start_url`
	// crawl set no limit.
	MaxDepth int64 `json:"max_depth" api:"required"`
	// The `maxUrls` submitted with the crawl. A sitemap crawl scrapes only the URLs
	// its sitemap lists, up to this many, so `input.reserved` can be lower.
	MaxPages int64 `json:"max_pages" api:"required"`
	// Matchers for the pages the crawl returned. `null` when it returned every fetched
	// page.
	Scrape []CrawlControlsScrapeUnion `json:"scrape" api:"required"`
	// Where the crawl started.
	Source CrawlControlsSourceUnion `json:"source" api:"required"`
	// Regex that URLs had to match to be crawled. `null` when the crawl set none.
	URLPattern string `json:"url_pattern" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Exclude           respjson.Field
		ExcludeURLPattern respjson.Field
		Follow            respjson.Field
		FollowSubdomains  respjson.Field
		Languages         respjson.Field
		MaxDepth          respjson.Field
		MaxPages          respjson.Field
		Scrape            respjson.Field
		Source            respjson.Field
		URLPattern        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CrawlControls) RawJSON() string { return r.JSON.raw }
func (r *CrawlControls) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CrawlControlsExcludeUnion contains all possible properties and values from
// [CrawlControlsExcludeRegex], [CrawlControlsExcludeProduct].
//
// Use the [CrawlControlsExcludeUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CrawlControlsExcludeUnion struct {
	// Any of "regex", "product".
	Kind string `json:"kind"`
	// This field is from variant [CrawlControlsExcludeRegex].
	Pattern string `json:"pattern"`
	JSON    struct {
		Kind    respjson.Field
		Pattern respjson.Field
		raw     string
	} `json:"-"`
}

// anyCrawlControlsExclude is implemented by each variant of
// [CrawlControlsExcludeUnion] to add type safety for the return type of
// [CrawlControlsExcludeUnion.AsAny]
type anyCrawlControlsExclude interface {
	implCrawlControlsExcludeUnion()
}

func (CrawlControlsExcludeRegex) implCrawlControlsExcludeUnion()   {}
func (CrawlControlsExcludeProduct) implCrawlControlsExcludeUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := CrawlControlsExcludeUnion.AsAny().(type) {
//	case contextdev.CrawlControlsExcludeRegex:
//	case contextdev.CrawlControlsExcludeProduct:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u CrawlControlsExcludeUnion) AsAny() anyCrawlControlsExclude {
	switch u.Kind {
	case "regex":
		return u.AsRegex()
	case "product":
		return u.AsProduct()
	}
	return nil
}

func (u CrawlControlsExcludeUnion) AsRegex() (v CrawlControlsExcludeRegex) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CrawlControlsExcludeUnion) AsProduct() (v CrawlControlsExcludeProduct) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CrawlControlsExcludeUnion) RawJSON() string { return u.JSON.raw }

func (r *CrawlControlsExcludeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pages whose URL matches a regex pattern.
type CrawlControlsExcludeRegex struct {
	// Match pages by URL pattern.
	Kind constant.Regex `json:"kind" default:"regex"`
	// Regular expression matched against the page URL.
	Pattern string `json:"pattern" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		Pattern     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CrawlControlsExcludeRegex) RawJSON() string { return r.JSON.raw }
func (r *CrawlControlsExcludeRegex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Product detail pages.
type CrawlControlsExcludeProduct struct {
	// Match product detail pages.
	Kind constant.Product `json:"kind" default:"product"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CrawlControlsExcludeProduct) RawJSON() string { return r.JSON.raw }
func (r *CrawlControlsExcludeProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CrawlControlsFollowUnion contains all possible properties and values from
// [CrawlControlsFollowRegex], [CrawlControlsFollowPagination],
// [CrawlControlsFollowProduct].
//
// Use the [CrawlControlsFollowUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CrawlControlsFollowUnion struct {
	// Any of "regex", "pagination", "product".
	Kind string `json:"kind"`
	// This field is from variant [CrawlControlsFollowRegex].
	Pattern string `json:"pattern"`
	JSON    struct {
		Kind    respjson.Field
		Pattern respjson.Field
		raw     string
	} `json:"-"`
}

// anyCrawlControlsFollow is implemented by each variant of
// [CrawlControlsFollowUnion] to add type safety for the return type of
// [CrawlControlsFollowUnion.AsAny]
type anyCrawlControlsFollow interface {
	implCrawlControlsFollowUnion()
}

func (CrawlControlsFollowRegex) implCrawlControlsFollowUnion()      {}
func (CrawlControlsFollowPagination) implCrawlControlsFollowUnion() {}
func (CrawlControlsFollowProduct) implCrawlControlsFollowUnion()    {}

// Use the following switch statement to find the correct variant
//
//	switch variant := CrawlControlsFollowUnion.AsAny().(type) {
//	case contextdev.CrawlControlsFollowRegex:
//	case contextdev.CrawlControlsFollowPagination:
//	case contextdev.CrawlControlsFollowProduct:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u CrawlControlsFollowUnion) AsAny() anyCrawlControlsFollow {
	switch u.Kind {
	case "regex":
		return u.AsRegex()
	case "pagination":
		return u.AsPagination()
	case "product":
		return u.AsProduct()
	}
	return nil
}

func (u CrawlControlsFollowUnion) AsRegex() (v CrawlControlsFollowRegex) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CrawlControlsFollowUnion) AsPagination() (v CrawlControlsFollowPagination) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CrawlControlsFollowUnion) AsProduct() (v CrawlControlsFollowProduct) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CrawlControlsFollowUnion) RawJSON() string { return u.JSON.raw }

func (r *CrawlControlsFollowUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pages whose URL matches a regex pattern.
type CrawlControlsFollowRegex struct {
	// Match pages by URL pattern.
	Kind constant.Regex `json:"kind" default:"regex"`
	// Regular expression matched against the page URL.
	Pattern string `json:"pattern" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		Pattern     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CrawlControlsFollowRegex) RawJSON() string { return r.JSON.raw }
func (r *CrawlControlsFollowRegex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Further pages of a paginated list, such as page 2 of a catalog.
type CrawlControlsFollowPagination struct {
	// Match the further pages of a paginated list.
	Kind constant.Pagination `json:"kind" default:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CrawlControlsFollowPagination) RawJSON() string { return r.JSON.raw }
func (r *CrawlControlsFollowPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Product detail pages.
type CrawlControlsFollowProduct struct {
	// Match product detail pages.
	Kind constant.Product `json:"kind" default:"product"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CrawlControlsFollowProduct) RawJSON() string { return r.JSON.raw }
func (r *CrawlControlsFollowProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CrawlControlsLanguagesUnion contains all possible properties and values from
// [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfCrawlControlsLanguagesString
// OfCrawlControlsLanguagesArrayItemArray]
type CrawlControlsLanguagesUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfCrawlControlsLanguagesString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfCrawlControlsLanguagesArrayItemArray []string `json:",inline"`
	JSON                                   struct {
		OfCrawlControlsLanguagesString         respjson.Field
		OfCrawlControlsLanguagesArrayItemArray respjson.Field
		raw                                    string
	} `json:"-"`
}

func (u CrawlControlsLanguagesUnion) AsCrawlControlsLanguagesString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CrawlControlsLanguagesUnion) AsCrawlControlsLanguagesArrayItemArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CrawlControlsLanguagesUnion) RawJSON() string { return u.JSON.raw }

func (r *CrawlControlsLanguagesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlControlsLanguagesString string

const (
	CrawlControlsLanguagesStringSameAsStartURL CrawlControlsLanguagesString = "same_as_start_url"
)

// CrawlControlsScrapeUnion contains all possible properties and values from
// [CrawlControlsScrapeRegex], [CrawlControlsScrapeProduct].
//
// Use the [CrawlControlsScrapeUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CrawlControlsScrapeUnion struct {
	// Any of "regex", "product".
	Kind string `json:"kind"`
	// This field is from variant [CrawlControlsScrapeRegex].
	Pattern string `json:"pattern"`
	JSON    struct {
		Kind    respjson.Field
		Pattern respjson.Field
		raw     string
	} `json:"-"`
}

// anyCrawlControlsScrape is implemented by each variant of
// [CrawlControlsScrapeUnion] to add type safety for the return type of
// [CrawlControlsScrapeUnion.AsAny]
type anyCrawlControlsScrape interface {
	implCrawlControlsScrapeUnion()
}

func (CrawlControlsScrapeRegex) implCrawlControlsScrapeUnion()   {}
func (CrawlControlsScrapeProduct) implCrawlControlsScrapeUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := CrawlControlsScrapeUnion.AsAny().(type) {
//	case contextdev.CrawlControlsScrapeRegex:
//	case contextdev.CrawlControlsScrapeProduct:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u CrawlControlsScrapeUnion) AsAny() anyCrawlControlsScrape {
	switch u.Kind {
	case "regex":
		return u.AsRegex()
	case "product":
		return u.AsProduct()
	}
	return nil
}

func (u CrawlControlsScrapeUnion) AsRegex() (v CrawlControlsScrapeRegex) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CrawlControlsScrapeUnion) AsProduct() (v CrawlControlsScrapeProduct) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CrawlControlsScrapeUnion) RawJSON() string { return u.JSON.raw }

func (r *CrawlControlsScrapeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pages whose URL matches a regex pattern.
type CrawlControlsScrapeRegex struct {
	// Match pages by URL pattern.
	Kind constant.Regex `json:"kind" default:"regex"`
	// Regular expression matched against the page URL.
	Pattern string `json:"pattern" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		Pattern     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CrawlControlsScrapeRegex) RawJSON() string { return r.JSON.raw }
func (r *CrawlControlsScrapeRegex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Product detail pages.
type CrawlControlsScrapeProduct struct {
	// Match product detail pages.
	Kind constant.Product `json:"kind" default:"product"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CrawlControlsScrapeProduct) RawJSON() string { return r.JSON.raw }
func (r *CrawlControlsScrapeProduct) UnmarshalJSON(data []byte) error {
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

// Page counts from your submission.
type Intake struct {
	// URLs dropped because another entry resolved to the same page. Sitemap crawls
	// report duplicates too, since sitemaps often list a page more than once.
	Duplicates int64 `json:"duplicates" api:"required"`
	// URLs rejected as invalid. `null` for a crawl.
	Invalid int64 `json:"invalid" api:"required"`
	// Pages accepted for scraping. `progress` counts toward this total.
	Reserved int64 `json:"reserved" api:"required"`
	// `true` when `reserved` is a crawl ceiling, `false` when it is an exact URL
	// count.
	ReservedIsCeiling bool `json:"reserved_is_ceiling" api:"required"`
	// URLs in your list before validation and de-duplication. `null` for a crawl.
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
	// Unique ID of the batch.
	ID string `json:"id" api:"required"`
	// Crawl settings as submitted.
	Crawl CrawlControls `json:"crawl" api:"required"`
	// Credits reserved, refunded, and charged for the batch.
	Credits BatchGetResponseCredits `json:"credits" api:"required"`
	// Why the batch as a whole failed. Individual page failures are counted in
	// `page_errors`.
	Failure Failure `json:"failure" api:"required"`
	// Formats requested for each page, as submitted in `input.data.formats`.
	Formats BatchGetResponseFormats `json:"formats" api:"required"`
	// Page counts from your submission.
	Input Intake `json:"input" api:"required"`
	// URLs rejected from your list, up to the first 100.
	InvalidURLs []BatchGetResponseInvalidURL `json:"invalid_urls" api:"required"`
	// `scrape` for a URL list, `crawl` for a site crawl.
	//
	// Any of "scrape", "crawl".
	Mode BatchGetResponseMode `json:"mode" api:"required"`
	// Page failures grouped by error code, most common first. For the batch itself
	// failing, see `failure`.
	PageErrors []PageErrorCount `json:"page_errors" api:"required"`
	// Page counts so far. Check `status` to know when the batch is done.
	Progress BatchGetResponseProgress `json:"progress" api:"required"`
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Links to download the results. `null` until the batch finishes; the files are
	// deleted 180 days after that.
	Results BatchGetResponseResults `json:"results" api:"required"`
	// Current state. `pending` batches are being validated or are waiting for a slot
	// under your plan's concurrent batch limit, `queued` batches are ready to start,
	// and `running` batches are scraping. `completed`, `cancelled`, and `failed` are
	// final.
	//
	// Any of "pending", "queued", "running", "cancelling", "completed", "cancelled",
	// "failed".
	Status BatchGetResponseStatus `json:"status" api:"required"`
	// Tags set when the batch was submitted.
	Tags []string `json:"tags" api:"required"`
	// When the batch was created, started, and finished.
	Timing BatchGetResponseTiming `json:"timing" api:"required"`
	// Credits this request used and your remaining balance.
	KeyMetadata BatchGetResponseKeyMetadata `json:"key_metadata"`
	// Access policy IDs applied to every page.
	Policies []string `json:"policies"`
	// ID of the batch's final-status webhook delivery, for
	// `GET /webhooks/deliveries/{delivery_id}`. Omitted until a delivery exists.
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
		Policies          respjson.Field
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

// Credits reserved, refunded, and charged for the batch.
type BatchGetResponseCredits struct {
	// `reserved` minus `refunded` plus `ocr_charged`.
	Net int64 `json:"net" api:"required"`
	// Credits charged for OCR when the batch finishes.
	OcrCharged int64 `json:"ocr_charged" api:"required"`
	// Credits returned when the batch finishes, such as for failed or unattempted
	// pages.
	Refunded int64 `json:"refunded" api:"required"`
	// Credits reserved before scraping starts. `0` while `pending`.
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

// Formats requested for each page, as submitted in `input.data.formats`.
type BatchGetResponseFormats struct {
	// Return the original HTTP response body.
	Bytes bool `json:"bytes"`
	// Return Markdown excerpts relevant to `highlightsParams.query`.
	Highlights bool `json:"highlights"`
	// Return the rendered HTML.
	HTML bool `json:"html"`
	// Return the images found on the page.
	Images bool `json:"images"`
	// Return an object matching `jsonParams.schema`, extracted from the page text and
	// embedded data such as JSON-LD, meta tags, and framework state.
	Json bool `json:"json"`
	// Return the page content, or a video transcript, as Markdown.
	Markdown bool `json:"markdown"`
	// Return fields extracted with custom rules or a prebuilt parser, as `parsed`.
	Parse bool `json:"parse"`
	// Return product details such as name, price, and availability.
	Product bool `json:"product"`
	// Return a screenshot of the page.
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

// `scrape` for a URL list, `crawl` for a site crawl.
type BatchGetResponseMode string

const (
	BatchGetResponseModeScrape BatchGetResponseMode = "scrape"
	BatchGetResponseModeCrawl  BatchGetResponseMode = "crawl"
)

// Page counts so far. Check `status` to know when the batch is done.
type BatchGetResponseProgress struct {
	// Pages that could not be scraped.
	Failed int64 `json:"failed" api:"required"`
	// Accepted pages not yet attempted. A link-following crawl reports `0` once it
	// finishes, since its unused page budget is not counted.
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

// Links to download the results. `null` until the batch finishes; the files are
// deleted 180 days after that.
type BatchGetResponseResults struct {
	// When these links expire, 24 hours after this response. Retrieve the batch again
	// for fresh links.
	ExpiresAt string `json:"expires_at" api:"required"`
	// Result files, in no particular order.
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
	// Number of result records in this file.
	Items int64 `json:"items" api:"required"`
	// Temporary download URL for a gzipped NDJSON file of result records.
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

// Current state. `pending` batches are being validated or are waiting for a slot
// under your plan's concurrent batch limit, `queued` batches are ready to start,
// and `running` batches are scraping. `completed`, `cancelled`, and `failed` are
// final.
type BatchGetResponseStatus string

const (
	BatchGetResponseStatusPending    BatchGetResponseStatus = "pending"
	BatchGetResponseStatusQueued     BatchGetResponseStatus = "queued"
	BatchGetResponseStatusRunning    BatchGetResponseStatus = "running"
	BatchGetResponseStatusCancelling BatchGetResponseStatus = "cancelling"
	BatchGetResponseStatusCompleted  BatchGetResponseStatus = "completed"
	BatchGetResponseStatusCancelled  BatchGetResponseStatus = "cancelled"
	BatchGetResponseStatusFailed     BatchGetResponseStatus = "failed"
)

// When the batch was created, started, and finished.
type BatchGetResponseTiming struct {
	// When the batch reached a final status. `null` until then.
	CompletedAt string `json:"completed_at" api:"required"`
	// When the batch was created.
	CreatedAt string `json:"created_at" api:"required"`
	// When scraping started. `null` until the batch is `running`.
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

// Credits this request used and your remaining balance.
type BatchGetResponseKeyMetadata struct {
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
	// Whether more results are available after this page.
	HasMore bool `json:"has_more"`
	// Credits this request used and your remaining balance.
	KeyMetadata BatchListResponseKeyMetadata `json:"key_metadata"`
	// Pass as `cursor` to fetch the next page. `null` on the last page.
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

// An asynchronous scrape or crawl job.
type BatchListResponseData struct {
	// Unique ID of the batch.
	ID string `json:"id" api:"required"`
	// Crawl settings as submitted.
	Crawl CrawlControls `json:"crawl" api:"required"`
	// Credits reserved, refunded, and charged for the batch.
	Credits BatchListResponseDataCredits `json:"credits" api:"required"`
	// Why the batch as a whole failed. Individual page failures are counted in
	// `page_errors`.
	Failure Failure `json:"failure" api:"required"`
	// Formats requested for each page, as submitted in `input.data.formats`.
	Formats BatchListResponseDataFormats `json:"formats" api:"required"`
	// Page counts from your submission.
	Input Intake `json:"input" api:"required"`
	// `scrape` for a URL list, `crawl` for a site crawl.
	//
	// Any of "scrape", "crawl".
	Mode string `json:"mode" api:"required"`
	// Page failures grouped by error code, most common first. For the batch itself
	// failing, see `failure`.
	PageErrors []PageErrorCount `json:"page_errors" api:"required"`
	// Page counts so far. Check `status` to know when the batch is done.
	Progress BatchListResponseDataProgress `json:"progress" api:"required"`
	// Links to download the results. `null` until the batch finishes; the files are
	// deleted 180 days after that.
	Results BatchListResponseDataResults `json:"results" api:"required"`
	// Current state. `pending` batches are being validated or are waiting for a slot
	// under your plan's concurrent batch limit, `queued` batches are ready to start,
	// and `running` batches are scraping. `completed`, `cancelled`, and `failed` are
	// final.
	//
	// Any of "pending", "queued", "running", "cancelling", "completed", "cancelled",
	// "failed".
	Status string `json:"status" api:"required"`
	// Tags set when the batch was submitted.
	Tags []string `json:"tags" api:"required"`
	// When the batch was created, started, and finished.
	Timing BatchListResponseDataTiming `json:"timing" api:"required"`
	// Access policy IDs applied to every page.
	Policies []string `json:"policies"`
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
		Policies    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchListResponseData) RawJSON() string { return r.JSON.raw }
func (r *BatchListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credits reserved, refunded, and charged for the batch.
type BatchListResponseDataCredits struct {
	// `reserved` minus `refunded` plus `ocr_charged`.
	Net int64 `json:"net" api:"required"`
	// Credits charged for OCR when the batch finishes.
	OcrCharged int64 `json:"ocr_charged" api:"required"`
	// Credits returned when the batch finishes, such as for failed or unattempted
	// pages.
	Refunded int64 `json:"refunded" api:"required"`
	// Credits reserved before scraping starts. `0` while `pending`.
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

// Formats requested for each page, as submitted in `input.data.formats`.
type BatchListResponseDataFormats struct {
	// Return the original HTTP response body.
	Bytes bool `json:"bytes"`
	// Return Markdown excerpts relevant to `highlightsParams.query`.
	Highlights bool `json:"highlights"`
	// Return the rendered HTML.
	HTML bool `json:"html"`
	// Return the images found on the page.
	Images bool `json:"images"`
	// Return an object matching `jsonParams.schema`, extracted from the page text and
	// embedded data such as JSON-LD, meta tags, and framework state.
	Json bool `json:"json"`
	// Return the page content, or a video transcript, as Markdown.
	Markdown bool `json:"markdown"`
	// Return fields extracted with custom rules or a prebuilt parser, as `parsed`.
	Parse bool `json:"parse"`
	// Return product details such as name, price, and availability.
	Product bool `json:"product"`
	// Return a screenshot of the page.
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

// Page counts so far. Check `status` to know when the batch is done.
type BatchListResponseDataProgress struct {
	// Pages that could not be scraped.
	Failed int64 `json:"failed" api:"required"`
	// Accepted pages not yet attempted. A link-following crawl reports `0` once it
	// finishes, since its unused page budget is not counted.
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

// Links to download the results. `null` until the batch finishes; the files are
// deleted 180 days after that.
type BatchListResponseDataResults struct {
	// When these links expire, 24 hours after this response. Retrieve the batch again
	// for fresh links.
	ExpiresAt string `json:"expires_at" api:"required"`
	// Result files, in no particular order.
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
	// Number of result records in this file.
	Items int64 `json:"items" api:"required"`
	// Temporary download URL for a gzipped NDJSON file of result records.
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

// When the batch was created, started, and finished.
type BatchListResponseDataTiming struct {
	// When the batch reached a final status. `null` until then.
	CompletedAt string `json:"completed_at" api:"required"`
	// When the batch was created.
	CreatedAt string `json:"created_at" api:"required"`
	// When scraping started. `null` until the batch is `running`.
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
	// Always `true`.
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
func (r BatchDeleteResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *BatchDeleteResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchCancelResponse struct {
	// Unique ID of the batch.
	ID string `json:"id" api:"required"`
	// Crawl settings as submitted.
	Crawl CrawlControls `json:"crawl" api:"required"`
	// Credits reserved for the batch.
	Credits BatchCancelResponseCredits `json:"credits" api:"required"`
	// Formats requested for each page, as submitted in `input.data.formats`.
	Formats BatchCancelResponseFormats `json:"formats" api:"required"`
	// Page counts from your submission.
	Input Intake `json:"input" api:"required"`
	// `scrape` for a URL list, `crawl` for a site crawl.
	//
	// Any of "scrape", "crawl".
	Mode BatchCancelResponseMode `json:"mode" api:"required"`
	// Page failures so far, grouped by error code, most common first.
	PageErrors []PageErrorCount `json:"page_errors" api:"required"`
	// How far the batch got before cancellation.
	Progress BatchCancelResponseProgress `json:"progress" api:"required"`
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Always `cancelling`. Pages in progress finish, then the batch becomes
	// `cancelled`.
	//
	// Any of "cancelling".
	Status BatchCancelResponseStatus `json:"status" api:"required"`
	// Tags set when the batch was submitted.
	Tags []string `json:"tags" api:"required"`
	// When the batch was created and started.
	Timing BatchCancelResponseTiming `json:"timing" api:"required"`
	// Credits this request used and your remaining balance.
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

// Credits reserved for the batch.
type BatchCancelResponseCredits struct {
	// Credits reserved for the batch. Unused credits are refunded once it reaches
	// `cancelled`.
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

// Formats requested for each page, as submitted in `input.data.formats`.
type BatchCancelResponseFormats struct {
	// Return the original HTTP response body.
	Bytes bool `json:"bytes"`
	// Return Markdown excerpts relevant to `highlightsParams.query`.
	Highlights bool `json:"highlights"`
	// Return the rendered HTML.
	HTML bool `json:"html"`
	// Return the images found on the page.
	Images bool `json:"images"`
	// Return an object matching `jsonParams.schema`, extracted from the page text and
	// embedded data such as JSON-LD, meta tags, and framework state.
	Json bool `json:"json"`
	// Return the page content, or a video transcript, as Markdown.
	Markdown bool `json:"markdown"`
	// Return fields extracted with custom rules or a prebuilt parser, as `parsed`.
	Parse bool `json:"parse"`
	// Return product details such as name, price, and availability.
	Product bool `json:"product"`
	// Return a screenshot of the page.
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

// `scrape` for a URL list, `crawl` for a site crawl.
type BatchCancelResponseMode string

const (
	BatchCancelResponseModeScrape BatchCancelResponseMode = "scrape"
	BatchCancelResponseModeCrawl  BatchCancelResponseMode = "crawl"
)

// How far the batch got before cancellation.
type BatchCancelResponseProgress struct {
	// Pages that could not be scraped before cancellation.
	Failed int64 `json:"failed" api:"required"`
	// Pages not yet attempted, which will be skipped.
	Pending int64 `json:"pending" api:"required"`
	// Pages scraped successfully before cancellation.
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

// Always `cancelling`. Pages in progress finish, then the batch becomes
// `cancelled`.
type BatchCancelResponseStatus string

const (
	BatchCancelResponseStatusCancelling BatchCancelResponseStatus = "cancelling"
)

// When the batch was created and started.
type BatchCancelResponseTiming struct {
	// When the batch was created.
	CreatedAt string `json:"created_at" api:"required"`
	// When scraping started. `null` if the batch was cancelled before it started.
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

// Credits this request used and your remaining balance.
type BatchCancelResponseKeyMetadata struct {
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
func (r BatchCancelResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *BatchCancelResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchGetResultsResponse struct {
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Results on this page.
	Data []BatchGetResultsResponseDataUnion `json:"data"`
	// Whether more results are available after this page.
	HasMore bool `json:"has_more"`
	// Credits this request used and your remaining balance.
	KeyMetadata BatchGetResultsResponseKeyMetadata `json:"key_metadata"`
	// Pass as `cursor` to fetch the next page. `null` on the last page.
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

// A page where at least one requested format succeeded. Check `result.isPartial`
// and each format's `success` for incomplete results.
type BatchGetResultsResponseDataOk struct {
	// Credits charged for this page, including OCR. Cached results are priced the same
	// as in `POST /web/scrape`.
	CreditsUsed int64 `json:"credits_used" api:"required"`
	// Scrape output, in the same shape as a `POST /web/scrape` response.
	Result BatchGetResultsResponseDataOkResult `json:"result" api:"required"`
	// The page was scraped.
	Status constant.Ok `json:"status" default:"ok"`
	// URL as submitted, or as discovered by the crawl.
	URL string `json:"url" api:"required"`
	// The `itemId` you submitted with this URL.
	ItemID string `json:"itemId"`
	// The `meta` you submitted with this URL.
	Meta map[string]any `json:"meta"`
	// Number of PDF pages read with OCR. Omitted when OCR did not run.
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

// Scrape output, in the same shape as a `POST /web/scrape` response.
type BatchGetResultsResponseDataOkResult struct {
	// The original HTTP response body, unchanged by waits, actions, and filters.
	Bytes BatchGetResultsResponseDataOkResultBytes `json:"bytes" api:"required"`
	// Cache status of this response.
	CacheMetadata BatchGetResultsResponseDataOkResultCacheMetadata `json:"cache_metadata" api:"required"`
	// Relevant Markdown excerpts in page order. `[Heading]` adds context; `…` marks
	// omitted text.
	Highlights BatchGetResultsResponseDataOkResultHighlights `json:"highlights" api:"required"`
	// Rendered HTML of the page, after content filters.
	HTML BatchGetResultsResponseDataOkResultHTML `json:"html" api:"required"`
	// Images found on the page, after content filters. `[]` when there are none.
	Images BatchGetResultsResponseDataOkResultImages `json:"images" api:"required"`
	// Object matching `jsonParams.schema`.
	Json BatchGetResultsResponseDataOkResultJson `json:"json" api:"required"`
	// Page content as Markdown, after content filters.
	Markdown BatchGetResultsResponseDataOkResultMarkdown `json:"markdown" api:"required"`
	// Page metadata. Fields are omitted when not found.
	Metadata BatchGetResultsResponseDataOkResultMetadata `json:"metadata" api:"required"`
	// Fields extracted by `parseParams`. Custom rules return each field after content
	// filters, with `null` when nothing matches (`[]` for `list` rules). Prebuilt
	// parsers return the record or lists described under `parseParams.parser`.
	Parsed BatchGetResultsResponseDataOkResultParsed `json:"parsed" api:"required"`
	// Product details found on the page.
	Product BatchGetResultsResponseDataOkResultProduct `json:"product" api:"required"`
	// Browser profile used for this scrape. `null` when no profile was requested.
	Profile BatchGetResultsResponseDataOkResultProfile `json:"profile" api:"required"`
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Screenshot as a base64 image data URL.
	Screenshot BatchGetResultsResponseDataOkResultScreenshot `json:"screenshot" api:"required"`
	// Final URL after redirects and browser actions.
	URL string `json:"url" api:"required" format:"uri"`
	// `true` when at least one requested output succeeded but others failed or are
	// incomplete. Absent otherwise.
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
	Data BatchGetResultsResponseDataOkResultBytesData `json:"data" api:"required"`
	// Whether you requested this output in `formats`.
	Requested bool `json:"requested" api:"required"`
	// `true` when the output succeeded, `false` when it failed, `null` when it was not
	// requested.
	Success bool `json:"success" api:"required"`
	// Machine-readable error code. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Human-readable explanation of the error and how to resolve it.
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
	// Response body as base64, after HTTP decompression. Up to 50 MiB decoded.
	Base64 string `json:"base64" api:"required"`
	// `Content-Type` of the response.
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

// Cache status of this response.
type BatchGetResultsResponseDataOkResultCacheMetadata struct {
	// Age of the cached data, in milliseconds. `0` for `miss` and `zdr`.
	AgeMs int64 `json:"age_ms" api:"required"`
	// `hit` when served from cache, `miss` when fetched fresh, `zdr` when zero data
	// retention bypassed the cache.
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
	Data []string `json:"data" api:"required"`
	// Whether you requested this output in `formats`.
	Requested bool `json:"requested" api:"required"`
	// `true` when the output succeeded, `false` when it failed, `null` when it was not
	// requested.
	Success bool `json:"success" api:"required"`
	// Machine-readable error code. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Human-readable explanation of the error and how to resolve it.
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

// Rendered HTML of the page, after content filters.
type BatchGetResultsResponseDataOkResultHTML struct {
	Data string `json:"data" api:"required"`
	// Whether you requested this output in `formats`.
	Requested bool `json:"requested" api:"required"`
	// `true` when the output succeeded, `false` when it failed, `null` when it was not
	// requested.
	Success bool `json:"success" api:"required"`
	// Machine-readable error code. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Human-readable explanation of the error and how to resolve it.
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

// Images found on the page, after content filters. `[]` when there are none.
type BatchGetResultsResponseDataOkResultImages struct {
	Data []BatchGetResultsResponseDataOkResultImagesData `json:"data" api:"required"`
	// Whether you requested this output in `formats`.
	Requested bool `json:"requested" api:"required"`
	// `true` when the output succeeded, `false` when it failed, `null` when it was not
	// requested.
	Success bool `json:"success" api:"required"`
	// Machine-readable error code. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Human-readable explanation of the error and how to resolve it.
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
	// Alt text. `null` when the image has none.
	Alt string `json:"alt" api:"required"`
	// Image URL, or a data URI for inline images.
	URL string `json:"url" api:"required" format:"uri"`
	// What the image shows, such as a photo, logo, or icon. Requires `classification`
	// enrichment.
	//
	// Any of "photography", "illustration", "logo", "wordmark", "icon", "pattern",
	// "graphic", "other".
	Classification string `json:"classification"`
	// Hosted copy of the image, valid for 24 hours after capture. Requires `file`
	// enrichment; not returned with ZDR.
	FileURL string `json:"fileUrl" format:"uri"`
	// Image height in pixels. Requires `dimensions` enrichment.
	Height int64 `json:"height"`
	// Image width in pixels. Requires `dimensions` enrichment.
	Width int64 `json:"width"`
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
	Data map[string]any `json:"data" api:"required"`
	// Whether you requested this output in `formats`.
	Requested bool `json:"requested" api:"required"`
	// `true` when the output succeeded, `false` when it failed, `null` when it was not
	// requested.
	Success bool `json:"success" api:"required"`
	// Machine-readable error code. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// `true` when the page was too long to read in full, so values that appear only in
	// the unread part may be missing.
	//
	// Any of true.
	IsTruncated bool `json:"isTruncated"`
	// Human-readable explanation of the error and how to resolve it.
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

// Page content as Markdown, after content filters.
type BatchGetResultsResponseDataOkResultMarkdown struct {
	Data string `json:"data" api:"required"`
	// Whether you requested this output in `formats`.
	Requested bool `json:"requested" api:"required"`
	// `true` when the output succeeded, `false` when it failed, `null` when it was not
	// requested.
	Success bool `json:"success" api:"required"`
	// Machine-readable error code. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Human-readable explanation of the error and how to resolve it.
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
	// Other non-social meta tags that don't have their own field.
	AdditionalMeta map[string]BatchGetResultsResponseDataOkResultMetadataAdditionalMetaUnion `json:"additionalMeta"`
	// Alternate links from `link rel="alternate"` tags, as absolute URLs.
	Alternates []BatchGetResultsResponseDataOkResultMetadataAlternate `json:"alternates"`
	// Author, when present.
	Author string `json:"author"`
	// Absolute canonical URL, when present.
	CanonicalURL string `json:"canonicalUrl"`
	// Best available description, from standard, Open Graph, or Twitter meta tags.
	Description string `json:"description"`
	// Absolute favicon URL, when present.
	Favicon string `json:"favicon"`
	// Headings (`h1` to `h6`) in document order, from the full page before any
	// filtering. Up to 500. Omitted when the page has none.
	Headings []BatchGetResultsResponseDataOkResultMetadataHeading `json:"headings"`
	// Main preview image URL, from Open Graph, Twitter, or image meta tags.
	Image string `json:"image"`
	// JSON-LD structured data blocks parsed from the page.
	JsonLd []map[string]any `json:"jsonLd"`
	// Keywords from the page's `keywords` meta tag.
	Keywords []string `json:"keywords"`
	// Page language, from the `html` `lang` attribute or language meta tags.
	Language string `json:"language"`
	// Last-modified date or timestamp from page metadata, when present.
	ModifiedTime string `json:"modifiedTime"`
	// Open Graph tags, with the `og:` prefix removed and keys in camelCase.
	OpenGraph map[string]BatchGetResultsResponseDataOkResultMetadataOpenGraphUnion `json:"openGraph"`
	// Publish date or timestamp from page metadata, when present.
	PublishedTime string `json:"publishedTime"`
	// `robots` meta directive, when present.
	Robots string `json:"robots"`
	// Site or app name from page metadata.
	SiteName string `json:"siteName"`
	// Best available page title.
	Title string `json:"title"`
	// Twitter card tags, with the `twitter:` prefix removed and keys in camelCase.
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
	// Absolute URL of the alternate version.
	Href string `json:"href" api:"required"`
	// Language or locale of the alternate version, when present.
	Hreflang string `json:"hreflang"`
	// Title of the alternate version, when present.
	Title string `json:"title"`
	// MIME type of the alternate version, when present.
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
	// Heading level, from 1 (`h1`) to 6 (`h6`).
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

// Fields extracted by `parseParams`. Custom rules return each field after content
// filters, with `null` when nothing matches (`[]` for `list` rules). Prebuilt
// parsers return the record or lists described under `parseParams.parser`.
type BatchGetResultsResponseDataOkResultParsed struct {
	Data map[string]any `json:"data" api:"required"`
	// Whether you requested this output in `formats`.
	Requested bool `json:"requested" api:"required"`
	// `true` when the output succeeded, `false` when it failed, `null` when it was not
	// requested.
	Success bool `json:"success" api:"required"`
	// Prebuilt parser that ran, or `custom` for custom rules or when `parse` was not
	// requested.
	//
	// Any of "custom", "google-maps", "google-maps-reviews", "google-play",
	// "google-play-reviews", "app-store", "app-store-reviews", "g2-reviews",
	// "instagram-profile", "instagram-post", "instagram-feed", "tiktok-profile",
	// "tiktok-video", "tiktok-feed", "tiktok-comments", "tiktok-followers",
	// "tiktok-following", "tiktok-user-search", "reddit-post", "reddit-search",
	// "reddit-subreddit", "reddit-user", "yelp-search", "yelp-reviews",
	// "tripadvisor-reviews", "car-rental-availability", "car-rental-locations",
	// "emails", "phone-numbers", "links".
	Type string `json:"type" api:"required"`
	// Machine-readable error code. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Human-readable explanation of the error and how to resolve it.
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
	Data BatchGetResultsResponseDataOkResultProductData `json:"data" api:"required"`
	// Whether you requested this output in `formats`.
	Requested bool `json:"requested" api:"required"`
	// `true` when the output succeeded, `false` when it failed, `null` when it was not
	// requested.
	Success bool `json:"success" api:"required"`
	// Machine-readable error code. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Human-readable explanation of the error and how to resolve it.
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
	// The extracted product. `null` when the page is not a product detail page.
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

// The extracted product. `null` when the page is not a product detail page.
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
	// ISO 4217 currency code of `price`.
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
	// Product variants, such as colors or sizes, with their attributes and images.
	// `[]` when none are found. May not include every variant the store offers.
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
	// Attributes that define this variant, such as color, size, material, or pattern,
	// as stated on the page.
	Attributes map[string]string `json:"attributes" api:"required"`
	// Image URLs the source attaches to this variant.
	Images []string `json:"images" api:"required"`
	// SKU of this variant. `null` when the source does not state one.
	SKU string `json:"sku" api:"required"`
	// URL of this variant or offer, when the source provides one. Several variants can
	// share a URL.
	URL string `json:"url" api:"required"`
	// Stock or ordering availability of this variant.
	//
	// Any of "in_stock", "out_of_stock", "limited_availability", "preorder",
	// "backorder", "made_to_order", "discontinued".
	Availability string `json:"availability" api:"nullable"`
	// Currency code of `price`, such as `USD` or `EUR`, when stated.
	Currency string `json:"currency" api:"nullable"`
	// Price of this variant, when the source states one per variant.
	Price float64 `json:"price" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attributes   respjson.Field
		Images       respjson.Field
		SKU          respjson.Field
		URL          respjson.Field
		Availability respjson.Field
		Currency     respjson.Field
		Price        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataOkResultProductDataProductVariant) RawJSON() string {
	return r.JSON.raw
}
func (r *BatchGetResultsResponseDataOkResultProductDataProductVariant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Browser profile used for this scrape. `null` when no profile was requested.
type BatchGetResultsResponseDataOkResultProfile struct {
	// ID of the browser profile. After creating a profile, save it and pass it as
	// `sharedParams.profile.profileId` with `type: "existing"`.
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
	Data string `json:"data" api:"required" format:"uri"`
	// Whether you requested this output in `formats`.
	Requested bool `json:"requested" api:"required"`
	// `true` when the output succeeded, `false` when it failed, `null` when it was not
	// requested.
	Success bool `json:"success" api:"required"`
	// Machine-readable error code. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Human-readable explanation of the error and how to resolve it.
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
func (r BatchGetResultsResponseDataOkResultKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataOkResultKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A page that could not be scraped.
type BatchGetResultsResponseDataError struct {
	// Machine-readable error code, such as `NOT_FOUND`.
	ErrorCode string `json:"error_code" api:"required"`
	// Human-readable explanation of the error.
	Message string `json:"message" api:"required"`
	// The page could not be scraped.
	Status constant.Error `json:"status" default:"error"`
	// URL as submitted, or as discovered by the crawl.
	URL string `json:"url" api:"required"`
	// Credits charged for this page, including OCR. Cached results are priced the same
	// as in `POST /web/scrape`.
	CreditsUsed int64 `json:"credits_used"`
	// The `itemId` you submitted with this URL.
	ItemID string `json:"itemId"`
	// The `meta` you submitted with this URL.
	Meta map[string]any `json:"meta"`
	// Number of PDF pages read with OCR. Omitted when OCR did not run.
	OcrPages int64 `json:"ocr_pages"`
	// Per-format errors, when the scrape returned a response. Omitted when the page
	// could not be scraped at all.
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

// Per-format errors, when the scrape returned a response. Omitted when the page
// could not be scraped at all.
type BatchGetResultsResponseDataErrorResult struct {
	// The original HTTP response body, unchanged by waits, actions, and filters.
	Bytes BatchGetResultsResponseDataErrorResultBytes `json:"bytes" api:"required"`
	// Cache status of this response.
	CacheMetadata BatchGetResultsResponseDataErrorResultCacheMetadata `json:"cache_metadata" api:"required"`
	// Relevant Markdown excerpts in page order. `[Heading]` adds context; `…` marks
	// omitted text.
	Highlights BatchGetResultsResponseDataErrorResultHighlights `json:"highlights" api:"required"`
	// Rendered HTML of the page, after content filters.
	HTML BatchGetResultsResponseDataErrorResultHTML `json:"html" api:"required"`
	// Images found on the page, after content filters. `[]` when there are none.
	Images BatchGetResultsResponseDataErrorResultImages `json:"images" api:"required"`
	// Object matching `jsonParams.schema`.
	Json BatchGetResultsResponseDataErrorResultJson `json:"json" api:"required"`
	// Page content as Markdown, after content filters.
	Markdown BatchGetResultsResponseDataErrorResultMarkdown `json:"markdown" api:"required"`
	// Page metadata. Fields are omitted when not found.
	Metadata BatchGetResultsResponseDataErrorResultMetadata `json:"metadata" api:"required"`
	// Fields extracted by `parseParams`. Custom rules return each field after content
	// filters, with `null` when nothing matches (`[]` for `list` rules). Prebuilt
	// parsers return the record or lists described under `parseParams.parser`.
	Parsed BatchGetResultsResponseDataErrorResultParsed `json:"parsed" api:"required"`
	// Product details found on the page.
	Product BatchGetResultsResponseDataErrorResultProduct `json:"product" api:"required"`
	// Browser profile used for this scrape. `null` when no profile was requested.
	Profile BatchGetResultsResponseDataErrorResultProfile `json:"profile" api:"required"`
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Screenshot as a base64 image data URL.
	Screenshot BatchGetResultsResponseDataErrorResultScreenshot `json:"screenshot" api:"required"`
	// Final URL after redirects and browser actions.
	URL string `json:"url" api:"required" format:"uri"`
	// `true` when at least one requested output succeeded but others failed or are
	// incomplete. Absent otherwise.
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
	Data BatchGetResultsResponseDataErrorResultBytesData `json:"data" api:"required"`
	// Whether you requested this output in `formats`.
	Requested bool `json:"requested" api:"required"`
	// `true` when the output succeeded, `false` when it failed, `null` when it was not
	// requested.
	Success bool `json:"success" api:"required"`
	// Machine-readable error code. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Human-readable explanation of the error and how to resolve it.
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
	// Response body as base64, after HTTP decompression. Up to 50 MiB decoded.
	Base64 string `json:"base64" api:"required"`
	// `Content-Type` of the response.
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

// Cache status of this response.
type BatchGetResultsResponseDataErrorResultCacheMetadata struct {
	// Age of the cached data, in milliseconds. `0` for `miss` and `zdr`.
	AgeMs int64 `json:"age_ms" api:"required"`
	// `hit` when served from cache, `miss` when fetched fresh, `zdr` when zero data
	// retention bypassed the cache.
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
	Data []string `json:"data" api:"required"`
	// Whether you requested this output in `formats`.
	Requested bool `json:"requested" api:"required"`
	// `true` when the output succeeded, `false` when it failed, `null` when it was not
	// requested.
	Success bool `json:"success" api:"required"`
	// Machine-readable error code. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Human-readable explanation of the error and how to resolve it.
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

// Rendered HTML of the page, after content filters.
type BatchGetResultsResponseDataErrorResultHTML struct {
	Data string `json:"data" api:"required"`
	// Whether you requested this output in `formats`.
	Requested bool `json:"requested" api:"required"`
	// `true` when the output succeeded, `false` when it failed, `null` when it was not
	// requested.
	Success bool `json:"success" api:"required"`
	// Machine-readable error code. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Human-readable explanation of the error and how to resolve it.
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

// Images found on the page, after content filters. `[]` when there are none.
type BatchGetResultsResponseDataErrorResultImages struct {
	Data []BatchGetResultsResponseDataErrorResultImagesData `json:"data" api:"required"`
	// Whether you requested this output in `formats`.
	Requested bool `json:"requested" api:"required"`
	// `true` when the output succeeded, `false` when it failed, `null` when it was not
	// requested.
	Success bool `json:"success" api:"required"`
	// Machine-readable error code. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Human-readable explanation of the error and how to resolve it.
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
	// Alt text. `null` when the image has none.
	Alt string `json:"alt" api:"required"`
	// Image URL, or a data URI for inline images.
	URL string `json:"url" api:"required" format:"uri"`
	// What the image shows, such as a photo, logo, or icon. Requires `classification`
	// enrichment.
	//
	// Any of "photography", "illustration", "logo", "wordmark", "icon", "pattern",
	// "graphic", "other".
	Classification string `json:"classification"`
	// Hosted copy of the image, valid for 24 hours after capture. Requires `file`
	// enrichment; not returned with ZDR.
	FileURL string `json:"fileUrl" format:"uri"`
	// Image height in pixels. Requires `dimensions` enrichment.
	Height int64 `json:"height"`
	// Image width in pixels. Requires `dimensions` enrichment.
	Width int64 `json:"width"`
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
	Data map[string]any `json:"data" api:"required"`
	// Whether you requested this output in `formats`.
	Requested bool `json:"requested" api:"required"`
	// `true` when the output succeeded, `false` when it failed, `null` when it was not
	// requested.
	Success bool `json:"success" api:"required"`
	// Machine-readable error code. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// `true` when the page was too long to read in full, so values that appear only in
	// the unread part may be missing.
	//
	// Any of true.
	IsTruncated bool `json:"isTruncated"`
	// Human-readable explanation of the error and how to resolve it.
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

// Page content as Markdown, after content filters.
type BatchGetResultsResponseDataErrorResultMarkdown struct {
	Data string `json:"data" api:"required"`
	// Whether you requested this output in `formats`.
	Requested bool `json:"requested" api:"required"`
	// `true` when the output succeeded, `false` when it failed, `null` when it was not
	// requested.
	Success bool `json:"success" api:"required"`
	// Machine-readable error code. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Human-readable explanation of the error and how to resolve it.
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
	// Other non-social meta tags that don't have their own field.
	AdditionalMeta map[string]BatchGetResultsResponseDataErrorResultMetadataAdditionalMetaUnion `json:"additionalMeta"`
	// Alternate links from `link rel="alternate"` tags, as absolute URLs.
	Alternates []BatchGetResultsResponseDataErrorResultMetadataAlternate `json:"alternates"`
	// Author, when present.
	Author string `json:"author"`
	// Absolute canonical URL, when present.
	CanonicalURL string `json:"canonicalUrl"`
	// Best available description, from standard, Open Graph, or Twitter meta tags.
	Description string `json:"description"`
	// Absolute favicon URL, when present.
	Favicon string `json:"favicon"`
	// Headings (`h1` to `h6`) in document order, from the full page before any
	// filtering. Up to 500. Omitted when the page has none.
	Headings []BatchGetResultsResponseDataErrorResultMetadataHeading `json:"headings"`
	// Main preview image URL, from Open Graph, Twitter, or image meta tags.
	Image string `json:"image"`
	// JSON-LD structured data blocks parsed from the page.
	JsonLd []map[string]any `json:"jsonLd"`
	// Keywords from the page's `keywords` meta tag.
	Keywords []string `json:"keywords"`
	// Page language, from the `html` `lang` attribute or language meta tags.
	Language string `json:"language"`
	// Last-modified date or timestamp from page metadata, when present.
	ModifiedTime string `json:"modifiedTime"`
	// Open Graph tags, with the `og:` prefix removed and keys in camelCase.
	OpenGraph map[string]BatchGetResultsResponseDataErrorResultMetadataOpenGraphUnion `json:"openGraph"`
	// Publish date or timestamp from page metadata, when present.
	PublishedTime string `json:"publishedTime"`
	// `robots` meta directive, when present.
	Robots string `json:"robots"`
	// Site or app name from page metadata.
	SiteName string `json:"siteName"`
	// Best available page title.
	Title string `json:"title"`
	// Twitter card tags, with the `twitter:` prefix removed and keys in camelCase.
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
	// Absolute URL of the alternate version.
	Href string `json:"href" api:"required"`
	// Language or locale of the alternate version, when present.
	Hreflang string `json:"hreflang"`
	// Title of the alternate version, when present.
	Title string `json:"title"`
	// MIME type of the alternate version, when present.
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
	// Heading level, from 1 (`h1`) to 6 (`h6`).
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

// Fields extracted by `parseParams`. Custom rules return each field after content
// filters, with `null` when nothing matches (`[]` for `list` rules). Prebuilt
// parsers return the record or lists described under `parseParams.parser`.
type BatchGetResultsResponseDataErrorResultParsed struct {
	Data map[string]any `json:"data" api:"required"`
	// Whether you requested this output in `formats`.
	Requested bool `json:"requested" api:"required"`
	// `true` when the output succeeded, `false` when it failed, `null` when it was not
	// requested.
	Success bool `json:"success" api:"required"`
	// Prebuilt parser that ran, or `custom` for custom rules or when `parse` was not
	// requested.
	//
	// Any of "custom", "google-maps", "google-maps-reviews", "google-play",
	// "google-play-reviews", "app-store", "app-store-reviews", "g2-reviews",
	// "instagram-profile", "instagram-post", "instagram-feed", "tiktok-profile",
	// "tiktok-video", "tiktok-feed", "tiktok-comments", "tiktok-followers",
	// "tiktok-following", "tiktok-user-search", "reddit-post", "reddit-search",
	// "reddit-subreddit", "reddit-user", "yelp-search", "yelp-reviews",
	// "tripadvisor-reviews", "car-rental-availability", "car-rental-locations",
	// "emails", "phone-numbers", "links".
	Type string `json:"type" api:"required"`
	// Machine-readable error code. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Human-readable explanation of the error and how to resolve it.
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
	Data BatchGetResultsResponseDataErrorResultProductData `json:"data" api:"required"`
	// Whether you requested this output in `formats`.
	Requested bool `json:"requested" api:"required"`
	// `true` when the output succeeded, `false` when it failed, `null` when it was not
	// requested.
	Success bool `json:"success" api:"required"`
	// Machine-readable error code. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Human-readable explanation of the error and how to resolve it.
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
	// The extracted product. `null` when the page is not a product detail page.
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

// The extracted product. `null` when the page is not a product detail page.
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
	// ISO 4217 currency code of `price`.
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
	// Product variants, such as colors or sizes, with their attributes and images.
	// `[]` when none are found. May not include every variant the store offers.
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
	// Attributes that define this variant, such as color, size, material, or pattern,
	// as stated on the page.
	Attributes map[string]string `json:"attributes" api:"required"`
	// Image URLs the source attaches to this variant.
	Images []string `json:"images" api:"required"`
	// SKU of this variant. `null` when the source does not state one.
	SKU string `json:"sku" api:"required"`
	// URL of this variant or offer, when the source provides one. Several variants can
	// share a URL.
	URL string `json:"url" api:"required"`
	// Stock or ordering availability of this variant.
	//
	// Any of "in_stock", "out_of_stock", "limited_availability", "preorder",
	// "backorder", "made_to_order", "discontinued".
	Availability string `json:"availability" api:"nullable"`
	// Currency code of `price`, such as `USD` or `EUR`, when stated.
	Currency string `json:"currency" api:"nullable"`
	// Price of this variant, when the source states one per variant.
	Price float64 `json:"price" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attributes   respjson.Field
		Images       respjson.Field
		SKU          respjson.Field
		URL          respjson.Field
		Availability respjson.Field
		Currency     respjson.Field
		Price        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResultsResponseDataErrorResultProductDataProductVariant) RawJSON() string {
	return r.JSON.raw
}
func (r *BatchGetResultsResponseDataErrorResultProductDataProductVariant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Browser profile used for this scrape. `null` when no profile was requested.
type BatchGetResultsResponseDataErrorResultProfile struct {
	// ID of the browser profile. After creating a profile, save it and pass it as
	// `sharedParams.profile.profileId` with `type: "existing"`.
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
	Data string `json:"data" api:"required" format:"uri"`
	// Whether you requested this output in `formats`.
	Requested bool `json:"requested" api:"required"`
	// `true` when the output succeeded, `false` when it failed, `null` when it was not
	// requested.
	Success bool `json:"success" api:"required"`
	// Machine-readable error code. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Human-readable explanation of the error and how to resolve it.
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
func (r BatchGetResultsResponseDataErrorResultKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseDataErrorResultKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credits this request used and your remaining balance.
type BatchGetResultsResponseKeyMetadata struct {
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
func (r BatchGetResultsResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResultsResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchSubmitResponse struct {
	// Unique ID of the batch. Pass it to `GET /batch/{batch_id}` to track progress.
	ID string `json:"id" api:"required"`
	// Cache status of this response.
	CacheMetadata BatchSubmitResponseCacheMetadata `json:"cache_metadata" api:"required"`
	// Crawl settings as submitted.
	Crawl CrawlControls `json:"crawl" api:"required"`
	// When the batch was created.
	CreatedAt string `json:"created_at" api:"required"`
	// Credits reserved for the batch.
	Credits BatchSubmitResponseCredits `json:"credits" api:"required"`
	// Formats requested for each page, as submitted in `input.data.formats`.
	Formats BatchSubmitResponseFormats `json:"formats" api:"required"`
	// Page counts from your submission.
	Input Intake `json:"input" api:"required"`
	// URLs rejected from your list, up to the first 100.
	InvalidURLs []BatchSubmitResponseInvalidURL `json:"invalid_urls" api:"required"`
	// `scrape` for a URL list, `crawl` for a site crawl.
	//
	// Any of "scrape", "crawl".
	Mode BatchSubmitResponseMode `json:"mode" api:"required"`
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Current state. `pending` batches are being validated or are waiting for a slot
	// under your plan's concurrent batch limit, `queued` batches are ready to start,
	// and `running` batches are scraping. `completed`, `cancelled`, and `failed` are
	// final.
	//
	// Any of "pending", "queued", "running", "cancelling", "completed", "cancelled",
	// "failed".
	Status BatchSubmitResponseStatus `json:"status" api:"required"`
	// Tags set when the batch was submitted.
	Tags []string `json:"tags" api:"required"`
	// Credits this request used and your remaining balance.
	KeyMetadata BatchSubmitResponseKeyMetadata `json:"key_metadata"`
	// Secret for verifying `X-Context-Signature` on this batch's webhooks. Returned
	// only here and only when you set a webhook, so store it.
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

// Cache status of this response.
type BatchSubmitResponseCacheMetadata struct {
	// Age of the cached data, in milliseconds. `0` for `miss` and `zdr`.
	AgeMs int64 `json:"age_ms" api:"required"`
	// `hit` when served from cache, `miss` when fetched fresh, `zdr` when zero data
	// retention bypassed the cache.
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

// Credits reserved for the batch.
type BatchSubmitResponseCredits struct {
	// Credits reserved after validation, before scraping starts. `0` while `pending`.
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

// Formats requested for each page, as submitted in `input.data.formats`.
type BatchSubmitResponseFormats struct {
	// Return the original HTTP response body.
	Bytes bool `json:"bytes"`
	// Return Markdown excerpts relevant to `highlightsParams.query`.
	Highlights bool `json:"highlights"`
	// Return the rendered HTML.
	HTML bool `json:"html"`
	// Return the images found on the page.
	Images bool `json:"images"`
	// Return an object matching `jsonParams.schema`, extracted from the page text and
	// embedded data such as JSON-LD, meta tags, and framework state.
	Json bool `json:"json"`
	// Return the page content, or a video transcript, as Markdown.
	Markdown bool `json:"markdown"`
	// Return fields extracted with custom rules or a prebuilt parser, as `parsed`.
	Parse bool `json:"parse"`
	// Return product details such as name, price, and availability.
	Product bool `json:"product"`
	// Return a screenshot of the page.
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

// `scrape` for a URL list, `crawl` for a site crawl.
type BatchSubmitResponseMode string

const (
	BatchSubmitResponseModeScrape BatchSubmitResponseMode = "scrape"
	BatchSubmitResponseModeCrawl  BatchSubmitResponseMode = "crawl"
)

// Current state. `pending` batches are being validated or are waiting for a slot
// under your plan's concurrent batch limit, `queued` batches are ready to start,
// and `running` batches are scraping. `completed`, `cancelled`, and `failed` are
// final.
type BatchSubmitResponseStatus string

const (
	BatchSubmitResponseStatusPending    BatchSubmitResponseStatus = "pending"
	BatchSubmitResponseStatusQueued     BatchSubmitResponseStatus = "queued"
	BatchSubmitResponseStatusRunning    BatchSubmitResponseStatus = "running"
	BatchSubmitResponseStatusCancelling BatchSubmitResponseStatus = "cancelling"
	BatchSubmitResponseStatusCompleted  BatchSubmitResponseStatus = "completed"
	BatchSubmitResponseStatusCancelled  BatchSubmitResponseStatus = "cancelled"
	BatchSubmitResponseStatusFailed     BatchSubmitResponseStatus = "failed"
)

// Credits this request used and your remaining balance.
type BatchSubmitResponseKeyMetadata struct {
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
func (r BatchSubmitResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *BatchSubmitResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchListParams struct {
	// Cursor from the previous page's `next_cursor`. Omit for the first page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of batches to return per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Search term matched against batch IDs, crawl sources (start URL or sitemap
	// domain), and tags.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// How `q` matches: `prefix` matches the start of a term, for search-as-you-type,
	// and `exact` matches whole terms. Defaults to `prefix`.
	//
	// Any of "exact", "prefix".
	SearchType BatchListParamsSearchType `query:"search_type,omitzero" json:"-"`
	// Filter by batch status.
	//
	// Any of "pending", "queued", "running", "cancelling", "completed", "cancelled",
	// "failed".
	Status BatchListParamsStatus `query:"status,omitzero" json:"-"`
	// Comma-separated tags. Returns batches that have any of them. You can also repeat
	// the `tags` parameter.
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

// How `q` matches: `prefix` matches the start of a term, for search-as-you-type,
// and `exact` matches whole terms. Defaults to `prefix`.
type BatchListParamsSearchType string

const (
	BatchListParamsSearchTypeExact  BatchListParamsSearchType = "exact"
	BatchListParamsSearchTypePrefix BatchListParamsSearchType = "prefix"
)

// Filter by batch status.
type BatchListParamsStatus string

const (
	BatchListParamsStatusPending    BatchListParamsStatus = "pending"
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
	// Cursor from the previous page's `next_cursor`. Omit for the first page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results to return per page. A page can hold fewer to stay
	// under about 8 MB, so page with `next_cursor` instead of counting records.
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
	// What to scrape: a list of URLs (`scrape`) or a site crawl (`crawl`).
	Input BatchSubmitParamsInputUnion `json:"input,omitzero" api:"required"`
	// Unique key for this submission. Retrying with the same key and body returns the
	// original batch instead of creating another; a different body returns `409`.
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	// Access policy IDs to apply to every page; each URL must pass all of them.
	// Requires access policies enabled for your organization; otherwise returns
	// `ACCESS_POLICIES_NOT_ENABLED`.
	Policies []string `json:"policies,omitzero"`
	// Tags to store on the batch. Filter `GET /batch/list` by them later.
	Tags []string `json:"tags,omitzero"`
	// Webhook to notify when the batch finishes. Omit `retry` for a single attempt, or
	// pass `retry: {}` for the default retry schedule.
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
	// URLs to scrape, with the formats and options to apply to each page.
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
	OfScrapeAnything *BatchSubmitParamsInputScrapeDataScrapeAnything `json:",omitzero,inline"`
	OfMarkdown       *BatchSubmitParamsInputScrapeDataMarkdown       `json:",omitzero,inline"`
	OfHTML           *BatchSubmitParamsInputScrapeDataHTML           `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputScrapeDataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfScrapeAnything, u.OfMarkdown, u.OfHTML)
}
func (u *BatchSubmitParamsInputScrapeDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The same formats and options as `POST /web/scrape`, with `urls` in place of
// `url`. Options apply to every page.
//
// The properties Formats, URLs are required.
type BatchSubmitParamsInputScrapeDataScrapeAnything struct {
	// Outputs to return. Set at least one to `true`.
	Formats BatchSubmitParamsInputScrapeDataScrapeAnythingFormats `json:"formats,omitzero" api:"required"`
	// Pages to scrape, up to 100,000.
	URLs []BatchSubmitParamsInputScrapeDataScrapeAnythingURL `json:"urls,omitzero" api:"required"`
	// Maximum age of a cached result, in milliseconds. `0` always fetches fresh.
	// Defaults to 3 days (259200000 ms). Maximum: 1 year (31536000000 ms).
	MaxAgeMs param.Opt[int64] `json:"maxAgeMs,omitzero"`
	// Query for the `highlights` output. Required when `formats.highlights` is `true`,
	// and allowed only then.
	HighlightsParams BatchSubmitParamsInputScrapeDataScrapeAnythingHighlightsParams `json:"highlightsParams,omitzero"`
	// Options for the `images` output. Requires `formats.images: true`.
	ImageParams BatchSubmitParamsInputScrapeDataScrapeAnythingImageParams `json:"imageParams,omitzero"`
	// Schema and instructions for the `json` output. Required when `formats.json` is
	// `true`, and allowed only then.
	JsonParams BatchSubmitParamsInputScrapeDataScrapeAnythingJsonParams `json:"jsonParams,omitzero"`
	// Options for the `markdown` output. Requires `formats.markdown: true`.
	MarkdownParams BatchSubmitParamsInputScrapeDataScrapeAnythingMarkdownParams `json:"markdownParams,omitzero"`
	// Custom rules or a prebuilt parser for the `parsed` output. Required when
	// `formats.parse` is `true`, and allowed only then.
	ParseParams BatchSubmitParamsInputScrapeDataScrapeAnythingParseParamsUnion `json:"parseParams,omitzero"`
	// Options for the `product` output. Requires `formats.product: true`.
	ProductParams BatchSubmitParamsInputScrapeDataScrapeAnythingProductParams `json:"productParams,omitzero"`
	// Options for the `screenshot` output. Requires `formats.screenshot: true`.
	ScreenshotParams BatchSubmitParamsInputScrapeDataScrapeAnythingScreenshotParams `json:"screenshotParams,omitzero"`
	// Browser and content settings shared by all outputs.
	SharedParams BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParams `json:"sharedParams,omitzero"`
	// Labels for this request. Filter usage and request logs by them in the dashboard
	// or with `GET /org/logs`.
	Tags []string `json:"tags,omitzero"`
	// Time limit for each page, with the same behavior and defaults as in
	// `POST /web/scrape`.
	TimeoutOpts BatchSubmitParamsInputScrapeDataScrapeAnythingTimeoutOpts `json:"timeoutOpts,omitzero"`
	// Zero data retention is not available for batches, because batches store their
	// results. Sending `enabled` returns `ZDR_NOT_SUPPORTED`.
	//
	// Any of "disabled".
	Zdr string `json:"zdr,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataScrapeAnything) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataScrapeAnything
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataScrapeAnything) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputScrapeDataScrapeAnything](
		"zdr", "disabled",
	)
}

// Outputs to return. Set at least one to `true`.
type BatchSubmitParamsInputScrapeDataScrapeAnythingFormats struct {
	// Return the original HTTP response body.
	Bytes param.Opt[bool] `json:"bytes,omitzero"`
	// Return Markdown excerpts relevant to `highlightsParams.query`.
	Highlights param.Opt[bool] `json:"highlights,omitzero"`
	// Return the rendered HTML.
	HTML param.Opt[bool] `json:"html,omitzero"`
	// Return the images found on the page.
	Images param.Opt[bool] `json:"images,omitzero"`
	// Return an object matching `jsonParams.schema`, extracted from the page text and
	// embedded data such as JSON-LD, meta tags, and framework state.
	Json param.Opt[bool] `json:"json,omitzero"`
	// Return the page content, or a video transcript, as Markdown.
	Markdown param.Opt[bool] `json:"markdown,omitzero"`
	// Return fields extracted with custom rules or a prebuilt parser, as `parsed`.
	Parse param.Opt[bool] `json:"parse,omitzero"`
	// Return product details such as name, price, and availability.
	Product param.Opt[bool] `json:"product,omitzero"`
	// Return a screenshot of the page.
	Screenshot param.Opt[bool] `json:"screenshot,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataScrapeAnythingFormats) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataScrapeAnythingFormats
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataScrapeAnythingFormats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A page to scrape, with optional `itemId` and `meta` to match it to its result.
//
// The property URL is required.
type BatchSubmitParamsInputScrapeDataScrapeAnythingURL struct {
	// URL of the page.
	URL string `json:"url" api:"required"`
	// Your ID for this page, returned with its result. The same URL can use different
	// IDs.
	ItemID param.Opt[string] `json:"itemId,omitzero"`
	// Custom JSON returned unchanged with this page's result.
	Meta map[string]any `json:"meta,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataScrapeAnythingURL) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataScrapeAnythingURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataScrapeAnythingURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Query for the `highlights` output. Required when `formats.highlights` is `true`,
// and allowed only then.
//
// The property Query is required.
type BatchSubmitParamsInputScrapeDataScrapeAnythingHighlightsParams struct {
	// Question or topic to find relevant passages for.
	Query string `json:"query" api:"required"`
	// Maximum combined length of the returned passages, in characters.
	MaxCharacters param.Opt[int64] `json:"maxCharacters,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataScrapeAnythingHighlightsParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataScrapeAnythingHighlightsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataScrapeAnythingHighlightsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Options for the `images` output. Requires `formats.images: true`.
type BatchSubmitParamsInputScrapeDataScrapeAnythingImageParams struct {
	// Set to `visual` to drop visually duplicate images, keeping the largest copy.
	//
	// Any of "none", "visual".
	Dedupe string `json:"dedupe,omitzero"`
	// Extra data to add to each image: `dimensions` adds `width` and `height`,
	// `classification` adds `classification`, and `file` adds a hosted `fileUrl`.
	//
	// Any of "dimensions", "classification", "file".
	Enrich []string `json:"enrich,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataScrapeAnythingImageParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataScrapeAnythingImageParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataScrapeAnythingImageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputScrapeDataScrapeAnythingImageParams](
		"dedupe", "none", "visual",
	)
}

// Schema and instructions for the `json` output. Required when `formats.json` is
// `true`, and allowed only then.
//
// The property Schema is required.
type BatchSubmitParamsInputScrapeDataScrapeAnythingJsonParams struct {
	// JSON Schema for the object to extract, not an example object. Must describe a
	// top-level object, up to 50 KB. Make fields optional or nullable when the page
	// may not state them.
	Schema map[string]any `json:"schema,omitzero" api:"required"`
	// Extra guidance, such as which facts to prefer or how to read a field.
	Instructions param.Opt[string] `json:"instructions,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataScrapeAnythingJsonParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataScrapeAnythingJsonParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataScrapeAnythingJsonParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Options for the `markdown` output. Requires `formats.markdown: true`.
type BatchSubmitParamsInputScrapeDataScrapeAnythingMarkdownParams struct {
	// Include image references in the Markdown (`html-conversion` only).
	IncludeImages param.Opt[bool] `json:"includeImages,omitzero"`
	// Keep link URLs in the Markdown (`html-conversion` only).
	IncludeLinks param.Opt[bool] `json:"includeLinks,omitzero"`
	// How to write base64 images when `includeImages` is `true`: `placeholder`
	// shortens the data, `preserve` keeps it in full (`html-conversion` only).
	//
	// Any of "placeholder", "preserve".
	InlineImages string `json:"inlineImages,omitzero"`
	// Markdown sources in order of preference: `html-conversion` converts the page
	// HTML, `markdown-direct` uses Markdown the site publishes.
	//
	// Any of "html-conversion", "markdown-direct".
	SourcePreference []string `json:"sourcePreference,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataScrapeAnythingMarkdownParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataScrapeAnythingMarkdownParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataScrapeAnythingMarkdownParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputScrapeDataScrapeAnythingMarkdownParams](
		"inlineImages", "placeholder", "preserve",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputScrapeDataScrapeAnythingParseParamsUnion struct {
	OfCustomParser   *BatchSubmitParamsInputScrapeDataScrapeAnythingParseParamsCustomParser   `json:",omitzero,inline"`
	OfPrebuiltParser *BatchSubmitParamsInputScrapeDataScrapeAnythingParseParamsPrebuiltParser `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputScrapeDataScrapeAnythingParseParamsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCustomParser, u.OfPrebuiltParser)
}
func (u *BatchSubmitParamsInputScrapeDataScrapeAnythingParseParamsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property Rules is required.
type BatchSubmitParamsInputScrapeDataScrapeAnythingParseParamsCustomParser struct {
	// Output field names mapped to a CSS selector (such as `h1` or `a@href`) or a rule
	// object. Up to 100 fields, nested up to 5 levels.
	Rules map[string]BatchSubmitParamsInputScrapeDataScrapeAnythingParseParamsCustomParserRuleUnion `json:"rules,omitzero" api:"required"`
	// Extract fields with custom CSS rules. The default when `type` is omitted.
	//
	// Any of "custom".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataScrapeAnythingParseParamsCustomParser) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataScrapeAnythingParseParamsCustomParser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataScrapeAnythingParseParamsCustomParser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputScrapeDataScrapeAnythingParseParamsCustomParser](
		"type", "custom",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputScrapeDataScrapeAnythingParseParamsCustomParserRuleUnion struct {
	OfString                                                                     param.Opt[string]                                                                `json:",omitzero,inline"`
	OfBatchSubmitsInputScrapeDataScrapeAnythingParseParamsCustomParserRuleObject *BatchSubmitParamsInputScrapeDataScrapeAnythingParseParamsCustomParserRuleObject `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputScrapeDataScrapeAnythingParseParamsCustomParserRuleUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfBatchSubmitsInputScrapeDataScrapeAnythingParseParamsCustomParserRuleObject)
}
func (u *BatchSubmitParamsInputScrapeDataScrapeAnythingParseParamsCustomParserRuleUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property Selector is required.
type BatchSubmitParamsInputScrapeDataScrapeAnythingParseParamsCustomParserRuleObject struct {
	// CSS selector to match within the current page or parent rule.
	Selector string `json:"selector" api:"required"`
	// What to return: `text`, `html`, an attribute such as `@href`, or nested field
	// rules. Defaults to `text`.
	Output string `json:"output,omitzero"`
	// Return the first match with `item` or all matches with `list`.
	//
	// Any of "item", "list".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataScrapeAnythingParseParamsCustomParserRuleObject) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataScrapeAnythingParseParamsCustomParserRuleObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataScrapeAnythingParseParamsCustomParserRuleObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputScrapeDataScrapeAnythingParseParamsCustomParserRuleObject](
		"type", "item", "list",
	)
}

// The properties Parser, Type are required.
type BatchSubmitParamsInputScrapeDataScrapeAnythingParseParamsPrebuiltParser struct {
	// Prebuilt parser to run. Each one accepts specific URLs and returns its result in
	// `parsed.data`. `emails`, `phone-numbers`, and `links` read the rendered page and
	// respect content filters; the others read their source directly, independently of
	// content filters and other outputs.
	//
	//   - `google-maps`: Google Maps place, search, and internal data URLs. Returns
	//     businesses in `data.results`; a search URL returns up to 200 places. Results
	//     use the country set in the URL, or the US when none is set, and come back with
	//     English labels.
	//   - `google-maps-reviews`: Google Maps place, search, and `?cid=` URLs, plus
	//     `search.google.com/local/reviews` URLs, so the `google_maps_url` and
	//     `google_reviews_url` fields of `google-maps` results work as input. Returns
	//     the reviews of the place in `data.reviews` (`rating`, `date`, `isoDate`,
	//     `snippet`, `translatedSnippet`, `likes`, `user`, `media`, `response`, `link`,
	//     `id`, `language`, guided ratings in `details`, and `user.isLocalGuide`), plus
	//     the place `featureId`, `cid`, and `sort` in `data`. `sort` and `limit` set the
	//     order and count; `oldest` is exact. An unknown place returns no reviews.
	//   - `google-play`: Play Store app details URLs
	//     (`play.google.com/store/apps/details?id=…`). Returns one app in `data`:
	//     details, developer contact, ratings, price, permissions, data safety, the
	//     storefront country the record describes, and `is_available` (whether the app
	//     is offered there). Ratings are counted per storefront, so `rating`,
	//     `rating_count`, and `rating_histogram` are `null` where nobody rated the app.
	//     The language and storefront country set in the URL apply (default English and
	//     the US); without a storefront country, an app not offered in the US is read
	//     from the country of its developer.
	//   - `google-play-reviews`: Same URLs as `google-play`. Returns up to the 200
	//     newest reviews in the language set in the URL, with developer replies, in
	//     `data.reviews`; an app with fewer such reviews returns them all.
	//   - `app-store`: App Store app URLs (`apps.apple.com/us/app/…/id…`). Returns one
	//     app in `data`: details, developer, ratings, price, version, and screenshots.
	//     The country in the path sets the storefront (default `us`).
	//   - `app-store-reviews`: Same URLs as `app-store`. Returns up to the 500 newest
	//     reviews in `data.reviews`.
	//   - `g2-reviews`: G2 product URLs (`g2.com/products/…`), following renamed
	//     products. Returns the 25 most recently published reviews in `data.reviews`
	//     (rating, title, likes, dislikes, `problems_solved`, reviewer role, company
	//     size, and industry), plus `data.product`. Keeps star-rating filters from the
	//     URL (`filters[nps_score][]=1` to `5`) and ignores other filters, sort orders,
	//     and pages.
	//   - `instagram-profile`: Public profile URLs (`instagram.com/{username}`). Returns
	//     one profile in `data`: bio, links, follower and post counts, category,
	//     business address, `similar_accounts`, and `latest_posts` (up to the 12 posts
	//     at the top of the profile, newest first). Instagram lists pinned posts first,
	//     so posts with `is_pinned: true` can be older. When Instagram shows only a
	//     rounded post count (such as 32K), `post_count` is that rounded number and
	//     `is_post_count_rounded` is `true`.
	//   - `instagram-post`: Post and reel URLs (`/p/…`, `/reel/…`). Returns one post in
	//     `data`: caption, hashtags, mentions, like, comment, and view counts
	//     (`like_count` and `view_count` are `null` when the owner hides them), media
	//     URLs, tagged users, collaborators, location, and about the top 14 comments in
	//     `data.comments`.
	//   - `instagram-feed`: Hashtag and keyword URLs (`/explore/tags/…`, `/popular/…`)
	//     return the top posts; location URLs (`/explore/locations/{id}/…`) return
	//     recent posts. Posts are in `data.posts`; keyword posts carry only the
	//     shortcode, owner, thumbnail, and views, plus the caption when Instagram shows
	//     it. All three Instagram parsers read public data without signing in, so
	//     private accounts return no posts and follower lists are unavailable.
	//   - `tiktok-profile`: Public profile URLs (`tiktok.com/@{username}`). Returns one
	//     profile in `data`: nickname, bio, bio link, follower, following, like, and
	//     video counts, verified, private, and business flags, and business category.
	//     Does not list videos.
	//   - `tiktok-video`: Video and photo post URLs (`/@{username}/video/{id}`,
	//     `/@{username}/photo/{id}`). Returns one post in `data`: caption, hashtags,
	//     mentions, author, view, like, comment, share, save, and repost counts,
	//     duration, thumbnail, a `video_url` that downloads without cookies,
	//     `image_urls` for photo posts, music, location, and ad, paid-partnership, and
	//     AI-generated flags.
	//   - `tiktok-feed`: Returns videos in the `tiktok-video` shape in `data.videos`:
	//     the latest 15 from a profile URL, the top 30 from a hashtag URL
	//     (`/tag/{name}`), or the top 12 from a search URL (`/search?q={query}`).
	//   - `tiktok-comments`: Video and photo post URLs. Returns the first 20 comments in
	//     `data.comments`: text, `posted_at`, like and reply counts, whether the author
	//     liked it, and the commenter.
	//   - `tiktok-followers`, `tiktok-following`: Profile URLs. Return up to 30 accounts
	//     in `data.users`, in the `tiktok-profile` shape.
	//   - `tiktok-user-search`: User search URLs (`/search/user?q={query}`). Returns the
	//     top 10 accounts in `data.users`: `id`, `username`, `nickname`, `biography`,
	//     `avatar_url`, `sec_uid`, `follower_count`, `like_count`, and `is_verified`.
	//     The TikTok parsers accept `tiktok.com`, `www.tiktok.com`, and `m.tiktok.com`
	//     URLs, not short links (`vm.tiktok.com`, `vt.tiktok.com`, `tiktok.com/t/…`).
	//   - `reddit-post`: Post URLs (`reddit.com/r/{subreddit}/comments/{id}`,
	//     `reddit.com/comments/{id}`, `redd.it/{id}`). Returns the post in `data.post`:
	//     title, body as Markdown, type (`text`, `link`, `image`, `gallery`, or
	//     `video`), `link_url`, subreddit, author, flairs, score, `upvote_ratio`,
	//     comment and crosspost counts, NSFW, spoiler, locked, stickied, and archived
	//     flags, images (every gallery image in order), and video (an MP4 `url` plus an
	//     `hls_url` that carries the audio). `data.comments` holds the comment tree;
	//     each comment has `id`, `author`, `body` (Markdown), `score`, `posted_at`,
	//     `depth`, `parent_id`, `url`, `is_submitter`, `is_deleted`, and its `replies`.
	//     `limit` caps the total number of comments (default 100), filled shallowest
	//     first so whole conversations are kept, and `data.has_more_comments` says
	//     whether the thread has more. The `sort` in the URL (`best`, `top`, `new`,
	//     `controversial`, `old`, `qa`) orders the comments.
	//   - `reddit-search`: Search URLs (`reddit.com/search?q={query}`, or
	//     `reddit.com/r/{subreddit}/search?q={query}` to search one subreddit). With
	//     `type=link` or no `type`, returns posts in the `reddit-post` shape in
	//     `data.posts`, ordered by the `sort` in the URL (`relevance`, `hot`, `top`,
	//     `new`, `comments`) within its `t` time range (`hour`, `day`, `week`, `month`,
	//     `year`, `all`). `type=sr` returns communities in `data.communities` (name,
	//     title, description, sidebar, `subscriber_count`, `created_at`, type, NSFW and
	//     quarantine flags, language, category, icon and banner URLs, and allowed post
	//     types). `type=user` returns accounts in `data.users` (username,
	//     `display_name`, bio, `created_at`, post and comment karma, avatar and banner
	//     URLs, and verified, employee, moderator, premium, and NSFW flags). `limit`
	//     caps the number of results (default 100), in the order Reddit returns them,
	//     and `data.has_more` says whether Reddit has more.
	//   - `reddit-subreddit`: Subreddit URLs (`reddit.com/r/{subreddit}`, optionally
	//     followed by `/hot`, `/new`, `/top`, `/rising`, or `/controversial`, with `t`
	//     setting the time range of `top` and `controversial`, default `day`). Returns
	//     `data.subreddit` in the community shape of `reddit-search`, `data.rules`
	//     (`name`, `description`, `applies_to`, `violation_reason`, `created_at`), and
	//     `data.posts` in the `reddit-post` shape, ordered like the tab (default `hot`)
	//     with stickied posts flagged. `r/all` and `r/popular` return posts with `null`
	//     `subreddit` and `rules`. `limit` caps the posts (default 100) and
	//     `data.has_more` says whether Reddit has more. Private and quarantined
	//     subreddits return `WEBSITE_ACCESS_ERROR`.
	//   - `reddit-user`: Profile URLs (`reddit.com/user/{username}` or
	//     `reddit.com/u/{username}`, optionally followed by `/overview`, `/submitted`,
	//     or `/comments`, with `sort` set to `new`, `hot`, `top`, or `controversial` and
	//     `t` setting the time range of `top` and `controversial`, default `all`).
	//     Returns `data.user` (`id`, `username`, `display_name`, bio, `created_at`,
	//     total, post, comment, awardee, and awarder karma, avatar and banner URLs, and
	//     verified, suspended, employee, moderator, premium, and NSFW flags),
	//     `data.trophies` (`name`, `description`, `icon_url`, `granted_at`, `url`),
	//     `data.moderated_communities` (`name`, `title`, `url`, `subscriber_count`,
	//     `icon_url`, `is_nsfw`), `data.posts` in the `reddit-post` shape, and
	//     `data.comments`, each with `id`, `author`, `body` (Markdown), `score`,
	//     `posted_at`, `edited_at`, `url`, `subreddit`, `parent_id` (`null` for a
	//     top-level comment), `is_submitter`, `is_stickied`, `distinguished`,
	//     `award_count`, and the post it was left on (`id`, `title`, `url`, `author`,
	//     `is_nsfw`). The overview tab returns both posts and comments; `submitted`
	//     returns `null` comments and `comments` returns `null` posts. `limit` caps
	//     posts and comments together (default 100) and `data.has_more` says whether
	//     Reddit has more. Users who hide their activity return empty lists, and
	//     suspended accounts return their profile with `null` posts and comments.
	//   - `yelp-search`: Yelp search URLs with `find_desc` and `find_loc`; `start` sets
	//     the result offset. Returns up to 20 ranked organic businesses in
	//     `data.businesses`, with `total` and `pagination`.
	//   - `yelp-reviews`: Yelp business URLs (`yelp.com/biz/{alias}`); optional `start`
	//     sets the offset. Returns up to 20 of the latest public reviews, always newest
	//     first, in `data.reviews`, with business details and `pagination.next_offset`.
	//     Each review has rating, full text, `updated_at`, and author (`id` from the
	//     reviewer profile URL, `url` linking to that profile, `name`, `location`, and
	//     `review_count`); owner replies and photos are not included. Unsupported sort
	//     and review filters are rejected; tracking parameters are ignored.
	//   - `tripadvisor-reviews`: Tripadvisor hotel, restaurant, and attraction review
	//     URLs (`tripadvisor.com/Hotel_Review-g{geo}-d{id}-Reviews-….html`, on any
	//     Tripadvisor country domain); `-or{N}-` in the path sets the offset. Returns 20
	//     reviews in `data.reviews` in the order Tripadvisor shows them, with
	//     `data.location` (`id`, `name`, `type`, `url`, `review_count`) and
	//     `pagination.next_offset`. Each review has `id`, `url`, `title`, full text,
	//     `rating`, `language`, `published_at`, `visit_month`, `helpful_votes`,
	//     `subratings`, `photo_urls`, `owner_response` (`text`, `published_at`,
	//     `author_name`), and `author` (`id`, `username`, `name`, `url`,
	//     `contribution_count`). An unknown location returns no reviews.
	//   - `car-rental-availability`: avis.com and budget.com vehicle-availability search
	//     URLs. Returns every vehicle with pay-now and pay-later prices in
	//     `data.vehicles`.
	//   - `car-rental-locations`: avis.com and budget.com `/locations?keyword=` URLs.
	//     Returns rental branches with codes, addresses, coordinates, and hours in
	//     `data.locations`.
	//   - `emails`: Any page. Returns its email addresses in `data.emails`
	//     (`[{ email }]`), read from `mailto:` links, Cloudflare-protected addresses,
	//     and visible text.
	//   - `phone-numbers`: Any page. Returns its valid phone numbers in E.164 format in
	//     `data.phone_numbers` (`[{ phone }]`), read from `tel:` links and visible text.
	//     Numbers without a country code are read for the country of the site TLD (US by
	//     default).
	//   - `links`: Any page. Returns its unique HTTP and HTTPS links as absolute URLs in
	//     `data.links` (`[{ url, text }]`). `text` falls back to `aria-label` or image
	//     alt text and is `null` when a link has none.
	//
	// Any of "google-maps", "google-maps-reviews", "google-play",
	// "google-play-reviews", "app-store", "app-store-reviews", "g2-reviews",
	// "instagram-profile", "instagram-post", "instagram-feed", "tiktok-profile",
	// "tiktok-video", "tiktok-feed", "tiktok-comments", "tiktok-followers",
	// "tiktok-following", "tiktok-user-search", "reddit-post", "reddit-search",
	// "reddit-subreddit", "reddit-user", "yelp-search", "yelp-reviews",
	// "tripadvisor-reviews", "car-rental-availability", "car-rental-locations",
	// "emails", "phone-numbers", "links".
	Parser string `json:"parser,omitzero" api:"required"`
	// Maximum number of reviews (`google-maps-reviews`, default 100), comments
	// (`reddit-post`, default 100), or results (`reddit-search`, `reddit-subreddit`,
	// and `reddit-user`, default 100) to return. Only these parsers accept it.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// Order of `data.reviews` for `google-maps-reviews`. `newest`, `relevant`,
	// `highest`, and `lowest` use the same orders as Google Maps; `oldest` returns the
	// earliest reviews first.
	//
	// Any of "newest", "oldest", "relevant", "highest", "lowest".
	Sort string `json:"sort,omitzero"`
	// Use a prebuilt parser instead of custom rules.
	//
	// This field can be elided, and will marshal its zero value as "prebuilt".
	Type constant.Prebuilt `json:"type" default:"prebuilt"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataScrapeAnythingParseParamsPrebuiltParser) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataScrapeAnythingParseParamsPrebuiltParser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataScrapeAnythingParseParamsPrebuiltParser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputScrapeDataScrapeAnythingParseParamsPrebuiltParser](
		"parser", "google-maps", "google-maps-reviews", "google-play", "google-play-reviews", "app-store", "app-store-reviews", "g2-reviews", "instagram-profile", "instagram-post", "instagram-feed", "tiktok-profile", "tiktok-video", "tiktok-feed", "tiktok-comments", "tiktok-followers", "tiktok-following", "tiktok-user-search", "reddit-post", "reddit-search", "reddit-subreddit", "reddit-user", "yelp-search", "yelp-reviews", "tripadvisor-reviews", "car-rental-availability", "car-rental-locations", "emails", "phone-numbers", "links",
	)
	apijson.RegisterFieldValidator[BatchSubmitParamsInputScrapeDataScrapeAnythingParseParamsPrebuiltParser](
		"sort", "newest", "oldest", "relevant", "highest", "lowest",
	)
}

// Options for the `product` output. Requires `formats.product: true`.
type BatchSubmitParamsInputScrapeDataScrapeAnythingProductParams struct {
	// Drop visually duplicate product images, keeping the largest copy.
	DedupeImages param.Opt[bool] `json:"dedupeImages,omitzero"`
	// Ignored. The AI fallback runs automatically.
	//
	// Deprecated: deprecated
	UseAIFallback param.Opt[bool] `json:"useAIFallback,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataScrapeAnythingProductParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataScrapeAnythingProductParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataScrapeAnythingProductParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Options for the `screenshot` output. Requires `formats.screenshot: true`.
type BatchSubmitParamsInputScrapeDataScrapeAnythingScreenshotParams struct {
	// What to capture: `viewport`, `fullPage`, one element, or a rectangle. Up to 40
	// megapixels.
	Area BatchSubmitParamsInputScrapeDataScrapeAnythingScreenshotParamsAreaUnion `json:"area,omitzero"`
	// Image format for the screenshot.
	//
	// Any of "png", "jpeg", "webp".
	Format string `json:"format,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataScrapeAnythingScreenshotParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataScrapeAnythingScreenshotParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataScrapeAnythingScreenshotParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputScrapeDataScrapeAnythingScreenshotParams](
		"format", "png", "jpeg", "webp",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputScrapeDataScrapeAnythingScreenshotParamsAreaUnion struct {
	// Check if union is this variant with !param.IsOmitted(union.OfPage)
	OfPage      param.Opt[string]                                                            `json:",omitzero,inline"`
	OfElement   *BatchSubmitParamsInputScrapeDataScrapeAnythingScreenshotParamsAreaElement   `json:",omitzero,inline"`
	OfRectangle *BatchSubmitParamsInputScrapeDataScrapeAnythingScreenshotParamsAreaRectangle `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputScrapeDataScrapeAnythingScreenshotParamsAreaUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPage, u.OfElement, u.OfRectangle)
}
func (u *BatchSubmitParamsInputScrapeDataScrapeAnythingScreenshotParamsAreaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type BatchSubmitParamsInputScrapeDataScrapeAnythingScreenshotParamsAreaPage string

const (
	BatchSubmitParamsInputScrapeDataScrapeAnythingScreenshotParamsAreaPageViewport BatchSubmitParamsInputScrapeDataScrapeAnythingScreenshotParamsAreaPage = "viewport"
	BatchSubmitParamsInputScrapeDataScrapeAnythingScreenshotParamsAreaPageFullPage BatchSubmitParamsInputScrapeDataScrapeAnythingScreenshotParamsAreaPage = "fullPage"
)

// The property Selector is required.
type BatchSubmitParamsInputScrapeDataScrapeAnythingScreenshotParamsAreaElement struct {
	// CSS selector matching exactly one visible element.
	Selector string `json:"selector" api:"required"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataScrapeAnythingScreenshotParamsAreaElement) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataScrapeAnythingScreenshotParamsAreaElement
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataScrapeAnythingScreenshotParamsAreaElement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rectangle to capture, in pixels from the document origin.
//
// The properties Height, Width, X, Y are required.
type BatchSubmitParamsInputScrapeDataScrapeAnythingScreenshotParamsAreaRectangle struct {
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

func (r BatchSubmitParamsInputScrapeDataScrapeAnythingScreenshotParamsAreaRectangle) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataScrapeAnythingScreenshotParamsAreaRectangle
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataScrapeAnythingScreenshotParamsAreaRectangle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Browser and content settings shared by all outputs.
type BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParams struct {
	// Country to fetch from, as an ISO 3166-1 alpha-2 code such as `US`.
	// Case-insensitive.
	Country param.Opt[string] `json:"country,omitzero"`
	// Accept cookie banners before actions and capture.
	DismissCookies param.Opt[bool] `json:"dismissCookies,omitzero"`
	// Close other popups before actions and capture.
	DismissPopups param.Opt[bool] `json:"dismissPopups,omitzero"`
	// Keep only the main content, without navigation, headers, footers, or sidebars.
	// Does not affect `screenshot`, `bytes`, or `product`.
	MainContentOnly param.Opt[bool] `json:"mainContentOnly,omitzero"`
	// Wait for CSS animations to finish before capture. Defaults to `true` when
	// `formats.screenshot` is `true`.
	SettleAnimations param.Opt[bool] `json:"settleAnimations,omitzero"`
	// Browser steps to run in order before capture. Requires a paid plan. Skips the
	// cache.
	Actions []BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionUnion `json:"actions,omitzero"`
	// Remove elements matching these CSS selectors. Overrides `includeSelectors`.
	ExcludeSelectors []string `json:"excludeSelectors,omitzero"`
	// HTTP headers to send to the target site. Requests with headers skip the cache.
	Headers map[string]string `json:"headers,omitzero"`
	// Which iframes to inline in HTML and text outputs. Screenshots always show
	// visible frames. `auto` inlines only iframes that hold page content, such as
	// same-site documents, framesets, job boards, forms, and embedded documents, and
	// skips ads and widgets. `true` inlines every iframe; `false` inlines none.
	IncludeFrames BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsIncludeFramesUnion `json:"includeFrames,omitzero"`
	// Keep only elements matching these CSS selectors.
	IncludeSelectors []string `json:"includeSelectors,omitzero"`
	// Options for parsing documents such as PDFs.
	Parsers BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsParsers `json:"parsers,omitzero"`
	// Scrape as a signed-in user. Create a profile with `type: "new"` and your
	// credentials, then reuse it with `type: "existing"`. Requires a paid plan and an
	// HTTPS URL. Profiles keep their session between requests, skip the cache, and
	// cannot be combined with ZDR, custom headers, `formats.bytes`, or prebuilt
	// parsers other than `emails`, `phone-numbers`, and `links`.
	Profile BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsProfileUnion `json:"profile,omitzero"`
	// Emulate a light or dark color scheme.
	//
	// Any of "light", "dark".
	Theme string `json:"theme,omitzero"`
	// Browser viewport in pixels. Defaults to 1920 × 1080. If you pass `viewport`, a
	// missing `width` defaults to 1440 and a missing `height` to 900.
	Viewport BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsViewport `json:"viewport,omitzero"`
	// Time to wait after actions, in milliseconds, or a CSS selector to wait for.
	// Defaults to 500 ms, or 2000 ms when `includeFrames` is `true` or the URL ends in
	// `.xml`.
	WaitFor BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsWaitForUnion `json:"waitFor,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParams](
		"theme", "light", "dark",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionUnion struct {
	OfPerform *BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionPerform `json:",omitzero,inline"`
	OfScroll  *BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionScroll  `json:",omitzero,inline"`
	OfWait    *BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionWait    `json:",omitzero,inline"`
	OfWaitFor *BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionWaitFor `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPerform, u.OfScroll, u.OfWait, u.OfWaitFor)
}
func (u *BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionUnion](
		"type",
		apijson.Discriminator[BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionPerform]("perform"),
		apijson.Discriminator[BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionScroll]("scroll"),
		apijson.Discriminator[BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionWait]("wait"),
		apijson.Discriminator[BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionWaitFor]("waitFor"),
	)
}

// The properties Action, Type are required.
type BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionPerform struct {
	// One browser instruction, such as clicking a button or entering text.
	Action string `json:"action" api:"required"`
	// Follow a plain-language instruction.
	//
	// This field can be elided, and will marshal its zero value as "perform".
	Type constant.Perform `json:"type" default:"perform"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionPerform) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionPerform
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionPerform) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionScroll struct {
	// Maximum number of scroll steps for this action.
	MaxScrolls param.Opt[int64] `json:"maxScrolls,omitzero"`
	// CSS selector of the container to scroll. Omit to scroll the page.
	Selector param.Opt[string] `json:"selector,omitzero"`
	// Distance per scroll step: a number of pixels, one `viewport`, or `max` to reach
	// the end.
	Amount BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionScrollAmountUnion `json:"amount,omitzero"`
	// Direction to scroll.
	//
	// Any of "down", "up", "left", "right".
	Direction string `json:"direction,omitzero"`
	// Scroll the page or a container.
	//
	// This field can be elided, and will marshal its zero value as "scroll".
	Type constant.Scroll `json:"type" default:"scroll"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionScroll) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionScroll
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionScroll) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionScroll](
		"direction", "down", "up", "left", "right",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionScrollAmountUnion struct {
	OfInt param.Opt[int64] `json:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfBatchSubmitsInputScrapeDataScrapeAnythingSharedParamsActionScrollAmountString)
	OfBatchSubmitsInputScrapeDataScrapeAnythingSharedParamsActionScrollAmountString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionScrollAmountUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfBatchSubmitsInputScrapeDataScrapeAnythingSharedParamsActionScrollAmountString)
}
func (u *BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionScrollAmountUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionScrollAmountString string

const (
	BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionScrollAmountStringViewport BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionScrollAmountString = "viewport"
	BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionScrollAmountStringMax      BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionScrollAmountString = "max"
)

// The properties Milliseconds, Type are required.
type BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionWait struct {
	// How long to pause before the next action, in milliseconds.
	Milliseconds int64 `json:"milliseconds" api:"required"`
	// Pause for a fixed time.
	//
	// This field can be elided, and will marshal its zero value as "wait".
	Type constant.Wait `json:"type" default:"wait"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionWait) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionWait
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionWait) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Selector, Type are required.
type BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionWaitFor struct {
	// CSS selector to wait for before continuing.
	Selector string `json:"selector" api:"required"`
	// Wait for an element matching `selector`.
	//
	// This field can be elided, and will marshal its zero value as "waitFor".
	Type constant.WaitFor `json:"type" default:"waitFor"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionWaitFor) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionWaitFor
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsActionWaitFor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsIncludeFramesUnion struct {
	OfBool param.Opt[bool] `json:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfBatchSubmitsInputScrapeDataScrapeAnythingSharedParamsIncludeFramesString)
	OfBatchSubmitsInputScrapeDataScrapeAnythingSharedParamsIncludeFramesString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsIncludeFramesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfBatchSubmitsInputScrapeDataScrapeAnythingSharedParamsIncludeFramesString)
}
func (u *BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsIncludeFramesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsIncludeFramesString string

const (
	BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsIncludeFramesStringAuto BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsIncludeFramesString = "auto"
)

// Options for parsing documents such as PDFs.
type BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsParsers struct {
	// Page range, OCR, and page markers for PDFs.
	Pdf BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsParsersPdf `json:"pdf,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsParsers) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsParsers
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsParsers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Page range, OCR, and page markers for PDFs.
type BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsParsersPdf struct {
	// Last page to parse. Defaults to the last page. Must be at least `startPage`.
	EndPage param.Opt[int64] `json:"endPage,omitzero"`
	// Start each PDF page in the Markdown with a `<!-- page N -->` line, where `N` is
	// the original 1-based page number. These responses are never served from cache.
	IncludeNumberedPages param.Opt[bool] `json:"includeNumberedPages,omitzero"`
	// First page to parse, starting at 1.
	StartPage param.Opt[int64] `json:"startPage,omitzero"`
	// Set to `auto` to read scanned pages with OCR. Adds 1 credit per page read.
	//
	// Any of "off", "auto".
	Ocr string `json:"ocr,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsParsersPdf) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsParsersPdf
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsParsersPdf) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsParsersPdf](
		"ocr", "off", "auto",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsProfileUnion struct {
	OfExisting *BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsProfileExisting `json:",omitzero,inline"`
	OfNew      *BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsProfileNew      `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsProfileUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExisting, u.OfNew)
}
func (u *BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsProfileUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsProfileUnion](
		"type",
		apijson.Discriminator[BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsProfileExisting]("existing"),
		apijson.Discriminator[BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsProfileNew]("new"),
	)
}

// The properties ProfileID, Type are required.
type BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsProfileExisting struct {
	// ID of the browser profile, from `profile.id` in an earlier scrape response.
	ProfileID string `json:"profileId" api:"required"`
	// Reuse a saved browser profile.
	//
	// This field can be elided, and will marshal its zero value as "existing".
	Type constant.Existing `json:"type" default:"existing"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsProfileExisting) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsProfileExisting
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsProfileExisting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Identifier, Password, Type are required.
type BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsProfileNew struct {
	// Username or email address for the website.
	Identifier string `json:"identifier" api:"required"`
	// Password for the website. Used to sign in and never stored.
	Password string `json:"password" api:"required"`
	// Label for the profile, shown by `GET /org/browser/profiles`.
	Description param.Opt[string] `json:"description,omitzero"`
	// Sign in and save a new browser profile.
	//
	// This field can be elided, and will marshal its zero value as "new".
	Type constant.New `json:"type" default:"new"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsProfileNew) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsProfileNew
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsProfileNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Browser viewport in pixels. Defaults to 1920 × 1080. If you pass `viewport`, a
// missing `width` defaults to 1440 and a missing `height` to 900.
type BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsViewport struct {
	// Browser viewport height in pixels.
	Height param.Opt[int64] `json:"height,omitzero"`
	// Browser viewport width in pixels.
	Width param.Opt[int64] `json:"width,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsViewport) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsViewport
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsViewport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsWaitForUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsWaitForUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *BatchSubmitParamsInputScrapeDataScrapeAnythingSharedParamsWaitForUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Time limit for each page, with the same behavior and defaults as in
// `POST /web/scrape`.
//
// The property Milliseconds is required.
type BatchSubmitParamsInputScrapeDataScrapeAnythingTimeoutOpts struct {
	// Deadline in milliseconds.
	Milliseconds int64 `json:"milliseconds" api:"required"`
	// `fail` returns 408 at the deadline. `return-partial` returns the results
	// collected so far and marks the response as partial. `return-partial` requires at
	// least 5000 ms. On `POST /web/scrape`, an output that is not ready in time fails
	// with `REQUEST_TIMEOUT` inside a 200 response under either behavior; 408 means
	// the whole request could not respond by the deadline.
	//
	// Any of "fail", "return-partial".
	Behavior string `json:"behavior,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataScrapeAnythingTimeoutOpts) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataScrapeAnythingTimeoutOpts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataScrapeAnythingTimeoutOpts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputScrapeDataScrapeAnythingTimeoutOpts](
		"behavior", "fail", "return-partial",
	)
}

// Scrape the listed pages as Markdown.
//
// Deprecated: deprecated
//
// The properties Format, URLs are required.
type BatchSubmitParamsInputScrapeDataMarkdown struct {
	// Return page content as Markdown.
	//
	// Any of "markdown".
	Format string `json:"format,omitzero" api:"required"`
	// Pages to scrape, up to 100,000.
	URLs []BatchSubmitParamsInputScrapeDataMarkdownURL `json:"urls,omitzero" api:"required"`
	// Options for Markdown output.
	Options BatchSubmitParamsInputScrapeDataMarkdownOptions `json:"options,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataMarkdown) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataMarkdown
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataMarkdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputScrapeDataMarkdown](
		"format", "markdown",
	)
}

// A page to scrape, with optional `itemId` and `meta` to match it to its result.
//
// The property URL is required.
type BatchSubmitParamsInputScrapeDataMarkdownURL struct {
	// URL of the page.
	URL string `json:"url" api:"required"`
	// Your ID for this page, returned with its result. The same URL can use different
	// IDs.
	ItemID param.Opt[string] `json:"itemId,omitzero"`
	// Custom JSON returned unchanged with this page's result.
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
	// Maximum cache age in milliseconds. Defaults to 3 days (259200000 ms). Maximum: 1
	// year (31536000000 ms). `0` fetches fresh.
	MaxAgeMs param.Opt[int64] `json:"maxAgeMs,omitzero"`
	// Also return each page's HTML in `html`.
	IncludeHTML param.Opt[bool] `json:"includeHTML,omitzero"`
	// Include image references in the Markdown (`html-conversion` only).
	IncludeImages param.Opt[bool] `json:"includeImages,omitzero"`
	// Include links in the Markdown (`html-conversion` only).
	IncludeLinks param.Opt[bool] `json:"includeLinks,omitzero"`
	// Wait for CSS animations to finish before extracting, on browser-rendered pages.
	SettleAnimations param.Opt[bool] `json:"settleAnimations,omitzero"`
	// Shorten inline base64 image data (`html-conversion` only).
	ShortenBase64Images param.Opt[bool] `json:"shortenBase64Images,omitzero"`
	// Return the main content without navigation or footers.
	UseMainContentOnly param.Opt[bool] `json:"useMainContentOnly,omitzero"`
	// How long to wait after initial page load, in milliseconds. `0` waits 500 ms.
	WaitForMs param.Opt[int64] `json:"waitForMs,omitzero"`
	// Remove elements matching these CSS selectors. Applied after `includeSelectors`,
	// so an element matching both is removed.
	ExcludeSelectors []string `json:"excludeSelectors,omitzero"`
	// Keep only elements matching these CSS selectors. Filtered pages ignore
	// `maxAgeMs`.
	IncludeSelectors []string `json:"includeSelectors,omitzero"`
	// Country to fetch from, as an ISO 3166-1 alpha-2 code.
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
	// Whether to parse PDFs and which pages to read.
	Pdf BatchSubmitParamsInputScrapeDataMarkdownOptionsPdf `json:"pdf,omitzero"`
	// Markdown sources in order of preference: `html-conversion` converts the page
	// HTML, `markdown-direct` uses Markdown the site publishes.
	//
	// Any of "html-conversion", "markdown-direct".
	SourcePreference []string `json:"sourcePreference,omitzero"`
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

// Whether to parse PDFs and which pages to read.
type BatchSubmitParamsInputScrapeDataMarkdownOptionsPdf struct {
	// Last page to parse. Defaults to the last page. Must be at least `start`.
	End param.Opt[int64] `json:"end,omitzero"`
	// Read scanned PDF pages with OCR. Pages that already have text are kept as is.
	Ocr param.Opt[bool] `json:"ocr,omitzero"`
	// Parse PDF URLs. When false, PDFs fail with `PDF_SKIPPED`.
	ShouldParse param.Opt[bool] `json:"shouldParse,omitzero"`
	// First page to parse, starting at 1.
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
// Deprecated: deprecated
//
// The properties Format, URLs are required.
type BatchSubmitParamsInputScrapeDataHTML struct {
	// Return page content as HTML.
	//
	// Any of "html".
	Format string `json:"format,omitzero" api:"required"`
	// Pages to scrape, up to 100,000.
	URLs []BatchSubmitParamsInputScrapeDataHTMLURL `json:"urls,omitzero" api:"required"`
	// Options for HTML output.
	Options BatchSubmitParamsInputScrapeDataHTMLOptions `json:"options,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputScrapeDataHTML) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputScrapeDataHTML
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputScrapeDataHTML) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputScrapeDataHTML](
		"format", "html",
	)
}

// A page to scrape, with optional `itemId` and `meta` to match it to its result.
//
// The property URL is required.
type BatchSubmitParamsInputScrapeDataHTMLURL struct {
	// URL of the page.
	URL string `json:"url" api:"required"`
	// Your ID for this page, returned with its result. The same URL can use different
	// IDs.
	ItemID param.Opt[string] `json:"itemId,omitzero"`
	// Custom JSON returned unchanged with this page's result.
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
	// Maximum cache age in milliseconds. Defaults to 3 days (259200000 ms). Maximum: 1
	// year (31536000000 ms). `0` fetches fresh.
	MaxAgeMs param.Opt[int64] `json:"maxAgeMs,omitzero"`
	// Wait for CSS animations to finish before extracting, on browser-rendered pages.
	SettleAnimations param.Opt[bool] `json:"settleAnimations,omitzero"`
	// Return the main content without navigation or footers.
	UseMainContentOnly param.Opt[bool] `json:"useMainContentOnly,omitzero"`
	// How long to wait after initial page load, in milliseconds. `0` waits 500 ms.
	WaitForMs param.Opt[int64] `json:"waitForMs,omitzero"`
	// Remove elements matching these CSS selectors. Applied after `includeSelectors`,
	// so an element matching both is removed.
	ExcludeSelectors []string `json:"excludeSelectors,omitzero"`
	// Keep only elements matching these CSS selectors. Filtered pages ignore
	// `maxAgeMs`.
	IncludeSelectors []string `json:"includeSelectors,omitzero"`
	// Country to fetch from, as an ISO 3166-1 alpha-2 code.
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
	// Whether to parse PDFs and which pages to read.
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

// Whether to parse PDFs and which pages to read.
type BatchSubmitParamsInputScrapeDataHTMLOptionsPdf struct {
	// Last page to parse. Defaults to the last page. Must be at least `start`.
	End param.Opt[int64] `json:"end,omitzero"`
	// Read scanned PDF pages with OCR. Pages that already have text are kept as is.
	Ocr param.Opt[bool] `json:"ocr,omitzero"`
	// Parse PDF URLs. When false, PDFs fail with `PDF_SKIPPED`.
	ShouldParse param.Opt[bool] `json:"shouldParse,omitzero"`
	// First page to parse, starting at 1.
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
	// Where to find pages, with the formats and options to apply to each page.
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
	OfCrawlWithScrapeAnything *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnything `json:",omitzero,inline"`
	OfMarkdown                *BatchSubmitParamsInputCrawlDataMarkdown                `json:",omitzero,inline"`
	OfHTML                    *BatchSubmitParamsInputCrawlDataHTML                    `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlWithScrapeAnything, u.OfMarkdown, u.OfHTML)
}
func (u *BatchSubmitParamsInputCrawlDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The same formats and options as `POST /web/scrape`, with `source` in place of
// `url`. Options apply to every page.
//
// The properties Formats, Source are required.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnything struct {
	// Outputs to return. Set at least one to `true`.
	Formats BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingFormats `json:"formats,omitzero" api:"required"`
	// How to find pages to crawl.
	Source BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceUnion `json:"source,omitzero" api:"required"`
	// Maximum age of a cached result, in milliseconds. `0` always fetches fresh.
	// Defaults to 3 days (259200000 ms). Maximum: 1 year (31536000000 ms).
	MaxAgeMs param.Opt[int64] `json:"maxAgeMs,omitzero"`
	// Query for the `highlights` output. Required when `formats.highlights` is `true`,
	// and allowed only then.
	HighlightsParams BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingHighlightsParams `json:"highlightsParams,omitzero"`
	// Options for the `images` output. Requires `formats.images: true`.
	ImageParams BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingImageParams `json:"imageParams,omitzero"`
	// Schema and instructions for the `json` output. Required when `formats.json` is
	// `true`, and allowed only then.
	JsonParams BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingJsonParams `json:"jsonParams,omitzero"`
	// Options for the `markdown` output. Requires `formats.markdown: true`.
	MarkdownParams BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingMarkdownParams `json:"markdownParams,omitzero"`
	// Custom rules or a prebuilt parser for the `parsed` output. Required when
	// `formats.parse` is `true`, and allowed only then.
	ParseParams BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingParseParamsUnion `json:"parseParams,omitzero"`
	// Options for the `product` output. Requires `formats.product: true`.
	ProductParams BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingProductParams `json:"productParams,omitzero"`
	// Options for the `screenshot` output. Requires `formats.screenshot: true`.
	ScreenshotParams BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingScreenshotParams `json:"screenshotParams,omitzero"`
	// Browser and content settings shared by all outputs.
	SharedParams BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParams `json:"sharedParams,omitzero"`
	// Labels for this request. Filter usage and request logs by them in the dashboard
	// or with `GET /org/logs`.
	Tags []string `json:"tags,omitzero"`
	// Time limit for each page, with the same behavior and defaults as in
	// `POST /web/scrape`.
	TimeoutOpts BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingTimeoutOpts `json:"timeoutOpts,omitzero"`
	// Zero data retention is not available for batches, because batches store their
	// results. Sending `enabled` returns `ZDR_NOT_SUPPORTED`.
	//
	// Any of "disabled".
	Zdr string `json:"zdr,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnything) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnything
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnything) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnything](
		"zdr", "disabled",
	)
}

// Outputs to return. Set at least one to `true`.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingFormats struct {
	// Return the original HTTP response body.
	Bytes param.Opt[bool] `json:"bytes,omitzero"`
	// Return Markdown excerpts relevant to `highlightsParams.query`.
	Highlights param.Opt[bool] `json:"highlights,omitzero"`
	// Return the rendered HTML.
	HTML param.Opt[bool] `json:"html,omitzero"`
	// Return the images found on the page.
	Images param.Opt[bool] `json:"images,omitzero"`
	// Return an object matching `jsonParams.schema`, extracted from the page text and
	// embedded data such as JSON-LD, meta tags, and framework state.
	Json param.Opt[bool] `json:"json,omitzero"`
	// Return the page content, or a video transcript, as Markdown.
	Markdown param.Opt[bool] `json:"markdown,omitzero"`
	// Return fields extracted with custom rules or a prebuilt parser, as `parsed`.
	Parse param.Opt[bool] `json:"parse,omitzero"`
	// Return product details such as name, price, and availability.
	Product param.Opt[bool] `json:"product,omitzero"`
	// Return a screenshot of the page.
	Screenshot param.Opt[bool] `json:"screenshot,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingFormats) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingFormats
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingFormats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceUnion struct {
	OfStartURL *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURL `json:",omitzero,inline"`
	OfSitemap  *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemap  `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfStartURL, u.OfSitemap)
}
func (u *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceUnion](
		"type",
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURL]("start_url"),
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemap]("sitemap"),
	)
}

// Discover pages by following links from one URL.
//
// The properties Type, URL are required.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURL struct {
	// Page where crawling begins. A URL without a scheme is treated as `https://`.
	URL string `json:"url" api:"required"`
	// Limits and filters for page discovery.
	Controls BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControls `json:"controls,omitzero"`
	// Follow links from one URL.
	//
	// This field can be elided, and will marshal its zero value as "start_url".
	Type constant.StartURL `json:"type" default:"start_url"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURL) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Limits and filters for page discovery.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControls struct {
	// Follow links to subdomains.
	FollowSubdomains param.Opt[bool] `json:"followSubdomains,omitzero"`
	// Maximum link depth, where the start page is depth `0`. No limit when omitted.
	MaxDepth param.Opt[int64] `json:"maxDepth,omitzero"`
	// Maximum number of pages to fetch, up to 100,000. Credits are reserved for this
	// many pages and the unused part is refunded.
	MaxURLs param.Opt[int64] `json:"maxUrls,omitzero"`
	// Pages to skip entirely, even when they match `follow` or `scrape`.
	Exclude []BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsExcludeUnion `json:"exclude,omitzero"`
	// Links to follow from each page; every followed page is checked for links again.
	// Follows all links when omitted.
	Follow []BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsFollowUnion `json:"follow,omitzero"`
	// Only follow pages in these languages, or pass `same_as_start_url` to match the
	// start page. Follows all languages when omitted.
	Languages BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsLanguagesUnion `json:"languages,omitzero"`
	// Pages to return in the results. Returns every fetched page when omitted.
	Scrape []BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsScrapeUnion `json:"scrape,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControls) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControls
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControls) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsExcludeUnion struct {
	OfRegex   *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsExcludeRegex   `json:",omitzero,inline"`
	OfProduct *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsExcludeProduct `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsExcludeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfRegex, u.OfProduct)
}
func (u *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsExcludeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsExcludeUnion](
		"kind",
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsExcludeRegex]("regex"),
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsExcludeProduct]("product"),
	)
}

// Pages whose URL matches a regex pattern.
//
// The properties Kind, Pattern are required.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsExcludeRegex struct {
	// Regular expression matched against the page URL.
	Pattern string `json:"pattern" api:"required"`
	// Match pages by URL pattern.
	//
	// This field can be elided, and will marshal its zero value as "regex".
	Kind constant.Regex `json:"kind" default:"regex"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsExcludeRegex) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsExcludeRegex
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsExcludeRegex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewBatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsExcludeProduct() BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsExcludeProduct {
	return BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsExcludeProduct{
		Kind: "product",
	}
}

// Product detail pages.
//
// This struct has a constant value, construct it with
// [NewBatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsExcludeProduct].
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsExcludeProduct struct {
	// Match product detail pages.
	Kind constant.Product `json:"kind" default:"product"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsExcludeProduct) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsExcludeProduct
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsExcludeProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsFollowUnion struct {
	OfRegex      *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsFollowRegex      `json:",omitzero,inline"`
	OfPagination *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsFollowPagination `json:",omitzero,inline"`
	OfProduct    *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsFollowProduct    `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsFollowUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfRegex, u.OfPagination, u.OfProduct)
}
func (u *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsFollowUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsFollowUnion](
		"kind",
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsFollowRegex]("regex"),
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsFollowPagination]("pagination"),
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsFollowProduct]("product"),
	)
}

// Pages whose URL matches a regex pattern.
//
// The properties Kind, Pattern are required.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsFollowRegex struct {
	// Regular expression matched against the page URL.
	Pattern string `json:"pattern" api:"required"`
	// Match pages by URL pattern.
	//
	// This field can be elided, and will marshal its zero value as "regex".
	Kind constant.Regex `json:"kind" default:"regex"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsFollowRegex) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsFollowRegex
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsFollowRegex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewBatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsFollowPagination() BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsFollowPagination {
	return BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsFollowPagination{
		Kind: "pagination",
	}
}

// Further pages of a paginated list, such as page 2 of a catalog.
//
// This struct has a constant value, construct it with
// [NewBatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsFollowPagination].
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsFollowPagination struct {
	// Match the further pages of a paginated list.
	Kind constant.Pagination `json:"kind" default:"pagination"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsFollowPagination) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsFollowPagination
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsFollowPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewBatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsFollowProduct() BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsFollowProduct {
	return BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsFollowProduct{
		Kind: "product",
	}
}

// Product detail pages.
//
// This struct has a constant value, construct it with
// [NewBatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsFollowProduct].
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsFollowProduct struct {
	// Match product detail pages.
	Kind constant.Product `json:"kind" default:"product"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsFollowProduct) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsFollowProduct
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsFollowProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsLanguagesUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfBatchSubmitsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsLanguagesString)
	OfBatchSubmitsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsLanguagesString         param.Opt[string] `json:",omitzero,inline"`
	OfBatchSubmitsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsLanguagesArrayItemArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsLanguagesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBatchSubmitsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsLanguagesString, u.OfBatchSubmitsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsLanguagesArrayItemArray)
}
func (u *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsLanguagesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsLanguagesString string

const (
	BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsLanguagesStringSameAsStartURL BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsLanguagesString = "same_as_start_url"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsScrapeUnion struct {
	OfRegex   *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsScrapeRegex   `json:",omitzero,inline"`
	OfProduct *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsScrapeProduct `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsScrapeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfRegex, u.OfProduct)
}
func (u *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsScrapeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsScrapeUnion](
		"kind",
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsScrapeRegex]("regex"),
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsScrapeProduct]("product"),
	)
}

// Pages whose URL matches a regex pattern.
//
// The properties Kind, Pattern are required.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsScrapeRegex struct {
	// Regular expression matched against the page URL.
	Pattern string `json:"pattern" api:"required"`
	// Match pages by URL pattern.
	//
	// This field can be elided, and will marshal its zero value as "regex".
	Kind constant.Regex `json:"kind" default:"regex"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsScrapeRegex) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsScrapeRegex
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsScrapeRegex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewBatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsScrapeProduct() BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsScrapeProduct {
	return BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsScrapeProduct{
		Kind: "product",
	}
}

// Product detail pages.
//
// This struct has a constant value, construct it with
// [NewBatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsScrapeProduct].
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsScrapeProduct struct {
	// Match product detail pages.
	Kind constant.Product `json:"kind" default:"product"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsScrapeProduct) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsScrapeProduct
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceStartURLControlsScrapeProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Scrape the pages listed in a domain's sitemap. Links on those pages are not
// followed.
//
// The properties Domain, Type are required.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemap struct {
	// Domain whose sitemap lists the pages to scrape. A full URL is reduced to its
	// domain.
	Domain string `json:"domain" api:"required"`
	// Limits and filters for the sitemap URLs. Links on sitemap pages are never
	// followed.
	Controls BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControls `json:"controls,omitzero"`
	// Scrape the URLs in the domain's sitemap.
	//
	// This field can be elided, and will marshal its zero value as "sitemap".
	Type constant.Sitemap `json:"type" default:"sitemap"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemap) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemap
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Limits and filters for the sitemap URLs. Links on sitemap pages are never
// followed.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControls struct {
	// Maximum number of pages to fetch, up to 100,000. Credits are reserved for this
	// many pages and the unused part is refunded.
	MaxURLs param.Opt[int64] `json:"maxUrls,omitzero"`
	// Sitemap pages to skip, even when they match `scrape`.
	Exclude []BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsExcludeUnion `json:"exclude,omitzero"`
	// Path fragments that exclude a sitemap URL. Applies in addition to `scrape` and
	// `exclude`; pass `[]` to keep every listed URL.
	ExcludePaths []string `json:"excludePaths,omitzero"`
	// Sitemap pages to return in the results. Returns every listed page when omitted.
	Scrape []BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsScrapeUnion `json:"scrape,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControls) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControls
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControls) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsExcludeUnion struct {
	OfRegex   *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsExcludeRegex   `json:",omitzero,inline"`
	OfProduct *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsExcludeProduct `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsExcludeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfRegex, u.OfProduct)
}
func (u *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsExcludeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsExcludeUnion](
		"kind",
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsExcludeRegex]("regex"),
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsExcludeProduct]("product"),
	)
}

// Pages whose URL matches a regex pattern.
//
// The properties Kind, Pattern are required.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsExcludeRegex struct {
	// Regular expression matched against the page URL.
	Pattern string `json:"pattern" api:"required"`
	// Match pages by URL pattern.
	//
	// This field can be elided, and will marshal its zero value as "regex".
	Kind constant.Regex `json:"kind" default:"regex"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsExcludeRegex) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsExcludeRegex
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsExcludeRegex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewBatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsExcludeProduct() BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsExcludeProduct {
	return BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsExcludeProduct{
		Kind: "product",
	}
}

// Product detail pages.
//
// This struct has a constant value, construct it with
// [NewBatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsExcludeProduct].
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsExcludeProduct struct {
	// Match product detail pages.
	Kind constant.Product `json:"kind" default:"product"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsExcludeProduct) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsExcludeProduct
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsExcludeProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsScrapeUnion struct {
	OfRegex   *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsScrapeRegex   `json:",omitzero,inline"`
	OfProduct *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsScrapeProduct `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsScrapeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfRegex, u.OfProduct)
}
func (u *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsScrapeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsScrapeUnion](
		"kind",
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsScrapeRegex]("regex"),
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsScrapeProduct]("product"),
	)
}

// Pages whose URL matches a regex pattern.
//
// The properties Kind, Pattern are required.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsScrapeRegex struct {
	// Regular expression matched against the page URL.
	Pattern string `json:"pattern" api:"required"`
	// Match pages by URL pattern.
	//
	// This field can be elided, and will marshal its zero value as "regex".
	Kind constant.Regex `json:"kind" default:"regex"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsScrapeRegex) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsScrapeRegex
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsScrapeRegex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewBatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsScrapeProduct() BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsScrapeProduct {
	return BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsScrapeProduct{
		Kind: "product",
	}
}

// Product detail pages.
//
// This struct has a constant value, construct it with
// [NewBatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsScrapeProduct].
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsScrapeProduct struct {
	// Match product detail pages.
	Kind constant.Product `json:"kind" default:"product"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsScrapeProduct) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsScrapeProduct
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSourceSitemapControlsScrapeProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Query for the `highlights` output. Required when `formats.highlights` is `true`,
// and allowed only then.
//
// The property Query is required.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingHighlightsParams struct {
	// Question or topic to find relevant passages for.
	Query string `json:"query" api:"required"`
	// Maximum combined length of the returned passages, in characters.
	MaxCharacters param.Opt[int64] `json:"maxCharacters,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingHighlightsParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingHighlightsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingHighlightsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Options for the `images` output. Requires `formats.images: true`.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingImageParams struct {
	// Set to `visual` to drop visually duplicate images, keeping the largest copy.
	//
	// Any of "none", "visual".
	Dedupe string `json:"dedupe,omitzero"`
	// Extra data to add to each image: `dimensions` adds `width` and `height`,
	// `classification` adds `classification`, and `file` adds a hosted `fileUrl`.
	//
	// Any of "dimensions", "classification", "file".
	Enrich []string `json:"enrich,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingImageParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingImageParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingImageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingImageParams](
		"dedupe", "none", "visual",
	)
}

// Schema and instructions for the `json` output. Required when `formats.json` is
// `true`, and allowed only then.
//
// The property Schema is required.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingJsonParams struct {
	// JSON Schema for the object to extract, not an example object. Must describe a
	// top-level object, up to 50 KB. Make fields optional or nullable when the page
	// may not state them.
	Schema map[string]any `json:"schema,omitzero" api:"required"`
	// Extra guidance, such as which facts to prefer or how to read a field.
	Instructions param.Opt[string] `json:"instructions,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingJsonParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingJsonParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingJsonParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Options for the `markdown` output. Requires `formats.markdown: true`.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingMarkdownParams struct {
	// Include image references in the Markdown (`html-conversion` only).
	IncludeImages param.Opt[bool] `json:"includeImages,omitzero"`
	// Keep link URLs in the Markdown (`html-conversion` only).
	IncludeLinks param.Opt[bool] `json:"includeLinks,omitzero"`
	// How to write base64 images when `includeImages` is `true`: `placeholder`
	// shortens the data, `preserve` keeps it in full (`html-conversion` only).
	//
	// Any of "placeholder", "preserve".
	InlineImages string `json:"inlineImages,omitzero"`
	// Markdown sources in order of preference: `html-conversion` converts the page
	// HTML, `markdown-direct` uses Markdown the site publishes.
	//
	// Any of "html-conversion", "markdown-direct".
	SourcePreference []string `json:"sourcePreference,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingMarkdownParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingMarkdownParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingMarkdownParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingMarkdownParams](
		"inlineImages", "placeholder", "preserve",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingParseParamsUnion struct {
	OfCustomParser   *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingParseParamsCustomParser   `json:",omitzero,inline"`
	OfPrebuiltParser *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingParseParamsPrebuiltParser `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingParseParamsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCustomParser, u.OfPrebuiltParser)
}
func (u *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingParseParamsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property Rules is required.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingParseParamsCustomParser struct {
	// Output field names mapped to a CSS selector (such as `h1` or `a@href`) or a rule
	// object. Up to 100 fields, nested up to 5 levels.
	Rules map[string]BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingParseParamsCustomParserRuleUnion `json:"rules,omitzero" api:"required"`
	// Extract fields with custom CSS rules. The default when `type` is omitted.
	//
	// Any of "custom".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingParseParamsCustomParser) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingParseParamsCustomParser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingParseParamsCustomParser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingParseParamsCustomParser](
		"type", "custom",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingParseParamsCustomParserRuleUnion struct {
	OfString                                                                             param.Opt[string]                                                                        `json:",omitzero,inline"`
	OfBatchSubmitsInputCrawlDataCrawlWithScrapeAnythingParseParamsCustomParserRuleObject *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingParseParamsCustomParserRuleObject `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingParseParamsCustomParserRuleUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfBatchSubmitsInputCrawlDataCrawlWithScrapeAnythingParseParamsCustomParserRuleObject)
}
func (u *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingParseParamsCustomParserRuleUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property Selector is required.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingParseParamsCustomParserRuleObject struct {
	// CSS selector to match within the current page or parent rule.
	Selector string `json:"selector" api:"required"`
	// What to return: `text`, `html`, an attribute such as `@href`, or nested field
	// rules. Defaults to `text`.
	Output string `json:"output,omitzero"`
	// Return the first match with `item` or all matches with `list`.
	//
	// Any of "item", "list".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingParseParamsCustomParserRuleObject) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingParseParamsCustomParserRuleObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingParseParamsCustomParserRuleObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingParseParamsCustomParserRuleObject](
		"type", "item", "list",
	)
}

// The properties Parser, Type are required.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingParseParamsPrebuiltParser struct {
	// Prebuilt parser to run. Each one accepts specific URLs and returns its result in
	// `parsed.data`. `emails`, `phone-numbers`, and `links` read the rendered page and
	// respect content filters; the others read their source directly, independently of
	// content filters and other outputs.
	//
	//   - `google-maps`: Google Maps place, search, and internal data URLs. Returns
	//     businesses in `data.results`; a search URL returns up to 200 places. Results
	//     use the country set in the URL, or the US when none is set, and come back with
	//     English labels.
	//   - `google-maps-reviews`: Google Maps place, search, and `?cid=` URLs, plus
	//     `search.google.com/local/reviews` URLs, so the `google_maps_url` and
	//     `google_reviews_url` fields of `google-maps` results work as input. Returns
	//     the reviews of the place in `data.reviews` (`rating`, `date`, `isoDate`,
	//     `snippet`, `translatedSnippet`, `likes`, `user`, `media`, `response`, `link`,
	//     `id`, `language`, guided ratings in `details`, and `user.isLocalGuide`), plus
	//     the place `featureId`, `cid`, and `sort` in `data`. `sort` and `limit` set the
	//     order and count; `oldest` is exact. An unknown place returns no reviews.
	//   - `google-play`: Play Store app details URLs
	//     (`play.google.com/store/apps/details?id=…`). Returns one app in `data`:
	//     details, developer contact, ratings, price, permissions, data safety, the
	//     storefront country the record describes, and `is_available` (whether the app
	//     is offered there). Ratings are counted per storefront, so `rating`,
	//     `rating_count`, and `rating_histogram` are `null` where nobody rated the app.
	//     The language and storefront country set in the URL apply (default English and
	//     the US); without a storefront country, an app not offered in the US is read
	//     from the country of its developer.
	//   - `google-play-reviews`: Same URLs as `google-play`. Returns up to the 200
	//     newest reviews in the language set in the URL, with developer replies, in
	//     `data.reviews`; an app with fewer such reviews returns them all.
	//   - `app-store`: App Store app URLs (`apps.apple.com/us/app/…/id…`). Returns one
	//     app in `data`: details, developer, ratings, price, version, and screenshots.
	//     The country in the path sets the storefront (default `us`).
	//   - `app-store-reviews`: Same URLs as `app-store`. Returns up to the 500 newest
	//     reviews in `data.reviews`.
	//   - `g2-reviews`: G2 product URLs (`g2.com/products/…`), following renamed
	//     products. Returns the 25 most recently published reviews in `data.reviews`
	//     (rating, title, likes, dislikes, `problems_solved`, reviewer role, company
	//     size, and industry), plus `data.product`. Keeps star-rating filters from the
	//     URL (`filters[nps_score][]=1` to `5`) and ignores other filters, sort orders,
	//     and pages.
	//   - `instagram-profile`: Public profile URLs (`instagram.com/{username}`). Returns
	//     one profile in `data`: bio, links, follower and post counts, category,
	//     business address, `similar_accounts`, and `latest_posts` (up to the 12 posts
	//     at the top of the profile, newest first). Instagram lists pinned posts first,
	//     so posts with `is_pinned: true` can be older. When Instagram shows only a
	//     rounded post count (such as 32K), `post_count` is that rounded number and
	//     `is_post_count_rounded` is `true`.
	//   - `instagram-post`: Post and reel URLs (`/p/…`, `/reel/…`). Returns one post in
	//     `data`: caption, hashtags, mentions, like, comment, and view counts
	//     (`like_count` and `view_count` are `null` when the owner hides them), media
	//     URLs, tagged users, collaborators, location, and about the top 14 comments in
	//     `data.comments`.
	//   - `instagram-feed`: Hashtag and keyword URLs (`/explore/tags/…`, `/popular/…`)
	//     return the top posts; location URLs (`/explore/locations/{id}/…`) return
	//     recent posts. Posts are in `data.posts`; keyword posts carry only the
	//     shortcode, owner, thumbnail, and views, plus the caption when Instagram shows
	//     it. All three Instagram parsers read public data without signing in, so
	//     private accounts return no posts and follower lists are unavailable.
	//   - `tiktok-profile`: Public profile URLs (`tiktok.com/@{username}`). Returns one
	//     profile in `data`: nickname, bio, bio link, follower, following, like, and
	//     video counts, verified, private, and business flags, and business category.
	//     Does not list videos.
	//   - `tiktok-video`: Video and photo post URLs (`/@{username}/video/{id}`,
	//     `/@{username}/photo/{id}`). Returns one post in `data`: caption, hashtags,
	//     mentions, author, view, like, comment, share, save, and repost counts,
	//     duration, thumbnail, a `video_url` that downloads without cookies,
	//     `image_urls` for photo posts, music, location, and ad, paid-partnership, and
	//     AI-generated flags.
	//   - `tiktok-feed`: Returns videos in the `tiktok-video` shape in `data.videos`:
	//     the latest 15 from a profile URL, the top 30 from a hashtag URL
	//     (`/tag/{name}`), or the top 12 from a search URL (`/search?q={query}`).
	//   - `tiktok-comments`: Video and photo post URLs. Returns the first 20 comments in
	//     `data.comments`: text, `posted_at`, like and reply counts, whether the author
	//     liked it, and the commenter.
	//   - `tiktok-followers`, `tiktok-following`: Profile URLs. Return up to 30 accounts
	//     in `data.users`, in the `tiktok-profile` shape.
	//   - `tiktok-user-search`: User search URLs (`/search/user?q={query}`). Returns the
	//     top 10 accounts in `data.users`: `id`, `username`, `nickname`, `biography`,
	//     `avatar_url`, `sec_uid`, `follower_count`, `like_count`, and `is_verified`.
	//     The TikTok parsers accept `tiktok.com`, `www.tiktok.com`, and `m.tiktok.com`
	//     URLs, not short links (`vm.tiktok.com`, `vt.tiktok.com`, `tiktok.com/t/…`).
	//   - `reddit-post`: Post URLs (`reddit.com/r/{subreddit}/comments/{id}`,
	//     `reddit.com/comments/{id}`, `redd.it/{id}`). Returns the post in `data.post`:
	//     title, body as Markdown, type (`text`, `link`, `image`, `gallery`, or
	//     `video`), `link_url`, subreddit, author, flairs, score, `upvote_ratio`,
	//     comment and crosspost counts, NSFW, spoiler, locked, stickied, and archived
	//     flags, images (every gallery image in order), and video (an MP4 `url` plus an
	//     `hls_url` that carries the audio). `data.comments` holds the comment tree;
	//     each comment has `id`, `author`, `body` (Markdown), `score`, `posted_at`,
	//     `depth`, `parent_id`, `url`, `is_submitter`, `is_deleted`, and its `replies`.
	//     `limit` caps the total number of comments (default 100), filled shallowest
	//     first so whole conversations are kept, and `data.has_more_comments` says
	//     whether the thread has more. The `sort` in the URL (`best`, `top`, `new`,
	//     `controversial`, `old`, `qa`) orders the comments.
	//   - `reddit-search`: Search URLs (`reddit.com/search?q={query}`, or
	//     `reddit.com/r/{subreddit}/search?q={query}` to search one subreddit). With
	//     `type=link` or no `type`, returns posts in the `reddit-post` shape in
	//     `data.posts`, ordered by the `sort` in the URL (`relevance`, `hot`, `top`,
	//     `new`, `comments`) within its `t` time range (`hour`, `day`, `week`, `month`,
	//     `year`, `all`). `type=sr` returns communities in `data.communities` (name,
	//     title, description, sidebar, `subscriber_count`, `created_at`, type, NSFW and
	//     quarantine flags, language, category, icon and banner URLs, and allowed post
	//     types). `type=user` returns accounts in `data.users` (username,
	//     `display_name`, bio, `created_at`, post and comment karma, avatar and banner
	//     URLs, and verified, employee, moderator, premium, and NSFW flags). `limit`
	//     caps the number of results (default 100), in the order Reddit returns them,
	//     and `data.has_more` says whether Reddit has more.
	//   - `reddit-subreddit`: Subreddit URLs (`reddit.com/r/{subreddit}`, optionally
	//     followed by `/hot`, `/new`, `/top`, `/rising`, or `/controversial`, with `t`
	//     setting the time range of `top` and `controversial`, default `day`). Returns
	//     `data.subreddit` in the community shape of `reddit-search`, `data.rules`
	//     (`name`, `description`, `applies_to`, `violation_reason`, `created_at`), and
	//     `data.posts` in the `reddit-post` shape, ordered like the tab (default `hot`)
	//     with stickied posts flagged. `r/all` and `r/popular` return posts with `null`
	//     `subreddit` and `rules`. `limit` caps the posts (default 100) and
	//     `data.has_more` says whether Reddit has more. Private and quarantined
	//     subreddits return `WEBSITE_ACCESS_ERROR`.
	//   - `reddit-user`: Profile URLs (`reddit.com/user/{username}` or
	//     `reddit.com/u/{username}`, optionally followed by `/overview`, `/submitted`,
	//     or `/comments`, with `sort` set to `new`, `hot`, `top`, or `controversial` and
	//     `t` setting the time range of `top` and `controversial`, default `all`).
	//     Returns `data.user` (`id`, `username`, `display_name`, bio, `created_at`,
	//     total, post, comment, awardee, and awarder karma, avatar and banner URLs, and
	//     verified, suspended, employee, moderator, premium, and NSFW flags),
	//     `data.trophies` (`name`, `description`, `icon_url`, `granted_at`, `url`),
	//     `data.moderated_communities` (`name`, `title`, `url`, `subscriber_count`,
	//     `icon_url`, `is_nsfw`), `data.posts` in the `reddit-post` shape, and
	//     `data.comments`, each with `id`, `author`, `body` (Markdown), `score`,
	//     `posted_at`, `edited_at`, `url`, `subreddit`, `parent_id` (`null` for a
	//     top-level comment), `is_submitter`, `is_stickied`, `distinguished`,
	//     `award_count`, and the post it was left on (`id`, `title`, `url`, `author`,
	//     `is_nsfw`). The overview tab returns both posts and comments; `submitted`
	//     returns `null` comments and `comments` returns `null` posts. `limit` caps
	//     posts and comments together (default 100) and `data.has_more` says whether
	//     Reddit has more. Users who hide their activity return empty lists, and
	//     suspended accounts return their profile with `null` posts and comments.
	//   - `yelp-search`: Yelp search URLs with `find_desc` and `find_loc`; `start` sets
	//     the result offset. Returns up to 20 ranked organic businesses in
	//     `data.businesses`, with `total` and `pagination`.
	//   - `yelp-reviews`: Yelp business URLs (`yelp.com/biz/{alias}`); optional `start`
	//     sets the offset. Returns up to 20 of the latest public reviews, always newest
	//     first, in `data.reviews`, with business details and `pagination.next_offset`.
	//     Each review has rating, full text, `updated_at`, and author (`id` from the
	//     reviewer profile URL, `url` linking to that profile, `name`, `location`, and
	//     `review_count`); owner replies and photos are not included. Unsupported sort
	//     and review filters are rejected; tracking parameters are ignored.
	//   - `tripadvisor-reviews`: Tripadvisor hotel, restaurant, and attraction review
	//     URLs (`tripadvisor.com/Hotel_Review-g{geo}-d{id}-Reviews-….html`, on any
	//     Tripadvisor country domain); `-or{N}-` in the path sets the offset. Returns 20
	//     reviews in `data.reviews` in the order Tripadvisor shows them, with
	//     `data.location` (`id`, `name`, `type`, `url`, `review_count`) and
	//     `pagination.next_offset`. Each review has `id`, `url`, `title`, full text,
	//     `rating`, `language`, `published_at`, `visit_month`, `helpful_votes`,
	//     `subratings`, `photo_urls`, `owner_response` (`text`, `published_at`,
	//     `author_name`), and `author` (`id`, `username`, `name`, `url`,
	//     `contribution_count`). An unknown location returns no reviews.
	//   - `car-rental-availability`: avis.com and budget.com vehicle-availability search
	//     URLs. Returns every vehicle with pay-now and pay-later prices in
	//     `data.vehicles`.
	//   - `car-rental-locations`: avis.com and budget.com `/locations?keyword=` URLs.
	//     Returns rental branches with codes, addresses, coordinates, and hours in
	//     `data.locations`.
	//   - `emails`: Any page. Returns its email addresses in `data.emails`
	//     (`[{ email }]`), read from `mailto:` links, Cloudflare-protected addresses,
	//     and visible text.
	//   - `phone-numbers`: Any page. Returns its valid phone numbers in E.164 format in
	//     `data.phone_numbers` (`[{ phone }]`), read from `tel:` links and visible text.
	//     Numbers without a country code are read for the country of the site TLD (US by
	//     default).
	//   - `links`: Any page. Returns its unique HTTP and HTTPS links as absolute URLs in
	//     `data.links` (`[{ url, text }]`). `text` falls back to `aria-label` or image
	//     alt text and is `null` when a link has none.
	//
	// Any of "google-maps", "google-maps-reviews", "google-play",
	// "google-play-reviews", "app-store", "app-store-reviews", "g2-reviews",
	// "instagram-profile", "instagram-post", "instagram-feed", "tiktok-profile",
	// "tiktok-video", "tiktok-feed", "tiktok-comments", "tiktok-followers",
	// "tiktok-following", "tiktok-user-search", "reddit-post", "reddit-search",
	// "reddit-subreddit", "reddit-user", "yelp-search", "yelp-reviews",
	// "tripadvisor-reviews", "car-rental-availability", "car-rental-locations",
	// "emails", "phone-numbers", "links".
	Parser string `json:"parser,omitzero" api:"required"`
	// Maximum number of reviews (`google-maps-reviews`, default 100), comments
	// (`reddit-post`, default 100), or results (`reddit-search`, `reddit-subreddit`,
	// and `reddit-user`, default 100) to return. Only these parsers accept it.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// Order of `data.reviews` for `google-maps-reviews`. `newest`, `relevant`,
	// `highest`, and `lowest` use the same orders as Google Maps; `oldest` returns the
	// earliest reviews first.
	//
	// Any of "newest", "oldest", "relevant", "highest", "lowest".
	Sort string `json:"sort,omitzero"`
	// Use a prebuilt parser instead of custom rules.
	//
	// This field can be elided, and will marshal its zero value as "prebuilt".
	Type constant.Prebuilt `json:"type" default:"prebuilt"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingParseParamsPrebuiltParser) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingParseParamsPrebuiltParser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingParseParamsPrebuiltParser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingParseParamsPrebuiltParser](
		"parser", "google-maps", "google-maps-reviews", "google-play", "google-play-reviews", "app-store", "app-store-reviews", "g2-reviews", "instagram-profile", "instagram-post", "instagram-feed", "tiktok-profile", "tiktok-video", "tiktok-feed", "tiktok-comments", "tiktok-followers", "tiktok-following", "tiktok-user-search", "reddit-post", "reddit-search", "reddit-subreddit", "reddit-user", "yelp-search", "yelp-reviews", "tripadvisor-reviews", "car-rental-availability", "car-rental-locations", "emails", "phone-numbers", "links",
	)
	apijson.RegisterFieldValidator[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingParseParamsPrebuiltParser](
		"sort", "newest", "oldest", "relevant", "highest", "lowest",
	)
}

// Options for the `product` output. Requires `formats.product: true`.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingProductParams struct {
	// Drop visually duplicate product images, keeping the largest copy.
	DedupeImages param.Opt[bool] `json:"dedupeImages,omitzero"`
	// Ignored. The AI fallback runs automatically.
	//
	// Deprecated: deprecated
	UseAIFallback param.Opt[bool] `json:"useAIFallback,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingProductParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingProductParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingProductParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Options for the `screenshot` output. Requires `formats.screenshot: true`.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingScreenshotParams struct {
	// What to capture: `viewport`, `fullPage`, one element, or a rectangle. Up to 40
	// megapixels.
	Area BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingScreenshotParamsAreaUnion `json:"area,omitzero"`
	// Image format for the screenshot.
	//
	// Any of "png", "jpeg", "webp".
	Format string `json:"format,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingScreenshotParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingScreenshotParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingScreenshotParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingScreenshotParams](
		"format", "png", "jpeg", "webp",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingScreenshotParamsAreaUnion struct {
	// Check if union is this variant with !param.IsOmitted(union.OfPage)
	OfPage      param.Opt[string]                                                                    `json:",omitzero,inline"`
	OfElement   *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingScreenshotParamsAreaElement   `json:",omitzero,inline"`
	OfRectangle *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingScreenshotParamsAreaRectangle `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingScreenshotParamsAreaUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPage, u.OfElement, u.OfRectangle)
}
func (u *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingScreenshotParamsAreaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingScreenshotParamsAreaPage string

const (
	BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingScreenshotParamsAreaPageViewport BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingScreenshotParamsAreaPage = "viewport"
	BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingScreenshotParamsAreaPageFullPage BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingScreenshotParamsAreaPage = "fullPage"
)

// The property Selector is required.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingScreenshotParamsAreaElement struct {
	// CSS selector matching exactly one visible element.
	Selector string `json:"selector" api:"required"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingScreenshotParamsAreaElement) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingScreenshotParamsAreaElement
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingScreenshotParamsAreaElement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rectangle to capture, in pixels from the document origin.
//
// The properties Height, Width, X, Y are required.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingScreenshotParamsAreaRectangle struct {
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

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingScreenshotParamsAreaRectangle) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingScreenshotParamsAreaRectangle
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingScreenshotParamsAreaRectangle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Browser and content settings shared by all outputs.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParams struct {
	// Country to fetch from, as an ISO 3166-1 alpha-2 code such as `US`.
	// Case-insensitive.
	Country param.Opt[string] `json:"country,omitzero"`
	// Accept cookie banners before actions and capture.
	DismissCookies param.Opt[bool] `json:"dismissCookies,omitzero"`
	// Close other popups before actions and capture.
	DismissPopups param.Opt[bool] `json:"dismissPopups,omitzero"`
	// Keep only the main content, without navigation, headers, footers, or sidebars.
	// Does not affect `screenshot`, `bytes`, or `product`.
	MainContentOnly param.Opt[bool] `json:"mainContentOnly,omitzero"`
	// Wait for CSS animations to finish before capture. Defaults to `true` when
	// `formats.screenshot` is `true`.
	SettleAnimations param.Opt[bool] `json:"settleAnimations,omitzero"`
	// Browser steps to run in order before capture. Requires a paid plan. Skips the
	// cache.
	Actions []BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionUnion `json:"actions,omitzero"`
	// Remove elements matching these CSS selectors. Overrides `includeSelectors`.
	ExcludeSelectors []string `json:"excludeSelectors,omitzero"`
	// HTTP headers to send to the target site. Requests with headers skip the cache.
	Headers map[string]string `json:"headers,omitzero"`
	// Which iframes to inline in HTML and text outputs. Screenshots always show
	// visible frames. `auto` inlines only iframes that hold page content, such as
	// same-site documents, framesets, job boards, forms, and embedded documents, and
	// skips ads and widgets. `true` inlines every iframe; `false` inlines none.
	IncludeFrames BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsIncludeFramesUnion `json:"includeFrames,omitzero"`
	// Keep only elements matching these CSS selectors.
	IncludeSelectors []string `json:"includeSelectors,omitzero"`
	// Options for parsing documents such as PDFs.
	Parsers BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsParsers `json:"parsers,omitzero"`
	// Scrape as a signed-in user. Create a profile with `type: "new"` and your
	// credentials, then reuse it with `type: "existing"`. Requires a paid plan and an
	// HTTPS URL. Profiles keep their session between requests, skip the cache, and
	// cannot be combined with ZDR, custom headers, `formats.bytes`, or prebuilt
	// parsers other than `emails`, `phone-numbers`, and `links`.
	Profile BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsProfileUnion `json:"profile,omitzero"`
	// Emulate a light or dark color scheme.
	//
	// Any of "light", "dark".
	Theme string `json:"theme,omitzero"`
	// Browser viewport in pixels. Defaults to 1920 × 1080. If you pass `viewport`, a
	// missing `width` defaults to 1440 and a missing `height` to 900.
	Viewport BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsViewport `json:"viewport,omitzero"`
	// Time to wait after actions, in milliseconds, or a CSS selector to wait for.
	// Defaults to 500 ms, or 2000 ms when `includeFrames` is `true` or the URL ends in
	// `.xml`.
	WaitFor BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsWaitForUnion `json:"waitFor,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParams](
		"theme", "light", "dark",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionUnion struct {
	OfPerform *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionPerform `json:",omitzero,inline"`
	OfScroll  *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionScroll  `json:",omitzero,inline"`
	OfWait    *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionWait    `json:",omitzero,inline"`
	OfWaitFor *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionWaitFor `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPerform, u.OfScroll, u.OfWait, u.OfWaitFor)
}
func (u *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionUnion](
		"type",
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionPerform]("perform"),
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionScroll]("scroll"),
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionWait]("wait"),
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionWaitFor]("waitFor"),
	)
}

// The properties Action, Type are required.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionPerform struct {
	// One browser instruction, such as clicking a button or entering text.
	Action string `json:"action" api:"required"`
	// Follow a plain-language instruction.
	//
	// This field can be elided, and will marshal its zero value as "perform".
	Type constant.Perform `json:"type" default:"perform"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionPerform) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionPerform
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionPerform) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionScroll struct {
	// Maximum number of scroll steps for this action.
	MaxScrolls param.Opt[int64] `json:"maxScrolls,omitzero"`
	// CSS selector of the container to scroll. Omit to scroll the page.
	Selector param.Opt[string] `json:"selector,omitzero"`
	// Distance per scroll step: a number of pixels, one `viewport`, or `max` to reach
	// the end.
	Amount BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionScrollAmountUnion `json:"amount,omitzero"`
	// Direction to scroll.
	//
	// Any of "down", "up", "left", "right".
	Direction string `json:"direction,omitzero"`
	// Scroll the page or a container.
	//
	// This field can be elided, and will marshal its zero value as "scroll".
	Type constant.Scroll `json:"type" default:"scroll"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionScroll) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionScroll
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionScroll) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionScroll](
		"direction", "down", "up", "left", "right",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionScrollAmountUnion struct {
	OfInt param.Opt[int64] `json:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfBatchSubmitsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionScrollAmountString)
	OfBatchSubmitsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionScrollAmountString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionScrollAmountUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfBatchSubmitsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionScrollAmountString)
}
func (u *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionScrollAmountUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionScrollAmountString string

const (
	BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionScrollAmountStringViewport BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionScrollAmountString = "viewport"
	BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionScrollAmountStringMax      BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionScrollAmountString = "max"
)

// The properties Milliseconds, Type are required.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionWait struct {
	// How long to pause before the next action, in milliseconds.
	Milliseconds int64 `json:"milliseconds" api:"required"`
	// Pause for a fixed time.
	//
	// This field can be elided, and will marshal its zero value as "wait".
	Type constant.Wait `json:"type" default:"wait"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionWait) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionWait
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionWait) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Selector, Type are required.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionWaitFor struct {
	// CSS selector to wait for before continuing.
	Selector string `json:"selector" api:"required"`
	// Wait for an element matching `selector`.
	//
	// This field can be elided, and will marshal its zero value as "waitFor".
	Type constant.WaitFor `json:"type" default:"waitFor"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionWaitFor) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionWaitFor
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsActionWaitFor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsIncludeFramesUnion struct {
	OfBool param.Opt[bool] `json:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfBatchSubmitsInputCrawlDataCrawlWithScrapeAnythingSharedParamsIncludeFramesString)
	OfBatchSubmitsInputCrawlDataCrawlWithScrapeAnythingSharedParamsIncludeFramesString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsIncludeFramesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfBatchSubmitsInputCrawlDataCrawlWithScrapeAnythingSharedParamsIncludeFramesString)
}
func (u *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsIncludeFramesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsIncludeFramesString string

const (
	BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsIncludeFramesStringAuto BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsIncludeFramesString = "auto"
)

// Options for parsing documents such as PDFs.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsParsers struct {
	// Page range, OCR, and page markers for PDFs.
	Pdf BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsParsersPdf `json:"pdf,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsParsers) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsParsers
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsParsers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Page range, OCR, and page markers for PDFs.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsParsersPdf struct {
	// Last page to parse. Defaults to the last page. Must be at least `startPage`.
	EndPage param.Opt[int64] `json:"endPage,omitzero"`
	// Start each PDF page in the Markdown with a `<!-- page N -->` line, where `N` is
	// the original 1-based page number. These responses are never served from cache.
	IncludeNumberedPages param.Opt[bool] `json:"includeNumberedPages,omitzero"`
	// First page to parse, starting at 1.
	StartPage param.Opt[int64] `json:"startPage,omitzero"`
	// Set to `auto` to read scanned pages with OCR. Adds 1 credit per page read.
	//
	// Any of "off", "auto".
	Ocr string `json:"ocr,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsParsersPdf) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsParsersPdf
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsParsersPdf) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsParsersPdf](
		"ocr", "off", "auto",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsProfileUnion struct {
	OfExisting *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsProfileExisting `json:",omitzero,inline"`
	OfNew      *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsProfileNew      `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsProfileUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExisting, u.OfNew)
}
func (u *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsProfileUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsProfileUnion](
		"type",
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsProfileExisting]("existing"),
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsProfileNew]("new"),
	)
}

// The properties ProfileID, Type are required.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsProfileExisting struct {
	// ID of the browser profile, from `profile.id` in an earlier scrape response.
	ProfileID string `json:"profileId" api:"required"`
	// Reuse a saved browser profile.
	//
	// This field can be elided, and will marshal its zero value as "existing".
	Type constant.Existing `json:"type" default:"existing"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsProfileExisting) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsProfileExisting
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsProfileExisting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Identifier, Password, Type are required.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsProfileNew struct {
	// Username or email address for the website.
	Identifier string `json:"identifier" api:"required"`
	// Password for the website. Used to sign in and never stored.
	Password string `json:"password" api:"required"`
	// Label for the profile, shown by `GET /org/browser/profiles`.
	Description param.Opt[string] `json:"description,omitzero"`
	// Sign in and save a new browser profile.
	//
	// This field can be elided, and will marshal its zero value as "new".
	Type constant.New `json:"type" default:"new"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsProfileNew) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsProfileNew
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsProfileNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Browser viewport in pixels. Defaults to 1920 × 1080. If you pass `viewport`, a
// missing `width` defaults to 1440 and a missing `height` to 900.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsViewport struct {
	// Browser viewport height in pixels.
	Height param.Opt[int64] `json:"height,omitzero"`
	// Browser viewport width in pixels.
	Width param.Opt[int64] `json:"width,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsViewport) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsViewport
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsViewport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsWaitForUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsWaitForUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingSharedParamsWaitForUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Time limit for each page, with the same behavior and defaults as in
// `POST /web/scrape`.
//
// The property Milliseconds is required.
type BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingTimeoutOpts struct {
	// Deadline in milliseconds.
	Milliseconds int64 `json:"milliseconds" api:"required"`
	// `fail` returns 408 at the deadline. `return-partial` returns the results
	// collected so far and marks the response as partial. `return-partial` requires at
	// least 5000 ms. On `POST /web/scrape`, an output that is not ready in time fails
	// with `REQUEST_TIMEOUT` inside a 200 response under either behavior; 408 means
	// the whole request could not respond by the deadline.
	//
	// Any of "fail", "return-partial".
	Behavior string `json:"behavior,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingTimeoutOpts) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingTimeoutOpts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingTimeoutOpts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputCrawlDataCrawlWithScrapeAnythingTimeoutOpts](
		"behavior", "fail", "return-partial",
	)
}

// Crawl pages and return Markdown.
//
// Deprecated: deprecated
//
// The properties Format, Source are required.
type BatchSubmitParamsInputCrawlDataMarkdown struct {
	// Return page content as Markdown.
	//
	// Any of "markdown".
	Format string `json:"format,omitzero" api:"required"`
	// How to find pages to crawl.
	Source BatchSubmitParamsInputCrawlDataMarkdownSourceUnion `json:"source,omitzero" api:"required"`
	// Options for Markdown output.
	Options BatchSubmitParamsInputCrawlDataMarkdownOptions `json:"options,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataMarkdown) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataMarkdown
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataMarkdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputCrawlDataMarkdown](
		"format", "markdown",
	)
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
	// Page where crawling begins. A URL without a scheme is treated as `https://`.
	URL string `json:"url" api:"required"`
	// Limits and filters for page discovery.
	Controls BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControls `json:"controls,omitzero"`
	// Follow links from one URL.
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
	// Maximum link depth, where the start page is depth `0`. No limit when omitted.
	MaxDepth param.Opt[int64] `json:"maxDepth,omitzero"`
	// Maximum number of pages to fetch, up to 100,000. Credits are reserved for this
	// many pages and the unused part is refunded.
	MaxURLs param.Opt[int64] `json:"maxUrls,omitzero"`
	// Pages to skip entirely, even when they match `follow` or `scrape`.
	Exclude []BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsExcludeUnion `json:"exclude,omitzero"`
	// Links to follow from each page; every followed page is checked for links again.
	// Follows all links when omitted.
	Follow []BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsFollowUnion `json:"follow,omitzero"`
	// Only follow pages in these languages, or pass `same_as_start_url` to match the
	// start page. Follows all languages when omitted.
	Languages BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsLanguagesUnion `json:"languages,omitzero"`
	// Pages to return in the results. Returns every fetched page when omitted.
	Scrape []BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsScrapeUnion `json:"scrape,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControls) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControls
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControls) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsExcludeUnion struct {
	OfRegex   *BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsExcludeRegex   `json:",omitzero,inline"`
	OfProduct *BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsExcludeProduct `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsExcludeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfRegex, u.OfProduct)
}
func (u *BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsExcludeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsExcludeUnion](
		"kind",
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsExcludeRegex]("regex"),
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsExcludeProduct]("product"),
	)
}

// Pages whose URL matches a regex pattern.
//
// The properties Kind, Pattern are required.
type BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsExcludeRegex struct {
	// Regular expression matched against the page URL.
	Pattern string `json:"pattern" api:"required"`
	// Match pages by URL pattern.
	//
	// This field can be elided, and will marshal its zero value as "regex".
	Kind constant.Regex `json:"kind" default:"regex"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsExcludeRegex) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsExcludeRegex
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsExcludeRegex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewBatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsExcludeProduct() BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsExcludeProduct {
	return BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsExcludeProduct{
		Kind: "product",
	}
}

// Product detail pages.
//
// This struct has a constant value, construct it with
// [NewBatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsExcludeProduct].
type BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsExcludeProduct struct {
	// Match product detail pages.
	Kind constant.Product `json:"kind" default:"product"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsExcludeProduct) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsExcludeProduct
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsExcludeProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsFollowUnion struct {
	OfRegex      *BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsFollowRegex      `json:",omitzero,inline"`
	OfPagination *BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsFollowPagination `json:",omitzero,inline"`
	OfProduct    *BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsFollowProduct    `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsFollowUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfRegex, u.OfPagination, u.OfProduct)
}
func (u *BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsFollowUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsFollowUnion](
		"kind",
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsFollowRegex]("regex"),
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsFollowPagination]("pagination"),
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsFollowProduct]("product"),
	)
}

// Pages whose URL matches a regex pattern.
//
// The properties Kind, Pattern are required.
type BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsFollowRegex struct {
	// Regular expression matched against the page URL.
	Pattern string `json:"pattern" api:"required"`
	// Match pages by URL pattern.
	//
	// This field can be elided, and will marshal its zero value as "regex".
	Kind constant.Regex `json:"kind" default:"regex"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsFollowRegex) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsFollowRegex
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsFollowRegex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewBatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsFollowPagination() BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsFollowPagination {
	return BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsFollowPagination{
		Kind: "pagination",
	}
}

// Further pages of a paginated list, such as page 2 of a catalog.
//
// This struct has a constant value, construct it with
// [NewBatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsFollowPagination].
type BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsFollowPagination struct {
	// Match the further pages of a paginated list.
	Kind constant.Pagination `json:"kind" default:"pagination"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsFollowPagination) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsFollowPagination
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsFollowPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewBatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsFollowProduct() BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsFollowProduct {
	return BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsFollowProduct{
		Kind: "product",
	}
}

// Product detail pages.
//
// This struct has a constant value, construct it with
// [NewBatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsFollowProduct].
type BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsFollowProduct struct {
	// Match product detail pages.
	Kind constant.Product `json:"kind" default:"product"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsFollowProduct) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsFollowProduct
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsFollowProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsLanguagesUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfBatchSubmitsInputCrawlDataMarkdownSourceStartURLControlsLanguagesString)
	OfBatchSubmitsInputCrawlDataMarkdownSourceStartURLControlsLanguagesString         param.Opt[string] `json:",omitzero,inline"`
	OfBatchSubmitsInputCrawlDataMarkdownSourceStartURLControlsLanguagesArrayItemArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsLanguagesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBatchSubmitsInputCrawlDataMarkdownSourceStartURLControlsLanguagesString, u.OfBatchSubmitsInputCrawlDataMarkdownSourceStartURLControlsLanguagesArrayItemArray)
}
func (u *BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsLanguagesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsLanguagesString string

const (
	BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsLanguagesStringSameAsStartURL BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsLanguagesString = "same_as_start_url"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsScrapeUnion struct {
	OfRegex   *BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsScrapeRegex   `json:",omitzero,inline"`
	OfProduct *BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsScrapeProduct `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsScrapeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfRegex, u.OfProduct)
}
func (u *BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsScrapeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsScrapeUnion](
		"kind",
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsScrapeRegex]("regex"),
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsScrapeProduct]("product"),
	)
}

// Pages whose URL matches a regex pattern.
//
// The properties Kind, Pattern are required.
type BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsScrapeRegex struct {
	// Regular expression matched against the page URL.
	Pattern string `json:"pattern" api:"required"`
	// Match pages by URL pattern.
	//
	// This field can be elided, and will marshal its zero value as "regex".
	Kind constant.Regex `json:"kind" default:"regex"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsScrapeRegex) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsScrapeRegex
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsScrapeRegex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewBatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsScrapeProduct() BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsScrapeProduct {
	return BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsScrapeProduct{
		Kind: "product",
	}
}

// Product detail pages.
//
// This struct has a constant value, construct it with
// [NewBatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsScrapeProduct].
type BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsScrapeProduct struct {
	// Match product detail pages.
	Kind constant.Product `json:"kind" default:"product"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsScrapeProduct) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsScrapeProduct
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataMarkdownSourceStartURLControlsScrapeProduct) UnmarshalJSON(data []byte) error {
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
	// Limits and filters for the sitemap URLs. Links on sitemap pages are never
	// followed.
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

// Limits and filters for the sitemap URLs. Links on sitemap pages are never
// followed.
type BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControls struct {
	// Maximum number of pages to fetch, up to 100,000. Credits are reserved for this
	// many pages and the unused part is refunded.
	MaxURLs param.Opt[int64] `json:"maxUrls,omitzero"`
	// Sitemap pages to skip, even when they match `scrape`.
	Exclude []BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsExcludeUnion `json:"exclude,omitzero"`
	// Path fragments that exclude a sitemap URL. Applies in addition to `scrape` and
	// `exclude`; pass `[]` to keep every listed URL.
	ExcludePaths []string `json:"excludePaths,omitzero"`
	// Sitemap pages to return in the results. Returns every listed page when omitted.
	Scrape []BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsScrapeUnion `json:"scrape,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControls) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControls
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControls) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsExcludeUnion struct {
	OfRegex   *BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsExcludeRegex   `json:",omitzero,inline"`
	OfProduct *BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsExcludeProduct `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsExcludeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfRegex, u.OfProduct)
}
func (u *BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsExcludeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsExcludeUnion](
		"kind",
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsExcludeRegex]("regex"),
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsExcludeProduct]("product"),
	)
}

// Pages whose URL matches a regex pattern.
//
// The properties Kind, Pattern are required.
type BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsExcludeRegex struct {
	// Regular expression matched against the page URL.
	Pattern string `json:"pattern" api:"required"`
	// Match pages by URL pattern.
	//
	// This field can be elided, and will marshal its zero value as "regex".
	Kind constant.Regex `json:"kind" default:"regex"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsExcludeRegex) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsExcludeRegex
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsExcludeRegex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewBatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsExcludeProduct() BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsExcludeProduct {
	return BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsExcludeProduct{
		Kind: "product",
	}
}

// Product detail pages.
//
// This struct has a constant value, construct it with
// [NewBatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsExcludeProduct].
type BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsExcludeProduct struct {
	// Match product detail pages.
	Kind constant.Product `json:"kind" default:"product"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsExcludeProduct) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsExcludeProduct
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsExcludeProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsScrapeUnion struct {
	OfRegex   *BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsScrapeRegex   `json:",omitzero,inline"`
	OfProduct *BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsScrapeProduct `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsScrapeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfRegex, u.OfProduct)
}
func (u *BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsScrapeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsScrapeUnion](
		"kind",
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsScrapeRegex]("regex"),
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsScrapeProduct]("product"),
	)
}

// Pages whose URL matches a regex pattern.
//
// The properties Kind, Pattern are required.
type BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsScrapeRegex struct {
	// Regular expression matched against the page URL.
	Pattern string `json:"pattern" api:"required"`
	// Match pages by URL pattern.
	//
	// This field can be elided, and will marshal its zero value as "regex".
	Kind constant.Regex `json:"kind" default:"regex"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsScrapeRegex) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsScrapeRegex
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsScrapeRegex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewBatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsScrapeProduct() BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsScrapeProduct {
	return BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsScrapeProduct{
		Kind: "product",
	}
}

// Product detail pages.
//
// This struct has a constant value, construct it with
// [NewBatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsScrapeProduct].
type BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsScrapeProduct struct {
	// Match product detail pages.
	Kind constant.Product `json:"kind" default:"product"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsScrapeProduct) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsScrapeProduct
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataMarkdownSourceSitemapControlsScrapeProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Options for Markdown output.
type BatchSubmitParamsInputCrawlDataMarkdownOptions struct {
	// Maximum cache age in milliseconds. Defaults to 3 days (259200000 ms). Maximum: 1
	// year (31536000000 ms). `0` fetches fresh.
	MaxAgeMs param.Opt[int64] `json:"maxAgeMs,omitzero"`
	// Also return each page's HTML in `html`.
	IncludeHTML param.Opt[bool] `json:"includeHTML,omitzero"`
	// Include image references in the Markdown (`html-conversion` only).
	IncludeImages param.Opt[bool] `json:"includeImages,omitzero"`
	// Include links in the Markdown (`html-conversion` only).
	IncludeLinks param.Opt[bool] `json:"includeLinks,omitzero"`
	// Wait for CSS animations to finish before extracting, on browser-rendered pages.
	SettleAnimations param.Opt[bool] `json:"settleAnimations,omitzero"`
	// Shorten inline base64 image data (`html-conversion` only).
	ShortenBase64Images param.Opt[bool] `json:"shortenBase64Images,omitzero"`
	// Return the main content without navigation or footers.
	UseMainContentOnly param.Opt[bool] `json:"useMainContentOnly,omitzero"`
	// How long to wait after initial page load, in milliseconds. `0` waits 500 ms.
	WaitForMs param.Opt[int64] `json:"waitForMs,omitzero"`
	// Remove elements matching these CSS selectors. Applied after `includeSelectors`,
	// so an element matching both is removed.
	ExcludeSelectors []string `json:"excludeSelectors,omitzero"`
	// Keep only elements matching these CSS selectors. Filtered pages ignore
	// `maxAgeMs`.
	IncludeSelectors []string `json:"includeSelectors,omitzero"`
	// Country to fetch from, as an ISO 3166-1 alpha-2 code.
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
	// Whether to parse PDFs and which pages to read.
	Pdf BatchSubmitParamsInputCrawlDataMarkdownOptionsPdf `json:"pdf,omitzero"`
	// Markdown sources in order of preference: `html-conversion` converts the page
	// HTML, `markdown-direct` uses Markdown the site publishes.
	//
	// Any of "html-conversion", "markdown-direct".
	SourcePreference []string `json:"sourcePreference,omitzero"`
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

// Whether to parse PDFs and which pages to read.
type BatchSubmitParamsInputCrawlDataMarkdownOptionsPdf struct {
	// Last page to parse. Defaults to the last page. Must be at least `start`.
	End param.Opt[int64] `json:"end,omitzero"`
	// Read scanned PDF pages with OCR. Pages that already have text are kept as is.
	Ocr param.Opt[bool] `json:"ocr,omitzero"`
	// Parse PDF URLs. When false, PDFs fail with `PDF_SKIPPED`.
	ShouldParse param.Opt[bool] `json:"shouldParse,omitzero"`
	// First page to parse, starting at 1.
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
// Deprecated: deprecated
//
// The properties Format, Source are required.
type BatchSubmitParamsInputCrawlDataHTML struct {
	// Return page content as HTML.
	//
	// Any of "html".
	Format string `json:"format,omitzero" api:"required"`
	// How to find pages to crawl.
	Source BatchSubmitParamsInputCrawlDataHTMLSourceUnion `json:"source,omitzero" api:"required"`
	// Options for HTML output.
	Options BatchSubmitParamsInputCrawlDataHTMLOptions `json:"options,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataHTML) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataHTML
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataHTML) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BatchSubmitParamsInputCrawlDataHTML](
		"format", "html",
	)
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
	// Page where crawling begins. A URL without a scheme is treated as `https://`.
	URL string `json:"url" api:"required"`
	// Limits and filters for page discovery.
	Controls BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControls `json:"controls,omitzero"`
	// Follow links from one URL.
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
	// Maximum link depth, where the start page is depth `0`. No limit when omitted.
	MaxDepth param.Opt[int64] `json:"maxDepth,omitzero"`
	// Maximum number of pages to fetch, up to 100,000. Credits are reserved for this
	// many pages and the unused part is refunded.
	MaxURLs param.Opt[int64] `json:"maxUrls,omitzero"`
	// Pages to skip entirely, even when they match `follow` or `scrape`.
	Exclude []BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsExcludeUnion `json:"exclude,omitzero"`
	// Links to follow from each page; every followed page is checked for links again.
	// Follows all links when omitted.
	Follow []BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsFollowUnion `json:"follow,omitzero"`
	// Only follow pages in these languages, or pass `same_as_start_url` to match the
	// start page. Follows all languages when omitted.
	Languages BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsLanguagesUnion `json:"languages,omitzero"`
	// Pages to return in the results. Returns every fetched page when omitted.
	Scrape []BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsScrapeUnion `json:"scrape,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControls) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControls
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControls) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsExcludeUnion struct {
	OfRegex   *BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsExcludeRegex   `json:",omitzero,inline"`
	OfProduct *BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsExcludeProduct `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsExcludeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfRegex, u.OfProduct)
}
func (u *BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsExcludeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsExcludeUnion](
		"kind",
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsExcludeRegex]("regex"),
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsExcludeProduct]("product"),
	)
}

// Pages whose URL matches a regex pattern.
//
// The properties Kind, Pattern are required.
type BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsExcludeRegex struct {
	// Regular expression matched against the page URL.
	Pattern string `json:"pattern" api:"required"`
	// Match pages by URL pattern.
	//
	// This field can be elided, and will marshal its zero value as "regex".
	Kind constant.Regex `json:"kind" default:"regex"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsExcludeRegex) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsExcludeRegex
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsExcludeRegex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewBatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsExcludeProduct() BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsExcludeProduct {
	return BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsExcludeProduct{
		Kind: "product",
	}
}

// Product detail pages.
//
// This struct has a constant value, construct it with
// [NewBatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsExcludeProduct].
type BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsExcludeProduct struct {
	// Match product detail pages.
	Kind constant.Product `json:"kind" default:"product"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsExcludeProduct) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsExcludeProduct
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsExcludeProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsFollowUnion struct {
	OfRegex      *BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsFollowRegex      `json:",omitzero,inline"`
	OfPagination *BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsFollowPagination `json:",omitzero,inline"`
	OfProduct    *BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsFollowProduct    `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsFollowUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfRegex, u.OfPagination, u.OfProduct)
}
func (u *BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsFollowUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsFollowUnion](
		"kind",
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsFollowRegex]("regex"),
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsFollowPagination]("pagination"),
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsFollowProduct]("product"),
	)
}

// Pages whose URL matches a regex pattern.
//
// The properties Kind, Pattern are required.
type BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsFollowRegex struct {
	// Regular expression matched against the page URL.
	Pattern string `json:"pattern" api:"required"`
	// Match pages by URL pattern.
	//
	// This field can be elided, and will marshal its zero value as "regex".
	Kind constant.Regex `json:"kind" default:"regex"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsFollowRegex) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsFollowRegex
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsFollowRegex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewBatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsFollowPagination() BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsFollowPagination {
	return BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsFollowPagination{
		Kind: "pagination",
	}
}

// Further pages of a paginated list, such as page 2 of a catalog.
//
// This struct has a constant value, construct it with
// [NewBatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsFollowPagination].
type BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsFollowPagination struct {
	// Match the further pages of a paginated list.
	Kind constant.Pagination `json:"kind" default:"pagination"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsFollowPagination) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsFollowPagination
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsFollowPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewBatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsFollowProduct() BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsFollowProduct {
	return BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsFollowProduct{
		Kind: "product",
	}
}

// Product detail pages.
//
// This struct has a constant value, construct it with
// [NewBatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsFollowProduct].
type BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsFollowProduct struct {
	// Match product detail pages.
	Kind constant.Product `json:"kind" default:"product"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsFollowProduct) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsFollowProduct
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsFollowProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsLanguagesUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfBatchSubmitsInputCrawlDataHTMLSourceStartURLControlsLanguagesString)
	OfBatchSubmitsInputCrawlDataHTMLSourceStartURLControlsLanguagesString         param.Opt[string] `json:",omitzero,inline"`
	OfBatchSubmitsInputCrawlDataHTMLSourceStartURLControlsLanguagesArrayItemArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsLanguagesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBatchSubmitsInputCrawlDataHTMLSourceStartURLControlsLanguagesString, u.OfBatchSubmitsInputCrawlDataHTMLSourceStartURLControlsLanguagesArrayItemArray)
}
func (u *BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsLanguagesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsLanguagesString string

const (
	BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsLanguagesStringSameAsStartURL BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsLanguagesString = "same_as_start_url"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsScrapeUnion struct {
	OfRegex   *BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsScrapeRegex   `json:",omitzero,inline"`
	OfProduct *BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsScrapeProduct `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsScrapeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfRegex, u.OfProduct)
}
func (u *BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsScrapeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsScrapeUnion](
		"kind",
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsScrapeRegex]("regex"),
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsScrapeProduct]("product"),
	)
}

// Pages whose URL matches a regex pattern.
//
// The properties Kind, Pattern are required.
type BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsScrapeRegex struct {
	// Regular expression matched against the page URL.
	Pattern string `json:"pattern" api:"required"`
	// Match pages by URL pattern.
	//
	// This field can be elided, and will marshal its zero value as "regex".
	Kind constant.Regex `json:"kind" default:"regex"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsScrapeRegex) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsScrapeRegex
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsScrapeRegex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewBatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsScrapeProduct() BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsScrapeProduct {
	return BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsScrapeProduct{
		Kind: "product",
	}
}

// Product detail pages.
//
// This struct has a constant value, construct it with
// [NewBatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsScrapeProduct].
type BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsScrapeProduct struct {
	// Match product detail pages.
	Kind constant.Product `json:"kind" default:"product"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsScrapeProduct) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsScrapeProduct
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataHTMLSourceStartURLControlsScrapeProduct) UnmarshalJSON(data []byte) error {
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
	// Limits and filters for the sitemap URLs. Links on sitemap pages are never
	// followed.
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

// Limits and filters for the sitemap URLs. Links on sitemap pages are never
// followed.
type BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControls struct {
	// Maximum number of pages to fetch, up to 100,000. Credits are reserved for this
	// many pages and the unused part is refunded.
	MaxURLs param.Opt[int64] `json:"maxUrls,omitzero"`
	// Sitemap pages to skip, even when they match `scrape`.
	Exclude []BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsExcludeUnion `json:"exclude,omitzero"`
	// Path fragments that exclude a sitemap URL. Applies in addition to `scrape` and
	// `exclude`; pass `[]` to keep every listed URL.
	ExcludePaths []string `json:"excludePaths,omitzero"`
	// Sitemap pages to return in the results. Returns every listed page when omitted.
	Scrape []BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsScrapeUnion `json:"scrape,omitzero"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControls) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControls
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControls) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsExcludeUnion struct {
	OfRegex   *BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsExcludeRegex   `json:",omitzero,inline"`
	OfProduct *BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsExcludeProduct `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsExcludeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfRegex, u.OfProduct)
}
func (u *BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsExcludeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsExcludeUnion](
		"kind",
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsExcludeRegex]("regex"),
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsExcludeProduct]("product"),
	)
}

// Pages whose URL matches a regex pattern.
//
// The properties Kind, Pattern are required.
type BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsExcludeRegex struct {
	// Regular expression matched against the page URL.
	Pattern string `json:"pattern" api:"required"`
	// Match pages by URL pattern.
	//
	// This field can be elided, and will marshal its zero value as "regex".
	Kind constant.Regex `json:"kind" default:"regex"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsExcludeRegex) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsExcludeRegex
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsExcludeRegex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewBatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsExcludeProduct() BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsExcludeProduct {
	return BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsExcludeProduct{
		Kind: "product",
	}
}

// Product detail pages.
//
// This struct has a constant value, construct it with
// [NewBatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsExcludeProduct].
type BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsExcludeProduct struct {
	// Match product detail pages.
	Kind constant.Product `json:"kind" default:"product"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsExcludeProduct) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsExcludeProduct
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsExcludeProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsScrapeUnion struct {
	OfRegex   *BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsScrapeRegex   `json:",omitzero,inline"`
	OfProduct *BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsScrapeProduct `json:",omitzero,inline"`
	paramUnion
}

func (u BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsScrapeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfRegex, u.OfProduct)
}
func (u *BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsScrapeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsScrapeUnion](
		"kind",
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsScrapeRegex]("regex"),
		apijson.Discriminator[BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsScrapeProduct]("product"),
	)
}

// Pages whose URL matches a regex pattern.
//
// The properties Kind, Pattern are required.
type BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsScrapeRegex struct {
	// Regular expression matched against the page URL.
	Pattern string `json:"pattern" api:"required"`
	// Match pages by URL pattern.
	//
	// This field can be elided, and will marshal its zero value as "regex".
	Kind constant.Regex `json:"kind" default:"regex"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsScrapeRegex) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsScrapeRegex
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsScrapeRegex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewBatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsScrapeProduct() BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsScrapeProduct {
	return BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsScrapeProduct{
		Kind: "product",
	}
}

// Product detail pages.
//
// This struct has a constant value, construct it with
// [NewBatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsScrapeProduct].
type BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsScrapeProduct struct {
	// Match product detail pages.
	Kind constant.Product `json:"kind" default:"product"`
	paramObj
}

func (r BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsScrapeProduct) MarshalJSON() (data []byte, err error) {
	type shadow BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsScrapeProduct
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchSubmitParamsInputCrawlDataHTMLSourceSitemapControlsScrapeProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Options for HTML output.
type BatchSubmitParamsInputCrawlDataHTMLOptions struct {
	// Maximum cache age in milliseconds. Defaults to 3 days (259200000 ms). Maximum: 1
	// year (31536000000 ms). `0` fetches fresh.
	MaxAgeMs param.Opt[int64] `json:"maxAgeMs,omitzero"`
	// Wait for CSS animations to finish before extracting, on browser-rendered pages.
	SettleAnimations param.Opt[bool] `json:"settleAnimations,omitzero"`
	// Return the main content without navigation or footers.
	UseMainContentOnly param.Opt[bool] `json:"useMainContentOnly,omitzero"`
	// How long to wait after initial page load, in milliseconds. `0` waits 500 ms.
	WaitForMs param.Opt[int64] `json:"waitForMs,omitzero"`
	// Remove elements matching these CSS selectors. Applied after `includeSelectors`,
	// so an element matching both is removed.
	ExcludeSelectors []string `json:"excludeSelectors,omitzero"`
	// Keep only elements matching these CSS selectors. Filtered pages ignore
	// `maxAgeMs`.
	IncludeSelectors []string `json:"includeSelectors,omitzero"`
	// Country to fetch from, as an ISO 3166-1 alpha-2 code.
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
	// Whether to parse PDFs and which pages to read.
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

// Whether to parse PDFs and which pages to read.
type BatchSubmitParamsInputCrawlDataHTMLOptionsPdf struct {
	// Last page to parse. Defaults to the last page. Must be at least `start`.
	End param.Opt[int64] `json:"end,omitzero"`
	// Read scanned PDF pages with OCR. Pages that already have text are kept as is.
	Ocr param.Opt[bool] `json:"ocr,omitzero"`
	// Parse PDF URLs. When false, PDFs fail with `PDF_SKIPPED`.
	ShouldParse param.Opt[bool] `json:"shouldParse,omitzero"`
	// First page to parse, starting at 1.
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

// Webhook to notify when the batch finishes. Omit `retry` for a single attempt, or
// pass `retry: {}` for the default retry schedule.
//
// The property URL is required.
type BatchSubmitParamsWebhook struct {
	// Public HTTP or HTTPS URL that receives the `batch.completed`, `batch.failed`, or
	// `batch.cancelled` event.
	URL string `json:"url" api:"required" format:"uri"`
	// Automatic retry settings. Pass `{}` for the default schedule.
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
