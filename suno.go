// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package taamcloud

import (
	"github.com/taamsoftadmin/taam-cloud-go-sdk/option"
)

// SunoService contains methods and other services that help with interacting with
// the taam-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSunoService] method instead.
type SunoService struct {
	Options []option.RequestOption
	Music   *SunoMusicService
}

// NewSunoService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSunoService(opts ...option.RequestOption) (r *SunoService) {
	r = &SunoService{}
	r.Options = opts
	r.Music = NewSunoMusicService(opts...)
	return
}
