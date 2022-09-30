package cmd

import (
	"CherryPatch/internal/project"

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
	projectName := ctx.Args().First()
	return project.CreateProject(projectName)
}
