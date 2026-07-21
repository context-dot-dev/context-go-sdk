// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contextdev

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"

	"github.com/context-dot-dev/context-go-sdk/v2/internal/apiform"
	"github.com/context-dot-dev/context-go-sdk/v2/internal/apijson"
	"github.com/context-dot-dev/context-go-sdk/v2/internal/apiquery"
	"github.com/context-dot-dev/context-go-sdk/v2/internal/requestconfig"
	"github.com/context-dot-dev/context-go-sdk/v2/option"
	"github.com/context-dot-dev/context-go-sdk/v2/packages/param"
	"github.com/context-dot-dev/context-go-sdk/v2/packages/respjson"
)

// ParseService contains methods and other services that help with interacting with
// the context.dev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewParseService] method instead.
type ParseService struct {
	options []option.RequestOption
}

// NewParseService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewParseService(opts ...option.RequestOption) (r ParseService) {
	r = ParseService{}
	r.options = opts
	return
}

// Converts raw text, source code, web/data, PDF, Microsoft Office, and image bytes
// into LLM-usable Markdown.
func (r *ParseService) Handle(ctx context.Context, body io.Reader, params ParseHandleParams, opts ...option.RequestOption) (res *ParseHandleResponse, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithRequestBody("application/octet-stream", body)}, opts...)
	path := "parse"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type ParseHandleResponse struct {
	// Input bytes converted to GitHub Flavored Markdown
	Markdown string `json:"markdown" api:"required"`
	// Indicates success
	//
	// Any of true.
	Success bool `json:"success" api:"required"`
	// Detected content type used for parsing
	//
	// Any of "html", "xml", "json", "jsonl", "text", "csv", "tsv", "markdown", "yaml",
	// "python", "java", "javascript", "php", "shell", "ruby", "typescript", "rtf",
	// "srt", "css", "scss", "less", "stylus", "sass", "svg", "pdf", "docx", "doc",
	// "xlsx", "xls", "pptx", "ppt", "jpg", "png", "gif", "bmp", "tiff", "webp", "ppm",
	// "pbm", "pgm", "pnm".
	Type ParseHandleResponseType `json:"type" api:"required"`
	// Metadata about the API key used for the request. Included in every response
	// whenever a valid API key is provided, even when the response status is not 200.
	KeyMetadata ParseHandleResponseKeyMetadata `json:"key_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Markdown    respjson.Field
		Success     respjson.Field
		Type        respjson.Field
		KeyMetadata respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParseHandleResponse) RawJSON() string { return r.JSON.raw }
func (r *ParseHandleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detected content type used for parsing
type ParseHandleResponseType string

const (
	ParseHandleResponseTypeHTML       ParseHandleResponseType = "html"
	ParseHandleResponseTypeXml        ParseHandleResponseType = "xml"
	ParseHandleResponseTypeJson       ParseHandleResponseType = "json"
	ParseHandleResponseTypeJSONL      ParseHandleResponseType = "jsonl"
	ParseHandleResponseTypeText       ParseHandleResponseType = "text"
	ParseHandleResponseTypeCsv        ParseHandleResponseType = "csv"
	ParseHandleResponseTypeTsv        ParseHandleResponseType = "tsv"
	ParseHandleResponseTypeMarkdown   ParseHandleResponseType = "markdown"
	ParseHandleResponseTypeYaml       ParseHandleResponseType = "yaml"
	ParseHandleResponseTypePython     ParseHandleResponseType = "python"
	ParseHandleResponseTypeJava       ParseHandleResponseType = "java"
	ParseHandleResponseTypeJavascript ParseHandleResponseType = "javascript"
	ParseHandleResponseTypePhp        ParseHandleResponseType = "php"
	ParseHandleResponseTypeShell      ParseHandleResponseType = "shell"
	ParseHandleResponseTypeRuby       ParseHandleResponseType = "ruby"
	ParseHandleResponseTypeTypescript ParseHandleResponseType = "typescript"
	ParseHandleResponseTypeRtf        ParseHandleResponseType = "rtf"
	ParseHandleResponseTypeSrt        ParseHandleResponseType = "srt"
	ParseHandleResponseTypeCss        ParseHandleResponseType = "css"
	ParseHandleResponseTypeScss       ParseHandleResponseType = "scss"
	ParseHandleResponseTypeLess       ParseHandleResponseType = "less"
	ParseHandleResponseTypeStylus     ParseHandleResponseType = "stylus"
	ParseHandleResponseTypeSass       ParseHandleResponseType = "sass"
	ParseHandleResponseTypeSvg        ParseHandleResponseType = "svg"
	ParseHandleResponseTypePdf        ParseHandleResponseType = "pdf"
	ParseHandleResponseTypeDocx       ParseHandleResponseType = "docx"
	ParseHandleResponseTypeDoc        ParseHandleResponseType = "doc"
	ParseHandleResponseTypeXlsx       ParseHandleResponseType = "xlsx"
	ParseHandleResponseTypeXls        ParseHandleResponseType = "xls"
	ParseHandleResponseTypePptx       ParseHandleResponseType = "pptx"
	ParseHandleResponseTypePpt        ParseHandleResponseType = "ppt"
	ParseHandleResponseTypeJpg        ParseHandleResponseType = "jpg"
	ParseHandleResponseTypePng        ParseHandleResponseType = "png"
	ParseHandleResponseTypeGif        ParseHandleResponseType = "gif"
	ParseHandleResponseTypeBmp        ParseHandleResponseType = "bmp"
	ParseHandleResponseTypeTiff       ParseHandleResponseType = "tiff"
	ParseHandleResponseTypeWebp       ParseHandleResponseType = "webp"
	ParseHandleResponseTypePpm        ParseHandleResponseType = "ppm"
	ParseHandleResponseTypePbm        ParseHandleResponseType = "pbm"
	ParseHandleResponseTypePgm        ParseHandleResponseType = "pgm"
	ParseHandleResponseTypePnm        ParseHandleResponseType = "pnm"
)

// Metadata about the API key used for the request. Included in every response
// whenever a valid API key is provided, even when the response status is not 200.
type ParseHandleResponseKeyMetadata struct {
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
func (r ParseHandleResponseKeyMetadata) RawJSON() string { return r.JSON.raw }
func (r *ParseHandleResponseKeyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ParseHandleParams struct {
	// Optional client identifier used for usage attribution.
	Client param.Opt[string] `query:"client,omitzero" json:"-"`
	// Optional file extension hint, such as pdf, docx, xlsx, pptx, html, json, csv,
	// md, py, rtf, jpg, png, or txt.
	//
	// Any of "txt", "text", "md", "markdown", "html", "htm", "xhtml", "xml", "rss",
	// "atom", "csv", "tsv", "yaml", "yml", "py", "java", "js", "jsx", "mjs", "cjs",
	// "json", "jsonl", "ndjson", "php", "sh", "bash", "zsh", "fish", "rb", "ts",
	// "tsx", "rtf", "srt", "css", "scss", "less", "styl", "sass", "svg", "pdf",
	// "docx", "doc", "xlsx", "xlsm", "xlsb", "xltx", "xltm", "xls", "pptx", "pptm",
	// "ppsx", "ppsm", "potx", "potm", "ppt", "pps", "pot", "jpg", "jpeg", "jpe",
	// "png", "gif", "bmp", "tiff", "tif", "webp", "ppm", "pbm", "pgm", "pnm".
	Extension ParseHandleParamsExtension `query:"extension,omitzero" json:"-"`
	// Include image references in Markdown output
	IncludeImages ParseHandleParamsIncludeImagesUnion `query:"includeImages,omitzero" json:"-"`
	// Preserve hyperlinks in Markdown output
	IncludeLinks ParseHandleParamsIncludeLinksUnion `query:"includeLinks,omitzero" json:"-"`
	// When true for PDF inputs, detect and OCR images embedded in the selected pages,
	// inserting recognized text at each image's position in page reading order while
	// preserving the PDF text layer. pdf.start/pdf.end limit the inclusive page range.
	// When false, all OCR is disabled, including the automatic scanned-PDF fallback.
	Ocr ParseHandleParamsOcrUnion `query:"ocr,omitzero" json:"-"`
	// PDF page-range options as a JSON object, e.g. {"start": 2, "end": 5}.
	Pdf ParseHandleParamsPdf `query:"pdf,omitzero" json:"-"`
	// Shorten base64-encoded image data in the Markdown output
	ShortenBase64Images ParseHandleParamsShortenBase64ImagesUnion `query:"shortenBase64Images,omitzero" json:"-"`
	// Optional comma-separated caller-defined tags for tracking this request. Tags are
	// recorded on the request's usage log and can be used to filter usage on the
	// dashboard usage page. Up to 20 tags, each 1-50 characters.
	Tags []string `query:"tags,omitzero" json:"-"`
	// Extract only the main content from HTML-like inputs
	UseMainContentOnly ParseHandleParamsUseMainContentOnlyUnion `query:"useMainContentOnly,omitzero" json:"-"`
	// Set to enabled to bypass shared caches and omit request and response content
	// from retained usage logs. Requires zero data retention to be enabled for your
	// organization (contact support@context.dev), otherwise the request fails with
	// ZDR_NOT_ENABLED. Successful ZDR responses include X-Context-ZDR: true.
	//
	// Any of "enabled", "disabled".
	Zdr ParseHandleParamsZdr `query:"zdr,omitzero" json:"-"`
	paramObj
}

func (r ParseHandleParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// URLQuery serializes [ParseHandleParams]'s query parameters as `url.Values`.
func (r ParseHandleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Optional file extension hint, such as pdf, docx, xlsx, pptx, html, json, csv,
// md, py, rtf, jpg, png, or txt.
type ParseHandleParamsExtension string

const (
	ParseHandleParamsExtensionTxt      ParseHandleParamsExtension = "txt"
	ParseHandleParamsExtensionText     ParseHandleParamsExtension = "text"
	ParseHandleParamsExtensionMd       ParseHandleParamsExtension = "md"
	ParseHandleParamsExtensionMarkdown ParseHandleParamsExtension = "markdown"
	ParseHandleParamsExtensionHTML     ParseHandleParamsExtension = "html"
	ParseHandleParamsExtensionHtm      ParseHandleParamsExtension = "htm"
	ParseHandleParamsExtensionXhtml    ParseHandleParamsExtension = "xhtml"
	ParseHandleParamsExtensionXml      ParseHandleParamsExtension = "xml"
	ParseHandleParamsExtensionRss      ParseHandleParamsExtension = "rss"
	ParseHandleParamsExtensionAtom     ParseHandleParamsExtension = "atom"
	ParseHandleParamsExtensionCsv      ParseHandleParamsExtension = "csv"
	ParseHandleParamsExtensionTsv      ParseHandleParamsExtension = "tsv"
	ParseHandleParamsExtensionYaml     ParseHandleParamsExtension = "yaml"
	ParseHandleParamsExtensionYml      ParseHandleParamsExtension = "yml"
	ParseHandleParamsExtensionPy       ParseHandleParamsExtension = "py"
	ParseHandleParamsExtensionJava     ParseHandleParamsExtension = "java"
	ParseHandleParamsExtensionJs       ParseHandleParamsExtension = "js"
	ParseHandleParamsExtensionJsx      ParseHandleParamsExtension = "jsx"
	ParseHandleParamsExtensionMjs      ParseHandleParamsExtension = "mjs"
	ParseHandleParamsExtensionCjs      ParseHandleParamsExtension = "cjs"
	ParseHandleParamsExtensionJson     ParseHandleParamsExtension = "json"
	ParseHandleParamsExtensionJSONL    ParseHandleParamsExtension = "jsonl"
	ParseHandleParamsExtensionNdjson   ParseHandleParamsExtension = "ndjson"
	ParseHandleParamsExtensionPhp      ParseHandleParamsExtension = "php"
	ParseHandleParamsExtensionSh       ParseHandleParamsExtension = "sh"
	ParseHandleParamsExtensionBash     ParseHandleParamsExtension = "bash"
	ParseHandleParamsExtensionZsh      ParseHandleParamsExtension = "zsh"
	ParseHandleParamsExtensionFish     ParseHandleParamsExtension = "fish"
	ParseHandleParamsExtensionRb       ParseHandleParamsExtension = "rb"
	ParseHandleParamsExtensionTs       ParseHandleParamsExtension = "ts"
	ParseHandleParamsExtensionTsx      ParseHandleParamsExtension = "tsx"
	ParseHandleParamsExtensionRtf      ParseHandleParamsExtension = "rtf"
	ParseHandleParamsExtensionSrt      ParseHandleParamsExtension = "srt"
	ParseHandleParamsExtensionCss      ParseHandleParamsExtension = "css"
	ParseHandleParamsExtensionScss     ParseHandleParamsExtension = "scss"
	ParseHandleParamsExtensionLess     ParseHandleParamsExtension = "less"
	ParseHandleParamsExtensionStyl     ParseHandleParamsExtension = "styl"
	ParseHandleParamsExtensionSass     ParseHandleParamsExtension = "sass"
	ParseHandleParamsExtensionSvg      ParseHandleParamsExtension = "svg"
	ParseHandleParamsExtensionPdf      ParseHandleParamsExtension = "pdf"
	ParseHandleParamsExtensionDocx     ParseHandleParamsExtension = "docx"
	ParseHandleParamsExtensionDoc      ParseHandleParamsExtension = "doc"
	ParseHandleParamsExtensionXlsx     ParseHandleParamsExtension = "xlsx"
	ParseHandleParamsExtensionXlsm     ParseHandleParamsExtension = "xlsm"
	ParseHandleParamsExtensionXlsb     ParseHandleParamsExtension = "xlsb"
	ParseHandleParamsExtensionXltx     ParseHandleParamsExtension = "xltx"
	ParseHandleParamsExtensionXltm     ParseHandleParamsExtension = "xltm"
	ParseHandleParamsExtensionXls      ParseHandleParamsExtension = "xls"
	ParseHandleParamsExtensionPptx     ParseHandleParamsExtension = "pptx"
	ParseHandleParamsExtensionPptm     ParseHandleParamsExtension = "pptm"
	ParseHandleParamsExtensionPpsx     ParseHandleParamsExtension = "ppsx"
	ParseHandleParamsExtensionPpsm     ParseHandleParamsExtension = "ppsm"
	ParseHandleParamsExtensionPotx     ParseHandleParamsExtension = "potx"
	ParseHandleParamsExtensionPotm     ParseHandleParamsExtension = "potm"
	ParseHandleParamsExtensionPpt      ParseHandleParamsExtension = "ppt"
	ParseHandleParamsExtensionPps      ParseHandleParamsExtension = "pps"
	ParseHandleParamsExtensionPot      ParseHandleParamsExtension = "pot"
	ParseHandleParamsExtensionJpg      ParseHandleParamsExtension = "jpg"
	ParseHandleParamsExtensionJpeg     ParseHandleParamsExtension = "jpeg"
	ParseHandleParamsExtensionJpe      ParseHandleParamsExtension = "jpe"
	ParseHandleParamsExtensionPng      ParseHandleParamsExtension = "png"
	ParseHandleParamsExtensionGif      ParseHandleParamsExtension = "gif"
	ParseHandleParamsExtensionBmp      ParseHandleParamsExtension = "bmp"
	ParseHandleParamsExtensionTiff     ParseHandleParamsExtension = "tiff"
	ParseHandleParamsExtensionTif      ParseHandleParamsExtension = "tif"
	ParseHandleParamsExtensionWebp     ParseHandleParamsExtension = "webp"
	ParseHandleParamsExtensionPpm      ParseHandleParamsExtension = "ppm"
	ParseHandleParamsExtensionPbm      ParseHandleParamsExtension = "pbm"
	ParseHandleParamsExtensionPgm      ParseHandleParamsExtension = "pgm"
	ParseHandleParamsExtensionPnm      ParseHandleParamsExtension = "pnm"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParseHandleParamsIncludeImagesUnion struct {
	OfBool param.Opt[bool] `query:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfParseHandlesIncludeImagesString)
	OfParseHandlesIncludeImagesString param.Opt[string] `query:",omitzero,inline"`
	paramUnion
}

