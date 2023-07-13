package model

import (
	codegenStorage "github.com/deryrahman/secondbrain/codegen/storage"
	"github.com/google/uuid"
)

type GetRecordByTagsResponse struct {
	ID      uuid.UUID
	Content string
	Tags    []string
}

type CreateRecordWithTagsResponse struct {
	ID      uuid.UUID
	Content string
}

func (r *GetRecordByTagsResponse) From(record *codegenStorage.GetRecordsByTagRow) {
	if record == nil {
		return
	}
	r.ID = record.ID
	r.Content = record.Content
	r.Tags = record.Tags
}
