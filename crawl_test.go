// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package taamcloud_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/taamsoftadmin/taam-cloud-go-sdk"
	"github.com/taamsoftadmin/taam-cloud-go-sdk/internal/testutil"
	"github.com/taamsoftadmin/taam-cloud-go-sdk/option"
)

func TestCrawlNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Crawl.New(context.TODO(), taamcloud.CrawlNewParams{
		URL:                taamcloud.F("url"),
		AllowBackwardLinks: taamcloud.F(true),
		AllowExternalLinks: taamcloud.F(true),
		ExcludePaths:       taamcloud.F([]string{"string"}),
		IgnoreSitemap:      taamcloud.F(true),
		IncludePaths:       taamcloud.F([]string{"string"}),
		Limit:              taamcloud.F(int64(0)),
		MaxDepth:           taamcloud.F(int64(0)),
		ScrapeOptions: taamcloud.F(taamcloud.CrawlNewParamsScrapeOptions{
			Actions: taamcloud.F([]taamcloud.CrawlNewParamsScrapeOptionsActionUnion{taamcloud.CrawlNewParamsScrapeOptionsActionsObject{
				Milliseconds: taamcloud.F(int64(0)),
				Type:         taamcloud.F(taamcloud.CrawlNewParamsScrapeOptionsActionsObjectTypeWait),
				Selector:     taamcloud.F("selector"),
			}}),
			ExcludeTags: taamcloud.F([]string{"string"}),
			Formats:     taamcloud.F([]taamcloud.CrawlNewParamsScrapeOptionsFormat{taamcloud.CrawlNewParamsScrapeOptionsFormatMarkdown}),
			Headers: taamcloud.F(map[string]interface{}{
				"foo": "bar",
			}),
			IncludeTags: taamcloud.F([]string{"string"}),
			JsonOptions: taamcloud.F(taamcloud.CrawlNewParamsScrapeOptionsJsonOptions{
				Prompt:       taamcloud.F("prompt"),
				Schema:       taamcloud.F[any](map[string]interface{}{}),
				SystemPrompt: taamcloud.F("systemPrompt"),
			}),
			Location: taamcloud.F(taamcloud.CrawlNewParamsScrapeOptionsLocation{
				Country:   taamcloud.F("country"),
				Languages: taamcloud.F([]string{"string"}),
			}),
			Mobile:              taamcloud.F(true),
			OnlyMainContent:     taamcloud.F(true),
			RemoveBase64Images:  taamcloud.F(true),
			SkipTlsVerification: taamcloud.F(true),
			Timeout:             taamcloud.F(int64(0)),
			WaitFor:             taamcloud.F(int64(0)),
		}),
		Webhook: taamcloud.F("webhook"),
	})
	if err != nil {
		var apierr *taamcloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCrawlGet(t *testing.T) {
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
	_, err := client.Crawl.Get(context.TODO(), "id")
	if err != nil {
		var apierr *taamcloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
