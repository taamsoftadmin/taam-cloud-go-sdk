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

func TestImageGenerateWithOptionalParams(t *testing.T) {
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
	_, err := client.Images.Generate(context.TODO(), taamcloud.ImageGenerateParams{
		Prompt:  taamcloud.F("A beautiful sunset over a calm ocean"),
		Model:   taamcloud.F(taamcloud.ImageGenerateParamsModelDallE3),
		N:       taamcloud.F(int64(1)),
		Quality: taamcloud.F(taamcloud.ImageGenerateParamsQualityStandard),
		Size:    taamcloud.F(taamcloud.ImageGenerateParamsSize1024x1024),
		Style:   taamcloud.F(taamcloud.ImageGenerateParamsStyleNatural),
	})
	if err != nil {
		var apierr *taamcloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
