// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package taamcloud

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/taam-cloud-go/internal/apijson"
	"github.com/stainless-sdks/taam-cloud-go/internal/param"
	"github.com/stainless-sdks/taam-cloud-go/internal/requestconfig"
	"github.com/stainless-sdks/taam-cloud-go/option"
)

// ChatCompletionService contains methods and other services that help with
// interacting with the taam-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChatCompletionService] method instead.
type ChatCompletionService struct {
	Options []option.RequestOption
}

// NewChatCompletionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewChatCompletionService(opts ...option.RequestOption) (r *ChatCompletionService) {
	r = &ChatCompletionService{}
	r.Options = opts
	return
}

// Chat completions
func (r *ChatCompletionService) New(ctx context.Context, body ChatCompletionNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/chat/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type ChatCompletionNewParams struct {
	Messages  param.Field[[]ChatCompletionNewParamsMessage] `json:"messages,required"`
	Model     param.Field[string]                           `json:"model,required"`
	MaxTokens param.Field[int64]                            `json:"max_tokens"`
	// Whether to stream the response
	Stream      param.Field[bool]    `json:"stream"`
	Temperature param.Field[float64] `json:"temperature"`
}

func (r ChatCompletionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChatCompletionNewParamsMessage struct {
	Content param.Field[string]                              `json:"content,required"`
	Role    param.Field[ChatCompletionNewParamsMessagesRole] `json:"role,required"`
}

func (r ChatCompletionNewParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChatCompletionNewParamsMessagesRole string

const (
	ChatCompletionNewParamsMessagesRoleUser      ChatCompletionNewParamsMessagesRole = "user"
	ChatCompletionNewParamsMessagesRoleAssistant ChatCompletionNewParamsMessagesRole = "assistant"
	ChatCompletionNewParamsMessagesRoleSystem    ChatCompletionNewParamsMessagesRole = "system"
)

func (r ChatCompletionNewParamsMessagesRole) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesRoleUser, ChatCompletionNewParamsMessagesRoleAssistant, ChatCompletionNewParamsMessagesRoleSystem:
		return true
	}
	return false
}
