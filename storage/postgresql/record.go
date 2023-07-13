package storage

import (
	"context"
	"fmt"

	codegenStorage "github.com/deryrahman/secondbrain/codegen/storage"
	model "github.com/deryrahman/secondbrain/model/storage"
	"github.com/deryrahman/secondbrain/pkg/errors"
	"github.com/deryrahman/secondbrain/storage"
	"github.com/google/uuid"
)

var (
	_ storage.RecordModifier = (*recordStorage)(nil)
	_ storage.RecordGetter   = (*recordStorage)(nil)
)

type recordStorage struct {
	db      storage.DB
	querier codegenStorage.Querier
}

func NewRecordStoragePSQL(db storage.DB, querier codegenStorage.Querier) (*recordStorage, error) {
	var err error
	if db == nil {
		err = errors.Join(err, fmt.Errorf("db is nil"))
	}
	if querier == nil {
		err = errors.Join(err, fmt.Errorf("querier is nil"))
	}
	if err != nil {
		return nil, errors.RootCause(err)
	}

	return &recordStorage{
		db:      db,
		querier: querier,
	}, nil
}

func (s *recordStorage) CreateRecordWithTags(ctx context.Context, id uuid.UUID, content string, tags ...string) (*model.CreateRecordWithTagsResponse, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.RootCause(err)
	}

	createRecordParams := codegenStorage.CreateRecordParams{ID: id, Content: content}
	recordID, err := s.querier.CreateRecord(ctx, tx, createRecordParams)
	if err != nil {
		if e := tx.Rollback(); e != nil {
			err = errors.Join(err, e)
		}
		return nil, errors.RootCause(err)
	}

	for _, tag := range tags {
		upsertTagPArams := codegenStorage.UpsertTagParams{ID: tag}
		if err := s.querier.UpsertTag(ctx, tx, upsertTagPArams); err != nil {
			if e := tx.Rollback(); e != nil {
				err = errors.Join(err, e)
			}
			return nil, errors.RootCause(err)
		}
		associateRecordToTagParams := codegenStorage.AssociateNoteToTagParams{RecordID: recordID, TagID: tag}
		if err := s.querier.AssociateNoteToTag(ctx, tx, associateRecordToTagParams); err != nil {
			if e := tx.Rollback(); e != nil {
				err = errors.Join(err, e)
			}
			return nil, errors.RootCause(err)
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, errors.RootCause(err)
	}
	return &model.CreateRecordWithTagsResponse{ID: id, Content: content}, nil
}

func (s *recordStorage) GetRecordsByTags(ctx context.Context, tags ...string) ([]*model.GetRecordByTagsResponse, error) {
	getRecordsByTagParams := codegenStorage.GetRecordsByTagParams{Column1: tags}
	records, err := s.querier.GetRecordsByTag(ctx, s.db, getRecordsByTagParams)
	if err != nil {
		return nil, errors.RootCause(err)
	}

	recordsResponse := make([]*model.GetRecordByTagsResponse, len(records))
	for i, record := range records {
		recordsResponse[i] = &model.GetRecordByTagsResponse{}
		recordsResponse[i].From(record)
	}

	return recordsResponse, nil
}
