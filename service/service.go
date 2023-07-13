package service

import (
	"context"

	model "github.com/deryrahman/secondbrain/model/service"
)

type RecordService interface {
	CreateRecord(ctx context.Context, content string, tags ...string) (string, error)
	GetRecords(ctx context.Context, tags ...string) ([]*model.GetRecordResponse, error)
}
