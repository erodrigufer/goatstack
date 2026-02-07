package main

import (
	"context"
	"fmt"
	"os"

	"github.com/erodrigufer/goatstack/internal/cli"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	cliApp := cli.CreateCLI()

	if err := cliApp.Run(context.Background(), os.Args); err != nil {
		return fmt.Errorf("unable to run cmd: %w", err)
	}
	return nil
}
