package storage

import (
	"context"
	"database/sql"

	model "github.com/deryrahman/secondbrain/model/storage"
	"github.com/google/uuid"
)

type DB interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

type RecordModifier interface {
	CreateRecordWithTags(ctx context.Context, id uuid.UUID, content string, tags ...string) (*model.CreateRecordWithTagsResponse, error)
}

type RecordGetter interface {
	GetRecordsByTags(ctx context.Context, tags ...string) ([]*model.GetRecordByTagsResponse, error)
}
