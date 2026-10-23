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

// Find a person from identity clues and return their profile with a match score.
// Available on all plans; free or disposable email addresses return 422.
func (r *PersonService) Enrich(ctx context.Context, body PersonEnrichParams, opts ...option.RequestOption) (res *PersonEnrichResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "people/enrich"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Find people with a natural-language query.
func (r *PersonService) Search(ctx context.Context, body PersonSearchParams, opts ...option.RequestOption) (res *PersonSearchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "people/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type PersonEnrichResponse struct {
	// The highest-scoring person candidate.
	Match PersonEnrichResponseMatchUnion `json:"match" api:"required"`
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Credits this request used and your remaining balance.
	KeyMetadata PersonEnrichResponseKeyMetadata `json:"key_metadata"`
	// True when the timeout ended processing and this response contains the usable
	// data completed so far. Unfinished fields are omitted.
	Partial bool `json:"partial"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Match       respjson.Field
		RequestID   respjson.Field
		KeyMetadata respjson.Field
		Partial     respjson.Field
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

// Credits this request used and your remaining balance.
type PersonEnrichResponseKeyMetadata struct {
	// Credits charged for this request.
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
func (r PersonEnrichResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *PersonEnrichResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonSearchResponse struct {
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Up to 10 matching people, most relevant first.
	Results []PersonSearchResponseResult `json:"results" api:"required"`
	// Credits this request used and your remaining balance.
	KeyMetadata PersonSearchResponseKeyMetadata `json:"key_metadata"`
	// Pass as cursor to fetch the next page; null when there are no more results.
	NextCursor string `json:"next_cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestID   respjson.Field
		Results     respjson.Field
		KeyMetadata respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *PersonSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A person matching your search; unavailable details are null or empty lists.
type PersonSearchResponseResult struct {
	// What the person does.
	Description string `json:"description" api:"required"`
	// Whether the person is currently a founder or co-founder.
	IsFounder bool `json:"is_founder" api:"required"`
	// Where the person lives.
	Location PersonSearchResponseResultLocation `json:"location" api:"required"`
	// Full name.
	Name string `json:"name" api:"required"`
	// Public professional profile URLs.
	Profiles []string `json:"profiles" api:"required"`
	// The person’s current role.
	Role PersonSearchResponseResultRole `json:"role" api:"required"`
	// Schools attended.
	Schools []string `json:"schools" api:"required"`
	// Professional skills.
	Skills []string `json:"skills" api:"required"`
	// Years of professional experience.
	YearsOfExperience int64 `json:"years_of_experience" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description       respjson.Field
		IsFounder         respjson.Field
		Location          respjson.Field
		Name              respjson.Field
		Profiles          respjson.Field
		Role              respjson.Field
		Schools           respjson.Field
		Skills            respjson.Field
		YearsOfExperience respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonSearchResponseResult) RawJSON() string { return r.JSON.raw }
func (r *PersonSearchResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Where the person lives.
type PersonSearchResponseResultLocation struct {
	// City.
	City string `json:"city" api:"required"`
	// Country as a two-letter code, such as US or GB.
	CountryCode string `json:"country_code" api:"required"`
	// State or region.
	Region string `json:"region" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		CountryCode respjson.Field
		Region      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonSearchResponseResultLocation) RawJSON() string { return r.JSON.raw }
func (r *PersonSearchResponseResultLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The person’s current role.
type PersonSearchResponseResultRole struct {
	// Current employer, when known.
	Company PersonSearchResponseResultRoleCompany `json:"company" api:"required"`
	// Areas of responsibility.
	Functions []string `json:"functions" api:"required"`
	// Seniority level.
	Seniority string `json:"seniority" api:"required"`
	// Role start date (YYYY-MM-DD).
	StartedOn string `json:"started_on" api:"required"`
	// Job title.
	Title string `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Company     respjson.Field
		Functions   respjson.Field
		Seniority   respjson.Field
		StartedOn   respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonSearchResponseResultRole) RawJSON() string { return r.JSON.raw }
func (r *PersonSearchResponseResultRole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current employer, when known.
type PersonSearchResponseResultRoleCompany struct {
	// The company’s website domain.
	Domain string `json:"domain" api:"required"`
	// Company name.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonSearchResponseResultRoleCompany) RawJSON() string { return r.JSON.raw }
func (r *PersonSearchResponseResultRoleCompany) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credits this request used and your remaining balance.
type PersonSearchResponseKeyMetadata struct {
	// Credits charged for this request.
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
func (r PersonSearchResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *PersonSearchResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonEnrichParams struct {
	// Email address of the person to find.
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// Company context to help identify the person. Provide a name or domain.
	Company PersonEnrichParamsCompany `json:"company,omitzero"`
	// Education history to help distinguish people with similar names.
	Education []PersonEnrichParamsEducation `json:"education,omitzero"`
	// Location context to help identify the person. Provide a city, region, or
	// country.
	Location PersonEnrichParamsLocation `json:"location,omitzero"`
	// Person name. Without an email or person-profile URL, provide both first and last
	// name plus company, education, or location.
	Name PersonEnrichParamsName `json:"name,omitzero"`
	// Public profile URLs for the person. A person-profile URL can identify the person
	// without a name.
	SocialURLs []string `json:"social_urls,omitzero" format:"uri"`
	// Labels for filtering usage in the dashboard.
	Tags []string `json:"tags,omitzero"`
	// Request deadline and what to return when it passes.
	TimeoutOpts PersonEnrichParamsTimeoutOpts `json:"timeoutOpts,omitzero"`
	// `enabled` turns on zero data retention. Returns 403 `ZDR_NOT_ENABLED` unless
	// your organization has ZDR.
	//
	// Any of "enabled", "disabled".
	Zdr PersonEnrichParamsZdr `json:"zdr,omitzero"`
	paramObj
}

func (r PersonEnrichParams) MarshalJSON() (data []byte, err error) {
	type shadow PersonEnrichParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonEnrichParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Company context to help identify the person. Provide a name or domain.
type PersonEnrichParamsCompany struct {
	// Website domain of a company associated with the person.
	Domain param.Opt[string] `json:"domain,omitzero"`
	// Name of a company associated with the person.
	Name param.Opt[string] `json:"name,omitzero"`
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
	// Degree or qualification earned.
	Degree param.Opt[string] `json:"degree,omitzero"`
	// Subject or major studied.
	FieldOfStudy param.Opt[string] `json:"field_of_study,omitzero"`
	// Four-digit graduation year.
	GraduationYear param.Opt[int64] `json:"graduation_year,omitzero"`
	// School or university, identified by name or domain.
	Institution PersonEnrichParamsEducationInstitution `json:"institution,omitzero"`
	paramObj
}

func (r PersonEnrichParamsEducation) MarshalJSON() (data []byte, err error) {
	type shadow PersonEnrichParamsEducation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonEnrichParamsEducation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// School or university, identified by name or domain.
type PersonEnrichParamsEducationInstitution struct {
	// Website domain of the school or university.
	Domain param.Opt[string] `json:"domain,omitzero"`
	// Name of the school or university.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r PersonEnrichParamsEducationInstitution) MarshalJSON() (data []byte, err error) {
	type shadow PersonEnrichParamsEducationInstitution
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonEnrichParamsEducationInstitution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Location context to help identify the person. Provide a city, region, or
// country.
type PersonEnrichParamsLocation struct {
	// City associated with the person.
	City param.Opt[string] `json:"city,omitzero"`
	// Country associated with the person.
	Country param.Opt[string] `json:"country,omitzero"`
	// State, province, or region associated with the person.
	Region param.Opt[string] `json:"region,omitzero"`
	paramObj
}

func (r PersonEnrichParamsLocation) MarshalJSON() (data []byte, err error) {
	type shadow PersonEnrichParamsLocation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonEnrichParamsLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Person name. Without an email or person-profile URL, provide both first and last
// name plus company, education, or location.
type PersonEnrichParamsName struct {
	// First or given name.
	First param.Opt[string] `json:"first,omitzero"`
	// Last or family name.
	Last param.Opt[string] `json:"last,omitzero"`
	paramObj
}

func (r PersonEnrichParamsName) MarshalJSON() (data []byte, err error) {
	type shadow PersonEnrichParamsName
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonEnrichParamsName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request deadline and what to return when it passes.
//
// The property Milliseconds is required.
type PersonEnrichParamsTimeoutOpts struct {
	// Deadline in milliseconds.
	Milliseconds int64 `json:"milliseconds" api:"required"`
	// "fail" returns 408 at the deadline. "return-partial" returns available results;
	// inspect the response’s partial flag.
	//
	// Any of "fail", "return-partial".
	Behavior string `json:"behavior,omitzero"`
	paramObj
}

func (r PersonEnrichParamsTimeoutOpts) MarshalJSON() (data []byte, err error) {
	type shadow PersonEnrichParamsTimeoutOpts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonEnrichParamsTimeoutOpts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PersonEnrichParamsTimeoutOpts](
		"behavior", "fail", "return-partial",
	)
}

// `enabled` turns on zero data retention. Returns 403 `ZDR_NOT_ENABLED` unless
// your organization has ZDR.
type PersonEnrichParamsZdr string

const (
	PersonEnrichParamsZdrEnabled  PersonEnrichParamsZdr = "enabled"
	PersonEnrichParamsZdrDisabled PersonEnrichParamsZdr = "disabled"
)

type PersonSearchParams struct {
	// Describe the people to find.
	Query string `json:"query" api:"required"`
	// Use next_cursor with the same query and filters to fetch the next page.
	Cursor param.Opt[string] `json:"cursor,omitzero"`
	// Additional constraints every result must match.
	Filters PersonSearchParamsFilters `json:"filters,omitzero"`
	// Labels for filtering usage in the dashboard.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r PersonSearchParams) MarshalJSON() (data []byte, err error) {
	type shadow PersonSearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional constraints every result must match.
type PersonSearchParamsFilters struct {
	// Maximum years of professional experience.
	YearsOfExperienceMax param.Opt[int64] `json:"years_of_experience_max,omitzero"`
	// Minimum years of professional experience.
	YearsOfExperienceMin param.Opt[int64] `json:"years_of_experience_min,omitzero"`
	// Countries people live in as two-letter codes, such as US or GB.
	Countries []string `json:"countries,omitzero"`
	paramObj
}

func (r PersonSearchParamsFilters) MarshalJSON() (data []byte, err error) {
	type shadow PersonSearchParamsFilters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonSearchParamsFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
