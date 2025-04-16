// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package taamcloud

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"reflect"

	"github.com/taamsoftadmin/taam-cloud-go-sdk/internal/apiform"
	"github.com/taamsoftadmin/taam-cloud-go-sdk/internal/apijson"
	"github.com/taamsoftadmin/taam-cloud-go-sdk/internal/param"
	"github.com/taamsoftadmin/taam-cloud-go-sdk/internal/requestconfig"
	"github.com/taamsoftadmin/taam-cloud-go-sdk/option"
	"github.com/tidwall/gjson"
)

// UploadService contains methods and other services that help with interacting
// with the taam-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUploadService] method instead.
type UploadService struct {
	Options []option.RequestOption
}

// NewUploadService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUploadService(opts ...option.RequestOption) (r *UploadService) {
	r = &UploadService{}
	r.Options = opts
	return
}

// Process various file types including documents, images, and audio files with
// advanced extraction options
func (r *UploadService) New(ctx context.Context, body UploadNewParams, opts ...option.RequestOption) (res *UploadNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "upload"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type UploadNewResponse struct {
	// Unique identifier for the request
	ID string `json:"id"`
	// Extracted text or data URL
	Content string `json:"content"`
	// Unix timestamp when the request was created
	Created int64 `json:"created"`
	// This field can have the runtime type of
	// [UploadNewResponseFileEmbeddingsResponseData].
	Data interface{} `json:"data"`
	// Error message (if any)
	Error string `json:"error"`
	// Object type (e.g., 'chunks')
	Object string `json:"object"`
	// Operation status
	Status bool `json:"status"`
	// Type of the processed file
	Type string `json:"type"`
	// This field can have the runtime type of
	// [UploadNewResponseFileEmbeddingsResponseUsage].
	Usage interface{}           `json:"usage"`
	JSON  uploadNewResponseJSON `json:"-"`
	union UploadNewResponseUnion
}

// uploadNewResponseJSON contains the JSON metadata for the struct
// [UploadNewResponse]
type uploadNewResponseJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	Created     apijson.Field
	Data        apijson.Field
	Error       apijson.Field
	Object      apijson.Field
	Status      apijson.Field
	Type        apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r uploadNewResponseJSON) RawJSON() string {
	return r.raw
}

func (r *UploadNewResponse) UnmarshalJSON(data []byte) (err error) {
	*r = UploadNewResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [UploadNewResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [UploadNewResponseFileUploadResponse],
// [UploadNewResponseFileEmbeddingsResponse].
func (r UploadNewResponse) AsUnion() UploadNewResponseUnion {
	return r.union
}

// Union satisfied by [UploadNewResponseFileUploadResponse] or
// [UploadNewResponseFileEmbeddingsResponse].
type UploadNewResponseUnion interface {
	implementsUploadNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UploadNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UploadNewResponseFileUploadResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UploadNewResponseFileEmbeddingsResponse{}),
		},
	)
}

type UploadNewResponseFileUploadResponse struct {
	// Extracted text or data URL
	Content string `json:"content"`
	// Error message (if any)
	Error string `json:"error"`
	// Operation status
	Status bool `json:"status"`
	// Type of the processed file
	Type UploadNewResponseFileUploadResponseType `json:"type"`
	JSON uploadNewResponseFileUploadResponseJSON `json:"-"`
}

// uploadNewResponseFileUploadResponseJSON contains the JSON metadata for the
// struct [UploadNewResponseFileUploadResponse]
type uploadNewResponseFileUploadResponseJSON struct {
	Content     apijson.Field
	Error       apijson.Field
	Status      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UploadNewResponseFileUploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uploadNewResponseFileUploadResponseJSON) RawJSON() string {
	return r.raw
}

func (r UploadNewResponseFileUploadResponse) implementsUploadNewResponse() {}

// Type of the processed file
type UploadNewResponseFileUploadResponseType string

