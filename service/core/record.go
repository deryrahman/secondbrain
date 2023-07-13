package core

import (
	"context"
	"fmt"

	model "github.com/deryrahman/secondbrain/model/service"
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

func (s *recordService) GetRecords(ctx context.Context, tags ...string) ([]*model.GetRecordResponse, error) {
	records, err := s.recordStorager.GetRecordsByTags(ctx, tags...)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	recordsResponse := make([]*model.GetRecordResponse, len(records))
	for i, record := range records {
		recordsResponse[i] = &model.GetRecordResponse{}
		recordsResponse[i].From(record)
	}

	return recordsResponse, nil
}
