package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

type ProjectCommand struct{}

func (c *ProjectCommand) GetName() string {
	return "project"
}

func (c *ProjectCommand) GetUsage() string {
	return "change or create project"
}

func (c *ProjectCommand) Run(ctx *cli.Context) error {
	// return filesystem.CherryInitIfNotExist()
	fmt.Println(ctx.Args().Get(0))
	return nil
}