const (
	UploadNewResponseFileUploadResponseTypePdf   UploadNewResponseFileUploadResponseType = "pdf"
	UploadNewResponseFileUploadResponseTypeDocx  UploadNewResponseFileUploadResponseType = "docx"
	UploadNewResponseFileUploadResponseTypePptx  UploadNewResponseFileUploadResponseType = "pptx"
	UploadNewResponseFileUploadResponseTypeXlsx  UploadNewResponseFileUploadResponseType = "xlsx"
	UploadNewResponseFileUploadResponseTypeImage UploadNewResponseFileUploadResponseType = "image"
	UploadNewResponseFileUploadResponseTypeAudio UploadNewResponseFileUploadResponseType = "audio"
	UploadNewResponseFileUploadResponseTypeText  UploadNewResponseFileUploadResponseType = "text"
	UploadNewResponseFileUploadResponseTypeError UploadNewResponseFileUploadResponseType = "error"
)

func (r UploadNewResponseFileUploadResponseType) IsKnown() bool {
	switch r {
	case UploadNewResponseFileUploadResponseTypePdf, UploadNewResponseFileUploadResponseTypeDocx, UploadNewResponseFileUploadResponseTypePptx, UploadNewResponseFileUploadResponseTypeXlsx, UploadNewResponseFileUploadResponseTypeImage, UploadNewResponseFileUploadResponseTypeAudio, UploadNewResponseFileUploadResponseTypeText, UploadNewResponseFileUploadResponseTypeError:
		return true
	}
	return false
}

type UploadNewResponseFileEmbeddingsResponse struct {
	// Unique identifier for the request
	ID string `json:"id"`
	// Unix timestamp when the request was created
	Created int64                                       `json:"created"`
	Data    UploadNewResponseFileEmbeddingsResponseData `json:"data"`
	// Error message (if any)
	Error string `json:"error"`
	// Object type (e.g., 'chunks')
	Object string `json:"object"`
	// Type of the processed file
	Type  string                                       `json:"type"`
	Usage UploadNewResponseFileEmbeddingsResponseUsage `json:"usage"`
	JSON  uploadNewResponseFileEmbeddingsResponseJSON  `json:"-"`
}

// uploadNewResponseFileEmbeddingsResponseJSON contains the JSON metadata for the
// struct [UploadNewResponseFileEmbeddingsResponse]
type uploadNewResponseFileEmbeddingsResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Data        apijson.Field
	Error       apijson.Field
	Object      apijson.Field
	Type        apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UploadNewResponseFileEmbeddingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uploadNewResponseFileEmbeddingsResponseJSON) RawJSON() string {
	return r.raw
}

func (r UploadNewResponseFileEmbeddingsResponse) implementsUploadNewResponse() {}

type UploadNewResponseFileEmbeddingsResponseData struct {
	Data    UploadNewResponseFileEmbeddingsResponseDataData `json:"data"`
	Success bool                                            `json:"success"`
	JSON    uploadNewResponseFileEmbeddingsResponseDataJSON `json:"-"`
}

// uploadNewResponseFileEmbeddingsResponseDataJSON contains the JSON metadata for
// the struct [UploadNewResponseFileEmbeddingsResponseData]
type uploadNewResponseFileEmbeddingsResponseDataJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UploadNewResponseFileEmbeddingsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uploadNewResponseFileEmbeddingsResponseDataJSON) RawJSON() string {
	return r.raw
}

type UploadNewResponseFileEmbeddingsResponseDataData struct {
	Chunks   []UploadNewResponseFileEmbeddingsResponseDataDataChunk  `json:"chunks"`
	Metadata UploadNewResponseFileEmbeddingsResponseDataDataMetadata `json:"metadata"`
	JSON     uploadNewResponseFileEmbeddingsResponseDataDataJSON     `json:"-"`
}

// uploadNewResponseFileEmbeddingsResponseDataDataJSON contains the JSON metadata
// for the struct [UploadNewResponseFileEmbeddingsResponseDataData]
type uploadNewResponseFileEmbeddingsResponseDataDataJSON struct {
	Chunks      apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UploadNewResponseFileEmbeddingsResponseDataData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uploadNewResponseFileEmbeddingsResponseDataDataJSON) RawJSON() string {
	return r.raw
}

type UploadNewResponseFileEmbeddingsResponseDataDataChunk struct {
	// Text content of the chunk
	Content string `json:"content"`
	// Page number the chunk is from
	FromPage int64 `json:"from_page"`
	// Number of tokens in the chunk
	TotalTokens int64                                                    `json:"total_tokens"`
	JSON        uploadNewResponseFileEmbeddingsResponseDataDataChunkJSON `json:"-"`
}

