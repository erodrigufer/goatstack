// Package cli configures the CLI.
package cli

import (
	"context"
	"fmt"
	"strings"

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
			Action: func(ctx context.Context, cmd *cli.Command, s string) error {
				return checkFlagValue("app", s)
			},
		},
		&cli.StringFlag{
			Name:     "module",
			Aliases:  []string{"m"},
			Usage:    "Module name on go.mod file",
			Required: true,
			Action: func(ctx context.Context, cmd *cli.Command, s string) error {
				return checkFlagValue("module", s)
			},
		},
		&cli.StringFlag{
			Name:     "daemon",
			Aliases:  []string{"d"},
			Usage:    "Daemon name to be used on deployment",
			Required: true,
			Action: func(ctx context.Context, cmd *cli.Command, s string) error {
				return checkFlagValue("daemon", s)
			},
		},
		&cli.StringFlag{
			Name:        "db",
			Usage:       "DB type used on app, e.g. 'sqlite', 'postgres'",
			Value:       "postgres",
			DefaultText: "postgres",
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

		fmt.Printf("Creating...\n\n •Application name: %s\n •Module name: %s\n •Daemon name: %s\n •DB Engine: %s\n\n", data.MainName, data.ModuleName, data.DaemonName, data.DB)

		err := templates.Populate(data)
		if err != nil {
			return fmt.Errorf("unable to populate project: %w", err)
		}
		fmt.Printf("✅ %s has been created.\n → Run `just dev` to start developing.\n\n", data.MainName)
		return nil
	}
}

func checkFlagValue(flagName, flagValue string) error {
	if strings.Contains(flagValue, " ") {
		return fmt.Errorf("flag `%s` value `%s` should not contain whitespaces", flagName, flagValue)
	}
	if strings.ToLower(flagValue) != flagValue {
		return fmt.Errorf("flag `%s` value `%s` should not contain uppercase letters", flagName, flagValue)
	}
	return nil
}
