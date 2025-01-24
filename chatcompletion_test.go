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

func TestChatCompletionNewWithOptionalParams(t *testing.T) {
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
	err := client.Chat.Completions.New(context.TODO(), taamcloud.ChatCompletionNewParams{
		Messages: taamcloud.F([]taamcloud.ChatCompletionNewParamsMessage{{
			Content: taamcloud.F("content"),
			Role:    taamcloud.F(taamcloud.ChatCompletionNewParamsMessagesRoleUser),
		}}),
		Model:       taamcloud.F("model"),
		MaxTokens:   taamcloud.F(int64(0)),
		Stream:      taamcloud.F(true),
		Temperature: taamcloud.F(0.000000),
	})
	if err != nil {
		var apierr *taamcloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
