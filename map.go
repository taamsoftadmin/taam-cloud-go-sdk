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

// MapService contains methods and other services that help with interacting with
// the taam-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMapService] method instead.
type MapService struct {
	Options []option.RequestOption
}

// NewMapService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMapService(opts ...option.RequestOption) (r *MapService) {
	r = &MapService{}
	r.Options = opts
	return
}

// Discover and map all links on a website
func (r *MapService) Discover(ctx context.Context, body MapDiscoverParams, opts ...option.RequestOption) (res *MapResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/map"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type MapResponse struct {
	Links   []string        `json:"links"`
	Success bool            `json:"success"`
	JSON    mapResponseJSON `json:"-"`
}

// mapResponseJSON contains the JSON metadata for the struct [MapResponse]
type mapResponseJSON struct {
	Links       apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MapResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mapResponseJSON) RawJSON() string {
	return r.raw
}

type MapDiscoverParams struct {
	// The base URL to start mapping from
	URL param.Field[string] `json:"url,required"`
	// Ignore the website sitemap when crawling
	IgnoreSitemap param.Field[bool] `json:"ignoreSitemap"`
	// Include subdomains of the website
	IncludeSubdomains param.Field[bool] `json:"includeSubdomains"`
	// Maximum number of links to return
	Limit param.Field[int64] `json:"limit"`
	// Search query to use for mapping
	Search param.Field[string] `json:"search"`
	// Only return links found in the website sitemap
	SitemapOnly param.Field[bool] `json:"sitemapOnly"`
}

func (r MapDiscoverParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
