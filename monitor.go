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
	"time"

	"github.com/context-dot-dev/context-go-sdk/v2/internal/apijson"
	"github.com/context-dot-dev/context-go-sdk/v2/internal/apiquery"
	"github.com/context-dot-dev/context-go-sdk/v2/internal/requestconfig"
	"github.com/context-dot-dev/context-go-sdk/v2/option"
	"github.com/context-dot-dev/context-go-sdk/v2/packages/param"
	"github.com/context-dot-dev/context-go-sdk/v2/packages/respjson"
	"github.com/context-dot-dev/context-go-sdk/v2/shared/constant"
)

// Monitor pages, sitemaps, and extracted website data for exact or semantic
// changes. Webhook payloads are documented by the
// MonitorsChangeDetectedWebhookPayload and MonitorsRunCompletedWebhookPayload
// schemas.
//
// MonitorService contains methods and other services that help with interacting
// with the context.dev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMonitorService] method instead.
type MonitorService struct {
	options []option.RequestOption
}

// NewMonitorService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMonitorService(opts ...option.RequestOption) (r MonitorService) {
	r = MonitorService{}
	r.options = opts
	return
}

// Creates a monitor. The request body is a union of the supported target/change
// detection combinations. The monitor runs immediately after creation to create
// its initial baseline.
func (r *MonitorService) New(ctx context.Context, body MonitorNewParams, opts ...option.RequestOption) (res *MonitorNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "monitors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a monitor
func (r *MonitorService) Get(ctx context.Context, monitorID string, opts ...option.RequestOption) (res *MonitorGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if monitorID == "" {
		err = errors.New("missing required monitor_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("monitors/%s", url.PathEscape(monitorID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a monitor. If `target` or `change_detection` changes, the monitor
// creates a new baseline. Unsupported target/change detection combinations are
// rejected.
func (r *MonitorService) Update(ctx context.Context, monitorID string, body MonitorUpdateParams, opts ...option.RequestOption) (res *MonitorUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if monitorID == "" {
		err = errors.New("missing required monitor_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("monitors/%s", url.PathEscape(monitorID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Lists monitors for the authenticated organization. Supports free-text search
// (`q` over `search_by` fields, `prefix` or `exact` via `search_type`) plus
// status/type/tag filters. Results are paginated via the opaque `cursor`.
func (r *MonitorService) List(ctx context.Context, query MonitorListParams, opts ...option.RequestOption) (res *MonitorListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "monitors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a monitor
func (r *MonitorService) Delete(ctx context.Context, monitorID string, opts ...option.RequestOption) (res *MonitorDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if monitorID == "" {
		err = errors.New("missing required monitor_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("monitors/%s", url.PathEscape(monitorID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Returns credits charged per monitor over an optional [since, until] window,
// newest spenders first.
func (r *MonitorService) GetCreditUsage(ctx context.Context, query MonitorGetCreditUsageParams, opts ...option.RequestOption) (res *MonitorGetCreditUsageResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "monitors/credit-usage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns how many monitors the account has and the maximum it allows.
func (r *MonitorService) GetLimits(ctx context.Context, opts ...option.RequestOption) (res *MonitorGetLimitsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "monitors/limits"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns an account-wide feed of detected changes across monitors.
func (r *MonitorService) ListAccountChanges(ctx context.Context, query MonitorListAccountChangesParams, opts ...option.RequestOption) (res *MonitorListAccountChangesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "monitors/changes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns an account-wide feed of monitor runs across all monitors.
func (r *MonitorService) ListAccountRuns(ctx context.Context, query MonitorListAccountRunsParams, opts ...option.RequestOption) (res *MonitorListAccountRunsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "monitors/runs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List changes for a monitor
func (r *MonitorService) ListChanges(ctx context.Context, monitorID string, query MonitorListChangesParams, opts ...option.RequestOption) (res *MonitorListChangesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if monitorID == "" {
		err = errors.New("missing required monitor_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("monitors/%s/changes", url.PathEscape(monitorID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List monitor runs
func (r *MonitorService) ListRuns(ctx context.Context, monitorID string, query MonitorListRunsParams, opts ...option.RequestOption) (res *MonitorListRunsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if monitorID == "" {
		err = errors.New("missing required monitor_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("monitors/%s/runs", url.PathEscape(monitorID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get a change
func (r *MonitorService) GetChange(ctx context.Context, changeID string, opts ...option.RequestOption) (res *MonitorGetChangeResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if changeID == "" {
		err = errors.New("missing required change_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("monitors/changes/%s", url.PathEscape(changeID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Triggers an immediate run of the monitor outside its normal schedule. The run is
// queued and processed asynchronously.
func (r *MonitorService) Run(ctx context.Context, monitorID string, opts ...option.RequestOption) (res *MonitorRunResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if monitorID == "" {
		err = errors.New("missing required monitor_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("monitors/%s/run", url.PathEscape(monitorID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// A web monitor. `mode` is the constant `web`; behavior is described by `target`
// (page/sitemap/extract) and `change_detection` (exact/semantic).
type MonitorNewResponse struct {
	ID string `json:"id" api:"required"`
	// Discriminated union describing how changes are detected.
	ChangeDetection MonitorNewResponseChangeDetectionUnion `json:"change_detection" api:"required"`
	CreatedAt       time.Time                              `json:"created_at" api:"required" format:"date-time"`
	// Top-level monitor category. Always `web` today; the concrete behavior is
	// described by `target` and `change_detection`.
	//
	// Any of "web".
	Mode MonitorNewResponseMode `json:"mode" api:"required"`
	Name string                 `json:"name" api:"required"`
	// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
	// every 6 hours or every 2 days. The total interval (frequency × unit) must be
	// between 10 minutes and 1 year.
	Schedule MonitorNewResponseSchedule `json:"schedule" api:"required"`
	// Monitor lifecycle status. `failed` means the most recent run failed (see the
	// monitor's `last_error`); failed monitors keep running on schedule and flip back
	// to `active` on the next successful run. Monitors are auto-`paused` after
	// repeated consecutive failures or insufficient-credit skips; resume by PATCHing
	// status to `active`.
	//
	// Any of "active", "paused", "failed".
	Status MonitorNewResponseStatus `json:"status" api:"required"`
	// Discriminated union describing what the monitor watches.
	Target    MonitorNewResponseTargetUnion `json:"target" api:"required"`
	UpdatedAt time.Time                     `json:"updated_at" api:"required" format:"date-time"`
	// Current baseline: the last observed value the monitor compares new snapshots
	// against. Its shape follows `target.type` (page/sitemap/extract). Only populated
	// on GET /monitors/{monitor_id}; null until the first baseline run completes (and
	// after a target or change_detection update, which resets the baseline).
	Baseline     MonitorNewResponseBaselineUnion `json:"baseline" api:"nullable"`
	LastChangeAt time.Time                       `json:"last_change_at" api:"nullable" format:"date-time"`
	// Error from the most recent failed run; null when the last run succeeded.
	LastError MonitorNewResponseLastError `json:"last_error" api:"nullable"`
	LastRunAt time.Time                   `json:"last_run_at" api:"nullable" format:"date-time"`
	// When the next scheduled run is due.
	NextRunAt time.Time `json:"next_run_at" api:"nullable" format:"date-time"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags    []string                  `json:"tags"`
	Webhook MonitorNewResponseWebhook `json:"webhook" api:"nullable"`
	// Present while webhook deliveries are failing consecutively; null when deliveries
	// are healthy or no webhook is configured. Cleared on the next successful delivery
	// and when the webhook URL changes.
	WebhookFailure MonitorNewResponseWebhookFailure `json:"webhook_failure" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ChangeDetection respjson.Field
		CreatedAt       respjson.Field
		Mode            respjson.Field
		Name            respjson.Field
		Schedule        respjson.Field
		Status          respjson.Field
		Target          respjson.Field
		UpdatedAt       respjson.Field
		Baseline        respjson.Field
		LastChangeAt    respjson.Field
		LastError       respjson.Field
		LastRunAt       respjson.Field
		NextRunAt       respjson.Field
		Tags            respjson.Field
		Webhook         respjson.Field
		WebhookFailure  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorNewResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MonitorNewResponseChangeDetectionUnion contains all possible properties and
// values from [MonitorNewResponseChangeDetectionExact],
// [MonitorNewResponseChangeDetectionSemantic].
//
// Use the [MonitorNewResponseChangeDetectionUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MonitorNewResponseChangeDetectionUnion struct {
	// Any of "exact", "semantic".
	Type string `json:"type"`
	// This field is from variant [MonitorNewResponseChangeDetectionSemantic].
	ConfidenceThreshold float64 `json:"confidence_threshold"`
	JSON                struct {
		Type                respjson.Field
		ConfidenceThreshold respjson.Field
		raw                 string
	} `json:"-"`
}

// anyMonitorNewResponseChangeDetection is implemented by each variant of
// [MonitorNewResponseChangeDetectionUnion] to add type safety for the return type
// of [MonitorNewResponseChangeDetectionUnion.AsAny]
type anyMonitorNewResponseChangeDetection interface {
	implMonitorNewResponseChangeDetectionUnion()
}

func (MonitorNewResponseChangeDetectionExact) implMonitorNewResponseChangeDetectionUnion()    {}
func (MonitorNewResponseChangeDetectionSemantic) implMonitorNewResponseChangeDetectionUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := MonitorNewResponseChangeDetectionUnion.AsAny().(type) {
//	case contextdev.MonitorNewResponseChangeDetectionExact:
//	case contextdev.MonitorNewResponseChangeDetectionSemantic:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u MonitorNewResponseChangeDetectionUnion) AsAny() anyMonitorNewResponseChangeDetection {
	switch u.Type {
	case "exact":
		return u.AsExact()
	case "semantic":
		return u.AsSemantic()
	}
	return nil
}

func (u MonitorNewResponseChangeDetectionUnion) AsExact() (v MonitorNewResponseChangeDetectionExact) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorNewResponseChangeDetectionUnion) AsSemantic() (v MonitorNewResponseChangeDetectionSemantic) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MonitorNewResponseChangeDetectionUnion) RawJSON() string { return u.JSON.raw }

func (r *MonitorNewResponseChangeDetectionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detect exact changes. For page targets, this means visible text diffs. For
// sitemap targets, this means URL additions and removals.
type MonitorNewResponseChangeDetectionExact struct {
	Type constant.Exact `json:"type" default:"exact"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorNewResponseChangeDetectionExact) RawJSON() string { return r.JSON.raw }
func (r *MonitorNewResponseChangeDetectionExact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detect meaning-level changes to tracked page content, ignoring cosmetic or
// paraphrase-only differences. Which changes are meaningful is judged against the
// extract target's `instructions` (and `schema`, when provided).
type MonitorNewResponseChangeDetectionSemantic struct {
	Type                constant.Semantic `json:"type" default:"semantic"`
	ConfidenceThreshold float64           `json:"confidence_threshold"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type                respjson.Field
		ConfidenceThreshold respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorNewResponseChangeDetectionSemantic) RawJSON() string { return r.JSON.raw }
func (r *MonitorNewResponseChangeDetectionSemantic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Top-level monitor category. Always `web` today; the concrete behavior is
// described by `target` and `change_detection`.
type MonitorNewResponseMode string

const (
	MonitorNewResponseModeWeb MonitorNewResponseMode = "web"
)

// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
// every 6 hours or every 2 days. The total interval (frequency × unit) must be
// between 10 minutes and 1 year.
type MonitorNewResponseSchedule struct {
	// Number of units between runs. The resulting interval (frequency × unit) must be
	// at least 10 minutes and at most 1 year (e.g. minimum 10 when unit is minutes;
	// maximum 365 when unit is days).
	Frequency int64 `json:"frequency" api:"required"`
	// Any of "interval".
	Type string `json:"type" api:"required"`
	// Any of "minutes", "hours", "days".
	Unit string `json:"unit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Frequency   respjson.Field
		Type        respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorNewResponseSchedule) RawJSON() string { return r.JSON.raw }
func (r *MonitorNewResponseSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monitor lifecycle status. `failed` means the most recent run failed (see the
// monitor's `last_error`); failed monitors keep running on schedule and flip back
// to `active` on the next successful run. Monitors are auto-`paused` after
// repeated consecutive failures or insufficient-credit skips; resume by PATCHing
// status to `active`.
type MonitorNewResponseStatus string

const (
	MonitorNewResponseStatusActive MonitorNewResponseStatus = "active"
	MonitorNewResponseStatusPaused MonitorNewResponseStatus = "paused"
	MonitorNewResponseStatusFailed MonitorNewResponseStatus = "failed"
)

// MonitorNewResponseTargetUnion contains all possible properties and values from
// [MonitorNewResponseTargetPage], [MonitorNewResponseTargetSitemap],
// [MonitorNewResponseTargetExtract].
//
// Use the [MonitorNewResponseTargetUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MonitorNewResponseTargetUnion struct {
	// Any of "page", "sitemap", "extract".
	Type string `json:"type"`
	URL  string `json:"url"`
	// This field is from variant [MonitorNewResponseTargetPage].
	NormalizeWhitespace bool `json:"normalize_whitespace"`
	// This field is from variant [MonitorNewResponseTargetSitemap].
	Exclude []string `json:"exclude"`
	// This field is from variant [MonitorNewResponseTargetSitemap].
	Include []string `json:"include"`
	// This field is from variant [MonitorNewResponseTargetSitemap].
	MaxURLs int64 `json:"max_urls"`
	// This field is from variant [MonitorNewResponseTargetExtract].
	Instructions string `json:"instructions"`
	// This field is from variant [MonitorNewResponseTargetExtract].
	FollowSubdomains bool `json:"follow_subdomains"`
	// This field is from variant [MonitorNewResponseTargetExtract].
	MaxDepth int64 `json:"max_depth"`
	// This field is from variant [MonitorNewResponseTargetExtract].
	MaxPages int64 `json:"max_pages"`
	// This field is from variant [MonitorNewResponseTargetExtract].
	Schema map[string]any `json:"schema"`
	JSON   struct {
		Type                respjson.Field
		URL                 respjson.Field
		NormalizeWhitespace respjson.Field
		Exclude             respjson.Field
		Include             respjson.Field
		MaxURLs             respjson.Field
		Instructions        respjson.Field
		FollowSubdomains    respjson.Field
		MaxDepth            respjson.Field
		MaxPages            respjson.Field
		Schema              respjson.Field
		raw                 string
	} `json:"-"`
}

// anyMonitorNewResponseTarget is implemented by each variant of
// [MonitorNewResponseTargetUnion] to add type safety for the return type of
// [MonitorNewResponseTargetUnion.AsAny]
type anyMonitorNewResponseTarget interface {
	implMonitorNewResponseTargetUnion()
}

func (MonitorNewResponseTargetPage) implMonitorNewResponseTargetUnion()    {}
func (MonitorNewResponseTargetSitemap) implMonitorNewResponseTargetUnion() {}
func (MonitorNewResponseTargetExtract) implMonitorNewResponseTargetUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := MonitorNewResponseTargetUnion.AsAny().(type) {
//	case contextdev.MonitorNewResponseTargetPage:
//	case contextdev.MonitorNewResponseTargetSitemap:
//	case contextdev.MonitorNewResponseTargetExtract:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u MonitorNewResponseTargetUnion) AsAny() anyMonitorNewResponseTarget {
	switch u.Type {
	case "page":
		return u.AsPage()
	case "sitemap":
		return u.AsSitemap()
	case "extract":
		return u.AsExtract()
	}
	return nil
}

func (u MonitorNewResponseTargetUnion) AsPage() (v MonitorNewResponseTargetPage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorNewResponseTargetUnion) AsSitemap() (v MonitorNewResponseTargetSitemap) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorNewResponseTargetUnion) AsExtract() (v MonitorNewResponseTargetExtract) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MonitorNewResponseTargetUnion) RawJSON() string { return u.JSON.raw }

func (r *MonitorNewResponseTargetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Watch a single web page.
type MonitorNewResponseTargetPage struct {
	Type constant.Page `json:"type" default:"page"`
	URL  string        `json:"url" api:"required" format:"uri"`
	// Normalize whitespace before comparing or analyzing text.
	NormalizeWhitespace bool `json:"normalize_whitespace"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type                respjson.Field
		URL                 respjson.Field
		NormalizeWhitespace respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorNewResponseTargetPage) RawJSON() string { return r.JSON.raw }
func (r *MonitorNewResponseTargetPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Watch a sitemap for URL additions and removals. Crawled URLs are normalized
// (lowercased host, no trailing slash/fragment) and scoped to the monitored site
// and its subdomains before comparison. On a detected difference the sitemap is
// re-fetched within the same run and only URLs both observations agree on are
// reported, suppressing transient crawl flaps.
type MonitorNewResponseTargetSitemap struct {
	Type constant.Sitemap `json:"type" default:"sitemap"`
	// Sitemap URL to monitor.
	URL string `json:"url" api:"required" format:"uri"`
	// URL path patterns to exclude.
	Exclude []string `json:"exclude"`
	// URL path patterns to include.
	Include []string `json:"include"`
	// Maximum number of sitemap URLs to track (capped at 10,000).
	MaxURLs int64 `json:"max_urls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URL         respjson.Field
		Exclude     respjson.Field
		Include     respjson.Field
		MaxURLs     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorNewResponseTargetSitemap) RawJSON() string { return r.JSON.raw }
func (r *MonitorNewResponseTargetSitemap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Watch the monitor-relevant pages of a site for meaningful changes. A crawl
// guided by `schema`/`instructions` selects up to `max_pages` relevant pages to
// track; each run re-checks exactly those pages, and confirmed content changes are
// judged for relevance against the monitor's `instructions` (and `schema`, when
// provided). The tracked page set is refreshed by a periodic re-discovery crawl.
type MonitorNewResponseTargetExtract struct {
	// Natural-language instructions guiding which pages and facts to track and which
	// changes to report.
	Instructions string           `json:"instructions" api:"required"`
	Type         constant.Extract `json:"type" default:"extract"`
	// Root URL to extract structured data from.
	URL              string `json:"url" api:"required" format:"uri"`
	FollowSubdomains bool   `json:"follow_subdomains"`
	// Optional maximum link depth from the starting URL (0 = only the starting page).
	MaxDepth int64 `json:"max_depth"`
	// Maximum number of pages to track.
	MaxPages int64 `json:"max_pages"`
	// JSON Schema describing the data you care about. It is used three ways: it guides
	// which pages are selected for tracking, it gives the change judge extra context
	// on which changes matter (alongside `instructions`), and it defines the shape of
	// the baseline `data` snapshot on GET /monitors/{monitor_id} (refreshed at most
	// about once a day). It is not a response format for changes: change events and
	// webhook payloads always contain diffs, summaries, and evidence excerpts — never
	// data in this schema's shape. If omitted, a default summary + key-points schema
	// is used.
	Schema map[string]any `json:"schema"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Instructions     respjson.Field
		Type             respjson.Field
		URL              respjson.Field
		FollowSubdomains respjson.Field
		MaxDepth         respjson.Field
		MaxPages         respjson.Field
		Schema           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorNewResponseTargetExtract) RawJSON() string { return r.JSON.raw }
func (r *MonitorNewResponseTargetExtract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MonitorNewResponseBaselineUnion contains all possible properties and values from
// [MonitorNewResponseBaselinePageBaseline],
// [MonitorNewResponseBaselineSitemapBaseline],
// [MonitorNewResponseBaselineExtractBaseline].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MonitorNewResponseBaselineUnion struct {
	CapturedAt time.Time `json:"captured_at"`
	// This field is from variant [MonitorNewResponseBaselinePageBaseline].
	Text string `json:"text"`
	// This field is from variant [MonitorNewResponseBaselineSitemapBaseline].
	URLCount int64 `json:"url_count"`
	// This field is from variant [MonitorNewResponseBaselineSitemapBaseline].
	URLs []string `json:"urls"`
	// This field is from variant [MonitorNewResponseBaselineExtractBaseline].
	Data any `json:"data"`
	// This field is from variant [MonitorNewResponseBaselineExtractBaseline].
	URLsAnalyzed []string `json:"urls_analyzed"`
	JSON         struct {
		CapturedAt   respjson.Field
		Text         respjson.Field
		URLCount     respjson.Field
		URLs         respjson.Field
		Data         respjson.Field
		URLsAnalyzed respjson.Field
		raw          string
	} `json:"-"`
}

func (u MonitorNewResponseBaselineUnion) AsPageBaseline() (v MonitorNewResponseBaselinePageBaseline) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorNewResponseBaselineUnion) AsSitemapBaseline() (v MonitorNewResponseBaselineSitemapBaseline) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorNewResponseBaselineUnion) AsExtractBaseline() (v MonitorNewResponseBaselineExtractBaseline) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MonitorNewResponseBaselineUnion) RawJSON() string { return u.JSON.raw }

func (r *MonitorNewResponseBaselineUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current baseline of a `page` monitor: the visible page text as last observed.
type MonitorNewResponseBaselinePageBaseline struct {
	// When this baseline was last captured or replaced.
	CapturedAt time.Time `json:"captured_at" api:"required" format:"date-time"`
	// The page's visible text as last observed.
	Text string `json:"text" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CapturedAt  respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorNewResponseBaselinePageBaseline) RawJSON() string { return r.JSON.raw }
func (r *MonitorNewResponseBaselinePageBaseline) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current baseline of a `sitemap` monitor: the normalized URL set as last
// observed.
type MonitorNewResponseBaselineSitemapBaseline struct {
	// When this baseline was last captured or replaced.
	CapturedAt time.Time `json:"captured_at" api:"required" format:"date-time"`
	// Number of URLs in the baseline.
	URLCount int64 `json:"url_count" api:"required"`
	// The sitemap URLs as last observed (sorted, normalized).
	URLs []string `json:"urls" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CapturedAt  respjson.Field
		URLCount    respjson.Field
		URLs        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorNewResponseBaselineSitemapBaseline) RawJSON() string { return r.JSON.raw }
func (r *MonitorNewResponseBaselineSitemapBaseline) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current baseline of an `extract` monitor: the pages it tracks and the structured
// data as last extracted.
type MonitorNewResponseBaselineExtractBaseline struct {
	// When this baseline was last captured or replaced.
	CapturedAt time.Time `json:"captured_at" api:"required" format:"date-time"`
	// The extracted structured data, matching the monitor's extraction schema (same
	// shape as the /web/extract endpoint's `data`). Refreshed when the monitor
	// re-discovers its page set (at most about once a day); `null` when no extraction
	// has been captured yet.
	Data any `json:"data" api:"required"`
	// The page URLs the monitor tracks and analyzes for changes.
	URLsAnalyzed []string `json:"urls_analyzed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CapturedAt   respjson.Field
		Data         respjson.Field
		URLsAnalyzed respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorNewResponseBaselineExtractBaseline) RawJSON() string { return r.JSON.raw }
func (r *MonitorNewResponseBaselineExtractBaseline) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error from the most recent failed run; null when the last run succeeded.
type MonitorNewResponseLastError struct {
	Code    string `json:"code" api:"required"`
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
func (r MonitorNewResponseLastError) RawJSON() string { return r.JSON.raw }
func (r *MonitorNewResponseLastError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorNewResponseWebhook struct {
	// Webhook URL events are delivered to.
	URL string `json:"url" api:"required" format:"uri"`
	// Events delivered to this endpoint. `change.detected` fires only when a run
	// detects a change; `run.completed` fires on every completed run — including runs
	// that detected no change — and embeds the change when one was detected. Defaults
	// to `["change.detected"]` when omitted.
	//
	// Any of "change.detected", "run.completed".
	Events []string `json:"events"`
	// Signing secret used to verify webhook authenticity. Each delivery includes an
	// `X-Context-Signature: t=<unix>,v1=<hmac>` header, where the HMAC is SHA-256 over
	// `"{t}.{rawRequestBody}"` keyed by this secret. Recompute it with a constant-time
	// compare and reject stale timestamps to prevent replay. Generated by the API;
	// cannot be set by clients.
	Secret string `json:"secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Events      respjson.Field
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorNewResponseWebhook) RawJSON() string { return r.JSON.raw }
func (r *MonitorNewResponseWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Present while webhook deliveries are failing consecutively; null when deliveries
// are healthy or no webhook is configured. Cleared on the next successful delivery
// and when the webhook URL changes.
type MonitorNewResponseWebhookFailure struct {
	// Number of consecutive delivery attempts that did not succeed.
	ConsecutiveFailures int64     `json:"consecutive_failures" api:"required"`
	LastFailedAt        time.Time `json:"last_failed_at" api:"required" format:"date-time"`
	// Human-readable description of the most recent failure.
	LastMessage string `json:"last_message" api:"required"`
	// Outcome of the most recent failed delivery. rejected means a non-2xx response;
	// failed means no HTTP response was received; skipped_unsafe_url means the URL
	// failed the public-endpoint safety check.
	//
	// Any of "rejected", "failed", "skipped_unsafe_url".
	LastStatus string `json:"last_status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConsecutiveFailures respjson.Field
		LastFailedAt        respjson.Field
		LastMessage         respjson.Field
		LastStatus          respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorNewResponseWebhookFailure) RawJSON() string { return r.JSON.raw }
func (r *MonitorNewResponseWebhookFailure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A web monitor. `mode` is the constant `web`; behavior is described by `target`
// (page/sitemap/extract) and `change_detection` (exact/semantic).
type MonitorGetResponse struct {
	ID string `json:"id" api:"required"`
	// Discriminated union describing how changes are detected.
	ChangeDetection MonitorGetResponseChangeDetectionUnion `json:"change_detection" api:"required"`
	CreatedAt       time.Time                              `json:"created_at" api:"required" format:"date-time"`
	// Top-level monitor category. Always `web` today; the concrete behavior is
	// described by `target` and `change_detection`.
	//
	// Any of "web".
	Mode MonitorGetResponseMode `json:"mode" api:"required"`
	Name string                 `json:"name" api:"required"`
	// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
	// every 6 hours or every 2 days. The total interval (frequency × unit) must be
	// between 10 minutes and 1 year.
	Schedule MonitorGetResponseSchedule `json:"schedule" api:"required"`
	// Monitor lifecycle status. `failed` means the most recent run failed (see the
	// monitor's `last_error`); failed monitors keep running on schedule and flip back
	// to `active` on the next successful run. Monitors are auto-`paused` after
	// repeated consecutive failures or insufficient-credit skips; resume by PATCHing
	// status to `active`.
	//
	// Any of "active", "paused", "failed".
	Status MonitorGetResponseStatus `json:"status" api:"required"`
	// Discriminated union describing what the monitor watches.
	Target    MonitorGetResponseTargetUnion `json:"target" api:"required"`
	UpdatedAt time.Time                     `json:"updated_at" api:"required" format:"date-time"`
	// Current baseline: the last observed value the monitor compares new snapshots
	// against. Its shape follows `target.type` (page/sitemap/extract). Only populated
	// on GET /monitors/{monitor_id}; null until the first baseline run completes (and
	// after a target or change_detection update, which resets the baseline).
	Baseline     MonitorGetResponseBaselineUnion `json:"baseline" api:"nullable"`
	LastChangeAt time.Time                       `json:"last_change_at" api:"nullable" format:"date-time"`
	// Error from the most recent failed run; null when the last run succeeded.
	LastError MonitorGetResponseLastError `json:"last_error" api:"nullable"`
	LastRunAt time.Time                   `json:"last_run_at" api:"nullable" format:"date-time"`
	// When the next scheduled run is due.
	NextRunAt time.Time `json:"next_run_at" api:"nullable" format:"date-time"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags    []string                  `json:"tags"`
	Webhook MonitorGetResponseWebhook `json:"webhook" api:"nullable"`
	// Present while webhook deliveries are failing consecutively; null when deliveries
	// are healthy or no webhook is configured. Cleared on the next successful delivery
	// and when the webhook URL changes.
	WebhookFailure MonitorGetResponseWebhookFailure `json:"webhook_failure" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ChangeDetection respjson.Field
		CreatedAt       respjson.Field
		Mode            respjson.Field
		Name            respjson.Field
		Schedule        respjson.Field
		Status          respjson.Field
		Target          respjson.Field
		UpdatedAt       respjson.Field
		Baseline        respjson.Field
		LastChangeAt    respjson.Field
		LastError       respjson.Field
		LastRunAt       respjson.Field
		NextRunAt       respjson.Field
		Tags            respjson.Field
		Webhook         respjson.Field
		WebhookFailure  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MonitorGetResponseChangeDetectionUnion contains all possible properties and
// values from [MonitorGetResponseChangeDetectionExact],
// [MonitorGetResponseChangeDetectionSemantic].
//
// Use the [MonitorGetResponseChangeDetectionUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MonitorGetResponseChangeDetectionUnion struct {
	// Any of "exact", "semantic".
	Type string `json:"type"`
	// This field is from variant [MonitorGetResponseChangeDetectionSemantic].
	ConfidenceThreshold float64 `json:"confidence_threshold"`
	JSON                struct {
		Type                respjson.Field
		ConfidenceThreshold respjson.Field
		raw                 string
	} `json:"-"`
}

// anyMonitorGetResponseChangeDetection is implemented by each variant of
// [MonitorGetResponseChangeDetectionUnion] to add type safety for the return type
// of [MonitorGetResponseChangeDetectionUnion.AsAny]
type anyMonitorGetResponseChangeDetection interface {
	implMonitorGetResponseChangeDetectionUnion()
}

func (MonitorGetResponseChangeDetectionExact) implMonitorGetResponseChangeDetectionUnion()    {}
func (MonitorGetResponseChangeDetectionSemantic) implMonitorGetResponseChangeDetectionUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := MonitorGetResponseChangeDetectionUnion.AsAny().(type) {
//	case contextdev.MonitorGetResponseChangeDetectionExact:
//	case contextdev.MonitorGetResponseChangeDetectionSemantic:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u MonitorGetResponseChangeDetectionUnion) AsAny() anyMonitorGetResponseChangeDetection {
	switch u.Type {
	case "exact":
		return u.AsExact()
	case "semantic":
		return u.AsSemantic()
	}
	return nil
}

func (u MonitorGetResponseChangeDetectionUnion) AsExact() (v MonitorGetResponseChangeDetectionExact) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorGetResponseChangeDetectionUnion) AsSemantic() (v MonitorGetResponseChangeDetectionSemantic) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MonitorGetResponseChangeDetectionUnion) RawJSON() string { return u.JSON.raw }

func (r *MonitorGetResponseChangeDetectionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detect exact changes. For page targets, this means visible text diffs. For
// sitemap targets, this means URL additions and removals.
type MonitorGetResponseChangeDetectionExact struct {
	Type constant.Exact `json:"type" default:"exact"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetResponseChangeDetectionExact) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetResponseChangeDetectionExact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detect meaning-level changes to tracked page content, ignoring cosmetic or
// paraphrase-only differences. Which changes are meaningful is judged against the
// extract target's `instructions` (and `schema`, when provided).
type MonitorGetResponseChangeDetectionSemantic struct {
	Type                constant.Semantic `json:"type" default:"semantic"`
	ConfidenceThreshold float64           `json:"confidence_threshold"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type                respjson.Field
		ConfidenceThreshold respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetResponseChangeDetectionSemantic) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetResponseChangeDetectionSemantic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Top-level monitor category. Always `web` today; the concrete behavior is
// described by `target` and `change_detection`.
type MonitorGetResponseMode string

const (
	MonitorGetResponseModeWeb MonitorGetResponseMode = "web"
)

// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
// every 6 hours or every 2 days. The total interval (frequency × unit) must be
// between 10 minutes and 1 year.
type MonitorGetResponseSchedule struct {
	// Number of units between runs. The resulting interval (frequency × unit) must be
	// at least 10 minutes and at most 1 year (e.g. minimum 10 when unit is minutes;
	// maximum 365 when unit is days).
	Frequency int64 `json:"frequency" api:"required"`
	// Any of "interval".
	Type string `json:"type" api:"required"`
	// Any of "minutes", "hours", "days".
	Unit string `json:"unit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Frequency   respjson.Field
		Type        respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetResponseSchedule) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetResponseSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monitor lifecycle status. `failed` means the most recent run failed (see the
// monitor's `last_error`); failed monitors keep running on schedule and flip back
// to `active` on the next successful run. Monitors are auto-`paused` after
// repeated consecutive failures or insufficient-credit skips; resume by PATCHing
// status to `active`.
type MonitorGetResponseStatus string

const (
	MonitorGetResponseStatusActive MonitorGetResponseStatus = "active"
	MonitorGetResponseStatusPaused MonitorGetResponseStatus = "paused"
	MonitorGetResponseStatusFailed MonitorGetResponseStatus = "failed"
)

// MonitorGetResponseTargetUnion contains all possible properties and values from
// [MonitorGetResponseTargetPage], [MonitorGetResponseTargetSitemap],
// [MonitorGetResponseTargetExtract].
//
// Use the [MonitorGetResponseTargetUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MonitorGetResponseTargetUnion struct {
	// Any of "page", "sitemap", "extract".
	Type string `json:"type"`
	URL  string `json:"url"`
	// This field is from variant [MonitorGetResponseTargetPage].
	NormalizeWhitespace bool `json:"normalize_whitespace"`
	// This field is from variant [MonitorGetResponseTargetSitemap].
	Exclude []string `json:"exclude"`
	// This field is from variant [MonitorGetResponseTargetSitemap].
	Include []string `json:"include"`
	// This field is from variant [MonitorGetResponseTargetSitemap].
	MaxURLs int64 `json:"max_urls"`
	// This field is from variant [MonitorGetResponseTargetExtract].
	Instructions string `json:"instructions"`
	// This field is from variant [MonitorGetResponseTargetExtract].
	FollowSubdomains bool `json:"follow_subdomains"`
	// This field is from variant [MonitorGetResponseTargetExtract].
	MaxDepth int64 `json:"max_depth"`
	// This field is from variant [MonitorGetResponseTargetExtract].
	MaxPages int64 `json:"max_pages"`
	// This field is from variant [MonitorGetResponseTargetExtract].
	Schema map[string]any `json:"schema"`
	JSON   struct {
		Type                respjson.Field
		URL                 respjson.Field
		NormalizeWhitespace respjson.Field
		Exclude             respjson.Field
		Include             respjson.Field
		MaxURLs             respjson.Field
		Instructions        respjson.Field
		FollowSubdomains    respjson.Field
		MaxDepth            respjson.Field
		MaxPages            respjson.Field
		Schema              respjson.Field
		raw                 string
	} `json:"-"`
}

// anyMonitorGetResponseTarget is implemented by each variant of
// [MonitorGetResponseTargetUnion] to add type safety for the return type of
// [MonitorGetResponseTargetUnion.AsAny]
type anyMonitorGetResponseTarget interface {
	implMonitorGetResponseTargetUnion()
}

func (MonitorGetResponseTargetPage) implMonitorGetResponseTargetUnion()    {}
func (MonitorGetResponseTargetSitemap) implMonitorGetResponseTargetUnion() {}
func (MonitorGetResponseTargetExtract) implMonitorGetResponseTargetUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := MonitorGetResponseTargetUnion.AsAny().(type) {
//	case contextdev.MonitorGetResponseTargetPage:
//	case contextdev.MonitorGetResponseTargetSitemap:
//	case contextdev.MonitorGetResponseTargetExtract:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u MonitorGetResponseTargetUnion) AsAny() anyMonitorGetResponseTarget {
	switch u.Type {
	case "page":
		return u.AsPage()
	case "sitemap":
		return u.AsSitemap()
	case "extract":
		return u.AsExtract()
	}
	return nil
}

func (u MonitorGetResponseTargetUnion) AsPage() (v MonitorGetResponseTargetPage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorGetResponseTargetUnion) AsSitemap() (v MonitorGetResponseTargetSitemap) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorGetResponseTargetUnion) AsExtract() (v MonitorGetResponseTargetExtract) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MonitorGetResponseTargetUnion) RawJSON() string { return u.JSON.raw }

func (r *MonitorGetResponseTargetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Watch a single web page.
type MonitorGetResponseTargetPage struct {
	Type constant.Page `json:"type" default:"page"`
	URL  string        `json:"url" api:"required" format:"uri"`
	// Normalize whitespace before comparing or analyzing text.
	NormalizeWhitespace bool `json:"normalize_whitespace"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type                respjson.Field
		URL                 respjson.Field
		NormalizeWhitespace respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetResponseTargetPage) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetResponseTargetPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Watch a sitemap for URL additions and removals. Crawled URLs are normalized
// (lowercased host, no trailing slash/fragment) and scoped to the monitored site
// and its subdomains before comparison. On a detected difference the sitemap is
// re-fetched within the same run and only URLs both observations agree on are
// reported, suppressing transient crawl flaps.
type MonitorGetResponseTargetSitemap struct {
	Type constant.Sitemap `json:"type" default:"sitemap"`
	// Sitemap URL to monitor.
	URL string `json:"url" api:"required" format:"uri"`
	// URL path patterns to exclude.
	Exclude []string `json:"exclude"`
	// URL path patterns to include.
	Include []string `json:"include"`
	// Maximum number of sitemap URLs to track (capped at 10,000).
	MaxURLs int64 `json:"max_urls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URL         respjson.Field
		Exclude     respjson.Field
		Include     respjson.Field
		MaxURLs     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetResponseTargetSitemap) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetResponseTargetSitemap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Watch the monitor-relevant pages of a site for meaningful changes. A crawl
// guided by `schema`/`instructions` selects up to `max_pages` relevant pages to
// track; each run re-checks exactly those pages, and confirmed content changes are
// judged for relevance against the monitor's `instructions` (and `schema`, when
// provided). The tracked page set is refreshed by a periodic re-discovery crawl.
type MonitorGetResponseTargetExtract struct {
	// Natural-language instructions guiding which pages and facts to track and which
	// changes to report.
	Instructions string           `json:"instructions" api:"required"`
	Type         constant.Extract `json:"type" default:"extract"`
	// Root URL to extract structured data from.
	URL              string `json:"url" api:"required" format:"uri"`
	FollowSubdomains bool   `json:"follow_subdomains"`
	// Optional maximum link depth from the starting URL (0 = only the starting page).
	MaxDepth int64 `json:"max_depth"`
	// Maximum number of pages to track.
	MaxPages int64 `json:"max_pages"`
	// JSON Schema describing the data you care about. It is used three ways: it guides
	// which pages are selected for tracking, it gives the change judge extra context
	// on which changes matter (alongside `instructions`), and it defines the shape of
	// the baseline `data` snapshot on GET /monitors/{monitor_id} (refreshed at most
	// about once a day). It is not a response format for changes: change events and
	// webhook payloads always contain diffs, summaries, and evidence excerpts — never
	// data in this schema's shape. If omitted, a default summary + key-points schema
	// is used.
	Schema map[string]any `json:"schema"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Instructions     respjson.Field
		Type             respjson.Field
		URL              respjson.Field
		FollowSubdomains respjson.Field
		MaxDepth         respjson.Field
		MaxPages         respjson.Field
		Schema           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetResponseTargetExtract) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetResponseTargetExtract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MonitorGetResponseBaselineUnion contains all possible properties and values from
// [MonitorGetResponseBaselinePageBaseline],
// [MonitorGetResponseBaselineSitemapBaseline],
// [MonitorGetResponseBaselineExtractBaseline].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MonitorGetResponseBaselineUnion struct {
	CapturedAt time.Time `json:"captured_at"`
	// This field is from variant [MonitorGetResponseBaselinePageBaseline].
	Text string `json:"text"`
	// This field is from variant [MonitorGetResponseBaselineSitemapBaseline].
	URLCount int64 `json:"url_count"`
	// This field is from variant [MonitorGetResponseBaselineSitemapBaseline].
	URLs []string `json:"urls"`
	// This field is from variant [MonitorGetResponseBaselineExtractBaseline].
	Data any `json:"data"`
	// This field is from variant [MonitorGetResponseBaselineExtractBaseline].
	URLsAnalyzed []string `json:"urls_analyzed"`
	JSON         struct {
		CapturedAt   respjson.Field
		Text         respjson.Field
		URLCount     respjson.Field
		URLs         respjson.Field
		Data         respjson.Field
		URLsAnalyzed respjson.Field
		raw          string
	} `json:"-"`
}

func (u MonitorGetResponseBaselineUnion) AsPageBaseline() (v MonitorGetResponseBaselinePageBaseline) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorGetResponseBaselineUnion) AsSitemapBaseline() (v MonitorGetResponseBaselineSitemapBaseline) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorGetResponseBaselineUnion) AsExtractBaseline() (v MonitorGetResponseBaselineExtractBaseline) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MonitorGetResponseBaselineUnion) RawJSON() string { return u.JSON.raw }

func (r *MonitorGetResponseBaselineUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current baseline of a `page` monitor: the visible page text as last observed.
type MonitorGetResponseBaselinePageBaseline struct {
	// When this baseline was last captured or replaced.
	CapturedAt time.Time `json:"captured_at" api:"required" format:"date-time"`
	// The page's visible text as last observed.
	Text string `json:"text" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CapturedAt  respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetResponseBaselinePageBaseline) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetResponseBaselinePageBaseline) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current baseline of a `sitemap` monitor: the normalized URL set as last
// observed.
type MonitorGetResponseBaselineSitemapBaseline struct {
	// When this baseline was last captured or replaced.
	CapturedAt time.Time `json:"captured_at" api:"required" format:"date-time"`
	// Number of URLs in the baseline.
	URLCount int64 `json:"url_count" api:"required"`
	// The sitemap URLs as last observed (sorted, normalized).
	URLs []string `json:"urls" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CapturedAt  respjson.Field
		URLCount    respjson.Field
		URLs        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetResponseBaselineSitemapBaseline) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetResponseBaselineSitemapBaseline) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current baseline of an `extract` monitor: the pages it tracks and the structured
// data as last extracted.
type MonitorGetResponseBaselineExtractBaseline struct {
	// When this baseline was last captured or replaced.
	CapturedAt time.Time `json:"captured_at" api:"required" format:"date-time"`
	// The extracted structured data, matching the monitor's extraction schema (same
	// shape as the /web/extract endpoint's `data`). Refreshed when the monitor
	// re-discovers its page set (at most about once a day); `null` when no extraction
	// has been captured yet.
	Data any `json:"data" api:"required"`
	// The page URLs the monitor tracks and analyzes for changes.
	URLsAnalyzed []string `json:"urls_analyzed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CapturedAt   respjson.Field
		Data         respjson.Field
		URLsAnalyzed respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetResponseBaselineExtractBaseline) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetResponseBaselineExtractBaseline) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error from the most recent failed run; null when the last run succeeded.
type MonitorGetResponseLastError struct {
	Code    string `json:"code" api:"required"`
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
func (r MonitorGetResponseLastError) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetResponseLastError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorGetResponseWebhook struct {
	// Webhook URL events are delivered to.
	URL string `json:"url" api:"required" format:"uri"`
	// Events delivered to this endpoint. `change.detected` fires only when a run
	// detects a change; `run.completed` fires on every completed run — including runs
	// that detected no change — and embeds the change when one was detected. Defaults
	// to `["change.detected"]` when omitted.
	//
	// Any of "change.detected", "run.completed".
	Events []string `json:"events"`
	// Signing secret used to verify webhook authenticity. Each delivery includes an
	// `X-Context-Signature: t=<unix>,v1=<hmac>` header, where the HMAC is SHA-256 over
	// `"{t}.{rawRequestBody}"` keyed by this secret. Recompute it with a constant-time
	// compare and reject stale timestamps to prevent replay. Generated by the API;
	// cannot be set by clients.
	Secret string `json:"secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Events      respjson.Field
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetResponseWebhook) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetResponseWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Present while webhook deliveries are failing consecutively; null when deliveries
// are healthy or no webhook is configured. Cleared on the next successful delivery
// and when the webhook URL changes.
type MonitorGetResponseWebhookFailure struct {
	// Number of consecutive delivery attempts that did not succeed.
	ConsecutiveFailures int64     `json:"consecutive_failures" api:"required"`
	LastFailedAt        time.Time `json:"last_failed_at" api:"required" format:"date-time"`
	// Human-readable description of the most recent failure.
	LastMessage string `json:"last_message" api:"required"`
	// Outcome of the most recent failed delivery. rejected means a non-2xx response;
	// failed means no HTTP response was received; skipped_unsafe_url means the URL
	// failed the public-endpoint safety check.
	//
	// Any of "rejected", "failed", "skipped_unsafe_url".
	LastStatus string `json:"last_status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConsecutiveFailures respjson.Field
		LastFailedAt        respjson.Field
		LastMessage         respjson.Field
		LastStatus          respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetResponseWebhookFailure) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetResponseWebhookFailure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A web monitor. `mode` is the constant `web`; behavior is described by `target`
// (page/sitemap/extract) and `change_detection` (exact/semantic).
type MonitorUpdateResponse struct {
	ID string `json:"id" api:"required"`
	// Discriminated union describing how changes are detected.
	ChangeDetection MonitorUpdateResponseChangeDetectionUnion `json:"change_detection" api:"required"`
	CreatedAt       time.Time                                 `json:"created_at" api:"required" format:"date-time"`
	// Top-level monitor category. Always `web` today; the concrete behavior is
	// described by `target` and `change_detection`.
	//
	// Any of "web".
	Mode MonitorUpdateResponseMode `json:"mode" api:"required"`
	Name string                    `json:"name" api:"required"`
	// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
	// every 6 hours or every 2 days. The total interval (frequency × unit) must be
	// between 10 minutes and 1 year.
	Schedule MonitorUpdateResponseSchedule `json:"schedule" api:"required"`
	// Monitor lifecycle status. `failed` means the most recent run failed (see the
	// monitor's `last_error`); failed monitors keep running on schedule and flip back
	// to `active` on the next successful run. Monitors are auto-`paused` after
	// repeated consecutive failures or insufficient-credit skips; resume by PATCHing
	// status to `active`.
	//
	// Any of "active", "paused", "failed".
	Status MonitorUpdateResponseStatus `json:"status" api:"required"`
	// Discriminated union describing what the monitor watches.
	Target    MonitorUpdateResponseTargetUnion `json:"target" api:"required"`
	UpdatedAt time.Time                        `json:"updated_at" api:"required" format:"date-time"`
	// Current baseline: the last observed value the monitor compares new snapshots
	// against. Its shape follows `target.type` (page/sitemap/extract). Only populated
	// on GET /monitors/{monitor_id}; null until the first baseline run completes (and
	// after a target or change_detection update, which resets the baseline).
	Baseline     MonitorUpdateResponseBaselineUnion `json:"baseline" api:"nullable"`
	LastChangeAt time.Time                          `json:"last_change_at" api:"nullable" format:"date-time"`
	// Error from the most recent failed run; null when the last run succeeded.
	LastError MonitorUpdateResponseLastError `json:"last_error" api:"nullable"`
	LastRunAt time.Time                      `json:"last_run_at" api:"nullable" format:"date-time"`
	// When the next scheduled run is due.
	NextRunAt time.Time `json:"next_run_at" api:"nullable" format:"date-time"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags    []string                     `json:"tags"`
	Webhook MonitorUpdateResponseWebhook `json:"webhook" api:"nullable"`
	// Present while webhook deliveries are failing consecutively; null when deliveries
	// are healthy or no webhook is configured. Cleared on the next successful delivery
	// and when the webhook URL changes.
	WebhookFailure MonitorUpdateResponseWebhookFailure `json:"webhook_failure" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ChangeDetection respjson.Field
		CreatedAt       respjson.Field
		Mode            respjson.Field
		Name            respjson.Field
		Schedule        respjson.Field
		Status          respjson.Field
		Target          respjson.Field
		UpdatedAt       respjson.Field
		Baseline        respjson.Field
		LastChangeAt    respjson.Field
		LastError       respjson.Field
		LastRunAt       respjson.Field
		NextRunAt       respjson.Field
		Tags            respjson.Field
		Webhook         respjson.Field
		WebhookFailure  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MonitorUpdateResponseChangeDetectionUnion contains all possible properties and
// values from [MonitorUpdateResponseChangeDetectionExact],
// [MonitorUpdateResponseChangeDetectionSemantic].
//
// Use the [MonitorUpdateResponseChangeDetectionUnion.AsAny] method to switch on
// the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MonitorUpdateResponseChangeDetectionUnion struct {
	// Any of "exact", "semantic".
	Type string `json:"type"`
	// This field is from variant [MonitorUpdateResponseChangeDetectionSemantic].
	ConfidenceThreshold float64 `json:"confidence_threshold"`
	JSON                struct {
		Type                respjson.Field
		ConfidenceThreshold respjson.Field
		raw                 string
	} `json:"-"`
}

// anyMonitorUpdateResponseChangeDetection is implemented by each variant of
// [MonitorUpdateResponseChangeDetectionUnion] to add type safety for the return
// type of [MonitorUpdateResponseChangeDetectionUnion.AsAny]
type anyMonitorUpdateResponseChangeDetection interface {
	implMonitorUpdateResponseChangeDetectionUnion()
}

func (MonitorUpdateResponseChangeDetectionExact) implMonitorUpdateResponseChangeDetectionUnion()    {}
func (MonitorUpdateResponseChangeDetectionSemantic) implMonitorUpdateResponseChangeDetectionUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := MonitorUpdateResponseChangeDetectionUnion.AsAny().(type) {
//	case contextdev.MonitorUpdateResponseChangeDetectionExact:
//	case contextdev.MonitorUpdateResponseChangeDetectionSemantic:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u MonitorUpdateResponseChangeDetectionUnion) AsAny() anyMonitorUpdateResponseChangeDetection {
	switch u.Type {
	case "exact":
		return u.AsExact()
	case "semantic":
		return u.AsSemantic()
	}
	return nil
}

func (u MonitorUpdateResponseChangeDetectionUnion) AsExact() (v MonitorUpdateResponseChangeDetectionExact) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorUpdateResponseChangeDetectionUnion) AsSemantic() (v MonitorUpdateResponseChangeDetectionSemantic) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MonitorUpdateResponseChangeDetectionUnion) RawJSON() string { return u.JSON.raw }

func (r *MonitorUpdateResponseChangeDetectionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detect exact changes. For page targets, this means visible text diffs. For
// sitemap targets, this means URL additions and removals.
type MonitorUpdateResponseChangeDetectionExact struct {
	Type constant.Exact `json:"type" default:"exact"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorUpdateResponseChangeDetectionExact) RawJSON() string { return r.JSON.raw }
func (r *MonitorUpdateResponseChangeDetectionExact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detect meaning-level changes to tracked page content, ignoring cosmetic or
// paraphrase-only differences. Which changes are meaningful is judged against the
// extract target's `instructions` (and `schema`, when provided).
type MonitorUpdateResponseChangeDetectionSemantic struct {
	Type                constant.Semantic `json:"type" default:"semantic"`
	ConfidenceThreshold float64           `json:"confidence_threshold"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type                respjson.Field
		ConfidenceThreshold respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorUpdateResponseChangeDetectionSemantic) RawJSON() string { return r.JSON.raw }
func (r *MonitorUpdateResponseChangeDetectionSemantic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Top-level monitor category. Always `web` today; the concrete behavior is
// described by `target` and `change_detection`.
type MonitorUpdateResponseMode string

const (
	MonitorUpdateResponseModeWeb MonitorUpdateResponseMode = "web"
)

// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
// every 6 hours or every 2 days. The total interval (frequency × unit) must be
// between 10 minutes and 1 year.
type MonitorUpdateResponseSchedule struct {
	// Number of units between runs. The resulting interval (frequency × unit) must be
	// at least 10 minutes and at most 1 year (e.g. minimum 10 when unit is minutes;
	// maximum 365 when unit is days).
	Frequency int64 `json:"frequency" api:"required"`
	// Any of "interval".
	Type string `json:"type" api:"required"`
	// Any of "minutes", "hours", "days".
	Unit string `json:"unit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Frequency   respjson.Field
		Type        respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorUpdateResponseSchedule) RawJSON() string { return r.JSON.raw }
func (r *MonitorUpdateResponseSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monitor lifecycle status. `failed` means the most recent run failed (see the
// monitor's `last_error`); failed monitors keep running on schedule and flip back
// to `active` on the next successful run. Monitors are auto-`paused` after
// repeated consecutive failures or insufficient-credit skips; resume by PATCHing
// status to `active`.
type MonitorUpdateResponseStatus string

const (
	MonitorUpdateResponseStatusActive MonitorUpdateResponseStatus = "active"
	MonitorUpdateResponseStatusPaused MonitorUpdateResponseStatus = "paused"
	MonitorUpdateResponseStatusFailed MonitorUpdateResponseStatus = "failed"
)

// MonitorUpdateResponseTargetUnion contains all possible properties and values
// from [MonitorUpdateResponseTargetPage], [MonitorUpdateResponseTargetSitemap],
// [MonitorUpdateResponseTargetExtract].
//
// Use the [MonitorUpdateResponseTargetUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MonitorUpdateResponseTargetUnion struct {
	// Any of "page", "sitemap", "extract".
	Type string `json:"type"`
	URL  string `json:"url"`
	// This field is from variant [MonitorUpdateResponseTargetPage].
	NormalizeWhitespace bool `json:"normalize_whitespace"`
	// This field is from variant [MonitorUpdateResponseTargetSitemap].
	Exclude []string `json:"exclude"`
	// This field is from variant [MonitorUpdateResponseTargetSitemap].
	Include []string `json:"include"`
	// This field is from variant [MonitorUpdateResponseTargetSitemap].
	MaxURLs int64 `json:"max_urls"`
	// This field is from variant [MonitorUpdateResponseTargetExtract].
	Instructions string `json:"instructions"`
	// This field is from variant [MonitorUpdateResponseTargetExtract].
	FollowSubdomains bool `json:"follow_subdomains"`
	// This field is from variant [MonitorUpdateResponseTargetExtract].
	MaxDepth int64 `json:"max_depth"`
	// This field is from variant [MonitorUpdateResponseTargetExtract].
	MaxPages int64 `json:"max_pages"`
	// This field is from variant [MonitorUpdateResponseTargetExtract].
	Schema map[string]any `json:"schema"`
	JSON   struct {
		Type                respjson.Field
		URL                 respjson.Field
		NormalizeWhitespace respjson.Field
		Exclude             respjson.Field
		Include             respjson.Field
		MaxURLs             respjson.Field
		Instructions        respjson.Field
		FollowSubdomains    respjson.Field
		MaxDepth            respjson.Field
		MaxPages            respjson.Field
		Schema              respjson.Field
		raw                 string
	} `json:"-"`
}

// anyMonitorUpdateResponseTarget is implemented by each variant of
// [MonitorUpdateResponseTargetUnion] to add type safety for the return type of
// [MonitorUpdateResponseTargetUnion.AsAny]
type anyMonitorUpdateResponseTarget interface {
	implMonitorUpdateResponseTargetUnion()
}

func (MonitorUpdateResponseTargetPage) implMonitorUpdateResponseTargetUnion()    {}
func (MonitorUpdateResponseTargetSitemap) implMonitorUpdateResponseTargetUnion() {}
func (MonitorUpdateResponseTargetExtract) implMonitorUpdateResponseTargetUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := MonitorUpdateResponseTargetUnion.AsAny().(type) {
//	case contextdev.MonitorUpdateResponseTargetPage:
//	case contextdev.MonitorUpdateResponseTargetSitemap:
//	case contextdev.MonitorUpdateResponseTargetExtract:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u MonitorUpdateResponseTargetUnion) AsAny() anyMonitorUpdateResponseTarget {
	switch u.Type {
	case "page":
		return u.AsPage()
	case "sitemap":
		return u.AsSitemap()
	case "extract":
		return u.AsExtract()
	}
	return nil
}

func (u MonitorUpdateResponseTargetUnion) AsPage() (v MonitorUpdateResponseTargetPage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorUpdateResponseTargetUnion) AsSitemap() (v MonitorUpdateResponseTargetSitemap) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorUpdateResponseTargetUnion) AsExtract() (v MonitorUpdateResponseTargetExtract) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MonitorUpdateResponseTargetUnion) RawJSON() string { return u.JSON.raw }

func (r *MonitorUpdateResponseTargetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Watch a single web page.
type MonitorUpdateResponseTargetPage struct {
	Type constant.Page `json:"type" default:"page"`
	URL  string        `json:"url" api:"required" format:"uri"`
	// Normalize whitespace before comparing or analyzing text.
	NormalizeWhitespace bool `json:"normalize_whitespace"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type                respjson.Field
		URL                 respjson.Field
		NormalizeWhitespace respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorUpdateResponseTargetPage) RawJSON() string { return r.JSON.raw }
func (r *MonitorUpdateResponseTargetPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Watch a sitemap for URL additions and removals. Crawled URLs are normalized
// (lowercased host, no trailing slash/fragment) and scoped to the monitored site
// and its subdomains before comparison. On a detected difference the sitemap is
// re-fetched within the same run and only URLs both observations agree on are
// reported, suppressing transient crawl flaps.
type MonitorUpdateResponseTargetSitemap struct {
	Type constant.Sitemap `json:"type" default:"sitemap"`
	// Sitemap URL to monitor.
	URL string `json:"url" api:"required" format:"uri"`
	// URL path patterns to exclude.
	Exclude []string `json:"exclude"`
	// URL path patterns to include.
	Include []string `json:"include"`
	// Maximum number of sitemap URLs to track (capped at 10,000).
	MaxURLs int64 `json:"max_urls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URL         respjson.Field
		Exclude     respjson.Field
		Include     respjson.Field
		MaxURLs     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorUpdateResponseTargetSitemap) RawJSON() string { return r.JSON.raw }
func (r *MonitorUpdateResponseTargetSitemap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Watch the monitor-relevant pages of a site for meaningful changes. A crawl
// guided by `schema`/`instructions` selects up to `max_pages` relevant pages to
// track; each run re-checks exactly those pages, and confirmed content changes are
// judged for relevance against the monitor's `instructions` (and `schema`, when
// provided). The tracked page set is refreshed by a periodic re-discovery crawl.
type MonitorUpdateResponseTargetExtract struct {
	// Natural-language instructions guiding which pages and facts to track and which
	// changes to report.
	Instructions string           `json:"instructions" api:"required"`
	Type         constant.Extract `json:"type" default:"extract"`
	// Root URL to extract structured data from.
	URL              string `json:"url" api:"required" format:"uri"`
	FollowSubdomains bool   `json:"follow_subdomains"`
	// Optional maximum link depth from the starting URL (0 = only the starting page).
	MaxDepth int64 `json:"max_depth"`
	// Maximum number of pages to track.
	MaxPages int64 `json:"max_pages"`
	// JSON Schema describing the data you care about. It is used three ways: it guides
	// which pages are selected for tracking, it gives the change judge extra context
	// on which changes matter (alongside `instructions`), and it defines the shape of
	// the baseline `data` snapshot on GET /monitors/{monitor_id} (refreshed at most
	// about once a day). It is not a response format for changes: change events and
	// webhook payloads always contain diffs, summaries, and evidence excerpts — never
	// data in this schema's shape. If omitted, a default summary + key-points schema
	// is used.
	Schema map[string]any `json:"schema"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Instructions     respjson.Field
		Type             respjson.Field
		URL              respjson.Field
		FollowSubdomains respjson.Field
		MaxDepth         respjson.Field
		MaxPages         respjson.Field
		Schema           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorUpdateResponseTargetExtract) RawJSON() string { return r.JSON.raw }
func (r *MonitorUpdateResponseTargetExtract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MonitorUpdateResponseBaselineUnion contains all possible properties and values
// from [MonitorUpdateResponseBaselinePageBaseline],
// [MonitorUpdateResponseBaselineSitemapBaseline],
// [MonitorUpdateResponseBaselineExtractBaseline].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MonitorUpdateResponseBaselineUnion struct {
	CapturedAt time.Time `json:"captured_at"`
	// This field is from variant [MonitorUpdateResponseBaselinePageBaseline].
	Text string `json:"text"`
	// This field is from variant [MonitorUpdateResponseBaselineSitemapBaseline].
	URLCount int64 `json:"url_count"`
	// This field is from variant [MonitorUpdateResponseBaselineSitemapBaseline].
	URLs []string `json:"urls"`
	// This field is from variant [MonitorUpdateResponseBaselineExtractBaseline].
	Data any `json:"data"`
	// This field is from variant [MonitorUpdateResponseBaselineExtractBaseline].
	URLsAnalyzed []string `json:"urls_analyzed"`
	JSON         struct {
		CapturedAt   respjson.Field
		Text         respjson.Field
		URLCount     respjson.Field
		URLs         respjson.Field
		Data         respjson.Field
		URLsAnalyzed respjson.Field
		raw          string
	} `json:"-"`
}

func (u MonitorUpdateResponseBaselineUnion) AsPageBaseline() (v MonitorUpdateResponseBaselinePageBaseline) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorUpdateResponseBaselineUnion) AsSitemapBaseline() (v MonitorUpdateResponseBaselineSitemapBaseline) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorUpdateResponseBaselineUnion) AsExtractBaseline() (v MonitorUpdateResponseBaselineExtractBaseline) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MonitorUpdateResponseBaselineUnion) RawJSON() string { return u.JSON.raw }

func (r *MonitorUpdateResponseBaselineUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current baseline of a `page` monitor: the visible page text as last observed.
type MonitorUpdateResponseBaselinePageBaseline struct {
	// When this baseline was last captured or replaced.
	CapturedAt time.Time `json:"captured_at" api:"required" format:"date-time"`
	// The page's visible text as last observed.
	Text string `json:"text" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CapturedAt  respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorUpdateResponseBaselinePageBaseline) RawJSON() string { return r.JSON.raw }
func (r *MonitorUpdateResponseBaselinePageBaseline) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current baseline of a `sitemap` monitor: the normalized URL set as last
// observed.
type MonitorUpdateResponseBaselineSitemapBaseline struct {
	// When this baseline was last captured or replaced.
	CapturedAt time.Time `json:"captured_at" api:"required" format:"date-time"`
	// Number of URLs in the baseline.
	URLCount int64 `json:"url_count" api:"required"`
	// The sitemap URLs as last observed (sorted, normalized).
	URLs []string `json:"urls" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CapturedAt  respjson.Field
		URLCount    respjson.Field
		URLs        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorUpdateResponseBaselineSitemapBaseline) RawJSON() string { return r.JSON.raw }
func (r *MonitorUpdateResponseBaselineSitemapBaseline) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current baseline of an `extract` monitor: the pages it tracks and the structured
// data as last extracted.
type MonitorUpdateResponseBaselineExtractBaseline struct {
	// When this baseline was last captured or replaced.
	CapturedAt time.Time `json:"captured_at" api:"required" format:"date-time"`
	// The extracted structured data, matching the monitor's extraction schema (same
	// shape as the /web/extract endpoint's `data`). Refreshed when the monitor
	// re-discovers its page set (at most about once a day); `null` when no extraction
	// has been captured yet.
	Data any `json:"data" api:"required"`
	// The page URLs the monitor tracks and analyzes for changes.
	URLsAnalyzed []string `json:"urls_analyzed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CapturedAt   respjson.Field
		Data         respjson.Field
		URLsAnalyzed respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorUpdateResponseBaselineExtractBaseline) RawJSON() string { return r.JSON.raw }
func (r *MonitorUpdateResponseBaselineExtractBaseline) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error from the most recent failed run; null when the last run succeeded.
type MonitorUpdateResponseLastError struct {
	Code    string `json:"code" api:"required"`
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
func (r MonitorUpdateResponseLastError) RawJSON() string { return r.JSON.raw }
func (r *MonitorUpdateResponseLastError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorUpdateResponseWebhook struct {
	// Webhook URL events are delivered to.
	URL string `json:"url" api:"required" format:"uri"`
	// Events delivered to this endpoint. `change.detected` fires only when a run
	// detects a change; `run.completed` fires on every completed run — including runs
	// that detected no change — and embeds the change when one was detected. Defaults
	// to `["change.detected"]` when omitted.
	//
	// Any of "change.detected", "run.completed".
	Events []string `json:"events"`
	// Signing secret used to verify webhook authenticity. Each delivery includes an
	// `X-Context-Signature: t=<unix>,v1=<hmac>` header, where the HMAC is SHA-256 over
	// `"{t}.{rawRequestBody}"` keyed by this secret. Recompute it with a constant-time
	// compare and reject stale timestamps to prevent replay. Generated by the API;
	// cannot be set by clients.
	Secret string `json:"secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Events      respjson.Field
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorUpdateResponseWebhook) RawJSON() string { return r.JSON.raw }
func (r *MonitorUpdateResponseWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Present while webhook deliveries are failing consecutively; null when deliveries
// are healthy or no webhook is configured. Cleared on the next successful delivery
// and when the webhook URL changes.
type MonitorUpdateResponseWebhookFailure struct {
	// Number of consecutive delivery attempts that did not succeed.
	ConsecutiveFailures int64     `json:"consecutive_failures" api:"required"`
	LastFailedAt        time.Time `json:"last_failed_at" api:"required" format:"date-time"`
	// Human-readable description of the most recent failure.
	LastMessage string `json:"last_message" api:"required"`
	// Outcome of the most recent failed delivery. rejected means a non-2xx response;
	// failed means no HTTP response was received; skipped_unsafe_url means the URL
	// failed the public-endpoint safety check.
	//
	// Any of "rejected", "failed", "skipped_unsafe_url".
	LastStatus string `json:"last_status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConsecutiveFailures respjson.Field
		LastFailedAt        respjson.Field
		LastMessage         respjson.Field
		LastStatus          respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorUpdateResponseWebhookFailure) RawJSON() string { return r.JSON.raw }
func (r *MonitorUpdateResponseWebhookFailure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListResponse struct {
	Data       []MonitorListResponseData `json:"data" api:"required"`
	HasMore    bool                      `json:"has_more" api:"required"`
	NextCursor string                    `json:"next_cursor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A web monitor. `mode` is the constant `web`; behavior is described by `target`
// (page/sitemap/extract) and `change_detection` (exact/semantic).
type MonitorListResponseData struct {
	ID string `json:"id" api:"required"`
	// Discriminated union describing how changes are detected.
	ChangeDetection MonitorListResponseDataChangeDetectionUnion `json:"change_detection" api:"required"`
	CreatedAt       time.Time                                   `json:"created_at" api:"required" format:"date-time"`
	// Top-level monitor category. Always `web` today; the concrete behavior is
	// described by `target` and `change_detection`.
	//
	// Any of "web".
	Mode string `json:"mode" api:"required"`
	Name string `json:"name" api:"required"`
	// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
	// every 6 hours or every 2 days. The total interval (frequency × unit) must be
	// between 10 minutes and 1 year.
	Schedule MonitorListResponseDataSchedule `json:"schedule" api:"required"`
	// Monitor lifecycle status. `failed` means the most recent run failed (see the
	// monitor's `last_error`); failed monitors keep running on schedule and flip back
	// to `active` on the next successful run. Monitors are auto-`paused` after
	// repeated consecutive failures or insufficient-credit skips; resume by PATCHing
	// status to `active`.
	//
	// Any of "active", "paused", "failed".
	Status string `json:"status" api:"required"`
	// Discriminated union describing what the monitor watches.
	Target    MonitorListResponseDataTargetUnion `json:"target" api:"required"`
	UpdatedAt time.Time                          `json:"updated_at" api:"required" format:"date-time"`
	// Current baseline: the last observed value the monitor compares new snapshots
	// against. Its shape follows `target.type` (page/sitemap/extract). Only populated
	// on GET /monitors/{monitor_id}; null until the first baseline run completes (and
	// after a target or change_detection update, which resets the baseline).
	Baseline     MonitorListResponseDataBaselineUnion `json:"baseline" api:"nullable"`
	LastChangeAt time.Time                            `json:"last_change_at" api:"nullable" format:"date-time"`
	// Error from the most recent failed run; null when the last run succeeded.
	LastError MonitorListResponseDataLastError `json:"last_error" api:"nullable"`
	LastRunAt time.Time                        `json:"last_run_at" api:"nullable" format:"date-time"`
	// When the next scheduled run is due.
	NextRunAt time.Time `json:"next_run_at" api:"nullable" format:"date-time"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags    []string                       `json:"tags"`
	Webhook MonitorListResponseDataWebhook `json:"webhook" api:"nullable"`
	// Present while webhook deliveries are failing consecutively; null when deliveries
	// are healthy or no webhook is configured. Cleared on the next successful delivery
	// and when the webhook URL changes.
	WebhookFailure MonitorListResponseDataWebhookFailure `json:"webhook_failure" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ChangeDetection respjson.Field
		CreatedAt       respjson.Field
		Mode            respjson.Field
		Name            respjson.Field
		Schedule        respjson.Field
		Status          respjson.Field
		Target          respjson.Field
		UpdatedAt       respjson.Field
		Baseline        respjson.Field
		LastChangeAt    respjson.Field
		LastError       respjson.Field
		LastRunAt       respjson.Field
		NextRunAt       respjson.Field
		Tags            respjson.Field
		Webhook         respjson.Field
		WebhookFailure  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListResponseData) RawJSON() string { return r.JSON.raw }
func (r *MonitorListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MonitorListResponseDataChangeDetectionUnion contains all possible properties and
// values from [MonitorListResponseDataChangeDetectionExact],
// [MonitorListResponseDataChangeDetectionSemantic].
//
// Use the [MonitorListResponseDataChangeDetectionUnion.AsAny] method to switch on
// the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MonitorListResponseDataChangeDetectionUnion struct {
	// Any of "exact", "semantic".
	Type string `json:"type"`
	// This field is from variant [MonitorListResponseDataChangeDetectionSemantic].
	ConfidenceThreshold float64 `json:"confidence_threshold"`
	JSON                struct {
		Type                respjson.Field
		ConfidenceThreshold respjson.Field
		raw                 string
	} `json:"-"`
}

// anyMonitorListResponseDataChangeDetection is implemented by each variant of
// [MonitorListResponseDataChangeDetectionUnion] to add type safety for the return
// type of [MonitorListResponseDataChangeDetectionUnion.AsAny]
type anyMonitorListResponseDataChangeDetection interface {
	implMonitorListResponseDataChangeDetectionUnion()
}

func (MonitorListResponseDataChangeDetectionExact) implMonitorListResponseDataChangeDetectionUnion() {
}
func (MonitorListResponseDataChangeDetectionSemantic) implMonitorListResponseDataChangeDetectionUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := MonitorListResponseDataChangeDetectionUnion.AsAny().(type) {
//	case contextdev.MonitorListResponseDataChangeDetectionExact:
//	case contextdev.MonitorListResponseDataChangeDetectionSemantic:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u MonitorListResponseDataChangeDetectionUnion) AsAny() anyMonitorListResponseDataChangeDetection {
	switch u.Type {
	case "exact":
		return u.AsExact()
	case "semantic":
		return u.AsSemantic()
	}
	return nil
}

func (u MonitorListResponseDataChangeDetectionUnion) AsExact() (v MonitorListResponseDataChangeDetectionExact) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorListResponseDataChangeDetectionUnion) AsSemantic() (v MonitorListResponseDataChangeDetectionSemantic) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MonitorListResponseDataChangeDetectionUnion) RawJSON() string { return u.JSON.raw }

func (r *MonitorListResponseDataChangeDetectionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detect exact changes. For page targets, this means visible text diffs. For
// sitemap targets, this means URL additions and removals.
type MonitorListResponseDataChangeDetectionExact struct {
	Type constant.Exact `json:"type" default:"exact"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListResponseDataChangeDetectionExact) RawJSON() string { return r.JSON.raw }
func (r *MonitorListResponseDataChangeDetectionExact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detect meaning-level changes to tracked page content, ignoring cosmetic or
// paraphrase-only differences. Which changes are meaningful is judged against the
// extract target's `instructions` (and `schema`, when provided).
type MonitorListResponseDataChangeDetectionSemantic struct {
	Type                constant.Semantic `json:"type" default:"semantic"`
	ConfidenceThreshold float64           `json:"confidence_threshold"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type                respjson.Field
		ConfidenceThreshold respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListResponseDataChangeDetectionSemantic) RawJSON() string { return r.JSON.raw }
func (r *MonitorListResponseDataChangeDetectionSemantic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
// every 6 hours or every 2 days. The total interval (frequency × unit) must be
// between 10 minutes and 1 year.
type MonitorListResponseDataSchedule struct {
	// Number of units between runs. The resulting interval (frequency × unit) must be
	// at least 10 minutes and at most 1 year (e.g. minimum 10 when unit is minutes;
	// maximum 365 when unit is days).
	Frequency int64 `json:"frequency" api:"required"`
	// Any of "interval".
	Type string `json:"type" api:"required"`
	// Any of "minutes", "hours", "days".
	Unit string `json:"unit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Frequency   respjson.Field
		Type        respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListResponseDataSchedule) RawJSON() string { return r.JSON.raw }
func (r *MonitorListResponseDataSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MonitorListResponseDataTargetUnion contains all possible properties and values
// from [MonitorListResponseDataTargetPage],
// [MonitorListResponseDataTargetSitemap], [MonitorListResponseDataTargetExtract].
//
// Use the [MonitorListResponseDataTargetUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MonitorListResponseDataTargetUnion struct {
	// Any of "page", "sitemap", "extract".
	Type string `json:"type"`
	URL  string `json:"url"`
	// This field is from variant [MonitorListResponseDataTargetPage].
	NormalizeWhitespace bool `json:"normalize_whitespace"`
	// This field is from variant [MonitorListResponseDataTargetSitemap].
	Exclude []string `json:"exclude"`
	// This field is from variant [MonitorListResponseDataTargetSitemap].
	Include []string `json:"include"`
	// This field is from variant [MonitorListResponseDataTargetSitemap].
	MaxURLs int64 `json:"max_urls"`
	// This field is from variant [MonitorListResponseDataTargetExtract].
	Instructions string `json:"instructions"`
	// This field is from variant [MonitorListResponseDataTargetExtract].
	FollowSubdomains bool `json:"follow_subdomains"`
	// This field is from variant [MonitorListResponseDataTargetExtract].
	MaxDepth int64 `json:"max_depth"`
	// This field is from variant [MonitorListResponseDataTargetExtract].
	MaxPages int64 `json:"max_pages"`
	// This field is from variant [MonitorListResponseDataTargetExtract].
	Schema map[string]any `json:"schema"`
	JSON   struct {
		Type                respjson.Field
		URL                 respjson.Field
		NormalizeWhitespace respjson.Field
		Exclude             respjson.Field
		Include             respjson.Field
		MaxURLs             respjson.Field
		Instructions        respjson.Field
		FollowSubdomains    respjson.Field
		MaxDepth            respjson.Field
		MaxPages            respjson.Field
		Schema              respjson.Field
		raw                 string
	} `json:"-"`
}

// anyMonitorListResponseDataTarget is implemented by each variant of
// [MonitorListResponseDataTargetUnion] to add type safety for the return type of
// [MonitorListResponseDataTargetUnion.AsAny]
type anyMonitorListResponseDataTarget interface {
	implMonitorListResponseDataTargetUnion()
}

func (MonitorListResponseDataTargetPage) implMonitorListResponseDataTargetUnion()    {}
func (MonitorListResponseDataTargetSitemap) implMonitorListResponseDataTargetUnion() {}
func (MonitorListResponseDataTargetExtract) implMonitorListResponseDataTargetUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := MonitorListResponseDataTargetUnion.AsAny().(type) {
//	case contextdev.MonitorListResponseDataTargetPage:
//	case contextdev.MonitorListResponseDataTargetSitemap:
//	case contextdev.MonitorListResponseDataTargetExtract:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u MonitorListResponseDataTargetUnion) AsAny() anyMonitorListResponseDataTarget {
	switch u.Type {
	case "page":
		return u.AsPage()
	case "sitemap":
		return u.AsSitemap()
	case "extract":
		return u.AsExtract()
	}
	return nil
}

func (u MonitorListResponseDataTargetUnion) AsPage() (v MonitorListResponseDataTargetPage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorListResponseDataTargetUnion) AsSitemap() (v MonitorListResponseDataTargetSitemap) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorListResponseDataTargetUnion) AsExtract() (v MonitorListResponseDataTargetExtract) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MonitorListResponseDataTargetUnion) RawJSON() string { return u.JSON.raw }

func (r *MonitorListResponseDataTargetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Watch a single web page.
type MonitorListResponseDataTargetPage struct {
	Type constant.Page `json:"type" default:"page"`
	URL  string        `json:"url" api:"required" format:"uri"`
	// Normalize whitespace before comparing or analyzing text.
	NormalizeWhitespace bool `json:"normalize_whitespace"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type                respjson.Field
		URL                 respjson.Field
		NormalizeWhitespace respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListResponseDataTargetPage) RawJSON() string { return r.JSON.raw }
func (r *MonitorListResponseDataTargetPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Watch a sitemap for URL additions and removals. Crawled URLs are normalized
// (lowercased host, no trailing slash/fragment) and scoped to the monitored site
// and its subdomains before comparison. On a detected difference the sitemap is
// re-fetched within the same run and only URLs both observations agree on are
// reported, suppressing transient crawl flaps.
type MonitorListResponseDataTargetSitemap struct {
	Type constant.Sitemap `json:"type" default:"sitemap"`
	// Sitemap URL to monitor.
	URL string `json:"url" api:"required" format:"uri"`
	// URL path patterns to exclude.
	Exclude []string `json:"exclude"`
	// URL path patterns to include.
	Include []string `json:"include"`
	// Maximum number of sitemap URLs to track (capped at 10,000).
	MaxURLs int64 `json:"max_urls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URL         respjson.Field
		Exclude     respjson.Field
		Include     respjson.Field
		MaxURLs     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListResponseDataTargetSitemap) RawJSON() string { return r.JSON.raw }
func (r *MonitorListResponseDataTargetSitemap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Watch the monitor-relevant pages of a site for meaningful changes. A crawl
// guided by `schema`/`instructions` selects up to `max_pages` relevant pages to
// track; each run re-checks exactly those pages, and confirmed content changes are
// judged for relevance against the monitor's `instructions` (and `schema`, when
// provided). The tracked page set is refreshed by a periodic re-discovery crawl.
type MonitorListResponseDataTargetExtract struct {
	// Natural-language instructions guiding which pages and facts to track and which
	// changes to report.
	Instructions string           `json:"instructions" api:"required"`
	Type         constant.Extract `json:"type" default:"extract"`
	// Root URL to extract structured data from.
	URL              string `json:"url" api:"required" format:"uri"`
	FollowSubdomains bool   `json:"follow_subdomains"`
	// Optional maximum link depth from the starting URL (0 = only the starting page).
	MaxDepth int64 `json:"max_depth"`
	// Maximum number of pages to track.
	MaxPages int64 `json:"max_pages"`
	// JSON Schema describing the data you care about. It is used three ways: it guides
	// which pages are selected for tracking, it gives the change judge extra context
	// on which changes matter (alongside `instructions`), and it defines the shape of
	// the baseline `data` snapshot on GET /monitors/{monitor_id} (refreshed at most
	// about once a day). It is not a response format for changes: change events and
	// webhook payloads always contain diffs, summaries, and evidence excerpts — never
	// data in this schema's shape. If omitted, a default summary + key-points schema
	// is used.
	Schema map[string]any `json:"schema"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Instructions     respjson.Field
		Type             respjson.Field
		URL              respjson.Field
		FollowSubdomains respjson.Field
		MaxDepth         respjson.Field
		MaxPages         respjson.Field
		Schema           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListResponseDataTargetExtract) RawJSON() string { return r.JSON.raw }
func (r *MonitorListResponseDataTargetExtract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MonitorListResponseDataBaselineUnion contains all possible properties and values
// from [MonitorListResponseDataBaselinePageBaseline],
// [MonitorListResponseDataBaselineSitemapBaseline],
// [MonitorListResponseDataBaselineExtractBaseline].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MonitorListResponseDataBaselineUnion struct {
	CapturedAt time.Time `json:"captured_at"`
	// This field is from variant [MonitorListResponseDataBaselinePageBaseline].
	Text string `json:"text"`
	// This field is from variant [MonitorListResponseDataBaselineSitemapBaseline].
	URLCount int64 `json:"url_count"`
	// This field is from variant [MonitorListResponseDataBaselineSitemapBaseline].
	URLs []string `json:"urls"`
	// This field is from variant [MonitorListResponseDataBaselineExtractBaseline].
	Data any `json:"data"`
	// This field is from variant [MonitorListResponseDataBaselineExtractBaseline].
	URLsAnalyzed []string `json:"urls_analyzed"`
	JSON         struct {
		CapturedAt   respjson.Field
		Text         respjson.Field
		URLCount     respjson.Field
		URLs         respjson.Field
		Data         respjson.Field
		URLsAnalyzed respjson.Field
		raw          string
	} `json:"-"`
}

func (u MonitorListResponseDataBaselineUnion) AsPageBaseline() (v MonitorListResponseDataBaselinePageBaseline) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorListResponseDataBaselineUnion) AsSitemapBaseline() (v MonitorListResponseDataBaselineSitemapBaseline) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorListResponseDataBaselineUnion) AsExtractBaseline() (v MonitorListResponseDataBaselineExtractBaseline) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MonitorListResponseDataBaselineUnion) RawJSON() string { return u.JSON.raw }

func (r *MonitorListResponseDataBaselineUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current baseline of a `page` monitor: the visible page text as last observed.
type MonitorListResponseDataBaselinePageBaseline struct {
	// When this baseline was last captured or replaced.
	CapturedAt time.Time `json:"captured_at" api:"required" format:"date-time"`
	// The page's visible text as last observed.
	Text string `json:"text" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CapturedAt  respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListResponseDataBaselinePageBaseline) RawJSON() string { return r.JSON.raw }
func (r *MonitorListResponseDataBaselinePageBaseline) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current baseline of a `sitemap` monitor: the normalized URL set as last
// observed.
type MonitorListResponseDataBaselineSitemapBaseline struct {
	// When this baseline was last captured or replaced.
	CapturedAt time.Time `json:"captured_at" api:"required" format:"date-time"`
	// Number of URLs in the baseline.
	URLCount int64 `json:"url_count" api:"required"`
	// The sitemap URLs as last observed (sorted, normalized).
	URLs []string `json:"urls" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CapturedAt  respjson.Field
		URLCount    respjson.Field
		URLs        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListResponseDataBaselineSitemapBaseline) RawJSON() string { return r.JSON.raw }
func (r *MonitorListResponseDataBaselineSitemapBaseline) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current baseline of an `extract` monitor: the pages it tracks and the structured
// data as last extracted.
type MonitorListResponseDataBaselineExtractBaseline struct {
	// When this baseline was last captured or replaced.
	CapturedAt time.Time `json:"captured_at" api:"required" format:"date-time"`
	// The extracted structured data, matching the monitor's extraction schema (same
	// shape as the /web/extract endpoint's `data`). Refreshed when the monitor
	// re-discovers its page set (at most about once a day); `null` when no extraction
	// has been captured yet.
	Data any `json:"data" api:"required"`
	// The page URLs the monitor tracks and analyzes for changes.
	URLsAnalyzed []string `json:"urls_analyzed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CapturedAt   respjson.Field
		Data         respjson.Field
		URLsAnalyzed respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListResponseDataBaselineExtractBaseline) RawJSON() string { return r.JSON.raw }
func (r *MonitorListResponseDataBaselineExtractBaseline) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error from the most recent failed run; null when the last run succeeded.
type MonitorListResponseDataLastError struct {
	Code    string `json:"code" api:"required"`
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
func (r MonitorListResponseDataLastError) RawJSON() string { return r.JSON.raw }
func (r *MonitorListResponseDataLastError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListResponseDataWebhook struct {
	// Webhook URL events are delivered to.
	URL string `json:"url" api:"required" format:"uri"`
	// Events delivered to this endpoint. `change.detected` fires only when a run
	// detects a change; `run.completed` fires on every completed run — including runs
	// that detected no change — and embeds the change when one was detected. Defaults
	// to `["change.detected"]` when omitted.
	//
	// Any of "change.detected", "run.completed".
	Events []string `json:"events"`
	// Signing secret used to verify webhook authenticity. Each delivery includes an
	// `X-Context-Signature: t=<unix>,v1=<hmac>` header, where the HMAC is SHA-256 over
	// `"{t}.{rawRequestBody}"` keyed by this secret. Recompute it with a constant-time
	// compare and reject stale timestamps to prevent replay. Generated by the API;
	// cannot be set by clients.
	Secret string `json:"secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Events      respjson.Field
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListResponseDataWebhook) RawJSON() string { return r.JSON.raw }
func (r *MonitorListResponseDataWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Present while webhook deliveries are failing consecutively; null when deliveries
// are healthy or no webhook is configured. Cleared on the next successful delivery
// and when the webhook URL changes.
type MonitorListResponseDataWebhookFailure struct {
	// Number of consecutive delivery attempts that did not succeed.
	ConsecutiveFailures int64     `json:"consecutive_failures" api:"required"`
	LastFailedAt        time.Time `json:"last_failed_at" api:"required" format:"date-time"`
	// Human-readable description of the most recent failure.
	LastMessage string `json:"last_message" api:"required"`
	// Outcome of the most recent failed delivery. rejected means a non-2xx response;
	// failed means no HTTP response was received; skipped_unsafe_url means the URL
	// failed the public-endpoint safety check.
	//
	// Any of "rejected", "failed", "skipped_unsafe_url".
	LastStatus string `json:"last_status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConsecutiveFailures respjson.Field
		LastFailedAt        respjson.Field
		LastMessage         respjson.Field
		LastStatus          respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListResponseDataWebhookFailure) RawJSON() string { return r.JSON.raw }
func (r *MonitorListResponseDataWebhookFailure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorDeleteResponse struct {
	ID      string `json:"id" api:"required"`
	Deleted bool   `json:"deleted" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Deleted     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorGetCreditUsageResponse struct {
	Data []MonitorGetCreditUsageResponseData `json:"data" api:"required"`
	// Sum of credits across all monitors in the window.
	TotalCredits int64 `json:"total_credits" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data         respjson.Field
		TotalCredits respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetCreditUsageResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetCreditUsageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorGetCreditUsageResponseData struct {
	// Credits charged to this monitor over the window.
	Credits   int64  `json:"credits" api:"required"`
	MonitorID string `json:"monitor_id" api:"required"`
	// Monitor name (falls back to the id when the monitor was deleted).
	Name string `json:"name" api:"required"`
	// Number of billed runs over the window.
	Runs int64 `json:"runs" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Credits     respjson.Field
		MonitorID   respjson.Field
		Name        respjson.Field
		Runs        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetCreditUsageResponseData) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetCreditUsageResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorGetLimitsResponse struct {
	// Maximum number of monitors allowed for the account. Defaults to the plan
	// allowance unless a custom limit is set for the organization.
	MonitorsLimit int64 `json:"monitors_limit" api:"required"`
	// Number of monitors the account currently has.
	MonitorsUsed int64 `json:"monitors_used" api:"required"`
	// The plan tier the limit was resolved from.
	//
	// Any of "free", "starter", "pro", "scale".
	Plan MonitorGetLimitsResponsePlan `json:"plan" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MonitorsLimit respjson.Field
		MonitorsUsed  respjson.Field
		Plan          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetLimitsResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetLimitsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The plan tier the limit was resolved from.
type MonitorGetLimitsResponsePlan string

const (
	MonitorGetLimitsResponsePlanFree    MonitorGetLimitsResponsePlan = "free"
	MonitorGetLimitsResponsePlanStarter MonitorGetLimitsResponsePlan = "starter"
	MonitorGetLimitsResponsePlanPro     MonitorGetLimitsResponsePlan = "pro"
	MonitorGetLimitsResponsePlanScale   MonitorGetLimitsResponsePlan = "scale"
)

type MonitorListAccountChangesResponse struct {
	Data       []MonitorListAccountChangesResponseData `json:"data" api:"required"`
	HasMore    bool                                    `json:"has_more" api:"required"`
	NextCursor string                                  `json:"next_cursor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListAccountChangesResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorListAccountChangesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A lightweight change summary. `mode` is the constant `web`; `target_type` and
// `change_detection_type` describe the change, and which optional fields are
// present depends on them (e.g. sitemap changes include
// `added_url_count`/`removed_url_count`; semantic changes include
// `confidence`/`importance`).
type MonitorListAccountChangesResponseData struct {
	ID string `json:"id" api:"required"`
	// Any of "exact", "semantic".
	ChangeDetectionType string    `json:"change_detection_type" api:"required"`
	DetectedAt          time.Time `json:"detected_at" api:"required" format:"date-time"`
	// Top-level monitor category. Always `web` today; the concrete behavior is
	// described by `target` and `change_detection`.
	//
	// Any of "web".
	Mode      string `json:"mode" api:"required"`
	MonitorID string `json:"monitor_id" api:"required"`
	Summary   string `json:"summary" api:"required"`
	// Any of "page", "sitemap", "extract".
	TargetType    string  `json:"target_type" api:"required"`
	Title         string  `json:"title" api:"required"`
	URL           string  `json:"url" api:"required" format:"uri"`
	AddedURLCount int64   `json:"added_url_count"`
	Confidence    float64 `json:"confidence"`
	// Any of "low", "medium", "high".
	Importance      string `json:"importance"`
	MatchedURLCount int64  `json:"matched_url_count"`
	RemovedURLCount int64  `json:"removed_url_count"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags []string `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		ChangeDetectionType respjson.Field
		DetectedAt          respjson.Field
		Mode                respjson.Field
		MonitorID           respjson.Field
		Summary             respjson.Field
		TargetType          respjson.Field
		Title               respjson.Field
		URL                 respjson.Field
		AddedURLCount       respjson.Field
		Confidence          respjson.Field
		Importance          respjson.Field
		MatchedURLCount     respjson.Field
		RemovedURLCount     respjson.Field
		Tags                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListAccountChangesResponseData) RawJSON() string { return r.JSON.raw }
func (r *MonitorListAccountChangesResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListAccountRunsResponse struct {
	Data       []MonitorListAccountRunsResponseData `json:"data" api:"required"`
	HasMore    bool                                 `json:"has_more" api:"required"`
	NextCursor string                               `json:"next_cursor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListAccountRunsResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorListAccountRunsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListAccountRunsResponseData struct {
	ID string `json:"id" api:"required"`
	// True when this run established the monitor's initial baseline; baseline runs
	// perform no change detection.
	BaselineCreated bool `json:"baseline_created" api:"required"`
	ChangeDetected  bool `json:"change_detected" api:"required"`
	// Any of "exact", "semantic".
	ChangeDetectionType string `json:"change_detection_type" api:"required"`
	// Credits charged for this run (0 for skipped/failed runs).
	CreditsCharged int64  `json:"credits_charged" api:"required"`
	MonitorID      string `json:"monitor_id" api:"required"`
	// The first run after monitor creation is a baseline run.
	//
	// Any of "baseline", "scheduled".
	RunType string `json:"run_type" api:"required"`
	// Lifecycle status of a run. `skipped` runs never executed — see `skip_reason`
	// (insufficient credits, monitor paused, or superseded by a concurrent run).
	//
	// Any of "queued", "running", "completed", "failed", "skipped".
	Status string `json:"status" api:"required"`
	// Any of "page", "sitemap", "extract".
	TargetType  string                                  `json:"target_type" api:"required"`
	ChangeID    string                                  `json:"change_id" api:"nullable"`
	CompletedAt time.Time                               `json:"completed_at" api:"nullable" format:"date-time"`
	Error       MonitorListAccountRunsResponseDataError `json:"error" api:"nullable"`
	// Why a skipped run never executed; null unless status is `skipped`.
	//
	// Any of "insufficient_credits", "monitor_paused", "superseded".
	SkipReason string    `json:"skip_reason" api:"nullable"`
	StartedAt  time.Time `json:"started_at" api:"nullable" format:"date-time"`
	// All webhook deliveries attempted by this run — one per subscribed event that
	// fired. Omitted when no webhook was attempted, including runs created before
	// event selection was added.
	WebhookDeliveries []MonitorListAccountRunsResponseDataWebhookDelivery `json:"webhook_deliveries"`
	// Deprecated: use `webhook_deliveries`, which records every attempt now that a run
	// can deliver multiple events. Omitted when no webhook was attempted, including
	// historical runs created before delivery tracking was added.
	//
	// Deprecated: deprecated
	WebhookDelivery MonitorListAccountRunsResponseDataWebhookDelivery `json:"webhook_delivery"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		BaselineCreated     respjson.Field
		ChangeDetected      respjson.Field
		ChangeDetectionType respjson.Field
		CreditsCharged      respjson.Field
		MonitorID           respjson.Field
		RunType             respjson.Field
		Status              respjson.Field
		TargetType          respjson.Field
		ChangeID            respjson.Field
		CompletedAt         respjson.Field
		Error               respjson.Field
		SkipReason          respjson.Field
		StartedAt           respjson.Field
		WebhookDeliveries   respjson.Field
		WebhookDelivery     respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListAccountRunsResponseData) RawJSON() string { return r.JSON.raw }
func (r *MonitorListAccountRunsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListAccountRunsResponseDataError struct {
	Code    string `json:"code" api:"required"`
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
func (r MonitorListAccountRunsResponseDataError) RawJSON() string { return r.JSON.raw }
func (r *MonitorListAccountRunsResponseDataError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListAccountRunsResponseDataWebhookDelivery struct {
	AttemptedAt time.Time                                              `json:"attempted_at" api:"required" format:"date-time"`
	Error       MonitorListAccountRunsResponseDataWebhookDeliveryError `json:"error" api:"required"`
	// The event this delivery carried. Deliveries recorded before event selection
	// existed report change.detected.
	//
	// Any of "change.detected", "run.completed".
	Event string `json:"event" api:"required"`
	// Identifier sent in the X-Context-Id header.
	EventID string `json:"event_id" api:"required"`
	// The endpoint's final HTTP response status, or null when no response was
	// received.
	HTTPStatus int64 `json:"http_status" api:"required"`
	// Delivery outcome. delivered means any 2xx response; rejected means a non-2xx
	// response; failed means no HTTP response was received; skipped_unsafe_url means
	// the URL failed the public-endpoint safety check.
	//
	// Any of "delivered", "rejected", "failed", "skipped_unsafe_url".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttemptedAt respjson.Field
		Error       respjson.Field
		Event       respjson.Field
		EventID     respjson.Field
		HTTPStatus  respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListAccountRunsResponseDataWebhookDelivery) RawJSON() string { return r.JSON.raw }
func (r *MonitorListAccountRunsResponseDataWebhookDelivery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListAccountRunsResponseDataWebhookDeliveryError struct {
	Code    string `json:"code" api:"required"`
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
func (r MonitorListAccountRunsResponseDataWebhookDeliveryError) RawJSON() string { return r.JSON.raw }
func (r *MonitorListAccountRunsResponseDataWebhookDeliveryError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListChangesResponse struct {
	Data       []MonitorListChangesResponseData `json:"data" api:"required"`
	HasMore    bool                             `json:"has_more" api:"required"`
	NextCursor string                           `json:"next_cursor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListChangesResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorListChangesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A lightweight change summary. `mode` is the constant `web`; `target_type` and
// `change_detection_type` describe the change, and which optional fields are
// present depends on them (e.g. sitemap changes include
// `added_url_count`/`removed_url_count`; semantic changes include
// `confidence`/`importance`).
type MonitorListChangesResponseData struct {
	ID string `json:"id" api:"required"`
	// Any of "exact", "semantic".
	ChangeDetectionType string    `json:"change_detection_type" api:"required"`
	DetectedAt          time.Time `json:"detected_at" api:"required" format:"date-time"`
	// Top-level monitor category. Always `web` today; the concrete behavior is
	// described by `target` and `change_detection`.
	//
	// Any of "web".
	Mode      string `json:"mode" api:"required"`
	MonitorID string `json:"monitor_id" api:"required"`
	Summary   string `json:"summary" api:"required"`
	// Any of "page", "sitemap", "extract".
	TargetType    string  `json:"target_type" api:"required"`
	Title         string  `json:"title" api:"required"`
	URL           string  `json:"url" api:"required" format:"uri"`
	AddedURLCount int64   `json:"added_url_count"`
	Confidence    float64 `json:"confidence"`
	// Any of "low", "medium", "high".
	Importance      string `json:"importance"`
	MatchedURLCount int64  `json:"matched_url_count"`
	RemovedURLCount int64  `json:"removed_url_count"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags []string `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		ChangeDetectionType respjson.Field
		DetectedAt          respjson.Field
		Mode                respjson.Field
		MonitorID           respjson.Field
		Summary             respjson.Field
		TargetType          respjson.Field
		Title               respjson.Field
		URL                 respjson.Field
		AddedURLCount       respjson.Field
		Confidence          respjson.Field
		Importance          respjson.Field
		MatchedURLCount     respjson.Field
		RemovedURLCount     respjson.Field
		Tags                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListChangesResponseData) RawJSON() string { return r.JSON.raw }
func (r *MonitorListChangesResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListRunsResponse struct {
	Data       []MonitorListRunsResponseData `json:"data" api:"required"`
	HasMore    bool                          `json:"has_more" api:"required"`
	NextCursor string                        `json:"next_cursor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListRunsResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorListRunsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListRunsResponseData struct {
	ID string `json:"id" api:"required"`
	// True when this run established the monitor's initial baseline; baseline runs
	// perform no change detection.
	BaselineCreated bool `json:"baseline_created" api:"required"`
	ChangeDetected  bool `json:"change_detected" api:"required"`
	// Any of "exact", "semantic".
	ChangeDetectionType string `json:"change_detection_type" api:"required"`
	// Credits charged for this run (0 for skipped/failed runs).
	CreditsCharged int64  `json:"credits_charged" api:"required"`
	MonitorID      string `json:"monitor_id" api:"required"`
	// The first run after monitor creation is a baseline run.
	//
	// Any of "baseline", "scheduled".
	RunType string `json:"run_type" api:"required"`
	// Lifecycle status of a run. `skipped` runs never executed — see `skip_reason`
	// (insufficient credits, monitor paused, or superseded by a concurrent run).
	//
	// Any of "queued", "running", "completed", "failed", "skipped".
	Status string `json:"status" api:"required"`
	// Any of "page", "sitemap", "extract".
	TargetType  string                           `json:"target_type" api:"required"`
	ChangeID    string                           `json:"change_id" api:"nullable"`
	CompletedAt time.Time                        `json:"completed_at" api:"nullable" format:"date-time"`
	Error       MonitorListRunsResponseDataError `json:"error" api:"nullable"`
	// Why a skipped run never executed; null unless status is `skipped`.
	//
	// Any of "insufficient_credits", "monitor_paused", "superseded".
	SkipReason string    `json:"skip_reason" api:"nullable"`
	StartedAt  time.Time `json:"started_at" api:"nullable" format:"date-time"`
	// All webhook deliveries attempted by this run — one per subscribed event that
	// fired. Omitted when no webhook was attempted, including runs created before
	// event selection was added.
	WebhookDeliveries []MonitorListRunsResponseDataWebhookDelivery `json:"webhook_deliveries"`
	// Deprecated: use `webhook_deliveries`, which records every attempt now that a run
	// can deliver multiple events. Omitted when no webhook was attempted, including
	// historical runs created before delivery tracking was added.
	//
	// Deprecated: deprecated
	WebhookDelivery MonitorListRunsResponseDataWebhookDelivery `json:"webhook_delivery"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		BaselineCreated     respjson.Field
		ChangeDetected      respjson.Field
		ChangeDetectionType respjson.Field
		CreditsCharged      respjson.Field
		MonitorID           respjson.Field
		RunType             respjson.Field
		Status              respjson.Field
		TargetType          respjson.Field
		ChangeID            respjson.Field
		CompletedAt         respjson.Field
		Error               respjson.Field
		SkipReason          respjson.Field
		StartedAt           respjson.Field
		WebhookDeliveries   respjson.Field
		WebhookDelivery     respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListRunsResponseData) RawJSON() string { return r.JSON.raw }
func (r *MonitorListRunsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListRunsResponseDataError struct {
	Code    string `json:"code" api:"required"`
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
func (r MonitorListRunsResponseDataError) RawJSON() string { return r.JSON.raw }
func (r *MonitorListRunsResponseDataError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListRunsResponseDataWebhookDelivery struct {
	AttemptedAt time.Time                                       `json:"attempted_at" api:"required" format:"date-time"`
	Error       MonitorListRunsResponseDataWebhookDeliveryError `json:"error" api:"required"`
	// The event this delivery carried. Deliveries recorded before event selection
	// existed report change.detected.
	//
	// Any of "change.detected", "run.completed".
	Event string `json:"event" api:"required"`
	// Identifier sent in the X-Context-Id header.
	EventID string `json:"event_id" api:"required"`
	// The endpoint's final HTTP response status, or null when no response was
	// received.
	HTTPStatus int64 `json:"http_status" api:"required"`
	// Delivery outcome. delivered means any 2xx response; rejected means a non-2xx
	// response; failed means no HTTP response was received; skipped_unsafe_url means
	// the URL failed the public-endpoint safety check.
	//
	// Any of "delivered", "rejected", "failed", "skipped_unsafe_url".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttemptedAt respjson.Field
		Error       respjson.Field
		Event       respjson.Field
		EventID     respjson.Field
		HTTPStatus  respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListRunsResponseDataWebhookDelivery) RawJSON() string { return r.JSON.raw }
func (r *MonitorListRunsResponseDataWebhookDelivery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListRunsResponseDataWebhookDeliveryError struct {
	Code    string `json:"code" api:"required"`
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
func (r MonitorListRunsResponseDataWebhookDeliveryError) RawJSON() string { return r.JSON.raw }
func (r *MonitorListRunsResponseDataWebhookDeliveryError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A detected change. `mode` is the constant `web`; `target_type` and
// `change_detection_type` describe the change, and which optional fields are
// present depends on them (page: `diff` + excerpts; sitemap:
// `added_urls`/`removed_urls`; semantic:
// `confidence`/`importance`/`evidence`/`matched_urls`).
type MonitorGetChangeResponse struct {
	ID string `json:"id" api:"required"`
	// Any of "exact", "semantic".
	ChangeDetectionType MonitorGetChangeResponseChangeDetectionType `json:"change_detection_type" api:"required"`
	DetectedAt          time.Time                                   `json:"detected_at" api:"required" format:"date-time"`
	// Top-level monitor category. Always `web` today; the concrete behavior is
	// described by `target` and `change_detection`.
	//
	// Any of "web".
	Mode      MonitorGetChangeResponseMode `json:"mode" api:"required"`
	MonitorID string                       `json:"monitor_id" api:"required"`
	// The run that detected this change.
	RunID   string `json:"run_id" api:"required"`
	Summary string `json:"summary" api:"required"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags []string `json:"tags" api:"required"`
	// Any of "page", "sitemap", "extract".
	TargetType    MonitorGetChangeResponseTargetType `json:"target_type" api:"required"`
	Title         string                             `json:"title" api:"required"`
	URL           string                             `json:"url" api:"required" format:"uri"`
	AddedURLCount int64                              `json:"added_url_count"`
	// At most 500 URLs are included; the corresponding count field is always exact.
	AddedURLs         []string `json:"added_urls" format:"uri"`
	AfterTextExcerpt  string   `json:"after_text_excerpt"`
	BeforeTextExcerpt string   `json:"before_text_excerpt"`
	Confidence        float64  `json:"confidence"`
	// Text diff between the previous and current page baseline (page targets).
	Diff     string                             `json:"diff"`
	Evidence []MonitorGetChangeResponseEvidence `json:"evidence"`
	// Any of "low", "medium", "high".
	Importance      MonitorGetChangeResponseImportance `json:"importance"`
	MatchedURLCount int64                              `json:"matched_url_count"`
	// At most 500 URLs are included; the corresponding count field is always exact.
	MatchedURLs     []string `json:"matched_urls" format:"uri"`
	RemovedURLCount int64    `json:"removed_url_count"`
	// At most 500 URLs are included; the corresponding count field is always exact.
	RemovedURLs []string `json:"removed_urls" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		ChangeDetectionType respjson.Field
		DetectedAt          respjson.Field
		Mode                respjson.Field
		MonitorID           respjson.Field
		RunID               respjson.Field
		Summary             respjson.Field
		Tags                respjson.Field
		TargetType          respjson.Field
		Title               respjson.Field
		URL                 respjson.Field
		AddedURLCount       respjson.Field
		AddedURLs           respjson.Field
		AfterTextExcerpt    respjson.Field
		BeforeTextExcerpt   respjson.Field
		Confidence          respjson.Field
		Diff                respjson.Field
		Evidence            respjson.Field
		Importance          respjson.Field
		MatchedURLCount     respjson.Field
		MatchedURLs         respjson.Field
		RemovedURLCount     respjson.Field
		RemovedURLs         respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetChangeResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetChangeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorGetChangeResponseChangeDetectionType string

const (
	MonitorGetChangeResponseChangeDetectionTypeExact    MonitorGetChangeResponseChangeDetectionType = "exact"
	MonitorGetChangeResponseChangeDetectionTypeSemantic MonitorGetChangeResponseChangeDetectionType = "semantic"
)

// Top-level monitor category. Always `web` today; the concrete behavior is
// described by `target` and `change_detection`.
type MonitorGetChangeResponseMode string

const (
	MonitorGetChangeResponseModeWeb MonitorGetChangeResponseMode = "web"
)

type MonitorGetChangeResponseTargetType string

const (
	MonitorGetChangeResponseTargetTypePage    MonitorGetChangeResponseTargetType = "page"
	MonitorGetChangeResponseTargetTypeSitemap MonitorGetChangeResponseTargetType = "sitemap"
	MonitorGetChangeResponseTargetTypeExtract MonitorGetChangeResponseTargetType = "extract"
)

type MonitorGetChangeResponseEvidence struct {
	// Snapshot of the content after the change.
	After string `json:"after" api:"required"`
	// Snapshot of the content before the change.
	Before string `json:"before" api:"required"`
	// Optional URL the evidence relates to. Absent for whole-target diffs.
	URL string `json:"url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		After       respjson.Field
		Before      respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetChangeResponseEvidence) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetChangeResponseEvidence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorGetChangeResponseImportance string

const (
	MonitorGetChangeResponseImportanceLow    MonitorGetChangeResponseImportance = "low"
	MonitorGetChangeResponseImportanceMedium MonitorGetChangeResponseImportance = "medium"
	MonitorGetChangeResponseImportanceHigh   MonitorGetChangeResponseImportance = "high"
)

type MonitorRunResponse struct {
	MonitorID string `json:"monitor_id" api:"required"`
	Queued    bool   `json:"queued" api:"required"`
	// The queued run. Poll GET /monitors/{monitor_id}/runs or use it to correlate
	// results.
	RunID string `json:"run_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MonitorID   respjson.Field
		Queued      respjson.Field
		RunID       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorRunResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorNewParams struct {
	// Discriminated union describing how changes are detected.
	ChangeDetection MonitorNewParamsChangeDetectionUnion `json:"change_detection,omitzero" api:"required"`
	Name            string                               `json:"name" api:"required"`
	// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
	// every 6 hours or every 2 days. The total interval (frequency × unit) must be
	// between 10 minutes and 1 year.
	Schedule MonitorNewParamsSchedule `json:"schedule,omitzero" api:"required"`
	// Discriminated union describing what the monitor watches.
	Target  MonitorNewParamsTargetUnion `json:"target,omitzero" api:"required"`
	Webhook MonitorNewParamsWebhook     `json:"webhook,omitzero"`
	// Top-level monitor category. Always `web` today; the concrete behavior is
	// described by `target` and `change_detection`.
	//
	// Any of "web".
	Mode MonitorNewParamsMode `json:"mode,omitzero"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r MonitorNewParams) MarshalJSON() (data []byte, err error) {
	type shadow MonitorNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MonitorNewParamsChangeDetectionUnion struct {
	OfExact    *MonitorNewParamsChangeDetectionExact    `json:",omitzero,inline"`
	OfSemantic *MonitorNewParamsChangeDetectionSemantic `json:",omitzero,inline"`
	paramUnion
}

func (u MonitorNewParamsChangeDetectionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExact, u.OfSemantic)
}
func (u *MonitorNewParamsChangeDetectionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[MonitorNewParamsChangeDetectionUnion](
		"type",
		apijson.Discriminator[MonitorNewParamsChangeDetectionExact]("exact"),
		apijson.Discriminator[MonitorNewParamsChangeDetectionSemantic]("semantic"),
	)
}

func NewMonitorNewParamsChangeDetectionExact() MonitorNewParamsChangeDetectionExact {
	return MonitorNewParamsChangeDetectionExact{
		Type: "exact",
	}
}

// Detect exact changes. For page targets, this means visible text diffs. For
// sitemap targets, this means URL additions and removals.
//
// This struct has a constant value, construct it with
// [NewMonitorNewParamsChangeDetectionExact].
type MonitorNewParamsChangeDetectionExact struct {
	Type constant.Exact `json:"type" default:"exact"`
	paramObj
}

func (r MonitorNewParamsChangeDetectionExact) MarshalJSON() (data []byte, err error) {
	type shadow MonitorNewParamsChangeDetectionExact
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorNewParamsChangeDetectionExact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detect meaning-level changes to tracked page content, ignoring cosmetic or
// paraphrase-only differences. Which changes are meaningful is judged against the
// extract target's `instructions` (and `schema`, when provided).
//
// The property Type is required.
type MonitorNewParamsChangeDetectionSemantic struct {
	ConfidenceThreshold param.Opt[float64] `json:"confidence_threshold,omitzero"`
	// This field can be elided, and will marshal its zero value as "semantic".
	Type constant.Semantic `json:"type" default:"semantic"`
	paramObj
}

func (r MonitorNewParamsChangeDetectionSemantic) MarshalJSON() (data []byte, err error) {
	type shadow MonitorNewParamsChangeDetectionSemantic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorNewParamsChangeDetectionSemantic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
// every 6 hours or every 2 days. The total interval (frequency × unit) must be
// between 10 minutes and 1 year.
//
// The properties Frequency, Type, Unit are required.
type MonitorNewParamsSchedule struct {
	// Number of units between runs. The resulting interval (frequency × unit) must be
	// at least 10 minutes and at most 1 year (e.g. minimum 10 when unit is minutes;
	// maximum 365 when unit is days).
	Frequency int64 `json:"frequency" api:"required"`
	// Any of "interval".
	Type string `json:"type,omitzero" api:"required"`
	// Any of "minutes", "hours", "days".
	Unit string `json:"unit,omitzero" api:"required"`
	paramObj
}

func (r MonitorNewParamsSchedule) MarshalJSON() (data []byte, err error) {
	type shadow MonitorNewParamsSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorNewParamsSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MonitorNewParamsSchedule](
		"type", "interval",
	)
	apijson.RegisterFieldValidator[MonitorNewParamsSchedule](
		"unit", "minutes", "hours", "days",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MonitorNewParamsTargetUnion struct {
	OfPage    *MonitorNewParamsTargetPage    `json:",omitzero,inline"`
	OfSitemap *MonitorNewParamsTargetSitemap `json:",omitzero,inline"`
	OfExtract *MonitorNewParamsTargetExtract `json:",omitzero,inline"`
	paramUnion
}

func (u MonitorNewParamsTargetUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPage, u.OfSitemap, u.OfExtract)
}
func (u *MonitorNewParamsTargetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[MonitorNewParamsTargetUnion](
		"type",
		apijson.Discriminator[MonitorNewParamsTargetPage]("page"),
		apijson.Discriminator[MonitorNewParamsTargetSitemap]("sitemap"),
		apijson.Discriminator[MonitorNewParamsTargetExtract]("extract"),
	)
}

// Watch a single web page.
//
// The properties Type, URL are required.
type MonitorNewParamsTargetPage struct {
	URL string `json:"url" api:"required" format:"uri"`
	// Normalize whitespace before comparing or analyzing text.
	NormalizeWhitespace param.Opt[bool] `json:"normalize_whitespace,omitzero"`
	// This field can be elided, and will marshal its zero value as "page".
	Type constant.Page `json:"type" default:"page"`
	paramObj
}

func (r MonitorNewParamsTargetPage) MarshalJSON() (data []byte, err error) {
	type shadow MonitorNewParamsTargetPage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorNewParamsTargetPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Watch a sitemap for URL additions and removals. Crawled URLs are normalized
// (lowercased host, no trailing slash/fragment) and scoped to the monitored site
// and its subdomains before comparison. On a detected difference the sitemap is
// re-fetched within the same run and only URLs both observations agree on are
// reported, suppressing transient crawl flaps.
//
// The properties Type, URL are required.
type MonitorNewParamsTargetSitemap struct {
	// Sitemap URL to monitor.
	URL string `json:"url" api:"required" format:"uri"`
	// Maximum number of sitemap URLs to track (capped at 10,000).
	MaxURLs param.Opt[int64] `json:"max_urls,omitzero"`
	// URL path patterns to exclude.
	Exclude []string `json:"exclude,omitzero"`
	// URL path patterns to include.
	Include []string `json:"include,omitzero"`
	// This field can be elided, and will marshal its zero value as "sitemap".
	Type constant.Sitemap `json:"type" default:"sitemap"`
	paramObj
}

func (r MonitorNewParamsTargetSitemap) MarshalJSON() (data []byte, err error) {
	type shadow MonitorNewParamsTargetSitemap
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorNewParamsTargetSitemap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Watch the monitor-relevant pages of a site for meaningful changes. A crawl
// guided by `schema`/`instructions` selects up to `max_pages` relevant pages to
// track; each run re-checks exactly those pages, and confirmed content changes are
// judged for relevance against the monitor's `instructions` (and `schema`, when
// provided). The tracked page set is refreshed by a periodic re-discovery crawl.
//
// The properties Instructions, Type, URL are required.
type MonitorNewParamsTargetExtract struct {
	// Natural-language instructions guiding which pages and facts to track and which
	// changes to report.
	Instructions string `json:"instructions" api:"required"`
	// Root URL to extract structured data from.
	URL              string          `json:"url" api:"required" format:"uri"`
	FollowSubdomains param.Opt[bool] `json:"follow_subdomains,omitzero"`
	// Optional maximum link depth from the starting URL (0 = only the starting page).
	MaxDepth param.Opt[int64] `json:"max_depth,omitzero"`
	// Maximum number of pages to track.
	MaxPages param.Opt[int64] `json:"max_pages,omitzero"`
	// JSON Schema describing the data you care about. It is used three ways: it guides
	// which pages are selected for tracking, it gives the change judge extra context
	// on which changes matter (alongside `instructions`), and it defines the shape of
	// the baseline `data` snapshot on GET /monitors/{monitor_id} (refreshed at most
	// about once a day). It is not a response format for changes: change events and
	// webhook payloads always contain diffs, summaries, and evidence excerpts — never
	// data in this schema's shape. If omitted, a default summary + key-points schema
	// is used.
	Schema map[string]any `json:"schema,omitzero"`
	// This field can be elided, and will marshal its zero value as "extract".
	Type constant.Extract `json:"type" default:"extract"`
	paramObj
}

func (r MonitorNewParamsTargetExtract) MarshalJSON() (data []byte, err error) {
	type shadow MonitorNewParamsTargetExtract
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorNewParamsTargetExtract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Top-level monitor category. Always `web` today; the concrete behavior is
// described by `target` and `change_detection`.
type MonitorNewParamsMode string

const (
	MonitorNewParamsModeWeb MonitorNewParamsMode = "web"
)

// The property URL is required.
type MonitorNewParamsWebhook struct {
	// Webhook URL events are delivered to.
	URL string `json:"url" api:"required" format:"uri"`
	// Events delivered to this endpoint. `change.detected` fires only when a run
	// detects a change; `run.completed` fires on every completed run — including runs
	// that detected no change — and embeds the change when one was detected. Defaults
	// to `["change.detected"]` when omitted.
	//
	// Any of "change.detected", "run.completed".
	Events []string `json:"events,omitzero"`
	paramObj
}

func (r MonitorNewParamsWebhook) MarshalJSON() (data []byte, err error) {
	type shadow MonitorNewParamsWebhook
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorNewParamsWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorUpdateParams struct {
	Name param.Opt[string] `json:"name,omitzero"`
	// Set to null to remove the webhook.
	Webhook MonitorUpdateParamsWebhook `json:"webhook,omitzero"`
	// Discriminated union describing how changes are detected.
	ChangeDetection MonitorUpdateParamsChangeDetectionUnion `json:"change_detection,omitzero"`
	// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
	// every 6 hours or every 2 days. The total interval (frequency × unit) must be
	// between 10 minutes and 1 year.
	Schedule MonitorUpdateParamsSchedule `json:"schedule,omitzero"`
	// Any of "active", "paused".
	Status MonitorUpdateParamsStatus `json:"status,omitzero"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags []string `json:"tags,omitzero"`
	// Discriminated union describing what the monitor watches.
	Target MonitorUpdateParamsTargetUnion `json:"target,omitzero"`
	paramObj
}

func (r MonitorUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow MonitorUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MonitorUpdateParamsChangeDetectionUnion struct {
	OfExact    *MonitorUpdateParamsChangeDetectionExact    `json:",omitzero,inline"`
	OfSemantic *MonitorUpdateParamsChangeDetectionSemantic `json:",omitzero,inline"`
	paramUnion
}

func (u MonitorUpdateParamsChangeDetectionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExact, u.OfSemantic)
}
func (u *MonitorUpdateParamsChangeDetectionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[MonitorUpdateParamsChangeDetectionUnion](
		"type",
		apijson.Discriminator[MonitorUpdateParamsChangeDetectionExact]("exact"),
		apijson.Discriminator[MonitorUpdateParamsChangeDetectionSemantic]("semantic"),
	)
}

func NewMonitorUpdateParamsChangeDetectionExact() MonitorUpdateParamsChangeDetectionExact {
	return MonitorUpdateParamsChangeDetectionExact{
		Type: "exact",
	}
}

// Detect exact changes. For page targets, this means visible text diffs. For
// sitemap targets, this means URL additions and removals.
//
// This struct has a constant value, construct it with
// [NewMonitorUpdateParamsChangeDetectionExact].
type MonitorUpdateParamsChangeDetectionExact struct {
	Type constant.Exact `json:"type" default:"exact"`
	paramObj
}

func (r MonitorUpdateParamsChangeDetectionExact) MarshalJSON() (data []byte, err error) {
	type shadow MonitorUpdateParamsChangeDetectionExact
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorUpdateParamsChangeDetectionExact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detect meaning-level changes to tracked page content, ignoring cosmetic or
// paraphrase-only differences. Which changes are meaningful is judged against the
// extract target's `instructions` (and `schema`, when provided).
//
// The property Type is required.
type MonitorUpdateParamsChangeDetectionSemantic struct {
	ConfidenceThreshold param.Opt[float64] `json:"confidence_threshold,omitzero"`
	// This field can be elided, and will marshal its zero value as "semantic".
	Type constant.Semantic `json:"type" default:"semantic"`
	paramObj
}

func (r MonitorUpdateParamsChangeDetectionSemantic) MarshalJSON() (data []byte, err error) {
	type shadow MonitorUpdateParamsChangeDetectionSemantic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorUpdateParamsChangeDetectionSemantic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
// every 6 hours or every 2 days. The total interval (frequency × unit) must be
// between 10 minutes and 1 year.
//
// The properties Frequency, Type, Unit are required.
type MonitorUpdateParamsSchedule struct {
	// Number of units between runs. The resulting interval (frequency × unit) must be
	// at least 10 minutes and at most 1 year (e.g. minimum 10 when unit is minutes;
	// maximum 365 when unit is days).
	Frequency int64 `json:"frequency" api:"required"`
	// Any of "interval".
	Type string `json:"type,omitzero" api:"required"`
	// Any of "minutes", "hours", "days".
	Unit string `json:"unit,omitzero" api:"required"`
	paramObj
}

func (r MonitorUpdateParamsSchedule) MarshalJSON() (data []byte, err error) {
	type shadow MonitorUpdateParamsSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorUpdateParamsSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MonitorUpdateParamsSchedule](
		"type", "interval",
	)
	apijson.RegisterFieldValidator[MonitorUpdateParamsSchedule](
		"unit", "minutes", "hours", "days",
	)
}

type MonitorUpdateParamsStatus string

const (
	MonitorUpdateParamsStatusActive MonitorUpdateParamsStatus = "active"
	MonitorUpdateParamsStatusPaused MonitorUpdateParamsStatus = "paused"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MonitorUpdateParamsTargetUnion struct {
	OfPage    *MonitorUpdateParamsTargetPage    `json:",omitzero,inline"`
	OfSitemap *MonitorUpdateParamsTargetSitemap `json:",omitzero,inline"`
	OfExtract *MonitorUpdateParamsTargetExtract `json:",omitzero,inline"`
	paramUnion
}

func (u MonitorUpdateParamsTargetUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPage, u.OfSitemap, u.OfExtract)
}
func (u *MonitorUpdateParamsTargetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[MonitorUpdateParamsTargetUnion](
		"type",
		apijson.Discriminator[MonitorUpdateParamsTargetPage]("page"),
		apijson.Discriminator[MonitorUpdateParamsTargetSitemap]("sitemap"),
		apijson.Discriminator[MonitorUpdateParamsTargetExtract]("extract"),
	)
}

// Watch a single web page.
//
// The properties Type, URL are required.
type MonitorUpdateParamsTargetPage struct {
	URL string `json:"url" api:"required" format:"uri"`
	// Normalize whitespace before comparing or analyzing text.
	NormalizeWhitespace param.Opt[bool] `json:"normalize_whitespace,omitzero"`
	// This field can be elided, and will marshal its zero value as "page".
	Type constant.Page `json:"type" default:"page"`
	paramObj
}

func (r MonitorUpdateParamsTargetPage) MarshalJSON() (data []byte, err error) {
	type shadow MonitorUpdateParamsTargetPage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorUpdateParamsTargetPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Watch a sitemap for URL additions and removals. Crawled URLs are normalized
// (lowercased host, no trailing slash/fragment) and scoped to the monitored site
// and its subdomains before comparison. On a detected difference the sitemap is
// re-fetched within the same run and only URLs both observations agree on are
// reported, suppressing transient crawl flaps.
//
// The properties Type, URL are required.
type MonitorUpdateParamsTargetSitemap struct {
	// Sitemap URL to monitor.
	URL string `json:"url" api:"required" format:"uri"`
	// Maximum number of sitemap URLs to track (capped at 10,000).
	MaxURLs param.Opt[int64] `json:"max_urls,omitzero"`
	// URL path patterns to exclude.
	Exclude []string `json:"exclude,omitzero"`
	// URL path patterns to include.
	Include []string `json:"include,omitzero"`
	// This field can be elided, and will marshal its zero value as "sitemap".
	Type constant.Sitemap `json:"type" default:"sitemap"`
	paramObj
}

func (r MonitorUpdateParamsTargetSitemap) MarshalJSON() (data []byte, err error) {
	type shadow MonitorUpdateParamsTargetSitemap
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorUpdateParamsTargetSitemap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Watch the monitor-relevant pages of a site for meaningful changes. A crawl
// guided by `schema`/`instructions` selects up to `max_pages` relevant pages to
// track; each run re-checks exactly those pages, and confirmed content changes are
// judged for relevance against the monitor's `instructions` (and `schema`, when
// provided). The tracked page set is refreshed by a periodic re-discovery crawl.
//
// The properties Instructions, Type, URL are required.
type MonitorUpdateParamsTargetExtract struct {
	// Natural-language instructions guiding which pages and facts to track and which
	// changes to report.
	Instructions string `json:"instructions" api:"required"`
	// Root URL to extract structured data from.
	URL              string          `json:"url" api:"required" format:"uri"`
	FollowSubdomains param.Opt[bool] `json:"follow_subdomains,omitzero"`
	// Optional maximum link depth from the starting URL (0 = only the starting page).
	MaxDepth param.Opt[int64] `json:"max_depth,omitzero"`
	// Maximum number of pages to track.
	MaxPages param.Opt[int64] `json:"max_pages,omitzero"`
	// JSON Schema describing the data you care about. It is used three ways: it guides
	// which pages are selected for tracking, it gives the change judge extra context
	// on which changes matter (alongside `instructions`), and it defines the shape of
	// the baseline `data` snapshot on GET /monitors/{monitor_id} (refreshed at most
	// about once a day). It is not a response format for changes: change events and
	// webhook payloads always contain diffs, summaries, and evidence excerpts — never
	// data in this schema's shape. If omitted, a default summary + key-points schema
	// is used.
	Schema map[string]any `json:"schema,omitzero"`
	// This field can be elided, and will marshal its zero value as "extract".
	Type constant.Extract `json:"type" default:"extract"`
	paramObj
}

func (r MonitorUpdateParamsTargetExtract) MarshalJSON() (data []byte, err error) {
	type shadow MonitorUpdateParamsTargetExtract
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorUpdateParamsTargetExtract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Set to null to remove the webhook.
//
// The property URL is required.
type MonitorUpdateParamsWebhook struct {
	// Webhook URL events are delivered to.
	URL string `json:"url" api:"required" format:"uri"`
	// Events delivered to this endpoint. `change.detected` fires only when a run
	// detects a change; `run.completed` fires on every completed run — including runs
	// that detected no change — and embeds the change when one was detected. Defaults
	// to `["change.detected"]` when omitted.
	//
	// Any of "change.detected", "run.completed".
	Events []string `json:"events,omitzero"`
	paramObj
}

func (r MonitorUpdateParamsWebhook) MarshalJSON() (data []byte, err error) {
	type shadow MonitorUpdateParamsWebhook
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorUpdateParamsWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListParams struct {
	// Opaque pagination cursor from a previous response.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return per page (1-100). Defaults to 25.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Free-text search term, matched against the fields named in `search_by`.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Filter to items that have this tag.
	Tag param.Opt[string] `query:"tag,omitzero" json:"-"`
	// Comma-separated fields to search with `q`. Defaults to all of them. Note
	// `instructions` only exists on extract monitors.
	//
	// Any of "name", "url", "instructions", "tags".
	SearchBy []string `query:"search_by,omitzero" json:"-"`
	// Comma-separated list of tags to filter by (matches monitors having any of them).
	Tags []string `query:"tags,omitzero" json:"-"`
	// Filter by change detection type.
	//
	// Any of "exact", "semantic".
	ChangeDetectionType MonitorListParamsChangeDetectionType `query:"change_detection_type,omitzero" json:"-"`
	// `prefix` for as-you-type prefix matching (default), `exact` for full-token
	// matching.
	//
	// Any of "exact", "prefix".
	SearchType MonitorListParamsSearchType `query:"search_type,omitzero" json:"-"`
	// Filter monitors by lifecycle status.
	//
	// Any of "active", "paused", "failed".
	Status MonitorListParamsStatus `query:"status,omitzero" json:"-"`
	// Filter by target type.
	//
	// Any of "page", "sitemap", "extract".
	TargetType MonitorListParamsTargetType `query:"target_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MonitorListParams]'s query parameters as `url.Values`.
func (r MonitorListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by change detection type.
type MonitorListParamsChangeDetectionType string

const (
	MonitorListParamsChangeDetectionTypeExact    MonitorListParamsChangeDetectionType = "exact"
	MonitorListParamsChangeDetectionTypeSemantic MonitorListParamsChangeDetectionType = "semantic"
)

// `prefix` for as-you-type prefix matching (default), `exact` for full-token
// matching.
type MonitorListParamsSearchType string

const (
	MonitorListParamsSearchTypeExact  MonitorListParamsSearchType = "exact"
	MonitorListParamsSearchTypePrefix MonitorListParamsSearchType = "prefix"
)

// Filter monitors by lifecycle status.
type MonitorListParamsStatus string

const (
	MonitorListParamsStatusActive MonitorListParamsStatus = "active"
	MonitorListParamsStatusPaused MonitorListParamsStatus = "paused"
	MonitorListParamsStatusFailed MonitorListParamsStatus = "failed"
)

// Filter by target type.
type MonitorListParamsTargetType string

const (
	MonitorListParamsTargetTypePage    MonitorListParamsTargetType = "page"
	MonitorListParamsTargetTypeSitemap MonitorListParamsTargetType = "sitemap"
	MonitorListParamsTargetTypeExtract MonitorListParamsTargetType = "extract"
)

type MonitorGetCreditUsageParams struct {
	// Only include items at or after this ISO 8601 timestamp.
	Since param.Opt[time.Time] `query:"since,omitzero" format:"date-time" json:"-"`
	// Only include items before this ISO 8601 timestamp.
	Until param.Opt[time.Time] `query:"until,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [MonitorGetCreditUsageParams]'s query parameters as
// `url.Values`.
func (r MonitorGetCreditUsageParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MonitorListAccountChangesParams struct {
	// Opaque pagination cursor from a previous response.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return per page (1-100). Defaults to 25.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter changes to a single monitor.
	MonitorID param.Opt[string] `query:"monitor_id,omitzero" json:"-"`
	// Only include items at or after this ISO 8601 timestamp.
	Since param.Opt[time.Time] `query:"since,omitzero" format:"date-time" json:"-"`
	// Filter to items that have this tag.
	Tag param.Opt[string] `query:"tag,omitzero" json:"-"`
	// Only include items before this ISO 8601 timestamp.
	Until param.Opt[time.Time] `query:"until,omitzero" format:"date-time" json:"-"`
	// Filter by change detection type.
	//
	// Any of "exact", "semantic".
	ChangeDetectionType MonitorListAccountChangesParamsChangeDetectionType `query:"change_detection_type,omitzero" json:"-"`
	// Filter by target type.
	//
	// Any of "page", "sitemap", "extract".
	TargetType MonitorListAccountChangesParamsTargetType `query:"target_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MonitorListAccountChangesParams]'s query parameters as
// `url.Values`.
func (r MonitorListAccountChangesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by change detection type.
type MonitorListAccountChangesParamsChangeDetectionType string

const (
	MonitorListAccountChangesParamsChangeDetectionTypeExact    MonitorListAccountChangesParamsChangeDetectionType = "exact"
	MonitorListAccountChangesParamsChangeDetectionTypeSemantic MonitorListAccountChangesParamsChangeDetectionType = "semantic"
)

// Filter by target type.
type MonitorListAccountChangesParamsTargetType string

const (
	MonitorListAccountChangesParamsTargetTypePage    MonitorListAccountChangesParamsTargetType = "page"
	MonitorListAccountChangesParamsTargetTypeSitemap MonitorListAccountChangesParamsTargetType = "sitemap"
	MonitorListAccountChangesParamsTargetTypeExtract MonitorListAccountChangesParamsTargetType = "extract"
)

type MonitorListAccountRunsParams struct {
	// Opaque pagination cursor from a previous response.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return per page (1-100). Defaults to 25.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter runs by lifecycle status.
	//
	// Any of "queued", "running", "completed", "failed", "skipped".
	Status MonitorListAccountRunsParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MonitorListAccountRunsParams]'s query parameters as
// `url.Values`.
func (r MonitorListAccountRunsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter runs by lifecycle status.
type MonitorListAccountRunsParamsStatus string

const (
	MonitorListAccountRunsParamsStatusQueued    MonitorListAccountRunsParamsStatus = "queued"
	MonitorListAccountRunsParamsStatusRunning   MonitorListAccountRunsParamsStatus = "running"
	MonitorListAccountRunsParamsStatusCompleted MonitorListAccountRunsParamsStatus = "completed"
	MonitorListAccountRunsParamsStatusFailed    MonitorListAccountRunsParamsStatus = "failed"
	MonitorListAccountRunsParamsStatusSkipped   MonitorListAccountRunsParamsStatus = "skipped"
)

type MonitorListChangesParams struct {
	// Opaque pagination cursor from a previous response.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return per page (1-100). Defaults to 25.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Only include items at or after this ISO 8601 timestamp.
	Since param.Opt[time.Time] `query:"since,omitzero" format:"date-time" json:"-"`
	// Filter to items that have this tag.
	Tag param.Opt[string] `query:"tag,omitzero" json:"-"`
	// Only include items before this ISO 8601 timestamp.
	Until param.Opt[time.Time] `query:"until,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [MonitorListChangesParams]'s query parameters as
// `url.Values`.
func (r MonitorListChangesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MonitorListRunsParams struct {
	// Opaque pagination cursor from a previous response.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return per page (1-100). Defaults to 25.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter runs by lifecycle status.
	//
	// Any of "queued", "running", "completed", "failed", "skipped".
	Status MonitorListRunsParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MonitorListRunsParams]'s query parameters as `url.Values`.
func (r MonitorListRunsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter runs by lifecycle status.
type MonitorListRunsParamsStatus string

const (
	MonitorListRunsParamsStatusQueued    MonitorListRunsParamsStatus = "queued"
	MonitorListRunsParamsStatusRunning   MonitorListRunsParamsStatus = "running"
	MonitorListRunsParamsStatusCompleted MonitorListRunsParamsStatus = "completed"
	MonitorListRunsParamsStatusFailed    MonitorListRunsParamsStatus = "failed"
	MonitorListRunsParamsStatusSkipped   MonitorListRunsParamsStatus = "skipped"
)
