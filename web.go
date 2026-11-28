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

// Research a task across the live web and Context.dev's company and people data,
// and get the answer in the JSON shape you provide, with the URLs it used. `fast`
// (10 credits) suits short lookups; `ultra` (100 credits, the default) reads more
// sources for deeper research.
func (r *WebService) Answers(ctx context.Context, body WebAnswersParams, opts ...option.RequestOption) (res *WebAnswersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "web/answers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Find a company's direct competitors from its website and web search results.
// Each competitor comes with a confidence level and the search results behind it.
func (r *WebService) ExtractCompetitors(ctx context.Context, query WebExtractCompetitorsParams, opts ...option.RequestOption) (res *WebExtractCompetitorsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "web/competitors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Extract a website's design system: colors, typography, spacing, shadows, button
// and card styles, and downloadable font files. Every component style includes
// ready-to-use CSS.
func (r *WebService) ExtractStyleguide(ctx context.Context, query WebExtractStyleguideParams, opts ...option.RequestOption) (res *WebExtractStyleguideResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "web/styleguide"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Map up to 100,000 URLs on a domain, with titles, descriptions, keywords, and
// language when available. Filter with regex patterns, or rank by topic with
// `search` for 1 extra credit. Newly discovered URLs can lack metadata.
func (r *WebService) MapURLs(ctx context.Context, query WebMapURLsParams, opts ...option.RequestOption) (res *WebMapURLsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "web/urls"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Turn any URL into Markdown, HTML, a screenshot, images, raw bytes, parsed
// fields, product data, highlights, or schema-matched JSON in one request. Choose
// the outputs in `formats`; each one succeeds or fails on its own. Handles PDF,
// Word, Excel, PowerPoint, and 40 other file formats.
func (r *WebService) Scrape(ctx context.Context, body WebScrapeParams, opts ...option.RequestOption) (res *WebScrapeResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "web/scrape"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Capture a screenshot of a website's home page, a page type such as pricing, or
// an exact URL. Returns the visible viewport, a scroll slice, or the full page.
func (r *WebService) Screenshot(ctx context.Context, query WebScreenshotParams, opts ...option.RequestOption) (res *WebScreenshotResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "web/screenshot"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Search the web and get ranked results, with optional page Markdown or
// query-relevant passages for each one. Supports search operators and filters for
// domain, freshness, country, and language.
func (r *WebService) Search(ctx context.Context, body WebSearchParams, opts ...option.RequestOption) (res *WebSearchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "web/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Crawl up to 500 pages from a start URL and get each page back as Markdown with
// metadata. Control scope with depth limits, regex patterns, and subdomain rules.
// For larger crawls, use a batch.
func (r *WebService) WebCrawlMd(ctx context.Context, body WebWebCrawlMdParams, opts ...option.RequestOption) (res *WebWebCrawlMdResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "web/crawl"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type WebAnswersResponse struct {
	// The answer, in the shape of `json_format`.
	JsonContent map[string]any `json:"json_content" api:"required"`
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// URLs of the pages and company or profile records read for the answer, followed
	// by URLs cited in `json_content`.
	Sources []string `json:"sources" api:"required"`
	// Credits this request used and your remaining balance.
	KeyMetadata WebAnswersResponseKeyMetadata `json:"key_metadata"`
	// `true` when the answer is incomplete because the deadline ended research or a
	// requested list item has no researched values.
	Partial bool `json:"partial"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JsonContent respjson.Field
		RequestID   respjson.Field
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

// Credits this request used and your remaining balance.
type WebAnswersResponseKeyMetadata struct {
	// Credits charged for this request.
	CreditsConsumed int64 `json:"credits_consumed" api:"required"`
	// Credits your organization has left after this request.
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
	// Direct competitors, ordered by relevance and confidence.
	Competitors []WebExtractCompetitorsResponseCompetitor `json:"competitors" api:"required"`
	// Normalized domain from the request.
	Domain string `json:"domain" api:"required"`
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Always `ok` on success.
	//
	// Any of "ok".
	Status WebExtractCompetitorsResponseStatus `json:"status" api:"required"`
	// Profile of the requested company, inferred from its website.
	Target WebExtractCompetitorsResponseTarget `json:"target" api:"required"`
	// Credits this request used and your remaining balance.
	KeyMetadata WebExtractCompetitorsResponseKeyMetadata `json:"key_metadata"`
	// `true` when the deadline passed with `timeoutOpts.behavior` set to
	// `return-partial`. Only results finished in time are included.
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
	// Short description of what the competitor offers.
	Description string `json:"description" api:"required"`
	// Competitor's official domain, normalized.
	Domain string `json:"domain" api:"required"`
	// Competitor company or product name.
	Name string `json:"name" api:"required"`
	// Search result URLs used as evidence for this competitor.
	SourceURLs []string `json:"sourceUrls" api:"required"`
	// Competitor's website URL.
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

// Always `ok` on success.
type WebExtractCompetitorsResponseStatus string

const (
	WebExtractCompetitorsResponseStatusOk WebExtractCompetitorsResponseStatus = "ok"
)

// Profile of the requested company, inferred from its website.
type WebExtractCompetitorsResponseTarget struct {
	// Company or product name, inferred from its website.
	CompanyName string `json:"companyName" api:"required"`
	// Product category or market the company operates in.
	Field string `json:"field" api:"required"`
	// One-sentence summary of what the company sells and who it serves.
	FieldDescription string `json:"fieldDescription" api:"required"`
	// Website URL that was analyzed.
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

// Credits this request used and your remaining balance.
type WebExtractCompetitorsResponseKeyMetadata struct {
	// Credits charged for this request.
	CreditsConsumed int64 `json:"credits_consumed" api:"required"`
	// Credits your organization has left after this request.
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
	// Cache status of this response.
	CacheMetadata WebExtractStyleguideResponseCacheMetadata `json:"cache_metadata" api:"required"`
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// HTTP status code, repeated in the body.
	Code int64 `json:"code"`
	// Normalized domain that was analyzed.
	Domain string `json:"domain"`
	// `loaded`, or `still-loading` when capture ended before the page finished
	// loading.
	//
	// Any of "loaded", "still-loading".
	FinalDomState WebExtractStyleguideResponseFinalDomState `json:"finalDOMState"`
	// Credits this request used and your remaining balance.
	KeyMetadata WebExtractStyleguideResponseKeyMetadata `json:"key_metadata"`
	// Always `ok` on success.
	Status string `json:"status"`
	// Design tokens and component styles extracted from the website.
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

// Cache status of this response.
type WebExtractStyleguideResponseCacheMetadata struct {
	// Age of the cached data, in milliseconds. `0` for `miss` and `zdr`.
	AgeMs int64 `json:"age_ms" api:"required"`
	// `hit` when served from cache, `miss` when fetched fresh, `zdr` when zero data
	// retention bypassed the cache.
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

// `loaded`, or `still-loading` when capture ended before the page finished
// loading.
type WebExtractStyleguideResponseFinalDomState string

const (
	WebExtractStyleguideResponseFinalDomStateLoaded       WebExtractStyleguideResponseFinalDomState = "loaded"
	WebExtractStyleguideResponseFinalDomStateStillLoading WebExtractStyleguideResponseFinalDomState = "still-loading"
)

// Credits this request used and your remaining balance.
type WebExtractStyleguideResponseKeyMetadata struct {
	// Credits charged for this request.
	CreditsConsumed int64 `json:"credits_consumed" api:"required"`
	// Credits your organization has left after this request.
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

// Design tokens and component styles extracted from the website.
type WebExtractStyleguideResponseStyleguide struct {
	// Core colors used on the website.
	Colors WebExtractStyleguideResponseStyleguideColors `json:"colors" api:"required"`
	// Styles for buttons and cards.
	Components WebExtractStyleguideResponseStyleguideComponents `json:"components" api:"required"`
	// Spacing scale used on the website, from `xs` to `xl`.
	ElementSpacing WebExtractStyleguideResponseStyleguideElementSpacing `json:"elementSpacing" api:"required"`
	// Font files keyed by family name, as it appears in `fontFamily` and
	// `fontFallbacks`. Generic families are left out. Match style and weight against
	// `faces` to pick a file.
	FontLinks map[string]WebExtractStyleguideResponseStyleguideFontLink `json:"fontLinks" api:"required"`
	// Whether the site's design is mainly light or dark.
	//
	// Any of "light", "dark".
	Mode string `json:"mode" api:"required"`
	// Box shadow scale used on the website, from `sm` to `xl`, plus `inner`.
	Shadows WebExtractStyleguideResponseStyleguideShadows `json:"shadows" api:"required"`
	// Text styles for headings and body text.
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

// Core colors used on the website.
type WebExtractStyleguideResponseStyleguideColors struct {
	// Accent color, in hex.
	Accent string `json:"accent" api:"required"`
	// Background color, in hex.
	Background string `json:"background" api:"required"`
	// Text color, in hex.
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

// Styles for buttons and cards.
type WebExtractStyleguideResponseStyleguideComponents struct {
	// Primary, secondary, and link button styles.
	Button WebExtractStyleguideResponseStyleguideComponentsButton `json:"button" api:"required"`
	// Card style.
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

// Primary, secondary, and link button styles.
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
	// Border color in hex: `#RRGGBB`, or `#RRGGBBAA` when the border is translucent.
	BorderColor  string `json:"borderColor" api:"required"`
	BorderRadius string `json:"borderRadius" api:"required"`
	BorderStyle  string `json:"borderStyle" api:"required"`
	BorderWidth  string `json:"borderWidth" api:"required"`
	// Computed `box-shadow`, with layers separated by commas.
	BoxShadow string `json:"boxShadow" api:"required"`
	Color     string `json:"color" api:"required"`
	// Ready-to-use CSS declarations for this style.
	Css        string  `json:"css" api:"required"`
	FontSize   string  `json:"fontSize" api:"required"`
	FontWeight float64 `json:"fontWeight" api:"required"`
	// Minimum height as a CSS length, usually in pixels.
	MinHeight string `json:"minHeight" api:"required"`
	// Minimum width as a CSS length, usually in pixels.
	MinWidth       string `json:"minWidth" api:"required"`
	Padding        string `json:"padding" api:"required"`
	TextDecoration string `json:"textDecoration" api:"required"`
	// Full font stack from the computed `font-family`, in order.
	FontFallbacks []string `json:"fontFallbacks"`
	// Button typeface, the first entry in `fontFallbacks`.
	FontFamily string `json:"fontFamily"`
	// Underline color in hex, when it differs from the text color.
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
	// Border color in hex: `#RRGGBB`, or `#RRGGBBAA` when the border is translucent.
	BorderColor  string `json:"borderColor" api:"required"`
	BorderRadius string `json:"borderRadius" api:"required"`
	BorderStyle  string `json:"borderStyle" api:"required"`
	BorderWidth  string `json:"borderWidth" api:"required"`
	// Computed `box-shadow`, with layers separated by commas.
	BoxShadow string `json:"boxShadow" api:"required"`
	Color     string `json:"color" api:"required"`
	// Ready-to-use CSS declarations for this style.
	Css        string  `json:"css" api:"required"`
	FontSize   string  `json:"fontSize" api:"required"`
	FontWeight float64 `json:"fontWeight" api:"required"`
	// Minimum height as a CSS length, usually in pixels.
	MinHeight string `json:"minHeight" api:"required"`
	// Minimum width as a CSS length, usually in pixels.
	MinWidth       string `json:"minWidth" api:"required"`
	Padding        string `json:"padding" api:"required"`
	TextDecoration string `json:"textDecoration" api:"required"`
	// Full font stack from the computed `font-family`, in order.
	FontFallbacks []string `json:"fontFallbacks"`
	// Button typeface, the first entry in `fontFallbacks`.
	FontFamily string `json:"fontFamily"`
	// Underline color in hex, when it differs from the text color.
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
	// Border color in hex: `#RRGGBB`, or `#RRGGBBAA` when the border is translucent.
	BorderColor  string `json:"borderColor" api:"required"`
	BorderRadius string `json:"borderRadius" api:"required"`
	BorderStyle  string `json:"borderStyle" api:"required"`
	BorderWidth  string `json:"borderWidth" api:"required"`
	// Computed `box-shadow`, with layers separated by commas.
	BoxShadow string `json:"boxShadow" api:"required"`
	Color     string `json:"color" api:"required"`
	// Ready-to-use CSS declarations for this style.
	Css        string  `json:"css" api:"required"`
	FontSize   string  `json:"fontSize" api:"required"`
	FontWeight float64 `json:"fontWeight" api:"required"`
	// Minimum height as a CSS length, usually in pixels.
	MinHeight string `json:"minHeight" api:"required"`
	// Minimum width as a CSS length, usually in pixels.
	MinWidth       string `json:"minWidth" api:"required"`
	Padding        string `json:"padding" api:"required"`
	TextDecoration string `json:"textDecoration" api:"required"`
	// Full font stack from the computed `font-family`, in order.
	FontFallbacks []string `json:"fontFallbacks"`
	// Button typeface, the first entry in `fontFallbacks`.
	FontFamily string `json:"fontFamily"`
	// Underline color in hex, when it differs from the text color.
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

// Card style.
type WebExtractStyleguideResponseStyleguideComponentsCard struct {
	BackgroundColor string `json:"backgroundColor" api:"required"`
	// Border color in hex: `#RRGGBB`, or `#RRGGBBAA` when the border is translucent.
	BorderColor  string `json:"borderColor" api:"required"`
	BorderRadius string `json:"borderRadius" api:"required"`
	BorderStyle  string `json:"borderStyle" api:"required"`
	BorderWidth  string `json:"borderWidth" api:"required"`
	BoxShadow    string `json:"boxShadow" api:"required"`
	// Ready-to-use CSS declarations for this style.
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

// Spacing scale used on the website, from `xs` to `xl`.
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
	// `google` for Google Fonts, `custom` for any other font source.
	//
	// Any of "google", "custom".
	Type string `json:"type" api:"required"`
	// Google Fonts category, such as `sans-serif`, `serif`, `monospace`, `display`, or
	// `handwriting`. Omitted for custom fonts when unknown.
	Category string `json:"category"`
	// Human-readable family name: the `fontLinks` key for custom fonts, or the
	// official Google Fonts name for Google fonts.
	DisplayName string `json:"displayName"`
	// Font faces found for this family, one per style, weight, and Unicode subset.
	Faces []WebExtractStyleguideResponseStyleguideFontLinkFace `json:"faces"`
	// Deprecated upright font lookup keyed by weight string. A variable font can map
	// several standard weight keys to the same URL; faces retains its actual weight
	// range.
	//
	// Deprecated: deprecated
	Files map[string]string `json:"files"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Category    respjson.Field
		DisplayName respjson.Field
		Faces       respjson.Field
		Files       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebExtractStyleguideResponseStyleguideFontLink) RawJSON() string { return r.JSON.raw }
func (r *WebExtractStyleguideResponseStyleguideFontLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebExtractStyleguideResponseStyleguideFontLinkFace struct {
	// CSS `font-style` descriptor: `normal`, `italic`, or `oblique`, with an optional
	// angle or angle range such as `oblique 10deg 20deg`.
	Style string `json:"style" api:"required"`
	// URL of the font file, or an inline data URL.
	URL string `json:"url" api:"required"`
	// CSS `font-weight` descriptor. A variable font can have a range such as
	// `300 500`, which covers every weight in between with one file.
	Weight string `json:"weight" api:"required"`
	// CSS `unicode-range` descriptor.
	UnicodeRange string `json:"unicodeRange"`
	// CSS `font-width` or `font-stretch` descriptor.
	Width string `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Style        respjson.Field
		URL          respjson.Field
		Weight       respjson.Field
		UnicodeRange respjson.Field
		Width        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebExtractStyleguideResponseStyleguideFontLinkFace) RawJSON() string { return r.JSON.raw }
func (r *WebExtractStyleguideResponseStyleguideFontLinkFace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Box shadow scale used on the website, from `sm` to `xl`, plus `inner`.
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

// Text styles for headings and body text.
type WebExtractStyleguideResponseStyleguideTypography struct {
	// Styles for `h1` through `h4` headings.
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

// Styles for `h1` through `h4` headings.
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
	// Full font stack from the computed `font-family`, in order.
	FontFallbacks []string `json:"fontFallbacks" api:"required"`
	// First font in the stack.
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
	// Full font stack from the computed `font-family`, in order.
	FontFallbacks []string `json:"fontFallbacks" api:"required"`
	// First font in the stack.
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
	// Full font stack from the computed `font-family`, in order.
	FontFallbacks []string `json:"fontFallbacks" api:"required"`
	// First font in the stack.
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
	// Full font stack from the computed `font-family`, in order.
	FontFallbacks []string `json:"fontFallbacks" api:"required"`
	// First font in the stack.
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
	// Full font stack from the computed `font-family`, in order.
	FontFallbacks []string `json:"fontFallbacks" api:"required"`
	// First font in the stack.
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
	// Normalized domain that was mapped.
	Domain string `json:"domain" api:"required"`
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Always `true` on success.
	//
	// Any of true.
	Success bool `json:"success" api:"required"`
	// URLs found on the domain, with page metadata when available.
	URLs []WebMapURLsResponseURL `json:"urls" api:"required"`
	// Credits this request used and your remaining balance.
	KeyMetadata WebMapURLsResponseKeyMetadata `json:"key_metadata"`
	// `true` when the deadline passed with `timeoutOpts.behavior` set to
	// `return-partial`. Only results finished in time are included.
	Partial bool `json:"partial"`
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
	// Page URL.
	URL string `json:"url" api:"required"`
	// Page description, when available.
	Description string `json:"description"`
	// Keywords from the page's `keywords` meta tag.
	Keywords []string `json:"keywords"`
	// Page language, when available.
	Language string `json:"language"`
	// Page title, when available.
	Title string `json:"title"`
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

// Credits this request used and your remaining balance.
type WebMapURLsResponseKeyMetadata struct {
	// Credits charged for this request.
	CreditsConsumed int64 `json:"credits_consumed" api:"required"`
	// Credits your organization has left after this request.
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
	// The original HTTP response body, unchanged by waits, actions, and filters.
	Bytes WebScrapeResponseBytes `json:"bytes" api:"required"`
	// Cache status of this response.
	CacheMetadata WebScrapeResponseCacheMetadata `json:"cache_metadata" api:"required"`
	// Relevant Markdown excerpts in page order. `[Heading]` adds context; `…` marks
	// omitted text.
	Highlights WebScrapeResponseHighlights `json:"highlights" api:"required"`
	// Rendered HTML of the page, after content filters.
	HTML WebScrapeResponseHTML `json:"html" api:"required"`
	// Images found on the page, after content filters. `[]` when there are none.
	Images WebScrapeResponseImages `json:"images" api:"required"`
	// Object matching `jsonParams.schema`.
	Json WebScrapeResponseJson `json:"json" api:"required"`
	// Page content as Markdown, after content filters.
	Markdown WebScrapeResponseMarkdown `json:"markdown" api:"required"`
	// Page metadata. Fields are omitted when not found.
	Metadata WebScrapeResponseMetadata `json:"metadata" api:"required"`
	// Fields extracted by `parseParams`. Custom rules return each field after content
	// filters, with `null` when nothing matches (`[]` for `list` rules). Prebuilt
	// parsers return the record or lists described under `parseParams.parser`.
	Parsed WebScrapeResponseParsed `json:"parsed" api:"required"`
	// Product details found on the page.
	Product WebScrapeResponseProduct `json:"product" api:"required"`
	// Browser profile used for this scrape. `null` when no profile was requested.
	Profile WebScrapeResponseProfile `json:"profile" api:"required"`
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Screenshot as a base64 image data URL.
	Screenshot WebScrapeResponseScreenshot `json:"screenshot" api:"required"`
	// Final URL after redirects and browser actions.
	URL string `json:"url" api:"required" format:"uri"`
	// `true` when at least one requested output succeeded but others failed or are
	// incomplete. Absent otherwise.
	//
	// Any of true.
	IsPartial bool `json:"isPartial"`
	// Credits this request used and your remaining balance.
	KeyMetadata WebScrapeResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bytes         respjson.Field
		CacheMetadata respjson.Field
		Highlights    respjson.Field
		HTML          respjson.Field
		Images        respjson.Field
		Json          respjson.Field
		Markdown      respjson.Field
		Metadata      respjson.Field
		Parsed        respjson.Field
		Product       respjson.Field
		Profile       respjson.Field
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

// The original HTTP response body, unchanged by waits, actions, and filters.
type WebScrapeResponseBytes struct {
	Data WebScrapeResponseBytesData `json:"data" api:"required"`
	// Whether you requested this output in `formats`.
	Requested bool `json:"requested" api:"required"`
	// `true` when the output succeeded, `false` when it failed, `null` when it was not
	// requested.
	Success bool `json:"success" api:"required"`
	// Machine-readable error code. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Human-readable explanation of the error and how to resolve it.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		Success     respjson.Field
		ErrorCode   respjson.Field
		Message     respjson.Field
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
	// Response body as base64, after HTTP decompression. Up to 50 MiB decoded.
	Base64 string `json:"base64" api:"required"`
	// `Content-Type` of the response.
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

// Cache status of this response.
type WebScrapeResponseCacheMetadata struct {
	// Age of the cached data, in milliseconds. `0` for `miss` and `zdr`.
	AgeMs int64 `json:"age_ms" api:"required"`
	// `hit` when served from cache, `miss` when fetched fresh, `zdr` when zero data
	// retention bypassed the cache.
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

// Relevant Markdown excerpts in page order. `[Heading]` adds context; `…` marks
// omitted text.
type WebScrapeResponseHighlights struct {
	Data []string `json:"data" api:"required"`
	// Whether you requested this output in `formats`.
	Requested bool `json:"requested" api:"required"`
	// `true` when the output succeeded, `false` when it failed, `null` when it was not
	// requested.
	Success bool `json:"success" api:"required"`
	// Machine-readable error code. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Human-readable explanation of the error and how to resolve it.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		Success     respjson.Field
		ErrorCode   respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScrapeResponseHighlights) RawJSON() string { return r.JSON.raw }
func (r *WebScrapeResponseHighlights) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rendered HTML of the page, after content filters.
type WebScrapeResponseHTML struct {
	Data string `json:"data" api:"required"`
	// Whether you requested this output in `formats`.
	Requested bool `json:"requested" api:"required"`
	// `true` when the output succeeded, `false` when it failed, `null` when it was not
	// requested.
	Success bool `json:"success" api:"required"`
	// Machine-readable error code. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Human-readable explanation of the error and how to resolve it.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		Success     respjson.Field
		ErrorCode   respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScrapeResponseHTML) RawJSON() string { return r.JSON.raw }
func (r *WebScrapeResponseHTML) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Images found on the page, after content filters. `[]` when there are none.
type WebScrapeResponseImages struct {
	Data []WebScrapeResponseImagesData `json:"data" api:"required"`
	// Whether you requested this output in `formats`.
	Requested bool `json:"requested" api:"required"`
	// `true` when the output succeeded, `false` when it failed, `null` when it was not
	// requested.
	Success bool `json:"success" api:"required"`
	// Machine-readable error code. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Human-readable explanation of the error and how to resolve it.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		Success     respjson.Field
		ErrorCode   respjson.Field
		Message     respjson.Field
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
	// Alt text. `null` when the image has none.
	Alt string `json:"alt" api:"required"`
	// Image URL, or a data URI for inline images.
	URL string `json:"url" api:"required" format:"uri"`
	// What the image shows, such as a photo, logo, or icon. Requires `classification`
	// enrichment.
	//
	// Any of "photography", "illustration", "logo", "wordmark", "icon", "pattern",
	// "graphic", "other".
	Classification string `json:"classification"`
	// Hosted copy of the image, valid for 24 hours after capture. Requires `file`
	// enrichment; not returned with ZDR.
	FileURL string `json:"fileUrl" format:"uri"`
	// Image height in pixels. Requires `dimensions` enrichment.
	Height int64 `json:"height"`
	// Image width in pixels. Requires `dimensions` enrichment.
	Width int64 `json:"width"`
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

// Object matching `jsonParams.schema`.
type WebScrapeResponseJson struct {
	Data map[string]any `json:"data" api:"required"`
	// Whether you requested this output in `formats`.
	Requested bool `json:"requested" api:"required"`
	// `true` when the output succeeded, `false` when it failed, `null` when it was not
	// requested.
	Success bool `json:"success" api:"required"`
	// Machine-readable error code. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// `true` when the page was too long to read in full, so values that appear only in
	// the unread part may be missing.
	//
	// Any of true.
	IsTruncated bool `json:"isTruncated"`
	// Human-readable explanation of the error and how to resolve it.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		Success     respjson.Field
		ErrorCode   respjson.Field
		IsTruncated respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScrapeResponseJson) RawJSON() string { return r.JSON.raw }
func (r *WebScrapeResponseJson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Page content as Markdown, after content filters.
type WebScrapeResponseMarkdown struct {
	Data string `json:"data" api:"required"`
	// Whether you requested this output in `formats`.
	Requested bool `json:"requested" api:"required"`
	// `true` when the output succeeded, `false` when it failed, `null` when it was not
	// requested.
	Success bool `json:"success" api:"required"`
	// Machine-readable error code. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Human-readable explanation of the error and how to resolve it.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		Success     respjson.Field
		ErrorCode   respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScrapeResponseMarkdown) RawJSON() string { return r.JSON.raw }
func (r *WebScrapeResponseMarkdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Page metadata. Fields are omitted when not found.
type WebScrapeResponseMetadata struct {
	// Other non-social meta tags that don't have their own field.
	AdditionalMeta map[string]WebScrapeResponseMetadataAdditionalMetaUnion `json:"additionalMeta"`
	// Alternate links from `link rel="alternate"` tags, as absolute URLs.
	Alternates []WebScrapeResponseMetadataAlternate `json:"alternates"`
	// Author, when present.
	Author string `json:"author"`
	// Absolute canonical URL, when present.
	CanonicalURL string `json:"canonicalUrl"`
	// Best available description, from standard, Open Graph, or Twitter meta tags.
	Description string `json:"description"`
	// Absolute favicon URL, when present.
	Favicon string `json:"favicon"`
	// Headings (`h1` to `h6`) in document order, from the full page before any
	// filtering. Up to 500. Omitted when the page has none.
	Headings []WebScrapeResponseMetadataHeading `json:"headings"`
	// Main preview image URL, from Open Graph, Twitter, or image meta tags.
	Image string `json:"image"`
	// JSON-LD structured data blocks parsed from the page.
	JsonLd []map[string]any `json:"jsonLd"`
	// Keywords from the page's `keywords` meta tag.
	Keywords []string `json:"keywords"`
	// Page language, from the `html` `lang` attribute or language meta tags.
	Language string `json:"language"`
	// Last-modified date or timestamp from page metadata, when present.
	ModifiedTime string `json:"modifiedTime"`
	// Open Graph tags, with the `og:` prefix removed and keys in camelCase.
	OpenGraph map[string]WebScrapeResponseMetadataOpenGraphUnion `json:"openGraph"`
	// Publish date or timestamp from page metadata, when present.
	PublishedTime string `json:"publishedTime"`
	// `robots` meta directive, when present.
	Robots string `json:"robots"`
	// Site or app name from page metadata.
	SiteName string `json:"siteName"`
	// Best available page title.
	Title string `json:"title"`
	// Twitter card tags, with the `twitter:` prefix removed and keys in camelCase.
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
	// Absolute URL of the alternate version.
	Href string `json:"href" api:"required"`
	// Language or locale of the alternate version, when present.
	Hreflang string `json:"hreflang"`
	// Title of the alternate version, when present.
	Title string `json:"title"`
	// MIME type of the alternate version, when present.
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
	// Heading level, from 1 (`h1`) to 6 (`h6`).
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

// Fields extracted by `parseParams`. Custom rules return each field after content
// filters, with `null` when nothing matches (`[]` for `list` rules). Prebuilt
// parsers return the record or lists described under `parseParams.parser`.
type WebScrapeResponseParsed struct {
	Data map[string]any `json:"data" api:"required"`
	// Whether you requested this output in `formats`.
	Requested bool `json:"requested" api:"required"`
	// `true` when the output succeeded, `false` when it failed, `null` when it was not
	// requested.
	Success bool `json:"success" api:"required"`
	// Prebuilt parser that ran, or `custom` for custom rules or when `parse` was not
	// requested.
	//
	// Any of "custom", "google-maps", "google-maps-reviews", "google-play",
	// "google-play-reviews", "app-store", "app-store-reviews", "g2-reviews",
	// "instagram-profile", "instagram-post", "instagram-feed", "tiktok-profile",
	// "tiktok-video", "tiktok-feed", "tiktok-comments", "tiktok-followers",
	// "tiktok-following", "tiktok-user-search", "reddit-post", "reddit-search",
	// "reddit-subreddit", "reddit-user", "yelp-search", "yelp-reviews",
	// "tripadvisor-reviews", "car-rental-availability", "car-rental-locations",
	// "emails", "phone-numbers", "links".
	Type string `json:"type" api:"required"`
	// Machine-readable error code. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Human-readable explanation of the error and how to resolve it.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		Success     respjson.Field
		Type        respjson.Field
		ErrorCode   respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScrapeResponseParsed) RawJSON() string { return r.JSON.raw }
func (r *WebScrapeResponseParsed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Product details found on the page.
type WebScrapeResponseProduct struct {
	Data WebScrapeResponseProductData `json:"data" api:"required"`
	// Whether you requested this output in `formats`.
	Requested bool `json:"requested" api:"required"`
	// `true` when the output succeeded, `false` when it failed, `null` when it was not
	// requested.
	Success bool `json:"success" api:"required"`
	// Machine-readable error code. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Human-readable explanation of the error and how to resolve it.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		Success     respjson.Field
		ErrorCode   respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScrapeResponseProduct) RawJSON() string { return r.JSON.raw }
func (r *WebScrapeResponseProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScrapeResponseProductData struct {
	// Whether the page is a product detail page.
	IsProductPage bool `json:"isProductPage" api:"required"`
	// The extracted product. `null` when the page is not a product detail page.
	Product WebScrapeResponseProductDataProduct `json:"product" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsProductPage respjson.Field
		Product       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScrapeResponseProductData) RawJSON() string { return r.JSON.raw }
func (r *WebScrapeResponseProductData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The extracted product. `null` when the page is not a product detail page.
type WebScrapeResponseProductDataProduct struct {
	// Stock or ordering availability.
	//
	// Any of "in_stock", "out_of_stock", "limited_availability", "preorder",
	// "backorder", "made_to_order", "discontinued".
	Availability string `json:"availability" api:"required"`
	// Brand or vendor.
	Brand string `json:"brand" api:"required"`
	// Product category.
	Category string `json:"category" api:"required"`
	// ISO 4217 currency code of `price`.
	Currency string `json:"currency" api:"required"`
	// Product description.
	Description string `json:"description" api:"required"`
	// Product dimensions as shown on the page.
	Dimensions []string `json:"dimensions" api:"required"`
	// Key features and specifications.
	Features []string `json:"features" api:"required"`
	// Product image URLs, main image first.
	Images []string `json:"images" api:"required" format:"uri"`
	// Main product image URL.
	ImageURL string `json:"imageUrl" api:"required" format:"uri"`
	// Product name.
	Name string `json:"name" api:"required"`
	// Current price.
	Price float64 `json:"price" api:"required"`
	// List price before any discount.
	RegularPrice float64 `json:"regularPrice" api:"required"`
	// Product identifier such as a SKU or model number.
	SKU string `json:"sku" api:"required"`
	// Product tags.
	Tags []string `json:"tags" api:"required"`
	// Intended audience.
	TargetAudience []string `json:"targetAudience" api:"required"`
	// Product variants, such as colors or sizes, with their attributes and images.
	// `[]` when none are found. May not include every variant the store offers.
	Variants []WebScrapeResponseProductDataProductVariant `json:"variants" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Availability   respjson.Field
		Brand          respjson.Field
		Category       respjson.Field
		Currency       respjson.Field
		Description    respjson.Field
		Dimensions     respjson.Field
		Features       respjson.Field
		Images         respjson.Field
		ImageURL       respjson.Field
		Name           respjson.Field
		Price          respjson.Field
		RegularPrice   respjson.Field
		SKU            respjson.Field
		Tags           respjson.Field
		TargetAudience respjson.Field
		Variants       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScrapeResponseProductDataProduct) RawJSON() string { return r.JSON.raw }
func (r *WebScrapeResponseProductDataProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScrapeResponseProductDataProductVariant struct {
	// Attributes that define this variant, such as color, size, material, or pattern,
	// as stated on the page.
	Attributes map[string]string `json:"attributes" api:"required"`
	// Image URLs the source attaches to this variant.
	Images []string `json:"images" api:"required"`
	// SKU of this variant. `null` when the source does not state one.
	SKU string `json:"sku" api:"required"`
	// URL of this variant or offer, when the source provides one. Several variants can
	// share a URL.
	URL string `json:"url" api:"required"`
	// Stock or ordering availability of this variant.
	//
	// Any of "in_stock", "out_of_stock", "limited_availability", "preorder",
	// "backorder", "made_to_order", "discontinued".
	Availability string `json:"availability" api:"nullable"`
	// Currency code of `price`, such as `USD` or `EUR`, when stated.
	Currency string `json:"currency" api:"nullable"`
	// Price of this variant, when the source states one per variant.
	Price float64 `json:"price" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attributes   respjson.Field
		Images       respjson.Field
		SKU          respjson.Field
		URL          respjson.Field
		Availability respjson.Field
		Currency     respjson.Field
		Price        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScrapeResponseProductDataProductVariant) RawJSON() string { return r.JSON.raw }
func (r *WebScrapeResponseProductDataProductVariant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Browser profile used for this scrape. `null` when no profile was requested.
type WebScrapeResponseProfile struct {
	// ID of the browser profile. After creating a profile, save it and pass it as
	// `sharedParams.profile.profileId` with `type: "existing"`.
	ID string `json:"id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScrapeResponseProfile) RawJSON() string { return r.JSON.raw }
func (r *WebScrapeResponseProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot as a base64 image data URL.
type WebScrapeResponseScreenshot struct {
	Data string `json:"data" api:"required" format:"uri"`
	// Whether you requested this output in `formats`.
	Requested bool `json:"requested" api:"required"`
	// `true` when the output succeeded, `false` when it failed, `null` when it was not
	// requested.
	Success bool `json:"success" api:"required"`
	// Machine-readable error code. Present only when `success` is `false`.
	ErrorCode string `json:"error_code"`
	// Human-readable explanation of the error and how to resolve it.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Requested   respjson.Field
		Success     respjson.Field
		ErrorCode   respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScrapeResponseScreenshot) RawJSON() string { return r.JSON.raw }
func (r *WebScrapeResponseScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credits this request used and your remaining balance.
type WebScrapeResponseKeyMetadata struct {
	// Credits charged for this request.
	CreditsConsumed int64 `json:"credits_consumed" api:"required"`
	// Credits your organization has left after this request.
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
	// Cache status of this response.
	CacheMetadata WebScreenshotResponseCacheMetadata `json:"cache_metadata" api:"required"`
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// HTTP status code, repeated in the body.
	Code int64 `json:"code"`
	// Normalized domain that was captured.
	Domain string `json:"domain"`
	// `loaded`, or `still-loading` when capture ended before the page finished
	// loading.
	//
	// Any of "loaded", "still-loading".
	FinalDomState WebScreenshotResponseFinalDomState `json:"finalDOMState"`
	// Height of the returned image, in pixels.
	Height int64 `json:"height"`
	// Credits this request used and your remaining balance.
	KeyMetadata WebScreenshotResponseKeyMetadata `json:"key_metadata"`
	// URL of the screenshot image. A data URL instead when `zdr` is `enabled` or
	// `headers` are sent.
	Screenshot string `json:"screenshot"`
	// Whether the image covers the viewport or the full page.
	//
	// Any of "viewport", "fullPage".
	ScreenshotType WebScreenshotResponseScreenshotType `json:"screenshotType"`
	// Always `ok` on success.
	Status string `json:"status"`
	// Width of the returned image, in pixels.
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

// Cache status of this response.
type WebScreenshotResponseCacheMetadata struct {
	// Age of the cached data, in milliseconds. `0` for `miss` and `zdr`.
	AgeMs int64 `json:"age_ms" api:"required"`
	// `hit` when served from cache, `miss` when fetched fresh, `zdr` when zero data
	// retention bypassed the cache.
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

// `loaded`, or `still-loading` when capture ended before the page finished
// loading.
type WebScreenshotResponseFinalDomState string

const (
	WebScreenshotResponseFinalDomStateLoaded       WebScreenshotResponseFinalDomState = "loaded"
	WebScreenshotResponseFinalDomStateStillLoading WebScreenshotResponseFinalDomState = "still-loading"
)

// Credits this request used and your remaining balance.
type WebScreenshotResponseKeyMetadata struct {
	// Credits charged for this request.
	CreditsConsumed int64 `json:"credits_consumed" api:"required"`
	// Credits your organization has left after this request.
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

// Whether the image covers the viewport or the full page.
type WebScreenshotResponseScreenshotType string

const (
	WebScreenshotResponseScreenshotTypeViewport WebScreenshotResponseScreenshotType = "viewport"
	WebScreenshotResponseScreenshotTypeFullPage WebScreenshotResponseScreenshotType = "fullPage"
)

type WebSearchResponse struct {
	// Cache status of this response.
	CacheMetadata WebSearchResponseCacheMetadata `json:"cache_metadata" api:"required"`
	// The query you sent.
	Query string `json:"query" api:"required"`
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string                    `json:"request_id" api:"required" format:"uuid"`
	Results   []WebSearchResponseResult `json:"results" api:"required"`
	// Credits this request used and your remaining balance.
	KeyMetadata WebSearchResponseKeyMetadata `json:"key_metadata"`
	// `true` when the deadline passed with `timeoutOpts.behavior` set to
	// `return-partial`. Only results finished in time are included. Partial results
	// are not cached.
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

// Cache status of this response.
type WebSearchResponseCacheMetadata struct {
	// Age of the cached data, in milliseconds. `0` for `miss` and `zdr`.
	AgeMs int64 `json:"age_ms" api:"required"`
	// `hit` when served from cache, `miss` when fetched fresh, `zdr` when zero data
	// retention bypassed the cache.
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
	// Snippet from the page. Empty string when no snippet is available.
	Description string `json:"description" api:"required"`
	// Highlights status and passages for this result.
	Highlights WebSearchResponseResultHighlights `json:"highlights" api:"required"`
	// Markdown scrape status and content for this result.
	Markdown WebSearchResponseResultMarkdown `json:"markdown" api:"required"`
	// How relevant the result is to the query.
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
		Highlights  respjson.Field
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

// Highlights status and passages for this result.
type WebSearchResponseResultHighlights struct {
	// Highlights outcome for this result. Check it before reading `highlights`.
	//
	// Any of "SUCCESS", "NOT_REQUESTED", "TIMEOUT", "CONTENT_TOO_LARGE",
	// "WEBSITE_ACCESS_ERROR", "WEBSITE_BLOCKED", "WEBSITE_NOT_FOUND", "ERROR".
	Code string `json:"code" api:"required"`
	// Passages relevant to the query, in page order. `null` unless
	// `highlightsOptions.enabled` is `true` and the page was read.
	Highlights []string `json:"highlights" api:"required"`
	// Why this output is missing. Present only when `code` is a failure.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Highlights  respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebSearchResponseResultHighlights) RawJSON() string { return r.JSON.raw }
func (r *WebSearchResponseResultHighlights) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Markdown scrape status and content for this result.
type WebSearchResponseResultMarkdown struct {
	// Scrape outcome for this result. Check it before reading `markdown`.
	//
	// Any of "SUCCESS", "NOT_REQUESTED", "TIMEOUT", "CONTENT_TOO_LARGE",
	// "WEBSITE_ACCESS_ERROR", "WEBSITE_BLOCKED", "WEBSITE_NOT_FOUND", "ERROR".
	Code string `json:"code" api:"required"`
	// Page content as GitHub-flavored Markdown. `null` unless
	// `markdownOptions.enabled` is `true` and the scrape succeeded.
	Markdown string `json:"markdown" api:"required"`
	// `loaded`, or `still-loading` when capture ended before the page finished
	// loading.
	//
	// Any of "loaded", "still-loading".
	FinalDomState string `json:"finalDOMState"`
	// Why this output is missing. Present only when `code` is a failure.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code          respjson.Field
		Markdown      respjson.Field
		FinalDomState respjson.Field
		Message       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebSearchResponseResultMarkdown) RawJSON() string { return r.JSON.raw }
func (r *WebSearchResponseResultMarkdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credits this request used and your remaining balance.
type WebSearchResponseKeyMetadata struct {
	// Credits charged for this request.
	CreditsConsumed int64 `json:"credits_consumed" api:"required"`
	// Credits your organization has left after this request.
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
	// Cache status of this response.
	CacheMetadata WebWebCrawlMdResponseCacheMetadata `json:"cache_metadata" api:"required"`
	Metadata      WebWebCrawlMdResponseMetadata      `json:"metadata" api:"required"`
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string                        `json:"request_id" api:"required" format:"uuid"`
	Results   []WebWebCrawlMdResponseResult `json:"results" api:"required"`
	// Credits this request used and your remaining balance.
	KeyMetadata WebWebCrawlMdResponseKeyMetadata `json:"key_metadata"`
	// `true` when the deadline passed with `timeoutOpts.behavior` set to
	// `return-partial`. Only results finished in time are included. Partial results
	// are not cached.
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

// Cache status of this response.
type WebWebCrawlMdResponseCacheMetadata struct {
	// Age of the cached data, in milliseconds. `0` for `miss` and `zdr`.
	AgeMs int64 `json:"age_ms" api:"required"`
	// `hit` when served from cache, `miss` when fetched fresh, `zdr` when zero data
	// retention bypassed the cache.
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
	// Deepest link depth reached.
	MaxCrawlDepth int64 `json:"maxCrawlDepth" api:"required"`
	// Number of pages that failed.
	NumFailed int64 `json:"numFailed" api:"required"`
	// Number of URLs skipped: PDFs when `pdf.shouldParse` is `false`, URLs that do not
	// match `urlRegex`, and URLs that match `excludeUrlRegex`.
	NumSkipped int64 `json:"numSkipped" api:"required"`
	// Number of pages crawled successfully.
	NumSucceeded int64 `json:"numSucceeded" api:"required"`
	// Number of URLs crawled.
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
	// Page content as Markdown. Empty string when the page failed.
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
	// Link depth from the start URL. `0` is the start URL, `1` is one link away.
	CrawlDepth int64 `json:"crawlDepth" api:"required"`
	// URL that was scraped, after redirects. Same as `sourceUrl` when unknown.
	FinalURL string `json:"finalUrl" api:"required"`
	// URL that was requested, before redirects.
	SourceURL string `json:"sourceUrl" api:"required"`
	// HTTP status code of the page. `0` when the page could not be retrieved.
	StatusCode int64 `json:"statusCode" api:"required"`
	// Whether the page was fetched and parsed.
	Success bool `json:"success" api:"required"`
	// Page title. Empty string when unavailable.
	Title string `json:"title" api:"required"`
	// URL fetched for this page.
	URL string `json:"url" api:"required"`
	// Other non-social meta tags that don't have their own field.
	AdditionalMeta map[string]WebWebCrawlMdResponseResultMetadataAdditionalMetaUnion `json:"additionalMeta"`
	// Alternate links from `link rel="alternate"` tags, as absolute URLs.
	Alternates []WebWebCrawlMdResponseResultMetadataAlternate `json:"alternates"`
	// Author, when present.
	Author string `json:"author"`
	// Absolute canonical URL, when present.
	CanonicalURL string `json:"canonicalUrl"`
	// Best available description, from standard, Open Graph, or Twitter meta tags.
	Description string `json:"description"`
	// Why the page failed. Present only when `success` is `false`: `WEBSITE_BLOCKED`
	// (the site served an anti-bot challenge or access wall), `WEBSITE_ACCESS_ERROR`
	// (the page could not be retrieved or had no usable content), `NOT_FOUND`,
	// `UNSUPPORTED_CONTENT`, `NO_PUBLISHED_MARKDOWN`, or `INTERNAL_ERROR`.
	ErrorCode string `json:"error_code"`
	// Absolute favicon URL, when present.
	Favicon string `json:"favicon"`
	// Headings (`h1` to `h6`) in document order, from the full page before any
	// filtering. Up to 500. Omitted when the page has none.
	Headings []WebWebCrawlMdResponseResultMetadataHeading `json:"headings"`
	// Main preview image URL, from Open Graph, Twitter, or image meta tags.
	Image string `json:"image"`
	// JSON-LD structured data blocks parsed from the page.
	JsonLd []map[string]any `json:"jsonLd"`
	// Keywords from the page's `keywords` meta tag.
	Keywords []string `json:"keywords"`
	// Page language, from the `html` `lang` attribute or language meta tags.
	Language string `json:"language"`
	// Last-modified date or timestamp from page metadata, when present.
	ModifiedTime string `json:"modifiedTime"`
	// Open Graph tags, with the `og:` prefix removed and keys in camelCase.
	OpenGraph map[string]WebWebCrawlMdResponseResultMetadataOpenGraphUnion `json:"openGraph"`
	// Publish date or timestamp from page metadata, when present.
	PublishedTime string `json:"publishedTime"`
	// `robots` meta directive, when present.
	Robots string `json:"robots"`
	// Site or app name from page metadata.
	SiteName string `json:"siteName"`
	// Twitter card tags, with the `twitter:` prefix removed and keys in camelCase.
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
		ErrorCode      respjson.Field
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
	// Absolute URL of the alternate version.
	Href string `json:"href" api:"required"`
	// Language or locale of the alternate version, when present.
	Hreflang string `json:"hreflang"`
	// Title of the alternate version, when present.
	Title string `json:"title"`
	// MIME type of the alternate version, when present.
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
	// Heading level, from 1 (`h1`) to 6 (`h6`).
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

// Credits this request used and your remaining balance.
type WebWebCrawlMdResponseKeyMetadata struct {
	// Credits charged for this request.
	CreditsConsumed int64 `json:"credits_consumed" api:"required"`
	// Credits your organization has left after this request.
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
	// What to research, in plain language, up to 2000 characters. Include domains or
	// URLs to focus the research.
	Task string `json:"task" api:"required"`
	// Limit which sites the research can search, read, and cite.
	Controls WebAnswersParamsControls `json:"controls,omitzero"`
	// Example of the answer object to return, not a JSON Schema. Up to 8 levels deep,
	// 500 values, and 16,000 characters. Values that cannot be found come back as
	// `null`. Defaults to `{ "result": "" }`.
	JsonFormat map[string]any `json:"json_format,omitzero"`
	// `fast` prioritizes speed and adds extra verification for people and companies;
	// `ultra` searches and reads more sources for deeper research. Defaults to
	// `ultra`.
	//
	// Any of "fast", "ultra".
	Mode WebAnswersParamsMode `json:"mode,omitzero"`
	// Access policy IDs to apply; each URL must pass all of them. Requires access
	// policies enabled for your organization; otherwise returns
	// `ACCESS_POLICIES_NOT_ENABLED`.
	Policies []string `json:"policies,omitzero"`
	// Labels for this request. Filter usage and request logs by them in the dashboard
	// or with `GET /org/logs`.
	Tags []string `json:"tags,omitzero"`
	// Request deadline and what to return when it passes.
	TimeoutOpts WebAnswersParamsTimeoutOpts `json:"timeoutOpts,omitzero"`
	// `enabled` turns on zero data retention. Returns 403 `ZDR_NOT_ENABLED` unless
	// your organization has ZDR.
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

// Limit which sites the research can search, read, and cite.
type WebAnswersParamsControls struct {
	// Never search, read, or cite these sites, sections, or pages. Same format as
	// `sources`.
	ExcludeSources []string `json:"excludeSources,omitzero"`
	// Only search, read, and cite these sites, sections, or pages. Each entry is a URL
	// prefix: `example.com` covers the site and its subdomains, `example.com/docs`
	// covers that section, and `example.com/docs/setup.html` is one page.
	Sources []string `json:"sources,omitzero"`
	paramObj
}

func (r WebAnswersParamsControls) MarshalJSON() (data []byte, err error) {
	type shadow WebAnswersParamsControls
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebAnswersParamsControls) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `fast` prioritizes speed and adds extra verification for people and companies;
// `ultra` searches and reads more sources for deeper research. Defaults to
// `ultra`.
type WebAnswersParamsMode string

const (
	WebAnswersParamsModeFast  WebAnswersParamsMode = "fast"
	WebAnswersParamsModeUltra WebAnswersParamsMode = "ultra"
)

// Request deadline and what to return when it passes.
//
// The property Milliseconds is required.
type WebAnswersParamsTimeoutOpts struct {
	// Deadline in milliseconds.
	Milliseconds int64 `json:"milliseconds" api:"required"`
	// `fail` returns 408 at the deadline. `return-partial` returns the results
	// collected so far and marks the response as partial.
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

// `enabled` turns on zero data retention. Returns 403 `ZDR_NOT_ENABLED` unless
// your organization has ZDR.
type WebAnswersParamsZdr string

const (
	WebAnswersParamsZdrEnabled  WebAnswersParamsZdr = "enabled"
	WebAnswersParamsZdrDisabled WebAnswersParamsZdr = "disabled"
)

type WebExtractCompetitorsParams struct {
	// Company domain, such as `stripe.com`. Full URLs are reduced to their domain.
	Domain string `query:"domain" api:"required" json:"-"`
	// Exact number of direct competitors to return.
	NumCompetitors param.Opt[int64] `query:"numCompetitors,omitzero" json:"-"`
	// Comma-separated labels for this request, such as `production,team-alpha`. Filter
	// usage and request logs by them.
	Tags []string `query:"tags,omitzero" json:"-"`
	// Request deadline and what to return when it passes.
	TimeoutOpts WebExtractCompetitorsParamsTimeoutOpts `query:"timeoutOpts,omitzero" json:"-"`
	// `enabled` turns on zero data retention. Returns 403 `ZDR_NOT_ENABLED` unless
	// your organization has ZDR.
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

// Request deadline and what to return when it passes.
//
// The property Milliseconds is required.
type WebExtractCompetitorsParamsTimeoutOpts struct {
	// Deadline in milliseconds.
	Milliseconds int64 `query:"milliseconds" api:"required" json:"-"`
	// `fail` returns 408 at the deadline. `return-partial` returns the results
	// collected so far and marks the response as partial.
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

// `enabled` turns on zero data retention. Returns 403 `ZDR_NOT_ENABLED` unless
// your organization has ZDR.
type WebExtractCompetitorsParamsZdr string

const (
	WebExtractCompetitorsParamsZdrEnabled  WebExtractCompetitorsParamsZdr = "enabled"
	WebExtractCompetitorsParamsZdrDisabled WebExtractCompetitorsParamsZdr = "disabled"
)

type WebExtractStyleguideParams struct {
	// Maximum age of a cached result, in milliseconds. `0` always fetches fresh.
	// Defaults to 90 days (7776000000 ms). Values above 1 year (31536000000 ms) are
	// capped at 1 year.
	MaxAgeMs param.Opt[int64] `query:"maxAgeMs,omitzero" json:"-"`
	// Exact page URL to analyze. Provide either `domain` or `directUrl`, not both.
	DirectURL param.Opt[string] `query:"directUrl,omitzero" format:"uri" json:"-"`
	// Domain to analyze, such as `stripe.com`. Provide either `domain` or `directUrl`,
	// not both.
	Domain param.Opt[string] `query:"domain,omitzero" json:"-"`
	// Emulate a light or dark color scheme. Each scheme is cached separately.
	//
	// Any of "light", "dark".
	ColorScheme WebExtractStyleguideParamsColorScheme `query:"colorScheme,omitzero" json:"-"`
	// Comma-separated access policy IDs to apply; each URL must pass all of them.
	// Requires access policies enabled for your organization; otherwise returns
	// `ACCESS_POLICIES_NOT_ENABLED`.
	Policies []string `query:"policies,omitzero" json:"-"`
	// Comma-separated labels for this request, such as `production,team-alpha`. Filter
	// usage and request logs by them.
	Tags []string `query:"tags,omitzero" json:"-"`
	// Request deadline and what to return when it passes.
	TimeoutOpts WebExtractStyleguideParamsTimeoutOpts `query:"timeoutOpts,omitzero" json:"-"`
	// `enabled` turns off request logs and traces and keeps request and response
	// content out of usage logs. Brand data and hosted assets may still be stored.
	// Returns 403 `ZDR_NOT_ENABLED` unless your organization has ZDR.
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

// Emulate a light or dark color scheme. Each scheme is cached separately.
type WebExtractStyleguideParamsColorScheme string

const (
	WebExtractStyleguideParamsColorSchemeLight WebExtractStyleguideParamsColorScheme = "light"
	WebExtractStyleguideParamsColorSchemeDark  WebExtractStyleguideParamsColorScheme = "dark"
)

// Request deadline and what to return when it passes.
//
// The property Milliseconds is required.
type WebExtractStyleguideParamsTimeoutOpts struct {
	// Deadline in milliseconds.
	Milliseconds int64 `query:"milliseconds" api:"required" json:"-"`
	// `fail` returns 408 at the deadline. `return-partial` returns the results
	// collected so far and marks the response as partial. `return-partial` requires at
	// least 5000 ms. On `POST /web/scrape`, an output that is not ready in time fails
	// with `REQUEST_TIMEOUT` inside a 200 response under either behavior; 408 means
	// the whole request could not respond by the deadline.
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

// `enabled` turns off request logs and traces and keeps request and response
// content out of usage logs. Brand data and hosted assets may still be stored.
// Returns 403 `ZDR_NOT_ENABLED` unless your organization has ZDR.
type WebExtractStyleguideParamsZdr string

const (
	WebExtractStyleguideParamsZdrEnabled  WebExtractStyleguideParamsZdr = "enabled"
	WebExtractStyleguideParamsZdrDisabled WebExtractStyleguideParamsZdr = "disabled"
)

type WebMapURLsParams struct {
	// Domain to map, such as `stripe.com`.
	Domain string `query:"domain" api:"required" json:"-"`
	// Leave out URLs matching this regex, even when they match `urlRegex`.
	ExcludeURLRegex param.Opt[string] `query:"excludeUrlRegex,omitzero" json:"-"`
	// Include URLs on subdomains.
	IncludeSubdomains param.Opt[bool] `query:"includeSubdomains,omitzero" json:"-"`
	// Maximum number of URLs to return.
	MaxLinks param.Opt[int64] `query:"maxLinks,omitzero" json:"-"`
	// Filter URLs by a topic or phrase, most relevant first.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// Fetch this sitemap instead of discovering sitemaps. Must belong to the domain or
	// a subdomain.
	SitemapURL param.Opt[string] `query:"sitemapUrl,omitzero" format:"uri" json:"-"`
	// Only return URLs matching this regex. Only matches count toward `maxLinks`.
	URLRegex param.Opt[string] `query:"urlRegex,omitzero" json:"-"`
	// HTTP headers to send to the target site. Requests with headers skip the cache.
	Headers map[string]string `query:"headers,omitzero" json:"-"`
	// Comma-separated access policy IDs to apply; each URL must pass all of them.
	// Requires access policies enabled for your organization; otherwise returns
	// `ACCESS_POLICIES_NOT_ENABLED`.
	Policies []string `query:"policies,omitzero" json:"-"`
	// Comma-separated labels for this request, such as `production,team-alpha`. Filter
	// usage and request logs by them.
	Tags []string `query:"tags,omitzero" json:"-"`
	// Request deadline and what to return when it passes.
	TimeoutOpts WebMapURLsParamsTimeoutOpts `query:"timeoutOpts,omitzero" json:"-"`
	// `enabled` turns on zero data retention. Returns 403 `ZDR_NOT_ENABLED` unless
	// your organization has ZDR.
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

// Request deadline and what to return when it passes.
//
// The property Milliseconds is required.
type WebMapURLsParamsTimeoutOpts struct {
	// Deadline in milliseconds.
	Milliseconds int64 `query:"milliseconds" api:"required" json:"-"`
	// `fail` returns 408 at the deadline. `return-partial` returns the results
	// collected so far and marks the response as partial.
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

// `enabled` turns on zero data retention. Returns 403 `ZDR_NOT_ENABLED` unless
// your organization has ZDR.
type WebMapURLsParamsZdr string

const (
	WebMapURLsParamsZdrEnabled  WebMapURLsParamsZdr = "enabled"
	WebMapURLsParamsZdrDisabled WebMapURLsParamsZdr = "disabled"
)

type WebScrapeParams struct {
	// Outputs to return. Set at least one to `true`.
	Formats WebScrapeParamsFormats `json:"formats,omitzero" api:"required"`
	// Public HTTP or HTTPS URL to scrape.
	URL string `json:"url" api:"required" format:"uri"`
	// Maximum age of a cached result, in milliseconds. `0` always fetches fresh.
	// Defaults to 3 days (259200000 ms). Maximum: 1 year (31536000000 ms).
	MaxAgeMs param.Opt[int64] `json:"maxAgeMs,omitzero"`
	// Query for the `highlights` output. Required when `formats.highlights` is `true`,
	// and allowed only then.
	HighlightsParams WebScrapeParamsHighlightsParams `json:"highlightsParams,omitzero"`
	// Options for the `images` output. Requires `formats.images: true`.
	ImageParams WebScrapeParamsImageParams `json:"imageParams,omitzero"`
	// Schema and instructions for the `json` output. Required when `formats.json` is
	// `true`, and allowed only then.
	JsonParams WebScrapeParamsJsonParams `json:"jsonParams,omitzero"`
	// Options for the `markdown` output. Requires `formats.markdown: true`.
	MarkdownParams WebScrapeParamsMarkdownParams `json:"markdownParams,omitzero"`
	// Custom rules or a prebuilt parser for the `parsed` output. Required when
	// `formats.parse` is `true`, and allowed only then.
	ParseParams WebScrapeParamsParseParamsUnion `json:"parseParams,omitzero"`
	// Access policy IDs to apply; each URL must pass all of them. Requires access
	// policies enabled for your organization; otherwise returns
	// `ACCESS_POLICIES_NOT_ENABLED`.
	Policies []string `json:"policies,omitzero"`
	// Options for the `product` output. Requires `formats.product: true`.
	ProductParams WebScrapeParamsProductParams `json:"productParams,omitzero"`
	// Options for the `screenshot` output. Requires `formats.screenshot: true`.
	ScreenshotParams WebScrapeParamsScreenshotParams `json:"screenshotParams,omitzero"`
	// Browser and content settings shared by all outputs.
	SharedParams WebScrapeParamsSharedParams `json:"sharedParams,omitzero"`
	// Labels for this request. Filter usage and request logs by them in the dashboard
	// or with `GET /org/logs`.
	Tags []string `json:"tags,omitzero"`
	// Request deadline and what to return when it passes.
	TimeoutOpts WebScrapeParamsTimeoutOpts `json:"timeoutOpts,omitzero"`
	// `enabled` turns on zero data retention. Returns 403 `ZDR_NOT_ENABLED` unless
	// your organization has ZDR.
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

// Outputs to return. Set at least one to `true`.
type WebScrapeParamsFormats struct {
	// Return the original HTTP response body.
	Bytes param.Opt[bool] `json:"bytes,omitzero"`
	// Return Markdown excerpts relevant to `highlightsParams.query`.
	Highlights param.Opt[bool] `json:"highlights,omitzero"`
	// Return the rendered HTML.
	HTML param.Opt[bool] `json:"html,omitzero"`
	// Return the images found on the page.
	Images param.Opt[bool] `json:"images,omitzero"`
	// Return an object matching `jsonParams.schema`, extracted from the page text and
	// embedded data such as JSON-LD, meta tags, and framework state.
	Json param.Opt[bool] `json:"json,omitzero"`
	// Return the page content, or a video transcript, as Markdown.
	Markdown param.Opt[bool] `json:"markdown,omitzero"`
	// Return fields extracted with custom rules or a prebuilt parser, as `parsed`.
	Parse param.Opt[bool] `json:"parse,omitzero"`
	// Return product details such as name, price, and availability.
	Product param.Opt[bool] `json:"product,omitzero"`
	// Return a screenshot of the page.
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

// Query for the `highlights` output. Required when `formats.highlights` is `true`,
// and allowed only then.
//
// The property Query is required.
type WebScrapeParamsHighlightsParams struct {
	// Question or topic to find relevant passages for.
	Query string `json:"query" api:"required"`
	// Maximum combined length of the returned passages, in characters.
	MaxCharacters param.Opt[int64] `json:"maxCharacters,omitzero"`
	paramObj
}

func (r WebScrapeParamsHighlightsParams) MarshalJSON() (data []byte, err error) {
	type shadow WebScrapeParamsHighlightsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScrapeParamsHighlightsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Options for the `images` output. Requires `formats.images: true`.
type WebScrapeParamsImageParams struct {
	// Set to `visual` to drop visually duplicate images, keeping the largest copy.
	//
	// Any of "none", "visual".
	Dedupe string `json:"dedupe,omitzero"`
	// Extra data to add to each image: `dimensions` adds `width` and `height`,
	// `classification` adds `classification`, and `file` adds a hosted `fileUrl`.
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

// Schema and instructions for the `json` output. Required when `formats.json` is
// `true`, and allowed only then.
//
// The property Schema is required.
type WebScrapeParamsJsonParams struct {
	// JSON Schema for the object to extract, not an example object. Must describe a
	// top-level object, up to 50 KB. Make fields optional or nullable when the page
	// may not state them.
	Schema map[string]any `json:"schema,omitzero" api:"required"`
	// Extra guidance, such as which facts to prefer or how to read a field.
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

// Options for the `markdown` output. Requires `formats.markdown: true`.
type WebScrapeParamsMarkdownParams struct {
	// Include image references in the Markdown (`html-conversion` only).
	IncludeImages param.Opt[bool] `json:"includeImages,omitzero"`
	// Keep link URLs in the Markdown (`html-conversion` only).
	IncludeLinks param.Opt[bool] `json:"includeLinks,omitzero"`
	// How to write base64 images when `includeImages` is `true`: `placeholder`
	// shortens the data, `preserve` keeps it in full (`html-conversion` only).
	//
	// Any of "placeholder", "preserve".
	InlineImages string `json:"inlineImages,omitzero"`
	// Markdown sources in order of preference: `html-conversion` converts the page
	// HTML, `markdown-direct` uses Markdown the site publishes.
	//
	// Any of "html-conversion", "markdown-direct".
	SourcePreference []string `json:"sourcePreference,omitzero"`
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebScrapeParamsParseParamsUnion struct {
	OfCustomParser   *WebScrapeParamsParseParamsCustomParser   `json:",omitzero,inline"`
	OfPrebuiltParser *WebScrapeParamsParseParamsPrebuiltParser `json:",omitzero,inline"`
	paramUnion
}

func (u WebScrapeParamsParseParamsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCustomParser, u.OfPrebuiltParser)
}
func (u *WebScrapeParamsParseParamsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property Rules is required.
type WebScrapeParamsParseParamsCustomParser struct {
	// Output field names mapped to a CSS selector (such as `h1` or `a@href`) or a rule
	// object. Up to 100 fields, nested up to 5 levels.
	Rules map[string]WebScrapeParamsParseParamsCustomParserRuleUnion `json:"rules,omitzero" api:"required"`
	// Extract fields with custom CSS rules. The default when `type` is omitted.
	//
	// Any of "custom".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r WebScrapeParamsParseParamsCustomParser) MarshalJSON() (data []byte, err error) {
	type shadow WebScrapeParamsParseParamsCustomParser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScrapeParamsParseParamsCustomParser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebScrapeParamsParseParamsCustomParser](
		"type", "custom",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebScrapeParamsParseParamsCustomParserRuleUnion struct {
	OfString                                      param.Opt[string]                                 `json:",omitzero,inline"`
	OfWebScrapesParseParamsCustomParserRuleObject *WebScrapeParamsParseParamsCustomParserRuleObject `json:",omitzero,inline"`
	paramUnion
}

func (u WebScrapeParamsParseParamsCustomParserRuleUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfWebScrapesParseParamsCustomParserRuleObject)
}
func (u *WebScrapeParamsParseParamsCustomParserRuleUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property Selector is required.
type WebScrapeParamsParseParamsCustomParserRuleObject struct {
	// CSS selector to match within the current page or parent rule.
	Selector string `json:"selector" api:"required"`
	// What to return: `text`, `html`, an attribute such as `@href`, or nested field
	// rules. Defaults to `text`.
	Output string `json:"output,omitzero"`
	// Return the first match with `item` or all matches with `list`.
	//
	// Any of "item", "list".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r WebScrapeParamsParseParamsCustomParserRuleObject) MarshalJSON() (data []byte, err error) {
	type shadow WebScrapeParamsParseParamsCustomParserRuleObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScrapeParamsParseParamsCustomParserRuleObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebScrapeParamsParseParamsCustomParserRuleObject](
		"type", "item", "list",
	)
}

// The properties Parser, Type are required.
type WebScrapeParamsParseParamsPrebuiltParser struct {
	// Prebuilt parser to run. Each one accepts specific URLs and returns its result in
	// `parsed.data`. `emails`, `phone-numbers`, and `links` read the rendered page and
	// respect content filters; the others read their source directly, independently of
	// content filters and other outputs.
	//
	//   - `google-maps`: Google Maps place, search, and internal data URLs. Returns
	//     businesses in `data.results`; a search URL returns up to 200 places. Results
	//     use the country set in the URL, or the US when none is set, and come back with
	//     English labels.
	//   - `google-maps-reviews`: Google Maps place, search, and `?cid=` URLs, plus
	//     `search.google.com/local/reviews` URLs, so the `google_maps_url` and
	//     `google_reviews_url` fields of `google-maps` results work as input. Returns
	//     the reviews of the place in `data.reviews` (`rating`, `date`, `isoDate`,
	//     `snippet`, `translatedSnippet`, `likes`, `user`, `media`, `response`, `link`,
	//     `id`, `language`, guided ratings in `details`, and `user.isLocalGuide`), plus
	//     the place `featureId`, `cid`, and `sort` in `data`. `sort` and `limit` set the
	//     order and count; `oldest` is exact. An unknown place returns no reviews.
	//   - `google-play`: Play Store app details URLs
	//     (`play.google.com/store/apps/details?id=…`). Returns one app in `data`:
	//     details, developer contact, ratings, price, permissions, data safety, the
	//     storefront country the record describes, and `is_available` (whether the app
	//     is offered there). Ratings are counted per storefront, so `rating`,
	//     `rating_count`, and `rating_histogram` are `null` where nobody rated the app.
	//     The language and storefront country set in the URL apply (default English and
	//     the US); without a storefront country, an app not offered in the US is read
	//     from the country of its developer.
	//   - `google-play-reviews`: Same URLs as `google-play`. Returns up to the 200
	//     newest reviews in the language set in the URL, with developer replies, in
	//     `data.reviews`; an app with fewer such reviews returns them all.
	//   - `app-store`: App Store app URLs (`apps.apple.com/us/app/…/id…`). Returns one
	//     app in `data`: details, developer, ratings, price, version, and screenshots.
	//     The country in the path sets the storefront (default `us`).
	//   - `app-store-reviews`: Same URLs as `app-store`. Returns up to the 500 newest
	//     reviews in `data.reviews`.
	//   - `g2-reviews`: G2 product URLs (`g2.com/products/…`), following renamed
	//     products. Returns the 25 most recently published reviews in `data.reviews`
	//     (rating, title, likes, dislikes, `problems_solved`, reviewer role, company
	//     size, and industry), plus `data.product`. Keeps star-rating filters from the
	//     URL (`filters[nps_score][]=1` to `5`) and ignores other filters, sort orders,
	//     and pages.
	//   - `instagram-profile`: Public profile URLs (`instagram.com/{username}`). Returns
	//     one profile in `data`: bio, links, follower and post counts, category,
	//     business address, `similar_accounts`, and `latest_posts` (up to the 12 posts
	//     at the top of the profile, newest first). Instagram lists pinned posts first,
	//     so posts with `is_pinned: true` can be older. When Instagram shows only a
	//     rounded post count (such as 32K), `post_count` is that rounded number and
	//     `is_post_count_rounded` is `true`.
	//   - `instagram-post`: Post and reel URLs (`/p/…`, `/reel/…`). Returns one post in
	//     `data`: caption, hashtags, mentions, like, comment, and view counts
	//     (`like_count` and `view_count` are `null` when the owner hides them), media
	//     URLs, tagged users, collaborators, location, and about the top 14 comments in
	//     `data.comments`.
	//   - `instagram-feed`: Hashtag and keyword URLs (`/explore/tags/…`, `/popular/…`)
	//     return the top posts; location URLs (`/explore/locations/{id}/…`) return
	//     recent posts. Posts are in `data.posts`; keyword posts carry only the
	//     shortcode, owner, thumbnail, and views, plus the caption when Instagram shows
	//     it. All three Instagram parsers read public data without signing in, so
	//     private accounts return no posts and follower lists are unavailable.
	//   - `tiktok-profile`: Public profile URLs (`tiktok.com/@{username}`). Returns one
	//     profile in `data`: nickname, bio, bio link, follower, following, like, and
	//     video counts, verified, private, and business flags, and business category.
	//     Does not list videos.
	//   - `tiktok-video`: Video and photo post URLs (`/@{username}/video/{id}`,
	//     `/@{username}/photo/{id}`). Returns one post in `data`: caption, hashtags,
	//     mentions, author, view, like, comment, share, save, and repost counts,
	//     duration, thumbnail, a `video_url` that downloads without cookies,
	//     `image_urls` for photo posts, music, location, and ad, paid-partnership, and
	//     AI-generated flags.
	//   - `tiktok-feed`: Returns videos in the `tiktok-video` shape in `data.videos`:
	//     the latest 15 from a profile URL, the top 30 from a hashtag URL
	//     (`/tag/{name}`), or the top 12 from a search URL (`/search?q={query}`).
	//   - `tiktok-comments`: Video and photo post URLs. Returns the first 20 comments in
	//     `data.comments`: text, `posted_at`, like and reply counts, whether the author
	//     liked it, and the commenter.
	//   - `tiktok-followers`, `tiktok-following`: Profile URLs. Return up to 30 accounts
	//     in `data.users`, in the `tiktok-profile` shape.
	//   - `tiktok-user-search`: User search URLs (`/search/user?q={query}`). Returns the
	//     top 10 accounts in `data.users`: `id`, `username`, `nickname`, `biography`,
	//     `avatar_url`, `sec_uid`, `follower_count`, `like_count`, and `is_verified`.
	//     The TikTok parsers accept `tiktok.com`, `www.tiktok.com`, and `m.tiktok.com`
	//     URLs, not short links (`vm.tiktok.com`, `vt.tiktok.com`, `tiktok.com/t/…`).
	//   - `reddit-post`: Post URLs (`reddit.com/r/{subreddit}/comments/{id}`,
	//     `reddit.com/comments/{id}`, `redd.it/{id}`). Returns the post in `data.post`:
	//     title, body as Markdown, type (`text`, `link`, `image`, `gallery`, or
	//     `video`), `link_url`, subreddit, author, flairs, score, `upvote_ratio`,
	//     comment and crosspost counts, NSFW, spoiler, locked, stickied, and archived
	//     flags, images (every gallery image in order), and video (an MP4 `url` plus an
	//     `hls_url` that carries the audio). `data.comments` holds the comment tree;
	//     each comment has `id`, `author`, `body` (Markdown), `score`, `posted_at`,
	//     `depth`, `parent_id`, `url`, `is_submitter`, `is_deleted`, and its `replies`.
	//     `limit` caps the total number of comments (default 100), filled shallowest
	//     first so whole conversations are kept, and `data.has_more_comments` says
	//     whether the thread has more. The `sort` in the URL (`best`, `top`, `new`,
	//     `controversial`, `old`, `qa`) orders the comments.
	//   - `reddit-search`: Search URLs (`reddit.com/search?q={query}`, or
	//     `reddit.com/r/{subreddit}/search?q={query}` to search one subreddit). With
	//     `type=link` or no `type`, returns posts in the `reddit-post` shape in
	//     `data.posts`, ordered by the `sort` in the URL (`relevance`, `hot`, `top`,
	//     `new`, `comments`) within its `t` time range (`hour`, `day`, `week`, `month`,
	//     `year`, `all`). `type=sr` returns communities in `data.communities` (name,
	//     title, description, sidebar, `subscriber_count`, `created_at`, type, NSFW and
	//     quarantine flags, language, category, icon and banner URLs, and allowed post
	//     types). `type=user` returns accounts in `data.users` (username,
	//     `display_name`, bio, `created_at`, post and comment karma, avatar and banner
	//     URLs, and verified, employee, moderator, premium, and NSFW flags). `limit`
	//     caps the number of results (default 100), in the order Reddit returns them,
	//     and `data.has_more` says whether Reddit has more.
	//   - `reddit-subreddit`: Subreddit URLs (`reddit.com/r/{subreddit}`, optionally
	//     followed by `/hot`, `/new`, `/top`, `/rising`, or `/controversial`, with `t`
	//     setting the time range of `top` and `controversial`, default `day`). Returns
	//     `data.subreddit` in the community shape of `reddit-search`, `data.rules`
	//     (`name`, `description`, `applies_to`, `violation_reason`, `created_at`), and
	//     `data.posts` in the `reddit-post` shape, ordered like the tab (default `hot`)
	//     with stickied posts flagged. `r/all` and `r/popular` return posts with `null`
	//     `subreddit` and `rules`. `limit` caps the posts (default 100) and
	//     `data.has_more` says whether Reddit has more. Private and quarantined
	//     subreddits return `WEBSITE_ACCESS_ERROR`.
	//   - `reddit-user`: Profile URLs (`reddit.com/user/{username}` or
	//     `reddit.com/u/{username}`, optionally followed by `/overview`, `/submitted`,
	//     or `/comments`, with `sort` set to `new`, `hot`, `top`, or `controversial` and
	//     `t` setting the time range of `top` and `controversial`, default `all`).
	//     Returns `data.user` (`id`, `username`, `display_name`, bio, `created_at`,
	//     total, post, comment, awardee, and awarder karma, avatar and banner URLs, and
	//     verified, suspended, employee, moderator, premium, and NSFW flags),
	//     `data.trophies` (`name`, `description`, `icon_url`, `granted_at`, `url`),
	//     `data.moderated_communities` (`name`, `title`, `url`, `subscriber_count`,
	//     `icon_url`, `is_nsfw`), `data.posts` in the `reddit-post` shape, and
	//     `data.comments`, each with `id`, `author`, `body` (Markdown), `score`,
	//     `posted_at`, `edited_at`, `url`, `subreddit`, `parent_id` (`null` for a
	//     top-level comment), `is_submitter`, `is_stickied`, `distinguished`,
	//     `award_count`, and the post it was left on (`id`, `title`, `url`, `author`,
	//     `is_nsfw`). The overview tab returns both posts and comments; `submitted`
	//     returns `null` comments and `comments` returns `null` posts. `limit` caps
	//     posts and comments together (default 100) and `data.has_more` says whether
	//     Reddit has more. Users who hide their activity return empty lists, and
	//     suspended accounts return their profile with `null` posts and comments.
	//   - `yelp-search`: Yelp search URLs with `find_desc` and `find_loc`; `start` sets
	//     the result offset. Returns up to 20 ranked organic businesses in
	//     `data.businesses`, with `total` and `pagination`.
	//   - `yelp-reviews`: Yelp business URLs (`yelp.com/biz/{alias}`); optional `start`
	//     sets the offset. Returns up to 20 of the latest public reviews, always newest
	//     first, in `data.reviews`, with business details and `pagination.next_offset`.
	//     Each review has rating, full text, `updated_at`, and author (`id` from the
	//     reviewer profile URL, `url` linking to that profile, `name`, `location`, and
	//     `review_count`); owner replies and photos are not included. Unsupported sort
	//     and review filters are rejected; tracking parameters are ignored.
	//   - `tripadvisor-reviews`: Tripadvisor hotel, restaurant, and attraction review
	//     URLs (`tripadvisor.com/Hotel_Review-g{geo}-d{id}-Reviews-….html`, on any
	//     Tripadvisor country domain); `-or{N}-` in the path sets the offset. Returns 20
	//     reviews in `data.reviews` in the order Tripadvisor shows them, with
	//     `data.location` (`id`, `name`, `type`, `url`, `review_count`) and
	//     `pagination.next_offset`. Each review has `id`, `url`, `title`, full text,
	//     `rating`, `language`, `published_at`, `visit_month`, `helpful_votes`,
	//     `subratings`, `photo_urls`, `owner_response` (`text`, `published_at`,
	//     `author_name`), and `author` (`id`, `username`, `name`, `url`,
	//     `contribution_count`). An unknown location returns no reviews.
	//   - `car-rental-availability`: avis.com and budget.com vehicle-availability search
	//     URLs. Returns every vehicle with pay-now and pay-later prices in
	//     `data.vehicles`.
	//   - `car-rental-locations`: avis.com and budget.com `/locations?keyword=` URLs.
	//     Returns rental branches with codes, addresses, coordinates, and hours in
	//     `data.locations`.
	//   - `emails`: Any page. Returns its email addresses in `data.emails`
	//     (`[{ email }]`), read from `mailto:` links, Cloudflare-protected addresses,
	//     and visible text.
	//   - `phone-numbers`: Any page. Returns its valid phone numbers in E.164 format in
	//     `data.phone_numbers` (`[{ phone }]`), read from `tel:` links and visible text.
	//     Numbers without a country code are read for the country of the site TLD (US by
	//     default).
	//   - `links`: Any page. Returns its unique HTTP and HTTPS links as absolute URLs in
	//     `data.links` (`[{ url, text }]`). `text` falls back to `aria-label` or image
	//     alt text and is `null` when a link has none.
	//
	// Any of "google-maps", "google-maps-reviews", "google-play",
	// "google-play-reviews", "app-store", "app-store-reviews", "g2-reviews",
	// "instagram-profile", "instagram-post", "instagram-feed", "tiktok-profile",
	// "tiktok-video", "tiktok-feed", "tiktok-comments", "tiktok-followers",
	// "tiktok-following", "tiktok-user-search", "reddit-post", "reddit-search",
	// "reddit-subreddit", "reddit-user", "yelp-search", "yelp-reviews",
	// "tripadvisor-reviews", "car-rental-availability", "car-rental-locations",
	// "emails", "phone-numbers", "links".
	Parser string `json:"parser,omitzero" api:"required"`
	// Maximum number of reviews (`google-maps-reviews`, default 100), comments
	// (`reddit-post`, default 100), or results (`reddit-search`, `reddit-subreddit`,
	// and `reddit-user`, default 100) to return. Only these parsers accept it.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// Order of `data.reviews` for `google-maps-reviews`. `newest`, `relevant`,
	// `highest`, and `lowest` use the same orders as Google Maps; `oldest` returns the
	// earliest reviews first.
	//
	// Any of "newest", "oldest", "relevant", "highest", "lowest".
	Sort string `json:"sort,omitzero"`
	// Use a prebuilt parser instead of custom rules.
	//
	// This field can be elided, and will marshal its zero value as "prebuilt".
	Type constant.Prebuilt `json:"type" default:"prebuilt"`
	paramObj
}

func (r WebScrapeParamsParseParamsPrebuiltParser) MarshalJSON() (data []byte, err error) {
	type shadow WebScrapeParamsParseParamsPrebuiltParser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScrapeParamsParseParamsPrebuiltParser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebScrapeParamsParseParamsPrebuiltParser](
		"parser", "google-maps", "google-maps-reviews", "google-play", "google-play-reviews", "app-store", "app-store-reviews", "g2-reviews", "instagram-profile", "instagram-post", "instagram-feed", "tiktok-profile", "tiktok-video", "tiktok-feed", "tiktok-comments", "tiktok-followers", "tiktok-following", "tiktok-user-search", "reddit-post", "reddit-search", "reddit-subreddit", "reddit-user", "yelp-search", "yelp-reviews", "tripadvisor-reviews", "car-rental-availability", "car-rental-locations", "emails", "phone-numbers", "links",
	)
	apijson.RegisterFieldValidator[WebScrapeParamsParseParamsPrebuiltParser](
		"sort", "newest", "oldest", "relevant", "highest", "lowest",
	)
}

// Options for the `product` output. Requires `formats.product: true`.
type WebScrapeParamsProductParams struct {
	// Drop visually duplicate product images, keeping the largest copy.
	DedupeImages param.Opt[bool] `json:"dedupeImages,omitzero"`
	// Ignored. The AI fallback runs automatically.
	//
	// Deprecated: deprecated
	UseAIFallback param.Opt[bool] `json:"useAIFallback,omitzero"`
	paramObj
}

func (r WebScrapeParamsProductParams) MarshalJSON() (data []byte, err error) {
	type shadow WebScrapeParamsProductParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScrapeParamsProductParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Options for the `screenshot` output. Requires `formats.screenshot: true`.
type WebScrapeParamsScreenshotParams struct {
	// What to capture: `viewport`, `fullPage`, one element, or a rectangle. Up to 40
	// megapixels.
	Area WebScrapeParamsScreenshotParamsAreaUnion `json:"area,omitzero"`
	// Image format for the screenshot.
	//
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
	// CSS selector matching exactly one visible element.
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

// Rectangle to capture, in pixels from the document origin.
//
// The properties Height, Width, X, Y are required.
type WebScrapeParamsScreenshotParamsAreaRectangle struct {
	// Height of the capture in pixels.
	Height int64 `json:"height" api:"required"`
	// Width of the capture in pixels.
	Width int64 `json:"width" api:"required"`
	// Left edge of the capture, in pixels from the document origin.
	X int64 `json:"x" api:"required"`
	// Top edge of the capture, in pixels from the document origin.
	Y int64 `json:"y" api:"required"`
	paramObj
}

func (r WebScrapeParamsScreenshotParamsAreaRectangle) MarshalJSON() (data []byte, err error) {
	type shadow WebScrapeParamsScreenshotParamsAreaRectangle
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScrapeParamsScreenshotParamsAreaRectangle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Browser and content settings shared by all outputs.
type WebScrapeParamsSharedParams struct {
	// Country to fetch from, as an ISO 3166-1 alpha-2 code such as `US`.
	// Case-insensitive.
	Country param.Opt[string] `json:"country,omitzero"`
	// Accept cookie banners before actions and capture.
	DismissCookies param.Opt[bool] `json:"dismissCookies,omitzero"`
	// Close other popups before actions and capture.
	DismissPopups param.Opt[bool] `json:"dismissPopups,omitzero"`
	// Keep only the main content, without navigation, headers, footers, or sidebars.
	// Does not affect `screenshot`, `bytes`, or `product`.
	MainContentOnly param.Opt[bool] `json:"mainContentOnly,omitzero"`
	// Wait for CSS animations to finish before capture. Defaults to `true` when
	// `formats.screenshot` is `true`.
	SettleAnimations param.Opt[bool] `json:"settleAnimations,omitzero"`
	// Browser steps to run in order before capture. Requires a paid plan. Skips the
	// cache.
	Actions []WebScrapeParamsSharedParamsActionUnion `json:"actions,omitzero"`
	// Remove elements matching these CSS selectors. Overrides `includeSelectors`.
	ExcludeSelectors []string `json:"excludeSelectors,omitzero"`
	// HTTP headers to send to the target site. Requests with headers skip the cache.
	Headers map[string]string `json:"headers,omitzero"`
	// Which iframes to inline in HTML and text outputs. Screenshots always show
	// visible frames. `auto` inlines only iframes that hold page content, such as
	// same-site documents, framesets, job boards, forms, and embedded documents, and
	// skips ads and widgets. `true` inlines every iframe; `false` inlines none.
	IncludeFrames WebScrapeParamsSharedParamsIncludeFramesUnion `json:"includeFrames,omitzero"`
	// Keep only elements matching these CSS selectors.
	IncludeSelectors []string `json:"includeSelectors,omitzero"`
	// Options for parsing documents such as PDFs.
	Parsers WebScrapeParamsSharedParamsParsers `json:"parsers,omitzero"`
	// Scrape as a signed-in user. Create a profile with `type: "new"` and your
	// credentials, then reuse it with `type: "existing"`. Requires a paid plan and an
	// HTTPS URL. Profiles keep their session between requests, skip the cache, and
	// cannot be combined with ZDR, custom headers, `formats.bytes`, or prebuilt
	// parsers other than `emails`, `phone-numbers`, and `links`.
	Profile WebScrapeParamsSharedParamsProfileUnion `json:"profile,omitzero"`
	// Emulate a light or dark color scheme.
	//
	// Any of "light", "dark".
	Theme string `json:"theme,omitzero"`
	// Browser viewport in pixels. Defaults to 1920 × 1080. If you pass `viewport`, a
	// missing `width` defaults to 1440 and a missing `height` to 900.
	Viewport WebScrapeParamsSharedParamsViewport `json:"viewport,omitzero"`
	// Time to wait after actions, in milliseconds, or a CSS selector to wait for.
	// Defaults to 500 ms, or 2000 ms when `includeFrames` is `true` or the URL ends in
	// `.xml`.
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
	// One browser instruction, such as clicking a button or entering text.
	Action string `json:"action" api:"required"`
	// Follow a plain-language instruction.
	//
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
	// Maximum number of scroll steps for this action.
	MaxScrolls param.Opt[int64] `json:"maxScrolls,omitzero"`
	// CSS selector of the container to scroll. Omit to scroll the page.
	Selector param.Opt[string] `json:"selector,omitzero"`
	// Distance per scroll step: a number of pixels, one `viewport`, or `max` to reach
	// the end.
	Amount WebScrapeParamsSharedParamsActionScrollAmountUnion `json:"amount,omitzero"`
	// Direction to scroll.
	//
	// Any of "down", "up", "left", "right".
	Direction string `json:"direction,omitzero"`
	// Scroll the page or a container.
	//
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
	// How long to pause before the next action, in milliseconds.
	Milliseconds int64 `json:"milliseconds" api:"required"`
	// Pause for a fixed time.
	//
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
	// CSS selector to wait for before continuing.
	Selector string `json:"selector" api:"required"`
	// Wait for an element matching `selector`.
	//
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebScrapeParamsSharedParamsIncludeFramesUnion struct {
	OfBool param.Opt[bool] `json:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWebScrapesSharedParamsIncludeFramesString)
	OfWebScrapesSharedParamsIncludeFramesString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u WebScrapeParamsSharedParamsIncludeFramesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfWebScrapesSharedParamsIncludeFramesString)
}
func (u *WebScrapeParamsSharedParamsIncludeFramesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type WebScrapeParamsSharedParamsIncludeFramesString string

const (
	WebScrapeParamsSharedParamsIncludeFramesStringAuto WebScrapeParamsSharedParamsIncludeFramesString = "auto"
)

// Options for parsing documents such as PDFs.
type WebScrapeParamsSharedParamsParsers struct {
	// Page range, OCR, and page markers for PDFs.
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

// Page range, OCR, and page markers for PDFs.
type WebScrapeParamsSharedParamsParsersPdf struct {
	// Last page to parse. Defaults to the last page. Must be at least `startPage`.
	EndPage param.Opt[int64] `json:"endPage,omitzero"`
	// Start each PDF page in the Markdown with a `<!-- page N -->` line, where `N` is
	// the original 1-based page number. These responses are never served from cache.
	IncludeNumberedPages param.Opt[bool] `json:"includeNumberedPages,omitzero"`
	// First page to parse, starting at 1.
	StartPage param.Opt[int64] `json:"startPage,omitzero"`
	// Set to `auto` to read scanned pages with OCR. Adds 1 credit per page read.
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebScrapeParamsSharedParamsProfileUnion struct {
	OfExisting *WebScrapeParamsSharedParamsProfileExisting `json:",omitzero,inline"`
	OfNew      *WebScrapeParamsSharedParamsProfileNew      `json:",omitzero,inline"`
	paramUnion
}

func (u WebScrapeParamsSharedParamsProfileUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExisting, u.OfNew)
}
func (u *WebScrapeParamsSharedParamsProfileUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[WebScrapeParamsSharedParamsProfileUnion](
		"type",
		apijson.Discriminator[WebScrapeParamsSharedParamsProfileExisting]("existing"),
		apijson.Discriminator[WebScrapeParamsSharedParamsProfileNew]("new"),
	)
}

// The properties ProfileID, Type are required.
type WebScrapeParamsSharedParamsProfileExisting struct {
	// ID of the browser profile, from `profile.id` in an earlier scrape response.
	ProfileID string `json:"profileId" api:"required"`
	// Reuse a saved browser profile.
	//
	// This field can be elided, and will marshal its zero value as "existing".
	Type constant.Existing `json:"type" default:"existing"`
	paramObj
}

func (r WebScrapeParamsSharedParamsProfileExisting) MarshalJSON() (data []byte, err error) {
	type shadow WebScrapeParamsSharedParamsProfileExisting
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScrapeParamsSharedParamsProfileExisting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Identifier, Password, Type are required.
type WebScrapeParamsSharedParamsProfileNew struct {
	// Username or email address for the website.
	Identifier string `json:"identifier" api:"required"`
	// Password for the website. Used to sign in and never stored.
	Password string `json:"password" api:"required"`
	// Label for the profile, shown by `GET /org/browser/profiles`.
	Description param.Opt[string] `json:"description,omitzero"`
	// Sign in and save a new browser profile.
	//
	// This field can be elided, and will marshal its zero value as "new".
	Type constant.New `json:"type" default:"new"`
	paramObj
}

func (r WebScrapeParamsSharedParamsProfileNew) MarshalJSON() (data []byte, err error) {
	type shadow WebScrapeParamsSharedParamsProfileNew
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScrapeParamsSharedParamsProfileNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Browser viewport in pixels. Defaults to 1920 × 1080. If you pass `viewport`, a
// missing `width` defaults to 1440 and a missing `height` to 900.
type WebScrapeParamsSharedParamsViewport struct {
	// Browser viewport height in pixels.
	Height param.Opt[int64] `json:"height,omitzero"`
	// Browser viewport width in pixels.
	Width param.Opt[int64] `json:"width,omitzero"`
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

// Request deadline and what to return when it passes.
//
// The property Milliseconds is required.
type WebScrapeParamsTimeoutOpts struct {
	// Deadline in milliseconds.
	Milliseconds int64 `json:"milliseconds" api:"required"`
	// `fail` returns 408 at the deadline. `return-partial` returns the results
	// collected so far and marks the response as partial. `return-partial` requires at
	// least 5000 ms. On `POST /web/scrape`, an output that is not ready in time fails
	// with `REQUEST_TIMEOUT` inside a 200 response under either behavior; 408 means
	// the whole request could not respond by the deadline.
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

// `enabled` turns on zero data retention. Returns 403 `ZDR_NOT_ENABLED` unless
// your organization has ZDR.
type WebScrapeParamsZdr string

const (
	WebScrapeParamsZdrEnabled  WebScrapeParamsZdr = "enabled"
	WebScrapeParamsZdrDisabled WebScrapeParamsZdr = "disabled"
)

type WebScreenshotParams struct {
	// Maximum age of a cached result, in milliseconds. `0` always fetches fresh.
	// Defaults to 1 day (86400000 ms). Maximum: 30 days (2592000000 ms).
	MaxAgeMs param.Opt[int64] `query:"maxAgeMs,omitzero" json:"-"`
	// Return the viewport-sized slice of the full page that starts this many pixels
	// from the top. Step by the viewport height, such as `0`, `1080`, `2160`, to walk
	// a long page. The last slice can be shorter. Overrides `fullScreenshot`.
	ScrollOffset param.Opt[int64] `query:"scrollOffset,omitzero" json:"-"`
	// Time to wait after the page loads before capturing, in milliseconds. `0` waits
	// 500 ms. `timeoutOpts.milliseconds` must be at least `waitForMs` plus 10000 ms,
	// or the request fails with `TIMEOUT_TOO_SHORT_FOR_WAIT`.
	WaitForMs param.Opt[int64] `query:"waitForMs,omitzero" json:"-"`
	// Dismiss cookie banners and close other popups and overlays before capture.
	ClearPopups param.Opt[bool] `query:"clearPopups,omitzero" json:"-"`
	// Exact URL to capture, such as `https://example.com/pricing`. Provide either
	// `domain` or `directUrl`, not both.
	DirectURL param.Opt[string] `query:"directUrl,omitzero" format:"uri" json:"-"`
	// Domain to capture, such as `stripe.com`. Provide either `domain` or `directUrl`,
	// not both.
	Domain param.Opt[string] `query:"domain,omitzero" json:"-"`
	// Dismiss cookie consent banners before capture.
	HandleCookiePopup param.Opt[bool] `query:"handleCookiePopup,omitzero" json:"-"`
	// Emulate a light or dark color scheme.
	//
	// Any of "light", "dark".
	ColorScheme WebScreenshotParamsColorScheme `query:"colorScheme,omitzero" json:"-"`
	// Country to fetch from, as an ISO 3166-1 alpha-2 code.
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
	// Capture the full scrollable page instead of only the viewport.
	//
	// Any of "true", "false".
	FullScreenshot WebScreenshotParamsFullScreenshot `query:"fullScreenshot,omitzero" json:"-"`
	// HTTP headers to send to the target origin, such as
	// `headers[Authorization]=Bearer token`. Requests with headers skip the cache and
	// return the image as a data URL.
	Headers map[string]string `query:"headers,omitzero" json:"-"`
	// Capture this type of page instead of the home page. Context.dev finds it among
	// the domain's links, in 30 languages. Only with `domain`.
	//
	// Any of "login", "signup", "blog", "careers", "pricing", "terms", "privacy",
	// "contact".
	Page WebScreenshotParamsPage `query:"page,omitzero" json:"-"`
	// Comma-separated labels for this request, such as `production,team-alpha`. Filter
	// usage and request logs by them.
	Tags []string `query:"tags,omitzero" json:"-"`
	// Request deadline and what to return when it passes.
	TimeoutOpts WebScreenshotParamsTimeoutOpts `query:"timeoutOpts,omitzero" json:"-"`
	// Browser viewport size in pixels.
	Viewport WebScreenshotParamsViewport `query:"viewport,omitzero" json:"-"`
	// `enabled` turns on zero data retention. Returns 403 `ZDR_NOT_ENABLED` unless
	// your organization has ZDR.
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

// Emulate a light or dark color scheme.
type WebScreenshotParamsColorScheme string

const (
	WebScreenshotParamsColorSchemeLight WebScreenshotParamsColorScheme = "light"
	WebScreenshotParamsColorSchemeDark  WebScreenshotParamsColorScheme = "dark"
)

// Country to fetch from, as an ISO 3166-1 alpha-2 code.
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

// Capture the full scrollable page instead of only the viewport.
type WebScreenshotParamsFullScreenshot string

const (
	WebScreenshotParamsFullScreenshotTrue  WebScreenshotParamsFullScreenshot = "true"
	WebScreenshotParamsFullScreenshotFalse WebScreenshotParamsFullScreenshot = "false"
)

// Capture this type of page instead of the home page. Context.dev finds it among
// the domain's links, in 30 languages. Only with `domain`.
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

// Request deadline and what to return when it passes.
//
// The property Milliseconds is required.
type WebScreenshotParamsTimeoutOpts struct {
	// Deadline in milliseconds.
	Milliseconds int64 `query:"milliseconds" api:"required" json:"-"`
	// `fail` returns 408 at the deadline. `return-partial` returns the results
	// collected so far and marks the response as partial. `return-partial` requires at
	// least 5000 ms. On `POST /web/scrape`, an output that is not ready in time fails
	// with `REQUEST_TIMEOUT` inside a 200 response under either behavior; 408 means
	// the whole request could not respond by the deadline.
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

// Browser viewport size in pixels.
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

// `enabled` turns on zero data retention. Returns 403 `ZDR_NOT_ENABLED` unless
// your organization has ZDR.
type WebScreenshotParamsZdr string

const (
	WebScreenshotParamsZdrEnabled  WebScreenshotParamsZdr = "enabled"
	WebScreenshotParamsZdrDisabled WebScreenshotParamsZdr = "disabled"
)

type WebSearchParams struct {
	// Search query. Accepts natural language and search operators such as `site:`,
	// `-site:`, `inurl:`, `intitle:`, quoted phrases, and `OR`.
	Query string `json:"query" api:"required"`
	// Maximum length of each result's `description`, in characters.
	DescriptionMaxCharacters param.Opt[int64] `json:"descriptionMaxCharacters,omitzero"`
	// Number of results to return.
	NumResults param.Opt[int64] `json:"numResults,omitzero"`
	// Return passages from each result page that are relevant to the query. Pages are
	// read with the `markdownOptions` settings.
	HighlightsOptions WebSearchParamsHighlightsOptions `json:"highlightsOptions,omitzero"`
	// Scrape each result page to Markdown. Set `enabled` to `true` to turn it on.
	MarkdownOptions WebSearchParamsMarkdownOptions `json:"markdownOptions,omitzero"`
	// Localize results to this country, as an ISO 3166-1 alpha-2 code such as `us` or
	// `de`.
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
	// Leave out results from these domains, such as `pinterest.com`.
	ExcludeDomains []string `json:"excludeDomains,omitzero"`
	// Restrict results to content published within this window.
	//
	// Any of "last_24_hours", "last_week", "last_month", "last_year".
	Freshness WebSearchParamsFreshness `json:"freshness,omitzero"`
	// Only return results from these domains, such as `arxiv.org`.
	IncludeDomains []string `json:"includeDomains,omitzero"`
	// Languages to return results in, as BCP 47 codes such as `en` or `de`, or `all`
	// for any language.
	Languages WebSearchParamsLanguagesUnion `json:"languages,omitzero"`
	// Access policy IDs to apply; each URL must pass all of them. Requires access
	// policies enabled for your organization; otherwise returns
	// `ACCESS_POLICIES_NOT_ENABLED`.
	Policies []string `json:"policies,omitzero"`
	// Labels for this request. Filter usage and request logs by them in the dashboard
	// or with `GET /org/logs`.
	Tags []string `json:"tags,omitzero"`
	// Request deadline and what to return when it passes.
	TimeoutOpts WebSearchParamsTimeoutOpts `json:"timeoutOpts,omitzero"`
	// `enabled` turns on zero data retention. Returns 403 `ZDR_NOT_ENABLED` unless
	// your organization has ZDR.
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

// Localize results to this country, as an ISO 3166-1 alpha-2 code such as `us` or
// `de`.
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

// Return passages from each result page that are relevant to the query. Pages are
// read with the `markdownOptions` settings.
type WebSearchParamsHighlightsOptions struct {
	// Return relevant passages for each result. Adds 1 credit per 10 results.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Maximum combined length of passages per result, in characters.
	MaxCharacters param.Opt[int64] `json:"maxCharacters,omitzero"`
	paramObj
}

func (r WebSearchParamsHighlightsOptions) MarshalJSON() (data []byte, err error) {
	type shadow WebSearchParamsHighlightsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebSearchParamsHighlightsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebSearchParamsLanguagesUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWebSearchsLanguagesString)
	OfWebSearchsLanguagesString         param.Opt[string] `json:",omitzero,inline"`
	OfWebSearchsLanguagesArrayItemArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u WebSearchParamsLanguagesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfWebSearchsLanguagesString, u.OfWebSearchsLanguagesArrayItemArray)
}
func (u *WebSearchParamsLanguagesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type WebSearchParamsLanguagesString string

const (
	WebSearchParamsLanguagesStringAll WebSearchParamsLanguagesString = "all"
)

// Scrape each result page to Markdown. Set `enabled` to `true` to turn it on.
type WebSearchParamsMarkdownOptions struct {
	// Maximum age of a cached result, in milliseconds. `0` always fetches fresh.
	// Defaults to 180 days (15552000000 ms) with Markdown, or 365 days (31536000000
	// ms) with highlights only.
	MaxAgeMs param.Opt[int64] `json:"maxAgeMs,omitzero"`
	// Time to wait after each result page loads before reading it, in milliseconds.
	// `0` waits 500 ms. Helps with JavaScript-heavy pages.
	WaitForMs param.Opt[int64] `json:"waitForMs,omitzero"`
	// Scrape each result page to Markdown. Adds 1 credit per result that returns
	// Markdown.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Include image references in the Markdown.
	IncludeImages param.Opt[bool] `json:"includeImages,omitzero"`
	// Keep link URLs in the Markdown.
	IncludeLinks param.Opt[bool] `json:"includeLinks,omitzero"`
	// Shorten inline base64 image data.
	ShortenBase64Images param.Opt[bool] `json:"shortenBase64Images,omitzero"`
	// Keep only the main content, without navigation, headers, footers, or sidebars.
	UseMainContentOnly param.Opt[bool] `json:"useMainContentOnly,omitzero"`
	// Which iframe contents to include in the Markdown. `auto` inlines only iframes
	// that hold page content, such as same-site documents, framesets, job boards,
	// forms, and embedded documents, and skips ads and widgets. `true` inlines every
	// iframe; `false` inlines none.
	IncludeFrames WebSearchParamsMarkdownOptionsIncludeFramesUnion `json:"includeFrames,omitzero"`
	// Whether to parse PDFs and which pages to read.
	Pdf WebSearchParamsMarkdownOptionsPdf `json:"pdf,omitzero"`
	// Request deadline and what to return when it passes.
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebSearchParamsMarkdownOptionsIncludeFramesUnion struct {
	OfBool param.Opt[bool] `json:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWebSearchsMarkdownOptionsIncludeFramesString)
	OfWebSearchsMarkdownOptionsIncludeFramesString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u WebSearchParamsMarkdownOptionsIncludeFramesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfWebSearchsMarkdownOptionsIncludeFramesString)
}
func (u *WebSearchParamsMarkdownOptionsIncludeFramesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type WebSearchParamsMarkdownOptionsIncludeFramesString string

const (
	WebSearchParamsMarkdownOptionsIncludeFramesStringAuto WebSearchParamsMarkdownOptionsIncludeFramesString = "auto"
)

// Whether to parse PDFs and which pages to read.
type WebSearchParamsMarkdownOptionsPdf struct {
	// Last page to parse. Defaults to the last page. Must be at least `start`.
	End param.Opt[int64] `json:"end,omitzero"`
	// Parse PDF results. When `false`, PDF results fail with `WEBSITE_ACCESS_ERROR`.
	ShouldParse param.Opt[bool] `json:"shouldParse,omitzero"`
	// First page to parse, starting at 1.
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

// Request deadline and what to return when it passes.
//
// The property Milliseconds is required.
type WebSearchParamsMarkdownOptionsTimeoutOpts struct {
	// Deadline in milliseconds.
	Milliseconds int64 `json:"milliseconds" api:"required"`
	// `fail` returns 408 at the deadline. `return-partial` returns the results
	// collected so far and marks the response as partial. `return-partial` requires at
	// least 5000 ms. On `POST /web/scrape`, an output that is not ready in time fails
	// with `REQUEST_TIMEOUT` inside a 200 response under either behavior; 408 means
	// the whole request could not respond by the deadline.
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

// Request deadline and what to return when it passes.
//
// The property Milliseconds is required.
type WebSearchParamsTimeoutOpts struct {
	// Deadline in milliseconds.
	Milliseconds int64 `json:"milliseconds" api:"required"`
	// `fail` returns 408 at the deadline. `return-partial` returns the results
	// collected so far and marks the response as partial.
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

// `enabled` turns on zero data retention. Returns 403 `ZDR_NOT_ENABLED` unless
// your organization has ZDR.
type WebSearchParamsZdr string

const (
	WebSearchParamsZdrEnabled  WebSearchParamsZdr = "enabled"
	WebSearchParamsZdrDisabled WebSearchParamsZdr = "disabled"
)

type WebWebCrawlMdParams struct {
	// Start URL, including `http://` or `https://`.
	URL string `json:"url" api:"required" format:"uri"`
	// Maximum age of a cached result, in milliseconds. `0` always fetches fresh.
	// Defaults to 1 day (86400000 ms). Maximum: 30 days (2592000000 ms).
	MaxAgeMs param.Opt[int64] `json:"maxAgeMs,omitzero"`
	// Time to wait after each page loads before reading it, in milliseconds. `0` waits
	// 500 ms. `timeoutOpts.milliseconds` must be at least `waitForMs` plus 10000 ms,
	// or the request fails with `TIMEOUT_TOO_SHORT_FOR_WAIT`.
	WaitForMs param.Opt[int64] `json:"waitForMs,omitzero"`
	// Never follow or scrape URLs matching this regex, even when they match
	// `urlRegex`. The start URL is always scraped.
	ExcludeURLRegex param.Opt[string] `json:"excludeUrlRegex,omitzero"`
	// Follow links to subdomains of the start URL's domain, such as `docs.example.com`
	// from `example.com`. `www` and the apex domain always count as the same site.
	FollowSubdomains param.Opt[bool] `json:"followSubdomains,omitzero"`
	// Include image references in the Markdown (`html-conversion` only).
	IncludeImages param.Opt[bool] `json:"includeImages,omitzero"`
	// Keep link URLs in the Markdown (`html-conversion` only).
	IncludeLinks param.Opt[bool] `json:"includeLinks,omitzero"`
	// Maximum link depth from the start URL. `0` crawls only the start URL. Omit for
	// no depth limit.
	MaxDepth param.Opt[int64] `json:"maxDepth,omitzero"`
	// Maximum number of pages to crawl.
	MaxPages param.Opt[int64] `json:"maxPages,omitzero"`
	// Wait briefly for CSS animations and transitions to settle before reading each
	// page.
	SettleAnimations param.Opt[bool] `json:"settleAnimations,omitzero"`
	// Shorten inline base64 image data (`html-conversion` only).
	ShortenBase64Images param.Opt[bool] `json:"shortenBase64Images,omitzero"`
	// Soft time budget for the crawl, in milliseconds. When it runs out, the crawl
	// returns the pages collected so far. Larger values are capped at 110 seconds.
	StopAfterMs param.Opt[int64] `json:"stopAfterMs,omitzero"`
	// Only follow and scrape URLs matching this regex. A prefix pattern of the form
	// `^<start URL>` also follows a redirect of the start page.
	URLRegex param.Opt[string] `json:"urlRegex,omitzero"`
	// Keep only the main content, without navigation, headers, footers, or sidebars.
	UseMainContentOnly param.Opt[bool] `json:"useMainContentOnly,omitzero"`
	// Remove elements matching these CSS selectors. Overrides `includeSelectors`.
	ExcludeSelectors []string `json:"excludeSelectors,omitzero"`
	// Keep only elements matching these CSS selectors.
	IncludeSelectors []string `json:"includeSelectors,omitzero"`
	// Country to fetch from, as an ISO 3166-1 alpha-2 code.
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
	// Which iframe contents to include in each page's Markdown. `auto` inlines only
	// iframes that hold page content, such as same-site documents, framesets, job
	// boards, forms, and embedded documents, and skips ads and widgets. `true` inlines
	// every iframe; `false` inlines none.
	IncludeFrames WebWebCrawlMdParamsIncludeFramesUnion `json:"includeFrames,omitzero"`
	// Whether to parse PDFs and which pages to read.
	Pdf WebWebCrawlMdParamsPdf `json:"pdf,omitzero"`
	// Access policy IDs to apply; each URL must pass all of them. Requires access
	// policies enabled for your organization; otherwise returns
	// `ACCESS_POLICIES_NOT_ENABLED`.
	Policies []string `json:"policies,omitzero"`
	// Markdown sources in order of preference: `html-conversion` converts the page
	// HTML, `markdown-direct` uses Markdown the site publishes.
	//
	// Any of "html-conversion", "markdown-direct".
	SourcePreference []string `json:"sourcePreference,omitzero"`
	// Labels for this request. Filter usage and request logs by them in the dashboard
	// or with `GET /org/logs`.
	Tags []string `json:"tags,omitzero"`
	// Request deadline and what to return when it passes.
	TimeoutOpts WebWebCrawlMdParamsTimeoutOpts `json:"timeoutOpts,omitzero"`
	// `enabled` turns on zero data retention. Returns 403 `ZDR_NOT_ENABLED` unless
	// your organization has ZDR.
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

// Country to fetch from, as an ISO 3166-1 alpha-2 code.
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebWebCrawlMdParamsIncludeFramesUnion struct {
	OfBool param.Opt[bool] `json:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWebWebCrawlMdsIncludeFramesString)
	OfWebWebCrawlMdsIncludeFramesString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u WebWebCrawlMdParamsIncludeFramesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfWebWebCrawlMdsIncludeFramesString)
}
func (u *WebWebCrawlMdParamsIncludeFramesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type WebWebCrawlMdParamsIncludeFramesString string

const (
	WebWebCrawlMdParamsIncludeFramesStringAuto WebWebCrawlMdParamsIncludeFramesString = "auto"
)

// Whether to parse PDFs and which pages to read.
type WebWebCrawlMdParamsPdf struct {
	// Last page to parse. Defaults to the last page. Must be at least `start`.
	End param.Opt[int64] `json:"end,omitzero"`
	// Read scanned PDF pages with OCR. Pages that already have text are kept as is.
	Ocr param.Opt[bool] `json:"ocr,omitzero"`
	// Parse PDF pages. When `false`, PDFs are left out of the results and not counted
	// as failures.
	ShouldParse param.Opt[bool] `json:"shouldParse,omitzero"`
	// First page to parse, starting at 1.
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

// Request deadline and what to return when it passes.
//
// The property Milliseconds is required.
type WebWebCrawlMdParamsTimeoutOpts struct {
	// Deadline in milliseconds.
	Milliseconds int64 `json:"milliseconds" api:"required"`
	// `fail` returns 408 at the deadline. `return-partial` returns the results
	// collected so far and marks the response as partial.
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

// `enabled` turns on zero data retention. Returns 403 `ZDR_NOT_ENABLED` unless
// your organization has ZDR.
type WebWebCrawlMdParamsZdr string

const (
	WebWebCrawlMdParamsZdrEnabled  WebWebCrawlMdParamsZdr = "enabled"
	WebWebCrawlMdParamsZdrDisabled WebWebCrawlMdParamsZdr = "disabled"
)
