// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package taamcloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/taam-cloud-go/internal/apijson"
	"github.com/stainless-sdks/taam-cloud-go/internal/param"
	"github.com/stainless-sdks/taam-cloud-go/internal/requestconfig"
	"github.com/stainless-sdks/taam-cloud-go/option"
)

// CrawlService contains methods and other services that help with interacting with
// the taam-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCrawlService] method instead.
type CrawlService struct {
	Options []option.RequestOption
}

// NewCrawlService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCrawlService(opts ...option.RequestOption) (r *CrawlService) {
	r = &CrawlService{}
	r.Options = opts
	return
}

// Crawl a website with customizable options
func (r *CrawlService) New(ctx context.Context, body CrawlNewParams, opts ...option.RequestOption) (res *CrawlResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/crawl"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve the status and results of a crawl job
func (r *CrawlService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *CrawlStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/crawl/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CrawlResponse struct {
	ID      string            `json:"id"`
	Success bool              `json:"success"`
	URL     string            `json:"url"`
	JSON    crawlResponseJSON `json:"-"`
}

// crawlResponseJSON contains the JSON metadata for the struct [CrawlResponse]
type crawlResponseJSON struct {
	ID          apijson.Field
	Success     apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CrawlResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r crawlResponseJSON) RawJSON() string {
	return r.raw
}

type CrawlStatusResponse struct {
	// Number of successfully crawled pages
	Completed int64 `json:"completed"`
	// Credits consumed by the crawl
	CreditsUsed int64                     `json:"creditsUsed"`
	Data        []CrawlStatusResponseData `json:"data"`
	// Expiration date of the crawl data
	ExpiresAt time.Time `json:"expiresAt" format:"date-time"`
	// URL to retrieve next batch of data
	Next string `json:"next,nullable"`
	// Current status of the crawl
	Status CrawlStatusResponseStatus `json:"status"`
	// Total number of pages attempted
	Total int64                   `json:"total"`
	JSON  crawlStatusResponseJSON `json:"-"`
}

// crawlStatusResponseJSON contains the JSON metadata for the struct
// [CrawlStatusResponse]
type crawlStatusResponseJSON struct {
	Completed   apijson.Field
	CreditsUsed apijson.Field
	Data        apijson.Field
	ExpiresAt   apijson.Field
	Next        apijson.Field
	Status      apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CrawlStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r crawlStatusResponseJSON) RawJSON() string {
	return r.raw
}

type CrawlStatusResponseData struct {
	HTML       string                          `json:"html"`
	Links      []string                        `json:"links"`
	Markdown   string                          `json:"markdown"`
	Metadata   CrawlStatusResponseDataMetadata `json:"metadata"`
	RawHTML    string                          `json:"rawHtml"`
	Screenshot string                          `json:"screenshot"`
	JSON       crawlStatusResponseDataJSON     `json:"-"`
}

// crawlStatusResponseDataJSON contains the JSON metadata for the struct
// [CrawlStatusResponseData]
type crawlStatusResponseDataJSON struct {
	HTML        apijson.Field
	Links       apijson.Field
	Markdown    apijson.Field
	Metadata    apijson.Field
	RawHTML     apijson.Field
	Screenshot  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CrawlStatusResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r crawlStatusResponseDataJSON) RawJSON() string {
	return r.raw
}

type CrawlStatusResponseDataMetadata struct {
	Description string                              `json:"description"`
	Error       string                              `json:"error"`
	Language    string                              `json:"language"`
	SourceURL   string                              `json:"sourceURL"`
	StatusCode  int64                               `json:"statusCode"`
	Title       string                              `json:"title"`
	ExtraFields map[string]interface{}              `json:"-,extras"`
	JSON        crawlStatusResponseDataMetadataJSON `json:"-"`
}

// crawlStatusResponseDataMetadataJSON contains the JSON metadata for the struct
// [CrawlStatusResponseDataMetadata]
type crawlStatusResponseDataMetadataJSON struct {
	Description apijson.Field
	Error       apijson.Field
	Language    apijson.Field
	SourceURL   apijson.Field
	StatusCode  apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CrawlStatusResponseDataMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r crawlStatusResponseDataMetadataJSON) RawJSON() string {
	return r.raw
}

// Current status of the crawl
type CrawlStatusResponseStatus string

const (
	CrawlStatusResponseStatusScraping  CrawlStatusResponseStatus = "scraping"
	CrawlStatusResponseStatusCompleted CrawlStatusResponseStatus = "completed"
	CrawlStatusResponseStatusFailed    CrawlStatusResponseStatus = "failed"
)

func (r CrawlStatusResponseStatus) IsKnown() bool {
	switch r {
	case CrawlStatusResponseStatusScraping, CrawlStatusResponseStatusCompleted, CrawlStatusResponseStatusFailed:
		return true
	}
	return false
}

type CrawlNewParams struct {
	// The base URL to start crawling from
	URL param.Field[string] `json:"url,required"`
	// Allow navigation to previously linked pages
	AllowBackwardLinks param.Field[bool] `json:"allowBackwardLinks"`
	// Allow following external links
	AllowExternalLinks param.Field[bool] `json:"allowExternalLinks"`
	// URL patterns to exclude from crawl
	ExcludePaths param.Field[[]string] `json:"excludePaths"`
	// Ignore website sitemap
	IgnoreSitemap param.Field[bool] `json:"ignoreSitemap"`
	// URL patterns to include in crawl
	IncludePaths param.Field[[]string] `json:"includePaths"`
	// Maximum number of pages to crawl
	Limit param.Field[int64] `json:"limit"`
	// Maximum depth to crawl
	MaxDepth      param.Field[int64]                       `json:"maxDepth"`
	ScrapeOptions param.Field[CrawlNewParamsScrapeOptions] `json:"scrapeOptions"`
	// Webhook URL for crawl events
	Webhook param.Field[string] `json:"webhook"`
}

func (r CrawlNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CrawlNewParamsScrapeOptions struct {
	Actions             param.Field[[]CrawlNewParamsScrapeOptionsActionUnion] `json:"actions"`
	ExcludeTags         param.Field[[]string]                                 `json:"excludeTags"`
	Formats             param.Field[[]CrawlNewParamsScrapeOptionsFormat]      `json:"formats"`
	Headers             param.Field[map[string]interface{}]                   `json:"headers"`
	IncludeTags         param.Field[[]string]                                 `json:"includeTags"`
	JsonOptions         param.Field[CrawlNewParamsScrapeOptionsJsonOptions]   `json:"jsonOptions"`
	Location            param.Field[CrawlNewParamsScrapeOptionsLocation]      `json:"location"`
	Mobile              param.Field[bool]                                     `json:"mobile"`
	OnlyMainContent     param.Field[bool]                                     `json:"onlyMainContent"`
	RemoveBase64Images  param.Field[bool]                                     `json:"removeBase64Images"`
	SkipTlsVerification param.Field[bool]                                     `json:"skipTlsVerification"`
	Timeout             param.Field[int64]                                    `json:"timeout"`
	WaitFor             param.Field[int64]                                    `json:"waitFor"`
}

func (r CrawlNewParamsScrapeOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CrawlNewParamsScrapeOptionsAction struct {
	Type         param.Field[CrawlNewParamsScrapeOptionsActionsType] `json:"type,required"`
	Milliseconds param.Field[int64]                                  `json:"milliseconds"`
	Selector     param.Field[string]                                 `json:"selector"`
	Text         param.Field[string]                                 `json:"text"`
}

func (r CrawlNewParamsScrapeOptionsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CrawlNewParamsScrapeOptionsAction) implementsCrawlNewParamsScrapeOptionsActionUnion() {}

// Satisfied by [CrawlNewParamsScrapeOptionsActionsObject],
// [CrawlNewParamsScrapeOptionsActionsObject], [CrawlNewParamsScrapeOptionsAction].
type CrawlNewParamsScrapeOptionsActionUnion interface {
	implementsCrawlNewParamsScrapeOptionsActionUnion()
}

type CrawlNewParamsScrapeOptionsActionsObject struct {
	Milliseconds param.Field[int64]                                        `json:"milliseconds,required"`
	Type         param.Field[CrawlNewParamsScrapeOptionsActionsObjectType] `json:"type,required"`
	Selector     param.Field[string]                                       `json:"selector"`
}

func (r CrawlNewParamsScrapeOptionsActionsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CrawlNewParamsScrapeOptionsActionsObject) implementsCrawlNewParamsScrapeOptionsActionUnion() {
}

type CrawlNewParamsScrapeOptionsActionsObjectType string

const (
	CrawlNewParamsScrapeOptionsActionsObjectTypeWait CrawlNewParamsScrapeOptionsActionsObjectType = "wait"
)

func (r CrawlNewParamsScrapeOptionsActionsObjectType) IsKnown() bool {
	switch r {
	case CrawlNewParamsScrapeOptionsActionsObjectTypeWait:
		return true
	}
	return false
}

type CrawlNewParamsScrapeOptionsActionsType string

const (
	CrawlNewParamsScrapeOptionsActionsTypeWait       CrawlNewParamsScrapeOptionsActionsType = "wait"
	CrawlNewParamsScrapeOptionsActionsTypeScreenshot CrawlNewParamsScrapeOptionsActionsType = "screenshot"
	CrawlNewParamsScrapeOptionsActionsTypeClick      CrawlNewParamsScrapeOptionsActionsType = "click"
	CrawlNewParamsScrapeOptionsActionsTypeWrite      CrawlNewParamsScrapeOptionsActionsType = "write"
	CrawlNewParamsScrapeOptionsActionsTypePress      CrawlNewParamsScrapeOptionsActionsType = "press"
	CrawlNewParamsScrapeOptionsActionsTypeScroll     CrawlNewParamsScrapeOptionsActionsType = "scroll"
	CrawlNewParamsScrapeOptionsActionsTypeScrape     CrawlNewParamsScrapeOptionsActionsType = "scrape"
	CrawlNewParamsScrapeOptionsActionsTypeExecute    CrawlNewParamsScrapeOptionsActionsType = "execute"
)

func (r CrawlNewParamsScrapeOptionsActionsType) IsKnown() bool {
	switch r {
	case CrawlNewParamsScrapeOptionsActionsTypeWait, CrawlNewParamsScrapeOptionsActionsTypeScreenshot, CrawlNewParamsScrapeOptionsActionsTypeClick, CrawlNewParamsScrapeOptionsActionsTypeWrite, CrawlNewParamsScrapeOptionsActionsTypePress, CrawlNewParamsScrapeOptionsActionsTypeScroll, CrawlNewParamsScrapeOptionsActionsTypeScrape, CrawlNewParamsScrapeOptionsActionsTypeExecute:
		return true
	}
	return false
}

type CrawlNewParamsScrapeOptionsFormat string

const (
	CrawlNewParamsScrapeOptionsFormatMarkdown           CrawlNewParamsScrapeOptionsFormat = "markdown"
	CrawlNewParamsScrapeOptionsFormatHTML               CrawlNewParamsScrapeOptionsFormat = "html"
	CrawlNewParamsScrapeOptionsFormatRawHTML            CrawlNewParamsScrapeOptionsFormat = "rawHtml"
	CrawlNewParamsScrapeOptionsFormatLinks              CrawlNewParamsScrapeOptionsFormat = "links"
	CrawlNewParamsScrapeOptionsFormatScreenshot         CrawlNewParamsScrapeOptionsFormat = "screenshot"
	CrawlNewParamsScrapeOptionsFormatScreenshotFullPage CrawlNewParamsScrapeOptionsFormat = "screenshot@fullPage"
	CrawlNewParamsScrapeOptionsFormatJson               CrawlNewParamsScrapeOptionsFormat = "json"
)

func (r CrawlNewParamsScrapeOptionsFormat) IsKnown() bool {
	switch r {
	case CrawlNewParamsScrapeOptionsFormatMarkdown, CrawlNewParamsScrapeOptionsFormatHTML, CrawlNewParamsScrapeOptionsFormatRawHTML, CrawlNewParamsScrapeOptionsFormatLinks, CrawlNewParamsScrapeOptionsFormatScreenshot, CrawlNewParamsScrapeOptionsFormatScreenshotFullPage, CrawlNewParamsScrapeOptionsFormatJson:
		return true
	}
	return false
}

type CrawlNewParamsScrapeOptionsJsonOptions struct {
	Prompt       param.Field[string]      `json:"prompt"`
	Schema       param.Field[interface{}] `json:"schema"`
	SystemPrompt param.Field[string]      `json:"systemPrompt"`
}

func (r CrawlNewParamsScrapeOptionsJsonOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CrawlNewParamsScrapeOptionsLocation struct {
	Country   param.Field[string]   `json:"country"`
	Languages param.Field[[]string] `json:"languages"`
}

func (r CrawlNewParamsScrapeOptionsLocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
