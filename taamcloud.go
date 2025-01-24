// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package taamcloud

import (
	"bytes"
	"io"
	"mime/multipart"

	"github.com/stainless-sdks/taam-cloud-go/internal/apiform"
	"github.com/stainless-sdks/taam-cloud-go/internal/apijson"
	"github.com/stainless-sdks/taam-cloud-go/internal/param"
)

type UploadResponse struct {
	Content string             `json:"content"`
	Error   string             `json:"error"`
	Status  bool               `json:"status"`
	Type    string             `json:"type"`
	JSON    uploadResponseJSON `json:"-"`
}

// uploadResponseJSON contains the JSON metadata for the struct [UploadResponse]
type uploadResponseJSON struct {
	Content     apijson.Field
	Error       apijson.Field
	Status      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uploadResponseJSON) RawJSON() string {
	return r.raw
}

type UploadParams struct {
	// File to upload
	File param.Field[io.Reader] `json:"file,required" format:"binary"`
	// Enable OCR processing
	EnableOcr param.Field[bool] `json:"enable_ocr"`
	// Enable Vision processing
	EnableVision param.Field[bool] `json:"enable_vision"`
	// Save raw files
	SaveAll param.Field[bool] `json:"save_all"`
}

func (r UploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
