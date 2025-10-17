# Me

Response Types:

- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#MeGetResponse">MeGetResponse</a>

Methods:

- <code title="get /v1/api/me">client.Me.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#MeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#MeGetResponse">MeGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Contacts

Response Types:

- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ContactCheckCapabilitiesResponse">ContactCheckCapabilitiesResponse</a>

Methods:

- <code title="get /v1/api/contacts/{contact}/capabilities">client.Contacts.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ContactService.CheckCapabilities">CheckCapabilities</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, contact <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ContactCheckCapabilitiesResponse">ContactCheckCapabilitiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Messages

Response Types:

- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#MessageGetResponse">MessageGetResponse</a>
- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#MessageCancelResponse">MessageCancelResponse</a>
- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#MessageGetStatusResponse">MessageGetStatusResponse</a>
- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#MessageSendResponse">MessageSendResponse</a>

Methods:

- <code title="get /v1/api/messages/{messageId}">client.Messages.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#MessageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#MessageGetResponse">MessageGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/api/messages/{messageId}">client.Messages.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#MessageService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#MessageCancelResponse">MessageCancelResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/api/messages/{messageId}/status">client.Messages.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#MessageService.GetStatus">GetStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#MessageGetStatusResponse">MessageGetStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/api/messages">client.Messages.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#MessageService.Send">Send</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#MessageSendParams">MessageSendParams</a>) (<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#MessageSendResponse">MessageSendResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Config

## Webhook

Response Types:

- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ConfigWebhookGetResponse">ConfigWebhookGetResponse</a>
- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ConfigWebhookUpdateResponse">ConfigWebhookUpdateResponse</a>

Methods:

- <code title="get /v1/api/config/webhook">client.Config.Webhook.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ConfigWebhookService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ConfigWebhookGetResponse">ConfigWebhookGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/api/config/webhook">client.Config.Webhook.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ConfigWebhookService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ConfigWebhookUpdateParams">ConfigWebhookUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ConfigWebhookUpdateResponse">ConfigWebhookUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Batches

Methods:

- <code title="post /v1/api/batches">client.Batches.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#BatchService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/api/batches/{batchId}">client.Batches.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#BatchService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, batchID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/api/batches/{batchId}/messages">client.Batches.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#BatchService.ListMessages">ListMessages</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, batchID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/api/batches/{batchId}/status">client.Batches.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#BatchService.GetStatus">GetStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, batchID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
