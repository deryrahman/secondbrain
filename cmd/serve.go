package main

import (
	"fmt"
	"net/http"

	server "github.com/deryrahman/secondbrain/server/http"
	"github.com/deryrahman/secondbrain/service/core"
	storage "github.com/deryrahman/secondbrain/storage/postgresql"
	"github.com/urfave/cli/v2"
)

func serveCommand() *cli.Command {
	flags := []cli.Flag{
		&cli.StringFlag{Name: "dsn", Usage: "dsn connection string", Required: true},
		&cli.StringFlag{Name: "port", Usage: "port in which the server would run", Value: "8080"},
	}
	return &cli.Command{
		Name:   "serve",
		Flags:  flags,
		Action: serveActionFunc(),
	}
}

func serveActionFunc() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		dsn := ctx.String("dsn")
		port := fmt.Sprintf(":%s", ctx.String("port"))

		db, err := storage.NewPSQLDB(dsn)
		if err != nil {
			return err
		}
		recordStorage, _ := storage.NewRecordStoragePSQL(db, storage.NewPSQLQuerier())
		recordService, _ := core.NewRecordService(recordStorage)
		httpServer, _ := server.NewHTTPServer("/api/v0.0.1", recordService)

		return http.ListenAndServe(port, httpServer)
	}
}
