# [RFC 1] Initial App Architecture
Status: Draft

This RFC covers detail of how the app should be implemented in minimal fashion. It includes the architecture of the app implementa†ion should follow, the minimal API details which RFC 0 is mentioned, and tech stack choices (programming language and database engine).

## Architecture implementation
The implementation should follow 3 layer architecture: handler, service, and persistence. The scopes of each layer are follow:
- **handler**: contribute to validate / modify the request and response, give appropriate http code status (IETF RFC7231), define contract definition for request and response, and become a model adapter between outside world contract and service layer contract.
- **service**: contains business related logic and model adapter between service layer contract and persistence layer contract. Service is not necessarily the name of the resource, it can be any object.
- **persistence**: handling the data management to interact with persistence object. Including storing, accessing, deleting, and updating. The name in persistence layer should be a resource name, and the behaviour is limited on CRUD.

Another layer but important:
- **sdk**: contains any components which essential for another developer to develop their own client/addon to interact with secondbrain app.

## Initial API details
Before jump in into the endpoint details, it's worth to mention that secondbrain app uses api with semantic versioning scheme. These following rules are applied:
- the endpoint convention should be like this /api/v1.0.0/<resource_name>
- endpoint /api/<resource_name> should be pointing out to latest api version
- endpoint with only mention major version, eg /api/v1/<resource_name> should be pointing out to latest version which major version resides
- resource name should be prural

### Create note with tags

**Endpoint**: POST /api/v0.1.0/notes
**Request Body**:
```json
{
    "content": <string>,
    "tags": [<string>,]
}
```
**Response**:
```json
{
    "id": <string>
}
```

**Response Status Codes**:
- 201: notes are created successfully
- 5xx: indicates something wrong in server

### Get list of notes

**Endpoint**: GET /api/v0.1.0/notes
**Request Params**:
- tag: specify the tag to fetch the list of note which match with this tag
**Response**:
```json
{
    "note_snippets": [
        {
            "id": <string>,
            "excerpt": <string>,
            "tags": [<string>,]
        },
    ]
}
```

**Response Status Codes**:
- 200: list of notes are fetched successfully
- 404: indicates the notes are not exist
- 5xx: indicates something wrong in server

## Tech Stack Choices

Shortlisted the tech stack uses:
- Go
- PostgreSQL
