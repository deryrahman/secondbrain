package model

import codegenHandler "github.com/deryrahman/secondbrain/codegen/handler"

type PostRecordsJSONRequestBody codegenHandler.PostRecordsJSONRequestBody

type PostRecordsJSONResponseBody struct {
	Id string `json:"id,omitempty"`
}

type GetRecordsJSONResponseBody struct {
	RecordSnippets []RecordSnippet `json:"record_snippets,omitempty"`
}

type RecordSnippet codegenHandler.RecordSnippet
