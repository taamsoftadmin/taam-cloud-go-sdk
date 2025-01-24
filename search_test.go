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

func TestSearchPerformWithOptionalParams(t *testing.T) {
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
	_, err := client.Searches.Perform(context.TODO(), taamcloud.SearchPerformParams{
		ChatModel: taamcloud.F(taamcloud.SearchPerformParamsChatModel{
			Model:               taamcloud.F(taamcloud.SearchPerformParamsChatModelModelGpt3_5Turbo),
			Provider:            taamcloud.F(taamcloud.SearchPerformParamsChatModelProviderCustomOpenAI),
			CustomOpenAIBaseURL: taamcloud.F("customOpenAIBaseURL"),
			CustomOpenAIKey:     taamcloud.F("customOpenAIKey"),
		}),
		EmbeddingModel: taamcloud.F(taamcloud.SearchPerformParamsEmbeddingModel{
			Model:    taamcloud.F(taamcloud.SearchPerformParamsEmbeddingModelModelTextEmbedding3Large),
			Provider: taamcloud.F(taamcloud.SearchPerformParamsEmbeddingModelProviderCustomOpenAI),
		}),
		FocusMode: taamcloud.F(taamcloud.SearchPerformParamsFocusModeWebSearch),
	})
	if err != nil {
		var apierr *taamcloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
