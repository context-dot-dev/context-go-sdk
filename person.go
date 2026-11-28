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

// Enrich a person from identity clues, or find people with a natural-language
// query.
//
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

// Match a person from a work email, a profile URL, or a name with company,
// education, or location context, and get their profile with an identity-match
// score from 0 to 100. The best candidate is returned even when the match is weak,
// so set a score threshold for your workflow. `not_found` results are free.
func (r *PersonService) Enrich(ctx context.Context, body PersonEnrichParams, opts ...option.RequestOption) (res *PersonEnrichResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "people/enrich"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Find people from a plain-language description, such as
// `VPs of sales at solar companies in California`. Each result includes the
// current role and employer, location, years of experience, skills, schools, and
// public profile URLs. Pages with no matches are free.
func (r *PersonService) Search(ctx context.Context, body PersonSearchParams, opts ...option.RequestOption) (res *PersonSearchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "people/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type PersonEnrichResponse struct {
	// The best-matching person, with an identity-match score.
	Match PersonEnrichResponseMatchUnion `json:"match" api:"required"`
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Credits this request used and your remaining balance.
	KeyMetadata PersonEnrichResponseKeyMetadata `json:"key_metadata"`
	// `true` when the `timeoutOpts` deadline ended processing early. The response
	// holds the data completed so far; unfinished fields are omitted.
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

// The best-matching person, with an identity-match score.
type PersonEnrichResponseMatchCandidate struct {
	// Profile of the matched person.
	Person PersonEnrichResponseMatchCandidatePerson `json:"person" api:"required"`
	// Identity-match score from 0 to 100. Higher is a stronger match; weak candidates
	// are still returned.
	Score int64 `json:"score" api:"required"`
	// Always `candidate` for a match.
	Status constant.Candidate `json:"status" default:"candidate"`
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

// Profile of the matched person.
type PersonEnrichResponseMatchCandidatePerson struct {
	// Whether the person's current role is known. `present` means `current_role` is
	// filled in; `none` means the work history shows every role has ended; `unknown`
	// means it could not be confirmed either way, so treat a missing `current_role` as
	// unverified, not vacant.
	//
	// Any of "present", "none", "unknown".
	CurrentRoleStatus string `json:"current_role_status" api:"required"`
	// Education history.
	Education []PersonEnrichResponseMatchCandidatePersonEducation `json:"education" api:"required"`
	// Work history. Each entry has the same fields as `current_role`.
	Experience []PersonEnrichResponseMatchCandidatePersonExperience `json:"experience" api:"required"`
	// Professional skills.
	Skills []string `json:"skills" api:"required"`
	// Public profile URLs for the person.
	SocialURLs []string `json:"social_urls" api:"required" format:"uri"`
	// Other websites associated with the person.
	WebsiteURLs []string `json:"website_urls" api:"required" format:"uri"`
	// URL of the hosted profile picture. Omitted with zero data retention.
	AvatarURL string `json:"avatar_url"`
	// Short bio from the person's profile.
	Bio string `json:"bio"`
	// When Context.dev last refreshed the profile from its sources, as an ISO 8601
	// timestamp.
	CheckedAt string `json:"checked_at"`
	// The person's current job. See `current_role_status` when it is missing.
	CurrentRole PersonEnrichResponseMatchCandidatePersonCurrentRole `json:"current_role"`
	// Email address of the person.
	Email string `json:"email" format:"email"`
	// When the profile data last changed at its sources, as an ISO 8601 timestamp.
	// Omitted when unknown.
	LastUpdated string `json:"last_updated"`
	// Where the person is based.
	Location PersonEnrichResponseMatchCandidatePersonLocation `json:"location"`
	// Name of the person.
	Name PersonEnrichResponseMatchCandidatePersonName `json:"name"`
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
	// School or university.
	Institution PersonEnrichResponseMatchCandidatePersonEducationInstitution `json:"institution" api:"required"`
	// Degree or qualification.
	Degree string `json:"degree"`
	// Description of the studies.
	Description string `json:"description"`
	// When the studies ended.
	EndDate PersonEnrichResponseMatchCandidatePersonEducationEndDate `json:"end_date"`
	// Subject or major studied.
	FieldOfStudy string `json:"field_of_study"`
	// When the studies started.
	StartDate PersonEnrichResponseMatchCandidatePersonEducationStartDate `json:"start_date"`
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

// School or university.
type PersonEnrichResponseMatchCandidatePersonEducationInstitution struct {
	// Name of the organization.
	Name string `json:"name" api:"required"`
	// Website domain of the organization. Omitted when unknown.
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

// When the studies ended.
type PersonEnrichResponseMatchCandidatePersonEducationEndDate struct {
	// Four-digit year.
	Year int64 `json:"year" api:"required"`
	// Day of the month. Omitted when unknown.
	Day int64 `json:"day"`
	// Month. Omitted when unknown.
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

// When the studies started.
type PersonEnrichResponseMatchCandidatePersonEducationStartDate struct {
	// Four-digit year.
	Year int64 `json:"year" api:"required"`
	// Day of the month. Omitted when unknown.
	Day int64 `json:"day"`
	// Month. Omitted when unknown.
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
	// Employer for this role.
	Organization PersonEnrichResponseMatchCandidatePersonExperienceOrganization `json:"organization" api:"required"`
	// Job title.
	Title string `json:"title" api:"required"`
	// Description of the role.
	Description string `json:"description"`
	// When the role ended.
	EndDate PersonEnrichResponseMatchCandidatePersonExperienceEndDate `json:"end_date"`
	// Whether the person still holds this role.
	IsCurrent bool `json:"is_current"`
	// Where the role is based.
	Location string `json:"location"`
	// When the role started.
	StartDate PersonEnrichResponseMatchCandidatePersonExperienceStartDate `json:"start_date"`
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

// Employer for this role.
type PersonEnrichResponseMatchCandidatePersonExperienceOrganization struct {
	// Name of the organization.
	Name string `json:"name" api:"required"`
	// Website domain of the organization. Omitted when unknown.
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

// When the role ended.
type PersonEnrichResponseMatchCandidatePersonExperienceEndDate struct {
	// Four-digit year.
	Year int64 `json:"year" api:"required"`
	// Day of the month. Omitted when unknown.
	Day int64 `json:"day"`
	// Month. Omitted when unknown.
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

// When the role started.
type PersonEnrichResponseMatchCandidatePersonExperienceStartDate struct {
	// Four-digit year.
	Year int64 `json:"year" api:"required"`
	// Day of the month. Omitted when unknown.
	Day int64 `json:"day"`
	// Month. Omitted when unknown.
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

// The person's current job. See `current_role_status` when it is missing.
type PersonEnrichResponseMatchCandidatePersonCurrentRole struct {
	// Employer for this role.
	Organization PersonEnrichResponseMatchCandidatePersonCurrentRoleOrganization `json:"organization" api:"required"`
	// Job title.
	Title string `json:"title" api:"required"`
	// Description of the role.
	Description string `json:"description"`
	// When the role ended.
	EndDate PersonEnrichResponseMatchCandidatePersonCurrentRoleEndDate `json:"end_date"`
	// Whether the person still holds this role.
	IsCurrent bool `json:"is_current"`
	// Where the role is based.
	Location string `json:"location"`
	// When the role started.
	StartDate PersonEnrichResponseMatchCandidatePersonCurrentRoleStartDate `json:"start_date"`
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

// Employer for this role.
type PersonEnrichResponseMatchCandidatePersonCurrentRoleOrganization struct {
	// Name of the organization.
	Name string `json:"name" api:"required"`
	// Website domain of the organization. Omitted when unknown.
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

// When the role ended.
type PersonEnrichResponseMatchCandidatePersonCurrentRoleEndDate struct {
	// Four-digit year.
	Year int64 `json:"year" api:"required"`
	// Day of the month. Omitted when unknown.
	Day int64 `json:"day"`
	// Month. Omitted when unknown.
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

// When the role started.
type PersonEnrichResponseMatchCandidatePersonCurrentRoleStartDate struct {
	// Four-digit year.
	Year int64 `json:"year" api:"required"`
	// Day of the month. Omitted when unknown.
	Day int64 `json:"day"`
	// Month. Omitted when unknown.
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

// Where the person is based.
type PersonEnrichResponseMatchCandidatePersonLocation struct {
	// City.
	City string `json:"city"`
	// Country name.
	Country string `json:"country"`
	// ISO 3166-1 alpha-2 country code.
	CountryCode string `json:"country_code"`
	// Location as a single display string.
	Display string `json:"display"`
	// State, province, or region.
	Region string `json:"region"`
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

// Name of the person.
type PersonEnrichResponseMatchCandidatePersonName struct {
	// First or given name.
	First string `json:"first"`
	// Full name.
	Full string `json:"full"`
	// Last or family name.
	Last string `json:"last"`
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

// No usable candidate was found.
type PersonEnrichResponseMatchNotFound struct {
	// Always `null`.
	Person any `json:"person" api:"required"`
	// Always `null`.
	Score any `json:"score" api:"required"`
	// Always `not_found` when no candidate exists.
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
	// Pass as `cursor` to fetch the next page. `null` on the last page.
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

// A person who matches your search. Unknown details are `null` or empty lists.
type PersonSearchResponseResult struct {
	// What the person does.
	Description string `json:"description" api:"required"`
	// Whether the person is currently a founder or co-founder. `null` when unknown.
	IsFounder bool `json:"is_founder" api:"required"`
	// Where the person lives.
	Location PersonSearchResponseResultLocation `json:"location" api:"required"`
	// Full name.
	Name string `json:"name" api:"required"`
	// Public professional profile URLs. Send one to `POST /people/enrich` as
	// `social_urls` for the full profile.
	Profiles []string `json:"profiles" api:"required"`
	// The person's current role.
	Role PersonSearchResponseResultRole `json:"role" api:"required"`
	// Schools attended.
	Schools []string `json:"schools" api:"required"`
	// Professional skills.
	Skills []string `json:"skills" api:"required"`
	// Years of professional experience. `null` when unknown.
	YearsOfExperience int64 `json:"years_of_experience" api:"required"`
	// How well the person matches the query. Returned only in `ultra` mode.
	Match PersonSearchResponseResultMatch `json:"match"`
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
		Match             respjson.Field
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
	// ISO 3166-1 alpha-2 country code, such as `US`.
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

// The person's current role.
type PersonSearchResponseResultRole struct {
	// Current employer. `null` when unknown.
	Company PersonSearchResponseResultRoleCompany `json:"company" api:"required"`
	// Areas of responsibility, such as `sales`.
	Functions []string `json:"functions" api:"required"`
	// Seniority level, such as `executive`.
	Seniority string `json:"seniority" api:"required"`
	// When the role started, as `YYYY-MM-DD`.
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

// Current employer. `null` when unknown.
type PersonSearchResponseResultRoleCompany struct {
	// Website domain of the company.
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

// How well the person matches the query. Returned only in `ultra` mode.
type PersonSearchResponseResultMatch struct {
	// How well the person matches the query, from 0 to 1.
	Confidence float64 `json:"confidence" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Confidence  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonSearchResponseResultMatch) RawJSON() string { return r.JSON.raw }
func (r *PersonSearchResponseResultMatch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credits this request used and your remaining balance.
type PersonSearchResponseKeyMetadata struct {
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
func (r PersonSearchResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *PersonSearchResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonEnrichParams struct {
	// Work email address of the person. Free and disposable email providers
	// return 422.
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// Company associated with the person. Send a `name`, a `domain`, or both.
	Company PersonEnrichParamsCompany `json:"company,omitzero"`
	// Education history to help distinguish people with similar names.
	Education []PersonEnrichParamsEducation `json:"education,omitzero"`
	// Location associated with the person. Send at least one of `city`, `region`, or
	// `country`.
	Location PersonEnrichParamsLocation `json:"location,omitzero"`
	// Name of the person. Without an email or profile URL, send both `first` and
	// `last` plus `company`, `education`, or `location`.
	Name PersonEnrichParamsName `json:"name,omitzero"`
	// Public profile URLs for the person. A person-profile URL is enough to identify
	// them without a name.
	SocialURLs []string `json:"social_urls,omitzero" format:"uri"`
	// Labels for this request. Filter usage and request logs by them in the dashboard
	// or with `GET /org/logs`.
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

// Company associated with the person. Send a `name`, a `domain`, or both.
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

// Location associated with the person. Send at least one of `city`, `region`, or
// `country`.
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

// Name of the person. Without an email or profile URL, send both `first` and
// `last` plus `company`, `education`, or `location`.
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
	// `fail` returns 408 at the deadline. `return-partial` returns the results
	// collected so far and marks the response as partial.
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
	// Plain-language description of the people to find. Up to 500 characters in `fast`
	// mode and 5000 in `ultra` mode.
	Query string `json:"query" api:"required"`
	// Cursor from the previous page's `next_cursor`. Omit for the first page. Send the
	// same `query`, `mode`, and `filters` with it.
	Cursor param.Opt[string] `json:"cursor,omitzero"`
	// Constraints every result must match, in addition to `query`.
	Filters PersonSearchParamsFilters `json:"filters,omitzero"`
	// `fast` matches people on basic criteria such as role, location, or background;
	// `ultra` runs a deeper search, checks every candidate against each criterion in
	// your query, and adds a `match` score. Complex `ultra` searches can take up to 30
	// seconds, so set your client timeout to match.
	//
	// Any of "fast", "ultra".
	Mode PersonSearchParamsMode `json:"mode,omitzero"`
	// Labels for this request. Filter usage and request logs by them in the dashboard
	// or with `GET /org/logs`.
	Tags []string `json:"tags,omitzero"`
	// `enabled` turns on zero data retention. Returns 403 `ZDR_NOT_ENABLED` unless
	// your organization has ZDR.
	//
	// Any of "enabled", "disabled".
	Zdr PersonSearchParamsZdr `json:"zdr,omitzero"`
	paramObj
}

func (r PersonSearchParams) MarshalJSON() (data []byte, err error) {
	type shadow PersonSearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constraints every result must match, in addition to `query`.
type PersonSearchParamsFilters struct {
	// Maximum years of professional experience.
	YearsOfExperienceMax param.Opt[int64] `json:"years_of_experience_max,omitzero"`
	// Minimum years of professional experience.
	YearsOfExperienceMin param.Opt[int64] `json:"years_of_experience_min,omitzero"`
	// Only return people who live in these countries, as ISO 3166-1 alpha-2 codes such
	// as `US`.
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

// `fast` matches people on basic criteria such as role, location, or background;
// `ultra` runs a deeper search, checks every candidate against each criterion in
// your query, and adds a `match` score. Complex `ultra` searches can take up to 30
// seconds, so set your client timeout to match.
type PersonSearchParamsMode string

const (
	PersonSearchParamsModeFast  PersonSearchParamsMode = "fast"
	PersonSearchParamsModeUltra PersonSearchParamsMode = "ultra"
)

// `enabled` turns on zero data retention. Returns 403 `ZDR_NOT_ENABLED` unless
// your organization has ZDR.
type PersonSearchParamsZdr string

const (
	PersonSearchParamsZdrEnabled  PersonSearchParamsZdr = "enabled"
	PersonSearchParamsZdrDisabled PersonSearchParamsZdr = "disabled"
)
