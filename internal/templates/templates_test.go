package templates

import (
	"testing"
)

func Test_parseTemplates(t *testing.T) {
	files, err := getTemplateFSFilesBasePath()
	if err != nil {
		t.Fatalf("unable to retrieve files from templateFS: %s", err.Error())
	}
	totalTemplates := len(files)
	t.Run("Amount of templates", func(t *testing.T) {
		templates, err := parseTemplates()
		if err != nil {
			t.Fatalf("parseTemplates() returned an unexpected error: %s", err.Error())
		}
		amountOfTemplates := len(templates.Templates())
		if amountOfTemplates != totalTemplates {
			t.Errorf("amount of templates = %d, want %d", amountOfTemplates, totalTemplates)
		}
	})
}
