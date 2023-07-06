package storage

import (
	"database/sql"

	"github.com/deryrahman/secondbrain/pkg/errors"
	"github.com/deryrahman/secondbrain/storage"
	_ "github.com/lib/pq"
)

func NewPSQLDB(dsn string) (storage.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, errors.RootCause(err)
	}
	return db, nil
}
