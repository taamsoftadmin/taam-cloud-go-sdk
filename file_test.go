// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package taamcloud_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/taamsoftadmin/taam-cloud-go-sdk"
	"github.com/taamsoftadmin/taam-cloud-go-sdk/internal/testutil"
	"github.com/taamsoftadmin/taam-cloud-go-sdk/option"
)

func TestFileGet(t *testing.T) {
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
	_, err := client.Files.Get(context.TODO(), taamcloud.FileGetParams{
		FileID: taamcloud.F("file_id"),
	})
	if err != nil {
		var apierr *taamcloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFileUploadWithOptionalParams(t *testing.T) {
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
	_, err := client.Files.Upload(context.TODO(), taamcloud.FileUploadParams{
		File:          taamcloud.F(io.Reader(bytes.NewBuffer([]byte("some file contents")))),
		EnableOcr:     taamcloud.F(taamcloud.FileUploadParamsEnableOcrTrue),
		EnableVision:  taamcloud.F(taamcloud.FileUploadParamsEnableVisionTrue),
		ExtractMode:   taamcloud.F(taamcloud.FileUploadParamsExtractModeDefault),
		ImagesOnly:    taamcloud.F(taamcloud.FileUploadParamsImagesOnlyTrue),
		PageBased:     taamcloud.F(taamcloud.FileUploadParamsPageBasedTrue),
		RemoveHeaders: taamcloud.F(taamcloud.FileUploadParamsRemoveHeadersTrue),
		SaveAll:       taamcloud.F(taamcloud.FileUploadParamsSaveAllTrue),
		TextOnly:      taamcloud.F(taamcloud.FileUploadParamsTextOnlyTrue),
		VisionOnly:    taamcloud.F(taamcloud.FileUploadParamsVisionOnlyTrue),
	})
	if err != nil {
		var apierr *taamcloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
