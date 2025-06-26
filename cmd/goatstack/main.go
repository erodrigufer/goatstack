package main

import (
	"fmt"
	"os"

	"github.com/erodrigufer/goatstack/internal/templates"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	data := templates.TemplateData{
		ModuleName: "erodriguez.de/app",
		MainName:   "webserver",
		DaemonName: "webserverd",
	}

	err := templates.Populate(data)
	if err != nil {
		return fmt.Errorf("unable to populate project: %w", err)
	}
	return nil
}
