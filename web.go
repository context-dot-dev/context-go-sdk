// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contextdev

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"slices"

	"github.com/context-dot-dev/context-go-sdk/v2/internal/apijson"
	"github.com/context-dot-dev/context-go-sdk/v2/internal/apiquery"
	"github.com/context-dot-dev/context-go-sdk/v2/internal/requestconfig"
	"github.com/context-dot-dev/context-go-sdk/v2/option"
	"github.com/context-dot-dev/context-go-sdk/v2/packages/param"
	"github.com/context-dot-dev/context-go-sdk/v2/packages/respjson"
	"github.com/context-dot-dev/context-go-sdk/v2/shared/constant"
)

// WebService contains methods and other services that help with interacting with
// the context.dev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebService] method instead.
type WebService struct {
	options []option.RequestOption
}

// NewWebService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebService(opts ...option.RequestOption) (r WebService) {
	r = WebService{}
	r.options = opts
	return
}

// Crawl a website, use the provided JSON Schema and instructions to prioritize
// relevant internal links, and extract structured data from the selected pages.
func (r *WebService) Extract(ctx context.Context, body WebExtractParams, opts ...option.RequestOption) (res *WebExtractResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "web/extract"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Analyze a company's landing page and web search evidence to return direct
// competitors for the same product or market.
func (r *WebService) ExtractCompetitors(ctx context.Context, query WebExtractCompetitorsParams, opts ...option.RequestOption) (res *WebExtractCompetitorsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "web/competitors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Scrape font information from a website including font families, usage
// statistics, fallbacks, and element/word counts.
func (r *WebService) ExtractFonts(ctx context.Context, query WebExtractFontsParams, opts ...option.RequestOption) (res *WebExtractFontsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "web/fonts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Extract a comprehensive design system from a website including colors,
// typography, spacing, shadows, and UI components.
func (r *WebService) ExtractStyleguide(ctx context.Context, query WebExtractStyleguideParams, opts ...option.RequestOption) (res *WebExtractStyleguideResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "web/styleguide"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Capture a screenshot of a website.
func (r *WebService) Screenshot(ctx context.Context, query WebScreenshotParams, opts ...option.RequestOption) (res *WebScreenshotResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "web/screenshot"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Search the web and optionally scrape each result to Markdown in one round-trip.
func (r *WebService) Search(ctx context.Context, body WebSearchParams, opts ...option.RequestOption) (res *WebSearchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "web/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Performs a crawl starting from a given URL, extracts page content as Markdown,
// and returns results for all crawled pages.
func (r *WebService) WebCrawlMd(ctx context.Context, body WebWebCrawlMdParams, opts ...option.RequestOption) (res *WebWebCrawlMdResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "web/crawl"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Scrapes the given URL and returns the raw HTML content of the page. The base
// request costs 1 credit; requests with browser actions cost 2 credits.
func (r *WebService) WebScrapeHTML(ctx context.Context, query WebWebScrapeHTMLParams, opts ...option.RequestOption) (res *WebWebScrapeHTMLResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "web/scrape/html"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Extract image assets from a web page, including standard URLs, inline SVGs, data
// URIs, responsive image sources, metadata, CSS backgrounds, video posters, and
// embeds. The base request costs 1 credit, or 2 credits with browser actions. When
// enrichment is enabled, the entire call costs 5 credits, including requests that
// also use actions.
func (r *WebService) WebScrapeImages(ctx context.Context, query WebWebScrapeImagesParams, opts ...option.RequestOption) (res *WebWebScrapeImagesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "web/scrape/images"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Scrapes the given URL into LLM usable Markdown. Inspect key_metadata on JSON
// responses from a recognized API key; use error_code to distinguish stable
// failure categories.
//
// ### Billing & errors
//
// | HTTP status | Billed?                                   | Meaning                                                                                  |
// | ----------- | ----------------------------------------- | ---------------------------------------------------------------------------------------- |
// | 200         | Yes — 1 credit, or 2 credits with actions | Successful scrape, including a zero-length result when includeSelectors matched nothing  |
// | 400         | No                                        | Invalid input, skipped PDF, or the page could not be scraped                             |
// | 401 / 403   | No                                        | Invalid/disabled key, insufficient permissions, or credits exhausted; inspect error_code |
// | 404         | No                                        | Target page returned or fingerprinted as not found                                       |
// | 408         | No                                        | Request timed out                                                                        |
// | 413         | No                                        | Target content exceeds the maximum supported size (20 MB)                                |
// | 415         | No                                        | Unsupported content type                                                                 |
// | 429         | No                                        | Per-minute rate limit exceeded; honor Retry-After                                        |
// | 500         | No                                        | Internal error                                                                           |
func (r *WebService) WebScrapeMd(ctx context.Context, query WebWebScrapeMdParams, opts ...option.RequestOption) (res *WebWebScrapeMdResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "web/scrape/markdown"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Crawl an entire website's sitemap and return all discovered page URLs.
func (r *WebService) WebScrapeSitemap(ctx context.Context, query WebWebScrapeSitemapParams, opts ...option.RequestOption) (res *WebWebScrapeSitemapResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "web/scrape/sitemap"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type WebExtractResponse struct {
	// Extracted data matching the request schema
	Data     map[string]any             `json:"data" api:"required"`
	Metadata WebExtractResponseMetadata `json:"metadata" api:"required"`
	// Status of the response, e.g., 'ok'
	Status string `json:"status" api:"required"`
	// The starting URL that was analyzed
	URL string `json:"url" api:"required"`
	// List of URLs whose Markdown was used for extraction
	URLsAnalyzed []string `json:"urls_analyzed" api:"required"`
	// Metadata about the API key used for the request. Included in every response
	// whenever a valid API key is provided, even when the response status is not 200.
	KeyMetadata WebExtractResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data         respjson.Field
		Metadata     respjson.Field
		Status       respjson.Field
		URL          respjson.Field
		URLsAnalyzed respjson.Field
		KeyMetadata  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebExtractResponse) RawJSON() string { return r.JSON.raw }
func (r *WebExtractResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebExtractResponseMetadata struct {
	MaxCrawlDepth int64 `json:"maxCrawlDepth" api:"required"`
	// Number of crawled pages excluded because they were anti-bot challenges, error
	// pages, or parked-domain placeholders.
	NumBlocked   int64 `json:"numBlocked" api:"required"`
	NumFailed    int64 `json:"numFailed" api:"required"`
	NumSkipped   int64 `json:"numSkipped" api:"required"`
	NumSucceeded int64 `json:"numSucceeded" api:"required"`
	NumURLs      int64 `json:"numUrls" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxCrawlDepth respjson.Field
		NumBlocked    respjson.Field
		NumFailed     respjson.Field
		NumSkipped    respjson.Field
		NumSucceeded  respjson.Field
		NumURLs       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebExtractResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebExtractResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the API key used for the request. Included in every response
// whenever a valid API key is provided, even when the response status is not 200.
type WebExtractResponseKeyMetadata struct {
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
func (r WebExtractResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebExtractResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebExtractCompetitorsResponse struct {
	// Direct competitors ordered by relevance and confidence.
	Competitors []WebExtractCompetitorsResponseCompetitor `json:"competitors" api:"required"`
	// Normalized input domain.
	Domain string `json:"domain" api:"required"`
	// Status of the response.
	//
	// Any of "ok".
	Status WebExtractCompetitorsResponseStatus `json:"status" api:"required"`
	// Target company profile inferred from the landing page.
	Target WebExtractCompetitorsResponseTarget `json:"target" api:"required"`
	// Metadata about the API key used for the request. Included in every response
	// whenever a valid API key is provided, even when the response status is not 200.
	KeyMetadata WebExtractCompetitorsResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Competitors respjson.Field
		Domain      respjson.Field
		Status      respjson.Field
		Target      respjson.Field
		KeyMetadata respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebExtractCompetitorsResponse) RawJSON() string { return r.JSON.raw }
func (r *WebExtractCompetitorsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebExtractCompetitorsResponseCompetitor struct {
	// Confidence that this company is a direct competitor.
	//
	// Any of "high", "medium".
	Confidence string `json:"confidence" api:"required"`
	// Short description of the competitor.
	Description string `json:"description" api:"required"`
	// Competitor's normalized official domain.
	Domain string `json:"domain" api:"required"`
	// Competitor company or product name.
	Name string `json:"name" api:"required"`
	// Search result URLs used as evidence for this competitor.
	SourceURLs []string `json:"sourceUrls" api:"required"`
	// Competitor website URL.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Confidence  respjson.Field
		Description respjson.Field
		Domain      respjson.Field
		Name        respjson.Field
		SourceURLs  respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebExtractCompetitorsResponseCompetitor) RawJSON() string { return r.JSON.raw }
func (r *WebExtractCompetitorsResponseCompetitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the response.
type WebExtractCompetitorsResponseStatus string

const (
	WebExtractCompetitorsResponseStatusOk WebExtractCompetitorsResponseStatus = "ok"
)

// Target company profile inferred from the landing page.
type WebExtractCompetitorsResponseTarget struct {
	// Company or product name inferred from the landing page.
	CompanyName string `json:"companyName" api:"required"`
	// Specific operating field, product category, or market.
	Field string `json:"field" api:"required"`
	// One-sentence description of what the target company sells and who it serves.
	FieldDescription string `json:"fieldDescription" api:"required"`
	// Resolved URL used for the landing page analysis.
	WebsiteURL string `json:"websiteUrl" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompanyName      respjson.Field
		Field            respjson.Field
		FieldDescription respjson.Field
		WebsiteURL       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebExtractCompetitorsResponseTarget) RawJSON() string { return r.JSON.raw }
func (r *WebExtractCompetitorsResponseTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the API key used for the request. Included in every response
// whenever a valid API key is provided, even when the response status is not 200.
type WebExtractCompetitorsResponseKeyMetadata struct {
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
func (r WebExtractCompetitorsResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebExtractCompetitorsResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebExtractFontsResponse struct {
	// HTTP status code, e.g., 200
	Code int64 `json:"code" api:"required"`
	// The normalized domain that was processed
	Domain string `json:"domain" api:"required"`
	// Array of font usage information
	Fonts []WebExtractFontsResponseFont `json:"fonts" api:"required"`
	// Status of the response, e.g., 'ok'
	Status string `json:"status" api:"required"`
	// Font assets keyed by family name as it appears in the fonts array (non-generic
	// names only). Clients match entries in fonts to pick a file URL from files.
	// Omitted when no families resolve to Google or custom @font-face URLs.
	FontLinks map[string]WebExtractFontsResponseFontLink `json:"fontLinks"`
	// Metadata about the API key used for the request. Included in every response
	// whenever a valid API key is provided, even when the response status is not 200.
	KeyMetadata WebExtractFontsResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Domain      respjson.Field
		Fonts       respjson.Field
		Status      respjson.Field
		FontLinks   respjson.Field
		KeyMetadata respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebExtractFontsResponse) RawJSON() string { return r.JSON.raw }
func (r *WebExtractFontsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebExtractFontsResponseFont struct {
	// Array of fallback font families
	Fallbacks []string `json:"fallbacks" api:"required"`
	// Font family name
	Font string `json:"font" api:"required"`
	// Number of elements using this font
	NumElements float64 `json:"num_elements" api:"required"`
	// Number of words using this font
	NumWords float64 `json:"num_words" api:"required"`
	// Percentage of elements using this font
	PercentElements float64 `json:"percent_elements" api:"required"`
	// Percentage of words using this font
	PercentWords float64 `json:"percent_words" api:"required"`
	// Array of CSS selectors or element types where this font is used
	Uses []string `json:"uses" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Fallbacks       respjson.Field
		Font            respjson.Field
		NumElements     respjson.Field
		NumWords        respjson.Field
		PercentElements respjson.Field
		PercentWords    respjson.Field
		Uses            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebExtractFontsResponseFont) RawJSON() string { return r.JSON.raw }
func (r *WebExtractFontsResponseFont) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebExtractFontsResponseFontLink struct {
	// Upright font files keyed by weight string (e.g. "400" for regular, "500",
	// "700"). Values are absolute URLs.
	Files map[string]string `json:"files" api:"required"`
	// Any of "google", "custom".
	Type string `json:"type" api:"required"`
	// Google Fonts category when type is google (e.g. sans-serif, serif, monospace,
	// display, handwriting). Omitted for custom fonts when unknown.
	Category string `json:"category"`
	// Present when type is custom: human-readable name derived from the fontLinks key
	// (strip build/hash suffixes, split camelCase / PascalCase, normalize separators).
	// Google entries omit this.
	DisplayName string `json:"displayName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Files       respjson.Field
		Type        respjson.Field
		Category    respjson.Field
		DisplayName respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebExtractFontsResponseFontLink) RawJSON() string { return r.JSON.raw }
func (r *WebExtractFontsResponseFontLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the API key used for the request. Included in every response
// whenever a valid API key is provided, even when the response status is not 200.
type WebExtractFontsResponseKeyMetadata struct {
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
func (r WebExtractFontsResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebExtractFontsResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebExtractStyleguideResponse struct {
	// HTTP status code
	Code int64 `json:"code"`
	// The normalized domain that was processed
	Domain string `json:"domain"`
	// Metadata about the API key used for the request. Included in every response
	// whenever a valid API key is provided, even when the response status is not 200.
	KeyMetadata WebExtractStyleguideResponseKeyMetadata `json:"key_metadata"`
	// Status of the response, e.g., 'ok'
	Status string `json:"status"`
	// Comprehensive styleguide data extracted from the website
	Styleguide WebExtractStyleguideResponseStyleguide `json:"styleguide"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Domain      respjson.Field
		KeyMetadata respjson.Field
		Status      respjson.Field
		Styleguide  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebExtractStyleguideResponse) RawJSON() string { return r.JSON.raw }
func (r *WebExtractStyleguideResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the API key used for the request. Included in every response
// whenever a valid API key is provided, even when the response status is not 200.
type WebExtractStyleguideResponseKeyMetadata struct {
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
func (r WebExtractStyleguideResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebExtractStyleguideResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Comprehensive styleguide data extracted from the website
type WebExtractStyleguideResponseStyleguide struct {
	// Primary colors used on the website
	Colors WebExtractStyleguideResponseStyleguideColors `json:"colors" api:"required"`
	// UI component styles
	Components WebExtractStyleguideResponseStyleguideComponents `json:"components" api:"required"`
	// Spacing system used on the website
	ElementSpacing WebExtractStyleguideResponseStyleguideElementSpacing `json:"elementSpacing" api:"required"`
	// Font assets keyed by family name as it appears in fontFamily/fontFallbacks
	// (non-generic names only). Clients match typography.fontFamily / fontWeight or
	// button styles to pick a file URL from files.
	FontLinks map[string]WebExtractStyleguideResponseStyleguideFontLink `json:"fontLinks" api:"required"`
	// The primary color mode of the website design
	//
	// Any of "light", "dark".
	Mode string `json:"mode" api:"required"`
	// Shadow styles used on the website
	Shadows WebExtractStyleguideResponseStyleguideShadows `json:"shadows" api:"required"`
	// Typography styles used on the website
	Typography WebExtractStyleguideResponseStyleguideTypography `json:"typography" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Colors         respjson.Field
		Components     respjson.Field
		ElementSpacing respjson.Field
		FontLinks      respjson.Field
		Mode           respjson.Field
		Shadows        respjson.Field
		Typography     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebExtractStyleguideResponseStyleguide) RawJSON() string { return r.JSON.raw }
func (r *WebExtractStyleguideResponseStyleguide) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Primary colors used on the website
type WebExtractStyleguideResponseStyleguideColors struct {
	// Accent color (hex format)
	Accent string `json:"accent" api:"required"`
	// Background color (hex format)
	Background string `json:"background" api:"required"`
	// Text color (hex format)
	Text string `json:"text" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accent      respjson.Field
		Background  respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebExtractStyleguideResponseStyleguideColors) RawJSON() string { return r.JSON.raw }
func (r *WebExtractStyleguideResponseStyleguideColors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UI component styles
type WebExtractStyleguideResponseStyleguideComponents struct {
	// Button component styles
	Button WebExtractStyleguideResponseStyleguideComponentsButton `json:"button" api:"required"`
	// Card component style
	Card WebExtractStyleguideResponseStyleguideComponentsCard `json:"card"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Button      respjson.Field
		Card        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebExtractStyleguideResponseStyleguideComponents) RawJSON() string { return r.JSON.raw }
func (r *WebExtractStyleguideResponseStyleguideComponents) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Button component styles
type WebExtractStyleguideResponseStyleguideComponentsButton struct {
	Link      WebExtractStyleguideResponseStyleguideComponentsButtonLink      `json:"link"`
	Primary   WebExtractStyleguideResponseStyleguideComponentsButtonPrimary   `json:"primary"`
	Secondary WebExtractStyleguideResponseStyleguideComponentsButtonSecondary `json:"secondary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Link        respjson.Field
		Primary     respjson.Field
		Secondary   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebExtractStyleguideResponseStyleguideComponentsButton) RawJSON() string { return r.JSON.raw }
func (r *WebExtractStyleguideResponseStyleguideComponentsButton) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebExtractStyleguideResponseStyleguideComponentsButtonLink struct {
	BackgroundColor string `json:"backgroundColor" api:"required"`
	// Border color as CSS hex (#RRGGBB or #RRGGBBAA when computed border-color has
	// alpha)
	BorderColor  string `json:"borderColor" api:"required"`
	BorderRadius string `json:"borderRadius" api:"required"`
	BorderStyle  string `json:"borderStyle" api:"required"`
	BorderWidth  string `json:"borderWidth" api:"required"`
	// Computed box-shadow (comma-separated layers when present)
	BoxShadow string `json:"boxShadow" api:"required"`
	Color     string `json:"color" api:"required"`
	// Ready-to-use CSS declaration block for this component style
	Css        string  `json:"css" api:"required"`
	FontSize   string  `json:"fontSize" api:"required"`
	FontWeight float64 `json:"fontWeight" api:"required"`
	// Sampled minimum height of the button box (typically px)
	MinHeight string `json:"minHeight" api:"required"`
	// Sampled minimum width of the button box (typically px)
	MinWidth       string `json:"minWidth" api:"required"`
	Padding        string `json:"padding" api:"required"`
	TextDecoration string `json:"textDecoration" api:"required"`
	// Full ordered font list from computed font-family
	FontFallbacks []string `json:"fontFallbacks"`
	// Primary button typeface (first in fontFallbacks)
	FontFamily string `json:"fontFamily"`
	// Hex color of the underline when it differs from the text color
	TextDecorationColor string `json:"textDecorationColor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BackgroundColor     respjson.Field
		BorderColor         respjson.Field
		BorderRadius        respjson.Field
		BorderStyle         respjson.Field
		BorderWidth         respjson.Field
		BoxShadow           respjson.Field
		Color               respjson.Field
		Css                 respjson.Field
		FontSize            respjson.Field
		FontWeight          respjson.Field
		MinHeight           respjson.Field
		MinWidth            respjson.Field
		Padding             respjson.Field
		TextDecoration      respjson.Field
		FontFallbacks       respjson.Field
		FontFamily          respjson.Field
		TextDecorationColor respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebExtractStyleguideResponseStyleguideComponentsButtonLink) RawJSON() string {
	return r.JSON.raw
}
func (r *WebExtractStyleguideResponseStyleguideComponentsButtonLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebExtractStyleguideResponseStyleguideComponentsButtonPrimary struct {
	BackgroundColor string `json:"backgroundColor" api:"required"`
	// Border color as CSS hex (#RRGGBB or #RRGGBBAA when computed border-color has
	// alpha)
	BorderColor  string `json:"borderColor" api:"required"`
	BorderRadius string `json:"borderRadius" api:"required"`
	BorderStyle  string `json:"borderStyle" api:"required"`
	BorderWidth  string `json:"borderWidth" api:"required"`
	// Computed box-shadow (comma-separated layers when present)
	BoxShadow string `json:"boxShadow" api:"required"`
	Color     string `json:"color" api:"required"`
	// Ready-to-use CSS declaration block for this component style
	Css        string  `json:"css" api:"required"`
	FontSize   string  `json:"fontSize" api:"required"`
	FontWeight float64 `json:"fontWeight" api:"required"`
	// Sampled minimum height of the button box (typically px)
	MinHeight string `json:"minHeight" api:"required"`
	// Sampled minimum width of the button box (typically px)
	MinWidth       string `json:"minWidth" api:"required"`
	Padding        string `json:"padding" api:"required"`
	TextDecoration string `json:"textDecoration" api:"required"`
	// Full ordered font list from computed font-family
	FontFallbacks []string `json:"fontFallbacks"`
	// Primary button typeface (first in fontFallbacks)
	FontFamily string `json:"fontFamily"`
	// Hex color of the underline when it differs from the text color
	TextDecorationColor string `json:"textDecorationColor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BackgroundColor     respjson.Field
		BorderColor         respjson.Field
		BorderRadius        respjson.Field
		BorderStyle         respjson.Field
		BorderWidth         respjson.Field
		BoxShadow           respjson.Field
		Color               respjson.Field
		Css                 respjson.Field
		FontSize            respjson.Field
		FontWeight          respjson.Field
		MinHeight           respjson.Field
		MinWidth            respjson.Field
		Padding             respjson.Field
		TextDecoration      respjson.Field
		FontFallbacks       respjson.Field
		FontFamily          respjson.Field
		TextDecorationColor respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebExtractStyleguideResponseStyleguideComponentsButtonPrimary) RawJSON() string {
	return r.JSON.raw
}
func (r *WebExtractStyleguideResponseStyleguideComponentsButtonPrimary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebExtractStyleguideResponseStyleguideComponentsButtonSecondary struct {
	BackgroundColor string `json:"backgroundColor" api:"required"`
	// Border color as CSS hex (#RRGGBB or #RRGGBBAA when computed border-color has
	// alpha)
	BorderColor  string `json:"borderColor" api:"required"`
	BorderRadius string `json:"borderRadius" api:"required"`
	BorderStyle  string `json:"borderStyle" api:"required"`
	BorderWidth  string `json:"borderWidth" api:"required"`
	// Computed box-shadow (comma-separated layers when present)
	BoxShadow string `json:"boxShadow" api:"required"`
	Color     string `json:"color" api:"required"`
	// Ready-to-use CSS declaration block for this component style
	Css        string  `json:"css" api:"required"`
	FontSize   string  `json:"fontSize" api:"required"`
	FontWeight float64 `json:"fontWeight" api:"required"`
	// Sampled minimum height of the button box (typically px)
	MinHeight string `json:"minHeight" api:"required"`
	// Sampled minimum width of the button box (typically px)
	MinWidth       string `json:"minWidth" api:"required"`
	Padding        string `json:"padding" api:"required"`
	TextDecoration string `json:"textDecoration" api:"required"`
	// Full ordered font list from computed font-family
	FontFallbacks []string `json:"fontFallbacks"`
	// Primary button typeface (first in fontFallbacks)
	FontFamily string `json:"fontFamily"`
	// Hex color of the underline when it differs from the text color
	TextDecorationColor string `json:"textDecorationColor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BackgroundColor     respjson.Field
		BorderColor         respjson.Field
		BorderRadius        respjson.Field
		BorderStyle         respjson.Field
		BorderWidth         respjson.Field
		BoxShadow           respjson.Field
		Color               respjson.Field
		Css                 respjson.Field
		FontSize            respjson.Field
		FontWeight          respjson.Field
		MinHeight           respjson.Field
		MinWidth            respjson.Field
		Padding             respjson.Field
		TextDecoration      respjson.Field
		FontFallbacks       respjson.Field
		FontFamily          respjson.Field
		TextDecorationColor respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebExtractStyleguideResponseStyleguideComponentsButtonSecondary) RawJSON() string {
	return r.JSON.raw
}
func (r *WebExtractStyleguideResponseStyleguideComponentsButtonSecondary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Card component style
type WebExtractStyleguideResponseStyleguideComponentsCard struct {
	BackgroundColor string `json:"backgroundColor" api:"required"`
	// Border color as CSS hex (#RRGGBB or #RRGGBBAA when computed border-color has
	// alpha)
	BorderColor  string `json:"borderColor" api:"required"`
	BorderRadius string `json:"borderRadius" api:"required"`
	BorderStyle  string `json:"borderStyle" api:"required"`
	BorderWidth  string `json:"borderWidth" api:"required"`
	BoxShadow    string `json:"boxShadow" api:"required"`
	// Ready-to-use CSS declaration block for this component style
	Css       string `json:"css" api:"required"`
	Padding   string `json:"padding" api:"required"`
	TextColor string `json:"textColor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BackgroundColor respjson.Field
		BorderColor     respjson.Field
		BorderRadius    respjson.Field
		BorderStyle     respjson.Field
		BorderWidth     respjson.Field
		BoxShadow       respjson.Field
		Css             respjson.Field
		Padding         respjson.Field
		TextColor       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebExtractStyleguideResponseStyleguideComponentsCard) RawJSON() string { return r.JSON.raw }
func (r *WebExtractStyleguideResponseStyleguideComponentsCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Spacing system used on the website
type WebExtractStyleguideResponseStyleguideElementSpacing struct {
	Lg string `json:"lg" api:"required"`
	Md string `json:"md" api:"required"`
	Sm string `json:"sm" api:"required"`
	Xl string `json:"xl" api:"required"`
	Xs string `json:"xs" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Lg          respjson.Field
		Md          respjson.Field
		Sm          respjson.Field
		Xl          respjson.Field
		Xs          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebExtractStyleguideResponseStyleguideElementSpacing) RawJSON() string { return r.JSON.raw }
func (r *WebExtractStyleguideResponseStyleguideElementSpacing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebExtractStyleguideResponseStyleguideFontLink struct {
	// Upright font files keyed by weight string (e.g. "400" for regular, "500",
	// "700"). Values are absolute URLs.
	Files map[string]string `json:"files" api:"required"`
	// Any of "google", "custom".
	Type string `json:"type" api:"required"`
	// Google Fonts category when type is google (e.g. sans-serif, serif, monospace,
	// display, handwriting). Omitted for custom fonts when unknown.
	Category string `json:"category"`
	// Present when type is custom: human-readable name derived from the fontLinks key
	// (strip build/hash suffixes, split camelCase / PascalCase, normalize separators).
	// Google entries omit this.
	DisplayName string `json:"displayName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Files       respjson.Field
		Type        respjson.Field
		Category    respjson.Field
		DisplayName respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebExtractStyleguideResponseStyleguideFontLink) RawJSON() string { return r.JSON.raw }
func (r *WebExtractStyleguideResponseStyleguideFontLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Shadow styles used on the website
type WebExtractStyleguideResponseStyleguideShadows struct {
	Inner string `json:"inner" api:"required"`
	Lg    string `json:"lg" api:"required"`
	Md    string `json:"md" api:"required"`
	Sm    string `json:"sm" api:"required"`
	Xl    string `json:"xl" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Inner       respjson.Field
		Lg          respjson.Field
		Md          respjson.Field
		Sm          respjson.Field
		Xl          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebExtractStyleguideResponseStyleguideShadows) RawJSON() string { return r.JSON.raw }
func (r *WebExtractStyleguideResponseStyleguideShadows) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typography styles used on the website
type WebExtractStyleguideResponseStyleguideTypography struct {
	// Heading styles
	Headings WebExtractStyleguideResponseStyleguideTypographyHeadings `json:"headings" api:"required"`
	P        WebExtractStyleguideResponseStyleguideTypographyP        `json:"p"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Headings    respjson.Field
		P           respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebExtractStyleguideResponseStyleguideTypography) RawJSON() string { return r.JSON.raw }
func (r *WebExtractStyleguideResponseStyleguideTypography) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Heading styles
type WebExtractStyleguideResponseStyleguideTypographyHeadings struct {
	H1 WebExtractStyleguideResponseStyleguideTypographyHeadingsH1 `json:"h1"`
	H2 WebExtractStyleguideResponseStyleguideTypographyHeadingsH2 `json:"h2"`
	H3 WebExtractStyleguideResponseStyleguideTypographyHeadingsH3 `json:"h3"`
	H4 WebExtractStyleguideResponseStyleguideTypographyHeadingsH4 `json:"h4"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		H1          respjson.Field
		H2          respjson.Field
		H3          respjson.Field
		H4          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebExtractStyleguideResponseStyleguideTypographyHeadings) RawJSON() string { return r.JSON.raw }
func (r *WebExtractStyleguideResponseStyleguideTypographyHeadings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebExtractStyleguideResponseStyleguideTypographyHeadingsH1 struct {
	// Full ordered font list from resolved computed font-family
	FontFallbacks []string `json:"fontFallbacks" api:"required"`
	// Primary face (first family in the computed stack)
	FontFamily    string  `json:"fontFamily" api:"required"`
	FontSize      string  `json:"fontSize" api:"required"`
	FontWeight    float64 `json:"fontWeight" api:"required"`
	LetterSpacing string  `json:"letterSpacing" api:"required"`
	LineHeight    string  `json:"lineHeight" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FontFallbacks respjson.Field
		FontFamily    respjson.Field
		FontSize      respjson.Field
		FontWeight    respjson.Field
		LetterSpacing respjson.Field
		LineHeight    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebExtractStyleguideResponseStyleguideTypographyHeadingsH1) RawJSON() string {
	return r.JSON.raw
}
func (r *WebExtractStyleguideResponseStyleguideTypographyHeadingsH1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebExtractStyleguideResponseStyleguideTypographyHeadingsH2 struct {
	// Full ordered font list from resolved computed font-family
	FontFallbacks []string `json:"fontFallbacks" api:"required"`
	// Primary face (first family in the computed stack)
	FontFamily    string  `json:"fontFamily" api:"required"`
	FontSize      string  `json:"fontSize" api:"required"`
	FontWeight    float64 `json:"fontWeight" api:"required"`
	LetterSpacing string  `json:"letterSpacing" api:"required"`
	LineHeight    string  `json:"lineHeight" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FontFallbacks respjson.Field
		FontFamily    respjson.Field
		FontSize      respjson.Field
		FontWeight    respjson.Field
		LetterSpacing respjson.Field
		LineHeight    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebExtractStyleguideResponseStyleguideTypographyHeadingsH2) RawJSON() string {
	return r.JSON.raw
}
func (r *WebExtractStyleguideResponseStyleguideTypographyHeadingsH2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebExtractStyleguideResponseStyleguideTypographyHeadingsH3 struct {
	// Full ordered font list from resolved computed font-family
	FontFallbacks []string `json:"fontFallbacks" api:"required"`
	// Primary face (first family in the computed stack)
	FontFamily    string  `json:"fontFamily" api:"required"`
	FontSize      string  `json:"fontSize" api:"required"`
	FontWeight    float64 `json:"fontWeight" api:"required"`
	LetterSpacing string  `json:"letterSpacing" api:"required"`
	LineHeight    string  `json:"lineHeight" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FontFallbacks respjson.Field
		FontFamily    respjson.Field
		FontSize      respjson.Field
		FontWeight    respjson.Field
		LetterSpacing respjson.Field
		LineHeight    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebExtractStyleguideResponseStyleguideTypographyHeadingsH3) RawJSON() string {
	return r.JSON.raw
}
func (r *WebExtractStyleguideResponseStyleguideTypographyHeadingsH3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebExtractStyleguideResponseStyleguideTypographyHeadingsH4 struct {
	// Full ordered font list from resolved computed font-family
	FontFallbacks []string `json:"fontFallbacks" api:"required"`
	// Primary face (first family in the computed stack)
	FontFamily    string  `json:"fontFamily" api:"required"`
	FontSize      string  `json:"fontSize" api:"required"`
	FontWeight    float64 `json:"fontWeight" api:"required"`
	LetterSpacing string  `json:"letterSpacing" api:"required"`
	LineHeight    string  `json:"lineHeight" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FontFallbacks respjson.Field
		FontFamily    respjson.Field
		FontSize      respjson.Field
		FontWeight    respjson.Field
		LetterSpacing respjson.Field
		LineHeight    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebExtractStyleguideResponseStyleguideTypographyHeadingsH4) RawJSON() string {
	return r.JSON.raw
}
func (r *WebExtractStyleguideResponseStyleguideTypographyHeadingsH4) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebExtractStyleguideResponseStyleguideTypographyP struct {
	// Full ordered font list from resolved computed font-family
	FontFallbacks []string `json:"fontFallbacks" api:"required"`
	// Primary face (first family in the computed stack)
	FontFamily    string  `json:"fontFamily" api:"required"`
	FontSize      string  `json:"fontSize" api:"required"`
	FontWeight    float64 `json:"fontWeight" api:"required"`
	LetterSpacing string  `json:"letterSpacing" api:"required"`
	LineHeight    string  `json:"lineHeight" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FontFallbacks respjson.Field
		FontFamily    respjson.Field
		FontSize      respjson.Field
		FontWeight    respjson.Field
		LetterSpacing respjson.Field
		LineHeight    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebExtractStyleguideResponseStyleguideTypographyP) RawJSON() string { return r.JSON.raw }
func (r *WebExtractStyleguideResponseStyleguideTypographyP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScreenshotResponse struct {
	// HTTP status code
	Code int64 `json:"code"`
	// The normalized domain that was processed
	Domain string `json:"domain"`
	// Height in pixels of the returned screenshot image
	Height int64 `json:"height"`
	// Metadata about the API key used for the request. Included in every response
	// whenever a valid API key is provided, even when the response status is not 200.
	KeyMetadata WebScreenshotResponseKeyMetadata `json:"key_metadata"`
	// Public image URL for standard requests, or an in-memory data URL when ZDR is
	// enabled.
	Screenshot string `json:"screenshot"`
	// Type of screenshot that was captured
	//
	// Any of "viewport", "fullPage".
	ScreenshotType WebScreenshotResponseScreenshotType `json:"screenshotType"`
	// Status of the response, e.g., 'ok'
	Status string `json:"status"`
	// Width in pixels of the returned screenshot image
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code           respjson.Field
		Domain         respjson.Field
		Height         respjson.Field
		KeyMetadata    respjson.Field
		Screenshot     respjson.Field
		ScreenshotType respjson.Field
		Status         respjson.Field
		Width          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScreenshotResponse) RawJSON() string { return r.JSON.raw }
func (r *WebScreenshotResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the API key used for the request. Included in every response
// whenever a valid API key is provided, even when the response status is not 200.
type WebScreenshotResponseKeyMetadata struct {
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
func (r WebScreenshotResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebScreenshotResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of screenshot that was captured
type WebScreenshotResponseScreenshotType string

const (
	WebScreenshotResponseScreenshotTypeViewport WebScreenshotResponseScreenshotType = "viewport"
	WebScreenshotResponseScreenshotTypeFullPage WebScreenshotResponseScreenshotType = "fullPage"
)

type WebSearchResponse struct {
	// Echo of the original query (useful when fanout was enabled).
	Query   string                    `json:"query" api:"required"`
	Results []WebSearchResponseResult `json:"results" api:"required"`
	// Metadata about the API key used for the request. Included in every response
	// whenever a valid API key is provided, even when the response status is not 200.
	KeyMetadata WebSearchResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Query       respjson.Field
		Results     respjson.Field
		KeyMetadata respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *WebSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebSearchResponseResult struct {
	// Snippet excerpt from the page.
	Description string `json:"description" api:"required"`
	// Markdown scrape status and content for this result.
	Markdown WebSearchResponseResultMarkdown `json:"markdown" api:"required"`
	// Relevance to the original query.
	//
	// Any of "high", "medium", "low".
	Relevance string `json:"relevance" api:"required"`
	// Page title.
	Title string `json:"title" api:"required"`
	// Canonical result URL.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Markdown    respjson.Field
		Relevance   respjson.Field
		Title       respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebSearchResponseResult) RawJSON() string { return r.JSON.raw }
func (r *WebSearchResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Markdown scrape status and content for this result.
type WebSearchResponseResultMarkdown struct {
	// Per-result scrape outcome. Inspect this before reading `markdown`.
	//
	// Any of "SUCCESS", "NOT_REQUESTED", "TIMEOUT", "WEBSITE_ACCESS_ERROR", "ERROR".
	Code string `json:"code" api:"required"`
	// GFM Markdown of the page. Null unless markdownOptions.enabled is true and
	// scraping succeeded.
	Markdown string `json:"markdown" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Markdown    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebSearchResponseResultMarkdown) RawJSON() string { return r.JSON.raw }
func (r *WebSearchResponseResultMarkdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the API key used for the request. Included in every response
// whenever a valid API key is provided, even when the response status is not 200.
type WebSearchResponseKeyMetadata struct {
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
func (r WebSearchResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebSearchResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebWebCrawlMdResponse struct {
	Metadata WebWebCrawlMdResponseMetadata `json:"metadata" api:"required"`
	Results  []WebWebCrawlMdResponseResult `json:"results" api:"required"`
	// Metadata about the API key used for the request. Included in every response
	// whenever a valid API key is provided, even when the response status is not 200.
	KeyMetadata WebWebCrawlMdResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Results     respjson.Field
		KeyMetadata respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebWebCrawlMdResponse) RawJSON() string { return r.JSON.raw }
func (r *WebWebCrawlMdResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebWebCrawlMdResponseMetadata struct {
	// Maximum crawl depth reached during the crawl
	MaxCrawlDepth int64 `json:"maxCrawlDepth" api:"required"`
	// Number of pages that failed to crawl
	NumFailed int64 `json:"numFailed" api:"required"`
	// Number of URLs skipped (PDFs when pdf.shouldParse=false, or URLs not matching
	// urlRegex)
	NumSkipped int64 `json:"numSkipped" api:"required"`
	// Number of pages successfully crawled
	NumSucceeded int64 `json:"numSucceeded" api:"required"`
	// Total number of URLs crawled
	NumURLs int64 `json:"numUrls" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxCrawlDepth respjson.Field
		NumFailed     respjson.Field
		NumSkipped    respjson.Field
		NumSucceeded  respjson.Field
		NumURLs       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebWebCrawlMdResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebWebCrawlMdResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebWebCrawlMdResponseResult struct {
	// Extracted page content as Markdown (empty string on failure)
	Markdown string                              `json:"markdown" api:"required"`
	Metadata WebWebCrawlMdResponseResultMetadata `json:"metadata" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Markdown    respjson.Field
		Metadata    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebWebCrawlMdResponseResult) RawJSON() string { return r.JSON.raw }
func (r *WebWebCrawlMdResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebWebCrawlMdResponseResultMetadata struct {
	// Depth relative to the start URL. 0 = start URL, 1 = one link away.
	CrawlDepth int64 `json:"crawlDepth" api:"required"`
	// Final URL scraped after redirects or scraper fallback, when known. Falls back to
	// sourceUrl when unavailable.
	FinalURL string `json:"finalUrl" api:"required"`
	// Original URL requested by the caller.
	SourceURL string `json:"sourceUrl" api:"required"`
	// HTTP status code of the response
	StatusCode int64 `json:"statusCode" api:"required"`
	// true if the page was fetched and parsed successfully
	Success bool `json:"success" api:"required"`
	// Best page title extracted from the page (empty string if unavailable).
	Title string `json:"title" api:"required"`
	// The crawl URL fetched for this page.
	URL string `json:"url" api:"required"`
	// Additional non-social meta tags not promoted to top-level metadata fields.
	AdditionalMeta map[string]WebWebCrawlMdResponseResultMetadataAdditionalMetaUnion `json:"additionalMeta"`
	// Resolved alternate links from link rel=alternate tags.
	Alternates []WebWebCrawlMdResponseResultMetadataAlternate `json:"alternates"`
	// Author metadata, when present.
	Author string `json:"author"`
	// Resolved canonical URL, when present.
	CanonicalURL string `json:"canonicalUrl"`
	// Best description extracted from standard, Open Graph, or Twitter metadata.
	Description string `json:"description"`
	// Resolved favicon URL, when present.
	Favicon string `json:"favicon"`
	// Primary resolved preview image from Open Graph, Twitter, or image metadata.
	Image string `json:"image"`
	// JSON-LD structured data blocks parsed from the page.
	JsonLd []map[string]any `json:"jsonLd"`
	// Keywords extracted from the page's keywords meta tag.
	Keywords []string `json:"keywords"`
	// Language extracted from html lang or language meta tags.
	Language string `json:"language"`
	// Modified timestamp/date from page metadata, when present.
	ModifiedTime string `json:"modifiedTime"`
	// Open Graph metadata with the og: prefix removed and keys camel-cased.
	OpenGraph map[string]WebWebCrawlMdResponseResultMetadataOpenGraphUnion `json:"openGraph"`
	// Published timestamp/date from page metadata, when present.
	PublishedTime string `json:"publishedTime"`
	// Robots meta directive, when present.
	Robots string `json:"robots"`
	// Site or application name from page metadata.
	SiteName string `json:"siteName"`
	// Twitter card metadata with the twitter: prefix removed and keys camel-cased.
	Twitter map[string]WebWebCrawlMdResponseResultMetadataTwitterUnion `json:"twitter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CrawlDepth     respjson.Field
		FinalURL       respjson.Field
		SourceURL      respjson.Field
		StatusCode     respjson.Field
		Success        respjson.Field
		Title          respjson.Field
		URL            respjson.Field
		AdditionalMeta respjson.Field
		Alternates     respjson.Field
		Author         respjson.Field
		CanonicalURL   respjson.Field
		Description    respjson.Field
		Favicon        respjson.Field
		Image          respjson.Field
		JsonLd         respjson.Field
		Keywords       respjson.Field
		Language       respjson.Field
		ModifiedTime   respjson.Field
		OpenGraph      respjson.Field
		PublishedTime  respjson.Field
		Robots         respjson.Field
		SiteName       respjson.Field
		Twitter        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebWebCrawlMdResponseResultMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebWebCrawlMdResponseResultMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WebWebCrawlMdResponseResultMetadataAdditionalMetaUnion contains all possible
// properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type WebWebCrawlMdResponseResultMetadataAdditionalMetaUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u WebWebCrawlMdResponseResultMetadataAdditionalMetaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebWebCrawlMdResponseResultMetadataAdditionalMetaUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WebWebCrawlMdResponseResultMetadataAdditionalMetaUnion) RawJSON() string { return u.JSON.raw }

func (r *WebWebCrawlMdResponseResultMetadataAdditionalMetaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebWebCrawlMdResponseResultMetadataAlternate struct {
	// Resolved alternate URL.
	Href string `json:"href" api:"required"`
	// Language or locale for the alternate URL, when present.
	Hreflang string `json:"hreflang"`
	// Alternate resource title, when present.
	Title string `json:"title"`
	// Alternate resource MIME type, when present.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Href        respjson.Field
		Hreflang    respjson.Field
		Title       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebWebCrawlMdResponseResultMetadataAlternate) RawJSON() string { return r.JSON.raw }
func (r *WebWebCrawlMdResponseResultMetadataAlternate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WebWebCrawlMdResponseResultMetadataOpenGraphUnion contains all possible
// properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type WebWebCrawlMdResponseResultMetadataOpenGraphUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u WebWebCrawlMdResponseResultMetadataOpenGraphUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebWebCrawlMdResponseResultMetadataOpenGraphUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WebWebCrawlMdResponseResultMetadataOpenGraphUnion) RawJSON() string { return u.JSON.raw }

func (r *WebWebCrawlMdResponseResultMetadataOpenGraphUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WebWebCrawlMdResponseResultMetadataTwitterUnion contains all possible properties
// and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type WebWebCrawlMdResponseResultMetadataTwitterUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u WebWebCrawlMdResponseResultMetadataTwitterUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebWebCrawlMdResponseResultMetadataTwitterUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WebWebCrawlMdResponseResultMetadataTwitterUnion) RawJSON() string { return u.JSON.raw }

func (r *WebWebCrawlMdResponseResultMetadataTwitterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the API key used for the request. Included in every response
// whenever a valid API key is provided, even when the response status is not 200.
type WebWebCrawlMdResponseKeyMetadata struct {
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
func (r WebWebCrawlMdResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebWebCrawlMdResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebWebScrapeHTMLResponse struct {
	// The scraped content of the page. For normal pages this is the raw HTML. When the
	// page is a sitemap or feed served behind an XSL stylesheet (which browsers render
	// into HTML), this is the underlying XML instead — see the `type` field.
	HTML string `json:"html" api:"required"`
	// Metadata extracted from the scraped page HTML.
	Metadata WebWebScrapeHTMLResponseMetadata `json:"metadata" api:"required"`
	// Indicates success
	//
	// Any of true.
	Success bool `json:"success" api:"required"`
	// Detected content type of the returned `html` field. Sitemaps and feeds are
	// surfaced as `xml`; ordinary pages are `html`. Excel workbooks are surfaced as
	// `xlsx`/`xls` with the extracted sheets as HTML tables; PowerPoint presentations
	// are surfaced as `pptx`/`ppt` with the extracted slides as HTML.
	//
	// Any of "html", "xml", "json", "text", "csv", "markdown", "svg", "pdf", "docx",
	// "doc", "xlsx", "xls", "pptx", "ppt".
	Type WebWebScrapeHTMLResponseType `json:"type" api:"required"`
	// The URL that was scraped
	URL string `json:"url" api:"required"`
	// One verified outcome per requested browser action, in request order.
	ActionsApplied []WebWebScrapeHTMLResponseActionsApplied `json:"actionsApplied"`
	// True when an action was applied but the returned content could not be refreshed
	// afterward.
	ActionsHTMLStale bool `json:"actionsHtmlStale"`
	// Metadata about the API key used for the request. Included in every response
	// whenever a valid API key is provided, even when the response status is not 200.
	KeyMetadata WebWebScrapeHTMLResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HTML             respjson.Field
		Metadata         respjson.Field
		Success          respjson.Field
		Type             respjson.Field
		URL              respjson.Field
		ActionsApplied   respjson.Field
		ActionsHTMLStale respjson.Field
		KeyMetadata      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebWebScrapeHTMLResponse) RawJSON() string { return r.JSON.raw }
func (r *WebWebScrapeHTMLResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata extracted from the scraped page HTML.
type WebWebScrapeHTMLResponseMetadata struct {
	// Final URL scraped after redirects or scraper fallback, when known. Falls back to
	// sourceUrl when unavailable.
	FinalURL string `json:"finalUrl" api:"required"`
	// Original URL requested by the caller.
	SourceURL string `json:"sourceUrl" api:"required"`
	// Additional non-social meta tags not promoted to top-level metadata fields.
	AdditionalMeta map[string]WebWebScrapeHTMLResponseMetadataAdditionalMetaUnion `json:"additionalMeta"`
	// Resolved alternate links from link rel=alternate tags.
	Alternates []WebWebScrapeHTMLResponseMetadataAlternate `json:"alternates"`
	// Author metadata, when present.
	Author string `json:"author"`
	// Resolved canonical URL, when present.
	CanonicalURL string `json:"canonicalUrl"`
	// Best description extracted from standard, Open Graph, or Twitter metadata.
	Description string `json:"description"`
	// Resolved favicon URL, when present.
	Favicon string `json:"favicon"`
	// Primary resolved preview image from Open Graph, Twitter, or image metadata.
	Image string `json:"image"`
	// JSON-LD structured data blocks parsed from the page.
	JsonLd []map[string]any `json:"jsonLd"`
	// Keywords extracted from the page's keywords meta tag.
	Keywords []string `json:"keywords"`
	// Language extracted from html lang or language meta tags.
	Language string `json:"language"`
	// Modified timestamp/date from page metadata, when present.
	ModifiedTime string `json:"modifiedTime"`
	// Open Graph metadata with the og: prefix removed and keys camel-cased.
	OpenGraph map[string]WebWebScrapeHTMLResponseMetadataOpenGraphUnion `json:"openGraph"`
	// Published timestamp/date from page metadata, when present.
	PublishedTime string `json:"publishedTime"`
	// Robots meta directive, when present.
	Robots string `json:"robots"`
	// Site or application name from page metadata.
	SiteName string `json:"siteName"`
	// Best title extracted from the page.
	Title string `json:"title"`
	// Twitter card metadata with the twitter: prefix removed and keys camel-cased.
	Twitter map[string]WebWebScrapeHTMLResponseMetadataTwitterUnion `json:"twitter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FinalURL       respjson.Field
		SourceURL      respjson.Field
		AdditionalMeta respjson.Field
		Alternates     respjson.Field
		Author         respjson.Field
		CanonicalURL   respjson.Field
		Description    respjson.Field
		Favicon        respjson.Field
		Image          respjson.Field
		JsonLd         respjson.Field
		Keywords       respjson.Field
		Language       respjson.Field
		ModifiedTime   respjson.Field
		OpenGraph      respjson.Field
		PublishedTime  respjson.Field
		Robots         respjson.Field
		SiteName       respjson.Field
		Title          respjson.Field
		Twitter        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebWebScrapeHTMLResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebWebScrapeHTMLResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WebWebScrapeHTMLResponseMetadataAdditionalMetaUnion contains all possible
// properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type WebWebScrapeHTMLResponseMetadataAdditionalMetaUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u WebWebScrapeHTMLResponseMetadataAdditionalMetaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebWebScrapeHTMLResponseMetadataAdditionalMetaUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WebWebScrapeHTMLResponseMetadataAdditionalMetaUnion) RawJSON() string { return u.JSON.raw }

func (r *WebWebScrapeHTMLResponseMetadataAdditionalMetaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebWebScrapeHTMLResponseMetadataAlternate struct {
	// Resolved alternate URL.
	Href string `json:"href" api:"required"`
	// Language or locale for the alternate URL, when present.
	Hreflang string `json:"hreflang"`
	// Alternate resource title, when present.
	Title string `json:"title"`
	// Alternate resource MIME type, when present.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Href        respjson.Field
		Hreflang    respjson.Field
		Title       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebWebScrapeHTMLResponseMetadataAlternate) RawJSON() string { return r.JSON.raw }
func (r *WebWebScrapeHTMLResponseMetadataAlternate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WebWebScrapeHTMLResponseMetadataOpenGraphUnion contains all possible properties
// and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type WebWebScrapeHTMLResponseMetadataOpenGraphUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u WebWebScrapeHTMLResponseMetadataOpenGraphUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebWebScrapeHTMLResponseMetadataOpenGraphUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WebWebScrapeHTMLResponseMetadataOpenGraphUnion) RawJSON() string { return u.JSON.raw }

func (r *WebWebScrapeHTMLResponseMetadataOpenGraphUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WebWebScrapeHTMLResponseMetadataTwitterUnion contains all possible properties
// and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type WebWebScrapeHTMLResponseMetadataTwitterUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u WebWebScrapeHTMLResponseMetadataTwitterUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebWebScrapeHTMLResponseMetadataTwitterUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WebWebScrapeHTMLResponseMetadataTwitterUnion) RawJSON() string { return u.JSON.raw }

func (r *WebWebScrapeHTMLResponseMetadataTwitterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detected content type of the returned `html` field. Sitemaps and feeds are
// surfaced as `xml`; ordinary pages are `html`. Excel workbooks are surfaced as
// `xlsx`/`xls` with the extracted sheets as HTML tables; PowerPoint presentations
// are surfaced as `pptx`/`ppt` with the extracted slides as HTML.
type WebWebScrapeHTMLResponseType string

const (
	WebWebScrapeHTMLResponseTypeHTML     WebWebScrapeHTMLResponseType = "html"
	WebWebScrapeHTMLResponseTypeXml      WebWebScrapeHTMLResponseType = "xml"
	WebWebScrapeHTMLResponseTypeJson     WebWebScrapeHTMLResponseType = "json"
	WebWebScrapeHTMLResponseTypeText     WebWebScrapeHTMLResponseType = "text"
	WebWebScrapeHTMLResponseTypeCsv      WebWebScrapeHTMLResponseType = "csv"
	WebWebScrapeHTMLResponseTypeMarkdown WebWebScrapeHTMLResponseType = "markdown"
	WebWebScrapeHTMLResponseTypeSvg      WebWebScrapeHTMLResponseType = "svg"
	WebWebScrapeHTMLResponseTypePdf      WebWebScrapeHTMLResponseType = "pdf"
	WebWebScrapeHTMLResponseTypeDocx     WebWebScrapeHTMLResponseType = "docx"
	WebWebScrapeHTMLResponseTypeDoc      WebWebScrapeHTMLResponseType = "doc"
	WebWebScrapeHTMLResponseTypeXlsx     WebWebScrapeHTMLResponseType = "xlsx"
	WebWebScrapeHTMLResponseTypeXls      WebWebScrapeHTMLResponseType = "xls"
	WebWebScrapeHTMLResponseTypePptx     WebWebScrapeHTMLResponseType = "pptx"
	WebWebScrapeHTMLResponseTypePpt      WebWebScrapeHTMLResponseType = "ppt"
)

type WebWebScrapeHTMLResponseActionsApplied struct {
	Instruction string `json:"instruction" api:"required"`
	// Applied means the requested page state was visibly verified. Failed means it was
	// not verified. Skipped means it was not attempted.
	//
	// Any of "applied", "failed", "skipped".
	Status string `json:"status" api:"required"`
	// Visible page evidence used to verify an applied action.
	CompletionEvidence string  `json:"completionEvidence"`
	DurationMs         float64 `json:"durationMs"`
	Error              string  `json:"error"`
	Method             string  `json:"method"`
	TargetDescription  string  `json:"targetDescription"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Instruction        respjson.Field
		Status             respjson.Field
		CompletionEvidence respjson.Field
		DurationMs         respjson.Field
		Error              respjson.Field
		Method             respjson.Field
		TargetDescription  respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebWebScrapeHTMLResponseActionsApplied) RawJSON() string { return r.JSON.raw }
func (r *WebWebScrapeHTMLResponseActionsApplied) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the API key used for the request. Included in every response
// whenever a valid API key is provided, even when the response status is not 200.
type WebWebScrapeHTMLResponseKeyMetadata struct {
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
func (r WebWebScrapeHTMLResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebWebScrapeHTMLResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebWebScrapeImagesResponse struct {
	// Images found on the page.
	Images []WebWebScrapeImagesResponseImage `json:"images" api:"required"`
	// Always true on success.
	//
	// Any of true.
	Success bool `json:"success" api:"required"`
	// Page URL that was scraped.
	URL string `json:"url" api:"required"`
	// Metadata about the API key used for the request. Included in every response
	// whenever a valid API key is provided, even when the response status is not 200.
	KeyMetadata WebWebScrapeImagesResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Images      respjson.Field
		Success     respjson.Field
		URL         respjson.Field
		KeyMetadata respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebWebScrapeImagesResponse) RawJSON() string { return r.JSON.raw }
func (r *WebWebScrapeImagesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebWebScrapeImagesResponseImage struct {
	// Image alt text, or null when unavailable.
	Alt string `json:"alt" api:"required"`
	// Where the image was found.
	//
	// Any of "img", "svg", "link", "source", "video", "css", "object", "meta",
	// "background".
	Element string `json:"element" api:"required"`
	// Original image value: URL, inline SVG or HTML, or base64 data URI.
	Src string `json:"src" api:"required"`
	// Format of src.
	//
	// Any of "url", "html", "base64".
	Type string `json:"type" api:"required"`
	// Requested metadata for images that could be processed.
	Enrichment WebWebScrapeImagesResponseImageEnrichment `json:"enrichment"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Alt         respjson.Field
		Element     respjson.Field
		Src         respjson.Field
		Type        respjson.Field
		Enrichment  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebWebScrapeImagesResponseImage) RawJSON() string { return r.JSON.raw }
func (r *WebWebScrapeImagesResponseImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Requested metadata for images that could be processed.
type WebWebScrapeImagesResponseImageEnrichment struct {
	// Image height in pixels, when measured.
	Height int64 `json:"height"`
	// Detected MIME type, when hosted.
	Mimetype string `json:"mimetype"`
	// Visual asset category, when classified.
	//
	// Any of "photography", "illustration", "logo", "wordmark", "icon", "pattern",
	// "graphic", "other".
	Type string `json:"type"`
	// Brand.dev CDN URL, when hosted.
	URL string `json:"url" format:"uri"`
	// Image width in pixels, when measured.
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Mimetype    respjson.Field
		Type        respjson.Field
		URL         respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebWebScrapeImagesResponseImageEnrichment) RawJSON() string { return r.JSON.raw }
func (r *WebWebScrapeImagesResponseImageEnrichment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the API key used for the request. Included in every response
// whenever a valid API key is provided, even when the response status is not 200.
type WebWebScrapeImagesResponseKeyMetadata struct {
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
func (r WebWebScrapeImagesResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebWebScrapeImagesResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebWebScrapeMdResponse struct {
	// UTF-8 byte length of the returned Markdown. Use 0 to identify an empty result
	// and compare small values against your workload's minimum useful-content
	// threshold.
	ContentLength int64 `json:"contentLength" api:"required"`
	// Page content converted to GitHub Flavored Markdown
	Markdown string `json:"markdown" api:"required"`
	// Metadata extracted from the scraped page HTML.
	Metadata WebWebScrapeMdResponseMetadata `json:"metadata" api:"required"`
	// Indicates success
	//
	// Any of true.
	Success bool `json:"success" api:"required"`
	// The URL that was scraped
	URL string `json:"url" api:"required"`
	// One verified outcome per requested browser action, in request order.
	ActionsApplied []WebWebScrapeMdResponseActionsApplied `json:"actionsApplied"`
	// True when an action was applied but the returned content could not be refreshed
	// afterward.
	ActionsHTMLStale bool `json:"actionsHtmlStale"`
	// Metadata about the API key used for the request. Included in every response
	// whenever a valid API key is provided, even when the response status is not 200.
	KeyMetadata WebWebScrapeMdResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentLength    respjson.Field
		Markdown         respjson.Field
		Metadata         respjson.Field
		Success          respjson.Field
		URL              respjson.Field
		ActionsApplied   respjson.Field
		ActionsHTMLStale respjson.Field
		KeyMetadata      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebWebScrapeMdResponse) RawJSON() string { return r.JSON.raw }
func (r *WebWebScrapeMdResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata extracted from the scraped page HTML.
type WebWebScrapeMdResponseMetadata struct {
	// Final URL scraped after redirects or scraper fallback, when known. Falls back to
	// sourceUrl when unavailable.
	FinalURL string `json:"finalUrl" api:"required"`
	// Original URL requested by the caller.
	SourceURL string `json:"sourceUrl" api:"required"`
	// Additional non-social meta tags not promoted to top-level metadata fields.
	AdditionalMeta map[string]WebWebScrapeMdResponseMetadataAdditionalMetaUnion `json:"additionalMeta"`
	// Resolved alternate links from link rel=alternate tags.
	Alternates []WebWebScrapeMdResponseMetadataAlternate `json:"alternates"`
	// Author metadata, when present.
	Author string `json:"author"`
	// Resolved canonical URL, when present.
	CanonicalURL string `json:"canonicalUrl"`
	// Best description extracted from standard, Open Graph, or Twitter metadata.
	Description string `json:"description"`
	// Resolved favicon URL, when present.
	Favicon string `json:"favicon"`
	// Primary resolved preview image from Open Graph, Twitter, or image metadata.
	Image string `json:"image"`
	// JSON-LD structured data blocks parsed from the page.
	JsonLd []map[string]any `json:"jsonLd"`
	// Keywords extracted from the page's keywords meta tag.
	Keywords []string `json:"keywords"`
	// Language extracted from html lang or language meta tags.
	Language string `json:"language"`
	// Modified timestamp/date from page metadata, when present.
	ModifiedTime string `json:"modifiedTime"`
	// Open Graph metadata with the og: prefix removed and keys camel-cased.
	OpenGraph map[string]WebWebScrapeMdResponseMetadataOpenGraphUnion `json:"openGraph"`
	// Published timestamp/date from page metadata, when present.
	PublishedTime string `json:"publishedTime"`
	// Robots meta directive, when present.
	Robots string `json:"robots"`
	// Site or application name from page metadata.
	SiteName string `json:"siteName"`
	// Best title extracted from the page.
	Title string `json:"title"`
	// Twitter card metadata with the twitter: prefix removed and keys camel-cased.
	Twitter map[string]WebWebScrapeMdResponseMetadataTwitterUnion `json:"twitter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FinalURL       respjson.Field
		SourceURL      respjson.Field
		AdditionalMeta respjson.Field
		Alternates     respjson.Field
		Author         respjson.Field
		CanonicalURL   respjson.Field
		Description    respjson.Field
		Favicon        respjson.Field
		Image          respjson.Field
		JsonLd         respjson.Field
		Keywords       respjson.Field
		Language       respjson.Field
		ModifiedTime   respjson.Field
		OpenGraph      respjson.Field
		PublishedTime  respjson.Field
		Robots         respjson.Field
		SiteName       respjson.Field
		Title          respjson.Field
		Twitter        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebWebScrapeMdResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebWebScrapeMdResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WebWebScrapeMdResponseMetadataAdditionalMetaUnion contains all possible
// properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type WebWebScrapeMdResponseMetadataAdditionalMetaUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u WebWebScrapeMdResponseMetadataAdditionalMetaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebWebScrapeMdResponseMetadataAdditionalMetaUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WebWebScrapeMdResponseMetadataAdditionalMetaUnion) RawJSON() string { return u.JSON.raw }

func (r *WebWebScrapeMdResponseMetadataAdditionalMetaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebWebScrapeMdResponseMetadataAlternate struct {
	// Resolved alternate URL.
	Href string `json:"href" api:"required"`
	// Language or locale for the alternate URL, when present.
	Hreflang string `json:"hreflang"`
	// Alternate resource title, when present.
	Title string `json:"title"`
	// Alternate resource MIME type, when present.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Href        respjson.Field
		Hreflang    respjson.Field
		Title       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebWebScrapeMdResponseMetadataAlternate) RawJSON() string { return r.JSON.raw }
func (r *WebWebScrapeMdResponseMetadataAlternate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WebWebScrapeMdResponseMetadataOpenGraphUnion contains all possible properties
// and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type WebWebScrapeMdResponseMetadataOpenGraphUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u WebWebScrapeMdResponseMetadataOpenGraphUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebWebScrapeMdResponseMetadataOpenGraphUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WebWebScrapeMdResponseMetadataOpenGraphUnion) RawJSON() string { return u.JSON.raw }

func (r *WebWebScrapeMdResponseMetadataOpenGraphUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WebWebScrapeMdResponseMetadataTwitterUnion contains all possible properties and
// values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type WebWebScrapeMdResponseMetadataTwitterUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u WebWebScrapeMdResponseMetadataTwitterUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebWebScrapeMdResponseMetadataTwitterUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WebWebScrapeMdResponseMetadataTwitterUnion) RawJSON() string { return u.JSON.raw }

func (r *WebWebScrapeMdResponseMetadataTwitterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebWebScrapeMdResponseActionsApplied struct {
	Instruction string `json:"instruction" api:"required"`
	// Applied means the requested page state was visibly verified. Failed means it was
	// not verified. Skipped means it was not attempted.
	//
	// Any of "applied", "failed", "skipped".
	Status string `json:"status" api:"required"`
	// Visible page evidence used to verify an applied action.
	CompletionEvidence string  `json:"completionEvidence"`
	DurationMs         float64 `json:"durationMs"`
	Error              string  `json:"error"`
	Method             string  `json:"method"`
	TargetDescription  string  `json:"targetDescription"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Instruction        respjson.Field
		Status             respjson.Field
		CompletionEvidence respjson.Field
		DurationMs         respjson.Field
		Error              respjson.Field
		Method             respjson.Field
		TargetDescription  respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebWebScrapeMdResponseActionsApplied) RawJSON() string { return r.JSON.raw }
func (r *WebWebScrapeMdResponseActionsApplied) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the API key used for the request. Included in every response
// whenever a valid API key is provided, even when the response status is not 200.
type WebWebScrapeMdResponseKeyMetadata struct {
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
func (r WebWebScrapeMdResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebWebScrapeMdResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebWebScrapeSitemapResponse struct {
	// The normalized domain that was crawled
	Domain string `json:"domain" api:"required"`
	// Metadata about the sitemap crawl operation
	Meta WebWebScrapeSitemapResponseMeta `json:"meta" api:"required"`
	// Indicates success
	//
	// Any of true.
	Success bool `json:"success" api:"required"`
	// Array of discovered page URLs from the sitemap (max 500)
	URLs []string `json:"urls" api:"required"`
	// Metadata about the API key used for the request. Included in every response
	// whenever a valid API key is provided, even when the response status is not 200.
	KeyMetadata WebWebScrapeSitemapResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		Meta        respjson.Field
		Success     respjson.Field
		URLs        respjson.Field
		KeyMetadata respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebWebScrapeSitemapResponse) RawJSON() string { return r.JSON.raw }
func (r *WebWebScrapeSitemapResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the sitemap crawl operation
type WebWebScrapeSitemapResponseMeta struct {
	// Number of errors encountered during crawling
	Errors int64 `json:"errors" api:"required"`
	// Total number of sitemap files discovered
	SitemapsDiscovered int64 `json:"sitemapsDiscovered" api:"required"`
	// Number of sitemap files successfully fetched and parsed
	SitemapsFetched int64 `json:"sitemapsFetched" api:"required"`
	// Number of sitemap files skipped (due to errors, timeouts, or limits)
	SitemapsSkipped int64 `json:"sitemapsSkipped" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Errors             respjson.Field
		SitemapsDiscovered respjson.Field
		SitemapsFetched    respjson.Field
		SitemapsSkipped    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebWebScrapeSitemapResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *WebWebScrapeSitemapResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the API key used for the request. Included in every response
// whenever a valid API key is provided, even when the response status is not 200.
type WebWebScrapeSitemapResponseKeyMetadata struct {
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
func (r WebWebScrapeSitemapResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebWebScrapeSitemapResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebExtractParams struct {
	// JSON Schema for the returned data object. TypeScript Zod users can pass a JSON
	// Schema generated from a Zod object; Python users can pass the equivalent JSON
	// Schema object.
	Schema map[string]any `json:"schema,omitzero" api:"required"`
	// The starting website URL to crawl and extract from. Must include http:// or
	// https://.
	URL string `json:"url" api:"required" format:"uri"`
	// When true, every returned value must be grounded in facts stated on the page;
	// fields that cannot be supported by the page are returned as null/empty. When
	// false (default), the model may make reasonable inferences and derivations from
	// the page content (e.g. ideal customer, competitor analysis, recommendations)
	// while keeping verifiable specifics (names, quotes, URLs, dates, metrics)
	// faithful to the source.
	FactCheck param.Opt[bool] `json:"factCheck,omitzero"`
	// When true, follow links on subdomains of the starting URL's domain.
	FollowSubdomains param.Opt[bool] `json:"followSubdomains,omitzero"`
	// When true, iframe contents are included in Markdown before extraction.
	IncludeFrames param.Opt[bool] `json:"includeFrames,omitzero"`
	// Optional extraction guidance, such as which facts to prioritize or how to
	// interpret fields in the schema.
	Instructions param.Opt[string] `json:"instructions,omitzero"`
	// Return cached scrape results if a prior scrape for the same parameters is
	// younger than this many milliseconds. Defaults to 7 days (604800000 ms).
	MaxAgeMs param.Opt[int64] `json:"maxAgeMs,omitzero"`
	// Optional maximum link depth from the starting URL (0 = only the starting page).
	// If omitted, there is no crawl depth limit.
	MaxDepth param.Opt[int64] `json:"maxDepth,omitzero"`
	// Maximum number of pages to analyze for extraction. Hard cap: 50. Defaults to 5.
	MaxPages param.Opt[int64] `json:"maxPages,omitzero"`
	// When true, waits briefly for CSS and transition animations to settle before
	// extracting each crawled page. Defaults to false. This adds a bit of latency in
	// exchange for more stable output on animated pages.
	SettleAnimations param.Opt[bool] `json:"settleAnimations,omitzero"`
	// Soft time budget for the crawl in milliseconds. Min: 10000 (10s). Max: 110000
	// (110s). Default: 80000 (80s).
	StopAfterMs param.Opt[int64] `json:"stopAfterMs,omitzero"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `json:"timeoutMS,omitzero"`
	// Optional browser wait time in milliseconds after initial page load for each
	// crawled page.
	WaitForMs param.Opt[int64]    `json:"waitForMs,omitzero"`
	Pdf       WebExtractParamsPdf `json:"pdf,omitzero"`
	// Optional tags for tracking usage. Up to 20 tags, each 1 to 50 characters.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r WebExtractParams) MarshalJSON() (data []byte, err error) {
	type shadow WebExtractParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebExtractParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebExtractParamsPdf struct {
	// Last 1-based PDF page to parse. Must be greater than or equal to start when both
	// are provided.
	End param.Opt[int64] `json:"end,omitzero"`
	// When true, PDF pages are fetched and parsed. When false, PDF pages are skipped.
	ShouldParse param.Opt[bool] `json:"shouldParse,omitzero"`
	// First 1-based PDF page to parse.
	Start param.Opt[int64] `json:"start,omitzero"`
	paramObj
}

func (r WebExtractParamsPdf) MarshalJSON() (data []byte, err error) {
	type shadow WebExtractParamsPdf
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebExtractParamsPdf) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebExtractCompetitorsParams struct {
	// Company domain to analyze, such as `stripe.com`. Full http(s) URLs are accepted
	// and normalized to their domain.
	Domain string `query:"domain" api:"required" json:"-"`
	// Exact number of direct competitors to return. Defaults to 5.
	NumCompetitors param.Opt[int64] `query:"numCompetitors,omitzero" json:"-"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `query:"timeoutMS,omitzero" json:"-"`
	// Optional comma-separated caller-defined tags for tracking this request. Tags are
	// recorded on the request's usage log and can be used to filter usage on the
	// dashboard usage page. Up to 20 tags, each 1-50 characters.
	Tags []string `query:"tags,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebExtractCompetitorsParams]'s query parameters as
// `url.Values`.
func (r WebExtractCompetitorsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebExtractFontsParams struct {
	// Maximum age in milliseconds for cached brand data before the API performs a hard
	// refresh. Defaults to 3 months (7776000000 ms). Values below 1 day (86400000 ms)
	// are clamped to 1 day; values above 1 year (31536000000 ms) are clamped to 1
	// year.
	MaxAgeMs param.Opt[int64] `query:"maxAgeMs,omitzero" json:"-"`
	// A specific URL to fetch fonts from directly, bypassing domain resolution (e.g.,
	// 'https://example.com/design-system'). When provided, fonts are extracted from
	// this exact URL. You must provide either 'domain' or 'directUrl', but not both.
	DirectURL param.Opt[string] `query:"directUrl,omitzero" format:"uri" json:"-"`
	// Domain name to extract fonts from (e.g., 'example.com', 'google.com'). The
	// domain will be automatically normalized and validated. You must provide either
	// 'domain' or 'directUrl', but not both.
	Domain param.Opt[string] `query:"domain,omitzero" json:"-"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `query:"timeoutMS,omitzero" json:"-"`
	// Optional comma-separated caller-defined tags for tracking this request. Tags are
	// recorded on the request's usage log and can be used to filter usage on the
	// dashboard usage page. Up to 20 tags, each 1-50 characters.
	Tags []string `query:"tags,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebExtractFontsParams]'s query parameters as `url.Values`.
func (r WebExtractFontsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebExtractStyleguideParams struct {
	// Maximum age in milliseconds for cached brand data before the API performs a hard
	// refresh. Defaults to 3 months (7776000000 ms). Values below 1 day (86400000 ms)
	// are clamped to 1 day; values above 1 year (31536000000 ms) are clamped to 1
	// year.
	MaxAgeMs param.Opt[int64] `query:"maxAgeMs,omitzero" json:"-"`
	// A specific URL to fetch the styleguide from directly, bypassing domain
	// resolution (e.g., 'https://example.com/design-system'). When provided, the
	// styleguide is extracted from this exact URL. You must provide either 'domain' or
	// 'directUrl', but not both.
	DirectURL param.Opt[string] `query:"directUrl,omitzero" format:"uri" json:"-"`
	// Domain name to extract styleguide from (e.g., 'example.com', 'google.com'). The
	// domain will be automatically normalized and validated. You must provide either
	// 'domain' or 'directUrl', but not both.
	Domain param.Opt[string] `query:"domain,omitzero" json:"-"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `query:"timeoutMS,omitzero" json:"-"`
	// Optional browser color scheme to emulate for websites that respond to
	// prefers-color-scheme. This value is part of the styleguide cache key.
	//
	// Any of "light", "dark".
	ColorScheme WebExtractStyleguideParamsColorScheme `query:"colorScheme,omitzero" json:"-"`
	// Optional comma-separated caller-defined tags for tracking this request. Tags are
	// recorded on the request's usage log and can be used to filter usage on the
	// dashboard usage page. Up to 20 tags, each 1-50 characters.
	Tags []string `query:"tags,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebExtractStyleguideParams]'s query parameters as
// `url.Values`.
func (r WebExtractStyleguideParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Optional browser color scheme to emulate for websites that respond to
// prefers-color-scheme. This value is part of the styleguide cache key.
type WebExtractStyleguideParamsColorScheme string

const (
	WebExtractStyleguideParamsColorSchemeLight WebExtractStyleguideParamsColorScheme = "light"
	WebExtractStyleguideParamsColorSchemeDark  WebExtractStyleguideParamsColorScheme = "dark"
)

type WebScreenshotParams struct {
	// Return a cached screenshot if a prior screenshot for the same parameters exists
	// and is younger than this many milliseconds. Defaults to 1 day (86400000 ms) when
	// omitted. Max is 30 days (2592000000 ms). Set to 0 to always capture fresh.
	MaxAgeMs param.Opt[int64] `query:"maxAgeMs,omitzero" json:"-"`
	// Optional vertical scroll offset in pixels for capturing a long page in
	// viewport-sized chunks. When provided, the full page is captured once and the
	// returned image is the viewport-sized slice that begins at this Y offset (e.g.
	// request scrollOffset=0, then 1080, then 2160 to walk a 1920x1080 landing page
	// top to bottom). The final slice may be shorter than the viewport height. Takes
	// precedence over fullScreenshot. Max: 100000.
	ScrollOffset param.Opt[int64] `query:"scrollOffset,omitzero" json:"-"`
	// Optional browser wait time in milliseconds after initial page load before taking
	// the screenshot. Min: 0. Max: 30000 (30 seconds). Defaults to 3000 ms when
	// omitted.
	WaitForMs param.Opt[int64] `query:"waitForMs,omitzero" json:"-"`
	// A specific URL to screenshot directly, bypassing domain resolution (e.g.,
	// 'https://example.com/pricing'). When provided, the screenshot is taken of this
	// exact URL. You must provide either 'domain' or 'directUrl', but not both.
	DirectURL param.Opt[string] `query:"directUrl,omitzero" format:"uri" json:"-"`
	// Domain name to take screenshot of (e.g., 'example.com', 'google.com'). The
	// domain will be automatically normalized and validated. You must provide either
	// 'domain' or 'directUrl', but not both.
	Domain param.Opt[string] `query:"domain,omitzero" json:"-"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `query:"timeoutMS,omitzero" json:"-"`
	// Optional parameter to choose the site's visual theme in the screenshot. Use
	// 'light' or 'dark' when the site offers both appearances.
	//
	// Any of "light", "dark".
	ColorScheme WebScreenshotParamsColorScheme `query:"colorScheme,omitzero" json:"-"`
	// Fetch the target page through a residential proxy in this country (ISO 3166-1
	// alpha-2).
	//
	// Any of "ad", "ae", "af", "ag", "ai", "al", "am", "ao", "ar", "at", "au", "aw",
	// "az", "ba", "bb", "bd", "be", "bf", "bg", "bh", "bi", "bj", "bm", "bn", "bo",
	// "bq", "br", "bs", "bw", "by", "bz", "ca", "cd", "cf", "cg", "ch", "ci", "cl",
	// "cm", "cn", "co", "cr", "cv", "cw", "cy", "cz", "de", "dj", "dk", "dm", "do",
	// "dz", "ec", "ee", "eg", "es", "et", "fi", "fj", "fr", "ga", "gb", "gd", "ge",
	// "gf", "gg", "gh", "gm", "gn", "gp", "gq", "gr", "gt", "gu", "gw", "gy", "hk",
	// "hn", "hr", "ht", "hu", "id", "ie", "il", "im", "in", "iq", "ir", "is", "it",
	// "je", "jm", "jo", "jp", "ke", "kg", "kh", "kn", "kr", "kw", "ky", "kz", "la",
	// "lb", "lc", "lk", "lr", "ls", "lt", "lu", "lv", "ly", "ma", "mc", "md", "me",
	// "mf", "mg", "mk", "ml", "mm", "mn", "mo", "mq", "mr", "mt", "mu", "mv", "mw",
	// "mx", "my", "mz", "na", "nc", "ne", "ng", "ni", "nl", "no", "np", "nz", "om",
	// "pa", "pe", "pf", "pg", "ph", "pk", "pl", "pr", "ps", "pt", "py", "qa", "re",
	// "ro", "rs", "ru", "rw", "sa", "sc", "sd", "se", "sg", "si", "sk", "sl", "sm",
	// "sn", "so", "sr", "ss", "st", "sv", "sx", "sy", "sz", "tc", "td", "tg", "th",
	// "tj", "tl", "tm", "tn", "tr", "tt", "tw", "tz", "ua", "ug", "us", "uy", "uz",
	// "vc", "ve", "vg", "vi", "vn", "ye", "yt", "za", "zm", "zw".
	Country WebScreenshotParamsCountry `query:"country,omitzero" json:"-"`
	// Optional parameter to determine screenshot type. If 'true', takes a full page
	// screenshot capturing all content. If 'false' or not provided, takes a viewport
	// screenshot (standard browser view).
	//
	// Any of "true", "false".
	FullScreenshot WebScreenshotParamsFullScreenshot `query:"fullScreenshot,omitzero" json:"-"`
	// Optional parameter to control cookie/consent popup handling. If 'true', we
	// dismiss cookie banner before capture. If 'false' or not provided, captures the
	// page without that step.
	HandleCookiePopup WebScreenshotParamsHandleCookiePopupUnion `query:"handleCookiePopup,omitzero" json:"-"`
	// Optional parameter to specify which page type to screenshot. If provided, the
	// system will scrape the domain's links and use heuristics to find the most
	// appropriate URL for the specified page type (30 supported languages). If not
	// provided, screenshots the main domain landing page. Only applicable when using
	// 'domain', not 'directUrl'.
	//
	// Any of "login", "signup", "blog", "careers", "pricing", "terms", "privacy",
	// "contact".
	Page WebScreenshotParamsPage `query:"page,omitzero" json:"-"`
	// Optional comma-separated caller-defined tags for tracking this request. Tags are
	// recorded on the request's usage log and can be used to filter usage on the
	// dashboard usage page. Up to 20 tags, each 1-50 characters.
	Tags []string `query:"tags,omitzero" json:"-"`
	// Optional browser viewport dimensions for the screenshot. Defaults to 1920x1080.
	Viewport WebScreenshotParamsViewport `query:"viewport,omitzero" json:"-"`
	// Set to enabled to bypass shared caches and omit request and response content
	// from retained usage logs. Requires zero data retention to be enabled for your
	// organization (contact support@context.dev), otherwise the request fails with
	// ZDR_NOT_ENABLED. Successful ZDR responses include X-Context-ZDR: true.
	//
	// Any of "enabled", "disabled".
	Zdr WebScreenshotParamsZdr `query:"zdr,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebScreenshotParams]'s query parameters as `url.Values`.
func (r WebScreenshotParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Optional parameter to choose the site's visual theme in the screenshot. Use
// 'light' or 'dark' when the site offers both appearances.
type WebScreenshotParamsColorScheme string

const (
	WebScreenshotParamsColorSchemeLight WebScreenshotParamsColorScheme = "light"
	WebScreenshotParamsColorSchemeDark  WebScreenshotParamsColorScheme = "dark"
)

// Fetch the target page through a residential proxy in this country (ISO 3166-1
// alpha-2).
type WebScreenshotParamsCountry string

const (
	WebScreenshotParamsCountryAd WebScreenshotParamsCountry = "ad"
	WebScreenshotParamsCountryAe WebScreenshotParamsCountry = "ae"
	WebScreenshotParamsCountryAf WebScreenshotParamsCountry = "af"
	WebScreenshotParamsCountryAg WebScreenshotParamsCountry = "ag"
	WebScreenshotParamsCountryAI WebScreenshotParamsCountry = "ai"
	WebScreenshotParamsCountryAl WebScreenshotParamsCountry = "al"
	WebScreenshotParamsCountryAm WebScreenshotParamsCountry = "am"
	WebScreenshotParamsCountryAo WebScreenshotParamsCountry = "ao"
	WebScreenshotParamsCountryAr WebScreenshotParamsCountry = "ar"
	WebScreenshotParamsCountryAt WebScreenshotParamsCountry = "at"
	WebScreenshotParamsCountryAu WebScreenshotParamsCountry = "au"
	WebScreenshotParamsCountryAw WebScreenshotParamsCountry = "aw"
	WebScreenshotParamsCountryAz WebScreenshotParamsCountry = "az"
	WebScreenshotParamsCountryBa WebScreenshotParamsCountry = "ba"
	WebScreenshotParamsCountryBb WebScreenshotParamsCountry = "bb"
	WebScreenshotParamsCountryBd WebScreenshotParamsCountry = "bd"
	WebScreenshotParamsCountryBe WebScreenshotParamsCountry = "be"
	WebScreenshotParamsCountryBf WebScreenshotParamsCountry = "bf"
	WebScreenshotParamsCountryBg WebScreenshotParamsCountry = "bg"
	WebScreenshotParamsCountryBh WebScreenshotParamsCountry = "bh"
	WebScreenshotParamsCountryBi WebScreenshotParamsCountry = "bi"
	WebScreenshotParamsCountryBj WebScreenshotParamsCountry = "bj"
	WebScreenshotParamsCountryBm WebScreenshotParamsCountry = "bm"
	WebScreenshotParamsCountryBn WebScreenshotParamsCountry = "bn"
	WebScreenshotParamsCountryBo WebScreenshotParamsCountry = "bo"
	WebScreenshotParamsCountryBq WebScreenshotParamsCountry = "bq"
	WebScreenshotParamsCountryBr WebScreenshotParamsCountry = "br"
	WebScreenshotParamsCountryBs WebScreenshotParamsCountry = "bs"
	WebScreenshotParamsCountryBw WebScreenshotParamsCountry = "bw"
	WebScreenshotParamsCountryBy WebScreenshotParamsCountry = "by"
	WebScreenshotParamsCountryBz WebScreenshotParamsCountry = "bz"
	WebScreenshotParamsCountryCa WebScreenshotParamsCountry = "ca"
	WebScreenshotParamsCountryCd WebScreenshotParamsCountry = "cd"
	WebScreenshotParamsCountryCf WebScreenshotParamsCountry = "cf"
	WebScreenshotParamsCountryCg WebScreenshotParamsCountry = "cg"
	WebScreenshotParamsCountryCh WebScreenshotParamsCountry = "ch"
	WebScreenshotParamsCountryCi WebScreenshotParamsCountry = "ci"
	WebScreenshotParamsCountryCl WebScreenshotParamsCountry = "cl"
	WebScreenshotParamsCountryCm WebScreenshotParamsCountry = "cm"
	WebScreenshotParamsCountryCn WebScreenshotParamsCountry = "cn"
	WebScreenshotParamsCountryCo WebScreenshotParamsCountry = "co"
	WebScreenshotParamsCountryCr WebScreenshotParamsCountry = "cr"
	WebScreenshotParamsCountryCv WebScreenshotParamsCountry = "cv"
	WebScreenshotParamsCountryCw WebScreenshotParamsCountry = "cw"
	WebScreenshotParamsCountryCy WebScreenshotParamsCountry = "cy"
	WebScreenshotParamsCountryCz WebScreenshotParamsCountry = "cz"
	WebScreenshotParamsCountryDe WebScreenshotParamsCountry = "de"
	WebScreenshotParamsCountryDj WebScreenshotParamsCountry = "dj"
	WebScreenshotParamsCountryDk WebScreenshotParamsCountry = "dk"
	WebScreenshotParamsCountryDm WebScreenshotParamsCountry = "dm"
	WebScreenshotParamsCountryDo WebScreenshotParamsCountry = "do"
	WebScreenshotParamsCountryDz WebScreenshotParamsCountry = "dz"
	WebScreenshotParamsCountryEc WebScreenshotParamsCountry = "ec"
	WebScreenshotParamsCountryEe WebScreenshotParamsCountry = "ee"
	WebScreenshotParamsCountryEg WebScreenshotParamsCountry = "eg"
	WebScreenshotParamsCountryEs WebScreenshotParamsCountry = "es"
	WebScreenshotParamsCountryEt WebScreenshotParamsCountry = "et"
	WebScreenshotParamsCountryFi WebScreenshotParamsCountry = "fi"
	WebScreenshotParamsCountryFj WebScreenshotParamsCountry = "fj"
	WebScreenshotParamsCountryFr WebScreenshotParamsCountry = "fr"
	WebScreenshotParamsCountryGa WebScreenshotParamsCountry = "ga"
	WebScreenshotParamsCountryGB WebScreenshotParamsCountry = "gb"
	WebScreenshotParamsCountryGd WebScreenshotParamsCountry = "gd"
	WebScreenshotParamsCountryGe WebScreenshotParamsCountry = "ge"
	WebScreenshotParamsCountryGf WebScreenshotParamsCountry = "gf"
	WebScreenshotParamsCountryGg WebScreenshotParamsCountry = "gg"
	WebScreenshotParamsCountryGh WebScreenshotParamsCountry = "gh"
	WebScreenshotParamsCountryGm WebScreenshotParamsCountry = "gm"
	WebScreenshotParamsCountryGn WebScreenshotParamsCountry = "gn"
	WebScreenshotParamsCountryGp WebScreenshotParamsCountry = "gp"
	WebScreenshotParamsCountryGq WebScreenshotParamsCountry = "gq"
	WebScreenshotParamsCountryGr WebScreenshotParamsCountry = "gr"
	WebScreenshotParamsCountryGt WebScreenshotParamsCountry = "gt"
	WebScreenshotParamsCountryGu WebScreenshotParamsCountry = "gu"
	WebScreenshotParamsCountryGw WebScreenshotParamsCountry = "gw"
	WebScreenshotParamsCountryGy WebScreenshotParamsCountry = "gy"
	WebScreenshotParamsCountryHk WebScreenshotParamsCountry = "hk"
	WebScreenshotParamsCountryHn WebScreenshotParamsCountry = "hn"
	WebScreenshotParamsCountryHr WebScreenshotParamsCountry = "hr"
	WebScreenshotParamsCountryHt WebScreenshotParamsCountry = "ht"
	WebScreenshotParamsCountryHu WebScreenshotParamsCountry = "hu"
	WebScreenshotParamsCountryID WebScreenshotParamsCountry = "id"
	WebScreenshotParamsCountryIe WebScreenshotParamsCountry = "ie"
	WebScreenshotParamsCountryIl WebScreenshotParamsCountry = "il"
	WebScreenshotParamsCountryIm WebScreenshotParamsCountry = "im"
	WebScreenshotParamsCountryIn WebScreenshotParamsCountry = "in"
	WebScreenshotParamsCountryIq WebScreenshotParamsCountry = "iq"
	WebScreenshotParamsCountryIr WebScreenshotParamsCountry = "ir"
	WebScreenshotParamsCountryIs WebScreenshotParamsCountry = "is"
	WebScreenshotParamsCountryIt WebScreenshotParamsCountry = "it"
	WebScreenshotParamsCountryJe WebScreenshotParamsCountry = "je"
	WebScreenshotParamsCountryJm WebScreenshotParamsCountry = "jm"
	WebScreenshotParamsCountryJo WebScreenshotParamsCountry = "jo"
	WebScreenshotParamsCountryJp WebScreenshotParamsCountry = "jp"
	WebScreenshotParamsCountryKe WebScreenshotParamsCountry = "ke"
	WebScreenshotParamsCountryKg WebScreenshotParamsCountry = "kg"
	WebScreenshotParamsCountryKh WebScreenshotParamsCountry = "kh"
	WebScreenshotParamsCountryKn WebScreenshotParamsCountry = "kn"
	WebScreenshotParamsCountryKr WebScreenshotParamsCountry = "kr"
	WebScreenshotParamsCountryKw WebScreenshotParamsCountry = "kw"
	WebScreenshotParamsCountryKy WebScreenshotParamsCountry = "ky"
	WebScreenshotParamsCountryKz WebScreenshotParamsCountry = "kz"
	WebScreenshotParamsCountryLa WebScreenshotParamsCountry = "la"
	WebScreenshotParamsCountryLb WebScreenshotParamsCountry = "lb"
	WebScreenshotParamsCountryLc WebScreenshotParamsCountry = "lc"
	WebScreenshotParamsCountryLk WebScreenshotParamsCountry = "lk"
	WebScreenshotParamsCountryLr WebScreenshotParamsCountry = "lr"
	WebScreenshotParamsCountryLs WebScreenshotParamsCountry = "ls"
	WebScreenshotParamsCountryLt WebScreenshotParamsCountry = "lt"
	WebScreenshotParamsCountryLu WebScreenshotParamsCountry = "lu"
	WebScreenshotParamsCountryLv WebScreenshotParamsCountry = "lv"
	WebScreenshotParamsCountryLy WebScreenshotParamsCountry = "ly"
	WebScreenshotParamsCountryMa WebScreenshotParamsCountry = "ma"
	WebScreenshotParamsCountryMc WebScreenshotParamsCountry = "mc"
	WebScreenshotParamsCountryMd WebScreenshotParamsCountry = "md"
	WebScreenshotParamsCountryMe WebScreenshotParamsCountry = "me"
	WebScreenshotParamsCountryMf WebScreenshotParamsCountry = "mf"
	WebScreenshotParamsCountryMg WebScreenshotParamsCountry = "mg"
	WebScreenshotParamsCountryMk WebScreenshotParamsCountry = "mk"
	WebScreenshotParamsCountryMl WebScreenshotParamsCountry = "ml"
	WebScreenshotParamsCountryMm WebScreenshotParamsCountry = "mm"
	WebScreenshotParamsCountryMn WebScreenshotParamsCountry = "mn"
	WebScreenshotParamsCountryMo WebScreenshotParamsCountry = "mo"
	WebScreenshotParamsCountryMq WebScreenshotParamsCountry = "mq"
	WebScreenshotParamsCountryMr WebScreenshotParamsCountry = "mr"
	WebScreenshotParamsCountryMt WebScreenshotParamsCountry = "mt"
	WebScreenshotParamsCountryMu WebScreenshotParamsCountry = "mu"
	WebScreenshotParamsCountryMv WebScreenshotParamsCountry = "mv"
	WebScreenshotParamsCountryMw WebScreenshotParamsCountry = "mw"
	WebScreenshotParamsCountryMx WebScreenshotParamsCountry = "mx"
	WebScreenshotParamsCountryMy WebScreenshotParamsCountry = "my"
	WebScreenshotParamsCountryMz WebScreenshotParamsCountry = "mz"
	WebScreenshotParamsCountryNa WebScreenshotParamsCountry = "na"
	WebScreenshotParamsCountryNc WebScreenshotParamsCountry = "nc"
	WebScreenshotParamsCountryNe WebScreenshotParamsCountry = "ne"
	WebScreenshotParamsCountryNg WebScreenshotParamsCountry = "ng"
	WebScreenshotParamsCountryNi WebScreenshotParamsCountry = "ni"
	WebScreenshotParamsCountryNl WebScreenshotParamsCountry = "nl"
	WebScreenshotParamsCountryNo WebScreenshotParamsCountry = "no"
	WebScreenshotParamsCountryNp WebScreenshotParamsCountry = "np"
	WebScreenshotParamsCountryNz WebScreenshotParamsCountry = "nz"
	WebScreenshotParamsCountryOm WebScreenshotParamsCountry = "om"
	WebScreenshotParamsCountryPa WebScreenshotParamsCountry = "pa"
	WebScreenshotParamsCountryPe WebScreenshotParamsCountry = "pe"
	WebScreenshotParamsCountryPf WebScreenshotParamsCountry = "pf"
	WebScreenshotParamsCountryPg WebScreenshotParamsCountry = "pg"
	WebScreenshotParamsCountryPh WebScreenshotParamsCountry = "ph"
	WebScreenshotParamsCountryPk WebScreenshotParamsCountry = "pk"
	WebScreenshotParamsCountryPl WebScreenshotParamsCountry = "pl"
	WebScreenshotParamsCountryPr WebScreenshotParamsCountry = "pr"
	WebScreenshotParamsCountryPs WebScreenshotParamsCountry = "ps"
	WebScreenshotParamsCountryPt WebScreenshotParamsCountry = "pt"
	WebScreenshotParamsCountryPy WebScreenshotParamsCountry = "py"
	WebScreenshotParamsCountryQa WebScreenshotParamsCountry = "qa"
	WebScreenshotParamsCountryRe WebScreenshotParamsCountry = "re"
	WebScreenshotParamsCountryRo WebScreenshotParamsCountry = "ro"
	WebScreenshotParamsCountryRs WebScreenshotParamsCountry = "rs"
	WebScreenshotParamsCountryRu WebScreenshotParamsCountry = "ru"
	WebScreenshotParamsCountryRw WebScreenshotParamsCountry = "rw"
	WebScreenshotParamsCountrySa WebScreenshotParamsCountry = "sa"
	WebScreenshotParamsCountrySc WebScreenshotParamsCountry = "sc"
	WebScreenshotParamsCountrySd WebScreenshotParamsCountry = "sd"
	WebScreenshotParamsCountrySe WebScreenshotParamsCountry = "se"
	WebScreenshotParamsCountrySg WebScreenshotParamsCountry = "sg"
	WebScreenshotParamsCountrySi WebScreenshotParamsCountry = "si"
	WebScreenshotParamsCountrySk WebScreenshotParamsCountry = "sk"
	WebScreenshotParamsCountrySl WebScreenshotParamsCountry = "sl"
	WebScreenshotParamsCountrySm WebScreenshotParamsCountry = "sm"
	WebScreenshotParamsCountrySn WebScreenshotParamsCountry = "sn"
	WebScreenshotParamsCountrySo WebScreenshotParamsCountry = "so"
	WebScreenshotParamsCountrySr WebScreenshotParamsCountry = "sr"
	WebScreenshotParamsCountrySS WebScreenshotParamsCountry = "ss"
	WebScreenshotParamsCountrySt WebScreenshotParamsCountry = "st"
	WebScreenshotParamsCountrySv WebScreenshotParamsCountry = "sv"
	WebScreenshotParamsCountrySx WebScreenshotParamsCountry = "sx"
	WebScreenshotParamsCountrySy WebScreenshotParamsCountry = "sy"
	WebScreenshotParamsCountrySz WebScreenshotParamsCountry = "sz"
	WebScreenshotParamsCountryTc WebScreenshotParamsCountry = "tc"
	WebScreenshotParamsCountryTd WebScreenshotParamsCountry = "td"
	WebScreenshotParamsCountryTg WebScreenshotParamsCountry = "tg"
	WebScreenshotParamsCountryTh WebScreenshotParamsCountry = "th"
	WebScreenshotParamsCountryTj WebScreenshotParamsCountry = "tj"
	WebScreenshotParamsCountryTl WebScreenshotParamsCountry = "tl"
	WebScreenshotParamsCountryTm WebScreenshotParamsCountry = "tm"
	WebScreenshotParamsCountryTn WebScreenshotParamsCountry = "tn"
	WebScreenshotParamsCountryTr WebScreenshotParamsCountry = "tr"
	WebScreenshotParamsCountryTt WebScreenshotParamsCountry = "tt"
	WebScreenshotParamsCountryTw WebScreenshotParamsCountry = "tw"
	WebScreenshotParamsCountryTz WebScreenshotParamsCountry = "tz"
	WebScreenshotParamsCountryUa WebScreenshotParamsCountry = "ua"
	WebScreenshotParamsCountryUg WebScreenshotParamsCountry = "ug"
	WebScreenshotParamsCountryUs WebScreenshotParamsCountry = "us"
	WebScreenshotParamsCountryUy WebScreenshotParamsCountry = "uy"
	WebScreenshotParamsCountryUz WebScreenshotParamsCountry = "uz"
	WebScreenshotParamsCountryVc WebScreenshotParamsCountry = "vc"
	WebScreenshotParamsCountryVe WebScreenshotParamsCountry = "ve"
	WebScreenshotParamsCountryVg WebScreenshotParamsCountry = "vg"
	WebScreenshotParamsCountryVi WebScreenshotParamsCountry = "vi"
	WebScreenshotParamsCountryVn WebScreenshotParamsCountry = "vn"
	WebScreenshotParamsCountryYe WebScreenshotParamsCountry = "ye"
	WebScreenshotParamsCountryYt WebScreenshotParamsCountry = "yt"
	WebScreenshotParamsCountryZa WebScreenshotParamsCountry = "za"
	WebScreenshotParamsCountryZm WebScreenshotParamsCountry = "zm"
	WebScreenshotParamsCountryZw WebScreenshotParamsCountry = "zw"
)

// Optional parameter to determine screenshot type. If 'true', takes a full page
// screenshot capturing all content. If 'false' or not provided, takes a viewport
// screenshot (standard browser view).
type WebScreenshotParamsFullScreenshot string

const (
	WebScreenshotParamsFullScreenshotTrue  WebScreenshotParamsFullScreenshot = "true"
	WebScreenshotParamsFullScreenshotFalse WebScreenshotParamsFullScreenshot = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebScreenshotParamsHandleCookiePopupUnion struct {
	OfBool param.Opt[bool] `query:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWebScreenshotsHandleCookiePopupString)
	OfWebScreenshotsHandleCookiePopupString param.Opt[string] `query:",omitzero,inline"`
	paramUnion
}

type WebScreenshotParamsHandleCookiePopupString string

const (
	WebScreenshotParamsHandleCookiePopupStringTrue  WebScreenshotParamsHandleCookiePopupString = "true"
	WebScreenshotParamsHandleCookiePopupStringFalse WebScreenshotParamsHandleCookiePopupString = "false"
)

// Optional parameter to specify which page type to screenshot. If provided, the
// system will scrape the domain's links and use heuristics to find the most
// appropriate URL for the specified page type (30 supported languages). If not
// provided, screenshots the main domain landing page. Only applicable when using
// 'domain', not 'directUrl'.
type WebScreenshotParamsPage string

const (
	WebScreenshotParamsPageLogin   WebScreenshotParamsPage = "login"
	WebScreenshotParamsPageSignup  WebScreenshotParamsPage = "signup"
	WebScreenshotParamsPageBlog    WebScreenshotParamsPage = "blog"
	WebScreenshotParamsPageCareers WebScreenshotParamsPage = "careers"
	WebScreenshotParamsPagePricing WebScreenshotParamsPage = "pricing"
	WebScreenshotParamsPageTerms   WebScreenshotParamsPage = "terms"
	WebScreenshotParamsPagePrivacy WebScreenshotParamsPage = "privacy"
	WebScreenshotParamsPageContact WebScreenshotParamsPage = "contact"
)

// Optional browser viewport dimensions for the screenshot. Defaults to 1920x1080.
type WebScreenshotParamsViewport struct {
	// Viewport height in pixels.
	Height param.Opt[int64] `query:"height,omitzero" json:"-"`
	// Viewport width in pixels.
	Width param.Opt[int64] `query:"width,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebScreenshotParamsViewport]'s query parameters as
// `url.Values`.
func (r WebScreenshotParamsViewport) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Set to enabled to bypass shared caches and omit request and response content
// from retained usage logs. Requires zero data retention to be enabled for your
// organization (contact support@context.dev), otherwise the request fails with
// ZDR_NOT_ENABLED. Successful ZDR responses include X-Context-ZDR: true.
type WebScreenshotParamsZdr string

const (
	WebScreenshotParamsZdrEnabled  WebScreenshotParamsZdr = "enabled"
	WebScreenshotParamsZdrDisabled WebScreenshotParamsZdr = "disabled"
)

type WebSearchParams struct {
	// Search query. Accepts natural language as well as Google-style search operators
	// such as `site:`, `-site:`, `inurl:`, `intitle:`, quoted phrases, and `OR`.
	Query string `json:"query" api:"required"`
	// Number of results to request and return (10–100). Defaults to 10.
	NumResults param.Opt[int64] `json:"numResults,omitzero"`
	// Expand the query into multiple parallel variants for broader recall.
	QueryFanout param.Opt[bool] `json:"queryFanout,omitzero"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `json:"timeoutMS,omitzero"`
	// Two-letter ISO 3166-1 alpha-2 country code to localize results to a specific
	// country (maps to Google's `gl` parameter). Example: "us", "gb", "de".
	//
	// Any of "af", "al", "dz", "as", "ad", "ao", "ai", "aq", "ag", "ar", "am", "aw",
	// "au", "at", "az", "bs", "bh", "bd", "bb", "by", "be", "bz", "bj", "bm", "bt",
	// "bo", "ba", "bw", "bv", "br", "io", "bn", "bg", "bf", "bi", "kh", "cm", "ca",
	// "cv", "ky", "cf", "td", "cl", "cn", "cx", "cc", "co", "km", "cg", "cd", "ck",
	// "cr", "ci", "hr", "cu", "cy", "cz", "dk", "dj", "dm", "do", "ec", "eg", "sv",
	// "gq", "er", "ee", "et", "fk", "fo", "fj", "fi", "fr", "gf", "pf", "tf", "ga",
	// "gm", "ge", "de", "gh", "gi", "gr", "gl", "gd", "gp", "gu", "gt", "gn", "gw",
	// "gy", "ht", "hm", "va", "hn", "hk", "hu", "is", "in", "id", "ir", "iq", "ie",
	// "il", "it", "jm", "jp", "jo", "kz", "ke", "ki", "kp", "kr", "kw", "kg", "la",
	// "lv", "lb", "ls", "lr", "ly", "li", "lt", "lu", "mo", "mk", "mg", "mw", "my",
	// "mv", "ml", "mt", "mh", "mq", "mr", "mu", "yt", "mx", "fm", "md", "mc", "mn",
	// "ms", "ma", "mz", "mm", "na", "nr", "np", "nl", "an", "nc", "nz", "ni", "ne",
	// "ng", "nu", "nf", "mp", "no", "om", "pk", "pw", "ps", "pa", "pg", "py", "pe",
	// "ph", "pn", "pl", "pt", "pr", "qa", "re", "ro", "ru", "rw", "sh", "kn", "lc",
	// "pm", "vc", "ws", "sm", "st", "sa", "sn", "rs", "sc", "sl", "sg", "sk", "si",
	// "sb", "so", "za", "gs", "es", "lk", "sd", "sr", "sj", "sz", "se", "ch", "sy",
	// "tw", "tj", "tz", "th", "tl", "tg", "tk", "to", "tt", "tn", "tr", "tm", "tc",
	// "tv", "ug", "ua", "ae", "gb", "us", "um", "uy", "uz", "vu", "ve", "vn", "vg",
	// "vi", "wf", "eh", "ye", "zm", "zw".
	Country WebSearchParamsCountry `json:"country,omitzero"`
	// Blocklist — drop results from these domains. Example: ["pinterest.com",
	// "reddit.com"].
	ExcludeDomains []string `json:"excludeDomains,omitzero"`
	// Restrict results to content published within this window.
	//
	// Any of "last_24_hours", "last_week", "last_month", "last_year".
	Freshness WebSearchParamsFreshness `json:"freshness,omitzero"`
	// Allowlist — only return results from these domains. Example: ["arxiv.org",
	// "github.com"].
	IncludeDomains []string `json:"includeDomains,omitzero"`
	// Inline Markdown scraping for each result. Set `enabled: true` to activate.
	MarkdownOptions WebSearchParamsMarkdownOptions `json:"markdownOptions,omitzero"`
	// Optional tags for tracking usage. Up to 20 tags, each 1 to 50 characters.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r WebSearchParams) MarshalJSON() (data []byte, err error) {
	type shadow WebSearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Two-letter ISO 3166-1 alpha-2 country code to localize results to a specific
// country (maps to Google's `gl` parameter). Example: "us", "gb", "de".
type WebSearchParamsCountry string

const (
	WebSearchParamsCountryAf WebSearchParamsCountry = "af"
	WebSearchParamsCountryAl WebSearchParamsCountry = "al"
	WebSearchParamsCountryDz WebSearchParamsCountry = "dz"
	WebSearchParamsCountryAs WebSearchParamsCountry = "as"
	WebSearchParamsCountryAd WebSearchParamsCountry = "ad"
	WebSearchParamsCountryAo WebSearchParamsCountry = "ao"
	WebSearchParamsCountryAI WebSearchParamsCountry = "ai"
	WebSearchParamsCountryAq WebSearchParamsCountry = "aq"
	WebSearchParamsCountryAg WebSearchParamsCountry = "ag"
	WebSearchParamsCountryAr WebSearchParamsCountry = "ar"
	WebSearchParamsCountryAm WebSearchParamsCountry = "am"
	WebSearchParamsCountryAw WebSearchParamsCountry = "aw"
	WebSearchParamsCountryAu WebSearchParamsCountry = "au"
	WebSearchParamsCountryAt WebSearchParamsCountry = "at"
	WebSearchParamsCountryAz WebSearchParamsCountry = "az"
	WebSearchParamsCountryBs WebSearchParamsCountry = "bs"
	WebSearchParamsCountryBh WebSearchParamsCountry = "bh"
	WebSearchParamsCountryBd WebSearchParamsCountry = "bd"
	WebSearchParamsCountryBb WebSearchParamsCountry = "bb"
	WebSearchParamsCountryBy WebSearchParamsCountry = "by"
	WebSearchParamsCountryBe WebSearchParamsCountry = "be"
	WebSearchParamsCountryBz WebSearchParamsCountry = "bz"
	WebSearchParamsCountryBj WebSearchParamsCountry = "bj"
	WebSearchParamsCountryBm WebSearchParamsCountry = "bm"
	WebSearchParamsCountryBt WebSearchParamsCountry = "bt"
	WebSearchParamsCountryBo WebSearchParamsCountry = "bo"
	WebSearchParamsCountryBa WebSearchParamsCountry = "ba"
	WebSearchParamsCountryBw WebSearchParamsCountry = "bw"
	WebSearchParamsCountryBv WebSearchParamsCountry = "bv"
	WebSearchParamsCountryBr WebSearchParamsCountry = "br"
	WebSearchParamsCountryIo WebSearchParamsCountry = "io"
	WebSearchParamsCountryBn WebSearchParamsCountry = "bn"
	WebSearchParamsCountryBg WebSearchParamsCountry = "bg"
	WebSearchParamsCountryBf WebSearchParamsCountry = "bf"
	WebSearchParamsCountryBi WebSearchParamsCountry = "bi"
	WebSearchParamsCountryKh WebSearchParamsCountry = "kh"
	WebSearchParamsCountryCm WebSearchParamsCountry = "cm"
	WebSearchParamsCountryCa WebSearchParamsCountry = "ca"
	WebSearchParamsCountryCv WebSearchParamsCountry = "cv"
	WebSearchParamsCountryKy WebSearchParamsCountry = "ky"
	WebSearchParamsCountryCf WebSearchParamsCountry = "cf"
	WebSearchParamsCountryTd WebSearchParamsCountry = "td"
	WebSearchParamsCountryCl WebSearchParamsCountry = "cl"
	WebSearchParamsCountryCn WebSearchParamsCountry = "cn"
	WebSearchParamsCountryCx WebSearchParamsCountry = "cx"
	WebSearchParamsCountryCc WebSearchParamsCountry = "cc"
	WebSearchParamsCountryCo WebSearchParamsCountry = "co"
	WebSearchParamsCountryKm WebSearchParamsCountry = "km"
	WebSearchParamsCountryCg WebSearchParamsCountry = "cg"
	WebSearchParamsCountryCd WebSearchParamsCountry = "cd"
	WebSearchParamsCountryCk WebSearchParamsCountry = "ck"
	WebSearchParamsCountryCr WebSearchParamsCountry = "cr"
	WebSearchParamsCountryCi WebSearchParamsCountry = "ci"
	WebSearchParamsCountryHr WebSearchParamsCountry = "hr"
	WebSearchParamsCountryCu WebSearchParamsCountry = "cu"
	WebSearchParamsCountryCy WebSearchParamsCountry = "cy"
	WebSearchParamsCountryCz WebSearchParamsCountry = "cz"
	WebSearchParamsCountryDk WebSearchParamsCountry = "dk"
	WebSearchParamsCountryDj WebSearchParamsCountry = "dj"
	WebSearchParamsCountryDm WebSearchParamsCountry = "dm"
	WebSearchParamsCountryDo WebSearchParamsCountry = "do"
	WebSearchParamsCountryEc WebSearchParamsCountry = "ec"
	WebSearchParamsCountryEg WebSearchParamsCountry = "eg"
	WebSearchParamsCountrySv WebSearchParamsCountry = "sv"
	WebSearchParamsCountryGq WebSearchParamsCountry = "gq"
	WebSearchParamsCountryEr WebSearchParamsCountry = "er"
	WebSearchParamsCountryEe WebSearchParamsCountry = "ee"
	WebSearchParamsCountryEt WebSearchParamsCountry = "et"
	WebSearchParamsCountryFk WebSearchParamsCountry = "fk"
	WebSearchParamsCountryFo WebSearchParamsCountry = "fo"
	WebSearchParamsCountryFj WebSearchParamsCountry = "fj"
	WebSearchParamsCountryFi WebSearchParamsCountry = "fi"
	WebSearchParamsCountryFr WebSearchParamsCountry = "fr"
	WebSearchParamsCountryGf WebSearchParamsCountry = "gf"
	WebSearchParamsCountryPf WebSearchParamsCountry = "pf"
	WebSearchParamsCountryTf WebSearchParamsCountry = "tf"
	WebSearchParamsCountryGa WebSearchParamsCountry = "ga"
	WebSearchParamsCountryGm WebSearchParamsCountry = "gm"
	WebSearchParamsCountryGe WebSearchParamsCountry = "ge"
	WebSearchParamsCountryDe WebSearchParamsCountry = "de"
	WebSearchParamsCountryGh WebSearchParamsCountry = "gh"
	WebSearchParamsCountryGi WebSearchParamsCountry = "gi"
	WebSearchParamsCountryGr WebSearchParamsCountry = "gr"
	WebSearchParamsCountryGl WebSearchParamsCountry = "gl"
	WebSearchParamsCountryGd WebSearchParamsCountry = "gd"
	WebSearchParamsCountryGp WebSearchParamsCountry = "gp"
	WebSearchParamsCountryGu WebSearchParamsCountry = "gu"
	WebSearchParamsCountryGt WebSearchParamsCountry = "gt"
	WebSearchParamsCountryGn WebSearchParamsCountry = "gn"
	WebSearchParamsCountryGw WebSearchParamsCountry = "gw"
	WebSearchParamsCountryGy WebSearchParamsCountry = "gy"
	WebSearchParamsCountryHt WebSearchParamsCountry = "ht"
	WebSearchParamsCountryHm WebSearchParamsCountry = "hm"
	WebSearchParamsCountryVa WebSearchParamsCountry = "va"
	WebSearchParamsCountryHn WebSearchParamsCountry = "hn"
	WebSearchParamsCountryHk WebSearchParamsCountry = "hk"
	WebSearchParamsCountryHu WebSearchParamsCountry = "hu"
	WebSearchParamsCountryIs WebSearchParamsCountry = "is"
	WebSearchParamsCountryIn WebSearchParamsCountry = "in"
	WebSearchParamsCountryID WebSearchParamsCountry = "id"
	WebSearchParamsCountryIr WebSearchParamsCountry = "ir"
	WebSearchParamsCountryIq WebSearchParamsCountry = "iq"
	WebSearchParamsCountryIe WebSearchParamsCountry = "ie"
	WebSearchParamsCountryIl WebSearchParamsCountry = "il"
	WebSearchParamsCountryIt WebSearchParamsCountry = "it"
	WebSearchParamsCountryJm WebSearchParamsCountry = "jm"
	WebSearchParamsCountryJp WebSearchParamsCountry = "jp"
	WebSearchParamsCountryJo WebSearchParamsCountry = "jo"
	WebSearchParamsCountryKz WebSearchParamsCountry = "kz"
	WebSearchParamsCountryKe WebSearchParamsCountry = "ke"
	WebSearchParamsCountryKi WebSearchParamsCountry = "ki"
	WebSearchParamsCountryKp WebSearchParamsCountry = "kp"
	WebSearchParamsCountryKr WebSearchParamsCountry = "kr"
	WebSearchParamsCountryKw WebSearchParamsCountry = "kw"
	WebSearchParamsCountryKg WebSearchParamsCountry = "kg"
	WebSearchParamsCountryLa WebSearchParamsCountry = "la"
	WebSearchParamsCountryLv WebSearchParamsCountry = "lv"
	WebSearchParamsCountryLb WebSearchParamsCountry = "lb"
	WebSearchParamsCountryLs WebSearchParamsCountry = "ls"
	WebSearchParamsCountryLr WebSearchParamsCountry = "lr"
	WebSearchParamsCountryLy WebSearchParamsCountry = "ly"
	WebSearchParamsCountryLi WebSearchParamsCountry = "li"
	WebSearchParamsCountryLt WebSearchParamsCountry = "lt"
	WebSearchParamsCountryLu WebSearchParamsCountry = "lu"
	WebSearchParamsCountryMo WebSearchParamsCountry = "mo"
	WebSearchParamsCountryMk WebSearchParamsCountry = "mk"
	WebSearchParamsCountryMg WebSearchParamsCountry = "mg"
	WebSearchParamsCountryMw WebSearchParamsCountry = "mw"
	WebSearchParamsCountryMy WebSearchParamsCountry = "my"
	WebSearchParamsCountryMv WebSearchParamsCountry = "mv"
	WebSearchParamsCountryMl WebSearchParamsCountry = "ml"
	WebSearchParamsCountryMt WebSearchParamsCountry = "mt"
	WebSearchParamsCountryMh WebSearchParamsCountry = "mh"
	WebSearchParamsCountryMq WebSearchParamsCountry = "mq"
	WebSearchParamsCountryMr WebSearchParamsCountry = "mr"
	WebSearchParamsCountryMu WebSearchParamsCountry = "mu"
	WebSearchParamsCountryYt WebSearchParamsCountry = "yt"
	WebSearchParamsCountryMx WebSearchParamsCountry = "mx"
	WebSearchParamsCountryFm WebSearchParamsCountry = "fm"
	WebSearchParamsCountryMd WebSearchParamsCountry = "md"
	WebSearchParamsCountryMc WebSearchParamsCountry = "mc"
	WebSearchParamsCountryMn WebSearchParamsCountry = "mn"
	WebSearchParamsCountryMs WebSearchParamsCountry = "ms"
	WebSearchParamsCountryMa WebSearchParamsCountry = "ma"
	WebSearchParamsCountryMz WebSearchParamsCountry = "mz"
	WebSearchParamsCountryMm WebSearchParamsCountry = "mm"
	WebSearchParamsCountryNa WebSearchParamsCountry = "na"
	WebSearchParamsCountryNr WebSearchParamsCountry = "nr"
	WebSearchParamsCountryNp WebSearchParamsCountry = "np"
	WebSearchParamsCountryNl WebSearchParamsCountry = "nl"
	WebSearchParamsCountryAn WebSearchParamsCountry = "an"
	WebSearchParamsCountryNc WebSearchParamsCountry = "nc"
	WebSearchParamsCountryNz WebSearchParamsCountry = "nz"
	WebSearchParamsCountryNi WebSearchParamsCountry = "ni"
	WebSearchParamsCountryNe WebSearchParamsCountry = "ne"
	WebSearchParamsCountryNg WebSearchParamsCountry = "ng"
	WebSearchParamsCountryNu WebSearchParamsCountry = "nu"
	WebSearchParamsCountryNf WebSearchParamsCountry = "nf"
	WebSearchParamsCountryMp WebSearchParamsCountry = "mp"
	WebSearchParamsCountryNo WebSearchParamsCountry = "no"
	WebSearchParamsCountryOm WebSearchParamsCountry = "om"
	WebSearchParamsCountryPk WebSearchParamsCountry = "pk"
	WebSearchParamsCountryPw WebSearchParamsCountry = "pw"
	WebSearchParamsCountryPs WebSearchParamsCountry = "ps"
	WebSearchParamsCountryPa WebSearchParamsCountry = "pa"
	WebSearchParamsCountryPg WebSearchParamsCountry = "pg"
	WebSearchParamsCountryPy WebSearchParamsCountry = "py"
	WebSearchParamsCountryPe WebSearchParamsCountry = "pe"
	WebSearchParamsCountryPh WebSearchParamsCountry = "ph"
	WebSearchParamsCountryPn WebSearchParamsCountry = "pn"
	WebSearchParamsCountryPl WebSearchParamsCountry = "pl"
	WebSearchParamsCountryPt WebSearchParamsCountry = "pt"
	WebSearchParamsCountryPr WebSearchParamsCountry = "pr"
	WebSearchParamsCountryQa WebSearchParamsCountry = "qa"
	WebSearchParamsCountryRe WebSearchParamsCountry = "re"
	WebSearchParamsCountryRo WebSearchParamsCountry = "ro"
	WebSearchParamsCountryRu WebSearchParamsCountry = "ru"
	WebSearchParamsCountryRw WebSearchParamsCountry = "rw"
	WebSearchParamsCountrySh WebSearchParamsCountry = "sh"
	WebSearchParamsCountryKn WebSearchParamsCountry = "kn"
	WebSearchParamsCountryLc WebSearchParamsCountry = "lc"
	WebSearchParamsCountryPm WebSearchParamsCountry = "pm"
	WebSearchParamsCountryVc WebSearchParamsCountry = "vc"
	WebSearchParamsCountryWs WebSearchParamsCountry = "ws"
	WebSearchParamsCountrySm WebSearchParamsCountry = "sm"
	WebSearchParamsCountrySt WebSearchParamsCountry = "st"
	WebSearchParamsCountrySa WebSearchParamsCountry = "sa"
	WebSearchParamsCountrySn WebSearchParamsCountry = "sn"
	WebSearchParamsCountryRs WebSearchParamsCountry = "rs"
	WebSearchParamsCountrySc WebSearchParamsCountry = "sc"
	WebSearchParamsCountrySl WebSearchParamsCountry = "sl"
	WebSearchParamsCountrySg WebSearchParamsCountry = "sg"
	WebSearchParamsCountrySk WebSearchParamsCountry = "sk"
	WebSearchParamsCountrySi WebSearchParamsCountry = "si"
	WebSearchParamsCountrySb WebSearchParamsCountry = "sb"
	WebSearchParamsCountrySo WebSearchParamsCountry = "so"
	WebSearchParamsCountryZa WebSearchParamsCountry = "za"
	WebSearchParamsCountryGs WebSearchParamsCountry = "gs"
	WebSearchParamsCountryEs WebSearchParamsCountry = "es"
	WebSearchParamsCountryLk WebSearchParamsCountry = "lk"
	WebSearchParamsCountrySd WebSearchParamsCountry = "sd"
	WebSearchParamsCountrySr WebSearchParamsCountry = "sr"
	WebSearchParamsCountrySj WebSearchParamsCountry = "sj"
	WebSearchParamsCountrySz WebSearchParamsCountry = "sz"
	WebSearchParamsCountrySe WebSearchParamsCountry = "se"
	WebSearchParamsCountryCh WebSearchParamsCountry = "ch"
	WebSearchParamsCountrySy WebSearchParamsCountry = "sy"
	WebSearchParamsCountryTw WebSearchParamsCountry = "tw"
	WebSearchParamsCountryTj WebSearchParamsCountry = "tj"
	WebSearchParamsCountryTz WebSearchParamsCountry = "tz"
	WebSearchParamsCountryTh WebSearchParamsCountry = "th"
	WebSearchParamsCountryTl WebSearchParamsCountry = "tl"
	WebSearchParamsCountryTg WebSearchParamsCountry = "tg"
	WebSearchParamsCountryTk WebSearchParamsCountry = "tk"
	WebSearchParamsCountryTo WebSearchParamsCountry = "to"
	WebSearchParamsCountryTt WebSearchParamsCountry = "tt"
	WebSearchParamsCountryTn WebSearchParamsCountry = "tn"
	WebSearchParamsCountryTr WebSearchParamsCountry = "tr"
	WebSearchParamsCountryTm WebSearchParamsCountry = "tm"
	WebSearchParamsCountryTc WebSearchParamsCountry = "tc"
	WebSearchParamsCountryTv WebSearchParamsCountry = "tv"
	WebSearchParamsCountryUg WebSearchParamsCountry = "ug"
	WebSearchParamsCountryUa WebSearchParamsCountry = "ua"
	WebSearchParamsCountryAe WebSearchParamsCountry = "ae"
	WebSearchParamsCountryGB WebSearchParamsCountry = "gb"
	WebSearchParamsCountryUs WebSearchParamsCountry = "us"
	WebSearchParamsCountryUm WebSearchParamsCountry = "um"
	WebSearchParamsCountryUy WebSearchParamsCountry = "uy"
	WebSearchParamsCountryUz WebSearchParamsCountry = "uz"
	WebSearchParamsCountryVu WebSearchParamsCountry = "vu"
	WebSearchParamsCountryVe WebSearchParamsCountry = "ve"
	WebSearchParamsCountryVn WebSearchParamsCountry = "vn"
	WebSearchParamsCountryVg WebSearchParamsCountry = "vg"
	WebSearchParamsCountryVi WebSearchParamsCountry = "vi"
	WebSearchParamsCountryWf WebSearchParamsCountry = "wf"
	WebSearchParamsCountryEh WebSearchParamsCountry = "eh"
	WebSearchParamsCountryYe WebSearchParamsCountry = "ye"
	WebSearchParamsCountryZm WebSearchParamsCountry = "zm"
	WebSearchParamsCountryZw WebSearchParamsCountry = "zw"
)

// Restrict results to content published within this window.
type WebSearchParamsFreshness string

const (
	WebSearchParamsFreshnessLast24Hours WebSearchParamsFreshness = "last_24_hours"
	WebSearchParamsFreshnessLastWeek    WebSearchParamsFreshness = "last_week"
	WebSearchParamsFreshnessLastMonth   WebSearchParamsFreshness = "last_month"
	WebSearchParamsFreshnessLastYear    WebSearchParamsFreshness = "last_year"
)

// Inline Markdown scraping for each result. Set `enabled: true` to activate.
type WebSearchParamsMarkdownOptions struct {
	// Scrape each result to Markdown. Off by default to keep search cheap and fast.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Render iframe contents into the Markdown.
	IncludeFrames param.Opt[bool] `json:"includeFrames,omitzero"`
	// Emit image references in the Markdown.
	IncludeImages param.Opt[bool] `json:"includeImages,omitzero"`
	// Keep hyperlinks in the Markdown.
	IncludeLinks param.Opt[bool] `json:"includeLinks,omitzero"`
	// Cache TTL in ms for scraped Markdown keyed by URL + options. Default 1 day, max
	// 30 days. Set to 0 to force a fresh scrape.
	MaxAgeMs param.Opt[int64] `json:"maxAgeMs,omitzero"`
	// Truncate inline base64 image payloads to keep responses small.
	ShortenBase64Images param.Opt[bool] `json:"shortenBase64Images,omitzero"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `json:"timeoutMS,omitzero"`
	// Strip nav, header, footer, and sidebar — keep only the primary article content.
	UseMainContentOnly param.Opt[bool] `json:"useMainContentOnly,omitzero"`
	// Extra wait after page load before rendering, in ms (0–30000). Useful for
	// JS-heavy pages.
	WaitForMs param.Opt[int64] `json:"waitForMs,omitzero"`
	// PDF handling. Use start/end to bound text extraction and OCR to a page range.
	Pdf WebSearchParamsMarkdownOptionsPdf `json:"pdf,omitzero"`
	paramObj
}

func (r WebSearchParamsMarkdownOptions) MarshalJSON() (data []byte, err error) {
	type shadow WebSearchParamsMarkdownOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebSearchParamsMarkdownOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PDF handling. Use start/end to bound text extraction and OCR to a page range.
type WebSearchParamsMarkdownOptionsPdf struct {
	// Last PDF page to parse (1-based, inclusive). Defaults to the final page. Must
	// be >= start.
	End param.Opt[int64] `json:"end,omitzero"`
	// Parse PDF URLs. When false, PDF results are skipped with WEBSITE_ACCESS_ERROR.
	ShouldParse param.Opt[bool] `json:"shouldParse,omitzero"`
	// First PDF page to parse (1-based, inclusive). Defaults to page 1.
	Start param.Opt[int64] `json:"start,omitzero"`
	paramObj
}

func (r WebSearchParamsMarkdownOptionsPdf) MarshalJSON() (data []byte, err error) {
	type shadow WebSearchParamsMarkdownOptionsPdf
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebSearchParamsMarkdownOptionsPdf) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebWebCrawlMdParams struct {
	// The starting URL for the crawl (must include http:// or https:// protocol)
	URL string `json:"url" api:"required" format:"uri"`
	// When true, follow links on subdomains of the starting URL's domain (e.g.
	// docs.example.com when starting from example.com). www and apex are always
	// treated as equivalent.
	FollowSubdomains param.Opt[bool] `json:"followSubdomains,omitzero"`
	// When true, the contents of iframes are rendered to Markdown for each crawled
	// page.
	IncludeFrames param.Opt[bool] `json:"includeFrames,omitzero"`
	// Include image references in the Markdown output
	IncludeImages param.Opt[bool] `json:"includeImages,omitzero"`
	// Preserve hyperlinks in the Markdown output
	IncludeLinks param.Opt[bool] `json:"includeLinks,omitzero"`
	// Return a cached result if a prior scrape for the same parameters exists and is
	// younger than this many milliseconds. Defaults to 1 day (86400000 ms) when
	// omitted. Max is 30 days (2592000000 ms). Set to 0 to always scrape fresh.
	MaxAgeMs param.Opt[int64] `json:"maxAgeMs,omitzero"`
	// Maximum link depth from the starting URL (0 = only the starting page)
	MaxDepth param.Opt[int64] `json:"maxDepth,omitzero"`
	// Maximum number of pages to crawl. Hard cap: 500.
	MaxPages param.Opt[int64] `json:"maxPages,omitzero"`
	// When true, waits briefly for CSS and transition animations to settle before
	// extracting each crawled page. Defaults to false. This adds a bit of latency in
	// exchange for more stable output on animated pages.
	SettleAnimations param.Opt[bool] `json:"settleAnimations,omitzero"`
	// Truncate base64-encoded image data in the Markdown output
	ShortenBase64Images param.Opt[bool] `json:"shortenBase64Images,omitzero"`
	// Soft time budget for the crawl in milliseconds. After each scrape, the crawler
	// checks the elapsed time and, if exceeded, returns the pages collected so far
	// instead of continuing. Min: 10000 (10s). Max: 110000 (110s). Default: 80000
	// (80s).
	StopAfterMs param.Opt[int64] `json:"stopAfterMs,omitzero"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `json:"timeoutMS,omitzero"`
	// Regex pattern. Only URLs matching this pattern will be followed and scraped.
	URLRegex param.Opt[string] `json:"urlRegex,omitzero"`
	// Extract only the main content, stripping headers, footers, sidebars, and
	// navigation
	UseMainContentOnly param.Opt[bool] `json:"useMainContentOnly,omitzero"`
	// Optional browser wait time in milliseconds after initial page load for each
	// crawled page. Min: 0. Max: 30000 (30 seconds).
	WaitForMs param.Opt[int64] `json:"waitForMs,omitzero"`
	// Fetch the target page through a residential proxy in this country (ISO 3166-1
	// alpha-2).
	//
	// Any of "ad", "ae", "af", "ag", "ai", "al", "am", "ao", "ar", "at", "au", "aw",
	// "az", "ba", "bb", "bd", "be", "bf", "bg", "bh", "bi", "bj", "bm", "bn", "bo",
	// "bq", "br", "bs", "bw", "by", "bz", "ca", "cd", "cf", "cg", "ch", "ci", "cl",
	// "cm", "cn", "co", "cr", "cv", "cw", "cy", "cz", "de", "dj", "dk", "dm", "do",
	// "dz", "ec", "ee", "eg", "es", "et", "fi", "fj", "fr", "ga", "gb", "gd", "ge",
	// "gf", "gg", "gh", "gm", "gn", "gp", "gq", "gr", "gt", "gu", "gw", "gy", "hk",
	// "hn", "hr", "ht", "hu", "id", "ie", "il", "im", "in", "iq", "ir", "is", "it",
	// "je", "jm", "jo", "jp", "ke", "kg", "kh", "kn", "kr", "kw", "ky", "kz", "la",
	// "lb", "lc", "lk", "lr", "ls", "lt", "lu", "lv", "ly", "ma", "mc", "md", "me",
	// "mf", "mg", "mk", "ml", "mm", "mn", "mo", "mq", "mr", "mt", "mu", "mv", "mw",
	// "mx", "my", "mz", "na", "nc", "ne", "ng", "ni", "nl", "no", "np", "nz", "om",
	// "pa", "pe", "pf", "pg", "ph", "pk", "pl", "pr", "ps", "pt", "py", "qa", "re",
	// "ro", "rs", "ru", "rw", "sa", "sc", "sd", "se", "sg", "si", "sk", "sl", "sm",
	// "sn", "so", "sr", "ss", "st", "sv", "sx", "sy", "sz", "tc", "td", "tg", "th",
	// "tj", "tl", "tm", "tn", "tr", "tt", "tw", "tz", "ua", "ug", "us", "uy", "uz",
	// "vc", "ve", "vg", "vi", "vn", "ye", "yt", "za", "zm", "zw".
	Country WebWebCrawlMdParamsCountry `json:"country,omitzero"`
	// CSS selectors to remove before each crawled page is converted to Markdown.
	// Applied after includeSelectors. Exclusion takes precedence: an element matching
	// both is removed. Examples: "nav", "footer", ".ad-banner", "[aria-hidden=true]".
	ExcludeSelectors []string `json:"excludeSelectors,omitzero"`
	// CSS selectors. When provided, only matching HTML subtrees (and their
	// descendants) are kept before each crawled page is converted to Markdown. When
	// omitted, the entire document is kept. Examples: "article.main", "#content",
	// "[role=main]".
	IncludeSelectors []string `json:"includeSelectors,omitzero"`
	// PDF parsing controls. Use start/end to limit text extraction and embedded-image
	// detection/OCR to an inclusive 1-based page range.
	Pdf WebWebCrawlMdParamsPdf `json:"pdf,omitzero"`
	// Optional tags for tracking usage. Up to 20 tags, each 1 to 50 characters.
	Tags []string `json:"tags,omitzero"`
	// Set to enabled to bypass shared caches and omit request and response content
	// from retained usage logs. Requires zero data retention to be enabled for your
	// organization (contact support@context.dev), otherwise the request fails with
	// ZDR_NOT_ENABLED. Successful ZDR responses include X-Context-ZDR: true.
	//
	// Any of "enabled", "disabled".
	Zdr WebWebCrawlMdParamsZdr `json:"zdr,omitzero"`
	paramObj
}

func (r WebWebCrawlMdParams) MarshalJSON() (data []byte, err error) {
	type shadow WebWebCrawlMdParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebWebCrawlMdParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Fetch the target page through a residential proxy in this country (ISO 3166-1
// alpha-2).
type WebWebCrawlMdParamsCountry string

const (
	WebWebCrawlMdParamsCountryAd WebWebCrawlMdParamsCountry = "ad"
	WebWebCrawlMdParamsCountryAe WebWebCrawlMdParamsCountry = "ae"
	WebWebCrawlMdParamsCountryAf WebWebCrawlMdParamsCountry = "af"
	WebWebCrawlMdParamsCountryAg WebWebCrawlMdParamsCountry = "ag"
	WebWebCrawlMdParamsCountryAI WebWebCrawlMdParamsCountry = "ai"
	WebWebCrawlMdParamsCountryAl WebWebCrawlMdParamsCountry = "al"
	WebWebCrawlMdParamsCountryAm WebWebCrawlMdParamsCountry = "am"
	WebWebCrawlMdParamsCountryAo WebWebCrawlMdParamsCountry = "ao"
	WebWebCrawlMdParamsCountryAr WebWebCrawlMdParamsCountry = "ar"
	WebWebCrawlMdParamsCountryAt WebWebCrawlMdParamsCountry = "at"
	WebWebCrawlMdParamsCountryAu WebWebCrawlMdParamsCountry = "au"
	WebWebCrawlMdParamsCountryAw WebWebCrawlMdParamsCountry = "aw"
	WebWebCrawlMdParamsCountryAz WebWebCrawlMdParamsCountry = "az"
	WebWebCrawlMdParamsCountryBa WebWebCrawlMdParamsCountry = "ba"
	WebWebCrawlMdParamsCountryBb WebWebCrawlMdParamsCountry = "bb"
	WebWebCrawlMdParamsCountryBd WebWebCrawlMdParamsCountry = "bd"
	WebWebCrawlMdParamsCountryBe WebWebCrawlMdParamsCountry = "be"
	WebWebCrawlMdParamsCountryBf WebWebCrawlMdParamsCountry = "bf"
	WebWebCrawlMdParamsCountryBg WebWebCrawlMdParamsCountry = "bg"
	WebWebCrawlMdParamsCountryBh WebWebCrawlMdParamsCountry = "bh"
	WebWebCrawlMdParamsCountryBi WebWebCrawlMdParamsCountry = "bi"
	WebWebCrawlMdParamsCountryBj WebWebCrawlMdParamsCountry = "bj"
	WebWebCrawlMdParamsCountryBm WebWebCrawlMdParamsCountry = "bm"
	WebWebCrawlMdParamsCountryBn WebWebCrawlMdParamsCountry = "bn"
	WebWebCrawlMdParamsCountryBo WebWebCrawlMdParamsCountry = "bo"
	WebWebCrawlMdParamsCountryBq WebWebCrawlMdParamsCountry = "bq"
	WebWebCrawlMdParamsCountryBr WebWebCrawlMdParamsCountry = "br"
	WebWebCrawlMdParamsCountryBs WebWebCrawlMdParamsCountry = "bs"
	WebWebCrawlMdParamsCountryBw WebWebCrawlMdParamsCountry = "bw"
	WebWebCrawlMdParamsCountryBy WebWebCrawlMdParamsCountry = "by"
	WebWebCrawlMdParamsCountryBz WebWebCrawlMdParamsCountry = "bz"
	WebWebCrawlMdParamsCountryCa WebWebCrawlMdParamsCountry = "ca"
	WebWebCrawlMdParamsCountryCd WebWebCrawlMdParamsCountry = "cd"
	WebWebCrawlMdParamsCountryCf WebWebCrawlMdParamsCountry = "cf"
	WebWebCrawlMdParamsCountryCg WebWebCrawlMdParamsCountry = "cg"
	WebWebCrawlMdParamsCountryCh WebWebCrawlMdParamsCountry = "ch"
	WebWebCrawlMdParamsCountryCi WebWebCrawlMdParamsCountry = "ci"
	WebWebCrawlMdParamsCountryCl WebWebCrawlMdParamsCountry = "cl"
	WebWebCrawlMdParamsCountryCm WebWebCrawlMdParamsCountry = "cm"
	WebWebCrawlMdParamsCountryCn WebWebCrawlMdParamsCountry = "cn"
	WebWebCrawlMdParamsCountryCo WebWebCrawlMdParamsCountry = "co"
	WebWebCrawlMdParamsCountryCr WebWebCrawlMdParamsCountry = "cr"
	WebWebCrawlMdParamsCountryCv WebWebCrawlMdParamsCountry = "cv"
	WebWebCrawlMdParamsCountryCw WebWebCrawlMdParamsCountry = "cw"
	WebWebCrawlMdParamsCountryCy WebWebCrawlMdParamsCountry = "cy"
	WebWebCrawlMdParamsCountryCz WebWebCrawlMdParamsCountry = "cz"
	WebWebCrawlMdParamsCountryDe WebWebCrawlMdParamsCountry = "de"
	WebWebCrawlMdParamsCountryDj WebWebCrawlMdParamsCountry = "dj"
	WebWebCrawlMdParamsCountryDk WebWebCrawlMdParamsCountry = "dk"
	WebWebCrawlMdParamsCountryDm WebWebCrawlMdParamsCountry = "dm"
	WebWebCrawlMdParamsCountryDo WebWebCrawlMdParamsCountry = "do"
	WebWebCrawlMdParamsCountryDz WebWebCrawlMdParamsCountry = "dz"
	WebWebCrawlMdParamsCountryEc WebWebCrawlMdParamsCountry = "ec"
	WebWebCrawlMdParamsCountryEe WebWebCrawlMdParamsCountry = "ee"
	WebWebCrawlMdParamsCountryEg WebWebCrawlMdParamsCountry = "eg"
	WebWebCrawlMdParamsCountryEs WebWebCrawlMdParamsCountry = "es"
	WebWebCrawlMdParamsCountryEt WebWebCrawlMdParamsCountry = "et"
	WebWebCrawlMdParamsCountryFi WebWebCrawlMdParamsCountry = "fi"
	WebWebCrawlMdParamsCountryFj WebWebCrawlMdParamsCountry = "fj"
	WebWebCrawlMdParamsCountryFr WebWebCrawlMdParamsCountry = "fr"
	WebWebCrawlMdParamsCountryGa WebWebCrawlMdParamsCountry = "ga"
	WebWebCrawlMdParamsCountryGB WebWebCrawlMdParamsCountry = "gb"
	WebWebCrawlMdParamsCountryGd WebWebCrawlMdParamsCountry = "gd"
	WebWebCrawlMdParamsCountryGe WebWebCrawlMdParamsCountry = "ge"
	WebWebCrawlMdParamsCountryGf WebWebCrawlMdParamsCountry = "gf"
	WebWebCrawlMdParamsCountryGg WebWebCrawlMdParamsCountry = "gg"
	WebWebCrawlMdParamsCountryGh WebWebCrawlMdParamsCountry = "gh"
	WebWebCrawlMdParamsCountryGm WebWebCrawlMdParamsCountry = "gm"
	WebWebCrawlMdParamsCountryGn WebWebCrawlMdParamsCountry = "gn"
	WebWebCrawlMdParamsCountryGp WebWebCrawlMdParamsCountry = "gp"
	WebWebCrawlMdParamsCountryGq WebWebCrawlMdParamsCountry = "gq"
	WebWebCrawlMdParamsCountryGr WebWebCrawlMdParamsCountry = "gr"
	WebWebCrawlMdParamsCountryGt WebWebCrawlMdParamsCountry = "gt"
	WebWebCrawlMdParamsCountryGu WebWebCrawlMdParamsCountry = "gu"
	WebWebCrawlMdParamsCountryGw WebWebCrawlMdParamsCountry = "gw"
	WebWebCrawlMdParamsCountryGy WebWebCrawlMdParamsCountry = "gy"
	WebWebCrawlMdParamsCountryHk WebWebCrawlMdParamsCountry = "hk"
	WebWebCrawlMdParamsCountryHn WebWebCrawlMdParamsCountry = "hn"
	WebWebCrawlMdParamsCountryHr WebWebCrawlMdParamsCountry = "hr"
	WebWebCrawlMdParamsCountryHt WebWebCrawlMdParamsCountry = "ht"
	WebWebCrawlMdParamsCountryHu WebWebCrawlMdParamsCountry = "hu"
	WebWebCrawlMdParamsCountryID WebWebCrawlMdParamsCountry = "id"
	WebWebCrawlMdParamsCountryIe WebWebCrawlMdParamsCountry = "ie"
	WebWebCrawlMdParamsCountryIl WebWebCrawlMdParamsCountry = "il"
	WebWebCrawlMdParamsCountryIm WebWebCrawlMdParamsCountry = "im"
	WebWebCrawlMdParamsCountryIn WebWebCrawlMdParamsCountry = "in"
	WebWebCrawlMdParamsCountryIq WebWebCrawlMdParamsCountry = "iq"
	WebWebCrawlMdParamsCountryIr WebWebCrawlMdParamsCountry = "ir"
	WebWebCrawlMdParamsCountryIs WebWebCrawlMdParamsCountry = "is"
	WebWebCrawlMdParamsCountryIt WebWebCrawlMdParamsCountry = "it"
	WebWebCrawlMdParamsCountryJe WebWebCrawlMdParamsCountry = "je"
	WebWebCrawlMdParamsCountryJm WebWebCrawlMdParamsCountry = "jm"
	WebWebCrawlMdParamsCountryJo WebWebCrawlMdParamsCountry = "jo"
	WebWebCrawlMdParamsCountryJp WebWebCrawlMdParamsCountry = "jp"
	WebWebCrawlMdParamsCountryKe WebWebCrawlMdParamsCountry = "ke"
	WebWebCrawlMdParamsCountryKg WebWebCrawlMdParamsCountry = "kg"
	WebWebCrawlMdParamsCountryKh WebWebCrawlMdParamsCountry = "kh"
	WebWebCrawlMdParamsCountryKn WebWebCrawlMdParamsCountry = "kn"
	WebWebCrawlMdParamsCountryKr WebWebCrawlMdParamsCountry = "kr"
	WebWebCrawlMdParamsCountryKw WebWebCrawlMdParamsCountry = "kw"
	WebWebCrawlMdParamsCountryKy WebWebCrawlMdParamsCountry = "ky"
	WebWebCrawlMdParamsCountryKz WebWebCrawlMdParamsCountry = "kz"
	WebWebCrawlMdParamsCountryLa WebWebCrawlMdParamsCountry = "la"
	WebWebCrawlMdParamsCountryLb WebWebCrawlMdParamsCountry = "lb"
	WebWebCrawlMdParamsCountryLc WebWebCrawlMdParamsCountry = "lc"
	WebWebCrawlMdParamsCountryLk WebWebCrawlMdParamsCountry = "lk"
	WebWebCrawlMdParamsCountryLr WebWebCrawlMdParamsCountry = "lr"
	WebWebCrawlMdParamsCountryLs WebWebCrawlMdParamsCountry = "ls"
	WebWebCrawlMdParamsCountryLt WebWebCrawlMdParamsCountry = "lt"
	WebWebCrawlMdParamsCountryLu WebWebCrawlMdParamsCountry = "lu"
	WebWebCrawlMdParamsCountryLv WebWebCrawlMdParamsCountry = "lv"
	WebWebCrawlMdParamsCountryLy WebWebCrawlMdParamsCountry = "ly"
	WebWebCrawlMdParamsCountryMa WebWebCrawlMdParamsCountry = "ma"
	WebWebCrawlMdParamsCountryMc WebWebCrawlMdParamsCountry = "mc"
	WebWebCrawlMdParamsCountryMd WebWebCrawlMdParamsCountry = "md"
	WebWebCrawlMdParamsCountryMe WebWebCrawlMdParamsCountry = "me"
	WebWebCrawlMdParamsCountryMf WebWebCrawlMdParamsCountry = "mf"
	WebWebCrawlMdParamsCountryMg WebWebCrawlMdParamsCountry = "mg"
	WebWebCrawlMdParamsCountryMk WebWebCrawlMdParamsCountry = "mk"
	WebWebCrawlMdParamsCountryMl WebWebCrawlMdParamsCountry = "ml"
	WebWebCrawlMdParamsCountryMm WebWebCrawlMdParamsCountry = "mm"
	WebWebCrawlMdParamsCountryMn WebWebCrawlMdParamsCountry = "mn"
	WebWebCrawlMdParamsCountryMo WebWebCrawlMdParamsCountry = "mo"
	WebWebCrawlMdParamsCountryMq WebWebCrawlMdParamsCountry = "mq"
	WebWebCrawlMdParamsCountryMr WebWebCrawlMdParamsCountry = "mr"
	WebWebCrawlMdParamsCountryMt WebWebCrawlMdParamsCountry = "mt"
	WebWebCrawlMdParamsCountryMu WebWebCrawlMdParamsCountry = "mu"
	WebWebCrawlMdParamsCountryMv WebWebCrawlMdParamsCountry = "mv"
	WebWebCrawlMdParamsCountryMw WebWebCrawlMdParamsCountry = "mw"
	WebWebCrawlMdParamsCountryMx WebWebCrawlMdParamsCountry = "mx"
	WebWebCrawlMdParamsCountryMy WebWebCrawlMdParamsCountry = "my"
	WebWebCrawlMdParamsCountryMz WebWebCrawlMdParamsCountry = "mz"
	WebWebCrawlMdParamsCountryNa WebWebCrawlMdParamsCountry = "na"
	WebWebCrawlMdParamsCountryNc WebWebCrawlMdParamsCountry = "nc"
	WebWebCrawlMdParamsCountryNe WebWebCrawlMdParamsCountry = "ne"
	WebWebCrawlMdParamsCountryNg WebWebCrawlMdParamsCountry = "ng"
	WebWebCrawlMdParamsCountryNi WebWebCrawlMdParamsCountry = "ni"
	WebWebCrawlMdParamsCountryNl WebWebCrawlMdParamsCountry = "nl"
	WebWebCrawlMdParamsCountryNo WebWebCrawlMdParamsCountry = "no"
	WebWebCrawlMdParamsCountryNp WebWebCrawlMdParamsCountry = "np"
	WebWebCrawlMdParamsCountryNz WebWebCrawlMdParamsCountry = "nz"
	WebWebCrawlMdParamsCountryOm WebWebCrawlMdParamsCountry = "om"
	WebWebCrawlMdParamsCountryPa WebWebCrawlMdParamsCountry = "pa"
	WebWebCrawlMdParamsCountryPe WebWebCrawlMdParamsCountry = "pe"
	WebWebCrawlMdParamsCountryPf WebWebCrawlMdParamsCountry = "pf"
	WebWebCrawlMdParamsCountryPg WebWebCrawlMdParamsCountry = "pg"
	WebWebCrawlMdParamsCountryPh WebWebCrawlMdParamsCountry = "ph"
	WebWebCrawlMdParamsCountryPk WebWebCrawlMdParamsCountry = "pk"
	WebWebCrawlMdParamsCountryPl WebWebCrawlMdParamsCountry = "pl"
	WebWebCrawlMdParamsCountryPr WebWebCrawlMdParamsCountry = "pr"
	WebWebCrawlMdParamsCountryPs WebWebCrawlMdParamsCountry = "ps"
	WebWebCrawlMdParamsCountryPt WebWebCrawlMdParamsCountry = "pt"
	WebWebCrawlMdParamsCountryPy WebWebCrawlMdParamsCountry = "py"
	WebWebCrawlMdParamsCountryQa WebWebCrawlMdParamsCountry = "qa"
	WebWebCrawlMdParamsCountryRe WebWebCrawlMdParamsCountry = "re"
	WebWebCrawlMdParamsCountryRo WebWebCrawlMdParamsCountry = "ro"
	WebWebCrawlMdParamsCountryRs WebWebCrawlMdParamsCountry = "rs"
	WebWebCrawlMdParamsCountryRu WebWebCrawlMdParamsCountry = "ru"
	WebWebCrawlMdParamsCountryRw WebWebCrawlMdParamsCountry = "rw"
	WebWebCrawlMdParamsCountrySa WebWebCrawlMdParamsCountry = "sa"
	WebWebCrawlMdParamsCountrySc WebWebCrawlMdParamsCountry = "sc"
	WebWebCrawlMdParamsCountrySd WebWebCrawlMdParamsCountry = "sd"
	WebWebCrawlMdParamsCountrySe WebWebCrawlMdParamsCountry = "se"
	WebWebCrawlMdParamsCountrySg WebWebCrawlMdParamsCountry = "sg"
	WebWebCrawlMdParamsCountrySi WebWebCrawlMdParamsCountry = "si"
	WebWebCrawlMdParamsCountrySk WebWebCrawlMdParamsCountry = "sk"
	WebWebCrawlMdParamsCountrySl WebWebCrawlMdParamsCountry = "sl"
	WebWebCrawlMdParamsCountrySm WebWebCrawlMdParamsCountry = "sm"
	WebWebCrawlMdParamsCountrySn WebWebCrawlMdParamsCountry = "sn"
	WebWebCrawlMdParamsCountrySo WebWebCrawlMdParamsCountry = "so"
	WebWebCrawlMdParamsCountrySr WebWebCrawlMdParamsCountry = "sr"
	WebWebCrawlMdParamsCountrySS WebWebCrawlMdParamsCountry = "ss"
	WebWebCrawlMdParamsCountrySt WebWebCrawlMdParamsCountry = "st"
	WebWebCrawlMdParamsCountrySv WebWebCrawlMdParamsCountry = "sv"
	WebWebCrawlMdParamsCountrySx WebWebCrawlMdParamsCountry = "sx"
	WebWebCrawlMdParamsCountrySy WebWebCrawlMdParamsCountry = "sy"
	WebWebCrawlMdParamsCountrySz WebWebCrawlMdParamsCountry = "sz"
	WebWebCrawlMdParamsCountryTc WebWebCrawlMdParamsCountry = "tc"
	WebWebCrawlMdParamsCountryTd WebWebCrawlMdParamsCountry = "td"
	WebWebCrawlMdParamsCountryTg WebWebCrawlMdParamsCountry = "tg"
	WebWebCrawlMdParamsCountryTh WebWebCrawlMdParamsCountry = "th"
	WebWebCrawlMdParamsCountryTj WebWebCrawlMdParamsCountry = "tj"
	WebWebCrawlMdParamsCountryTl WebWebCrawlMdParamsCountry = "tl"
	WebWebCrawlMdParamsCountryTm WebWebCrawlMdParamsCountry = "tm"
	WebWebCrawlMdParamsCountryTn WebWebCrawlMdParamsCountry = "tn"
	WebWebCrawlMdParamsCountryTr WebWebCrawlMdParamsCountry = "tr"
	WebWebCrawlMdParamsCountryTt WebWebCrawlMdParamsCountry = "tt"
	WebWebCrawlMdParamsCountryTw WebWebCrawlMdParamsCountry = "tw"
	WebWebCrawlMdParamsCountryTz WebWebCrawlMdParamsCountry = "tz"
	WebWebCrawlMdParamsCountryUa WebWebCrawlMdParamsCountry = "ua"
	WebWebCrawlMdParamsCountryUg WebWebCrawlMdParamsCountry = "ug"
	WebWebCrawlMdParamsCountryUs WebWebCrawlMdParamsCountry = "us"
	WebWebCrawlMdParamsCountryUy WebWebCrawlMdParamsCountry = "uy"
	WebWebCrawlMdParamsCountryUz WebWebCrawlMdParamsCountry = "uz"
	WebWebCrawlMdParamsCountryVc WebWebCrawlMdParamsCountry = "vc"
	WebWebCrawlMdParamsCountryVe WebWebCrawlMdParamsCountry = "ve"
	WebWebCrawlMdParamsCountryVg WebWebCrawlMdParamsCountry = "vg"
	WebWebCrawlMdParamsCountryVi WebWebCrawlMdParamsCountry = "vi"
	WebWebCrawlMdParamsCountryVn WebWebCrawlMdParamsCountry = "vn"
	WebWebCrawlMdParamsCountryYe WebWebCrawlMdParamsCountry = "ye"
	WebWebCrawlMdParamsCountryYt WebWebCrawlMdParamsCountry = "yt"
	WebWebCrawlMdParamsCountryZa WebWebCrawlMdParamsCountry = "za"
	WebWebCrawlMdParamsCountryZm WebWebCrawlMdParamsCountry = "zm"
	WebWebCrawlMdParamsCountryZw WebWebCrawlMdParamsCountry = "zw"
)

// PDF parsing controls. Use start/end to limit text extraction and embedded-image
// detection/OCR to an inclusive 1-based page range.
type WebWebCrawlMdParamsPdf struct {
	// Last 1-based PDF page to parse. When omitted, parsing ends at the last page.
	// Must be greater than or equal to start when both are provided.
	End param.Opt[int64] `json:"end,omitzero"`
	// When true, detect and OCR images embedded in the selected PDF pages, inserting
	// recognized text at each image's position in page reading order while preserving
	// the PDF text layer. This is separate from automatic scanned-PDF OCR fallback.
	Ocr param.Opt[bool] `json:"ocr,omitzero"`
	// When true, PDF pages are fetched and parsed. When false, PDF pages are skipped
	// entirely (not included in results and not counted as failures).
	ShouldParse param.Opt[bool] `json:"shouldParse,omitzero"`
	// First 1-based PDF page to parse. When omitted, parsing starts at the first page.
	Start param.Opt[int64] `json:"start,omitzero"`
	paramObj
}

func (r WebWebCrawlMdParamsPdf) MarshalJSON() (data []byte, err error) {
	type shadow WebWebCrawlMdParamsPdf
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebWebCrawlMdParamsPdf) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Set to enabled to bypass shared caches and omit request and response content
// from retained usage logs. Requires zero data retention to be enabled for your
// organization (contact support@context.dev), otherwise the request fails with
// ZDR_NOT_ENABLED. Successful ZDR responses include X-Context-ZDR: true.
type WebWebCrawlMdParamsZdr string

const (
	WebWebCrawlMdParamsZdrEnabled  WebWebCrawlMdParamsZdr = "enabled"
	WebWebCrawlMdParamsZdrDisabled WebWebCrawlMdParamsZdr = "disabled"
)

type WebWebScrapeHTMLParams struct {
	// Full URL to scrape (must include http:// or https:// protocol)
	URL string `query:"url" api:"required" format:"uri" json:"-"`
	// Return a cached result if a prior scrape for the same parameters exists and is
	// younger than this many milliseconds. Defaults to 1 day (86400000 ms) when
	// omitted. Max is 30 days (2592000000 ms). Set to 0 to always scrape fresh.
	MaxAgeMs param.Opt[int64] `query:"maxAgeMs,omitzero" json:"-"`
	// Optional browser wait time in milliseconds after initial page load. Min: 0. Max:
	// 30000 (30 seconds).
	WaitForMs param.Opt[int64] `query:"waitForMs,omitzero" json:"-"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `query:"timeoutMS,omitzero" json:"-"`
	// Optional browser actions executed in array order after the page loads and before
	// content is captured. Requires a paid plan. Send a JSON array in the query
	// parameter. Maximum: 5 actions.
	Actions []WebWebScrapeHTMLParamsActionUnion `query:"actions,omitzero" json:"-"`
	// CSS selectors to remove from the result. Applied after includeSelectors.
	// Exclusion takes precedence: an element matching both is removed. Examples:
	// "nav", "footer", ".ad-banner", "[aria-hidden=true]".
	ExcludeSelectors []string `query:"excludeSelectors,omitzero" json:"-"`
	// CSS selectors. When provided, only matching subtrees (and their descendants) are
	// kept and everything else is dropped. When omitted, the entire document is kept.
	// Examples: "article.main", "#content", "[role=main]".
	IncludeSelectors []string `query:"includeSelectors,omitzero" json:"-"`
	// Fetch the target page through a residential proxy in this country (ISO 3166-1
	// alpha-2).
	//
	// Any of "ad", "ae", "af", "ag", "ai", "al", "am", "ao", "ar", "at", "au", "aw",
	// "az", "ba", "bb", "bd", "be", "bf", "bg", "bh", "bi", "bj", "bm", "bn", "bo",
	// "bq", "br", "bs", "bw", "by", "bz", "ca", "cd", "cf", "cg", "ch", "ci", "cl",
	// "cm", "cn", "co", "cr", "cv", "cw", "cy", "cz", "de", "dj", "dk", "dm", "do",
	// "dz", "ec", "ee", "eg", "es", "et", "fi", "fj", "fr", "ga", "gb", "gd", "ge",
	// "gf", "gg", "gh", "gm", "gn", "gp", "gq", "gr", "gt", "gu", "gw", "gy", "hk",
	// "hn", "hr", "ht", "hu", "id", "ie", "il", "im", "in", "iq", "ir", "is", "it",
	// "je", "jm", "jo", "jp", "ke", "kg", "kh", "kn", "kr", "kw", "ky", "kz", "la",
	// "lb", "lc", "lk", "lr", "ls", "lt", "lu", "lv", "ly", "ma", "mc", "md", "me",
	// "mf", "mg", "mk", "ml", "mm", "mn", "mo", "mq", "mr", "mt", "mu", "mv", "mw",
	// "mx", "my", "mz", "na", "nc", "ne", "ng", "ni", "nl", "no", "np", "nz", "om",
	// "pa", "pe", "pf", "pg", "ph", "pk", "pl", "pr", "ps", "pt", "py", "qa", "re",
	// "ro", "rs", "ru", "rw", "sa", "sc", "sd", "se", "sg", "si", "sk", "sl", "sm",
	// "sn", "so", "sr", "ss", "st", "sv", "sx", "sy", "sz", "tc", "td", "tg", "th",
	// "tj", "tl", "tm", "tn", "tr", "tt", "tw", "tz", "ua", "ug", "us", "uy", "uz",
	// "vc", "ve", "vg", "vi", "vn", "ye", "yt", "za", "zm", "zw".
	Country WebWebScrapeHTMLParamsCountry `query:"country,omitzero" json:"-"`
	// Optional outbound HTTP headers forwarded only to the target URL, sent as
	// deep-object query params such as headers[X-Custom]=value. When provided, caching
	// is bypassed: the result is neither read from nor written to cache.
	Headers map[string]string `query:"headers,omitzero" json:"-"`
	// When true, iframes are rendered inline into the returned HTML.
	IncludeFrames WebWebScrapeHTMLParamsIncludeFramesUnion `query:"includeFrames,omitzero" json:"-"`
	// PDF parsing controls. Use start/end to limit text extraction and embedded-image
	// detection/OCR to an inclusive 1-based page range.
	Pdf WebWebScrapeHTMLParamsPdf `query:"pdf,omitzero" json:"-"`
	// When true, waits briefly for CSS and transition animations to settle before
	// extracting HTML. Defaults to false. This adds a bit of latency in exchange for
	// more stable output on animated pages.
	SettleAnimations WebWebScrapeHTMLParamsSettleAnimationsUnion `query:"settleAnimations,omitzero" json:"-"`
	// Optional comma-separated caller-defined tags for tracking this request. Tags are
	// recorded on the request's usage log and can be used to filter usage on the
	// dashboard usage page. Up to 20 tags, each 1-50 characters.
	Tags []string `query:"tags,omitzero" json:"-"`
	// When true, return only the page's main content in the HTML response, excluding
	// headers, footers, sidebars, and navigation when detectable.
	UseMainContentOnly WebWebScrapeHTMLParamsUseMainContentOnlyUnion `query:"useMainContentOnly,omitzero" json:"-"`
	// Set to enabled to bypass shared caches and omit request and response content
	// from retained usage logs. Requires zero data retention to be enabled for your
	// organization (contact support@context.dev), otherwise the request fails with
	// ZDR_NOT_ENABLED. Successful ZDR responses include X-Context-ZDR: true.
	//
	// Any of "enabled", "disabled".
	Zdr WebWebScrapeHTMLParamsZdr `query:"zdr,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebWebScrapeHTMLParams]'s query parameters as `url.Values`.
func (r WebWebScrapeHTMLParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebWebScrapeHTMLParamsActionUnion struct {
	OfWait    *WebWebScrapeHTMLParamsActionWait    `query:",omitzero,inline"`
	OfPerform *WebWebScrapeHTMLParamsActionPerform `query:",omitzero,inline"`
	paramUnion
}

func init() {
	apijson.RegisterUnion[WebWebScrapeHTMLParamsActionUnion](
		"do",
		apijson.Discriminator[WebWebScrapeHTMLParamsActionWait]("wait"),
		apijson.Discriminator[WebWebScrapeHTMLParamsActionPerform]("perform"),
	)
}

// Pause for a fixed number of milliseconds before continuing to the next action.
//
// The properties Do, TimeMs are required.
type WebWebScrapeHTMLParamsActionWait struct {
	TimeMs int64 `query:"timeMs" api:"required" json:"-"`
	// This field can be elided, and will marshal its zero value as "wait".
	Do constant.Wait `query:"do" json:"-" default:"wait"`
	paramObj
}

// URLQuery serializes [WebWebScrapeHTMLParamsActionWait]'s query parameters as
// `url.Values`.
func (r WebWebScrapeHTMLParamsActionWait) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Resolve and perform one natural-language browser action.
//
// The properties Action, Do are required.
type WebWebScrapeHTMLParamsActionPerform struct {
	Action string `query:"action" api:"required" json:"-"`
	// This field can be elided, and will marshal its zero value as "perform".
	Do constant.Perform `query:"do" json:"-" default:"perform"`
	paramObj
}

// URLQuery serializes [WebWebScrapeHTMLParamsActionPerform]'s query parameters as
// `url.Values`.
func (r WebWebScrapeHTMLParamsActionPerform) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Fetch the target page through a residential proxy in this country (ISO 3166-1
// alpha-2).
type WebWebScrapeHTMLParamsCountry string

const (
	WebWebScrapeHTMLParamsCountryAd WebWebScrapeHTMLParamsCountry = "ad"
	WebWebScrapeHTMLParamsCountryAe WebWebScrapeHTMLParamsCountry = "ae"
	WebWebScrapeHTMLParamsCountryAf WebWebScrapeHTMLParamsCountry = "af"
	WebWebScrapeHTMLParamsCountryAg WebWebScrapeHTMLParamsCountry = "ag"
	WebWebScrapeHTMLParamsCountryAI WebWebScrapeHTMLParamsCountry = "ai"
	WebWebScrapeHTMLParamsCountryAl WebWebScrapeHTMLParamsCountry = "al"
	WebWebScrapeHTMLParamsCountryAm WebWebScrapeHTMLParamsCountry = "am"
	WebWebScrapeHTMLParamsCountryAo WebWebScrapeHTMLParamsCountry = "ao"
	WebWebScrapeHTMLParamsCountryAr WebWebScrapeHTMLParamsCountry = "ar"
	WebWebScrapeHTMLParamsCountryAt WebWebScrapeHTMLParamsCountry = "at"
	WebWebScrapeHTMLParamsCountryAu WebWebScrapeHTMLParamsCountry = "au"
	WebWebScrapeHTMLParamsCountryAw WebWebScrapeHTMLParamsCountry = "aw"
	WebWebScrapeHTMLParamsCountryAz WebWebScrapeHTMLParamsCountry = "az"
	WebWebScrapeHTMLParamsCountryBa WebWebScrapeHTMLParamsCountry = "ba"
	WebWebScrapeHTMLParamsCountryBb WebWebScrapeHTMLParamsCountry = "bb"
	WebWebScrapeHTMLParamsCountryBd WebWebScrapeHTMLParamsCountry = "bd"
	WebWebScrapeHTMLParamsCountryBe WebWebScrapeHTMLParamsCountry = "be"
	WebWebScrapeHTMLParamsCountryBf WebWebScrapeHTMLParamsCountry = "bf"
	WebWebScrapeHTMLParamsCountryBg WebWebScrapeHTMLParamsCountry = "bg"
	WebWebScrapeHTMLParamsCountryBh WebWebScrapeHTMLParamsCountry = "bh"
	WebWebScrapeHTMLParamsCountryBi WebWebScrapeHTMLParamsCountry = "bi"
	WebWebScrapeHTMLParamsCountryBj WebWebScrapeHTMLParamsCountry = "bj"
	WebWebScrapeHTMLParamsCountryBm WebWebScrapeHTMLParamsCountry = "bm"
	WebWebScrapeHTMLParamsCountryBn WebWebScrapeHTMLParamsCountry = "bn"
	WebWebScrapeHTMLParamsCountryBo WebWebScrapeHTMLParamsCountry = "bo"
	WebWebScrapeHTMLParamsCountryBq WebWebScrapeHTMLParamsCountry = "bq"
	WebWebScrapeHTMLParamsCountryBr WebWebScrapeHTMLParamsCountry = "br"
	WebWebScrapeHTMLParamsCountryBs WebWebScrapeHTMLParamsCountry = "bs"
	WebWebScrapeHTMLParamsCountryBw WebWebScrapeHTMLParamsCountry = "bw"
	WebWebScrapeHTMLParamsCountryBy WebWebScrapeHTMLParamsCountry = "by"
	WebWebScrapeHTMLParamsCountryBz WebWebScrapeHTMLParamsCountry = "bz"
	WebWebScrapeHTMLParamsCountryCa WebWebScrapeHTMLParamsCountry = "ca"
	WebWebScrapeHTMLParamsCountryCd WebWebScrapeHTMLParamsCountry = "cd"
	WebWebScrapeHTMLParamsCountryCf WebWebScrapeHTMLParamsCountry = "cf"
	WebWebScrapeHTMLParamsCountryCg WebWebScrapeHTMLParamsCountry = "cg"
	WebWebScrapeHTMLParamsCountryCh WebWebScrapeHTMLParamsCountry = "ch"
	WebWebScrapeHTMLParamsCountryCi WebWebScrapeHTMLParamsCountry = "ci"
	WebWebScrapeHTMLParamsCountryCl WebWebScrapeHTMLParamsCountry = "cl"
	WebWebScrapeHTMLParamsCountryCm WebWebScrapeHTMLParamsCountry = "cm"
	WebWebScrapeHTMLParamsCountryCn WebWebScrapeHTMLParamsCountry = "cn"
	WebWebScrapeHTMLParamsCountryCo WebWebScrapeHTMLParamsCountry = "co"
	WebWebScrapeHTMLParamsCountryCr WebWebScrapeHTMLParamsCountry = "cr"
	WebWebScrapeHTMLParamsCountryCv WebWebScrapeHTMLParamsCountry = "cv"
	WebWebScrapeHTMLParamsCountryCw WebWebScrapeHTMLParamsCountry = "cw"
	WebWebScrapeHTMLParamsCountryCy WebWebScrapeHTMLParamsCountry = "cy"
	WebWebScrapeHTMLParamsCountryCz WebWebScrapeHTMLParamsCountry = "cz"
	WebWebScrapeHTMLParamsCountryDe WebWebScrapeHTMLParamsCountry = "de"
	WebWebScrapeHTMLParamsCountryDj WebWebScrapeHTMLParamsCountry = "dj"
	WebWebScrapeHTMLParamsCountryDk WebWebScrapeHTMLParamsCountry = "dk"
	WebWebScrapeHTMLParamsCountryDm WebWebScrapeHTMLParamsCountry = "dm"
	WebWebScrapeHTMLParamsCountryDo WebWebScrapeHTMLParamsCountry = "do"
	WebWebScrapeHTMLParamsCountryDz WebWebScrapeHTMLParamsCountry = "dz"
	WebWebScrapeHTMLParamsCountryEc WebWebScrapeHTMLParamsCountry = "ec"
	WebWebScrapeHTMLParamsCountryEe WebWebScrapeHTMLParamsCountry = "ee"
	WebWebScrapeHTMLParamsCountryEg WebWebScrapeHTMLParamsCountry = "eg"
	WebWebScrapeHTMLParamsCountryEs WebWebScrapeHTMLParamsCountry = "es"
	WebWebScrapeHTMLParamsCountryEt WebWebScrapeHTMLParamsCountry = "et"
	WebWebScrapeHTMLParamsCountryFi WebWebScrapeHTMLParamsCountry = "fi"
	WebWebScrapeHTMLParamsCountryFj WebWebScrapeHTMLParamsCountry = "fj"
	WebWebScrapeHTMLParamsCountryFr WebWebScrapeHTMLParamsCountry = "fr"
	WebWebScrapeHTMLParamsCountryGa WebWebScrapeHTMLParamsCountry = "ga"
	WebWebScrapeHTMLParamsCountryGB WebWebScrapeHTMLParamsCountry = "gb"
	WebWebScrapeHTMLParamsCountryGd WebWebScrapeHTMLParamsCountry = "gd"
	WebWebScrapeHTMLParamsCountryGe WebWebScrapeHTMLParamsCountry = "ge"
	WebWebScrapeHTMLParamsCountryGf WebWebScrapeHTMLParamsCountry = "gf"
	WebWebScrapeHTMLParamsCountryGg WebWebScrapeHTMLParamsCountry = "gg"
	WebWebScrapeHTMLParamsCountryGh WebWebScrapeHTMLParamsCountry = "gh"
	WebWebScrapeHTMLParamsCountryGm WebWebScrapeHTMLParamsCountry = "gm"
	WebWebScrapeHTMLParamsCountryGn WebWebScrapeHTMLParamsCountry = "gn"
	WebWebScrapeHTMLParamsCountryGp WebWebScrapeHTMLParamsCountry = "gp"
	WebWebScrapeHTMLParamsCountryGq WebWebScrapeHTMLParamsCountry = "gq"
	WebWebScrapeHTMLParamsCountryGr WebWebScrapeHTMLParamsCountry = "gr"
	WebWebScrapeHTMLParamsCountryGt WebWebScrapeHTMLParamsCountry = "gt"
	WebWebScrapeHTMLParamsCountryGu WebWebScrapeHTMLParamsCountry = "gu"
	WebWebScrapeHTMLParamsCountryGw WebWebScrapeHTMLParamsCountry = "gw"
	WebWebScrapeHTMLParamsCountryGy WebWebScrapeHTMLParamsCountry = "gy"
	WebWebScrapeHTMLParamsCountryHk WebWebScrapeHTMLParamsCountry = "hk"
	WebWebScrapeHTMLParamsCountryHn WebWebScrapeHTMLParamsCountry = "hn"
	WebWebScrapeHTMLParamsCountryHr WebWebScrapeHTMLParamsCountry = "hr"
	WebWebScrapeHTMLParamsCountryHt WebWebScrapeHTMLParamsCountry = "ht"
	WebWebScrapeHTMLParamsCountryHu WebWebScrapeHTMLParamsCountry = "hu"
	WebWebScrapeHTMLParamsCountryID WebWebScrapeHTMLParamsCountry = "id"
	WebWebScrapeHTMLParamsCountryIe WebWebScrapeHTMLParamsCountry = "ie"
	WebWebScrapeHTMLParamsCountryIl WebWebScrapeHTMLParamsCountry = "il"
	WebWebScrapeHTMLParamsCountryIm WebWebScrapeHTMLParamsCountry = "im"
	WebWebScrapeHTMLParamsCountryIn WebWebScrapeHTMLParamsCountry = "in"
	WebWebScrapeHTMLParamsCountryIq WebWebScrapeHTMLParamsCountry = "iq"
	WebWebScrapeHTMLParamsCountryIr WebWebScrapeHTMLParamsCountry = "ir"
	WebWebScrapeHTMLParamsCountryIs WebWebScrapeHTMLParamsCountry = "is"
	WebWebScrapeHTMLParamsCountryIt WebWebScrapeHTMLParamsCountry = "it"
	WebWebScrapeHTMLParamsCountryJe WebWebScrapeHTMLParamsCountry = "je"
	WebWebScrapeHTMLParamsCountryJm WebWebScrapeHTMLParamsCountry = "jm"
	WebWebScrapeHTMLParamsCountryJo WebWebScrapeHTMLParamsCountry = "jo"
	WebWebScrapeHTMLParamsCountryJp WebWebScrapeHTMLParamsCountry = "jp"
	WebWebScrapeHTMLParamsCountryKe WebWebScrapeHTMLParamsCountry = "ke"
	WebWebScrapeHTMLParamsCountryKg WebWebScrapeHTMLParamsCountry = "kg"
	WebWebScrapeHTMLParamsCountryKh WebWebScrapeHTMLParamsCountry = "kh"
	WebWebScrapeHTMLParamsCountryKn WebWebScrapeHTMLParamsCountry = "kn"
	WebWebScrapeHTMLParamsCountryKr WebWebScrapeHTMLParamsCountry = "kr"
	WebWebScrapeHTMLParamsCountryKw WebWebScrapeHTMLParamsCountry = "kw"
	WebWebScrapeHTMLParamsCountryKy WebWebScrapeHTMLParamsCountry = "ky"
	WebWebScrapeHTMLParamsCountryKz WebWebScrapeHTMLParamsCountry = "kz"
	WebWebScrapeHTMLParamsCountryLa WebWebScrapeHTMLParamsCountry = "la"
	WebWebScrapeHTMLParamsCountryLb WebWebScrapeHTMLParamsCountry = "lb"
	WebWebScrapeHTMLParamsCountryLc WebWebScrapeHTMLParamsCountry = "lc"
	WebWebScrapeHTMLParamsCountryLk WebWebScrapeHTMLParamsCountry = "lk"
	WebWebScrapeHTMLParamsCountryLr WebWebScrapeHTMLParamsCountry = "lr"
	WebWebScrapeHTMLParamsCountryLs WebWebScrapeHTMLParamsCountry = "ls"
	WebWebScrapeHTMLParamsCountryLt WebWebScrapeHTMLParamsCountry = "lt"
	WebWebScrapeHTMLParamsCountryLu WebWebScrapeHTMLParamsCountry = "lu"
	WebWebScrapeHTMLParamsCountryLv WebWebScrapeHTMLParamsCountry = "lv"
	WebWebScrapeHTMLParamsCountryLy WebWebScrapeHTMLParamsCountry = "ly"
	WebWebScrapeHTMLParamsCountryMa WebWebScrapeHTMLParamsCountry = "ma"
	WebWebScrapeHTMLParamsCountryMc WebWebScrapeHTMLParamsCountry = "mc"
	WebWebScrapeHTMLParamsCountryMd WebWebScrapeHTMLParamsCountry = "md"
	WebWebScrapeHTMLParamsCountryMe WebWebScrapeHTMLParamsCountry = "me"
	WebWebScrapeHTMLParamsCountryMf WebWebScrapeHTMLParamsCountry = "mf"
	WebWebScrapeHTMLParamsCountryMg WebWebScrapeHTMLParamsCountry = "mg"
	WebWebScrapeHTMLParamsCountryMk WebWebScrapeHTMLParamsCountry = "mk"
	WebWebScrapeHTMLParamsCountryMl WebWebScrapeHTMLParamsCountry = "ml"
	WebWebScrapeHTMLParamsCountryMm WebWebScrapeHTMLParamsCountry = "mm"
	WebWebScrapeHTMLParamsCountryMn WebWebScrapeHTMLParamsCountry = "mn"
	WebWebScrapeHTMLParamsCountryMo WebWebScrapeHTMLParamsCountry = "mo"
	WebWebScrapeHTMLParamsCountryMq WebWebScrapeHTMLParamsCountry = "mq"
	WebWebScrapeHTMLParamsCountryMr WebWebScrapeHTMLParamsCountry = "mr"
	WebWebScrapeHTMLParamsCountryMt WebWebScrapeHTMLParamsCountry = "mt"
	WebWebScrapeHTMLParamsCountryMu WebWebScrapeHTMLParamsCountry = "mu"
	WebWebScrapeHTMLParamsCountryMv WebWebScrapeHTMLParamsCountry = "mv"
	WebWebScrapeHTMLParamsCountryMw WebWebScrapeHTMLParamsCountry = "mw"
	WebWebScrapeHTMLParamsCountryMx WebWebScrapeHTMLParamsCountry = "mx"
	WebWebScrapeHTMLParamsCountryMy WebWebScrapeHTMLParamsCountry = "my"
	WebWebScrapeHTMLParamsCountryMz WebWebScrapeHTMLParamsCountry = "mz"
	WebWebScrapeHTMLParamsCountryNa WebWebScrapeHTMLParamsCountry = "na"
	WebWebScrapeHTMLParamsCountryNc WebWebScrapeHTMLParamsCountry = "nc"
	WebWebScrapeHTMLParamsCountryNe WebWebScrapeHTMLParamsCountry = "ne"
	WebWebScrapeHTMLParamsCountryNg WebWebScrapeHTMLParamsCountry = "ng"
	WebWebScrapeHTMLParamsCountryNi WebWebScrapeHTMLParamsCountry = "ni"
	WebWebScrapeHTMLParamsCountryNl WebWebScrapeHTMLParamsCountry = "nl"
	WebWebScrapeHTMLParamsCountryNo WebWebScrapeHTMLParamsCountry = "no"
	WebWebScrapeHTMLParamsCountryNp WebWebScrapeHTMLParamsCountry = "np"
	WebWebScrapeHTMLParamsCountryNz WebWebScrapeHTMLParamsCountry = "nz"
	WebWebScrapeHTMLParamsCountryOm WebWebScrapeHTMLParamsCountry = "om"
	WebWebScrapeHTMLParamsCountryPa WebWebScrapeHTMLParamsCountry = "pa"
	WebWebScrapeHTMLParamsCountryPe WebWebScrapeHTMLParamsCountry = "pe"
	WebWebScrapeHTMLParamsCountryPf WebWebScrapeHTMLParamsCountry = "pf"
	WebWebScrapeHTMLParamsCountryPg WebWebScrapeHTMLParamsCountry = "pg"
	WebWebScrapeHTMLParamsCountryPh WebWebScrapeHTMLParamsCountry = "ph"
	WebWebScrapeHTMLParamsCountryPk WebWebScrapeHTMLParamsCountry = "pk"
	WebWebScrapeHTMLParamsCountryPl WebWebScrapeHTMLParamsCountry = "pl"
	WebWebScrapeHTMLParamsCountryPr WebWebScrapeHTMLParamsCountry = "pr"
	WebWebScrapeHTMLParamsCountryPs WebWebScrapeHTMLParamsCountry = "ps"
	WebWebScrapeHTMLParamsCountryPt WebWebScrapeHTMLParamsCountry = "pt"
	WebWebScrapeHTMLParamsCountryPy WebWebScrapeHTMLParamsCountry = "py"
	WebWebScrapeHTMLParamsCountryQa WebWebScrapeHTMLParamsCountry = "qa"
	WebWebScrapeHTMLParamsCountryRe WebWebScrapeHTMLParamsCountry = "re"
	WebWebScrapeHTMLParamsCountryRo WebWebScrapeHTMLParamsCountry = "ro"
	WebWebScrapeHTMLParamsCountryRs WebWebScrapeHTMLParamsCountry = "rs"
	WebWebScrapeHTMLParamsCountryRu WebWebScrapeHTMLParamsCountry = "ru"
	WebWebScrapeHTMLParamsCountryRw WebWebScrapeHTMLParamsCountry = "rw"
	WebWebScrapeHTMLParamsCountrySa WebWebScrapeHTMLParamsCountry = "sa"
	WebWebScrapeHTMLParamsCountrySc WebWebScrapeHTMLParamsCountry = "sc"
	WebWebScrapeHTMLParamsCountrySd WebWebScrapeHTMLParamsCountry = "sd"
	WebWebScrapeHTMLParamsCountrySe WebWebScrapeHTMLParamsCountry = "se"
	WebWebScrapeHTMLParamsCountrySg WebWebScrapeHTMLParamsCountry = "sg"
	WebWebScrapeHTMLParamsCountrySi WebWebScrapeHTMLParamsCountry = "si"
	WebWebScrapeHTMLParamsCountrySk WebWebScrapeHTMLParamsCountry = "sk"
	WebWebScrapeHTMLParamsCountrySl WebWebScrapeHTMLParamsCountry = "sl"
	WebWebScrapeHTMLParamsCountrySm WebWebScrapeHTMLParamsCountry = "sm"
	WebWebScrapeHTMLParamsCountrySn WebWebScrapeHTMLParamsCountry = "sn"
	WebWebScrapeHTMLParamsCountrySo WebWebScrapeHTMLParamsCountry = "so"
	WebWebScrapeHTMLParamsCountrySr WebWebScrapeHTMLParamsCountry = "sr"
	WebWebScrapeHTMLParamsCountrySS WebWebScrapeHTMLParamsCountry = "ss"
	WebWebScrapeHTMLParamsCountrySt WebWebScrapeHTMLParamsCountry = "st"
	WebWebScrapeHTMLParamsCountrySv WebWebScrapeHTMLParamsCountry = "sv"
	WebWebScrapeHTMLParamsCountrySx WebWebScrapeHTMLParamsCountry = "sx"
	WebWebScrapeHTMLParamsCountrySy WebWebScrapeHTMLParamsCountry = "sy"
	WebWebScrapeHTMLParamsCountrySz WebWebScrapeHTMLParamsCountry = "sz"
	WebWebScrapeHTMLParamsCountryTc WebWebScrapeHTMLParamsCountry = "tc"
	WebWebScrapeHTMLParamsCountryTd WebWebScrapeHTMLParamsCountry = "td"
	WebWebScrapeHTMLParamsCountryTg WebWebScrapeHTMLParamsCountry = "tg"
	WebWebScrapeHTMLParamsCountryTh WebWebScrapeHTMLParamsCountry = "th"
	WebWebScrapeHTMLParamsCountryTj WebWebScrapeHTMLParamsCountry = "tj"
	WebWebScrapeHTMLParamsCountryTl WebWebScrapeHTMLParamsCountry = "tl"
	WebWebScrapeHTMLParamsCountryTm WebWebScrapeHTMLParamsCountry = "tm"
	WebWebScrapeHTMLParamsCountryTn WebWebScrapeHTMLParamsCountry = "tn"
	WebWebScrapeHTMLParamsCountryTr WebWebScrapeHTMLParamsCountry = "tr"
	WebWebScrapeHTMLParamsCountryTt WebWebScrapeHTMLParamsCountry = "tt"
	WebWebScrapeHTMLParamsCountryTw WebWebScrapeHTMLParamsCountry = "tw"
	WebWebScrapeHTMLParamsCountryTz WebWebScrapeHTMLParamsCountry = "tz"
	WebWebScrapeHTMLParamsCountryUa WebWebScrapeHTMLParamsCountry = "ua"
	WebWebScrapeHTMLParamsCountryUg WebWebScrapeHTMLParamsCountry = "ug"
	WebWebScrapeHTMLParamsCountryUs WebWebScrapeHTMLParamsCountry = "us"
	WebWebScrapeHTMLParamsCountryUy WebWebScrapeHTMLParamsCountry = "uy"
	WebWebScrapeHTMLParamsCountryUz WebWebScrapeHTMLParamsCountry = "uz"
	WebWebScrapeHTMLParamsCountryVc WebWebScrapeHTMLParamsCountry = "vc"
	WebWebScrapeHTMLParamsCountryVe WebWebScrapeHTMLParamsCountry = "ve"
	WebWebScrapeHTMLParamsCountryVg WebWebScrapeHTMLParamsCountry = "vg"
	WebWebScrapeHTMLParamsCountryVi WebWebScrapeHTMLParamsCountry = "vi"
	WebWebScrapeHTMLParamsCountryVn WebWebScrapeHTMLParamsCountry = "vn"
	WebWebScrapeHTMLParamsCountryYe WebWebScrapeHTMLParamsCountry = "ye"
	WebWebScrapeHTMLParamsCountryYt WebWebScrapeHTMLParamsCountry = "yt"
	WebWebScrapeHTMLParamsCountryZa WebWebScrapeHTMLParamsCountry = "za"
	WebWebScrapeHTMLParamsCountryZm WebWebScrapeHTMLParamsCountry = "zm"
	WebWebScrapeHTMLParamsCountryZw WebWebScrapeHTMLParamsCountry = "zw"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebWebScrapeHTMLParamsIncludeFramesUnion struct {
	OfBool param.Opt[bool] `query:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWebWebScrapeHTMLsIncludeFramesString)
	OfWebWebScrapeHTMLsIncludeFramesString param.Opt[string] `query:",omitzero,inline"`
	paramUnion
}

type WebWebScrapeHTMLParamsIncludeFramesString string

const (
	WebWebScrapeHTMLParamsIncludeFramesStringTrue  WebWebScrapeHTMLParamsIncludeFramesString = "true"
	WebWebScrapeHTMLParamsIncludeFramesStringFalse WebWebScrapeHTMLParamsIncludeFramesString = "false"
)

// PDF parsing controls. Use start/end to limit text extraction and embedded-image
// detection/OCR to an inclusive 1-based page range.
type WebWebScrapeHTMLParamsPdf struct {
	// Last 1-based PDF page to parse. When omitted, parsing ends at the last page.
	// Must be greater than or equal to start when both are provided.
	End param.Opt[int64] `query:"end,omitzero" json:"-"`
	// First 1-based PDF page to parse. When omitted, parsing starts at the first page.
	Start param.Opt[int64] `query:"start,omitzero" json:"-"`
	// When true, detect and OCR images embedded in the selected PDF pages, inserting
	// recognized text at each image's position in page reading order while preserving
	// the PDF text layer. This is separate from automatic scanned-PDF OCR fallback.
	Ocr WebWebScrapeHTMLParamsPdfOcrUnion `query:"ocr,omitzero" json:"-"`
	// When true, PDF URLs are fetched and parsed. When false, PDF URLs are skipped and
	// a 400 WEBSITE_ACCESS_ERROR is returned.
	ShouldParse WebWebScrapeHTMLParamsPdfShouldParseUnion `query:"shouldParse,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebWebScrapeHTMLParamsPdf]'s query parameters as
// `url.Values`.
func (r WebWebScrapeHTMLParamsPdf) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebWebScrapeHTMLParamsPdfOcrUnion struct {
	OfBool param.Opt[bool] `query:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWebWebScrapeHTMLsPdfOcrString)
	OfWebWebScrapeHTMLsPdfOcrString param.Opt[string] `query:",omitzero,inline"`
	paramUnion
}

type WebWebScrapeHTMLParamsPdfOcrString string

const (
	WebWebScrapeHTMLParamsPdfOcrStringTrue  WebWebScrapeHTMLParamsPdfOcrString = "true"
	WebWebScrapeHTMLParamsPdfOcrStringFalse WebWebScrapeHTMLParamsPdfOcrString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebWebScrapeHTMLParamsPdfShouldParseUnion struct {
	OfBool param.Opt[bool] `query:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWebWebScrapeHTMLsPdfShouldParseString)
	OfWebWebScrapeHTMLsPdfShouldParseString param.Opt[string] `query:",omitzero,inline"`
	paramUnion
}

type WebWebScrapeHTMLParamsPdfShouldParseString string

const (
	WebWebScrapeHTMLParamsPdfShouldParseStringTrue  WebWebScrapeHTMLParamsPdfShouldParseString = "true"
	WebWebScrapeHTMLParamsPdfShouldParseStringFalse WebWebScrapeHTMLParamsPdfShouldParseString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebWebScrapeHTMLParamsSettleAnimationsUnion struct {
	OfBool param.Opt[bool] `query:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWebWebScrapeHTMLsSettleAnimationsString)
	OfWebWebScrapeHTMLsSettleAnimationsString param.Opt[string] `query:",omitzero,inline"`
	paramUnion
}

type WebWebScrapeHTMLParamsSettleAnimationsString string

const (
	WebWebScrapeHTMLParamsSettleAnimationsStringTrue  WebWebScrapeHTMLParamsSettleAnimationsString = "true"
	WebWebScrapeHTMLParamsSettleAnimationsStringFalse WebWebScrapeHTMLParamsSettleAnimationsString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebWebScrapeHTMLParamsUseMainContentOnlyUnion struct {
	OfBool param.Opt[bool] `query:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWebWebScrapeHTMLsUseMainContentOnlyString)
	OfWebWebScrapeHTMLsUseMainContentOnlyString param.Opt[string] `query:",omitzero,inline"`
	paramUnion
}

type WebWebScrapeHTMLParamsUseMainContentOnlyString string

const (
	WebWebScrapeHTMLParamsUseMainContentOnlyStringTrue  WebWebScrapeHTMLParamsUseMainContentOnlyString = "true"
	WebWebScrapeHTMLParamsUseMainContentOnlyStringFalse WebWebScrapeHTMLParamsUseMainContentOnlyString = "false"
)

// Set to enabled to bypass shared caches and omit request and response content
// from retained usage logs. Requires zero data retention to be enabled for your
// organization (contact support@context.dev), otherwise the request fails with
// ZDR_NOT_ENABLED. Successful ZDR responses include X-Context-ZDR: true.
type WebWebScrapeHTMLParamsZdr string

const (
	WebWebScrapeHTMLParamsZdrEnabled  WebWebScrapeHTMLParamsZdr = "enabled"
	WebWebScrapeHTMLParamsZdrDisabled WebWebScrapeHTMLParamsZdr = "disabled"
)

type WebWebScrapeImagesParams struct {
	// Page URL to inspect. Must include http:// or https://.
	URL string `query:"url" api:"required" format:"uri" json:"-"`
	// Reuse a cached result this many milliseconds old or newer. Default: 86400000 (1
	// day). Set to 0 to bypass cache. Maximum: 2592000000 (30 days).
	MaxAgeMs param.Opt[int64] `query:"maxAgeMs,omitzero" json:"-"`
	// Optional browser wait time in milliseconds after initial page load before
	// collecting images. Min: 0. Max: 30000 (30 seconds).
	WaitForMs param.Opt[int64] `query:"waitForMs,omitzero" json:"-"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `query:"timeoutMS,omitzero" json:"-"`
	// Optional browser actions executed in array order after the page loads and before
	// content is captured. Requires a paid plan. Send a JSON array in the query
	// parameter. Maximum: 5 actions.
	Actions []WebWebScrapeImagesParamsActionUnion `query:"actions,omitzero" json:"-"`
	// Optional per-image processing, sent as deep-object query params such as
	// enrichment[resolution]=true.
	Enrichment WebWebScrapeImagesParamsEnrichment `query:"enrichment,omitzero" json:"-"`
	// When true, visually duplicate images are removed: every image is loaded and
	// perceptually hashed, and only the highest-resolution copy of each duplicate
	// group is kept. Images that cannot be downloaded or hashed are kept. Default:
	// false.
	Dedupe WebWebScrapeImagesParamsDedupeUnion `query:"dedupe,omitzero" json:"-"`
	// Optional outbound HTTP headers forwarded only to the target URL, sent as
	// deep-object query params such as headers[X-Custom]=value. When provided, caching
	// is bypassed: the result is neither read from nor written to cache.
	Headers map[string]string `query:"headers,omitzero" json:"-"`
	// Optional comma-separated caller-defined tags for tracking this request. Tags are
	// recorded on the request's usage log and can be used to filter usage on the
	// dashboard usage page. Up to 20 tags, each 1-50 characters.
	Tags []string `query:"tags,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebWebScrapeImagesParams]'s query parameters as
// `url.Values`.
func (r WebWebScrapeImagesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebWebScrapeImagesParamsActionUnion struct {
	OfWait    *WebWebScrapeImagesParamsActionWait    `query:",omitzero,inline"`
	OfPerform *WebWebScrapeImagesParamsActionPerform `query:",omitzero,inline"`
	paramUnion
}

func init() {
	apijson.RegisterUnion[WebWebScrapeImagesParamsActionUnion](
		"do",
		apijson.Discriminator[WebWebScrapeImagesParamsActionWait]("wait"),
		apijson.Discriminator[WebWebScrapeImagesParamsActionPerform]("perform"),
	)
}

// Pause for a fixed number of milliseconds before continuing to the next action.
//
// The properties Do, TimeMs are required.
type WebWebScrapeImagesParamsActionWait struct {
	TimeMs int64 `query:"timeMs" api:"required" json:"-"`
	// This field can be elided, and will marshal its zero value as "wait".
	Do constant.Wait `query:"do" json:"-" default:"wait"`
	paramObj
}

// URLQuery serializes [WebWebScrapeImagesParamsActionWait]'s query parameters as
// `url.Values`.
func (r WebWebScrapeImagesParamsActionWait) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Resolve and perform one natural-language browser action.
//
// The properties Action, Do are required.
type WebWebScrapeImagesParamsActionPerform struct {
	Action string `query:"action" api:"required" json:"-"`
	// This field can be elided, and will marshal its zero value as "perform".
	Do constant.Perform `query:"do" json:"-" default:"perform"`
	paramObj
}

// URLQuery serializes [WebWebScrapeImagesParamsActionPerform]'s query parameters
// as `url.Values`.
func (r WebWebScrapeImagesParamsActionPerform) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebWebScrapeImagesParamsDedupeUnion struct {
	OfBool param.Opt[bool] `query:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWebWebScrapeImagessDedupeString)
	OfWebWebScrapeImagessDedupeString param.Opt[string] `query:",omitzero,inline"`
	paramUnion
}

type WebWebScrapeImagesParamsDedupeString string

const (
	WebWebScrapeImagesParamsDedupeStringTrue  WebWebScrapeImagesParamsDedupeString = "true"
	WebWebScrapeImagesParamsDedupeStringFalse WebWebScrapeImagesParamsDedupeString = "false"
)

// Optional per-image processing, sent as deep-object query params such as
// enrichment[resolution]=true.
type WebWebScrapeImagesParamsEnrichment struct {
	// Per-image enrichment timeout in milliseconds. Default: 30000. Maximum: 60000.
	MaxTimePerMs param.Opt[int64] `query:"maxTimePerMs,omitzero" json:"-"`
	// Classify each image by visual asset type.
	Classification WebWebScrapeImagesParamsEnrichmentClassificationUnion `query:"classification,omitzero" json:"-"`
	// Host materializable images on the Brand.dev CDN and return their URL and MIME
	// type.
	HostedURL WebWebScrapeImagesParamsEnrichmentHostedURLUnion `query:"hostedUrl,omitzero" json:"-"`
	// Measure image width and height when possible.
	Resolution WebWebScrapeImagesParamsEnrichmentResolutionUnion `query:"resolution,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebWebScrapeImagesParamsEnrichment]'s query parameters as
// `url.Values`.
func (r WebWebScrapeImagesParamsEnrichment) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebWebScrapeImagesParamsEnrichmentClassificationUnion struct {
	OfBool param.Opt[bool] `query:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWebWebScrapeImagessEnrichmentClassificationString)
	OfWebWebScrapeImagessEnrichmentClassificationString param.Opt[string] `query:",omitzero,inline"`
	paramUnion
}

type WebWebScrapeImagesParamsEnrichmentClassificationString string

const (
	WebWebScrapeImagesParamsEnrichmentClassificationStringTrue  WebWebScrapeImagesParamsEnrichmentClassificationString = "true"
	WebWebScrapeImagesParamsEnrichmentClassificationStringFalse WebWebScrapeImagesParamsEnrichmentClassificationString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebWebScrapeImagesParamsEnrichmentHostedURLUnion struct {
	OfBool param.Opt[bool] `query:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWebWebScrapeImagessEnrichmentHostedURLString)
	OfWebWebScrapeImagessEnrichmentHostedURLString param.Opt[string] `query:",omitzero,inline"`
	paramUnion
}

type WebWebScrapeImagesParamsEnrichmentHostedURLString string

const (
	WebWebScrapeImagesParamsEnrichmentHostedURLStringTrue  WebWebScrapeImagesParamsEnrichmentHostedURLString = "true"
	WebWebScrapeImagesParamsEnrichmentHostedURLStringFalse WebWebScrapeImagesParamsEnrichmentHostedURLString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebWebScrapeImagesParamsEnrichmentResolutionUnion struct {
	OfBool param.Opt[bool] `query:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWebWebScrapeImagessEnrichmentResolutionString)
	OfWebWebScrapeImagessEnrichmentResolutionString param.Opt[string] `query:",omitzero,inline"`
	paramUnion
}

type WebWebScrapeImagesParamsEnrichmentResolutionString string

const (
	WebWebScrapeImagesParamsEnrichmentResolutionStringTrue  WebWebScrapeImagesParamsEnrichmentResolutionString = "true"
	WebWebScrapeImagesParamsEnrichmentResolutionStringFalse WebWebScrapeImagesParamsEnrichmentResolutionString = "false"
)

type WebWebScrapeMdParams struct {
	// Full URL to scrape into LLM usable Markdown (must include http:// or https://
	// protocol)
	URL string `query:"url" api:"required" format:"uri" json:"-"`
	// Return a cached result if a prior scrape for the same parameters exists and is
	// younger than this many milliseconds. Defaults to 1 day (86400000 ms) when
	// omitted. Max is 30 days (2592000000 ms). Set to 0 to always scrape fresh.
	MaxAgeMs param.Opt[int64] `query:"maxAgeMs,omitzero" json:"-"`
	// Optional browser wait time in milliseconds after initial page load before
	// converting the page to Markdown. Min: 0. Max: 30000 (30 seconds).
	WaitForMs param.Opt[int64] `query:"waitForMs,omitzero" json:"-"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `query:"timeoutMS,omitzero" json:"-"`
	// Optional browser actions executed in array order after the page loads and before
	// content is captured. Requires a paid plan. Send a JSON array in the query
	// parameter. Maximum: 5 actions.
	Actions []WebWebScrapeMdParamsActionUnion `query:"actions,omitzero" json:"-"`
	// CSS selectors to remove before conversion to Markdown. Applied after
	// includeSelectors. Exclusion takes precedence: an element matching both is
	// removed. Examples: "nav", "footer", ".ad-banner", "[aria-hidden=true]".
	ExcludeSelectors []string `query:"excludeSelectors,omitzero" json:"-"`
	// CSS selectors. When provided, only matching HTML subtrees (and their
	// descendants) are kept before conversion to Markdown. When omitted, the entire
	// document is kept. Examples: "article.main", "#content", "[role=main]".
	IncludeSelectors []string `query:"includeSelectors,omitzero" json:"-"`
	// Fetch the target page through a residential proxy in this country (ISO 3166-1
	// alpha-2).
	//
	// Any of "ad", "ae", "af", "ag", "ai", "al", "am", "ao", "ar", "at", "au", "aw",
	// "az", "ba", "bb", "bd", "be", "bf", "bg", "bh", "bi", "bj", "bm", "bn", "bo",
	// "bq", "br", "bs", "bw", "by", "bz", "ca", "cd", "cf", "cg", "ch", "ci", "cl",
	// "cm", "cn", "co", "cr", "cv", "cw", "cy", "cz", "de", "dj", "dk", "dm", "do",
	// "dz", "ec", "ee", "eg", "es", "et", "fi", "fj", "fr", "ga", "gb", "gd", "ge",
	// "gf", "gg", "gh", "gm", "gn", "gp", "gq", "gr", "gt", "gu", "gw", "gy", "hk",
	// "hn", "hr", "ht", "hu", "id", "ie", "il", "im", "in", "iq", "ir", "is", "it",
	// "je", "jm", "jo", "jp", "ke", "kg", "kh", "kn", "kr", "kw", "ky", "kz", "la",
	// "lb", "lc", "lk", "lr", "ls", "lt", "lu", "lv", "ly", "ma", "mc", "md", "me",
	// "mf", "mg", "mk", "ml", "mm", "mn", "mo", "mq", "mr", "mt", "mu", "mv", "mw",
	// "mx", "my", "mz", "na", "nc", "ne", "ng", "ni", "nl", "no", "np", "nz", "om",
	// "pa", "pe", "pf", "pg", "ph", "pk", "pl", "pr", "ps", "pt", "py", "qa", "re",
	// "ro", "rs", "ru", "rw", "sa", "sc", "sd", "se", "sg", "si", "sk", "sl", "sm",
	// "sn", "so", "sr", "ss", "st", "sv", "sx", "sy", "sz", "tc", "td", "tg", "th",
	// "tj", "tl", "tm", "tn", "tr", "tt", "tw", "tz", "ua", "ug", "us", "uy", "uz",
	// "vc", "ve", "vg", "vi", "vn", "ye", "yt", "za", "zm", "zw".
	Country WebWebScrapeMdParamsCountry `query:"country,omitzero" json:"-"`
	// Optional outbound HTTP headers forwarded only to the target URL, sent as
	// deep-object query params such as headers[X-Custom]=value. When provided, caching
	// is bypassed: the result is neither read from nor written to cache.
	Headers map[string]string `query:"headers,omitzero" json:"-"`
	// When true, the contents of iframes are rendered to Markdown.
	IncludeFrames WebWebScrapeMdParamsIncludeFramesUnion `query:"includeFrames,omitzero" json:"-"`
	// Include image references in Markdown output
	IncludeImages WebWebScrapeMdParamsIncludeImagesUnion `query:"includeImages,omitzero" json:"-"`
	// Preserve hyperlinks in Markdown output
	IncludeLinks WebWebScrapeMdParamsIncludeLinksUnion `query:"includeLinks,omitzero" json:"-"`
	// PDF parsing controls. Use start/end to limit text extraction and embedded-image
	// detection/OCR to an inclusive 1-based page range.
	Pdf WebWebScrapeMdParamsPdf `query:"pdf,omitzero" json:"-"`
	// When true, waits briefly for CSS and transition animations to settle before
	// converting to Markdown. Defaults to false. This adds a bit of latency in
	// exchange for more stable output on animated pages.
	SettleAnimations WebWebScrapeMdParamsSettleAnimationsUnion `query:"settleAnimations,omitzero" json:"-"`
	// Shorten base64-encoded image data in the Markdown output
	ShortenBase64Images WebWebScrapeMdParamsShortenBase64ImagesUnion `query:"shortenBase64Images,omitzero" json:"-"`
	// Optional comma-separated caller-defined tags for tracking this request. Tags are
	// recorded on the request's usage log and can be used to filter usage on the
	// dashboard usage page. Up to 20 tags, each 1-50 characters.
	Tags []string `query:"tags,omitzero" json:"-"`
	// Extract only the main content of the page, excluding headers, footers, sidebars,
	// and navigation
	UseMainContentOnly WebWebScrapeMdParamsUseMainContentOnlyUnion `query:"useMainContentOnly,omitzero" json:"-"`
	// Set to enabled to bypass shared caches and omit request and response content
	// from retained usage logs. Requires zero data retention to be enabled for your
	// organization (contact support@context.dev), otherwise the request fails with
	// ZDR_NOT_ENABLED. Successful ZDR responses include X-Context-ZDR: true.
	//
	// Any of "enabled", "disabled".
	Zdr WebWebScrapeMdParamsZdr `query:"zdr,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebWebScrapeMdParams]'s query parameters as `url.Values`.
func (r WebWebScrapeMdParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebWebScrapeMdParamsActionUnion struct {
	OfWait    *WebWebScrapeMdParamsActionWait    `query:",omitzero,inline"`
	OfPerform *WebWebScrapeMdParamsActionPerform `query:",omitzero,inline"`
	paramUnion
}

func init() {
	apijson.RegisterUnion[WebWebScrapeMdParamsActionUnion](
		"do",
		apijson.Discriminator[WebWebScrapeMdParamsActionWait]("wait"),
		apijson.Discriminator[WebWebScrapeMdParamsActionPerform]("perform"),
	)
}

// Pause for a fixed number of milliseconds before continuing to the next action.
//
// The properties Do, TimeMs are required.
type WebWebScrapeMdParamsActionWait struct {
	TimeMs int64 `query:"timeMs" api:"required" json:"-"`
	// This field can be elided, and will marshal its zero value as "wait".
	Do constant.Wait `query:"do" json:"-" default:"wait"`
	paramObj
}

// URLQuery serializes [WebWebScrapeMdParamsActionWait]'s query parameters as
// `url.Values`.
func (r WebWebScrapeMdParamsActionWait) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Resolve and perform one natural-language browser action.
//
// The properties Action, Do are required.
type WebWebScrapeMdParamsActionPerform struct {
	Action string `query:"action" api:"required" json:"-"`
	// This field can be elided, and will marshal its zero value as "perform".
	Do constant.Perform `query:"do" json:"-" default:"perform"`
	paramObj
}

// URLQuery serializes [WebWebScrapeMdParamsActionPerform]'s query parameters as
// `url.Values`.
func (r WebWebScrapeMdParamsActionPerform) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Fetch the target page through a residential proxy in this country (ISO 3166-1
// alpha-2).
type WebWebScrapeMdParamsCountry string

const (
	WebWebScrapeMdParamsCountryAd WebWebScrapeMdParamsCountry = "ad"
	WebWebScrapeMdParamsCountryAe WebWebScrapeMdParamsCountry = "ae"
	WebWebScrapeMdParamsCountryAf WebWebScrapeMdParamsCountry = "af"
	WebWebScrapeMdParamsCountryAg WebWebScrapeMdParamsCountry = "ag"
	WebWebScrapeMdParamsCountryAI WebWebScrapeMdParamsCountry = "ai"
	WebWebScrapeMdParamsCountryAl WebWebScrapeMdParamsCountry = "al"
	WebWebScrapeMdParamsCountryAm WebWebScrapeMdParamsCountry = "am"
	WebWebScrapeMdParamsCountryAo WebWebScrapeMdParamsCountry = "ao"
	WebWebScrapeMdParamsCountryAr WebWebScrapeMdParamsCountry = "ar"
	WebWebScrapeMdParamsCountryAt WebWebScrapeMdParamsCountry = "at"
	WebWebScrapeMdParamsCountryAu WebWebScrapeMdParamsCountry = "au"
	WebWebScrapeMdParamsCountryAw WebWebScrapeMdParamsCountry = "aw"
	WebWebScrapeMdParamsCountryAz WebWebScrapeMdParamsCountry = "az"
	WebWebScrapeMdParamsCountryBa WebWebScrapeMdParamsCountry = "ba"
	WebWebScrapeMdParamsCountryBb WebWebScrapeMdParamsCountry = "bb"
	WebWebScrapeMdParamsCountryBd WebWebScrapeMdParamsCountry = "bd"
	WebWebScrapeMdParamsCountryBe WebWebScrapeMdParamsCountry = "be"
	WebWebScrapeMdParamsCountryBf WebWebScrapeMdParamsCountry = "bf"
	WebWebScrapeMdParamsCountryBg WebWebScrapeMdParamsCountry = "bg"
	WebWebScrapeMdParamsCountryBh WebWebScrapeMdParamsCountry = "bh"
	WebWebScrapeMdParamsCountryBi WebWebScrapeMdParamsCountry = "bi"
	WebWebScrapeMdParamsCountryBj WebWebScrapeMdParamsCountry = "bj"
	WebWebScrapeMdParamsCountryBm WebWebScrapeMdParamsCountry = "bm"
	WebWebScrapeMdParamsCountryBn WebWebScrapeMdParamsCountry = "bn"
	WebWebScrapeMdParamsCountryBo WebWebScrapeMdParamsCountry = "bo"
	WebWebScrapeMdParamsCountryBq WebWebScrapeMdParamsCountry = "bq"
	WebWebScrapeMdParamsCountryBr WebWebScrapeMdParamsCountry = "br"
	WebWebScrapeMdParamsCountryBs WebWebScrapeMdParamsCountry = "bs"
	WebWebScrapeMdParamsCountryBw WebWebScrapeMdParamsCountry = "bw"
	WebWebScrapeMdParamsCountryBy WebWebScrapeMdParamsCountry = "by"
	WebWebScrapeMdParamsCountryBz WebWebScrapeMdParamsCountry = "bz"
	WebWebScrapeMdParamsCountryCa WebWebScrapeMdParamsCountry = "ca"
	WebWebScrapeMdParamsCountryCd WebWebScrapeMdParamsCountry = "cd"
	WebWebScrapeMdParamsCountryCf WebWebScrapeMdParamsCountry = "cf"
	WebWebScrapeMdParamsCountryCg WebWebScrapeMdParamsCountry = "cg"
	WebWebScrapeMdParamsCountryCh WebWebScrapeMdParamsCountry = "ch"
	WebWebScrapeMdParamsCountryCi WebWebScrapeMdParamsCountry = "ci"
	WebWebScrapeMdParamsCountryCl WebWebScrapeMdParamsCountry = "cl"
	WebWebScrapeMdParamsCountryCm WebWebScrapeMdParamsCountry = "cm"
	WebWebScrapeMdParamsCountryCn WebWebScrapeMdParamsCountry = "cn"
	WebWebScrapeMdParamsCountryCo WebWebScrapeMdParamsCountry = "co"
	WebWebScrapeMdParamsCountryCr WebWebScrapeMdParamsCountry = "cr"
	WebWebScrapeMdParamsCountryCv WebWebScrapeMdParamsCountry = "cv"
	WebWebScrapeMdParamsCountryCw WebWebScrapeMdParamsCountry = "cw"
	WebWebScrapeMdParamsCountryCy WebWebScrapeMdParamsCountry = "cy"
	WebWebScrapeMdParamsCountryCz WebWebScrapeMdParamsCountry = "cz"
	WebWebScrapeMdParamsCountryDe WebWebScrapeMdParamsCountry = "de"
	WebWebScrapeMdParamsCountryDj WebWebScrapeMdParamsCountry = "dj"
	WebWebScrapeMdParamsCountryDk WebWebScrapeMdParamsCountry = "dk"
	WebWebScrapeMdParamsCountryDm WebWebScrapeMdParamsCountry = "dm"
	WebWebScrapeMdParamsCountryDo WebWebScrapeMdParamsCountry = "do"
	WebWebScrapeMdParamsCountryDz WebWebScrapeMdParamsCountry = "dz"
	WebWebScrapeMdParamsCountryEc WebWebScrapeMdParamsCountry = "ec"
	WebWebScrapeMdParamsCountryEe WebWebScrapeMdParamsCountry = "ee"
	WebWebScrapeMdParamsCountryEg WebWebScrapeMdParamsCountry = "eg"
	WebWebScrapeMdParamsCountryEs WebWebScrapeMdParamsCountry = "es"
	WebWebScrapeMdParamsCountryEt WebWebScrapeMdParamsCountry = "et"
	WebWebScrapeMdParamsCountryFi WebWebScrapeMdParamsCountry = "fi"
	WebWebScrapeMdParamsCountryFj WebWebScrapeMdParamsCountry = "fj"
	WebWebScrapeMdParamsCountryFr WebWebScrapeMdParamsCountry = "fr"
	WebWebScrapeMdParamsCountryGa WebWebScrapeMdParamsCountry = "ga"
	WebWebScrapeMdParamsCountryGB WebWebScrapeMdParamsCountry = "gb"
	WebWebScrapeMdParamsCountryGd WebWebScrapeMdParamsCountry = "gd"
	WebWebScrapeMdParamsCountryGe WebWebScrapeMdParamsCountry = "ge"
	WebWebScrapeMdParamsCountryGf WebWebScrapeMdParamsCountry = "gf"
	WebWebScrapeMdParamsCountryGg WebWebScrapeMdParamsCountry = "gg"
	WebWebScrapeMdParamsCountryGh WebWebScrapeMdParamsCountry = "gh"
	WebWebScrapeMdParamsCountryGm WebWebScrapeMdParamsCountry = "gm"
	WebWebScrapeMdParamsCountryGn WebWebScrapeMdParamsCountry = "gn"
	WebWebScrapeMdParamsCountryGp WebWebScrapeMdParamsCountry = "gp"
	WebWebScrapeMdParamsCountryGq WebWebScrapeMdParamsCountry = "gq"
	WebWebScrapeMdParamsCountryGr WebWebScrapeMdParamsCountry = "gr"
	WebWebScrapeMdParamsCountryGt WebWebScrapeMdParamsCountry = "gt"
	WebWebScrapeMdParamsCountryGu WebWebScrapeMdParamsCountry = "gu"
	WebWebScrapeMdParamsCountryGw WebWebScrapeMdParamsCountry = "gw"
	WebWebScrapeMdParamsCountryGy WebWebScrapeMdParamsCountry = "gy"
	WebWebScrapeMdParamsCountryHk WebWebScrapeMdParamsCountry = "hk"
	WebWebScrapeMdParamsCountryHn WebWebScrapeMdParamsCountry = "hn"
	WebWebScrapeMdParamsCountryHr WebWebScrapeMdParamsCountry = "hr"
	WebWebScrapeMdParamsCountryHt WebWebScrapeMdParamsCountry = "ht"
	WebWebScrapeMdParamsCountryHu WebWebScrapeMdParamsCountry = "hu"
	WebWebScrapeMdParamsCountryID WebWebScrapeMdParamsCountry = "id"
	WebWebScrapeMdParamsCountryIe WebWebScrapeMdParamsCountry = "ie"
	WebWebScrapeMdParamsCountryIl WebWebScrapeMdParamsCountry = "il"
	WebWebScrapeMdParamsCountryIm WebWebScrapeMdParamsCountry = "im"
	WebWebScrapeMdParamsCountryIn WebWebScrapeMdParamsCountry = "in"
	WebWebScrapeMdParamsCountryIq WebWebScrapeMdParamsCountry = "iq"
	WebWebScrapeMdParamsCountryIr WebWebScrapeMdParamsCountry = "ir"
	WebWebScrapeMdParamsCountryIs WebWebScrapeMdParamsCountry = "is"
	WebWebScrapeMdParamsCountryIt WebWebScrapeMdParamsCountry = "it"
	WebWebScrapeMdParamsCountryJe WebWebScrapeMdParamsCountry = "je"
	WebWebScrapeMdParamsCountryJm WebWebScrapeMdParamsCountry = "jm"
	WebWebScrapeMdParamsCountryJo WebWebScrapeMdParamsCountry = "jo"
	WebWebScrapeMdParamsCountryJp WebWebScrapeMdParamsCountry = "jp"
	WebWebScrapeMdParamsCountryKe WebWebScrapeMdParamsCountry = "ke"
	WebWebScrapeMdParamsCountryKg WebWebScrapeMdParamsCountry = "kg"
	WebWebScrapeMdParamsCountryKh WebWebScrapeMdParamsCountry = "kh"
	WebWebScrapeMdParamsCountryKn WebWebScrapeMdParamsCountry = "kn"
	WebWebScrapeMdParamsCountryKr WebWebScrapeMdParamsCountry = "kr"
	WebWebScrapeMdParamsCountryKw WebWebScrapeMdParamsCountry = "kw"
	WebWebScrapeMdParamsCountryKy WebWebScrapeMdParamsCountry = "ky"
	WebWebScrapeMdParamsCountryKz WebWebScrapeMdParamsCountry = "kz"
	WebWebScrapeMdParamsCountryLa WebWebScrapeMdParamsCountry = "la"
	WebWebScrapeMdParamsCountryLb WebWebScrapeMdParamsCountry = "lb"
	WebWebScrapeMdParamsCountryLc WebWebScrapeMdParamsCountry = "lc"
	WebWebScrapeMdParamsCountryLk WebWebScrapeMdParamsCountry = "lk"
	WebWebScrapeMdParamsCountryLr WebWebScrapeMdParamsCountry = "lr"
	WebWebScrapeMdParamsCountryLs WebWebScrapeMdParamsCountry = "ls"
	WebWebScrapeMdParamsCountryLt WebWebScrapeMdParamsCountry = "lt"
	WebWebScrapeMdParamsCountryLu WebWebScrapeMdParamsCountry = "lu"
	WebWebScrapeMdParamsCountryLv WebWebScrapeMdParamsCountry = "lv"
	WebWebScrapeMdParamsCountryLy WebWebScrapeMdParamsCountry = "ly"
	WebWebScrapeMdParamsCountryMa WebWebScrapeMdParamsCountry = "ma"
	WebWebScrapeMdParamsCountryMc WebWebScrapeMdParamsCountry = "mc"
	WebWebScrapeMdParamsCountryMd WebWebScrapeMdParamsCountry = "md"
	WebWebScrapeMdParamsCountryMe WebWebScrapeMdParamsCountry = "me"
	WebWebScrapeMdParamsCountryMf WebWebScrapeMdParamsCountry = "mf"
	WebWebScrapeMdParamsCountryMg WebWebScrapeMdParamsCountry = "mg"
	WebWebScrapeMdParamsCountryMk WebWebScrapeMdParamsCountry = "mk"
	WebWebScrapeMdParamsCountryMl WebWebScrapeMdParamsCountry = "ml"
	WebWebScrapeMdParamsCountryMm WebWebScrapeMdParamsCountry = "mm"
	WebWebScrapeMdParamsCountryMn WebWebScrapeMdParamsCountry = "mn"
	WebWebScrapeMdParamsCountryMo WebWebScrapeMdParamsCountry = "mo"
	WebWebScrapeMdParamsCountryMq WebWebScrapeMdParamsCountry = "mq"
	WebWebScrapeMdParamsCountryMr WebWebScrapeMdParamsCountry = "mr"
	WebWebScrapeMdParamsCountryMt WebWebScrapeMdParamsCountry = "mt"
	WebWebScrapeMdParamsCountryMu WebWebScrapeMdParamsCountry = "mu"
	WebWebScrapeMdParamsCountryMv WebWebScrapeMdParamsCountry = "mv"
	WebWebScrapeMdParamsCountryMw WebWebScrapeMdParamsCountry = "mw"
	WebWebScrapeMdParamsCountryMx WebWebScrapeMdParamsCountry = "mx"
	WebWebScrapeMdParamsCountryMy WebWebScrapeMdParamsCountry = "my"
	WebWebScrapeMdParamsCountryMz WebWebScrapeMdParamsCountry = "mz"
	WebWebScrapeMdParamsCountryNa WebWebScrapeMdParamsCountry = "na"
	WebWebScrapeMdParamsCountryNc WebWebScrapeMdParamsCountry = "nc"
	WebWebScrapeMdParamsCountryNe WebWebScrapeMdParamsCountry = "ne"
	WebWebScrapeMdParamsCountryNg WebWebScrapeMdParamsCountry = "ng"
	WebWebScrapeMdParamsCountryNi WebWebScrapeMdParamsCountry = "ni"
	WebWebScrapeMdParamsCountryNl WebWebScrapeMdParamsCountry = "nl"
	WebWebScrapeMdParamsCountryNo WebWebScrapeMdParamsCountry = "no"
	WebWebScrapeMdParamsCountryNp WebWebScrapeMdParamsCountry = "np"
	WebWebScrapeMdParamsCountryNz WebWebScrapeMdParamsCountry = "nz"
	WebWebScrapeMdParamsCountryOm WebWebScrapeMdParamsCountry = "om"
	WebWebScrapeMdParamsCountryPa WebWebScrapeMdParamsCountry = "pa"
	WebWebScrapeMdParamsCountryPe WebWebScrapeMdParamsCountry = "pe"
	WebWebScrapeMdParamsCountryPf WebWebScrapeMdParamsCountry = "pf"
	WebWebScrapeMdParamsCountryPg WebWebScrapeMdParamsCountry = "pg"
	WebWebScrapeMdParamsCountryPh WebWebScrapeMdParamsCountry = "ph"
	WebWebScrapeMdParamsCountryPk WebWebScrapeMdParamsCountry = "pk"
	WebWebScrapeMdParamsCountryPl WebWebScrapeMdParamsCountry = "pl"
	WebWebScrapeMdParamsCountryPr WebWebScrapeMdParamsCountry = "pr"
	WebWebScrapeMdParamsCountryPs WebWebScrapeMdParamsCountry = "ps"
	WebWebScrapeMdParamsCountryPt WebWebScrapeMdParamsCountry = "pt"
	WebWebScrapeMdParamsCountryPy WebWebScrapeMdParamsCountry = "py"
	WebWebScrapeMdParamsCountryQa WebWebScrapeMdParamsCountry = "qa"
	WebWebScrapeMdParamsCountryRe WebWebScrapeMdParamsCountry = "re"
	WebWebScrapeMdParamsCountryRo WebWebScrapeMdParamsCountry = "ro"
	WebWebScrapeMdParamsCountryRs WebWebScrapeMdParamsCountry = "rs"
	WebWebScrapeMdParamsCountryRu WebWebScrapeMdParamsCountry = "ru"
	WebWebScrapeMdParamsCountryRw WebWebScrapeMdParamsCountry = "rw"
	WebWebScrapeMdParamsCountrySa WebWebScrapeMdParamsCountry = "sa"
	WebWebScrapeMdParamsCountrySc WebWebScrapeMdParamsCountry = "sc"
	WebWebScrapeMdParamsCountrySd WebWebScrapeMdParamsCountry = "sd"
	WebWebScrapeMdParamsCountrySe WebWebScrapeMdParamsCountry = "se"
	WebWebScrapeMdParamsCountrySg WebWebScrapeMdParamsCountry = "sg"
	WebWebScrapeMdParamsCountrySi WebWebScrapeMdParamsCountry = "si"
	WebWebScrapeMdParamsCountrySk WebWebScrapeMdParamsCountry = "sk"
	WebWebScrapeMdParamsCountrySl WebWebScrapeMdParamsCountry = "sl"
	WebWebScrapeMdParamsCountrySm WebWebScrapeMdParamsCountry = "sm"
	WebWebScrapeMdParamsCountrySn WebWebScrapeMdParamsCountry = "sn"
	WebWebScrapeMdParamsCountrySo WebWebScrapeMdParamsCountry = "so"
	WebWebScrapeMdParamsCountrySr WebWebScrapeMdParamsCountry = "sr"
	WebWebScrapeMdParamsCountrySS WebWebScrapeMdParamsCountry = "ss"
	WebWebScrapeMdParamsCountrySt WebWebScrapeMdParamsCountry = "st"
	WebWebScrapeMdParamsCountrySv WebWebScrapeMdParamsCountry = "sv"
	WebWebScrapeMdParamsCountrySx WebWebScrapeMdParamsCountry = "sx"
	WebWebScrapeMdParamsCountrySy WebWebScrapeMdParamsCountry = "sy"
	WebWebScrapeMdParamsCountrySz WebWebScrapeMdParamsCountry = "sz"
	WebWebScrapeMdParamsCountryTc WebWebScrapeMdParamsCountry = "tc"
	WebWebScrapeMdParamsCountryTd WebWebScrapeMdParamsCountry = "td"
	WebWebScrapeMdParamsCountryTg WebWebScrapeMdParamsCountry = "tg"
	WebWebScrapeMdParamsCountryTh WebWebScrapeMdParamsCountry = "th"
	WebWebScrapeMdParamsCountryTj WebWebScrapeMdParamsCountry = "tj"
	WebWebScrapeMdParamsCountryTl WebWebScrapeMdParamsCountry = "tl"
	WebWebScrapeMdParamsCountryTm WebWebScrapeMdParamsCountry = "tm"
	WebWebScrapeMdParamsCountryTn WebWebScrapeMdParamsCountry = "tn"
	WebWebScrapeMdParamsCountryTr WebWebScrapeMdParamsCountry = "tr"
	WebWebScrapeMdParamsCountryTt WebWebScrapeMdParamsCountry = "tt"
	WebWebScrapeMdParamsCountryTw WebWebScrapeMdParamsCountry = "tw"
	WebWebScrapeMdParamsCountryTz WebWebScrapeMdParamsCountry = "tz"
	WebWebScrapeMdParamsCountryUa WebWebScrapeMdParamsCountry = "ua"
	WebWebScrapeMdParamsCountryUg WebWebScrapeMdParamsCountry = "ug"
	WebWebScrapeMdParamsCountryUs WebWebScrapeMdParamsCountry = "us"
	WebWebScrapeMdParamsCountryUy WebWebScrapeMdParamsCountry = "uy"
	WebWebScrapeMdParamsCountryUz WebWebScrapeMdParamsCountry = "uz"
	WebWebScrapeMdParamsCountryVc WebWebScrapeMdParamsCountry = "vc"
	WebWebScrapeMdParamsCountryVe WebWebScrapeMdParamsCountry = "ve"
	WebWebScrapeMdParamsCountryVg WebWebScrapeMdParamsCountry = "vg"
	WebWebScrapeMdParamsCountryVi WebWebScrapeMdParamsCountry = "vi"
	WebWebScrapeMdParamsCountryVn WebWebScrapeMdParamsCountry = "vn"
	WebWebScrapeMdParamsCountryYe WebWebScrapeMdParamsCountry = "ye"
	WebWebScrapeMdParamsCountryYt WebWebScrapeMdParamsCountry = "yt"
	WebWebScrapeMdParamsCountryZa WebWebScrapeMdParamsCountry = "za"
	WebWebScrapeMdParamsCountryZm WebWebScrapeMdParamsCountry = "zm"
	WebWebScrapeMdParamsCountryZw WebWebScrapeMdParamsCountry = "zw"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebWebScrapeMdParamsIncludeFramesUnion struct {
	OfBool param.Opt[bool] `query:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWebWebScrapeMdsIncludeFramesString)
	OfWebWebScrapeMdsIncludeFramesString param.Opt[string] `query:",omitzero,inline"`
	paramUnion
}

type WebWebScrapeMdParamsIncludeFramesString string

const (
	WebWebScrapeMdParamsIncludeFramesStringTrue  WebWebScrapeMdParamsIncludeFramesString = "true"
	WebWebScrapeMdParamsIncludeFramesStringFalse WebWebScrapeMdParamsIncludeFramesString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebWebScrapeMdParamsIncludeImagesUnion struct {
	OfBool param.Opt[bool] `query:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWebWebScrapeMdsIncludeImagesString)
	OfWebWebScrapeMdsIncludeImagesString param.Opt[string] `query:",omitzero,inline"`
	paramUnion
}

type WebWebScrapeMdParamsIncludeImagesString string

const (
	WebWebScrapeMdParamsIncludeImagesStringTrue  WebWebScrapeMdParamsIncludeImagesString = "true"
	WebWebScrapeMdParamsIncludeImagesStringFalse WebWebScrapeMdParamsIncludeImagesString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebWebScrapeMdParamsIncludeLinksUnion struct {
	OfBool param.Opt[bool] `query:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWebWebScrapeMdsIncludeLinksString)
	OfWebWebScrapeMdsIncludeLinksString param.Opt[string] `query:",omitzero,inline"`
	paramUnion
}

type WebWebScrapeMdParamsIncludeLinksString string

const (
	WebWebScrapeMdParamsIncludeLinksStringTrue  WebWebScrapeMdParamsIncludeLinksString = "true"
	WebWebScrapeMdParamsIncludeLinksStringFalse WebWebScrapeMdParamsIncludeLinksString = "false"
)

// PDF parsing controls. Use start/end to limit text extraction and embedded-image
// detection/OCR to an inclusive 1-based page range.
type WebWebScrapeMdParamsPdf struct {
	// Last 1-based PDF page to parse. When omitted, parsing ends at the last page.
	// Must be greater than or equal to start when both are provided.
	End param.Opt[int64] `query:"end,omitzero" json:"-"`
	// First 1-based PDF page to parse. When omitted, parsing starts at the first page.
	Start param.Opt[int64] `query:"start,omitzero" json:"-"`
	// When true, detect and OCR images embedded in the selected PDF pages, inserting
	// recognized text at each image's position in page reading order while preserving
	// the PDF text layer. This is separate from automatic scanned-PDF OCR fallback.
	Ocr WebWebScrapeMdParamsPdfOcrUnion `query:"ocr,omitzero" json:"-"`
	// When true, PDF URLs are fetched and parsed. When false, PDF URLs are skipped and
	// a 400 WEBSITE_ACCESS_ERROR is returned.
	ShouldParse WebWebScrapeMdParamsPdfShouldParseUnion `query:"shouldParse,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebWebScrapeMdParamsPdf]'s query parameters as
// `url.Values`.
func (r WebWebScrapeMdParamsPdf) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebWebScrapeMdParamsPdfOcrUnion struct {
	OfBool param.Opt[bool] `query:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWebWebScrapeMdsPdfOcrString)
	OfWebWebScrapeMdsPdfOcrString param.Opt[string] `query:",omitzero,inline"`
	paramUnion
}

type WebWebScrapeMdParamsPdfOcrString string

const (
	WebWebScrapeMdParamsPdfOcrStringTrue  WebWebScrapeMdParamsPdfOcrString = "true"
	WebWebScrapeMdParamsPdfOcrStringFalse WebWebScrapeMdParamsPdfOcrString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebWebScrapeMdParamsPdfShouldParseUnion struct {
	OfBool param.Opt[bool] `query:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWebWebScrapeMdsPdfShouldParseString)
	OfWebWebScrapeMdsPdfShouldParseString param.Opt[string] `query:",omitzero,inline"`
	paramUnion
}

type WebWebScrapeMdParamsPdfShouldParseString string

const (
	WebWebScrapeMdParamsPdfShouldParseStringTrue  WebWebScrapeMdParamsPdfShouldParseString = "true"
	WebWebScrapeMdParamsPdfShouldParseStringFalse WebWebScrapeMdParamsPdfShouldParseString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebWebScrapeMdParamsSettleAnimationsUnion struct {
	OfBool param.Opt[bool] `query:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWebWebScrapeMdsSettleAnimationsString)
	OfWebWebScrapeMdsSettleAnimationsString param.Opt[string] `query:",omitzero,inline"`
	paramUnion
}

type WebWebScrapeMdParamsSettleAnimationsString string

const (
	WebWebScrapeMdParamsSettleAnimationsStringTrue  WebWebScrapeMdParamsSettleAnimationsString = "true"
	WebWebScrapeMdParamsSettleAnimationsStringFalse WebWebScrapeMdParamsSettleAnimationsString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebWebScrapeMdParamsShortenBase64ImagesUnion struct {
	OfBool param.Opt[bool] `query:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWebWebScrapeMdsShortenBase64ImagesString)
	OfWebWebScrapeMdsShortenBase64ImagesString param.Opt[string] `query:",omitzero,inline"`
	paramUnion
}

type WebWebScrapeMdParamsShortenBase64ImagesString string

const (
	WebWebScrapeMdParamsShortenBase64ImagesStringTrue  WebWebScrapeMdParamsShortenBase64ImagesString = "true"
	WebWebScrapeMdParamsShortenBase64ImagesStringFalse WebWebScrapeMdParamsShortenBase64ImagesString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebWebScrapeMdParamsUseMainContentOnlyUnion struct {
	OfBool param.Opt[bool] `query:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWebWebScrapeMdsUseMainContentOnlyString)
	OfWebWebScrapeMdsUseMainContentOnlyString param.Opt[string] `query:",omitzero,inline"`
	paramUnion
}

type WebWebScrapeMdParamsUseMainContentOnlyString string

const (
	WebWebScrapeMdParamsUseMainContentOnlyStringTrue  WebWebScrapeMdParamsUseMainContentOnlyString = "true"
	WebWebScrapeMdParamsUseMainContentOnlyStringFalse WebWebScrapeMdParamsUseMainContentOnlyString = "false"
)

// Set to enabled to bypass shared caches and omit request and response content
// from retained usage logs. Requires zero data retention to be enabled for your
// organization (contact support@context.dev), otherwise the request fails with
// ZDR_NOT_ENABLED. Successful ZDR responses include X-Context-ZDR: true.
type WebWebScrapeMdParamsZdr string

const (
	WebWebScrapeMdParamsZdrEnabled  WebWebScrapeMdParamsZdr = "enabled"
	WebWebScrapeMdParamsZdrDisabled WebWebScrapeMdParamsZdr = "disabled"
)

type WebWebScrapeSitemapParams struct {
	// Domain to build a sitemap for
	Domain string `query:"domain" api:"required" json:"-"`
	// Maximum number of links to return from the sitemap crawl. Defaults to 10,000.
	// Minimum is 1, maximum is 100,000.
	MaxLinks param.Opt[int64] `query:"maxLinks,omitzero" json:"-"`
	// Optional explicit sitemap URL. When provided, exactly this sitemap is crawled
	// instead of discovering the domain's sitemaps.
	SitemapURL param.Opt[string] `query:"sitemapUrl,omitzero" format:"uri" json:"-"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `query:"timeoutMS,omitzero" json:"-"`
	// Optional RE2-compatible regex pattern. Only URLs matching this pattern are
	// returned and counted against maxLinks.
	URLRegex param.Opt[string] `query:"urlRegex,omitzero" json:"-"`
	// Optional outbound HTTP headers forwarded only to the target URL, sent as
	// deep-object query params such as headers[X-Custom]=value. When provided, caching
	// is bypassed: the result is neither read from nor written to cache.
	Headers map[string]string `query:"headers,omitzero" json:"-"`
	// Optional comma-separated caller-defined tags for tracking this request. Tags are
	// recorded on the request's usage log and can be used to filter usage on the
	// dashboard usage page. Up to 20 tags, each 1-50 characters.
	Tags []string `query:"tags,omitzero" json:"-"`
	// Set to enabled to bypass shared caches and omit request and response content
	// from retained usage logs. Requires zero data retention to be enabled for your
	// organization (contact support@context.dev), otherwise the request fails with
	// ZDR_NOT_ENABLED. Successful ZDR responses include X-Context-ZDR: true.
	//
	// Any of "enabled", "disabled".
	Zdr WebWebScrapeSitemapParamsZdr `query:"zdr,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebWebScrapeSitemapParams]'s query parameters as
// `url.Values`.
func (r WebWebScrapeSitemapParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Set to enabled to bypass shared caches and omit request and response content
// from retained usage logs. Requires zero data retention to be enabled for your
// organization (contact support@context.dev), otherwise the request fails with
// ZDR_NOT_ENABLED. Successful ZDR responses include X-Context-ZDR: true.
type WebWebScrapeSitemapParamsZdr string

const (
	WebWebScrapeSitemapParamsZdrEnabled  WebWebScrapeSitemapParamsZdr = "enabled"
	WebWebScrapeSitemapParamsZdrDisabled WebWebScrapeSitemapParamsZdr = "disabled"
)
