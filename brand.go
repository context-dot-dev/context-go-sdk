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

// Retrieve logos, backdrops, colors, industry, description, and more. Provide
// exactly one lookup identifier in the request body: a domain, company name, email
// address, stock ticker, transaction descriptor, or direct URL. Note:
// `by_direct_url` fetches brand data only from the provided URL — not from the
// entire internet.
func (r *BrandService) Get(ctx context.Context, body BrandGetParams, opts ...option.RequestOption) (res *BrandGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "brand/retrieve"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a simplified version of brand data containing only essential
// information: domain, title, colors, logos, and backdrops. Optimized for faster
// responses and reduced data transfer.
func (r *BrandService) GetSimplified(ctx context.Context, query BrandGetSimplifiedParams, opts ...option.RequestOption) (res *BrandGetSimplifiedResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "brand/retrieve-simplified"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type BrandGetResponse struct {
	// Detailed brand information
	Brand BrandGetResponseBrand `json:"brand"`
	// HTTP status code
	Code int64 `json:"code"`
	// Metadata about the API key used for the request. Included in every response
	// whenever a valid API key is provided, even when the response status is not 200.
	KeyMetadata BrandGetResponseKeyMetadata `json:"key_metadata"`
	// Status of the response, e.g., 'ok'
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Brand       respjson.Field
		Code        respjson.Field
		KeyMetadata respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BrandGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detailed brand information
type BrandGetResponseBrand struct {
	// Physical address of the brand
	Address BrandGetResponseBrandAddress `json:"address"`
	// An array of backdrop images for the brand
	Backdrops []BrandGetResponseBrandBackdrop `json:"backdrops"`
	// An array of brand colors
	Colors []BrandGetResponseBrandColor `json:"colors"`
	// A brief description of the brand
	Description string `json:"description"`
	// The domain name of the brand
	Domain string `json:"domain"`
	// Company email address
	Email string `json:"email"`
	// Industry classification information for the brand
	Industries BrandGetResponseBrandIndustries `json:"industries"`
	// Indicates whether the brand content is not safe for work (NSFW)
	IsNsfw bool `json:"is_nsfw"`
	// Important website links for the brand
	Links BrandGetResponseBrandLinks `json:"links"`
	// An array of logos associated with the brand
	Logos []BrandGetResponseBrandLogo `json:"logos"`
	// Company phone number
	Phone string `json:"phone"`
	// The primary language of the brand's website content. Detected from the HTML lang
	// tag, page content analysis, or social media descriptions.
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
	// The brand's slogan
	Slogan string `json:"slogan"`
	// An array of social media links for the brand
	Socials []BrandGetResponseBrandSocial `json:"socials"`
	// Stock market information for this brand (will be null if not a publicly traded
	// company)
	Stock BrandGetResponseBrandStock `json:"stock"`
	// The title or name of the brand
	Title string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address         respjson.Field
		Backdrops       respjson.Field
		Colors          respjson.Field
		Description     respjson.Field
		Domain          respjson.Field
		Email           respjson.Field
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

// Physical address of the brand
type BrandGetResponseBrandAddress struct {
	// City name
	City string `json:"city"`
	// Country name
	Country string `json:"country"`
	// Country code
	CountryCode string `json:"country_code"`
	// Postal or ZIP code
	PostalCode string `json:"postal_code"`
	// State or province code
	StateCode string `json:"state_code"`
	// State or province name
	StateProvince string `json:"state_province"`
	// Street address
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
	// Array of colors in the backdrop image
	Colors []BrandGetResponseBrandBackdropColor `json:"colors"`
	// Resolution of the backdrop image
	Resolution BrandGetResponseBrandBackdropResolution `json:"resolution"`
	// URL of the backdrop image
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
	// Color in hexadecimal format
	Hex string `json:"hex"`
	// Name of the color
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

// Resolution of the backdrop image
type BrandGetResponseBrandBackdropResolution struct {
	// Aspect ratio of the image (width/height)
	AspectRatio float64 `json:"aspect_ratio"`
	// Height of the image in pixels
	Height int64 `json:"height"`
	// Width of the image in pixels
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
	// Color in hexadecimal format
	Hex string `json:"hex"`
	// Name of the color
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
func (r BrandGetResponseBrandColor) RawJSON() string { return r.JSON.raw }
func (r *BrandGetResponseBrandColor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Industry classification information for the brand
type BrandGetResponseBrandIndustries struct {
	// Easy Industry Classification - array of industry and subindustry pairs
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
	// Industry classification enum
	//
	// Any of "Aerospace & Defense", "Technology", "Finance", "Healthcare", "Retail &
	// E-commerce", "Entertainment", "Education", "Government & Nonprofit", "Industrial
	// & Energy", "Automotive & Transportation", "Lifestyle & Leisure", "Luxury &
	// Fashion", "News & Media", "Sports", "Real Estate & PropTech", "Legal &
	// Compliance", "Telecommunications", "Agriculture & Food", "Professional Services
	// & Agencies", "Chemicals & Materials", "Logistics & Supply Chain", "Hospitality &
	// Tourism", "Construction & Built Environment", "Consumer Packaged Goods (CPG)".
	Industry string `json:"industry" api:"required"`
	// Subindustry classification enum
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
	// Entertainment", "Creator Economy & Influencer Platforms", "Advertising, Adtech &
	// Media Buying", "Film, TV & Production Studios", "Events, Venues & Live
	// Entertainment", "Virtual Worlds & Metaverse Experiences", "K-12 Education
	// Platforms & Tools", "Higher Education & University Tech", "Online Learning &
	// MOOCs", "Test Prep & Certification", "Corporate Training & Upskilling",
	// "Tutoring & Supplemental Learning", "Education Management Systems (LMS/SIS)",
	// "Language Learning", "Creator-Led & Cohort-Based Courses", "Special Education &
	// Accessibility Tools", "Government Technology & Digital Services", "Civic
	// Engagement & Policy Platforms", "International Development & Humanitarian Aid",
	// "Philanthropy & Grantmaking", "Nonprofit Operations & Fundraising Tools",
	// "Public Health & Social Services", "Education & Youth Development Programs",
	// "Environmental & Climate Action Organizations", "Legal Aid & Social Justice
	// Advocacy", "Municipal & Infrastructure Services", "Manufacturing & Industrial
	// Automation", "Energy Production (Oil, Gas, Nuclear)", "Renewable Energy &
	// Cleantech", "Utilities & Grid Infrastructure", "Industrial IoT & Monitoring
	// Systems", "Construction & Heavy Equipment", "Mining & Natural Resources",
	// "Environmental Engineering & Sustainability", "Energy Storage & Battery
	// Technology", "Automotive OEMs & Vehicle Manufacturing", "Electric Vehicles (EVs)
	// & Charging Infrastructure", "Mobility-as-a-Service (MaaS)", "Fleet Management",
	// "Public Transit & Urban Mobility", "Autonomous Vehicles & ADAS", "Aftermarket
	// Parts & Services", "Telematics & Vehicle Connectivity", "Aviation & Aerospace
	// Transport", "Maritime Shipping", "Fitness & Wellness", "Beauty & Personal Care",
	// "Home & Living", "Dating & Relationships", "Hobbies, Crafts & DIY", "Outdoor &
	// Recreational Gear", "Events, Experiences & Ticketing Platforms", "Designer &
	// Luxury Apparel", "Accessories, Jewelry & Watches", "Footwear & Leather Goods",
	// "Beauty, Fragrance & Skincare", "Fashion Marketplaces & Retail Platforms",
	// "Sustainable & Ethical Fashion", "Resale, Vintage & Circular Fashion", "Fashion
	// Tech & Virtual Try-Ons", "Streetwear & Emerging Luxury", "Couture &
	// Made-to-Measure", "News Publishing & Journalism", "Digital Media & Content
	// Platforms", "Broadcasting (TV & Radio)", "Podcasting & Audio Media", "News
	// Aggregators & Curation Tools", "Independent & Creator-Led Media", "Newsletters &
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

// Important website links for the brand
type BrandGetResponseBrandLinks struct {
	// URL to the brand's blog or news page
	Blog string `json:"blog" api:"nullable"`
	// URL to the brand's careers or job opportunities page
	Careers string `json:"careers" api:"nullable"`
	// URL to the brand's contact or contact us page
	Contact string `json:"contact" api:"nullable"`
	// URL to the brand's pricing or plans page
	Pricing string `json:"pricing" api:"nullable"`
	// URL to the brand's privacy policy page
	Privacy string `json:"privacy" api:"nullable"`
	// URL to the brand's terms of service or terms and conditions page
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
	// Array of colors in the logo
	Colors []BrandGetResponseBrandLogoColor `json:"colors"`
	// Indicates when this logo is best used: 'light' = best for light mode, 'dark' =
	// best for dark mode, 'has_opaque_background' = can be used for either as image
	// has its own background
	//
	// Any of "light", "dark", "has_opaque_background".
	Mode string `json:"mode"`
	// Resolution of the logo image
	Resolution BrandGetResponseBrandLogoResolution `json:"resolution"`
	// Type of the logo based on resolution (e.g., 'icon', 'logo')
	//
	// Any of "icon", "logo".
	Type string `json:"type"`
	// CDN hosted url of the logo (ready for display)
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
	// Color in hexadecimal format
	Hex string `json:"hex"`
	// Name of the color
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

// Resolution of the logo image
type BrandGetResponseBrandLogoResolution struct {
	// Aspect ratio of the image (width/height)
	AspectRatio float64 `json:"aspect_ratio"`
	// Height of the image in pixels
	Height int64 `json:"height"`
	// Width of the image in pixels
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
	// Type of social media platform
	//
	// Any of "x", "facebook", "instagram", "linkedin", "youtube", "pinterest",
	// "tiktok", "dribbble", "github", "behance", "snapchat", "whatsapp", "telegram",
	// "line", "discord", "twitch", "vimeo", "imdb", "tumblr", "flickr", "giphy",
	// "medium", "spotify", "soundcloud", "tripadvisor", "yelp", "producthunt",
	// "reddit", "crunchbase", "appstore", "playstore".
	Type string `json:"type"`
	// URL of the social media page
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

// Stock market information for this brand (will be null if not a publicly traded
// company)
type BrandGetResponseBrandStock struct {
	// Stock exchange name
	Exchange string `json:"exchange"`
	// Stock ticker symbol
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

// Metadata about the API key used for the request. Included in every response
// whenever a valid API key is provided, even when the response status is not 200.
type BrandGetResponseKeyMetadata struct {
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
func (r BrandGetResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *BrandGetResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetSimplifiedResponse struct {
	// Simplified brand information
	Brand BrandGetSimplifiedResponseBrand `json:"brand"`
	// HTTP status code of the response
	Code int64 `json:"code"`
	// Metadata about the API key used for the request. Included in every response
	// whenever a valid API key is provided, even when the response status is not 200.
	KeyMetadata BrandGetSimplifiedResponseKeyMetadata `json:"key_metadata"`
	// Status of the response, e.g., 'ok'
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Brand       respjson.Field
		Code        respjson.Field
		KeyMetadata respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandGetSimplifiedResponse) RawJSON() string { return r.JSON.raw }
func (r *BrandGetSimplifiedResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Simplified brand information
type BrandGetSimplifiedResponseBrand struct {
	// An array of backdrop images for the brand
	Backdrops []BrandGetSimplifiedResponseBrandBackdrop `json:"backdrops"`
	// An array of brand colors
	Colors []BrandGetSimplifiedResponseBrandColor `json:"colors"`
	// The domain name of the brand
	Domain string `json:"domain"`
	// An array of logos associated with the brand
	Logos []BrandGetSimplifiedResponseBrandLogo `json:"logos"`
	// The title or name of the brand
	Title string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Backdrops   respjson.Field
		Colors      respjson.Field
		Domain      respjson.Field
		Logos       respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandGetSimplifiedResponseBrand) RawJSON() string { return r.JSON.raw }
func (r *BrandGetSimplifiedResponseBrand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetSimplifiedResponseBrandBackdrop struct {
	// Array of colors in the backdrop image
	Colors []BrandGetSimplifiedResponseBrandBackdropColor `json:"colors"`
	// Resolution of the backdrop image
	Resolution BrandGetSimplifiedResponseBrandBackdropResolution `json:"resolution"`
	// URL of the backdrop image
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
func (r BrandGetSimplifiedResponseBrandBackdrop) RawJSON() string { return r.JSON.raw }
func (r *BrandGetSimplifiedResponseBrandBackdrop) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetSimplifiedResponseBrandBackdropColor struct {
	// Color in hexadecimal format
	Hex string `json:"hex"`
	// Name of the color
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
func (r BrandGetSimplifiedResponseBrandBackdropColor) RawJSON() string { return r.JSON.raw }
func (r *BrandGetSimplifiedResponseBrandBackdropColor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resolution of the backdrop image
type BrandGetSimplifiedResponseBrandBackdropResolution struct {
	// Aspect ratio of the image (width/height)
	AspectRatio float64 `json:"aspect_ratio"`
	// Height of the image in pixels
	Height int64 `json:"height"`
	// Width of the image in pixels
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
func (r BrandGetSimplifiedResponseBrandBackdropResolution) RawJSON() string { return r.JSON.raw }
func (r *BrandGetSimplifiedResponseBrandBackdropResolution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetSimplifiedResponseBrandColor struct {
	// Color in hexadecimal format
	Hex string `json:"hex"`
	// Name of the color
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
func (r BrandGetSimplifiedResponseBrandColor) RawJSON() string { return r.JSON.raw }
func (r *BrandGetSimplifiedResponseBrandColor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetSimplifiedResponseBrandLogo struct {
	// Array of colors in the logo
	Colors []BrandGetSimplifiedResponseBrandLogoColor `json:"colors"`
	// Indicates when this logo is best used: 'light' = best for light mode, 'dark' =
	// best for dark mode, 'has_opaque_background' = can be used for either as image
	// has its own background
	//
	// Any of "light", "dark", "has_opaque_background".
	Mode string `json:"mode"`
	// Resolution of the logo image
	Resolution BrandGetSimplifiedResponseBrandLogoResolution `json:"resolution"`
	// Type of the logo based on resolution (e.g., 'icon', 'logo')
	//
	// Any of "icon", "logo".
	Type string `json:"type"`
	// CDN hosted url of the logo (ready for display)
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
func (r BrandGetSimplifiedResponseBrandLogo) RawJSON() string { return r.JSON.raw }
func (r *BrandGetSimplifiedResponseBrandLogo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetSimplifiedResponseBrandLogoColor struct {
	// Color in hexadecimal format
	Hex string `json:"hex"`
	// Name of the color
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
func (r BrandGetSimplifiedResponseBrandLogoColor) RawJSON() string { return r.JSON.raw }
func (r *BrandGetSimplifiedResponseBrandLogoColor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resolution of the logo image
type BrandGetSimplifiedResponseBrandLogoResolution struct {
	// Aspect ratio of the image (width/height)
	AspectRatio float64 `json:"aspect_ratio"`
	// Height of the image in pixels
	Height int64 `json:"height"`
	// Width of the image in pixels
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
func (r BrandGetSimplifiedResponseBrandLogoResolution) RawJSON() string { return r.JSON.raw }
func (r *BrandGetSimplifiedResponseBrandLogoResolution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the API key used for the request. Included in every response
// whenever a valid API key is provided, even when the response status is not 200.
type BrandGetSimplifiedResponseKeyMetadata struct {
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
func (r BrandGetSimplifiedResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *BrandGetSimplifiedResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	// Retrieve brand data by domain. Cannot be combined with name, email, or ticker.
	OfByDomain *BrandGetParamsBodyByDomain `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Retrieve brand data by company name. Cannot be combined with domain, email, or
	// ticker.
	OfByName *BrandGetParamsBodyByName `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Retrieve brand data by email address. The domain is extracted from the email.
	// Free and disposable email providers are rejected with 422. Cannot be combined
	// with domain, name, or ticker.
	OfByEmail *BrandGetParamsBodyByEmail `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Retrieve brand data by stock ticker. Cannot be combined with domain, name, or
	// email.
	OfByTicker *BrandGetParamsBodyByTicker `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Retrieve brand data by fetching the provided URL directly. Note: if you use
	// this, brand data is fetched only from the provided URL — not from the entire
	// internet — so results are limited to what that single page contains. No domain
	// resolution, database lookup, or cross-source enrichment is performed. Cannot be
	// combined with domain, name, email, or ticker.
	OfByDirectURL *BrandGetParamsBodyByDirectURL `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Identify brand data from a transaction descriptor. Cannot be combined with
	// domain, name, email, or ticker.
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

// Retrieve brand data by domain. Cannot be combined with name, email, or ticker.
//
// The properties Domain, Type are required.
type BrandGetParamsBodyByDomain struct {
	// Domain name to retrieve brand data for (e.g., 'stripe.com').
	Domain string `json:"domain" api:"required"`
	// Maximum age in milliseconds for cached brand data before the API performs a hard
	// refresh. Defaults to 3 months (7776000000 ms). Values below 1 day (86400000 ms)
	// are clamped to 1 day; values above 1 year (31536000000 ms) are clamped to 1
	// year.
	MaxAgeMs param.Opt[int64] `json:"maxAgeMs,omitzero"`
	// Optional parameter to optimize the API call for maximum speed. When set to true,
	// the API will skip time-consuming operations for faster response at the cost of
	// less comprehensive data.
	MaxSpeed param.Opt[bool] `json:"maxSpeed,omitzero"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `json:"timeoutMS,omitzero"`
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
	// Optional caller-defined tags for tracking this request. Tags are recorded on the
	// request's usage log and can be used to filter usage on the dashboard usage page.
	// Up to 20 tags, each 1-50 characters.
	Tags []string `json:"tags,omitzero"`
	// Discriminator for domain-based brand retrieval.
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
}

// Retrieve brand data by company name. Cannot be combined with domain, email, or
// ticker.
//
// The properties Name, Type are required.
type BrandGetParamsBodyByName struct {
	// Company name to retrieve brand data for (e.g., 'Apple Inc').
	Name string `json:"name" api:"required"`
	// Optional country code hint (GL parameter) to specify the country when looking up
	// by company name.
	CountryGl param.Opt[string] `json:"country_gl,omitzero"`
	// Maximum age in milliseconds for cached brand data before the API performs a hard
	// refresh. Defaults to 3 months (7776000000 ms). Values below 1 day (86400000 ms)
	// are clamped to 1 day; values above 1 year (31536000000 ms) are clamped to 1
	// year.
	MaxAgeMs param.Opt[int64] `json:"maxAgeMs,omitzero"`
	// Optional parameter to optimize the API call for maximum speed. When set to true,
	// the API will skip time-consuming operations for faster response at the cost of
	// less comprehensive data.
	MaxSpeed param.Opt[bool] `json:"maxSpeed,omitzero"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `json:"timeoutMS,omitzero"`
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
	// Optional caller-defined tags for tracking this request. Tags are recorded on the
	// request's usage log and can be used to filter usage on the dashboard usage page.
	// Up to 20 tags, each 1-50 characters.
	Tags []string `json:"tags,omitzero"`
	// Discriminator for name-based brand retrieval.
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
		"force_language", "afrikaans", "albanian", "amharic", "arabic", "armenian", "assamese", "aymara", "azeri", "basque", "belarusian", "bengali", "bosnian", "bulgarian", "burmese", "cantonese", "catalan", "cebuano", "chinese", "corsican", "croatian", "czech", "danish", "dutch", "english", "esperanto", "estonian", "farsi", "fijian", "finnish", "french", "galician", "georgian", "german", "greek", "guarani", "gujarati", "haitian-creole", "hausa", "hawaiian", "hebrew", "hindi", "hmong", "hungarian", "icelandic", "igbo", "indonesian", "irish", "italian", "japanese", "javanese", "kannada", "kazakh", "khmer", "kinyarwanda", "korean", "kurdish", "kyrgyz", "lao", "latin", "latvian", "lingala", "lithuanian", "luxembourgish", "macedonian", "malagasy", "malay", "malayalam", "maltese", "maori", "marathi", "mongolian", "nepali", "norwegian", "odia", "oromo", "pashto", "pidgin", "polish", "portuguese", "punjabi", "quechua", "romanian", "russian", "samoan", "scottish-gaelic", "serbian", "sesotho", "shona", "sindhi", "sinhala", "slovak", "slovene", "somali", "spanish", "sundanese", "swahili", "swedish", "tagalog", "tajik", "tamil", "tatar", "telugu", "thai", "tibetan", "tigrinya", "tongan", "tswana", "turkish", "turkmen", "ukrainian", "urdu", "uyghur", "uzbek", "vietnamese", "welsh", "wolof", "xhosa", "yiddish", "yoruba", "zulu",
	)
}

// Retrieve brand data by email address. The domain is extracted from the email.
// Free and disposable email providers are rejected with 422. Cannot be combined
// with domain, name, or ticker.
//
// The properties Email, Type are required.
type BrandGetParamsBodyByEmail struct {
	// Email address to retrieve brand data for (e.g., 'jane@stripe.com').
	Email string `json:"email" api:"required" format:"email"`
	// Maximum age in milliseconds for cached brand data before the API performs a hard
	// refresh. Defaults to 3 months (7776000000 ms). Values below 1 day (86400000 ms)
	// are clamped to 1 day; values above 1 year (31536000000 ms) are clamped to 1
	// year.
	MaxAgeMs param.Opt[int64] `json:"maxAgeMs,omitzero"`
	// Optional parameter to optimize the API call for maximum speed. When set to true,
	// the API will skip time-consuming operations for faster response at the cost of
	// less comprehensive data.
	MaxSpeed param.Opt[bool] `json:"maxSpeed,omitzero"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `json:"timeoutMS,omitzero"`
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
	// Optional caller-defined tags for tracking this request. Tags are recorded on the
	// request's usage log and can be used to filter usage on the dashboard usage page.
	// Up to 20 tags, each 1-50 characters.
	Tags []string `json:"tags,omitzero"`
	// Discriminator for email-based brand retrieval.
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
}

// Retrieve brand data by stock ticker. Cannot be combined with domain, name, or
// email.
//
// The properties Ticker, Type are required.
type BrandGetParamsBodyByTicker struct {
	// Stock ticker symbol to retrieve brand data for (e.g., 'AAPL').
	Ticker string `json:"ticker" api:"required"`
	// Maximum age in milliseconds for cached brand data before the API performs a hard
	// refresh. Defaults to 3 months (7776000000 ms). Values below 1 day (86400000 ms)
	// are clamped to 1 day; values above 1 year (31536000000 ms) are clamped to 1
	// year.
	MaxAgeMs param.Opt[int64] `json:"maxAgeMs,omitzero"`
	// Optional parameter to optimize the API call for maximum speed. When set to true,
	// the API will skip time-consuming operations for faster response at the cost of
	// less comprehensive data.
	MaxSpeed param.Opt[bool] `json:"maxSpeed,omitzero"`
	// Optional stock exchange for the ticker. Defaults to NASDAQ if not specified.
	TickerExchange param.Opt[string] `json:"ticker_exchange,omitzero"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `json:"timeoutMS,omitzero"`
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
	// Optional caller-defined tags for tracking this request. Tags are recorded on the
	// request's usage log and can be used to filter usage on the dashboard usage page.
	// Up to 20 tags, each 1-50 characters.
	Tags []string `json:"tags,omitzero"`
	// Discriminator for ticker-based brand retrieval.
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
}

// Retrieve brand data by fetching the provided URL directly. Note: if you use
// this, brand data is fetched only from the provided URL — not from the entire
// internet — so results are limited to what that single page contains. No domain
// resolution, database lookup, or cross-source enrichment is performed. Cannot be
// combined with domain, name, email, or ticker.
//
// The properties DirectURL, Type are required.
type BrandGetParamsBodyByDirectURL struct {
	// Full http(s) URL to fetch brand data from (e.g.,
	// 'https://stripe.com/enterprise'). Only this URL is fetched — not the entire
	// internet.
	DirectURL string `json:"direct_url" api:"required" format:"uri"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `json:"timeoutMS,omitzero"`
	// Optional caller-defined tags for tracking this request. Tags are recorded on the
	// request's usage log and can be used to filter usage on the dashboard usage page.
	// Up to 20 tags, each 1-50 characters.
	Tags []string `json:"tags,omitzero"`
	// Discriminator for direct-URL-based brand retrieval.
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

// Identify brand data from a transaction descriptor. Cannot be combined with
// domain, name, email, or ticker.
//
// The properties TransactionInfo, Type are required.
type BrandGetParamsBodyByTransaction struct {
	// Transaction information to identify the brand.
	TransactionInfo string `json:"transaction_info" api:"required"`
	// Optional city name to prioritize when searching for the brand.
	City param.Opt[string] `json:"city,omitzero"`
	// Optional country code hint (GL parameter) to specify the country when
	// identifying a transaction.
	CountryGl param.Opt[string] `json:"country_gl,omitzero"`
	// When set to true, the API performs additional verification to ensure the
	// identified brand matches the transaction with high confidence.
	HighConfidenceOnly param.Opt[bool] `json:"high_confidence_only,omitzero"`
	// Optional parameter to optimize the API call for maximum speed. When set to true,
	// the API will skip time-consuming operations for faster response at the cost of
	// less comprehensive data.
	MaxSpeed param.Opt[bool] `json:"maxSpeed,omitzero"`
	// Optional Merchant Category Code (MCC) to help identify the business category or
	// industry.
	Mcc param.Opt[int64] `json:"mcc,omitzero"`
	// Optional phone number from the transaction to help verify brand match.
	Phone param.Opt[float64] `json:"phone,omitzero"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `json:"timeoutMS,omitzero"`
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
	// Optional caller-defined tags for tracking this request. Tags are recorded on the
	// request's usage log and can be used to filter usage on the dashboard usage page.
	// Up to 20 tags, each 1-50 characters.
	Tags []string `json:"tags,omitzero"`
	// Discriminator for transaction-based brand retrieval.
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
		"force_language", "afrikaans", "albanian", "amharic", "arabic", "armenian", "assamese", "aymara", "azeri", "basque", "belarusian", "bengali", "bosnian", "bulgarian", "burmese", "cantonese", "catalan", "cebuano", "chinese", "corsican", "croatian", "czech", "danish", "dutch", "english", "esperanto", "estonian", "farsi", "fijian", "finnish", "french", "galician", "georgian", "german", "greek", "guarani", "gujarati", "haitian-creole", "hausa", "hawaiian", "hebrew", "hindi", "hmong", "hungarian", "icelandic", "igbo", "indonesian", "irish", "italian", "japanese", "javanese", "kannada", "kazakh", "khmer", "kinyarwanda", "korean", "kurdish", "kyrgyz", "lao", "latin", "latvian", "lingala", "lithuanian", "luxembourgish", "macedonian", "malagasy", "malay", "malayalam", "maltese", "maori", "marathi", "mongolian", "nepali", "norwegian", "odia", "oromo", "pashto", "pidgin", "polish", "portuguese", "punjabi", "quechua", "romanian", "russian", "samoan", "scottish-gaelic", "serbian", "sesotho", "shona", "sindhi", "sinhala", "slovak", "slovene", "somali", "spanish", "sundanese", "swahili", "swedish", "tagalog", "tajik", "tamil", "tatar", "telugu", "thai", "tibetan", "tigrinya", "tongan", "tswana", "turkish", "turkmen", "ukrainian", "urdu", "uyghur", "uzbek", "vietnamese", "welsh", "wolof", "xhosa", "yiddish", "yoruba", "zulu",
	)
}

type BrandGetSimplifiedParams struct {
	// Domain name to retrieve simplified brand data for
	Domain string `query:"domain" api:"required" json:"-"`
	// Maximum age in milliseconds for cached brand data before the API performs a hard
	// refresh. Defaults to 3 months (7776000000 ms). Values below 1 day (86400000 ms)
	// are clamped to 1 day; values above 1 year (31536000000 ms) are clamped to 1
	// year.
	MaxAgeMs param.Opt[int64] `query:"maxAgeMs,omitzero" json:"-"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `query:"timeoutMS,omitzero" json:"-"`
	// Optional comma-separated caller-defined tags for tracking this request. Tags are
	// recorded on the request's usage log and can be used to filter usage on the
	// dashboard usage page. Up to 20 tags, each 1-50 characters.
	Tags []string `query:"tags,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BrandGetSimplifiedParams]'s query parameters as
// `url.Values`.
func (r BrandGetSimplifiedParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
