package templates

import (
	"embed"
	"io"
	"text/template"
)

//go:embed all:content
var templatesFS embed.FS

type TemplateData struct {
	ModuleName string
	MainName   string
	DaemonName string
	UseDB      bool
}

func parseTemplates() (*template.Template, error) {
	templates, err := template.ParseFS(templatesFS, "content/*.tmpl", "content/backend/*.tmpl", "content/backend/*/*/*.tmpl", "content/backend/*/*/*/*.tmpl")
	if err != nil {
		return &template.Template{}, err
	}
	return templates, nil
}

func ExecuteTemplate(wr io.Writer, templateName string, data TemplateData) error {
	templates, err := parseTemplates()
	if err != nil {
		return err
	}
	return templates.ExecuteTemplate(wr, templateName, data)
}
