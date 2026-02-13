// Package templates handles all the template files that generate an app.
package templates

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"path/filepath"
	"text/template"
)

//go:embed all:content
var templatesFS embed.FS

type TemplateData struct {
	ModuleName string
	MainName   string
	DaemonName string
	DB         string
}

func parseTemplates() (*template.Template, error) {
	var templateFiles []string
	err := fs.WalkDir(templatesFS, "content", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && filepath.Ext(path) == ".tmpl" {
			templateFiles = append(templateFiles, path)
		}
		return nil
	})
	if err != nil {
		return &template.Template{}, fmt.Errorf("unable to walk content directory: %w", err)
	}

	templates, err := template.ParseFS(templatesFS, templateFiles...)
	if err != nil {
		return &template.Template{}, fmt.Errorf("unable to parse filesystem: %w", err)
	}
	return templates, nil
}

func executeTemplate(wr io.Writer, templateName string, data TemplateData) error {
	templates, err := parseTemplates()
	if err != nil {
		return fmt.Errorf("unable to parse templates: %w", err)
	}
	err = templates.ExecuteTemplate(wr, templateName, data)
	if err != nil {
		return fmt.Errorf("an error happened executing the templates: %w", err)
	}
	return nil
}
