# [RFC 1] Initial App Architecture
Status: Draft

This RFC covers detail of how the app should be implemented with minimal fashion. It includes the architecture of the app implementa†ion should follow, the minimal API details which RFC 0 is mentioned, and tech stack choices (programming language and database engine).

## Architecture implementation
The implementation should follow 3 layer architecture: handler, service, and persistence. The scopes of each layer are follow:
- **handler**: contribute to validate / modify the request and response, give appropriate http code status (IETF RFC7231), define contract definition for request and response, and become a model adapter between outside world contract and service layer contract.
- **service**: contains business related logic and model adapter between service layer contract and persistence layer contract. Service is not necessarily the name of the resource, it can be any object.
- **persistence**: handling the data management to interact with persistence object. Including storing, accessing, deleting, and updating. The name in persistence layer should be a resource name, and the behaviour is limited on CRUD.

Another layer but important:
- **sdk**: contains any components which essential for another developer to develop their own client/addon to interact with secondbrain app.
