package main

import (
	"database/sql"
	"log"
	"net/http"

	server "github.com/deryrahman/secondbrain/server/http"
	"github.com/deryrahman/secondbrain/service/core"
	storage "github.com/deryrahman/secondbrain/storage/postgresql"
)

func main() {
	db := sql.OpenDB(nil)
	recordStorage, _ := storage.NewRecordStoragePSQL(db, storage.NewPSQLQuerier())
	recordService, _ := core.NewRecordService(recordStorage)
	httpServer, _ := server.NewHTTPServer("/api/v0.0.1", recordService)

	log.Fatal(http.ListenAndServe(":8080", httpServer))
}
