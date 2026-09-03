# Me

Response Types:

- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#MeGetResponse">MeGetResponse</a>

Methods:

- <code title="get /me">client.Me.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#MeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#MeGetResponse">MeGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Numbers

Response Types:

- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#MeNumberListResponse">MeNumberListResponse</a>

Methods:

- <code title="get /me/numbers">client.Me.Numbers.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#MeNumberService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#MeNumberListResponse">MeNumberListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### ContactCard

Response Types:

- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#MeNumberContactCardGetResponse">MeNumberContactCardGetResponse</a>
- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#MeNumberContactCardUpdateResponse">MeNumberContactCardUpdateResponse</a>

Methods:

- <code title="get /me/numbers/{number}/contact-card">client.Me.Numbers.ContactCard.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#MeNumberContactCardService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, number <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#MeNumberContactCardGetResponse">MeNumberContactCardGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /me/numbers/{number}/contact-card">client.Me.Numbers.ContactCard.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#MeNumberContactCardService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, number <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#MeNumberContactCardUpdateParams">MeNumberContactCardUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#MeNumberContactCardUpdateResponse">MeNumberContactCardUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Contacts

Response Types:

- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#Contact">Contact</a>
- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#DeleteResponse">DeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#Pagination">Pagination</a>
- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ContactListResponse">ContactListResponse</a>
- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ContactCheckCapabilitiesResponse">ContactCheckCapabilitiesResponse</a>

Methods:

- <code title="post /contacts">client.Contacts.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ContactService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ContactNewParams">ContactNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#Contact">Contact</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /contacts/{contactId}">client.Contacts.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ContactService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, contactID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#Contact">Contact</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /contacts/{contactId}">client.Contacts.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ContactService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, contactID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ContactUpdateParams">ContactUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#Contact">Contact</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /contacts">client.Contacts.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ContactService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ContactListParams">ContactListParams</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ContactListResponse">ContactListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /contacts/{contactId}">client.Contacts.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ContactService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, contactID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#DeleteResponse">DeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /contacts/{contactId}/capabilities">client.Contacts.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ContactService.CheckCapabilities">CheckCapabilities</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, contactID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ContactCheckCapabilitiesResponse">ContactCheckCapabilitiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Tags

Response Types:

- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ContactTagListResponse">ContactTagListResponse</a>
- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ContactTagAddResponse">ContactTagAddResponse</a>

Methods:

- <code title="get /contacts/{contactId}/tags">client.Contacts.Tags.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ContactTagService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, contactID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ContactTagListResponse">ContactTagListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /contacts/{contactId}/tags">client.Contacts.Tags.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ContactTagService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, contactID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ContactTagAddParams">ContactTagAddParams</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ContactTagAddResponse">ContactTagAddResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /contacts/{contactId}/tags/{tag}">client.Contacts.Tags.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ContactTagService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, tag <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ContactTagRemoveParams">ContactTagRemoveParams</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#DeleteResponse">DeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Location

## Contacts

Response Types:

- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ContactLocation">ContactLocation</a>
- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#LocationContactListResponse">LocationContactListResponse</a>
- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#LocationContactRefreshResponse">LocationContactRefreshResponse</a>

Methods:

- <code title="get /location/contacts/{handle}">client.Location.Contacts.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#LocationContactService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, handle <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ContactLocation">ContactLocation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /location/contacts">client.Location.Contacts.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#LocationContactService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#LocationContactListResponse">LocationContactListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /location/contacts/refresh">client.Location.Contacts.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#LocationContactService.Refresh">Refresh</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#LocationContactRefreshResponse">LocationContactRefreshResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Facetime

Methods:

- <code title="post /facetime/calls">client.Facetime.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#FacetimeService.InitiateCall">InitiateCall</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#FacetimeInitiateCallParams">FacetimeInitiateCallParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Groups

Response Types:

- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#Group">Group</a>
- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#GroupNewResponse">GroupNewResponse</a>
- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#GroupUpdateResponse">GroupUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#GroupListResponse">GroupListResponse</a>
- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#GroupDeleteResponse">GroupDeleteResponse</a>

Methods:

