// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contextdev

import (
	"context"
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

// See the authenticated browser profiles your organization created.
//
// BrowserProfileService contains methods and other services that help with
// interacting with the context.dev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBrowserProfileService] method instead.
type BrowserProfileService struct {
	options []option.RequestOption
}

// NewBrowserProfileService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBrowserProfileService(opts ...option.RequestOption) (r BrowserProfileService) {
	r = BrowserProfileService{}
	r.options = opts
	return
}

// List the browser profiles your organization created with `/web/scrape`, newest
// first.
func (r *BrowserProfileService) List(ctx context.Context, query BrowserProfileListParams, opts ...option.RequestOption) (res *BrowserProfileListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "org/browser/profiles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type BrowserProfileListResponse struct {
	// Browser profiles, newest first.
	Data []BrowserProfileListResponseData `json:"data" api:"required"`
	// Whether a next page exists.
	HasMore bool `json:"has_more" api:"required"`
	// Pass as `cursor` to fetch the next page.
	NextCursor string `json:"next_cursor" api:"required"`
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Credits this request used and your remaining balance.
	KeyMetadata BrowserProfileListResponseKeyMetadata `json:"key_metadata"`
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
func (r BrowserProfileListResponse) RawJSON() string { return r.JSON.raw }
func (r *BrowserProfileListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrowserProfileListResponseData struct {
	// Profile ID to pass as `sharedParams.profile.profileId` on `/web/scrape`.
	ID string `json:"id" api:"required"`
	// When the profile was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The description supplied when the profile was created, or null.
	Description string `json:"description" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrowserProfileListResponseData) RawJSON() string { return r.JSON.raw }
func (r *BrowserProfileListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credits this request used and your remaining balance.
type BrowserProfileListResponseKeyMetadata struct {
	// Credits charged for this request.
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
func (r BrowserProfileListResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *BrowserProfileListResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrowserProfileListParams struct {
	// The `next_cursor` from the previous page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Number of profiles per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BrowserProfileListParams]'s query parameters as
// `url.Values`.
func (r BrowserProfileListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
