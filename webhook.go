// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contextdev

import (
	"encoding/json"

	"github.com/context-dot-dev/context-go-sdk/v2/internal/apijson"
	"github.com/context-dot-dev/context-go-sdk/v2/option"
	"github.com/context-dot-dev/context-go-sdk/v2/packages/param"
	"github.com/context-dot-dev/context-go-sdk/v2/packages/respjson"
)

// WebhookService contains methods and other services that help with interacting
// with the context.dev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	options []option.RequestOption
	// Inspect and retry batch and monitor webhook deliveries without rerunning the
	// underlying work.
	Deliveries WebhookDeliveryService
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r WebhookService) {
	r = WebhookService{}
	r.options = opts
	r.Deliveries = NewWebhookDeliveryService(opts...)
	return
}

// Opt into durable webhook delivery. An empty object uses the default retry
// schedule. Omit retry to preserve legacy delivery behavior. The policy is
// snapshotted for each event.
type RetryConfig struct {
	// Wait in seconds after each failed attempt. The first attempt is immediate. At
	// most 10 delays, each 1–86400 seconds, totaling at most 72 hours. Small jitter is
	// added automatically. An empty array disables automatic retries; manual retries
	// remain available.
	DelaysSeconds []int64 `json:"delays_seconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DelaysSeconds respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RetryConfig) RawJSON() string { return r.JSON.raw }
func (r *RetryConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RetryConfig to a RetryConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RetryConfigParam.Overrides()
func (r RetryConfig) ToParam() RetryConfigParam {
	return param.Override[RetryConfigParam](json.RawMessage(r.RawJSON()))
}

// Opt into durable webhook delivery. An empty object uses the default retry
// schedule. Omit retry to preserve legacy delivery behavior. The policy is
// snapshotted for each event.
type RetryConfigParam struct {
	// Wait in seconds after each failed attempt. The first attempt is immediate. At
	// most 10 delays, each 1–86400 seconds, totaling at most 72 hours. Small jitter is
	// added automatically. An empty array disables automatic retries; manual retries
	// remain available.
	DelaysSeconds []int64 `json:"delays_seconds,omitzero"`
	paramObj
}

func (r RetryConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow RetryConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RetryConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
