// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package taamcloud

import (
	"context"
	"net/http"
	"net/url"

	"github.com/taamsoftadmin/taam-cloud-go-sdk/internal/apijson"
	"github.com/taamsoftadmin/taam-cloud-go-sdk/internal/apiquery"
	"github.com/taamsoftadmin/taam-cloud-go-sdk/internal/param"
	"github.com/taamsoftadmin/taam-cloud-go-sdk/internal/requestconfig"
	"github.com/taamsoftadmin/taam-cloud-go-sdk/option"
)

// QueryService contains methods and other services that help with interacting with
// the taam-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQueryService] method instead.
type QueryService struct {
	Options []option.RequestOption
}

// NewQueryService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewQueryService(opts ...option.RequestOption) (r *QueryService) {
	r = &QueryService{}
	r.Options = opts
	return
}

// Query the status of a video generation task
func (r *QueryService) CheckVideoGenerationStatus(ctx context.Context, query QueryCheckVideoGenerationStatusParams, opts ...option.RequestOption) (res *QueryCheckVideoGenerationStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/query/video_generation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type QueryCheckVideoGenerationStatusResponse struct {
	BaseResp QueryCheckVideoGenerationStatusResponseBaseResp `json:"base_resp"`
	// File identifier for the generated video (only present when status is 'Success')
	FileID string `json:"file_id"`
	// Current status of the task
	Status QueryCheckVideoGenerationStatusResponseStatus `json:"status"`
	// Task identifier
	TaskID string                                      `json:"task_id"`
	JSON   queryCheckVideoGenerationStatusResponseJSON `json:"-"`
}

// queryCheckVideoGenerationStatusResponseJSON contains the JSON metadata for the
// struct [QueryCheckVideoGenerationStatusResponse]
type queryCheckVideoGenerationStatusResponseJSON struct {
	BaseResp    apijson.Field
	FileID      apijson.Field
	Status      apijson.Field
	TaskID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueryCheckVideoGenerationStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queryCheckVideoGenerationStatusResponseJSON) RawJSON() string {
	return r.raw
}

type QueryCheckVideoGenerationStatusResponseBaseResp struct {
	// Status code (0 for success)
	StatusCode int64 `json:"status_code"`
	// Status message
	StatusMsg string                                              `json:"status_msg"`
	JSON      queryCheckVideoGenerationStatusResponseBaseRespJSON `json:"-"`
}

// queryCheckVideoGenerationStatusResponseBaseRespJSON contains the JSON metadata
// for the struct [QueryCheckVideoGenerationStatusResponseBaseResp]
type queryCheckVideoGenerationStatusResponseBaseRespJSON struct {
	StatusCode  apijson.Field
	StatusMsg   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueryCheckVideoGenerationStatusResponseBaseResp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queryCheckVideoGenerationStatusResponseBaseRespJSON) RawJSON() string {
	return r.raw
}

// Current status of the task
type QueryCheckVideoGenerationStatusResponseStatus string

const (
	QueryCheckVideoGenerationStatusResponseStatusQueueing   QueryCheckVideoGenerationStatusResponseStatus = "Queueing"
	QueryCheckVideoGenerationStatusResponseStatusPreparing  QueryCheckVideoGenerationStatusResponseStatus = "Preparing"
	QueryCheckVideoGenerationStatusResponseStatusProcessing QueryCheckVideoGenerationStatusResponseStatus = "Processing"
	QueryCheckVideoGenerationStatusResponseStatusSuccess    QueryCheckVideoGenerationStatusResponseStatus = "Success"
	QueryCheckVideoGenerationStatusResponseStatusFail       QueryCheckVideoGenerationStatusResponseStatus = "Fail"
)

func (r QueryCheckVideoGenerationStatusResponseStatus) IsKnown() bool {
	switch r {
	case QueryCheckVideoGenerationStatusResponseStatusQueueing, QueryCheckVideoGenerationStatusResponseStatusPreparing, QueryCheckVideoGenerationStatusResponseStatusProcessing, QueryCheckVideoGenerationStatusResponseStatusSuccess, QueryCheckVideoGenerationStatusResponseStatusFail:
		return true
	}
	return false
}

type QueryCheckVideoGenerationStatusParams struct {
	// Task ID returned from video generation request
	TaskID param.Field[string] `query:"task_id,required"`
}

// URLQuery serializes [QueryCheckVideoGenerationStatusParams]'s query parameters
// as `url.Values`.
func (r QueryCheckVideoGenerationStatusParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
