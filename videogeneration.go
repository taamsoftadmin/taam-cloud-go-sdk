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

// VideoGenerationService contains methods and other services that help with
// interacting with the taam-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVideoGenerationService] method instead.
type VideoGenerationService struct {
	Options []option.RequestOption
}

// NewVideoGenerationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVideoGenerationService(opts ...option.RequestOption) (r *VideoGenerationService) {
	r = &VideoGenerationService{}
	r.Options = opts
	return
}

// Create dynamic videos from text descriptions or images
func (r *VideoGenerationService) New(ctx context.Context, body VideoGenerationNewParams, opts ...option.RequestOption) (res *VideoGenerationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/video_generation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type VideoGenerationNewResponse struct {
	BaseResp VideoGenerationNewResponseBaseResp `json:"base_resp"`
	// Unique identifier for the generation task
	TaskID string                         `json:"task_id"`
	JSON   videoGenerationNewResponseJSON `json:"-"`
}

// videoGenerationNewResponseJSON contains the JSON metadata for the struct
// [VideoGenerationNewResponse]
type videoGenerationNewResponseJSON struct {
	BaseResp    apijson.Field
	TaskID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VideoGenerationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r videoGenerationNewResponseJSON) RawJSON() string {
	return r.raw
}

type VideoGenerationNewResponseBaseResp struct {
	// Status code (0 for success)
	StatusCode int64 `json:"status_code"`
	// Status message
	StatusMsg string                                 `json:"status_msg"`
	JSON      videoGenerationNewResponseBaseRespJSON `json:"-"`
}

// videoGenerationNewResponseBaseRespJSON contains the JSON metadata for the struct
// [VideoGenerationNewResponseBaseResp]
type videoGenerationNewResponseBaseRespJSON struct {
	StatusCode  apijson.Field
	StatusMsg   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VideoGenerationNewResponseBaseResp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r videoGenerationNewResponseBaseRespJSON) RawJSON() string {
	return r.raw
}

type VideoGenerationNewParams struct {
	// Video generation model to use
	Model param.Field[VideoGenerationNewParamsModel] `json:"model,required"`
	// Text description of the video to generate. Can include camera movement
	// instructions in square brackets.
	Prompt param.Field[string] `json:"prompt,required"`
	// Base64-encoded image data for image-to-video generation
	FirstFrameImage param.Field[string] `json:"first_frame_image"`
}

func (r VideoGenerationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Video generation model to use
type VideoGenerationNewParamsModel string

const (
	VideoGenerationNewParamsModelT2V01Director VideoGenerationNewParamsModel = "T2V-01-Director"
	VideoGenerationNewParamsModelI2V01Director VideoGenerationNewParamsModel = "I2V-01-Director"
	VideoGenerationNewParamsModelS2V01         VideoGenerationNewParamsModel = "S2V-01"
	VideoGenerationNewParamsModelI2V01         VideoGenerationNewParamsModel = "I2V-01"
	VideoGenerationNewParamsModelI2V01Live     VideoGenerationNewParamsModel = "I2V-01-live"
	VideoGenerationNewParamsModelT2V01         VideoGenerationNewParamsModel = "T2V-01"
)

func (r VideoGenerationNewParamsModel) IsKnown() bool {
	switch r {
	case VideoGenerationNewParamsModelT2V01Director, VideoGenerationNewParamsModelI2V01Director, VideoGenerationNewParamsModelS2V01, VideoGenerationNewParamsModelI2V01, VideoGenerationNewParamsModelI2V01Live, VideoGenerationNewParamsModelT2V01:
		return true
	}
	return false
}
