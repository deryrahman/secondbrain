package storage

import (
	"database/sql"

	"github.com/deryrahman/secondbrain/storage"
	_ "github.com/lib/pq"
)

func NewPSQLDB(dsn string) (storage.DB, error) {
	return sql.Open("postgres", dsn)
}
