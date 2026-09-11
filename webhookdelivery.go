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

// Inspect and retry webhook deliveries. These endpoints cost no credits.
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

// Get a webhook delivery, including its status and latest attempt.
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

// List your batch or monitor webhook deliveries, newest first.
func (r *WebhookDeliveryService) List(ctx context.Context, body WebhookDeliveryListParams, opts ...option.RequestOption) (res *WebhookDeliveryListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "webhooks/deliveries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List delivery attempts, newest first.
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

// Retry a webhook delivery within seven days of creation.
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
	// Attempt number, starting at 1.
	Attempt int64 `json:"attempt" api:"required"`
	// Completion time, or null while in progress.
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	// Attempt error, or null if none.
	Error AttemptError `json:"error" api:"required"`
	// HTTP response status, or null if no response was received.
	HTTPStatus int64 `json:"http_status" api:"required"`
	// Attempt start time.
	StartedAt time.Time `json:"started_at" api:"required" format:"date-time"`
	// What started this attempt.
	//
	// Any of "initial", "automatic", "manual".
	Trigger AttemptTrigger `json:"trigger" api:"required"`
	// URL used for this attempt.
	URL string `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
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

// Attempt error, or null if none.
type AttemptError struct {
	// Error code.
	Code string `json:"code" api:"required"`
	// Error details.
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

// What started this attempt.
type AttemptTrigger string

const (
	AttemptTriggerInitial   AttemptTrigger = "initial"
	AttemptTriggerAutomatic AttemptTrigger = "automatic"
	AttemptTriggerManual    AttemptTrigger = "manual"
)

type Delivery struct {
	// Delivery ID.
	ID string `json:"id" api:"required"`
	// Event creation time.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Last successful delivery time, or null if never delivered.
	DeliveredAt time.Time `json:"delivered_at" api:"required" format:"date-time"`
	// Webhook event type.
	//
	// Any of "batch.completed", "batch.failed", "batch.cancelled", "change.detected",
	// "run.completed".
	Event DeliveryEvent `json:"event" api:"required"`
	// Stable event ID for deduplicating received webhooks.
	EventID string `json:"event_id" api:"required"`
	// Latest attempt, or null if none.
	LastAttempt Attempt `json:"last_attempt" api:"required"`
	// Latest delivery error, or null if none.
	LastError DeliveryLastError `json:"last_error" api:"required"`
	// Next scheduled attempt, or null if none.
	NextAttemptAt time.Time `json:"next_attempt_at" api:"required" format:"date-time"`
	// Webhook retry settings. Use {} for the default schedule.
	Retry RetryConfig `json:"retry" api:"required"`
	// Manual retry deadline, seven days after event creation.
	RetryExpiresAt time.Time `json:"retry_expires_at" api:"required" format:"date-time"`
	// Batch or monitor run that produced the event.
	Source DeliverySourceUnion `json:"source" api:"required"`
	// Current delivery status.
	//
	// Any of "pending", "delivering", "retrying", "delivered", "failed", "cancelled".
	Status DeliveryStatus `json:"status" api:"required"`
	// Webhook destination URL.
	URL string `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
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

// Webhook event type.
type DeliveryEvent string

const (
	DeliveryEventBatchCompleted DeliveryEvent = "batch.completed"
	DeliveryEventBatchFailed    DeliveryEvent = "batch.failed"
	DeliveryEventBatchCancelled DeliveryEvent = "batch.cancelled"
	DeliveryEventChangeDetected DeliveryEvent = "change.detected"
	DeliveryEventRunCompleted   DeliveryEvent = "run.completed"
)

// Latest delivery error, or null if none.
type DeliveryLastError struct {
	// Error code.
	Code string `json:"code" api:"required"`
	// Error details.
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
// [DeliverySourceBatch], [DeliverySourceMonitor].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type DeliverySourceUnion struct {
	// This field is from variant [DeliverySourceBatch].
	BatchID string `json:"batch_id"`
	Type    string `json:"type"`
	// This field is from variant [DeliverySourceMonitor].
	MonitorID string `json:"monitor_id"`
	// This field is from variant [DeliverySourceMonitor].
	RunID string `json:"run_id"`
	JSON  struct {
		BatchID   respjson.Field
		Type      respjson.Field
		MonitorID respjson.Field
		RunID     respjson.Field
		raw       string
	} `json:"-"`
}

func (u DeliverySourceUnion) AsBatch() (v DeliverySourceBatch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DeliverySourceUnion) AsMonitor() (v DeliverySourceMonitor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DeliverySourceUnion) RawJSON() string { return u.JSON.raw }

func (r *DeliverySourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeliverySourceBatch struct {
	// Batch ID.
	BatchID string `json:"batch_id" api:"required"`
	// Delivery source.
	//
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
func (r DeliverySourceBatch) RawJSON() string { return r.JSON.raw }
func (r *DeliverySourceBatch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeliverySourceMonitor struct {
	// Monitor ID.
	MonitorID string `json:"monitor_id" api:"required"`
	// Monitor run ID.
	RunID string `json:"run_id" api:"required"`
	// Delivery source.
	//
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
func (r DeliverySourceMonitor) RawJSON() string { return r.JSON.raw }
func (r *DeliverySourceMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current delivery status.
type DeliveryStatus string

const (
	DeliveryStatusPending    DeliveryStatus = "pending"
	DeliveryStatusDelivering DeliveryStatus = "delivering"
	DeliveryStatusRetrying   DeliveryStatus = "retrying"
	DeliveryStatusDelivered  DeliveryStatus = "delivered"
	DeliveryStatusFailed     DeliveryStatus = "failed"
	DeliveryStatusCancelled  DeliveryStatus = "cancelled"
)

type DeliverySummary struct {
	// Delivery ID.
	ID string `json:"id" api:"required"`
	// Event creation time.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Last successful delivery time, or null if never delivered.
	DeliveredAt time.Time `json:"delivered_at" api:"required" format:"date-time"`
	// Webhook event type.
	//
	// Any of "batch.completed", "batch.failed", "batch.cancelled", "change.detected",
	// "run.completed".
	Event DeliverySummaryEvent `json:"event" api:"required"`
	// Latest delivery error, or null if none.
	LastError DeliverySummaryLastError `json:"last_error" api:"required"`
	// Next scheduled attempt, or null if none.
	NextAttemptAt time.Time `json:"next_attempt_at" api:"required" format:"date-time"`
	// Manual retry deadline, seven days after event creation.
	RetryExpiresAt time.Time `json:"retry_expires_at" api:"required" format:"date-time"`
	// Batch or monitor run that produced the event.
	Source DeliverySummarySourceUnion `json:"source" api:"required"`
	// Current delivery status.
	//
	// Any of "pending", "delivering", "retrying", "delivered", "failed", "cancelled".
	Status DeliverySummaryStatus `json:"status" api:"required"`
	// Webhook destination URL.
	URL string `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		DeliveredAt    respjson.Field
		Event          respjson.Field
		LastError      respjson.Field
		NextAttemptAt  respjson.Field
		RetryExpiresAt respjson.Field
		Source         respjson.Field
		Status         respjson.Field
		URL            respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeliverySummary) RawJSON() string { return r.JSON.raw }
func (r *DeliverySummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Webhook event type.
type DeliverySummaryEvent string

const (
	DeliverySummaryEventBatchCompleted DeliverySummaryEvent = "batch.completed"
	DeliverySummaryEventBatchFailed    DeliverySummaryEvent = "batch.failed"
	DeliverySummaryEventBatchCancelled DeliverySummaryEvent = "batch.cancelled"
	DeliverySummaryEventChangeDetected DeliverySummaryEvent = "change.detected"
	DeliverySummaryEventRunCompleted   DeliverySummaryEvent = "run.completed"
)

// Latest delivery error, or null if none.
type DeliverySummaryLastError struct {
	// Error code.
	Code string `json:"code" api:"required"`
	// Error details.
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
func (r DeliverySummaryLastError) RawJSON() string { return r.JSON.raw }
func (r *DeliverySummaryLastError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DeliverySummarySourceUnion contains all possible properties and values from
// [DeliverySummarySourceBatch], [DeliverySummarySourceMonitor].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type DeliverySummarySourceUnion struct {
	// This field is from variant [DeliverySummarySourceBatch].
	BatchID string `json:"batch_id"`
	Type    string `json:"type"`
	// This field is from variant [DeliverySummarySourceMonitor].
	MonitorID string `json:"monitor_id"`
	// This field is from variant [DeliverySummarySourceMonitor].
	RunID string `json:"run_id"`
	JSON  struct {
		BatchID   respjson.Field
		Type      respjson.Field
		MonitorID respjson.Field
		RunID     respjson.Field
		raw       string
	} `json:"-"`
}

func (u DeliverySummarySourceUnion) AsBatch() (v DeliverySummarySourceBatch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DeliverySummarySourceUnion) AsMonitor() (v DeliverySummarySourceMonitor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DeliverySummarySourceUnion) RawJSON() string { return u.JSON.raw }

func (r *DeliverySummarySourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeliverySummarySourceBatch struct {
	// Batch ID.
	BatchID string `json:"batch_id" api:"required"`
	// Delivery source.
	//
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
func (r DeliverySummarySourceBatch) RawJSON() string { return r.JSON.raw }
func (r *DeliverySummarySourceBatch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeliverySummarySourceMonitor struct {
	// Monitor ID.
	MonitorID string `json:"monitor_id" api:"required"`
	// Monitor run ID.
	RunID string `json:"run_id" api:"required"`
	// Delivery source.
	//
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
func (r DeliverySummarySourceMonitor) RawJSON() string { return r.JSON.raw }
func (r *DeliverySummarySourceMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current delivery status.
type DeliverySummaryStatus string

const (
	DeliverySummaryStatusPending    DeliverySummaryStatus = "pending"
	DeliverySummaryStatusDelivering DeliverySummaryStatus = "delivering"
	DeliverySummaryStatusRetrying   DeliverySummaryStatus = "retrying"
	DeliverySummaryStatusDelivered  DeliverySummaryStatus = "delivered"
	DeliverySummaryStatusFailed     DeliverySummaryStatus = "failed"
	DeliverySummaryStatusCancelled  DeliverySummaryStatus = "cancelled"
)

type WebhookDeliveryGetResponse struct {
	// Unique id of this API call, also sent in the X-Request-Id response header. Quote
	// it when contacting support about a failed request.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Credit usage, included whenever a valid API key is provided.
	KeyMetadata WebhookDeliveryGetResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestID   respjson.Field
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

// Credit usage, included whenever a valid API key is provided.
type WebhookDeliveryGetResponseKeyMetadata struct {
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
func (r WebhookDeliveryGetResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryGetResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeliveryListResponse struct {
	// Webhook deliveries.
	Data []DeliverySummary `json:"data" api:"required"`
	// Whether more deliveries are available.
	HasMore bool `json:"has_more" api:"required"`
	// Next page cursor, or null on the last page.
	NextCursor string `json:"next_cursor" api:"required"`
	// Unique id of this API call, also sent in the X-Request-Id response header. Quote
	// it when contacting support about a failed request.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Credit usage, included whenever a valid API key is provided.
	KeyMetadata WebhookDeliveryListResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		NextCursor  respjson.Field
		RequestID   respjson.Field
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

// Credit usage, included whenever a valid API key is provided.
type WebhookDeliveryListResponseKeyMetadata struct {
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
func (r WebhookDeliveryListResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryListResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeliveryListAttemptsResponse struct {
	// Delivery attempts.
	Data []Attempt `json:"data" api:"required"`
	// Whether more attempts are available.
	HasMore bool `json:"has_more" api:"required"`
	// Next page cursor, or null on the last page.
	NextCursor string `json:"next_cursor" api:"required"`
	// Unique id of this API call, also sent in the X-Request-Id response header. Quote
	// it when contacting support about a failed request.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Credit usage, included whenever a valid API key is provided.
	KeyMetadata WebhookDeliveryListAttemptsResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		NextCursor  respjson.Field
		RequestID   respjson.Field
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

// Credit usage, included whenever a valid API key is provided.
type WebhookDeliveryListAttemptsResponseKeyMetadata struct {
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
func (r WebhookDeliveryListAttemptsResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryListAttemptsResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeliveryRetryResponse struct {
	// Delivery ID.
	ID string `json:"id" api:"required"`
	// Unique id of this API call, also sent in the X-Request-Id response header. Quote
	// it when contacting support about a failed request.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Credit usage, included whenever a valid API key is provided.
	KeyMetadata WebhookDeliveryRetryResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		RequestID   respjson.Field
		KeyMetadata respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeliveryRetryResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryRetryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credit usage, included whenever a valid API key is provided.
type WebhookDeliveryRetryResponseKeyMetadata struct {
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
func (r WebhookDeliveryRetryResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryRetryResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeliveryGetParams struct {
	// Comma-separated tags for tracking request usage. Up to 20 tags, each 1-50
	// characters.
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

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfBatch *WebhookDeliveryListParamsBodyBatch `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfMonitor *WebhookDeliveryListParamsBodyMonitor `json:",inline"`

	paramObj
}

func (u WebhookDeliveryListParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBatch, u.OfMonitor)
}
func (r *WebhookDeliveryListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type WebhookDeliveryListParamsBodyBatch struct {
	// Filter by batch ID.
	BatchID param.Opt[string] `json:"batch_id,omitzero"`
	// Only include events created after this ISO 8601 timestamp.
	CreatedAfter param.Opt[time.Time] `json:"created_after,omitzero" format:"date-time"`
	// The next_cursor from the previous response.
	Cursor param.Opt[string] `json:"cursor,omitzero"`
	// Number of deliveries to return.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// Filter by delivery status.
	//
	// Any of "pending", "delivering", "retrying", "delivered", "failed", "cancelled".
	Status string `json:"status,omitzero"`
	// Optional tags for tracking usage. Up to 20 tags, each 1 to 50 characters.
	Tags []string `json:"tags,omitzero"`
	// Delivery source.
	//
	// This field can be elided, and will marshal its zero value as "batch".
	Type constant.Batch `json:"type" default:"batch"`
	paramObj
}

func (r WebhookDeliveryListParamsBodyBatch) MarshalJSON() (data []byte, err error) {
	type shadow WebhookDeliveryListParamsBodyBatch
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookDeliveryListParamsBodyBatch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebhookDeliveryListParamsBodyBatch](
		"status", "pending", "delivering", "retrying", "delivered", "failed", "cancelled",
	)
}

// The property Type is required.
type WebhookDeliveryListParamsBodyMonitor struct {
	// Only include events created after this ISO 8601 timestamp.
	CreatedAfter param.Opt[time.Time] `json:"created_after,omitzero" format:"date-time"`
	// The next_cursor from the previous response.
	Cursor param.Opt[string] `json:"cursor,omitzero"`
	// Number of deliveries to return.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// Filter by monitor ID.
	MonitorID param.Opt[string] `json:"monitor_id,omitzero"`
	// Filter by monitor run ID.
	RunID param.Opt[string] `json:"run_id,omitzero"`
	// Filter by delivery status.
	//
	// Any of "pending", "delivering", "retrying", "delivered", "failed", "cancelled".
	Status string `json:"status,omitzero"`
	// Optional tags for tracking usage. Up to 20 tags, each 1 to 50 characters.
	Tags []string `json:"tags,omitzero"`
	// Delivery source.
	//
	// This field can be elided, and will marshal its zero value as "monitor".
	Type constant.Monitor `json:"type" default:"monitor"`
	paramObj
}

func (r WebhookDeliveryListParamsBodyMonitor) MarshalJSON() (data []byte, err error) {
	type shadow WebhookDeliveryListParamsBodyMonitor
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookDeliveryListParamsBodyMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebhookDeliveryListParamsBodyMonitor](
		"status", "pending", "delivering", "retrying", "delivered", "failed", "cancelled",
	)
}

type WebhookDeliveryListAttemptsParams struct {
	// The next_cursor from the previous response.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Number of attempts to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Comma-separated tags for tracking request usage. Up to 20 tags, each 1-50
	// characters.
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
	// Resend a delivery that already succeeded.
	Force param.Opt[bool] `json:"force,omitzero"`
	// Unique key to prevent duplicate retry requests.
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
