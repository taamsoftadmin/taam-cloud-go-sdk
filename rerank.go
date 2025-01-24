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

// RerankService contains methods and other services that help with interacting
// with the taam-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRerankService] method instead.
type RerankService struct {
	Options []option.RequestOption
}

// NewRerankService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRerankService(opts ...option.RequestOption) (r *RerankService) {
	r = &RerankService{}
	r.Options = opts
	return
}

// Rerank documents
func (r *RerankService) New(ctx context.Context, body RerankNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/rerank"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type RerankNewParams struct {
	Documents param.Field[[]string] `json:"documents,required"`
	Model     param.Field[string]   `json:"model,required"`
	Query     param.Field[string]   `json:"query,required"`
	TopN      param.Field[int64]    `json:"top_n"`
}

func (r RerankNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
