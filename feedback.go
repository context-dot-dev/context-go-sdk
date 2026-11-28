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

// Report bugs, wrong results, and docs mismatches to the Context.dev team.
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

// Report a bug, docs mismatch, or wrong result to the Context.dev team. Send
// `request_id`, `url`, or both. Any API key can submit, and a repeat report for
// the same `request_id` returns the original feedback.
func (r *FeedbackService) Submit(ctx context.Context, body FeedbackSubmitParams, opts ...option.RequestOption) (res *FeedbackSubmitResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "feedback"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type FeedbackSubmitResponse struct {
	// `true` when feedback for this `request_id` was already recorded. `feedback_id`
	// is then the original ID.
	AlreadySubmitted bool `json:"already_submitted" api:"required"`
	// Unique ID of the feedback.
	FeedbackID string `json:"feedback_id" api:"required"`
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Credits this request used and your remaining balance.
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

// Credits this request used and your remaining balance.
type FeedbackSubmitResponseKeyMetadata struct {
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
func (r FeedbackSubmitResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *FeedbackSubmitResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FeedbackSubmitParams struct {
	// Kind of issue. Use `quality_degradation` for wrong or incomplete results and
	// `bad_brand` for wrong or outdated brand data, such as a logo or colors.
	//
	// Any of "bug", "docs_mismatch", "friction", "feature_gap", "quality_degradation",
	// "bad_brand", "other".
	Category FeedbackSubmitParamsCategory `json:"category,omitzero" api:"required"`
	// What went wrong and what you expected instead. Leave out API keys and personal
	// data.
	Note string `json:"note" api:"required"`
	// ID of the API call the feedback is about, from its `request_id` or
	// `X-Request-Id` header. Send this, `url`, or both.
	RequestID param.Opt[string] `json:"request_id,omitzero" format:"uuid"`
	// URL the feedback is about, such as one page of a crawl or a docs page. Send
	// this, `request_id`, or both.
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	// Labels for this request. Filter usage and request logs by them in the dashboard
	// or with `GET /org/logs`.
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

// Kind of issue. Use `quality_degradation` for wrong or incomplete results and
// `bad_brand` for wrong or outdated brand data, such as a logo or colors.
type FeedbackSubmitParamsCategory string

const (
	FeedbackSubmitParamsCategoryBug                FeedbackSubmitParamsCategory = "bug"
	FeedbackSubmitParamsCategoryDocsMismatch       FeedbackSubmitParamsCategory = "docs_mismatch"
	FeedbackSubmitParamsCategoryFriction           FeedbackSubmitParamsCategory = "friction"
	FeedbackSubmitParamsCategoryFeatureGap         FeedbackSubmitParamsCategory = "feature_gap"
	FeedbackSubmitParamsCategoryQualityDegradation FeedbackSubmitParamsCategory = "quality_degradation"
	FeedbackSubmitParamsCategoryBadBrand           FeedbackSubmitParamsCategory = "bad_brand"
	FeedbackSubmitParamsCategoryOther              FeedbackSubmitParamsCategory = "other"
)
