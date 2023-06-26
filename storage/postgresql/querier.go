package storage

import codegenStorage "github.com/deryrahman/secondbrain/codegen/storage"

func NewPSQLQuerier() codegenStorage.Querier {
	return codegenStorage.New()
}
