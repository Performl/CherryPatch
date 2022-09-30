package cmd

import (
	"CherryPatch/internal/filesystem"

	"github.com/urfave/cli/v2"
)

type InitCommand struct{}

func (c *InitCommand) GetName() string {
	return "init"
}

func (c *InitCommand) GetUsage() string {
	return "this function initialises the Cherry Patch directory"
}

func (c *InitCommand) Run(ctx *cli.Context) error {
	return filesystem.CherryInitIfNotExist()
}
