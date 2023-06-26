package service

import (
	"context"

	"github.com/deryrahman/secondbrain/model"
)

type RecordService interface {
	CreateRecord(ctx context.Context, content string, tags ...string) (string, error)
	GetRecords(ctx context.Context, tags ...string) ([]model.RecordOnService, error)
}
