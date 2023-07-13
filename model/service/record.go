package model

import modelStorage "github.com/deryrahman/secondbrain/model/storage"

type GetRecordResponse struct {
	ID      string
	Content string
	Tags    []string
}

func (r *GetRecordResponse) From(record *modelStorage.GetRecordByTagsResponse) {
	if record == nil {
		return
	}
	r.ID = record.ID.String()
	r.Content = record.Content
	r.Tags = record.Tags
}
