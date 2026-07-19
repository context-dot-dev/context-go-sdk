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

// Signal that you may fetch brand data soon to improve latency. The type field
// selects what to prefetch (currently only 'brand') and identifier carries exactly
// one lookup key: a domain, or an email whose domain is extracted and validated
// (free email providers and disposable email addresses are not allowed).
func (r *UtilityService) Prefetch(ctx context.Context, body UtilityPrefetchParams, opts ...option.RequestOption) (res *UtilityPrefetchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "utility/prefetch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type UtilityPrefetchResponse struct {
	// The domain that was queued for prefetching
	Domain string `json:"domain"`
	// Metadata about the API key used for the request. Included in every response
	// whenever a valid API key is provided, even when the response status is not 200.
	KeyMetadata UtilityPrefetchResponseKeyMetadata `json:"key_metadata"`
	// Success message
	Message string `json:"message"`
	// Status of the response, e.g., 'ok'
	Status string `json:"status"`
	// The type of prefetch that was queued, echoed from the request (currently always
	// 'brand')
	//
	// Any of "brand".
	Type UtilityPrefetchResponseType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
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

// Metadata about the API key used for the request. Included in every response
// whenever a valid API key is provided, even when the response status is not 200.
type UtilityPrefetchResponseKeyMetadata struct {
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
func (r UtilityPrefetchResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *UtilityPrefetchResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of prefetch that was queued, echoed from the request (currently always
// 'brand')
type UtilityPrefetchResponseType string

const (
	UtilityPrefetchResponseTypeBrand UtilityPrefetchResponseType = "brand"
)

type UtilityPrefetchParams struct {
	// Identifier of the brand to prefetch. Provide exactly one of domain or email.
	Identifier UtilityPrefetchParamsIdentifierUnion `json:"identifier,omitzero" api:"required"`
	// What to prefetch. Currently only 'brand' is supported.
	//
	// Any of "brand".
	Type UtilityPrefetchParamsType `json:"type,omitzero" api:"required"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `json:"timeoutMS,omitzero"`
	// Optional tags for tracking usage. Up to 20 tags, each 1 to 50 characters.
	Tags []string `json:"tags,omitzero"`
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

// Prefetch brand data by domain.
//
// The property Domain is required.
type UtilityPrefetchParamsIdentifierByDomain struct {
	// Domain name to prefetch brand data for
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

// Prefetch brand data by email. The domain will be extracted and validated.
//
// The property Email is required.
type UtilityPrefetchParamsIdentifierByEmail struct {
	// Email address to prefetch brand data for. The domain will be extracted from the
	// email. Free email providers (gmail.com, yahoo.com, etc.) and disposable email
	// addresses are not allowed.
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

// What to prefetch. Currently only 'brand' is supported.
type UtilityPrefetchParamsType string

const (
	UtilityPrefetchParamsTypeBrand UtilityPrefetchParamsType = "brand"
)
