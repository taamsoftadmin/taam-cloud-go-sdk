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

func TestTaamcloudUploadWithOptionalParams(t *testing.T) {
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
	_, err := client.Upload(context.TODO(), taamcloud.UploadParams{
		File:         taamcloud.F(io.Reader(bytes.NewBuffer([]byte("some file contents")))),
		EnableOcr:    taamcloud.F(true),
		EnableVision: taamcloud.F(true),
		SaveAll:      taamcloud.F(true),
	})
	if err != nil {
		var apierr *taamcloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