// uploadNewResponseFileEmbeddingsResponseDataDataChunkJSON contains the JSON
// metadata for the struct [UploadNewResponseFileEmbeddingsResponseDataDataChunk]
type uploadNewResponseFileEmbeddingsResponseDataDataChunkJSON struct {
	Content     apijson.Field
	FromPage    apijson.Field
	TotalTokens apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UploadNewResponseFileEmbeddingsResponseDataDataChunk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uploadNewResponseFileEmbeddingsResponseDataDataChunkJSON) RawJSON() string {
	return r.raw
}

type UploadNewResponseFileEmbeddingsResponseDataDataMetadata struct {
	// Document description
	Description string `json:"description"`
	// Document language
	Language string `json:"language"`
	// Document title
	Title string                                                      `json:"title"`
	JSON  uploadNewResponseFileEmbeddingsResponseDataDataMetadataJSON `json:"-"`
}

// uploadNewResponseFileEmbeddingsResponseDataDataMetadataJSON contains the JSON
// metadata for the struct
// [UploadNewResponseFileEmbeddingsResponseDataDataMetadata]
type uploadNewResponseFileEmbeddingsResponseDataDataMetadataJSON struct {
	Description apijson.Field
	Language    apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UploadNewResponseFileEmbeddingsResponseDataDataMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uploadNewResponseFileEmbeddingsResponseDataDataMetadataJSON) RawJSON() string {
	return r.raw
}

type UploadNewResponseFileEmbeddingsResponseUsage struct {
	// Total number of chunks
	TotalChunks int64 `json:"total_chunks"`
	// Total characters of extracted text
	TotalExtractedText int64 `json:"total_extracted_text"`
	// Total number of pages
	TotalPages int64 `json:"total_pages"`
	// Total number of tokens
	TotalTokens int64                                            `json:"total_tokens"`
	JSON        uploadNewResponseFileEmbeddingsResponseUsageJSON `json:"-"`
}

// uploadNewResponseFileEmbeddingsResponseUsageJSON contains the JSON metadata for
// the struct [UploadNewResponseFileEmbeddingsResponseUsage]
type uploadNewResponseFileEmbeddingsResponseUsageJSON struct {
	TotalChunks        apijson.Field
	TotalExtractedText apijson.Field
	TotalPages         apijson.Field
	TotalTokens        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *UploadNewResponseFileEmbeddingsResponseUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uploadNewResponseFileEmbeddingsResponseUsageJSON) RawJSON() string {
	return r.raw
}

type UploadNewParams struct {
	// File to upload
	File param.Field[io.Reader] `json:"file,required" format:"binary"`
	// Enable OCR for image processing
	EnableOcr param.Field[UploadNewParamsEnableOcr] `json:"enable_ocr"`
	// Enable vision-based processing for images
	EnableVision param.Field[UploadNewParamsEnableVision] `json:"enable_vision"`
	// Extraction mode: 'default' or 'embeddings'
	ExtractMode param.Field[UploadNewParamsExtractMode] `json:"extract_mode"`
	// Extract only images from documents
	ImagesOnly param.Field[UploadNewParamsImagesOnly] `json:"images_only"`
	// Return page-based structured response
	PageBased param.Field[UploadNewParamsPageBased] `json:"page_based"`
	// Remove headers/footers from documents
	RemoveHeaders param.Field[UploadNewParamsRemoveHeaders] `json:"remove_headers"`
	// Save file to configured storage
	SaveAll param.Field[UploadNewParamsSaveAll] `json:"save_all"`
	// Extract only text content
	TextOnly param.Field[UploadNewParamsTextOnly] `json:"text_only"`
	// Process with vision only
	VisionOnly param.Field[UploadNewParamsVisionOnly] `json:"vision_only"`
}

