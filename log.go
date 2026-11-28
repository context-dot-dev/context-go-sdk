// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contextdev

import (
	"context"
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
)

// Inspect your organization's API requests, including retained inputs and
// responses.
//
// LogService contains methods and other services that help with interacting with
// the context.dev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLogService] method instead.
type LogService struct {
	options []option.RequestOption
}

// NewLogService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLogService(opts ...option.RequestOption) (r LogService) {
	r = LogService{}
	r.options = opts
	return
}

// Retrieve a request's metadata, retained input, and response. A request that just
// finished can return 404 `NOT_FOUND` for a few seconds, so retry shortly.
func (r *LogService) Get(ctx context.Context, requestID string, opts ...option.RequestOption) (res *LogGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if requestID == "" {
		err = errors.New("missing required request_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("org/logs/%s", requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List your organization's API requests, newest first. Filter by time, endpoint,
// API key, status, error code, tags, or text. Includes batch settlements and
// monitor runs.
func (r *LogService) List(ctx context.Context, query LogListParams, opts ...option.RequestOption) (res *LogListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "org/logs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type LogGetResponse struct {
	Data LogGetResponseData `json:"data" api:"required"`
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Credits this request used and your remaining balance.
	KeyMetadata LogGetResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		RequestID   respjson.Field
		KeyMetadata respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogGetResponse) RawJSON() string { return r.JSON.raw }
func (r *LogGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogGetResponseData struct {
	// Credits charged for this request.
	CreditsUsed int64 `json:"credits_used" api:"required"`
	// The response's `error_code`. `null` on success.
	ErrorCode string `json:"error_code" api:"required"`
	// Input retained from the request.
	Input LogGetResponseDataInput `json:"input" api:"required"`
	// ID of the API key that made the request. Can be `null` for `MONITOR` and `BATCH`
	// entries.
	KeyID string `json:"key_id" api:"required"`
	// Server-side processing time, in milliseconds.
	LatencyMs float64 `json:"latency_ms" api:"required"`
	// HTTP method, or `MONITOR` for a monitor run and `BATCH` for a batch settlement.
	Method string `json:"method" api:"required"`
	// Endpoint path that was called.
	Path string `json:"path" api:"required"`
	// ID of the logged request.
	RequestID string `json:"request_id" api:"required"`
	// HTTP status code of the response.
	StatusCode int64 `json:"status_code" api:"required"`
	// Tags sent with the request.
	Tags []string `json:"tags" api:"required"`
	// When the request completed.
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// The request's `User-Agent` header.
	UserAgent string `json:"user_agent" api:"required"`
	// Whether the request used zero data retention (ZDR).
	Zdr bool `json:"zdr" api:"required"`
	// Retained JSON response with credentials redacted. `null` when unavailable, such
	// as under ZDR.
	Response any `json:"response"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditsUsed respjson.Field
		ErrorCode   respjson.Field
		Input       respjson.Field
		KeyID       respjson.Field
		LatencyMs   respjson.Field
		Method      respjson.Field
		Path        respjson.Field
		RequestID   respjson.Field
		StatusCode  respjson.Field
		Tags        respjson.Field
		Timestamp   respjson.Field
		UserAgent   respjson.Field
		Zdr         respjson.Field
		Response    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *LogGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Input retained from the request.
type LogGetResponseDataInput struct {
	// Query parameters of the request.
	Query map[string]any `json:"query" api:"required"`
	// Request body with credentials and uploaded content redacted.
	Body any `json:"body"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Query       respjson.Field
		Body        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogGetResponseDataInput) RawJSON() string { return r.JSON.raw }
func (r *LogGetResponseDataInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credits this request used and your remaining balance.
type LogGetResponseKeyMetadata struct {
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
func (r LogGetResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *LogGetResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogListResponse struct {
	// Log entries on this page, newest first.
	Data []LogListResponseData `json:"data" api:"required"`
	// Whether more results are available after this page.
	HasMore bool `json:"has_more" api:"required"`
	// Maximum number of log entries to return per page.
	Limit int64 `json:"limit" api:"required"`
	// Current page number.
	Page int64 `json:"page" api:"required"`
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Credits this request used and your remaining balance.
	KeyMetadata LogListResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		Limit       respjson.Field
		Page        respjson.Field
		RequestID   respjson.Field
		KeyMetadata respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogListResponse) RawJSON() string { return r.JSON.raw }
func (r *LogListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogListResponseData struct {
	// Credits charged for this request.
	CreditsUsed int64 `json:"credits_used" api:"required"`
	// The response's `error_code`. `null` on success.
	ErrorCode string `json:"error_code" api:"required"`
	// ID of the API key that made the request. Can be `null` for `MONITOR` and `BATCH`
	// entries.
	KeyID string `json:"key_id" api:"required"`
	// Server-side processing time, in milliseconds.
	LatencyMs float64 `json:"latency_ms" api:"required"`
	// HTTP method, or `MONITOR` for a monitor run and `BATCH` for a batch settlement.
	Method string `json:"method" api:"required"`
	// Endpoint path that was called.
	Path string `json:"path" api:"required"`
	// ID of the logged request.
	RequestID string `json:"request_id" api:"required"`
	// HTTP status code of the response.
	StatusCode int64 `json:"status_code" api:"required"`
	// Tags sent with the request.
	Tags []string `json:"tags" api:"required"`
	// When the request completed.
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// Whether the request used zero data retention (ZDR).
	Zdr bool `json:"zdr" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditsUsed respjson.Field
		ErrorCode   respjson.Field
		KeyID       respjson.Field
		LatencyMs   respjson.Field
		Method      respjson.Field
		Path        respjson.Field
		RequestID   respjson.Field
		StatusCode  respjson.Field
		Tags        respjson.Field
		Timestamp   respjson.Field
		Zdr         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogListResponseData) RawJSON() string { return r.JSON.raw }
func (r *LogListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credits this request used and your remaining balance.
type LogListResponseKeyMetadata struct {
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
func (r LogListResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *LogListResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogListParams struct {
	// Filter by the `error_code` returned in the response.
	ErrorCode param.Opt[string] `query:"error_code,omitzero" json:"-"`
	// Return only requests that failed with a 4xx or 5xx status.
	ErrorsOnly param.Opt[bool] `query:"errors_only,omitzero" json:"-"`
	// Return requests at or after this time. Defaults to 24 hours before `to`.
	From param.Opt[time.Time] `query:"from,omitzero" format:"date-time" json:"-"`
	// Filter by the ID of the API key that made the request.
	KeyID param.Opt[string] `query:"key_id,omitzero" json:"-"`
	// Maximum number of log entries to return per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number, starting at 1.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Filter by endpoint path, with or without the `/v1` prefix.
	Path param.Opt[string] `query:"path,omitzero" json:"-"`
	// Case-insensitive text to find in the request query or body, such as a domain.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// Filter by exact HTTP status code.
	StatusCode param.Opt[int64] `query:"status_code,omitzero" json:"-"`
	// Comma-separated tags. Returns log entries that have any of them. Up to 20 tags
	// of 1 to 50 characters each.
	Tags param.Opt[string] `query:"tags,omitzero" json:"-"`
	// Return requests at or before this time. Defaults to now.
	To param.Opt[time.Time] `query:"to,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [LogListParams]'s query parameters as `url.Values`.
func (r LogListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
