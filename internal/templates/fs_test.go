package templates

import "testing"

func Test_checkNoDuplicateTemplateNames(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		files   []string
		wantErr bool
	}{
		{
			name:    "no duplicates",
			files:   []string{"justfile.tmpl", "README.md.tmpl"},
			wantErr: false,
		},
		{
			name:    "duplicated templates",
			files:   []string{"justfile.tmpl", "README.md.tmpl", "main.go.tmpl", "README.md.tmpl"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := checkNoDuplicateTemplateNames(tt.files)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("checkNoDuplicateTemplateNames() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("checkNoDuplicateTemplateNames() succeeded unexpectedly")
			}
		})
	}
}
