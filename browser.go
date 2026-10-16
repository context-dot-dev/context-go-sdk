// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contextdev

import (
	"github.com/context-dot-dev/context-go-sdk/v2/option"
)

// BrowserService contains methods and other services that help with interacting
// with the context.dev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBrowserService] method instead.
type BrowserService struct {
	options []option.RequestOption
	// See the authenticated browser profiles your organization created.
	Profiles BrowserProfileService
}

// NewBrowserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBrowserService(opts ...option.RequestOption) (r BrowserService) {
	r = BrowserService{}
	r.options = opts
	r.Profiles = NewBrowserProfileService(opts...)
	return
}
