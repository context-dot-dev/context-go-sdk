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
	// Cache outcome for this response. Composite responses are hits only when every
	// cache-controlled fetch contributing to the output was a hit; age_ms is the
	// oldest contributing hit.
	CacheMetadata AIExtractProductResponseCacheMetadata `json:"cache_metadata" api:"required"`
	// Unique id of this API call, also sent in the X-Request-Id response header. Quote
	// it when contacting support about a failed request.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Whether the given URL is a product detail page
	IsProductPage bool `json:"is_product_page"`
	// Credit usage, included whenever a valid API key is provided.
	KeyMetadata AIExtractProductResponseKeyMetadata `json:"key_metadata"`
	// True when the timeout ended processing and this response contains only usable
	// results completed so far. Unfinished results are omitted.
	Partial bool `json:"partial"`
	// The detected ecommerce platform, or null if not a product page
	//
	// Any of "amazon", "tiktok_shop", "etsy", "generic".
	Platform AIExtractProductResponsePlatform `json:"platform" api:"nullable"`
	// The extracted product data, or null if not a product page
	Product AIExtractProductResponseProduct `json:"product" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CacheMetadata respjson.Field
		RequestID     respjson.Field
		IsProductPage respjson.Field
		KeyMetadata   respjson.Field
		Partial       respjson.Field
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

// Cache outcome for this response. Composite responses are hits only when every
// cache-controlled fetch contributing to the output was a hit; age_ms is the
// oldest contributing hit.
type AIExtractProductResponseCacheMetadata struct {
	// Age of the cached data in milliseconds. Zero for miss and zdr responses.
	AgeMs int64 `json:"age_ms" api:"required"`
	// Whether the response was served from cache, required fresh work, or honored
	// zero-data-retention cache bypass.
	//
	// Any of "hit", "miss", "zdr".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgeMs       respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIExtractProductResponseCacheMetadata) RawJSON() string { return r.JSON.raw }
func (r *AIExtractProductResponseCacheMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credit usage, included whenever a valid API key is provided.
type AIExtractProductResponseKeyMetadata struct {
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
	// Normalized stock or ordering availability
	//
	// Any of "in_stock", "out_of_stock", "limited_availability", "preorder",
	// "backorder", "made_to_order", "discontinued".
	Availability string `json:"availability" api:"nullable"`
	// Billing frequency for the product
	//
	// Any of "monthly", "yearly", "one_time", "usage_based".
	BillingFrequency string `json:"billing_frequency" api:"nullable"`
	// Category of the product
	Category string `json:"category" api:"nullable"`
	// Currency code for the price (e.g., USD, EUR)
	Currency string `json:"currency" api:"nullable"`
	// Dimension statements shown for the product, preserving labels, values, and units
	Dimensions []string `json:"dimensions"`
	// URL to the product image
	ImageURL string `json:"image_url" api:"nullable"`
	// Price of the product
	Price float64 `json:"price" api:"nullable"`
	// Pricing model for the product
	//
	// Any of "per_seat", "flat", "tiered", "freemium", "custom".
	PricingModel string `json:"pricing_model" api:"nullable"`
	// Original or regular price before a displayed discount
	RegularPrice float64 `json:"regular_price" api:"nullable"`
	// URL to the product page
	URL string `json:"url" api:"nullable"`
	// Product variations, such as different colors or sizes, with their attributes and
	// images. Empty if none are found. May not include every variation offered by the
	// store.
	Variants []AIExtractProductResponseProductVariant `json:"variants"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description      respjson.Field
		Features         respjson.Field
		Images           respjson.Field
		Name             respjson.Field
		SKU              respjson.Field
		Tags             respjson.Field
		TargetAudience   respjson.Field
		Availability     respjson.Field
		BillingFrequency respjson.Field
		Category         respjson.Field
		Currency         respjson.Field
		Dimensions       respjson.Field
		ImageURL         respjson.Field
		Price            respjson.Field
		PricingModel     respjson.Field
		RegularPrice     respjson.Field
		URL              respjson.Field
		Variants         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIExtractProductResponseProduct) RawJSON() string { return r.JSON.raw }
