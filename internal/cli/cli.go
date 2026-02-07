// Package cli configures the CLI.
package cli

import (
	"context"
	"fmt"

	"github.com/erodrigufer/goatstack/internal/templates"
	"github.com/urfave/cli/v3"
)

func CreateCLI() *cli.Command {
	return &cli.Command{
		Name:   "create",
		Usage:  "Create a new project from template",
		Flags:  configCreateFlags(),
		Action: configCreateAction(),
	}
}

func configCreateFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "app",
			Aliases:  []string{"a"},
			Usage:    "Main app name",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "module",
			Aliases:  []string{"m"},
			Usage:    "Module name on go.mod file",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "daemon",
			Aliases:  []string{"d"},
			Usage:    "Daemon name to be used on deployment",
			Required: true,
		},
		&cli.StringFlag{
			Name:  "db",
			Usage: "DB type used on app, e.g. 'sqlite', 'postgres'",
		},
	}
}

func configCreateAction() cli.ActionFunc {
	return func(ctx context.Context, cmd *cli.Command) error {
		data := templates.TemplateData{
			ModuleName: cmd.String("module"),
			MainName:   cmd.String("app"),
			DaemonName: cmd.String("daemon"),
			DB:         cmd.String("db"),
		}

		err := templates.Populate(data)
		if err != nil {
			return fmt.Errorf("unable to populate project: %w", err)
		}
		return nil
	}
}
