// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contextdev

import (
	"context"
	"net/http"
	"slices"

	"github.com/context-dot-dev/context-go-sdk/v2/internal/apijson"
	"github.com/context-dot-dev/context-go-sdk/v2/internal/requestconfig"
	"github.com/context-dot-dev/context-go-sdk/v2/option"
	"github.com/context-dot-dev/context-go-sdk/v2/packages/param"
	"github.com/context-dot-dev/context-go-sdk/v2/packages/respjson"
)

// Prefetch brand and styleguide data so later lookups return sooner.
//
// UtilityService contains methods and other services that help with interacting
// with the context.dev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUtilityService] method instead.
type UtilityService struct {
	options []option.RequestOption
}

// NewUtilityService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUtilityService(opts ...option.RequestOption) (r UtilityService) {
	r = UtilityService{}
	r.options = opts
	return
}

// Start fetching a domain's brand or styleguide data in the background so a later
// lookup returns faster. Accepts a domain or a work email.
func (r *UtilityService) Prefetch(ctx context.Context, body UtilityPrefetchParams, opts ...option.RequestOption) (res *UtilityPrefetchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "utility/prefetch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type UtilityPrefetchResponse struct {
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Domain queued for prefetching.
	Domain string `json:"domain"`
	// Credits this request used and your remaining balance.
	KeyMetadata UtilityPrefetchResponseKeyMetadata `json:"key_metadata"`
	// Human-readable confirmation.
	Message string `json:"message"`
	// Always `ok` on success.
	Status string `json:"status"`
	// Data type that was queued, from the request.
	//
	// Any of "brand", "styleguide".
	Type UtilityPrefetchResponseType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestID   respjson.Field
		Domain      respjson.Field
		KeyMetadata respjson.Field
		Message     respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UtilityPrefetchResponse) RawJSON() string { return r.JSON.raw }
func (r *UtilityPrefetchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credits this request used and your remaining balance.
type UtilityPrefetchResponseKeyMetadata struct {
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
func (r UtilityPrefetchResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *UtilityPrefetchResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Data type that was queued, from the request.
type UtilityPrefetchResponseType string

const (
	UtilityPrefetchResponseTypeBrand      UtilityPrefetchResponseType = "brand"
	UtilityPrefetchResponseTypeStyleguide UtilityPrefetchResponseType = "styleguide"
)

type UtilityPrefetchParams struct {
	// Domain or work email to prefetch. Provide exactly one.
	Identifier UtilityPrefetchParamsIdentifierUnion `json:"identifier,omitzero" api:"required"`
	// Data to prefetch.
	//
	// Any of "brand", "styleguide".
	Type UtilityPrefetchParamsType `json:"type,omitzero" api:"required"`
	// Access policy IDs to apply; each URL must pass all of them. Requires access
	// policies enabled for your organization; otherwise returns
	// `ACCESS_POLICIES_NOT_ENABLED`.
	Policies []string `json:"policies,omitzero"`
	// Labels for this request. Filter usage and request logs by them in the dashboard
	// or with `GET /org/logs`.
	Tags []string `json:"tags,omitzero"`
	// Request deadline. The request fails with 408 when it passes.
	TimeoutOpts UtilityPrefetchParamsTimeoutOpts `json:"timeoutOpts,omitzero"`
	paramObj
}

func (r UtilityPrefetchParams) MarshalJSON() (data []byte, err error) {
	type shadow UtilityPrefetchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UtilityPrefetchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type UtilityPrefetchParamsIdentifierUnion struct {
	OfByDomain *UtilityPrefetchParamsIdentifierByDomain `json:",omitzero,inline"`
	OfByEmail  *UtilityPrefetchParamsIdentifierByEmail  `json:",omitzero,inline"`
	paramUnion
}

func (u UtilityPrefetchParamsIdentifierUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfByDomain, u.OfByEmail)
}
func (u *UtilityPrefetchParamsIdentifierUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Prefetch by domain.
//
// The property Domain is required.
type UtilityPrefetchParamsIdentifierByDomain struct {
	// Domain to prefetch, such as `stripe.com`.
	Domain string `json:"domain" api:"required"`
	paramObj
}

func (r UtilityPrefetchParamsIdentifierByDomain) MarshalJSON() (data []byte, err error) {
	type shadow UtilityPrefetchParamsIdentifierByDomain
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UtilityPrefetchParamsIdentifierByDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Prefetch by work email. Context.dev uses the email's domain.
//
// The property Email is required.
type UtilityPrefetchParamsIdentifierByEmail struct {
	// Work email whose domain to prefetch. Free email providers such as `gmail.com`
	// and disposable addresses are rejected with 422.
	Email string `json:"email" api:"required" format:"email"`
	paramObj
}

func (r UtilityPrefetchParamsIdentifierByEmail) MarshalJSON() (data []byte, err error) {
	type shadow UtilityPrefetchParamsIdentifierByEmail
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UtilityPrefetchParamsIdentifierByEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Data to prefetch.
type UtilityPrefetchParamsType string

const (
	UtilityPrefetchParamsTypeBrand      UtilityPrefetchParamsType = "brand"
	UtilityPrefetchParamsTypeStyleguide UtilityPrefetchParamsType = "styleguide"
)

// Request deadline. The request fails with 408 when it passes.
//
// The property Milliseconds is required.
type UtilityPrefetchParamsTimeoutOpts struct {
	// Deadline in milliseconds.
	Milliseconds int64 `json:"milliseconds" api:"required"`
	// Only `fail` is supported: the request returns 408 at the deadline.
	//
	// Any of "fail".
	Behavior string `json:"behavior,omitzero"`
	paramObj
}

func (r UtilityPrefetchParamsTimeoutOpts) MarshalJSON() (data []byte, err error) {
	type shadow UtilityPrefetchParamsTimeoutOpts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UtilityPrefetchParamsTimeoutOpts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UtilityPrefetchParamsTimeoutOpts](
		"behavior", "fail",
	)
}
