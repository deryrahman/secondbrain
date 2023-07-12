package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	// metadata for goreleaser ldflags
	version string
	commit  string
	date    string
)

func main() {
	cli.VersionPrinter = func(ctx *cli.Context) {
		fmt.Printf("%s\n", ctx.App.Version)
	}

	app := &cli.App{
		Name:    "secondbrain cli",
		Version: version,
	}

	app.Commands = []*cli.Command{serveCommand()}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
