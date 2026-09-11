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

// Read your organization's API request logs to debug failed calls. These endpoints
// cost no credits and use a separate rate limit.
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

// Get one logged API call, including its request input and response body.
func (r *LogService) Get(ctx context.Context, requestID string, opts ...option.RequestOption) (res *LogGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if requestID == "" {
		err = errors.New("missing required request_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("logs/%s", requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List your organization's API requests, newest first. Defaults to the last 24
// hours.
func (r *LogService) List(ctx context.Context, query LogListParams, opts ...option.RequestOption) (res *LogListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "logs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type LogGetResponse struct {
	Data LogGetResponseData `json:"data" api:"required"`
	// Unique id of this API call, also sent in the X-Request-Id response header. Quote
	// it when contacting support about a failed request.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Credit usage, included whenever a valid API key is provided.
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
	// The `error_code` from the response, or null on success.
	ErrorCode string `json:"error_code" api:"required"`
	// What was sent with the request.
	Input LogGetResponseDataInput `json:"input" api:"required"`
	// ID of the API key that made the request.
	KeyID string `json:"key_id" api:"required"`
	// Server-side processing time in milliseconds.
	LatencyMs float64 `json:"latency_ms" api:"required"`
	// HTTP method.
	Method string `json:"method" api:"required"`
	// Endpoint path as called.
	Path string `json:"path" api:"required"`
	// Request ID of the logged API call.
	RequestID string `json:"request_id" api:"required"`
	// HTTP status code returned.
	StatusCode int64 `json:"status_code" api:"required"`
	// Request tags supplied by the caller.
	Tags []string `json:"tags" api:"required"`
	// When the request completed.
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// User-Agent header of the request.
	UserAgent string `json:"user_agent" api:"required"`
	// Whether the request was made under zero data retention.
	Zdr bool `json:"zdr" api:"required"`
	// Credit usage, included whenever a valid API key is provided.
	KeyMetadata LogGetResponseDataKeyMetadata `json:"key_metadata"`
	// The retained JSON response with credentials redacted, or null when unavailable.
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
		KeyMetadata respjson.Field
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

// What was sent with the request.
type LogGetResponseDataInput struct {
	// Query parameters as sent.
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

// Credit usage, included whenever a valid API key is provided.
type LogGetResponseDataKeyMetadata struct {
	// Credits used by this request.
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
func (r LogGetResponseDataKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *LogGetResponseDataKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credit usage, included whenever a valid API key is provided.
type LogGetResponseKeyMetadata struct {
	// Credits used by this request.
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
func (r LogGetResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *LogGetResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogListResponse struct {
	// Log entries, newest first.
	Data []LogListResponseData `json:"data" api:"required"`
	// Whether a next page exists.
	HasMore bool `json:"has_more" api:"required"`
	// Entries per page.
	Limit int64 `json:"limit" api:"required"`
	// Current page number.
	Page int64 `json:"page" api:"required"`
	// Unique id of this API call, also sent in the X-Request-Id response header. Quote
	// it when contacting support about a failed request.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Credit usage, included whenever a valid API key is provided.
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
	// The `error_code` from the response, or null on success.
	ErrorCode string `json:"error_code" api:"required"`
	// ID of the API key that made the request.
	KeyID string `json:"key_id" api:"required"`
	// Server-side processing time in milliseconds.
	LatencyMs float64 `json:"latency_ms" api:"required"`
	// HTTP method.
	Method string `json:"method" api:"required"`
	// Endpoint path as called.
	Path string `json:"path" api:"required"`
	// Request ID of the logged API call.
	RequestID string `json:"request_id" api:"required"`
	// HTTP status code returned.
	StatusCode int64 `json:"status_code" api:"required"`
	// Request tags supplied by the caller.
	Tags []string `json:"tags" api:"required"`
	// When the request completed.
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// Whether the request was made under zero data retention.
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

// Credit usage, included whenever a valid API key is provided.
type LogListResponseKeyMetadata struct {
	// Credits used by this request.
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
func (r LogListResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *LogListResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogListParams struct {
	// Filter by the `error_code` returned in the response.
	ErrorCode param.Opt[string] `query:"error_code,omitzero" json:"-"`
	// Only include requests that returned a 4xx or 5xx status.
	ErrorsOnly param.Opt[bool] `query:"errors_only,omitzero" json:"-"`
	// Only include requests at or after this ISO 8601 timestamp. Defaults to 24 hours
	// before `to`.
	From param.Opt[time.Time] `query:"from,omitzero" format:"date-time" json:"-"`
	// Filter by the API key that made the request.
	KeyID param.Opt[string] `query:"key_id,omitzero" json:"-"`
	// Number of log entries per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number, starting at 1.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Filter by endpoint path, with or without the /v1 prefix.
	Path param.Opt[string] `query:"path,omitzero" json:"-"`
	// Case-insensitive substring match against the request query and body, e.g. a
	// domain.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// Filter by exact HTTP status code.
	StatusCode param.Opt[int64] `query:"status_code,omitzero" json:"-"`
	// Comma-separated request tags. Matches requests carrying any of them. Up to 20
	// tags, each 1-50 characters.
	Tags param.Opt[string] `query:"tags,omitzero" json:"-"`
	// Only include requests at or before this ISO 8601 timestamp. Defaults to now.
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
