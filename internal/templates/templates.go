// Package templates handles all the template files that generate an app.
package templates

import (
	"embed"
	"fmt"
	"io"
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
	templates, err := template.ParseFS(templatesFS, "content/*.tmpl", "content/backend/*.tmpl", "content/backend/*/*/*.tmpl", "content/backend/*/*/*/*.tmpl")
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
	return templates.ExecuteTemplate(wr, templateName, data)
}
