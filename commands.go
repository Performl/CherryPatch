package main

import (
	"CherryPatch/internal/filesystem"

	"github.com/urfave/cli/v2"
)

func InitCommands() *cli.App {
	return &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "init",
				Usage: "initialise cherrypatch root",
				Action: func(*cli.Context) error {
					filesystem.CherryInitIfNotExist()
					return nil
				},
			},
		},
	}
}
