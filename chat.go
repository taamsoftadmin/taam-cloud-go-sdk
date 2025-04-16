// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package taamcloud

import (
	"context"
	"net/http"

	"github.com/taamsoftadmin/taam-cloud-go-sdk/internal/apijson"
	"github.com/taamsoftadmin/taam-cloud-go-sdk/internal/param"
	"github.com/taamsoftadmin/taam-cloud-go-sdk/internal/requestconfig"
	"github.com/taamsoftadmin/taam-cloud-go-sdk/option"
)

// ChatService contains methods and other services that help with interacting with
// the taam-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChatService] method instead.
type ChatService struct {
	Options []option.RequestOption
}

// NewChatService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewChatService(opts ...option.RequestOption) (r *ChatService) {
	r = &ChatService{}
	r.Options = opts
	return
}

// Generate chat completions based on provided conversation history
func (r *ChatService) NewCompletion(ctx context.Context, body ChatNewCompletionParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/chat/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type ChatNewCompletionParams struct {
	Messages  param.Field[[]ChatNewCompletionParamsMessage] `json:"messages,required"`
	Model     param.Field[string]                           `json:"model,required"`
	MaxTokens param.Field[int64]                            `json:"max_tokens"`
	// Whether to stream the response
	Stream      param.Field[bool]    `json:"stream"`
	Temperature param.Field[float64] `json:"temperature"`
}

func (r ChatNewCompletionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChatNewCompletionParamsMessage struct {
	Content param.Field[string]                              `json:"content,required"`
	Role    param.Field[ChatNewCompletionParamsMessagesRole] `json:"role,required"`
}

func (r ChatNewCompletionParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChatNewCompletionParamsMessagesRole string

const (
	ChatNewCompletionParamsMessagesRoleUser      ChatNewCompletionParamsMessagesRole = "user"
	ChatNewCompletionParamsMessagesRoleAssistant ChatNewCompletionParamsMessagesRole = "assistant"
	ChatNewCompletionParamsMessagesRoleSystem    ChatNewCompletionParamsMessagesRole = "system"
)

func (r ChatNewCompletionParamsMessagesRole) IsKnown() bool {
	switch r {
	case ChatNewCompletionParamsMessagesRoleUser, ChatNewCompletionParamsMessagesRoleAssistant, ChatNewCompletionParamsMessagesRoleSystem:
		return true
	}
	return false
}
