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
	// Optional HTTP(S) source document URL used to resolve relative links and image
	// references. Relative references remain relative when omitted.
	BaseURL param.Opt[string] `query:"baseUrl,omitzero" format:"uri" json:"-"`
	// Optional file extension hint, such as pdf, docx, xlsx, pptx, html, json, csv,
	// md, py, rtf, jpg, png, or txt.
	Extension param.Opt[string] `query:"extension,omitzero" json:"-"`
	// Optional filename hint used to infer the extension when extension is omitted.
	Filename param.Opt[string] `query:"filename,omitzero" json:"-"`
	// Include image references in Markdown output
	IncludeImages param.Opt[bool] `query:"includeImages,omitzero" json:"-"`
	// Preserve hyperlinks in Markdown output
	IncludeLinks param.Opt[bool] `query:"includeLinks,omitzero" json:"-"`
	// When true for PDF inputs, detect and OCR images embedded in the selected pages,
	// inserting recognized text at each image's position in page reading order while
	// preserving the PDF text layer. pdfStart/pdfEnd limit the inclusive page range.
	// This is separate from automatic scanned-PDF OCR fallback.
	Ocr param.Opt[bool] `query:"ocr,omitzero" json:"-"`
	// Last 1-based PDF page to parse. When omitted, parsing ends at the last page.
	// Must be greater than or equal to pdfStart when both are provided.
	PdfEnd param.Opt[int64] `query:"pdfEnd,omitzero" json:"-"`
	// First 1-based PDF page to parse. When omitted, parsing starts at the first page.
	PdfStart param.Opt[int64] `query:"pdfStart,omitzero" json:"-"`
	// Shorten base64-encoded image data in the Markdown output
	ShortenBase64Images param.Opt[bool] `query:"shortenBase64Images,omitzero" json:"-"`
	// Extract only the main content from HTML-like inputs
	UseMainContentOnly param.Opt[bool] `query:"useMainContentOnly,omitzero" json:"-"`
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
