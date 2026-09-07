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
)

// Inspect and retry batch and monitor webhook deliveries without rerunning the
// underlying work.
//
// WebhookDeliveryService contains methods and other services that help with
// interacting with the context.dev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookDeliveryService] method instead.
type WebhookDeliveryService struct {
	options []option.RequestOption
}

// NewWebhookDeliveryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWebhookDeliveryService(opts ...option.RequestOption) (r WebhookDeliveryService) {
	r = WebhookDeliveryService{}
	r.options = opts
	return
}

// Get the live status, retry policy, latest attempt, and replay expiration for a
// retained delivery. Use the attempts endpoint for its complete paginated history.
// This endpoint costs no credits.
func (r *WebhookDeliveryService) Get(ctx context.Context, deliveryID string, query WebhookDeliveryGetParams, opts ...option.RequestOption) (res *WebhookDeliveryGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if deliveryID == "" {
		err = errors.New("missing required delivery_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks/deliveries/%s", url.PathEscape(deliveryID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List retained batch and monitor webhook deliveries for your organization, newest
// first. Filter by at most one of batch_id, monitor_id, or run_id, optionally
// combined with status. Historical events without retained payloads are not
// listed. This endpoint costs no credits.
func (r *WebhookDeliveryService) List(ctx context.Context, query WebhookDeliveryListParams, opts ...option.RequestOption) (res *WebhookDeliveryListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "webhooks/deliveries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List individual HTTP attempts for a delivery, newest first, including their
// destination, timestamps, HTTP status, and error. An interrupted attempt may have
// reached the endpoint even when its outcome is unknown. This endpoint costs no
// credits.
func (r *WebhookDeliveryService) ListAttempts(ctx context.Context, deliveryID string, query WebhookDeliveryListAttemptsParams, opts ...option.RequestOption) (res *WebhookDeliveryListAttemptsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if deliveryID == "" {
		err = errors.New("missing required delivery_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks/deliveries/%s/attempts", url.PathEscape(deliveryID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Queue an immediate attempt without rerunning or billing the underlying batch or
// monitor. A waiting retry is brought forward. A failed delivery gets one
// additional attempt without restarting its automatic retry budget. Set force:
// true to resend an acknowledged delivery. An in-progress attempt cannot be
// duplicated. The stored event body, event ID, and creation time remain unchanged;
// each attempt receives a fresh signature. Monitor retries use the current URL and
// secret; removing the webhook cancels pending deliveries. Batch result URLs in
// old payloads may have expired: retrieve the batch to get fresh URLs. Replay is
// available for seven days. A successful attempt cancels remaining automatic
// retries. Idempotency-Key is scoped to your organization and retained with the
// delivery metadata; repeating the same key and input returns the original
// accepted response.
func (r *WebhookDeliveryService) Retry(ctx context.Context, deliveryID string, params WebhookDeliveryRetryParams, opts ...option.RequestOption) (res *WebhookDeliveryRetryResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if deliveryID == "" {
		err = errors.New("missing required delivery_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks/deliveries/%s/retry", url.PathEscape(deliveryID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type Attempt struct {
	ID          string       `json:"id" api:"required"`
	Attempt     int64        `json:"attempt" api:"required"`
	CompletedAt time.Time    `json:"completed_at" api:"required" format:"date-time"`
	Error       AttemptError `json:"error" api:"required"`
	HTTPStatus  int64        `json:"http_status" api:"required"`
	StartedAt   time.Time    `json:"started_at" api:"required" format:"date-time"`
	// Any of "initial", "automatic", "manual".
	Trigger AttemptTrigger `json:"trigger" api:"required"`
	URL     string         `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Attempt     respjson.Field
		CompletedAt respjson.Field
		Error       respjson.Field
		HTTPStatus  respjson.Field
		StartedAt   respjson.Field
		Trigger     respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Attempt) RawJSON() string { return r.JSON.raw }
func (r *Attempt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttemptError struct {
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
func (r AttemptError) RawJSON() string { return r.JSON.raw }
func (r *AttemptError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttemptTrigger string

const (
	AttemptTriggerInitial   AttemptTrigger = "initial"
	AttemptTriggerAutomatic AttemptTrigger = "automatic"
	AttemptTriggerManual    AttemptTrigger = "manual"
)

type Delivery struct {
	ID string `json:"id" api:"required"`
	// Number of delivery attempts started, including any attempt in progress.
	AttemptCount int64     `json:"attempt_count" api:"required"`
	CreatedAt    time.Time `json:"created_at" api:"required" format:"date-time"`
	// Most recent successful acknowledgment; retained if a later forced resend fails.
	DeliveredAt time.Time `json:"delivered_at" api:"required" format:"date-time"`
	// Any of "batch.completed", "batch.failed", "batch.cancelled", "change.detected",
	// "run.completed".
	Event DeliveryEvent `json:"event" api:"required"`
	// Stable event ID. Unchanged across automatic and manual attempts; use it to
	// deduplicate events.
	EventID       string              `json:"event_id" api:"required"`
	LastAttempt   DeliveryLastAttempt `json:"last_attempt" api:"required"`
	LastError     DeliveryLastError   `json:"last_error" api:"required"`
	NextAttemptAt time.Time           `json:"next_attempt_at" api:"required" format:"date-time"`
	// Opt into durable webhook delivery. An empty object uses the default retry
	// schedule. Omit retry to preserve legacy delivery behavior. The policy is
	// snapshotted for each event.
	Retry RetryConfig `json:"retry" api:"required"`
	// Seven days after event creation. Manual retries after this time return 410.
	// Delivery and attempt metadata remain available for up to 30 days.
	RetryExpiresAt time.Time           `json:"retry_expires_at" api:"required" format:"date-time"`
	Source         DeliverySourceUnion `json:"source" api:"required"`
	// Any of "pending", "delivering", "retrying", "delivered", "failed", "cancelled".
	Status DeliveryStatus `json:"status" api:"required"`
	// Destination recorded for this delivery. Each attempt records the URL it used.
	// Monitor retries use the currently configured URL and signing secret.
	URL string `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		AttemptCount   respjson.Field
		CreatedAt      respjson.Field
		DeliveredAt    respjson.Field
		Event          respjson.Field
		EventID        respjson.Field
		LastAttempt    respjson.Field
		LastError      respjson.Field
		NextAttemptAt  respjson.Field
		Retry          respjson.Field
		RetryExpiresAt respjson.Field
		Source         respjson.Field
		Status         respjson.Field
		URL            respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Delivery) RawJSON() string { return r.JSON.raw }
func (r *Delivery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeliveryEvent string

const (
	DeliveryEventBatchCompleted DeliveryEvent = "batch.completed"
	DeliveryEventBatchFailed    DeliveryEvent = "batch.failed"
	DeliveryEventBatchCancelled DeliveryEvent = "batch.cancelled"
	DeliveryEventChangeDetected DeliveryEvent = "change.detected"
	DeliveryEventRunCompleted   DeliveryEvent = "run.completed"
)

type DeliveryLastAttempt struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Attempt
}

// Returns the unmodified JSON received from the API
func (r DeliveryLastAttempt) RawJSON() string { return r.JSON.raw }
func (r *DeliveryLastAttempt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeliveryLastError struct {
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
func (r DeliveryLastError) RawJSON() string { return r.JSON.raw }
func (r *DeliveryLastError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DeliverySourceUnion contains all possible properties and values from
// [DeliverySourceObject], [DeliverySourceObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type DeliverySourceUnion struct {
	// This field is from variant [DeliverySourceObject].
	BatchID string `json:"batch_id"`
	Type    string `json:"type"`
	// This field is from variant [DeliverySourceObject2].
	MonitorID string `json:"monitor_id"`
	// This field is from variant [DeliverySourceObject2].
	RunID string `json:"run_id"`
	JSON  struct {
		BatchID   respjson.Field
		Type      respjson.Field
		MonitorID respjson.Field
		RunID     respjson.Field
		raw       string
	} `json:"-"`
}

func (u DeliverySourceUnion) AsDeliverySourceObject() (v DeliverySourceObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DeliverySourceUnion) AsDeliverySourceObject2() (v DeliverySourceObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DeliverySourceUnion) RawJSON() string { return u.JSON.raw }

func (r *DeliverySourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeliverySourceObject struct {
	BatchID string `json:"batch_id" api:"required"`
	// Any of "batch".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BatchID     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeliverySourceObject) RawJSON() string { return r.JSON.raw }
func (r *DeliverySourceObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeliverySourceObject2 struct {
	MonitorID string `json:"monitor_id" api:"required"`
	RunID     string `json:"run_id" api:"required"`
	// Any of "monitor".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MonitorID   respjson.Field
		RunID       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeliverySourceObject2) RawJSON() string { return r.JSON.raw }
func (r *DeliverySourceObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeliveryStatus string

const (
	DeliveryStatusPending    DeliveryStatus = "pending"
	DeliveryStatusDelivering DeliveryStatus = "delivering"
	DeliveryStatusRetrying   DeliveryStatus = "retrying"
	DeliveryStatusDelivered  DeliveryStatus = "delivered"
	DeliveryStatusFailed     DeliveryStatus = "failed"
	DeliveryStatusCancelled  DeliveryStatus = "cancelled"
)

type WebhookDeliveryGetResponse struct {
	// Metadata about the API key used for the request. Included in every response
	// whenever a valid API key is provided, even when the response status is not 200.
	KeyMetadata WebhookDeliveryGetResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		KeyMetadata respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Delivery
}

// Returns the unmodified JSON received from the API
func (r WebhookDeliveryGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the API key used for the request. Included in every response
// whenever a valid API key is provided, even when the response status is not 200.
type WebhookDeliveryGetResponseKeyMetadata struct {
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
func (r WebhookDeliveryGetResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryGetResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeliveryListResponse struct {
	Data       []Delivery `json:"data" api:"required"`
	HasMore    bool       `json:"has_more" api:"required"`
	NextCursor string     `json:"next_cursor" api:"required"`
	// Metadata about the API key used for the request. Included in every response
	// whenever a valid API key is provided, even when the response status is not 200.
	KeyMetadata WebhookDeliveryListResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		NextCursor  respjson.Field
		KeyMetadata respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeliveryListResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the API key used for the request. Included in every response
// whenever a valid API key is provided, even when the response status is not 200.
type WebhookDeliveryListResponseKeyMetadata struct {
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
func (r WebhookDeliveryListResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryListResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeliveryListAttemptsResponse struct {
	Data       []Attempt `json:"data" api:"required"`
	HasMore    bool      `json:"has_more" api:"required"`
	NextCursor string    `json:"next_cursor" api:"required"`
	// Metadata about the API key used for the request. Included in every response
	// whenever a valid API key is provided, even when the response status is not 200.
	KeyMetadata WebhookDeliveryListAttemptsResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		NextCursor  respjson.Field
		KeyMetadata respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeliveryListAttemptsResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryListAttemptsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the API key used for the request. Included in every response
// whenever a valid API key is provided, even when the response status is not 200.
type WebhookDeliveryListAttemptsResponseKeyMetadata struct {
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
func (r WebhookDeliveryListAttemptsResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryListAttemptsResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeliveryRetryResponse struct {
	// Metadata about the API key used for the request. Included in every response
	// whenever a valid API key is provided, even when the response status is not 200.
	KeyMetadata WebhookDeliveryRetryResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		KeyMetadata respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Delivery
}

// Returns the unmodified JSON received from the API
func (r WebhookDeliveryRetryResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryRetryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the API key used for the request. Included in every response
// whenever a valid API key is provided, even when the response status is not 200.
type WebhookDeliveryRetryResponseKeyMetadata struct {
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
func (r WebhookDeliveryRetryResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryRetryResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeliveryGetParams struct {
	// Optional comma-separated caller-defined tags for tracking this request. Tags are
	// recorded on the request's usage log and can be used to filter usage on the
	// dashboard usage page. Up to 20 tags, each 1-50 characters.
	Tags []string `query:"tags,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookDeliveryGetParams]'s query parameters as
// `url.Values`.
func (r WebhookDeliveryGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookDeliveryListParams struct {
	BatchID   param.Opt[string] `query:"batch_id,omitzero" json:"-"`
	Cursor    param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit     param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	MonitorID param.Opt[string] `query:"monitor_id,omitzero" json:"-"`
	RunID     param.Opt[string] `query:"run_id,omitzero" json:"-"`
	// Any of "pending", "delivering", "retrying", "delivered", "failed", "cancelled".
	Status WebhookDeliveryListParamsStatus `query:"status,omitzero" json:"-"`
	// Optional comma-separated caller-defined tags for tracking this request. Tags are
	// recorded on the request's usage log and can be used to filter usage on the
	// dashboard usage page. Up to 20 tags, each 1-50 characters.
	Tags []string `query:"tags,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookDeliveryListParams]'s query parameters as
// `url.Values`.
func (r WebhookDeliveryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookDeliveryListParamsStatus string

const (
	WebhookDeliveryListParamsStatusPending    WebhookDeliveryListParamsStatus = "pending"
	WebhookDeliveryListParamsStatusDelivering WebhookDeliveryListParamsStatus = "delivering"
	WebhookDeliveryListParamsStatusRetrying   WebhookDeliveryListParamsStatus = "retrying"
	WebhookDeliveryListParamsStatusDelivered  WebhookDeliveryListParamsStatus = "delivered"
	WebhookDeliveryListParamsStatusFailed     WebhookDeliveryListParamsStatus = "failed"
	WebhookDeliveryListParamsStatusCancelled  WebhookDeliveryListParamsStatus = "cancelled"
)

type WebhookDeliveryListAttemptsParams struct {
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	// Optional comma-separated caller-defined tags for tracking this request. Tags are
	// recorded on the request's usage log and can be used to filter usage on the
	// dashboard usage page. Up to 20 tags, each 1-50 characters.
	Tags []string `query:"tags,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookDeliveryListAttemptsParams]'s query parameters as
// `url.Values`.
func (r WebhookDeliveryListAttemptsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookDeliveryRetryParams struct {
	Force          param.Opt[bool]   `json:"force,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	// Optional tags for tracking usage. Up to 20 tags, each 1 to 50 characters.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r WebhookDeliveryRetryParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookDeliveryRetryParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookDeliveryRetryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
