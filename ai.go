// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contextdev

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/stainless-sdks/context.dev-go/internal/apijson"
	"github.com/stainless-sdks/context.dev-go/internal/requestconfig"
	"github.com/stainless-sdks/context.dev-go/option"
	"github.com/stainless-sdks/context.dev-go/packages/param"
	"github.com/stainless-sdks/context.dev-go/packages/respjson"
)

// AIService contains methods and other services that help with interacting with
// the context.dev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIService] method instead.
type AIService struct {
	options []option.RequestOption
}

// NewAIService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAIService(opts ...option.RequestOption) (r AIService) {
	r = AIService{}
	r.options = opts
	return
}

// Use AI to extract specific data points from a brand's website. The AI will crawl
// the website and extract the requested information based on the provided data
// points.
func (r *AIService) AIQuery(ctx context.Context, body AIAIQueryParams, opts ...option.RequestOption) (res *AiaiQueryResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "brand/ai/query"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Given a single URL, determines if it is a product page and extracts the product
// information.
func (r *AIService) ExtractProduct(ctx context.Context, body AIExtractProductParams, opts ...option.RequestOption) (res *AIExtractProductResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "brand/ai/product"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Extract product information from a brand's website. We will analyze the website
// and return a list of products with details such as name, description, image,
// pricing, features, and more.
func (r *AIService) ExtractProducts(ctx context.Context, body AIExtractProductsParams, opts ...option.RequestOption) (res *AIExtractProductsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "brand/ai/products"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type AiaiQueryResponse struct {
	// Array of extracted data points
	DataExtracted []AiaiQueryResponseDataExtracted `json:"data_extracted"`
	// The domain that was analyzed
	Domain string `json:"domain"`
	// Status of the response, e.g., 'ok'
	Status string `json:"status"`
	// List of URLs that were analyzed
	URLsAnalyzed []string `json:"urls_analyzed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataExtracted respjson.Field
		Domain        respjson.Field
		Status        respjson.Field
		URLsAnalyzed  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AiaiQueryResponse) RawJSON() string { return r.JSON.raw }
func (r *AiaiQueryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AiaiQueryResponseDataExtracted struct {
	// Name of the extracted data point
	DatapointName string `json:"datapoint_name"`
	// Value of the extracted data point. Can be a primitive type, an array of
	// primitives, or an array of objects when datapoint_list_type is 'object'.
	DatapointValue AiaiQueryResponseDataExtractedDatapointValueUnion `json:"datapoint_value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DatapointName  respjson.Field
		DatapointValue respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AiaiQueryResponseDataExtracted) RawJSON() string { return r.JSON.raw }
func (r *AiaiQueryResponseDataExtracted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AiaiQueryResponseDataExtractedDatapointValueUnion contains all possible
// properties and values from [string], [float64], [bool], [[]string], [[]float64],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray OfFloatArray OfAnyArray]
type AiaiQueryResponseDataExtractedDatapointValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field will be present if the value is a [[]float64] instead of an object.
	OfFloatArray []float64 `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString      respjson.Field
		OfFloat       respjson.Field
		OfBool        respjson.Field
		OfStringArray respjson.Field
		OfFloatArray  respjson.Field
		OfAnyArray    respjson.Field
		raw           string
	} `json:"-"`
}

func (u AiaiQueryResponseDataExtractedDatapointValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AiaiQueryResponseDataExtractedDatapointValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AiaiQueryResponseDataExtractedDatapointValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AiaiQueryResponseDataExtractedDatapointValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AiaiQueryResponseDataExtractedDatapointValueUnion) AsFloatArray() (v []float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AiaiQueryResponseDataExtractedDatapointValueUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AiaiQueryResponseDataExtractedDatapointValueUnion) RawJSON() string { return u.JSON.raw }

func (r *AiaiQueryResponseDataExtractedDatapointValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIExtractProductResponse struct {
	// Whether the given URL is a product detail page
	IsProductPage bool `json:"is_product_page"`
	// The detected ecommerce platform, or null if not a product page
	//
	// Any of "amazon", "tiktok_shop", "etsy", "generic".
	Platform AIExtractProductResponsePlatform `json:"platform" api:"nullable"`
	// The extracted product data, or null if not a product page
	Product AIExtractProductResponseProduct `json:"product" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsProductPage respjson.Field
		Platform      respjson.Field
		Product       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIExtractProductResponse) RawJSON() string { return r.JSON.raw }
func (r *AIExtractProductResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The detected ecommerce platform, or null if not a product page
type AIExtractProductResponsePlatform string

const (
	AIExtractProductResponsePlatformAmazon     AIExtractProductResponsePlatform = "amazon"
	AIExtractProductResponsePlatformTiktokShop AIExtractProductResponsePlatform = "tiktok_shop"
	AIExtractProductResponsePlatformEtsy       AIExtractProductResponsePlatform = "etsy"
	AIExtractProductResponsePlatformGeneric    AIExtractProductResponsePlatform = "generic"
)

// The extracted product data, or null if not a product page
type AIExtractProductResponseProduct struct {
	// Description of the product
	Description string `json:"description" api:"required"`
	// List of product features
	Features []string `json:"features" api:"required"`
	// URLs to product images on the page (up to 7)
	Images []string `json:"images" api:"required"`
	// Name of the product
	Name string `json:"name" api:"required"`
	// Stock Keeping Unit (product identifier). Null if no identifier is found.
	SKU string `json:"sku" api:"required"`
	// Tags associated with the product
	Tags []string `json:"tags" api:"required"`
	// Target audience for the product (array of strings)
	TargetAudience []string `json:"target_audience" api:"required"`
	// Billing frequency for the product
	//
	// Any of "monthly", "yearly", "one_time", "usage_based".
	BillingFrequency string `json:"billing_frequency" api:"nullable"`
	// Category of the product
	Category string `json:"category" api:"nullable"`
	// Currency code for the price (e.g., USD, EUR)
	Currency string `json:"currency" api:"nullable"`
	// URL to the product image
	ImageURL string `json:"image_url" api:"nullable"`
	// Price of the product
	Price float64 `json:"price" api:"nullable"`
	// Pricing model for the product
	//
	// Any of "per_seat", "flat", "tiered", "freemium", "custom".
	PricingModel string `json:"pricing_model" api:"nullable"`
	// URL to the product page
	URL string `json:"url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description      respjson.Field
		Features         respjson.Field
		Images           respjson.Field
		Name             respjson.Field
		SKU              respjson.Field
		Tags             respjson.Field
		TargetAudience   respjson.Field
		BillingFrequency respjson.Field
		Category         respjson.Field
		Currency         respjson.Field
		ImageURL         respjson.Field
		Price            respjson.Field
		PricingModel     respjson.Field
		URL              respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIExtractProductResponseProduct) RawJSON() string { return r.JSON.raw }
func (r *AIExtractProductResponseProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIExtractProductsResponse struct {
	// Array of products extracted from the website
	Products []AIExtractProductsResponseProduct `json:"products"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Products    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIExtractProductsResponse) RawJSON() string { return r.JSON.raw }
func (r *AIExtractProductsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIExtractProductsResponseProduct struct {
	// Description of the product
	Description string `json:"description" api:"required"`
	// List of product features
	Features []string `json:"features" api:"required"`
	// URLs to product images on the page (up to 7)
	Images []string `json:"images" api:"required"`
	// Name of the product
	Name string `json:"name" api:"required"`
	// Stock Keeping Unit (product identifier). Null if no identifier is found.
	SKU string `json:"sku" api:"required"`
	// Tags associated with the product
	Tags []string `json:"tags" api:"required"`
	// Target audience for the product (array of strings)
	TargetAudience []string `json:"target_audience" api:"required"`
	// Billing frequency for the product
	//
	// Any of "monthly", "yearly", "one_time", "usage_based".
	BillingFrequency string `json:"billing_frequency" api:"nullable"`
	// Category of the product
	Category string `json:"category" api:"nullable"`
	// Currency code for the price (e.g., USD, EUR)
	Currency string `json:"currency" api:"nullable"`
	// URL to the product image
	ImageURL string `json:"image_url" api:"nullable"`
	// Price of the product
	Price float64 `json:"price" api:"nullable"`
	// Pricing model for the product
	//
	// Any of "per_seat", "flat", "tiered", "freemium", "custom".
	PricingModel string `json:"pricing_model" api:"nullable"`
	// URL to the product page
	URL string `json:"url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description      respjson.Field
		Features         respjson.Field
		Images           respjson.Field
		Name             respjson.Field
		SKU              respjson.Field
		Tags             respjson.Field
		TargetAudience   respjson.Field
		BillingFrequency respjson.Field
		Category         respjson.Field
		Currency         respjson.Field
		ImageURL         respjson.Field
		Price            respjson.Field
		PricingModel     respjson.Field
		URL              respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIExtractProductsResponseProduct) RawJSON() string { return r.JSON.raw }
func (r *AIExtractProductsResponseProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAIQueryParams struct {
	// Array of data points to extract from the website
	DataToExtract []AIAIQueryParamsDataToExtract `json:"data_to_extract,omitzero" api:"required"`
	// The domain name to analyze
	Domain string `json:"domain" api:"required"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `json:"timeoutMS,omitzero"`
	// Optional object specifying which pages to analyze
	SpecificPages AIAIQueryParamsSpecificPages `json:"specific_pages,omitzero"`
	paramObj
}

func (r AIAIQueryParams) MarshalJSON() (data []byte, err error) {
	type shadow AIAIQueryParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIAIQueryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DatapointDescription, DatapointExample, DatapointName,
// DatapointType are required.
type AIAIQueryParamsDataToExtract struct {
	// Description of what to extract
	DatapointDescription string `json:"datapoint_description" api:"required"`
	// Example of the expected value
	DatapointExample string `json:"datapoint_example" api:"required"`
	// Name of the data point to extract
	DatapointName string `json:"datapoint_name" api:"required"`
	// Type of the data point
	//
	// Any of "text", "number", "date", "boolean", "list", "url".
	DatapointType string `json:"datapoint_type,omitzero" api:"required"`
	// Type of items in the list when datapoint_type is 'list'. Defaults to 'string'.
	// Use 'object' to extract an array of objects matching a schema.
	//
	// Any of "string", "text", "number", "date", "boolean", "list", "url", "object".
	DatapointListType string `json:"datapoint_list_type,omitzero"`
	// Schema definition for objects when datapoint_list_type is 'object'. Provide a
	// map of field names to their scalar types.
	//
	// Any of "string", "number", "date", "boolean".
	DatapointObjectSchema map[string]string `json:"datapoint_object_schema,omitzero"`
	paramObj
}

func (r AIAIQueryParamsDataToExtract) MarshalJSON() (data []byte, err error) {
	type shadow AIAIQueryParamsDataToExtract
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIAIQueryParamsDataToExtract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AIAIQueryParamsDataToExtract](
		"datapoint_type", "text", "number", "date", "boolean", "list", "url",
	)
	apijson.RegisterFieldValidator[AIAIQueryParamsDataToExtract](
		"datapoint_list_type", "string", "text", "number", "date", "boolean", "list", "url", "object",
	)
}

// Optional object specifying which pages to analyze
type AIAIQueryParamsSpecificPages struct {
	// Whether to analyze the about us page
	AboutUs param.Opt[bool] `json:"about_us,omitzero"`
	// Whether to analyze the blog
	Blog param.Opt[bool] `json:"blog,omitzero"`
	// Whether to analyze the careers page
	Careers param.Opt[bool] `json:"careers,omitzero"`
	// Whether to analyze the contact us page
	ContactUs param.Opt[bool] `json:"contact_us,omitzero"`
	// Whether to analyze the FAQ page
	Faq param.Opt[bool] `json:"faq,omitzero"`
	// Whether to analyze the home page
	HomePage param.Opt[bool] `json:"home_page,omitzero"`
	// Whether to analyze the pricing page
	Pricing param.Opt[bool] `json:"pricing,omitzero"`
	// Whether to analyze the privacy policy page
	PrivacyPolicy param.Opt[bool] `json:"privacy_policy,omitzero"`
	// Whether to analyze the terms and conditions page
	TermsAndConditions param.Opt[bool] `json:"terms_and_conditions,omitzero"`
	paramObj
}

func (r AIAIQueryParamsSpecificPages) MarshalJSON() (data []byte, err error) {
	type shadow AIAIQueryParamsSpecificPages
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIAIQueryParamsSpecificPages) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIExtractProductParams struct {
	// The product page URL to extract product data from.
	URL string `json:"url" api:"required" format:"uri"`
	// Return a cached result if a prior scrape for the same parameters exists and is
	// younger than this many milliseconds. Defaults to 7 days (604800000 ms) when
	// omitted. Max is 30 days (2592000000 ms). Set to 0 to always scrape fresh.
	MaxAgeMs param.Opt[int64] `json:"maxAgeMs,omitzero"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `json:"timeoutMS,omitzero"`
	paramObj
}

func (r AIExtractProductParams) MarshalJSON() (data []byte, err error) {
	type shadow AIExtractProductParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIExtractProductParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIExtractProductsParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfByDomain *AIExtractProductsParamsBodyByDomain `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfByDirectURL *AIExtractProductsParamsBodyByDirectURL `json:",inline"`

	paramObj
}

func (u AIExtractProductsParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfByDomain, u.OfByDirectURL)
}
func (r *AIExtractProductsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Domain is required.
type AIExtractProductsParamsBodyByDomain struct {
	// The domain name to analyze.
	Domain string `json:"domain" api:"required"`
	// Return a cached result if a prior scrape for the same parameters exists and is
	// younger than this many milliseconds. Defaults to 7 days (604800000 ms) when
	// omitted. Max is 30 days (2592000000 ms). Set to 0 to always scrape fresh.
	MaxAgeMs param.Opt[int64] `json:"maxAgeMs,omitzero"`
	// Maximum number of products to extract.
	MaxProducts param.Opt[int64] `json:"maxProducts,omitzero"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `json:"timeoutMS,omitzero"`
	paramObj
}

func (r AIExtractProductsParamsBodyByDomain) MarshalJSON() (data []byte, err error) {
	type shadow AIExtractProductsParamsBodyByDomain
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIExtractProductsParamsBodyByDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property DirectURL is required.
type AIExtractProductsParamsBodyByDirectURL struct {
	// A specific URL to use directly as the starting point for extraction without
	// domain resolution.
	DirectURL string `json:"directUrl" api:"required" format:"uri"`
	// Return a cached result if a prior scrape for the same parameters exists and is
	// younger than this many milliseconds. Defaults to 7 days (604800000 ms) when
	// omitted. Max is 30 days (2592000000 ms). Set to 0 to always scrape fresh.
	MaxAgeMs param.Opt[int64] `json:"maxAgeMs,omitzero"`
	// Maximum number of products to extract.
	MaxProducts param.Opt[int64] `json:"maxProducts,omitzero"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `json:"timeoutMS,omitzero"`
	paramObj
}

func (r AIExtractProductsParamsBodyByDirectURL) MarshalJSON() (data []byte, err error) {
	type shadow AIExtractProductsParamsBodyByDirectURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIExtractProductsParamsBodyByDirectURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
