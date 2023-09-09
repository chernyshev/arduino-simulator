package cmd

import (
	"arduino-simulator/avr"
	"github.com/urfave/cli/v2"
)

var Disasm = &cli.Command{
	Name:  "disasm",
	Usage: "AVR disassembler",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "in",
			Usage:    "path to binary file",
			Required: true,
		},
	},
	Action: disassemble,
}

func disassemble(c *cli.Context) error {
	device := avr.New()
	err := device.LoadBinary(c.String("in"))
	if err != nil {
		return err
	}

	return device.Run()
}
