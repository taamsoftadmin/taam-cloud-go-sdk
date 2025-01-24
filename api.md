# taamcloud

Response Types:

- <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#UploadResponse">UploadResponse</a>

Methods:

- <code title="post /upload">client.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#TaamcloudService.Upload">Upload</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#UploadParams">UploadParams</a>) (<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#UploadResponse">UploadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Embeddings

Response Types:

- <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#EmbeddingsResponse">EmbeddingsResponse</a>

Methods:

- <code title="post /v1/embeddings">client.Embeddings.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#EmbeddingService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#EmbeddingNewParams">EmbeddingNewParams</a>) (<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#EmbeddingsResponse">EmbeddingsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Rerank

Methods:

- <code title="post /v1/rerank">client.Rerank.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#RerankService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#RerankNewParams">RerankNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Chat

## Completions

Methods:

- <code title="post /v1/chat/completions">client.Chat.Completions.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#ChatCompletionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#ChatCompletionNewParams">ChatCompletionNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Suno

## Music

Methods:

- <code title="post /suno/submit/music">client.Suno.Music.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#SunoMusicService.Submit">Submit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#SunoMusicSubmitParams">SunoMusicSubmitParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Models

Response Types:

- <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#ModelListResponse">ModelListResponse</a>

Methods:

- <code title="get /v1/models">client.Models.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#ModelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#ModelListResponse">ModelListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Images

## Generations

Response Types:

- <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#ImageGenerationResponse">ImageGenerationResponse</a>

Methods:

- <code title="post /v1/images/generations">client.Images.Generations.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#ImageGenerationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#ImageGenerationNewParams">ImageGenerationNewParams</a>) (<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#ImageGenerationResponse">ImageGenerationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Crawl

Response Types:

- <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#CrawlResponse">CrawlResponse</a>
- <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#CrawlStatusResponse">CrawlStatusResponse</a>

Methods:

- <code title="post /v1/crawl">client.Crawl.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#CrawlService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#CrawlNewParams">CrawlNewParams</a>) (<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#CrawlResponse">CrawlResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/crawl/{id}">client.Crawl.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#CrawlService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#CrawlStatusResponse">CrawlStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Scrape

Response Types:

- <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#ScrapeResponse">ScrapeResponse</a>

Methods:

- <code title="post /v1/scrape">client.Scrape.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#ScrapeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#ScrapeNewParams">ScrapeNewParams</a>) (<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#ScrapeResponse">ScrapeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Maps

Response Types:

- <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#MapResponse">MapResponse</a>

Methods:

- <code title="post /v1/map">client.Maps.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#MapService.Discover">Discover</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#MapDiscoverParams">MapDiscoverParams</a>) (<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#MapResponse">MapResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Searches

Response Types:

- <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#SearchResponse">SearchResponse</a>

Methods:

- <code title="post /api/search">client.Searches.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#SearchService.Perform">Perform</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#SearchPerformParams">SearchPerformParams</a>) (<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk">taamcloud</a>.<a href="https://pkg.go.dev/github.com/taamsoftadmin/taam-cloud-go-sdk#SearchResponse">SearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
