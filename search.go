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

// SearchService contains methods and other services that help with interacting
// with the taam-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSearchService] method instead.
type SearchService struct {
	Options []option.RequestOption
}

// NewSearchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSearchService(opts ...option.RequestOption) (r *SearchService) {
	r = &SearchService{}
	r.Options = opts
	return
}

// Perform intelligent AI-powered searches
func (r *SearchService) Perform(ctx context.Context, body SearchPerformParams, opts ...option.RequestOption) (res *SearchResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SearchResponse struct {
	Message string                 `json:"message"`
	Sources []SearchResponseSource `json:"sources"`
	JSON    searchResponseJSON     `json:"-"`
}

// searchResponseJSON contains the JSON metadata for the struct [SearchResponse]
type searchResponseJSON struct {
	Message     apijson.Field
	Sources     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchResponseJSON) RawJSON() string {
	return r.raw
}

type SearchResponseSource struct {
	Metadata    SearchResponseSourcesMetadata `json:"metadata"`
	PageContent string                        `json:"pageContent"`
	JSON        searchResponseSourceJSON      `json:"-"`
}

// searchResponseSourceJSON contains the JSON metadata for the struct
// [SearchResponseSource]
type searchResponseSourceJSON struct {
	Metadata    apijson.Field
	PageContent apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SearchResponseSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchResponseSourceJSON) RawJSON() string {
	return r.raw
}

type SearchResponseSourcesMetadata struct {
	Title string                            `json:"title"`
	URL   string                            `json:"url"`
	JSON  searchResponseSourcesMetadataJSON `json:"-"`
}

// searchResponseSourcesMetadataJSON contains the JSON metadata for the struct
// [SearchResponseSourcesMetadata]
type searchResponseSourcesMetadataJSON struct {
	Title       apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SearchResponseSourcesMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchResponseSourcesMetadataJSON) RawJSON() string {
	return r.raw
}

type SearchPerformParams struct {
	ChatModel      param.Field[SearchPerformParamsChatModel]      `json:"chatModel,required"`
	EmbeddingModel param.Field[SearchPerformParamsEmbeddingModel] `json:"embeddingModel,required"`
	// Specifies which focus mode to use for search
	FocusMode param.Field[SearchPerformParamsFocusMode] `json:"focusMode,required"`
}

func (r SearchPerformParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SearchPerformParamsChatModel struct {
	Model               param.Field[SearchPerformParamsChatModelModel]    `json:"model,required"`
	Provider            param.Field[SearchPerformParamsChatModelProvider] `json:"provider,required"`
	CustomOpenAIBaseURL param.Field[string]                               `json:"customOpenAIBaseURL"`
	CustomOpenAIKey     param.Field[string]                               `json:"customOpenAIKey"`
}

func (r SearchPerformParamsChatModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SearchPerformParamsChatModelModel string

const (
	SearchPerformParamsChatModelModelGpt3_5Turbo SearchPerformParamsChatModelModel = "gpt-3.5-turbo"
	SearchPerformParamsChatModelModelGpt4        SearchPerformParamsChatModelModel = "gpt-4"
	SearchPerformParamsChatModelModelGpt4o       SearchPerformParamsChatModelModel = "gpt-4o"
	SearchPerformParamsChatModelModelGpt4oMini   SearchPerformParamsChatModelModel = "gpt-4o-mini"
)

func (r SearchPerformParamsChatModelModel) IsKnown() bool {
	switch r {
	case SearchPerformParamsChatModelModelGpt3_5Turbo, SearchPerformParamsChatModelModelGpt4, SearchPerformParamsChatModelModelGpt4o, SearchPerformParamsChatModelModelGpt4oMini:
		return true
	}
	return false
}

type SearchPerformParamsChatModelProvider string

const (
	SearchPerformParamsChatModelProviderCustomOpenAI SearchPerformParamsChatModelProvider = "custom_openai"
)

func (r SearchPerformParamsChatModelProvider) IsKnown() bool {
	switch r {
	case SearchPerformParamsChatModelProviderCustomOpenAI:
		return true
	}
	return false
}

type SearchPerformParamsEmbeddingModel struct {
	Model    param.Field[SearchPerformParamsEmbeddingModelModel]    `json:"model,required"`
	Provider param.Field[SearchPerformParamsEmbeddingModelProvider] `json:"provider,required"`
}

func (r SearchPerformParamsEmbeddingModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SearchPerformParamsEmbeddingModelModel string

const (
	SearchPerformParamsEmbeddingModelModelTextEmbedding3Large      SearchPerformParamsEmbeddingModelModel = "text-embedding-3-large"
	SearchPerformParamsEmbeddingModelModelTextEmbedding3Small      SearchPerformParamsEmbeddingModelModel = "text-embedding-3-small"
	SearchPerformParamsEmbeddingModelModelTextEmbeddingAda002      SearchPerformParamsEmbeddingModelModel = "text-embedding-ada-002"
	SearchPerformParamsEmbeddingModelModelJinaEmbeddingsV3         SearchPerformParamsEmbeddingModelModel = "jina-embeddings-v3"
	SearchPerformParamsEmbeddingModelModelJinaEmbeddingsV2BaseEn   SearchPerformParamsEmbeddingModelModel = "jina-embeddings-v2-base-en"
	SearchPerformParamsEmbeddingModelModelJinaEmbeddingsV2BaseZh   SearchPerformParamsEmbeddingModelModel = "jina-embeddings-v2-base-zh"
	SearchPerformParamsEmbeddingModelModelJinaEmbeddingsV2BaseCode SearchPerformParamsEmbeddingModelModel = "jina-embeddings-v2-base-code"
)

func (r SearchPerformParamsEmbeddingModelModel) IsKnown() bool {
	switch r {
	case SearchPerformParamsEmbeddingModelModelTextEmbedding3Large, SearchPerformParamsEmbeddingModelModelTextEmbedding3Small, SearchPerformParamsEmbeddingModelModelTextEmbeddingAda002, SearchPerformParamsEmbeddingModelModelJinaEmbeddingsV3, SearchPerformParamsEmbeddingModelModelJinaEmbeddingsV2BaseEn, SearchPerformParamsEmbeddingModelModelJinaEmbeddingsV2BaseZh, SearchPerformParamsEmbeddingModelModelJinaEmbeddingsV2BaseCode:
		return true
	}
	return false
}

type SearchPerformParamsEmbeddingModelProvider string

const (
	SearchPerformParamsEmbeddingModelProviderCustomOpenAI SearchPerformParamsEmbeddingModelProvider = "custom_openai"
)

func (r SearchPerformParamsEmbeddingModelProvider) IsKnown() bool {
	switch r {
	case SearchPerformParamsEmbeddingModelProviderCustomOpenAI:
		return true
	}
	return false
}

// Specifies which focus mode to use for search
type SearchPerformParamsFocusMode string

const (
	SearchPerformParamsFocusModeWebSearch          SearchPerformParamsFocusMode = "webSearch"
	SearchPerformParamsFocusModeAcademicSearch     SearchPerformParamsFocusMode = "academicSearch"
	SearchPerformParamsFocusModeWritingAssistant   SearchPerformParamsFocusMode = "writingAssistant"
	SearchPerformParamsFocusModeWolframAlphaSearch SearchPerformParamsFocusMode = "wolframAlphaSearch"
	SearchPerformParamsFocusModeYoutubeSearch      SearchPerformParamsFocusMode = "youtubeSearch"
	SearchPerformParamsFocusModeRedditSearch       SearchPerformParamsFocusMode = "redditSearch"
)

func (r SearchPerformParamsFocusMode) IsKnown() bool {
	switch r {
	case SearchPerformParamsFocusModeWebSearch, SearchPerformParamsFocusModeAcademicSearch, SearchPerformParamsFocusModeWritingAssistant, SearchPerformParamsFocusModeWolframAlphaSearch, SearchPerformParamsFocusModeYoutubeSearch, SearchPerformParamsFocusModeRedditSearch:
		return true
	}
	return false
}
