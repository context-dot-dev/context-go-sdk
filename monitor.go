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

	"github.com/context-dot-dev/context-go-sdk/internal/apijson"
	"github.com/context-dot-dev/context-go-sdk/internal/apiquery"
	"github.com/context-dot-dev/context-go-sdk/internal/requestconfig"
	"github.com/context-dot-dev/context-go-sdk/option"
	"github.com/context-dot-dev/context-go-sdk/packages/param"
	"github.com/context-dot-dev/context-go-sdk/packages/respjson"
	"github.com/context-dot-dev/context-go-sdk/shared/constant"
)

// Monitor pages, sitemaps, and extracted website data for exact or semantic
// changes. The change.detected webhook payload is documented by the
// MonitorsChangeDetectedWebhookPayload schema.
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
func (r *MonitorService) New(ctx context.Context, body MonitorNewParams, opts ...option.RequestOption) (res *MonitorNewResponseUnion, err error) {
	opts = slices.Concat(r.options, opts)
	path := "monitors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a monitor
func (r *MonitorService) Get(ctx context.Context, monitorID string, opts ...option.RequestOption) (res *MonitorGetResponseUnion, err error) {
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
func (r *MonitorService) Update(ctx context.Context, monitorID string, body MonitorUpdateParams, opts ...option.RequestOption) (res *MonitorUpdateResponseUnion, err error) {
	opts = slices.Concat(r.options, opts)
	if monitorID == "" {
		err = errors.New("missing required monitor_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("monitors/%s", url.PathEscape(monitorID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List monitors
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
func (r *MonitorService) GetChange(ctx context.Context, changeID string, opts ...option.RequestOption) (res *MonitorGetChangeResponseUnion, err error) {
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

// MonitorNewResponseUnion contains all possible properties and values from
// [MonitorNewResponseMonitorsPageExactMonitor],
// [MonitorNewResponseMonitorsSitemapExactMonitor],
// [MonitorNewResponseMonitorsPageSemanticMonitor],
// [MonitorNewResponseMonitorsExtractSemanticMonitor].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MonitorNewResponseUnion struct {
	ID string `json:"id"`
	// This field is a union of
	// [MonitorNewResponseMonitorsPageExactMonitorChangeDetection],
	// [MonitorNewResponseMonitorsSitemapExactMonitorChangeDetection],
	// [MonitorNewResponseMonitorsPageSemanticMonitorChangeDetection],
	// [MonitorNewResponseMonitorsExtractSemanticMonitorChangeDetection]
	ChangeDetection MonitorNewResponseUnionChangeDetection `json:"change_detection"`
	CreatedAt       time.Time                              `json:"created_at"`
	Name            string                                 `json:"name"`
	// This field is a union of [MonitorNewResponseMonitorsPageExactMonitorSchedule],
	// [MonitorNewResponseMonitorsSitemapExactMonitorSchedule],
	// [MonitorNewResponseMonitorsPageSemanticMonitorSchedule],
	// [MonitorNewResponseMonitorsExtractSemanticMonitorSchedule]
	Schedule MonitorNewResponseUnionSchedule `json:"schedule"`
	Status   string                          `json:"status"`
	// This field is a union of [MonitorNewResponseMonitorsPageExactMonitorTarget],
	// [MonitorNewResponseMonitorsSitemapExactMonitorTarget],
	// [MonitorNewResponseMonitorsPageSemanticMonitorTarget],
	// [MonitorNewResponseMonitorsExtractSemanticMonitorTarget]
	Target       MonitorNewResponseUnionTarget `json:"target"`
	UpdatedAt    time.Time                     `json:"updated_at"`
	LastChangeAt time.Time                     `json:"last_change_at"`
	LastRunAt    time.Time                     `json:"last_run_at"`
	Tags         []string                      `json:"tags"`
	// This field is a union of [MonitorNewResponseMonitorsPageExactMonitorWebhook],
	// [MonitorNewResponseMonitorsSitemapExactMonitorWebhook],
	// [MonitorNewResponseMonitorsPageSemanticMonitorWebhook],
	// [MonitorNewResponseMonitorsExtractSemanticMonitorWebhook]
	Webhook MonitorNewResponseUnionWebhook `json:"webhook"`
	JSON    struct {
		ID              respjson.Field
		ChangeDetection respjson.Field
		CreatedAt       respjson.Field
		Name            respjson.Field
		Schedule        respjson.Field
		Status          respjson.Field
		Target          respjson.Field
		UpdatedAt       respjson.Field
		LastChangeAt    respjson.Field
		LastRunAt       respjson.Field
		Tags            respjson.Field
		Webhook         respjson.Field
		raw             string
	} `json:"-"`
}

func (u MonitorNewResponseUnion) AsMonitorNewResponseMonitorsPageExactMonitor() (v MonitorNewResponseMonitorsPageExactMonitor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorNewResponseUnion) AsMonitorNewResponseMonitorsSitemapExactMonitor() (v MonitorNewResponseMonitorsSitemapExactMonitor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorNewResponseUnion) AsMonitorNewResponseMonitorsPageSemanticMonitor() (v MonitorNewResponseMonitorsPageSemanticMonitor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorNewResponseUnion) AsMonitorNewResponseMonitorsExtractSemanticMonitor() (v MonitorNewResponseMonitorsExtractSemanticMonitor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MonitorNewResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *MonitorNewResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MonitorNewResponseUnionChangeDetection is an implicit subunion of
// [MonitorNewResponseUnion]. MonitorNewResponseUnionChangeDetection provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MonitorNewResponseUnion].
type MonitorNewResponseUnionChangeDetection struct {
	Type                string  `json:"type"`
	Query               string  `json:"query"`
	ConfidenceThreshold float64 `json:"confidence_threshold"`
	JSON                struct {
		Type                respjson.Field
		Query               respjson.Field
		ConfidenceThreshold respjson.Field
		raw                 string
	} `json:"-"`
}

func (r *MonitorNewResponseUnionChangeDetection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MonitorNewResponseUnionSchedule is an implicit subunion of
// [MonitorNewResponseUnion]. MonitorNewResponseUnionSchedule provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MonitorNewResponseUnion].
type MonitorNewResponseUnionSchedule struct {
	Frequency int64  `json:"frequency"`
	Type      string `json:"type"`
	Unit      string `json:"unit"`
	JSON      struct {
		Frequency respjson.Field
		Type      respjson.Field
		Unit      respjson.Field
		raw       string
	} `json:"-"`
}

func (r *MonitorNewResponseUnionSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MonitorNewResponseUnionTarget is an implicit subunion of
// [MonitorNewResponseUnion]. MonitorNewResponseUnionTarget provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MonitorNewResponseUnion].
type MonitorNewResponseUnionTarget struct {
	Type                string `json:"type"`
	URL                 string `json:"url"`
	NormalizeWhitespace bool   `json:"normalize_whitespace"`
	// This field is from variant
	// [MonitorNewResponseMonitorsSitemapExactMonitorTarget].
	Exclude []string `json:"exclude"`
	// This field is from variant
	// [MonitorNewResponseMonitorsSitemapExactMonitorTarget].
	Include []string `json:"include"`
	// This field is from variant
	// [MonitorNewResponseMonitorsSitemapExactMonitorTarget].
	MaxURLs int64 `json:"max_urls"`
	// This field is from variant
	// [MonitorNewResponseMonitorsExtractSemanticMonitorTarget].
	FollowSubdomains bool `json:"follow_subdomains"`
	// This field is from variant
	// [MonitorNewResponseMonitorsExtractSemanticMonitorTarget].
	Instructions string `json:"instructions"`
	// This field is from variant
	// [MonitorNewResponseMonitorsExtractSemanticMonitorTarget].
	MaxDepth int64 `json:"max_depth"`
	// This field is from variant
	// [MonitorNewResponseMonitorsExtractSemanticMonitorTarget].
	MaxPages int64 `json:"max_pages"`
	// This field is from variant
	// [MonitorNewResponseMonitorsExtractSemanticMonitorTarget].
	Schema map[string]any `json:"schema"`
	JSON   struct {
		Type                respjson.Field
		URL                 respjson.Field
		NormalizeWhitespace respjson.Field
		Exclude             respjson.Field
		Include             respjson.Field
		MaxURLs             respjson.Field
		FollowSubdomains    respjson.Field
		Instructions        respjson.Field
		MaxDepth            respjson.Field
		MaxPages            respjson.Field
		Schema              respjson.Field
		raw                 string
	} `json:"-"`
}

func (r *MonitorNewResponseUnionTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MonitorNewResponseUnionWebhook is an implicit subunion of
// [MonitorNewResponseUnion]. MonitorNewResponseUnionWebhook provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MonitorNewResponseUnion].
type MonitorNewResponseUnionWebhook struct {
	URL    string `json:"url"`
	Secret string `json:"secret"`
	JSON   struct {
		URL    respjson.Field
		Secret respjson.Field
		raw    string
	} `json:"-"`
}

func (r *MonitorNewResponseUnionWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A page monitor using exact change detection.
type MonitorNewResponseMonitorsPageExactMonitor struct {
	ID string `json:"id" api:"required"`
	// Detect exact changes. For page targets, this means visible text diffs. For
	// sitemap targets, this means URL additions and removals.
	ChangeDetection MonitorNewResponseMonitorsPageExactMonitorChangeDetection `json:"change_detection" api:"required"`
	CreatedAt       time.Time                                                 `json:"created_at" api:"required" format:"date-time"`
	Name            string                                                    `json:"name" api:"required"`
	// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
	// every 6 hours or every 2 days. The total interval (frequency × unit) must be
	// between 10 minutes and 1 year.
	Schedule MonitorNewResponseMonitorsPageExactMonitorSchedule `json:"schedule" api:"required"`
	// Any of "active", "paused", "failed".
	Status       string                                           `json:"status" api:"required"`
	Target       MonitorNewResponseMonitorsPageExactMonitorTarget `json:"target" api:"required"`
	UpdatedAt    time.Time                                        `json:"updated_at" api:"required" format:"date-time"`
	LastChangeAt time.Time                                        `json:"last_change_at" api:"nullable" format:"date-time"`
	LastRunAt    time.Time                                        `json:"last_run_at" api:"nullable" format:"date-time"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags    []string                                          `json:"tags"`
	Webhook MonitorNewResponseMonitorsPageExactMonitorWebhook `json:"webhook" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ChangeDetection respjson.Field
		CreatedAt       respjson.Field
		Name            respjson.Field
		Schedule        respjson.Field
		Status          respjson.Field
		Target          respjson.Field
		UpdatedAt       respjson.Field
		LastChangeAt    respjson.Field
		LastRunAt       respjson.Field
		Tags            respjson.Field
		Webhook         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorNewResponseMonitorsPageExactMonitor) RawJSON() string { return r.JSON.raw }
func (r *MonitorNewResponseMonitorsPageExactMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detect exact changes. For page targets, this means visible text diffs. For
// sitemap targets, this means URL additions and removals.
type MonitorNewResponseMonitorsPageExactMonitorChangeDetection struct {
	// Any of "exact".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorNewResponseMonitorsPageExactMonitorChangeDetection) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorNewResponseMonitorsPageExactMonitorChangeDetection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
// every 6 hours or every 2 days. The total interval (frequency × unit) must be
// between 10 minutes and 1 year.
type MonitorNewResponseMonitorsPageExactMonitorSchedule struct {
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
func (r MonitorNewResponseMonitorsPageExactMonitorSchedule) RawJSON() string { return r.JSON.raw }
func (r *MonitorNewResponseMonitorsPageExactMonitorSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorNewResponseMonitorsPageExactMonitorTarget struct {
	// Any of "page".
	Type string `json:"type" api:"required"`
	URL  string `json:"url" api:"required" format:"uri"`
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
func (r MonitorNewResponseMonitorsPageExactMonitorTarget) RawJSON() string { return r.JSON.raw }
func (r *MonitorNewResponseMonitorsPageExactMonitorTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorNewResponseMonitorsPageExactMonitorWebhook struct {
	// Webhook URL called when a change is detected.
	URL string `json:"url" api:"required" format:"uri"`
	// Signing secret used to verify webhook authenticity. Each delivery includes an
	// `X-Context-Signature: t=<unix>,v1=<hmac>` header, where the HMAC is SHA-256 over
	// `"{t}.{rawRequestBody}"` keyed by this secret. Recompute it with a constant-time
	// compare and reject stale timestamps to prevent replay. Generated by the API;
	// cannot be set by clients.
	Secret string `json:"secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorNewResponseMonitorsPageExactMonitorWebhook) RawJSON() string { return r.JSON.raw }
func (r *MonitorNewResponseMonitorsPageExactMonitorWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A sitemap monitor using exact change detection.
type MonitorNewResponseMonitorsSitemapExactMonitor struct {
	ID string `json:"id" api:"required"`
	// Detect exact changes. For page targets, this means visible text diffs. For
	// sitemap targets, this means URL additions and removals.
	ChangeDetection MonitorNewResponseMonitorsSitemapExactMonitorChangeDetection `json:"change_detection" api:"required"`
	CreatedAt       time.Time                                                    `json:"created_at" api:"required" format:"date-time"`
	Name            string                                                       `json:"name" api:"required"`
	// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
	// every 6 hours or every 2 days. The total interval (frequency × unit) must be
	// between 10 minutes and 1 year.
	Schedule MonitorNewResponseMonitorsSitemapExactMonitorSchedule `json:"schedule" api:"required"`
	// Any of "active", "paused", "failed".
	Status       string                                              `json:"status" api:"required"`
	Target       MonitorNewResponseMonitorsSitemapExactMonitorTarget `json:"target" api:"required"`
	UpdatedAt    time.Time                                           `json:"updated_at" api:"required" format:"date-time"`
	LastChangeAt time.Time                                           `json:"last_change_at" api:"nullable" format:"date-time"`
	LastRunAt    time.Time                                           `json:"last_run_at" api:"nullable" format:"date-time"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags    []string                                             `json:"tags"`
	Webhook MonitorNewResponseMonitorsSitemapExactMonitorWebhook `json:"webhook" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ChangeDetection respjson.Field
		CreatedAt       respjson.Field
		Name            respjson.Field
		Schedule        respjson.Field
		Status          respjson.Field
		Target          respjson.Field
		UpdatedAt       respjson.Field
		LastChangeAt    respjson.Field
		LastRunAt       respjson.Field
		Tags            respjson.Field
		Webhook         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorNewResponseMonitorsSitemapExactMonitor) RawJSON() string { return r.JSON.raw }
func (r *MonitorNewResponseMonitorsSitemapExactMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detect exact changes. For page targets, this means visible text diffs. For
// sitemap targets, this means URL additions and removals.
type MonitorNewResponseMonitorsSitemapExactMonitorChangeDetection struct {
	// Any of "exact".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorNewResponseMonitorsSitemapExactMonitorChangeDetection) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorNewResponseMonitorsSitemapExactMonitorChangeDetection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
// every 6 hours or every 2 days. The total interval (frequency × unit) must be
// between 10 minutes and 1 year.
type MonitorNewResponseMonitorsSitemapExactMonitorSchedule struct {
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
func (r MonitorNewResponseMonitorsSitemapExactMonitorSchedule) RawJSON() string { return r.JSON.raw }
func (r *MonitorNewResponseMonitorsSitemapExactMonitorSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorNewResponseMonitorsSitemapExactMonitorTarget struct {
	// Any of "sitemap".
	Type string `json:"type" api:"required"`
	// Sitemap URL to monitor.
	URL string `json:"url" api:"required" format:"uri"`
	// URL path patterns to exclude.
	Exclude []string `json:"exclude"`
	// URL path patterns to include.
	Include []string `json:"include"`
	MaxURLs int64    `json:"max_urls"`
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
func (r MonitorNewResponseMonitorsSitemapExactMonitorTarget) RawJSON() string { return r.JSON.raw }
func (r *MonitorNewResponseMonitorsSitemapExactMonitorTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorNewResponseMonitorsSitemapExactMonitorWebhook struct {
	// Webhook URL called when a change is detected.
	URL string `json:"url" api:"required" format:"uri"`
	// Signing secret used to verify webhook authenticity. Each delivery includes an
	// `X-Context-Signature: t=<unix>,v1=<hmac>` header, where the HMAC is SHA-256 over
	// `"{t}.{rawRequestBody}"` keyed by this secret. Recompute it with a constant-time
	// compare and reject stale timestamps to prevent replay. Generated by the API;
	// cannot be set by clients.
	Secret string `json:"secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorNewResponseMonitorsSitemapExactMonitorWebhook) RawJSON() string { return r.JSON.raw }
func (r *MonitorNewResponseMonitorsSitemapExactMonitorWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A page monitor using semantic change detection.
type MonitorNewResponseMonitorsPageSemanticMonitor struct {
	ID string `json:"id" api:"required"`
	// Detect meaning-level changes that match a natural language query.
	ChangeDetection MonitorNewResponseMonitorsPageSemanticMonitorChangeDetection `json:"change_detection" api:"required"`
	CreatedAt       time.Time                                                    `json:"created_at" api:"required" format:"date-time"`
	Name            string                                                       `json:"name" api:"required"`
	// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
	// every 6 hours or every 2 days. The total interval (frequency × unit) must be
	// between 10 minutes and 1 year.
	Schedule MonitorNewResponseMonitorsPageSemanticMonitorSchedule `json:"schedule" api:"required"`
	// Any of "active", "paused", "failed".
	Status       string                                              `json:"status" api:"required"`
	Target       MonitorNewResponseMonitorsPageSemanticMonitorTarget `json:"target" api:"required"`
	UpdatedAt    time.Time                                           `json:"updated_at" api:"required" format:"date-time"`
	LastChangeAt time.Time                                           `json:"last_change_at" api:"nullable" format:"date-time"`
	LastRunAt    time.Time                                           `json:"last_run_at" api:"nullable" format:"date-time"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags    []string                                             `json:"tags"`
	Webhook MonitorNewResponseMonitorsPageSemanticMonitorWebhook `json:"webhook" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ChangeDetection respjson.Field
		CreatedAt       respjson.Field
		Name            respjson.Field
		Schedule        respjson.Field
		Status          respjson.Field
		Target          respjson.Field
		UpdatedAt       respjson.Field
		LastChangeAt    respjson.Field
		LastRunAt       respjson.Field
		Tags            respjson.Field
		Webhook         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorNewResponseMonitorsPageSemanticMonitor) RawJSON() string { return r.JSON.raw }
func (r *MonitorNewResponseMonitorsPageSemanticMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detect meaning-level changes that match a natural language query.
type MonitorNewResponseMonitorsPageSemanticMonitorChangeDetection struct {
	Query string `json:"query" api:"required"`
	// Any of "semantic".
	Type                string  `json:"type" api:"required"`
	ConfidenceThreshold float64 `json:"confidence_threshold"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Query               respjson.Field
		Type                respjson.Field
		ConfidenceThreshold respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorNewResponseMonitorsPageSemanticMonitorChangeDetection) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorNewResponseMonitorsPageSemanticMonitorChangeDetection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
// every 6 hours or every 2 days. The total interval (frequency × unit) must be
// between 10 minutes and 1 year.
type MonitorNewResponseMonitorsPageSemanticMonitorSchedule struct {
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
func (r MonitorNewResponseMonitorsPageSemanticMonitorSchedule) RawJSON() string { return r.JSON.raw }
func (r *MonitorNewResponseMonitorsPageSemanticMonitorSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorNewResponseMonitorsPageSemanticMonitorTarget struct {
	// Any of "page".
	Type string `json:"type" api:"required"`
	URL  string `json:"url" api:"required" format:"uri"`
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
func (r MonitorNewResponseMonitorsPageSemanticMonitorTarget) RawJSON() string { return r.JSON.raw }
func (r *MonitorNewResponseMonitorsPageSemanticMonitorTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorNewResponseMonitorsPageSemanticMonitorWebhook struct {
	// Webhook URL called when a change is detected.
	URL string `json:"url" api:"required" format:"uri"`
	// Signing secret used to verify webhook authenticity. Each delivery includes an
	// `X-Context-Signature: t=<unix>,v1=<hmac>` header, where the HMAC is SHA-256 over
	// `"{t}.{rawRequestBody}"` keyed by this secret. Recompute it with a constant-time
	// compare and reject stale timestamps to prevent replay. Generated by the API;
	// cannot be set by clients.
	Secret string `json:"secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorNewResponseMonitorsPageSemanticMonitorWebhook) RawJSON() string { return r.JSON.raw }
func (r *MonitorNewResponseMonitorsPageSemanticMonitorWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An extract monitor using semantic change detection.
type MonitorNewResponseMonitorsExtractSemanticMonitor struct {
	ID string `json:"id" api:"required"`
	// Detect meaning-level changes that match a natural language query.
	ChangeDetection MonitorNewResponseMonitorsExtractSemanticMonitorChangeDetection `json:"change_detection" api:"required"`
	CreatedAt       time.Time                                                       `json:"created_at" api:"required" format:"date-time"`
	Name            string                                                          `json:"name" api:"required"`
	// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
	// every 6 hours or every 2 days. The total interval (frequency × unit) must be
	// between 10 minutes and 1 year.
	Schedule MonitorNewResponseMonitorsExtractSemanticMonitorSchedule `json:"schedule" api:"required"`
	// Any of "active", "paused", "failed".
	Status       string                                                 `json:"status" api:"required"`
	Target       MonitorNewResponseMonitorsExtractSemanticMonitorTarget `json:"target" api:"required"`
	UpdatedAt    time.Time                                              `json:"updated_at" api:"required" format:"date-time"`
	LastChangeAt time.Time                                              `json:"last_change_at" api:"nullable" format:"date-time"`
	LastRunAt    time.Time                                              `json:"last_run_at" api:"nullable" format:"date-time"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags    []string                                                `json:"tags"`
	Webhook MonitorNewResponseMonitorsExtractSemanticMonitorWebhook `json:"webhook" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ChangeDetection respjson.Field
		CreatedAt       respjson.Field
		Name            respjson.Field
		Schedule        respjson.Field
		Status          respjson.Field
		Target          respjson.Field
		UpdatedAt       respjson.Field
		LastChangeAt    respjson.Field
		LastRunAt       respjson.Field
		Tags            respjson.Field
		Webhook         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorNewResponseMonitorsExtractSemanticMonitor) RawJSON() string { return r.JSON.raw }
func (r *MonitorNewResponseMonitorsExtractSemanticMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detect meaning-level changes that match a natural language query.
type MonitorNewResponseMonitorsExtractSemanticMonitorChangeDetection struct {
	Query string `json:"query" api:"required"`
	// Any of "semantic".
	Type                string  `json:"type" api:"required"`
	ConfidenceThreshold float64 `json:"confidence_threshold"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Query               respjson.Field
		Type                respjson.Field
		ConfidenceThreshold respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorNewResponseMonitorsExtractSemanticMonitorChangeDetection) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorNewResponseMonitorsExtractSemanticMonitorChangeDetection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
// every 6 hours or every 2 days. The total interval (frequency × unit) must be
// between 10 minutes and 1 year.
type MonitorNewResponseMonitorsExtractSemanticMonitorSchedule struct {
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
func (r MonitorNewResponseMonitorsExtractSemanticMonitorSchedule) RawJSON() string { return r.JSON.raw }
func (r *MonitorNewResponseMonitorsExtractSemanticMonitorSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorNewResponseMonitorsExtractSemanticMonitorTarget struct {
	// Any of "extract".
	Type string `json:"type" api:"required"`
	// Root URL to extract structured data from.
	URL              string `json:"url" api:"required" format:"uri"`
	FollowSubdomains bool   `json:"follow_subdomains"`
	// Optional natural-language instructions guiding what to extract.
	Instructions string `json:"instructions"`
	// Optional maximum link depth from the starting URL (0 = only the starting page).
	MaxDepth int64 `json:"max_depth"`
	// Maximum number of pages to analyze during extraction.
	MaxPages int64 `json:"max_pages"`
	// JSON Schema describing the structured data to extract and watch for changes. If
	// omitted, a default summary + key-points schema is used.
	Schema map[string]any `json:"schema"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type             respjson.Field
		URL              respjson.Field
		FollowSubdomains respjson.Field
		Instructions     respjson.Field
		MaxDepth         respjson.Field
		MaxPages         respjson.Field
		Schema           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorNewResponseMonitorsExtractSemanticMonitorTarget) RawJSON() string { return r.JSON.raw }
func (r *MonitorNewResponseMonitorsExtractSemanticMonitorTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorNewResponseMonitorsExtractSemanticMonitorWebhook struct {
	// Webhook URL called when a change is detected.
	URL string `json:"url" api:"required" format:"uri"`
	// Signing secret used to verify webhook authenticity. Each delivery includes an
	// `X-Context-Signature: t=<unix>,v1=<hmac>` header, where the HMAC is SHA-256 over
	// `"{t}.{rawRequestBody}"` keyed by this secret. Recompute it with a constant-time
	// compare and reject stale timestamps to prevent replay. Generated by the API;
	// cannot be set by clients.
	Secret string `json:"secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorNewResponseMonitorsExtractSemanticMonitorWebhook) RawJSON() string { return r.JSON.raw }
func (r *MonitorNewResponseMonitorsExtractSemanticMonitorWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MonitorGetResponseUnion contains all possible properties and values from
// [MonitorGetResponseMonitorsPageExactMonitor],
// [MonitorGetResponseMonitorsSitemapExactMonitor],
// [MonitorGetResponseMonitorsPageSemanticMonitor],
// [MonitorGetResponseMonitorsExtractSemanticMonitor].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MonitorGetResponseUnion struct {
	ID string `json:"id"`
	// This field is a union of
	// [MonitorGetResponseMonitorsPageExactMonitorChangeDetection],
	// [MonitorGetResponseMonitorsSitemapExactMonitorChangeDetection],
	// [MonitorGetResponseMonitorsPageSemanticMonitorChangeDetection],
	// [MonitorGetResponseMonitorsExtractSemanticMonitorChangeDetection]
	ChangeDetection MonitorGetResponseUnionChangeDetection `json:"change_detection"`
	CreatedAt       time.Time                              `json:"created_at"`
	Name            string                                 `json:"name"`
	// This field is a union of [MonitorGetResponseMonitorsPageExactMonitorSchedule],
	// [MonitorGetResponseMonitorsSitemapExactMonitorSchedule],
	// [MonitorGetResponseMonitorsPageSemanticMonitorSchedule],
	// [MonitorGetResponseMonitorsExtractSemanticMonitorSchedule]
	Schedule MonitorGetResponseUnionSchedule `json:"schedule"`
	Status   string                          `json:"status"`
	// This field is a union of [MonitorGetResponseMonitorsPageExactMonitorTarget],
	// [MonitorGetResponseMonitorsSitemapExactMonitorTarget],
	// [MonitorGetResponseMonitorsPageSemanticMonitorTarget],
	// [MonitorGetResponseMonitorsExtractSemanticMonitorTarget]
	Target       MonitorGetResponseUnionTarget `json:"target"`
	UpdatedAt    time.Time                     `json:"updated_at"`
	LastChangeAt time.Time                     `json:"last_change_at"`
	LastRunAt    time.Time                     `json:"last_run_at"`
	Tags         []string                      `json:"tags"`
	// This field is a union of [MonitorGetResponseMonitorsPageExactMonitorWebhook],
	// [MonitorGetResponseMonitorsSitemapExactMonitorWebhook],
	// [MonitorGetResponseMonitorsPageSemanticMonitorWebhook],
	// [MonitorGetResponseMonitorsExtractSemanticMonitorWebhook]
	Webhook MonitorGetResponseUnionWebhook `json:"webhook"`
	JSON    struct {
		ID              respjson.Field
		ChangeDetection respjson.Field
		CreatedAt       respjson.Field
		Name            respjson.Field
		Schedule        respjson.Field
		Status          respjson.Field
		Target          respjson.Field
		UpdatedAt       respjson.Field
		LastChangeAt    respjson.Field
		LastRunAt       respjson.Field
		Tags            respjson.Field
		Webhook         respjson.Field
		raw             string
	} `json:"-"`
}

func (u MonitorGetResponseUnion) AsMonitorGetResponseMonitorsPageExactMonitor() (v MonitorGetResponseMonitorsPageExactMonitor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorGetResponseUnion) AsMonitorGetResponseMonitorsSitemapExactMonitor() (v MonitorGetResponseMonitorsSitemapExactMonitor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorGetResponseUnion) AsMonitorGetResponseMonitorsPageSemanticMonitor() (v MonitorGetResponseMonitorsPageSemanticMonitor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorGetResponseUnion) AsMonitorGetResponseMonitorsExtractSemanticMonitor() (v MonitorGetResponseMonitorsExtractSemanticMonitor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MonitorGetResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *MonitorGetResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MonitorGetResponseUnionChangeDetection is an implicit subunion of
// [MonitorGetResponseUnion]. MonitorGetResponseUnionChangeDetection provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MonitorGetResponseUnion].
type MonitorGetResponseUnionChangeDetection struct {
	Type                string  `json:"type"`
	Query               string  `json:"query"`
	ConfidenceThreshold float64 `json:"confidence_threshold"`
	JSON                struct {
		Type                respjson.Field
		Query               respjson.Field
		ConfidenceThreshold respjson.Field
		raw                 string
	} `json:"-"`
}

func (r *MonitorGetResponseUnionChangeDetection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MonitorGetResponseUnionSchedule is an implicit subunion of
// [MonitorGetResponseUnion]. MonitorGetResponseUnionSchedule provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MonitorGetResponseUnion].
type MonitorGetResponseUnionSchedule struct {
	Frequency int64  `json:"frequency"`
	Type      string `json:"type"`
	Unit      string `json:"unit"`
	JSON      struct {
		Frequency respjson.Field
		Type      respjson.Field
		Unit      respjson.Field
		raw       string
	} `json:"-"`
}

func (r *MonitorGetResponseUnionSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MonitorGetResponseUnionTarget is an implicit subunion of
// [MonitorGetResponseUnion]. MonitorGetResponseUnionTarget provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MonitorGetResponseUnion].
type MonitorGetResponseUnionTarget struct {
	Type                string `json:"type"`
	URL                 string `json:"url"`
	NormalizeWhitespace bool   `json:"normalize_whitespace"`
	// This field is from variant
	// [MonitorGetResponseMonitorsSitemapExactMonitorTarget].
	Exclude []string `json:"exclude"`
	// This field is from variant
	// [MonitorGetResponseMonitorsSitemapExactMonitorTarget].
	Include []string `json:"include"`
	// This field is from variant
	// [MonitorGetResponseMonitorsSitemapExactMonitorTarget].
	MaxURLs int64 `json:"max_urls"`
	// This field is from variant
	// [MonitorGetResponseMonitorsExtractSemanticMonitorTarget].
	FollowSubdomains bool `json:"follow_subdomains"`
	// This field is from variant
	// [MonitorGetResponseMonitorsExtractSemanticMonitorTarget].
	Instructions string `json:"instructions"`
	// This field is from variant
	// [MonitorGetResponseMonitorsExtractSemanticMonitorTarget].
	MaxDepth int64 `json:"max_depth"`
	// This field is from variant
	// [MonitorGetResponseMonitorsExtractSemanticMonitorTarget].
	MaxPages int64 `json:"max_pages"`
	// This field is from variant
	// [MonitorGetResponseMonitorsExtractSemanticMonitorTarget].
	Schema map[string]any `json:"schema"`
	JSON   struct {
		Type                respjson.Field
		URL                 respjson.Field
		NormalizeWhitespace respjson.Field
		Exclude             respjson.Field
		Include             respjson.Field
		MaxURLs             respjson.Field
		FollowSubdomains    respjson.Field
		Instructions        respjson.Field
		MaxDepth            respjson.Field
		MaxPages            respjson.Field
		Schema              respjson.Field
		raw                 string
	} `json:"-"`
}

func (r *MonitorGetResponseUnionTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MonitorGetResponseUnionWebhook is an implicit subunion of
// [MonitorGetResponseUnion]. MonitorGetResponseUnionWebhook provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MonitorGetResponseUnion].
type MonitorGetResponseUnionWebhook struct {
	URL    string `json:"url"`
	Secret string `json:"secret"`
	JSON   struct {
		URL    respjson.Field
		Secret respjson.Field
		raw    string
	} `json:"-"`
}

func (r *MonitorGetResponseUnionWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A page monitor using exact change detection.
type MonitorGetResponseMonitorsPageExactMonitor struct {
	ID string `json:"id" api:"required"`
	// Detect exact changes. For page targets, this means visible text diffs. For
	// sitemap targets, this means URL additions and removals.
	ChangeDetection MonitorGetResponseMonitorsPageExactMonitorChangeDetection `json:"change_detection" api:"required"`
	CreatedAt       time.Time                                                 `json:"created_at" api:"required" format:"date-time"`
	Name            string                                                    `json:"name" api:"required"`
	// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
	// every 6 hours or every 2 days. The total interval (frequency × unit) must be
	// between 10 minutes and 1 year.
	Schedule MonitorGetResponseMonitorsPageExactMonitorSchedule `json:"schedule" api:"required"`
	// Any of "active", "paused", "failed".
	Status       string                                           `json:"status" api:"required"`
	Target       MonitorGetResponseMonitorsPageExactMonitorTarget `json:"target" api:"required"`
	UpdatedAt    time.Time                                        `json:"updated_at" api:"required" format:"date-time"`
	LastChangeAt time.Time                                        `json:"last_change_at" api:"nullable" format:"date-time"`
	LastRunAt    time.Time                                        `json:"last_run_at" api:"nullable" format:"date-time"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags    []string                                          `json:"tags"`
	Webhook MonitorGetResponseMonitorsPageExactMonitorWebhook `json:"webhook" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ChangeDetection respjson.Field
		CreatedAt       respjson.Field
		Name            respjson.Field
		Schedule        respjson.Field
		Status          respjson.Field
		Target          respjson.Field
		UpdatedAt       respjson.Field
		LastChangeAt    respjson.Field
		LastRunAt       respjson.Field
		Tags            respjson.Field
		Webhook         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetResponseMonitorsPageExactMonitor) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetResponseMonitorsPageExactMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detect exact changes. For page targets, this means visible text diffs. For
// sitemap targets, this means URL additions and removals.
type MonitorGetResponseMonitorsPageExactMonitorChangeDetection struct {
	// Any of "exact".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetResponseMonitorsPageExactMonitorChangeDetection) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorGetResponseMonitorsPageExactMonitorChangeDetection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
// every 6 hours or every 2 days. The total interval (frequency × unit) must be
// between 10 minutes and 1 year.
type MonitorGetResponseMonitorsPageExactMonitorSchedule struct {
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
func (r MonitorGetResponseMonitorsPageExactMonitorSchedule) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetResponseMonitorsPageExactMonitorSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorGetResponseMonitorsPageExactMonitorTarget struct {
	// Any of "page".
	Type string `json:"type" api:"required"`
	URL  string `json:"url" api:"required" format:"uri"`
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
func (r MonitorGetResponseMonitorsPageExactMonitorTarget) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetResponseMonitorsPageExactMonitorTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorGetResponseMonitorsPageExactMonitorWebhook struct {
	// Webhook URL called when a change is detected.
	URL string `json:"url" api:"required" format:"uri"`
	// Signing secret used to verify webhook authenticity. Each delivery includes an
	// `X-Context-Signature: t=<unix>,v1=<hmac>` header, where the HMAC is SHA-256 over
	// `"{t}.{rawRequestBody}"` keyed by this secret. Recompute it with a constant-time
	// compare and reject stale timestamps to prevent replay. Generated by the API;
	// cannot be set by clients.
	Secret string `json:"secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetResponseMonitorsPageExactMonitorWebhook) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetResponseMonitorsPageExactMonitorWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A sitemap monitor using exact change detection.
type MonitorGetResponseMonitorsSitemapExactMonitor struct {
	ID string `json:"id" api:"required"`
	// Detect exact changes. For page targets, this means visible text diffs. For
	// sitemap targets, this means URL additions and removals.
	ChangeDetection MonitorGetResponseMonitorsSitemapExactMonitorChangeDetection `json:"change_detection" api:"required"`
	CreatedAt       time.Time                                                    `json:"created_at" api:"required" format:"date-time"`
	Name            string                                                       `json:"name" api:"required"`
	// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
	// every 6 hours or every 2 days. The total interval (frequency × unit) must be
	// between 10 minutes and 1 year.
	Schedule MonitorGetResponseMonitorsSitemapExactMonitorSchedule `json:"schedule" api:"required"`
	// Any of "active", "paused", "failed".
	Status       string                                              `json:"status" api:"required"`
	Target       MonitorGetResponseMonitorsSitemapExactMonitorTarget `json:"target" api:"required"`
	UpdatedAt    time.Time                                           `json:"updated_at" api:"required" format:"date-time"`
	LastChangeAt time.Time                                           `json:"last_change_at" api:"nullable" format:"date-time"`
	LastRunAt    time.Time                                           `json:"last_run_at" api:"nullable" format:"date-time"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags    []string                                             `json:"tags"`
	Webhook MonitorGetResponseMonitorsSitemapExactMonitorWebhook `json:"webhook" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ChangeDetection respjson.Field
		CreatedAt       respjson.Field
		Name            respjson.Field
		Schedule        respjson.Field
		Status          respjson.Field
		Target          respjson.Field
		UpdatedAt       respjson.Field
		LastChangeAt    respjson.Field
		LastRunAt       respjson.Field
		Tags            respjson.Field
		Webhook         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetResponseMonitorsSitemapExactMonitor) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetResponseMonitorsSitemapExactMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detect exact changes. For page targets, this means visible text diffs. For
// sitemap targets, this means URL additions and removals.
type MonitorGetResponseMonitorsSitemapExactMonitorChangeDetection struct {
	// Any of "exact".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetResponseMonitorsSitemapExactMonitorChangeDetection) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorGetResponseMonitorsSitemapExactMonitorChangeDetection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
// every 6 hours or every 2 days. The total interval (frequency × unit) must be
// between 10 minutes and 1 year.
type MonitorGetResponseMonitorsSitemapExactMonitorSchedule struct {
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
func (r MonitorGetResponseMonitorsSitemapExactMonitorSchedule) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetResponseMonitorsSitemapExactMonitorSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorGetResponseMonitorsSitemapExactMonitorTarget struct {
	// Any of "sitemap".
	Type string `json:"type" api:"required"`
	// Sitemap URL to monitor.
	URL string `json:"url" api:"required" format:"uri"`
	// URL path patterns to exclude.
	Exclude []string `json:"exclude"`
	// URL path patterns to include.
	Include []string `json:"include"`
	MaxURLs int64    `json:"max_urls"`
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
func (r MonitorGetResponseMonitorsSitemapExactMonitorTarget) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetResponseMonitorsSitemapExactMonitorTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorGetResponseMonitorsSitemapExactMonitorWebhook struct {
	// Webhook URL called when a change is detected.
	URL string `json:"url" api:"required" format:"uri"`
	// Signing secret used to verify webhook authenticity. Each delivery includes an
	// `X-Context-Signature: t=<unix>,v1=<hmac>` header, where the HMAC is SHA-256 over
	// `"{t}.{rawRequestBody}"` keyed by this secret. Recompute it with a constant-time
	// compare and reject stale timestamps to prevent replay. Generated by the API;
	// cannot be set by clients.
	Secret string `json:"secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetResponseMonitorsSitemapExactMonitorWebhook) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetResponseMonitorsSitemapExactMonitorWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A page monitor using semantic change detection.
type MonitorGetResponseMonitorsPageSemanticMonitor struct {
	ID string `json:"id" api:"required"`
	// Detect meaning-level changes that match a natural language query.
	ChangeDetection MonitorGetResponseMonitorsPageSemanticMonitorChangeDetection `json:"change_detection" api:"required"`
	CreatedAt       time.Time                                                    `json:"created_at" api:"required" format:"date-time"`
	Name            string                                                       `json:"name" api:"required"`
	// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
	// every 6 hours or every 2 days. The total interval (frequency × unit) must be
	// between 10 minutes and 1 year.
	Schedule MonitorGetResponseMonitorsPageSemanticMonitorSchedule `json:"schedule" api:"required"`
	// Any of "active", "paused", "failed".
	Status       string                                              `json:"status" api:"required"`
	Target       MonitorGetResponseMonitorsPageSemanticMonitorTarget `json:"target" api:"required"`
	UpdatedAt    time.Time                                           `json:"updated_at" api:"required" format:"date-time"`
	LastChangeAt time.Time                                           `json:"last_change_at" api:"nullable" format:"date-time"`
	LastRunAt    time.Time                                           `json:"last_run_at" api:"nullable" format:"date-time"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags    []string                                             `json:"tags"`
	Webhook MonitorGetResponseMonitorsPageSemanticMonitorWebhook `json:"webhook" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ChangeDetection respjson.Field
		CreatedAt       respjson.Field
		Name            respjson.Field
		Schedule        respjson.Field
		Status          respjson.Field
		Target          respjson.Field
		UpdatedAt       respjson.Field
		LastChangeAt    respjson.Field
		LastRunAt       respjson.Field
		Tags            respjson.Field
		Webhook         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetResponseMonitorsPageSemanticMonitor) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetResponseMonitorsPageSemanticMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detect meaning-level changes that match a natural language query.
type MonitorGetResponseMonitorsPageSemanticMonitorChangeDetection struct {
	Query string `json:"query" api:"required"`
	// Any of "semantic".
	Type                string  `json:"type" api:"required"`
	ConfidenceThreshold float64 `json:"confidence_threshold"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Query               respjson.Field
		Type                respjson.Field
		ConfidenceThreshold respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetResponseMonitorsPageSemanticMonitorChangeDetection) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorGetResponseMonitorsPageSemanticMonitorChangeDetection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
// every 6 hours or every 2 days. The total interval (frequency × unit) must be
// between 10 minutes and 1 year.
type MonitorGetResponseMonitorsPageSemanticMonitorSchedule struct {
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
func (r MonitorGetResponseMonitorsPageSemanticMonitorSchedule) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetResponseMonitorsPageSemanticMonitorSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorGetResponseMonitorsPageSemanticMonitorTarget struct {
	// Any of "page".
	Type string `json:"type" api:"required"`
	URL  string `json:"url" api:"required" format:"uri"`
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
func (r MonitorGetResponseMonitorsPageSemanticMonitorTarget) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetResponseMonitorsPageSemanticMonitorTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorGetResponseMonitorsPageSemanticMonitorWebhook struct {
	// Webhook URL called when a change is detected.
	URL string `json:"url" api:"required" format:"uri"`
	// Signing secret used to verify webhook authenticity. Each delivery includes an
	// `X-Context-Signature: t=<unix>,v1=<hmac>` header, where the HMAC is SHA-256 over
	// `"{t}.{rawRequestBody}"` keyed by this secret. Recompute it with a constant-time
	// compare and reject stale timestamps to prevent replay. Generated by the API;
	// cannot be set by clients.
	Secret string `json:"secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetResponseMonitorsPageSemanticMonitorWebhook) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetResponseMonitorsPageSemanticMonitorWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An extract monitor using semantic change detection.
type MonitorGetResponseMonitorsExtractSemanticMonitor struct {
	ID string `json:"id" api:"required"`
	// Detect meaning-level changes that match a natural language query.
	ChangeDetection MonitorGetResponseMonitorsExtractSemanticMonitorChangeDetection `json:"change_detection" api:"required"`
	CreatedAt       time.Time                                                       `json:"created_at" api:"required" format:"date-time"`
	Name            string                                                          `json:"name" api:"required"`
	// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
	// every 6 hours or every 2 days. The total interval (frequency × unit) must be
	// between 10 minutes and 1 year.
	Schedule MonitorGetResponseMonitorsExtractSemanticMonitorSchedule `json:"schedule" api:"required"`
	// Any of "active", "paused", "failed".
	Status       string                                                 `json:"status" api:"required"`
	Target       MonitorGetResponseMonitorsExtractSemanticMonitorTarget `json:"target" api:"required"`
	UpdatedAt    time.Time                                              `json:"updated_at" api:"required" format:"date-time"`
	LastChangeAt time.Time                                              `json:"last_change_at" api:"nullable" format:"date-time"`
	LastRunAt    time.Time                                              `json:"last_run_at" api:"nullable" format:"date-time"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags    []string                                                `json:"tags"`
	Webhook MonitorGetResponseMonitorsExtractSemanticMonitorWebhook `json:"webhook" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ChangeDetection respjson.Field
		CreatedAt       respjson.Field
		Name            respjson.Field
		Schedule        respjson.Field
		Status          respjson.Field
		Target          respjson.Field
		UpdatedAt       respjson.Field
		LastChangeAt    respjson.Field
		LastRunAt       respjson.Field
		Tags            respjson.Field
		Webhook         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetResponseMonitorsExtractSemanticMonitor) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetResponseMonitorsExtractSemanticMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detect meaning-level changes that match a natural language query.
type MonitorGetResponseMonitorsExtractSemanticMonitorChangeDetection struct {
	Query string `json:"query" api:"required"`
	// Any of "semantic".
	Type                string  `json:"type" api:"required"`
	ConfidenceThreshold float64 `json:"confidence_threshold"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Query               respjson.Field
		Type                respjson.Field
		ConfidenceThreshold respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetResponseMonitorsExtractSemanticMonitorChangeDetection) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorGetResponseMonitorsExtractSemanticMonitorChangeDetection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
// every 6 hours or every 2 days. The total interval (frequency × unit) must be
// between 10 minutes and 1 year.
type MonitorGetResponseMonitorsExtractSemanticMonitorSchedule struct {
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
func (r MonitorGetResponseMonitorsExtractSemanticMonitorSchedule) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetResponseMonitorsExtractSemanticMonitorSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorGetResponseMonitorsExtractSemanticMonitorTarget struct {
	// Any of "extract".
	Type string `json:"type" api:"required"`
	// Root URL to extract structured data from.
	URL              string `json:"url" api:"required" format:"uri"`
	FollowSubdomains bool   `json:"follow_subdomains"`
	// Optional natural-language instructions guiding what to extract.
	Instructions string `json:"instructions"`
	// Optional maximum link depth from the starting URL (0 = only the starting page).
	MaxDepth int64 `json:"max_depth"`
	// Maximum number of pages to analyze during extraction.
	MaxPages int64 `json:"max_pages"`
	// JSON Schema describing the structured data to extract and watch for changes. If
	// omitted, a default summary + key-points schema is used.
	Schema map[string]any `json:"schema"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type             respjson.Field
		URL              respjson.Field
		FollowSubdomains respjson.Field
		Instructions     respjson.Field
		MaxDepth         respjson.Field
		MaxPages         respjson.Field
		Schema           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetResponseMonitorsExtractSemanticMonitorTarget) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetResponseMonitorsExtractSemanticMonitorTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorGetResponseMonitorsExtractSemanticMonitorWebhook struct {
	// Webhook URL called when a change is detected.
	URL string `json:"url" api:"required" format:"uri"`
	// Signing secret used to verify webhook authenticity. Each delivery includes an
	// `X-Context-Signature: t=<unix>,v1=<hmac>` header, where the HMAC is SHA-256 over
	// `"{t}.{rawRequestBody}"` keyed by this secret. Recompute it with a constant-time
	// compare and reject stale timestamps to prevent replay. Generated by the API;
	// cannot be set by clients.
	Secret string `json:"secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetResponseMonitorsExtractSemanticMonitorWebhook) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetResponseMonitorsExtractSemanticMonitorWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MonitorUpdateResponseUnion contains all possible properties and values from
// [MonitorUpdateResponseMonitorsPageExactMonitor],
// [MonitorUpdateResponseMonitorsSitemapExactMonitor],
// [MonitorUpdateResponseMonitorsPageSemanticMonitor],
// [MonitorUpdateResponseMonitorsExtractSemanticMonitor].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MonitorUpdateResponseUnion struct {
	ID string `json:"id"`
	// This field is a union of
	// [MonitorUpdateResponseMonitorsPageExactMonitorChangeDetection],
	// [MonitorUpdateResponseMonitorsSitemapExactMonitorChangeDetection],
	// [MonitorUpdateResponseMonitorsPageSemanticMonitorChangeDetection],
	// [MonitorUpdateResponseMonitorsExtractSemanticMonitorChangeDetection]
	ChangeDetection MonitorUpdateResponseUnionChangeDetection `json:"change_detection"`
	CreatedAt       time.Time                                 `json:"created_at"`
	Name            string                                    `json:"name"`
	// This field is a union of
	// [MonitorUpdateResponseMonitorsPageExactMonitorSchedule],
	// [MonitorUpdateResponseMonitorsSitemapExactMonitorSchedule],
	// [MonitorUpdateResponseMonitorsPageSemanticMonitorSchedule],
	// [MonitorUpdateResponseMonitorsExtractSemanticMonitorSchedule]
	Schedule MonitorUpdateResponseUnionSchedule `json:"schedule"`
	Status   string                             `json:"status"`
	// This field is a union of [MonitorUpdateResponseMonitorsPageExactMonitorTarget],
	// [MonitorUpdateResponseMonitorsSitemapExactMonitorTarget],
	// [MonitorUpdateResponseMonitorsPageSemanticMonitorTarget],
	// [MonitorUpdateResponseMonitorsExtractSemanticMonitorTarget]
	Target       MonitorUpdateResponseUnionTarget `json:"target"`
	UpdatedAt    time.Time                        `json:"updated_at"`
	LastChangeAt time.Time                        `json:"last_change_at"`
	LastRunAt    time.Time                        `json:"last_run_at"`
	Tags         []string                         `json:"tags"`
	// This field is a union of [MonitorUpdateResponseMonitorsPageExactMonitorWebhook],
	// [MonitorUpdateResponseMonitorsSitemapExactMonitorWebhook],
	// [MonitorUpdateResponseMonitorsPageSemanticMonitorWebhook],
	// [MonitorUpdateResponseMonitorsExtractSemanticMonitorWebhook]
	Webhook MonitorUpdateResponseUnionWebhook `json:"webhook"`
	JSON    struct {
		ID              respjson.Field
		ChangeDetection respjson.Field
		CreatedAt       respjson.Field
		Name            respjson.Field
		Schedule        respjson.Field
		Status          respjson.Field
		Target          respjson.Field
		UpdatedAt       respjson.Field
		LastChangeAt    respjson.Field
		LastRunAt       respjson.Field
		Tags            respjson.Field
		Webhook         respjson.Field
		raw             string
	} `json:"-"`
}

func (u MonitorUpdateResponseUnion) AsMonitorUpdateResponseMonitorsPageExactMonitor() (v MonitorUpdateResponseMonitorsPageExactMonitor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorUpdateResponseUnion) AsMonitorUpdateResponseMonitorsSitemapExactMonitor() (v MonitorUpdateResponseMonitorsSitemapExactMonitor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorUpdateResponseUnion) AsMonitorUpdateResponseMonitorsPageSemanticMonitor() (v MonitorUpdateResponseMonitorsPageSemanticMonitor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorUpdateResponseUnion) AsMonitorUpdateResponseMonitorsExtractSemanticMonitor() (v MonitorUpdateResponseMonitorsExtractSemanticMonitor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MonitorUpdateResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *MonitorUpdateResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MonitorUpdateResponseUnionChangeDetection is an implicit subunion of
// [MonitorUpdateResponseUnion]. MonitorUpdateResponseUnionChangeDetection provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MonitorUpdateResponseUnion].
type MonitorUpdateResponseUnionChangeDetection struct {
	Type                string  `json:"type"`
	Query               string  `json:"query"`
	ConfidenceThreshold float64 `json:"confidence_threshold"`
	JSON                struct {
		Type                respjson.Field
		Query               respjson.Field
		ConfidenceThreshold respjson.Field
		raw                 string
	} `json:"-"`
}

func (r *MonitorUpdateResponseUnionChangeDetection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MonitorUpdateResponseUnionSchedule is an implicit subunion of
// [MonitorUpdateResponseUnion]. MonitorUpdateResponseUnionSchedule provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MonitorUpdateResponseUnion].
type MonitorUpdateResponseUnionSchedule struct {
	Frequency int64  `json:"frequency"`
	Type      string `json:"type"`
	Unit      string `json:"unit"`
	JSON      struct {
		Frequency respjson.Field
		Type      respjson.Field
		Unit      respjson.Field
		raw       string
	} `json:"-"`
}

func (r *MonitorUpdateResponseUnionSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MonitorUpdateResponseUnionTarget is an implicit subunion of
// [MonitorUpdateResponseUnion]. MonitorUpdateResponseUnionTarget provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MonitorUpdateResponseUnion].
type MonitorUpdateResponseUnionTarget struct {
	Type                string `json:"type"`
	URL                 string `json:"url"`
	NormalizeWhitespace bool   `json:"normalize_whitespace"`
	// This field is from variant
	// [MonitorUpdateResponseMonitorsSitemapExactMonitorTarget].
	Exclude []string `json:"exclude"`
	// This field is from variant
	// [MonitorUpdateResponseMonitorsSitemapExactMonitorTarget].
	Include []string `json:"include"`
	// This field is from variant
	// [MonitorUpdateResponseMonitorsSitemapExactMonitorTarget].
	MaxURLs int64 `json:"max_urls"`
	// This field is from variant
	// [MonitorUpdateResponseMonitorsExtractSemanticMonitorTarget].
	FollowSubdomains bool `json:"follow_subdomains"`
	// This field is from variant
	// [MonitorUpdateResponseMonitorsExtractSemanticMonitorTarget].
	Instructions string `json:"instructions"`
	// This field is from variant
	// [MonitorUpdateResponseMonitorsExtractSemanticMonitorTarget].
	MaxDepth int64 `json:"max_depth"`
	// This field is from variant
	// [MonitorUpdateResponseMonitorsExtractSemanticMonitorTarget].
	MaxPages int64 `json:"max_pages"`
	// This field is from variant
	// [MonitorUpdateResponseMonitorsExtractSemanticMonitorTarget].
	Schema map[string]any `json:"schema"`
	JSON   struct {
		Type                respjson.Field
		URL                 respjson.Field
		NormalizeWhitespace respjson.Field
		Exclude             respjson.Field
		Include             respjson.Field
		MaxURLs             respjson.Field
		FollowSubdomains    respjson.Field
		Instructions        respjson.Field
		MaxDepth            respjson.Field
		MaxPages            respjson.Field
		Schema              respjson.Field
		raw                 string
	} `json:"-"`
}

func (r *MonitorUpdateResponseUnionTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MonitorUpdateResponseUnionWebhook is an implicit subunion of
// [MonitorUpdateResponseUnion]. MonitorUpdateResponseUnionWebhook provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MonitorUpdateResponseUnion].
type MonitorUpdateResponseUnionWebhook struct {
	URL    string `json:"url"`
	Secret string `json:"secret"`
	JSON   struct {
		URL    respjson.Field
		Secret respjson.Field
		raw    string
	} `json:"-"`
}

func (r *MonitorUpdateResponseUnionWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A page monitor using exact change detection.
type MonitorUpdateResponseMonitorsPageExactMonitor struct {
	ID string `json:"id" api:"required"`
	// Detect exact changes. For page targets, this means visible text diffs. For
	// sitemap targets, this means URL additions and removals.
	ChangeDetection MonitorUpdateResponseMonitorsPageExactMonitorChangeDetection `json:"change_detection" api:"required"`
	CreatedAt       time.Time                                                    `json:"created_at" api:"required" format:"date-time"`
	Name            string                                                       `json:"name" api:"required"`
	// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
	// every 6 hours or every 2 days. The total interval (frequency × unit) must be
	// between 10 minutes and 1 year.
	Schedule MonitorUpdateResponseMonitorsPageExactMonitorSchedule `json:"schedule" api:"required"`
	// Any of "active", "paused", "failed".
	Status       string                                              `json:"status" api:"required"`
	Target       MonitorUpdateResponseMonitorsPageExactMonitorTarget `json:"target" api:"required"`
	UpdatedAt    time.Time                                           `json:"updated_at" api:"required" format:"date-time"`
	LastChangeAt time.Time                                           `json:"last_change_at" api:"nullable" format:"date-time"`
	LastRunAt    time.Time                                           `json:"last_run_at" api:"nullable" format:"date-time"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags    []string                                             `json:"tags"`
	Webhook MonitorUpdateResponseMonitorsPageExactMonitorWebhook `json:"webhook" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ChangeDetection respjson.Field
		CreatedAt       respjson.Field
		Name            respjson.Field
		Schedule        respjson.Field
		Status          respjson.Field
		Target          respjson.Field
		UpdatedAt       respjson.Field
		LastChangeAt    respjson.Field
		LastRunAt       respjson.Field
		Tags            respjson.Field
		Webhook         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorUpdateResponseMonitorsPageExactMonitor) RawJSON() string { return r.JSON.raw }
func (r *MonitorUpdateResponseMonitorsPageExactMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detect exact changes. For page targets, this means visible text diffs. For
// sitemap targets, this means URL additions and removals.
type MonitorUpdateResponseMonitorsPageExactMonitorChangeDetection struct {
	// Any of "exact".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorUpdateResponseMonitorsPageExactMonitorChangeDetection) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorUpdateResponseMonitorsPageExactMonitorChangeDetection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
// every 6 hours or every 2 days. The total interval (frequency × unit) must be
// between 10 minutes and 1 year.
type MonitorUpdateResponseMonitorsPageExactMonitorSchedule struct {
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
func (r MonitorUpdateResponseMonitorsPageExactMonitorSchedule) RawJSON() string { return r.JSON.raw }
func (r *MonitorUpdateResponseMonitorsPageExactMonitorSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorUpdateResponseMonitorsPageExactMonitorTarget struct {
	// Any of "page".
	Type string `json:"type" api:"required"`
	URL  string `json:"url" api:"required" format:"uri"`
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
func (r MonitorUpdateResponseMonitorsPageExactMonitorTarget) RawJSON() string { return r.JSON.raw }
func (r *MonitorUpdateResponseMonitorsPageExactMonitorTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorUpdateResponseMonitorsPageExactMonitorWebhook struct {
	// Webhook URL called when a change is detected.
	URL string `json:"url" api:"required" format:"uri"`
	// Signing secret used to verify webhook authenticity. Each delivery includes an
	// `X-Context-Signature: t=<unix>,v1=<hmac>` header, where the HMAC is SHA-256 over
	// `"{t}.{rawRequestBody}"` keyed by this secret. Recompute it with a constant-time
	// compare and reject stale timestamps to prevent replay. Generated by the API;
	// cannot be set by clients.
	Secret string `json:"secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorUpdateResponseMonitorsPageExactMonitorWebhook) RawJSON() string { return r.JSON.raw }
func (r *MonitorUpdateResponseMonitorsPageExactMonitorWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A sitemap monitor using exact change detection.
type MonitorUpdateResponseMonitorsSitemapExactMonitor struct {
	ID string `json:"id" api:"required"`
	// Detect exact changes. For page targets, this means visible text diffs. For
	// sitemap targets, this means URL additions and removals.
	ChangeDetection MonitorUpdateResponseMonitorsSitemapExactMonitorChangeDetection `json:"change_detection" api:"required"`
	CreatedAt       time.Time                                                       `json:"created_at" api:"required" format:"date-time"`
	Name            string                                                          `json:"name" api:"required"`
	// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
	// every 6 hours or every 2 days. The total interval (frequency × unit) must be
	// between 10 minutes and 1 year.
	Schedule MonitorUpdateResponseMonitorsSitemapExactMonitorSchedule `json:"schedule" api:"required"`
	// Any of "active", "paused", "failed".
	Status       string                                                 `json:"status" api:"required"`
	Target       MonitorUpdateResponseMonitorsSitemapExactMonitorTarget `json:"target" api:"required"`
	UpdatedAt    time.Time                                              `json:"updated_at" api:"required" format:"date-time"`
	LastChangeAt time.Time                                              `json:"last_change_at" api:"nullable" format:"date-time"`
	LastRunAt    time.Time                                              `json:"last_run_at" api:"nullable" format:"date-time"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags    []string                                                `json:"tags"`
	Webhook MonitorUpdateResponseMonitorsSitemapExactMonitorWebhook `json:"webhook" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ChangeDetection respjson.Field
		CreatedAt       respjson.Field
		Name            respjson.Field
		Schedule        respjson.Field
		Status          respjson.Field
		Target          respjson.Field
		UpdatedAt       respjson.Field
		LastChangeAt    respjson.Field
		LastRunAt       respjson.Field
		Tags            respjson.Field
		Webhook         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorUpdateResponseMonitorsSitemapExactMonitor) RawJSON() string { return r.JSON.raw }
func (r *MonitorUpdateResponseMonitorsSitemapExactMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detect exact changes. For page targets, this means visible text diffs. For
// sitemap targets, this means URL additions and removals.
type MonitorUpdateResponseMonitorsSitemapExactMonitorChangeDetection struct {
	// Any of "exact".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorUpdateResponseMonitorsSitemapExactMonitorChangeDetection) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorUpdateResponseMonitorsSitemapExactMonitorChangeDetection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
// every 6 hours or every 2 days. The total interval (frequency × unit) must be
// between 10 minutes and 1 year.
type MonitorUpdateResponseMonitorsSitemapExactMonitorSchedule struct {
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
func (r MonitorUpdateResponseMonitorsSitemapExactMonitorSchedule) RawJSON() string { return r.JSON.raw }
func (r *MonitorUpdateResponseMonitorsSitemapExactMonitorSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorUpdateResponseMonitorsSitemapExactMonitorTarget struct {
	// Any of "sitemap".
	Type string `json:"type" api:"required"`
	// Sitemap URL to monitor.
	URL string `json:"url" api:"required" format:"uri"`
	// URL path patterns to exclude.
	Exclude []string `json:"exclude"`
	// URL path patterns to include.
	Include []string `json:"include"`
	MaxURLs int64    `json:"max_urls"`
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
func (r MonitorUpdateResponseMonitorsSitemapExactMonitorTarget) RawJSON() string { return r.JSON.raw }
func (r *MonitorUpdateResponseMonitorsSitemapExactMonitorTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorUpdateResponseMonitorsSitemapExactMonitorWebhook struct {
	// Webhook URL called when a change is detected.
	URL string `json:"url" api:"required" format:"uri"`
	// Signing secret used to verify webhook authenticity. Each delivery includes an
	// `X-Context-Signature: t=<unix>,v1=<hmac>` header, where the HMAC is SHA-256 over
	// `"{t}.{rawRequestBody}"` keyed by this secret. Recompute it with a constant-time
	// compare and reject stale timestamps to prevent replay. Generated by the API;
	// cannot be set by clients.
	Secret string `json:"secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorUpdateResponseMonitorsSitemapExactMonitorWebhook) RawJSON() string { return r.JSON.raw }
func (r *MonitorUpdateResponseMonitorsSitemapExactMonitorWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A page monitor using semantic change detection.
type MonitorUpdateResponseMonitorsPageSemanticMonitor struct {
	ID string `json:"id" api:"required"`
	// Detect meaning-level changes that match a natural language query.
	ChangeDetection MonitorUpdateResponseMonitorsPageSemanticMonitorChangeDetection `json:"change_detection" api:"required"`
	CreatedAt       time.Time                                                       `json:"created_at" api:"required" format:"date-time"`
	Name            string                                                          `json:"name" api:"required"`
	// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
	// every 6 hours or every 2 days. The total interval (frequency × unit) must be
	// between 10 minutes and 1 year.
	Schedule MonitorUpdateResponseMonitorsPageSemanticMonitorSchedule `json:"schedule" api:"required"`
	// Any of "active", "paused", "failed".
	Status       string                                                 `json:"status" api:"required"`
	Target       MonitorUpdateResponseMonitorsPageSemanticMonitorTarget `json:"target" api:"required"`
	UpdatedAt    time.Time                                              `json:"updated_at" api:"required" format:"date-time"`
	LastChangeAt time.Time                                              `json:"last_change_at" api:"nullable" format:"date-time"`
	LastRunAt    time.Time                                              `json:"last_run_at" api:"nullable" format:"date-time"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags    []string                                                `json:"tags"`
	Webhook MonitorUpdateResponseMonitorsPageSemanticMonitorWebhook `json:"webhook" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ChangeDetection respjson.Field
		CreatedAt       respjson.Field
		Name            respjson.Field
		Schedule        respjson.Field
		Status          respjson.Field
		Target          respjson.Field
		UpdatedAt       respjson.Field
		LastChangeAt    respjson.Field
		LastRunAt       respjson.Field
		Tags            respjson.Field
		Webhook         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorUpdateResponseMonitorsPageSemanticMonitor) RawJSON() string { return r.JSON.raw }
func (r *MonitorUpdateResponseMonitorsPageSemanticMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detect meaning-level changes that match a natural language query.
type MonitorUpdateResponseMonitorsPageSemanticMonitorChangeDetection struct {
	Query string `json:"query" api:"required"`
	// Any of "semantic".
	Type                string  `json:"type" api:"required"`
	ConfidenceThreshold float64 `json:"confidence_threshold"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Query               respjson.Field
		Type                respjson.Field
		ConfidenceThreshold respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorUpdateResponseMonitorsPageSemanticMonitorChangeDetection) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorUpdateResponseMonitorsPageSemanticMonitorChangeDetection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
// every 6 hours or every 2 days. The total interval (frequency × unit) must be
// between 10 minutes and 1 year.
type MonitorUpdateResponseMonitorsPageSemanticMonitorSchedule struct {
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
func (r MonitorUpdateResponseMonitorsPageSemanticMonitorSchedule) RawJSON() string { return r.JSON.raw }
func (r *MonitorUpdateResponseMonitorsPageSemanticMonitorSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorUpdateResponseMonitorsPageSemanticMonitorTarget struct {
	// Any of "page".
	Type string `json:"type" api:"required"`
	URL  string `json:"url" api:"required" format:"uri"`
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
func (r MonitorUpdateResponseMonitorsPageSemanticMonitorTarget) RawJSON() string { return r.JSON.raw }
func (r *MonitorUpdateResponseMonitorsPageSemanticMonitorTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorUpdateResponseMonitorsPageSemanticMonitorWebhook struct {
	// Webhook URL called when a change is detected.
	URL string `json:"url" api:"required" format:"uri"`
	// Signing secret used to verify webhook authenticity. Each delivery includes an
	// `X-Context-Signature: t=<unix>,v1=<hmac>` header, where the HMAC is SHA-256 over
	// `"{t}.{rawRequestBody}"` keyed by this secret. Recompute it with a constant-time
	// compare and reject stale timestamps to prevent replay. Generated by the API;
	// cannot be set by clients.
	Secret string `json:"secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorUpdateResponseMonitorsPageSemanticMonitorWebhook) RawJSON() string { return r.JSON.raw }
func (r *MonitorUpdateResponseMonitorsPageSemanticMonitorWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An extract monitor using semantic change detection.
type MonitorUpdateResponseMonitorsExtractSemanticMonitor struct {
	ID string `json:"id" api:"required"`
	// Detect meaning-level changes that match a natural language query.
	ChangeDetection MonitorUpdateResponseMonitorsExtractSemanticMonitorChangeDetection `json:"change_detection" api:"required"`
	CreatedAt       time.Time                                                          `json:"created_at" api:"required" format:"date-time"`
	Name            string                                                             `json:"name" api:"required"`
	// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
	// every 6 hours or every 2 days. The total interval (frequency × unit) must be
	// between 10 minutes and 1 year.
	Schedule MonitorUpdateResponseMonitorsExtractSemanticMonitorSchedule `json:"schedule" api:"required"`
	// Any of "active", "paused", "failed".
	Status       string                                                    `json:"status" api:"required"`
	Target       MonitorUpdateResponseMonitorsExtractSemanticMonitorTarget `json:"target" api:"required"`
	UpdatedAt    time.Time                                                 `json:"updated_at" api:"required" format:"date-time"`
	LastChangeAt time.Time                                                 `json:"last_change_at" api:"nullable" format:"date-time"`
	LastRunAt    time.Time                                                 `json:"last_run_at" api:"nullable" format:"date-time"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags    []string                                                   `json:"tags"`
	Webhook MonitorUpdateResponseMonitorsExtractSemanticMonitorWebhook `json:"webhook" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ChangeDetection respjson.Field
		CreatedAt       respjson.Field
		Name            respjson.Field
		Schedule        respjson.Field
		Status          respjson.Field
		Target          respjson.Field
		UpdatedAt       respjson.Field
		LastChangeAt    respjson.Field
		LastRunAt       respjson.Field
		Tags            respjson.Field
		Webhook         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorUpdateResponseMonitorsExtractSemanticMonitor) RawJSON() string { return r.JSON.raw }
func (r *MonitorUpdateResponseMonitorsExtractSemanticMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detect meaning-level changes that match a natural language query.
type MonitorUpdateResponseMonitorsExtractSemanticMonitorChangeDetection struct {
	Query string `json:"query" api:"required"`
	// Any of "semantic".
	Type                string  `json:"type" api:"required"`
	ConfidenceThreshold float64 `json:"confidence_threshold"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Query               respjson.Field
		Type                respjson.Field
		ConfidenceThreshold respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorUpdateResponseMonitorsExtractSemanticMonitorChangeDetection) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorUpdateResponseMonitorsExtractSemanticMonitorChangeDetection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
// every 6 hours or every 2 days. The total interval (frequency × unit) must be
// between 10 minutes and 1 year.
type MonitorUpdateResponseMonitorsExtractSemanticMonitorSchedule struct {
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
func (r MonitorUpdateResponseMonitorsExtractSemanticMonitorSchedule) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorUpdateResponseMonitorsExtractSemanticMonitorSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorUpdateResponseMonitorsExtractSemanticMonitorTarget struct {
	// Any of "extract".
	Type string `json:"type" api:"required"`
	// Root URL to extract structured data from.
	URL              string `json:"url" api:"required" format:"uri"`
	FollowSubdomains bool   `json:"follow_subdomains"`
	// Optional natural-language instructions guiding what to extract.
	Instructions string `json:"instructions"`
	// Optional maximum link depth from the starting URL (0 = only the starting page).
	MaxDepth int64 `json:"max_depth"`
	// Maximum number of pages to analyze during extraction.
	MaxPages int64 `json:"max_pages"`
	// JSON Schema describing the structured data to extract and watch for changes. If
	// omitted, a default summary + key-points schema is used.
	Schema map[string]any `json:"schema"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type             respjson.Field
		URL              respjson.Field
		FollowSubdomains respjson.Field
		Instructions     respjson.Field
		MaxDepth         respjson.Field
		MaxPages         respjson.Field
		Schema           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorUpdateResponseMonitorsExtractSemanticMonitorTarget) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorUpdateResponseMonitorsExtractSemanticMonitorTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorUpdateResponseMonitorsExtractSemanticMonitorWebhook struct {
	// Webhook URL called when a change is detected.
	URL string `json:"url" api:"required" format:"uri"`
	// Signing secret used to verify webhook authenticity. Each delivery includes an
	// `X-Context-Signature: t=<unix>,v1=<hmac>` header, where the HMAC is SHA-256 over
	// `"{t}.{rawRequestBody}"` keyed by this secret. Recompute it with a constant-time
	// compare and reject stale timestamps to prevent replay. Generated by the API;
	// cannot be set by clients.
	Secret string `json:"secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorUpdateResponseMonitorsExtractSemanticMonitorWebhook) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorUpdateResponseMonitorsExtractSemanticMonitorWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListResponse struct {
	Data       []MonitorListResponseDataUnion `json:"data" api:"required"`
	HasMore    bool                           `json:"has_more" api:"required"`
	NextCursor string                         `json:"next_cursor" api:"required"`
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

// MonitorListResponseDataUnion contains all possible properties and values from
// [MonitorListResponseDataMonitorsPageExactMonitor],
// [MonitorListResponseDataMonitorsSitemapExactMonitor],
// [MonitorListResponseDataMonitorsPageSemanticMonitor],
// [MonitorListResponseDataMonitorsExtractSemanticMonitor].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MonitorListResponseDataUnion struct {
	ID string `json:"id"`
	// This field is a union of
	// [MonitorListResponseDataMonitorsPageExactMonitorChangeDetection],
	// [MonitorListResponseDataMonitorsSitemapExactMonitorChangeDetection],
	// [MonitorListResponseDataMonitorsPageSemanticMonitorChangeDetection],
	// [MonitorListResponseDataMonitorsExtractSemanticMonitorChangeDetection]
	ChangeDetection MonitorListResponseDataUnionChangeDetection `json:"change_detection"`
	CreatedAt       time.Time                                   `json:"created_at"`
	Name            string                                      `json:"name"`
	// This field is a union of
	// [MonitorListResponseDataMonitorsPageExactMonitorSchedule],
	// [MonitorListResponseDataMonitorsSitemapExactMonitorSchedule],
	// [MonitorListResponseDataMonitorsPageSemanticMonitorSchedule],
	// [MonitorListResponseDataMonitorsExtractSemanticMonitorSchedule]
	Schedule MonitorListResponseDataUnionSchedule `json:"schedule"`
	Status   string                               `json:"status"`
	// This field is a union of
	// [MonitorListResponseDataMonitorsPageExactMonitorTarget],
	// [MonitorListResponseDataMonitorsSitemapExactMonitorTarget],
	// [MonitorListResponseDataMonitorsPageSemanticMonitorTarget],
	// [MonitorListResponseDataMonitorsExtractSemanticMonitorTarget]
	Target       MonitorListResponseDataUnionTarget `json:"target"`
	UpdatedAt    time.Time                          `json:"updated_at"`
	LastChangeAt time.Time                          `json:"last_change_at"`
	LastRunAt    time.Time                          `json:"last_run_at"`
	Tags         []string                           `json:"tags"`
	// This field is a union of
	// [MonitorListResponseDataMonitorsPageExactMonitorWebhook],
	// [MonitorListResponseDataMonitorsSitemapExactMonitorWebhook],
	// [MonitorListResponseDataMonitorsPageSemanticMonitorWebhook],
	// [MonitorListResponseDataMonitorsExtractSemanticMonitorWebhook]
	Webhook MonitorListResponseDataUnionWebhook `json:"webhook"`
	JSON    struct {
		ID              respjson.Field
		ChangeDetection respjson.Field
		CreatedAt       respjson.Field
		Name            respjson.Field
		Schedule        respjson.Field
		Status          respjson.Field
		Target          respjson.Field
		UpdatedAt       respjson.Field
		LastChangeAt    respjson.Field
		LastRunAt       respjson.Field
		Tags            respjson.Field
		Webhook         respjson.Field
		raw             string
	} `json:"-"`
}

func (u MonitorListResponseDataUnion) AsMonitorListResponseDataMonitorsPageExactMonitor() (v MonitorListResponseDataMonitorsPageExactMonitor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorListResponseDataUnion) AsMonitorListResponseDataMonitorsSitemapExactMonitor() (v MonitorListResponseDataMonitorsSitemapExactMonitor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorListResponseDataUnion) AsMonitorListResponseDataMonitorsPageSemanticMonitor() (v MonitorListResponseDataMonitorsPageSemanticMonitor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorListResponseDataUnion) AsMonitorListResponseDataMonitorsExtractSemanticMonitor() (v MonitorListResponseDataMonitorsExtractSemanticMonitor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MonitorListResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *MonitorListResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MonitorListResponseDataUnionChangeDetection is an implicit subunion of
// [MonitorListResponseDataUnion]. MonitorListResponseDataUnionChangeDetection
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MonitorListResponseDataUnion].
type MonitorListResponseDataUnionChangeDetection struct {
	Type                string  `json:"type"`
	Query               string  `json:"query"`
	ConfidenceThreshold float64 `json:"confidence_threshold"`
	JSON                struct {
		Type                respjson.Field
		Query               respjson.Field
		ConfidenceThreshold respjson.Field
		raw                 string
	} `json:"-"`
}

func (r *MonitorListResponseDataUnionChangeDetection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MonitorListResponseDataUnionSchedule is an implicit subunion of
// [MonitorListResponseDataUnion]. MonitorListResponseDataUnionSchedule provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MonitorListResponseDataUnion].
type MonitorListResponseDataUnionSchedule struct {
	Frequency int64  `json:"frequency"`
	Type      string `json:"type"`
	Unit      string `json:"unit"`
	JSON      struct {
		Frequency respjson.Field
		Type      respjson.Field
		Unit      respjson.Field
		raw       string
	} `json:"-"`
}

func (r *MonitorListResponseDataUnionSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MonitorListResponseDataUnionTarget is an implicit subunion of
// [MonitorListResponseDataUnion]. MonitorListResponseDataUnionTarget provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MonitorListResponseDataUnion].
type MonitorListResponseDataUnionTarget struct {
	Type                string `json:"type"`
	URL                 string `json:"url"`
	NormalizeWhitespace bool   `json:"normalize_whitespace"`
	// This field is from variant
	// [MonitorListResponseDataMonitorsSitemapExactMonitorTarget].
	Exclude []string `json:"exclude"`
	// This field is from variant
	// [MonitorListResponseDataMonitorsSitemapExactMonitorTarget].
	Include []string `json:"include"`
	// This field is from variant
	// [MonitorListResponseDataMonitorsSitemapExactMonitorTarget].
	MaxURLs int64 `json:"max_urls"`
	// This field is from variant
	// [MonitorListResponseDataMonitorsExtractSemanticMonitorTarget].
	FollowSubdomains bool `json:"follow_subdomains"`
	// This field is from variant
	// [MonitorListResponseDataMonitorsExtractSemanticMonitorTarget].
	Instructions string `json:"instructions"`
	// This field is from variant
	// [MonitorListResponseDataMonitorsExtractSemanticMonitorTarget].
	MaxDepth int64 `json:"max_depth"`
	// This field is from variant
	// [MonitorListResponseDataMonitorsExtractSemanticMonitorTarget].
	MaxPages int64 `json:"max_pages"`
	// This field is from variant
	// [MonitorListResponseDataMonitorsExtractSemanticMonitorTarget].
	Schema map[string]any `json:"schema"`
	JSON   struct {
		Type                respjson.Field
		URL                 respjson.Field
		NormalizeWhitespace respjson.Field
		Exclude             respjson.Field
		Include             respjson.Field
		MaxURLs             respjson.Field
		FollowSubdomains    respjson.Field
		Instructions        respjson.Field
		MaxDepth            respjson.Field
		MaxPages            respjson.Field
		Schema              respjson.Field
		raw                 string
	} `json:"-"`
}

func (r *MonitorListResponseDataUnionTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MonitorListResponseDataUnionWebhook is an implicit subunion of
// [MonitorListResponseDataUnion]. MonitorListResponseDataUnionWebhook provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MonitorListResponseDataUnion].
type MonitorListResponseDataUnionWebhook struct {
	URL    string `json:"url"`
	Secret string `json:"secret"`
	JSON   struct {
		URL    respjson.Field
		Secret respjson.Field
		raw    string
	} `json:"-"`
}

func (r *MonitorListResponseDataUnionWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A page monitor using exact change detection.
type MonitorListResponseDataMonitorsPageExactMonitor struct {
	ID string `json:"id" api:"required"`
	// Detect exact changes. For page targets, this means visible text diffs. For
	// sitemap targets, this means URL additions and removals.
	ChangeDetection MonitorListResponseDataMonitorsPageExactMonitorChangeDetection `json:"change_detection" api:"required"`
	CreatedAt       time.Time                                                      `json:"created_at" api:"required" format:"date-time"`
	Name            string                                                         `json:"name" api:"required"`
	// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
	// every 6 hours or every 2 days. The total interval (frequency × unit) must be
	// between 10 minutes and 1 year.
	Schedule MonitorListResponseDataMonitorsPageExactMonitorSchedule `json:"schedule" api:"required"`
	// Any of "active", "paused", "failed".
	Status       string                                                `json:"status" api:"required"`
	Target       MonitorListResponseDataMonitorsPageExactMonitorTarget `json:"target" api:"required"`
	UpdatedAt    time.Time                                             `json:"updated_at" api:"required" format:"date-time"`
	LastChangeAt time.Time                                             `json:"last_change_at" api:"nullable" format:"date-time"`
	LastRunAt    time.Time                                             `json:"last_run_at" api:"nullable" format:"date-time"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags    []string                                               `json:"tags"`
	Webhook MonitorListResponseDataMonitorsPageExactMonitorWebhook `json:"webhook" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ChangeDetection respjson.Field
		CreatedAt       respjson.Field
		Name            respjson.Field
		Schedule        respjson.Field
		Status          respjson.Field
		Target          respjson.Field
		UpdatedAt       respjson.Field
		LastChangeAt    respjson.Field
		LastRunAt       respjson.Field
		Tags            respjson.Field
		Webhook         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListResponseDataMonitorsPageExactMonitor) RawJSON() string { return r.JSON.raw }
func (r *MonitorListResponseDataMonitorsPageExactMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detect exact changes. For page targets, this means visible text diffs. For
// sitemap targets, this means URL additions and removals.
type MonitorListResponseDataMonitorsPageExactMonitorChangeDetection struct {
	// Any of "exact".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListResponseDataMonitorsPageExactMonitorChangeDetection) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorListResponseDataMonitorsPageExactMonitorChangeDetection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
// every 6 hours or every 2 days. The total interval (frequency × unit) must be
// between 10 minutes and 1 year.
type MonitorListResponseDataMonitorsPageExactMonitorSchedule struct {
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
func (r MonitorListResponseDataMonitorsPageExactMonitorSchedule) RawJSON() string { return r.JSON.raw }
func (r *MonitorListResponseDataMonitorsPageExactMonitorSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListResponseDataMonitorsPageExactMonitorTarget struct {
	// Any of "page".
	Type string `json:"type" api:"required"`
	URL  string `json:"url" api:"required" format:"uri"`
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
func (r MonitorListResponseDataMonitorsPageExactMonitorTarget) RawJSON() string { return r.JSON.raw }
func (r *MonitorListResponseDataMonitorsPageExactMonitorTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListResponseDataMonitorsPageExactMonitorWebhook struct {
	// Webhook URL called when a change is detected.
	URL string `json:"url" api:"required" format:"uri"`
	// Signing secret used to verify webhook authenticity. Each delivery includes an
	// `X-Context-Signature: t=<unix>,v1=<hmac>` header, where the HMAC is SHA-256 over
	// `"{t}.{rawRequestBody}"` keyed by this secret. Recompute it with a constant-time
	// compare and reject stale timestamps to prevent replay. Generated by the API;
	// cannot be set by clients.
	Secret string `json:"secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListResponseDataMonitorsPageExactMonitorWebhook) RawJSON() string { return r.JSON.raw }
func (r *MonitorListResponseDataMonitorsPageExactMonitorWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A sitemap monitor using exact change detection.
type MonitorListResponseDataMonitorsSitemapExactMonitor struct {
	ID string `json:"id" api:"required"`
	// Detect exact changes. For page targets, this means visible text diffs. For
	// sitemap targets, this means URL additions and removals.
	ChangeDetection MonitorListResponseDataMonitorsSitemapExactMonitorChangeDetection `json:"change_detection" api:"required"`
	CreatedAt       time.Time                                                         `json:"created_at" api:"required" format:"date-time"`
	Name            string                                                            `json:"name" api:"required"`
	// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
	// every 6 hours or every 2 days. The total interval (frequency × unit) must be
	// between 10 minutes and 1 year.
	Schedule MonitorListResponseDataMonitorsSitemapExactMonitorSchedule `json:"schedule" api:"required"`
	// Any of "active", "paused", "failed".
	Status       string                                                   `json:"status" api:"required"`
	Target       MonitorListResponseDataMonitorsSitemapExactMonitorTarget `json:"target" api:"required"`
	UpdatedAt    time.Time                                                `json:"updated_at" api:"required" format:"date-time"`
	LastChangeAt time.Time                                                `json:"last_change_at" api:"nullable" format:"date-time"`
	LastRunAt    time.Time                                                `json:"last_run_at" api:"nullable" format:"date-time"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags    []string                                                  `json:"tags"`
	Webhook MonitorListResponseDataMonitorsSitemapExactMonitorWebhook `json:"webhook" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ChangeDetection respjson.Field
		CreatedAt       respjson.Field
		Name            respjson.Field
		Schedule        respjson.Field
		Status          respjson.Field
		Target          respjson.Field
		UpdatedAt       respjson.Field
		LastChangeAt    respjson.Field
		LastRunAt       respjson.Field
		Tags            respjson.Field
		Webhook         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListResponseDataMonitorsSitemapExactMonitor) RawJSON() string { return r.JSON.raw }
func (r *MonitorListResponseDataMonitorsSitemapExactMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detect exact changes. For page targets, this means visible text diffs. For
// sitemap targets, this means URL additions and removals.
type MonitorListResponseDataMonitorsSitemapExactMonitorChangeDetection struct {
	// Any of "exact".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListResponseDataMonitorsSitemapExactMonitorChangeDetection) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorListResponseDataMonitorsSitemapExactMonitorChangeDetection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
// every 6 hours or every 2 days. The total interval (frequency × unit) must be
// between 10 minutes and 1 year.
type MonitorListResponseDataMonitorsSitemapExactMonitorSchedule struct {
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
func (r MonitorListResponseDataMonitorsSitemapExactMonitorSchedule) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorListResponseDataMonitorsSitemapExactMonitorSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListResponseDataMonitorsSitemapExactMonitorTarget struct {
	// Any of "sitemap".
	Type string `json:"type" api:"required"`
	// Sitemap URL to monitor.
	URL string `json:"url" api:"required" format:"uri"`
	// URL path patterns to exclude.
	Exclude []string `json:"exclude"`
	// URL path patterns to include.
	Include []string `json:"include"`
	MaxURLs int64    `json:"max_urls"`
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
func (r MonitorListResponseDataMonitorsSitemapExactMonitorTarget) RawJSON() string { return r.JSON.raw }
func (r *MonitorListResponseDataMonitorsSitemapExactMonitorTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListResponseDataMonitorsSitemapExactMonitorWebhook struct {
	// Webhook URL called when a change is detected.
	URL string `json:"url" api:"required" format:"uri"`
	// Signing secret used to verify webhook authenticity. Each delivery includes an
	// `X-Context-Signature: t=<unix>,v1=<hmac>` header, where the HMAC is SHA-256 over
	// `"{t}.{rawRequestBody}"` keyed by this secret. Recompute it with a constant-time
	// compare and reject stale timestamps to prevent replay. Generated by the API;
	// cannot be set by clients.
	Secret string `json:"secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListResponseDataMonitorsSitemapExactMonitorWebhook) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorListResponseDataMonitorsSitemapExactMonitorWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A page monitor using semantic change detection.
type MonitorListResponseDataMonitorsPageSemanticMonitor struct {
	ID string `json:"id" api:"required"`
	// Detect meaning-level changes that match a natural language query.
	ChangeDetection MonitorListResponseDataMonitorsPageSemanticMonitorChangeDetection `json:"change_detection" api:"required"`
	CreatedAt       time.Time                                                         `json:"created_at" api:"required" format:"date-time"`
	Name            string                                                            `json:"name" api:"required"`
	// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
	// every 6 hours or every 2 days. The total interval (frequency × unit) must be
	// between 10 minutes and 1 year.
	Schedule MonitorListResponseDataMonitorsPageSemanticMonitorSchedule `json:"schedule" api:"required"`
	// Any of "active", "paused", "failed".
	Status       string                                                   `json:"status" api:"required"`
	Target       MonitorListResponseDataMonitorsPageSemanticMonitorTarget `json:"target" api:"required"`
	UpdatedAt    time.Time                                                `json:"updated_at" api:"required" format:"date-time"`
	LastChangeAt time.Time                                                `json:"last_change_at" api:"nullable" format:"date-time"`
	LastRunAt    time.Time                                                `json:"last_run_at" api:"nullable" format:"date-time"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags    []string                                                  `json:"tags"`
	Webhook MonitorListResponseDataMonitorsPageSemanticMonitorWebhook `json:"webhook" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ChangeDetection respjson.Field
		CreatedAt       respjson.Field
		Name            respjson.Field
		Schedule        respjson.Field
		Status          respjson.Field
		Target          respjson.Field
		UpdatedAt       respjson.Field
		LastChangeAt    respjson.Field
		LastRunAt       respjson.Field
		Tags            respjson.Field
		Webhook         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListResponseDataMonitorsPageSemanticMonitor) RawJSON() string { return r.JSON.raw }
func (r *MonitorListResponseDataMonitorsPageSemanticMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detect meaning-level changes that match a natural language query.
type MonitorListResponseDataMonitorsPageSemanticMonitorChangeDetection struct {
	Query string `json:"query" api:"required"`
	// Any of "semantic".
	Type                string  `json:"type" api:"required"`
	ConfidenceThreshold float64 `json:"confidence_threshold"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Query               respjson.Field
		Type                respjson.Field
		ConfidenceThreshold respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListResponseDataMonitorsPageSemanticMonitorChangeDetection) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorListResponseDataMonitorsPageSemanticMonitorChangeDetection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
// every 6 hours or every 2 days. The total interval (frequency × unit) must be
// between 10 minutes and 1 year.
type MonitorListResponseDataMonitorsPageSemanticMonitorSchedule struct {
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
func (r MonitorListResponseDataMonitorsPageSemanticMonitorSchedule) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorListResponseDataMonitorsPageSemanticMonitorSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListResponseDataMonitorsPageSemanticMonitorTarget struct {
	// Any of "page".
	Type string `json:"type" api:"required"`
	URL  string `json:"url" api:"required" format:"uri"`
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
func (r MonitorListResponseDataMonitorsPageSemanticMonitorTarget) RawJSON() string { return r.JSON.raw }
func (r *MonitorListResponseDataMonitorsPageSemanticMonitorTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListResponseDataMonitorsPageSemanticMonitorWebhook struct {
	// Webhook URL called when a change is detected.
	URL string `json:"url" api:"required" format:"uri"`
	// Signing secret used to verify webhook authenticity. Each delivery includes an
	// `X-Context-Signature: t=<unix>,v1=<hmac>` header, where the HMAC is SHA-256 over
	// `"{t}.{rawRequestBody}"` keyed by this secret. Recompute it with a constant-time
	// compare and reject stale timestamps to prevent replay. Generated by the API;
	// cannot be set by clients.
	Secret string `json:"secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListResponseDataMonitorsPageSemanticMonitorWebhook) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorListResponseDataMonitorsPageSemanticMonitorWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An extract monitor using semantic change detection.
type MonitorListResponseDataMonitorsExtractSemanticMonitor struct {
	ID string `json:"id" api:"required"`
	// Detect meaning-level changes that match a natural language query.
	ChangeDetection MonitorListResponseDataMonitorsExtractSemanticMonitorChangeDetection `json:"change_detection" api:"required"`
	CreatedAt       time.Time                                                            `json:"created_at" api:"required" format:"date-time"`
	Name            string                                                               `json:"name" api:"required"`
	// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
	// every 6 hours or every 2 days. The total interval (frequency × unit) must be
	// between 10 minutes and 1 year.
	Schedule MonitorListResponseDataMonitorsExtractSemanticMonitorSchedule `json:"schedule" api:"required"`
	// Any of "active", "paused", "failed".
	Status       string                                                      `json:"status" api:"required"`
	Target       MonitorListResponseDataMonitorsExtractSemanticMonitorTarget `json:"target" api:"required"`
	UpdatedAt    time.Time                                                   `json:"updated_at" api:"required" format:"date-time"`
	LastChangeAt time.Time                                                   `json:"last_change_at" api:"nullable" format:"date-time"`
	LastRunAt    time.Time                                                   `json:"last_run_at" api:"nullable" format:"date-time"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags    []string                                                     `json:"tags"`
	Webhook MonitorListResponseDataMonitorsExtractSemanticMonitorWebhook `json:"webhook" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ChangeDetection respjson.Field
		CreatedAt       respjson.Field
		Name            respjson.Field
		Schedule        respjson.Field
		Status          respjson.Field
		Target          respjson.Field
		UpdatedAt       respjson.Field
		LastChangeAt    respjson.Field
		LastRunAt       respjson.Field
		Tags            respjson.Field
		Webhook         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListResponseDataMonitorsExtractSemanticMonitor) RawJSON() string { return r.JSON.raw }
func (r *MonitorListResponseDataMonitorsExtractSemanticMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detect meaning-level changes that match a natural language query.
type MonitorListResponseDataMonitorsExtractSemanticMonitorChangeDetection struct {
	Query string `json:"query" api:"required"`
	// Any of "semantic".
	Type                string  `json:"type" api:"required"`
	ConfidenceThreshold float64 `json:"confidence_threshold"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Query               respjson.Field
		Type                respjson.Field
		ConfidenceThreshold respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListResponseDataMonitorsExtractSemanticMonitorChangeDetection) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorListResponseDataMonitorsExtractSemanticMonitorChangeDetection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
// every 6 hours or every 2 days. The total interval (frequency × unit) must be
// between 10 minutes and 1 year.
type MonitorListResponseDataMonitorsExtractSemanticMonitorSchedule struct {
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
func (r MonitorListResponseDataMonitorsExtractSemanticMonitorSchedule) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorListResponseDataMonitorsExtractSemanticMonitorSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListResponseDataMonitorsExtractSemanticMonitorTarget struct {
	// Any of "extract".
	Type string `json:"type" api:"required"`
	// Root URL to extract structured data from.
	URL              string `json:"url" api:"required" format:"uri"`
	FollowSubdomains bool   `json:"follow_subdomains"`
	// Optional natural-language instructions guiding what to extract.
	Instructions string `json:"instructions"`
	// Optional maximum link depth from the starting URL (0 = only the starting page).
	MaxDepth int64 `json:"max_depth"`
	// Maximum number of pages to analyze during extraction.
	MaxPages int64 `json:"max_pages"`
	// JSON Schema describing the structured data to extract and watch for changes. If
	// omitted, a default summary + key-points schema is used.
	Schema map[string]any `json:"schema"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type             respjson.Field
		URL              respjson.Field
		FollowSubdomains respjson.Field
		Instructions     respjson.Field
		MaxDepth         respjson.Field
		MaxPages         respjson.Field
		Schema           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListResponseDataMonitorsExtractSemanticMonitorTarget) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorListResponseDataMonitorsExtractSemanticMonitorTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListResponseDataMonitorsExtractSemanticMonitorWebhook struct {
	// Webhook URL called when a change is detected.
	URL string `json:"url" api:"required" format:"uri"`
	// Signing secret used to verify webhook authenticity. Each delivery includes an
	// `X-Context-Signature: t=<unix>,v1=<hmac>` header, where the HMAC is SHA-256 over
	// `"{t}.{rawRequestBody}"` keyed by this secret. Recompute it with a constant-time
	// compare and reject stale timestamps to prevent replay. Generated by the API;
	// cannot be set by clients.
	Secret string `json:"secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListResponseDataMonitorsExtractSemanticMonitorWebhook) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorListResponseDataMonitorsExtractSemanticMonitorWebhook) UnmarshalJSON(data []byte) error {
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

type MonitorListAccountChangesResponse struct {
	Data       []MonitorListAccountChangesResponseDataUnion `json:"data" api:"required"`
	HasMore    bool                                         `json:"has_more" api:"required"`
	NextCursor string                                       `json:"next_cursor" api:"required"`
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

// MonitorListAccountChangesResponseDataUnion contains all possible properties and
// values from
// [MonitorListAccountChangesResponseDataMonitorsPageExactChangeSummary],
// [MonitorListAccountChangesResponseDataMonitorsSitemapExactChangeSummary],
// [MonitorListAccountChangesResponseDataMonitorsPageSemanticChangeSummary],
// [MonitorListAccountChangesResponseDataMonitorsExtractSemanticChangeSummary].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MonitorListAccountChangesResponseDataUnion struct {
	ID                  string    `json:"id"`
	ChangeDetectionType string    `json:"change_detection_type"`
	DetectedAt          time.Time `json:"detected_at"`
	MonitorID           string    `json:"monitor_id"`
	Summary             string    `json:"summary"`
	TargetType          string    `json:"target_type"`
	Title               string    `json:"title"`
	URL                 string    `json:"url"`
	Tags                []string  `json:"tags"`
	// This field is from variant
	// [MonitorListAccountChangesResponseDataMonitorsSitemapExactChangeSummary].
	AddedURLCount int64 `json:"added_url_count"`
	// This field is from variant
	// [MonitorListAccountChangesResponseDataMonitorsSitemapExactChangeSummary].
	RemovedURLCount int64   `json:"removed_url_count"`
	Confidence      float64 `json:"confidence"`
	Importance      string  `json:"importance"`
	// This field is from variant
	// [MonitorListAccountChangesResponseDataMonitorsExtractSemanticChangeSummary].
	MatchedURLCount int64 `json:"matched_url_count"`
	JSON            struct {
		ID                  respjson.Field
		ChangeDetectionType respjson.Field
		DetectedAt          respjson.Field
		MonitorID           respjson.Field
		Summary             respjson.Field
		TargetType          respjson.Field
		Title               respjson.Field
		URL                 respjson.Field
		Tags                respjson.Field
		AddedURLCount       respjson.Field
		RemovedURLCount     respjson.Field
		Confidence          respjson.Field
		Importance          respjson.Field
		MatchedURLCount     respjson.Field
		raw                 string
	} `json:"-"`
}

func (u MonitorListAccountChangesResponseDataUnion) AsMonitorListAccountChangesResponseDataMonitorsPageExactChangeSummary() (v MonitorListAccountChangesResponseDataMonitorsPageExactChangeSummary) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorListAccountChangesResponseDataUnion) AsMonitorListAccountChangesResponseDataMonitorsSitemapExactChangeSummary() (v MonitorListAccountChangesResponseDataMonitorsSitemapExactChangeSummary) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorListAccountChangesResponseDataUnion) AsMonitorListAccountChangesResponseDataMonitorsPageSemanticChangeSummary() (v MonitorListAccountChangesResponseDataMonitorsPageSemanticChangeSummary) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorListAccountChangesResponseDataUnion) AsMonitorListAccountChangesResponseDataMonitorsExtractSemanticChangeSummary() (v MonitorListAccountChangesResponseDataMonitorsExtractSemanticChangeSummary) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MonitorListAccountChangesResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *MonitorListAccountChangesResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListAccountChangesResponseDataMonitorsPageExactChangeSummary struct {
	ID string `json:"id" api:"required"`
	// Any of "exact".
	ChangeDetectionType string    `json:"change_detection_type" api:"required"`
	DetectedAt          time.Time `json:"detected_at" api:"required" format:"date-time"`
	MonitorID           string    `json:"monitor_id" api:"required"`
	Summary             string    `json:"summary" api:"required"`
	// Any of "page".
	TargetType string `json:"target_type" api:"required"`
	Title      string `json:"title" api:"required"`
	URL        string `json:"url" api:"required" format:"uri"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags []string `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		ChangeDetectionType respjson.Field
		DetectedAt          respjson.Field
		MonitorID           respjson.Field
		Summary             respjson.Field
		TargetType          respjson.Field
		Title               respjson.Field
		URL                 respjson.Field
		Tags                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListAccountChangesResponseDataMonitorsPageExactChangeSummary) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorListAccountChangesResponseDataMonitorsPageExactChangeSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListAccountChangesResponseDataMonitorsSitemapExactChangeSummary struct {
	ID            string `json:"id" api:"required"`
	AddedURLCount int64  `json:"added_url_count" api:"required"`
	// Any of "exact".
	ChangeDetectionType string    `json:"change_detection_type" api:"required"`
	DetectedAt          time.Time `json:"detected_at" api:"required" format:"date-time"`
	MonitorID           string    `json:"monitor_id" api:"required"`
	RemovedURLCount     int64     `json:"removed_url_count" api:"required"`
	Summary             string    `json:"summary" api:"required"`
	// Any of "sitemap".
	TargetType string `json:"target_type" api:"required"`
	Title      string `json:"title" api:"required"`
	URL        string `json:"url" api:"required" format:"uri"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags []string `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		AddedURLCount       respjson.Field
		ChangeDetectionType respjson.Field
		DetectedAt          respjson.Field
		MonitorID           respjson.Field
		RemovedURLCount     respjson.Field
		Summary             respjson.Field
		TargetType          respjson.Field
		Title               respjson.Field
		URL                 respjson.Field
		Tags                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListAccountChangesResponseDataMonitorsSitemapExactChangeSummary) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorListAccountChangesResponseDataMonitorsSitemapExactChangeSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListAccountChangesResponseDataMonitorsPageSemanticChangeSummary struct {
	ID string `json:"id" api:"required"`
	// Any of "semantic".
	ChangeDetectionType string    `json:"change_detection_type" api:"required"`
	Confidence          float64   `json:"confidence" api:"required"`
	DetectedAt          time.Time `json:"detected_at" api:"required" format:"date-time"`
	// Any of "low", "medium", "high".
	Importance string `json:"importance" api:"required"`
	MonitorID  string `json:"monitor_id" api:"required"`
	Summary    string `json:"summary" api:"required"`
	// Any of "page".
	TargetType string `json:"target_type" api:"required"`
	Title      string `json:"title" api:"required"`
	URL        string `json:"url" api:"required" format:"uri"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags []string `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		ChangeDetectionType respjson.Field
		Confidence          respjson.Field
		DetectedAt          respjson.Field
		Importance          respjson.Field
		MonitorID           respjson.Field
		Summary             respjson.Field
		TargetType          respjson.Field
		Title               respjson.Field
		URL                 respjson.Field
		Tags                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListAccountChangesResponseDataMonitorsPageSemanticChangeSummary) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorListAccountChangesResponseDataMonitorsPageSemanticChangeSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListAccountChangesResponseDataMonitorsExtractSemanticChangeSummary struct {
	ID string `json:"id" api:"required"`
	// Any of "semantic".
	ChangeDetectionType string    `json:"change_detection_type" api:"required"`
	Confidence          float64   `json:"confidence" api:"required"`
	DetectedAt          time.Time `json:"detected_at" api:"required" format:"date-time"`
	// Any of "low", "medium", "high".
	Importance      string `json:"importance" api:"required"`
	MatchedURLCount int64  `json:"matched_url_count" api:"required"`
	MonitorID       string `json:"monitor_id" api:"required"`
	Summary         string `json:"summary" api:"required"`
	// Any of "extract".
	TargetType string `json:"target_type" api:"required"`
	Title      string `json:"title" api:"required"`
	URL        string `json:"url" api:"required" format:"uri"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags []string `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		ChangeDetectionType respjson.Field
		Confidence          respjson.Field
		DetectedAt          respjson.Field
		Importance          respjson.Field
		MatchedURLCount     respjson.Field
		MonitorID           respjson.Field
		Summary             respjson.Field
		TargetType          respjson.Field
		Title               respjson.Field
		URL                 respjson.Field
		Tags                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListAccountChangesResponseDataMonitorsExtractSemanticChangeSummary) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorListAccountChangesResponseDataMonitorsExtractSemanticChangeSummary) UnmarshalJSON(data []byte) error {
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
	MonitorID           string `json:"monitor_id" api:"required"`
	// The first run after monitor creation is a baseline run.
	//
	// Any of "baseline", "scheduled".
	RunType string `json:"run_type" api:"required"`
	// Any of "queued", "running", "completed", "failed".
	Status string `json:"status" api:"required"`
	// Any of "page", "sitemap", "extract".
	TargetType  string                                  `json:"target_type" api:"required"`
	ChangeID    string                                  `json:"change_id" api:"nullable"`
	CompletedAt time.Time                               `json:"completed_at" api:"nullable" format:"date-time"`
	Error       MonitorListAccountRunsResponseDataError `json:"error" api:"nullable"`
	StartedAt   time.Time                               `json:"started_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		BaselineCreated     respjson.Field
		ChangeDetected      respjson.Field
		ChangeDetectionType respjson.Field
		MonitorID           respjson.Field
		RunType             respjson.Field
		Status              respjson.Field
		TargetType          respjson.Field
		ChangeID            respjson.Field
		CompletedAt         respjson.Field
		Error               respjson.Field
		StartedAt           respjson.Field
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

type MonitorListChangesResponse struct {
	Data       []MonitorListChangesResponseDataUnion `json:"data" api:"required"`
	HasMore    bool                                  `json:"has_more" api:"required"`
	NextCursor string                                `json:"next_cursor" api:"required"`
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

// MonitorListChangesResponseDataUnion contains all possible properties and values
// from [MonitorListChangesResponseDataMonitorsPageExactChangeSummary],
// [MonitorListChangesResponseDataMonitorsSitemapExactChangeSummary],
// [MonitorListChangesResponseDataMonitorsPageSemanticChangeSummary],
// [MonitorListChangesResponseDataMonitorsExtractSemanticChangeSummary].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MonitorListChangesResponseDataUnion struct {
	ID                  string    `json:"id"`
	ChangeDetectionType string    `json:"change_detection_type"`
	DetectedAt          time.Time `json:"detected_at"`
	MonitorID           string    `json:"monitor_id"`
	Summary             string    `json:"summary"`
	TargetType          string    `json:"target_type"`
	Title               string    `json:"title"`
	URL                 string    `json:"url"`
	Tags                []string  `json:"tags"`
	// This field is from variant
	// [MonitorListChangesResponseDataMonitorsSitemapExactChangeSummary].
	AddedURLCount int64 `json:"added_url_count"`
	// This field is from variant
	// [MonitorListChangesResponseDataMonitorsSitemapExactChangeSummary].
	RemovedURLCount int64   `json:"removed_url_count"`
	Confidence      float64 `json:"confidence"`
	Importance      string  `json:"importance"`
	// This field is from variant
	// [MonitorListChangesResponseDataMonitorsExtractSemanticChangeSummary].
	MatchedURLCount int64 `json:"matched_url_count"`
	JSON            struct {
		ID                  respjson.Field
		ChangeDetectionType respjson.Field
		DetectedAt          respjson.Field
		MonitorID           respjson.Field
		Summary             respjson.Field
		TargetType          respjson.Field
		Title               respjson.Field
		URL                 respjson.Field
		Tags                respjson.Field
		AddedURLCount       respjson.Field
		RemovedURLCount     respjson.Field
		Confidence          respjson.Field
		Importance          respjson.Field
		MatchedURLCount     respjson.Field
		raw                 string
	} `json:"-"`
}

func (u MonitorListChangesResponseDataUnion) AsMonitorListChangesResponseDataMonitorsPageExactChangeSummary() (v MonitorListChangesResponseDataMonitorsPageExactChangeSummary) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorListChangesResponseDataUnion) AsMonitorListChangesResponseDataMonitorsSitemapExactChangeSummary() (v MonitorListChangesResponseDataMonitorsSitemapExactChangeSummary) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorListChangesResponseDataUnion) AsMonitorListChangesResponseDataMonitorsPageSemanticChangeSummary() (v MonitorListChangesResponseDataMonitorsPageSemanticChangeSummary) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorListChangesResponseDataUnion) AsMonitorListChangesResponseDataMonitorsExtractSemanticChangeSummary() (v MonitorListChangesResponseDataMonitorsExtractSemanticChangeSummary) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MonitorListChangesResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *MonitorListChangesResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListChangesResponseDataMonitorsPageExactChangeSummary struct {
	ID string `json:"id" api:"required"`
	// Any of "exact".
	ChangeDetectionType string    `json:"change_detection_type" api:"required"`
	DetectedAt          time.Time `json:"detected_at" api:"required" format:"date-time"`
	MonitorID           string    `json:"monitor_id" api:"required"`
	Summary             string    `json:"summary" api:"required"`
	// Any of "page".
	TargetType string `json:"target_type" api:"required"`
	Title      string `json:"title" api:"required"`
	URL        string `json:"url" api:"required" format:"uri"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags []string `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		ChangeDetectionType respjson.Field
		DetectedAt          respjson.Field
		MonitorID           respjson.Field
		Summary             respjson.Field
		TargetType          respjson.Field
		Title               respjson.Field
		URL                 respjson.Field
		Tags                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListChangesResponseDataMonitorsPageExactChangeSummary) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorListChangesResponseDataMonitorsPageExactChangeSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListChangesResponseDataMonitorsSitemapExactChangeSummary struct {
	ID            string `json:"id" api:"required"`
	AddedURLCount int64  `json:"added_url_count" api:"required"`
	// Any of "exact".
	ChangeDetectionType string    `json:"change_detection_type" api:"required"`
	DetectedAt          time.Time `json:"detected_at" api:"required" format:"date-time"`
	MonitorID           string    `json:"monitor_id" api:"required"`
	RemovedURLCount     int64     `json:"removed_url_count" api:"required"`
	Summary             string    `json:"summary" api:"required"`
	// Any of "sitemap".
	TargetType string `json:"target_type" api:"required"`
	Title      string `json:"title" api:"required"`
	URL        string `json:"url" api:"required" format:"uri"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags []string `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		AddedURLCount       respjson.Field
		ChangeDetectionType respjson.Field
		DetectedAt          respjson.Field
		MonitorID           respjson.Field
		RemovedURLCount     respjson.Field
		Summary             respjson.Field
		TargetType          respjson.Field
		Title               respjson.Field
		URL                 respjson.Field
		Tags                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListChangesResponseDataMonitorsSitemapExactChangeSummary) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorListChangesResponseDataMonitorsSitemapExactChangeSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListChangesResponseDataMonitorsPageSemanticChangeSummary struct {
	ID string `json:"id" api:"required"`
	// Any of "semantic".
	ChangeDetectionType string    `json:"change_detection_type" api:"required"`
	Confidence          float64   `json:"confidence" api:"required"`
	DetectedAt          time.Time `json:"detected_at" api:"required" format:"date-time"`
	// Any of "low", "medium", "high".
	Importance string `json:"importance" api:"required"`
	MonitorID  string `json:"monitor_id" api:"required"`
	Summary    string `json:"summary" api:"required"`
	// Any of "page".
	TargetType string `json:"target_type" api:"required"`
	Title      string `json:"title" api:"required"`
	URL        string `json:"url" api:"required" format:"uri"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags []string `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		ChangeDetectionType respjson.Field
		Confidence          respjson.Field
		DetectedAt          respjson.Field
		Importance          respjson.Field
		MonitorID           respjson.Field
		Summary             respjson.Field
		TargetType          respjson.Field
		Title               respjson.Field
		URL                 respjson.Field
		Tags                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListChangesResponseDataMonitorsPageSemanticChangeSummary) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorListChangesResponseDataMonitorsPageSemanticChangeSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListChangesResponseDataMonitorsExtractSemanticChangeSummary struct {
	ID string `json:"id" api:"required"`
	// Any of "semantic".
	ChangeDetectionType string    `json:"change_detection_type" api:"required"`
	Confidence          float64   `json:"confidence" api:"required"`
	DetectedAt          time.Time `json:"detected_at" api:"required" format:"date-time"`
	// Any of "low", "medium", "high".
	Importance      string `json:"importance" api:"required"`
	MatchedURLCount int64  `json:"matched_url_count" api:"required"`
	MonitorID       string `json:"monitor_id" api:"required"`
	Summary         string `json:"summary" api:"required"`
	// Any of "extract".
	TargetType string `json:"target_type" api:"required"`
	Title      string `json:"title" api:"required"`
	URL        string `json:"url" api:"required" format:"uri"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags []string `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		ChangeDetectionType respjson.Field
		Confidence          respjson.Field
		DetectedAt          respjson.Field
		Importance          respjson.Field
		MatchedURLCount     respjson.Field
		MonitorID           respjson.Field
		Summary             respjson.Field
		TargetType          respjson.Field
		Title               respjson.Field
		URL                 respjson.Field
		Tags                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListChangesResponseDataMonitorsExtractSemanticChangeSummary) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorListChangesResponseDataMonitorsExtractSemanticChangeSummary) UnmarshalJSON(data []byte) error {
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
	MonitorID           string `json:"monitor_id" api:"required"`
	// The first run after monitor creation is a baseline run.
	//
	// Any of "baseline", "scheduled".
	RunType string `json:"run_type" api:"required"`
	// Any of "queued", "running", "completed", "failed".
	Status string `json:"status" api:"required"`
	// Any of "page", "sitemap", "extract".
	TargetType  string                           `json:"target_type" api:"required"`
	ChangeID    string                           `json:"change_id" api:"nullable"`
	CompletedAt time.Time                        `json:"completed_at" api:"nullable" format:"date-time"`
	Error       MonitorListRunsResponseDataError `json:"error" api:"nullable"`
	StartedAt   time.Time                        `json:"started_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		BaselineCreated     respjson.Field
		ChangeDetected      respjson.Field
		ChangeDetectionType respjson.Field
		MonitorID           respjson.Field
		RunType             respjson.Field
		Status              respjson.Field
		TargetType          respjson.Field
		ChangeID            respjson.Field
		CompletedAt         respjson.Field
		Error               respjson.Field
		StartedAt           respjson.Field
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

// MonitorGetChangeResponseUnion contains all possible properties and values from
// [MonitorGetChangeResponseMonitorsPageExactChange],
// [MonitorGetChangeResponseMonitorsSitemapExactChange],
// [MonitorGetChangeResponseMonitorsPageSemanticChange],
// [MonitorGetChangeResponseMonitorsExtractSemanticChange].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MonitorGetChangeResponseUnion struct {
	ID                  string    `json:"id"`
	ChangeDetectionType string    `json:"change_detection_type"`
	DetectedAt          time.Time `json:"detected_at"`
	// This field is from variant [MonitorGetChangeResponseMonitorsPageExactChange].
	Diff       string `json:"diff"`
	MonitorID  string `json:"monitor_id"`
	Summary    string `json:"summary"`
	TargetType string `json:"target_type"`
	Title      string `json:"title"`
	URL        string `json:"url"`
	// This field is from variant [MonitorGetChangeResponseMonitorsPageExactChange].
	AfterTextExcerpt string `json:"after_text_excerpt"`
	// This field is from variant [MonitorGetChangeResponseMonitorsPageExactChange].
	BeforeTextExcerpt string   `json:"before_text_excerpt"`
	Tags              []string `json:"tags"`
	// This field is from variant [MonitorGetChangeResponseMonitorsSitemapExactChange].
	AddedURLCount int64 `json:"added_url_count"`
	// This field is from variant [MonitorGetChangeResponseMonitorsSitemapExactChange].
	AddedURLs []string `json:"added_urls"`
	// This field is from variant [MonitorGetChangeResponseMonitorsSitemapExactChange].
	RemovedURLCount int64 `json:"removed_url_count"`
	// This field is from variant [MonitorGetChangeResponseMonitorsSitemapExactChange].
	RemovedURLs []string `json:"removed_urls"`
	Confidence  float64  `json:"confidence"`
	// This field is a union of
	// [[]MonitorGetChangeResponseMonitorsPageSemanticChangeEvidence],
	// [[]MonitorGetChangeResponseMonitorsExtractSemanticChangeEvidence]
	Evidence   MonitorGetChangeResponseUnionEvidence `json:"evidence"`
	Importance string                                `json:"importance"`
	Query      string                                `json:"query"`
	// This field is from variant
	// [MonitorGetChangeResponseMonitorsExtractSemanticChange].
	MatchedURLCount int64 `json:"matched_url_count"`
	// This field is from variant
	// [MonitorGetChangeResponseMonitorsExtractSemanticChange].
	MatchedURLs []string `json:"matched_urls"`
	JSON        struct {
		ID                  respjson.Field
		ChangeDetectionType respjson.Field
		DetectedAt          respjson.Field
		Diff                respjson.Field
		MonitorID           respjson.Field
		Summary             respjson.Field
		TargetType          respjson.Field
		Title               respjson.Field
		URL                 respjson.Field
		AfterTextExcerpt    respjson.Field
		BeforeTextExcerpt   respjson.Field
		Tags                respjson.Field
		AddedURLCount       respjson.Field
		AddedURLs           respjson.Field
		RemovedURLCount     respjson.Field
		RemovedURLs         respjson.Field
		Confidence          respjson.Field
		Evidence            respjson.Field
		Importance          respjson.Field
		Query               respjson.Field
		MatchedURLCount     respjson.Field
		MatchedURLs         respjson.Field
		raw                 string
	} `json:"-"`
}

func (u MonitorGetChangeResponseUnion) AsMonitorGetChangeResponseMonitorsPageExactChange() (v MonitorGetChangeResponseMonitorsPageExactChange) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorGetChangeResponseUnion) AsMonitorGetChangeResponseMonitorsSitemapExactChange() (v MonitorGetChangeResponseMonitorsSitemapExactChange) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorGetChangeResponseUnion) AsMonitorGetChangeResponseMonitorsPageSemanticChange() (v MonitorGetChangeResponseMonitorsPageSemanticChange) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MonitorGetChangeResponseUnion) AsMonitorGetChangeResponseMonitorsExtractSemanticChange() (v MonitorGetChangeResponseMonitorsExtractSemanticChange) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MonitorGetChangeResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *MonitorGetChangeResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MonitorGetChangeResponseUnionEvidence is an implicit subunion of
// [MonitorGetChangeResponseUnion]. MonitorGetChangeResponseUnionEvidence provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MonitorGetChangeResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfMonitorGetChangeResponseMonitorsPageSemanticChangeEvidenceArray
// OfMonitorGetChangeResponseMonitorsExtractSemanticChangeEvidenceArray]
type MonitorGetChangeResponseUnionEvidence struct {
	// This field will be present if the value is a
	// [[]MonitorGetChangeResponseMonitorsPageSemanticChangeEvidence] instead of an
	// object.
	OfMonitorGetChangeResponseMonitorsPageSemanticChangeEvidenceArray []MonitorGetChangeResponseMonitorsPageSemanticChangeEvidence `json:",inline"`
	// This field will be present if the value is a
	// [[]MonitorGetChangeResponseMonitorsExtractSemanticChangeEvidence] instead of an
	// object.
	OfMonitorGetChangeResponseMonitorsExtractSemanticChangeEvidenceArray []MonitorGetChangeResponseMonitorsExtractSemanticChangeEvidence `json:",inline"`
	JSON                                                                 struct {
		OfMonitorGetChangeResponseMonitorsPageSemanticChangeEvidenceArray    respjson.Field
		OfMonitorGetChangeResponseMonitorsExtractSemanticChangeEvidenceArray respjson.Field
		raw                                                                  string
	} `json:"-"`
}

func (r *MonitorGetChangeResponseUnionEvidence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorGetChangeResponseMonitorsPageExactChange struct {
	ID string `json:"id" api:"required"`
	// Any of "exact".
	ChangeDetectionType string    `json:"change_detection_type" api:"required"`
	DetectedAt          time.Time `json:"detected_at" api:"required" format:"date-time"`
	// Text diff between the previous and current page baseline.
	Diff      string `json:"diff" api:"required"`
	MonitorID string `json:"monitor_id" api:"required"`
	Summary   string `json:"summary" api:"required"`
	// Any of "page".
	TargetType        string `json:"target_type" api:"required"`
	Title             string `json:"title" api:"required"`
	URL               string `json:"url" api:"required" format:"uri"`
	AfterTextExcerpt  string `json:"after_text_excerpt"`
	BeforeTextExcerpt string `json:"before_text_excerpt"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags []string `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		ChangeDetectionType respjson.Field
		DetectedAt          respjson.Field
		Diff                respjson.Field
		MonitorID           respjson.Field
		Summary             respjson.Field
		TargetType          respjson.Field
		Title               respjson.Field
		URL                 respjson.Field
		AfterTextExcerpt    respjson.Field
		BeforeTextExcerpt   respjson.Field
		Tags                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetChangeResponseMonitorsPageExactChange) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetChangeResponseMonitorsPageExactChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorGetChangeResponseMonitorsSitemapExactChange struct {
	ID            string   `json:"id" api:"required"`
	AddedURLCount int64    `json:"added_url_count" api:"required"`
	AddedURLs     []string `json:"added_urls" api:"required" format:"uri"`
	// Any of "exact".
	ChangeDetectionType string    `json:"change_detection_type" api:"required"`
	DetectedAt          time.Time `json:"detected_at" api:"required" format:"date-time"`
	MonitorID           string    `json:"monitor_id" api:"required"`
	RemovedURLCount     int64     `json:"removed_url_count" api:"required"`
	RemovedURLs         []string  `json:"removed_urls" api:"required" format:"uri"`
	Summary             string    `json:"summary" api:"required"`
	// Any of "sitemap".
	TargetType string `json:"target_type" api:"required"`
	Title      string `json:"title" api:"required"`
	URL        string `json:"url" api:"required" format:"uri"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags []string `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		AddedURLCount       respjson.Field
		AddedURLs           respjson.Field
		ChangeDetectionType respjson.Field
		DetectedAt          respjson.Field
		MonitorID           respjson.Field
		RemovedURLCount     respjson.Field
		RemovedURLs         respjson.Field
		Summary             respjson.Field
		TargetType          respjson.Field
		Title               respjson.Field
		URL                 respjson.Field
		Tags                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetChangeResponseMonitorsSitemapExactChange) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetChangeResponseMonitorsSitemapExactChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorGetChangeResponseMonitorsPageSemanticChange struct {
	ID string `json:"id" api:"required"`
	// Any of "semantic".
	ChangeDetectionType string                                                       `json:"change_detection_type" api:"required"`
	Confidence          float64                                                      `json:"confidence" api:"required"`
	DetectedAt          time.Time                                                    `json:"detected_at" api:"required" format:"date-time"`
	Evidence            []MonitorGetChangeResponseMonitorsPageSemanticChangeEvidence `json:"evidence" api:"required"`
	// Any of "low", "medium", "high".
	Importance string `json:"importance" api:"required"`
	MonitorID  string `json:"monitor_id" api:"required"`
	Query      string `json:"query" api:"required"`
	Summary    string `json:"summary" api:"required"`
	// Any of "page".
	TargetType string `json:"target_type" api:"required"`
	Title      string `json:"title" api:"required"`
	URL        string `json:"url" api:"required" format:"uri"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags []string `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		ChangeDetectionType respjson.Field
		Confidence          respjson.Field
		DetectedAt          respjson.Field
		Evidence            respjson.Field
		Importance          respjson.Field
		MonitorID           respjson.Field
		Query               respjson.Field
		Summary             respjson.Field
		TargetType          respjson.Field
		Title               respjson.Field
		URL                 respjson.Field
		Tags                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetChangeResponseMonitorsPageSemanticChange) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetChangeResponseMonitorsPageSemanticChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorGetChangeResponseMonitorsPageSemanticChangeEvidence struct {
	After  string `json:"after" api:"required"`
	Before string `json:"before" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		After       respjson.Field
		Before      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetChangeResponseMonitorsPageSemanticChangeEvidence) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorGetChangeResponseMonitorsPageSemanticChangeEvidence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorGetChangeResponseMonitorsExtractSemanticChange struct {
	ID string `json:"id" api:"required"`
	// Any of "semantic".
	ChangeDetectionType string                                                          `json:"change_detection_type" api:"required"`
	Confidence          float64                                                         `json:"confidence" api:"required"`
	DetectedAt          time.Time                                                       `json:"detected_at" api:"required" format:"date-time"`
	Evidence            []MonitorGetChangeResponseMonitorsExtractSemanticChangeEvidence `json:"evidence" api:"required"`
	// Any of "low", "medium", "high".
	Importance      string   `json:"importance" api:"required"`
	MatchedURLCount int64    `json:"matched_url_count" api:"required"`
	MatchedURLs     []string `json:"matched_urls" api:"required" format:"uri"`
	MonitorID       string   `json:"monitor_id" api:"required"`
	Query           string   `json:"query" api:"required"`
	Summary         string   `json:"summary" api:"required"`
	// Any of "extract".
	TargetType string `json:"target_type" api:"required"`
	Title      string `json:"title" api:"required"`
	// Root URL of the extract target.
	URL string `json:"url" api:"required" format:"uri"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags []string `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		ChangeDetectionType respjson.Field
		Confidence          respjson.Field
		DetectedAt          respjson.Field
		Evidence            respjson.Field
		Importance          respjson.Field
		MatchedURLCount     respjson.Field
		MatchedURLs         respjson.Field
		MonitorID           respjson.Field
		Query               respjson.Field
		Summary             respjson.Field
		TargetType          respjson.Field
		Title               respjson.Field
		URL                 respjson.Field
		Tags                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetChangeResponseMonitorsExtractSemanticChange) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetChangeResponseMonitorsExtractSemanticChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorGetChangeResponseMonitorsExtractSemanticChangeEvidence struct {
	// Snapshot of the extracted data after the change.
	After string `json:"after" api:"required"`
	// Snapshot of the extracted data before the change.
	Before string `json:"before" api:"required"`
	// Optional URL the evidence relates to. Absent for whole-target extract diffs.
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
func (r MonitorGetChangeResponseMonitorsExtractSemanticChangeEvidence) RawJSON() string {
	return r.JSON.raw
}
func (r *MonitorGetChangeResponseMonitorsExtractSemanticChangeEvidence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorRunResponse struct {
	MonitorID string `json:"monitor_id" api:"required"`
	Queued    bool   `json:"queued" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MonitorID   respjson.Field
		Queued      respjson.Field
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

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set. Monitor
	// a single page for exact visible text changes.
	OfMonitorsCreatePageExactMonitorRequest *MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequest `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Monitor
	// a sitemap for exact URL additions and removals.
	OfMonitorsCreateSitemapExactMonitorRequest *MonitorNewParamsBodyMonitorsCreateSitemapExactMonitorRequest `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Monitor
	// a single page for semantic changes described by a natural language query.
	OfMonitorsCreatePageSemanticMonitorRequest *MonitorNewParamsBodyMonitorsCreatePageSemanticMonitorRequest `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Monitor
	// a website's extracted structured data for semantic changes described by a
	// natural language query.
	OfMonitorsCreateExtractSemanticMonitorRequest *MonitorNewParamsBodyMonitorsCreateExtractSemanticMonitorRequest `json:",inline"`

	paramObj
}

func (u MonitorNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfMonitorsCreatePageExactMonitorRequest, u.OfMonitorsCreateSitemapExactMonitorRequest, u.OfMonitorsCreatePageSemanticMonitorRequest, u.OfMonitorsCreateExtractSemanticMonitorRequest)
}
func (r *MonitorNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monitor a single page for exact visible text changes.
//
// The properties ChangeDetection, Name, Schedule, Target are required.
type MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequest struct {
	// Detect exact changes. For page targets, this means visible text diffs. For
	// sitemap targets, this means URL additions and removals.
	ChangeDetection MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequestChangeDetection `json:"change_detection,omitzero" api:"required"`
	Name            string                                                                   `json:"name" api:"required"`
	// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
	// every 6 hours or every 2 days. The total interval (frequency × unit) must be
	// between 10 minutes and 1 year.
	Schedule MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequestSchedule `json:"schedule,omitzero" api:"required"`
	Target   MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequestTarget   `json:"target,omitzero" api:"required"`
	Webhook  MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequestWebhook  `json:"webhook,omitzero"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequest) MarshalJSON() (data []byte, err error) {
	type shadow MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detect exact changes. For page targets, this means visible text diffs. For
// sitemap targets, this means URL additions and removals.
//
// The property Type is required.
type MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequestChangeDetection struct {
	// Any of "exact".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequestChangeDetection) MarshalJSON() (data []byte, err error) {
	type shadow MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequestChangeDetection
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequestChangeDetection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequestChangeDetection](
		"type", "exact",
	)
}

// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
// every 6 hours or every 2 days. The total interval (frequency × unit) must be
// between 10 minutes and 1 year.
//
// The properties Frequency, Type, Unit are required.
type MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequestSchedule struct {
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

func (r MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequestSchedule) MarshalJSON() (data []byte, err error) {
	type shadow MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequestSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequestSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequestSchedule](
		"type", "interval",
	)
	apijson.RegisterFieldValidator[MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequestSchedule](
		"unit", "minutes", "hours", "days",
	)
}

// The properties Type, URL are required.
type MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequestTarget struct {
	// Any of "page".
	Type string `json:"type,omitzero" api:"required"`
	URL  string `json:"url" api:"required" format:"uri"`
	// Normalize whitespace before comparing or analyzing text.
	NormalizeWhitespace param.Opt[bool] `json:"normalize_whitespace,omitzero"`
	paramObj
}

func (r MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequestTarget) MarshalJSON() (data []byte, err error) {
	type shadow MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequestTarget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequestTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequestTarget](
		"type", "page",
	)
}

// The property URL is required.
type MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequestWebhook struct {
	// Webhook URL called when a change is detected.
	URL string `json:"url" api:"required" format:"uri"`
	paramObj
}

func (r MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequestWebhook) MarshalJSON() (data []byte, err error) {
	type shadow MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequestWebhook
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequestWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monitor a sitemap for exact URL additions and removals.
//
// The properties ChangeDetection, Name, Schedule, Target are required.
type MonitorNewParamsBodyMonitorsCreateSitemapExactMonitorRequest struct {
	// Detect exact changes. For page targets, this means visible text diffs. For
	// sitemap targets, this means URL additions and removals.
	ChangeDetection MonitorNewParamsBodyMonitorsCreateSitemapExactMonitorRequestChangeDetection `json:"change_detection,omitzero" api:"required"`
	Name            string                                                                      `json:"name" api:"required"`
	// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
	// every 6 hours or every 2 days. The total interval (frequency × unit) must be
	// between 10 minutes and 1 year.
	Schedule MonitorNewParamsBodyMonitorsCreateSitemapExactMonitorRequestSchedule `json:"schedule,omitzero" api:"required"`
	Target   MonitorNewParamsBodyMonitorsCreateSitemapExactMonitorRequestTarget   `json:"target,omitzero" api:"required"`
	Webhook  MonitorNewParamsBodyMonitorsCreateSitemapExactMonitorRequestWebhook  `json:"webhook,omitzero"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r MonitorNewParamsBodyMonitorsCreateSitemapExactMonitorRequest) MarshalJSON() (data []byte, err error) {
	type shadow MonitorNewParamsBodyMonitorsCreateSitemapExactMonitorRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorNewParamsBodyMonitorsCreateSitemapExactMonitorRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detect exact changes. For page targets, this means visible text diffs. For
// sitemap targets, this means URL additions and removals.
//
// The property Type is required.
type MonitorNewParamsBodyMonitorsCreateSitemapExactMonitorRequestChangeDetection struct {
	// Any of "exact".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r MonitorNewParamsBodyMonitorsCreateSitemapExactMonitorRequestChangeDetection) MarshalJSON() (data []byte, err error) {
	type shadow MonitorNewParamsBodyMonitorsCreateSitemapExactMonitorRequestChangeDetection
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorNewParamsBodyMonitorsCreateSitemapExactMonitorRequestChangeDetection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MonitorNewParamsBodyMonitorsCreateSitemapExactMonitorRequestChangeDetection](
		"type", "exact",
	)
}

// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
// every 6 hours or every 2 days. The total interval (frequency × unit) must be
// between 10 minutes and 1 year.
//
// The properties Frequency, Type, Unit are required.
type MonitorNewParamsBodyMonitorsCreateSitemapExactMonitorRequestSchedule struct {
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

func (r MonitorNewParamsBodyMonitorsCreateSitemapExactMonitorRequestSchedule) MarshalJSON() (data []byte, err error) {
	type shadow MonitorNewParamsBodyMonitorsCreateSitemapExactMonitorRequestSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorNewParamsBodyMonitorsCreateSitemapExactMonitorRequestSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MonitorNewParamsBodyMonitorsCreateSitemapExactMonitorRequestSchedule](
		"type", "interval",
	)
	apijson.RegisterFieldValidator[MonitorNewParamsBodyMonitorsCreateSitemapExactMonitorRequestSchedule](
		"unit", "minutes", "hours", "days",
	)
}

// The properties Type, URL are required.
type MonitorNewParamsBodyMonitorsCreateSitemapExactMonitorRequestTarget struct {
	// Any of "sitemap".
	Type string `json:"type,omitzero" api:"required"`
	// Sitemap URL to monitor.
	URL     string           `json:"url" api:"required" format:"uri"`
	MaxURLs param.Opt[int64] `json:"max_urls,omitzero"`
	// URL path patterns to exclude.
	Exclude []string `json:"exclude,omitzero"`
	// URL path patterns to include.
	Include []string `json:"include,omitzero"`
	paramObj
}

func (r MonitorNewParamsBodyMonitorsCreateSitemapExactMonitorRequestTarget) MarshalJSON() (data []byte, err error) {
	type shadow MonitorNewParamsBodyMonitorsCreateSitemapExactMonitorRequestTarget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorNewParamsBodyMonitorsCreateSitemapExactMonitorRequestTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MonitorNewParamsBodyMonitorsCreateSitemapExactMonitorRequestTarget](
		"type", "sitemap",
	)
}

// The property URL is required.
type MonitorNewParamsBodyMonitorsCreateSitemapExactMonitorRequestWebhook struct {
	// Webhook URL called when a change is detected.
	URL string `json:"url" api:"required" format:"uri"`
	paramObj
}

func (r MonitorNewParamsBodyMonitorsCreateSitemapExactMonitorRequestWebhook) MarshalJSON() (data []byte, err error) {
	type shadow MonitorNewParamsBodyMonitorsCreateSitemapExactMonitorRequestWebhook
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorNewParamsBodyMonitorsCreateSitemapExactMonitorRequestWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monitor a single page for semantic changes described by a natural language
// query.
//
// The properties ChangeDetection, Name, Schedule, Target are required.
type MonitorNewParamsBodyMonitorsCreatePageSemanticMonitorRequest struct {
	// Detect meaning-level changes that match a natural language query.
	ChangeDetection MonitorNewParamsBodyMonitorsCreatePageSemanticMonitorRequestChangeDetection `json:"change_detection,omitzero" api:"required"`
	Name            string                                                                      `json:"name" api:"required"`
	// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
	// every 6 hours or every 2 days. The total interval (frequency × unit) must be
	// between 10 minutes and 1 year.
	Schedule MonitorNewParamsBodyMonitorsCreatePageSemanticMonitorRequestSchedule `json:"schedule,omitzero" api:"required"`
	Target   MonitorNewParamsBodyMonitorsCreatePageSemanticMonitorRequestTarget   `json:"target,omitzero" api:"required"`
	Webhook  MonitorNewParamsBodyMonitorsCreatePageSemanticMonitorRequestWebhook  `json:"webhook,omitzero"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r MonitorNewParamsBodyMonitorsCreatePageSemanticMonitorRequest) MarshalJSON() (data []byte, err error) {
	type shadow MonitorNewParamsBodyMonitorsCreatePageSemanticMonitorRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorNewParamsBodyMonitorsCreatePageSemanticMonitorRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detect meaning-level changes that match a natural language query.
//
// The properties Query, Type are required.
type MonitorNewParamsBodyMonitorsCreatePageSemanticMonitorRequestChangeDetection struct {
	Query string `json:"query" api:"required"`
	// Any of "semantic".
	Type                string             `json:"type,omitzero" api:"required"`
	ConfidenceThreshold param.Opt[float64] `json:"confidence_threshold,omitzero"`
	paramObj
}

func (r MonitorNewParamsBodyMonitorsCreatePageSemanticMonitorRequestChangeDetection) MarshalJSON() (data []byte, err error) {
	type shadow MonitorNewParamsBodyMonitorsCreatePageSemanticMonitorRequestChangeDetection
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorNewParamsBodyMonitorsCreatePageSemanticMonitorRequestChangeDetection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MonitorNewParamsBodyMonitorsCreatePageSemanticMonitorRequestChangeDetection](
		"type", "semantic",
	)
}

// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
// every 6 hours or every 2 days. The total interval (frequency × unit) must be
// between 10 minutes and 1 year.
//
// The properties Frequency, Type, Unit are required.
type MonitorNewParamsBodyMonitorsCreatePageSemanticMonitorRequestSchedule struct {
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

func (r MonitorNewParamsBodyMonitorsCreatePageSemanticMonitorRequestSchedule) MarshalJSON() (data []byte, err error) {
	type shadow MonitorNewParamsBodyMonitorsCreatePageSemanticMonitorRequestSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorNewParamsBodyMonitorsCreatePageSemanticMonitorRequestSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MonitorNewParamsBodyMonitorsCreatePageSemanticMonitorRequestSchedule](
		"type", "interval",
	)
	apijson.RegisterFieldValidator[MonitorNewParamsBodyMonitorsCreatePageSemanticMonitorRequestSchedule](
		"unit", "minutes", "hours", "days",
	)
}

// The properties Type, URL are required.
type MonitorNewParamsBodyMonitorsCreatePageSemanticMonitorRequestTarget struct {
	// Any of "page".
	Type string `json:"type,omitzero" api:"required"`
	URL  string `json:"url" api:"required" format:"uri"`
	// Normalize whitespace before comparing or analyzing text.
	NormalizeWhitespace param.Opt[bool] `json:"normalize_whitespace,omitzero"`
	paramObj
}

func (r MonitorNewParamsBodyMonitorsCreatePageSemanticMonitorRequestTarget) MarshalJSON() (data []byte, err error) {
	type shadow MonitorNewParamsBodyMonitorsCreatePageSemanticMonitorRequestTarget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorNewParamsBodyMonitorsCreatePageSemanticMonitorRequestTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MonitorNewParamsBodyMonitorsCreatePageSemanticMonitorRequestTarget](
		"type", "page",
	)
}

// The property URL is required.
type MonitorNewParamsBodyMonitorsCreatePageSemanticMonitorRequestWebhook struct {
	// Webhook URL called when a change is detected.
	URL string `json:"url" api:"required" format:"uri"`
	paramObj
}

func (r MonitorNewParamsBodyMonitorsCreatePageSemanticMonitorRequestWebhook) MarshalJSON() (data []byte, err error) {
	type shadow MonitorNewParamsBodyMonitorsCreatePageSemanticMonitorRequestWebhook
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorNewParamsBodyMonitorsCreatePageSemanticMonitorRequestWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monitor a website's extracted structured data for semantic changes described by
// a natural language query.
//
// The properties ChangeDetection, Name, Schedule, Target are required.
type MonitorNewParamsBodyMonitorsCreateExtractSemanticMonitorRequest struct {
	// Detect meaning-level changes that match a natural language query.
	ChangeDetection MonitorNewParamsBodyMonitorsCreateExtractSemanticMonitorRequestChangeDetection `json:"change_detection,omitzero" api:"required"`
	Name            string                                                                         `json:"name" api:"required"`
	// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
	// every 6 hours or every 2 days. The total interval (frequency × unit) must be
	// between 10 minutes and 1 year.
	Schedule MonitorNewParamsBodyMonitorsCreateExtractSemanticMonitorRequestSchedule `json:"schedule,omitzero" api:"required"`
	Target   MonitorNewParamsBodyMonitorsCreateExtractSemanticMonitorRequestTarget   `json:"target,omitzero" api:"required"`
	Webhook  MonitorNewParamsBodyMonitorsCreateExtractSemanticMonitorRequestWebhook  `json:"webhook,omitzero"`
	// User-defined tags for grouping and filtering monitors and their changes.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r MonitorNewParamsBodyMonitorsCreateExtractSemanticMonitorRequest) MarshalJSON() (data []byte, err error) {
	type shadow MonitorNewParamsBodyMonitorsCreateExtractSemanticMonitorRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorNewParamsBodyMonitorsCreateExtractSemanticMonitorRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detect meaning-level changes that match a natural language query.
//
// The properties Query, Type are required.
type MonitorNewParamsBodyMonitorsCreateExtractSemanticMonitorRequestChangeDetection struct {
	Query string `json:"query" api:"required"`
	// Any of "semantic".
	Type                string             `json:"type,omitzero" api:"required"`
	ConfidenceThreshold param.Opt[float64] `json:"confidence_threshold,omitzero"`
	paramObj
}

func (r MonitorNewParamsBodyMonitorsCreateExtractSemanticMonitorRequestChangeDetection) MarshalJSON() (data []byte, err error) {
	type shadow MonitorNewParamsBodyMonitorsCreateExtractSemanticMonitorRequestChangeDetection
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorNewParamsBodyMonitorsCreateExtractSemanticMonitorRequestChangeDetection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MonitorNewParamsBodyMonitorsCreateExtractSemanticMonitorRequestChangeDetection](
		"type", "semantic",
	)
}

// Run the monitor on a fixed interval defined by a frequency and a unit, e.g.
// every 6 hours or every 2 days. The total interval (frequency × unit) must be
// between 10 minutes and 1 year.
//
// The properties Frequency, Type, Unit are required.
type MonitorNewParamsBodyMonitorsCreateExtractSemanticMonitorRequestSchedule struct {
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

func (r MonitorNewParamsBodyMonitorsCreateExtractSemanticMonitorRequestSchedule) MarshalJSON() (data []byte, err error) {
	type shadow MonitorNewParamsBodyMonitorsCreateExtractSemanticMonitorRequestSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorNewParamsBodyMonitorsCreateExtractSemanticMonitorRequestSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MonitorNewParamsBodyMonitorsCreateExtractSemanticMonitorRequestSchedule](
		"type", "interval",
	)
	apijson.RegisterFieldValidator[MonitorNewParamsBodyMonitorsCreateExtractSemanticMonitorRequestSchedule](
		"unit", "minutes", "hours", "days",
	)
}

// The properties Type, URL are required.
type MonitorNewParamsBodyMonitorsCreateExtractSemanticMonitorRequestTarget struct {
	// Any of "extract".
	Type string `json:"type,omitzero" api:"required"`
	// Root URL to extract structured data from.
	URL              string          `json:"url" api:"required" format:"uri"`
	FollowSubdomains param.Opt[bool] `json:"follow_subdomains,omitzero"`
	// Optional natural-language instructions guiding what to extract.
	Instructions param.Opt[string] `json:"instructions,omitzero"`
	// Optional maximum link depth from the starting URL (0 = only the starting page).
	MaxDepth param.Opt[int64] `json:"max_depth,omitzero"`
	// Maximum number of pages to analyze during extraction.
	MaxPages param.Opt[int64] `json:"max_pages,omitzero"`
	// JSON Schema describing the structured data to extract and watch for changes. If
	// omitted, a default summary + key-points schema is used.
	Schema map[string]any `json:"schema,omitzero"`
	paramObj
}

func (r MonitorNewParamsBodyMonitorsCreateExtractSemanticMonitorRequestTarget) MarshalJSON() (data []byte, err error) {
	type shadow MonitorNewParamsBodyMonitorsCreateExtractSemanticMonitorRequestTarget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorNewParamsBodyMonitorsCreateExtractSemanticMonitorRequestTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MonitorNewParamsBodyMonitorsCreateExtractSemanticMonitorRequestTarget](
		"type", "extract",
	)
}

// The property URL is required.
type MonitorNewParamsBodyMonitorsCreateExtractSemanticMonitorRequestWebhook struct {
	// Webhook URL called when a change is detected.
	URL string `json:"url" api:"required" format:"uri"`
	paramObj
}

func (r MonitorNewParamsBodyMonitorsCreateExtractSemanticMonitorRequestWebhook) MarshalJSON() (data []byte, err error) {
	type shadow MonitorNewParamsBodyMonitorsCreateExtractSemanticMonitorRequestWebhook
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorNewParamsBodyMonitorsCreateExtractSemanticMonitorRequestWebhook) UnmarshalJSON(data []byte) error {
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

// Detect meaning-level changes that match a natural language query.
//
// The properties Query, Type are required.
type MonitorUpdateParamsChangeDetectionSemantic struct {
	Query               string             `json:"query" api:"required"`
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

// The properties Type, URL are required.
type MonitorUpdateParamsTargetSitemap struct {
	// Sitemap URL to monitor.
	URL     string           `json:"url" api:"required" format:"uri"`
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

// The properties Type, URL are required.
type MonitorUpdateParamsTargetExtract struct {
	// Root URL to extract structured data from.
	URL              string          `json:"url" api:"required" format:"uri"`
	FollowSubdomains param.Opt[bool] `json:"follow_subdomains,omitzero"`
	// Optional natural-language instructions guiding what to extract.
	Instructions param.Opt[string] `json:"instructions,omitzero"`
	// Optional maximum link depth from the starting URL (0 = only the starting page).
	MaxDepth param.Opt[int64] `json:"max_depth,omitzero"`
	// Maximum number of pages to analyze during extraction.
	MaxPages param.Opt[int64] `json:"max_pages,omitzero"`
	// JSON Schema describing the structured data to extract and watch for changes. If
	// omitted, a default summary + key-points schema is used.
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
	// Webhook URL called when a change is detected.
	URL string `json:"url" api:"required" format:"uri"`
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
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	// Filter to items that have this tag.
	Tag param.Opt[string] `query:"tag,omitzero" json:"-"`
	// Any of "exact", "semantic".
	ChangeDetectionType MonitorListParamsChangeDetectionType `query:"change_detection_type,omitzero" json:"-"`
	// Any of "active", "paused", "failed".
	Status MonitorListParamsStatus `query:"status,omitzero" json:"-"`
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

type MonitorListParamsChangeDetectionType string

const (
	MonitorListParamsChangeDetectionTypeExact    MonitorListParamsChangeDetectionType = "exact"
	MonitorListParamsChangeDetectionTypeSemantic MonitorListParamsChangeDetectionType = "semantic"
)

type MonitorListParamsStatus string

const (
	MonitorListParamsStatusActive MonitorListParamsStatus = "active"
	MonitorListParamsStatusPaused MonitorListParamsStatus = "paused"
	MonitorListParamsStatusFailed MonitorListParamsStatus = "failed"
)

type MonitorListParamsTargetType string

const (
	MonitorListParamsTargetTypePage    MonitorListParamsTargetType = "page"
	MonitorListParamsTargetTypeSitemap MonitorListParamsTargetType = "sitemap"
	MonitorListParamsTargetTypeExtract MonitorListParamsTargetType = "extract"
)

type MonitorListAccountChangesParams struct {
	Cursor    param.Opt[string]    `query:"cursor,omitzero" json:"-"`
	Limit     param.Opt[int64]     `query:"limit,omitzero" json:"-"`
	MonitorID param.Opt[string]    `query:"monitor_id,omitzero" json:"-"`
	Since     param.Opt[time.Time] `query:"since,omitzero" format:"date-time" json:"-"`
	// Filter to items that have this tag.
	Tag   param.Opt[string]    `query:"tag,omitzero" json:"-"`
	Until param.Opt[time.Time] `query:"until,omitzero" format:"date-time" json:"-"`
	// Any of "exact", "semantic".
	ChangeDetectionType MonitorListAccountChangesParamsChangeDetectionType `query:"change_detection_type,omitzero" json:"-"`
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

type MonitorListAccountChangesParamsChangeDetectionType string

const (
	MonitorListAccountChangesParamsChangeDetectionTypeExact    MonitorListAccountChangesParamsChangeDetectionType = "exact"
	MonitorListAccountChangesParamsChangeDetectionTypeSemantic MonitorListAccountChangesParamsChangeDetectionType = "semantic"
)

type MonitorListAccountChangesParamsTargetType string

const (
	MonitorListAccountChangesParamsTargetTypePage    MonitorListAccountChangesParamsTargetType = "page"
	MonitorListAccountChangesParamsTargetTypeSitemap MonitorListAccountChangesParamsTargetType = "sitemap"
	MonitorListAccountChangesParamsTargetTypeExtract MonitorListAccountChangesParamsTargetType = "extract"
)

type MonitorListAccountRunsParams struct {
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	// Any of "queued", "running", "completed", "failed".
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

type MonitorListAccountRunsParamsStatus string

const (
	MonitorListAccountRunsParamsStatusQueued    MonitorListAccountRunsParamsStatus = "queued"
	MonitorListAccountRunsParamsStatusRunning   MonitorListAccountRunsParamsStatus = "running"
	MonitorListAccountRunsParamsStatusCompleted MonitorListAccountRunsParamsStatus = "completed"
	MonitorListAccountRunsParamsStatusFailed    MonitorListAccountRunsParamsStatus = "failed"
)

type MonitorListChangesParams struct {
	Cursor param.Opt[string]    `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]     `query:"limit,omitzero" json:"-"`
	Since  param.Opt[time.Time] `query:"since,omitzero" format:"date-time" json:"-"`
	// Filter to items that have this tag.
	Tag   param.Opt[string]    `query:"tag,omitzero" json:"-"`
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
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	// Any of "queued", "running", "completed", "failed".
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

type MonitorListRunsParamsStatus string

const (
	MonitorListRunsParamsStatusQueued    MonitorListRunsParamsStatus = "queued"
	MonitorListRunsParamsStatusRunning   MonitorListRunsParamsStatus = "running"
	MonitorListRunsParamsStatusCompleted MonitorListRunsParamsStatus = "completed"
	MonitorListRunsParamsStatusFailed    MonitorListRunsParamsStatus = "failed"
)
