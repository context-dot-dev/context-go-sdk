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

// Signal that you may fetch data soon to improve latency. The type field selects
// what to prefetch ('brand' queues a brand data fetch, 'styleguide' queues a
// styleguide extraction) and identifier carries exactly one lookup key: a domain,
// or an email whose domain is extracted and validated (free email providers and
// disposable email addresses are not allowed).
func (r *UtilityService) Prefetch(ctx context.Context, body UtilityPrefetchParams, opts ...option.RequestOption) (res *UtilityPrefetchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "utility/prefetch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type UtilityPrefetchResponse struct {
	// Unique id of this API call, also sent in the X-Request-Id response header. Quote
	// it when contacting support about a failed request.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// The domain that was queued for prefetching
	Domain string `json:"domain"`
	// Credit usage, included whenever a valid API key is provided.
	KeyMetadata UtilityPrefetchResponseKeyMetadata `json:"key_metadata"`
	// Success message
	Message string `json:"message"`
	// Status of the response, e.g., 'ok'
	Status string `json:"status"`
	// The type of prefetch that was queued, echoed from the request
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

// Credit usage, included whenever a valid API key is provided.
type UtilityPrefetchResponseKeyMetadata struct {
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
func (r UtilityPrefetchResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *UtilityPrefetchResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of prefetch that was queued, echoed from the request
type UtilityPrefetchResponseType string

const (
	UtilityPrefetchResponseTypeBrand      UtilityPrefetchResponseType = "brand"
	UtilityPrefetchResponseTypeStyleguide UtilityPrefetchResponseType = "styleguide"
)

type UtilityPrefetchParams struct {
	// Identifier of the target to prefetch. Provide exactly one of domain or email.
	Identifier UtilityPrefetchParamsIdentifierUnion `json:"identifier,omitzero" api:"required"`
	// What to prefetch: 'brand' warms the brand data cache, 'styleguide' warms the
	// styleguide cache.
	//
	// Any of "brand", "styleguide".
	Type UtilityPrefetchParamsType `json:"type,omitzero" api:"required"`
	// Optional tags for tracking usage. Up to 20 tags, each 1 to 50 characters.
	Tags []string `json:"tags,omitzero"`
	// Optional request deadline and behavior on timeout. For GET requests, use
	// timeoutOpts[milliseconds]=30000&timeoutOpts[behavior]=fail or a JSON-encoded
	// timeoutOpts object.
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
	// Domain name to prefetch data for
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

// Prefetch by email. The domain will be extracted and validated.
//
// The property Email is required.
type UtilityPrefetchParamsIdentifierByEmail struct {
	// Email address to prefetch data for. The domain will be extracted from the email.
	// Free email providers (gmail.com, yahoo.com, etc.) and disposable email addresses
	// are not allowed.
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

// What to prefetch: 'brand' warms the brand data cache, 'styleguide' warms the
// styleguide cache.
type UtilityPrefetchParamsType string

const (
	UtilityPrefetchParamsTypeBrand      UtilityPrefetchParamsType = "brand"
	UtilityPrefetchParamsTypeStyleguide UtilityPrefetchParamsType = "styleguide"
)

// Optional request deadline and behavior on timeout. For GET requests, use
// timeoutOpts[milliseconds]=30000&timeoutOpts[behavior]=fail or a JSON-encoded
// timeoutOpts object.
//
// The property Milliseconds is required.
type UtilityPrefetchParamsTimeoutOpts struct {
	// Request deadline in milliseconds. Maximum: 300000 (5 minutes).
	Milliseconds int64 `json:"milliseconds" api:"required"`
	// What to do at the deadline. This endpoint supports "fail": return 408
	// REQUEST_TIMEOUT without charging credits.
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
