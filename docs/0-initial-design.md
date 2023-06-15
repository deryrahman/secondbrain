# [RFC 0] Initial Design
Status: Draft

The purpose of this initial RFC is to define what the app's capabilities, terminologies, limitations, initial architectures, and the possible core components needed.

## Capabilities
- Seamless note taking for ideas, quotes, moments, and anything with minimal features which only essential for storing, recalling, and organizing
- Flexible to integrate with any client

## Terminologies
To standardize the development process, terminologies are given in this RFC. Changes for definition of these following terminologies is unlikely.
- **note**: refer to the entity which contains text-based information from user
- **tag**: identifiers which belongs to the note. One tag can be used for multiple note
- **filter**: a query to fetch the list of notes based on the given conditions / rules
- **client**: any entity which access this app

## Limitations
- Implementation detail for each component is not defined in this RFC
- This RFC only defines core specifications in which the capabilities defined above can be achieved
- The scope is only for building the functionality. How client use this functionality is out of scope

## Initial Architectures

```
+---+           +---+
| C | - HTTP -> | S |
+---+           +---|

C: client {cli, android app, ios app, web}
S: secondbrain app server
```

The app uses HTTP as protocol transport for information exchange between client and server. Client should be registered and verified by the server which is not covered in this RFC.

## Components
In the initial RFC, core components are minimal building blocks in which user can write note, give tags, store the note, and fetch the notes with specified filter. In order to achieve that, several endpoints need to be introduced which are not covered in this RFC

| Functionality | Endpoint | Request | Response |
|---|---|---|---|
| Create note with tags | POST: /v1alpha1/notes | {content:string,tags:[string]} | {id:string} |
| Get list of notes | GET: /v1alpha1/notes | ?filter=tag_name | {note_snippets:[{id:string,excerpt:string,tags:[string]}]} |

Another core component for extension capability is addon. Other developers can create their own addon and change the behaviour of the app functionality. The functionality which can be enhanced by an addon are storing and fetching (recalling / organizing). For example, addon to automatically add the date tag, everytime user create the note, then the date tag will be automatically appended. Another example is an addon to fetch with multiple filter logic, "fetch the note which contains word grocery in the tag a week ago but doesn't have tag done". The addon detail is not covered in this RFC.

Additional components which are not necessary but important are authentication + authorization mechanism. The plan is to use 3rd party auth to leverage simplicity of user sign up and log in. Detailed implementations are not covered in this RFC.
