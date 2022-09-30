package cmd

import "github.com/urfave/cli/v2"

type ICommand interface {
	GetName() string
	GetUsage() string
	Run(*cli.Context) error
}

func CommandFactory(commands []ICommand) []*cli.Command {
	cliCommands := []*cli.Command{}
	for _, command := range commands {
		cliCommands = append(cliCommands, &cli.Command{
			Name:  command.GetName(),
			Usage: command.GetUsage(),
			Action: func(ctx *cli.Context) error {
				return command.Run(ctx)
			},
		})
	}
	return cliCommands
}
