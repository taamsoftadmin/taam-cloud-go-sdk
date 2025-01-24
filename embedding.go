// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package taamcloud

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/taam-cloud-go/internal/apijson"
	"github.com/stainless-sdks/taam-cloud-go/internal/param"
	"github.com/stainless-sdks/taam-cloud-go/internal/requestconfig"
	"github.com/stainless-sdks/taam-cloud-go/option"
)

// EmbeddingService contains methods and other services that help with interacting
// with the taam-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmbeddingService] method instead.
type EmbeddingService struct {
	Options []option.RequestOption
}

// NewEmbeddingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmbeddingService(opts ...option.RequestOption) (r *EmbeddingService) {
	r = &EmbeddingService{}
	r.Options = opts
	return
}

// Generate embeddings
func (r *EmbeddingService) New(ctx context.Context, body EmbeddingNewParams, opts ...option.RequestOption) (res *EmbeddingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/embeddings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EmbeddingsResponse = interface{}

type EmbeddingNewParams struct {
	Input param.Field[[]string] `json:"input,required"`
	Model param.Field[string]   `json:"model,required"`
}

func (r EmbeddingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
