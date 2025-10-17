# Me

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#MeGetResponse">MeGetResponse</a>

Methods:

- <code title="get /v1/api/me">client.Me.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#MeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#MeGetResponse">MeGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Contacts

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ContactCheckCapabilitiesResponse">ContactCheckCapabilitiesResponse</a>

Methods:

- <code title="get /v1/api/contacts/{contact}/capabilities">client.Contacts.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ContactService.CheckCapabilities">CheckCapabilities</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, contact <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ContactCheckCapabilitiesResponse">ContactCheckCapabilitiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Messages

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#MessageGetResponse">MessageGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#MessageCancelResponse">MessageCancelResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#MessageGetStatusResponse">MessageGetStatusResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#MessageSendResponse">MessageSendResponse</a>

Methods:

- <code title="get /v1/api/messages/{messageId}">client.Messages.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#MessageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#MessageGetResponse">MessageGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/api/messages/{messageId}">client.Messages.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#MessageService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#MessageCancelResponse">MessageCancelResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/api/messages/{messageId}/status">client.Messages.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#MessageService.GetStatus">GetStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#MessageGetStatusResponse">MessageGetStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/api/messages">client.Messages.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#MessageService.Send">Send</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#MessageSendParams">MessageSendParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#MessageSendResponse">MessageSendResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Config

## Webhook

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ConfigWebhookGetResponse">ConfigWebhookGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ConfigWebhookUpdateResponse">ConfigWebhookUpdateResponse</a>

Methods:

- <code title="get /v1/api/config/webhook">client.Config.Webhook.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ConfigWebhookService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ConfigWebhookGetResponse">ConfigWebhookGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/api/config/webhook">client.Config.Webhook.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ConfigWebhookService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ConfigWebhookUpdateParams">ConfigWebhookUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ConfigWebhookUpdateResponse">ConfigWebhookUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Batches

Methods:

- <code title="post /v1/api/batches">client.Batches.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#BatchService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/api/batches/{batchId}">client.Batches.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#BatchService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, batchID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/api/batches/{batchId}/messages">client.Batches.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#BatchService.ListMessages">ListMessages</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, batchID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/api/batches/{batchId}/status">client.Batches.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#BatchService.GetStatus">GetStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, batchID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
