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

func TestChatNewCompletionWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	err := client.Chat.NewCompletion(context.TODO(), taamcloud.ChatNewCompletionParams{
		Messages: taamcloud.F([]taamcloud.ChatNewCompletionParamsMessage{{
			Content: taamcloud.F("You are a helpful assistant."),
			Role:    taamcloud.F(taamcloud.ChatNewCompletionParamsMessagesRoleSystem),
		}, {
			Content: taamcloud.F("Hello, how are you today?"),
			Role:    taamcloud.F(taamcloud.ChatNewCompletionParamsMessagesRoleUser),
		}}),
		Model:       taamcloud.F("gpt-4"),
		MaxTokens:   taamcloud.F(int64(150)),
		Stream:      taamcloud.F(true),
		Temperature: taamcloud.F(0.700000),
	})
	if err != nil {
		var apierr *taamcloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
