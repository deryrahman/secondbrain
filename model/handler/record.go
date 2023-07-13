package model

import (
	codegenHandler "github.com/deryrahman/secondbrain/codegen/handler"
	modelService "github.com/deryrahman/secondbrain/model/service"
)

type GetRecordsResponse struct {
	RecordSnippets []*codegenHandler.RecordSnippet `json:"record_snippets,omitempty"`
}

type PostRecordsRequest codegenHandler.PostRecordsJSONRequestBody

type PostRecordsResponse struct {
	ID string `json:"id,omitempty"`
}

func (r *GetRecordsResponse) From(records []*modelService.GetRecordResponse) {
	r.RecordSnippets = make([]*codegenHandler.RecordSnippet, len(records))
	for i, record := range records {
		r.RecordSnippets[i] = &codegenHandler.RecordSnippet{
			Id:      &record.ID,
			Excerpt: &record.Content,
			Tags:    &record.Tags,
		}
	}
}
