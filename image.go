// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package taamcloud

import (
	"context"
	"net/http"

	"github.com/taamsoftadmin/taam-cloud-go-sdk/internal/apijson"
	"github.com/taamsoftadmin/taam-cloud-go-sdk/internal/param"
	"github.com/taamsoftadmin/taam-cloud-go-sdk/internal/requestconfig"
	"github.com/taamsoftadmin/taam-cloud-go-sdk/option"
)

// ImageService contains methods and other services that help with interacting with
// the taam-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewImageService] method instead.
type ImageService struct {
	Options []option.RequestOption
}

// NewImageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewImageService(opts ...option.RequestOption) (r *ImageService) {
	r = &ImageService{}
	r.Options = opts
	return
}

// Create images from text descriptions
func (r *ImageService) Generate(ctx context.Context, body ImageGenerateParams, opts ...option.RequestOption) (res *ImageGenerateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/images/generations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ImageGenerateResponse struct {
	Created int64                       `json:"created"`
	Data    []ImageGenerateResponseData `json:"data"`
	JSON    imageGenerateResponseJSON   `json:"-"`
}

// imageGenerateResponseJSON contains the JSON metadata for the struct
// [ImageGenerateResponse]
type imageGenerateResponseJSON struct {
	Created     apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageGenerateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageGenerateResponseJSON) RawJSON() string {
	return r.raw
}

type ImageGenerateResponseData struct {
	RevisedPrompt string                        `json:"revised_prompt"`
	URL           string                        `json:"url"`
	JSON          imageGenerateResponseDataJSON `json:"-"`
}

// imageGenerateResponseDataJSON contains the JSON metadata for the struct
// [ImageGenerateResponseData]
type imageGenerateResponseDataJSON struct {
	RevisedPrompt apijson.Field
	URL           apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ImageGenerateResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageGenerateResponseDataJSON) RawJSON() string {
	return r.raw
}

type ImageGenerateParams struct {
	// Text description of the desired image
	Prompt param.Field[string]                   `json:"prompt,required"`
	Model  param.Field[ImageGenerateParamsModel] `json:"model"`
	// Number of images to generate
	N       param.Field[int64]                      `json:"n"`
	Quality param.Field[ImageGenerateParamsQuality] `json:"quality"`
	Size    param.Field[ImageGenerateParamsSize]    `json:"size"`
	Style   param.Field[ImageGenerateParamsStyle]   `json:"style"`
}

func (r ImageGenerateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ImageGenerateParamsModel string

const (
	ImageGenerateParamsModelDallE3 ImageGenerateParamsModel = "dall-e-3"
)

func (r ImageGenerateParamsModel) IsKnown() bool {
	switch r {
	case ImageGenerateParamsModelDallE3:
		return true
	}
	return false
}

type ImageGenerateParamsQuality string

const (
	ImageGenerateParamsQualityStandard ImageGenerateParamsQuality = "standard"
	ImageGenerateParamsQualityHD       ImageGenerateParamsQuality = "hd"
)

func (r ImageGenerateParamsQuality) IsKnown() bool {
	switch r {
	case ImageGenerateParamsQualityStandard, ImageGenerateParamsQualityHD:
		return true
	}
	return false
}

type ImageGenerateParamsSize string

const (
	ImageGenerateParamsSize1024x1024 ImageGenerateParamsSize = "1024x1024"
	ImageGenerateParamsSize1024x1792 ImageGenerateParamsSize = "1024x1792"
	ImageGenerateParamsSize1792x1024 ImageGenerateParamsSize = "1792x1024"
)

func (r ImageGenerateParamsSize) IsKnown() bool {
	switch r {
	case ImageGenerateParamsSize1024x1024, ImageGenerateParamsSize1024x1792, ImageGenerateParamsSize1792x1024:
		return true
	}
	return false
}

type ImageGenerateParamsStyle string

const (
	ImageGenerateParamsStyleNatural ImageGenerateParamsStyle = "natural"
	ImageGenerateParamsStyleVivid   ImageGenerateParamsStyle = "vivid"
)

func (r ImageGenerateParamsStyle) IsKnown() bool {
	switch r {
	case ImageGenerateParamsStyleNatural, ImageGenerateParamsStyleVivid:
		return true
	}
	return false
}