type ParseHandleParamsIncludeImagesString string

const (
	ParseHandleParamsIncludeImagesStringTrue  ParseHandleParamsIncludeImagesString = "true"
	ParseHandleParamsIncludeImagesStringFalse ParseHandleParamsIncludeImagesString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParseHandleParamsIncludeLinksUnion struct {
	OfBool param.Opt[bool] `query:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfParseHandlesIncludeLinksString)
	OfParseHandlesIncludeLinksString param.Opt[string] `query:",omitzero,inline"`
	paramUnion
}

type ParseHandleParamsIncludeLinksString string

const (
	ParseHandleParamsIncludeLinksStringTrue  ParseHandleParamsIncludeLinksString = "true"
	ParseHandleParamsIncludeLinksStringFalse ParseHandleParamsIncludeLinksString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParseHandleParamsOcrUnion struct {
	OfBool param.Opt[bool] `query:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfParseHandlesOcrString)
	OfParseHandlesOcrString param.Opt[string] `query:",omitzero,inline"`
	paramUnion
}

type ParseHandleParamsOcrString string

const (
	ParseHandleParamsOcrStringTrue  ParseHandleParamsOcrString = "true"
	ParseHandleParamsOcrStringFalse ParseHandleParamsOcrString = "false"
)

// PDF page-range options as a JSON object, e.g. {"start": 2, "end": 5}.
type ParseHandleParamsPdf struct {
	// Last 1-based PDF page to parse. When omitted, parsing ends at the last page.
	// Must be greater than or equal to start when both are provided.
	End param.Opt[int64] `query:"end,omitzero" json:"-"`
	// First 1-based PDF page to parse. When omitted, parsing starts at the first page.
	Start param.Opt[int64] `query:"start,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ParseHandleParamsPdf]'s query parameters as `url.Values`.
func (r ParseHandleParamsPdf) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParseHandleParamsShortenBase64ImagesUnion struct {
	OfBool param.Opt[bool] `query:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfParseHandlesShortenBase64ImagesString)
	OfParseHandlesShortenBase64ImagesString param.Opt[string] `query:",omitzero,inline"`
	paramUnion
}

