# Embeddings

Response Types:

- <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#EmbeddingNewResponse">EmbeddingNewResponse</a>

Methods:

- <code title="post /v1/embeddings">client.Embeddings.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#EmbeddingService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#EmbeddingNewParams">EmbeddingNewParams</a>) (<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#EmbeddingNewResponse">EmbeddingNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Rerank

Methods:

- <code title="post /v1/rerank">client.Rerank.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#RerankService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#RerankNewParams">RerankNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Chat

Methods:

- <code title="post /v1/chat/completions">client.Chat.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#ChatService.NewCompletion">NewCompletion</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#ChatNewCompletionParams">ChatNewCompletionParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Suno

## Submit

Methods:

- <code title="post /suno/submit/music">client.Suno.Submit.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#SunoSubmitService.GenerateMusic">GenerateMusic</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#SunoSubmitGenerateMusicParams">SunoSubmitGenerateMusicParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Models

Response Types:

- <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#ModelListResponse">ModelListResponse</a>

Methods:

- <code title="get /v1/models">client.Models.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#ModelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#ModelListResponse">ModelListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Images

Response Types:

- <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#ImageGenerateResponse">ImageGenerateResponse</a>

Methods:

- <code title="post /v1/images/generations">client.Images.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#ImageService.Generate">Generate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#ImageGenerateParams">ImageGenerateParams</a>) (<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#ImageGenerateResponse">ImageGenerateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Web

Response Types:

- <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#WebNewResponse">WebNewResponse</a>

Methods:

- <code title="post /v1/web">client.Web.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#WebService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#WebNewParams">WebNewParams</a>) (<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#WebNewResponse">WebNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Files

Response Types:

- <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#FileGetResponse">FileGetResponse</a>
- <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#FileUploadResponse">FileUploadResponse</a>

Methods:

- <code title="get /v1/files/retrieve">client.Files.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#FileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#FileGetParams">FileGetParams</a>) (<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#FileGetResponse">FileGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/files">client.Files.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#FileService.Upload">Upload</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#FileUploadParams">FileUploadParams</a>) (<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#FileUploadResponse">FileUploadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Upload

Response Types:

- <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#UploadNewResponse">UploadNewResponse</a>

Methods:

- <code title="post /upload">client.Upload.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#UploadService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#UploadNewParams">UploadNewParams</a>) (<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#UploadNewResponse">UploadNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# VideoGeneration

Response Types:

- <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#VideoGenerationNewResponse">VideoGenerationNewResponse</a>

Methods:

- <code title="post /v1/video_generation">client.VideoGeneration.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#VideoGenerationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#VideoGenerationNewParams">VideoGenerationNewParams</a>) (<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#VideoGenerationNewResponse">VideoGenerationNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Query

Response Types:

- <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#QueryCheckVideoGenerationStatusResponse">QueryCheckVideoGenerationStatusResponse</a>

Methods:

- <code title="get /v1/query/video_generation">client.Query.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#QueryService.CheckVideoGenerationStatus">CheckVideoGenerationStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#QueryCheckVideoGenerationStatusParams">QueryCheckVideoGenerationStatusParams</a>) (<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#QueryCheckVideoGenerationStatusResponse">QueryCheckVideoGenerationStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