func (r *AIExtractProductResponseProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIExtractProductResponseProductVariant struct {
	// Explicit variant attributes such as color, size, material, pattern and
	// properties declared by page.
	Attributes map[string]string `json:"attributes" api:"required"`
	// Original source image URLs explicitly attached to this variant.
	Images []string `json:"images" api:"required"`
	SKU    string   `json:"sku" api:"required"`
	// Variant or offer URL when provided by the source. May be shared by variants.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attributes  respjson.Field
		Images      respjson.Field
		SKU         respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIExtractProductResponseProductVariant) RawJSON() string { return r.JSON.raw }
func (r *AIExtractProductResponseProductVariant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIExtractProductsResponse struct {
	// Cache outcome for this response. Composite responses are hits only when every
	// cache-controlled fetch contributing to the output was a hit; age_ms is the
	// oldest contributing hit.
	CacheMetadata AIExtractProductsResponseCacheMetadata `json:"cache_metadata" api:"required"`
	// Unique id of this API call, also sent in the X-Request-Id response header. Quote
	// it when contacting support about a failed request.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Credit usage, included whenever a valid API key is provided.
	KeyMetadata AIExtractProductsResponseKeyMetadata `json:"key_metadata"`
	// True when timeoutOpts.behavior=return-partial returned the usable results
	// collected before the deadline. Partial collections are not cached as complete
	// results.
	Partial bool `json:"partial"`
	// Array of products extracted from the website
	Products []AIExtractProductsResponseProduct `json:"products"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CacheMetadata respjson.Field
		RequestID     respjson.Field
		KeyMetadata   respjson.Field
		Partial       respjson.Field
		Products      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIExtractProductsResponse) RawJSON() string { return r.JSON.raw }
func (r *AIExtractProductsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache outcome for this response. Composite responses are hits only when every
// cache-controlled fetch contributing to the output was a hit; age_ms is the
// oldest contributing hit.
type AIExtractProductsResponseCacheMetadata struct {
	// Age of the cached data in milliseconds. Zero for miss and zdr responses.
	AgeMs int64 `json:"age_ms" api:"required"`
	// Whether the response was served from cache, required fresh work, or honored
	// zero-data-retention cache bypass.
	//
	// Any of "hit", "miss", "zdr".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgeMs       respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIExtractProductsResponseCacheMetadata) RawJSON() string { return r.JSON.raw }
func (r *AIExtractProductsResponseCacheMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credit usage, included whenever a valid API key is provided.
type AIExtractProductsResponseKeyMetadata struct {
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
	// Normalized stock or ordering availability
	//
	// Any of "in_stock", "out_of_stock", "limited_availability", "preorder",
	// "backorder", "made_to_order", "discontinued".
	Availability string `json:"availability" api:"nullable"`
	// Billing frequency for the product
	//
	// Any of "monthly", "yearly", "one_time", "usage_based".
	BillingFrequency string `json:"billing_frequency" api:"nullable"`
	// Category of the product
	Category string `json:"category" api:"nullable"`
	// Currency code for the price (e.g., USD, EUR)
	Currency string `json:"currency" api:"nullable"`
	// Dimension statements shown for the product, preserving labels, values, and units
	Dimensions []string `json:"dimensions"`
	// URL to the product image
	ImageURL string `json:"image_url" api:"nullable"`
	// Price of the product
	Price float64 `json:"price" api:"nullable"`
	// Pricing model for the product
	//
	// Any of "per_seat", "flat", "tiered", "freemium", "custom".
	PricingModel string `json:"pricing_model" api:"nullable"`
	// Original or regular price before a displayed discount
	RegularPrice float64 `json:"regular_price" api:"nullable"`
	// URL to the product page
	URL string `json:"url" api:"nullable"`
	// Product variations, such as different colors or sizes, with their attributes and
	// images. Empty if none are found. May not include every variation offered by the
	// store.
	Variants []AIExtractProductsResponseProductVariant `json:"variants"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description      respjson.Field
		Features         respjson.Field
		Images           respjson.Field
		Name             respjson.Field
		SKU              respjson.Field
		Tags             respjson.Field
		TargetAudience   respjson.Field
		Availability     respjson.Field
		BillingFrequency respjson.Field
		Category         respjson.Field
		Currency         respjson.Field
		Dimensions       respjson.Field
		ImageURL         respjson.Field
		Price            respjson.Field
		PricingModel     respjson.Field
		RegularPrice     respjson.Field
		URL              respjson.Field
		Variants         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIExtractProductsResponseProduct) RawJSON() string { return r.JSON.raw }
func (r *AIExtractProductsResponseProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIExtractProductsResponseProductVariant struct {
	// Explicit variant attributes such as color, size, material, pattern and
	// properties declared by page.
	Attributes map[string]string `json:"attributes" api:"required"`
	// Original source image URLs explicitly attached to this variant.
	Images []string `json:"images" api:"required"`
	SKU    string   `json:"sku" api:"required"`
	// Variant or offer URL when provided by the source. May be shared by variants.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attributes  respjson.Field
		Images      respjson.Field
		SKU         respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIExtractProductsResponseProductVariant) RawJSON() string { return r.JSON.raw }
func (r *AIExtractProductsResponseProductVariant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIExtractProductParams struct {
	// The product page URL to extract product data from.
	URL string `json:"url" api:"required" format:"uri"`
	// Return a cached result if a prior scrape for the same parameters exists and is
	// younger than this many milliseconds. Defaults to 7 days (604800000 ms) when
	// omitted. Max is 30 days (2592000000 ms). Set to 0 to always scrape fresh.
	MaxAgeMs param.Opt[int64] `json:"maxAgeMs,omitzero"`
	// Optional tags for tracking usage. Up to 20 tags, each 1 to 50 characters.
	Tags []string `json:"tags,omitzero"`
	// Optional request deadline and behavior on timeout. For GET requests, use
	// timeoutOpts[milliseconds]=30000&timeoutOpts[behavior]=fail or a JSON-encoded
	// timeoutOpts object.
	TimeoutOpts AIExtractProductParamsTimeoutOpts `json:"timeoutOpts,omitzero"`
	// Set to enabled to bypass shared caches and omit request and response content
	// from retained usage logs. Asset uploads are skipped, so hosted image URLs are
	// omitted. Requires zero data retention to be enabled for your organization
	// (contact support@context.dev), otherwise the request fails with ZDR_NOT_ENABLED.
	// Successful ZDR responses include X-Context-ZDR: true.
	//
	// Any of "enabled", "disabled".
	Zdr AIExtractProductParamsZdr `json:"zdr,omitzero"`
	paramObj
}

func (r AIExtractProductParams) MarshalJSON() (data []byte, err error) {
	type shadow AIExtractProductParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIExtractProductParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional request deadline and behavior on timeout. For GET requests, use
// timeoutOpts[milliseconds]=30000&timeoutOpts[behavior]=fail or a JSON-encoded
// timeoutOpts object.
//
// The property Milliseconds is required.
type AIExtractProductParamsTimeoutOpts struct {
	// Request deadline in milliseconds. Maximum: 300000 (5 minutes).
	Milliseconds int64 `json:"milliseconds" api:"required"`
	// What to do at the deadline. "fail" returns 408 REQUEST_TIMEOUT without charging
	// credits. "return-partial" returns usable results collected so far; if none are
	// available, the request still fails without charging credits. Partial results are
	// not cached as complete results.
	//
	// Any of "fail", "return-partial".
	Behavior string `json:"behavior,omitzero"`
	paramObj
}

func (r AIExtractProductParamsTimeoutOpts) MarshalJSON() (data []byte, err error) {
	type shadow AIExtractProductParamsTimeoutOpts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIExtractProductParamsTimeoutOpts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AIExtractProductParamsTimeoutOpts](
		"behavior", "fail", "return-partial",
	)
}

// Set to enabled to bypass shared caches and omit request and response content
// from retained usage logs. Asset uploads are skipped, so hosted image URLs are
// omitted. Requires zero data retention to be enabled for your organization
// (contact support@context.dev), otherwise the request fails with ZDR_NOT_ENABLED.
// Successful ZDR responses include X-Context-ZDR: true.
type AIExtractProductParamsZdr string

const (
	AIExtractProductParamsZdrEnabled  AIExtractProductParamsZdr = "enabled"
	AIExtractProductParamsZdrDisabled AIExtractProductParamsZdr = "disabled"
)

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
	// Optional tags for tracking usage. Up to 20 tags, each 1 to 50 characters.
	Tags []string `json:"tags,omitzero"`
	// Optional request deadline and behavior on timeout. For GET requests, use
	// timeoutOpts[milliseconds]=30000&timeoutOpts[behavior]=fail or a JSON-encoded
	// timeoutOpts object.
	TimeoutOpts AIExtractProductsParamsBodyByDomainTimeoutOpts `json:"timeoutOpts,omitzero"`
	paramObj
}

func (r AIExtractProductsParamsBodyByDomain) MarshalJSON() (data []byte, err error) {
	type shadow AIExtractProductsParamsBodyByDomain
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIExtractProductsParamsBodyByDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional request deadline and behavior on timeout. For GET requests, use
// timeoutOpts[milliseconds]=30000&timeoutOpts[behavior]=fail or a JSON-encoded
// timeoutOpts object.
//
// The property Milliseconds is required.
type AIExtractProductsParamsBodyByDomainTimeoutOpts struct {
	// Request deadline in milliseconds. Maximum: 300000 (5 minutes).
	Milliseconds int64 `json:"milliseconds" api:"required"`
	// What to do at the deadline. "fail" returns 408 REQUEST_TIMEOUT without charging
	// credits. "return-partial" returns usable results collected so far; if none are
	// available, the request still fails without charging credits. Partial results are
	// not cached as complete results.
	//
	// Any of "fail", "return-partial".
	Behavior string `json:"behavior,omitzero"`
	paramObj
}

func (r AIExtractProductsParamsBodyByDomainTimeoutOpts) MarshalJSON() (data []byte, err error) {
	type shadow AIExtractProductsParamsBodyByDomainTimeoutOpts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIExtractProductsParamsBodyByDomainTimeoutOpts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AIExtractProductsParamsBodyByDomainTimeoutOpts](
		"behavior", "fail", "return-partial",
	)
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
	// Optional tags for tracking usage. Up to 20 tags, each 1 to 50 characters.
	Tags []string `json:"tags,omitzero"`
	// Optional request deadline and behavior on timeout. For GET requests, use
	// timeoutOpts[milliseconds]=30000&timeoutOpts[behavior]=fail or a JSON-encoded
	// timeoutOpts object.
	TimeoutOpts AIExtractProductsParamsBodyByDirectURLTimeoutOpts `json:"timeoutOpts,omitzero"`
	paramObj
}

func (r AIExtractProductsParamsBodyByDirectURL) MarshalJSON() (data []byte, err error) {
	type shadow AIExtractProductsParamsBodyByDirectURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIExtractProductsParamsBodyByDirectURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional request deadline and behavior on timeout. For GET requests, use
// timeoutOpts[milliseconds]=30000&timeoutOpts[behavior]=fail or a JSON-encoded
// timeoutOpts object.
//
// The property Milliseconds is required.
type AIExtractProductsParamsBodyByDirectURLTimeoutOpts struct {
	// Request deadline in milliseconds. Maximum: 300000 (5 minutes).
	Milliseconds int64 `json:"milliseconds" api:"required"`
	// What to do at the deadline. "fail" returns 408 REQUEST_TIMEOUT without charging
	// credits. "return-partial" returns usable results collected so far; if none are
	// available, the request still fails without charging credits. Partial results are
	// not cached as complete results.
	//
	// Any of "fail", "return-partial".
	Behavior string `json:"behavior,omitzero"`
	paramObj
}

func (r AIExtractProductsParamsBodyByDirectURLTimeoutOpts) MarshalJSON() (data []byte, err error) {
	type shadow AIExtractProductsParamsBodyByDirectURLTimeoutOpts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIExtractProductsParamsBodyByDirectURLTimeoutOpts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AIExtractProductsParamsBodyByDirectURLTimeoutOpts](
		"behavior", "fail", "return-partial",
	)
}
