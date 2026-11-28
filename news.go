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

// Search live and historical news by company or with a natural-language query.
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

// Find news about a company by name, domain, ticker, or ISIN, or search all news
// with a natural-language query. Each article includes its publisher, type of
// coverage, and, for company searches, whether the company is the main subject.
// Narrow results by publisher, country, language, article type, or publish date.
func (r *NewsService) Search(ctx context.Context, body NewsSearchParams, opts ...option.RequestOption) (res *NewsSearchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "news/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type NewsSearchResponse struct {
	// Articles on this page, in the requested order.
	Data []NewsSearchResponseData `json:"data" api:"required"`
	// Whether more results are available after this page.
	HasMore bool `json:"has_more" api:"required"`
	// Summary of this page.
	Meta NewsSearchResponseMeta `json:"meta" api:"required"`
	// Pass as `cursor` to fetch the next page. `null` on the last page.
	NextCursor string `json:"next_cursor" api:"required"`
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Credits this request used and your remaining balance.
	KeyMetadata NewsSearchResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		Meta        respjson.Field
		NextCursor  respjson.Field
		RequestID   respjson.Field
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
	// Unique ID of the article. It stays the same across requests, so you can use it
	// to deduplicate.
	ID string `json:"id" api:"required"`
	// Authors named in the byline. Empty when there is no byline.
	Authors []string `json:"authors" api:"required"`
	// Summary or excerpt of the article. For natural-language searches, it includes
	// the matching passages, up to 2,000 characters.
	Description string `json:"description" api:"required"`
	// URL of the article's lead image. `null` when there is none.
	ImageURL string `json:"image_url" api:"required"`
	// Language of the article, as an ISO 639-1 code such as `en`. `null` when unknown.
	Language string `json:"language" api:"required"`
	// How the article relates to the company you searched for. `null` for
	// natural-language searches.
	Match NewsSearchResponseDataMatch `json:"match" api:"required"`
	// When the article was published. `null` when the publisher does not give a
	// reliable date.
	PublishedAt time.Time `json:"published_at" api:"required" format:"date-time"`
	// Publication that published the article.
	Source NewsSearchResponseDataSource `json:"source" api:"required"`
	// ID shared by articles that cover the same story on the same day. Use it to group
	// syndicated copies of one announcement across outlets.
	StoryID string `json:"story_id" api:"required"`
	// Headline of the article.
	Title string `json:"title" api:"required"`
	// Type of coverage. Use it to separate independent reporting (`editorial`) from
	// company-issued content (`press_release`, `regulatory_filing`, `advisory`).
	//
	// Any of "editorial", "press_release", "regulatory_filing", "advisory".
	Type string `json:"type" api:"required"`
	// URL of the article on the publisher's site.
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

// How the article relates to the company you searched for. `null` for
// natural-language searches.
type NewsSearchResponseDataMatch struct {
	// Confidence in the match, from 0 to 1. `null` when no score is available.
	Confidence float64 `json:"confidence" api:"required"`
	// `primary` when the article is mainly about the company; `secondary` when the
	// company is mentioned but is not the main subject.
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

// Publication that published the article.
type NewsSearchResponseDataSource struct {
	// Whether the article came directly from the publisher's own feed rather than
	// through an aggregator.
	Direct bool `json:"direct" api:"required"`
	// Website domain of the publication, such as `reuters.com`.
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

// Summary of this page.
type NewsSearchResponseMeta struct {
	// Number of articles on this page.
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

// Credits this request used and your remaining balance.
type NewsSearchResponseKeyMetadata struct {
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
func (r NewsSearchResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *NewsSearchResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NewsSearchParams struct {
	// What to search for: news about a specific company, or a natural-language query
	// across all news.
	SearchBy NewsSearchParamsSearchByUnion `json:"searchBy,omitzero" api:"required"`
	// Cursor from the previous page's `next_cursor`. Omit for the first page.
	Cursor param.Opt[string] `json:"cursor,omitzero"`
	// Maximum number of articles to return per page.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// Filters for the results. Use at most one of `sourceDomain`, `sourceCountry`,
	// `articleLanguage`, or `articleType`, optionally combined with `date`.
	FilterBy NewsSearchParamsFilterBy `json:"filterBy,omitzero"`
	// Access policy IDs to apply; each URL must pass all of them. Requires access
	// policies enabled for your organization; otherwise returns
	// `ACCESS_POLICIES_NOT_ENABLED`.
	Policies []string `json:"policies,omitzero"`
	// How to order results. Company searches default to `newest`; natural-language
	// searches always use `relevance`.
	SortBy NewsSearchParamsSortBy `json:"sortBy,omitzero"`
	// Labels for this request. Filter usage and request logs by them in the dashboard
	// or with `GET /org/logs`.
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type NewsSearchParamsSearchByUnion struct {
	OfNaturalLanguage *NewsSearchParamsSearchByNaturalLanguage `json:",omitzero,inline"`
	OfEntity          *NewsSearchParamsSearchByEntity          `json:",omitzero,inline"`
	paramUnion
}

func (u NewsSearchParamsSearchByUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNaturalLanguage, u.OfEntity)
}
func (u *NewsSearchParamsSearchByUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[NewsSearchParamsSearchByUnion](
		"type",
		apijson.Discriminator[NewsSearchParamsSearchByNaturalLanguage]("natural_language"),
		apijson.Discriminator[NewsSearchParamsSearchByEntity]("entity"),
	)
}

// Search all news with a plain-language query. Returns up to 100 matches, ordered
// by relevance.
//
// The properties Query, Type are required.
type NewsSearchParamsSearchByNaturalLanguage struct {
	// What to search for, in plain language, such as
	// `chip export restrictions to China`.
	Query string `json:"query" api:"required"`
	// Set to `natural_language` to search all news with a plain-language query.
	//
	// This field can be elided, and will marshal its zero value as "natural_language".
	Type constant.NaturalLanguage `json:"type" default:"natural_language"`
	paramObj
}

func (r NewsSearchParamsSearchByNaturalLanguage) MarshalJSON() (data []byte, err error) {
	type shadow NewsSearchParamsSearchByNaturalLanguage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NewsSearchParamsSearchByNaturalLanguage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Find news about a specific company.
//
// The properties Entity, Type are required.
type NewsSearchParamsSearchByEntity struct {
	// The company to search news for, identified by name, domain, ticker, or ISIN.
	Entity NewsSearchParamsSearchByEntityEntityUnion `json:"entity,omitzero" api:"required"`
	// Set to `entity` to find news about a specific company.
	//
	// This field can be elided, and will marshal its zero value as "entity".
	Type constant.Entity `json:"type" default:"entity"`
	paramObj
}

func (r NewsSearchParamsSearchByEntity) MarshalJSON() (data []byte, err error) {
	type shadow NewsSearchParamsSearchByEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NewsSearchParamsSearchByEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type NewsSearchParamsSearchByEntityEntityUnion struct {
	OfName   *NewsSearchParamsSearchByEntityEntityName   `json:",omitzero,inline"`
	OfDomain *NewsSearchParamsSearchByEntityEntityDomain `json:",omitzero,inline"`
	OfTicker *NewsSearchParamsSearchByEntityEntityTicker `json:",omitzero,inline"`
	OfIsin   *NewsSearchParamsSearchByEntityEntityIsin   `json:",omitzero,inline"`
	paramUnion
}

func (u NewsSearchParamsSearchByEntityEntityUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfName, u.OfDomain, u.OfTicker, u.OfIsin)
}
func (u *NewsSearchParamsSearchByEntityEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[NewsSearchParamsSearchByEntityEntityUnion](
		"type",
		apijson.Discriminator[NewsSearchParamsSearchByEntityEntityName]("name"),
		apijson.Discriminator[NewsSearchParamsSearchByEntityEntityDomain]("domain"),
		apijson.Discriminator[NewsSearchParamsSearchByEntityEntityTicker]("ticker"),
		apijson.Discriminator[NewsSearchParamsSearchByEntityEntityIsin]("isin"),
	)
}

// Identify the company by name.
//
// The properties Name, Type are required.
type NewsSearchParamsSearchByEntityEntityName struct {
	// Company name, such as `Stripe`.
	Name string `json:"name" api:"required"`
	// Set to `name` to identify the company by name.
	//
	// This field can be elided, and will marshal its zero value as "name".
	Type constant.Name `json:"type" default:"name"`
	paramObj
}

func (r NewsSearchParamsSearchByEntityEntityName) MarshalJSON() (data []byte, err error) {
	type shadow NewsSearchParamsSearchByEntityEntityName
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NewsSearchParamsSearchByEntityEntityName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identify the company by website domain.
//
// The properties Domain, Type are required.
type NewsSearchParamsSearchByEntityEntityDomain struct {
	// Company website domain, such as `apple.com`.
	Domain string `json:"domain" api:"required"`
	// Set to `domain` to identify the company by website domain.
	//
	// This field can be elided, and will marshal its zero value as "domain".
	Type constant.Domain `json:"type" default:"domain"`
	paramObj
}

func (r NewsSearchParamsSearchByEntityEntityDomain) MarshalJSON() (data []byte, err error) {
	type shadow NewsSearchParamsSearchByEntityEntityDomain
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NewsSearchParamsSearchByEntityEntityDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identify the company by stock ticker, optionally on a specific exchange.
//
// The properties Ticker, Type are required.
type NewsSearchParamsSearchByEntityEntityTicker struct {
	// Stock ticker symbol, such as `AAPL`.
	Ticker string `json:"ticker" api:"required"`
	// Exchange the ticker trades on. Use it when the same ticker is listed on more
	// than one exchange.
	//
	// Any of "AMEX", "AMS", "AQS", "ASX", "ATH", "BER", "BME", "BRU", "BSE", "BUD",
	// "BUE", "BVC", "CBOE", "CNQ", "CPH", "DFM", "DOH", "DUB", "DUS", "DXE", "EGX",
	// "FSX", "HAM", "HEL", "HKSE", "HOSE", "ICE", "IOB", "IST", "JKT", "JNB", "JPX",
	// "KLS", "KOE", "KSC", "KUW", "LIS", "LSE", "MCX", "MEX", "MIL", "MUN", "NASDAQ",
	// "NEO", "NSE", "NYSE", "NZE", "OSL", "OTC", "PAR", "PNK", "PRA", "RIS", "SAO",
	// "SAU", "SES", "SET", "SGO", "SHH", "SHZ", "SIX", "STO", "STU", "TAI", "TAL",
	// "TLV", "TSX", "TSXV", "TWO", "VIE", "WSE", "XETRA".
	Exchange string `json:"exchange,omitzero"`
	// Set to `ticker` to identify a publicly traded company.
	//
	// This field can be elided, and will marshal its zero value as "ticker".
	Type constant.Ticker `json:"type" default:"ticker"`
	paramObj
}

func (r NewsSearchParamsSearchByEntityEntityTicker) MarshalJSON() (data []byte, err error) {
	type shadow NewsSearchParamsSearchByEntityEntityTicker
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NewsSearchParamsSearchByEntityEntityTicker) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[NewsSearchParamsSearchByEntityEntityTicker](
		"exchange", "AMEX", "AMS", "AQS", "ASX", "ATH", "BER", "BME", "BRU", "BSE", "BUD", "BUE", "BVC", "CBOE", "CNQ", "CPH", "DFM", "DOH", "DUB", "DUS", "DXE", "EGX", "FSX", "HAM", "HEL", "HKSE", "HOSE", "ICE", "IOB", "IST", "JKT", "JNB", "JPX", "KLS", "KOE", "KSC", "KUW", "LIS", "LSE", "MCX", "MEX", "MIL", "MUN", "NASDAQ", "NEO", "NSE", "NYSE", "NZE", "OSL", "OTC", "PAR", "PNK", "PRA", "RIS", "SAO", "SAU", "SES", "SET", "SGO", "SHH", "SHZ", "SIX", "STO", "STU", "TAI", "TAL", "TLV", "TSX", "TSXV", "TWO", "VIE", "WSE", "XETRA",
	)
}

// Identify the company by International Securities Identification Number (ISIN).
//
// The properties Isin, Type are required.
type NewsSearchParamsSearchByEntityEntityIsin struct {
	// International Securities Identification Number, such as `US0378331005`.
	Isin string `json:"isin" api:"required"`
	// Set to `isin` to identify the company by its International Securities
	// Identification Number.
	//
	// This field can be elided, and will marshal its zero value as "isin".
	Type constant.Isin `json:"type" default:"isin"`
	paramObj
}

func (r NewsSearchParamsSearchByEntityEntityIsin) MarshalJSON() (data []byte, err error) {
	type shadow NewsSearchParamsSearchByEntityEntityIsin
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NewsSearchParamsSearchByEntityEntityIsin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filters for the results. Use at most one of `sourceDomain`, `sourceCountry`,
// `articleLanguage`, or `articleType`, optionally combined with `date`.
type NewsSearchParamsFilterBy struct {
	// Only return articles in these languages, as ISO 639 codes such as `en`.
	//
	// Any of "am", "ar", "az", "be", "bg", "bn", "ca", "ceb", "cs", "cy", "da", "de",
	// "el", "en", "es", "et", "eu", "fa", "fi", "fr", "gu", "ha", "haw", "he", "hi",
	// "hr", "hu", "hy", "id", "is", "it", "ja", "ka", "kk", "kn", "ko", "ku", "ky",
	// "lo", "lt", "lv", "mk", "ml", "mn", "mr", "ms", "ne", "nl", "no", "or", "pa",
	// "pl", "ps", "pt", "ro", "ru", "sk", "sl", "so", "sq", "sr", "sv", "sw", "ta",
	// "te", "th", "tl", "tr", "uk", "ur", "uz", "vi", "yo", "zh".
	ArticleLanguage []string `json:"articleLanguage,omitzero"`
	// Only return these types of coverage, such as `press_release`.
	//
	// Any of "editorial", "press_release", "regulatory_filing", "advisory".
	ArticleType []string `json:"articleType,omitzero"`
	// Publish date range, in epoch milliseconds. `from` must not be later than `to`.
	Date NewsSearchParamsFilterByDate `json:"date,omitzero"`
	// Only return articles from publishers in these countries, as ISO 3166-1 alpha-2
	// codes such as `us`.
	//
	// Any of "ae", "ar", "au", "ca", "cg", "ch", "cl", "cz", "de", "fi", "fr", "gb",
	// "hk", "il", "in", "jp", "kr", "mx", "ng", "nl", "pk", "qa", "sa", "se", "sg",
	// "us", "za".
	SourceCountry []string `json:"sourceCountry,omitzero"`
	// Only return articles from these publisher domains, such as `reuters.com`.
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

// Publish date range, in epoch milliseconds. `from` must not be later than `to`.
type NewsSearchParamsFilterByDate struct {
	// Inclusive start of the publish date range, in epoch milliseconds.
	From param.Opt[int64] `json:"from,omitzero"`
	// Inclusive end of the publish date range, in epoch milliseconds.
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

// How to order results. Company searches default to `newest`; natural-language
// searches always use `relevance`.
//
// The property Type is required.
type NewsSearchParamsSortBy struct {
	// `newest` puts the most recently published articles first; `relevance` puts the
	// closest matches first.
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
