// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contextdev

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/context-dot-dev/context-go-sdk/v2/internal/apijson"
	"github.com/context-dot-dev/context-go-sdk/v2/internal/requestconfig"
	"github.com/context-dot-dev/context-go-sdk/v2/option"
	"github.com/context-dot-dev/context-go-sdk/v2/packages/param"
	"github.com/context-dot-dev/context-go-sdk/v2/packages/respjson"
	"github.com/context-dot-dev/context-go-sdk/v2/shared/constant"
)

// PersonService contains methods and other services that help with interacting
// with the context.dev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPersonService] method instead.
type PersonService struct {
	options []option.RequestOption
}

// NewPersonService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPersonService(opts ...option.RequestOption) (r PersonService) {
	r = PersonService{}
	r.options = opts
	return
}

// Finds and normalizes the best available person candidate from additive identity
// clues, then assigns an identity match score from 0 to 100. Available on all paid
// plans. Successful requests cost 20 credits. Disposable and free email addresses
// (like gmail.com, yahoo.com) will throw a 422 error.
func (r *PersonService) Enrich(ctx context.Context, body PersonEnrichParams, opts ...option.RequestOption) (res *PersonEnrichResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "people/enrich"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type PersonEnrichResponse struct {
	// The highest-scoring person candidate.
	Match PersonEnrichResponseMatchUnion `json:"match" api:"required"`
	// Metadata about the API key used for the request. Included in every response
	// whenever a valid API key is provided, even when the response status is not 200.
	KeyMetadata PersonEnrichResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Match       respjson.Field
		KeyMetadata respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonEnrichResponse) RawJSON() string { return r.JSON.raw }
func (r *PersonEnrichResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PersonEnrichResponseMatchUnion contains all possible properties and values from
// [PersonEnrichResponseMatchCandidate], [PersonEnrichResponseMatchNotFound].
//
// Use the [PersonEnrichResponseMatchUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PersonEnrichResponseMatchUnion struct {
	// This field is a union of [PersonEnrichResponseMatchCandidatePerson], [any]
	Person PersonEnrichResponseMatchUnionPerson `json:"person"`
	// This field is a union of [int64], [any]
	Score PersonEnrichResponseMatchUnionScore `json:"score"`
	// Any of "candidate", "not_found".
	Status string `json:"status"`
	JSON   struct {
		Person respjson.Field
		Score  respjson.Field
		Status respjson.Field
		raw    string
	} `json:"-"`
}

// anyPersonEnrichResponseMatch is implemented by each variant of
// [PersonEnrichResponseMatchUnion] to add type safety for the return type of
// [PersonEnrichResponseMatchUnion.AsAny]
type anyPersonEnrichResponseMatch interface {
	implPersonEnrichResponseMatchUnion()
}

func (PersonEnrichResponseMatchCandidate) implPersonEnrichResponseMatchUnion() {}
func (PersonEnrichResponseMatchNotFound) implPersonEnrichResponseMatchUnion()  {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PersonEnrichResponseMatchUnion.AsAny().(type) {
//	case contextdev.PersonEnrichResponseMatchCandidate:
//	case contextdev.PersonEnrichResponseMatchNotFound:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PersonEnrichResponseMatchUnion) AsAny() anyPersonEnrichResponseMatch {
	switch u.Status {
	case "candidate":
		return u.AsCandidate()
	case "not_found":
		return u.AsNotFound()
	}
	return nil
}

func (u PersonEnrichResponseMatchUnion) AsCandidate() (v PersonEnrichResponseMatchCandidate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PersonEnrichResponseMatchUnion) AsNotFound() (v PersonEnrichResponseMatchNotFound) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PersonEnrichResponseMatchUnion) RawJSON() string { return u.JSON.raw }

func (r *PersonEnrichResponseMatchUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PersonEnrichResponseMatchUnionPerson is an implicit subunion of
// [PersonEnrichResponseMatchUnion]. PersonEnrichResponseMatchUnionPerson provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PersonEnrichResponseMatchUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPersonEnrichResponseMatchNotFoundPerson]
type PersonEnrichResponseMatchUnionPerson struct {
	// This field will be present if the value is a [any] instead of an object.
	OfPersonEnrichResponseMatchNotFoundPerson any `json:",inline"`
	// This field is from variant [PersonEnrichResponseMatchCandidatePerson].
	CurrentRoleStatus string `json:"current_role_status"`
	// This field is from variant [PersonEnrichResponseMatchCandidatePerson].
	Education []PersonEnrichResponseMatchCandidatePersonEducation `json:"education"`
	// This field is from variant [PersonEnrichResponseMatchCandidatePerson].
	Experience []PersonEnrichResponseMatchCandidatePersonExperience `json:"experience"`
	// This field is from variant [PersonEnrichResponseMatchCandidatePerson].
	Skills []string `json:"skills"`
	// This field is from variant [PersonEnrichResponseMatchCandidatePerson].
	SocialURLs []string `json:"social_urls"`
	// This field is from variant [PersonEnrichResponseMatchCandidatePerson].
	WebsiteURLs []string `json:"website_urls"`
	// This field is from variant [PersonEnrichResponseMatchCandidatePerson].
	AvatarURL string `json:"avatar_url"`
	// This field is from variant [PersonEnrichResponseMatchCandidatePerson].
	Bio string `json:"bio"`
	// This field is from variant [PersonEnrichResponseMatchCandidatePerson].
	CheckedAt string `json:"checked_at"`
	// This field is from variant [PersonEnrichResponseMatchCandidatePerson].
	CurrentRole PersonEnrichResponseMatchCandidatePersonCurrentRole `json:"current_role"`
	// This field is from variant [PersonEnrichResponseMatchCandidatePerson].
	Email string `json:"email"`
	// This field is from variant [PersonEnrichResponseMatchCandidatePerson].
	LastUpdated string `json:"last_updated"`
	// This field is from variant [PersonEnrichResponseMatchCandidatePerson].
	Location PersonEnrichResponseMatchCandidatePersonLocation `json:"location"`
	// This field is from variant [PersonEnrichResponseMatchCandidatePerson].
	Name PersonEnrichResponseMatchCandidatePersonName `json:"name"`
	JSON struct {
		OfPersonEnrichResponseMatchNotFoundPerson respjson.Field
		CurrentRoleStatus                         respjson.Field
		Education                                 respjson.Field
		Experience                                respjson.Field
		Skills                                    respjson.Field
		SocialURLs                                respjson.Field
		WebsiteURLs                               respjson.Field
		AvatarURL                                 respjson.Field
		Bio                                       respjson.Field
		CheckedAt                                 respjson.Field
		CurrentRole                               respjson.Field
		Email                                     respjson.Field
		LastUpdated                               respjson.Field
		Location                                  respjson.Field
		Name                                      respjson.Field
		raw                                       string
	} `json:"-"`
}

func (r *PersonEnrichResponseMatchUnionPerson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PersonEnrichResponseMatchUnionScore is an implicit subunion of
// [PersonEnrichResponseMatchUnion]. PersonEnrichResponseMatchUnionScore provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PersonEnrichResponseMatchUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfPersonEnrichResponseMatchNotFoundScore]
type PersonEnrichResponseMatchUnionScore struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfPersonEnrichResponseMatchNotFoundScore any `json:",inline"`
	JSON                                     struct {
		OfInt                                    respjson.Field
		OfPersonEnrichResponseMatchNotFoundScore respjson.Field
		raw                                      string
	} `json:"-"`
}

func (r *PersonEnrichResponseMatchUnionScore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The highest-scoring person candidate.
type PersonEnrichResponseMatchCandidate struct {
	Person PersonEnrichResponseMatchCandidatePerson `json:"person" api:"required"`
	Score  int64                                    `json:"score" api:"required"`
	Status constant.Candidate                       `json:"status" default:"candidate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Person      respjson.Field
		Score       respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonEnrichResponseMatchCandidate) RawJSON() string { return r.JSON.raw }
func (r *PersonEnrichResponseMatchCandidate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonEnrichResponseMatchCandidatePerson struct {
	// Whether the person's current role is known. `present` — current_role is
	// populated. `none` — the work history explicitly shows every role has ended.
	// `unknown` — our data sources could not confirm either way; treat a missing
	// current_role as unverified rather than vacant.
	//
	// Any of "present", "none", "unknown".
	CurrentRoleStatus string                                               `json:"current_role_status" api:"required"`
	Education         []PersonEnrichResponseMatchCandidatePersonEducation  `json:"education" api:"required"`
	Experience        []PersonEnrichResponseMatchCandidatePersonExperience `json:"experience" api:"required"`
	Skills            []string                                             `json:"skills" api:"required"`
	SocialURLs        []string                                             `json:"social_urls" api:"required" format:"uri"`
	WebsiteURLs       []string                                             `json:"website_urls" api:"required" format:"uri"`
	AvatarURL         string                                               `json:"avatar_url"`
	Bio               string                                               `json:"bio"`
	// When we last refreshed this profile from our data sources (ISO 8601).
	CheckedAt   string                                              `json:"checked_at"`
	CurrentRole PersonEnrichResponseMatchCandidatePersonCurrentRole `json:"current_role"`
	Email       string                                              `json:"email" format:"email"`
	// When the underlying profile data last changed in our data sources (ISO 8601).
	// Omitted when unknown.
	LastUpdated string                                           `json:"last_updated"`
	Location    PersonEnrichResponseMatchCandidatePersonLocation `json:"location"`
	Name        PersonEnrichResponseMatchCandidatePersonName     `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrentRoleStatus respjson.Field
		Education         respjson.Field
		Experience        respjson.Field
		Skills            respjson.Field
		SocialURLs        respjson.Field
		WebsiteURLs       respjson.Field
		AvatarURL         respjson.Field
		Bio               respjson.Field
		CheckedAt         respjson.Field
		CurrentRole       respjson.Field
		Email             respjson.Field
		LastUpdated       respjson.Field
		Location          respjson.Field
		Name              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonEnrichResponseMatchCandidatePerson) RawJSON() string { return r.JSON.raw }
func (r *PersonEnrichResponseMatchCandidatePerson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonEnrichResponseMatchCandidatePersonEducation struct {
	Institution  PersonEnrichResponseMatchCandidatePersonEducationInstitution `json:"institution" api:"required"`
	Degree       string                                                       `json:"degree"`
	Description  string                                                       `json:"description"`
	EndDate      PersonEnrichResponseMatchCandidatePersonEducationEndDate     `json:"end_date"`
	FieldOfStudy string                                                       `json:"field_of_study"`
	StartDate    PersonEnrichResponseMatchCandidatePersonEducationStartDate   `json:"start_date"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Institution  respjson.Field
		Degree       respjson.Field
		Description  respjson.Field
		EndDate      respjson.Field
		FieldOfStudy respjson.Field
		StartDate    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonEnrichResponseMatchCandidatePersonEducation) RawJSON() string { return r.JSON.raw }
func (r *PersonEnrichResponseMatchCandidatePersonEducation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonEnrichResponseMatchCandidatePersonEducationInstitution struct {
	Name   string `json:"name" api:"required"`
	Domain string `json:"domain"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Domain      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonEnrichResponseMatchCandidatePersonEducationInstitution) RawJSON() string {
	return r.JSON.raw
}
func (r *PersonEnrichResponseMatchCandidatePersonEducationInstitution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonEnrichResponseMatchCandidatePersonEducationEndDate struct {
	Year  int64 `json:"year" api:"required"`
	Day   int64 `json:"day"`
	Month int64 `json:"month"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Year        respjson.Field
		Day         respjson.Field
		Month       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonEnrichResponseMatchCandidatePersonEducationEndDate) RawJSON() string { return r.JSON.raw }
func (r *PersonEnrichResponseMatchCandidatePersonEducationEndDate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonEnrichResponseMatchCandidatePersonEducationStartDate struct {
	Year  int64 `json:"year" api:"required"`
	Day   int64 `json:"day"`
	Month int64 `json:"month"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Year        respjson.Field
		Day         respjson.Field
		Month       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonEnrichResponseMatchCandidatePersonEducationStartDate) RawJSON() string {
	return r.JSON.raw
}
func (r *PersonEnrichResponseMatchCandidatePersonEducationStartDate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonEnrichResponseMatchCandidatePersonExperience struct {
	Organization PersonEnrichResponseMatchCandidatePersonExperienceOrganization `json:"organization" api:"required"`
	Title        string                                                         `json:"title" api:"required"`
	Description  string                                                         `json:"description"`
	EndDate      PersonEnrichResponseMatchCandidatePersonExperienceEndDate      `json:"end_date"`
	IsCurrent    bool                                                           `json:"is_current"`
	Location     string                                                         `json:"location"`
	StartDate    PersonEnrichResponseMatchCandidatePersonExperienceStartDate    `json:"start_date"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Organization respjson.Field
		Title        respjson.Field
		Description  respjson.Field
		EndDate      respjson.Field
		IsCurrent    respjson.Field
		Location     respjson.Field
		StartDate    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonEnrichResponseMatchCandidatePersonExperience) RawJSON() string { return r.JSON.raw }
func (r *PersonEnrichResponseMatchCandidatePersonExperience) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonEnrichResponseMatchCandidatePersonExperienceOrganization struct {
	Name   string `json:"name" api:"required"`
	Domain string `json:"domain"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Domain      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonEnrichResponseMatchCandidatePersonExperienceOrganization) RawJSON() string {
	return r.JSON.raw
}
func (r *PersonEnrichResponseMatchCandidatePersonExperienceOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonEnrichResponseMatchCandidatePersonExperienceEndDate struct {
	Year  int64 `json:"year" api:"required"`
	Day   int64 `json:"day"`
	Month int64 `json:"month"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Year        respjson.Field
		Day         respjson.Field
		Month       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonEnrichResponseMatchCandidatePersonExperienceEndDate) RawJSON() string {
	return r.JSON.raw
}
func (r *PersonEnrichResponseMatchCandidatePersonExperienceEndDate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonEnrichResponseMatchCandidatePersonExperienceStartDate struct {
	Year  int64 `json:"year" api:"required"`
	Day   int64 `json:"day"`
	Month int64 `json:"month"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Year        respjson.Field
		Day         respjson.Field
		Month       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonEnrichResponseMatchCandidatePersonExperienceStartDate) RawJSON() string {
	return r.JSON.raw
}
func (r *PersonEnrichResponseMatchCandidatePersonExperienceStartDate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonEnrichResponseMatchCandidatePersonCurrentRole struct {
	Organization PersonEnrichResponseMatchCandidatePersonCurrentRoleOrganization `json:"organization" api:"required"`
	Title        string                                                          `json:"title" api:"required"`
	Description  string                                                          `json:"description"`
	EndDate      PersonEnrichResponseMatchCandidatePersonCurrentRoleEndDate      `json:"end_date"`
	IsCurrent    bool                                                            `json:"is_current"`
	Location     string                                                          `json:"location"`
	StartDate    PersonEnrichResponseMatchCandidatePersonCurrentRoleStartDate    `json:"start_date"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Organization respjson.Field
		Title        respjson.Field
		Description  respjson.Field
		EndDate      respjson.Field
		IsCurrent    respjson.Field
		Location     respjson.Field
		StartDate    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonEnrichResponseMatchCandidatePersonCurrentRole) RawJSON() string { return r.JSON.raw }
func (r *PersonEnrichResponseMatchCandidatePersonCurrentRole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonEnrichResponseMatchCandidatePersonCurrentRoleOrganization struct {
	Name   string `json:"name" api:"required"`
	Domain string `json:"domain"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Domain      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonEnrichResponseMatchCandidatePersonCurrentRoleOrganization) RawJSON() string {
	return r.JSON.raw
}
func (r *PersonEnrichResponseMatchCandidatePersonCurrentRoleOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonEnrichResponseMatchCandidatePersonCurrentRoleEndDate struct {
	Year  int64 `json:"year" api:"required"`
	Day   int64 `json:"day"`
	Month int64 `json:"month"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Year        respjson.Field
		Day         respjson.Field
		Month       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonEnrichResponseMatchCandidatePersonCurrentRoleEndDate) RawJSON() string {
	return r.JSON.raw
}
func (r *PersonEnrichResponseMatchCandidatePersonCurrentRoleEndDate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonEnrichResponseMatchCandidatePersonCurrentRoleStartDate struct {
	Year  int64 `json:"year" api:"required"`
	Day   int64 `json:"day"`
	Month int64 `json:"month"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Year        respjson.Field
		Day         respjson.Field
		Month       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonEnrichResponseMatchCandidatePersonCurrentRoleStartDate) RawJSON() string {
	return r.JSON.raw
}
func (r *PersonEnrichResponseMatchCandidatePersonCurrentRoleStartDate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonEnrichResponseMatchCandidatePersonLocation struct {
	City        string `json:"city"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	Display     string `json:"display"`
	Region      string `json:"region"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		CountryCode respjson.Field
		Display     respjson.Field
		Region      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonEnrichResponseMatchCandidatePersonLocation) RawJSON() string { return r.JSON.raw }
func (r *PersonEnrichResponseMatchCandidatePersonLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonEnrichResponseMatchCandidatePersonName struct {
	First string `json:"first"`
	Full  string `json:"full"`
	Last  string `json:"last"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		First       respjson.Field
		Full        respjson.Field
		Last        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonEnrichResponseMatchCandidatePersonName) RawJSON() string { return r.JSON.raw }
func (r *PersonEnrichResponseMatchCandidatePersonName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// No usable person candidate was found.
type PersonEnrichResponseMatchNotFound struct {
	Person any               `json:"person" api:"required"`
	Score  any               `json:"score" api:"required"`
	Status constant.NotFound `json:"status" default:"not_found"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Person      respjson.Field
		Score       respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonEnrichResponseMatchNotFound) RawJSON() string { return r.JSON.raw }
func (r *PersonEnrichResponseMatchNotFound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the API key used for the request. Included in every response
// whenever a valid API key is provided, even when the response status is not 200.
type PersonEnrichResponseKeyMetadata struct {
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
func (r PersonEnrichResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *PersonEnrichResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonEnrichParams struct {
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs  param.Opt[int64]              `json:"timeoutMS,omitzero"`
	Company    PersonEnrichParamsCompany     `json:"company,omitzero"`
	Education  []PersonEnrichParamsEducation `json:"education,omitzero"`
	Location   PersonEnrichParamsLocation    `json:"location,omitzero"`
	Name       PersonEnrichParamsName        `json:"name,omitzero"`
	SocialURLs []string                      `json:"social_urls,omitzero" format:"uri"`
	// Optional tags for tracking usage. Up to 20 tags, each 1 to 50 characters.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r PersonEnrichParams) MarshalJSON() (data []byte, err error) {
	type shadow PersonEnrichParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonEnrichParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonEnrichParamsCompany struct {
	Domain param.Opt[string] `json:"domain,omitzero"`
	Name   param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r PersonEnrichParamsCompany) MarshalJSON() (data []byte, err error) {
	type shadow PersonEnrichParamsCompany
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonEnrichParamsCompany) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonEnrichParamsEducation struct {
	Degree         param.Opt[string]                      `json:"degree,omitzero"`
	FieldOfStudy   param.Opt[string]                      `json:"field_of_study,omitzero"`
	GraduationYear param.Opt[int64]                       `json:"graduation_year,omitzero"`
	Institution    PersonEnrichParamsEducationInstitution `json:"institution,omitzero"`
	paramObj
}

func (r PersonEnrichParamsEducation) MarshalJSON() (data []byte, err error) {
	type shadow PersonEnrichParamsEducation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonEnrichParamsEducation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonEnrichParamsEducationInstitution struct {
	Domain param.Opt[string] `json:"domain,omitzero"`
	Name   param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r PersonEnrichParamsEducationInstitution) MarshalJSON() (data []byte, err error) {
	type shadow PersonEnrichParamsEducationInstitution
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonEnrichParamsEducationInstitution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonEnrichParamsLocation struct {
	City    param.Opt[string] `json:"city,omitzero"`
	Country param.Opt[string] `json:"country,omitzero"`
	Region  param.Opt[string] `json:"region,omitzero"`
	paramObj
}

func (r PersonEnrichParamsLocation) MarshalJSON() (data []byte, err error) {
	type shadow PersonEnrichParamsLocation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonEnrichParamsLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonEnrichParamsName struct {
	First param.Opt[string] `json:"first,omitzero"`
	Last  param.Opt[string] `json:"last,omitzero"`
	paramObj
}

func (r PersonEnrichParamsName) MarshalJSON() (data []byte, err error) {
	type shadow PersonEnrichParamsName
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonEnrichParamsName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
