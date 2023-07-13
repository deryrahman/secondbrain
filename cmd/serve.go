package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/deryrahman/secondbrain/pkg/errors"
	l "github.com/deryrahman/secondbrain/pkg/log"
	log "github.com/deryrahman/secondbrain/pkg/log/slog"
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

		logger := log.NewSLog(slog.DebugLevel)
		db, err := storage.NewPSQLDB(dsn)
		fatalAndExit(logger, err)

		recordStorage, err := storage.NewRecordStoragePSQL(logger, db, storage.NewPSQLQuerier())
		fatalAndExit(logger, err)

		recordService, err := core.NewRecordService(logger, recordStorage)
		fatalAndExit(logger, err)

		apiPath := fmt.Sprintf("/api/v%s", strings.Split(version, "-")[0])
		httpServer, err := server.NewHTTPServer(apiPath, logger, recordService)
		fatalAndExit(logger, err)

		logger.Infof("server running on port %s", port)
		if err := http.ListenAndServe(port, httpServer); err != nil {
			return errors.Wrap(err)
		}
		return nil
	}
}

func fatalAndExit(logger l.Logger, err error) {
	if err != nil {
		logger.Fatal(err)
		os.Exit(1)
	}
}
