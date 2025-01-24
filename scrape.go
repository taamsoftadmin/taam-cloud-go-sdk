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

// ScrapeService contains methods and other services that help with interacting
// with the taam-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScrapeService] method instead.
type ScrapeService struct {
	Options []option.RequestOption
}

// NewScrapeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewScrapeService(opts ...option.RequestOption) (r *ScrapeService) {
	r = &ScrapeService{}
	r.Options = opts
	return
}

// Extract content from a webpage with customizable options
func (r *ScrapeService) New(ctx context.Context, body ScrapeNewParams, opts ...option.RequestOption) (res *ScrapeResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/scrape"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ScrapeResponse struct {
	Data    ScrapeResponseData `json:"data"`
	Success bool               `json:"success"`
	JSON    scrapeResponseJSON `json:"-"`
}

// scrapeResponseJSON contains the JSON metadata for the struct [ScrapeResponse]
type scrapeResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScrapeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scrapeResponseJSON) RawJSON() string {
	return r.raw
}

type ScrapeResponseData struct {
	Actions       ScrapeResponseDataActions  `json:"actions"`
	HTML          string                     `json:"html"`
	Links         []string                   `json:"links"`
	LlmExtraction interface{}                `json:"llm_extraction"`
	Markdown      string                     `json:"markdown"`
	Metadata      ScrapeResponseDataMetadata `json:"metadata"`
	RawHTML       string                     `json:"rawHtml"`
	Screenshot    string                     `json:"screenshot"`
	Warning       string                     `json:"warning"`
	JSON          scrapeResponseDataJSON     `json:"-"`
}

// scrapeResponseDataJSON contains the JSON metadata for the struct
// [ScrapeResponseData]
type scrapeResponseDataJSON struct {
	Actions       apijson.Field
	HTML          apijson.Field
	Links         apijson.Field
	LlmExtraction apijson.Field
	Markdown      apijson.Field
	Metadata      apijson.Field
	RawHTML       apijson.Field
	Screenshot    apijson.Field
	Warning       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScrapeResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scrapeResponseDataJSON) RawJSON() string {
	return r.raw
}

type ScrapeResponseDataActions struct {
	Screenshots []string                      `json:"screenshots"`
	JSON        scrapeResponseDataActionsJSON `json:"-"`
}

// scrapeResponseDataActionsJSON contains the JSON metadata for the struct
// [ScrapeResponseDataActions]
type scrapeResponseDataActionsJSON struct {
	Screenshots apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScrapeResponseDataActions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scrapeResponseDataActionsJSON) RawJSON() string {
	return r.raw
}

type ScrapeResponseDataMetadata struct {
	Description string                         `json:"description"`
	Error       string                         `json:"error"`
	Language    string                         `json:"language"`
	SourceURL   string                         `json:"sourceURL"`
	StatusCode  int64                          `json:"statusCode"`
	Title       string                         `json:"title"`
	JSON        scrapeResponseDataMetadataJSON `json:"-"`
}

// scrapeResponseDataMetadataJSON contains the JSON metadata for the struct
// [ScrapeResponseDataMetadata]
type scrapeResponseDataMetadataJSON struct {
	Description apijson.Field
	Error       apijson.Field
	Language    apijson.Field
	SourceURL   apijson.Field
	StatusCode  apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScrapeResponseDataMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scrapeResponseDataMetadataJSON) RawJSON() string {
	return r.raw
}

type ScrapeNewParams struct {
	// The URL to scrape
	URL             param.Field[string]                  `json:"url,required"`
	Formats         param.Field[[]ScrapeNewParamsFormat] `json:"formats"`
	OnlyMainContent param.Field[bool]                    `json:"onlyMainContent"`
}

func (r ScrapeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScrapeNewParamsFormat string

const (
	ScrapeNewParamsFormatMarkdown           ScrapeNewParamsFormat = "markdown"
	ScrapeNewParamsFormatHTML               ScrapeNewParamsFormat = "html"
	ScrapeNewParamsFormatRawHTML            ScrapeNewParamsFormat = "rawHtml"
	ScrapeNewParamsFormatLinks              ScrapeNewParamsFormat = "links"
	ScrapeNewParamsFormatScreenshot         ScrapeNewParamsFormat = "screenshot"
	ScrapeNewParamsFormatScreenshotFullPage ScrapeNewParamsFormat = "screenshot@fullPage"
	ScrapeNewParamsFormatJson               ScrapeNewParamsFormat = "json"
)

func (r ScrapeNewParamsFormat) IsKnown() bool {
	switch r {
	case ScrapeNewParamsFormatMarkdown, ScrapeNewParamsFormatHTML, ScrapeNewParamsFormatRawHTML, ScrapeNewParamsFormatLinks, ScrapeNewParamsFormatScreenshot, ScrapeNewParamsFormatScreenshotFullPage, ScrapeNewParamsFormatJson:
		return true
	}
	return false
}
