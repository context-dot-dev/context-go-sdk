// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contextdev

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/context-dot-dev/context-go-sdk/v2/internal/apijson"
	"github.com/context-dot-dev/context-go-sdk/v2/internal/apiquery"
	"github.com/context-dot-dev/context-go-sdk/v2/internal/requestconfig"
	"github.com/context-dot-dev/context-go-sdk/v2/option"
	"github.com/context-dot-dev/context-go-sdk/v2/packages/param"
	"github.com/context-dot-dev/context-go-sdk/v2/packages/respjson"
	"github.com/context-dot-dev/context-go-sdk/v2/shared/constant"
)

// Retrieve a brand's logos, colors, company details, and social links, or find
// brands by name or domain.
//
// BrandService contains methods and other services that help with interacting with
// the context.dev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBrandService] method instead.
type BrandService struct {
	options []option.RequestOption
}

// NewBrandService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBrandService(opts ...option.RequestOption) (r BrandService) {
	r = BrandService{}
	r.options = opts
	return
}

// Look up a brand by domain, company name, work email, stock ticker, page URL, or
// card transaction descriptor. Returns logos, colors, backdrops, description,
// industries, address, stock listing, and social links. A `by_direct_url` lookup
// reads only that page.
func (r *BrandService) Get(ctx context.Context, body BrandGetParams, opts ...option.RequestOption) (res *BrandGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "brand/retrieve"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Autocomplete brand names and domains as a user types, with up to 10 brands
// ranked by popularity and a ready-to-use logo URL for each. Filter by country and
// primary language. Pass a returned `domain` to `POST /brand/retrieve` for the
// full brand profile.
func (r *BrandService) Search(ctx context.Context, query BrandSearchParams, opts ...option.RequestOption) (res *BrandSearchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "brand/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type BrandGetResponse struct {
	// The brand's identity, visual assets, and company details.
	Brand BrandGetResponseBrand `json:"brand" api:"required"`
	// Cache status of this response.
	CacheMetadata BrandGetResponseCacheMetadata `json:"cache_metadata" api:"required"`
	// HTTP status code. Always `200` on success.
	Code int64 `json:"code" api:"required"`
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Always `ok` on success.
	Status string `json:"status" api:"required"`
	// Credits this request used and your remaining balance.
	KeyMetadata BrandGetResponseKeyMetadata `json:"key_metadata"`
	// `true` when the `timeoutOpts` deadline ended processing early. The response
	// holds the data completed so far; unfinished fields are omitted.
	Partial bool `json:"partial"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Brand         respjson.Field
		CacheMetadata respjson.Field
		Code          respjson.Field
		RequestID     respjson.Field
		Status        respjson.Field
		KeyMetadata   respjson.Field
		Partial       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BrandGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The brand's identity, visual assets, and company details.
type BrandGetResponseBrand struct {
	// Physical address of the brand. Omitted when unknown.
	Address BrandGetResponseBrandAddress `json:"address"`
	// Backdrop images for the brand, such as social profile banners.
	Backdrops []BrandGetResponseBrandBackdrop `json:"backdrops"`
	// Brand colors from the website and logos.
	Colors []BrandGetResponseBrandColor `json:"colors"`
	// Short description of the brand.
	Description string `json:"description"`
	// Domain of the brand.
	Domain string `json:"domain"`
	// Contact email address for the company.
	Email string `json:"email"`
	// Employee headcount. Omitted when unknown.
	Employees BrandGetResponseBrandEmployees `json:"employees"`
	// Industry classification of the brand.
	Industries BrandGetResponseBrandIndustries `json:"industries"`
	// Whether the brand's content is not safe for work (NSFW).
	IsNsfw bool `json:"is_nsfw"`
	// Key pages on the brand's website.
	Links BrandGetResponseBrandLinks `json:"links"`
	// Brand logos, most relevant first. An SVG is listed ahead of a raster logo of the
	// same shape.
	Logos []BrandGetResponseBrandLogo `json:"logos"`
	// Contact phone number for the company.
	Phone string `json:"phone"`
	// Lowercase language name, such as `english`.
	//
	// Any of "afrikaans", "albanian", "amharic", "arabic", "armenian", "assamese",
	// "aymara", "azeri", "basque", "belarusian", "bengali", "bosnian", "bulgarian",
	// "burmese", "cantonese", "catalan", "cebuano", "chinese", "corsican", "croatian",
	// "czech", "danish", "dutch", "english", "esperanto", "estonian", "farsi",
	// "fijian", "finnish", "french", "galician", "georgian", "german", "greek",
	// "guarani", "gujarati", "haitian-creole", "hausa", "hawaiian", "hebrew", "hindi",
	// "hmong", "hungarian", "icelandic", "igbo", "indonesian", "irish", "italian",
	// "japanese", "javanese", "kannada", "kazakh", "khmer", "kinyarwanda", "korean",
	// "kurdish", "kyrgyz", "lao", "latin", "latvian", "lingala", "lithuanian",
	// "luxembourgish", "macedonian", "malagasy", "malay", "malayalam", "maltese",
	// "maori", "marathi", "mongolian", "nepali", "norwegian", "odia", "oromo",
	// "pashto", "pidgin", "polish", "portuguese", "punjabi", "quechua", "romanian",
	// "russian", "samoan", "scottish-gaelic", "serbian", "sesotho", "shona", "sindhi",
	// "sinhala", "slovak", "slovene", "somali", "spanish", "sundanese", "swahili",
	// "swedish", "tagalog", "tajik", "tamil", "tatar", "telugu", "thai", "tibetan",
	// "tigrinya", "tongan", "tswana", "turkish", "turkmen", "ukrainian", "urdu",
	// "uyghur", "uzbek", "vietnamese", "welsh", "wolof", "xhosa", "yiddish", "yoruba",
	// "zulu".
	PrimaryLanguage string `json:"primary_language" api:"nullable"`
	// Brand slogan or tagline.
	Slogan string `json:"slogan"`
	// Social media profiles of the brand.
	Socials []BrandGetResponseBrandSocial `json:"socials"`
	// Stock listing. Omitted when the company is not publicly traded.
	Stock BrandGetResponseBrandStock `json:"stock"`
	// Brand name.
	Title string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address         respjson.Field
		Backdrops       respjson.Field
		Colors          respjson.Field
		Description     respjson.Field
		Domain          respjson.Field
		Email           respjson.Field
		Employees       respjson.Field
		Industries      respjson.Field
		IsNsfw          respjson.Field
		Links           respjson.Field
		Logos           respjson.Field
		Phone           respjson.Field
		PrimaryLanguage respjson.Field
		Slogan          respjson.Field
		Socials         respjson.Field
		Stock           respjson.Field
		Title           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandGetResponseBrand) RawJSON() string { return r.JSON.raw }
func (r *BrandGetResponseBrand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Physical address of the brand. Omitted when unknown.
type BrandGetResponseBrandAddress struct {
	// City.
	City string `json:"city"`
	// Country name.
	Country string `json:"country"`
	// Country code, such as `US`.
	CountryCode string `json:"country_code"`
	// Postal or ZIP code.
	PostalCode string `json:"postal_code"`
	// State or province code.
	StateCode string `json:"state_code"`
	// State or province name.
	StateProvince string `json:"state_province"`
	// Street address.
	Street string `json:"street"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City          respjson.Field
		Country       respjson.Field
		CountryCode   respjson.Field
		PostalCode    respjson.Field
		StateCode     respjson.Field
		StateProvince respjson.Field
		Street        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandGetResponseBrandAddress) RawJSON() string { return r.JSON.raw }
func (r *BrandGetResponseBrandAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetResponseBrandBackdrop struct {
	// Colors in the backdrop image.
	Colors []BrandGetResponseBrandBackdropColor `json:"colors"`
	// Size of the backdrop image.
	Resolution BrandGetResponseBrandBackdropResolution `json:"resolution"`
	// URL of the backdrop image.
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Colors      respjson.Field
		Resolution  respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandGetResponseBrandBackdrop) RawJSON() string { return r.JSON.raw }
func (r *BrandGetResponseBrandBackdrop) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetResponseBrandBackdropColor struct {
	// Hex color code, such as `#635BFF`.
	Hex string `json:"hex"`
	// Human-readable color name.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Hex         respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandGetResponseBrandBackdropColor) RawJSON() string { return r.JSON.raw }
func (r *BrandGetResponseBrandBackdropColor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Size of the backdrop image.
type BrandGetResponseBrandBackdropResolution struct {
	// Aspect ratio, as width divided by height.
	AspectRatio float64 `json:"aspect_ratio"`
	// Height, in pixels.
	Height int64 `json:"height"`
	// Width, in pixels.
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AspectRatio respjson.Field
		Height      respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandGetResponseBrandBackdropResolution) RawJSON() string { return r.JSON.raw }
func (r *BrandGetResponseBrandBackdropResolution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetResponseBrandColor struct {
	// Hex color code, such as `#635BFF`.
	Hex string `json:"hex"`
	// Human-readable color name.
	Name string `json:"name"`
	// `site` colors come from the website's own theme (rendered page colors, web app
	// manifest, `theme-color` meta tag); `logo` colors come from logo pixels.
	//
	// Any of "site", "logo".
	Source string `json:"source"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Hex         respjson.Field
		Name        respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandGetResponseBrandColor) RawJSON() string { return r.JSON.raw }
func (r *BrandGetResponseBrandColor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Employee headcount. Omitted when unknown.
type BrandGetResponseBrandEmployees struct {
	// Exact employee count, when known.
	Exact int64 `json:"exact"`
	// Headcount range, such as `11 to 50`.
	//
	// Any of "1 to 10", "11 to 50", "51 to 200", "201 to 500", "501 to 1000", "1001 to
	// 5000", "5001 to 10000", "10001+".
	Range string `json:"range"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Exact       respjson.Field
		Range       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandGetResponseBrandEmployees) RawJSON() string { return r.JSON.raw }
func (r *BrandGetResponseBrandEmployees) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Industry classification of the brand.
type BrandGetResponseBrandIndustries struct {
	// Industry and subindustry pairs from Context.dev's Easy Industry Classification
	// (EIC).
	Eic []BrandGetResponseBrandIndustriesEic `json:"eic"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Eic         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandGetResponseBrandIndustries) RawJSON() string { return r.JSON.raw }
func (r *BrandGetResponseBrandIndustries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetResponseBrandIndustriesEic struct {
	// Top-level industry in the Easy Industry Classification (EIC).
	//
	// Any of "Aerospace & Defense", "Technology", "Finance", "Healthcare", "Retail &
	// E-commerce", "Entertainment", "Education", "Government & Nonprofit", "Industrial
	// & Energy", "Automotive & Transportation", "Lifestyle & Leisure", "Luxury &
	// Fashion", "News & Media", "Sports", "Real Estate & PropTech", "Legal &
	// Compliance", "Telecommunications", "Agriculture & Food", "Professional Services
	// & Agencies", "Chemicals & Materials", "Logistics & Supply Chain", "Hospitality &
	// Tourism", "Construction & Built Environment", "Consumer Packaged Goods (CPG)".
	Industry string `json:"industry" api:"required"`
	// Subindustry in the Easy Industry Classification (EIC).
	//
	// Any of "Defense Systems & Military Hardware", "Aerospace Manufacturing",
	// "Avionics & Navigation Technology", "Subsea & Naval Defense Systems", "Space &
	// Satellite Technology", "Defense IT & Systems Integration", "Software (B2B)",
	// "Software (B2C)", "Cloud Infrastructure & DevOps", "Cybersecurity", "Artificial
	// Intelligence & Machine Learning", "Data Infrastructure & Analytics", "Hardware &
	// Semiconductors", "Fintech Infrastructure", "eCommerce & Marketplace Platforms",
	// "Developer Tools & APIs", "Web3 & Blockchain", "XR & Spatial Computing",
	// "Banking & Lending", "Investment Management & WealthTech", "Insurance &
	// InsurTech", "Payments & Money Movement", "Accounting, Tax & Financial Planning
	// Tools", "Capital Markets & Trading Platforms", "Financial Infrastructure &
	// APIs", "Credit Scoring & Risk Management", "Cryptocurrency & Digital Assets",
	// "BNPL & Alternative Financing", "Healthcare Providers & Services",
	// "Pharmaceuticals & Drug Development", "Medical Devices & Diagnostics",
	// "Biotechnology & Genomics", "Digital Health & Telemedicine", "Health Insurance &
	// Benefits Tech", "Clinical Trials & Research Platforms", "Mental Health &
	// Wellness", "Healthcare IT & EHR Systems", "Consumer Health & Wellness Products",
	// "Online Marketplaces", "Direct-to-Consumer (DTC) Brands", "Retail Tech &
	// Point-of-Sale Systems", "Omnichannel & In-Store Retail", "E-commerce Enablement
	// & Infrastructure", "Subscription & Membership Commerce", "Social Commerce &
	// Influencer Platforms", "Fashion & Apparel Retail", "Food, Beverage & Grocery
	// E-commerce", "Streaming Platforms (Video, Music, Audio)", "Gaming & Interactive
	// Entertainment", "Creator Economy & Influencer Platforms", "Film, TV & Production
	// Studios", "Events, Venues & Live Entertainment", "Virtual Worlds & Metaverse
	// Experiences", "K-12 Education Platforms & Tools", "Higher Education & University
	// Tech", "Online Learning & MOOCs", "Test Prep & Certification", "Corporate
	// Training & Upskilling", "Tutoring & Supplemental Learning", "Education
	// Management Systems (LMS/SIS)", "Language Learning", "Creator-Led & Cohort-Based
	// Courses", "Special Education & Accessibility Tools", "Government Technology &
	// Digital Services", "Civic Engagement & Policy Platforms", "International
	// Development & Humanitarian Aid", "Philanthropy & Grantmaking", "Nonprofit
	// Operations & Fundraising Tools", "Public Health & Social Services", "Education &
	// Youth Development Programs", "Environmental & Climate Action Organizations",
	// "Legal Aid & Social Justice Advocacy", "Municipal & Infrastructure Services",
	// "Manufacturing & Industrial Automation", "Energy Production (Oil, Gas,
	// Nuclear)", "Renewable Energy & Cleantech", "Utilities & Grid Infrastructure",
	// "Industrial IoT & Monitoring Systems", "Construction & Heavy Equipment", "Mining
	// & Natural Resources", "Environmental Engineering & Sustainability", "Energy
	// Storage & Battery Technology", "Automotive OEMs & Vehicle Manufacturing",
	// "Electric Vehicles (EVs) & Charging Infrastructure", "Mobility-as-a-Service
	// (MaaS)", "Fleet Management", "Public Transit & Urban Mobility", "Autonomous
	// Vehicles & ADAS", "Aftermarket Parts & Services", "Telematics & Vehicle
	// Connectivity", "Aviation & Aerospace Transport", "Maritime Shipping", "Fitness &
	// Wellness", "Beauty & Personal Care", "Home & Living", "Dating & Relationships",
	// "Hobbies, Crafts & DIY", "Outdoor & Recreational Gear", "Events, Experiences &
	// Ticketing Platforms", "Designer & Luxury Apparel", "Accessories, Jewelry &
	// Watches", "Footwear & Leather Goods", "Beauty, Fragrance & Skincare", "Fashion
	// Marketplaces & Retail Platforms", "Sustainable & Ethical Fashion", "Resale,
	// Vintage & Circular Fashion", "Fashion Tech & Virtual Try-Ons", "Streetwear &
	// Emerging Luxury", "Couture & Made-to-Measure", "News Publishing & Journalism",
	// "Advertising, Adtech & Media Buying", "Digital Media & Content Platforms",
	// "Broadcasting (TV & Radio)", "Podcasting & Audio Media", "News Aggregators &
	// Curation Tools", "Independent & Creator-Led Media", "Newsletters &
	// Substack-Style Platforms", "Political & Investigative Media", "Trade & Niche
	// Publications", "Media Monitoring & Analytics", "Professional Teams & Leagues",
	// "Sports Media & Broadcasting", "Sports Betting & Fantasy Sports", "Fitness &
	// Athletic Training Platforms", "Sportswear & Equipment", "Esports & Competitive
	// Gaming", "Sports Venues & Event Management", "Athlete Management & Talent
	// Agencies", "Sports Tech & Performance Analytics", "Youth, Amateur & Collegiate
	// Sports", "Real Estate Marketplaces", "Property Management Software", "Rental
	// Platforms", "Mortgage & Lending Tech", "Real Estate Investment Platforms", "Law
	// Firms & Legal Services", "Legal Tech & Automation", "Regulatory Compliance",
	// "E-Discovery & Litigation Tools", "Contract Management", "Governance, Risk &
	// Compliance (GRC)", "IP & Trademark Management", "Legal Research & Intelligence",
	// "Compliance Training & Certification", "Whistleblower & Ethics Reporting",
	// "Mobile & Wireless Networks (3G/4G/5G)", "Broadband & Fiber Internet",
	// "Satellite & Space-Based Communications", "Network Equipment & Infrastructure",
	// "Telecom Billing & OSS/BSS Systems", "VoIP & Unified Communications", "Internet
	// Service Providers (ISPs)", "Edge Computing & Network Virtualization", "IoT
	// Connectivity Platforms", "Precision Agriculture & AgTech", "Crop & Livestock
	// Production", "Food & Beverage Manufacturing & Processing", "Food Distribution",
	// "Restaurants & Food Service", "Agricultural Inputs & Equipment", "Sustainable &
	// Regenerative Agriculture", "Seafood & Aquaculture", "Management Consulting",
	// "Marketing & Advertising Agencies", "Design, Branding & Creative Studios", "IT
	// Services & Managed Services", "Staffing, Recruiting & Talent", "Accounting & Tax
	// Firms", "Public Relations & Communications", "Business Process Outsourcing
	// (BPO)", "Professional Training & Coaching", "Specialty Chemicals", "Commodity &
	// Petrochemicals", "Polymers, Plastics & Rubber", "Coatings, Adhesives &
	// Sealants", "Industrial Gases", "Advanced Materials & Composites", "Battery
	// Materials & Energy Storage", "Electronic Materials & Semiconductor Chemicals",
	// "Agrochemicals & Fertilizers", "Freight & Transportation Tech", "Last-Mile
	// Delivery", "Warehouse Automation", "Supply Chain Visibility Platforms",
	// "Logistics Marketplaces", "Shipping & Freight Forwarding", "Cold Chain
	// Logistics", "Reverse Logistics & Returns", "Cross-Border Trade Tech",
	// "Transportation Management Systems (TMS)", "Hotels & Accommodation", "Vacation
	// Rentals & Short-Term Stays", "Restaurant Tech & Management", "Travel Booking
	// Platforms", "Tourism Experiences & Activities", "Cruise Lines & Marine Tourism",
	// "Hospitality Management Systems", "Event & Venue Management", "Corporate Travel
	// Management", "Travel Insurance & Protection", "Construction Management
	// Software", "BIM/CAD & Design Tools", "Construction Marketplaces", "Equipment
	// Rental & Management", "Building Materials & Procurement", "Construction
	// Workforce Management", "Project Estimation & Bidding", "Modular & Prefab
	// Construction", "Construction Safety & Compliance", "Smart Building Technology",
	// "Food & Beverage CPG", "Home & Personal Care CPG", "CPG Analytics & Insights",
	// "Direct-to-Consumer CPG Brands", "CPG Supply Chain & Distribution", "Private
	// Label Manufacturing", "CPG Retail Intelligence", "Sustainable CPG & Packaging",
	// "Beauty & Cosmetics CPG", "Health & Wellness CPG".
	Subindustry string `json:"subindustry" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Industry    respjson.Field
		Subindustry respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandGetResponseBrandIndustriesEic) RawJSON() string { return r.JSON.raw }
func (r *BrandGetResponseBrandIndustriesEic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Key pages on the brand's website.
type BrandGetResponseBrandLinks struct {
	// URL of the blog or newsroom. `null` when not found.
	Blog string `json:"blog" api:"nullable"`
	// URL of the careers page. `null` when not found.
	Careers string `json:"careers" api:"nullable"`
	// URL of the contact page. `null` when not found.
	Contact string `json:"contact" api:"nullable"`
	// URL of the pricing page. `null` when not found.
	Pricing string `json:"pricing" api:"nullable"`
	// URL of the privacy policy. `null` when not found.
	Privacy string `json:"privacy" api:"nullable"`
	// URL of the terms of service. `null` when not found.
	Terms string `json:"terms" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blog        respjson.Field
		Careers     respjson.Field
		Contact     respjson.Field
		Pricing     respjson.Field
		Privacy     respjson.Field
		Terms       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandGetResponseBrandLinks) RawJSON() string { return r.JSON.raw }
func (r *BrandGetResponseBrandLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetResponseBrandLogo struct {
	// Colors in the logo.
	Colors []BrandGetResponseBrandLogoColor `json:"colors"`
	// `light` works best in light mode, `dark` works best in dark mode, and
	// `has_opaque_background` works in either because the image has its own
	// background.
	//
	// Any of "light", "dark", "has_opaque_background".
	Mode string `json:"mode"`
	// Size of the logo image.
	Resolution BrandGetResponseBrandLogoResolution `json:"resolution"`
	// Whether the image is an `icon` or a full `logo`, based on its dimensions.
	//
	// Any of "icon", "logo".
	Type string `json:"type"`
	// URL of the hosted logo image.
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Colors      respjson.Field
		Mode        respjson.Field
		Resolution  respjson.Field
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandGetResponseBrandLogo) RawJSON() string { return r.JSON.raw }
func (r *BrandGetResponseBrandLogo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetResponseBrandLogoColor struct {
	// Hex color code, such as `#635BFF`.
	Hex string `json:"hex"`
	// Human-readable color name.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Hex         respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandGetResponseBrandLogoColor) RawJSON() string { return r.JSON.raw }
func (r *BrandGetResponseBrandLogoColor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Size of the logo image.
type BrandGetResponseBrandLogoResolution struct {
	// Aspect ratio, as width divided by height.
	AspectRatio float64 `json:"aspect_ratio"`
	// Height, in pixels.
	Height int64 `json:"height"`
	// Width, in pixels.
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AspectRatio respjson.Field
		Height      respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandGetResponseBrandLogoResolution) RawJSON() string { return r.JSON.raw }
func (r *BrandGetResponseBrandLogoResolution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetResponseBrandSocial struct {
	// Social platform.
	//
	// Any of "x", "facebook", "instagram", "linkedin", "youtube", "pinterest",
	// "tiktok", "dribbble", "github", "behance", "snapchat", "whatsapp", "telegram",
	// "line", "discord", "twitch", "vimeo", "imdb", "tumblr", "flickr", "giphy",
	// "medium", "spotify", "soundcloud", "tripadvisor", "yelp", "producthunt",
	// "reddit", "crunchbase", "appstore", "playstore".
	Type string `json:"type"`
	// URL of the brand's profile on that platform.
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandGetResponseBrandSocial) RawJSON() string { return r.JSON.raw }
func (r *BrandGetResponseBrandSocial) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stock listing. Omitted when the company is not publicly traded.
type BrandGetResponseBrandStock struct {
	// Exchange the stock trades on.
	Exchange string `json:"exchange"`
	// Stock ticker symbol.
	Ticker string `json:"ticker"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Exchange    respjson.Field
		Ticker      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandGetResponseBrandStock) RawJSON() string { return r.JSON.raw }
func (r *BrandGetResponseBrandStock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache status of this response.
type BrandGetResponseCacheMetadata struct {
	// Age of the cached data, in milliseconds. `0` for `miss` and `zdr`.
	AgeMs int64 `json:"age_ms" api:"required"`
	// `hit` when served from cache, `miss` when fetched fresh, `zdr` when zero data
	// retention bypassed the cache.
	//
	// Any of "hit", "miss", "zdr".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgeMs       respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandGetResponseCacheMetadata) RawJSON() string { return r.JSON.raw }
func (r *BrandGetResponseCacheMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credits this request used and your remaining balance.
type BrandGetResponseKeyMetadata struct {
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
func (r BrandGetResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *BrandGetResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSearchResponse struct {
	// Unique ID of this request, also in `X-Request-Id`. Include it when contacting
	// support.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Up to 10 matching brands: name matches first, then domain matches, each group
	// ordered by popularity. Empty when nothing matches.
	Results []BrandSearchResponseResult `json:"results" api:"required"`
	// Credits this request used and your remaining balance.
	KeyMetadata BrandSearchResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestID   respjson.Field
		Results     respjson.Field
		KeyMetadata respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *BrandSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSearchResponseResult struct {
	// Domain of the brand.
	Domain string `json:"domain" api:"required"`
	// Logo Link URL that serves the brand's logo with your organization's client ID.
	// Use it directly as an image source.
	Logo string `json:"logo" api:"required"`
	// Name of the brand. Empty string when unknown.
	Name string `json:"name" api:"required"`
	// Country of the brand, as an uppercase ISO 3166-1 alpha-2 code. Omitted when
	// unknown.
	//
	// Any of "AD", "AE", "AF", "AG", "AI", "AL", "AM", "AO", "AQ", "AR", "AS", "AT",
	// "AU", "AW", "AX", "AZ", "BA", "BB", "BD", "BE", "BF", "BG", "BH", "BI", "BJ",
	// "BL", "BM", "BN", "BO", "BQ", "BR", "BS", "BT", "BV", "BW", "BY", "BZ", "CA",
	// "CC", "CD", "CF", "CG", "CH", "CI", "CK", "CL", "CM", "CN", "CO", "CR", "CU",
	// "CV", "CW", "CX", "CY", "CZ", "DE", "DJ", "DK", "DM", "DO", "DZ", "EC", "EE",
	// "EG", "EH", "ER", "ES", "ET", "FI", "FJ", "FK", "FM", "FO", "FR", "GA", "GB",
	// "GD", "GE", "GF", "GG", "GH", "GI", "GL", "GM", "GN", "GP", "GQ", "GR", "GS",
	// "GT", "GU", "GW", "GY", "HK", "HM", "HN", "HR", "HT", "HU", "ID", "IE", "IL",
	// "IM", "IN", "IO", "IQ", "IR", "IS", "IT", "JE", "JM", "JO", "JP", "KE", "KG",
	// "KH", "KI", "KM", "KN", "KP", "KR", "KW", "KY", "KZ", "LA", "LB", "LC", "LI",
	// "LK", "LR", "LS", "LT", "LU", "LV", "LY", "MA", "MC", "MD", "ME", "MF", "MG",
	// "MH", "MK", "ML", "MM", "MN", "MO", "MP", "MQ", "MR", "MS", "MT", "MU", "MV",
	// "MW", "MX", "MY", "MZ", "NA", "NC", "NE", "NF", "NG", "NI", "NL", "NO", "NP",
	// "NR", "NU", "NZ", "OM", "PA", "PE", "PF", "PG", "PH", "PK", "PL", "PM", "PN",
	// "PR", "PS", "PT", "PW", "PY", "QA", "RE", "RO", "RS", "RU", "RW", "SA", "SB",
	// "SC", "SD", "SE", "SG", "SH", "SI", "SJ", "SK", "SL", "SM", "SN", "SO", "SR",
	// "SS", "ST", "SV", "SX", "SY", "SZ", "TC", "TD", "TF", "TG", "TH", "TJ", "TK",
	// "TL", "TM", "TN", "TO", "TR", "TT", "TV", "TW", "TZ", "UA", "UG", "UM", "US",
	// "UY", "UZ", "VA", "VC", "VE", "VG", "VI", "VN", "VU", "WF", "WS", "YE", "YT",
	// "ZA", "ZM", "ZW".
	Country string `json:"country"`
	// Short description of the brand. Omitted when unknown.
	Description string `json:"description"`
	// Primary language of the brand. Omitted when unknown.
	//
	// Any of "afrikaans", "albanian", "amharic", "arabic", "armenian", "assamese",
	// "aymara", "azeri", "basque", "belarusian", "bengali", "bosnian", "bulgarian",
	// "burmese", "cantonese", "catalan", "cebuano", "chinese", "corsican", "croatian",
	// "czech", "danish", "dutch", "english", "esperanto", "estonian", "farsi",
	// "fijian", "finnish", "french", "galician", "georgian", "german", "greek",
	// "guarani", "gujarati", "haitian-creole", "hausa", "hawaiian", "hebrew", "hindi",
	// "hmong", "hungarian", "icelandic", "igbo", "indonesian", "irish", "italian",
	// "japanese", "javanese", "kannada", "kazakh", "khmer", "kinyarwanda", "korean",
	// "kurdish", "kyrgyz", "lao", "latin", "latvian", "lingala", "lithuanian",
	// "luxembourgish", "macedonian", "malagasy", "malay", "malayalam", "maltese",
	// "maori", "marathi", "mongolian", "nepali", "norwegian", "odia", "oromo",
	// "pashto", "pidgin", "polish", "portuguese", "punjabi", "quechua", "romanian",
	// "russian", "samoan", "scottish-gaelic", "serbian", "sesotho", "shona", "sindhi",
	// "sinhala", "slovak", "slovene", "somali", "spanish", "sundanese", "swahili",
	// "swedish", "tagalog", "tajik", "tamil", "tatar", "telugu", "thai", "tibetan",
	// "tigrinya", "tongan", "tswana", "turkish", "turkmen", "ukrainian", "urdu",
	// "uyghur", "uzbek", "vietnamese", "welsh", "wolof", "xhosa", "yiddish", "yoruba",
	// "zulu".
	Language string `json:"language"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		Logo        respjson.Field
		Name        respjson.Field
		Country     respjson.Field
		Description respjson.Field
		Language    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandSearchResponseResult) RawJSON() string { return r.JSON.raw }
func (r *BrandSearchResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credits this request used and your remaining balance.
type BrandSearchResponseKeyMetadata struct {
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
func (r BrandSearchResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *BrandSearchResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set. Look up
	// a brand by its domain.
	OfByDomain *BrandGetParamsBodyByDomain `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Look up
	// a brand by company name.
	OfByName *BrandGetParamsBodyByName `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Look up
	// a brand from a work email address. Free and disposable email providers
	// return 422.
	OfByEmail *BrandGetParamsBodyByEmail `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Look up
	// a publicly traded company by stock ticker.
	OfByTicker *BrandGetParamsBodyByTicker `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Extract
	// brand data from a single page. Only that URL is read, and no data from other
	// sources is added.
	OfByDirectURL *BrandGetParamsBodyByDirectURL `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Identify the merchant behind a card or bank transaction descriptor.
	OfByTransaction *BrandGetParamsBodyByTransaction `json:",inline"`

	paramObj
}

func (u BrandGetParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfByDomain,
		u.OfByName,
		u.OfByEmail,
		u.OfByTicker,
		u.OfByDirectURL,
		u.OfByTransaction)
}
func (r *BrandGetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Look up a brand by its domain.
//
// The properties Domain, Type are required.
type BrandGetParamsBodyByDomain struct {
	// Domain of the brand, such as `stripe.com`.
	Domain string `json:"domain" api:"required"`
	// Maximum age of cached brand data, in milliseconds. `0` always fetches fresh.
	// Defaults to 90 days (7776000000 ms). Values above 1 year (31536000000 ms) are
	// capped at 1 year.
	MaxAgeMs param.Opt[int64] `json:"maxAgeMs,omitzero"`
	// Skip slower enrichment steps for a faster, less detailed response.
	MaxSpeed param.Opt[bool] `json:"maxSpeed,omitzero"`
	// Translate the brand's `title`, `description`, and `slogan` into this language.
	//
	// Any of "afrikaans", "albanian", "amharic", "arabic", "armenian", "assamese",
	// "aymara", "azeri", "basque", "belarusian", "bengali", "bosnian", "bulgarian",
	// "burmese", "cantonese", "catalan", "cebuano", "chinese", "corsican", "croatian",
	// "czech", "danish", "dutch", "english", "esperanto", "estonian", "farsi",
	// "fijian", "finnish", "french", "galician", "georgian", "german", "greek",
	// "guarani", "gujarati", "haitian-creole", "hausa", "hawaiian", "hebrew", "hindi",
	// "hmong", "hungarian", "icelandic", "igbo", "indonesian", "irish", "italian",
	// "japanese", "javanese", "kannada", "kazakh", "khmer", "kinyarwanda", "korean",
	// "kurdish", "kyrgyz", "lao", "latin", "latvian", "lingala", "lithuanian",
	// "luxembourgish", "macedonian", "malagasy", "malay", "malayalam", "maltese",
	// "maori", "marathi", "mongolian", "nepali", "norwegian", "odia", "oromo",
	// "pashto", "pidgin", "polish", "portuguese", "punjabi", "quechua", "romanian",
	// "russian", "samoan", "scottish-gaelic", "serbian", "sesotho", "shona", "sindhi",
	// "sinhala", "slovak", "slovene", "somali", "spanish", "sundanese", "swahili",
	// "swedish", "tagalog", "tajik", "tamil", "tatar", "telugu", "thai", "tibetan",
	// "tigrinya", "tongan", "tswana", "turkish", "turkmen", "ukrainian", "urdu",
	// "uyghur", "uzbek", "vietnamese", "welsh", "wolof", "xhosa", "yiddish", "yoruba",
	// "zulu".
	ForceLanguage string `json:"force_language,omitzero"`
	// Access policy IDs to apply; each URL must pass all of them. Requires access
	// policies enabled for your organization; otherwise returns
	// `ACCESS_POLICIES_NOT_ENABLED`.
	Policies []string `json:"policies,omitzero"`
	// Labels for this request. Filter usage and request logs by them in the dashboard
	// or with `GET /org/logs`.
	Tags []string `json:"tags,omitzero"`
	// Request deadline and what to return when it passes.
	TimeoutOpts BrandGetParamsBodyByDomainTimeoutOpts `json:"timeoutOpts,omitzero"`
	// `enabled` turns off request logs and traces and keeps request and response
	// content out of usage logs. Brand data and hosted assets may still be stored.
	// Returns 403 `ZDR_NOT_ENABLED` unless your organization has ZDR.
	//
	// Any of "enabled", "disabled".
	Zdr string `json:"zdr,omitzero"`
	// Set to `by_domain` to look up a brand by domain.
	//
	// This field can be elided, and will marshal its zero value as "by_domain".
	Type constant.ByDomain `json:"type" default:"by_domain"`
	paramObj
}

func (r BrandGetParamsBodyByDomain) MarshalJSON() (data []byte, err error) {
	type shadow BrandGetParamsBodyByDomain
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandGetParamsBodyByDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BrandGetParamsBodyByDomain](
		"force_language", "afrikaans", "albanian", "amharic", "arabic", "armenian", "assamese", "aymara", "azeri", "basque", "belarusian", "bengali", "bosnian", "bulgarian", "burmese", "cantonese", "catalan", "cebuano", "chinese", "corsican", "croatian", "czech", "danish", "dutch", "english", "esperanto", "estonian", "farsi", "fijian", "finnish", "french", "galician", "georgian", "german", "greek", "guarani", "gujarati", "haitian-creole", "hausa", "hawaiian", "hebrew", "hindi", "hmong", "hungarian", "icelandic", "igbo", "indonesian", "irish", "italian", "japanese", "javanese", "kannada", "kazakh", "khmer", "kinyarwanda", "korean", "kurdish", "kyrgyz", "lao", "latin", "latvian", "lingala", "lithuanian", "luxembourgish", "macedonian", "malagasy", "malay", "malayalam", "maltese", "maori", "marathi", "mongolian", "nepali", "norwegian", "odia", "oromo", "pashto", "pidgin", "polish", "portuguese", "punjabi", "quechua", "romanian", "russian", "samoan", "scottish-gaelic", "serbian", "sesotho", "shona", "sindhi", "sinhala", "slovak", "slovene", "somali", "spanish", "sundanese", "swahili", "swedish", "tagalog", "tajik", "tamil", "tatar", "telugu", "thai", "tibetan", "tigrinya", "tongan", "tswana", "turkish", "turkmen", "ukrainian", "urdu", "uyghur", "uzbek", "vietnamese", "welsh", "wolof", "xhosa", "yiddish", "yoruba", "zulu",
	)
	apijson.RegisterFieldValidator[BrandGetParamsBodyByDomain](
		"zdr", "enabled", "disabled",
	)
}

// Request deadline and what to return when it passes.
//
// The property Milliseconds is required.
type BrandGetParamsBodyByDomainTimeoutOpts struct {
	// Deadline in milliseconds.
	Milliseconds int64 `json:"milliseconds" api:"required"`
	// `fail` returns 408 at the deadline. `return-partial` returns the results
	// collected so far and marks the response as partial.
	//
	// Any of "fail", "return-partial".
	Behavior string `json:"behavior,omitzero"`
	paramObj
}

func (r BrandGetParamsBodyByDomainTimeoutOpts) MarshalJSON() (data []byte, err error) {
	type shadow BrandGetParamsBodyByDomainTimeoutOpts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandGetParamsBodyByDomainTimeoutOpts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BrandGetParamsBodyByDomainTimeoutOpts](
		"behavior", "fail", "return-partial",
	)
}

// Look up a brand by company name.
//
// The properties Name, Type are required.
type BrandGetParamsBodyByName struct {
	// Company name, such as `Apple Inc`.
	Name string `json:"name" api:"required"`
	// Maximum age of cached brand data, in milliseconds. `0` always fetches fresh.
	// Defaults to 90 days (7776000000 ms). Values above 1 year (31536000000 ms) are
	// capped at 1 year.
	MaxAgeMs param.Opt[int64] `json:"maxAgeMs,omitzero"`
	// Skip slower enrichment steps for a faster, less detailed response.
	MaxSpeed param.Opt[bool] `json:"maxSpeed,omitzero"`
	// Translate the brand's `title`, `description`, and `slogan` into this language.
	//
	// Any of "afrikaans", "albanian", "amharic", "arabic", "armenian", "assamese",
	// "aymara", "azeri", "basque", "belarusian", "bengali", "bosnian", "bulgarian",
	// "burmese", "cantonese", "catalan", "cebuano", "chinese", "corsican", "croatian",
	// "czech", "danish", "dutch", "english", "esperanto", "estonian", "farsi",
	// "fijian", "finnish", "french", "galician", "georgian", "german", "greek",
	// "guarani", "gujarati", "haitian-creole", "hausa", "hawaiian", "hebrew", "hindi",
	// "hmong", "hungarian", "icelandic", "igbo", "indonesian", "irish", "italian",
	// "japanese", "javanese", "kannada", "kazakh", "khmer", "kinyarwanda", "korean",
	// "kurdish", "kyrgyz", "lao", "latin", "latvian", "lingala", "lithuanian",
	// "luxembourgish", "macedonian", "malagasy", "malay", "malayalam", "maltese",
	// "maori", "marathi", "mongolian", "nepali", "norwegian", "odia", "oromo",
	// "pashto", "pidgin", "polish", "portuguese", "punjabi", "quechua", "romanian",
	// "russian", "samoan", "scottish-gaelic", "serbian", "sesotho", "shona", "sindhi",
	// "sinhala", "slovak", "slovene", "somali", "spanish", "sundanese", "swahili",
	// "swedish", "tagalog", "tajik", "tamil", "tatar", "telugu", "thai", "tibetan",
	// "tigrinya", "tongan", "tswana", "turkish", "turkmen", "ukrainian", "urdu",
	// "uyghur", "uzbek", "vietnamese", "welsh", "wolof", "xhosa", "yiddish", "yoruba",
	// "zulu".
	ForceLanguage string `json:"force_language,omitzero"`
	// Lowercase ISO 3166-1 alpha-2 country code, such as `us`.
	//
	// Any of "af", "al", "dz", "as", "ad", "ao", "ai", "aq", "ag", "ar", "am", "aw",
	// "au", "at", "az", "bs", "bh", "bd", "bb", "by", "be", "bz", "bj", "bm", "bt",
	// "bo", "ba", "bw", "bv", "br", "io", "bn", "bg", "bf", "bi", "kh", "cm", "ca",
	// "cv", "ky", "cf", "td", "cl", "cn", "cx", "cc", "co", "km", "cg", "cd", "ck",
	// "cr", "ci", "hr", "cu", "cy", "cz", "dk", "dj", "dm", "do", "ec", "eg", "sv",
	// "gq", "er", "ee", "et", "fk", "fo", "fj", "fi", "fr", "gf", "pf", "tf", "ga",
	// "gm", "ge", "de", "gh", "gi", "gr", "gl", "gd", "gp", "gu", "gt", "gn", "gw",
	// "gy", "ht", "hm", "va", "hn", "hk", "hu", "is", "in", "id", "ir", "iq", "ie",
	// "il", "it", "jm", "jp", "jo", "kz", "ke", "ki", "kp", "kr", "kw", "kg", "la",
	// "lv", "lb", "ls", "lr", "ly", "li", "lt", "lu", "mo", "mk", "mg", "mw", "my",
	// "mv", "ml", "mt", "mh", "mq", "mr", "mu", "yt", "mx", "fm", "md", "mc", "mn",
	// "ms", "ma", "mz", "mm", "na", "nr", "np", "nl", "an", "nc", "nz", "ni", "ne",
	// "ng", "nu", "nf", "mp", "no", "om", "pk", "pw", "ps", "pa", "pg", "py", "pe",
	// "ph", "pn", "pl", "pt", "pr", "qa", "re", "ro", "ru", "rw", "sh", "kn", "lc",
	// "pm", "vc", "ws", "sm", "st", "sa", "sn", "rs", "sc", "sl", "sg", "sk", "si",
	// "sb", "so", "za", "gs", "es", "lk", "sd", "sr", "sj", "sz", "se", "ch", "sy",
	// "tw", "tj", "tz", "th", "tl", "tg", "tk", "to", "tt", "tn", "tr", "tm", "tc",
	// "tv", "ug", "ua", "ae", "gb", "us", "um", "uy", "uz", "vu", "ve", "vn", "vg",
	// "vi", "wf", "eh", "ye", "zm", "zw".
	Country string `json:"country,omitzero"`
	// Access policy IDs to apply; each URL must pass all of them. Requires access
	// policies enabled for your organization; otherwise returns
	// `ACCESS_POLICIES_NOT_ENABLED`.
	Policies []string `json:"policies,omitzero"`
	// Labels for this request. Filter usage and request logs by them in the dashboard
	// or with `GET /org/logs`.
	Tags []string `json:"tags,omitzero"`
	// Request deadline and what to return when it passes.
	TimeoutOpts BrandGetParamsBodyByNameTimeoutOpts `json:"timeoutOpts,omitzero"`
	// `enabled` turns off request logs and traces and keeps request and response
	// content out of usage logs. Brand data and hosted assets may still be stored.
	// Returns 403 `ZDR_NOT_ENABLED` unless your organization has ZDR.
	//
	// Any of "enabled", "disabled".
	Zdr string `json:"zdr,omitzero"`
	// Set to `by_name` to look up a brand by company name.
	//
	// This field can be elided, and will marshal its zero value as "by_name".
	Type constant.ByName `json:"type" default:"by_name"`
	paramObj
}

func (r BrandGetParamsBodyByName) MarshalJSON() (data []byte, err error) {
	type shadow BrandGetParamsBodyByName
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandGetParamsBodyByName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BrandGetParamsBodyByName](
		"country", "af", "al", "dz", "as", "ad", "ao", "ai", "aq", "ag", "ar", "am", "aw", "au", "at", "az", "bs", "bh", "bd", "bb", "by", "be", "bz", "bj", "bm", "bt", "bo", "ba", "bw", "bv", "br", "io", "bn", "bg", "bf", "bi", "kh", "cm", "ca", "cv", "ky", "cf", "td", "cl", "cn", "cx", "cc", "co", "km", "cg", "cd", "ck", "cr", "ci", "hr", "cu", "cy", "cz", "dk", "dj", "dm", "do", "ec", "eg", "sv", "gq", "er", "ee", "et", "fk", "fo", "fj", "fi", "fr", "gf", "pf", "tf", "ga", "gm", "ge", "de", "gh", "gi", "gr", "gl", "gd", "gp", "gu", "gt", "gn", "gw", "gy", "ht", "hm", "va", "hn", "hk", "hu", "is", "in", "id", "ir", "iq", "ie", "il", "it", "jm", "jp", "jo", "kz", "ke", "ki", "kp", "kr", "kw", "kg", "la", "lv", "lb", "ls", "lr", "ly", "li", "lt", "lu", "mo", "mk", "mg", "mw", "my", "mv", "ml", "mt", "mh", "mq", "mr", "mu", "yt", "mx", "fm", "md", "mc", "mn", "ms", "ma", "mz", "mm", "na", "nr", "np", "nl", "an", "nc", "nz", "ni", "ne", "ng", "nu", "nf", "mp", "no", "om", "pk", "pw", "ps", "pa", "pg", "py", "pe", "ph", "pn", "pl", "pt", "pr", "qa", "re", "ro", "ru", "rw", "sh", "kn", "lc", "pm", "vc", "ws", "sm", "st", "sa", "sn", "rs", "sc", "sl", "sg", "sk", "si", "sb", "so", "za", "gs", "es", "lk", "sd", "sr", "sj", "sz", "se", "ch", "sy", "tw", "tj", "tz", "th", "tl", "tg", "tk", "to", "tt", "tn", "tr", "tm", "tc", "tv", "ug", "ua", "ae", "gb", "us", "um", "uy", "uz", "vu", "ve", "vn", "vg", "vi", "wf", "eh", "ye", "zm", "zw",
	)
	apijson.RegisterFieldValidator[BrandGetParamsBodyByName](
		"force_language", "afrikaans", "albanian", "amharic", "arabic", "armenian", "assamese", "aymara", "azeri", "basque", "belarusian", "bengali", "bosnian", "bulgarian", "burmese", "cantonese", "catalan", "cebuano", "chinese", "corsican", "croatian", "czech", "danish", "dutch", "english", "esperanto", "estonian", "farsi", "fijian", "finnish", "french", "galician", "georgian", "german", "greek", "guarani", "gujarati", "haitian-creole", "hausa", "hawaiian", "hebrew", "hindi", "hmong", "hungarian", "icelandic", "igbo", "indonesian", "irish", "italian", "japanese", "javanese", "kannada", "kazakh", "khmer", "kinyarwanda", "korean", "kurdish", "kyrgyz", "lao", "latin", "latvian", "lingala", "lithuanian", "luxembourgish", "macedonian", "malagasy", "malay", "malayalam", "maltese", "maori", "marathi", "mongolian", "nepali", "norwegian", "odia", "oromo", "pashto", "pidgin", "polish", "portuguese", "punjabi", "quechua", "romanian", "russian", "samoan", "scottish-gaelic", "serbian", "sesotho", "shona", "sindhi", "sinhala", "slovak", "slovene", "somali", "spanish", "sundanese", "swahili", "swedish", "tagalog", "tajik", "tamil", "tatar", "telugu", "thai", "tibetan", "tigrinya", "tongan", "tswana", "turkish", "turkmen", "ukrainian", "urdu", "uyghur", "uzbek", "vietnamese", "welsh", "wolof", "xhosa", "yiddish", "yoruba", "zulu",
	)
	apijson.RegisterFieldValidator[BrandGetParamsBodyByName](
		"zdr", "enabled", "disabled",
	)
}

// Request deadline and what to return when it passes.
//
// The property Milliseconds is required.
type BrandGetParamsBodyByNameTimeoutOpts struct {
	// Deadline in milliseconds.
	Milliseconds int64 `json:"milliseconds" api:"required"`
	// `fail` returns 408 at the deadline. `return-partial` returns the results
	// collected so far and marks the response as partial.
	//
	// Any of "fail", "return-partial".
	Behavior string `json:"behavior,omitzero"`
	paramObj
}

func (r BrandGetParamsBodyByNameTimeoutOpts) MarshalJSON() (data []byte, err error) {
	type shadow BrandGetParamsBodyByNameTimeoutOpts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandGetParamsBodyByNameTimeoutOpts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BrandGetParamsBodyByNameTimeoutOpts](
		"behavior", "fail", "return-partial",
	)
}

// Look up a brand from a work email address. Free and disposable email providers
// return 422.
//
// The properties Email, Type are required.
type BrandGetParamsBodyByEmail struct {
	// Work email address, such as `jane@stripe.com`. The brand is looked up from its
	// domain.
	Email string `json:"email" api:"required" format:"email"`
	// Maximum age of cached brand data, in milliseconds. `0` always fetches fresh.
	// Defaults to 90 days (7776000000 ms). Values above 1 year (31536000000 ms) are
	// capped at 1 year.
	MaxAgeMs param.Opt[int64] `json:"maxAgeMs,omitzero"`
	// Skip slower enrichment steps for a faster, less detailed response.
	MaxSpeed param.Opt[bool] `json:"maxSpeed,omitzero"`
	// Translate the brand's `title`, `description`, and `slogan` into this language.
	//
	// Any of "afrikaans", "albanian", "amharic", "arabic", "armenian", "assamese",
	// "aymara", "azeri", "basque", "belarusian", "bengali", "bosnian", "bulgarian",
	// "burmese", "cantonese", "catalan", "cebuano", "chinese", "corsican", "croatian",
	// "czech", "danish", "dutch", "english", "esperanto", "estonian", "farsi",
	// "fijian", "finnish", "french", "galician", "georgian", "german", "greek",
	// "guarani", "gujarati", "haitian-creole", "hausa", "hawaiian", "hebrew", "hindi",
	// "hmong", "hungarian", "icelandic", "igbo", "indonesian", "irish", "italian",
	// "japanese", "javanese", "kannada", "kazakh", "khmer", "kinyarwanda", "korean",
	// "kurdish", "kyrgyz", "lao", "latin", "latvian", "lingala", "lithuanian",
	// "luxembourgish", "macedonian", "malagasy", "malay", "malayalam", "maltese",
	// "maori", "marathi", "mongolian", "nepali", "norwegian", "odia", "oromo",
	// "pashto", "pidgin", "polish", "portuguese", "punjabi", "quechua", "romanian",
	// "russian", "samoan", "scottish-gaelic", "serbian", "sesotho", "shona", "sindhi",
	// "sinhala", "slovak", "slovene", "somali", "spanish", "sundanese", "swahili",
	// "swedish", "tagalog", "tajik", "tamil", "tatar", "telugu", "thai", "tibetan",
	// "tigrinya", "tongan", "tswana", "turkish", "turkmen", "ukrainian", "urdu",
	// "uyghur", "uzbek", "vietnamese", "welsh", "wolof", "xhosa", "yiddish", "yoruba",
	// "zulu".
	ForceLanguage string `json:"force_language,omitzero"`
	// Access policy IDs to apply; each URL must pass all of them. Requires access
	// policies enabled for your organization; otherwise returns
	// `ACCESS_POLICIES_NOT_ENABLED`.
	Policies []string `json:"policies,omitzero"`
	// Labels for this request. Filter usage and request logs by them in the dashboard
	// or with `GET /org/logs`.
	Tags []string `json:"tags,omitzero"`
	// Request deadline and what to return when it passes.
	TimeoutOpts BrandGetParamsBodyByEmailTimeoutOpts `json:"timeoutOpts,omitzero"`
	// `enabled` turns off request logs and traces and keeps request and response
	// content out of usage logs. Brand data and hosted assets may still be stored.
	// Returns 403 `ZDR_NOT_ENABLED` unless your organization has ZDR.
	//
	// Any of "enabled", "disabled".
	Zdr string `json:"zdr,omitzero"`
	// Set to `by_email` to look up a brand from a work email address.
	//
	// This field can be elided, and will marshal its zero value as "by_email".
	Type constant.ByEmail `json:"type" default:"by_email"`
	paramObj
}

func (r BrandGetParamsBodyByEmail) MarshalJSON() (data []byte, err error) {
	type shadow BrandGetParamsBodyByEmail
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandGetParamsBodyByEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BrandGetParamsBodyByEmail](
		"force_language", "afrikaans", "albanian", "amharic", "arabic", "armenian", "assamese", "aymara", "azeri", "basque", "belarusian", "bengali", "bosnian", "bulgarian", "burmese", "cantonese", "catalan", "cebuano", "chinese", "corsican", "croatian", "czech", "danish", "dutch", "english", "esperanto", "estonian", "farsi", "fijian", "finnish", "french", "galician", "georgian", "german", "greek", "guarani", "gujarati", "haitian-creole", "hausa", "hawaiian", "hebrew", "hindi", "hmong", "hungarian", "icelandic", "igbo", "indonesian", "irish", "italian", "japanese", "javanese", "kannada", "kazakh", "khmer", "kinyarwanda", "korean", "kurdish", "kyrgyz", "lao", "latin", "latvian", "lingala", "lithuanian", "luxembourgish", "macedonian", "malagasy", "malay", "malayalam", "maltese", "maori", "marathi", "mongolian", "nepali", "norwegian", "odia", "oromo", "pashto", "pidgin", "polish", "portuguese", "punjabi", "quechua", "romanian", "russian", "samoan", "scottish-gaelic", "serbian", "sesotho", "shona", "sindhi", "sinhala", "slovak", "slovene", "somali", "spanish", "sundanese", "swahili", "swedish", "tagalog", "tajik", "tamil", "tatar", "telugu", "thai", "tibetan", "tigrinya", "tongan", "tswana", "turkish", "turkmen", "ukrainian", "urdu", "uyghur", "uzbek", "vietnamese", "welsh", "wolof", "xhosa", "yiddish", "yoruba", "zulu",
	)
	apijson.RegisterFieldValidator[BrandGetParamsBodyByEmail](
		"zdr", "enabled", "disabled",
	)
}

// Request deadline and what to return when it passes.
//
// The property Milliseconds is required.
type BrandGetParamsBodyByEmailTimeoutOpts struct {
	// Deadline in milliseconds.
	Milliseconds int64 `json:"milliseconds" api:"required"`
	// `fail` returns 408 at the deadline. `return-partial` returns the results
	// collected so far and marks the response as partial.
	//
	// Any of "fail", "return-partial".
	Behavior string `json:"behavior,omitzero"`
	paramObj
}

func (r BrandGetParamsBodyByEmailTimeoutOpts) MarshalJSON() (data []byte, err error) {
	type shadow BrandGetParamsBodyByEmailTimeoutOpts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandGetParamsBodyByEmailTimeoutOpts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BrandGetParamsBodyByEmailTimeoutOpts](
		"behavior", "fail", "return-partial",
	)
}

// Look up a publicly traded company by stock ticker.
//
// The properties Ticker, Type are required.
type BrandGetParamsBodyByTicker struct {
	// Stock ticker symbol, such as `AAPL`.
	Ticker string `json:"ticker" api:"required"`
	// Maximum age of cached brand data, in milliseconds. `0` always fetches fresh.
	// Defaults to 90 days (7776000000 ms). Values above 1 year (31536000000 ms) are
	// capped at 1 year.
	MaxAgeMs param.Opt[int64] `json:"maxAgeMs,omitzero"`
	// Skip slower enrichment steps for a faster, less detailed response.
	MaxSpeed param.Opt[bool] `json:"maxSpeed,omitzero"`
	// Translate the brand's `title`, `description`, and `slogan` into this language.
	//
	// Any of "afrikaans", "albanian", "amharic", "arabic", "armenian", "assamese",
	// "aymara", "azeri", "basque", "belarusian", "bengali", "bosnian", "bulgarian",
	// "burmese", "cantonese", "catalan", "cebuano", "chinese", "corsican", "croatian",
	// "czech", "danish", "dutch", "english", "esperanto", "estonian", "farsi",
	// "fijian", "finnish", "french", "galician", "georgian", "german", "greek",
	// "guarani", "gujarati", "haitian-creole", "hausa", "hawaiian", "hebrew", "hindi",
	// "hmong", "hungarian", "icelandic", "igbo", "indonesian", "irish", "italian",
	// "japanese", "javanese", "kannada", "kazakh", "khmer", "kinyarwanda", "korean",
	// "kurdish", "kyrgyz", "lao", "latin", "latvian", "lingala", "lithuanian",
	// "luxembourgish", "macedonian", "malagasy", "malay", "malayalam", "maltese",
	// "maori", "marathi", "mongolian", "nepali", "norwegian", "odia", "oromo",
	// "pashto", "pidgin", "polish", "portuguese", "punjabi", "quechua", "romanian",
	// "russian", "samoan", "scottish-gaelic", "serbian", "sesotho", "shona", "sindhi",
	// "sinhala", "slovak", "slovene", "somali", "spanish", "sundanese", "swahili",
	// "swedish", "tagalog", "tajik", "tamil", "tatar", "telugu", "thai", "tibetan",
	// "tigrinya", "tongan", "tswana", "turkish", "turkmen", "ukrainian", "urdu",
	// "uyghur", "uzbek", "vietnamese", "welsh", "wolof", "xhosa", "yiddish", "yoruba",
	// "zulu".
	ForceLanguage string `json:"force_language,omitzero"`
	// Access policy IDs to apply; each URL must pass all of them. Requires access
	// policies enabled for your organization; otherwise returns
	// `ACCESS_POLICIES_NOT_ENABLED`.
	Policies []string `json:"policies,omitzero"`
	// Labels for this request. Filter usage and request logs by them in the dashboard
	// or with `GET /org/logs`.
	Tags []string `json:"tags,omitzero"`
	// Stock exchange code, such as `NASDAQ` or `LSE`.
	//
	// Any of "AMEX", "AMS", "AQS", "ASX", "ATH", "BER", "BME", "BRU", "BSE", "BUD",
	// "BUE", "BVC", "CBOE", "CNQ", "CPH", "DFM", "DOH", "DUB", "DUS", "DXE", "EGX",
	// "FSX", "HAM", "HEL", "HKSE", "HOSE", "ICE", "IOB", "IST", "JKT", "JNB", "JPX",
	// "KLS", "KOE", "KSC", "KUW", "LIS", "LSE", "MCX", "MEX", "MIL", "MUN", "NASDAQ",
	// "NEO", "NSE", "NYSE", "NZE", "OSL", "OTC", "PAR", "PNK", "PRA", "RIS", "SAO",
	// "SAU", "SES", "SET", "SGO", "SHH", "SHZ", "SIX", "STO", "STU", "TAI", "TAL",
	// "TLV", "TSX", "TSXV", "TWO", "VIE", "WSE", "XETRA".
	TickerExchange string `json:"ticker_exchange,omitzero"`
	// Request deadline and what to return when it passes.
	TimeoutOpts BrandGetParamsBodyByTickerTimeoutOpts `json:"timeoutOpts,omitzero"`
	// `enabled` turns off request logs and traces and keeps request and response
	// content out of usage logs. Brand data and hosted assets may still be stored.
	// Returns 403 `ZDR_NOT_ENABLED` unless your organization has ZDR.
	//
	// Any of "enabled", "disabled".
	Zdr string `json:"zdr,omitzero"`
	// Set to `by_ticker` to look up a brand by stock ticker.
	//
	// This field can be elided, and will marshal its zero value as "by_ticker".
	Type constant.ByTicker `json:"type" default:"by_ticker"`
	paramObj
}

func (r BrandGetParamsBodyByTicker) MarshalJSON() (data []byte, err error) {
	type shadow BrandGetParamsBodyByTicker
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandGetParamsBodyByTicker) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BrandGetParamsBodyByTicker](
		"force_language", "afrikaans", "albanian", "amharic", "arabic", "armenian", "assamese", "aymara", "azeri", "basque", "belarusian", "bengali", "bosnian", "bulgarian", "burmese", "cantonese", "catalan", "cebuano", "chinese", "corsican", "croatian", "czech", "danish", "dutch", "english", "esperanto", "estonian", "farsi", "fijian", "finnish", "french", "galician", "georgian", "german", "greek", "guarani", "gujarati", "haitian-creole", "hausa", "hawaiian", "hebrew", "hindi", "hmong", "hungarian", "icelandic", "igbo", "indonesian", "irish", "italian", "japanese", "javanese", "kannada", "kazakh", "khmer", "kinyarwanda", "korean", "kurdish", "kyrgyz", "lao", "latin", "latvian", "lingala", "lithuanian", "luxembourgish", "macedonian", "malagasy", "malay", "malayalam", "maltese", "maori", "marathi", "mongolian", "nepali", "norwegian", "odia", "oromo", "pashto", "pidgin", "polish", "portuguese", "punjabi", "quechua", "romanian", "russian", "samoan", "scottish-gaelic", "serbian", "sesotho", "shona", "sindhi", "sinhala", "slovak", "slovene", "somali", "spanish", "sundanese", "swahili", "swedish", "tagalog", "tajik", "tamil", "tatar", "telugu", "thai", "tibetan", "tigrinya", "tongan", "tswana", "turkish", "turkmen", "ukrainian", "urdu", "uyghur", "uzbek", "vietnamese", "welsh", "wolof", "xhosa", "yiddish", "yoruba", "zulu",
	)
	apijson.RegisterFieldValidator[BrandGetParamsBodyByTicker](
		"ticker_exchange", "AMEX", "AMS", "AQS", "ASX", "ATH", "BER", "BME", "BRU", "BSE", "BUD", "BUE", "BVC", "CBOE", "CNQ", "CPH", "DFM", "DOH", "DUB", "DUS", "DXE", "EGX", "FSX", "HAM", "HEL", "HKSE", "HOSE", "ICE", "IOB", "IST", "JKT", "JNB", "JPX", "KLS", "KOE", "KSC", "KUW", "LIS", "LSE", "MCX", "MEX", "MIL", "MUN", "NASDAQ", "NEO", "NSE", "NYSE", "NZE", "OSL", "OTC", "PAR", "PNK", "PRA", "RIS", "SAO", "SAU", "SES", "SET", "SGO", "SHH", "SHZ", "SIX", "STO", "STU", "TAI", "TAL", "TLV", "TSX", "TSXV", "TWO", "VIE", "WSE", "XETRA",
	)
	apijson.RegisterFieldValidator[BrandGetParamsBodyByTicker](
		"zdr", "enabled", "disabled",
	)
}

// Request deadline and what to return when it passes.
//
// The property Milliseconds is required.
type BrandGetParamsBodyByTickerTimeoutOpts struct {
	// Deadline in milliseconds.
	Milliseconds int64 `json:"milliseconds" api:"required"`
	// `fail` returns 408 at the deadline. `return-partial` returns the results
	// collected so far and marks the response as partial.
	//
	// Any of "fail", "return-partial".
	Behavior string `json:"behavior,omitzero"`
	paramObj
}

func (r BrandGetParamsBodyByTickerTimeoutOpts) MarshalJSON() (data []byte, err error) {
	type shadow BrandGetParamsBodyByTickerTimeoutOpts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandGetParamsBodyByTickerTimeoutOpts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BrandGetParamsBodyByTickerTimeoutOpts](
		"behavior", "fail", "return-partial",
	)
}

// Extract brand data from a single page. Only that URL is read, and no data from
// other sources is added.
//
// The properties DirectURL, Type are required.
type BrandGetParamsBodyByDirectURL struct {
	// Full `http` or `https` URL of the page to read, such as
	// `https://stripe.com/enterprise`. Only this page is fetched.
	DirectURL string `json:"direct_url" api:"required" format:"uri"`
	// Access policy IDs to apply; each URL must pass all of them. Requires access
	// policies enabled for your organization; otherwise returns
	// `ACCESS_POLICIES_NOT_ENABLED`.
	Policies []string `json:"policies,omitzero"`
	// Labels for this request. Filter usage and request logs by them in the dashboard
	// or with `GET /org/logs`.
	Tags []string `json:"tags,omitzero"`
	// Request deadline and what to return when it passes.
	TimeoutOpts BrandGetParamsBodyByDirectURLTimeoutOpts `json:"timeoutOpts,omitzero"`
	// `enabled` turns off request logs and traces and keeps request and response
	// content out of usage logs. Brand data and hosted assets may still be stored.
	// Returns 403 `ZDR_NOT_ENABLED` unless your organization has ZDR.
	//
	// Any of "enabled", "disabled".
	Zdr string `json:"zdr,omitzero"`
	// Set to `by_direct_url` to extract brand data from a single page.
	//
	// This field can be elided, and will marshal its zero value as "by_direct_url".
	Type constant.ByDirectURL `json:"type" default:"by_direct_url"`
	paramObj
}

func (r BrandGetParamsBodyByDirectURL) MarshalJSON() (data []byte, err error) {
	type shadow BrandGetParamsBodyByDirectURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandGetParamsBodyByDirectURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BrandGetParamsBodyByDirectURL](
		"zdr", "enabled", "disabled",
	)
}

// Request deadline and what to return when it passes.
//
// The property Milliseconds is required.
type BrandGetParamsBodyByDirectURLTimeoutOpts struct {
	// Deadline in milliseconds.
	Milliseconds int64 `json:"milliseconds" api:"required"`
	// `fail` returns 408 at the deadline. `return-partial` returns the results
	// collected so far and marks the response as partial.
	//
	// Any of "fail", "return-partial".
	Behavior string `json:"behavior,omitzero"`
	paramObj
}

func (r BrandGetParamsBodyByDirectURLTimeoutOpts) MarshalJSON() (data []byte, err error) {
	type shadow BrandGetParamsBodyByDirectURLTimeoutOpts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandGetParamsBodyByDirectURLTimeoutOpts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BrandGetParamsBodyByDirectURLTimeoutOpts](
		"behavior", "fail", "return-partial",
	)
}

// Identify the merchant behind a card or bank transaction descriptor.
//
// The properties TransactionInfo, Type are required.
type BrandGetParamsBodyByTransaction struct {
	// Transaction descriptor from a card or bank statement, such as `SQ *COFFEE SHOP`.
	TransactionInfo string `json:"transaction_info" api:"required"`
	// City where the transaction took place. Brands in this city are preferred.
	City param.Opt[string] `json:"city,omitzero"`
	// Run extra verification so the returned brand matches the transaction with high
	// confidence.
	HighConfidenceOnly param.Opt[bool] `json:"high_confidence_only,omitzero"`
	// Skip slower enrichment steps for a faster, less detailed response.
	MaxSpeed param.Opt[bool] `json:"maxSpeed,omitzero"`
	// Translate the brand's `title`, `description`, and `slogan` into this language.
	//
	// Any of "afrikaans", "albanian", "amharic", "arabic", "armenian", "assamese",
	// "aymara", "azeri", "basque", "belarusian", "bengali", "bosnian", "bulgarian",
	// "burmese", "cantonese", "catalan", "cebuano", "chinese", "corsican", "croatian",
	// "czech", "danish", "dutch", "english", "esperanto", "estonian", "farsi",
	// "fijian", "finnish", "french", "galician", "georgian", "german", "greek",
	// "guarani", "gujarati", "haitian-creole", "hausa", "hawaiian", "hebrew", "hindi",
	// "hmong", "hungarian", "icelandic", "igbo", "indonesian", "irish", "italian",
	// "japanese", "javanese", "kannada", "kazakh", "khmer", "kinyarwanda", "korean",
	// "kurdish", "kyrgyz", "lao", "latin", "latvian", "lingala", "lithuanian",
	// "luxembourgish", "macedonian", "malagasy", "malay", "malayalam", "maltese",
	// "maori", "marathi", "mongolian", "nepali", "norwegian", "odia", "oromo",
	// "pashto", "pidgin", "polish", "portuguese", "punjabi", "quechua", "romanian",
	// "russian", "samoan", "scottish-gaelic", "serbian", "sesotho", "shona", "sindhi",
	// "sinhala", "slovak", "slovene", "somali", "spanish", "sundanese", "swahili",
	// "swedish", "tagalog", "tajik", "tamil", "tatar", "telugu", "thai", "tibetan",
	// "tigrinya", "tongan", "tswana", "turkish", "turkmen", "ukrainian", "urdu",
	// "uyghur", "uzbek", "vietnamese", "welsh", "wolof", "xhosa", "yiddish", "yoruba",
	// "zulu".
	ForceLanguage string `json:"force_language,omitzero"`
	// Lowercase ISO 3166-1 alpha-2 country code, such as `us`.
	//
	// Any of "af", "al", "dz", "as", "ad", "ao", "ai", "aq", "ag", "ar", "am", "aw",
	// "au", "at", "az", "bs", "bh", "bd", "bb", "by", "be", "bz", "bj", "bm", "bt",
	// "bo", "ba", "bw", "bv", "br", "io", "bn", "bg", "bf", "bi", "kh", "cm", "ca",
	// "cv", "ky", "cf", "td", "cl", "cn", "cx", "cc", "co", "km", "cg", "cd", "ck",
	// "cr", "ci", "hr", "cu", "cy", "cz", "dk", "dj", "dm", "do", "ec", "eg", "sv",
	// "gq", "er", "ee", "et", "fk", "fo", "fj", "fi", "fr", "gf", "pf", "tf", "ga",
	// "gm", "ge", "de", "gh", "gi", "gr", "gl", "gd", "gp", "gu", "gt", "gn", "gw",
	// "gy", "ht", "hm", "va", "hn", "hk", "hu", "is", "in", "id", "ir", "iq", "ie",
	// "il", "it", "jm", "jp", "jo", "kz", "ke", "ki", "kp", "kr", "kw", "kg", "la",
	// "lv", "lb", "ls", "lr", "ly", "li", "lt", "lu", "mo", "mk", "mg", "mw", "my",
	// "mv", "ml", "mt", "mh", "mq", "mr", "mu", "yt", "mx", "fm", "md", "mc", "mn",
	// "ms", "ma", "mz", "mm", "na", "nr", "np", "nl", "an", "nc", "nz", "ni", "ne",
	// "ng", "nu", "nf", "mp", "no", "om", "pk", "pw", "ps", "pa", "pg", "py", "pe",
	// "ph", "pn", "pl", "pt", "pr", "qa", "re", "ro", "ru", "rw", "sh", "kn", "lc",
	// "pm", "vc", "ws", "sm", "st", "sa", "sn", "rs", "sc", "sl", "sg", "sk", "si",
	// "sb", "so", "za", "gs", "es", "lk", "sd", "sr", "sj", "sz", "se", "ch", "sy",
	// "tw", "tj", "tz", "th", "tl", "tg", "tk", "to", "tt", "tn", "tr", "tm", "tc",
	// "tv", "ug", "ua", "ae", "gb", "us", "um", "uy", "uz", "vu", "ve", "vn", "vg",
	// "vi", "wf", "eh", "ye", "zm", "zw".
	Country string `json:"country,omitzero"`
	// Merchant Category Code (MCC) from the transaction. Helps identify the type of
	// business.
	Mcc BrandGetParamsBodyByTransactionMccUnion `json:"mcc,omitzero"`
	// Merchant phone number from the transaction. Helps confirm the match.
	Phone BrandGetParamsBodyByTransactionPhoneUnion `json:"phone,omitzero"`
	// Access policy IDs to apply; each URL must pass all of them. Requires access
	// policies enabled for your organization; otherwise returns
	// `ACCESS_POLICIES_NOT_ENABLED`.
	Policies []string `json:"policies,omitzero"`
	// Labels for this request. Filter usage and request logs by them in the dashboard
	// or with `GET /org/logs`.
	Tags []string `json:"tags,omitzero"`
	// Request deadline and what to return when it passes.
	TimeoutOpts BrandGetParamsBodyByTransactionTimeoutOpts `json:"timeoutOpts,omitzero"`
	// `enabled` turns off request logs and traces and keeps request and response
	// content out of usage logs. Brand data and hosted assets may still be stored.
	// Returns 403 `ZDR_NOT_ENABLED` unless your organization has ZDR.
	//
	// Any of "enabled", "disabled".
	Zdr string `json:"zdr,omitzero"`
	// Set to `by_transaction` to identify the merchant behind a transaction
	// descriptor.
	//
	// This field can be elided, and will marshal its zero value as "by_transaction".
	Type constant.ByTransaction `json:"type" default:"by_transaction"`
	paramObj
}

func (r BrandGetParamsBodyByTransaction) MarshalJSON() (data []byte, err error) {
	type shadow BrandGetParamsBodyByTransaction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandGetParamsBodyByTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BrandGetParamsBodyByTransaction](
		"country", "af", "al", "dz", "as", "ad", "ao", "ai", "aq", "ag", "ar", "am", "aw", "au", "at", "az", "bs", "bh", "bd", "bb", "by", "be", "bz", "bj", "bm", "bt", "bo", "ba", "bw", "bv", "br", "io", "bn", "bg", "bf", "bi", "kh", "cm", "ca", "cv", "ky", "cf", "td", "cl", "cn", "cx", "cc", "co", "km", "cg", "cd", "ck", "cr", "ci", "hr", "cu", "cy", "cz", "dk", "dj", "dm", "do", "ec", "eg", "sv", "gq", "er", "ee", "et", "fk", "fo", "fj", "fi", "fr", "gf", "pf", "tf", "ga", "gm", "ge", "de", "gh", "gi", "gr", "gl", "gd", "gp", "gu", "gt", "gn", "gw", "gy", "ht", "hm", "va", "hn", "hk", "hu", "is", "in", "id", "ir", "iq", "ie", "il", "it", "jm", "jp", "jo", "kz", "ke", "ki", "kp", "kr", "kw", "kg", "la", "lv", "lb", "ls", "lr", "ly", "li", "lt", "lu", "mo", "mk", "mg", "mw", "my", "mv", "ml", "mt", "mh", "mq", "mr", "mu", "yt", "mx", "fm", "md", "mc", "mn", "ms", "ma", "mz", "mm", "na", "nr", "np", "nl", "an", "nc", "nz", "ni", "ne", "ng", "nu", "nf", "mp", "no", "om", "pk", "pw", "ps", "pa", "pg", "py", "pe", "ph", "pn", "pl", "pt", "pr", "qa", "re", "ro", "ru", "rw", "sh", "kn", "lc", "pm", "vc", "ws", "sm", "st", "sa", "sn", "rs", "sc", "sl", "sg", "sk", "si", "sb", "so", "za", "gs", "es", "lk", "sd", "sr", "sj", "sz", "se", "ch", "sy", "tw", "tj", "tz", "th", "tl", "tg", "tk", "to", "tt", "tn", "tr", "tm", "tc", "tv", "ug", "ua", "ae", "gb", "us", "um", "uy", "uz", "vu", "ve", "vn", "vg", "vi", "wf", "eh", "ye", "zm", "zw",
	)
	apijson.RegisterFieldValidator[BrandGetParamsBodyByTransaction](
		"force_language", "afrikaans", "albanian", "amharic", "arabic", "armenian", "assamese", "aymara", "azeri", "basque", "belarusian", "bengali", "bosnian", "bulgarian", "burmese", "cantonese", "catalan", "cebuano", "chinese", "corsican", "croatian", "czech", "danish", "dutch", "english", "esperanto", "estonian", "farsi", "fijian", "finnish", "french", "galician", "georgian", "german", "greek", "guarani", "gujarati", "haitian-creole", "hausa", "hawaiian", "hebrew", "hindi", "hmong", "hungarian", "icelandic", "igbo", "indonesian", "irish", "italian", "japanese", "javanese", "kannada", "kazakh", "khmer", "kinyarwanda", "korean", "kurdish", "kyrgyz", "lao", "latin", "latvian", "lingala", "lithuanian", "luxembourgish", "macedonian", "malagasy", "malay", "malayalam", "maltese", "maori", "marathi", "mongolian", "nepali", "norwegian", "odia", "oromo", "pashto", "pidgin", "polish", "portuguese", "punjabi", "quechua", "romanian", "russian", "samoan", "scottish-gaelic", "serbian", "sesotho", "shona", "sindhi", "sinhala", "slovak", "slovene", "somali", "spanish", "sundanese", "swahili", "swedish", "tagalog", "tajik", "tamil", "tatar", "telugu", "thai", "tibetan", "tigrinya", "tongan", "tswana", "turkish", "turkmen", "ukrainian", "urdu", "uyghur", "uzbek", "vietnamese", "welsh", "wolof", "xhosa", "yiddish", "yoruba", "zulu",
	)
	apijson.RegisterFieldValidator[BrandGetParamsBodyByTransaction](
		"zdr", "enabled", "disabled",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BrandGetParamsBodyByTransactionMccUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	paramUnion
}

func (u BrandGetParamsBodyByTransactionMccUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat)
}
func (u *BrandGetParamsBodyByTransactionMccUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BrandGetParamsBodyByTransactionPhoneUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	paramUnion
}

func (u BrandGetParamsBodyByTransactionPhoneUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat)
}
func (u *BrandGetParamsBodyByTransactionPhoneUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Request deadline and what to return when it passes.
//
// The property Milliseconds is required.
type BrandGetParamsBodyByTransactionTimeoutOpts struct {
	// Deadline in milliseconds.
	Milliseconds int64 `json:"milliseconds" api:"required"`
	// `fail` returns 408 at the deadline. `return-partial` returns the results
	// collected so far and marks the response as partial.
	//
	// Any of "fail", "return-partial".
	Behavior string `json:"behavior,omitzero"`
	paramObj
}

func (r BrandGetParamsBodyByTransactionTimeoutOpts) MarshalJSON() (data []byte, err error) {
	type shadow BrandGetParamsBodyByTransactionTimeoutOpts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandGetParamsBodyByTransactionTimeoutOpts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BrandGetParamsBodyByTransactionTimeoutOpts](
		"behavior", "fail", "return-partial",
	)
}

type BrandSearchParams struct {
	// Search term, matched against the fields in `queryBy`, such as `nike`,
	// `nike.com`, or `nik`.
	Query string `query:"query" api:"required" json:"-"`
	// Match the search term as a prefix, so `nik` matches Nike while the user is still
	// typing. Set to `false` to match whole words only.
	Autocomplete param.Opt[bool] `query:"autocomplete,omitzero" json:"-"`
	// Maximum number of typos to tolerate when matching. `0` requires exact spelling.
	TypoTolerance param.Opt[int64] `query:"typoTolerance,omitzero" json:"-"`
	// Only return brands in this country, as an uppercase ISO 3166-1 alpha-2 code such
	// as `US` or `GB`.
	//
	// Any of "AD", "AE", "AF", "AG", "AI", "AL", "AM", "AO", "AQ", "AR", "AS", "AT",
	// "AU", "AW", "AX", "AZ", "BA", "BB", "BD", "BE", "BF", "BG", "BH", "BI", "BJ",
	// "BL", "BM", "BN", "BO", "BQ", "BR", "BS", "BT", "BV", "BW", "BY", "BZ", "CA",
	// "CC", "CD", "CF", "CG", "CH", "CI", "CK", "CL", "CM", "CN", "CO", "CR", "CU",
	// "CV", "CW", "CX", "CY", "CZ", "DE", "DJ", "DK", "DM", "DO", "DZ", "EC", "EE",
	// "EG", "EH", "ER", "ES", "ET", "FI", "FJ", "FK", "FM", "FO", "FR", "GA", "GB",
	// "GD", "GE", "GF", "GG", "GH", "GI", "GL", "GM", "GN", "GP", "GQ", "GR", "GS",
	// "GT", "GU", "GW", "GY", "HK", "HM", "HN", "HR", "HT", "HU", "ID", "IE", "IL",
	// "IM", "IN", "IO", "IQ", "IR", "IS", "IT", "JE", "JM", "JO", "JP", "KE", "KG",
	// "KH", "KI", "KM", "KN", "KP", "KR", "KW", "KY", "KZ", "LA", "LB", "LC", "LI",
	// "LK", "LR", "LS", "LT", "LU", "LV", "LY", "MA", "MC", "MD", "ME", "MF", "MG",
	// "MH", "MK", "ML", "MM", "MN", "MO", "MP", "MQ", "MR", "MS", "MT", "MU", "MV",
	// "MW", "MX", "MY", "MZ", "NA", "NC", "NE", "NF", "NG", "NI", "NL", "NO", "NP",
	// "NR", "NU", "NZ", "OM", "PA", "PE", "PF", "PG", "PH", "PK", "PL", "PM", "PN",
	// "PR", "PS", "PT", "PW", "PY", "QA", "RE", "RO", "RS", "RU", "RW", "SA", "SB",
	// "SC", "SD", "SE", "SG", "SH", "SI", "SJ", "SK", "SL", "SM", "SN", "SO", "SR",
	// "SS", "ST", "SV", "SX", "SY", "SZ", "TC", "TD", "TF", "TG", "TH", "TJ", "TK",
	// "TL", "TM", "TN", "TO", "TR", "TT", "TV", "TW", "TZ", "UA", "UG", "UM", "US",
	// "UY", "UZ", "VA", "VC", "VE", "VG", "VI", "VN", "VU", "WF", "WS", "YE", "YT",
	// "ZA", "ZM", "ZW".
	Country BrandSearchParamsCountry `query:"country,omitzero" json:"-"`
	// Only return brands with this primary language, such as `english` or `french`.
	//
	// Any of "afrikaans", "albanian", "amharic", "arabic", "armenian", "assamese",
	// "aymara", "azeri", "basque", "belarusian", "bengali", "bosnian", "bulgarian",
	// "burmese", "cantonese", "catalan", "cebuano", "chinese", "corsican", "croatian",
	// "czech", "danish", "dutch", "english", "esperanto", "estonian", "farsi",
	// "fijian", "finnish", "french", "galician", "georgian", "german", "greek",
	// "guarani", "gujarati", "haitian-creole", "hausa", "hawaiian", "hebrew", "hindi",
	// "hmong", "hungarian", "icelandic", "igbo", "indonesian", "irish", "italian",
	// "japanese", "javanese", "kannada", "kazakh", "khmer", "kinyarwanda", "korean",
	// "kurdish", "kyrgyz", "lao", "latin", "latvian", "lingala", "lithuanian",
	// "luxembourgish", "macedonian", "malagasy", "malay", "malayalam", "maltese",
	// "maori", "marathi", "mongolian", "nepali", "norwegian", "odia", "oromo",
	// "pashto", "pidgin", "polish", "portuguese", "punjabi", "quechua", "romanian",
	// "russian", "samoan", "scottish-gaelic", "serbian", "sesotho", "shona", "sindhi",
	// "sinhala", "slovak", "slovene", "somali", "spanish", "sundanese", "swahili",
	// "swedish", "tagalog", "tajik", "tamil", "tatar", "telugu", "thai", "tibetan",
	// "tigrinya", "tongan", "tswana", "turkish", "turkmen", "ukrainian", "urdu",
	// "uyghur", "uzbek", "vietnamese", "welsh", "wolof", "xhosa", "yiddish", "yoruba",
	// "zulu".
	Language BrandSearchParamsLanguage `query:"language,omitzero" json:"-"`
	// Fields to match the search term against: `name`, `domain`, or both. Send a
	// comma-separated list or repeat the parameter.
	//
	// Any of "name", "domain".
	QueryBy []string `query:"queryBy,omitzero" json:"-"`
	// Comma-separated labels for this request, such as `production,team-alpha`. Filter
	// usage and request logs by them.
	Tags []string `query:"tags,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BrandSearchParams]'s query parameters as `url.Values`.
func (r BrandSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only return brands in this country, as an uppercase ISO 3166-1 alpha-2 code such
// as `US` or `GB`.
type BrandSearchParamsCountry string

const (
	BrandSearchParamsCountryAd BrandSearchParamsCountry = "AD"
	BrandSearchParamsCountryAe BrandSearchParamsCountry = "AE"
	BrandSearchParamsCountryAf BrandSearchParamsCountry = "AF"
	BrandSearchParamsCountryAg BrandSearchParamsCountry = "AG"
	BrandSearchParamsCountryAI BrandSearchParamsCountry = "AI"
	BrandSearchParamsCountryAl BrandSearchParamsCountry = "AL"
	BrandSearchParamsCountryAm BrandSearchParamsCountry = "AM"
	BrandSearchParamsCountryAo BrandSearchParamsCountry = "AO"
	BrandSearchParamsCountryAq BrandSearchParamsCountry = "AQ"
	BrandSearchParamsCountryAr BrandSearchParamsCountry = "AR"
	BrandSearchParamsCountryAs BrandSearchParamsCountry = "AS"
	BrandSearchParamsCountryAt BrandSearchParamsCountry = "AT"
	BrandSearchParamsCountryAu BrandSearchParamsCountry = "AU"
	BrandSearchParamsCountryAw BrandSearchParamsCountry = "AW"
	BrandSearchParamsCountryAx BrandSearchParamsCountry = "AX"
	BrandSearchParamsCountryAz BrandSearchParamsCountry = "AZ"
	BrandSearchParamsCountryBa BrandSearchParamsCountry = "BA"
	BrandSearchParamsCountryBb BrandSearchParamsCountry = "BB"
	BrandSearchParamsCountryBd BrandSearchParamsCountry = "BD"
	BrandSearchParamsCountryBe BrandSearchParamsCountry = "BE"
	BrandSearchParamsCountryBf BrandSearchParamsCountry = "BF"
	BrandSearchParamsCountryBg BrandSearchParamsCountry = "BG"
	BrandSearchParamsCountryBh BrandSearchParamsCountry = "BH"
	BrandSearchParamsCountryBi BrandSearchParamsCountry = "BI"
	BrandSearchParamsCountryBj BrandSearchParamsCountry = "BJ"
	BrandSearchParamsCountryBl BrandSearchParamsCountry = "BL"
	BrandSearchParamsCountryBm BrandSearchParamsCountry = "BM"
	BrandSearchParamsCountryBn BrandSearchParamsCountry = "BN"
	BrandSearchParamsCountryBo BrandSearchParamsCountry = "BO"
	BrandSearchParamsCountryBq BrandSearchParamsCountry = "BQ"
	BrandSearchParamsCountryBr BrandSearchParamsCountry = "BR"
	BrandSearchParamsCountryBs BrandSearchParamsCountry = "BS"
	BrandSearchParamsCountryBt BrandSearchParamsCountry = "BT"
	BrandSearchParamsCountryBv BrandSearchParamsCountry = "BV"
	BrandSearchParamsCountryBw BrandSearchParamsCountry = "BW"
	BrandSearchParamsCountryBy BrandSearchParamsCountry = "BY"
	BrandSearchParamsCountryBz BrandSearchParamsCountry = "BZ"
	BrandSearchParamsCountryCa BrandSearchParamsCountry = "CA"
	BrandSearchParamsCountryCc BrandSearchParamsCountry = "CC"
	BrandSearchParamsCountryCd BrandSearchParamsCountry = "CD"
	BrandSearchParamsCountryCf BrandSearchParamsCountry = "CF"
	BrandSearchParamsCountryCg BrandSearchParamsCountry = "CG"
	BrandSearchParamsCountryCh BrandSearchParamsCountry = "CH"
	BrandSearchParamsCountryCi BrandSearchParamsCountry = "CI"
	BrandSearchParamsCountryCk BrandSearchParamsCountry = "CK"
	BrandSearchParamsCountryCl BrandSearchParamsCountry = "CL"
	BrandSearchParamsCountryCm BrandSearchParamsCountry = "CM"
	BrandSearchParamsCountryCn BrandSearchParamsCountry = "CN"
	BrandSearchParamsCountryCo BrandSearchParamsCountry = "CO"
	BrandSearchParamsCountryCr BrandSearchParamsCountry = "CR"
	BrandSearchParamsCountryCu BrandSearchParamsCountry = "CU"
	BrandSearchParamsCountryCv BrandSearchParamsCountry = "CV"
	BrandSearchParamsCountryCw BrandSearchParamsCountry = "CW"
	BrandSearchParamsCountryCx BrandSearchParamsCountry = "CX"
	BrandSearchParamsCountryCy BrandSearchParamsCountry = "CY"
	BrandSearchParamsCountryCz BrandSearchParamsCountry = "CZ"
	BrandSearchParamsCountryDe BrandSearchParamsCountry = "DE"
	BrandSearchParamsCountryDj BrandSearchParamsCountry = "DJ"
	BrandSearchParamsCountryDk BrandSearchParamsCountry = "DK"
	BrandSearchParamsCountryDm BrandSearchParamsCountry = "DM"
	BrandSearchParamsCountryDo BrandSearchParamsCountry = "DO"
	BrandSearchParamsCountryDz BrandSearchParamsCountry = "DZ"
	BrandSearchParamsCountryEc BrandSearchParamsCountry = "EC"
	BrandSearchParamsCountryEe BrandSearchParamsCountry = "EE"
	BrandSearchParamsCountryEg BrandSearchParamsCountry = "EG"
	BrandSearchParamsCountryEh BrandSearchParamsCountry = "EH"
	BrandSearchParamsCountryEr BrandSearchParamsCountry = "ER"
	BrandSearchParamsCountryEs BrandSearchParamsCountry = "ES"
	BrandSearchParamsCountryEt BrandSearchParamsCountry = "ET"
	BrandSearchParamsCountryFi BrandSearchParamsCountry = "FI"
	BrandSearchParamsCountryFj BrandSearchParamsCountry = "FJ"
	BrandSearchParamsCountryFk BrandSearchParamsCountry = "FK"
	BrandSearchParamsCountryFm BrandSearchParamsCountry = "FM"
	BrandSearchParamsCountryFo BrandSearchParamsCountry = "FO"
	BrandSearchParamsCountryFr BrandSearchParamsCountry = "FR"
	BrandSearchParamsCountryGa BrandSearchParamsCountry = "GA"
	BrandSearchParamsCountryGB BrandSearchParamsCountry = "GB"
	BrandSearchParamsCountryGd BrandSearchParamsCountry = "GD"
	BrandSearchParamsCountryGe BrandSearchParamsCountry = "GE"
	BrandSearchParamsCountryGf BrandSearchParamsCountry = "GF"
	BrandSearchParamsCountryGg BrandSearchParamsCountry = "GG"
	BrandSearchParamsCountryGh BrandSearchParamsCountry = "GH"
	BrandSearchParamsCountryGi BrandSearchParamsCountry = "GI"
	BrandSearchParamsCountryGl BrandSearchParamsCountry = "GL"
	BrandSearchParamsCountryGm BrandSearchParamsCountry = "GM"
	BrandSearchParamsCountryGn BrandSearchParamsCountry = "GN"
	BrandSearchParamsCountryGp BrandSearchParamsCountry = "GP"
	BrandSearchParamsCountryGq BrandSearchParamsCountry = "GQ"
	BrandSearchParamsCountryGr BrandSearchParamsCountry = "GR"
	BrandSearchParamsCountryGs BrandSearchParamsCountry = "GS"
	BrandSearchParamsCountryGt BrandSearchParamsCountry = "GT"
	BrandSearchParamsCountryGu BrandSearchParamsCountry = "GU"
	BrandSearchParamsCountryGw BrandSearchParamsCountry = "GW"
	BrandSearchParamsCountryGy BrandSearchParamsCountry = "GY"
	BrandSearchParamsCountryHk BrandSearchParamsCountry = "HK"
	BrandSearchParamsCountryHm BrandSearchParamsCountry = "HM"
	BrandSearchParamsCountryHn BrandSearchParamsCountry = "HN"
	BrandSearchParamsCountryHr BrandSearchParamsCountry = "HR"
	BrandSearchParamsCountryHt BrandSearchParamsCountry = "HT"
	BrandSearchParamsCountryHu BrandSearchParamsCountry = "HU"
	BrandSearchParamsCountryID BrandSearchParamsCountry = "ID"
	BrandSearchParamsCountryIe BrandSearchParamsCountry = "IE"
	BrandSearchParamsCountryIl BrandSearchParamsCountry = "IL"
	BrandSearchParamsCountryIm BrandSearchParamsCountry = "IM"
	BrandSearchParamsCountryIn BrandSearchParamsCountry = "IN"
	BrandSearchParamsCountryIo BrandSearchParamsCountry = "IO"
	BrandSearchParamsCountryIq BrandSearchParamsCountry = "IQ"
	BrandSearchParamsCountryIr BrandSearchParamsCountry = "IR"
	BrandSearchParamsCountryIs BrandSearchParamsCountry = "IS"
	BrandSearchParamsCountryIt BrandSearchParamsCountry = "IT"
	BrandSearchParamsCountryJe BrandSearchParamsCountry = "JE"
	BrandSearchParamsCountryJm BrandSearchParamsCountry = "JM"
	BrandSearchParamsCountryJo BrandSearchParamsCountry = "JO"
	BrandSearchParamsCountryJp BrandSearchParamsCountry = "JP"
	BrandSearchParamsCountryKe BrandSearchParamsCountry = "KE"
	BrandSearchParamsCountryKg BrandSearchParamsCountry = "KG"
	BrandSearchParamsCountryKh BrandSearchParamsCountry = "KH"
	BrandSearchParamsCountryKi BrandSearchParamsCountry = "KI"
	BrandSearchParamsCountryKm BrandSearchParamsCountry = "KM"
	BrandSearchParamsCountryKn BrandSearchParamsCountry = "KN"
	BrandSearchParamsCountryKp BrandSearchParamsCountry = "KP"
	BrandSearchParamsCountryKr BrandSearchParamsCountry = "KR"
	BrandSearchParamsCountryKw BrandSearchParamsCountry = "KW"
	BrandSearchParamsCountryKy BrandSearchParamsCountry = "KY"
	BrandSearchParamsCountryKz BrandSearchParamsCountry = "KZ"
	BrandSearchParamsCountryLa BrandSearchParamsCountry = "LA"
	BrandSearchParamsCountryLb BrandSearchParamsCountry = "LB"
	BrandSearchParamsCountryLc BrandSearchParamsCountry = "LC"
	BrandSearchParamsCountryLi BrandSearchParamsCountry = "LI"
	BrandSearchParamsCountryLk BrandSearchParamsCountry = "LK"
	BrandSearchParamsCountryLr BrandSearchParamsCountry = "LR"
	BrandSearchParamsCountryLs BrandSearchParamsCountry = "LS"
	BrandSearchParamsCountryLt BrandSearchParamsCountry = "LT"
	BrandSearchParamsCountryLu BrandSearchParamsCountry = "LU"
	BrandSearchParamsCountryLv BrandSearchParamsCountry = "LV"
	BrandSearchParamsCountryLy BrandSearchParamsCountry = "LY"
	BrandSearchParamsCountryMa BrandSearchParamsCountry = "MA"
	BrandSearchParamsCountryMc BrandSearchParamsCountry = "MC"
	BrandSearchParamsCountryMd BrandSearchParamsCountry = "MD"
	BrandSearchParamsCountryMe BrandSearchParamsCountry = "ME"
	BrandSearchParamsCountryMf BrandSearchParamsCountry = "MF"
	BrandSearchParamsCountryMg BrandSearchParamsCountry = "MG"
	BrandSearchParamsCountryMh BrandSearchParamsCountry = "MH"
	BrandSearchParamsCountryMk BrandSearchParamsCountry = "MK"
	BrandSearchParamsCountryMl BrandSearchParamsCountry = "ML"
	BrandSearchParamsCountryMm BrandSearchParamsCountry = "MM"
	BrandSearchParamsCountryMn BrandSearchParamsCountry = "MN"
	BrandSearchParamsCountryMo BrandSearchParamsCountry = "MO"
	BrandSearchParamsCountryMp BrandSearchParamsCountry = "MP"
	BrandSearchParamsCountryMq BrandSearchParamsCountry = "MQ"
	BrandSearchParamsCountryMr BrandSearchParamsCountry = "MR"
	BrandSearchParamsCountryMs BrandSearchParamsCountry = "MS"
	BrandSearchParamsCountryMt BrandSearchParamsCountry = "MT"
	BrandSearchParamsCountryMu BrandSearchParamsCountry = "MU"
	BrandSearchParamsCountryMv BrandSearchParamsCountry = "MV"
	BrandSearchParamsCountryMw BrandSearchParamsCountry = "MW"
	BrandSearchParamsCountryMx BrandSearchParamsCountry = "MX"
	BrandSearchParamsCountryMy BrandSearchParamsCountry = "MY"
	BrandSearchParamsCountryMz BrandSearchParamsCountry = "MZ"
	BrandSearchParamsCountryNa BrandSearchParamsCountry = "NA"
	BrandSearchParamsCountryNc BrandSearchParamsCountry = "NC"
	BrandSearchParamsCountryNe BrandSearchParamsCountry = "NE"
	BrandSearchParamsCountryNf BrandSearchParamsCountry = "NF"
	BrandSearchParamsCountryNg BrandSearchParamsCountry = "NG"
	BrandSearchParamsCountryNi BrandSearchParamsCountry = "NI"
	BrandSearchParamsCountryNl BrandSearchParamsCountry = "NL"
	BrandSearchParamsCountryNo BrandSearchParamsCountry = "NO"
	BrandSearchParamsCountryNp BrandSearchParamsCountry = "NP"
	BrandSearchParamsCountryNr BrandSearchParamsCountry = "NR"
	BrandSearchParamsCountryNu BrandSearchParamsCountry = "NU"
	BrandSearchParamsCountryNz BrandSearchParamsCountry = "NZ"
	BrandSearchParamsCountryOm BrandSearchParamsCountry = "OM"
	BrandSearchParamsCountryPa BrandSearchParamsCountry = "PA"
	BrandSearchParamsCountryPe BrandSearchParamsCountry = "PE"
	BrandSearchParamsCountryPf BrandSearchParamsCountry = "PF"
	BrandSearchParamsCountryPg BrandSearchParamsCountry = "PG"
	BrandSearchParamsCountryPh BrandSearchParamsCountry = "PH"
	BrandSearchParamsCountryPk BrandSearchParamsCountry = "PK"
	BrandSearchParamsCountryPl BrandSearchParamsCountry = "PL"
	BrandSearchParamsCountryPm BrandSearchParamsCountry = "PM"
	BrandSearchParamsCountryPn BrandSearchParamsCountry = "PN"
	BrandSearchParamsCountryPr BrandSearchParamsCountry = "PR"
	BrandSearchParamsCountryPs BrandSearchParamsCountry = "PS"
	BrandSearchParamsCountryPt BrandSearchParamsCountry = "PT"
	BrandSearchParamsCountryPw BrandSearchParamsCountry = "PW"
	BrandSearchParamsCountryPy BrandSearchParamsCountry = "PY"
	BrandSearchParamsCountryQa BrandSearchParamsCountry = "QA"
	BrandSearchParamsCountryRe BrandSearchParamsCountry = "RE"
	BrandSearchParamsCountryRo BrandSearchParamsCountry = "RO"
	BrandSearchParamsCountryRs BrandSearchParamsCountry = "RS"
	BrandSearchParamsCountryRu BrandSearchParamsCountry = "RU"
	BrandSearchParamsCountryRw BrandSearchParamsCountry = "RW"
	BrandSearchParamsCountrySa BrandSearchParamsCountry = "SA"
	BrandSearchParamsCountrySb BrandSearchParamsCountry = "SB"
	BrandSearchParamsCountrySc BrandSearchParamsCountry = "SC"
	BrandSearchParamsCountrySd BrandSearchParamsCountry = "SD"
	BrandSearchParamsCountrySe BrandSearchParamsCountry = "SE"
	BrandSearchParamsCountrySg BrandSearchParamsCountry = "SG"
	BrandSearchParamsCountrySh BrandSearchParamsCountry = "SH"
	BrandSearchParamsCountrySi BrandSearchParamsCountry = "SI"
	BrandSearchParamsCountrySj BrandSearchParamsCountry = "SJ"
	BrandSearchParamsCountrySk BrandSearchParamsCountry = "SK"
	BrandSearchParamsCountrySl BrandSearchParamsCountry = "SL"
	BrandSearchParamsCountrySm BrandSearchParamsCountry = "SM"
	BrandSearchParamsCountrySn BrandSearchParamsCountry = "SN"
	BrandSearchParamsCountrySo BrandSearchParamsCountry = "SO"
	BrandSearchParamsCountrySr BrandSearchParamsCountry = "SR"
	BrandSearchParamsCountrySS BrandSearchParamsCountry = "SS"
	BrandSearchParamsCountrySt BrandSearchParamsCountry = "ST"
	BrandSearchParamsCountrySv BrandSearchParamsCountry = "SV"
	BrandSearchParamsCountrySx BrandSearchParamsCountry = "SX"
	BrandSearchParamsCountrySy BrandSearchParamsCountry = "SY"
	BrandSearchParamsCountrySz BrandSearchParamsCountry = "SZ"
	BrandSearchParamsCountryTc BrandSearchParamsCountry = "TC"
	BrandSearchParamsCountryTd BrandSearchParamsCountry = "TD"
	BrandSearchParamsCountryTf BrandSearchParamsCountry = "TF"
	BrandSearchParamsCountryTg BrandSearchParamsCountry = "TG"
	BrandSearchParamsCountryTh BrandSearchParamsCountry = "TH"
	BrandSearchParamsCountryTj BrandSearchParamsCountry = "TJ"
	BrandSearchParamsCountryTk BrandSearchParamsCountry = "TK"
	BrandSearchParamsCountryTl BrandSearchParamsCountry = "TL"
	BrandSearchParamsCountryTm BrandSearchParamsCountry = "TM"
	BrandSearchParamsCountryTn BrandSearchParamsCountry = "TN"
	BrandSearchParamsCountryTo BrandSearchParamsCountry = "TO"
	BrandSearchParamsCountryTr BrandSearchParamsCountry = "TR"
	BrandSearchParamsCountryTt BrandSearchParamsCountry = "TT"
	BrandSearchParamsCountryTv BrandSearchParamsCountry = "TV"
	BrandSearchParamsCountryTw BrandSearchParamsCountry = "TW"
	BrandSearchParamsCountryTz BrandSearchParamsCountry = "TZ"
	BrandSearchParamsCountryUa BrandSearchParamsCountry = "UA"
	BrandSearchParamsCountryUg BrandSearchParamsCountry = "UG"
	BrandSearchParamsCountryUm BrandSearchParamsCountry = "UM"
	BrandSearchParamsCountryUs BrandSearchParamsCountry = "US"
	BrandSearchParamsCountryUy BrandSearchParamsCountry = "UY"
	BrandSearchParamsCountryUz BrandSearchParamsCountry = "UZ"
	BrandSearchParamsCountryVa BrandSearchParamsCountry = "VA"
	BrandSearchParamsCountryVc BrandSearchParamsCountry = "VC"
	BrandSearchParamsCountryVe BrandSearchParamsCountry = "VE"
	BrandSearchParamsCountryVg BrandSearchParamsCountry = "VG"
	BrandSearchParamsCountryVi BrandSearchParamsCountry = "VI"
	BrandSearchParamsCountryVn BrandSearchParamsCountry = "VN"
	BrandSearchParamsCountryVu BrandSearchParamsCountry = "VU"
	BrandSearchParamsCountryWf BrandSearchParamsCountry = "WF"
	BrandSearchParamsCountryWs BrandSearchParamsCountry = "WS"
	BrandSearchParamsCountryYe BrandSearchParamsCountry = "YE"
	BrandSearchParamsCountryYt BrandSearchParamsCountry = "YT"
	BrandSearchParamsCountryZa BrandSearchParamsCountry = "ZA"
	BrandSearchParamsCountryZm BrandSearchParamsCountry = "ZM"
	BrandSearchParamsCountryZw BrandSearchParamsCountry = "ZW"
)

// Only return brands with this primary language, such as `english` or `french`.
type BrandSearchParamsLanguage string

const (
	BrandSearchParamsLanguageAfrikaans      BrandSearchParamsLanguage = "afrikaans"
	BrandSearchParamsLanguageAlbanian       BrandSearchParamsLanguage = "albanian"
	BrandSearchParamsLanguageAmharic        BrandSearchParamsLanguage = "amharic"
	BrandSearchParamsLanguageArabic         BrandSearchParamsLanguage = "arabic"
	BrandSearchParamsLanguageArmenian       BrandSearchParamsLanguage = "armenian"
	BrandSearchParamsLanguageAssamese       BrandSearchParamsLanguage = "assamese"
	BrandSearchParamsLanguageAymara         BrandSearchParamsLanguage = "aymara"
	BrandSearchParamsLanguageAzeri          BrandSearchParamsLanguage = "azeri"
	BrandSearchParamsLanguageBasque         BrandSearchParamsLanguage = "basque"
	BrandSearchParamsLanguageBelarusian     BrandSearchParamsLanguage = "belarusian"
	BrandSearchParamsLanguageBengali        BrandSearchParamsLanguage = "bengali"
	BrandSearchParamsLanguageBosnian        BrandSearchParamsLanguage = "bosnian"
	BrandSearchParamsLanguageBulgarian      BrandSearchParamsLanguage = "bulgarian"
	BrandSearchParamsLanguageBurmese        BrandSearchParamsLanguage = "burmese"
	BrandSearchParamsLanguageCantonese      BrandSearchParamsLanguage = "cantonese"
	BrandSearchParamsLanguageCatalan        BrandSearchParamsLanguage = "catalan"
	BrandSearchParamsLanguageCebuano        BrandSearchParamsLanguage = "cebuano"
	BrandSearchParamsLanguageChinese        BrandSearchParamsLanguage = "chinese"
	BrandSearchParamsLanguageCorsican       BrandSearchParamsLanguage = "corsican"
	BrandSearchParamsLanguageCroatian       BrandSearchParamsLanguage = "croatian"
	BrandSearchParamsLanguageCzech          BrandSearchParamsLanguage = "czech"
	BrandSearchParamsLanguageDanish         BrandSearchParamsLanguage = "danish"
	BrandSearchParamsLanguageDutch          BrandSearchParamsLanguage = "dutch"
	BrandSearchParamsLanguageEnglish        BrandSearchParamsLanguage = "english"
	BrandSearchParamsLanguageEsperanto      BrandSearchParamsLanguage = "esperanto"
	BrandSearchParamsLanguageEstonian       BrandSearchParamsLanguage = "estonian"
	BrandSearchParamsLanguageFarsi          BrandSearchParamsLanguage = "farsi"
	BrandSearchParamsLanguageFijian         BrandSearchParamsLanguage = "fijian"
	BrandSearchParamsLanguageFinnish        BrandSearchParamsLanguage = "finnish"
	BrandSearchParamsLanguageFrench         BrandSearchParamsLanguage = "french"
	BrandSearchParamsLanguageGalician       BrandSearchParamsLanguage = "galician"
	BrandSearchParamsLanguageGeorgian       BrandSearchParamsLanguage = "georgian"
	BrandSearchParamsLanguageGerman         BrandSearchParamsLanguage = "german"
	BrandSearchParamsLanguageGreek          BrandSearchParamsLanguage = "greek"
	BrandSearchParamsLanguageGuarani        BrandSearchParamsLanguage = "guarani"
	BrandSearchParamsLanguageGujarati       BrandSearchParamsLanguage = "gujarati"
	BrandSearchParamsLanguageHaitianCreole  BrandSearchParamsLanguage = "haitian-creole"
	BrandSearchParamsLanguageHausa          BrandSearchParamsLanguage = "hausa"
	BrandSearchParamsLanguageHawaiian       BrandSearchParamsLanguage = "hawaiian"
	BrandSearchParamsLanguageHebrew         BrandSearchParamsLanguage = "hebrew"
	BrandSearchParamsLanguageHindi          BrandSearchParamsLanguage = "hindi"
	BrandSearchParamsLanguageHmong          BrandSearchParamsLanguage = "hmong"
	BrandSearchParamsLanguageHungarian      BrandSearchParamsLanguage = "hungarian"
	BrandSearchParamsLanguageIcelandic      BrandSearchParamsLanguage = "icelandic"
	BrandSearchParamsLanguageIgbo           BrandSearchParamsLanguage = "igbo"
	BrandSearchParamsLanguageIndonesian     BrandSearchParamsLanguage = "indonesian"
	BrandSearchParamsLanguageIrish          BrandSearchParamsLanguage = "irish"
	BrandSearchParamsLanguageItalian        BrandSearchParamsLanguage = "italian"
	BrandSearchParamsLanguageJapanese       BrandSearchParamsLanguage = "japanese"
	BrandSearchParamsLanguageJavanese       BrandSearchParamsLanguage = "javanese"
	BrandSearchParamsLanguageKannada        BrandSearchParamsLanguage = "kannada"
	BrandSearchParamsLanguageKazakh         BrandSearchParamsLanguage = "kazakh"
	BrandSearchParamsLanguageKhmer          BrandSearchParamsLanguage = "khmer"
	BrandSearchParamsLanguageKinyarwanda    BrandSearchParamsLanguage = "kinyarwanda"
	BrandSearchParamsLanguageKorean         BrandSearchParamsLanguage = "korean"
	BrandSearchParamsLanguageKurdish        BrandSearchParamsLanguage = "kurdish"
	BrandSearchParamsLanguageKyrgyz         BrandSearchParamsLanguage = "kyrgyz"
	BrandSearchParamsLanguageLao            BrandSearchParamsLanguage = "lao"
	BrandSearchParamsLanguageLatin          BrandSearchParamsLanguage = "latin"
	BrandSearchParamsLanguageLatvian        BrandSearchParamsLanguage = "latvian"
	BrandSearchParamsLanguageLingala        BrandSearchParamsLanguage = "lingala"
	BrandSearchParamsLanguageLithuanian     BrandSearchParamsLanguage = "lithuanian"
	BrandSearchParamsLanguageLuxembourgish  BrandSearchParamsLanguage = "luxembourgish"
	BrandSearchParamsLanguageMacedonian     BrandSearchParamsLanguage = "macedonian"
	BrandSearchParamsLanguageMalagasy       BrandSearchParamsLanguage = "malagasy"
	BrandSearchParamsLanguageMalay          BrandSearchParamsLanguage = "malay"
	BrandSearchParamsLanguageMalayalam      BrandSearchParamsLanguage = "malayalam"
	BrandSearchParamsLanguageMaltese        BrandSearchParamsLanguage = "maltese"
	BrandSearchParamsLanguageMaori          BrandSearchParamsLanguage = "maori"
	BrandSearchParamsLanguageMarathi        BrandSearchParamsLanguage = "marathi"
	BrandSearchParamsLanguageMongolian      BrandSearchParamsLanguage = "mongolian"
	BrandSearchParamsLanguageNepali         BrandSearchParamsLanguage = "nepali"
	BrandSearchParamsLanguageNorwegian      BrandSearchParamsLanguage = "norwegian"
	BrandSearchParamsLanguageOdia           BrandSearchParamsLanguage = "odia"
	BrandSearchParamsLanguageOromo          BrandSearchParamsLanguage = "oromo"
	BrandSearchParamsLanguagePashto         BrandSearchParamsLanguage = "pashto"
	BrandSearchParamsLanguagePidgin         BrandSearchParamsLanguage = "pidgin"
	BrandSearchParamsLanguagePolish         BrandSearchParamsLanguage = "polish"
	BrandSearchParamsLanguagePortuguese     BrandSearchParamsLanguage = "portuguese"
	BrandSearchParamsLanguagePunjabi        BrandSearchParamsLanguage = "punjabi"
	BrandSearchParamsLanguageQuechua        BrandSearchParamsLanguage = "quechua"
	BrandSearchParamsLanguageRomanian       BrandSearchParamsLanguage = "romanian"
	BrandSearchParamsLanguageRussian        BrandSearchParamsLanguage = "russian"
	BrandSearchParamsLanguageSamoan         BrandSearchParamsLanguage = "samoan"
	BrandSearchParamsLanguageScottishGaelic BrandSearchParamsLanguage = "scottish-gaelic"
	BrandSearchParamsLanguageSerbian        BrandSearchParamsLanguage = "serbian"
	BrandSearchParamsLanguageSesotho        BrandSearchParamsLanguage = "sesotho"
	BrandSearchParamsLanguageShona          BrandSearchParamsLanguage = "shona"
	BrandSearchParamsLanguageSindhi         BrandSearchParamsLanguage = "sindhi"
	BrandSearchParamsLanguageSinhala        BrandSearchParamsLanguage = "sinhala"
	BrandSearchParamsLanguageSlovak         BrandSearchParamsLanguage = "slovak"
	BrandSearchParamsLanguageSlovene        BrandSearchParamsLanguage = "slovene"
	BrandSearchParamsLanguageSomali         BrandSearchParamsLanguage = "somali"
	BrandSearchParamsLanguageSpanish        BrandSearchParamsLanguage = "spanish"
	BrandSearchParamsLanguageSundanese      BrandSearchParamsLanguage = "sundanese"
	BrandSearchParamsLanguageSwahili        BrandSearchParamsLanguage = "swahili"
	BrandSearchParamsLanguageSwedish        BrandSearchParamsLanguage = "swedish"
	BrandSearchParamsLanguageTagalog        BrandSearchParamsLanguage = "tagalog"
	BrandSearchParamsLanguageTajik          BrandSearchParamsLanguage = "tajik"
	BrandSearchParamsLanguageTamil          BrandSearchParamsLanguage = "tamil"
	BrandSearchParamsLanguageTatar          BrandSearchParamsLanguage = "tatar"
	BrandSearchParamsLanguageTelugu         BrandSearchParamsLanguage = "telugu"
	BrandSearchParamsLanguageThai           BrandSearchParamsLanguage = "thai"
	BrandSearchParamsLanguageTibetan        BrandSearchParamsLanguage = "tibetan"
	BrandSearchParamsLanguageTigrinya       BrandSearchParamsLanguage = "tigrinya"
	BrandSearchParamsLanguageTongan         BrandSearchParamsLanguage = "tongan"
	BrandSearchParamsLanguageTswana         BrandSearchParamsLanguage = "tswana"
	BrandSearchParamsLanguageTurkish        BrandSearchParamsLanguage = "turkish"
	BrandSearchParamsLanguageTurkmen        BrandSearchParamsLanguage = "turkmen"
	BrandSearchParamsLanguageUkrainian      BrandSearchParamsLanguage = "ukrainian"
	BrandSearchParamsLanguageUrdu           BrandSearchParamsLanguage = "urdu"
	BrandSearchParamsLanguageUyghur         BrandSearchParamsLanguage = "uyghur"
	BrandSearchParamsLanguageUzbek          BrandSearchParamsLanguage = "uzbek"
	BrandSearchParamsLanguageVietnamese     BrandSearchParamsLanguage = "vietnamese"
	BrandSearchParamsLanguageWelsh          BrandSearchParamsLanguage = "welsh"
	BrandSearchParamsLanguageWolof          BrandSearchParamsLanguage = "wolof"
	BrandSearchParamsLanguageXhosa          BrandSearchParamsLanguage = "xhosa"
	BrandSearchParamsLanguageYiddish        BrandSearchParamsLanguage = "yiddish"
	BrandSearchParamsLanguageYoruba         BrandSearchParamsLanguage = "yoruba"
	BrandSearchParamsLanguageZulu           BrandSearchParamsLanguage = "zulu"
)