- <code title="post /groups">client.Groups.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#GroupService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#GroupNewParams">GroupNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#GroupNewResponse">GroupNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /groups/{groupId}">client.Groups.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#GroupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, groupID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#Group">Group</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /groups/{groupId}">client.Groups.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#GroupService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, groupID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#GroupUpdateParams">GroupUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#GroupUpdateResponse">GroupUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /groups">client.Groups.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#GroupService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#GroupListParams">GroupListParams</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#GroupListResponse">GroupListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /groups/{groupId}">client.Groups.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#GroupService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, groupID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#GroupDeleteResponse">GroupDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Members

Response Types:

- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#GroupMember">GroupMember</a>
- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#GroupMemberListResponse">GroupMemberListResponse</a>

Methods:

- <code title="get /groups/{groupId}/members">client.Groups.Members.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#GroupMemberService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, groupID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#GroupMemberListParams">GroupMemberListParams</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#GroupMemberListResponse">GroupMemberListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /groups/{groupId}/members">client.Groups.Members.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#GroupMemberService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, groupID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#GroupMemberAddParams">GroupMemberAddParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="delete /groups/{groupId}/members/{contactId}">client.Groups.Members.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#GroupMemberService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, contactID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#GroupMemberRemoveParams">GroupMemberRemoveParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Icon

Response Types:

- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#GroupIcon">GroupIcon</a>

Methods:

- <code title="delete /groups/{groupId}/icon">client.Groups.Icon.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#GroupIconService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, groupID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#GroupIcon">GroupIcon</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /groups/{groupId}/icon">client.Groups.Icon.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#GroupIconService.Set">Set</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, groupID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#GroupIconSetParams">GroupIconSetParams</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#GroupIcon">GroupIcon</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Webhooks

Response Types:

- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#Webhook">Webhook</a>
- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#WebhookNewResponse">WebhookNewResponse</a>
- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#WebhookListResponse">WebhookListResponse</a>
- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#WebhookDeleteResponse">WebhookDeleteResponse</a>

Methods:

- <code title="post /webhooks">client.Webhooks.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#WebhookService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#WebhookNewParams">WebhookNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#WebhookNewResponse">WebhookNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /webhooks/{webhookId}">client.Webhooks.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#WebhookService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#Webhook">Webhook</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /webhooks/{webhookId}">client.Webhooks.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#WebhookService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#WebhookUpdateParams">WebhookUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#Webhook">Webhook</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /webhooks">client.Webhooks.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#WebhookService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#WebhookListResponse">WebhookListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /webhooks/{webhookId}">client.Webhooks.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#WebhookService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#WebhookDeleteResponse">WebhookDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Secret

Response Types:

- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#WebhookSecretRotateResponse">WebhookSecretRotateResponse</a>

Methods:

- <code title="post /webhooks/{webhookId}/secret/rotate">client.Webhooks.Secret.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#WebhookSecretService.Rotate">Rotate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#WebhookSecretRotateResponse">WebhookSecretRotateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Logs

Response Types:

- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#WebhookLogListResponse">WebhookLogListResponse</a>
- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#WebhookLogReplayResponse">WebhookLogReplayResponse</a>

Methods:

- <code title="get /webhooks/{webhookId}/logs">client.Webhooks.Logs.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#WebhookLogService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#WebhookLogListParams">WebhookLogListParams</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#WebhookLogListResponse">WebhookLogListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /webhooks/{webhookId}/logs/{eventId}/replay">client.Webhooks.Logs.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#WebhookLogService.Replay">Replay</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, eventID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#WebhookLogReplayParams">WebhookLogReplayParams</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#WebhookLogReplayResponse">WebhookLogReplayResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Chats

Response Types:

- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#LastMessage">LastMessage</a>
- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatGetResponse">ChatGetResponse</a>
- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatListResponse">ChatListResponse</a>
- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatMarkAsReadResponse">ChatMarkAsReadResponse</a>
- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatShareContactCardResponse">ChatShareContactCardResponse</a>

Methods:

- <code title="get /chats/{chatId}">client.Chats.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatGetResponse">ChatGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /chats">client.Chats.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatListParams">ChatListParams</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatListResponse">ChatListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /chats/{chatId}/read">client.Chats.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatService.MarkAsRead">MarkAsRead</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatMarkAsReadResponse">ChatMarkAsReadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /chats/{chatId}/contact-card">client.Chats.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatService.ShareContactCard">ShareContactCard</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatShareContactCardResponse">ChatShareContactCardResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Messages

