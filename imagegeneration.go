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

// ImageGenerationService contains methods and other services that help with
// interacting with the taam-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewImageGenerationService] method instead.
type ImageGenerationService struct {
	Options []option.RequestOption
}

// NewImageGenerationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewImageGenerationService(opts ...option.RequestOption) (r *ImageGenerationService) {
	r = &ImageGenerationService{}
	r.Options = opts
	return
}

// Create images from text descriptions
func (r *ImageGenerationService) New(ctx context.Context, body ImageGenerationNewParams, opts ...option.RequestOption) (res *ImageGenerationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/images/generations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ImageGenerationResponse struct {
	Created int64                         `json:"created"`
	Data    []ImageGenerationResponseData `json:"data"`
	JSON    imageGenerationResponseJSON   `json:"-"`
}

// imageGenerationResponseJSON contains the JSON metadata for the struct
// [ImageGenerationResponse]
type imageGenerationResponseJSON struct {
	Created     apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageGenerationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageGenerationResponseJSON) RawJSON() string {
	return r.raw
}

type ImageGenerationResponseData struct {
	RevisedPrompt string                          `json:"revised_prompt"`
	URL           string                          `json:"url"`
	JSON          imageGenerationResponseDataJSON `json:"-"`
}

// imageGenerationResponseDataJSON contains the JSON metadata for the struct
// [ImageGenerationResponseData]
type imageGenerationResponseDataJSON struct {
	RevisedPrompt apijson.Field
	URL           apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ImageGenerationResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageGenerationResponseDataJSON) RawJSON() string {
	return r.raw
}

type ImageGenerationNewParams struct {
	// Text description of the desired image
	Prompt param.Field[string]                        `json:"prompt,required"`
	Model  param.Field[ImageGenerationNewParamsModel] `json:"model"`
	// Number of images to generate
	N       param.Field[int64]                           `json:"n"`
	Quality param.Field[ImageGenerationNewParamsQuality] `json:"quality"`
	Size    param.Field[ImageGenerationNewParamsSize]    `json:"size"`
	Style   param.Field[ImageGenerationNewParamsStyle]   `json:"style"`
}

func (r ImageGenerationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ImageGenerationNewParamsModel string

const (
	ImageGenerationNewParamsModelDallE3 ImageGenerationNewParamsModel = "dall-e-3"
)

func (r ImageGenerationNewParamsModel) IsKnown() bool {
	switch r {
	case ImageGenerationNewParamsModelDallE3:
		return true
	}
	return false
}

type ImageGenerationNewParamsQuality string

const (
	ImageGenerationNewParamsQualityStandard ImageGenerationNewParamsQuality = "standard"
	ImageGenerationNewParamsQualityHD       ImageGenerationNewParamsQuality = "hd"
)

func (r ImageGenerationNewParamsQuality) IsKnown() bool {
	switch r {
	case ImageGenerationNewParamsQualityStandard, ImageGenerationNewParamsQualityHD:
		return true
	}
	return false
}

type ImageGenerationNewParamsSize string

const (
	ImageGenerationNewParamsSize1024x1024 ImageGenerationNewParamsSize = "1024x1024"
	ImageGenerationNewParamsSize1024x1792 ImageGenerationNewParamsSize = "1024x1792"
	ImageGenerationNewParamsSize1792x1024 ImageGenerationNewParamsSize = "1792x1024"
)

func (r ImageGenerationNewParamsSize) IsKnown() bool {
	switch r {
	case ImageGenerationNewParamsSize1024x1024, ImageGenerationNewParamsSize1024x1792, ImageGenerationNewParamsSize1792x1024:
		return true
	}
	return false
}

type ImageGenerationNewParamsStyle string

const (
	ImageGenerationNewParamsStyleNatural ImageGenerationNewParamsStyle = "natural"
	ImageGenerationNewParamsStyleVivid   ImageGenerationNewParamsStyle = "vivid"
)

func (r ImageGenerationNewParamsStyle) IsKnown() bool {
	switch r {
	case ImageGenerationNewParamsStyleNatural, ImageGenerationNewParamsStyleVivid:
		return true
	}
	return false
}
