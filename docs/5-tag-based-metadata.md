# [RFC 5] Tag Based Metadata
Status: Draft

This RFC aims to provide sufficient specification for building a metadata based on tag. The capability of this metadata is used for generic search, filtering, sorting and storing any useful information for the extended capability through plugin (which will be explained on the different RFC later). The basic idea of tags is to support the `recalling`, `organizing`, and `extending` capability. Tag should be named in a very short term and should not contain any full information but keeps complete and compact. For the initial specification, tag will have 3 major categories, first is `user` which are defined by end-user, second is `core` which to be used on the core system, third is `plugin` which to be used on the plugin specified functions. This RFC scoped only for the high level design.

## Purposes
- Generic search / filtering: search based on date, category, etc (the feature of search will be explain on the different RFC later)
- Sorting: sort is limited to sort in lexicographically order
- Extension: to extend the capability through plugin

## Tag Syntax
Three types of tags are structured in the following syntax:

| No | Type   | Proposed Syntax          |
|----|--------|--------------------------|
| 1  | user   | user;<tag_name>          |
| 2  | core   | core;<data_type>;<value> |
| 3  | plugin | plugin;<plugin_syntax>   |

The entity is separated by semi-colon `;`, hence the tag_name are not allowed to contains semi-colon.

**User Defined Tags**

This is the simplest tag's type. The value can only be defined from the end-user. The usage of this tag's type is limited on search, filter, and sort action.

**Core Defined Tags**

The core tag's type contains information on what are the system can do on that records. The creation of this tag is automatically created by the system whenever the record is created. The system defined the action (on this case search / filter, sort). In API perspective, user are not allowed to use this tag under the `tag` query params. The system should provide the query params for corresponding api endpoint based on the `data_type` introduced by this tag. For example, system provided the `data_type` date, then the corresponding api endpoints should have `date` query params for supporting this kind of tag.

**Plugin Defined Tags**

This RFC doesn't provide any comprehensive detail about the plugin mechanism, but one thing for sure that the plugin will behave on the record level based on the metadata provided from tag. Things need to consider for brainstorming the plugin:
- on which case this tag is created on the records? does the mechanism handled on the system or plugin?
- on where the plugin should live? on remote or within the same machine
- tag is being used to make sure which records can be taken an action from which plugin. The action or capability of the plugin is beyond this RFC

## Implementation details
The implementation is only for user defined tag (`user`) and system defined tag (`core`). Tag service provide the generic capability for those 3 type of tag, which contains this following function:
- Create
- Update

### User defined tag
Under the prefix `user;`
- provide the entry point for user creating the tag
- validate the tag name (should not have semi-colon)

These following are useful function on user defined tag:
- Filter

### System defined tag
Under the prefix `core;`
- provide autotag generation for every record created
- validation (should not have semi-colon is handled on system level)
Each data type provide unique functions on which they will handle the corresponding value
