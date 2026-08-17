// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contextdev

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/context-dot-dev/context-go-sdk/v2/internal/apijson"
	"github.com/context-dot-dev/context-go-sdk/v2/internal/requestconfig"
	"github.com/context-dot-dev/context-go-sdk/v2/option"
	"github.com/context-dot-dev/context-go-sdk/v2/packages/param"
	"github.com/context-dot-dev/context-go-sdk/v2/packages/respjson"
	"github.com/context-dot-dev/context-go-sdk/v2/shared/constant"
)

// Search live first-party RSS and free historical news data by company identity.
//
// NewsService contains methods and other services that help with interacting with
// the context.dev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNewsService] method instead.
type NewsService struct {
	options []option.RequestOption
}

// NewNewsService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewNewsService(opts ...option.RequestOption) (r NewsService) {
	r = NewsService{}
	r.options = opts
	return
}

// Searches live and historical company news for one company, identified in
// searchBy by name, domain, ticker (optionally disambiguated by exchange), or
// ISIN. Results can be filtered by publisher domain, publisher country, article
// language, article type, and published-at date, and include stable story IDs,
// source metadata, verified entity relevance, and cursor pagination.
func (r *NewsService) Search(ctx context.Context, body NewsSearchParams, opts ...option.RequestOption) (res *NewsSearchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "news/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type NewsSearchResponse struct {
	// Articles matching the search, in the requested order.
	Data []NewsSearchResponseData `json:"data" api:"required"`
	// True when more results are available beyond this page.
	HasMore bool `json:"has_more" api:"required"`
	// Summary information about this response.
	Meta NewsSearchResponseMeta `json:"meta" api:"required"`
	// Pass as cursor in the next request to fetch the following page. Null when there
	// are no more results.
	NextCursor string `json:"next_cursor" api:"required"`
	// Metadata about the API key used for the request. Included in every response
	// whenever a valid API key is provided, even when the response status is not 200.
	KeyMetadata NewsSearchResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		Meta        respjson.Field
		NextCursor  respjson.Field
		KeyMetadata respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewsSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *NewsSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NewsSearchResponseData struct {
	// Stable unique identifier for this article. Use it to deduplicate or reference an
	// article across requests.
	ID string `json:"id" api:"required"`
	// Bylined authors. Empty when no byline is available.
	Authors []string `json:"authors" api:"required"`
	// Short summary or excerpt of the article, when the publisher provides one.
	Description string `json:"description" api:"required"`
	// Lead image for the article, when one is available.
	ImageURL string `json:"image_url" api:"required"`
	// Language the article is written in, as a lowercase ISO 639-1 code such as en.
	// Null when unknown.
	Language string `json:"language" api:"required"`
	// How the article relates to the company you searched for.
	Match NewsSearchResponseDataMatch `json:"match" api:"required"`
	// When the article was published, as an ISO 8601 timestamp. Null when the
	// publisher does not state a reliable date.
	PublishedAt time.Time `json:"published_at" api:"required" format:"date-time"`
	// The publication that published the article.
	Source NewsSearchResponseDataSource `json:"source" api:"required"`
	// Shared by articles covering the same story on the same day. Use it to group or
	// collapse syndicated copies of one announcement across outlets.
	StoryID string `json:"story_id" api:"required"`
	// Article headline.
	Title string `json:"title" api:"required"`
	// Kind of coverage. Use it to separate independent reporting (editorial) from
	// company-issued content (press_release, regulatory_filing, advisory).
	//
	// Any of "editorial", "press_release", "regulatory_filing", "advisory".
	Type string `json:"type" api:"required"`
	// Link to the article on the publisher site.
	URL string `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Authors     respjson.Field
		Description respjson.Field
		ImageURL    respjson.Field
		Language    respjson.Field
		Match       respjson.Field
		PublishedAt respjson.Field
		Source      respjson.Field
		StoryID     respjson.Field
		Title       respjson.Field
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewsSearchResponseData) RawJSON() string { return r.JSON.raw }
func (r *NewsSearchResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How the article relates to the company you searched for.
type NewsSearchResponseDataMatch struct {
	// How confident the match is, from 0 to 1. Null when a score is unavailable.
	Confidence float64 `json:"confidence" api:"required"`
	// primary when the article is mainly about the company, secondary when the company
	// is mentioned but is not the main subject.
	//
	// Any of "primary", "secondary".
	Level string `json:"level" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Confidence  respjson.Field
		Level       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewsSearchResponseDataMatch) RawJSON() string { return r.JSON.raw }
func (r *NewsSearchResponseDataMatch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The publication that published the article.
type NewsSearchResponseDataSource struct {
	// True when Context observed this article in the publisher-owned feed.
	Direct bool `json:"direct" api:"required"`
	// Website domain of the publication.
	Domain string `json:"domain" api:"required"`
	// Name of the publication, such as Reuters.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Direct      respjson.Field
		Domain      respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewsSearchResponseDataSource) RawJSON() string { return r.JSON.raw }
func (r *NewsSearchResponseDataSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Summary information about this response.
type NewsSearchResponseMeta struct {
	// Number of articles in this page.
	Count int64 `json:"count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewsSearchResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *NewsSearchResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the API key used for the request. Included in every response
// whenever a valid API key is provided, even when the response status is not 200.
type NewsSearchResponseKeyMetadata struct {
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
func (r NewsSearchResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *NewsSearchResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NewsSearchParams struct {
	// What to search for.
	SearchBy NewsSearchParamsSearchBy `json:"searchBy,omitzero" api:"required"`
	// Opaque next_cursor from the previous response, or null for the first page.
	Cursor param.Opt[string] `json:"cursor,omitzero"`
	// Maximum results to return. Defaults to 10.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// Optional result filters.
	FilterBy NewsSearchParamsFilterBy `json:"filterBy,omitzero"`
	// Result ordering. Defaults to newest.
	SortBy NewsSearchParamsSortBy `json:"sortBy,omitzero"`
	// Optional tags for tracking usage. Up to 20 tags, each 1 to 50 characters.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r NewsSearchParams) MarshalJSON() (data []byte, err error) {
	type shadow NewsSearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NewsSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What to search for.
//
// The properties Entity, Type are required.
type NewsSearchParamsSearchBy struct {
	// The company to search news for, identified by name, domain, ticker, or ISIN.
	Entity NewsSearchParamsSearchByEntityUnion `json:"entity,omitzero" api:"required"`
	// How to search. Only entity search is supported.
	//
	// Any of "entity".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r NewsSearchParamsSearchBy) MarshalJSON() (data []byte, err error) {
	type shadow NewsSearchParamsSearchBy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NewsSearchParamsSearchBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[NewsSearchParamsSearchBy](
		"type", "entity",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type NewsSearchParamsSearchByEntityUnion struct {
	OfName   *NewsSearchParamsSearchByEntityName   `json:",omitzero,inline"`
	OfDomain *NewsSearchParamsSearchByEntityDomain `json:",omitzero,inline"`
	OfTicker *NewsSearchParamsSearchByEntityTicker `json:",omitzero,inline"`
	OfIsin   *NewsSearchParamsSearchByEntityIsin   `json:",omitzero,inline"`
	paramUnion
}

func (u NewsSearchParamsSearchByEntityUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfName, u.OfDomain, u.OfTicker, u.OfIsin)
}
func (u *NewsSearchParamsSearchByEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[NewsSearchParamsSearchByEntityUnion](
		"type",
		apijson.Discriminator[NewsSearchParamsSearchByEntityName]("name"),
		apijson.Discriminator[NewsSearchParamsSearchByEntityDomain]("domain"),
		apijson.Discriminator[NewsSearchParamsSearchByEntityTicker]("ticker"),
		apijson.Discriminator[NewsSearchParamsSearchByEntityIsin]("isin"),
	)
}

// Identify the company by name.
//
// The properties Name, Type are required.
type NewsSearchParamsSearchByEntityName struct {
	// Company name.
	Name string `json:"name" api:"required"`
	// This field can be elided, and will marshal its zero value as "name".
	Type constant.Name `json:"type" default:"name"`
	paramObj
}

func (r NewsSearchParamsSearchByEntityName) MarshalJSON() (data []byte, err error) {
	type shadow NewsSearchParamsSearchByEntityName
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NewsSearchParamsSearchByEntityName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identify the company by website domain.
//
// The properties Domain, Type are required.
type NewsSearchParamsSearchByEntityDomain struct {
	// Company website domain, such as apple.com.
	Domain string `json:"domain" api:"required"`
	// This field can be elided, and will marshal its zero value as "domain".
	Type constant.Domain `json:"type" default:"domain"`
	paramObj
}

func (r NewsSearchParamsSearchByEntityDomain) MarshalJSON() (data []byte, err error) {
	type shadow NewsSearchParamsSearchByEntityDomain
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NewsSearchParamsSearchByEntityDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identify the company by stock ticker, optionally scoped to an exchange.
//
// The properties Ticker, Type are required.
type NewsSearchParamsSearchByEntityTicker struct {
	// Public-company ticker.
	Ticker string `json:"ticker" api:"required"`
	// Stock exchange the ticker trades on, used to disambiguate tickers listed on
	// multiple exchanges.
	//
	// Any of "AMEX", "AMS", "AQS", "ASX", "ATH", "BER", "BME", "BRU", "BSE", "BUD",
	// "BUE", "BVC", "CBOE", "CNQ", "CPH", "DFM", "DOH", "DUB", "DUS", "DXE", "EGX",
	// "FSX", "HAM", "HEL", "HKSE", "HOSE", "ICE", "IOB", "IST", "JKT", "JNB", "JPX",
	// "KLS", "KOE", "KSC", "KUW", "LIS", "LSE", "MCX", "MEX", "MIL", "MUN", "NASDAQ",
	// "NEO", "NSE", "NYSE", "NZE", "OSL", "OTC", "PAR", "PNK", "PRA", "RIS", "SAO",
	// "SAU", "SES", "SET", "SGO", "SHH", "SHZ", "SIX", "STO", "STU", "TAI", "TAL",
	// "TLV", "TSX", "TSXV", "TWO", "VIE", "WSE", "XETRA".
	Exchange string `json:"exchange,omitzero"`
	// This field can be elided, and will marshal its zero value as "ticker".
	Type constant.Ticker `json:"type" default:"ticker"`
	paramObj
}

func (r NewsSearchParamsSearchByEntityTicker) MarshalJSON() (data []byte, err error) {
	type shadow NewsSearchParamsSearchByEntityTicker
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NewsSearchParamsSearchByEntityTicker) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[NewsSearchParamsSearchByEntityTicker](
		"exchange", "AMEX", "AMS", "AQS", "ASX", "ATH", "BER", "BME", "BRU", "BSE", "BUD", "BUE", "BVC", "CBOE", "CNQ", "CPH", "DFM", "DOH", "DUB", "DUS", "DXE", "EGX", "FSX", "HAM", "HEL", "HKSE", "HOSE", "ICE", "IOB", "IST", "JKT", "JNB", "JPX", "KLS", "KOE", "KSC", "KUW", "LIS", "LSE", "MCX", "MEX", "MIL", "MUN", "NASDAQ", "NEO", "NSE", "NYSE", "NZE", "OSL", "OTC", "PAR", "PNK", "PRA", "RIS", "SAO", "SAU", "SES", "SET", "SGO", "SHH", "SHZ", "SIX", "STO", "STU", "TAI", "TAL", "TLV", "TSX", "TSXV", "TWO", "VIE", "WSE", "XETRA",
	)
}

// Identify the company by International Securities Identification Number.
//
// The properties Isin, Type are required.
type NewsSearchParamsSearchByEntityIsin struct {
	// International Securities Identification Number.
	Isin string `json:"isin" api:"required"`
	// This field can be elided, and will marshal its zero value as "isin".
	Type constant.Isin `json:"type" default:"isin"`
	paramObj
}

func (r NewsSearchParamsSearchByEntityIsin) MarshalJSON() (data []byte, err error) {
	type shadow NewsSearchParamsSearchByEntityIsin
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NewsSearchParamsSearchByEntityIsin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional result filters.
type NewsSearchParamsFilterBy struct {
	// Article languages to include. Up to 3.
	//
	// Any of "ar", "de", "en", "es", "fr", "hi", "it", "ja", "ko", "nl", "pt", "ru",
	// "zh".
	ArticleLanguage []string `json:"articleLanguage,omitzero"`
	// Article types to include. Up to 3.
	//
	// Any of "editorial", "press_release", "regulatory_filing", "advisory".
	ArticleType []string `json:"articleType,omitzero"`
	// Published-at window in epoch milliseconds.
	Date NewsSearchParamsFilterByDate `json:"date,omitzero"`
	// Publisher countries to include, as lowercase ISO 3166-1 alpha-2 codes. Up to 3.
	//
	// Any of "ae", "ar", "au", "ca", "cg", "ch", "cl", "de", "fi", "fr", "gb", "hk",
	// "il", "in", "jp", "kr", "mx", "ng", "nl", "qa", "sa", "se", "sg", "us", "za".
	SourceCountry []string `json:"sourceCountry,omitzero"`
	// Publisher domains to include. Up to 3.
	SourceDomain []string `json:"sourceDomain,omitzero"`
	paramObj
}

func (r NewsSearchParamsFilterBy) MarshalJSON() (data []byte, err error) {
	type shadow NewsSearchParamsFilterBy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NewsSearchParamsFilterBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Published-at window in epoch milliseconds.
type NewsSearchParamsFilterByDate struct {
	// Inclusive start of the published-at window, in epoch milliseconds.
	From param.Opt[int64] `json:"from,omitzero"`
	// Inclusive end of the published-at window, in epoch milliseconds.
	To param.Opt[int64] `json:"to,omitzero"`
	paramObj
}

func (r NewsSearchParamsFilterByDate) MarshalJSON() (data []byte, err error) {
	type shadow NewsSearchParamsFilterByDate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NewsSearchParamsFilterByDate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result ordering. Defaults to newest.
//
// The property Type is required.
type NewsSearchParamsSortBy struct {
	// Result ordering.
	//
	// Any of "relevance", "newest".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r NewsSearchParamsSortBy) MarshalJSON() (data []byte, err error) {
	type shadow NewsSearchParamsSortBy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NewsSearchParamsSortBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[NewsSearchParamsSortBy](
		"type", "relevance", "newest",
	)
}