type ParseHandleParamsShortenBase64ImagesString string

const (
	ParseHandleParamsShortenBase64ImagesStringTrue  ParseHandleParamsShortenBase64ImagesString = "true"
	ParseHandleParamsShortenBase64ImagesStringFalse ParseHandleParamsShortenBase64ImagesString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParseHandleParamsUseMainContentOnlyUnion struct {
	OfBool param.Opt[bool] `query:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfParseHandlesUseMainContentOnlyString)
	OfParseHandlesUseMainContentOnlyString param.Opt[string] `query:",omitzero,inline"`
	paramUnion
}

type ParseHandleParamsUseMainContentOnlyString string

const (
	ParseHandleParamsUseMainContentOnlyStringTrue  ParseHandleParamsUseMainContentOnlyString = "true"
	ParseHandleParamsUseMainContentOnlyStringFalse ParseHandleParamsUseMainContentOnlyString = "false"
)

// Set to enabled to bypass shared caches and omit request and response content
// from retained usage logs. Requires zero data retention to be enabled for your
// organization (contact support@context.dev), otherwise the request fails with
// ZDR_NOT_ENABLED. Successful ZDR responses include X-Context-ZDR: true.
type ParseHandleParamsZdr string

const (
	ParseHandleParamsZdrEnabled  ParseHandleParamsZdr = "enabled"
	ParseHandleParamsZdrDisabled ParseHandleParamsZdr = "disabled"
)
