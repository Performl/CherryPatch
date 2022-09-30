package main

import (
	cmd "CherryPatch/internal/command"

	"github.com/urfave/cli/v2"
)

func InitCommands() *cli.App {
	return &cli.App{
		EnableBashCompletion: true,
		Commands: cmd.CommandFactory([]cmd.ICommand{
			&cmd.InitCommand{},
			&cmd.ProjectCommand{},
		}),
	}
}
