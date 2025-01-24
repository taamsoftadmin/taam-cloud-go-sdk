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

func TestImageGenerationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Images.Generations.New(context.TODO(), taamcloud.ImageGenerationNewParams{
		Prompt:  taamcloud.F("prompt"),
		Model:   taamcloud.F(taamcloud.ImageGenerationNewParamsModelDallE3),
		N:       taamcloud.F(int64(0)),
		Quality: taamcloud.F(taamcloud.ImageGenerationNewParamsQualityStandard),
		Size:    taamcloud.F(taamcloud.ImageGenerationNewParamsSize1024x1024),
		Style:   taamcloud.F(taamcloud.ImageGenerationNewParamsStyleNatural),
	})
	if err != nil {
		var apierr *taamcloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
