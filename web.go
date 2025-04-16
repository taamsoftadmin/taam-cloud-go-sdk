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

// WebService contains methods and other services that help with interacting with
// the taam-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebService] method instead.
type WebService struct {
	Options []option.RequestOption
}

// NewWebService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebService(opts ...option.RequestOption) (r *WebService) {
	r = &WebService{}
	r.Options = opts
	return
}

// Unified endpoint for web scraping, crawling, mapping and AI search
func (r *WebService) New(ctx context.Context, body WebNewParams, opts ...option.RequestOption) (res *WebNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/web"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type WebNewResponse struct {
	// Unique identifier for the request
	ID string `json:"id"`
	// Unix timestamp of when the request was created
	Created int64 `json:"created"`
	// Model-specific response data
	Data interface{} `json:"data"`
	// Model used for the request
	Model string `json:"model"`
	// Type of completion (e.g., scrape.completion)
	Object            string              `json:"object"`
	SystemFingerprint string              `json:"system_fingerprint"`
	Usage             WebNewResponseUsage `json:"usage"`
	JSON              webNewResponseJSON  `json:"-"`
}

// webNewResponseJSON contains the JSON metadata for the struct [WebNewResponse]
type webNewResponseJSON struct {
	ID                apijson.Field
	Created           apijson.Field
	Data              apijson.Field
	Model             apijson.Field
	Object            apijson.Field
	SystemFingerprint apijson.Field
	Usage             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *WebNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webNewResponseJSON) RawJSON() string {
	return r.raw
}

type WebNewResponseUsage struct {
	CompletionTokens int64                   `json:"completion_tokens"`
	PromptTokens     int64                   `json:"prompt_tokens"`
	TotalTokens      int64                   `json:"total_tokens"`
	JSON             webNewResponseUsageJSON `json:"-"`
}

// webNewResponseUsageJSON contains the JSON metadata for the struct
// [WebNewResponseUsage]
type webNewResponseUsageJSON struct {
	CompletionTokens apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *WebNewResponseUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webNewResponseUsageJSON) RawJSON() string {
	return r.raw
}

type WebNewParams struct {
	// Type of web service to use
	Model param.Field[WebNewParamsModel] `json:"model,required"`
	// Parameters specific to the selected model
	Params param.Field[interface{}] `json:"params"`
}

func (r WebNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of web service to use
type WebNewParamsModel string

const (
	WebNewParamsModelScrape       WebNewParamsModel = "scrape"
	WebNewParamsModelCrawl        WebNewParamsModel = "crawl"
	WebNewParamsModelMap          WebNewParamsModel = "map"
	WebNewParamsModelTaamAISearch WebNewParamsModel = "taam-ai-search"
	WebNewParamsModelCrawlStatus  WebNewParamsModel = "crawl-status"
)

func (r WebNewParamsModel) IsKnown() bool {
	switch r {
	case WebNewParamsModelScrape, WebNewParamsModelCrawl, WebNewParamsModelMap, WebNewParamsModelTaamAISearch, WebNewParamsModelCrawlStatus:
		return true
	}
	return false
}
