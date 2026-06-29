// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/context-dot-dev/context-go-sdk/internal/encoding/json"
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

type Exact string    // Always "exact"
type Extract string  // Always "extract"
type Page string     // Always "page"
type Semantic string // Always "semantic"
type Sitemap string  // Always "sitemap"

func (c Exact) Default() Exact       { return "exact" }
func (c Extract) Default() Extract   { return "extract" }
func (c Page) Default() Page         { return "page" }
func (c Semantic) Default() Semantic { return "semantic" }
func (c Sitemap) Default() Sitemap   { return "sitemap" }

func (c Exact) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c Extract) MarshalJSON() ([]byte, error)  { return marshalString(c) }
func (c Page) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c Semantic) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Sitemap) MarshalJSON() ([]byte, error)  { return marshalString(c) }

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
