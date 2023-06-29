package main

import (
	"fmt"
	"net/http"

	slogx "github.com/deryrahman/secondbrain/pkg/log/slog"
	server "github.com/deryrahman/secondbrain/server/http"
	"github.com/deryrahman/secondbrain/service/core"
	storage "github.com/deryrahman/secondbrain/storage/postgresql"
	"github.com/gookit/slog"
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

		logger := slogx.NewSLog(slog.DebugLevel)
		db, err := storage.NewPSQLDB(dsn)
		if err != nil {
			logger.Fatal(err)
			return err
		}
		recordStorage, err := storage.NewRecordStoragePSQL(db, storage.NewPSQLQuerier())
		if err != nil {
			logger.Fatal(err)
			return err
		}
		recordService, err := core.NewRecordService(recordStorage)
		if err != nil {
			logger.Fatal(err)
			return err
		}
		httpServer, err := server.NewHTTPServer("/api/v0.0.1", logger, recordService)
		if err != nil {
			logger.Fatal(err)
			return err
		}

		return http.ListenAndServe(port, httpServer)
	}
}
