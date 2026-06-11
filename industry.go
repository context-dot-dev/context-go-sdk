// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contextdev

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/context-dot-dev/context-go-sdk/internal/apijson"
	"github.com/context-dot-dev/context-go-sdk/internal/apiquery"
	"github.com/context-dot-dev/context-go-sdk/internal/requestconfig"
	"github.com/context-dot-dev/context-go-sdk/option"
	"github.com/context-dot-dev/context-go-sdk/packages/param"
	"github.com/context-dot-dev/context-go-sdk/packages/respjson"
)

// IndustryService contains methods and other services that help with interacting
// with the context.dev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIndustryService] method instead.
type IndustryService struct {
	options []option.RequestOption
}

// NewIndustryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIndustryService(opts ...option.RequestOption) (r IndustryService) {
	r = IndustryService{}
	r.options = opts
	return
}

// Classify any brand into 2022 NAICS industry codes from its domain or name.
func (r *IndustryService) GetNaics(ctx context.Context, query IndustryGetNaicsParams, opts ...option.RequestOption) (res *IndustryGetNaicsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "web/naics"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Classify any brand into Standard Industrial Classification (SIC) codes from its
// domain or name. Choose between the original SIC system (`original_sic`) or the
// latest SIC list maintained by the SEC (`latest_sec`).
func (r *IndustryService) GetSic(ctx context.Context, query IndustryGetSicParams, opts ...option.RequestOption) (res *IndustryGetSicResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "web/sic"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type IndustryGetNaicsResponse struct {
	// Array of NAICS codes and titles.
	Codes []IndustryGetNaicsResponseCode `json:"codes"`
	// Domain found for the brand
	Domain string `json:"domain"`
	// Metadata about the API key used for the request. Included in every response
	// whenever a valid API key is provided, even when the response status is not 200.
	KeyMetadata IndustryGetNaicsResponseKeyMetadata `json:"key_metadata"`
	// Status of the response, e.g., 'ok'
	Status string `json:"status"`
	// Industry classification type, for naics api it will be `naics`
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Codes       respjson.Field
		Domain      respjson.Field
		KeyMetadata respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IndustryGetNaicsResponse) RawJSON() string { return r.JSON.raw }
func (r *IndustryGetNaicsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IndustryGetNaicsResponseCode struct {
	// NAICS code
	Code string `json:"code" api:"required"`
	// Confidence level for how well this NAICS code matches the company description
	//
	// Any of "high", "medium", "low".
	Confidence string `json:"confidence" api:"required"`
	// NAICS title
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Confidence  respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IndustryGetNaicsResponseCode) RawJSON() string { return r.JSON.raw }
func (r *IndustryGetNaicsResponseCode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the API key used for the request. Included in every response
// whenever a valid API key is provided, even when the response status is not 200.
type IndustryGetNaicsResponseKeyMetadata struct {
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
func (r IndustryGetNaicsResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *IndustryGetNaicsResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IndustryGetSicResponse struct {
	// Echoes back which SIC dataset was used to classify the brand.
	//
	// Any of "original_sic", "latest_sec".
	Classification IndustryGetSicResponseClassification `json:"classification"`
	// Array of SIC codes with confidence scores. Extra fields depend on the requested
	// classification: `original_sic` results include `majorGroup` and
	// `majorGroupName`; `latest_sec` results include `office`.
	Codes []IndustryGetSicResponseCode `json:"codes"`
	// Domain found for the brand
	Domain string `json:"domain"`
	// Metadata about the API key used for the request. Included in every response
	// whenever a valid API key is provided, even when the response status is not 200.
	KeyMetadata IndustryGetSicResponseKeyMetadata `json:"key_metadata"`
	// Status of the response, e.g., 'ok'
	Status string `json:"status"`
	// Industry classification type, for sic api it will be `sic`
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Classification respjson.Field
		Codes          respjson.Field
		Domain         respjson.Field
		KeyMetadata    respjson.Field
		Status         respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IndustryGetSicResponse) RawJSON() string { return r.JSON.raw }
func (r *IndustryGetSicResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Echoes back which SIC dataset was used to classify the brand.
type IndustryGetSicResponseClassification string

const (
	IndustryGetSicResponseClassificationOriginalSic IndustryGetSicResponseClassification = "original_sic"
	IndustryGetSicResponseClassificationLatestSec   IndustryGetSicResponseClassification = "latest_sec"
)

type IndustryGetSicResponseCode struct {
	// SIC code (4-digit).
	Code string `json:"code" api:"required"`
	// Confidence level for how well this SIC code matches the company description.
	//
	// Any of "high", "medium", "low".
	Confidence string `json:"confidence" api:"required"`
	// SIC industry title.
	Name string `json:"name" api:"required"`
	// 2-digit major group identifier (the leading two digits of the code). Only
	// present when `classification` is `original_sic`.
	MajorGroup string `json:"majorGroup"`
	// Description of the 2-digit major group. Only present when `classification` is
	// `original_sic`.
	MajorGroupName string `json:"majorGroupName"`
	// SEC review office responsible for filings under this code. Only present when
	// `classification` is `latest_sec`.
	Office string `json:"office"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code           respjson.Field
		Confidence     respjson.Field
		Name           respjson.Field
		MajorGroup     respjson.Field
		MajorGroupName respjson.Field
		Office         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IndustryGetSicResponseCode) RawJSON() string { return r.JSON.raw }
func (r *IndustryGetSicResponseCode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the API key used for the request. Included in every response
// whenever a valid API key is provided, even when the response status is not 200.
type IndustryGetSicResponseKeyMetadata struct {
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
func (r IndustryGetSicResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *IndustryGetSicResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IndustryGetNaicsParams struct {
	// Brand domain or title to retrieve NAICS code for. If a valid domain is provided,
	// it will be used for classification, otherwise, we will search for the brand
	// using the provided title.
	Input string `query:"input" api:"required" json:"-"`
	// Maximum number of NAICS codes to return. Must be between 1 and 10. Defaults
	// to 5.
	MaxResults param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	// Minimum number of NAICS codes to return. Must be at least 1. Defaults to 1.
	MinResults param.Opt[int64] `query:"minResults,omitzero" json:"-"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `query:"timeoutMS,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IndustryGetNaicsParams]'s query parameters as `url.Values`.
func (r IndustryGetNaicsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IndustryGetSicParams struct {
	// Brand domain or title to retrieve SIC code for. If a valid domain is provided,
	// it will be used for classification, otherwise, we will search for the brand
	// using the provided title.
	Input string `query:"input" api:"required" json:"-"`
	// Maximum number of SIC codes to return. Must be between 1 and 10. Defaults to 5.
	MaxResults param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	// Minimum number of SIC codes to return. Must be at least 1. Defaults to 1.
	MinResults param.Opt[int64] `query:"minResults,omitzero" json:"-"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `query:"timeoutMS,omitzero" json:"-"`
	// Which SIC dataset to classify against. `original_sic` uses the 1987 Standard
	// Industrial Classification system; `latest_sec` uses the current SIC list as
	// published by the SEC. Defaults to `original_sic`.
	//
	// Any of "original_sic", "latest_sec".
	Type IndustryGetSicParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IndustryGetSicParams]'s query parameters as `url.Values`.
func (r IndustryGetSicParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Which SIC dataset to classify against. `original_sic` uses the 1987 Standard
// Industrial Classification system; `latest_sec` uses the current SIC list as
// published by the SEC. Defaults to `original_sic`.
type IndustryGetSicParamsType string

const (
	IndustryGetSicParamsTypeOriginalSic IndustryGetSicParamsType = "original_sic"
	IndustryGetSicParamsTypeLatestSec   IndustryGetSicParamsType = "latest_sec"
)
