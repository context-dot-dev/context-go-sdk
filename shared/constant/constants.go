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

type ByDirectURL string   // Always "by_direct_url"
type ByDomain string      // Always "by_domain"
type ByEmail string       // Always "by_email"
type ByName string        // Always "by_name"
type ByTicker string      // Always "by_ticker"
type ByTransaction string // Always "by_transaction"
type Candidate string     // Always "candidate"
type Crawl string         // Always "crawl"
type Domain string        // Always "domain"
type Error string         // Always "error"
type Exact string         // Always "exact"
type Extract string       // Always "extract"
type HTML string          // Always "html"
type Isin string          // Always "isin"
type Markdown string      // Always "markdown"
type Name string          // Always "name"
type NotFound string      // Always "not_found"
type Ok string            // Always "ok"
type Page string          // Always "page"
type Perform string       // Always "perform"
type Scrape string        // Always "scrape"
type Semantic string      // Always "semantic"
type Sitemap string       // Always "sitemap"
type StartURL string      // Always "start_url"
type Ticker string        // Always "ticker"
type Wait string          // Always "wait"

func (c ByDirectURL) Default() ByDirectURL     { return "by_direct_url" }
func (c ByDomain) Default() ByDomain           { return "by_domain" }
func (c ByEmail) Default() ByEmail             { return "by_email" }
func (c ByName) Default() ByName               { return "by_name" }
func (c ByTicker) Default() ByTicker           { return "by_ticker" }
func (c ByTransaction) Default() ByTransaction { return "by_transaction" }
func (c Candidate) Default() Candidate         { return "candidate" }
func (c Crawl) Default() Crawl                 { return "crawl" }
func (c Domain) Default() Domain               { return "domain" }
func (c Error) Default() Error                 { return "error" }
func (c Exact) Default() Exact                 { return "exact" }
func (c Extract) Default() Extract             { return "extract" }
func (c HTML) Default() HTML                   { return "html" }
func (c Isin) Default() Isin                   { return "isin" }
func (c Markdown) Default() Markdown           { return "markdown" }
func (c Name) Default() Name                   { return "name" }
func (c NotFound) Default() NotFound           { return "not_found" }
func (c Ok) Default() Ok                       { return "ok" }
func (c Page) Default() Page                   { return "page" }
func (c Perform) Default() Perform             { return "perform" }
func (c Scrape) Default() Scrape               { return "scrape" }
func (c Semantic) Default() Semantic           { return "semantic" }
func (c Sitemap) Default() Sitemap             { return "sitemap" }
func (c StartURL) Default() StartURL           { return "start_url" }
func (c Ticker) Default() Ticker               { return "ticker" }
func (c Wait) Default() Wait                   { return "wait" }

func (c ByDirectURL) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c ByDomain) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c ByEmail) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c ByName) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c ByTicker) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c ByTransaction) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Candidate) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c Crawl) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c Domain) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c Error) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c Exact) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c Extract) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c HTML) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c Isin) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c Markdown) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c Name) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c NotFound) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c Ok) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c Page) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c Perform) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c Scrape) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c Semantic) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c Sitemap) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c StartURL) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c Ticker) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c Wait) MarshalJSON() ([]byte, error)          { return marshalString(c) }

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