Params Types:

- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#LinkPreviewParam">LinkPreviewParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#Reaction">Reaction</a>
- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatMessageGetResponse">ChatMessageGetResponse</a>
- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatMessageListResponse">ChatMessageListResponse</a>
- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatMessageGetStatusResponse">ChatMessageGetStatusResponse</a>
- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatMessageReactResponse">ChatMessageReactResponse</a>
- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatMessageSendResponse">ChatMessageSendResponse</a>

Methods:

- <code title="get /chats/{chatId}/messages/{messageId}">client.Chats.Messages.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatMessageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatMessageGetParams">ChatMessageGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatMessageGetResponse">ChatMessageGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /chats/{chatId}/messages">client.Chats.Messages.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatMessageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatMessageListParams">ChatMessageListParams</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatMessageListResponse">ChatMessageListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /chats/{chatId}/messages/{messageId}/status">client.Chats.Messages.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatMessageService.GetStatus">GetStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatMessageGetStatusParams">ChatMessageGetStatusParams</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatMessageGetStatusResponse">ChatMessageGetStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /chats/{chatId}/messages/{messageId}/reactions">client.Chats.Messages.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatMessageService.React">React</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatMessageReactParams">ChatMessageReactParams</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatMessageReactResponse">ChatMessageReactResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /chats/{chatId}/messages">client.Chats.Messages.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatMessageService.Send">Send</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatMessageSendParams">ChatMessageSendParams</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatMessageSendResponse">ChatMessageSendResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Polls

Response Types:

- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatPollGetResultsResponse">ChatPollGetResultsResponse</a>
- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatPollSendResponse">ChatPollSendResponse</a>

Methods:

- <code title="get /chats/{chatId}/polls/{pollId}">client.Chats.Polls.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatPollService.GetResults">GetResults</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pollID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatPollGetResultsParams">ChatPollGetResultsParams</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatPollGetResultsResponse">ChatPollGetResultsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /chats/{chatId}/polls">client.Chats.Polls.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatPollService.Send">Send</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatPollSendParams">ChatPollSendParams</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatPollSendResponse">ChatPollSendResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Typing

Response Types:

- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#TypingResponse">TypingResponse</a>

Methods:

- <code title="post /chats/{chatId}/typing">client.Chats.Typing.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatTypingService.Start">Start</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#TypingResponse">TypingResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /chats/{chatId}/typing">client.Chats.Typing.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatTypingService.Stop">Stop</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#TypingResponse">TypingResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Background

Response Types:

- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatBackgroundResponse">ChatBackgroundResponse</a>

Methods:

- <code title="get /chats/{chatId}/background">client.Chats.Background.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatBackgroundService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatBackgroundResponse">ChatBackgroundResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /chats/{chatId}/background">client.Chats.Background.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatBackgroundService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatBackgroundResponse">ChatBackgroundResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /chats/{chatId}/background">client.Chats.Background.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatBackgroundService.Set">Set</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatBackgroundSetParams">ChatBackgroundSetParams</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#ChatBackgroundResponse">ChatBackgroundResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# PhoneNumbers

Response Types:

- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#PhoneNumberBatchNewResponse">PhoneNumberBatchNewResponse</a>

Methods:

- <code title="post /phone-numbers/batch">client.PhoneNumbers.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#PhoneNumberService.BatchNew">BatchNew</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#PhoneNumberBatchNewParams">PhoneNumberBatchNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#PhoneNumberBatchNewResponse">PhoneNumberBatchNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Lookup

Response Types:

- <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#PhoneNumberLookupResult">PhoneNumberLookupResult</a>

Methods:

- <code title="post /phone-numbers/lookup">client.PhoneNumbers.Lookup.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#PhoneNumberLookupService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#PhoneNumberLookupNewParams">PhoneNumberLookupNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#PhoneNumberLookupResult">PhoneNumberLookupResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /phone-numbers/lookup">client.PhoneNumbers.Lookup.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#PhoneNumberLookupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#PhoneNumberLookupGetParams">PhoneNumberLookupGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk">blooio</a>.<a href="https://pkg.go.dev/github.com/Blooio/blooio-go-sdk#PhoneNumberLookupResult">PhoneNumberLookupResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