func (r UploadNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
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

// Enable OCR for image processing
type UploadNewParamsEnableOcr string

const (
	UploadNewParamsEnableOcrTrue  UploadNewParamsEnableOcr = "true"
	UploadNewParamsEnableOcrFalse UploadNewParamsEnableOcr = "false"
)

func (r UploadNewParamsEnableOcr) IsKnown() bool {
	switch r {
	case UploadNewParamsEnableOcrTrue, UploadNewParamsEnableOcrFalse:
		return true
	}
	return false
}

// Enable vision-based processing for images
type UploadNewParamsEnableVision string

const (
	UploadNewParamsEnableVisionTrue  UploadNewParamsEnableVision = "true"
	UploadNewParamsEnableVisionFalse UploadNewParamsEnableVision = "false"
)

func (r UploadNewParamsEnableVision) IsKnown() bool {
	switch r {
	case UploadNewParamsEnableVisionTrue, UploadNewParamsEnableVisionFalse:
		return true
	}
	return false
}

// Extraction mode: 'default' or 'embeddings'
type UploadNewParamsExtractMode string

const (
	UploadNewParamsExtractModeDefault    UploadNewParamsExtractMode = "default"
	UploadNewParamsExtractModeEmbeddings UploadNewParamsExtractMode = "embeddings"
)

func (r UploadNewParamsExtractMode) IsKnown() bool {
	switch r {
	case UploadNewParamsExtractModeDefault, UploadNewParamsExtractModeEmbeddings:
		return true
	}
	return false
}

// Extract only images from documents
type UploadNewParamsImagesOnly string

const (
	UploadNewParamsImagesOnlyTrue  UploadNewParamsImagesOnly = "true"
	UploadNewParamsImagesOnlyFalse UploadNewParamsImagesOnly = "false"
)

func (r UploadNewParamsImagesOnly) IsKnown() bool {
	switch r {
	case UploadNewParamsImagesOnlyTrue, UploadNewParamsImagesOnlyFalse:
		return true
	}
	return false
}

// Return page-based structured response
type UploadNewParamsPageBased string

const (
	UploadNewParamsPageBasedTrue  UploadNewParamsPageBased = "true"
	UploadNewParamsPageBasedFalse UploadNewParamsPageBased = "false"
)

func (r UploadNewParamsPageBased) IsKnown() bool {
	switch r {
	case UploadNewParamsPageBasedTrue, UploadNewParamsPageBasedFalse:
		return true
	}
	return false
}

// Remove headers/footers from documents
type UploadNewParamsRemoveHeaders string

const (
	UploadNewParamsRemoveHeadersTrue  UploadNewParamsRemoveHeaders = "true"
	UploadNewParamsRemoveHeadersFalse UploadNewParamsRemoveHeaders = "false"
)

func (r UploadNewParamsRemoveHeaders) IsKnown() bool {
	switch r {
	case UploadNewParamsRemoveHeadersTrue, UploadNewParamsRemoveHeadersFalse:
		return true
	}
	return false
}

// Save file to configured storage
type UploadNewParamsSaveAll string

const (
	UploadNewParamsSaveAllTrue  UploadNewParamsSaveAll = "true"
	UploadNewParamsSaveAllFalse UploadNewParamsSaveAll = "false"
)

func (r UploadNewParamsSaveAll) IsKnown() bool {
	switch r {
	case UploadNewParamsSaveAllTrue, UploadNewParamsSaveAllFalse:
		return true
	}
	return false
}

// Extract only text content
type UploadNewParamsTextOnly string

const (
	UploadNewParamsTextOnlyTrue  UploadNewParamsTextOnly = "true"
	UploadNewParamsTextOnlyFalse UploadNewParamsTextOnly = "false"
)

func (r UploadNewParamsTextOnly) IsKnown() bool {
	switch r {
	case UploadNewParamsTextOnlyTrue, UploadNewParamsTextOnlyFalse:
		return true
	}
	return false
}

// Process with vision only
type UploadNewParamsVisionOnly string

const (
	UploadNewParamsVisionOnlyTrue  UploadNewParamsVisionOnly = "true"
	UploadNewParamsVisionOnlyFalse UploadNewParamsVisionOnly = "false"
)

func (r UploadNewParamsVisionOnly) IsKnown() bool {
	switch r {
	case UploadNewParamsVisionOnlyTrue, UploadNewParamsVisionOnlyFalse:
		return true
	}
	return false
}
