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

// SunoSubmitService contains methods and other services that help with interacting
// with the taam-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSunoSubmitService] method instead.
type SunoSubmitService struct {
	Options []option.RequestOption
}

// NewSunoSubmitService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSunoSubmitService(opts ...option.RequestOption) (r *SunoSubmitService) {
	r = &SunoSubmitService{}
	r.Options = opts
	return
}

// Create AI-generated music based on text prompts
func (r *SunoSubmitService) GenerateMusic(ctx context.Context, body SunoSubmitGenerateMusicParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "suno/submit/music"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type SunoSubmitGenerateMusicParams struct {
	Mv     param.Field[string] `json:"mv"`
	Prompt param.Field[string] `json:"prompt"`
	Tags   param.Field[string] `json:"tags"`
	Title  param.Field[string] `json:"title"`
}

func (r SunoSubmitGenerateMusicParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
