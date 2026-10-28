// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/context-dot-dev/context-go-sdk/v2/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type Batch string           // Always "batch"
type ByDirectURL string     // Always "by_direct_url"
type ByDomain string        // Always "by_domain"
type ByEmail string         // Always "by_email"
type ByName string          // Always "by_name"
type ByTicker string        // Always "by_ticker"
type ByTransaction string   // Always "by_transaction"
type Candidate string       // Always "candidate"
type Crawl string           // Always "crawl"
type Domain string          // Always "domain"
type Entity string          // Always "entity"
type Error string           // Always "error"
type Exact string           // Always "exact"
type Existing string        // Always "existing"
type Extract string         // Always "extract"
type Isin string            // Always "isin"
type Monitor string         // Always "monitor"
type Name string            // Always "name"
type NaturalLanguage string // Always "natural_language"
type New string             // Always "new"
type NotFound string        // Always "not_found"
type Ok string              // Always "ok"
type Page string            // Always "page"
type Perform string         // Always "perform"
type Prebuilt string        // Always "prebuilt"
type Scrape string          // Always "scrape"
type Scroll string          // Always "scroll"
type Search string          // Always "search"
type Semantic string        // Always "semantic"
type Sitemap string         // Always "sitemap"
type StartURL string        // Always "start_url"
type Ticker string          // Always "ticker"
type Wait string            // Always "wait"
type WaitFor string         // Always "waitFor"

func (c Batch) Default() Batch                     { return "batch" }
func (c ByDirectURL) Default() ByDirectURL         { return "by_direct_url" }
func (c ByDomain) Default() ByDomain               { return "by_domain" }
func (c ByEmail) Default() ByEmail                 { return "by_email" }
func (c ByName) Default() ByName                   { return "by_name" }
func (c ByTicker) Default() ByTicker               { return "by_ticker" }
func (c ByTransaction) Default() ByTransaction     { return "by_transaction" }
func (c Candidate) Default() Candidate             { return "candidate" }
func (c Crawl) Default() Crawl                     { return "crawl" }
func (c Domain) Default() Domain                   { return "domain" }
func (c Entity) Default() Entity                   { return "entity" }
func (c Error) Default() Error                     { return "error" }
func (c Exact) Default() Exact                     { return "exact" }
func (c Existing) Default() Existing               { return "existing" }
func (c Extract) Default() Extract                 { return "extract" }
func (c Isin) Default() Isin                       { return "isin" }
func (c Monitor) Default() Monitor                 { return "monitor" }
func (c Name) Default() Name                       { return "name" }
func (c NaturalLanguage) Default() NaturalLanguage { return "natural_language" }
func (c New) Default() New                         { return "new" }
func (c NotFound) Default() NotFound               { return "not_found" }
func (c Ok) Default() Ok                           { return "ok" }
func (c Page) Default() Page                       { return "page" }
func (c Perform) Default() Perform                 { return "perform" }
func (c Prebuilt) Default() Prebuilt               { return "prebuilt" }
func (c Scrape) Default() Scrape                   { return "scrape" }
func (c Scroll) Default() Scroll                   { return "scroll" }
func (c Search) Default() Search                   { return "search" }
func (c Semantic) Default() Semantic               { return "semantic" }
func (c Sitemap) Default() Sitemap                 { return "sitemap" }
func (c StartURL) Default() StartURL               { return "start_url" }
func (c Ticker) Default() Ticker                   { return "ticker" }
func (c Wait) Default() Wait                       { return "wait" }
func (c WaitFor) Default() WaitFor                 { return "waitFor" }

func (c Batch) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c ByDirectURL) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c ByDomain) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c ByEmail) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c ByName) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c ByTicker) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c ByTransaction) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c Candidate) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c Crawl) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c Domain) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c Entity) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c Error) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c Exact) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c Existing) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c Extract) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c Isin) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c Monitor) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c Name) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c NaturalLanguage) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c New) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c NotFound) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c Ok) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c Page) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c Perform) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c Prebuilt) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c Scrape) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c Scroll) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c Search) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c Semantic) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c Sitemap) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c StartURL) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c Ticker) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c Wait) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c WaitFor) MarshalJSON() ([]byte, error)         { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
