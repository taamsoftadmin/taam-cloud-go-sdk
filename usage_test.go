// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package taamcloud_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/taam-cloud-go"
	"github.com/stainless-sdks/taam-cloud-go/internal/testutil"
	"github.com/stainless-sdks/taam-cloud-go/option"
)

func TestUsage(t *testing.T) {
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
	_, err := client.Embeddings.New(context.TODO(), taamcloud.EmbeddingNewParams{
		Input: taamcloud.F([]string{"string"}),
		Model: taamcloud.F("jina-embeddings-v3"),
	})
	if err != nil {
		t.Error(err)
	}
}
