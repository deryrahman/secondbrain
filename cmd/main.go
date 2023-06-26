package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	version = "v0.0.1"
)

func main() {
	app := &cli.App{
		Name:    "secondbrain cli 2",
		Version: version,
	}

	app.Commands = []*cli.Command{serveCommand()}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
