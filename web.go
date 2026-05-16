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

// Scrapes the given URL and returns the raw HTML content of the page.
func (r *WebService) WebScrapeHTML(ctx context.Context, query WebWebScrapeHTMLParams, opts ...option.RequestOption) (res *WebWebScrapeHTMLResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "web/scrape/html"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Extract image assets from a web page, including standard URLs, inline SVGs, data
// URIs, responsive image sources, metadata, CSS backgrounds, video posters, and
// embeds. The base request costs 1 credit; enrichment costs 1 credit per returned
// image.
func (r *WebService) WebScrapeImages(ctx context.Context, query WebWebScrapeImagesParams, opts ...option.RequestOption) (res *WebWebScrapeImagesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "web/scrape/images"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Scrapes the given URL into LLM usable Markdown.
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Domain      respjson.Field
		Fonts       respjson.Field
		Status      respjson.Field
		FontLinks   respjson.Field
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

type WebExtractStyleguideResponse struct {
	// HTTP status code
	Code int64 `json:"code"`
	// The normalized domain that was processed
	Domain string `json:"domain"`
	// Status of the response, e.g., 'ok'
	Status string `json:"status"`
	// Comprehensive styleguide data extracted from the website
	Styleguide WebExtractStyleguideResponseStyleguide `json:"styleguide"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Domain      respjson.Field
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
	// Public URL of the uploaded screenshot image
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Query       respjson.Field
		Results     respjson.Field
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

type WebWebCrawlMdResponse struct {
	Metadata WebWebCrawlMdResponseMetadata `json:"metadata" api:"required"`
	Results  []WebWebCrawlMdResponseResult `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Results     respjson.Field
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
	// HTTP status code of the response
	StatusCode int64 `json:"statusCode" api:"required"`
	// true if the page was fetched and parsed successfully
	Success bool `json:"success" api:"required"`
	// The page's <title> content (empty string if unavailable)
	Title string `json:"title" api:"required"`
	// The URL that was fetched
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CrawlDepth  respjson.Field
		StatusCode  respjson.Field
		Success     respjson.Field
		Title       respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebWebCrawlMdResponseResultMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebWebCrawlMdResponseResultMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebWebScrapeHTMLResponse struct {
	// Raw HTML content of the page
	HTML string `json:"html" api:"required"`
	// Indicates success
	//
	// Any of true.
	Success bool `json:"success" api:"required"`
	// The URL that was scraped
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HTML        respjson.Field
		Success     respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebWebScrapeHTMLResponse) RawJSON() string { return r.JSON.raw }
func (r *WebWebScrapeHTMLResponse) UnmarshalJSON(data []byte) error {
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Images      respjson.Field
		Success     respjson.Field
		URL         respjson.Field
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

type WebWebScrapeMdResponse struct {
	// Page content converted to GitHub Flavored Markdown
	Markdown string `json:"markdown" api:"required"`
	// Indicates success
	//
	// Any of true.
	Success bool `json:"success" api:"required"`
	// The URL that was scraped
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Markdown    respjson.Field
		Success     respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebWebScrapeMdResponse) RawJSON() string { return r.JSON.raw }
func (r *WebWebScrapeMdResponse) UnmarshalJSON(data []byte) error {
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		Meta        respjson.Field
		Success     respjson.Field
		URLs        respjson.Field
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

type WebExtractFontsParams struct {
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

type WebScreenshotParams struct {
	// A specific URL to screenshot directly, bypassing domain resolution (e.g.,
	// 'https://example.com/pricing'). When provided, the screenshot is taken of this
	// exact URL. You must provide either 'domain' or 'directUrl', but not both.
	DirectURL param.Opt[string] `query:"directUrl,omitzero" format:"uri" json:"-"`
	// Domain name to take screenshot of (e.g., 'example.com', 'google.com'). The
	// domain will be automatically normalized and validated. You must provide either
	// 'domain' or 'directUrl', but not both.
	Domain param.Opt[string] `query:"domain,omitzero" json:"-"`
	// Return a cached screenshot if a prior screenshot for the same parameters exists
	// and is younger than this many milliseconds. Defaults to 1 day (86400000 ms) when
	// omitted. Max is 30 days (2592000000 ms). Set to 0 to always capture fresh.
	MaxAgeMs param.Opt[int64] `query:"maxAgeMs,omitzero" json:"-"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `query:"timeoutMS,omitzero" json:"-"`
	// Optional browser wait time in milliseconds after initial page load before taking
	// the screenshot. Min: 0. Max: 30000 (30 seconds). Defaults to 3000 ms when
	// omitted.
	WaitForMs param.Opt[int64] `query:"waitForMs,omitzero" json:"-"`
	// Optional parameter to determine screenshot type. If 'true', takes a full page
	// screenshot capturing all content. If 'false' or not provided, takes a viewport
	// screenshot (standard browser view).
	//
	// Any of "true", "false".
	FullScreenshot WebScreenshotParamsFullScreenshot `query:"fullScreenshot,omitzero" json:"-"`
	// Optional parameter to control cookie/consent popup handling. If 'true', we
	// dismiss cookie banner before capture. If 'false' or not provided, captures the
	// page without that step.
	//
	// Any of "true", "false".
	HandleCookiePopup WebScreenshotParamsHandleCookiePopup `query:"handleCookiePopup,omitzero" json:"-"`
	// Optional parameter to specify which page type to screenshot. If provided, the
	// system will scrape the domain's links and use heuristics to find the most
	// appropriate URL for the specified page type (30 supported languages). If not
	// provided, screenshots the main domain landing page. Only applicable when using
	// 'domain', not 'directUrl'.
	//
	// Any of "login", "signup", "blog", "careers", "pricing", "terms", "privacy",
	// "contact".
	Page WebScreenshotParamsPage `query:"page,omitzero" json:"-"`
	// Optional browser viewport dimensions for the screenshot. Defaults to 1920x1080.
	Viewport WebScreenshotParamsViewport `query:"viewport,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebScreenshotParams]'s query parameters as `url.Values`.
func (r WebScreenshotParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Optional parameter to determine screenshot type. If 'true', takes a full page
// screenshot capturing all content. If 'false' or not provided, takes a viewport
// screenshot (standard browser view).
type WebScreenshotParamsFullScreenshot string

const (
	WebScreenshotParamsFullScreenshotTrue  WebScreenshotParamsFullScreenshot = "true"
	WebScreenshotParamsFullScreenshotFalse WebScreenshotParamsFullScreenshot = "false"
)

// Optional parameter to control cookie/consent popup handling. If 'true', we
// dismiss cookie banner before capture. If 'false' or not provided, captures the
// page without that step.
type WebScreenshotParamsHandleCookiePopup string

const (
	WebScreenshotParamsHandleCookiePopupTrue  WebScreenshotParamsHandleCookiePopup = "true"
	WebScreenshotParamsHandleCookiePopupFalse WebScreenshotParamsHandleCookiePopup = "false"
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

type WebSearchParams struct {
	// Natural-language search query.
	Query string `json:"query" api:"required"`
	// Expand the query into multiple parallel variants for broader recall.
	QueryFanout param.Opt[bool] `json:"queryFanout,omitzero"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `json:"timeoutMS,omitzero"`
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
	paramObj
}

func (r WebSearchParams) MarshalJSON() (data []byte, err error) {
	type shadow WebSearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
	// Truncate base64-encoded image data in the Markdown output
	ShortenBase64Images param.Opt[bool] `json:"shortenBase64Images,omitzero"`
	// Soft time budget for the crawl in milliseconds. After each scrape, the crawler
	// checks the elapsed time and, if exceeded, returns the pages collected so far
	// instead of continuing. Min: 10000 (10s). Max: 240000 (4 min). Default: 120000 (2
	// min).
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
	// PDF parsing controls. Use start/end to limit text extraction and OCR to an
	// inclusive 1-based page range.
	Pdf WebWebCrawlMdParamsPdf `json:"pdf,omitzero"`
	paramObj
}

func (r WebWebCrawlMdParams) MarshalJSON() (data []byte, err error) {
	type shadow WebWebCrawlMdParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebWebCrawlMdParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PDF parsing controls. Use start/end to limit text extraction and OCR to an
// inclusive 1-based page range.
type WebWebCrawlMdParamsPdf struct {
	// Last 1-based PDF page to parse. When omitted, parsing ends at the last page.
	// Must be greater than or equal to start when both are provided.
	End param.Opt[int64] `json:"end,omitzero"`
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

type WebWebScrapeHTMLParams struct {
	// Full URL to scrape (must include http:// or https:// protocol)
	URL string `query:"url" api:"required" format:"uri" json:"-"`
	// When true, iframes are rendered inline into the returned HTML.
	IncludeFrames param.Opt[bool] `query:"includeFrames,omitzero" json:"-"`
	// Return a cached result if a prior scrape for the same parameters exists and is
	// younger than this many milliseconds. Defaults to 1 day (86400000 ms) when
	// omitted. Max is 30 days (2592000000 ms). Set to 0 to always scrape fresh.
	MaxAgeMs param.Opt[int64] `query:"maxAgeMs,omitzero" json:"-"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `query:"timeoutMS,omitzero" json:"-"`
	// Optional browser wait time in milliseconds after initial page load. Min: 0. Max:
	// 30000 (30 seconds).
	WaitForMs param.Opt[int64] `query:"waitForMs,omitzero" json:"-"`
	// PDF parsing controls. Use start/end to limit text extraction and OCR to an
	// inclusive 1-based page range.
	Pdf WebWebScrapeHTMLParamsPdf `query:"pdf,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebWebScrapeHTMLParams]'s query parameters as `url.Values`.
func (r WebWebScrapeHTMLParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// PDF parsing controls. Use start/end to limit text extraction and OCR to an
// inclusive 1-based page range.
type WebWebScrapeHTMLParamsPdf struct {
	// Last 1-based PDF page to parse. When omitted, parsing ends at the last page.
	// Must be greater than or equal to start when both are provided.
	End param.Opt[int64] `query:"end,omitzero" json:"-"`
	// When true, PDF URLs are fetched and parsed. When false, PDF URLs are skipped and
	// a 400 WEBSITE_ACCESS_ERROR is returned.
	ShouldParse param.Opt[bool] `query:"shouldParse,omitzero" json:"-"`
	// First 1-based PDF page to parse. When omitted, parsing starts at the first page.
	Start param.Opt[int64] `query:"start,omitzero" json:"-"`
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

type WebWebScrapeImagesParams struct {
	// Page URL to inspect. Must include http:// or https://.
	URL string `query:"url" api:"required" format:"uri" json:"-"`
	// Reuse a cached result this many milliseconds old or newer. Default: 86400000 (1
	// day). Set to 0 to bypass cache. Maximum: 2592000000 (30 days).
	MaxAgeMs param.Opt[int64] `query:"maxAgeMs,omitzero" json:"-"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `query:"timeoutMS,omitzero" json:"-"`
	// Optional browser wait time in milliseconds after initial page load before
	// collecting images. Min: 0. Max: 30000 (30 seconds).
	WaitForMs param.Opt[int64] `query:"waitForMs,omitzero" json:"-"`
	// Optional per-image processing, sent as deep-object query params such as
	// enrichment[resolution]=true.
	Enrichment WebWebScrapeImagesParamsEnrichment `query:"enrichment,omitzero" json:"-"`
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

// Optional per-image processing, sent as deep-object query params such as
// enrichment[resolution]=true.
type WebWebScrapeImagesParamsEnrichment struct {
	// Classify each image by visual asset type.
	Classification param.Opt[bool] `query:"classification,omitzero" json:"-"`
	// Host materializable images on the Brand.dev CDN and return their URL and MIME
	// type.
	HostedURL param.Opt[bool] `query:"hostedUrl,omitzero" json:"-"`
	// Per-image enrichment timeout in milliseconds. Default: 30000. Maximum: 60000.
	MaxTimePerMs param.Opt[int64] `query:"maxTimePerMs,omitzero" json:"-"`
	// Measure image width and height when possible.
	Resolution param.Opt[bool] `query:"resolution,omitzero" json:"-"`
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

type WebWebScrapeMdParams struct {
	// Full URL to scrape into LLM usable Markdown (must include http:// or https://
	// protocol)
	URL string `query:"url" api:"required" format:"uri" json:"-"`
	// When true, the contents of iframes are rendered to Markdown.
	IncludeFrames param.Opt[bool] `query:"includeFrames,omitzero" json:"-"`
	// Include image references in Markdown output
	IncludeImages param.Opt[bool] `query:"includeImages,omitzero" json:"-"`
	// Preserve hyperlinks in Markdown output
	IncludeLinks param.Opt[bool] `query:"includeLinks,omitzero" json:"-"`
	// Return a cached result if a prior scrape for the same parameters exists and is
	// younger than this many milliseconds. Defaults to 1 day (86400000 ms) when
	// omitted. Max is 30 days (2592000000 ms). Set to 0 to always scrape fresh.
	MaxAgeMs param.Opt[int64] `query:"maxAgeMs,omitzero" json:"-"`
	// Shorten base64-encoded image data in the Markdown output
	ShortenBase64Images param.Opt[bool] `query:"shortenBase64Images,omitzero" json:"-"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `query:"timeoutMS,omitzero" json:"-"`
	// Extract only the main content of the page, excluding headers, footers, sidebars,
	// and navigation
	UseMainContentOnly param.Opt[bool] `query:"useMainContentOnly,omitzero" json:"-"`
	// Optional browser wait time in milliseconds after initial page load before
	// converting the page to Markdown. Min: 0. Max: 30000 (30 seconds).
	WaitForMs param.Opt[int64] `query:"waitForMs,omitzero" json:"-"`
	// PDF parsing controls. Use start/end to limit text extraction and OCR to an
	// inclusive 1-based page range.
	Pdf WebWebScrapeMdParamsPdf `query:"pdf,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebWebScrapeMdParams]'s query parameters as `url.Values`.
func (r WebWebScrapeMdParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// PDF parsing controls. Use start/end to limit text extraction and OCR to an
// inclusive 1-based page range.
type WebWebScrapeMdParamsPdf struct {
	// Last 1-based PDF page to parse. When omitted, parsing ends at the last page.
	// Must be greater than or equal to start when both are provided.
	End param.Opt[int64] `query:"end,omitzero" json:"-"`
	// When true, PDF URLs are fetched and parsed. When false, PDF URLs are skipped and
	// a 400 WEBSITE_ACCESS_ERROR is returned.
	ShouldParse param.Opt[bool] `query:"shouldParse,omitzero" json:"-"`
	// First 1-based PDF page to parse. When omitted, parsing starts at the first page.
	Start param.Opt[int64] `query:"start,omitzero" json:"-"`
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

type WebWebScrapeSitemapParams struct {
	// Domain to build a sitemap for
	Domain string `query:"domain" api:"required" json:"-"`
	// Maximum number of links to return from the sitemap crawl. Defaults to 10,000.
	// Minimum is 1, maximum is 100,000.
	MaxLinks param.Opt[int64] `query:"maxLinks,omitzero" json:"-"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `query:"timeoutMS,omitzero" json:"-"`
	// Optional RE2-compatible regex pattern. Only URLs matching this pattern are
	// returned and counted against maxLinks.
	URLRegex param.Opt[string] `query:"urlRegex,omitzero" json:"-"`
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
