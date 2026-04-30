# Me

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#MeGetResponse">MeGetResponse</a>

Methods:

- <code title="get /me">client.Me.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#MeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#MeGetResponse">MeGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Numbers

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#MeNumberListResponse">MeNumberListResponse</a>

Methods:

- <code title="get /me/numbers">client.Me.Numbers.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#MeNumberService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#MeNumberListResponse">MeNumberListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### ContactCard

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#MeNumberContactCardGetResponse">MeNumberContactCardGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#MeNumberContactCardUpdateResponse">MeNumberContactCardUpdateResponse</a>

Methods:

- <code title="get /me/numbers/{number}/contact-card">client.Me.Numbers.ContactCard.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#MeNumberContactCardService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, number <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#MeNumberContactCardGetResponse">MeNumberContactCardGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /me/numbers/{number}/contact-card">client.Me.Numbers.ContactCard.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#MeNumberContactCardService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, number <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#MeNumberContactCardUpdateParams">MeNumberContactCardUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#MeNumberContactCardUpdateResponse">MeNumberContactCardUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Contacts

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#Contact">Contact</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#DeleteResponse">DeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#Pagination">Pagination</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ContactListResponse">ContactListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ContactCheckCapabilitiesResponse">ContactCheckCapabilitiesResponse</a>

Methods:

- <code title="post /contacts">client.Contacts.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ContactService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ContactNewParams">ContactNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#Contact">Contact</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /contacts/{contactId}">client.Contacts.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ContactService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, contactID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#Contact">Contact</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /contacts/{contactId}">client.Contacts.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ContactService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, contactID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ContactUpdateParams">ContactUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#Contact">Contact</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /contacts">client.Contacts.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ContactService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ContactListParams">ContactListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ContactListResponse">ContactListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /contacts/{contactId}">client.Contacts.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ContactService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, contactID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#DeleteResponse">DeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /contacts/{contactId}/capabilities">client.Contacts.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ContactService.CheckCapabilities">CheckCapabilities</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, contactID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ContactCheckCapabilitiesResponse">ContactCheckCapabilitiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Tags

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ContactTagListResponse">ContactTagListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ContactTagAddResponse">ContactTagAddResponse</a>

Methods:

- <code title="get /contacts/{contactId}/tags">client.Contacts.Tags.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ContactTagService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, contactID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ContactTagListResponse">ContactTagListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /contacts/{contactId}/tags">client.Contacts.Tags.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ContactTagService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, contactID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ContactTagAddParams">ContactTagAddParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ContactTagAddResponse">ContactTagAddResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /contacts/{contactId}/tags/{tag}">client.Contacts.Tags.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ContactTagService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, tag <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ContactTagRemoveParams">ContactTagRemoveParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#DeleteResponse">DeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Location

## Contacts

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ContactLocation">ContactLocation</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#LocationContactListResponse">LocationContactListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#LocationContactRefreshResponse">LocationContactRefreshResponse</a>

Methods:

- <code title="get /location/contacts/{handle}">client.Location.Contacts.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#LocationContactService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, handle <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ContactLocation">ContactLocation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /location/contacts">client.Location.Contacts.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#LocationContactService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#LocationContactListResponse">LocationContactListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /location/contacts/refresh">client.Location.Contacts.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#LocationContactService.Refresh">Refresh</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#LocationContactRefreshResponse">LocationContactRefreshResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Facetime

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#FacetimeInitiateCallResponse">FacetimeInitiateCallResponse</a>

Methods:

- <code title="post /facetime/calls">client.Facetime.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#FacetimeService.InitiateCall">InitiateCall</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#FacetimeInitiateCallParams">FacetimeInitiateCallParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#FacetimeInitiateCallResponse">FacetimeInitiateCallResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Groups

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#Group">Group</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupNewResponse">GroupNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupUpdateResponse">GroupUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupListResponse">GroupListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupDeleteResponse">GroupDeleteResponse</a>

Methods:

- <code title="post /groups">client.Groups.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupNewParams">GroupNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupNewResponse">GroupNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /groups/{groupId}">client.Groups.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, groupID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#Group">Group</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /groups/{groupId}">client.Groups.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, groupID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupUpdateParams">GroupUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupUpdateResponse">GroupUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /groups">client.Groups.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupListParams">GroupListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupListResponse">GroupListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /groups/{groupId}">client.Groups.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, groupID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupDeleteResponse">GroupDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Members

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupMember">GroupMember</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupMemberListResponse">GroupMemberListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupMemberAddResponse">GroupMemberAddResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupMemberRemoveResponse">GroupMemberRemoveResponse</a>

Methods:

- <code title="get /groups/{groupId}/members">client.Groups.Members.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupMemberService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, groupID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupMemberListParams">GroupMemberListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupMemberListResponse">GroupMemberListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /groups/{groupId}/members">client.Groups.Members.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupMemberService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, groupID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupMemberAddParams">GroupMemberAddParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupMemberAddResponse">GroupMemberAddResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /groups/{groupId}/members/{contactId}">client.Groups.Members.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupMemberService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, contactID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupMemberRemoveParams">GroupMemberRemoveParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupMemberRemoveResponse">GroupMemberRemoveResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Icon

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupIcon">GroupIcon</a>

Methods:

- <code title="delete /groups/{groupId}/icon">client.Groups.Icon.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupIconService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, groupID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupIcon">GroupIcon</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /groups/{groupId}/icon">client.Groups.Icon.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupIconService.Set">Set</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, groupID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupIconSetParams">GroupIconSetParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#GroupIcon">GroupIcon</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Webhooks

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#Webhook">Webhook</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#WebhookNewResponse">WebhookNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#WebhookListResponse">WebhookListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#WebhookDeleteResponse">WebhookDeleteResponse</a>

Methods:

- <code title="post /webhooks">client.Webhooks.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#WebhookService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#WebhookNewParams">WebhookNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#WebhookNewResponse">WebhookNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /webhooks/{webhookId}">client.Webhooks.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#WebhookService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#Webhook">Webhook</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /webhooks/{webhookId}">client.Webhooks.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#WebhookService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#WebhookUpdateParams">WebhookUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#Webhook">Webhook</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /webhooks">client.Webhooks.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#WebhookService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#WebhookListResponse">WebhookListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /webhooks/{webhookId}">client.Webhooks.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#WebhookService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#WebhookDeleteResponse">WebhookDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Secret

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#WebhookSecretRotateResponse">WebhookSecretRotateResponse</a>

Methods:

- <code title="post /webhooks/{webhookId}/secret/rotate">client.Webhooks.Secret.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#WebhookSecretService.Rotate">Rotate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#WebhookSecretRotateResponse">WebhookSecretRotateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Logs

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#WebhookLogListResponse">WebhookLogListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#WebhookLogReplayResponse">WebhookLogReplayResponse</a>

Methods:

- <code title="get /webhooks/{webhookId}/logs">client.Webhooks.Logs.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#WebhookLogService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#WebhookLogListParams">WebhookLogListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#WebhookLogListResponse">WebhookLogListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /webhooks/{webhookId}/logs/{eventId}/replay">client.Webhooks.Logs.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#WebhookLogService.Replay">Replay</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, eventID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#WebhookLogReplayParams">WebhookLogReplayParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#WebhookLogReplayResponse">WebhookLogReplayResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Chats

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#LastMessage">LastMessage</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatGetResponse">ChatGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatListResponse">ChatListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatMarkAsReadResponse">ChatMarkAsReadResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatShareContactCardResponse">ChatShareContactCardResponse</a>

Methods:

