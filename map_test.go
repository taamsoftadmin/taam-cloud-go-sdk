// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package taamcloud_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/taam-cloud-go"
	"github.com/stainless-sdks/taam-cloud-go/internal/testutil"
	"github.com/stainless-sdks/taam-cloud-go/option"
)

func TestMapDiscoverWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := taamcloud.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Maps.Discover(context.TODO(), taamcloud.MapDiscoverParams{
		URL:               taamcloud.F("url"),
		IgnoreSitemap:     taamcloud.F(true),
		IncludeSubdomains: taamcloud.F(true),
		Limit:             taamcloud.F(int64(0)),
		Search:            taamcloud.F("search"),
		SitemapOnly:       taamcloud.F(true),
	})
	if err != nil {
		var apierr *taamcloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
