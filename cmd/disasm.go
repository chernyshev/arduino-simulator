package cmd

import (
	"github.com/urfave/cli/v2"
)

var Disasm = &cli.Command{
	Name:   "disasm",
	Usage:  "AVR disassembler",
	Action: disassemble,
}

func disassemble(c *cli.Context) error {
	return nil
}