- <code title="get /chats/{chatId}">client.Chats.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatGetResponse">ChatGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /chats">client.Chats.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatListParams">ChatListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatListResponse">ChatListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /chats/{chatId}/read">client.Chats.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatService.MarkAsRead">MarkAsRead</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatMarkAsReadResponse">ChatMarkAsReadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /chats/{chatId}/contact-card">client.Chats.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatService.ShareContactCard">ShareContactCard</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatShareContactCardResponse">ChatShareContactCardResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Messages

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#LinkPreviewParam">LinkPreviewParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#Reaction">Reaction</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatMessageGetResponse">ChatMessageGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatMessageListResponse">ChatMessageListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatMessageGetStatusResponse">ChatMessageGetStatusResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatMessageReactResponse">ChatMessageReactResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatMessageSendResponse">ChatMessageSendResponse</a>

Methods:

- <code title="get /chats/{chatId}/messages/{messageId}">client.Chats.Messages.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatMessageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatMessageGetParams">ChatMessageGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatMessageGetResponse">ChatMessageGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /chats/{chatId}/messages">client.Chats.Messages.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatMessageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatMessageListParams">ChatMessageListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatMessageListResponse">ChatMessageListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /chats/{chatId}/messages/{messageId}/status">client.Chats.Messages.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatMessageService.GetStatus">GetStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatMessageGetStatusParams">ChatMessageGetStatusParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatMessageGetStatusResponse">ChatMessageGetStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /chats/{chatId}/messages/{messageId}/reactions">client.Chats.Messages.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatMessageService.React">React</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatMessageReactParams">ChatMessageReactParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatMessageReactResponse">ChatMessageReactResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /chats/{chatId}/messages">client.Chats.Messages.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatMessageService.Send">Send</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatMessageSendParams">ChatMessageSendParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatMessageSendResponse">ChatMessageSendResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Polls

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatPollGetResultsResponse">ChatPollGetResultsResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatPollSendResponse">ChatPollSendResponse</a>

Methods:

- <code title="get /chats/{chatId}/polls/{pollId}">client.Chats.Polls.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatPollService.GetResults">GetResults</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pollID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatPollGetResultsParams">ChatPollGetResultsParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatPollGetResultsResponse">ChatPollGetResultsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /chats/{chatId}/polls">client.Chats.Polls.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatPollService.Send">Send</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatPollSendParams">ChatPollSendParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatPollSendResponse">ChatPollSendResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Typing

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#TypingResponse">TypingResponse</a>

Methods:

- <code title="post /chats/{chatId}/typing">client.Chats.Typing.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatTypingService.Start">Start</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#TypingResponse">TypingResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /chats/{chatId}/typing">client.Chats.Typing.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatTypingService.Stop">Stop</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#TypingResponse">TypingResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Background

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatBackgroundResponse">ChatBackgroundResponse</a>

Methods:

- <code title="get /chats/{chatId}/background">client.Chats.Background.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatBackgroundService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatBackgroundResponse">ChatBackgroundResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /chats/{chatId}/background">client.Chats.Background.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatBackgroundService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatBackgroundResponse">ChatBackgroundResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /chats/{chatId}/background">client.Chats.Background.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatBackgroundService.Set">Set</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatBackgroundSetParams">ChatBackgroundSetParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#ChatBackgroundResponse">ChatBackgroundResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# PhoneNumbers

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#PhoneNumberBatchNewResponse">PhoneNumberBatchNewResponse</a>

Methods:

- <code title="post /phone-numbers/batch">client.PhoneNumbers.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#PhoneNumberService.BatchNew">BatchNew</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#PhoneNumberBatchNewParams">PhoneNumberBatchNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#PhoneNumberBatchNewResponse">PhoneNumberBatchNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Lookup

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#PhoneNumberLookupResult">PhoneNumberLookupResult</a>

Methods:

- <code title="post /phone-numbers/lookup">client.PhoneNumbers.Lookup.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#PhoneNumberLookupService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#PhoneNumberLookupNewParams">PhoneNumberLookupNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#PhoneNumberLookupResult">PhoneNumberLookupResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /phone-numbers/lookup">client.PhoneNumbers.Lookup.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#PhoneNumberLookupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#PhoneNumberLookupGetParams">PhoneNumberLookupGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go">blooio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blooio-go#PhoneNumberLookupResult">PhoneNumberLookupResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
