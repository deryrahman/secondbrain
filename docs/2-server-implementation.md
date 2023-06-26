# [RFC 2] Server Implementation

Purpose of this RFC is to define the server implementation for minimal setup of the project. It includes explanation about the expectation of how the server runs, framework choices, config specification, database design, etc.

## Expectation
Running the server should be as seamless as possible by running the command `sb serve`. Default config and some flags configuration will be defined on config specification section.

## Framework Choices
To create as simple as generic possible, this project won't use any fancy web framework. Initial implementation will use net/http, later on fancy web framework can be used if use cases force it to doing so.
Persistence layer uses sqlc to achieve query based storage development. While go-migration is used to migrate the schema.
Handler layer uses open api generator to achieve design first implementation.

## Config Specs

Initial server implementation would have these following flags:
```
--dsn : dsn connection string
--port : port in which the server would run
```

## DB Design

```
| record 
+------
| id: uuid (PK)
| content: string
```

```
| tag
+------
| id:  (PK)
```

```
[record] *---* [tag]
```