// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contextdev

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/context.dev-go/internal/apijson"
	"github.com/stainless-sdks/context.dev-go/internal/apiquery"
	"github.com/stainless-sdks/context.dev-go/internal/requestconfig"
	"github.com/stainless-sdks/context.dev-go/option"
	"github.com/stainless-sdks/context.dev-go/packages/param"
	"github.com/stainless-sdks/context.dev-go/packages/respjson"
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

// Retrieve logos, backdrops, colors, industry, description, and more from any
// domain
func (r *BrandService) Get(ctx context.Context, query BrandGetParams, opts ...option.RequestOption) (res *BrandGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "brand/retrieve"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Endpoint specially designed for platforms that want to identify transaction data
// by the transaction title.
func (r *BrandService) IdentifyFromTransaction(ctx context.Context, query BrandIdentifyFromTransactionParams, opts ...option.RequestOption) (res *BrandIdentifyFromTransactionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "brand/transaction_identifier"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve brand information using an email address while detecting disposable and
// free email addresses. Disposable and free email addresses (like gmail.com,
// yahoo.com) will throw a 422 error.
func (r *BrandService) GetByEmail(ctx context.Context, query BrandGetByEmailParams, opts ...option.RequestOption) (res *BrandGetByEmailResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "brand/retrieve-by-email"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve brand information using an ISIN (International Securities
// Identification Number).
func (r *BrandService) GetByIsin(ctx context.Context, query BrandGetByIsinParams, opts ...option.RequestOption) (res *BrandGetByIsinResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "brand/retrieve-by-isin"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve brand information using a company name.
func (r *BrandService) GetByName(ctx context.Context, query BrandGetByNameParams, opts ...option.RequestOption) (res *BrandGetByNameResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "brand/retrieve-by-name"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve brand information using a stock ticker symbol.
func (r *BrandService) GetByTicker(ctx context.Context, query BrandGetByTickerParams, opts ...option.RequestOption) (res *BrandGetByTickerResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "brand/retrieve-by-ticker"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
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
	// Status of the response, e.g., 'ok'
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Brand       respjson.Field
		Code        respjson.Field
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

type BrandIdentifyFromTransactionResponse struct {
	// Detailed brand information
	Brand BrandIdentifyFromTransactionResponseBrand `json:"brand"`
	// HTTP status code
	Code int64 `json:"code"`
	// Status of the response, e.g., 'ok'
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Brand       respjson.Field
		Code        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandIdentifyFromTransactionResponse) RawJSON() string { return r.JSON.raw }
func (r *BrandIdentifyFromTransactionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detailed brand information
type BrandIdentifyFromTransactionResponseBrand struct {
	// Physical address of the brand
	Address BrandIdentifyFromTransactionResponseBrandAddress `json:"address"`
	// An array of backdrop images for the brand
	Backdrops []BrandIdentifyFromTransactionResponseBrandBackdrop `json:"backdrops"`
	// An array of brand colors
	Colors []BrandIdentifyFromTransactionResponseBrandColor `json:"colors"`
	// A brief description of the brand
	Description string `json:"description"`
	// The domain name of the brand
	Domain string `json:"domain"`
	// Company email address
	Email string `json:"email"`
	// Industry classification information for the brand
	Industries BrandIdentifyFromTransactionResponseBrandIndustries `json:"industries"`
	// Indicates whether the brand content is not safe for work (NSFW)
	IsNsfw bool `json:"is_nsfw"`
	// Important website links for the brand
	Links BrandIdentifyFromTransactionResponseBrandLinks `json:"links"`
	// An array of logos associated with the brand
	Logos []BrandIdentifyFromTransactionResponseBrandLogo `json:"logos"`
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
	Socials []BrandIdentifyFromTransactionResponseBrandSocial `json:"socials"`
	// Stock market information for this brand (will be null if not a publicly traded
	// company)
	Stock BrandIdentifyFromTransactionResponseBrandStock `json:"stock"`
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
func (r BrandIdentifyFromTransactionResponseBrand) RawJSON() string { return r.JSON.raw }
func (r *BrandIdentifyFromTransactionResponseBrand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Physical address of the brand
type BrandIdentifyFromTransactionResponseBrandAddress struct {
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
func (r BrandIdentifyFromTransactionResponseBrandAddress) RawJSON() string { return r.JSON.raw }
func (r *BrandIdentifyFromTransactionResponseBrandAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandIdentifyFromTransactionResponseBrandBackdrop struct {
	// Array of colors in the backdrop image
	Colors []BrandIdentifyFromTransactionResponseBrandBackdropColor `json:"colors"`
	// Resolution of the backdrop image
	Resolution BrandIdentifyFromTransactionResponseBrandBackdropResolution `json:"resolution"`
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
func (r BrandIdentifyFromTransactionResponseBrandBackdrop) RawJSON() string { return r.JSON.raw }
func (r *BrandIdentifyFromTransactionResponseBrandBackdrop) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandIdentifyFromTransactionResponseBrandBackdropColor struct {
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
func (r BrandIdentifyFromTransactionResponseBrandBackdropColor) RawJSON() string { return r.JSON.raw }
func (r *BrandIdentifyFromTransactionResponseBrandBackdropColor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resolution of the backdrop image
type BrandIdentifyFromTransactionResponseBrandBackdropResolution struct {
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
func (r BrandIdentifyFromTransactionResponseBrandBackdropResolution) RawJSON() string {
	return r.JSON.raw
}
func (r *BrandIdentifyFromTransactionResponseBrandBackdropResolution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandIdentifyFromTransactionResponseBrandColor struct {
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
func (r BrandIdentifyFromTransactionResponseBrandColor) RawJSON() string { return r.JSON.raw }
func (r *BrandIdentifyFromTransactionResponseBrandColor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Industry classification information for the brand
type BrandIdentifyFromTransactionResponseBrandIndustries struct {
	// Easy Industry Classification - array of industry and subindustry pairs
	Eic []BrandIdentifyFromTransactionResponseBrandIndustriesEic `json:"eic"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Eic         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandIdentifyFromTransactionResponseBrandIndustries) RawJSON() string { return r.JSON.raw }
func (r *BrandIdentifyFromTransactionResponseBrandIndustries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandIdentifyFromTransactionResponseBrandIndustriesEic struct {
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
func (r BrandIdentifyFromTransactionResponseBrandIndustriesEic) RawJSON() string { return r.JSON.raw }
func (r *BrandIdentifyFromTransactionResponseBrandIndustriesEic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Important website links for the brand
type BrandIdentifyFromTransactionResponseBrandLinks struct {
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
func (r BrandIdentifyFromTransactionResponseBrandLinks) RawJSON() string { return r.JSON.raw }
func (r *BrandIdentifyFromTransactionResponseBrandLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandIdentifyFromTransactionResponseBrandLogo struct {
	// Array of colors in the logo
	Colors []BrandIdentifyFromTransactionResponseBrandLogoColor `json:"colors"`
	// Indicates when this logo is best used: 'light' = best for light mode, 'dark' =
	// best for dark mode, 'has_opaque_background' = can be used for either as image
	// has its own background
	//
	// Any of "light", "dark", "has_opaque_background".
	Mode string `json:"mode"`
	// Resolution of the logo image
	Resolution BrandIdentifyFromTransactionResponseBrandLogoResolution `json:"resolution"`
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
func (r BrandIdentifyFromTransactionResponseBrandLogo) RawJSON() string { return r.JSON.raw }
func (r *BrandIdentifyFromTransactionResponseBrandLogo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandIdentifyFromTransactionResponseBrandLogoColor struct {
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
func (r BrandIdentifyFromTransactionResponseBrandLogoColor) RawJSON() string { return r.JSON.raw }
func (r *BrandIdentifyFromTransactionResponseBrandLogoColor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resolution of the logo image
type BrandIdentifyFromTransactionResponseBrandLogoResolution struct {
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
func (r BrandIdentifyFromTransactionResponseBrandLogoResolution) RawJSON() string { return r.JSON.raw }
func (r *BrandIdentifyFromTransactionResponseBrandLogoResolution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandIdentifyFromTransactionResponseBrandSocial struct {
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
func (r BrandIdentifyFromTransactionResponseBrandSocial) RawJSON() string { return r.JSON.raw }
func (r *BrandIdentifyFromTransactionResponseBrandSocial) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stock market information for this brand (will be null if not a publicly traded
// company)
type BrandIdentifyFromTransactionResponseBrandStock struct {
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
func (r BrandIdentifyFromTransactionResponseBrandStock) RawJSON() string { return r.JSON.raw }
func (r *BrandIdentifyFromTransactionResponseBrandStock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetByEmailResponse struct {
	// Detailed brand information
	Brand BrandGetByEmailResponseBrand `json:"brand"`
	// HTTP status code
	Code int64 `json:"code"`
	// Status of the response, e.g., 'ok'
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Brand       respjson.Field
		Code        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandGetByEmailResponse) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByEmailResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detailed brand information
type BrandGetByEmailResponseBrand struct {
	// Physical address of the brand
	Address BrandGetByEmailResponseBrandAddress `json:"address"`
	// An array of backdrop images for the brand
	Backdrops []BrandGetByEmailResponseBrandBackdrop `json:"backdrops"`
	// An array of brand colors
	Colors []BrandGetByEmailResponseBrandColor `json:"colors"`
	// A brief description of the brand
	Description string `json:"description"`
	// The domain name of the brand
	Domain string `json:"domain"`
	// Company email address
	Email string `json:"email"`
	// Industry classification information for the brand
	Industries BrandGetByEmailResponseBrandIndustries `json:"industries"`
	// Indicates whether the brand content is not safe for work (NSFW)
	IsNsfw bool `json:"is_nsfw"`
	// Important website links for the brand
	Links BrandGetByEmailResponseBrandLinks `json:"links"`
	// An array of logos associated with the brand
	Logos []BrandGetByEmailResponseBrandLogo `json:"logos"`
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
	Socials []BrandGetByEmailResponseBrandSocial `json:"socials"`
	// Stock market information for this brand (will be null if not a publicly traded
	// company)
	Stock BrandGetByEmailResponseBrandStock `json:"stock"`
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
func (r BrandGetByEmailResponseBrand) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByEmailResponseBrand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Physical address of the brand
type BrandGetByEmailResponseBrandAddress struct {
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
func (r BrandGetByEmailResponseBrandAddress) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByEmailResponseBrandAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetByEmailResponseBrandBackdrop struct {
	// Array of colors in the backdrop image
	Colors []BrandGetByEmailResponseBrandBackdropColor `json:"colors"`
	// Resolution of the backdrop image
	Resolution BrandGetByEmailResponseBrandBackdropResolution `json:"resolution"`
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
func (r BrandGetByEmailResponseBrandBackdrop) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByEmailResponseBrandBackdrop) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetByEmailResponseBrandBackdropColor struct {
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
func (r BrandGetByEmailResponseBrandBackdropColor) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByEmailResponseBrandBackdropColor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resolution of the backdrop image
type BrandGetByEmailResponseBrandBackdropResolution struct {
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
func (r BrandGetByEmailResponseBrandBackdropResolution) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByEmailResponseBrandBackdropResolution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetByEmailResponseBrandColor struct {
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
func (r BrandGetByEmailResponseBrandColor) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByEmailResponseBrandColor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Industry classification information for the brand
type BrandGetByEmailResponseBrandIndustries struct {
	// Easy Industry Classification - array of industry and subindustry pairs
	Eic []BrandGetByEmailResponseBrandIndustriesEic `json:"eic"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Eic         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandGetByEmailResponseBrandIndustries) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByEmailResponseBrandIndustries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetByEmailResponseBrandIndustriesEic struct {
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
func (r BrandGetByEmailResponseBrandIndustriesEic) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByEmailResponseBrandIndustriesEic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Important website links for the brand
type BrandGetByEmailResponseBrandLinks struct {
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
func (r BrandGetByEmailResponseBrandLinks) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByEmailResponseBrandLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetByEmailResponseBrandLogo struct {
	// Array of colors in the logo
	Colors []BrandGetByEmailResponseBrandLogoColor `json:"colors"`
	// Indicates when this logo is best used: 'light' = best for light mode, 'dark' =
	// best for dark mode, 'has_opaque_background' = can be used for either as image
	// has its own background
	//
	// Any of "light", "dark", "has_opaque_background".
	Mode string `json:"mode"`
	// Resolution of the logo image
	Resolution BrandGetByEmailResponseBrandLogoResolution `json:"resolution"`
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
func (r BrandGetByEmailResponseBrandLogo) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByEmailResponseBrandLogo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetByEmailResponseBrandLogoColor struct {
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
func (r BrandGetByEmailResponseBrandLogoColor) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByEmailResponseBrandLogoColor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resolution of the logo image
type BrandGetByEmailResponseBrandLogoResolution struct {
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
func (r BrandGetByEmailResponseBrandLogoResolution) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByEmailResponseBrandLogoResolution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetByEmailResponseBrandSocial struct {
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
func (r BrandGetByEmailResponseBrandSocial) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByEmailResponseBrandSocial) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stock market information for this brand (will be null if not a publicly traded
// company)
type BrandGetByEmailResponseBrandStock struct {
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
func (r BrandGetByEmailResponseBrandStock) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByEmailResponseBrandStock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetByIsinResponse struct {
	// Detailed brand information
	Brand BrandGetByIsinResponseBrand `json:"brand"`
	// HTTP status code
	Code int64 `json:"code"`
	// Status of the response, e.g., 'ok'
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Brand       respjson.Field
		Code        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandGetByIsinResponse) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByIsinResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detailed brand information
type BrandGetByIsinResponseBrand struct {
	// Physical address of the brand
	Address BrandGetByIsinResponseBrandAddress `json:"address"`
	// An array of backdrop images for the brand
	Backdrops []BrandGetByIsinResponseBrandBackdrop `json:"backdrops"`
	// An array of brand colors
	Colors []BrandGetByIsinResponseBrandColor `json:"colors"`
	// A brief description of the brand
	Description string `json:"description"`
	// The domain name of the brand
	Domain string `json:"domain"`
	// Company email address
	Email string `json:"email"`
	// Industry classification information for the brand
	Industries BrandGetByIsinResponseBrandIndustries `json:"industries"`
	// Indicates whether the brand content is not safe for work (NSFW)
	IsNsfw bool `json:"is_nsfw"`
	// Important website links for the brand
	Links BrandGetByIsinResponseBrandLinks `json:"links"`
	// An array of logos associated with the brand
	Logos []BrandGetByIsinResponseBrandLogo `json:"logos"`
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
	Socials []BrandGetByIsinResponseBrandSocial `json:"socials"`
	// Stock market information for this brand (will be null if not a publicly traded
	// company)
	Stock BrandGetByIsinResponseBrandStock `json:"stock"`
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
func (r BrandGetByIsinResponseBrand) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByIsinResponseBrand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Physical address of the brand
type BrandGetByIsinResponseBrandAddress struct {
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
func (r BrandGetByIsinResponseBrandAddress) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByIsinResponseBrandAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetByIsinResponseBrandBackdrop struct {
	// Array of colors in the backdrop image
	Colors []BrandGetByIsinResponseBrandBackdropColor `json:"colors"`
	// Resolution of the backdrop image
	Resolution BrandGetByIsinResponseBrandBackdropResolution `json:"resolution"`
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
func (r BrandGetByIsinResponseBrandBackdrop) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByIsinResponseBrandBackdrop) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetByIsinResponseBrandBackdropColor struct {
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
func (r BrandGetByIsinResponseBrandBackdropColor) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByIsinResponseBrandBackdropColor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resolution of the backdrop image
type BrandGetByIsinResponseBrandBackdropResolution struct {
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
func (r BrandGetByIsinResponseBrandBackdropResolution) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByIsinResponseBrandBackdropResolution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetByIsinResponseBrandColor struct {
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
func (r BrandGetByIsinResponseBrandColor) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByIsinResponseBrandColor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Industry classification information for the brand
type BrandGetByIsinResponseBrandIndustries struct {
	// Easy Industry Classification - array of industry and subindustry pairs
	Eic []BrandGetByIsinResponseBrandIndustriesEic `json:"eic"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Eic         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandGetByIsinResponseBrandIndustries) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByIsinResponseBrandIndustries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetByIsinResponseBrandIndustriesEic struct {
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
func (r BrandGetByIsinResponseBrandIndustriesEic) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByIsinResponseBrandIndustriesEic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Important website links for the brand
type BrandGetByIsinResponseBrandLinks struct {
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
func (r BrandGetByIsinResponseBrandLinks) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByIsinResponseBrandLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetByIsinResponseBrandLogo struct {
	// Array of colors in the logo
	Colors []BrandGetByIsinResponseBrandLogoColor `json:"colors"`
	// Indicates when this logo is best used: 'light' = best for light mode, 'dark' =
	// best for dark mode, 'has_opaque_background' = can be used for either as image
	// has its own background
	//
	// Any of "light", "dark", "has_opaque_background".
	Mode string `json:"mode"`
	// Resolution of the logo image
	Resolution BrandGetByIsinResponseBrandLogoResolution `json:"resolution"`
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
func (r BrandGetByIsinResponseBrandLogo) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByIsinResponseBrandLogo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetByIsinResponseBrandLogoColor struct {
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
func (r BrandGetByIsinResponseBrandLogoColor) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByIsinResponseBrandLogoColor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resolution of the logo image
type BrandGetByIsinResponseBrandLogoResolution struct {
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
func (r BrandGetByIsinResponseBrandLogoResolution) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByIsinResponseBrandLogoResolution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetByIsinResponseBrandSocial struct {
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
func (r BrandGetByIsinResponseBrandSocial) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByIsinResponseBrandSocial) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stock market information for this brand (will be null if not a publicly traded
// company)
type BrandGetByIsinResponseBrandStock struct {
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
func (r BrandGetByIsinResponseBrandStock) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByIsinResponseBrandStock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetByNameResponse struct {
	// Detailed brand information
	Brand BrandGetByNameResponseBrand `json:"brand"`
	// HTTP status code
	Code int64 `json:"code"`
	// Status of the response, e.g., 'ok'
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Brand       respjson.Field
		Code        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandGetByNameResponse) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByNameResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detailed brand information
type BrandGetByNameResponseBrand struct {
	// Physical address of the brand
	Address BrandGetByNameResponseBrandAddress `json:"address"`
	// An array of backdrop images for the brand
	Backdrops []BrandGetByNameResponseBrandBackdrop `json:"backdrops"`
	// An array of brand colors
	Colors []BrandGetByNameResponseBrandColor `json:"colors"`
	// A brief description of the brand
	Description string `json:"description"`
	// The domain name of the brand
	Domain string `json:"domain"`
	// Company email address
	Email string `json:"email"`
	// Industry classification information for the brand
	Industries BrandGetByNameResponseBrandIndustries `json:"industries"`
	// Indicates whether the brand content is not safe for work (NSFW)
	IsNsfw bool `json:"is_nsfw"`
	// Important website links for the brand
	Links BrandGetByNameResponseBrandLinks `json:"links"`
	// An array of logos associated with the brand
	Logos []BrandGetByNameResponseBrandLogo `json:"logos"`
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
	Socials []BrandGetByNameResponseBrandSocial `json:"socials"`
	// Stock market information for this brand (will be null if not a publicly traded
	// company)
	Stock BrandGetByNameResponseBrandStock `json:"stock"`
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
func (r BrandGetByNameResponseBrand) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByNameResponseBrand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Physical address of the brand
type BrandGetByNameResponseBrandAddress struct {
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
func (r BrandGetByNameResponseBrandAddress) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByNameResponseBrandAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetByNameResponseBrandBackdrop struct {
	// Array of colors in the backdrop image
	Colors []BrandGetByNameResponseBrandBackdropColor `json:"colors"`
	// Resolution of the backdrop image
	Resolution BrandGetByNameResponseBrandBackdropResolution `json:"resolution"`
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
func (r BrandGetByNameResponseBrandBackdrop) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByNameResponseBrandBackdrop) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetByNameResponseBrandBackdropColor struct {
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
func (r BrandGetByNameResponseBrandBackdropColor) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByNameResponseBrandBackdropColor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resolution of the backdrop image
type BrandGetByNameResponseBrandBackdropResolution struct {
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
func (r BrandGetByNameResponseBrandBackdropResolution) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByNameResponseBrandBackdropResolution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetByNameResponseBrandColor struct {
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
func (r BrandGetByNameResponseBrandColor) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByNameResponseBrandColor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Industry classification information for the brand
type BrandGetByNameResponseBrandIndustries struct {
	// Easy Industry Classification - array of industry and subindustry pairs
	Eic []BrandGetByNameResponseBrandIndustriesEic `json:"eic"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Eic         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandGetByNameResponseBrandIndustries) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByNameResponseBrandIndustries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetByNameResponseBrandIndustriesEic struct {
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
func (r BrandGetByNameResponseBrandIndustriesEic) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByNameResponseBrandIndustriesEic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Important website links for the brand
type BrandGetByNameResponseBrandLinks struct {
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
func (r BrandGetByNameResponseBrandLinks) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByNameResponseBrandLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetByNameResponseBrandLogo struct {
	// Array of colors in the logo
	Colors []BrandGetByNameResponseBrandLogoColor `json:"colors"`
	// Indicates when this logo is best used: 'light' = best for light mode, 'dark' =
	// best for dark mode, 'has_opaque_background' = can be used for either as image
	// has its own background
	//
	// Any of "light", "dark", "has_opaque_background".
	Mode string `json:"mode"`
	// Resolution of the logo image
	Resolution BrandGetByNameResponseBrandLogoResolution `json:"resolution"`
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
func (r BrandGetByNameResponseBrandLogo) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByNameResponseBrandLogo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetByNameResponseBrandLogoColor struct {
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
func (r BrandGetByNameResponseBrandLogoColor) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByNameResponseBrandLogoColor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resolution of the logo image
type BrandGetByNameResponseBrandLogoResolution struct {
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
func (r BrandGetByNameResponseBrandLogoResolution) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByNameResponseBrandLogoResolution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetByNameResponseBrandSocial struct {
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
func (r BrandGetByNameResponseBrandSocial) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByNameResponseBrandSocial) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stock market information for this brand (will be null if not a publicly traded
// company)
type BrandGetByNameResponseBrandStock struct {
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
func (r BrandGetByNameResponseBrandStock) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByNameResponseBrandStock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetByTickerResponse struct {
	// Detailed brand information
	Brand BrandGetByTickerResponseBrand `json:"brand"`
	// HTTP status code
	Code int64 `json:"code"`
	// Status of the response, e.g., 'ok'
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Brand       respjson.Field
		Code        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandGetByTickerResponse) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByTickerResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detailed brand information
type BrandGetByTickerResponseBrand struct {
	// Physical address of the brand
	Address BrandGetByTickerResponseBrandAddress `json:"address"`
	// An array of backdrop images for the brand
	Backdrops []BrandGetByTickerResponseBrandBackdrop `json:"backdrops"`
	// An array of brand colors
	Colors []BrandGetByTickerResponseBrandColor `json:"colors"`
	// A brief description of the brand
	Description string `json:"description"`
	// The domain name of the brand
	Domain string `json:"domain"`
	// Company email address
	Email string `json:"email"`
	// Industry classification information for the brand
	Industries BrandGetByTickerResponseBrandIndustries `json:"industries"`
	// Indicates whether the brand content is not safe for work (NSFW)
	IsNsfw bool `json:"is_nsfw"`
	// Important website links for the brand
	Links BrandGetByTickerResponseBrandLinks `json:"links"`
	// An array of logos associated with the brand
	Logos []BrandGetByTickerResponseBrandLogo `json:"logos"`
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
	Socials []BrandGetByTickerResponseBrandSocial `json:"socials"`
	// Stock market information for this brand (will be null if not a publicly traded
	// company)
	Stock BrandGetByTickerResponseBrandStock `json:"stock"`
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
func (r BrandGetByTickerResponseBrand) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByTickerResponseBrand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Physical address of the brand
type BrandGetByTickerResponseBrandAddress struct {
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
func (r BrandGetByTickerResponseBrandAddress) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByTickerResponseBrandAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetByTickerResponseBrandBackdrop struct {
	// Array of colors in the backdrop image
	Colors []BrandGetByTickerResponseBrandBackdropColor `json:"colors"`
	// Resolution of the backdrop image
	Resolution BrandGetByTickerResponseBrandBackdropResolution `json:"resolution"`
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
func (r BrandGetByTickerResponseBrandBackdrop) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByTickerResponseBrandBackdrop) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetByTickerResponseBrandBackdropColor struct {
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
func (r BrandGetByTickerResponseBrandBackdropColor) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByTickerResponseBrandBackdropColor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resolution of the backdrop image
type BrandGetByTickerResponseBrandBackdropResolution struct {
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
func (r BrandGetByTickerResponseBrandBackdropResolution) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByTickerResponseBrandBackdropResolution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetByTickerResponseBrandColor struct {
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
func (r BrandGetByTickerResponseBrandColor) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByTickerResponseBrandColor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Industry classification information for the brand
type BrandGetByTickerResponseBrandIndustries struct {
	// Easy Industry Classification - array of industry and subindustry pairs
	Eic []BrandGetByTickerResponseBrandIndustriesEic `json:"eic"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Eic         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandGetByTickerResponseBrandIndustries) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByTickerResponseBrandIndustries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetByTickerResponseBrandIndustriesEic struct {
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
func (r BrandGetByTickerResponseBrandIndustriesEic) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByTickerResponseBrandIndustriesEic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Important website links for the brand
type BrandGetByTickerResponseBrandLinks struct {
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
func (r BrandGetByTickerResponseBrandLinks) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByTickerResponseBrandLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetByTickerResponseBrandLogo struct {
	// Array of colors in the logo
	Colors []BrandGetByTickerResponseBrandLogoColor `json:"colors"`
	// Indicates when this logo is best used: 'light' = best for light mode, 'dark' =
	// best for dark mode, 'has_opaque_background' = can be used for either as image
	// has its own background
	//
	// Any of "light", "dark", "has_opaque_background".
	Mode string `json:"mode"`
	// Resolution of the logo image
	Resolution BrandGetByTickerResponseBrandLogoResolution `json:"resolution"`
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
func (r BrandGetByTickerResponseBrandLogo) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByTickerResponseBrandLogo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetByTickerResponseBrandLogoColor struct {
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
func (r BrandGetByTickerResponseBrandLogoColor) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByTickerResponseBrandLogoColor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resolution of the logo image
type BrandGetByTickerResponseBrandLogoResolution struct {
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
func (r BrandGetByTickerResponseBrandLogoResolution) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByTickerResponseBrandLogoResolution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetByTickerResponseBrandSocial struct {
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
func (r BrandGetByTickerResponseBrandSocial) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByTickerResponseBrandSocial) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stock market information for this brand (will be null if not a publicly traded
// company)
type BrandGetByTickerResponseBrandStock struct {
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
func (r BrandGetByTickerResponseBrandStock) RawJSON() string { return r.JSON.raw }
func (r *BrandGetByTickerResponseBrandStock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandGetSimplifiedResponse struct {
	// Simplified brand information
	Brand BrandGetSimplifiedResponseBrand `json:"brand"`
	// HTTP status code of the response
	Code int64 `json:"code"`
	// Status of the response, e.g., 'ok'
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Brand       respjson.Field
		Code        respjson.Field
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

type BrandGetParams struct {
	// Domain name to retrieve brand data for (e.g., 'example.com', 'google.com').
	// Cannot be used with name or ticker parameters.
	Domain string `query:"domain" api:"required" json:"-"`
	// Maximum age in milliseconds for cached brand data before the API performs a hard
	// refresh. Defaults to 3 months (7776000000 ms). Values below 1 day (86400000 ms)
	// are clamped to 1 day; values above 1 year (31536000000 ms) are clamped to 1
	// year.
	MaxAgeMs param.Opt[int64] `query:"maxAgeMs,omitzero" json:"-"`
	// Optional parameter to optimize the API call for maximum speed. When set to true,
	// the API will skip time-consuming operations for faster response at the cost of
	// less comprehensive data. Works with all three lookup methods.
	MaxSpeed param.Opt[bool] `query:"maxSpeed,omitzero" json:"-"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `query:"timeoutMS,omitzero" json:"-"`
	// Optional parameter to force the language of the retrieved brand data.
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
	ForceLanguage BrandGetParamsForceLanguage `query:"force_language,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BrandGetParams]'s query parameters as `url.Values`.
func (r BrandGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Optional parameter to force the language of the retrieved brand data.
type BrandGetParamsForceLanguage string

const (
	BrandGetParamsForceLanguageAfrikaans      BrandGetParamsForceLanguage = "afrikaans"
	BrandGetParamsForceLanguageAlbanian       BrandGetParamsForceLanguage = "albanian"
	BrandGetParamsForceLanguageAmharic        BrandGetParamsForceLanguage = "amharic"
	BrandGetParamsForceLanguageArabic         BrandGetParamsForceLanguage = "arabic"
	BrandGetParamsForceLanguageArmenian       BrandGetParamsForceLanguage = "armenian"
	BrandGetParamsForceLanguageAssamese       BrandGetParamsForceLanguage = "assamese"
	BrandGetParamsForceLanguageAymara         BrandGetParamsForceLanguage = "aymara"
	BrandGetParamsForceLanguageAzeri          BrandGetParamsForceLanguage = "azeri"
	BrandGetParamsForceLanguageBasque         BrandGetParamsForceLanguage = "basque"
	BrandGetParamsForceLanguageBelarusian     BrandGetParamsForceLanguage = "belarusian"
	BrandGetParamsForceLanguageBengali        BrandGetParamsForceLanguage = "bengali"
	BrandGetParamsForceLanguageBosnian        BrandGetParamsForceLanguage = "bosnian"
	BrandGetParamsForceLanguageBulgarian      BrandGetParamsForceLanguage = "bulgarian"
	BrandGetParamsForceLanguageBurmese        BrandGetParamsForceLanguage = "burmese"
	BrandGetParamsForceLanguageCantonese      BrandGetParamsForceLanguage = "cantonese"
	BrandGetParamsForceLanguageCatalan        BrandGetParamsForceLanguage = "catalan"
	BrandGetParamsForceLanguageCebuano        BrandGetParamsForceLanguage = "cebuano"
	BrandGetParamsForceLanguageChinese        BrandGetParamsForceLanguage = "chinese"
	BrandGetParamsForceLanguageCorsican       BrandGetParamsForceLanguage = "corsican"
	BrandGetParamsForceLanguageCroatian       BrandGetParamsForceLanguage = "croatian"
	BrandGetParamsForceLanguageCzech          BrandGetParamsForceLanguage = "czech"
	BrandGetParamsForceLanguageDanish         BrandGetParamsForceLanguage = "danish"
	BrandGetParamsForceLanguageDutch          BrandGetParamsForceLanguage = "dutch"
	BrandGetParamsForceLanguageEnglish        BrandGetParamsForceLanguage = "english"
	BrandGetParamsForceLanguageEsperanto      BrandGetParamsForceLanguage = "esperanto"
	BrandGetParamsForceLanguageEstonian       BrandGetParamsForceLanguage = "estonian"
	BrandGetParamsForceLanguageFarsi          BrandGetParamsForceLanguage = "farsi"
	BrandGetParamsForceLanguageFijian         BrandGetParamsForceLanguage = "fijian"
	BrandGetParamsForceLanguageFinnish        BrandGetParamsForceLanguage = "finnish"
	BrandGetParamsForceLanguageFrench         BrandGetParamsForceLanguage = "french"
	BrandGetParamsForceLanguageGalician       BrandGetParamsForceLanguage = "galician"
	BrandGetParamsForceLanguageGeorgian       BrandGetParamsForceLanguage = "georgian"
	BrandGetParamsForceLanguageGerman         BrandGetParamsForceLanguage = "german"
	BrandGetParamsForceLanguageGreek          BrandGetParamsForceLanguage = "greek"
	BrandGetParamsForceLanguageGuarani        BrandGetParamsForceLanguage = "guarani"
	BrandGetParamsForceLanguageGujarati       BrandGetParamsForceLanguage = "gujarati"
	BrandGetParamsForceLanguageHaitianCreole  BrandGetParamsForceLanguage = "haitian-creole"
	BrandGetParamsForceLanguageHausa          BrandGetParamsForceLanguage = "hausa"
	BrandGetParamsForceLanguageHawaiian       BrandGetParamsForceLanguage = "hawaiian"
	BrandGetParamsForceLanguageHebrew         BrandGetParamsForceLanguage = "hebrew"
	BrandGetParamsForceLanguageHindi          BrandGetParamsForceLanguage = "hindi"
	BrandGetParamsForceLanguageHmong          BrandGetParamsForceLanguage = "hmong"
	BrandGetParamsForceLanguageHungarian      BrandGetParamsForceLanguage = "hungarian"
	BrandGetParamsForceLanguageIcelandic      BrandGetParamsForceLanguage = "icelandic"
	BrandGetParamsForceLanguageIgbo           BrandGetParamsForceLanguage = "igbo"
	BrandGetParamsForceLanguageIndonesian     BrandGetParamsForceLanguage = "indonesian"
	BrandGetParamsForceLanguageIrish          BrandGetParamsForceLanguage = "irish"
	BrandGetParamsForceLanguageItalian        BrandGetParamsForceLanguage = "italian"
	BrandGetParamsForceLanguageJapanese       BrandGetParamsForceLanguage = "japanese"
	BrandGetParamsForceLanguageJavanese       BrandGetParamsForceLanguage = "javanese"
	BrandGetParamsForceLanguageKannada        BrandGetParamsForceLanguage = "kannada"
	BrandGetParamsForceLanguageKazakh         BrandGetParamsForceLanguage = "kazakh"
	BrandGetParamsForceLanguageKhmer          BrandGetParamsForceLanguage = "khmer"
	BrandGetParamsForceLanguageKinyarwanda    BrandGetParamsForceLanguage = "kinyarwanda"
	BrandGetParamsForceLanguageKorean         BrandGetParamsForceLanguage = "korean"
	BrandGetParamsForceLanguageKurdish        BrandGetParamsForceLanguage = "kurdish"
	BrandGetParamsForceLanguageKyrgyz         BrandGetParamsForceLanguage = "kyrgyz"
	BrandGetParamsForceLanguageLao            BrandGetParamsForceLanguage = "lao"
	BrandGetParamsForceLanguageLatin          BrandGetParamsForceLanguage = "latin"
	BrandGetParamsForceLanguageLatvian        BrandGetParamsForceLanguage = "latvian"
	BrandGetParamsForceLanguageLingala        BrandGetParamsForceLanguage = "lingala"
	BrandGetParamsForceLanguageLithuanian     BrandGetParamsForceLanguage = "lithuanian"
	BrandGetParamsForceLanguageLuxembourgish  BrandGetParamsForceLanguage = "luxembourgish"
	BrandGetParamsForceLanguageMacedonian     BrandGetParamsForceLanguage = "macedonian"
	BrandGetParamsForceLanguageMalagasy       BrandGetParamsForceLanguage = "malagasy"
	BrandGetParamsForceLanguageMalay          BrandGetParamsForceLanguage = "malay"
	BrandGetParamsForceLanguageMalayalam      BrandGetParamsForceLanguage = "malayalam"
	BrandGetParamsForceLanguageMaltese        BrandGetParamsForceLanguage = "maltese"
	BrandGetParamsForceLanguageMaori          BrandGetParamsForceLanguage = "maori"
	BrandGetParamsForceLanguageMarathi        BrandGetParamsForceLanguage = "marathi"
	BrandGetParamsForceLanguageMongolian      BrandGetParamsForceLanguage = "mongolian"
	BrandGetParamsForceLanguageNepali         BrandGetParamsForceLanguage = "nepali"
	BrandGetParamsForceLanguageNorwegian      BrandGetParamsForceLanguage = "norwegian"
	BrandGetParamsForceLanguageOdia           BrandGetParamsForceLanguage = "odia"
	BrandGetParamsForceLanguageOromo          BrandGetParamsForceLanguage = "oromo"
	BrandGetParamsForceLanguagePashto         BrandGetParamsForceLanguage = "pashto"
	BrandGetParamsForceLanguagePidgin         BrandGetParamsForceLanguage = "pidgin"
	BrandGetParamsForceLanguagePolish         BrandGetParamsForceLanguage = "polish"
	BrandGetParamsForceLanguagePortuguese     BrandGetParamsForceLanguage = "portuguese"
	BrandGetParamsForceLanguagePunjabi        BrandGetParamsForceLanguage = "punjabi"
	BrandGetParamsForceLanguageQuechua        BrandGetParamsForceLanguage = "quechua"
	BrandGetParamsForceLanguageRomanian       BrandGetParamsForceLanguage = "romanian"
	BrandGetParamsForceLanguageRussian        BrandGetParamsForceLanguage = "russian"
	BrandGetParamsForceLanguageSamoan         BrandGetParamsForceLanguage = "samoan"
	BrandGetParamsForceLanguageScottishGaelic BrandGetParamsForceLanguage = "scottish-gaelic"
	BrandGetParamsForceLanguageSerbian        BrandGetParamsForceLanguage = "serbian"
	BrandGetParamsForceLanguageSesotho        BrandGetParamsForceLanguage = "sesotho"
	BrandGetParamsForceLanguageShona          BrandGetParamsForceLanguage = "shona"
	BrandGetParamsForceLanguageSindhi         BrandGetParamsForceLanguage = "sindhi"
	BrandGetParamsForceLanguageSinhala        BrandGetParamsForceLanguage = "sinhala"
	BrandGetParamsForceLanguageSlovak         BrandGetParamsForceLanguage = "slovak"
	BrandGetParamsForceLanguageSlovene        BrandGetParamsForceLanguage = "slovene"
	BrandGetParamsForceLanguageSomali         BrandGetParamsForceLanguage = "somali"
	BrandGetParamsForceLanguageSpanish        BrandGetParamsForceLanguage = "spanish"
	BrandGetParamsForceLanguageSundanese      BrandGetParamsForceLanguage = "sundanese"
	BrandGetParamsForceLanguageSwahili        BrandGetParamsForceLanguage = "swahili"
	BrandGetParamsForceLanguageSwedish        BrandGetParamsForceLanguage = "swedish"
	BrandGetParamsForceLanguageTagalog        BrandGetParamsForceLanguage = "tagalog"
	BrandGetParamsForceLanguageTajik          BrandGetParamsForceLanguage = "tajik"
	BrandGetParamsForceLanguageTamil          BrandGetParamsForceLanguage = "tamil"
	BrandGetParamsForceLanguageTatar          BrandGetParamsForceLanguage = "tatar"
	BrandGetParamsForceLanguageTelugu         BrandGetParamsForceLanguage = "telugu"
	BrandGetParamsForceLanguageThai           BrandGetParamsForceLanguage = "thai"
	BrandGetParamsForceLanguageTibetan        BrandGetParamsForceLanguage = "tibetan"
	BrandGetParamsForceLanguageTigrinya       BrandGetParamsForceLanguage = "tigrinya"
	BrandGetParamsForceLanguageTongan         BrandGetParamsForceLanguage = "tongan"
	BrandGetParamsForceLanguageTswana         BrandGetParamsForceLanguage = "tswana"
	BrandGetParamsForceLanguageTurkish        BrandGetParamsForceLanguage = "turkish"
	BrandGetParamsForceLanguageTurkmen        BrandGetParamsForceLanguage = "turkmen"
	BrandGetParamsForceLanguageUkrainian      BrandGetParamsForceLanguage = "ukrainian"
	BrandGetParamsForceLanguageUrdu           BrandGetParamsForceLanguage = "urdu"
	BrandGetParamsForceLanguageUyghur         BrandGetParamsForceLanguage = "uyghur"
	BrandGetParamsForceLanguageUzbek          BrandGetParamsForceLanguage = "uzbek"
	BrandGetParamsForceLanguageVietnamese     BrandGetParamsForceLanguage = "vietnamese"
	BrandGetParamsForceLanguageWelsh          BrandGetParamsForceLanguage = "welsh"
	BrandGetParamsForceLanguageWolof          BrandGetParamsForceLanguage = "wolof"
	BrandGetParamsForceLanguageXhosa          BrandGetParamsForceLanguage = "xhosa"
	BrandGetParamsForceLanguageYiddish        BrandGetParamsForceLanguage = "yiddish"
	BrandGetParamsForceLanguageYoruba         BrandGetParamsForceLanguage = "yoruba"
	BrandGetParamsForceLanguageZulu           BrandGetParamsForceLanguage = "zulu"
)

type BrandIdentifyFromTransactionParams struct {
	// Transaction information to identify the brand
	TransactionInfo string `query:"transaction_info" api:"required" json:"-"`
	// Optional city name to prioritize when searching for the brand.
	City param.Opt[string] `query:"city,omitzero" json:"-"`
	// When set to true, the API will perform an additional verification steps to
	// ensure the identified brand matches the transaction with high confidence.
	HighConfidenceOnly param.Opt[bool] `query:"high_confidence_only,omitzero" json:"-"`
	// Optional parameter to optimize the API call for maximum speed. When set to true,
	// the API will skip time-consuming operations for faster response at the cost of
	// less comprehensive data.
	MaxSpeed param.Opt[bool] `query:"maxSpeed,omitzero" json:"-"`
	// Optional Merchant Category Code (MCC) to help identify the business
	// category/industry.
	Mcc param.Opt[string] `query:"mcc,omitzero" json:"-"`
	// Optional phone number from the transaction to help verify brand match.
	Phone param.Opt[float64] `query:"phone,omitzero" json:"-"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `query:"timeoutMS,omitzero" json:"-"`
	// Optional country code (GL parameter) to specify the country. This affects the
	// geographic location used for search queries.
	//
	// Any of "ad", "ae", "af", "ag", "ai", "al", "am", "an", "ao", "aq", "ar", "as",
	// "at", "au", "aw", "az", "ba", "bb", "bd", "be", "bf", "bg", "bh", "bi", "bj",
	// "bm", "bn", "bo", "br", "bs", "bt", "bv", "bw", "by", "bz", "ca", "cc", "cd",
	// "cf", "cg", "ch", "ci", "ck", "cl", "cm", "cn", "co", "cr", "cu", "cv", "cx",
	// "cy", "cz", "de", "dj", "dk", "dm", "do", "dz", "ec", "ee", "eg", "eh", "er",
	// "es", "et", "fi", "fj", "fk", "fm", "fo", "fr", "ga", "gb", "gd", "ge", "gf",
	// "gh", "gi", "gl", "gm", "gn", "gp", "gq", "gr", "gs", "gt", "gu", "gw", "gy",
	// "hk", "hm", "hn", "hr", "ht", "hu", "id", "ie", "il", "in", "io", "iq", "ir",
	// "is", "it", "jm", "jo", "jp", "ke", "kg", "kh", "ki", "km", "kn", "kp", "kr",
	// "kw", "ky", "kz", "la", "lb", "lc", "li", "lk", "lr", "ls", "lt", "lu", "lv",
	// "ly", "ma", "mc", "md", "mg", "mh", "mk", "ml", "mm", "mn", "mo", "mp", "mq",
	// "mr", "ms", "mt", "mu", "mv", "mw", "mx", "my", "mz", "na", "nc", "ne", "nf",
	// "ng", "ni", "nl", "no", "np", "nr", "nu", "nz", "om", "pa", "pe", "pf", "pg",
	// "ph", "pk", "pl", "pm", "pn", "pr", "ps", "pt", "pw", "py", "qa", "re", "ro",
	// "rs", "ru", "rw", "sa", "sb", "sc", "sd", "se", "sg", "sh", "si", "sj", "sk",
	// "sl", "sm", "sn", "so", "sr", "st", "sv", "sy", "sz", "tc", "td", "tf", "tg",
	// "th", "tj", "tk", "tl", "tm", "tn", "to", "tr", "tt", "tv", "tw", "tz", "ua",
	// "ug", "um", "us", "uy", "uz", "va", "vc", "ve", "vg", "vi", "vn", "vu", "wf",
	// "ws", "ye", "yt", "za", "zm", "zw".
	CountryGl BrandIdentifyFromTransactionParamsCountryGl `query:"country_gl,omitzero" json:"-"`
	// Optional parameter to force the language of the retrieved brand data.
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
	ForceLanguage BrandIdentifyFromTransactionParamsForceLanguage `query:"force_language,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BrandIdentifyFromTransactionParams]'s query parameters as
// `url.Values`.
func (r BrandIdentifyFromTransactionParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Optional country code (GL parameter) to specify the country. This affects the
// geographic location used for search queries.
type BrandIdentifyFromTransactionParamsCountryGl string

const (
	BrandIdentifyFromTransactionParamsCountryGlAd BrandIdentifyFromTransactionParamsCountryGl = "ad"
	BrandIdentifyFromTransactionParamsCountryGlAe BrandIdentifyFromTransactionParamsCountryGl = "ae"
	BrandIdentifyFromTransactionParamsCountryGlAf BrandIdentifyFromTransactionParamsCountryGl = "af"
	BrandIdentifyFromTransactionParamsCountryGlAg BrandIdentifyFromTransactionParamsCountryGl = "ag"
	BrandIdentifyFromTransactionParamsCountryGlAI BrandIdentifyFromTransactionParamsCountryGl = "ai"
	BrandIdentifyFromTransactionParamsCountryGlAl BrandIdentifyFromTransactionParamsCountryGl = "al"
	BrandIdentifyFromTransactionParamsCountryGlAm BrandIdentifyFromTransactionParamsCountryGl = "am"
	BrandIdentifyFromTransactionParamsCountryGlAn BrandIdentifyFromTransactionParamsCountryGl = "an"
	BrandIdentifyFromTransactionParamsCountryGlAo BrandIdentifyFromTransactionParamsCountryGl = "ao"
	BrandIdentifyFromTransactionParamsCountryGlAq BrandIdentifyFromTransactionParamsCountryGl = "aq"
	BrandIdentifyFromTransactionParamsCountryGlAr BrandIdentifyFromTransactionParamsCountryGl = "ar"
	BrandIdentifyFromTransactionParamsCountryGlAs BrandIdentifyFromTransactionParamsCountryGl = "as"
	BrandIdentifyFromTransactionParamsCountryGlAt BrandIdentifyFromTransactionParamsCountryGl = "at"
	BrandIdentifyFromTransactionParamsCountryGlAu BrandIdentifyFromTransactionParamsCountryGl = "au"
	BrandIdentifyFromTransactionParamsCountryGlAw BrandIdentifyFromTransactionParamsCountryGl = "aw"
	BrandIdentifyFromTransactionParamsCountryGlAz BrandIdentifyFromTransactionParamsCountryGl = "az"
	BrandIdentifyFromTransactionParamsCountryGlBa BrandIdentifyFromTransactionParamsCountryGl = "ba"
	BrandIdentifyFromTransactionParamsCountryGlBb BrandIdentifyFromTransactionParamsCountryGl = "bb"
	BrandIdentifyFromTransactionParamsCountryGlBd BrandIdentifyFromTransactionParamsCountryGl = "bd"
	BrandIdentifyFromTransactionParamsCountryGlBe BrandIdentifyFromTransactionParamsCountryGl = "be"
	BrandIdentifyFromTransactionParamsCountryGlBf BrandIdentifyFromTransactionParamsCountryGl = "bf"
	BrandIdentifyFromTransactionParamsCountryGlBg BrandIdentifyFromTransactionParamsCountryGl = "bg"
	BrandIdentifyFromTransactionParamsCountryGlBh BrandIdentifyFromTransactionParamsCountryGl = "bh"
	BrandIdentifyFromTransactionParamsCountryGlBi BrandIdentifyFromTransactionParamsCountryGl = "bi"
	BrandIdentifyFromTransactionParamsCountryGlBj BrandIdentifyFromTransactionParamsCountryGl = "bj"
	BrandIdentifyFromTransactionParamsCountryGlBm BrandIdentifyFromTransactionParamsCountryGl = "bm"
	BrandIdentifyFromTransactionParamsCountryGlBn BrandIdentifyFromTransactionParamsCountryGl = "bn"
	BrandIdentifyFromTransactionParamsCountryGlBo BrandIdentifyFromTransactionParamsCountryGl = "bo"
	BrandIdentifyFromTransactionParamsCountryGlBr BrandIdentifyFromTransactionParamsCountryGl = "br"
	BrandIdentifyFromTransactionParamsCountryGlBs BrandIdentifyFromTransactionParamsCountryGl = "bs"
	BrandIdentifyFromTransactionParamsCountryGlBt BrandIdentifyFromTransactionParamsCountryGl = "bt"
	BrandIdentifyFromTransactionParamsCountryGlBv BrandIdentifyFromTransactionParamsCountryGl = "bv"
	BrandIdentifyFromTransactionParamsCountryGlBw BrandIdentifyFromTransactionParamsCountryGl = "bw"
	BrandIdentifyFromTransactionParamsCountryGlBy BrandIdentifyFromTransactionParamsCountryGl = "by"
	BrandIdentifyFromTransactionParamsCountryGlBz BrandIdentifyFromTransactionParamsCountryGl = "bz"
	BrandIdentifyFromTransactionParamsCountryGlCa BrandIdentifyFromTransactionParamsCountryGl = "ca"
	BrandIdentifyFromTransactionParamsCountryGlCc BrandIdentifyFromTransactionParamsCountryGl = "cc"
	BrandIdentifyFromTransactionParamsCountryGlCd BrandIdentifyFromTransactionParamsCountryGl = "cd"
	BrandIdentifyFromTransactionParamsCountryGlCf BrandIdentifyFromTransactionParamsCountryGl = "cf"
	BrandIdentifyFromTransactionParamsCountryGlCg BrandIdentifyFromTransactionParamsCountryGl = "cg"
	BrandIdentifyFromTransactionParamsCountryGlCh BrandIdentifyFromTransactionParamsCountryGl = "ch"
	BrandIdentifyFromTransactionParamsCountryGlCi BrandIdentifyFromTransactionParamsCountryGl = "ci"
	BrandIdentifyFromTransactionParamsCountryGlCk BrandIdentifyFromTransactionParamsCountryGl = "ck"
	BrandIdentifyFromTransactionParamsCountryGlCl BrandIdentifyFromTransactionParamsCountryGl = "cl"
	BrandIdentifyFromTransactionParamsCountryGlCm BrandIdentifyFromTransactionParamsCountryGl = "cm"
	BrandIdentifyFromTransactionParamsCountryGlCn BrandIdentifyFromTransactionParamsCountryGl = "cn"
	BrandIdentifyFromTransactionParamsCountryGlCo BrandIdentifyFromTransactionParamsCountryGl = "co"
	BrandIdentifyFromTransactionParamsCountryGlCr BrandIdentifyFromTransactionParamsCountryGl = "cr"
	BrandIdentifyFromTransactionParamsCountryGlCu BrandIdentifyFromTransactionParamsCountryGl = "cu"
	BrandIdentifyFromTransactionParamsCountryGlCv BrandIdentifyFromTransactionParamsCountryGl = "cv"
	BrandIdentifyFromTransactionParamsCountryGlCx BrandIdentifyFromTransactionParamsCountryGl = "cx"
	BrandIdentifyFromTransactionParamsCountryGlCy BrandIdentifyFromTransactionParamsCountryGl = "cy"
	BrandIdentifyFromTransactionParamsCountryGlCz BrandIdentifyFromTransactionParamsCountryGl = "cz"
	BrandIdentifyFromTransactionParamsCountryGlDe BrandIdentifyFromTransactionParamsCountryGl = "de"
	BrandIdentifyFromTransactionParamsCountryGlDj BrandIdentifyFromTransactionParamsCountryGl = "dj"
	BrandIdentifyFromTransactionParamsCountryGlDk BrandIdentifyFromTransactionParamsCountryGl = "dk"
	BrandIdentifyFromTransactionParamsCountryGlDm BrandIdentifyFromTransactionParamsCountryGl = "dm"
	BrandIdentifyFromTransactionParamsCountryGlDo BrandIdentifyFromTransactionParamsCountryGl = "do"
	BrandIdentifyFromTransactionParamsCountryGlDz BrandIdentifyFromTransactionParamsCountryGl = "dz"
	BrandIdentifyFromTransactionParamsCountryGlEc BrandIdentifyFromTransactionParamsCountryGl = "ec"
	BrandIdentifyFromTransactionParamsCountryGlEe BrandIdentifyFromTransactionParamsCountryGl = "ee"
	BrandIdentifyFromTransactionParamsCountryGlEg BrandIdentifyFromTransactionParamsCountryGl = "eg"
	BrandIdentifyFromTransactionParamsCountryGlEh BrandIdentifyFromTransactionParamsCountryGl = "eh"
	BrandIdentifyFromTransactionParamsCountryGlEr BrandIdentifyFromTransactionParamsCountryGl = "er"
	BrandIdentifyFromTransactionParamsCountryGlEs BrandIdentifyFromTransactionParamsCountryGl = "es"
	BrandIdentifyFromTransactionParamsCountryGlEt BrandIdentifyFromTransactionParamsCountryGl = "et"
	BrandIdentifyFromTransactionParamsCountryGlFi BrandIdentifyFromTransactionParamsCountryGl = "fi"
	BrandIdentifyFromTransactionParamsCountryGlFj BrandIdentifyFromTransactionParamsCountryGl = "fj"
	BrandIdentifyFromTransactionParamsCountryGlFk BrandIdentifyFromTransactionParamsCountryGl = "fk"
	BrandIdentifyFromTransactionParamsCountryGlFm BrandIdentifyFromTransactionParamsCountryGl = "fm"
	BrandIdentifyFromTransactionParamsCountryGlFo BrandIdentifyFromTransactionParamsCountryGl = "fo"
	BrandIdentifyFromTransactionParamsCountryGlFr BrandIdentifyFromTransactionParamsCountryGl = "fr"
	BrandIdentifyFromTransactionParamsCountryGlGa BrandIdentifyFromTransactionParamsCountryGl = "ga"
	BrandIdentifyFromTransactionParamsCountryGlGB BrandIdentifyFromTransactionParamsCountryGl = "gb"
	BrandIdentifyFromTransactionParamsCountryGlGd BrandIdentifyFromTransactionParamsCountryGl = "gd"
	BrandIdentifyFromTransactionParamsCountryGlGe BrandIdentifyFromTransactionParamsCountryGl = "ge"
	BrandIdentifyFromTransactionParamsCountryGlGf BrandIdentifyFromTransactionParamsCountryGl = "gf"
	BrandIdentifyFromTransactionParamsCountryGlGh BrandIdentifyFromTransactionParamsCountryGl = "gh"
	BrandIdentifyFromTransactionParamsCountryGlGi BrandIdentifyFromTransactionParamsCountryGl = "gi"
	BrandIdentifyFromTransactionParamsCountryGlGl BrandIdentifyFromTransactionParamsCountryGl = "gl"
	BrandIdentifyFromTransactionParamsCountryGlGm BrandIdentifyFromTransactionParamsCountryGl = "gm"
	BrandIdentifyFromTransactionParamsCountryGlGn BrandIdentifyFromTransactionParamsCountryGl = "gn"
	BrandIdentifyFromTransactionParamsCountryGlGp BrandIdentifyFromTransactionParamsCountryGl = "gp"
	BrandIdentifyFromTransactionParamsCountryGlGq BrandIdentifyFromTransactionParamsCountryGl = "gq"
	BrandIdentifyFromTransactionParamsCountryGlGr BrandIdentifyFromTransactionParamsCountryGl = "gr"
	BrandIdentifyFromTransactionParamsCountryGlGs BrandIdentifyFromTransactionParamsCountryGl = "gs"
	BrandIdentifyFromTransactionParamsCountryGlGt BrandIdentifyFromTransactionParamsCountryGl = "gt"
	BrandIdentifyFromTransactionParamsCountryGlGu BrandIdentifyFromTransactionParamsCountryGl = "gu"
	BrandIdentifyFromTransactionParamsCountryGlGw BrandIdentifyFromTransactionParamsCountryGl = "gw"
	BrandIdentifyFromTransactionParamsCountryGlGy BrandIdentifyFromTransactionParamsCountryGl = "gy"
	BrandIdentifyFromTransactionParamsCountryGlHk BrandIdentifyFromTransactionParamsCountryGl = "hk"
	BrandIdentifyFromTransactionParamsCountryGlHm BrandIdentifyFromTransactionParamsCountryGl = "hm"
	BrandIdentifyFromTransactionParamsCountryGlHn BrandIdentifyFromTransactionParamsCountryGl = "hn"
	BrandIdentifyFromTransactionParamsCountryGlHr BrandIdentifyFromTransactionParamsCountryGl = "hr"
	BrandIdentifyFromTransactionParamsCountryGlHt BrandIdentifyFromTransactionParamsCountryGl = "ht"
	BrandIdentifyFromTransactionParamsCountryGlHu BrandIdentifyFromTransactionParamsCountryGl = "hu"
	BrandIdentifyFromTransactionParamsCountryGlID BrandIdentifyFromTransactionParamsCountryGl = "id"
	BrandIdentifyFromTransactionParamsCountryGlIe BrandIdentifyFromTransactionParamsCountryGl = "ie"
	BrandIdentifyFromTransactionParamsCountryGlIl BrandIdentifyFromTransactionParamsCountryGl = "il"
	BrandIdentifyFromTransactionParamsCountryGlIn BrandIdentifyFromTransactionParamsCountryGl = "in"
	BrandIdentifyFromTransactionParamsCountryGlIo BrandIdentifyFromTransactionParamsCountryGl = "io"
	BrandIdentifyFromTransactionParamsCountryGlIq BrandIdentifyFromTransactionParamsCountryGl = "iq"
	BrandIdentifyFromTransactionParamsCountryGlIr BrandIdentifyFromTransactionParamsCountryGl = "ir"
	BrandIdentifyFromTransactionParamsCountryGlIs BrandIdentifyFromTransactionParamsCountryGl = "is"
	BrandIdentifyFromTransactionParamsCountryGlIt BrandIdentifyFromTransactionParamsCountryGl = "it"
	BrandIdentifyFromTransactionParamsCountryGlJm BrandIdentifyFromTransactionParamsCountryGl = "jm"
	BrandIdentifyFromTransactionParamsCountryGlJo BrandIdentifyFromTransactionParamsCountryGl = "jo"
	BrandIdentifyFromTransactionParamsCountryGlJp BrandIdentifyFromTransactionParamsCountryGl = "jp"
	BrandIdentifyFromTransactionParamsCountryGlKe BrandIdentifyFromTransactionParamsCountryGl = "ke"
	BrandIdentifyFromTransactionParamsCountryGlKg BrandIdentifyFromTransactionParamsCountryGl = "kg"
	BrandIdentifyFromTransactionParamsCountryGlKh BrandIdentifyFromTransactionParamsCountryGl = "kh"
	BrandIdentifyFromTransactionParamsCountryGlKi BrandIdentifyFromTransactionParamsCountryGl = "ki"
	BrandIdentifyFromTransactionParamsCountryGlKm BrandIdentifyFromTransactionParamsCountryGl = "km"
	BrandIdentifyFromTransactionParamsCountryGlKn BrandIdentifyFromTransactionParamsCountryGl = "kn"
	BrandIdentifyFromTransactionParamsCountryGlKp BrandIdentifyFromTransactionParamsCountryGl = "kp"
	BrandIdentifyFromTransactionParamsCountryGlKr BrandIdentifyFromTransactionParamsCountryGl = "kr"
	BrandIdentifyFromTransactionParamsCountryGlKw BrandIdentifyFromTransactionParamsCountryGl = "kw"
	BrandIdentifyFromTransactionParamsCountryGlKy BrandIdentifyFromTransactionParamsCountryGl = "ky"
	BrandIdentifyFromTransactionParamsCountryGlKz BrandIdentifyFromTransactionParamsCountryGl = "kz"
	BrandIdentifyFromTransactionParamsCountryGlLa BrandIdentifyFromTransactionParamsCountryGl = "la"
	BrandIdentifyFromTransactionParamsCountryGlLb BrandIdentifyFromTransactionParamsCountryGl = "lb"
	BrandIdentifyFromTransactionParamsCountryGlLc BrandIdentifyFromTransactionParamsCountryGl = "lc"
	BrandIdentifyFromTransactionParamsCountryGlLi BrandIdentifyFromTransactionParamsCountryGl = "li"
	BrandIdentifyFromTransactionParamsCountryGlLk BrandIdentifyFromTransactionParamsCountryGl = "lk"
	BrandIdentifyFromTransactionParamsCountryGlLr BrandIdentifyFromTransactionParamsCountryGl = "lr"
	BrandIdentifyFromTransactionParamsCountryGlLs BrandIdentifyFromTransactionParamsCountryGl = "ls"
	BrandIdentifyFromTransactionParamsCountryGlLt BrandIdentifyFromTransactionParamsCountryGl = "lt"
	BrandIdentifyFromTransactionParamsCountryGlLu BrandIdentifyFromTransactionParamsCountryGl = "lu"
	BrandIdentifyFromTransactionParamsCountryGlLv BrandIdentifyFromTransactionParamsCountryGl = "lv"
	BrandIdentifyFromTransactionParamsCountryGlLy BrandIdentifyFromTransactionParamsCountryGl = "ly"
	BrandIdentifyFromTransactionParamsCountryGlMa BrandIdentifyFromTransactionParamsCountryGl = "ma"
	BrandIdentifyFromTransactionParamsCountryGlMc BrandIdentifyFromTransactionParamsCountryGl = "mc"
	BrandIdentifyFromTransactionParamsCountryGlMd BrandIdentifyFromTransactionParamsCountryGl = "md"
	BrandIdentifyFromTransactionParamsCountryGlMg BrandIdentifyFromTransactionParamsCountryGl = "mg"
	BrandIdentifyFromTransactionParamsCountryGlMh BrandIdentifyFromTransactionParamsCountryGl = "mh"
	BrandIdentifyFromTransactionParamsCountryGlMk BrandIdentifyFromTransactionParamsCountryGl = "mk"
	BrandIdentifyFromTransactionParamsCountryGlMl BrandIdentifyFromTransactionParamsCountryGl = "ml"
	BrandIdentifyFromTransactionParamsCountryGlMm BrandIdentifyFromTransactionParamsCountryGl = "mm"
	BrandIdentifyFromTransactionParamsCountryGlMn BrandIdentifyFromTransactionParamsCountryGl = "mn"
	BrandIdentifyFromTransactionParamsCountryGlMo BrandIdentifyFromTransactionParamsCountryGl = "mo"
	BrandIdentifyFromTransactionParamsCountryGlMp BrandIdentifyFromTransactionParamsCountryGl = "mp"
	BrandIdentifyFromTransactionParamsCountryGlMq BrandIdentifyFromTransactionParamsCountryGl = "mq"
	BrandIdentifyFromTransactionParamsCountryGlMr BrandIdentifyFromTransactionParamsCountryGl = "mr"
	BrandIdentifyFromTransactionParamsCountryGlMs BrandIdentifyFromTransactionParamsCountryGl = "ms"
	BrandIdentifyFromTransactionParamsCountryGlMt BrandIdentifyFromTransactionParamsCountryGl = "mt"
	BrandIdentifyFromTransactionParamsCountryGlMu BrandIdentifyFromTransactionParamsCountryGl = "mu"
	BrandIdentifyFromTransactionParamsCountryGlMv BrandIdentifyFromTransactionParamsCountryGl = "mv"
	BrandIdentifyFromTransactionParamsCountryGlMw BrandIdentifyFromTransactionParamsCountryGl = "mw"
	BrandIdentifyFromTransactionParamsCountryGlMx BrandIdentifyFromTransactionParamsCountryGl = "mx"
	BrandIdentifyFromTransactionParamsCountryGlMy BrandIdentifyFromTransactionParamsCountryGl = "my"
	BrandIdentifyFromTransactionParamsCountryGlMz BrandIdentifyFromTransactionParamsCountryGl = "mz"
	BrandIdentifyFromTransactionParamsCountryGlNa BrandIdentifyFromTransactionParamsCountryGl = "na"
	BrandIdentifyFromTransactionParamsCountryGlNc BrandIdentifyFromTransactionParamsCountryGl = "nc"
	BrandIdentifyFromTransactionParamsCountryGlNe BrandIdentifyFromTransactionParamsCountryGl = "ne"
	BrandIdentifyFromTransactionParamsCountryGlNf BrandIdentifyFromTransactionParamsCountryGl = "nf"
	BrandIdentifyFromTransactionParamsCountryGlNg BrandIdentifyFromTransactionParamsCountryGl = "ng"
	BrandIdentifyFromTransactionParamsCountryGlNi BrandIdentifyFromTransactionParamsCountryGl = "ni"
	BrandIdentifyFromTransactionParamsCountryGlNl BrandIdentifyFromTransactionParamsCountryGl = "nl"
	BrandIdentifyFromTransactionParamsCountryGlNo BrandIdentifyFromTransactionParamsCountryGl = "no"
	BrandIdentifyFromTransactionParamsCountryGlNp BrandIdentifyFromTransactionParamsCountryGl = "np"
	BrandIdentifyFromTransactionParamsCountryGlNr BrandIdentifyFromTransactionParamsCountryGl = "nr"
	BrandIdentifyFromTransactionParamsCountryGlNu BrandIdentifyFromTransactionParamsCountryGl = "nu"
	BrandIdentifyFromTransactionParamsCountryGlNz BrandIdentifyFromTransactionParamsCountryGl = "nz"
	BrandIdentifyFromTransactionParamsCountryGlOm BrandIdentifyFromTransactionParamsCountryGl = "om"
	BrandIdentifyFromTransactionParamsCountryGlPa BrandIdentifyFromTransactionParamsCountryGl = "pa"
	BrandIdentifyFromTransactionParamsCountryGlPe BrandIdentifyFromTransactionParamsCountryGl = "pe"
	BrandIdentifyFromTransactionParamsCountryGlPf BrandIdentifyFromTransactionParamsCountryGl = "pf"
	BrandIdentifyFromTransactionParamsCountryGlPg BrandIdentifyFromTransactionParamsCountryGl = "pg"
	BrandIdentifyFromTransactionParamsCountryGlPh BrandIdentifyFromTransactionParamsCountryGl = "ph"
	BrandIdentifyFromTransactionParamsCountryGlPk BrandIdentifyFromTransactionParamsCountryGl = "pk"
	BrandIdentifyFromTransactionParamsCountryGlPl BrandIdentifyFromTransactionParamsCountryGl = "pl"
	BrandIdentifyFromTransactionParamsCountryGlPm BrandIdentifyFromTransactionParamsCountryGl = "pm"
	BrandIdentifyFromTransactionParamsCountryGlPn BrandIdentifyFromTransactionParamsCountryGl = "pn"
	BrandIdentifyFromTransactionParamsCountryGlPr BrandIdentifyFromTransactionParamsCountryGl = "pr"
	BrandIdentifyFromTransactionParamsCountryGlPs BrandIdentifyFromTransactionParamsCountryGl = "ps"
	BrandIdentifyFromTransactionParamsCountryGlPt BrandIdentifyFromTransactionParamsCountryGl = "pt"
	BrandIdentifyFromTransactionParamsCountryGlPw BrandIdentifyFromTransactionParamsCountryGl = "pw"
	BrandIdentifyFromTransactionParamsCountryGlPy BrandIdentifyFromTransactionParamsCountryGl = "py"
	BrandIdentifyFromTransactionParamsCountryGlQa BrandIdentifyFromTransactionParamsCountryGl = "qa"
	BrandIdentifyFromTransactionParamsCountryGlRe BrandIdentifyFromTransactionParamsCountryGl = "re"
	BrandIdentifyFromTransactionParamsCountryGlRo BrandIdentifyFromTransactionParamsCountryGl = "ro"
	BrandIdentifyFromTransactionParamsCountryGlRs BrandIdentifyFromTransactionParamsCountryGl = "rs"
	BrandIdentifyFromTransactionParamsCountryGlRu BrandIdentifyFromTransactionParamsCountryGl = "ru"
	BrandIdentifyFromTransactionParamsCountryGlRw BrandIdentifyFromTransactionParamsCountryGl = "rw"
	BrandIdentifyFromTransactionParamsCountryGlSa BrandIdentifyFromTransactionParamsCountryGl = "sa"
	BrandIdentifyFromTransactionParamsCountryGlSb BrandIdentifyFromTransactionParamsCountryGl = "sb"
	BrandIdentifyFromTransactionParamsCountryGlSc BrandIdentifyFromTransactionParamsCountryGl = "sc"
	BrandIdentifyFromTransactionParamsCountryGlSd BrandIdentifyFromTransactionParamsCountryGl = "sd"
	BrandIdentifyFromTransactionParamsCountryGlSe BrandIdentifyFromTransactionParamsCountryGl = "se"
	BrandIdentifyFromTransactionParamsCountryGlSg BrandIdentifyFromTransactionParamsCountryGl = "sg"
	BrandIdentifyFromTransactionParamsCountryGlSh BrandIdentifyFromTransactionParamsCountryGl = "sh"
	BrandIdentifyFromTransactionParamsCountryGlSi BrandIdentifyFromTransactionParamsCountryGl = "si"
	BrandIdentifyFromTransactionParamsCountryGlSj BrandIdentifyFromTransactionParamsCountryGl = "sj"
	BrandIdentifyFromTransactionParamsCountryGlSk BrandIdentifyFromTransactionParamsCountryGl = "sk"
	BrandIdentifyFromTransactionParamsCountryGlSl BrandIdentifyFromTransactionParamsCountryGl = "sl"
	BrandIdentifyFromTransactionParamsCountryGlSm BrandIdentifyFromTransactionParamsCountryGl = "sm"
	BrandIdentifyFromTransactionParamsCountryGlSn BrandIdentifyFromTransactionParamsCountryGl = "sn"
	BrandIdentifyFromTransactionParamsCountryGlSo BrandIdentifyFromTransactionParamsCountryGl = "so"
	BrandIdentifyFromTransactionParamsCountryGlSr BrandIdentifyFromTransactionParamsCountryGl = "sr"
	BrandIdentifyFromTransactionParamsCountryGlSt BrandIdentifyFromTransactionParamsCountryGl = "st"
	BrandIdentifyFromTransactionParamsCountryGlSv BrandIdentifyFromTransactionParamsCountryGl = "sv"
	BrandIdentifyFromTransactionParamsCountryGlSy BrandIdentifyFromTransactionParamsCountryGl = "sy"
	BrandIdentifyFromTransactionParamsCountryGlSz BrandIdentifyFromTransactionParamsCountryGl = "sz"
	BrandIdentifyFromTransactionParamsCountryGlTc BrandIdentifyFromTransactionParamsCountryGl = "tc"
	BrandIdentifyFromTransactionParamsCountryGlTd BrandIdentifyFromTransactionParamsCountryGl = "td"
	BrandIdentifyFromTransactionParamsCountryGlTf BrandIdentifyFromTransactionParamsCountryGl = "tf"
	BrandIdentifyFromTransactionParamsCountryGlTg BrandIdentifyFromTransactionParamsCountryGl = "tg"
	BrandIdentifyFromTransactionParamsCountryGlTh BrandIdentifyFromTransactionParamsCountryGl = "th"
	BrandIdentifyFromTransactionParamsCountryGlTj BrandIdentifyFromTransactionParamsCountryGl = "tj"
	BrandIdentifyFromTransactionParamsCountryGlTk BrandIdentifyFromTransactionParamsCountryGl = "tk"
	BrandIdentifyFromTransactionParamsCountryGlTl BrandIdentifyFromTransactionParamsCountryGl = "tl"
	BrandIdentifyFromTransactionParamsCountryGlTm BrandIdentifyFromTransactionParamsCountryGl = "tm"
	BrandIdentifyFromTransactionParamsCountryGlTn BrandIdentifyFromTransactionParamsCountryGl = "tn"
	BrandIdentifyFromTransactionParamsCountryGlTo BrandIdentifyFromTransactionParamsCountryGl = "to"
	BrandIdentifyFromTransactionParamsCountryGlTr BrandIdentifyFromTransactionParamsCountryGl = "tr"
	BrandIdentifyFromTransactionParamsCountryGlTt BrandIdentifyFromTransactionParamsCountryGl = "tt"
	BrandIdentifyFromTransactionParamsCountryGlTv BrandIdentifyFromTransactionParamsCountryGl = "tv"
	BrandIdentifyFromTransactionParamsCountryGlTw BrandIdentifyFromTransactionParamsCountryGl = "tw"
	BrandIdentifyFromTransactionParamsCountryGlTz BrandIdentifyFromTransactionParamsCountryGl = "tz"
	BrandIdentifyFromTransactionParamsCountryGlUa BrandIdentifyFromTransactionParamsCountryGl = "ua"
	BrandIdentifyFromTransactionParamsCountryGlUg BrandIdentifyFromTransactionParamsCountryGl = "ug"
	BrandIdentifyFromTransactionParamsCountryGlUm BrandIdentifyFromTransactionParamsCountryGl = "um"
	BrandIdentifyFromTransactionParamsCountryGlUs BrandIdentifyFromTransactionParamsCountryGl = "us"
	BrandIdentifyFromTransactionParamsCountryGlUy BrandIdentifyFromTransactionParamsCountryGl = "uy"
	BrandIdentifyFromTransactionParamsCountryGlUz BrandIdentifyFromTransactionParamsCountryGl = "uz"
	BrandIdentifyFromTransactionParamsCountryGlVa BrandIdentifyFromTransactionParamsCountryGl = "va"
	BrandIdentifyFromTransactionParamsCountryGlVc BrandIdentifyFromTransactionParamsCountryGl = "vc"
	BrandIdentifyFromTransactionParamsCountryGlVe BrandIdentifyFromTransactionParamsCountryGl = "ve"
	BrandIdentifyFromTransactionParamsCountryGlVg BrandIdentifyFromTransactionParamsCountryGl = "vg"
	BrandIdentifyFromTransactionParamsCountryGlVi BrandIdentifyFromTransactionParamsCountryGl = "vi"
	BrandIdentifyFromTransactionParamsCountryGlVn BrandIdentifyFromTransactionParamsCountryGl = "vn"
	BrandIdentifyFromTransactionParamsCountryGlVu BrandIdentifyFromTransactionParamsCountryGl = "vu"
	BrandIdentifyFromTransactionParamsCountryGlWf BrandIdentifyFromTransactionParamsCountryGl = "wf"
	BrandIdentifyFromTransactionParamsCountryGlWs BrandIdentifyFromTransactionParamsCountryGl = "ws"
	BrandIdentifyFromTransactionParamsCountryGlYe BrandIdentifyFromTransactionParamsCountryGl = "ye"
	BrandIdentifyFromTransactionParamsCountryGlYt BrandIdentifyFromTransactionParamsCountryGl = "yt"
	BrandIdentifyFromTransactionParamsCountryGlZa BrandIdentifyFromTransactionParamsCountryGl = "za"
	BrandIdentifyFromTransactionParamsCountryGlZm BrandIdentifyFromTransactionParamsCountryGl = "zm"
	BrandIdentifyFromTransactionParamsCountryGlZw BrandIdentifyFromTransactionParamsCountryGl = "zw"
)

// Optional parameter to force the language of the retrieved brand data.
type BrandIdentifyFromTransactionParamsForceLanguage string

const (
	BrandIdentifyFromTransactionParamsForceLanguageAfrikaans      BrandIdentifyFromTransactionParamsForceLanguage = "afrikaans"
	BrandIdentifyFromTransactionParamsForceLanguageAlbanian       BrandIdentifyFromTransactionParamsForceLanguage = "albanian"
	BrandIdentifyFromTransactionParamsForceLanguageAmharic        BrandIdentifyFromTransactionParamsForceLanguage = "amharic"
	BrandIdentifyFromTransactionParamsForceLanguageArabic         BrandIdentifyFromTransactionParamsForceLanguage = "arabic"
	BrandIdentifyFromTransactionParamsForceLanguageArmenian       BrandIdentifyFromTransactionParamsForceLanguage = "armenian"
	BrandIdentifyFromTransactionParamsForceLanguageAssamese       BrandIdentifyFromTransactionParamsForceLanguage = "assamese"
	BrandIdentifyFromTransactionParamsForceLanguageAymara         BrandIdentifyFromTransactionParamsForceLanguage = "aymara"
	BrandIdentifyFromTransactionParamsForceLanguageAzeri          BrandIdentifyFromTransactionParamsForceLanguage = "azeri"
	BrandIdentifyFromTransactionParamsForceLanguageBasque         BrandIdentifyFromTransactionParamsForceLanguage = "basque"
	BrandIdentifyFromTransactionParamsForceLanguageBelarusian     BrandIdentifyFromTransactionParamsForceLanguage = "belarusian"
	BrandIdentifyFromTransactionParamsForceLanguageBengali        BrandIdentifyFromTransactionParamsForceLanguage = "bengali"
	BrandIdentifyFromTransactionParamsForceLanguageBosnian        BrandIdentifyFromTransactionParamsForceLanguage = "bosnian"
	BrandIdentifyFromTransactionParamsForceLanguageBulgarian      BrandIdentifyFromTransactionParamsForceLanguage = "bulgarian"
	BrandIdentifyFromTransactionParamsForceLanguageBurmese        BrandIdentifyFromTransactionParamsForceLanguage = "burmese"
	BrandIdentifyFromTransactionParamsForceLanguageCantonese      BrandIdentifyFromTransactionParamsForceLanguage = "cantonese"
	BrandIdentifyFromTransactionParamsForceLanguageCatalan        BrandIdentifyFromTransactionParamsForceLanguage = "catalan"
	BrandIdentifyFromTransactionParamsForceLanguageCebuano        BrandIdentifyFromTransactionParamsForceLanguage = "cebuano"
	BrandIdentifyFromTransactionParamsForceLanguageChinese        BrandIdentifyFromTransactionParamsForceLanguage = "chinese"
	BrandIdentifyFromTransactionParamsForceLanguageCorsican       BrandIdentifyFromTransactionParamsForceLanguage = "corsican"
	BrandIdentifyFromTransactionParamsForceLanguageCroatian       BrandIdentifyFromTransactionParamsForceLanguage = "croatian"
	BrandIdentifyFromTransactionParamsForceLanguageCzech          BrandIdentifyFromTransactionParamsForceLanguage = "czech"
	BrandIdentifyFromTransactionParamsForceLanguageDanish         BrandIdentifyFromTransactionParamsForceLanguage = "danish"
	BrandIdentifyFromTransactionParamsForceLanguageDutch          BrandIdentifyFromTransactionParamsForceLanguage = "dutch"
	BrandIdentifyFromTransactionParamsForceLanguageEnglish        BrandIdentifyFromTransactionParamsForceLanguage = "english"
	BrandIdentifyFromTransactionParamsForceLanguageEsperanto      BrandIdentifyFromTransactionParamsForceLanguage = "esperanto"
	BrandIdentifyFromTransactionParamsForceLanguageEstonian       BrandIdentifyFromTransactionParamsForceLanguage = "estonian"
	BrandIdentifyFromTransactionParamsForceLanguageFarsi          BrandIdentifyFromTransactionParamsForceLanguage = "farsi"
	BrandIdentifyFromTransactionParamsForceLanguageFijian         BrandIdentifyFromTransactionParamsForceLanguage = "fijian"
	BrandIdentifyFromTransactionParamsForceLanguageFinnish        BrandIdentifyFromTransactionParamsForceLanguage = "finnish"
	BrandIdentifyFromTransactionParamsForceLanguageFrench         BrandIdentifyFromTransactionParamsForceLanguage = "french"
	BrandIdentifyFromTransactionParamsForceLanguageGalician       BrandIdentifyFromTransactionParamsForceLanguage = "galician"
	BrandIdentifyFromTransactionParamsForceLanguageGeorgian       BrandIdentifyFromTransactionParamsForceLanguage = "georgian"
	BrandIdentifyFromTransactionParamsForceLanguageGerman         BrandIdentifyFromTransactionParamsForceLanguage = "german"
	BrandIdentifyFromTransactionParamsForceLanguageGreek          BrandIdentifyFromTransactionParamsForceLanguage = "greek"
	BrandIdentifyFromTransactionParamsForceLanguageGuarani        BrandIdentifyFromTransactionParamsForceLanguage = "guarani"
	BrandIdentifyFromTransactionParamsForceLanguageGujarati       BrandIdentifyFromTransactionParamsForceLanguage = "gujarati"
	BrandIdentifyFromTransactionParamsForceLanguageHaitianCreole  BrandIdentifyFromTransactionParamsForceLanguage = "haitian-creole"
	BrandIdentifyFromTransactionParamsForceLanguageHausa          BrandIdentifyFromTransactionParamsForceLanguage = "hausa"
	BrandIdentifyFromTransactionParamsForceLanguageHawaiian       BrandIdentifyFromTransactionParamsForceLanguage = "hawaiian"
	BrandIdentifyFromTransactionParamsForceLanguageHebrew         BrandIdentifyFromTransactionParamsForceLanguage = "hebrew"
	BrandIdentifyFromTransactionParamsForceLanguageHindi          BrandIdentifyFromTransactionParamsForceLanguage = "hindi"
	BrandIdentifyFromTransactionParamsForceLanguageHmong          BrandIdentifyFromTransactionParamsForceLanguage = "hmong"
	BrandIdentifyFromTransactionParamsForceLanguageHungarian      BrandIdentifyFromTransactionParamsForceLanguage = "hungarian"
	BrandIdentifyFromTransactionParamsForceLanguageIcelandic      BrandIdentifyFromTransactionParamsForceLanguage = "icelandic"
	BrandIdentifyFromTransactionParamsForceLanguageIgbo           BrandIdentifyFromTransactionParamsForceLanguage = "igbo"
	BrandIdentifyFromTransactionParamsForceLanguageIndonesian     BrandIdentifyFromTransactionParamsForceLanguage = "indonesian"
	BrandIdentifyFromTransactionParamsForceLanguageIrish          BrandIdentifyFromTransactionParamsForceLanguage = "irish"
	BrandIdentifyFromTransactionParamsForceLanguageItalian        BrandIdentifyFromTransactionParamsForceLanguage = "italian"
	BrandIdentifyFromTransactionParamsForceLanguageJapanese       BrandIdentifyFromTransactionParamsForceLanguage = "japanese"
	BrandIdentifyFromTransactionParamsForceLanguageJavanese       BrandIdentifyFromTransactionParamsForceLanguage = "javanese"
	BrandIdentifyFromTransactionParamsForceLanguageKannada        BrandIdentifyFromTransactionParamsForceLanguage = "kannada"
	BrandIdentifyFromTransactionParamsForceLanguageKazakh         BrandIdentifyFromTransactionParamsForceLanguage = "kazakh"
	BrandIdentifyFromTransactionParamsForceLanguageKhmer          BrandIdentifyFromTransactionParamsForceLanguage = "khmer"
	BrandIdentifyFromTransactionParamsForceLanguageKinyarwanda    BrandIdentifyFromTransactionParamsForceLanguage = "kinyarwanda"
	BrandIdentifyFromTransactionParamsForceLanguageKorean         BrandIdentifyFromTransactionParamsForceLanguage = "korean"
	BrandIdentifyFromTransactionParamsForceLanguageKurdish        BrandIdentifyFromTransactionParamsForceLanguage = "kurdish"
	BrandIdentifyFromTransactionParamsForceLanguageKyrgyz         BrandIdentifyFromTransactionParamsForceLanguage = "kyrgyz"
	BrandIdentifyFromTransactionParamsForceLanguageLao            BrandIdentifyFromTransactionParamsForceLanguage = "lao"
	BrandIdentifyFromTransactionParamsForceLanguageLatin          BrandIdentifyFromTransactionParamsForceLanguage = "latin"
	BrandIdentifyFromTransactionParamsForceLanguageLatvian        BrandIdentifyFromTransactionParamsForceLanguage = "latvian"
	BrandIdentifyFromTransactionParamsForceLanguageLingala        BrandIdentifyFromTransactionParamsForceLanguage = "lingala"
	BrandIdentifyFromTransactionParamsForceLanguageLithuanian     BrandIdentifyFromTransactionParamsForceLanguage = "lithuanian"
	BrandIdentifyFromTransactionParamsForceLanguageLuxembourgish  BrandIdentifyFromTransactionParamsForceLanguage = "luxembourgish"
	BrandIdentifyFromTransactionParamsForceLanguageMacedonian     BrandIdentifyFromTransactionParamsForceLanguage = "macedonian"
	BrandIdentifyFromTransactionParamsForceLanguageMalagasy       BrandIdentifyFromTransactionParamsForceLanguage = "malagasy"
	BrandIdentifyFromTransactionParamsForceLanguageMalay          BrandIdentifyFromTransactionParamsForceLanguage = "malay"
	BrandIdentifyFromTransactionParamsForceLanguageMalayalam      BrandIdentifyFromTransactionParamsForceLanguage = "malayalam"
	BrandIdentifyFromTransactionParamsForceLanguageMaltese        BrandIdentifyFromTransactionParamsForceLanguage = "maltese"
	BrandIdentifyFromTransactionParamsForceLanguageMaori          BrandIdentifyFromTransactionParamsForceLanguage = "maori"
	BrandIdentifyFromTransactionParamsForceLanguageMarathi        BrandIdentifyFromTransactionParamsForceLanguage = "marathi"
	BrandIdentifyFromTransactionParamsForceLanguageMongolian      BrandIdentifyFromTransactionParamsForceLanguage = "mongolian"
	BrandIdentifyFromTransactionParamsForceLanguageNepali         BrandIdentifyFromTransactionParamsForceLanguage = "nepali"
	BrandIdentifyFromTransactionParamsForceLanguageNorwegian      BrandIdentifyFromTransactionParamsForceLanguage = "norwegian"
	BrandIdentifyFromTransactionParamsForceLanguageOdia           BrandIdentifyFromTransactionParamsForceLanguage = "odia"
	BrandIdentifyFromTransactionParamsForceLanguageOromo          BrandIdentifyFromTransactionParamsForceLanguage = "oromo"
	BrandIdentifyFromTransactionParamsForceLanguagePashto         BrandIdentifyFromTransactionParamsForceLanguage = "pashto"
	BrandIdentifyFromTransactionParamsForceLanguagePidgin         BrandIdentifyFromTransactionParamsForceLanguage = "pidgin"
	BrandIdentifyFromTransactionParamsForceLanguagePolish         BrandIdentifyFromTransactionParamsForceLanguage = "polish"
	BrandIdentifyFromTransactionParamsForceLanguagePortuguese     BrandIdentifyFromTransactionParamsForceLanguage = "portuguese"
	BrandIdentifyFromTransactionParamsForceLanguagePunjabi        BrandIdentifyFromTransactionParamsForceLanguage = "punjabi"
	BrandIdentifyFromTransactionParamsForceLanguageQuechua        BrandIdentifyFromTransactionParamsForceLanguage = "quechua"
	BrandIdentifyFromTransactionParamsForceLanguageRomanian       BrandIdentifyFromTransactionParamsForceLanguage = "romanian"
	BrandIdentifyFromTransactionParamsForceLanguageRussian        BrandIdentifyFromTransactionParamsForceLanguage = "russian"
	BrandIdentifyFromTransactionParamsForceLanguageSamoan         BrandIdentifyFromTransactionParamsForceLanguage = "samoan"
	BrandIdentifyFromTransactionParamsForceLanguageScottishGaelic BrandIdentifyFromTransactionParamsForceLanguage = "scottish-gaelic"
	BrandIdentifyFromTransactionParamsForceLanguageSerbian        BrandIdentifyFromTransactionParamsForceLanguage = "serbian"
	BrandIdentifyFromTransactionParamsForceLanguageSesotho        BrandIdentifyFromTransactionParamsForceLanguage = "sesotho"
	BrandIdentifyFromTransactionParamsForceLanguageShona          BrandIdentifyFromTransactionParamsForceLanguage = "shona"
	BrandIdentifyFromTransactionParamsForceLanguageSindhi         BrandIdentifyFromTransactionParamsForceLanguage = "sindhi"
	BrandIdentifyFromTransactionParamsForceLanguageSinhala        BrandIdentifyFromTransactionParamsForceLanguage = "sinhala"
	BrandIdentifyFromTransactionParamsForceLanguageSlovak         BrandIdentifyFromTransactionParamsForceLanguage = "slovak"
	BrandIdentifyFromTransactionParamsForceLanguageSlovene        BrandIdentifyFromTransactionParamsForceLanguage = "slovene"
	BrandIdentifyFromTransactionParamsForceLanguageSomali         BrandIdentifyFromTransactionParamsForceLanguage = "somali"
	BrandIdentifyFromTransactionParamsForceLanguageSpanish        BrandIdentifyFromTransactionParamsForceLanguage = "spanish"
	BrandIdentifyFromTransactionParamsForceLanguageSundanese      BrandIdentifyFromTransactionParamsForceLanguage = "sundanese"
	BrandIdentifyFromTransactionParamsForceLanguageSwahili        BrandIdentifyFromTransactionParamsForceLanguage = "swahili"
	BrandIdentifyFromTransactionParamsForceLanguageSwedish        BrandIdentifyFromTransactionParamsForceLanguage = "swedish"
	BrandIdentifyFromTransactionParamsForceLanguageTagalog        BrandIdentifyFromTransactionParamsForceLanguage = "tagalog"
	BrandIdentifyFromTransactionParamsForceLanguageTajik          BrandIdentifyFromTransactionParamsForceLanguage = "tajik"
	BrandIdentifyFromTransactionParamsForceLanguageTamil          BrandIdentifyFromTransactionParamsForceLanguage = "tamil"
	BrandIdentifyFromTransactionParamsForceLanguageTatar          BrandIdentifyFromTransactionParamsForceLanguage = "tatar"
	BrandIdentifyFromTransactionParamsForceLanguageTelugu         BrandIdentifyFromTransactionParamsForceLanguage = "telugu"
	BrandIdentifyFromTransactionParamsForceLanguageThai           BrandIdentifyFromTransactionParamsForceLanguage = "thai"
	BrandIdentifyFromTransactionParamsForceLanguageTibetan        BrandIdentifyFromTransactionParamsForceLanguage = "tibetan"
	BrandIdentifyFromTransactionParamsForceLanguageTigrinya       BrandIdentifyFromTransactionParamsForceLanguage = "tigrinya"
	BrandIdentifyFromTransactionParamsForceLanguageTongan         BrandIdentifyFromTransactionParamsForceLanguage = "tongan"
	BrandIdentifyFromTransactionParamsForceLanguageTswana         BrandIdentifyFromTransactionParamsForceLanguage = "tswana"
	BrandIdentifyFromTransactionParamsForceLanguageTurkish        BrandIdentifyFromTransactionParamsForceLanguage = "turkish"
	BrandIdentifyFromTransactionParamsForceLanguageTurkmen        BrandIdentifyFromTransactionParamsForceLanguage = "turkmen"
	BrandIdentifyFromTransactionParamsForceLanguageUkrainian      BrandIdentifyFromTransactionParamsForceLanguage = "ukrainian"
	BrandIdentifyFromTransactionParamsForceLanguageUrdu           BrandIdentifyFromTransactionParamsForceLanguage = "urdu"
	BrandIdentifyFromTransactionParamsForceLanguageUyghur         BrandIdentifyFromTransactionParamsForceLanguage = "uyghur"
	BrandIdentifyFromTransactionParamsForceLanguageUzbek          BrandIdentifyFromTransactionParamsForceLanguage = "uzbek"
	BrandIdentifyFromTransactionParamsForceLanguageVietnamese     BrandIdentifyFromTransactionParamsForceLanguage = "vietnamese"
	BrandIdentifyFromTransactionParamsForceLanguageWelsh          BrandIdentifyFromTransactionParamsForceLanguage = "welsh"
	BrandIdentifyFromTransactionParamsForceLanguageWolof          BrandIdentifyFromTransactionParamsForceLanguage = "wolof"
	BrandIdentifyFromTransactionParamsForceLanguageXhosa          BrandIdentifyFromTransactionParamsForceLanguage = "xhosa"
	BrandIdentifyFromTransactionParamsForceLanguageYiddish        BrandIdentifyFromTransactionParamsForceLanguage = "yiddish"
	BrandIdentifyFromTransactionParamsForceLanguageYoruba         BrandIdentifyFromTransactionParamsForceLanguage = "yoruba"
	BrandIdentifyFromTransactionParamsForceLanguageZulu           BrandIdentifyFromTransactionParamsForceLanguage = "zulu"
)

type BrandGetByEmailParams struct {
	// Email address to retrieve brand data for (e.g., 'contact@example.com'). The
	// domain will be extracted from the email. Free email providers (gmail.com,
	// yahoo.com, etc.) and disposable email addresses are not allowed.
	Email string `query:"email" api:"required" format:"email" json:"-"`
	// Maximum age in milliseconds for cached brand data before the API performs a hard
	// refresh. Defaults to 3 months (7776000000 ms). Values below 1 day (86400000 ms)
	// are clamped to 1 day; values above 1 year (31536000000 ms) are clamped to 1
	// year.
	MaxAgeMs param.Opt[int64] `query:"maxAgeMs,omitzero" json:"-"`
	// Optional parameter to optimize the API call for maximum speed. When set to true,
	// the API will skip time-consuming operations for faster response at the cost of
	// less comprehensive data.
	MaxSpeed param.Opt[bool] `query:"maxSpeed,omitzero" json:"-"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `query:"timeoutMS,omitzero" json:"-"`
	// Optional parameter to force the language of the retrieved brand data.
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
	ForceLanguage BrandGetByEmailParamsForceLanguage `query:"force_language,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BrandGetByEmailParams]'s query parameters as `url.Values`.
func (r BrandGetByEmailParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Optional parameter to force the language of the retrieved brand data.
type BrandGetByEmailParamsForceLanguage string

const (
	BrandGetByEmailParamsForceLanguageAfrikaans      BrandGetByEmailParamsForceLanguage = "afrikaans"
	BrandGetByEmailParamsForceLanguageAlbanian       BrandGetByEmailParamsForceLanguage = "albanian"
	BrandGetByEmailParamsForceLanguageAmharic        BrandGetByEmailParamsForceLanguage = "amharic"
	BrandGetByEmailParamsForceLanguageArabic         BrandGetByEmailParamsForceLanguage = "arabic"
	BrandGetByEmailParamsForceLanguageArmenian       BrandGetByEmailParamsForceLanguage = "armenian"
	BrandGetByEmailParamsForceLanguageAssamese       BrandGetByEmailParamsForceLanguage = "assamese"
	BrandGetByEmailParamsForceLanguageAymara         BrandGetByEmailParamsForceLanguage = "aymara"
	BrandGetByEmailParamsForceLanguageAzeri          BrandGetByEmailParamsForceLanguage = "azeri"
	BrandGetByEmailParamsForceLanguageBasque         BrandGetByEmailParamsForceLanguage = "basque"
	BrandGetByEmailParamsForceLanguageBelarusian     BrandGetByEmailParamsForceLanguage = "belarusian"
	BrandGetByEmailParamsForceLanguageBengali        BrandGetByEmailParamsForceLanguage = "bengali"
	BrandGetByEmailParamsForceLanguageBosnian        BrandGetByEmailParamsForceLanguage = "bosnian"
	BrandGetByEmailParamsForceLanguageBulgarian      BrandGetByEmailParamsForceLanguage = "bulgarian"
	BrandGetByEmailParamsForceLanguageBurmese        BrandGetByEmailParamsForceLanguage = "burmese"
	BrandGetByEmailParamsForceLanguageCantonese      BrandGetByEmailParamsForceLanguage = "cantonese"
	BrandGetByEmailParamsForceLanguageCatalan        BrandGetByEmailParamsForceLanguage = "catalan"
	BrandGetByEmailParamsForceLanguageCebuano        BrandGetByEmailParamsForceLanguage = "cebuano"
	BrandGetByEmailParamsForceLanguageChinese        BrandGetByEmailParamsForceLanguage = "chinese"
	BrandGetByEmailParamsForceLanguageCorsican       BrandGetByEmailParamsForceLanguage = "corsican"
	BrandGetByEmailParamsForceLanguageCroatian       BrandGetByEmailParamsForceLanguage = "croatian"
	BrandGetByEmailParamsForceLanguageCzech          BrandGetByEmailParamsForceLanguage = "czech"
	BrandGetByEmailParamsForceLanguageDanish         BrandGetByEmailParamsForceLanguage = "danish"
	BrandGetByEmailParamsForceLanguageDutch          BrandGetByEmailParamsForceLanguage = "dutch"
	BrandGetByEmailParamsForceLanguageEnglish        BrandGetByEmailParamsForceLanguage = "english"
	BrandGetByEmailParamsForceLanguageEsperanto      BrandGetByEmailParamsForceLanguage = "esperanto"
	BrandGetByEmailParamsForceLanguageEstonian       BrandGetByEmailParamsForceLanguage = "estonian"
	BrandGetByEmailParamsForceLanguageFarsi          BrandGetByEmailParamsForceLanguage = "farsi"
	BrandGetByEmailParamsForceLanguageFijian         BrandGetByEmailParamsForceLanguage = "fijian"
	BrandGetByEmailParamsForceLanguageFinnish        BrandGetByEmailParamsForceLanguage = "finnish"
	BrandGetByEmailParamsForceLanguageFrench         BrandGetByEmailParamsForceLanguage = "french"
	BrandGetByEmailParamsForceLanguageGalician       BrandGetByEmailParamsForceLanguage = "galician"
	BrandGetByEmailParamsForceLanguageGeorgian       BrandGetByEmailParamsForceLanguage = "georgian"
	BrandGetByEmailParamsForceLanguageGerman         BrandGetByEmailParamsForceLanguage = "german"
	BrandGetByEmailParamsForceLanguageGreek          BrandGetByEmailParamsForceLanguage = "greek"
	BrandGetByEmailParamsForceLanguageGuarani        BrandGetByEmailParamsForceLanguage = "guarani"
	BrandGetByEmailParamsForceLanguageGujarati       BrandGetByEmailParamsForceLanguage = "gujarati"
	BrandGetByEmailParamsForceLanguageHaitianCreole  BrandGetByEmailParamsForceLanguage = "haitian-creole"
	BrandGetByEmailParamsForceLanguageHausa          BrandGetByEmailParamsForceLanguage = "hausa"
	BrandGetByEmailParamsForceLanguageHawaiian       BrandGetByEmailParamsForceLanguage = "hawaiian"
	BrandGetByEmailParamsForceLanguageHebrew         BrandGetByEmailParamsForceLanguage = "hebrew"
	BrandGetByEmailParamsForceLanguageHindi          BrandGetByEmailParamsForceLanguage = "hindi"
	BrandGetByEmailParamsForceLanguageHmong          BrandGetByEmailParamsForceLanguage = "hmong"
	BrandGetByEmailParamsForceLanguageHungarian      BrandGetByEmailParamsForceLanguage = "hungarian"
	BrandGetByEmailParamsForceLanguageIcelandic      BrandGetByEmailParamsForceLanguage = "icelandic"
	BrandGetByEmailParamsForceLanguageIgbo           BrandGetByEmailParamsForceLanguage = "igbo"
	BrandGetByEmailParamsForceLanguageIndonesian     BrandGetByEmailParamsForceLanguage = "indonesian"
	BrandGetByEmailParamsForceLanguageIrish          BrandGetByEmailParamsForceLanguage = "irish"
	BrandGetByEmailParamsForceLanguageItalian        BrandGetByEmailParamsForceLanguage = "italian"
	BrandGetByEmailParamsForceLanguageJapanese       BrandGetByEmailParamsForceLanguage = "japanese"
	BrandGetByEmailParamsForceLanguageJavanese       BrandGetByEmailParamsForceLanguage = "javanese"
	BrandGetByEmailParamsForceLanguageKannada        BrandGetByEmailParamsForceLanguage = "kannada"
	BrandGetByEmailParamsForceLanguageKazakh         BrandGetByEmailParamsForceLanguage = "kazakh"
	BrandGetByEmailParamsForceLanguageKhmer          BrandGetByEmailParamsForceLanguage = "khmer"
	BrandGetByEmailParamsForceLanguageKinyarwanda    BrandGetByEmailParamsForceLanguage = "kinyarwanda"
	BrandGetByEmailParamsForceLanguageKorean         BrandGetByEmailParamsForceLanguage = "korean"
	BrandGetByEmailParamsForceLanguageKurdish        BrandGetByEmailParamsForceLanguage = "kurdish"
	BrandGetByEmailParamsForceLanguageKyrgyz         BrandGetByEmailParamsForceLanguage = "kyrgyz"
	BrandGetByEmailParamsForceLanguageLao            BrandGetByEmailParamsForceLanguage = "lao"
	BrandGetByEmailParamsForceLanguageLatin          BrandGetByEmailParamsForceLanguage = "latin"
	BrandGetByEmailParamsForceLanguageLatvian        BrandGetByEmailParamsForceLanguage = "latvian"
	BrandGetByEmailParamsForceLanguageLingala        BrandGetByEmailParamsForceLanguage = "lingala"
	BrandGetByEmailParamsForceLanguageLithuanian     BrandGetByEmailParamsForceLanguage = "lithuanian"
	BrandGetByEmailParamsForceLanguageLuxembourgish  BrandGetByEmailParamsForceLanguage = "luxembourgish"
	BrandGetByEmailParamsForceLanguageMacedonian     BrandGetByEmailParamsForceLanguage = "macedonian"
	BrandGetByEmailParamsForceLanguageMalagasy       BrandGetByEmailParamsForceLanguage = "malagasy"
	BrandGetByEmailParamsForceLanguageMalay          BrandGetByEmailParamsForceLanguage = "malay"
	BrandGetByEmailParamsForceLanguageMalayalam      BrandGetByEmailParamsForceLanguage = "malayalam"
	BrandGetByEmailParamsForceLanguageMaltese        BrandGetByEmailParamsForceLanguage = "maltese"
	BrandGetByEmailParamsForceLanguageMaori          BrandGetByEmailParamsForceLanguage = "maori"
	BrandGetByEmailParamsForceLanguageMarathi        BrandGetByEmailParamsForceLanguage = "marathi"
	BrandGetByEmailParamsForceLanguageMongolian      BrandGetByEmailParamsForceLanguage = "mongolian"
	BrandGetByEmailParamsForceLanguageNepali         BrandGetByEmailParamsForceLanguage = "nepali"
	BrandGetByEmailParamsForceLanguageNorwegian      BrandGetByEmailParamsForceLanguage = "norwegian"
	BrandGetByEmailParamsForceLanguageOdia           BrandGetByEmailParamsForceLanguage = "odia"
	BrandGetByEmailParamsForceLanguageOromo          BrandGetByEmailParamsForceLanguage = "oromo"
	BrandGetByEmailParamsForceLanguagePashto         BrandGetByEmailParamsForceLanguage = "pashto"
	BrandGetByEmailParamsForceLanguagePidgin         BrandGetByEmailParamsForceLanguage = "pidgin"
	BrandGetByEmailParamsForceLanguagePolish         BrandGetByEmailParamsForceLanguage = "polish"
	BrandGetByEmailParamsForceLanguagePortuguese     BrandGetByEmailParamsForceLanguage = "portuguese"
	BrandGetByEmailParamsForceLanguagePunjabi        BrandGetByEmailParamsForceLanguage = "punjabi"
	BrandGetByEmailParamsForceLanguageQuechua        BrandGetByEmailParamsForceLanguage = "quechua"
	BrandGetByEmailParamsForceLanguageRomanian       BrandGetByEmailParamsForceLanguage = "romanian"
	BrandGetByEmailParamsForceLanguageRussian        BrandGetByEmailParamsForceLanguage = "russian"
	BrandGetByEmailParamsForceLanguageSamoan         BrandGetByEmailParamsForceLanguage = "samoan"
	BrandGetByEmailParamsForceLanguageScottishGaelic BrandGetByEmailParamsForceLanguage = "scottish-gaelic"
	BrandGetByEmailParamsForceLanguageSerbian        BrandGetByEmailParamsForceLanguage = "serbian"
	BrandGetByEmailParamsForceLanguageSesotho        BrandGetByEmailParamsForceLanguage = "sesotho"
	BrandGetByEmailParamsForceLanguageShona          BrandGetByEmailParamsForceLanguage = "shona"
	BrandGetByEmailParamsForceLanguageSindhi         BrandGetByEmailParamsForceLanguage = "sindhi"
	BrandGetByEmailParamsForceLanguageSinhala        BrandGetByEmailParamsForceLanguage = "sinhala"
	BrandGetByEmailParamsForceLanguageSlovak         BrandGetByEmailParamsForceLanguage = "slovak"
	BrandGetByEmailParamsForceLanguageSlovene        BrandGetByEmailParamsForceLanguage = "slovene"
	BrandGetByEmailParamsForceLanguageSomali         BrandGetByEmailParamsForceLanguage = "somali"
	BrandGetByEmailParamsForceLanguageSpanish        BrandGetByEmailParamsForceLanguage = "spanish"
	BrandGetByEmailParamsForceLanguageSundanese      BrandGetByEmailParamsForceLanguage = "sundanese"
	BrandGetByEmailParamsForceLanguageSwahili        BrandGetByEmailParamsForceLanguage = "swahili"
	BrandGetByEmailParamsForceLanguageSwedish        BrandGetByEmailParamsForceLanguage = "swedish"
	BrandGetByEmailParamsForceLanguageTagalog        BrandGetByEmailParamsForceLanguage = "tagalog"
	BrandGetByEmailParamsForceLanguageTajik          BrandGetByEmailParamsForceLanguage = "tajik"
	BrandGetByEmailParamsForceLanguageTamil          BrandGetByEmailParamsForceLanguage = "tamil"
	BrandGetByEmailParamsForceLanguageTatar          BrandGetByEmailParamsForceLanguage = "tatar"
	BrandGetByEmailParamsForceLanguageTelugu         BrandGetByEmailParamsForceLanguage = "telugu"
	BrandGetByEmailParamsForceLanguageThai           BrandGetByEmailParamsForceLanguage = "thai"
	BrandGetByEmailParamsForceLanguageTibetan        BrandGetByEmailParamsForceLanguage = "tibetan"
	BrandGetByEmailParamsForceLanguageTigrinya       BrandGetByEmailParamsForceLanguage = "tigrinya"
	BrandGetByEmailParamsForceLanguageTongan         BrandGetByEmailParamsForceLanguage = "tongan"
	BrandGetByEmailParamsForceLanguageTswana         BrandGetByEmailParamsForceLanguage = "tswana"
	BrandGetByEmailParamsForceLanguageTurkish        BrandGetByEmailParamsForceLanguage = "turkish"
	BrandGetByEmailParamsForceLanguageTurkmen        BrandGetByEmailParamsForceLanguage = "turkmen"
	BrandGetByEmailParamsForceLanguageUkrainian      BrandGetByEmailParamsForceLanguage = "ukrainian"
	BrandGetByEmailParamsForceLanguageUrdu           BrandGetByEmailParamsForceLanguage = "urdu"
	BrandGetByEmailParamsForceLanguageUyghur         BrandGetByEmailParamsForceLanguage = "uyghur"
	BrandGetByEmailParamsForceLanguageUzbek          BrandGetByEmailParamsForceLanguage = "uzbek"
	BrandGetByEmailParamsForceLanguageVietnamese     BrandGetByEmailParamsForceLanguage = "vietnamese"
	BrandGetByEmailParamsForceLanguageWelsh          BrandGetByEmailParamsForceLanguage = "welsh"
	BrandGetByEmailParamsForceLanguageWolof          BrandGetByEmailParamsForceLanguage = "wolof"
	BrandGetByEmailParamsForceLanguageXhosa          BrandGetByEmailParamsForceLanguage = "xhosa"
	BrandGetByEmailParamsForceLanguageYiddish        BrandGetByEmailParamsForceLanguage = "yiddish"
	BrandGetByEmailParamsForceLanguageYoruba         BrandGetByEmailParamsForceLanguage = "yoruba"
	BrandGetByEmailParamsForceLanguageZulu           BrandGetByEmailParamsForceLanguage = "zulu"
)

type BrandGetByIsinParams struct {
	// ISIN (International Securities Identification Number) to retrieve brand data for
	// (e.g., 'AU000000IMD5', 'US0378331005'). Must be exactly 12 characters: 2 letters
	// followed by 9 alphanumeric characters and ending with a digit.
	Isin string `query:"isin" api:"required" json:"-"`
	// Maximum age in milliseconds for cached brand data before the API performs a hard
	// refresh. Defaults to 3 months (7776000000 ms). Values below 1 day (86400000 ms)
	// are clamped to 1 day; values above 1 year (31536000000 ms) are clamped to 1
	// year.
	MaxAgeMs param.Opt[int64] `query:"maxAgeMs,omitzero" json:"-"`
	// Optional parameter to optimize the API call for maximum speed. When set to true,
	// the API will skip time-consuming operations for faster response at the cost of
	// less comprehensive data.
	MaxSpeed param.Opt[bool] `query:"maxSpeed,omitzero" json:"-"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `query:"timeoutMS,omitzero" json:"-"`
	// Optional parameter to force the language of the retrieved brand data.
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
	ForceLanguage BrandGetByIsinParamsForceLanguage `query:"force_language,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BrandGetByIsinParams]'s query parameters as `url.Values`.
func (r BrandGetByIsinParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Optional parameter to force the language of the retrieved brand data.
type BrandGetByIsinParamsForceLanguage string

const (
	BrandGetByIsinParamsForceLanguageAfrikaans      BrandGetByIsinParamsForceLanguage = "afrikaans"
	BrandGetByIsinParamsForceLanguageAlbanian       BrandGetByIsinParamsForceLanguage = "albanian"
	BrandGetByIsinParamsForceLanguageAmharic        BrandGetByIsinParamsForceLanguage = "amharic"
	BrandGetByIsinParamsForceLanguageArabic         BrandGetByIsinParamsForceLanguage = "arabic"
	BrandGetByIsinParamsForceLanguageArmenian       BrandGetByIsinParamsForceLanguage = "armenian"
	BrandGetByIsinParamsForceLanguageAssamese       BrandGetByIsinParamsForceLanguage = "assamese"
	BrandGetByIsinParamsForceLanguageAymara         BrandGetByIsinParamsForceLanguage = "aymara"
	BrandGetByIsinParamsForceLanguageAzeri          BrandGetByIsinParamsForceLanguage = "azeri"
	BrandGetByIsinParamsForceLanguageBasque         BrandGetByIsinParamsForceLanguage = "basque"
	BrandGetByIsinParamsForceLanguageBelarusian     BrandGetByIsinParamsForceLanguage = "belarusian"
	BrandGetByIsinParamsForceLanguageBengali        BrandGetByIsinParamsForceLanguage = "bengali"
	BrandGetByIsinParamsForceLanguageBosnian        BrandGetByIsinParamsForceLanguage = "bosnian"
	BrandGetByIsinParamsForceLanguageBulgarian      BrandGetByIsinParamsForceLanguage = "bulgarian"
	BrandGetByIsinParamsForceLanguageBurmese        BrandGetByIsinParamsForceLanguage = "burmese"
	BrandGetByIsinParamsForceLanguageCantonese      BrandGetByIsinParamsForceLanguage = "cantonese"
	BrandGetByIsinParamsForceLanguageCatalan        BrandGetByIsinParamsForceLanguage = "catalan"
	BrandGetByIsinParamsForceLanguageCebuano        BrandGetByIsinParamsForceLanguage = "cebuano"
	BrandGetByIsinParamsForceLanguageChinese        BrandGetByIsinParamsForceLanguage = "chinese"
	BrandGetByIsinParamsForceLanguageCorsican       BrandGetByIsinParamsForceLanguage = "corsican"
	BrandGetByIsinParamsForceLanguageCroatian       BrandGetByIsinParamsForceLanguage = "croatian"
	BrandGetByIsinParamsForceLanguageCzech          BrandGetByIsinParamsForceLanguage = "czech"
	BrandGetByIsinParamsForceLanguageDanish         BrandGetByIsinParamsForceLanguage = "danish"
	BrandGetByIsinParamsForceLanguageDutch          BrandGetByIsinParamsForceLanguage = "dutch"
	BrandGetByIsinParamsForceLanguageEnglish        BrandGetByIsinParamsForceLanguage = "english"
	BrandGetByIsinParamsForceLanguageEsperanto      BrandGetByIsinParamsForceLanguage = "esperanto"
	BrandGetByIsinParamsForceLanguageEstonian       BrandGetByIsinParamsForceLanguage = "estonian"
	BrandGetByIsinParamsForceLanguageFarsi          BrandGetByIsinParamsForceLanguage = "farsi"
	BrandGetByIsinParamsForceLanguageFijian         BrandGetByIsinParamsForceLanguage = "fijian"
	BrandGetByIsinParamsForceLanguageFinnish        BrandGetByIsinParamsForceLanguage = "finnish"
	BrandGetByIsinParamsForceLanguageFrench         BrandGetByIsinParamsForceLanguage = "french"
	BrandGetByIsinParamsForceLanguageGalician       BrandGetByIsinParamsForceLanguage = "galician"
	BrandGetByIsinParamsForceLanguageGeorgian       BrandGetByIsinParamsForceLanguage = "georgian"
	BrandGetByIsinParamsForceLanguageGerman         BrandGetByIsinParamsForceLanguage = "german"
	BrandGetByIsinParamsForceLanguageGreek          BrandGetByIsinParamsForceLanguage = "greek"
	BrandGetByIsinParamsForceLanguageGuarani        BrandGetByIsinParamsForceLanguage = "guarani"
	BrandGetByIsinParamsForceLanguageGujarati       BrandGetByIsinParamsForceLanguage = "gujarati"
	BrandGetByIsinParamsForceLanguageHaitianCreole  BrandGetByIsinParamsForceLanguage = "haitian-creole"
	BrandGetByIsinParamsForceLanguageHausa          BrandGetByIsinParamsForceLanguage = "hausa"
	BrandGetByIsinParamsForceLanguageHawaiian       BrandGetByIsinParamsForceLanguage = "hawaiian"
	BrandGetByIsinParamsForceLanguageHebrew         BrandGetByIsinParamsForceLanguage = "hebrew"
	BrandGetByIsinParamsForceLanguageHindi          BrandGetByIsinParamsForceLanguage = "hindi"
	BrandGetByIsinParamsForceLanguageHmong          BrandGetByIsinParamsForceLanguage = "hmong"
	BrandGetByIsinParamsForceLanguageHungarian      BrandGetByIsinParamsForceLanguage = "hungarian"
	BrandGetByIsinParamsForceLanguageIcelandic      BrandGetByIsinParamsForceLanguage = "icelandic"
	BrandGetByIsinParamsForceLanguageIgbo           BrandGetByIsinParamsForceLanguage = "igbo"
	BrandGetByIsinParamsForceLanguageIndonesian     BrandGetByIsinParamsForceLanguage = "indonesian"
	BrandGetByIsinParamsForceLanguageIrish          BrandGetByIsinParamsForceLanguage = "irish"
	BrandGetByIsinParamsForceLanguageItalian        BrandGetByIsinParamsForceLanguage = "italian"
	BrandGetByIsinParamsForceLanguageJapanese       BrandGetByIsinParamsForceLanguage = "japanese"
	BrandGetByIsinParamsForceLanguageJavanese       BrandGetByIsinParamsForceLanguage = "javanese"
	BrandGetByIsinParamsForceLanguageKannada        BrandGetByIsinParamsForceLanguage = "kannada"
	BrandGetByIsinParamsForceLanguageKazakh         BrandGetByIsinParamsForceLanguage = "kazakh"
	BrandGetByIsinParamsForceLanguageKhmer          BrandGetByIsinParamsForceLanguage = "khmer"
	BrandGetByIsinParamsForceLanguageKinyarwanda    BrandGetByIsinParamsForceLanguage = "kinyarwanda"
	BrandGetByIsinParamsForceLanguageKorean         BrandGetByIsinParamsForceLanguage = "korean"
	BrandGetByIsinParamsForceLanguageKurdish        BrandGetByIsinParamsForceLanguage = "kurdish"
	BrandGetByIsinParamsForceLanguageKyrgyz         BrandGetByIsinParamsForceLanguage = "kyrgyz"
	BrandGetByIsinParamsForceLanguageLao            BrandGetByIsinParamsForceLanguage = "lao"
	BrandGetByIsinParamsForceLanguageLatin          BrandGetByIsinParamsForceLanguage = "latin"
	BrandGetByIsinParamsForceLanguageLatvian        BrandGetByIsinParamsForceLanguage = "latvian"
	BrandGetByIsinParamsForceLanguageLingala        BrandGetByIsinParamsForceLanguage = "lingala"
	BrandGetByIsinParamsForceLanguageLithuanian     BrandGetByIsinParamsForceLanguage = "lithuanian"
	BrandGetByIsinParamsForceLanguageLuxembourgish  BrandGetByIsinParamsForceLanguage = "luxembourgish"
	BrandGetByIsinParamsForceLanguageMacedonian     BrandGetByIsinParamsForceLanguage = "macedonian"
	BrandGetByIsinParamsForceLanguageMalagasy       BrandGetByIsinParamsForceLanguage = "malagasy"
	BrandGetByIsinParamsForceLanguageMalay          BrandGetByIsinParamsForceLanguage = "malay"
	BrandGetByIsinParamsForceLanguageMalayalam      BrandGetByIsinParamsForceLanguage = "malayalam"
	BrandGetByIsinParamsForceLanguageMaltese        BrandGetByIsinParamsForceLanguage = "maltese"
	BrandGetByIsinParamsForceLanguageMaori          BrandGetByIsinParamsForceLanguage = "maori"
	BrandGetByIsinParamsForceLanguageMarathi        BrandGetByIsinParamsForceLanguage = "marathi"
	BrandGetByIsinParamsForceLanguageMongolian      BrandGetByIsinParamsForceLanguage = "mongolian"
	BrandGetByIsinParamsForceLanguageNepali         BrandGetByIsinParamsForceLanguage = "nepali"
	BrandGetByIsinParamsForceLanguageNorwegian      BrandGetByIsinParamsForceLanguage = "norwegian"
	BrandGetByIsinParamsForceLanguageOdia           BrandGetByIsinParamsForceLanguage = "odia"
	BrandGetByIsinParamsForceLanguageOromo          BrandGetByIsinParamsForceLanguage = "oromo"
	BrandGetByIsinParamsForceLanguagePashto         BrandGetByIsinParamsForceLanguage = "pashto"
	BrandGetByIsinParamsForceLanguagePidgin         BrandGetByIsinParamsForceLanguage = "pidgin"
	BrandGetByIsinParamsForceLanguagePolish         BrandGetByIsinParamsForceLanguage = "polish"
	BrandGetByIsinParamsForceLanguagePortuguese     BrandGetByIsinParamsForceLanguage = "portuguese"
	BrandGetByIsinParamsForceLanguagePunjabi        BrandGetByIsinParamsForceLanguage = "punjabi"
	BrandGetByIsinParamsForceLanguageQuechua        BrandGetByIsinParamsForceLanguage = "quechua"
	BrandGetByIsinParamsForceLanguageRomanian       BrandGetByIsinParamsForceLanguage = "romanian"
	BrandGetByIsinParamsForceLanguageRussian        BrandGetByIsinParamsForceLanguage = "russian"
	BrandGetByIsinParamsForceLanguageSamoan         BrandGetByIsinParamsForceLanguage = "samoan"
	BrandGetByIsinParamsForceLanguageScottishGaelic BrandGetByIsinParamsForceLanguage = "scottish-gaelic"
	BrandGetByIsinParamsForceLanguageSerbian        BrandGetByIsinParamsForceLanguage = "serbian"
	BrandGetByIsinParamsForceLanguageSesotho        BrandGetByIsinParamsForceLanguage = "sesotho"
	BrandGetByIsinParamsForceLanguageShona          BrandGetByIsinParamsForceLanguage = "shona"
	BrandGetByIsinParamsForceLanguageSindhi         BrandGetByIsinParamsForceLanguage = "sindhi"
	BrandGetByIsinParamsForceLanguageSinhala        BrandGetByIsinParamsForceLanguage = "sinhala"
	BrandGetByIsinParamsForceLanguageSlovak         BrandGetByIsinParamsForceLanguage = "slovak"
	BrandGetByIsinParamsForceLanguageSlovene        BrandGetByIsinParamsForceLanguage = "slovene"
	BrandGetByIsinParamsForceLanguageSomali         BrandGetByIsinParamsForceLanguage = "somali"
	BrandGetByIsinParamsForceLanguageSpanish        BrandGetByIsinParamsForceLanguage = "spanish"
	BrandGetByIsinParamsForceLanguageSundanese      BrandGetByIsinParamsForceLanguage = "sundanese"
	BrandGetByIsinParamsForceLanguageSwahili        BrandGetByIsinParamsForceLanguage = "swahili"
	BrandGetByIsinParamsForceLanguageSwedish        BrandGetByIsinParamsForceLanguage = "swedish"
	BrandGetByIsinParamsForceLanguageTagalog        BrandGetByIsinParamsForceLanguage = "tagalog"
	BrandGetByIsinParamsForceLanguageTajik          BrandGetByIsinParamsForceLanguage = "tajik"
	BrandGetByIsinParamsForceLanguageTamil          BrandGetByIsinParamsForceLanguage = "tamil"
	BrandGetByIsinParamsForceLanguageTatar          BrandGetByIsinParamsForceLanguage = "tatar"
	BrandGetByIsinParamsForceLanguageTelugu         BrandGetByIsinParamsForceLanguage = "telugu"
	BrandGetByIsinParamsForceLanguageThai           BrandGetByIsinParamsForceLanguage = "thai"
	BrandGetByIsinParamsForceLanguageTibetan        BrandGetByIsinParamsForceLanguage = "tibetan"
	BrandGetByIsinParamsForceLanguageTigrinya       BrandGetByIsinParamsForceLanguage = "tigrinya"
	BrandGetByIsinParamsForceLanguageTongan         BrandGetByIsinParamsForceLanguage = "tongan"
	BrandGetByIsinParamsForceLanguageTswana         BrandGetByIsinParamsForceLanguage = "tswana"
	BrandGetByIsinParamsForceLanguageTurkish        BrandGetByIsinParamsForceLanguage = "turkish"
	BrandGetByIsinParamsForceLanguageTurkmen        BrandGetByIsinParamsForceLanguage = "turkmen"
	BrandGetByIsinParamsForceLanguageUkrainian      BrandGetByIsinParamsForceLanguage = "ukrainian"
	BrandGetByIsinParamsForceLanguageUrdu           BrandGetByIsinParamsForceLanguage = "urdu"
	BrandGetByIsinParamsForceLanguageUyghur         BrandGetByIsinParamsForceLanguage = "uyghur"
	BrandGetByIsinParamsForceLanguageUzbek          BrandGetByIsinParamsForceLanguage = "uzbek"
	BrandGetByIsinParamsForceLanguageVietnamese     BrandGetByIsinParamsForceLanguage = "vietnamese"
	BrandGetByIsinParamsForceLanguageWelsh          BrandGetByIsinParamsForceLanguage = "welsh"
	BrandGetByIsinParamsForceLanguageWolof          BrandGetByIsinParamsForceLanguage = "wolof"
	BrandGetByIsinParamsForceLanguageXhosa          BrandGetByIsinParamsForceLanguage = "xhosa"
	BrandGetByIsinParamsForceLanguageYiddish        BrandGetByIsinParamsForceLanguage = "yiddish"
	BrandGetByIsinParamsForceLanguageYoruba         BrandGetByIsinParamsForceLanguage = "yoruba"
	BrandGetByIsinParamsForceLanguageZulu           BrandGetByIsinParamsForceLanguage = "zulu"
)

type BrandGetByNameParams struct {
	// Company name to retrieve brand data for (e.g., 'Apple Inc', 'Microsoft
	// Corporation'). Must be 3-30 characters.
	Name string `query:"name" api:"required" json:"-"`
	// Maximum age in milliseconds for cached brand data before the API performs a hard
	// refresh. Defaults to 3 months (7776000000 ms). Values below 1 day (86400000 ms)
	// are clamped to 1 day; values above 1 year (31536000000 ms) are clamped to 1
	// year.
	MaxAgeMs param.Opt[int64] `query:"maxAgeMs,omitzero" json:"-"`
	// Optional parameter to optimize the API call for maximum speed. When set to true,
	// the API will skip time-consuming operations for faster response at the cost of
	// less comprehensive data.
	MaxSpeed param.Opt[bool] `query:"maxSpeed,omitzero" json:"-"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `query:"timeoutMS,omitzero" json:"-"`
	// Optional country code hint (GL parameter) to specify the country for the company
	// name.
	//
	// Any of "ad", "ae", "af", "ag", "ai", "al", "am", "an", "ao", "aq", "ar", "as",
	// "at", "au", "aw", "az", "ba", "bb", "bd", "be", "bf", "bg", "bh", "bi", "bj",
	// "bm", "bn", "bo", "br", "bs", "bt", "bv", "bw", "by", "bz", "ca", "cc", "cd",
	// "cf", "cg", "ch", "ci", "ck", "cl", "cm", "cn", "co", "cr", "cu", "cv", "cx",
	// "cy", "cz", "de", "dj", "dk", "dm", "do", "dz", "ec", "ee", "eg", "eh", "er",
	// "es", "et", "fi", "fj", "fk", "fm", "fo", "fr", "ga", "gb", "gd", "ge", "gf",
	// "gh", "gi", "gl", "gm", "gn", "gp", "gq", "gr", "gs", "gt", "gu", "gw", "gy",
	// "hk", "hm", "hn", "hr", "ht", "hu", "id", "ie", "il", "in", "io", "iq", "ir",
	// "is", "it", "jm", "jo", "jp", "ke", "kg", "kh", "ki", "km", "kn", "kp", "kr",
	// "kw", "ky", "kz", "la", "lb", "lc", "li", "lk", "lr", "ls", "lt", "lu", "lv",
	// "ly", "ma", "mc", "md", "mg", "mh", "mk", "ml", "mm", "mn", "mo", "mp", "mq",
	// "mr", "ms", "mt", "mu", "mv", "mw", "mx", "my", "mz", "na", "nc", "ne", "nf",
	// "ng", "ni", "nl", "no", "np", "nr", "nu", "nz", "om", "pa", "pe", "pf", "pg",
	// "ph", "pk", "pl", "pm", "pn", "pr", "ps", "pt", "pw", "py", "qa", "re", "ro",
	// "rs", "ru", "rw", "sa", "sb", "sc", "sd", "se", "sg", "sh", "si", "sj", "sk",
	// "sl", "sm", "sn", "so", "sr", "st", "sv", "sy", "sz", "tc", "td", "tf", "tg",
	// "th", "tj", "tk", "tl", "tm", "tn", "to", "tr", "tt", "tv", "tw", "tz", "ua",
	// "ug", "um", "us", "uy", "uz", "va", "vc", "ve", "vg", "vi", "vn", "vu", "wf",
	// "ws", "ye", "yt", "za", "zm", "zw".
	CountryGl BrandGetByNameParamsCountryGl `query:"country_gl,omitzero" json:"-"`
	// Optional parameter to force the language of the retrieved brand data.
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
	ForceLanguage BrandGetByNameParamsForceLanguage `query:"force_language,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BrandGetByNameParams]'s query parameters as `url.Values`.
func (r BrandGetByNameParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Optional country code hint (GL parameter) to specify the country for the company
// name.
type BrandGetByNameParamsCountryGl string

const (
	BrandGetByNameParamsCountryGlAd BrandGetByNameParamsCountryGl = "ad"
	BrandGetByNameParamsCountryGlAe BrandGetByNameParamsCountryGl = "ae"
	BrandGetByNameParamsCountryGlAf BrandGetByNameParamsCountryGl = "af"
	BrandGetByNameParamsCountryGlAg BrandGetByNameParamsCountryGl = "ag"
	BrandGetByNameParamsCountryGlAI BrandGetByNameParamsCountryGl = "ai"
	BrandGetByNameParamsCountryGlAl BrandGetByNameParamsCountryGl = "al"
	BrandGetByNameParamsCountryGlAm BrandGetByNameParamsCountryGl = "am"
	BrandGetByNameParamsCountryGlAn BrandGetByNameParamsCountryGl = "an"
	BrandGetByNameParamsCountryGlAo BrandGetByNameParamsCountryGl = "ao"
	BrandGetByNameParamsCountryGlAq BrandGetByNameParamsCountryGl = "aq"
	BrandGetByNameParamsCountryGlAr BrandGetByNameParamsCountryGl = "ar"
	BrandGetByNameParamsCountryGlAs BrandGetByNameParamsCountryGl = "as"
	BrandGetByNameParamsCountryGlAt BrandGetByNameParamsCountryGl = "at"
	BrandGetByNameParamsCountryGlAu BrandGetByNameParamsCountryGl = "au"
	BrandGetByNameParamsCountryGlAw BrandGetByNameParamsCountryGl = "aw"
	BrandGetByNameParamsCountryGlAz BrandGetByNameParamsCountryGl = "az"
	BrandGetByNameParamsCountryGlBa BrandGetByNameParamsCountryGl = "ba"
	BrandGetByNameParamsCountryGlBb BrandGetByNameParamsCountryGl = "bb"
	BrandGetByNameParamsCountryGlBd BrandGetByNameParamsCountryGl = "bd"
	BrandGetByNameParamsCountryGlBe BrandGetByNameParamsCountryGl = "be"
	BrandGetByNameParamsCountryGlBf BrandGetByNameParamsCountryGl = "bf"
	BrandGetByNameParamsCountryGlBg BrandGetByNameParamsCountryGl = "bg"
	BrandGetByNameParamsCountryGlBh BrandGetByNameParamsCountryGl = "bh"
	BrandGetByNameParamsCountryGlBi BrandGetByNameParamsCountryGl = "bi"
	BrandGetByNameParamsCountryGlBj BrandGetByNameParamsCountryGl = "bj"
	BrandGetByNameParamsCountryGlBm BrandGetByNameParamsCountryGl = "bm"
	BrandGetByNameParamsCountryGlBn BrandGetByNameParamsCountryGl = "bn"
	BrandGetByNameParamsCountryGlBo BrandGetByNameParamsCountryGl = "bo"
	BrandGetByNameParamsCountryGlBr BrandGetByNameParamsCountryGl = "br"
	BrandGetByNameParamsCountryGlBs BrandGetByNameParamsCountryGl = "bs"
	BrandGetByNameParamsCountryGlBt BrandGetByNameParamsCountryGl = "bt"
	BrandGetByNameParamsCountryGlBv BrandGetByNameParamsCountryGl = "bv"
	BrandGetByNameParamsCountryGlBw BrandGetByNameParamsCountryGl = "bw"
	BrandGetByNameParamsCountryGlBy BrandGetByNameParamsCountryGl = "by"
	BrandGetByNameParamsCountryGlBz BrandGetByNameParamsCountryGl = "bz"
	BrandGetByNameParamsCountryGlCa BrandGetByNameParamsCountryGl = "ca"
	BrandGetByNameParamsCountryGlCc BrandGetByNameParamsCountryGl = "cc"
	BrandGetByNameParamsCountryGlCd BrandGetByNameParamsCountryGl = "cd"
	BrandGetByNameParamsCountryGlCf BrandGetByNameParamsCountryGl = "cf"
	BrandGetByNameParamsCountryGlCg BrandGetByNameParamsCountryGl = "cg"
	BrandGetByNameParamsCountryGlCh BrandGetByNameParamsCountryGl = "ch"
	BrandGetByNameParamsCountryGlCi BrandGetByNameParamsCountryGl = "ci"
	BrandGetByNameParamsCountryGlCk BrandGetByNameParamsCountryGl = "ck"
	BrandGetByNameParamsCountryGlCl BrandGetByNameParamsCountryGl = "cl"
	BrandGetByNameParamsCountryGlCm BrandGetByNameParamsCountryGl = "cm"
	BrandGetByNameParamsCountryGlCn BrandGetByNameParamsCountryGl = "cn"
	BrandGetByNameParamsCountryGlCo BrandGetByNameParamsCountryGl = "co"
	BrandGetByNameParamsCountryGlCr BrandGetByNameParamsCountryGl = "cr"
	BrandGetByNameParamsCountryGlCu BrandGetByNameParamsCountryGl = "cu"
	BrandGetByNameParamsCountryGlCv BrandGetByNameParamsCountryGl = "cv"
	BrandGetByNameParamsCountryGlCx BrandGetByNameParamsCountryGl = "cx"
	BrandGetByNameParamsCountryGlCy BrandGetByNameParamsCountryGl = "cy"
	BrandGetByNameParamsCountryGlCz BrandGetByNameParamsCountryGl = "cz"
	BrandGetByNameParamsCountryGlDe BrandGetByNameParamsCountryGl = "de"
	BrandGetByNameParamsCountryGlDj BrandGetByNameParamsCountryGl = "dj"
	BrandGetByNameParamsCountryGlDk BrandGetByNameParamsCountryGl = "dk"
	BrandGetByNameParamsCountryGlDm BrandGetByNameParamsCountryGl = "dm"
	BrandGetByNameParamsCountryGlDo BrandGetByNameParamsCountryGl = "do"
	BrandGetByNameParamsCountryGlDz BrandGetByNameParamsCountryGl = "dz"
	BrandGetByNameParamsCountryGlEc BrandGetByNameParamsCountryGl = "ec"
	BrandGetByNameParamsCountryGlEe BrandGetByNameParamsCountryGl = "ee"
	BrandGetByNameParamsCountryGlEg BrandGetByNameParamsCountryGl = "eg"
	BrandGetByNameParamsCountryGlEh BrandGetByNameParamsCountryGl = "eh"
	BrandGetByNameParamsCountryGlEr BrandGetByNameParamsCountryGl = "er"
	BrandGetByNameParamsCountryGlEs BrandGetByNameParamsCountryGl = "es"
	BrandGetByNameParamsCountryGlEt BrandGetByNameParamsCountryGl = "et"
	BrandGetByNameParamsCountryGlFi BrandGetByNameParamsCountryGl = "fi"
	BrandGetByNameParamsCountryGlFj BrandGetByNameParamsCountryGl = "fj"
	BrandGetByNameParamsCountryGlFk BrandGetByNameParamsCountryGl = "fk"
	BrandGetByNameParamsCountryGlFm BrandGetByNameParamsCountryGl = "fm"
	BrandGetByNameParamsCountryGlFo BrandGetByNameParamsCountryGl = "fo"
	BrandGetByNameParamsCountryGlFr BrandGetByNameParamsCountryGl = "fr"
	BrandGetByNameParamsCountryGlGa BrandGetByNameParamsCountryGl = "ga"
	BrandGetByNameParamsCountryGlGB BrandGetByNameParamsCountryGl = "gb"
	BrandGetByNameParamsCountryGlGd BrandGetByNameParamsCountryGl = "gd"
	BrandGetByNameParamsCountryGlGe BrandGetByNameParamsCountryGl = "ge"
	BrandGetByNameParamsCountryGlGf BrandGetByNameParamsCountryGl = "gf"
	BrandGetByNameParamsCountryGlGh BrandGetByNameParamsCountryGl = "gh"
	BrandGetByNameParamsCountryGlGi BrandGetByNameParamsCountryGl = "gi"
	BrandGetByNameParamsCountryGlGl BrandGetByNameParamsCountryGl = "gl"
	BrandGetByNameParamsCountryGlGm BrandGetByNameParamsCountryGl = "gm"
	BrandGetByNameParamsCountryGlGn BrandGetByNameParamsCountryGl = "gn"
	BrandGetByNameParamsCountryGlGp BrandGetByNameParamsCountryGl = "gp"
	BrandGetByNameParamsCountryGlGq BrandGetByNameParamsCountryGl = "gq"
	BrandGetByNameParamsCountryGlGr BrandGetByNameParamsCountryGl = "gr"
	BrandGetByNameParamsCountryGlGs BrandGetByNameParamsCountryGl = "gs"
	BrandGetByNameParamsCountryGlGt BrandGetByNameParamsCountryGl = "gt"
	BrandGetByNameParamsCountryGlGu BrandGetByNameParamsCountryGl = "gu"
	BrandGetByNameParamsCountryGlGw BrandGetByNameParamsCountryGl = "gw"
	BrandGetByNameParamsCountryGlGy BrandGetByNameParamsCountryGl = "gy"
	BrandGetByNameParamsCountryGlHk BrandGetByNameParamsCountryGl = "hk"
	BrandGetByNameParamsCountryGlHm BrandGetByNameParamsCountryGl = "hm"
	BrandGetByNameParamsCountryGlHn BrandGetByNameParamsCountryGl = "hn"
	BrandGetByNameParamsCountryGlHr BrandGetByNameParamsCountryGl = "hr"
	BrandGetByNameParamsCountryGlHt BrandGetByNameParamsCountryGl = "ht"
	BrandGetByNameParamsCountryGlHu BrandGetByNameParamsCountryGl = "hu"
	BrandGetByNameParamsCountryGlID BrandGetByNameParamsCountryGl = "id"
	BrandGetByNameParamsCountryGlIe BrandGetByNameParamsCountryGl = "ie"
	BrandGetByNameParamsCountryGlIl BrandGetByNameParamsCountryGl = "il"
	BrandGetByNameParamsCountryGlIn BrandGetByNameParamsCountryGl = "in"
	BrandGetByNameParamsCountryGlIo BrandGetByNameParamsCountryGl = "io"
	BrandGetByNameParamsCountryGlIq BrandGetByNameParamsCountryGl = "iq"
	BrandGetByNameParamsCountryGlIr BrandGetByNameParamsCountryGl = "ir"
	BrandGetByNameParamsCountryGlIs BrandGetByNameParamsCountryGl = "is"
	BrandGetByNameParamsCountryGlIt BrandGetByNameParamsCountryGl = "it"
	BrandGetByNameParamsCountryGlJm BrandGetByNameParamsCountryGl = "jm"
	BrandGetByNameParamsCountryGlJo BrandGetByNameParamsCountryGl = "jo"
	BrandGetByNameParamsCountryGlJp BrandGetByNameParamsCountryGl = "jp"
	BrandGetByNameParamsCountryGlKe BrandGetByNameParamsCountryGl = "ke"
	BrandGetByNameParamsCountryGlKg BrandGetByNameParamsCountryGl = "kg"
	BrandGetByNameParamsCountryGlKh BrandGetByNameParamsCountryGl = "kh"
	BrandGetByNameParamsCountryGlKi BrandGetByNameParamsCountryGl = "ki"
	BrandGetByNameParamsCountryGlKm BrandGetByNameParamsCountryGl = "km"
	BrandGetByNameParamsCountryGlKn BrandGetByNameParamsCountryGl = "kn"
	BrandGetByNameParamsCountryGlKp BrandGetByNameParamsCountryGl = "kp"
	BrandGetByNameParamsCountryGlKr BrandGetByNameParamsCountryGl = "kr"
	BrandGetByNameParamsCountryGlKw BrandGetByNameParamsCountryGl = "kw"
	BrandGetByNameParamsCountryGlKy BrandGetByNameParamsCountryGl = "ky"
	BrandGetByNameParamsCountryGlKz BrandGetByNameParamsCountryGl = "kz"
	BrandGetByNameParamsCountryGlLa BrandGetByNameParamsCountryGl = "la"
	BrandGetByNameParamsCountryGlLb BrandGetByNameParamsCountryGl = "lb"
	BrandGetByNameParamsCountryGlLc BrandGetByNameParamsCountryGl = "lc"
	BrandGetByNameParamsCountryGlLi BrandGetByNameParamsCountryGl = "li"
	BrandGetByNameParamsCountryGlLk BrandGetByNameParamsCountryGl = "lk"
	BrandGetByNameParamsCountryGlLr BrandGetByNameParamsCountryGl = "lr"
	BrandGetByNameParamsCountryGlLs BrandGetByNameParamsCountryGl = "ls"
	BrandGetByNameParamsCountryGlLt BrandGetByNameParamsCountryGl = "lt"
	BrandGetByNameParamsCountryGlLu BrandGetByNameParamsCountryGl = "lu"
	BrandGetByNameParamsCountryGlLv BrandGetByNameParamsCountryGl = "lv"
	BrandGetByNameParamsCountryGlLy BrandGetByNameParamsCountryGl = "ly"
	BrandGetByNameParamsCountryGlMa BrandGetByNameParamsCountryGl = "ma"
	BrandGetByNameParamsCountryGlMc BrandGetByNameParamsCountryGl = "mc"
	BrandGetByNameParamsCountryGlMd BrandGetByNameParamsCountryGl = "md"
	BrandGetByNameParamsCountryGlMg BrandGetByNameParamsCountryGl = "mg"
	BrandGetByNameParamsCountryGlMh BrandGetByNameParamsCountryGl = "mh"
	BrandGetByNameParamsCountryGlMk BrandGetByNameParamsCountryGl = "mk"
	BrandGetByNameParamsCountryGlMl BrandGetByNameParamsCountryGl = "ml"
	BrandGetByNameParamsCountryGlMm BrandGetByNameParamsCountryGl = "mm"
	BrandGetByNameParamsCountryGlMn BrandGetByNameParamsCountryGl = "mn"
	BrandGetByNameParamsCountryGlMo BrandGetByNameParamsCountryGl = "mo"
	BrandGetByNameParamsCountryGlMp BrandGetByNameParamsCountryGl = "mp"
	BrandGetByNameParamsCountryGlMq BrandGetByNameParamsCountryGl = "mq"
	BrandGetByNameParamsCountryGlMr BrandGetByNameParamsCountryGl = "mr"
	BrandGetByNameParamsCountryGlMs BrandGetByNameParamsCountryGl = "ms"
	BrandGetByNameParamsCountryGlMt BrandGetByNameParamsCountryGl = "mt"
	BrandGetByNameParamsCountryGlMu BrandGetByNameParamsCountryGl = "mu"
	BrandGetByNameParamsCountryGlMv BrandGetByNameParamsCountryGl = "mv"
	BrandGetByNameParamsCountryGlMw BrandGetByNameParamsCountryGl = "mw"
	BrandGetByNameParamsCountryGlMx BrandGetByNameParamsCountryGl = "mx"
	BrandGetByNameParamsCountryGlMy BrandGetByNameParamsCountryGl = "my"
	BrandGetByNameParamsCountryGlMz BrandGetByNameParamsCountryGl = "mz"
	BrandGetByNameParamsCountryGlNa BrandGetByNameParamsCountryGl = "na"
	BrandGetByNameParamsCountryGlNc BrandGetByNameParamsCountryGl = "nc"
	BrandGetByNameParamsCountryGlNe BrandGetByNameParamsCountryGl = "ne"
	BrandGetByNameParamsCountryGlNf BrandGetByNameParamsCountryGl = "nf"
	BrandGetByNameParamsCountryGlNg BrandGetByNameParamsCountryGl = "ng"
	BrandGetByNameParamsCountryGlNi BrandGetByNameParamsCountryGl = "ni"
	BrandGetByNameParamsCountryGlNl BrandGetByNameParamsCountryGl = "nl"
	BrandGetByNameParamsCountryGlNo BrandGetByNameParamsCountryGl = "no"
	BrandGetByNameParamsCountryGlNp BrandGetByNameParamsCountryGl = "np"
	BrandGetByNameParamsCountryGlNr BrandGetByNameParamsCountryGl = "nr"
	BrandGetByNameParamsCountryGlNu BrandGetByNameParamsCountryGl = "nu"
	BrandGetByNameParamsCountryGlNz BrandGetByNameParamsCountryGl = "nz"
	BrandGetByNameParamsCountryGlOm BrandGetByNameParamsCountryGl = "om"
	BrandGetByNameParamsCountryGlPa BrandGetByNameParamsCountryGl = "pa"
	BrandGetByNameParamsCountryGlPe BrandGetByNameParamsCountryGl = "pe"
	BrandGetByNameParamsCountryGlPf BrandGetByNameParamsCountryGl = "pf"
	BrandGetByNameParamsCountryGlPg BrandGetByNameParamsCountryGl = "pg"
	BrandGetByNameParamsCountryGlPh BrandGetByNameParamsCountryGl = "ph"
	BrandGetByNameParamsCountryGlPk BrandGetByNameParamsCountryGl = "pk"
	BrandGetByNameParamsCountryGlPl BrandGetByNameParamsCountryGl = "pl"
	BrandGetByNameParamsCountryGlPm BrandGetByNameParamsCountryGl = "pm"
	BrandGetByNameParamsCountryGlPn BrandGetByNameParamsCountryGl = "pn"
	BrandGetByNameParamsCountryGlPr BrandGetByNameParamsCountryGl = "pr"
	BrandGetByNameParamsCountryGlPs BrandGetByNameParamsCountryGl = "ps"
	BrandGetByNameParamsCountryGlPt BrandGetByNameParamsCountryGl = "pt"
	BrandGetByNameParamsCountryGlPw BrandGetByNameParamsCountryGl = "pw"
	BrandGetByNameParamsCountryGlPy BrandGetByNameParamsCountryGl = "py"
	BrandGetByNameParamsCountryGlQa BrandGetByNameParamsCountryGl = "qa"
	BrandGetByNameParamsCountryGlRe BrandGetByNameParamsCountryGl = "re"
	BrandGetByNameParamsCountryGlRo BrandGetByNameParamsCountryGl = "ro"
	BrandGetByNameParamsCountryGlRs BrandGetByNameParamsCountryGl = "rs"
	BrandGetByNameParamsCountryGlRu BrandGetByNameParamsCountryGl = "ru"
	BrandGetByNameParamsCountryGlRw BrandGetByNameParamsCountryGl = "rw"
	BrandGetByNameParamsCountryGlSa BrandGetByNameParamsCountryGl = "sa"
	BrandGetByNameParamsCountryGlSb BrandGetByNameParamsCountryGl = "sb"
	BrandGetByNameParamsCountryGlSc BrandGetByNameParamsCountryGl = "sc"
	BrandGetByNameParamsCountryGlSd BrandGetByNameParamsCountryGl = "sd"
	BrandGetByNameParamsCountryGlSe BrandGetByNameParamsCountryGl = "se"
	BrandGetByNameParamsCountryGlSg BrandGetByNameParamsCountryGl = "sg"
	BrandGetByNameParamsCountryGlSh BrandGetByNameParamsCountryGl = "sh"
	BrandGetByNameParamsCountryGlSi BrandGetByNameParamsCountryGl = "si"
	BrandGetByNameParamsCountryGlSj BrandGetByNameParamsCountryGl = "sj"
	BrandGetByNameParamsCountryGlSk BrandGetByNameParamsCountryGl = "sk"
	BrandGetByNameParamsCountryGlSl BrandGetByNameParamsCountryGl = "sl"
	BrandGetByNameParamsCountryGlSm BrandGetByNameParamsCountryGl = "sm"
	BrandGetByNameParamsCountryGlSn BrandGetByNameParamsCountryGl = "sn"
	BrandGetByNameParamsCountryGlSo BrandGetByNameParamsCountryGl = "so"
	BrandGetByNameParamsCountryGlSr BrandGetByNameParamsCountryGl = "sr"
	BrandGetByNameParamsCountryGlSt BrandGetByNameParamsCountryGl = "st"
	BrandGetByNameParamsCountryGlSv BrandGetByNameParamsCountryGl = "sv"
	BrandGetByNameParamsCountryGlSy BrandGetByNameParamsCountryGl = "sy"
	BrandGetByNameParamsCountryGlSz BrandGetByNameParamsCountryGl = "sz"
	BrandGetByNameParamsCountryGlTc BrandGetByNameParamsCountryGl = "tc"
	BrandGetByNameParamsCountryGlTd BrandGetByNameParamsCountryGl = "td"
	BrandGetByNameParamsCountryGlTf BrandGetByNameParamsCountryGl = "tf"
	BrandGetByNameParamsCountryGlTg BrandGetByNameParamsCountryGl = "tg"
	BrandGetByNameParamsCountryGlTh BrandGetByNameParamsCountryGl = "th"
	BrandGetByNameParamsCountryGlTj BrandGetByNameParamsCountryGl = "tj"
	BrandGetByNameParamsCountryGlTk BrandGetByNameParamsCountryGl = "tk"
	BrandGetByNameParamsCountryGlTl BrandGetByNameParamsCountryGl = "tl"
	BrandGetByNameParamsCountryGlTm BrandGetByNameParamsCountryGl = "tm"
	BrandGetByNameParamsCountryGlTn BrandGetByNameParamsCountryGl = "tn"
	BrandGetByNameParamsCountryGlTo BrandGetByNameParamsCountryGl = "to"
	BrandGetByNameParamsCountryGlTr BrandGetByNameParamsCountryGl = "tr"
	BrandGetByNameParamsCountryGlTt BrandGetByNameParamsCountryGl = "tt"
	BrandGetByNameParamsCountryGlTv BrandGetByNameParamsCountryGl = "tv"
	BrandGetByNameParamsCountryGlTw BrandGetByNameParamsCountryGl = "tw"
	BrandGetByNameParamsCountryGlTz BrandGetByNameParamsCountryGl = "tz"
	BrandGetByNameParamsCountryGlUa BrandGetByNameParamsCountryGl = "ua"
	BrandGetByNameParamsCountryGlUg BrandGetByNameParamsCountryGl = "ug"
	BrandGetByNameParamsCountryGlUm BrandGetByNameParamsCountryGl = "um"
	BrandGetByNameParamsCountryGlUs BrandGetByNameParamsCountryGl = "us"
	BrandGetByNameParamsCountryGlUy BrandGetByNameParamsCountryGl = "uy"
	BrandGetByNameParamsCountryGlUz BrandGetByNameParamsCountryGl = "uz"
	BrandGetByNameParamsCountryGlVa BrandGetByNameParamsCountryGl = "va"
	BrandGetByNameParamsCountryGlVc BrandGetByNameParamsCountryGl = "vc"
	BrandGetByNameParamsCountryGlVe BrandGetByNameParamsCountryGl = "ve"
	BrandGetByNameParamsCountryGlVg BrandGetByNameParamsCountryGl = "vg"
	BrandGetByNameParamsCountryGlVi BrandGetByNameParamsCountryGl = "vi"
	BrandGetByNameParamsCountryGlVn BrandGetByNameParamsCountryGl = "vn"
	BrandGetByNameParamsCountryGlVu BrandGetByNameParamsCountryGl = "vu"
	BrandGetByNameParamsCountryGlWf BrandGetByNameParamsCountryGl = "wf"
	BrandGetByNameParamsCountryGlWs BrandGetByNameParamsCountryGl = "ws"
	BrandGetByNameParamsCountryGlYe BrandGetByNameParamsCountryGl = "ye"
	BrandGetByNameParamsCountryGlYt BrandGetByNameParamsCountryGl = "yt"
	BrandGetByNameParamsCountryGlZa BrandGetByNameParamsCountryGl = "za"
	BrandGetByNameParamsCountryGlZm BrandGetByNameParamsCountryGl = "zm"
	BrandGetByNameParamsCountryGlZw BrandGetByNameParamsCountryGl = "zw"
)

// Optional parameter to force the language of the retrieved brand data.
type BrandGetByNameParamsForceLanguage string

const (
	BrandGetByNameParamsForceLanguageAfrikaans      BrandGetByNameParamsForceLanguage = "afrikaans"
	BrandGetByNameParamsForceLanguageAlbanian       BrandGetByNameParamsForceLanguage = "albanian"
	BrandGetByNameParamsForceLanguageAmharic        BrandGetByNameParamsForceLanguage = "amharic"
	BrandGetByNameParamsForceLanguageArabic         BrandGetByNameParamsForceLanguage = "arabic"
	BrandGetByNameParamsForceLanguageArmenian       BrandGetByNameParamsForceLanguage = "armenian"
	BrandGetByNameParamsForceLanguageAssamese       BrandGetByNameParamsForceLanguage = "assamese"
	BrandGetByNameParamsForceLanguageAymara         BrandGetByNameParamsForceLanguage = "aymara"
	BrandGetByNameParamsForceLanguageAzeri          BrandGetByNameParamsForceLanguage = "azeri"
	BrandGetByNameParamsForceLanguageBasque         BrandGetByNameParamsForceLanguage = "basque"
	BrandGetByNameParamsForceLanguageBelarusian     BrandGetByNameParamsForceLanguage = "belarusian"
	BrandGetByNameParamsForceLanguageBengali        BrandGetByNameParamsForceLanguage = "bengali"
	BrandGetByNameParamsForceLanguageBosnian        BrandGetByNameParamsForceLanguage = "bosnian"
	BrandGetByNameParamsForceLanguageBulgarian      BrandGetByNameParamsForceLanguage = "bulgarian"
	BrandGetByNameParamsForceLanguageBurmese        BrandGetByNameParamsForceLanguage = "burmese"
	BrandGetByNameParamsForceLanguageCantonese      BrandGetByNameParamsForceLanguage = "cantonese"
	BrandGetByNameParamsForceLanguageCatalan        BrandGetByNameParamsForceLanguage = "catalan"
	BrandGetByNameParamsForceLanguageCebuano        BrandGetByNameParamsForceLanguage = "cebuano"
	BrandGetByNameParamsForceLanguageChinese        BrandGetByNameParamsForceLanguage = "chinese"
	BrandGetByNameParamsForceLanguageCorsican       BrandGetByNameParamsForceLanguage = "corsican"
	BrandGetByNameParamsForceLanguageCroatian       BrandGetByNameParamsForceLanguage = "croatian"
	BrandGetByNameParamsForceLanguageCzech          BrandGetByNameParamsForceLanguage = "czech"
	BrandGetByNameParamsForceLanguageDanish         BrandGetByNameParamsForceLanguage = "danish"
	BrandGetByNameParamsForceLanguageDutch          BrandGetByNameParamsForceLanguage = "dutch"
	BrandGetByNameParamsForceLanguageEnglish        BrandGetByNameParamsForceLanguage = "english"
	BrandGetByNameParamsForceLanguageEsperanto      BrandGetByNameParamsForceLanguage = "esperanto"
	BrandGetByNameParamsForceLanguageEstonian       BrandGetByNameParamsForceLanguage = "estonian"
	BrandGetByNameParamsForceLanguageFarsi          BrandGetByNameParamsForceLanguage = "farsi"
	BrandGetByNameParamsForceLanguageFijian         BrandGetByNameParamsForceLanguage = "fijian"
	BrandGetByNameParamsForceLanguageFinnish        BrandGetByNameParamsForceLanguage = "finnish"
	BrandGetByNameParamsForceLanguageFrench         BrandGetByNameParamsForceLanguage = "french"
	BrandGetByNameParamsForceLanguageGalician       BrandGetByNameParamsForceLanguage = "galician"
	BrandGetByNameParamsForceLanguageGeorgian       BrandGetByNameParamsForceLanguage = "georgian"
	BrandGetByNameParamsForceLanguageGerman         BrandGetByNameParamsForceLanguage = "german"
	BrandGetByNameParamsForceLanguageGreek          BrandGetByNameParamsForceLanguage = "greek"
	BrandGetByNameParamsForceLanguageGuarani        BrandGetByNameParamsForceLanguage = "guarani"
	BrandGetByNameParamsForceLanguageGujarati       BrandGetByNameParamsForceLanguage = "gujarati"
	BrandGetByNameParamsForceLanguageHaitianCreole  BrandGetByNameParamsForceLanguage = "haitian-creole"
	BrandGetByNameParamsForceLanguageHausa          BrandGetByNameParamsForceLanguage = "hausa"
	BrandGetByNameParamsForceLanguageHawaiian       BrandGetByNameParamsForceLanguage = "hawaiian"
	BrandGetByNameParamsForceLanguageHebrew         BrandGetByNameParamsForceLanguage = "hebrew"
	BrandGetByNameParamsForceLanguageHindi          BrandGetByNameParamsForceLanguage = "hindi"
	BrandGetByNameParamsForceLanguageHmong          BrandGetByNameParamsForceLanguage = "hmong"
	BrandGetByNameParamsForceLanguageHungarian      BrandGetByNameParamsForceLanguage = "hungarian"
	BrandGetByNameParamsForceLanguageIcelandic      BrandGetByNameParamsForceLanguage = "icelandic"
	BrandGetByNameParamsForceLanguageIgbo           BrandGetByNameParamsForceLanguage = "igbo"
	BrandGetByNameParamsForceLanguageIndonesian     BrandGetByNameParamsForceLanguage = "indonesian"
	BrandGetByNameParamsForceLanguageIrish          BrandGetByNameParamsForceLanguage = "irish"
	BrandGetByNameParamsForceLanguageItalian        BrandGetByNameParamsForceLanguage = "italian"
	BrandGetByNameParamsForceLanguageJapanese       BrandGetByNameParamsForceLanguage = "japanese"
	BrandGetByNameParamsForceLanguageJavanese       BrandGetByNameParamsForceLanguage = "javanese"
	BrandGetByNameParamsForceLanguageKannada        BrandGetByNameParamsForceLanguage = "kannada"
	BrandGetByNameParamsForceLanguageKazakh         BrandGetByNameParamsForceLanguage = "kazakh"
	BrandGetByNameParamsForceLanguageKhmer          BrandGetByNameParamsForceLanguage = "khmer"
	BrandGetByNameParamsForceLanguageKinyarwanda    BrandGetByNameParamsForceLanguage = "kinyarwanda"
	BrandGetByNameParamsForceLanguageKorean         BrandGetByNameParamsForceLanguage = "korean"
	BrandGetByNameParamsForceLanguageKurdish        BrandGetByNameParamsForceLanguage = "kurdish"
	BrandGetByNameParamsForceLanguageKyrgyz         BrandGetByNameParamsForceLanguage = "kyrgyz"
	BrandGetByNameParamsForceLanguageLao            BrandGetByNameParamsForceLanguage = "lao"
	BrandGetByNameParamsForceLanguageLatin          BrandGetByNameParamsForceLanguage = "latin"
	BrandGetByNameParamsForceLanguageLatvian        BrandGetByNameParamsForceLanguage = "latvian"
	BrandGetByNameParamsForceLanguageLingala        BrandGetByNameParamsForceLanguage = "lingala"
	BrandGetByNameParamsForceLanguageLithuanian     BrandGetByNameParamsForceLanguage = "lithuanian"
	BrandGetByNameParamsForceLanguageLuxembourgish  BrandGetByNameParamsForceLanguage = "luxembourgish"
	BrandGetByNameParamsForceLanguageMacedonian     BrandGetByNameParamsForceLanguage = "macedonian"
	BrandGetByNameParamsForceLanguageMalagasy       BrandGetByNameParamsForceLanguage = "malagasy"
	BrandGetByNameParamsForceLanguageMalay          BrandGetByNameParamsForceLanguage = "malay"
	BrandGetByNameParamsForceLanguageMalayalam      BrandGetByNameParamsForceLanguage = "malayalam"
	BrandGetByNameParamsForceLanguageMaltese        BrandGetByNameParamsForceLanguage = "maltese"
	BrandGetByNameParamsForceLanguageMaori          BrandGetByNameParamsForceLanguage = "maori"
	BrandGetByNameParamsForceLanguageMarathi        BrandGetByNameParamsForceLanguage = "marathi"
	BrandGetByNameParamsForceLanguageMongolian      BrandGetByNameParamsForceLanguage = "mongolian"
	BrandGetByNameParamsForceLanguageNepali         BrandGetByNameParamsForceLanguage = "nepali"
	BrandGetByNameParamsForceLanguageNorwegian      BrandGetByNameParamsForceLanguage = "norwegian"
	BrandGetByNameParamsForceLanguageOdia           BrandGetByNameParamsForceLanguage = "odia"
	BrandGetByNameParamsForceLanguageOromo          BrandGetByNameParamsForceLanguage = "oromo"
	BrandGetByNameParamsForceLanguagePashto         BrandGetByNameParamsForceLanguage = "pashto"
	BrandGetByNameParamsForceLanguagePidgin         BrandGetByNameParamsForceLanguage = "pidgin"
	BrandGetByNameParamsForceLanguagePolish         BrandGetByNameParamsForceLanguage = "polish"
	BrandGetByNameParamsForceLanguagePortuguese     BrandGetByNameParamsForceLanguage = "portuguese"
	BrandGetByNameParamsForceLanguagePunjabi        BrandGetByNameParamsForceLanguage = "punjabi"
	BrandGetByNameParamsForceLanguageQuechua        BrandGetByNameParamsForceLanguage = "quechua"
	BrandGetByNameParamsForceLanguageRomanian       BrandGetByNameParamsForceLanguage = "romanian"
	BrandGetByNameParamsForceLanguageRussian        BrandGetByNameParamsForceLanguage = "russian"
	BrandGetByNameParamsForceLanguageSamoan         BrandGetByNameParamsForceLanguage = "samoan"
	BrandGetByNameParamsForceLanguageScottishGaelic BrandGetByNameParamsForceLanguage = "scottish-gaelic"
	BrandGetByNameParamsForceLanguageSerbian        BrandGetByNameParamsForceLanguage = "serbian"
	BrandGetByNameParamsForceLanguageSesotho        BrandGetByNameParamsForceLanguage = "sesotho"
	BrandGetByNameParamsForceLanguageShona          BrandGetByNameParamsForceLanguage = "shona"
	BrandGetByNameParamsForceLanguageSindhi         BrandGetByNameParamsForceLanguage = "sindhi"
	BrandGetByNameParamsForceLanguageSinhala        BrandGetByNameParamsForceLanguage = "sinhala"
	BrandGetByNameParamsForceLanguageSlovak         BrandGetByNameParamsForceLanguage = "slovak"
	BrandGetByNameParamsForceLanguageSlovene        BrandGetByNameParamsForceLanguage = "slovene"
	BrandGetByNameParamsForceLanguageSomali         BrandGetByNameParamsForceLanguage = "somali"
	BrandGetByNameParamsForceLanguageSpanish        BrandGetByNameParamsForceLanguage = "spanish"
	BrandGetByNameParamsForceLanguageSundanese      BrandGetByNameParamsForceLanguage = "sundanese"
	BrandGetByNameParamsForceLanguageSwahili        BrandGetByNameParamsForceLanguage = "swahili"
	BrandGetByNameParamsForceLanguageSwedish        BrandGetByNameParamsForceLanguage = "swedish"
	BrandGetByNameParamsForceLanguageTagalog        BrandGetByNameParamsForceLanguage = "tagalog"
	BrandGetByNameParamsForceLanguageTajik          BrandGetByNameParamsForceLanguage = "tajik"
	BrandGetByNameParamsForceLanguageTamil          BrandGetByNameParamsForceLanguage = "tamil"
	BrandGetByNameParamsForceLanguageTatar          BrandGetByNameParamsForceLanguage = "tatar"
	BrandGetByNameParamsForceLanguageTelugu         BrandGetByNameParamsForceLanguage = "telugu"
	BrandGetByNameParamsForceLanguageThai           BrandGetByNameParamsForceLanguage = "thai"
	BrandGetByNameParamsForceLanguageTibetan        BrandGetByNameParamsForceLanguage = "tibetan"
	BrandGetByNameParamsForceLanguageTigrinya       BrandGetByNameParamsForceLanguage = "tigrinya"
	BrandGetByNameParamsForceLanguageTongan         BrandGetByNameParamsForceLanguage = "tongan"
	BrandGetByNameParamsForceLanguageTswana         BrandGetByNameParamsForceLanguage = "tswana"
	BrandGetByNameParamsForceLanguageTurkish        BrandGetByNameParamsForceLanguage = "turkish"
	BrandGetByNameParamsForceLanguageTurkmen        BrandGetByNameParamsForceLanguage = "turkmen"
	BrandGetByNameParamsForceLanguageUkrainian      BrandGetByNameParamsForceLanguage = "ukrainian"
	BrandGetByNameParamsForceLanguageUrdu           BrandGetByNameParamsForceLanguage = "urdu"
	BrandGetByNameParamsForceLanguageUyghur         BrandGetByNameParamsForceLanguage = "uyghur"
	BrandGetByNameParamsForceLanguageUzbek          BrandGetByNameParamsForceLanguage = "uzbek"
	BrandGetByNameParamsForceLanguageVietnamese     BrandGetByNameParamsForceLanguage = "vietnamese"
	BrandGetByNameParamsForceLanguageWelsh          BrandGetByNameParamsForceLanguage = "welsh"
	BrandGetByNameParamsForceLanguageWolof          BrandGetByNameParamsForceLanguage = "wolof"
	BrandGetByNameParamsForceLanguageXhosa          BrandGetByNameParamsForceLanguage = "xhosa"
	BrandGetByNameParamsForceLanguageYiddish        BrandGetByNameParamsForceLanguage = "yiddish"
	BrandGetByNameParamsForceLanguageYoruba         BrandGetByNameParamsForceLanguage = "yoruba"
	BrandGetByNameParamsForceLanguageZulu           BrandGetByNameParamsForceLanguage = "zulu"
)

type BrandGetByTickerParams struct {
	// Stock ticker symbol to retrieve brand data for (e.g., 'AAPL', 'GOOGL', 'BRK.A').
	// Must be 1-15 characters, letters/numbers/dots only.
	Ticker string `query:"ticker" api:"required" json:"-"`
	// Maximum age in milliseconds for cached brand data before the API performs a hard
	// refresh. Defaults to 3 months (7776000000 ms). Values below 1 day (86400000 ms)
	// are clamped to 1 day; values above 1 year (31536000000 ms) are clamped to 1
	// year.
	MaxAgeMs param.Opt[int64] `query:"maxAgeMs,omitzero" json:"-"`
	// Optional parameter to optimize the API call for maximum speed. When set to true,
	// the API will skip time-consuming operations for faster response at the cost of
	// less comprehensive data.
	MaxSpeed param.Opt[bool] `query:"maxSpeed,omitzero" json:"-"`
	// Optional timeout in milliseconds for the request. If the request takes longer
	// than this value, it will be aborted with a 408 status code. Maximum allowed
	// value is 300000ms (5 minutes).
	TimeoutMs param.Opt[int64] `query:"timeoutMS,omitzero" json:"-"`
	// Optional parameter to force the language of the retrieved brand data.
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
	ForceLanguage BrandGetByTickerParamsForceLanguage `query:"force_language,omitzero" json:"-"`
	// Optional stock exchange for the ticker. Defaults to NASDAQ if not specified.
	//
	// Any of "AMEX", "AMS", "AQS", "ASX", "ATH", "BER", "BME", "BRU", "BSE", "BUD",
	// "BUE", "BVC", "CBOE", "CNQ", "CPH", "DFM", "DOH", "DUB", "DUS", "DXE", "EGX",
	// "FSX", "HAM", "HEL", "HKSE", "HOSE", "ICE", "IOB", "IST", "JKT", "JNB", "JPX",
	// "KLS", "KOE", "KSC", "KUW", "LIS", "LSE", "MCX", "MEX", "MIL", "MUN", "NASDAQ",
	// "NEO", "NSE", "NYSE", "NZE", "OSL", "OTC", "PAR", "PNK", "PRA", "RIS", "SAO",
	// "SAU", "SES", "SET", "SGO", "SHH", "SHZ", "SIX", "STO", "STU", "TAI", "TAL",
	// "TLV", "TSX", "TSXV", "TWO", "VIE", "WSE", "XETRA".
	TickerExchange BrandGetByTickerParamsTickerExchange `query:"ticker_exchange,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BrandGetByTickerParams]'s query parameters as `url.Values`.
func (r BrandGetByTickerParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Optional parameter to force the language of the retrieved brand data.
type BrandGetByTickerParamsForceLanguage string

const (
	BrandGetByTickerParamsForceLanguageAfrikaans      BrandGetByTickerParamsForceLanguage = "afrikaans"
	BrandGetByTickerParamsForceLanguageAlbanian       BrandGetByTickerParamsForceLanguage = "albanian"
	BrandGetByTickerParamsForceLanguageAmharic        BrandGetByTickerParamsForceLanguage = "amharic"
	BrandGetByTickerParamsForceLanguageArabic         BrandGetByTickerParamsForceLanguage = "arabic"
	BrandGetByTickerParamsForceLanguageArmenian       BrandGetByTickerParamsForceLanguage = "armenian"
	BrandGetByTickerParamsForceLanguageAssamese       BrandGetByTickerParamsForceLanguage = "assamese"
	BrandGetByTickerParamsForceLanguageAymara         BrandGetByTickerParamsForceLanguage = "aymara"
	BrandGetByTickerParamsForceLanguageAzeri          BrandGetByTickerParamsForceLanguage = "azeri"
	BrandGetByTickerParamsForceLanguageBasque         BrandGetByTickerParamsForceLanguage = "basque"
	BrandGetByTickerParamsForceLanguageBelarusian     BrandGetByTickerParamsForceLanguage = "belarusian"
	BrandGetByTickerParamsForceLanguageBengali        BrandGetByTickerParamsForceLanguage = "bengali"
	BrandGetByTickerParamsForceLanguageBosnian        BrandGetByTickerParamsForceLanguage = "bosnian"
	BrandGetByTickerParamsForceLanguageBulgarian      BrandGetByTickerParamsForceLanguage = "bulgarian"
	BrandGetByTickerParamsForceLanguageBurmese        BrandGetByTickerParamsForceLanguage = "burmese"
	BrandGetByTickerParamsForceLanguageCantonese      BrandGetByTickerParamsForceLanguage = "cantonese"
	BrandGetByTickerParamsForceLanguageCatalan        BrandGetByTickerParamsForceLanguage = "catalan"
	BrandGetByTickerParamsForceLanguageCebuano        BrandGetByTickerParamsForceLanguage = "cebuano"
	BrandGetByTickerParamsForceLanguageChinese        BrandGetByTickerParamsForceLanguage = "chinese"
	BrandGetByTickerParamsForceLanguageCorsican       BrandGetByTickerParamsForceLanguage = "corsican"
	BrandGetByTickerParamsForceLanguageCroatian       BrandGetByTickerParamsForceLanguage = "croatian"
	BrandGetByTickerParamsForceLanguageCzech          BrandGetByTickerParamsForceLanguage = "czech"
	BrandGetByTickerParamsForceLanguageDanish         BrandGetByTickerParamsForceLanguage = "danish"
	BrandGetByTickerParamsForceLanguageDutch          BrandGetByTickerParamsForceLanguage = "dutch"
	BrandGetByTickerParamsForceLanguageEnglish        BrandGetByTickerParamsForceLanguage = "english"
	BrandGetByTickerParamsForceLanguageEsperanto      BrandGetByTickerParamsForceLanguage = "esperanto"
	BrandGetByTickerParamsForceLanguageEstonian       BrandGetByTickerParamsForceLanguage = "estonian"
	BrandGetByTickerParamsForceLanguageFarsi          BrandGetByTickerParamsForceLanguage = "farsi"
	BrandGetByTickerParamsForceLanguageFijian         BrandGetByTickerParamsForceLanguage = "fijian"
	BrandGetByTickerParamsForceLanguageFinnish        BrandGetByTickerParamsForceLanguage = "finnish"
	BrandGetByTickerParamsForceLanguageFrench         BrandGetByTickerParamsForceLanguage = "french"
	BrandGetByTickerParamsForceLanguageGalician       BrandGetByTickerParamsForceLanguage = "galician"
	BrandGetByTickerParamsForceLanguageGeorgian       BrandGetByTickerParamsForceLanguage = "georgian"
	BrandGetByTickerParamsForceLanguageGerman         BrandGetByTickerParamsForceLanguage = "german"
	BrandGetByTickerParamsForceLanguageGreek          BrandGetByTickerParamsForceLanguage = "greek"
	BrandGetByTickerParamsForceLanguageGuarani        BrandGetByTickerParamsForceLanguage = "guarani"
	BrandGetByTickerParamsForceLanguageGujarati       BrandGetByTickerParamsForceLanguage = "gujarati"
	BrandGetByTickerParamsForceLanguageHaitianCreole  BrandGetByTickerParamsForceLanguage = "haitian-creole"
	BrandGetByTickerParamsForceLanguageHausa          BrandGetByTickerParamsForceLanguage = "hausa"
	BrandGetByTickerParamsForceLanguageHawaiian       BrandGetByTickerParamsForceLanguage = "hawaiian"
	BrandGetByTickerParamsForceLanguageHebrew         BrandGetByTickerParamsForceLanguage = "hebrew"
	BrandGetByTickerParamsForceLanguageHindi          BrandGetByTickerParamsForceLanguage = "hindi"
	BrandGetByTickerParamsForceLanguageHmong          BrandGetByTickerParamsForceLanguage = "hmong"
	BrandGetByTickerParamsForceLanguageHungarian      BrandGetByTickerParamsForceLanguage = "hungarian"
	BrandGetByTickerParamsForceLanguageIcelandic      BrandGetByTickerParamsForceLanguage = "icelandic"
	BrandGetByTickerParamsForceLanguageIgbo           BrandGetByTickerParamsForceLanguage = "igbo"
	BrandGetByTickerParamsForceLanguageIndonesian     BrandGetByTickerParamsForceLanguage = "indonesian"
	BrandGetByTickerParamsForceLanguageIrish          BrandGetByTickerParamsForceLanguage = "irish"
	BrandGetByTickerParamsForceLanguageItalian        BrandGetByTickerParamsForceLanguage = "italian"
	BrandGetByTickerParamsForceLanguageJapanese       BrandGetByTickerParamsForceLanguage = "japanese"
	BrandGetByTickerParamsForceLanguageJavanese       BrandGetByTickerParamsForceLanguage = "javanese"
	BrandGetByTickerParamsForceLanguageKannada        BrandGetByTickerParamsForceLanguage = "kannada"
	BrandGetByTickerParamsForceLanguageKazakh         BrandGetByTickerParamsForceLanguage = "kazakh"
	BrandGetByTickerParamsForceLanguageKhmer          BrandGetByTickerParamsForceLanguage = "khmer"
	BrandGetByTickerParamsForceLanguageKinyarwanda    BrandGetByTickerParamsForceLanguage = "kinyarwanda"
	BrandGetByTickerParamsForceLanguageKorean         BrandGetByTickerParamsForceLanguage = "korean"
	BrandGetByTickerParamsForceLanguageKurdish        BrandGetByTickerParamsForceLanguage = "kurdish"
	BrandGetByTickerParamsForceLanguageKyrgyz         BrandGetByTickerParamsForceLanguage = "kyrgyz"
	BrandGetByTickerParamsForceLanguageLao            BrandGetByTickerParamsForceLanguage = "lao"
	BrandGetByTickerParamsForceLanguageLatin          BrandGetByTickerParamsForceLanguage = "latin"
	BrandGetByTickerParamsForceLanguageLatvian        BrandGetByTickerParamsForceLanguage = "latvian"
	BrandGetByTickerParamsForceLanguageLingala        BrandGetByTickerParamsForceLanguage = "lingala"
	BrandGetByTickerParamsForceLanguageLithuanian     BrandGetByTickerParamsForceLanguage = "lithuanian"
	BrandGetByTickerParamsForceLanguageLuxembourgish  BrandGetByTickerParamsForceLanguage = "luxembourgish"
	BrandGetByTickerParamsForceLanguageMacedonian     BrandGetByTickerParamsForceLanguage = "macedonian"
	BrandGetByTickerParamsForceLanguageMalagasy       BrandGetByTickerParamsForceLanguage = "malagasy"
	BrandGetByTickerParamsForceLanguageMalay          BrandGetByTickerParamsForceLanguage = "malay"
	BrandGetByTickerParamsForceLanguageMalayalam      BrandGetByTickerParamsForceLanguage = "malayalam"
	BrandGetByTickerParamsForceLanguageMaltese        BrandGetByTickerParamsForceLanguage = "maltese"
	BrandGetByTickerParamsForceLanguageMaori          BrandGetByTickerParamsForceLanguage = "maori"
	BrandGetByTickerParamsForceLanguageMarathi        BrandGetByTickerParamsForceLanguage = "marathi"
	BrandGetByTickerParamsForceLanguageMongolian      BrandGetByTickerParamsForceLanguage = "mongolian"
	BrandGetByTickerParamsForceLanguageNepali         BrandGetByTickerParamsForceLanguage = "nepali"
	BrandGetByTickerParamsForceLanguageNorwegian      BrandGetByTickerParamsForceLanguage = "norwegian"
	BrandGetByTickerParamsForceLanguageOdia           BrandGetByTickerParamsForceLanguage = "odia"
	BrandGetByTickerParamsForceLanguageOromo          BrandGetByTickerParamsForceLanguage = "oromo"
	BrandGetByTickerParamsForceLanguagePashto         BrandGetByTickerParamsForceLanguage = "pashto"
	BrandGetByTickerParamsForceLanguagePidgin         BrandGetByTickerParamsForceLanguage = "pidgin"
	BrandGetByTickerParamsForceLanguagePolish         BrandGetByTickerParamsForceLanguage = "polish"
	BrandGetByTickerParamsForceLanguagePortuguese     BrandGetByTickerParamsForceLanguage = "portuguese"
	BrandGetByTickerParamsForceLanguagePunjabi        BrandGetByTickerParamsForceLanguage = "punjabi"
	BrandGetByTickerParamsForceLanguageQuechua        BrandGetByTickerParamsForceLanguage = "quechua"
	BrandGetByTickerParamsForceLanguageRomanian       BrandGetByTickerParamsForceLanguage = "romanian"
	BrandGetByTickerParamsForceLanguageRussian        BrandGetByTickerParamsForceLanguage = "russian"
	BrandGetByTickerParamsForceLanguageSamoan         BrandGetByTickerParamsForceLanguage = "samoan"
	BrandGetByTickerParamsForceLanguageScottishGaelic BrandGetByTickerParamsForceLanguage = "scottish-gaelic"
	BrandGetByTickerParamsForceLanguageSerbian        BrandGetByTickerParamsForceLanguage = "serbian"
	BrandGetByTickerParamsForceLanguageSesotho        BrandGetByTickerParamsForceLanguage = "sesotho"
	BrandGetByTickerParamsForceLanguageShona          BrandGetByTickerParamsForceLanguage = "shona"
	BrandGetByTickerParamsForceLanguageSindhi         BrandGetByTickerParamsForceLanguage = "sindhi"
	BrandGetByTickerParamsForceLanguageSinhala        BrandGetByTickerParamsForceLanguage = "sinhala"
	BrandGetByTickerParamsForceLanguageSlovak         BrandGetByTickerParamsForceLanguage = "slovak"
	BrandGetByTickerParamsForceLanguageSlovene        BrandGetByTickerParamsForceLanguage = "slovene"
	BrandGetByTickerParamsForceLanguageSomali         BrandGetByTickerParamsForceLanguage = "somali"
	BrandGetByTickerParamsForceLanguageSpanish        BrandGetByTickerParamsForceLanguage = "spanish"
	BrandGetByTickerParamsForceLanguageSundanese      BrandGetByTickerParamsForceLanguage = "sundanese"
	BrandGetByTickerParamsForceLanguageSwahili        BrandGetByTickerParamsForceLanguage = "swahili"
	BrandGetByTickerParamsForceLanguageSwedish        BrandGetByTickerParamsForceLanguage = "swedish"
	BrandGetByTickerParamsForceLanguageTagalog        BrandGetByTickerParamsForceLanguage = "tagalog"
	BrandGetByTickerParamsForceLanguageTajik          BrandGetByTickerParamsForceLanguage = "tajik"
	BrandGetByTickerParamsForceLanguageTamil          BrandGetByTickerParamsForceLanguage = "tamil"
	BrandGetByTickerParamsForceLanguageTatar          BrandGetByTickerParamsForceLanguage = "tatar"
	BrandGetByTickerParamsForceLanguageTelugu         BrandGetByTickerParamsForceLanguage = "telugu"
	BrandGetByTickerParamsForceLanguageThai           BrandGetByTickerParamsForceLanguage = "thai"
	BrandGetByTickerParamsForceLanguageTibetan        BrandGetByTickerParamsForceLanguage = "tibetan"
	BrandGetByTickerParamsForceLanguageTigrinya       BrandGetByTickerParamsForceLanguage = "tigrinya"
	BrandGetByTickerParamsForceLanguageTongan         BrandGetByTickerParamsForceLanguage = "tongan"
	BrandGetByTickerParamsForceLanguageTswana         BrandGetByTickerParamsForceLanguage = "tswana"
	BrandGetByTickerParamsForceLanguageTurkish        BrandGetByTickerParamsForceLanguage = "turkish"
	BrandGetByTickerParamsForceLanguageTurkmen        BrandGetByTickerParamsForceLanguage = "turkmen"
	BrandGetByTickerParamsForceLanguageUkrainian      BrandGetByTickerParamsForceLanguage = "ukrainian"
	BrandGetByTickerParamsForceLanguageUrdu           BrandGetByTickerParamsForceLanguage = "urdu"
	BrandGetByTickerParamsForceLanguageUyghur         BrandGetByTickerParamsForceLanguage = "uyghur"
	BrandGetByTickerParamsForceLanguageUzbek          BrandGetByTickerParamsForceLanguage = "uzbek"
	BrandGetByTickerParamsForceLanguageVietnamese     BrandGetByTickerParamsForceLanguage = "vietnamese"
	BrandGetByTickerParamsForceLanguageWelsh          BrandGetByTickerParamsForceLanguage = "welsh"
	BrandGetByTickerParamsForceLanguageWolof          BrandGetByTickerParamsForceLanguage = "wolof"
	BrandGetByTickerParamsForceLanguageXhosa          BrandGetByTickerParamsForceLanguage = "xhosa"
	BrandGetByTickerParamsForceLanguageYiddish        BrandGetByTickerParamsForceLanguage = "yiddish"
	BrandGetByTickerParamsForceLanguageYoruba         BrandGetByTickerParamsForceLanguage = "yoruba"
	BrandGetByTickerParamsForceLanguageZulu           BrandGetByTickerParamsForceLanguage = "zulu"
)

// Optional stock exchange for the ticker. Defaults to NASDAQ if not specified.
type BrandGetByTickerParamsTickerExchange string

const (
	BrandGetByTickerParamsTickerExchangeAmex   BrandGetByTickerParamsTickerExchange = "AMEX"
	BrandGetByTickerParamsTickerExchangeAms    BrandGetByTickerParamsTickerExchange = "AMS"
	BrandGetByTickerParamsTickerExchangeAqs    BrandGetByTickerParamsTickerExchange = "AQS"
	BrandGetByTickerParamsTickerExchangeAsx    BrandGetByTickerParamsTickerExchange = "ASX"
	BrandGetByTickerParamsTickerExchangeAth    BrandGetByTickerParamsTickerExchange = "ATH"
	BrandGetByTickerParamsTickerExchangeBer    BrandGetByTickerParamsTickerExchange = "BER"
	BrandGetByTickerParamsTickerExchangeBme    BrandGetByTickerParamsTickerExchange = "BME"
	BrandGetByTickerParamsTickerExchangeBru    BrandGetByTickerParamsTickerExchange = "BRU"
	BrandGetByTickerParamsTickerExchangeBse    BrandGetByTickerParamsTickerExchange = "BSE"
	BrandGetByTickerParamsTickerExchangeBud    BrandGetByTickerParamsTickerExchange = "BUD"
	BrandGetByTickerParamsTickerExchangeBue    BrandGetByTickerParamsTickerExchange = "BUE"
	BrandGetByTickerParamsTickerExchangeBvc    BrandGetByTickerParamsTickerExchange = "BVC"
	BrandGetByTickerParamsTickerExchangeCboe   BrandGetByTickerParamsTickerExchange = "CBOE"
	BrandGetByTickerParamsTickerExchangeCnq    BrandGetByTickerParamsTickerExchange = "CNQ"
	BrandGetByTickerParamsTickerExchangeCph    BrandGetByTickerParamsTickerExchange = "CPH"
	BrandGetByTickerParamsTickerExchangeDfm    BrandGetByTickerParamsTickerExchange = "DFM"
	BrandGetByTickerParamsTickerExchangeDoh    BrandGetByTickerParamsTickerExchange = "DOH"
	BrandGetByTickerParamsTickerExchangeDub    BrandGetByTickerParamsTickerExchange = "DUB"
	BrandGetByTickerParamsTickerExchangeDus    BrandGetByTickerParamsTickerExchange = "DUS"
	BrandGetByTickerParamsTickerExchangeDxe    BrandGetByTickerParamsTickerExchange = "DXE"
	BrandGetByTickerParamsTickerExchangeEgx    BrandGetByTickerParamsTickerExchange = "EGX"
	BrandGetByTickerParamsTickerExchangeFsx    BrandGetByTickerParamsTickerExchange = "FSX"
	BrandGetByTickerParamsTickerExchangeHam    BrandGetByTickerParamsTickerExchange = "HAM"
	BrandGetByTickerParamsTickerExchangeHel    BrandGetByTickerParamsTickerExchange = "HEL"
	BrandGetByTickerParamsTickerExchangeHkse   BrandGetByTickerParamsTickerExchange = "HKSE"
	BrandGetByTickerParamsTickerExchangeHose   BrandGetByTickerParamsTickerExchange = "HOSE"
	BrandGetByTickerParamsTickerExchangeIce    BrandGetByTickerParamsTickerExchange = "ICE"
	BrandGetByTickerParamsTickerExchangeIob    BrandGetByTickerParamsTickerExchange = "IOB"
	BrandGetByTickerParamsTickerExchangeIst    BrandGetByTickerParamsTickerExchange = "IST"
	BrandGetByTickerParamsTickerExchangeJkt    BrandGetByTickerParamsTickerExchange = "JKT"
	BrandGetByTickerParamsTickerExchangeJnb    BrandGetByTickerParamsTickerExchange = "JNB"
	BrandGetByTickerParamsTickerExchangeJpx    BrandGetByTickerParamsTickerExchange = "JPX"
	BrandGetByTickerParamsTickerExchangeKls    BrandGetByTickerParamsTickerExchange = "KLS"
	BrandGetByTickerParamsTickerExchangeKoe    BrandGetByTickerParamsTickerExchange = "KOE"
	BrandGetByTickerParamsTickerExchangeKsc    BrandGetByTickerParamsTickerExchange = "KSC"
	BrandGetByTickerParamsTickerExchangeKuw    BrandGetByTickerParamsTickerExchange = "KUW"
	BrandGetByTickerParamsTickerExchangeLis    BrandGetByTickerParamsTickerExchange = "LIS"
	BrandGetByTickerParamsTickerExchangeLse    BrandGetByTickerParamsTickerExchange = "LSE"
	BrandGetByTickerParamsTickerExchangeMcx    BrandGetByTickerParamsTickerExchange = "MCX"
	BrandGetByTickerParamsTickerExchangeMex    BrandGetByTickerParamsTickerExchange = "MEX"
	BrandGetByTickerParamsTickerExchangeMil    BrandGetByTickerParamsTickerExchange = "MIL"
	BrandGetByTickerParamsTickerExchangeMun    BrandGetByTickerParamsTickerExchange = "MUN"
	BrandGetByTickerParamsTickerExchangeNasdaq BrandGetByTickerParamsTickerExchange = "NASDAQ"
	BrandGetByTickerParamsTickerExchangeNeo    BrandGetByTickerParamsTickerExchange = "NEO"
	BrandGetByTickerParamsTickerExchangeNse    BrandGetByTickerParamsTickerExchange = "NSE"
	BrandGetByTickerParamsTickerExchangeNyse   BrandGetByTickerParamsTickerExchange = "NYSE"
	BrandGetByTickerParamsTickerExchangeNze    BrandGetByTickerParamsTickerExchange = "NZE"
	BrandGetByTickerParamsTickerExchangeOsl    BrandGetByTickerParamsTickerExchange = "OSL"
	BrandGetByTickerParamsTickerExchangeOtc    BrandGetByTickerParamsTickerExchange = "OTC"
	BrandGetByTickerParamsTickerExchangePar    BrandGetByTickerParamsTickerExchange = "PAR"
	BrandGetByTickerParamsTickerExchangePnk    BrandGetByTickerParamsTickerExchange = "PNK"
	BrandGetByTickerParamsTickerExchangePra    BrandGetByTickerParamsTickerExchange = "PRA"
	BrandGetByTickerParamsTickerExchangeRis    BrandGetByTickerParamsTickerExchange = "RIS"
	BrandGetByTickerParamsTickerExchangeSao    BrandGetByTickerParamsTickerExchange = "SAO"
	BrandGetByTickerParamsTickerExchangeSau    BrandGetByTickerParamsTickerExchange = "SAU"
	BrandGetByTickerParamsTickerExchangeSes    BrandGetByTickerParamsTickerExchange = "SES"
	BrandGetByTickerParamsTickerExchangeSet    BrandGetByTickerParamsTickerExchange = "SET"
	BrandGetByTickerParamsTickerExchangeSgo    BrandGetByTickerParamsTickerExchange = "SGO"
	BrandGetByTickerParamsTickerExchangeShh    BrandGetByTickerParamsTickerExchange = "SHH"
	BrandGetByTickerParamsTickerExchangeShz    BrandGetByTickerParamsTickerExchange = "SHZ"
	BrandGetByTickerParamsTickerExchangeSix    BrandGetByTickerParamsTickerExchange = "SIX"
	BrandGetByTickerParamsTickerExchangeSto    BrandGetByTickerParamsTickerExchange = "STO"
	BrandGetByTickerParamsTickerExchangeStu    BrandGetByTickerParamsTickerExchange = "STU"
	BrandGetByTickerParamsTickerExchangeTai    BrandGetByTickerParamsTickerExchange = "TAI"
	BrandGetByTickerParamsTickerExchangeTal    BrandGetByTickerParamsTickerExchange = "TAL"
	BrandGetByTickerParamsTickerExchangeTlv    BrandGetByTickerParamsTickerExchange = "TLV"
	BrandGetByTickerParamsTickerExchangeTsx    BrandGetByTickerParamsTickerExchange = "TSX"
	BrandGetByTickerParamsTickerExchangeTsxv   BrandGetByTickerParamsTickerExchange = "TSXV"
	BrandGetByTickerParamsTickerExchangeTwo    BrandGetByTickerParamsTickerExchange = "TWO"
	BrandGetByTickerParamsTickerExchangeVie    BrandGetByTickerParamsTickerExchange = "VIE"
	BrandGetByTickerParamsTickerExchangeWse    BrandGetByTickerParamsTickerExchange = "WSE"
	BrandGetByTickerParamsTickerExchangeXetra  BrandGetByTickerParamsTickerExchange = "XETRA"
)

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
