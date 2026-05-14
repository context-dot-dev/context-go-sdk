// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contextdev

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/context.dev-go/internal/apijson"
	"github.com/stainless-sdks/context.dev-go/internal/requestconfig"
	"github.com/stainless-sdks/context.dev-go/option"
	"github.com/stainless-sdks/context.dev-go/packages/param"
	"github.com/stainless-sdks/context.dev-go/packages/respjson"
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

// Signal that you may fetch brand data for a particular domain soon to improve
// latency.
func (r *UtilityService) Prefetch(ctx context.Context, body UtilityPrefetchParams, opts ...option.RequestOption) (res *UtilityPrefetchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "brand/prefetch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Signal that you may fetch brand data for a particular domain soon to improve
// latency. This endpoint accepts an email address, extracts the domain from it,
// validates that it's not a disposable or free email provider, and queues the
// domain for prefetching.
func (r *UtilityService) PrefetchByEmail(ctx context.Context, body UtilityPrefetchByEmailParams, opts ...option.RequestOption) (res *UtilityPrefetchByEmailResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "brand/prefetch-by-email"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type UtilityPrefetchResponse struct {
	// The domain that was queued for prefetching
	Domain string `json:"domain"`
	// Success message
	Message string `json:"message"`
	// Status of the response, e.g., 'ok'
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		Message     respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UtilityPrefetchResponse) RawJSON() string { return r.JSON.raw }
func (r *UtilityPrefetchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UtilityPrefetchByEmailResponse struct {
	// The domain that was queued for prefetching
	Domain string `json:"domain"`
	// Success message
	Message string `json:"message"`
	// Status of the response, e.g., 'ok'
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		Message     respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UtilityPrefetchByEmailResponse) RawJSON() string { return r.JSON.raw }
func (r *UtilityPrefetchByEmailResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UtilityPrefetchParams struct {
	// Domain name to prefetch brand data for
	Domain string `json:"domain" api:"required"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `json:"timeoutMS,omitzero"`
	paramObj
}

func (r UtilityPrefetchParams) MarshalJSON() (data []byte, err error) {
	type shadow UtilityPrefetchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UtilityPrefetchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UtilityPrefetchByEmailParams struct {
	// Email address to prefetch brand data for. The domain will be extracted from the
	// email. Free email providers (gmail.com, yahoo.com, etc.) and disposable email
	// addresses are not allowed.
	Email string `json:"email" api:"required" format:"email"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `json:"timeoutMS,omitzero"`
	paramObj
}

func (r UtilityPrefetchByEmailParams) MarshalJSON() (data []byte, err error) {
	type shadow UtilityPrefetchByEmailParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UtilityPrefetchByEmailParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
