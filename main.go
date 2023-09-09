package main

import (
	"arduino-simulator/cmd"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var version string

func main() {
	app := &cli.App{
		Name:           "sim",
		Usage:          "Arduino simulator",
		Version:        version,
		DefaultCommand: "disasm",
		Commands: []*cli.Command{
			cmd.Disasm,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
