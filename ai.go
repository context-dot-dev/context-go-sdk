// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contextdev

import (
	"context"
	"net/http"
	"slices"

	"github.com/context-dot-dev/context-go-sdk/internal/apijson"
	"github.com/context-dot-dev/context-go-sdk/internal/requestconfig"
	"github.com/context-dot-dev/context-go-sdk/option"
	"github.com/context-dot-dev/context-go-sdk/packages/param"
	"github.com/context-dot-dev/context-go-sdk/packages/respjson"
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

type AIExtractProductResponse struct {
	// Whether the given URL is a product detail page
	IsProductPage bool `json:"is_product_page"`
	// Metadata about the API key used for the request. Included in every response
	// whenever a valid API key is provided, even when the response status is not 200.
	KeyMetadata AIExtractProductResponseKeyMetadata `json:"key_metadata"`
	// The detected ecommerce platform, or null if not a product page
	//
	// Any of "amazon", "tiktok_shop", "etsy", "generic".
	Platform AIExtractProductResponsePlatform `json:"platform" api:"nullable"`
	// The extracted product data, or null if not a product page
	Product AIExtractProductResponseProduct `json:"product" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsProductPage respjson.Field
		KeyMetadata   respjson.Field
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

// Metadata about the API key used for the request. Included in every response
// whenever a valid API key is provided, even when the response status is not 200.
type AIExtractProductResponseKeyMetadata struct {
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
func (r AIExtractProductResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *AIExtractProductResponseKeyMetadata) UnmarshalJSON(data []byte) error {
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
	// Metadata about the API key used for the request. Included in every response
	// whenever a valid API key is provided, even when the response status is not 200.
	KeyMetadata AIExtractProductsResponseKeyMetadata `json:"key_metadata"`
	// Array of products extracted from the website
	Products []AIExtractProductsResponseProduct `json:"products"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		KeyMetadata respjson.Field
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

// Metadata about the API key used for the request. Included in every response
// whenever a valid API key is provided, even when the response status is not 200.
type AIExtractProductsResponseKeyMetadata struct {
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
func (r AIExtractProductsResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *AIExtractProductsResponseKeyMetadata) UnmarshalJSON(data []byte) error {
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
