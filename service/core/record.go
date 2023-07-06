package core

import (
	"context"
	"fmt"

	"github.com/deryrahman/secondbrain/model"
	"github.com/deryrahman/secondbrain/pkg/errors"
	"github.com/deryrahman/secondbrain/service"
	"github.com/deryrahman/secondbrain/storage"
	"github.com/google/uuid"
)

type RecordStorager interface {
	storage.RecordModifier
	storage.RecordGetter
}

var _ service.RecordService = (*recordService)(nil)

type recordService struct {
	recordStorager RecordStorager
}

func NewRecordService(recordStorager RecordStorager) (*recordService, error) {
	if recordStorager == nil {
		return nil, fmt.Errorf("recordStorager is nil")
	}
	return &recordService{
		recordStorager: recordStorager,
	}, nil
}

func (s *recordService) CreateRecord(ctx context.Context, content string, tags ...string) (string, error) {
	id := uuid.New()
	_, err := s.recordStorager.CreateRecordWithTags(ctx, id, content, tags...)
	if err != nil {
		return "", errors.Wrap(err)
	}
	return id.String(), nil
}

func (s *recordService) GetRecords(ctx context.Context, tags ...string) ([]model.RecordOnService, error) {
	return nil, nil
}
