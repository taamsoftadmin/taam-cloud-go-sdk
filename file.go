// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package taamcloud

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/taamsoftadmin/taam-cloud-go-sdk/internal/apiform"
	"github.com/taamsoftadmin/taam-cloud-go-sdk/internal/apijson"
	"github.com/taamsoftadmin/taam-cloud-go-sdk/internal/apiquery"
	"github.com/taamsoftadmin/taam-cloud-go-sdk/internal/param"
	"github.com/taamsoftadmin/taam-cloud-go-sdk/internal/requestconfig"
	"github.com/taamsoftadmin/taam-cloud-go-sdk/option"
)

// FileService contains methods and other services that help with interacting with
// the taam-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileService] method instead.
type FileService struct {
	Options []option.RequestOption
}

// NewFileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFileService(opts ...option.RequestOption) (r *FileService) {
	r = &FileService{}
	r.Options = opts
	return
}

// Download a generated video file
func (r *FileService) Get(ctx context.Context, query FileGetParams, opts ...option.RequestOption) (res *FileGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/files/retrieve"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Upload a file for processing with AI models
func (r *FileService) Upload(ctx context.Context, body FileUploadParams, opts ...option.RequestOption) (res *FileUploadResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/files"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type FileGetResponse struct {
	BaseResp FileGetResponseBaseResp `json:"base_resp"`
	File     FileGetResponseFile     `json:"file"`
	JSON     fileGetResponseJSON     `json:"-"`
}

// fileGetResponseJSON contains the JSON metadata for the struct [FileGetResponse]
type fileGetResponseJSON struct {
	BaseResp    apijson.Field
	File        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FileGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileGetResponseJSON) RawJSON() string {
	return r.raw
}

type FileGetResponseBaseResp struct {
	// Status code (0 for success)
	StatusCode int64 `json:"status_code"`
	// Status message
	StatusMsg string                      `json:"status_msg"`
	JSON      fileGetResponseBaseRespJSON `json:"-"`
}

// fileGetResponseBaseRespJSON contains the JSON metadata for the struct
// [FileGetResponseBaseResp]
type fileGetResponseBaseRespJSON struct {
	StatusCode  apijson.Field
	StatusMsg   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FileGetResponseBaseResp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileGetResponseBaseRespJSON) RawJSON() string {
	return r.raw
}

type FileGetResponseFile struct {
	// File size in bytes
	Bytes int64 `json:"bytes"`
	// Unix timestamp of when the file was created
	CreatedAt int64 `json:"created_at"`
	// URL to download the file
	DownloadURL string `json:"download_url"`
	// File identifier
	FileID string `json:"file_id"`
	// Name of the file
	Filename string `json:"filename"`
	// Purpose of the file
	Purpose string                  `json:"purpose"`
	JSON    fileGetResponseFileJSON `json:"-"`
}

// fileGetResponseFileJSON contains the JSON metadata for the struct
// [FileGetResponseFile]
type fileGetResponseFileJSON struct {
	Bytes       apijson.Field
	CreatedAt   apijson.Field
	DownloadURL apijson.Field
	FileID      apijson.Field
	Filename    apijson.Field
	Purpose     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FileGetResponseFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileGetResponseFileJSON) RawJSON() string {
	return r.raw
}

type FileUploadResponse struct {
	// Extracted text or data URL
	Content string `json:"content"`
	// Error message (if any)
	Error string `json:"error"`
	// Operation status
	Status bool `json:"status"`
	// Type of the processed file
	Type FileUploadResponseType `json:"type"`
	JSON fileUploadResponseJSON `json:"-"`
}

// fileUploadResponseJSON contains the JSON metadata for the struct
// [FileUploadResponse]
type fileUploadResponseJSON struct {
	Content     apijson.Field
	Error       apijson.Field
	Status      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FileUploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileUploadResponseJSON) RawJSON() string {
	return r.raw
}

// Type of the processed file
type FileUploadResponseType string

const (
	FileUploadResponseTypePdf   FileUploadResponseType = "pdf"
	FileUploadResponseTypeDocx  FileUploadResponseType = "docx"
	FileUploadResponseTypePptx  FileUploadResponseType = "pptx"
	FileUploadResponseTypeXlsx  FileUploadResponseType = "xlsx"
	FileUploadResponseTypeImage FileUploadResponseType = "image"
	FileUploadResponseTypeAudio FileUploadResponseType = "audio"
	FileUploadResponseTypeText  FileUploadResponseType = "text"
	FileUploadResponseTypeError FileUploadResponseType = "error"
)

func (r FileUploadResponseType) IsKnown() bool {
	switch r {
	case FileUploadResponseTypePdf, FileUploadResponseTypeDocx, FileUploadResponseTypePptx, FileUploadResponseTypeXlsx, FileUploadResponseTypeImage, FileUploadResponseTypeAudio, FileUploadResponseTypeText, FileUploadResponseTypeError:
		return true
	}
	return false
}

type FileGetParams struct {
	// File ID returned from successful video generation
	FileID param.Field[string] `query:"file_id,required"`
}

// URLQuery serializes [FileGetParams]'s query parameters as `url.Values`.
func (r FileGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FileUploadParams struct {
	// File to upload
	File param.Field[io.Reader] `json:"file,required" format:"binary"`
	// Enable OCR for image processing
	EnableOcr param.Field[FileUploadParamsEnableOcr] `json:"enable_ocr"`
	// Enable vision-based processing for images
	EnableVision param.Field[FileUploadParamsEnableVision] `json:"enable_vision"`
	// Extraction mode: 'default' or 'embeddings'
	ExtractMode param.Field[FileUploadParamsExtractMode] `json:"extract_mode"`
	// Extract only images from documents
	ImagesOnly param.Field[FileUploadParamsImagesOnly] `json:"images_only"`
	// Return page-based structured response
	PageBased param.Field[FileUploadParamsPageBased] `json:"page_based"`
	// Remove headers/footers from documents
	RemoveHeaders param.Field[FileUploadParamsRemoveHeaders] `json:"remove_headers"`
	// Save file to configured storage
	SaveAll param.Field[FileUploadParamsSaveAll] `json:"save_all"`
	// Extract only text content
	TextOnly param.Field[FileUploadParamsTextOnly] `json:"text_only"`
	// Process with vision only
	VisionOnly param.Field[FileUploadParamsVisionOnly] `json:"vision_only"`
}

func (r FileUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
type FileUploadParamsEnableOcr string

const (
	FileUploadParamsEnableOcrTrue  FileUploadParamsEnableOcr = "true"
	FileUploadParamsEnableOcrFalse FileUploadParamsEnableOcr = "false"
)

func (r FileUploadParamsEnableOcr) IsKnown() bool {
	switch r {
	case FileUploadParamsEnableOcrTrue, FileUploadParamsEnableOcrFalse:
		return true
	}
	return false
}

// Enable vision-based processing for images
type FileUploadParamsEnableVision string

const (
	FileUploadParamsEnableVisionTrue  FileUploadParamsEnableVision = "true"
	FileUploadParamsEnableVisionFalse FileUploadParamsEnableVision = "false"
)

func (r FileUploadParamsEnableVision) IsKnown() bool {
	switch r {
	case FileUploadParamsEnableVisionTrue, FileUploadParamsEnableVisionFalse:
		return true
	}
	return false
}

// Extraction mode: 'default' or 'embeddings'
type FileUploadParamsExtractMode string

const (
	FileUploadParamsExtractModeDefault    FileUploadParamsExtractMode = "default"
	FileUploadParamsExtractModeEmbeddings FileUploadParamsExtractMode = "embeddings"
)

func (r FileUploadParamsExtractMode) IsKnown() bool {
	switch r {
	case FileUploadParamsExtractModeDefault, FileUploadParamsExtractModeEmbeddings:
		return true
	}
	return false
}

// Extract only images from documents
type FileUploadParamsImagesOnly string

const (
	FileUploadParamsImagesOnlyTrue  FileUploadParamsImagesOnly = "true"
	FileUploadParamsImagesOnlyFalse FileUploadParamsImagesOnly = "false"
)

func (r FileUploadParamsImagesOnly) IsKnown() bool {
	switch r {
	case FileUploadParamsImagesOnlyTrue, FileUploadParamsImagesOnlyFalse:
		return true
	}
	return false
}

// Return page-based structured response
type FileUploadParamsPageBased string

const (
	FileUploadParamsPageBasedTrue  FileUploadParamsPageBased = "true"
	FileUploadParamsPageBasedFalse FileUploadParamsPageBased = "false"
)

func (r FileUploadParamsPageBased) IsKnown() bool {
	switch r {
	case FileUploadParamsPageBasedTrue, FileUploadParamsPageBasedFalse:
		return true
	}
	return false
}

// Remove headers/footers from documents
type FileUploadParamsRemoveHeaders string

const (
	FileUploadParamsRemoveHeadersTrue  FileUploadParamsRemoveHeaders = "true"
	FileUploadParamsRemoveHeadersFalse FileUploadParamsRemoveHeaders = "false"
)

func (r FileUploadParamsRemoveHeaders) IsKnown() bool {
	switch r {
	case FileUploadParamsRemoveHeadersTrue, FileUploadParamsRemoveHeadersFalse:
		return true
	}
	return false
}

// Save file to configured storage
type FileUploadParamsSaveAll string

const (
	FileUploadParamsSaveAllTrue  FileUploadParamsSaveAll = "true"
	FileUploadParamsSaveAllFalse FileUploadParamsSaveAll = "false"
)

func (r FileUploadParamsSaveAll) IsKnown() bool {
	switch r {
	case FileUploadParamsSaveAllTrue, FileUploadParamsSaveAllFalse:
		return true
	}
	return false
}

// Extract only text content
type FileUploadParamsTextOnly string

const (
	FileUploadParamsTextOnlyTrue  FileUploadParamsTextOnly = "true"
	FileUploadParamsTextOnlyFalse FileUploadParamsTextOnly = "false"
)

func (r FileUploadParamsTextOnly) IsKnown() bool {
	switch r {
	case FileUploadParamsTextOnlyTrue, FileUploadParamsTextOnlyFalse:
		return true
	}
	return false
}

// Process with vision only
type FileUploadParamsVisionOnly string

const (
	FileUploadParamsVisionOnlyTrue  FileUploadParamsVisionOnly = "true"
	FileUploadParamsVisionOnlyFalse FileUploadParamsVisionOnly = "false"
)

func (r FileUploadParamsVisionOnly) IsKnown() bool {
	switch r {
	case FileUploadParamsVisionOnlyTrue, FileUploadParamsVisionOnlyFalse:
		return true
	}
	return false
}
