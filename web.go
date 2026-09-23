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

// Researches the live web and returns a sourced answer in your requested JSON
// shape. Select fast for a smaller research budget at 10 credits or ultra for
// deeper reasoning at 100 credits. Defaults to ultra. Fast research is limited to
// 30 seconds and ultra to 50 seconds; timeoutOpts.milliseconds can shorten either
// deadline.
func (r *WebService) Answers(ctx context.Context, body WebAnswersParams, opts ...option.RequestOption) (res *WebAnswersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "web/answers"
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

// Extract a comprehensive design system from a website including colors,
// typography, spacing, shadows, and UI components.
func (r *WebService) ExtractStyleguide(ctx context.Context, query WebExtractStyleguideParams, opts ...option.RequestOption) (res *WebExtractStyleguideResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "web/styleguide"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Discovers URLs using the same sitemap crawl, filters, and limits as
// /web/scrape/sitemap. Each URL includes its available title, description,
// keywords, and language. URLs without stored enrichment are returned immediately
// with only the URL and queued for background HTML scraping, so later requests can
// include their metadata. Responses are never cached as a whole; every request
// reads the current per-URL enrichment. Zero data retention and credential-bearing
// discovery requests return URLs without reading or storing shared enrichment or
// queuing background scrapes. Costs 1 credit, or 2 credits with search.
func (r *WebService) MapURLs(ctx context.Context, query WebMapURLsParams, opts ...option.RequestOption) (res *WebMapURLsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "web/urls"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Reuse cached outputs independently and capture missing formats in one page
// visit. Each cache key includes only the settings that affect that output. HTML
// is shared with Markdown, parsed fields, and JSON extraction. Cached outputs can
// come from different visits within maxAgeMs; use 0 for a fresh capture. HTML-only
// requests use the existing fast acquisition path. One credit per request,
// including cache hits and missing pages, or two with browser actions; JSON
// extraction adds four credits and runs an LLM over the page Markdown on every
// request that has text to extract; PDF OCR adds one credit per recovered page on
// fresh extraction. Original response bytes and screenshots are limited to 20 MiB
// each, screenshots to 40 megapixels, and the combined browser capture to 60 MiB.
func (r *WebService) Scrape(ctx context.Context, body WebScrapeParams, opts ...option.RequestOption) (res *WebScrapeResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "web/scrape"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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

type WebAnswersResponse struct {
	// The answer, in the shape requested by json_format.
	JsonContent map[string]any `json:"json_content" api:"required"`
	// URLs that supplied search results or readable page content, in first-seen order.
	// Unreadable pages are excluded.
	Sources []string `json:"sources" api:"required"`
	// Credit usage, included whenever a valid API key is provided.
	KeyMetadata WebAnswersResponseKeyMetadata `json:"key_metadata"`
	// True when the request deadline ended research and the answer uses the evidence
	// collected so far.
	Partial bool `json:"partial"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JsonContent respjson.Field
		Sources     respjson.Field
		KeyMetadata respjson.Field
		Partial     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebAnswersResponse) RawJSON() string { return r.JSON.raw }
func (r *WebAnswersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credit usage, included whenever a valid API key is provided.
type WebAnswersResponseKeyMetadata struct {
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
func (r WebAnswersResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebAnswersResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebExtractCompetitorsResponse struct {
	// Direct competitors ordered by relevance and confidence.
	Competitors []WebExtractCompetitorsResponseCompetitor `json:"competitors" api:"required"`
	// Normalized input domain.
	Domain string `json:"domain" api:"required"`
	// Unique id of this API call, also sent in the X-Request-Id response header. Quote
	// it when contacting support about a failed request.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Status of the response.
	//
	// Any of "ok".
	Status WebExtractCompetitorsResponseStatus `json:"status" api:"required"`
	// Target company profile inferred from the landing page.
	Target WebExtractCompetitorsResponseTarget `json:"target" api:"required"`
	// Credit usage, included whenever a valid API key is provided.
	KeyMetadata WebExtractCompetitorsResponseKeyMetadata `json:"key_metadata"`
	// True when the timeout ended processing and this response contains only usable
	// results completed so far. Unfinished results are omitted.
	Partial bool `json:"partial"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Competitors respjson.Field
		Domain      respjson.Field
		RequestID   respjson.Field
		Status      respjson.Field
		Target      respjson.Field
		KeyMetadata respjson.Field
		Partial     respjson.Field
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

// Credit usage, included whenever a valid API key is provided.
type WebExtractCompetitorsResponseKeyMetadata struct {
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
func (r WebExtractCompetitorsResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebExtractCompetitorsResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebExtractStyleguideResponse struct {
	// Cache outcome for this response. Composite responses are hits only when every
	// cache-controlled fetch contributing to the output was a hit; age_ms is the
	// oldest contributing hit.
	CacheMetadata WebExtractStyleguideResponseCacheMetadata `json:"cache_metadata" api:"required"`
	// Unique id of this API call, also sent in the X-Request-Id response header. Quote
	// it when contacting support about a failed request.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// HTTP status code
	Code int64 `json:"code"`
	// The normalized domain that was processed
	Domain string `json:"domain"`
	// How complete the returned content is. `loaded` means the page finished the waits
	// the request asked for. `still-loading` only occurs with
	// timeoutOpts.behavior=return-partial: the timeoutOpts.milliseconds deadline was
	// reached first, so the content reflects the DOM at that moment and late-rendering
	// parts may be missing. Partial results are billed at the base request cost.
	//
	// Any of "loaded", "still-loading".
	FinalDomState WebExtractStyleguideResponseFinalDomState `json:"finalDOMState"`
	// Credit usage, included whenever a valid API key is provided.
	KeyMetadata WebExtractStyleguideResponseKeyMetadata `json:"key_metadata"`
	// Status of the response, e.g., 'ok'
	Status string `json:"status"`
	// Comprehensive styleguide data extracted from the website
	Styleguide WebExtractStyleguideResponseStyleguide `json:"styleguide"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CacheMetadata respjson.Field
		RequestID     respjson.Field
		Code          respjson.Field
		Domain        respjson.Field
		FinalDomState respjson.Field
		KeyMetadata   respjson.Field
		Status        respjson.Field
		Styleguide    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebExtractStyleguideResponse) RawJSON() string { return r.JSON.raw }
func (r *WebExtractStyleguideResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache outcome for this response. Composite responses are hits only when every
// cache-controlled fetch contributing to the output was a hit; age_ms is the
// oldest contributing hit.
type WebExtractStyleguideResponseCacheMetadata struct {
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
func (r WebExtractStyleguideResponseCacheMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebExtractStyleguideResponseCacheMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How complete the returned content is. `loaded` means the page finished the waits
// the request asked for. `still-loading` only occurs with
// timeoutOpts.behavior=return-partial: the timeoutOpts.milliseconds deadline was
// reached first, so the content reflects the DOM at that moment and late-rendering
// parts may be missing. Partial results are billed at the base request cost.
type WebExtractStyleguideResponseFinalDomState string

const (
	WebExtractStyleguideResponseFinalDomStateLoaded       WebExtractStyleguideResponseFinalDomState = "loaded"
	WebExtractStyleguideResponseFinalDomStateStillLoading WebExtractStyleguideResponseFinalDomState = "still-loading"
)

// Credit usage, included whenever a valid API key is provided.
type WebExtractStyleguideResponseKeyMetadata struct {
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

type WebMapURLsResponse struct {
	Domain string `json:"domain" api:"required"`
	// Unique id of this API call, also sent in the X-Request-Id response header. Quote
	// it when contacting support about a failed request.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Any of true.
	Success bool                    `json:"success" api:"required"`
	URLs    []WebMapURLsResponseURL `json:"urls" api:"required"`
	// Credit usage, included whenever a valid API key is provided.
	KeyMetadata WebMapURLsResponseKeyMetadata `json:"key_metadata"`
	Partial     bool                          `json:"partial"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		RequestID   respjson.Field
		Success     respjson.Field
		URLs        respjson.Field
		KeyMetadata respjson.Field
		Partial     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebMapURLsResponse) RawJSON() string { return r.JSON.raw }
func (r *WebMapURLsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebMapURLsResponseURL struct {
	URL         string   `json:"url" api:"required"`
	Description string   `json:"description"`
	Keywords    []string `json:"keywords"`
	Language    string   `json:"language"`
	Title       string   `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Description respjson.Field
		Keywords    respjson.Field
		Language    respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebMapURLsResponseURL) RawJSON() string { return r.JSON.raw }
func (r *WebMapURLsResponseURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credit usage, included whenever a valid API key is provided.
type WebMapURLsResponseKeyMetadata struct {
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
func (r WebMapURLsResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebMapURLsResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScrapeResponse struct {
	// Original HTTP response body. Waiting, actions, and content filters never change
	// it.
	Bytes WebScrapeResponseBytes `json:"bytes" api:"required"`
	// Cache outcome for this response. Composite responses are hits only when every
	// cache-controlled fetch contributing to the output was a hit; age_ms is the
	// oldest contributing hit.
	CacheMetadata WebScrapeResponseCacheMetadata `json:"cache_metadata" api:"required"`
	// Rendered HTML after content filters.
	HTML WebScrapeResponseHTML `json:"html" api:"required"`
	// Images after content filters. Empty when none are found.
	Images WebScrapeResponseImages `json:"images" api:"required"`
	// Page data extracted into jsonParams.schema, after shared content filters. Values
	// are grounded in the page; optional fields the page does not state are omitted,
	// or null when their type allows null. An empty object when the filters leave no
	// text.
	Json WebScrapeResponseJson `json:"json" api:"required"`
	// Markdown after content filters.
	Markdown WebScrapeResponseMarkdown `json:"markdown" api:"required"`
	// Page details, when available.
	Metadata WebScrapeResponseMetadata `json:"metadata" api:"required"`
	// Fields produced by parseParams.rules, after shared content filters.
	Parsed WebScrapeResponseParsed `json:"parsed" api:"required"`
	// Unique id of this API call, also sent in the X-Request-Id response header. Quote
	// it when contacting support about a failed request.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// An image data URL. Use directly as an image src.
	Screenshot WebScrapeResponseScreenshot `json:"screenshot" api:"required"`
	// Final URL after redirects and browser actions.
	URL string `json:"url" api:"required" format:"uri"`
	// Present when return-partial captures a page that is still loading or returns
	// images before image processing finishes. Partial responses are not cached.
	//
	// Any of true.
	IsPartial bool `json:"isPartial"`
	// Credit usage, included whenever a valid API key is provided.
	KeyMetadata WebScrapeResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bytes         respjson.Field
		CacheMetadata respjson.Field
		HTML          respjson.Field
		Images        respjson.Field
		Json          respjson.Field
		Markdown      respjson.Field
		Metadata      respjson.Field
		Parsed        respjson.Field
		RequestID     respjson.Field
		Screenshot    respjson.Field
		URL           respjson.Field
		IsPartial     respjson.Field
		KeyMetadata   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScrapeResponse) RawJSON() string { return r.JSON.raw }
func (r *WebScrapeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Original HTTP response body. Waiting, actions, and content filters never change
// it.
type WebScrapeResponseBytes struct {
	Data      WebScrapeResponseBytesData `json:"data" api:"required"`
	Requested bool                       `json:"requested" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScrapeResponseBytes) RawJSON() string { return r.JSON.raw }
func (r *WebScrapeResponseBytes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScrapeResponseBytesData struct {
	// Original response body as base64, after HTTP decompression. Maximum decoded
	// size: 20 MiB.
	Base64      string `json:"base64" api:"required"`
	ContentType string `json:"contentType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Base64      respjson.Field
		ContentType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScrapeResponseBytesData) RawJSON() string { return r.JSON.raw }
func (r *WebScrapeResponseBytesData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache outcome for this response. Composite responses are hits only when every
// cache-controlled fetch contributing to the output was a hit; age_ms is the
// oldest contributing hit.
type WebScrapeResponseCacheMetadata struct {
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
func (r WebScrapeResponseCacheMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebScrapeResponseCacheMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rendered HTML after content filters.
type WebScrapeResponseHTML struct {
	Data      string `json:"data" api:"required"`
	Requested bool   `json:"requested" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScrapeResponseHTML) RawJSON() string { return r.JSON.raw }
func (r *WebScrapeResponseHTML) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Images after content filters. Empty when none are found.
type WebScrapeResponseImages struct {
	Data      []WebScrapeResponseImagesData `json:"data" api:"required"`
	Requested bool                          `json:"requested" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScrapeResponseImages) RawJSON() string { return r.JSON.raw }
func (r *WebScrapeResponseImages) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScrapeResponseImagesData struct {
	// Alt text, if present.
	Alt string `json:"alt" api:"required"`
	// Image URL, or a data URI for inline images.
	URL string `json:"url" api:"required" format:"uri"`
	// Any of "photography", "illustration", "logo", "wordmark", "icon", "pattern",
	// "graphic", "other".
	Classification string `json:"classification"`
	// Hosted copy when file enrichment is requested and zdr is disabled. Valid for 24
	// hours from the original capture.
	FileURL string `json:"fileUrl" format:"uri"`
	Height  int64  `json:"height"`
	Width   int64  `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Alt            respjson.Field
		URL            respjson.Field
		Classification respjson.Field
		FileURL        respjson.Field
		Height         respjson.Field
		Width          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScrapeResponseImagesData) RawJSON() string { return r.JSON.raw }
func (r *WebScrapeResponseImagesData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Page data extracted into jsonParams.schema, after shared content filters. Values
// are grounded in the page; optional fields the page does not state are omitted,
// or null when their type allows null. An empty object when the filters leave no
// text.
type WebScrapeResponseJson struct {
	Data      map[string]any `json:"data" api:"required"`
	Requested bool           `json:"requested" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScrapeResponseJson) RawJSON() string { return r.JSON.raw }
func (r *WebScrapeResponseJson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Markdown after content filters.
type WebScrapeResponseMarkdown struct {
	Data      string `json:"data" api:"required"`
	Requested bool   `json:"requested" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScrapeResponseMarkdown) RawJSON() string { return r.JSON.raw }
func (r *WebScrapeResponseMarkdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Page details, when available.
type WebScrapeResponseMetadata struct {
	// Additional non-social meta tags not promoted to top-level metadata fields.
	AdditionalMeta map[string]WebScrapeResponseMetadataAdditionalMetaUnion `json:"additionalMeta"`
	// Resolved alternate links from link rel=alternate tags.
	Alternates []WebScrapeResponseMetadataAlternate `json:"alternates"`
	// Author metadata, when present.
	Author string `json:"author"`
	// Resolved canonical URL, when present.
	CanonicalURL string `json:"canonicalUrl"`
	// Best description extracted from standard, Open Graph, or Twitter metadata.
	Description string `json:"description"`
	// Resolved favicon URL, when present.
	Favicon string `json:"favicon"`
	// Page headings (h1–h6) in document order, extracted from the unfiltered document.
	// Capped at the first 500 headings. Omitted when the page has none.
	Headings []WebScrapeResponseMetadataHeading `json:"headings"`
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
	OpenGraph map[string]WebScrapeResponseMetadataOpenGraphUnion `json:"openGraph"`
	// Published timestamp/date from page metadata, when present.
	PublishedTime string `json:"publishedTime"`
	// Robots meta directive, when present.
	Robots string `json:"robots"`
	// Site or application name from page metadata.
	SiteName string `json:"siteName"`
	// Best title extracted from the page.
	Title string `json:"title"`
	// Twitter card metadata with the twitter: prefix removed and keys camel-cased.
	Twitter map[string]WebScrapeResponseMetadataTwitterUnion `json:"twitter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdditionalMeta respjson.Field
		Alternates     respjson.Field
		Author         respjson.Field
		CanonicalURL   respjson.Field
		Description    respjson.Field
		Favicon        respjson.Field
		Headings       respjson.Field
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
func (r WebScrapeResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebScrapeResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WebScrapeResponseMetadataAdditionalMetaUnion contains all possible properties
// and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type WebScrapeResponseMetadataAdditionalMetaUnion struct {
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

func (u WebScrapeResponseMetadataAdditionalMetaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebScrapeResponseMetadataAdditionalMetaUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WebScrapeResponseMetadataAdditionalMetaUnion) RawJSON() string { return u.JSON.raw }

func (r *WebScrapeResponseMetadataAdditionalMetaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScrapeResponseMetadataAlternate struct {
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
func (r WebScrapeResponseMetadataAlternate) RawJSON() string { return r.JSON.raw }
func (r *WebScrapeResponseMetadataAlternate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScrapeResponseMetadataHeading struct {
	// Heading level, 1–6 (from h1–h6).
	Level int64 `json:"level" api:"required"`
	// Heading text with whitespace collapsed, truncated to 1000 characters.
	Text string `json:"text" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Level       respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScrapeResponseMetadataHeading) RawJSON() string { return r.JSON.raw }
func (r *WebScrapeResponseMetadataHeading) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WebScrapeResponseMetadataOpenGraphUnion contains all possible properties and
// values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type WebScrapeResponseMetadataOpenGraphUnion struct {
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

func (u WebScrapeResponseMetadataOpenGraphUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebScrapeResponseMetadataOpenGraphUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WebScrapeResponseMetadataOpenGraphUnion) RawJSON() string { return u.JSON.raw }

func (r *WebScrapeResponseMetadataOpenGraphUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WebScrapeResponseMetadataTwitterUnion contains all possible properties and
// values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type WebScrapeResponseMetadataTwitterUnion struct {
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

func (u WebScrapeResponseMetadataTwitterUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebScrapeResponseMetadataTwitterUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WebScrapeResponseMetadataTwitterUnion) RawJSON() string { return u.JSON.raw }

func (r *WebScrapeResponseMetadataTwitterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Fields produced by parseParams.rules, after shared content filters.
type WebScrapeResponseParsed struct {
	Data      map[string]any `json:"data" api:"required"`
	Requested bool           `json:"requested" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScrapeResponseParsed) RawJSON() string { return r.JSON.raw }
func (r *WebScrapeResponseParsed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An image data URL. Use directly as an image src.
type WebScrapeResponseScreenshot struct {
	Data      string `json:"data" api:"required" format:"uri"`
	Requested bool   `json:"requested" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScrapeResponseScreenshot) RawJSON() string { return r.JSON.raw }
func (r *WebScrapeResponseScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credit usage, included whenever a valid API key is provided.
type WebScrapeResponseKeyMetadata struct {
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
func (r WebScrapeResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebScrapeResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScreenshotResponse struct {
	// Cache outcome for this response. Composite responses are hits only when every
	// cache-controlled fetch contributing to the output was a hit; age_ms is the
	// oldest contributing hit.
	CacheMetadata WebScreenshotResponseCacheMetadata `json:"cache_metadata" api:"required"`
	// Unique id of this API call, also sent in the X-Request-Id response header. Quote
	// it when contacting support about a failed request.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// HTTP status code
	Code int64 `json:"code"`
	// The normalized domain that was processed
	Domain string `json:"domain"`
	// How complete the returned content is. `loaded` means the page finished the waits
	// the request asked for. `still-loading` only occurs with
	// timeoutOpts.behavior=return-partial: the timeoutOpts.milliseconds deadline was
	// reached first, so the content reflects the DOM at that moment and late-rendering
	// parts may be missing. Partial results are billed at the base request cost.
	//
	// Any of "loaded", "still-loading".
	FinalDomState WebScreenshotResponseFinalDomState `json:"finalDOMState"`
	// Height in pixels of the returned screenshot image
	Height int64 `json:"height"`
	// Credit usage, included whenever a valid API key is provided.
	KeyMetadata WebScreenshotResponseKeyMetadata `json:"key_metadata"`
	// Public image URL for standard requests, or an in-memory data URL when ZDR or
	// non-empty custom headers are supplied.
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
		CacheMetadata  respjson.Field
		RequestID      respjson.Field
		Code           respjson.Field
		Domain         respjson.Field
		FinalDomState  respjson.Field
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

// Cache outcome for this response. Composite responses are hits only when every
// cache-controlled fetch contributing to the output was a hit; age_ms is the
// oldest contributing hit.
type WebScreenshotResponseCacheMetadata struct {
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
func (r WebScreenshotResponseCacheMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebScreenshotResponseCacheMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How complete the returned content is. `loaded` means the page finished the waits
// the request asked for. `still-loading` only occurs with
// timeoutOpts.behavior=return-partial: the timeoutOpts.milliseconds deadline was
// reached first, so the content reflects the DOM at that moment and late-rendering
// parts may be missing. Partial results are billed at the base request cost.
type WebScreenshotResponseFinalDomState string

const (
	WebScreenshotResponseFinalDomStateLoaded       WebScreenshotResponseFinalDomState = "loaded"
	WebScreenshotResponseFinalDomStateStillLoading WebScreenshotResponseFinalDomState = "still-loading"
)

// Credit usage, included whenever a valid API key is provided.
type WebScreenshotResponseKeyMetadata struct {
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
	// Cache outcome for this response. Composite responses are hits only when every
	// cache-controlled fetch contributing to the output was a hit; age_ms is the
	// oldest contributing hit.
	CacheMetadata WebSearchResponseCacheMetadata `json:"cache_metadata" api:"required"`
	// Echo of the original query (useful when fanout was enabled).
	Query string `json:"query" api:"required"`
	// Unique id of this API call, also sent in the X-Request-Id response header. Quote
	// it when contacting support about a failed request.
	RequestID string                    `json:"request_id" api:"required" format:"uuid"`
	Results   []WebSearchResponseResult `json:"results" api:"required"`
	// Credit usage, included whenever a valid API key is provided.
	KeyMetadata WebSearchResponseKeyMetadata `json:"key_metadata"`
	// True when timeoutOpts.behavior=return-partial returned the usable results
	// collected before the deadline. Partial collections are not cached as complete
	// results.
	Partial bool `json:"partial"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CacheMetadata respjson.Field
		Query         respjson.Field
		RequestID     respjson.Field
		Results       respjson.Field
		KeyMetadata   respjson.Field
		Partial       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *WebSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache outcome for this response. Composite responses are hits only when every
// cache-controlled fetch contributing to the output was a hit; age_ms is the
// oldest contributing hit.
type WebSearchResponseCacheMetadata struct {
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
func (r WebSearchResponseCacheMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebSearchResponseCacheMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebSearchResponseResult struct {
	// Snippet excerpt from the page. Empty string when the search provider does not
	// supply a snippet.
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
	// Any of "SUCCESS", "NOT_REQUESTED", "TIMEOUT", "CONTENT_TOO_LARGE",
	// "WEBSITE_ACCESS_ERROR", "ERROR".
	Code string `json:"code" api:"required"`
	// GFM Markdown of the page. Null unless markdownOptions.enabled is true and
	// scraping succeeded.
	Markdown string `json:"markdown" api:"required"`
	// How complete the returned content is. `loaded` means the page finished the waits
	// the request asked for. `still-loading` only occurs with
	// timeoutOpts.behavior=return-partial: the timeoutOpts.milliseconds deadline was
	// reached first, so the content reflects the DOM at that moment and late-rendering
	// parts may be missing. Partial results are billed at the base request cost.
	//
	// Any of "loaded", "still-loading".
	FinalDomState string `json:"finalDOMState"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code          respjson.Field
		Markdown      respjson.Field
		FinalDomState respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebSearchResponseResultMarkdown) RawJSON() string { return r.JSON.raw }
func (r *WebSearchResponseResultMarkdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credit usage, included whenever a valid API key is provided.
type WebSearchResponseKeyMetadata struct {
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
func (r WebSearchResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebSearchResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebWebCrawlMdResponse struct {
	// Cache outcome for this response. Composite responses are hits only when every
	// cache-controlled fetch contributing to the output was a hit; age_ms is the
	// oldest contributing hit.
	CacheMetadata WebWebCrawlMdResponseCacheMetadata `json:"cache_metadata" api:"required"`
	Metadata      WebWebCrawlMdResponseMetadata      `json:"metadata" api:"required"`
	// Unique id of this API call, also sent in the X-Request-Id response header. Quote
	// it when contacting support about a failed request.
	RequestID string                        `json:"request_id" api:"required" format:"uuid"`
	Results   []WebWebCrawlMdResponseResult `json:"results" api:"required"`
	// Credit usage, included whenever a valid API key is provided.
	KeyMetadata WebWebCrawlMdResponseKeyMetadata `json:"key_metadata"`
	// True when timeoutOpts.behavior=return-partial returned the usable results
	// collected before the deadline. Partial collections are not cached as complete
	// results.
	Partial bool `json:"partial"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CacheMetadata respjson.Field
		Metadata      respjson.Field
		RequestID     respjson.Field
		Results       respjson.Field
		KeyMetadata   respjson.Field
		Partial       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebWebCrawlMdResponse) RawJSON() string { return r.JSON.raw }
func (r *WebWebCrawlMdResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache outcome for this response. Composite responses are hits only when every
// cache-controlled fetch contributing to the output was a hit; age_ms is the
// oldest contributing hit.
type WebWebCrawlMdResponseCacheMetadata struct {
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
func (r WebWebCrawlMdResponseCacheMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebWebCrawlMdResponseCacheMetadata) UnmarshalJSON(data []byte) error {
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
	// Page headings (h1–h6) in document order, extracted from the unfiltered document.
	// Capped at the first 500 headings. Omitted when the page has none.
	Headings []WebWebCrawlMdResponseResultMetadataHeading `json:"headings"`
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
		Headings       respjson.Field
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

type WebWebCrawlMdResponseResultMetadataHeading struct {
	// Heading level, 1–6 (from h1–h6).
	Level int64 `json:"level" api:"required"`
	// Heading text with whitespace collapsed, truncated to 1000 characters.
	Text string `json:"text" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Level       respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebWebCrawlMdResponseResultMetadataHeading) RawJSON() string { return r.JSON.raw }
func (r *WebWebCrawlMdResponseResultMetadataHeading) UnmarshalJSON(data []byte) error {
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

// Credit usage, included whenever a valid API key is provided.
type WebWebCrawlMdResponseKeyMetadata struct {
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
func (r WebWebCrawlMdResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebWebCrawlMdResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnswersParams struct {
	// What to research and answer, in plain language. Naming a domain in the task (for
	// example "pricing on context.dev") makes the agent read that site before it
	// searches.
	Task string `json:"task" api:"required"`
	// An example object with placeholder values (for example {"pricing_page_url": "",
	// "plans": [{"name": "", "price": 0}]}). Object keys and value types are
	// preserved; unknown values may be null. Empty arrays accept any JSON items.
	// Defaults to {"result": ""}. Maximum 8 levels, 500 values, and 16000 characters.
	JsonFormat map[string]any `json:"json_format,omitzero"`
	// Research level: fast uses a smaller model and research budget for 10 credits;
	// ultra uses deeper reasoning and research for 100 credits. Defaults to ultra.
	// Only successful requests consume credits.
	//
	// Any of "fast", "ultra".
	Mode WebAnswersParamsMode `json:"mode,omitzero"`
	// Optional tags for tracking usage. Up to 20 tags, each 1 to 50 characters.
	Tags []string `json:"tags,omitzero"`
	// Optional request deadline and behavior on timeout. For GET requests, use
	// timeoutOpts[milliseconds]=30000&timeoutOpts[behavior]=fail or a JSON-encoded
	// timeoutOpts object.
	TimeoutOpts WebAnswersParamsTimeoutOpts `json:"timeoutOpts,omitzero"`
	// Set to enabled to bypass shared caches and omit request and response content
	// from retained usage logs. Asset uploads are skipped, so hosted image URLs are
	// omitted. Requires zero data retention to be enabled for your organization
	// (contact support@context.dev), otherwise the request fails with ZDR_NOT_ENABLED.
	// Successful ZDR responses include X-Context-ZDR: true.
	//
	// Any of "enabled", "disabled".
	Zdr WebAnswersParamsZdr `json:"zdr,omitzero"`
	paramObj
}

func (r WebAnswersParams) MarshalJSON() (data []byte, err error) {
	type shadow WebAnswersParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebAnswersParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Research level: fast uses a smaller model and research budget for 10 credits;
// ultra uses deeper reasoning and research for 100 credits. Defaults to ultra.
// Only successful requests consume credits.
type WebAnswersParamsMode string

const (
	WebAnswersParamsModeFast  WebAnswersParamsMode = "fast"
	WebAnswersParamsModeUltra WebAnswersParamsMode = "ultra"
)

// Optional request deadline and behavior on timeout. For GET requests, use
// timeoutOpts[milliseconds]=30000&timeoutOpts[behavior]=fail or a JSON-encoded
// timeoutOpts object.
//
// The property Milliseconds is required.
type WebAnswersParamsTimeoutOpts struct {
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

func (r WebAnswersParamsTimeoutOpts) MarshalJSON() (data []byte, err error) {
	type shadow WebAnswersParamsTimeoutOpts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebAnswersParamsTimeoutOpts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebAnswersParamsTimeoutOpts](
		"behavior", "fail", "return-partial",
	)
}

// Set to enabled to bypass shared caches and omit request and response content
// from retained usage logs. Asset uploads are skipped, so hosted image URLs are
// omitted. Requires zero data retention to be enabled for your organization
// (contact support@context.dev), otherwise the request fails with ZDR_NOT_ENABLED.
// Successful ZDR responses include X-Context-ZDR: true.
type WebAnswersParamsZdr string

const (
	WebAnswersParamsZdrEnabled  WebAnswersParamsZdr = "enabled"
	WebAnswersParamsZdrDisabled WebAnswersParamsZdr = "disabled"
)

type WebExtractCompetitorsParams struct {
	// Company domain to analyze, such as `stripe.com`. Full http(s) URLs are accepted
	// and normalized to their domain.
	Domain string `query:"domain" api:"required" json:"-"`
	// Exact number of direct competitors to return. Defaults to 5.
	NumCompetitors param.Opt[int64] `query:"numCompetitors,omitzero" json:"-"`
	// Comma-separated tags for tracking request usage. Up to 20 tags, each 1-50
	// characters.
	Tags []string `query:"tags,omitzero" json:"-"`
	// Optional request deadline and behavior on timeout. For GET requests, use
	// timeoutOpts[milliseconds]=30000&timeoutOpts[behavior]=fail or a JSON-encoded
	// timeoutOpts object.
	TimeoutOpts WebExtractCompetitorsParamsTimeoutOpts `query:"timeoutOpts,omitzero" json:"-"`
	// Set to enabled to bypass shared caches and omit request and response content
	// from retained usage logs. Asset uploads are skipped, so hosted image URLs are
	// omitted. Requires zero data retention to be enabled for your organization
	// (contact support@context.dev), otherwise the request fails with ZDR_NOT_ENABLED.
	// Successful ZDR responses include X-Context-ZDR: true.
	//
	// Any of "enabled", "disabled".
	Zdr WebExtractCompetitorsParamsZdr `query:"zdr,omitzero" json:"-"`
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

// Optional request deadline and behavior on timeout. For GET requests, use
// timeoutOpts[milliseconds]=30000&timeoutOpts[behavior]=fail or a JSON-encoded
// timeoutOpts object.
//
// The property Milliseconds is required.
type WebExtractCompetitorsParamsTimeoutOpts struct {
	// Request deadline in milliseconds. Maximum: 300000 (5 minutes).
	Milliseconds int64 `query:"milliseconds" api:"required" json:"-"`
	// What to do at the deadline. "fail" returns 408 REQUEST_TIMEOUT without charging
	// credits. "return-partial" returns usable results collected so far; if none are
	// available, the request still fails without charging credits. Partial results are
	// not cached as complete results.
	//
	// Any of "fail", "return-partial".
	Behavior string `query:"behavior,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebExtractCompetitorsParamsTimeoutOpts]'s query parameters
// as `url.Values`.
func (r WebExtractCompetitorsParamsTimeoutOpts) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Set to enabled to bypass shared caches and omit request and response content
// from retained usage logs. Asset uploads are skipped, so hosted image URLs are
// omitted. Requires zero data retention to be enabled for your organization
// (contact support@context.dev), otherwise the request fails with ZDR_NOT_ENABLED.
// Successful ZDR responses include X-Context-ZDR: true.
type WebExtractCompetitorsParamsZdr string

const (
	WebExtractCompetitorsParamsZdrEnabled  WebExtractCompetitorsParamsZdr = "enabled"
	WebExtractCompetitorsParamsZdrDisabled WebExtractCompetitorsParamsZdr = "disabled"
)

type WebExtractStyleguideParams struct {
	// Maximum age in milliseconds for cached brand data before the API performs a hard
	// refresh. Defaults to 3 months (7776000000 ms). Set to 0 to always perform a hard
	// refresh. Negative values are clamped to 0; values above 1 year (31536000000 ms)
	// are clamped to 1 year.
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
	// Optional browser color scheme to emulate for websites that respond to
	// prefers-color-scheme. This value is part of the styleguide cache key.
	//
	// Any of "light", "dark".
	ColorScheme WebExtractStyleguideParamsColorScheme `query:"colorScheme,omitzero" json:"-"`
	// Comma-separated tags for tracking request usage. Up to 20 tags, each 1-50
	// characters.
	Tags []string `query:"tags,omitzero" json:"-"`
	// Optional request deadline and behavior on timeout. For GET requests, use
	// timeoutOpts[milliseconds]=30000&timeoutOpts[behavior]=fail or a JSON-encoded
	// timeoutOpts object.
	TimeoutOpts WebExtractStyleguideParamsTimeoutOpts `query:"timeoutOpts,omitzero" json:"-"`
	// Set to enabled to bypass shared caches and omit request and response content
	// from retained usage logs. Asset uploads are skipped, so hosted image URLs are
	// omitted. Requires zero data retention to be enabled for your organization
	// (contact support@context.dev), otherwise the request fails with ZDR_NOT_ENABLED.
	// Successful ZDR responses include X-Context-ZDR: true.
	//
	// Any of "enabled", "disabled".
	Zdr WebExtractStyleguideParamsZdr `query:"zdr,omitzero" json:"-"`
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

// Optional request deadline and behavior on timeout. For GET requests, use
// timeoutOpts[milliseconds]=30000&timeoutOpts[behavior]=fail or a JSON-encoded
// timeoutOpts object.
//
// The property Milliseconds is required.
type WebExtractStyleguideParamsTimeoutOpts struct {
	// Request deadline in milliseconds. Maximum: 300000 (5 minutes).
	Milliseconds int64 `query:"milliseconds" api:"required" json:"-"`
	// What to do at the deadline. "fail" returns 408 REQUEST_TIMEOUT without charging
	// credits. "return-partial" returns usable results collected so far; if none are
	// available, the request still fails without charging credits. Partial results are
	// not cached as complete results. "return-partial" requires milliseconds of at
	// least 5000.
	//
	// Any of "fail", "return-partial".
	Behavior string `query:"behavior,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebExtractStyleguideParamsTimeoutOpts]'s query parameters
// as `url.Values`.
func (r WebExtractStyleguideParamsTimeoutOpts) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Set to enabled to bypass shared caches and omit request and response content
// from retained usage logs. Asset uploads are skipped, so hosted image URLs are
// omitted. Requires zero data retention to be enabled for your organization
// (contact support@context.dev), otherwise the request fails with ZDR_NOT_ENABLED.
// Successful ZDR responses include X-Context-ZDR: true.
type WebExtractStyleguideParamsZdr string

const (
	WebExtractStyleguideParamsZdrEnabled  WebExtractStyleguideParamsZdr = "enabled"
	WebExtractStyleguideParamsZdrDisabled WebExtractStyleguideParamsZdr = "disabled"
)

type WebMapURLsParams struct {
	// Domain to build a sitemap for
	Domain string `query:"domain" api:"required" json:"-"`
	// When true, discover and include public pages and sitemaps on subdomains of the
	// requested domain. Defaults to false.
	IncludeSubdomains param.Opt[bool] `query:"includeSubdomains,omitzero" json:"-"`
	// Maximum number of links to return from the sitemap crawl. Defaults to 10,000.
	// Minimum is 1, maximum is 100,000.
	MaxLinks param.Opt[int64] `query:"maxLinks,omitzero" json:"-"`
	// Optional search phrase. When provided, the crawled sitemap is filtered to the
	// pages whose URLs are about that phrase, most relevant first, and the request
	// costs 2 credits instead of 1.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// Optional explicit sitemap URL. When provided, exactly this sitemap is crawled
	// instead of discovering the domain's sitemaps.
	SitemapURL param.Opt[string] `query:"sitemapUrl,omitzero" format:"uri" json:"-"`
	// Optional RE2-compatible regex pattern. Only URLs matching this pattern are
	// returned and counted against maxLinks.
	URLRegex param.Opt[string] `query:"urlRegex,omitzero" json:"-"`
	// Optional outbound HTTP headers forwarded only to the target URL, sent as
	// deep-object query params such as headers[X-Custom]=value. When provided, caching
	// is bypassed: the result is neither read from nor written to cache.
	Headers map[string]string `query:"headers,omitzero" json:"-"`
	// Comma-separated tags for tracking request usage. Up to 20 tags, each 1-50
	// characters.
	Tags []string `query:"tags,omitzero" json:"-"`
	// Optional request deadline and behavior on timeout. For GET requests, use
	// timeoutOpts[milliseconds]=30000&timeoutOpts[behavior]=fail or a JSON-encoded
	// timeoutOpts object.
	TimeoutOpts WebMapURLsParamsTimeoutOpts `query:"timeoutOpts,omitzero" json:"-"`
	// Set to enabled to bypass shared caches and omit request and response content
	// from retained usage logs. Asset uploads are skipped, so hosted image URLs are
	// omitted. Requires zero data retention to be enabled for your organization
	// (contact support@context.dev), otherwise the request fails with ZDR_NOT_ENABLED.
	// Successful ZDR responses include X-Context-ZDR: true.
	//
	// Any of "enabled", "disabled".
	Zdr WebMapURLsParamsZdr `query:"zdr,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebMapURLsParams]'s query parameters as `url.Values`.
func (r WebMapURLsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Optional request deadline and behavior on timeout. For GET requests, use
// timeoutOpts[milliseconds]=30000&timeoutOpts[behavior]=fail or a JSON-encoded
// timeoutOpts object.
//
// The property Milliseconds is required.
type WebMapURLsParamsTimeoutOpts struct {
	// Request deadline in milliseconds. Maximum: 300000 (5 minutes).
	Milliseconds int64 `query:"milliseconds" api:"required" json:"-"`
	// What to do at the deadline. "fail" returns 408 REQUEST_TIMEOUT without charging
	// credits. "return-partial" returns usable results collected so far; if none are
	// available, the request still fails without charging credits. Partial results are
	// not cached as complete results.
	//
	// Any of "fail", "return-partial".
	Behavior string `query:"behavior,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebMapURLsParamsTimeoutOpts]'s query parameters as
// `url.Values`.
func (r WebMapURLsParamsTimeoutOpts) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Set to enabled to bypass shared caches and omit request and response content
// from retained usage logs. Asset uploads are skipped, so hosted image URLs are
// omitted. Requires zero data retention to be enabled for your organization
// (contact support@context.dev), otherwise the request fails with ZDR_NOT_ENABLED.
// Successful ZDR responses include X-Context-ZDR: true.
type WebMapURLsParamsZdr string

const (
	WebMapURLsParamsZdrEnabled  WebMapURLsParamsZdr = "enabled"
	WebMapURLsParamsZdrDisabled WebMapURLsParamsZdr = "disabled"
)

type WebScrapeParams struct {
	// Outputs to return. Enable at least one; omitted formats are false.
	Formats WebScrapeParamsFormats `json:"formats,omitzero" api:"required"`
	// The URL to scrape.
	URL string `json:"url" api:"required" format:"uri"`
	// Maximum age of each cached output. Defaults to 1 day; 0 fetches fresh and
	// updates the requested outputs. Compatible outputs are shared with the individual
	// scrape endpoints. Image results with hosted files refresh after 23 hours; other
	// outputs retain their own freshness.
	MaxAgeMs param.Opt[int64] `json:"maxAgeMs,omitzero"`
	// Image options. Requires formats.images: true.
	ImageParams WebScrapeParamsImageParams `json:"imageParams,omitzero"`
	// Required when formats.json is true.
	JsonParams WebScrapeParamsJsonParams `json:"jsonParams,omitzero"`
	// Markdown options. Requires formats.markdown: true.
	MarkdownParams WebScrapeParamsMarkdownParams `json:"markdownParams,omitzero"`
	// Required when formats.parse is true.
	ParseParams WebScrapeParamsParseParams `json:"parseParams,omitzero"`
	// Screenshot options. Requires formats.screenshot: true.
	ScreenshotParams WebScrapeParamsScreenshotParams `json:"screenshotParams,omitzero"`
	// Shared browser and content settings. Content filters leave screenshots and
	// original bytes unchanged.
	SharedParams WebScrapeParamsSharedParams `json:"sharedParams,omitzero"`
	// Labels for tracking request usage. Not retained when zdr is enabled.
	Tags []string `json:"tags,omitzero"`
	// Total deadline, including navigation, actions, waiting, and all outputs.
	// Defaults to 60000 milliseconds with behavior fail. Use return-partial to capture
	// the current page state and return captured images if image processing cannot
	// finish before the deadline; these responses set isPartial and are not cached.
	// Every requested format must still be available. Fixed waits must fit before a
	// response reserve of up to 5000 milliseconds (at most one quarter of the timeout)
	// when using return-partial.
	TimeoutOpts WebScrapeParamsTimeoutOpts `json:"timeoutOpts,omitzero"`
	// Zero data retention. Bypasses caches and uploads; excludes request/response
	// content and tags from logs. Must be enabled for your organization.
	//
	// Any of "enabled", "disabled".
	Zdr WebScrapeParamsZdr `json:"zdr,omitzero"`
	paramObj
}

func (r WebScrapeParams) MarshalJSON() (data []byte, err error) {
	type shadow WebScrapeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScrapeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Outputs to return. Enable at least one; omitted formats are false.
type WebScrapeParamsFormats struct {
	// The original HTTP response body.
	Bytes param.Opt[bool] `json:"bytes,omitzero"`
	// Rendered HTML.
	HTML param.Opt[bool] `json:"html,omitzero"`
	// Images found on the page.
	Images param.Opt[bool] `json:"images,omitzero"`
	// Page data extracted by an LLM from the page Markdown into jsonParams.schema;
	// values carried only in attributes or CSS classes need formats.parse instead.
	// Adds four credits when the page has text to extract; when shared content filters
	// leave no text the result is an empty object and only the base price applies.
	Json param.Opt[bool] `json:"json,omitzero"`
	// Page content as Markdown.
	Markdown param.Opt[bool] `json:"markdown,omitzero"`
	// Fields selected by parseParams.rules.
	Parse param.Opt[bool] `json:"parse,omitzero"`
	// An inline image of the page.
	Screenshot param.Opt[bool] `json:"screenshot,omitzero"`
	paramObj
}

func (r WebScrapeParamsFormats) MarshalJSON() (data []byte, err error) {
	type shadow WebScrapeParamsFormats
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScrapeParamsFormats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image options. Requires formats.images: true.
type WebScrapeParamsImageParams struct {
	// For visual duplicates, keep the largest image.
	//
	// Any of "none", "visual".
	Dedupe string `json:"dedupe,omitzero"`
	// Add dimensions, a visual category, or a hosted file URL. Each image has a
	// maximum processing time of 30000 milliseconds, bounded by the remaining request
	// deadline.
	//
	// Any of "dimensions", "classification", "file".
	Enrich []string `json:"enrich,omitzero"`
	paramObj
}

func (r WebScrapeParamsImageParams) MarshalJSON() (data []byte, err error) {
	type shadow WebScrapeParamsImageParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScrapeParamsImageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebScrapeParamsImageParams](
		"dedupe", "none", "visual",
	)
}

// Required when formats.json is true.
//
// The property Schema is required.
type WebScrapeParamsJsonParams struct {
	// JSON Schema for the returned object. Must describe a top-level object; at most
	// 50 KB serialized. Optional fields the page does not state are omitted, or null
	// when their type allows null, while required non-nullable fields always receive a
	// best-effort value, so prefer nullable or optional fields for data a page may
	// omit. Zod users can pass the output of z.toJSONSchema().
	Schema map[string]any `json:"schema,omitzero" api:"required"`
	// Optional guidance on which facts to prioritize or how to interpret schema
	// fields.
	Instructions param.Opt[string] `json:"instructions,omitzero"`
	paramObj
}

func (r WebScrapeParamsJsonParams) MarshalJSON() (data []byte, err error) {
	type shadow WebScrapeParamsJsonParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScrapeParamsJsonParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Markdown options. Requires formats.markdown: true.
type WebScrapeParamsMarkdownParams struct {
	IncludeImages param.Opt[bool] `json:"includeImages,omitzero"`
	IncludeLinks  param.Opt[bool] `json:"includeLinks,omitzero"`
	// Base64 images use placeholders by default. Requires includeImages: true.
	//
	// Any of "placeholder", "preserve".
	InlineImages string `json:"inlineImages,omitzero"`
	paramObj
}

func (r WebScrapeParamsMarkdownParams) MarshalJSON() (data []byte, err error) {
	type shadow WebScrapeParamsMarkdownParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScrapeParamsMarkdownParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebScrapeParamsMarkdownParams](
		"inlineImages", "placeholder", "preserve",
	)
}

// Required when formats.parse is true.
//
// The property Rules is required.
type WebScrapeParamsParseParams struct {
	// Map field names to CSS selectors or rules. Missing items return null; missing
	// lists return [].
	Rules map[string]WebScrapeParamsParseParamsRuleUnion `json:"rules,omitzero" api:"required"`
	paramObj
}

func (r WebScrapeParamsParseParams) MarshalJSON() (data []byte, err error) {
	type shadow WebScrapeParamsParseParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScrapeParamsParseParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebScrapeParamsParseParamsRuleUnion struct {
	OfString                          param.Opt[string]                     `json:",omitzero,inline"`
	OfWebScrapesParseParamsRuleObject *WebScrapeParamsParseParamsRuleObject `json:",omitzero,inline"`
	paramUnion
}

func (u WebScrapeParamsParseParamsRuleUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfWebScrapesParseParamsRuleObject)
}
func (u *WebScrapeParamsParseParamsRuleUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property Selector is required.
type WebScrapeParamsParseParamsRuleObject struct {
	Selector string `json:"selector" api:"required"`
	Output   string `json:"output,omitzero"`
	// Any of "item", "list".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r WebScrapeParamsParseParamsRuleObject) MarshalJSON() (data []byte, err error) {
	type shadow WebScrapeParamsParseParamsRuleObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScrapeParamsParseParamsRuleObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebScrapeParamsParseParamsRuleObject](
		"type", "item", "list",
	)
}

// Screenshot options. Requires formats.screenshot: true.
type WebScrapeParamsScreenshotParams struct {
	// Viewport, full page, one visible element, or a rectangle. Maximum 40 megapixels.
	Area WebScrapeParamsScreenshotParamsAreaUnion `json:"area,omitzero"`
	// Any of "png", "jpeg", "webp".
	Format string `json:"format,omitzero"`
	paramObj
}

func (r WebScrapeParamsScreenshotParams) MarshalJSON() (data []byte, err error) {
	type shadow WebScrapeParamsScreenshotParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScrapeParamsScreenshotParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebScrapeParamsScreenshotParams](
		"format", "png", "jpeg", "webp",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebScrapeParamsScreenshotParamsAreaUnion struct {
	// Check if union is this variant with !param.IsOmitted(union.OfPage)
	OfPage      param.Opt[string]                             `json:",omitzero,inline"`
	OfElement   *WebScrapeParamsScreenshotParamsAreaElement   `json:",omitzero,inline"`
	OfRectangle *WebScrapeParamsScreenshotParamsAreaRectangle `json:",omitzero,inline"`
	paramUnion
}

func (u WebScrapeParamsScreenshotParamsAreaUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPage, u.OfElement, u.OfRectangle)
}
func (u *WebScrapeParamsScreenshotParamsAreaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type WebScrapeParamsScreenshotParamsAreaPage string

const (
	WebScrapeParamsScreenshotParamsAreaPageViewport WebScrapeParamsScreenshotParamsAreaPage = "viewport"
	WebScrapeParamsScreenshotParamsAreaPageFullPage WebScrapeParamsScreenshotParamsAreaPage = "fullPage"
)

// The property Selector is required.
type WebScrapeParamsScreenshotParamsAreaElement struct {
	// Must match one visible element.
	Selector string `json:"selector" api:"required"`
	paramObj
}

func (r WebScrapeParamsScreenshotParamsAreaElement) MarshalJSON() (data []byte, err error) {
	type shadow WebScrapeParamsScreenshotParamsAreaElement
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScrapeParamsScreenshotParamsAreaElement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pixels from the document origin.
//
// The properties Height, Width, X, Y are required.
type WebScrapeParamsScreenshotParamsAreaRectangle struct {
	Height int64 `json:"height" api:"required"`
	Width  int64 `json:"width" api:"required"`
	X      int64 `json:"x" api:"required"`
	Y      int64 `json:"y" api:"required"`
	paramObj
}

func (r WebScrapeParamsScreenshotParamsAreaRectangle) MarshalJSON() (data []byte, err error) {
	type shadow WebScrapeParamsScreenshotParamsAreaRectangle
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScrapeParamsScreenshotParamsAreaRectangle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Shared browser and content settings. Content filters leave screenshots and
// original bytes unchanged.
type WebScrapeParamsSharedParams struct {
	// Supported two-letter country code, case-insensitive. Applies to every output,
	// including image downloads.
	Country param.Opt[string] `json:"country,omitzero"`
	// Dismiss cookie banners by accepting cookies before actions.
	DismissCookies param.Opt[bool] `json:"dismissCookies,omitzero"`
	// Dismiss other popups before actions.
	DismissPopups param.Opt[bool] `json:"dismissPopups,omitzero"`
	// Include iframe content in extraction. Screenshots show visible frames
	// regardless.
	IncludeFrames param.Opt[bool] `json:"includeFrames,omitzero"`
	// Keep only main content in HTML, Markdown, images, and parsed fields.
	MainContentOnly param.Opt[bool] `json:"mainContentOnly,omitzero"`
	// Settle animations before capture. Defaults to true with screenshots, otherwise
	// false.
	SettleAnimations param.Opt[bool] `json:"settleAnimations,omitzero"`
	// Run in order before capture. A failed action fails the request. Bypasses
	// caching.
	Actions []WebScrapeParamsSharedParamsActionUnion `json:"actions,omitzero"`
	// Remove matching content. Exclusions win.
	ExcludeSelectors []string `json:"excludeSelectors,omitzero"`
	// Headers for the target origin. Requests with custom headers bypass caching.
	Headers map[string]string `json:"headers,omitzero"`
	// Keep matching content after mainContentOnly.
	IncludeSelectors []string `json:"includeSelectors,omitzero"`
	// Document parsing options.
	Parsers WebScrapeParamsSharedParamsParsers `json:"parsers,omitzero"`
	// Override the browser color scheme.
	//
	// Any of "light", "dark".
	Theme string `json:"theme,omitzero"`
	// Browser dimensions in pixels.
	Viewport WebScrapeParamsSharedParamsViewport `json:"viewport,omitzero"`
	// After actions, wait this many milliseconds or until a CSS selector is visible.
	// Defaults to 500 ms, or 2000 ms with frames or an XML URL. Set 0 to skip.
	WaitFor WebScrapeParamsSharedParamsWaitForUnion `json:"waitFor,omitzero"`
	paramObj
}

func (r WebScrapeParamsSharedParams) MarshalJSON() (data []byte, err error) {
	type shadow WebScrapeParamsSharedParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScrapeParamsSharedParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebScrapeParamsSharedParams](
		"theme", "light", "dark",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebScrapeParamsSharedParamsActionUnion struct {
	OfPerform *WebScrapeParamsSharedParamsActionPerform `json:",omitzero,inline"`
	OfScroll  *WebScrapeParamsSharedParamsActionScroll  `json:",omitzero,inline"`
	OfWait    *WebScrapeParamsSharedParamsActionWait    `json:",omitzero,inline"`
	OfWaitFor *WebScrapeParamsSharedParamsActionWaitFor `json:",omitzero,inline"`
	paramUnion
}

func (u WebScrapeParamsSharedParamsActionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPerform, u.OfScroll, u.OfWait, u.OfWaitFor)
}
func (u *WebScrapeParamsSharedParamsActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[WebScrapeParamsSharedParamsActionUnion](
		"type",
		apijson.Discriminator[WebScrapeParamsSharedParamsActionPerform]("perform"),
		apijson.Discriminator[WebScrapeParamsSharedParamsActionScroll]("scroll"),
		apijson.Discriminator[WebScrapeParamsSharedParamsActionWait]("wait"),
		apijson.Discriminator[WebScrapeParamsSharedParamsActionWaitFor]("waitFor"),
	)
}

// The properties Action, Type are required.
type WebScrapeParamsSharedParamsActionPerform struct {
	Action string `json:"action" api:"required"`
	// This field can be elided, and will marshal its zero value as "perform".
	Type constant.Perform `json:"type" default:"perform"`
	paramObj
}

func (r WebScrapeParamsSharedParamsActionPerform) MarshalJSON() (data []byte, err error) {
	type shadow WebScrapeParamsSharedParamsActionPerform
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScrapeParamsSharedParamsActionPerform) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type WebScrapeParamsSharedParamsActionScroll struct {
	MaxScrolls param.Opt[int64] `json:"maxScrolls,omitzero"`
	// Scroll this container. Omit to scroll the page.
	Selector param.Opt[string]                                  `json:"selector,omitzero"`
	Amount   WebScrapeParamsSharedParamsActionScrollAmountUnion `json:"amount,omitzero"`
	// Any of "down", "up", "left", "right".
	Direction string `json:"direction,omitzero"`
	// This field can be elided, and will marshal its zero value as "scroll".
	Type constant.Scroll `json:"type" default:"scroll"`
	paramObj
}

func (r WebScrapeParamsSharedParamsActionScroll) MarshalJSON() (data []byte, err error) {
	type shadow WebScrapeParamsSharedParamsActionScroll
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScrapeParamsSharedParamsActionScroll) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebScrapeParamsSharedParamsActionScroll](
		"direction", "down", "up", "left", "right",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebScrapeParamsSharedParamsActionScrollAmountUnion struct {
	OfInt param.Opt[int64] `json:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWebScrapesSharedParamsActionScrollAmountString)
	OfWebScrapesSharedParamsActionScrollAmountString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u WebScrapeParamsSharedParamsActionScrollAmountUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfWebScrapesSharedParamsActionScrollAmountString)
}
func (u *WebScrapeParamsSharedParamsActionScrollAmountUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type WebScrapeParamsSharedParamsActionScrollAmountString string

const (
	WebScrapeParamsSharedParamsActionScrollAmountStringViewport WebScrapeParamsSharedParamsActionScrollAmountString = "viewport"
	WebScrapeParamsSharedParamsActionScrollAmountStringMax      WebScrapeParamsSharedParamsActionScrollAmountString = "max"
)

// The properties Milliseconds, Type are required.
type WebScrapeParamsSharedParamsActionWait struct {
	Milliseconds int64 `json:"milliseconds" api:"required"`
	// This field can be elided, and will marshal its zero value as "wait".
	Type constant.Wait `json:"type" default:"wait"`
	paramObj
}

func (r WebScrapeParamsSharedParamsActionWait) MarshalJSON() (data []byte, err error) {
	type shadow WebScrapeParamsSharedParamsActionWait
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScrapeParamsSharedParamsActionWait) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Selector, Type are required.
type WebScrapeParamsSharedParamsActionWaitFor struct {
	Selector string `json:"selector" api:"required"`
	// This field can be elided, and will marshal its zero value as "waitFor".
	Type constant.WaitFor `json:"type" default:"waitFor"`
	paramObj
}

func (r WebScrapeParamsSharedParamsActionWaitFor) MarshalJSON() (data []byte, err error) {
	type shadow WebScrapeParamsSharedParamsActionWaitFor
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScrapeParamsSharedParamsActionWaitFor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Document parsing options.
type WebScrapeParamsSharedParamsParsers struct {
	// PDF text options for HTML, Markdown, and parsed fields.
	Pdf WebScrapeParamsSharedParamsParsersPdf `json:"pdf,omitzero"`
	paramObj
}

func (r WebScrapeParamsSharedParamsParsers) MarshalJSON() (data []byte, err error) {
	type shadow WebScrapeParamsSharedParamsParsers
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScrapeParamsSharedParamsParsers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PDF text options for HTML, Markdown, and parsed fields.
type WebScrapeParamsSharedParamsParsersPdf struct {
	// Last page to parse. Must be at least startPage.
	EndPage param.Opt[int64] `json:"endPage,omitzero"`
	// First page to parse, starting at 1.
	StartPage param.Opt[int64] `json:"startPage,omitzero"`
	// Read text from scanned pages.
	//
	// Any of "off", "auto".
	Ocr string `json:"ocr,omitzero"`
	paramObj
}

func (r WebScrapeParamsSharedParamsParsersPdf) MarshalJSON() (data []byte, err error) {
	type shadow WebScrapeParamsSharedParamsParsersPdf
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScrapeParamsSharedParamsParsersPdf) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebScrapeParamsSharedParamsParsersPdf](
		"ocr", "off", "auto",
	)
}

// Browser dimensions in pixels.
type WebScrapeParamsSharedParamsViewport struct {
	Height param.Opt[int64] `json:"height,omitzero"`
	Width  param.Opt[int64] `json:"width,omitzero"`
	paramObj
}

func (r WebScrapeParamsSharedParamsViewport) MarshalJSON() (data []byte, err error) {
	type shadow WebScrapeParamsSharedParamsViewport
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScrapeParamsSharedParamsViewport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebScrapeParamsSharedParamsWaitForUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u WebScrapeParamsSharedParamsWaitForUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *WebScrapeParamsSharedParamsWaitForUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Total deadline, including navigation, actions, waiting, and all outputs.
// Defaults to 60000 milliseconds with behavior fail. Use return-partial to capture
// the current page state and return captured images if image processing cannot
// finish before the deadline; these responses set isPartial and are not cached.
// Every requested format must still be available. Fixed waits must fit before a
// response reserve of up to 5000 milliseconds (at most one quarter of the timeout)
// when using return-partial.
//
// The property Milliseconds is required.
type WebScrapeParamsTimeoutOpts struct {
	// Request deadline in milliseconds. Maximum: 300000 (5 minutes).
	Milliseconds int64 `json:"milliseconds" api:"required"`
	// What to do at the deadline. "fail" returns 408 REQUEST_TIMEOUT without charging
	// credits. "return-partial" returns usable results collected so far; if none are
	// available, the request still fails without charging credits. Partial results are
	// not cached as complete results. "return-partial" requires milliseconds of at
	// least 5000.
	//
	// Any of "fail", "return-partial".
	Behavior string `json:"behavior,omitzero"`
	paramObj
}

func (r WebScrapeParamsTimeoutOpts) MarshalJSON() (data []byte, err error) {
	type shadow WebScrapeParamsTimeoutOpts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScrapeParamsTimeoutOpts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebScrapeParamsTimeoutOpts](
		"behavior", "fail", "return-partial",
	)
}

// Zero data retention. Bypasses caches and uploads; excludes request/response
// content and tags from logs. Must be enabled for your organization.
type WebScrapeParamsZdr string

const (
	WebScrapeParamsZdrEnabled  WebScrapeParamsZdr = "enabled"
	WebScrapeParamsZdrDisabled WebScrapeParamsZdr = "disabled"
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
	// omitted. When combined with timeoutOpts, timeoutOpts.milliseconds must be at
	// least waitForMs + 10000 ms; a shorter deadline is rejected with 400
	// TIMEOUT_TOO_SHORT_FOR_WAIT.
	WaitForMs param.Opt[int64] `query:"waitForMs,omitzero" json:"-"`
	// Optional parameter for comprehensive popup cleanup. If 'true', the browser
	// dismisses detected cookie/consent UI and clears other detected obstructive
	// popups and overlays before capture. If 'false' or not provided, this parameter
	// requests no cleanup; handleCookiePopup can still request cookie/consent handling
	// independently.
	ClearPopups param.Opt[bool] `query:"clearPopups,omitzero" json:"-"`
	// A specific URL to screenshot directly, bypassing domain resolution (e.g.,
	// 'https://example.com/pricing'). When provided, the screenshot is taken of this
	// exact URL. You must provide either 'domain' or 'directUrl', but not both.
	DirectURL param.Opt[string] `query:"directUrl,omitzero" format:"uri" json:"-"`
	// Domain name to take screenshot of (e.g., 'example.com', 'google.com'). The
	// domain will be automatically normalized and validated. You must provide either
	// 'domain' or 'directUrl', but not both.
	Domain param.Opt[string] `query:"domain,omitzero" json:"-"`
	// Optional parameter to control cookie/consent popup handling. If 'true', we
	// dismiss cookie banner before capture. If 'false' or not provided, captures the
	// page without that step.
	HandleCookiePopup param.Opt[bool] `query:"handleCookiePopup,omitzero" json:"-"`
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
	// Optional outbound HTTP headers, using the same JSON object or deep-object query
	// format as other scrape endpoints (for example headers[Authorization]=Bearer
	// token). Headers are scoped to the target origin during capture. For domain/page
	// requests, discovery receives no custom headers and only pages on the resolved
	// origin are eligible. Non-empty headers bypass screenshot caching and return an
	// in-memory data URL; no screenshot is uploaded. Empty objects behave like omitted
	// headers.
	Headers map[string]string `query:"headers,omitzero" json:"-"`
	// Optional parameter to specify which page type to screenshot. If provided, the
	// system will scrape the domain's links and use heuristics to find the most
	// appropriate URL for the specified page type (30 supported languages). If not
	// provided, screenshots the main domain landing page. Only applicable when using
	// 'domain', not 'directUrl'.
	//
	// Any of "login", "signup", "blog", "careers", "pricing", "terms", "privacy",
	// "contact".
	Page WebScreenshotParamsPage `query:"page,omitzero" json:"-"`
	// Comma-separated tags for tracking request usage. Up to 20 tags, each 1-50
	// characters.
	Tags []string `query:"tags,omitzero" json:"-"`
	// Optional request deadline and behavior on timeout. For GET requests, use
	// timeoutOpts[milliseconds]=30000&timeoutOpts[behavior]=fail or a JSON-encoded
	// timeoutOpts object.
	TimeoutOpts WebScreenshotParamsTimeoutOpts `query:"timeoutOpts,omitzero" json:"-"`
	// Optional browser viewport dimensions for the screenshot. Defaults to 1920x1080.
	Viewport WebScreenshotParamsViewport `query:"viewport,omitzero" json:"-"`
	// Set to enabled to bypass shared caches and omit request and response content
	// from retained usage logs. Asset uploads are skipped, so hosted image URLs are
	// omitted. Requires zero data retention to be enabled for your organization
	// (contact support@context.dev), otherwise the request fails with ZDR_NOT_ENABLED.
	// Successful ZDR responses include X-Context-ZDR: true.
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

// Optional request deadline and behavior on timeout. For GET requests, use
// timeoutOpts[milliseconds]=30000&timeoutOpts[behavior]=fail or a JSON-encoded
// timeoutOpts object.
//
// The property Milliseconds is required.
type WebScreenshotParamsTimeoutOpts struct {
	// Request deadline in milliseconds. Maximum: 300000 (5 minutes).
	Milliseconds int64 `query:"milliseconds" api:"required" json:"-"`
	// What to do at the deadline. "fail" returns 408 REQUEST_TIMEOUT without charging
	// credits. "return-partial" returns usable results collected so far; if none are
	// available, the request still fails without charging credits. Partial results are
	// not cached as complete results. "return-partial" requires milliseconds of at
	// least 5000.
	//
	// Any of "fail", "return-partial".
	Behavior string `query:"behavior,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebScreenshotParamsTimeoutOpts]'s query parameters as
// `url.Values`.
func (r WebScreenshotParamsTimeoutOpts) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

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
// from retained usage logs. Asset uploads are skipped, so hosted image URLs are
// omitted. Requires zero data retention to be enabled for your organization
// (contact support@context.dev), otherwise the request fails with ZDR_NOT_ENABLED.
// Successful ZDR responses include X-Context-ZDR: true.
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
	// Optional request deadline and behavior on timeout. For GET requests, use
	// timeoutOpts[milliseconds]=30000&timeoutOpts[behavior]=fail or a JSON-encoded
	// timeoutOpts object.
	TimeoutOpts WebSearchParamsTimeoutOpts `json:"timeoutOpts,omitzero"`
	// Set to enabled to bypass shared caches and omit request and response content
	// from retained usage logs. Asset uploads are skipped, so hosted image URLs are
	// omitted. Requires zero data retention to be enabled for your organization
	// (contact support@context.dev), otherwise the request fails with ZDR_NOT_ENABLED.
	// Successful ZDR responses include X-Context-ZDR: true.
	//
	// Any of "enabled", "disabled".
	Zdr WebSearchParamsZdr `json:"zdr,omitzero"`
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
	// Strip nav, header, footer, and sidebar — keep only the primary article content.
	UseMainContentOnly param.Opt[bool] `json:"useMainContentOnly,omitzero"`
	// Extra wait after page load before rendering, in ms (0–30000). Useful for
	// JS-heavy pages.
	WaitForMs param.Opt[int64] `json:"waitForMs,omitzero"`
	// PDF handling. Use start/end to bound text extraction and OCR to a page range.
	Pdf WebSearchParamsMarkdownOptionsPdf `json:"pdf,omitzero"`
	// Optional request deadline and behavior on timeout. For GET requests, use
	// timeoutOpts[milliseconds]=30000&timeoutOpts[behavior]=fail or a JSON-encoded
	// timeoutOpts object.
	TimeoutOpts WebSearchParamsMarkdownOptionsTimeoutOpts `json:"timeoutOpts,omitzero"`
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

// Optional request deadline and behavior on timeout. For GET requests, use
// timeoutOpts[milliseconds]=30000&timeoutOpts[behavior]=fail or a JSON-encoded
// timeoutOpts object.
//
// The property Milliseconds is required.
type WebSearchParamsMarkdownOptionsTimeoutOpts struct {
	// Request deadline in milliseconds. Maximum: 300000 (5 minutes).
	Milliseconds int64 `json:"milliseconds" api:"required"`
	// What to do at the deadline. "fail" returns 408 REQUEST_TIMEOUT without charging
	// credits. "return-partial" returns usable results collected so far; if none are
	// available, the request still fails without charging credits. Partial results are
	// not cached as complete results. "return-partial" requires milliseconds of at
	// least 5000.
	//
	// Any of "fail", "return-partial".
	Behavior string `json:"behavior,omitzero"`
	paramObj
}

func (r WebSearchParamsMarkdownOptionsTimeoutOpts) MarshalJSON() (data []byte, err error) {
	type shadow WebSearchParamsMarkdownOptionsTimeoutOpts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebSearchParamsMarkdownOptionsTimeoutOpts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebSearchParamsMarkdownOptionsTimeoutOpts](
		"behavior", "fail", "return-partial",
	)
}

// Optional request deadline and behavior on timeout. For GET requests, use
// timeoutOpts[milliseconds]=30000&timeoutOpts[behavior]=fail or a JSON-encoded
// timeoutOpts object.
//
// The property Milliseconds is required.
type WebSearchParamsTimeoutOpts struct {
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

func (r WebSearchParamsTimeoutOpts) MarshalJSON() (data []byte, err error) {
	type shadow WebSearchParamsTimeoutOpts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebSearchParamsTimeoutOpts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebSearchParamsTimeoutOpts](
		"behavior", "fail", "return-partial",
	)
}

// Set to enabled to bypass shared caches and omit request and response content
// from retained usage logs. Asset uploads are skipped, so hosted image URLs are
// omitted. Requires zero data retention to be enabled for your organization
// (contact support@context.dev), otherwise the request fails with ZDR_NOT_ENABLED.
// Successful ZDR responses include X-Context-ZDR: true.
type WebSearchParamsZdr string

const (
	WebSearchParamsZdrEnabled  WebSearchParamsZdr = "enabled"
	WebSearchParamsZdrDisabled WebSearchParamsZdr = "disabled"
)

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
	// Regex pattern. Only URLs matching this pattern will be followed and scraped. An
	// automatic prefix scope in the form ^<starting URL> follows a redirect of the
	// starting page.
	URLRegex param.Opt[string] `json:"urlRegex,omitzero"`
	// Extract only the main content, stripping headers, footers, sidebars, and
	// navigation
	UseMainContentOnly param.Opt[bool] `json:"useMainContentOnly,omitzero"`
	// Browser wait time in milliseconds after initial page load for each crawled page.
	// Defaults to 3500 (3.5 seconds). Min: 0. Max: 30000 (30 seconds).
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
	// Optional request deadline and behavior on timeout. For GET requests, use
	// timeoutOpts[milliseconds]=30000&timeoutOpts[behavior]=fail or a JSON-encoded
	// timeoutOpts object.
	TimeoutOpts WebWebCrawlMdParamsTimeoutOpts `json:"timeoutOpts,omitzero"`
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
	// When true, OCR the selected PDF pages that have no usable text layer (scans),
	// replacing each recovered page's text with the OCR result while pages with a real
	// text layer keep it. Billed at 1 credit per page OCR actually recovered, on top
	// of the base request cost.
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

// Optional request deadline and behavior on timeout. For GET requests, use
// timeoutOpts[milliseconds]=30000&timeoutOpts[behavior]=fail or a JSON-encoded
// timeoutOpts object.
//
// The property Milliseconds is required.
type WebWebCrawlMdParamsTimeoutOpts struct {
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

func (r WebWebCrawlMdParamsTimeoutOpts) MarshalJSON() (data []byte, err error) {
	type shadow WebWebCrawlMdParamsTimeoutOpts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebWebCrawlMdParamsTimeoutOpts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebWebCrawlMdParamsTimeoutOpts](
		"behavior", "fail", "return-partial",
	)
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
