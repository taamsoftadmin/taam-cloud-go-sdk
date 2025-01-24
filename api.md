# taamcloud

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go">taamcloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#UploadResponse">UploadResponse</a>

Methods:

- <code title="post /upload">client.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#TaamcloudService.Upload">Upload</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go">taamcloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#UploadParams">UploadParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go">taamcloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#UploadResponse">UploadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Embeddings

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go">taamcloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#EmbeddingsResponse">EmbeddingsResponse</a>

Methods:

- <code title="post /v1/embeddings">client.Embeddings.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#EmbeddingService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go">taamcloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#EmbeddingNewParams">EmbeddingNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go">taamcloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#EmbeddingsResponse">EmbeddingsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Rerank

Methods:

- <code title="post /v1/rerank">client.Rerank.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#RerankService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go">taamcloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#RerankNewParams">RerankNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Chat

## Completions

Methods:

- <code title="post /v1/chat/completions">client.Chat.Completions.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#ChatCompletionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go">taamcloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#ChatCompletionNewParams">ChatCompletionNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Suno

## Music

Methods:

- <code title="post /suno/submit/music">client.Suno.Music.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#SunoMusicService.Submit">Submit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go">taamcloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#SunoMusicSubmitParams">SunoMusicSubmitParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Models

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go">taamcloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#ModelListResponse">ModelListResponse</a>

Methods:

- <code title="get /v1/models">client.Models.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#ModelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go">taamcloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#ModelListResponse">ModelListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Images

## Generations

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go">taamcloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#ImageGenerationResponse">ImageGenerationResponse</a>

Methods:

- <code title="post /v1/images/generations">client.Images.Generations.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#ImageGenerationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go">taamcloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#ImageGenerationNewParams">ImageGenerationNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go">taamcloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#ImageGenerationResponse">ImageGenerationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Crawl

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go">taamcloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#CrawlResponse">CrawlResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go">taamcloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#CrawlStatusResponse">CrawlStatusResponse</a>

Methods:

- <code title="post /v1/crawl">client.Crawl.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#CrawlService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go">taamcloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#CrawlNewParams">CrawlNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go">taamcloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#CrawlResponse">CrawlResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/crawl/{id}">client.Crawl.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#CrawlService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go">taamcloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#CrawlStatusResponse">CrawlStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Scrape

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go">taamcloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#ScrapeResponse">ScrapeResponse</a>

Methods:

- <code title="post /v1/scrape">client.Scrape.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#ScrapeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go">taamcloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#ScrapeNewParams">ScrapeNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go">taamcloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#ScrapeResponse">ScrapeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Maps

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go">taamcloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#MapResponse">MapResponse</a>

Methods:

- <code title="post /v1/map">client.Maps.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#MapService.Discover">Discover</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go">taamcloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#MapDiscoverParams">MapDiscoverParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go">taamcloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#MapResponse">MapResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Searches

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go">taamcloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#SearchResponse">SearchResponse</a>

Methods:

- <code title="post /api/search">client.Searches.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#SearchService.Perform">Perform</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go">taamcloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#SearchPerformParams">SearchPerformParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go">taamcloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/taam-cloud-go#SearchResponse">SearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
