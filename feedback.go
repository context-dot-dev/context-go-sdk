// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contextdev

import (
	"context"
	"net/http"
	"slices"

	"github.com/context-dot-dev/context-go-sdk/v2/internal/apijson"
	"github.com/context-dot-dev/context-go-sdk/v2/internal/requestconfig"
	"github.com/context-dot-dev/context-go-sdk/v2/option"
	"github.com/context-dot-dev/context-go-sdk/v2/packages/param"
	"github.com/context-dot-dev/context-go-sdk/v2/packages/respjson"
)

// Report bugs, docs mismatches, and friction with any Context.dev API. Submissions
// cost no credits and use a separate rate limit.
//
// FeedbackService contains methods and other services that help with interacting
// with the context.dev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFeedbackService] method instead.
type FeedbackService struct {
	options []option.RequestOption
}

// NewFeedbackService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFeedbackService(opts ...option.RequestOption) (r FeedbackService) {
	r = FeedbackService{}
	r.options = opts
	return
}

// Report a problem with a Context.dev API call, docs page, SDK, or CLI. Include
// request_id, url, or both.
func (r *FeedbackService) Submit(ctx context.Context, body FeedbackSubmitParams, opts ...option.RequestOption) (res *FeedbackSubmitResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "feedback"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type FeedbackSubmitResponse struct {
	// True when feedback for this request_id was already recorded; the original
	// feedback_id is returned.
	AlreadySubmitted bool `json:"already_submitted" api:"required"`
	// ID of the stored feedback.
	FeedbackID string `json:"feedback_id" api:"required"`
	// Unique id of this API call, also sent in the X-Request-Id response header. Quote
	// it when contacting support about a failed request.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Credit usage, included whenever a valid API key is provided.
	KeyMetadata FeedbackSubmitResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AlreadySubmitted respjson.Field
		FeedbackID       respjson.Field
		RequestID        respjson.Field
		KeyMetadata      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FeedbackSubmitResponse) RawJSON() string { return r.JSON.raw }
func (r *FeedbackSubmitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credit usage, included whenever a valid API key is provided.
type FeedbackSubmitResponseKeyMetadata struct {
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
func (r FeedbackSubmitResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *FeedbackSubmitResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FeedbackSubmitParams struct {
	// Kind of issue.
	//
	// Any of "bug", "docs_mismatch", "friction", "feature_gap", "quality_degradation",
	// "other".
	Category FeedbackSubmitParamsCategory `json:"category,omitzero" api:"required"`
	// What went wrong and what you expected instead.
	Note string `json:"note" api:"required"`
	// The request_id of the API call the feedback is about, from its response body or
	// X-Request-Id header.
	RequestID param.Opt[string] `json:"request_id,omitzero" format:"uuid"`
	// The page the feedback is about, such as one page of a crawl or a docs page.
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	// Optional tags for tracking usage. Up to 20 tags, each 1 to 50 characters.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r FeedbackSubmitParams) MarshalJSON() (data []byte, err error) {
	type shadow FeedbackSubmitParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FeedbackSubmitParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Kind of issue.
type FeedbackSubmitParamsCategory string

const (
	FeedbackSubmitParamsCategoryBug                FeedbackSubmitParamsCategory = "bug"
	FeedbackSubmitParamsCategoryDocsMismatch       FeedbackSubmitParamsCategory = "docs_mismatch"
	FeedbackSubmitParamsCategoryFriction           FeedbackSubmitParamsCategory = "friction"
	FeedbackSubmitParamsCategoryFeatureGap         FeedbackSubmitParamsCategory = "feature_gap"
	FeedbackSubmitParamsCategoryQualityDegradation FeedbackSubmitParamsCategory = "quality_degradation"
	FeedbackSubmitParamsCategoryOther              FeedbackSubmitParamsCategory = "other"
)
