package templates

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func walkTemplateFS(walkDirFunc fs.WalkDirFunc) error {
	subTemplateFS, err := fs.Sub(templatesFS, "content")
	if err != nil {
		return fmt.Errorf("unable to create subtree FS: %w", err)
	}
	return fs.WalkDir(subTemplateFS, ".", walkDirFunc)
}

// getTemplateFSFilesFullPath returns a slice with the full path of all
// non-directory files on the templates FS.
func getTemplateFSFilesFullPath() ([]string, error) {
	files := make([]string, 0)
	walkDirFunc := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("an error happened while traversing templateFS: %w", err)
		}
		if !d.IsDir() {
			files = append(files, path)
		}
		return nil
	}
	err := walkTemplateFS(walkDirFunc)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func getTemplateFSFilesBasePath() ([]string, error) {
	files, err := getTemplateFSFilesFullPath()
	if err != nil {
		return nil, err
	}
	for i, path := range files {
		files[i] = filepath.Base(path)
	}
	return files, nil
}

func checkNoDuplicateTemplateNames(files []string) error {
	fileNames := make(map[string]string, 0)
	for _, file := range files {
		_, found := fileNames[file]
		if found {
			return fmt.Errorf("template name %s is duplicated", file)
		}
		fileNames[file] = file
	}
	return nil
}

func Populate(data TemplateData) error {
	templatesBaseNames, err := getTemplateFSFilesBasePath()
	if err != nil {
		return fmt.Errorf("unable to get templates base names: %w", err)
	}
	err = checkNoDuplicateTemplateNames(templatesBaseNames)
	if err != nil {
		return fmt.Errorf("template files do not have unique names: %w", err)
	}
	root, err := os.OpenRoot(".")
	if err != nil {
		return fmt.Errorf("unable to create os.Root: %w", err)
	}
	defer root.Close()

	walkDirFunc := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		switch isDir := d.IsDir(); isDir {
		case true:
			root.Mkdir(path, 0o755)
		case false:
			pathWithoutSuffix, err := removeFilenameSuffix(path, ".tmpl")
			if err != nil {
				return fmt.Errorf("unable to remove file suffix: %w", err)
			}
			file, err := root.Create(pathWithoutSuffix)
			if err != nil {
				return fmt.Errorf("unable to create file: %w", err)
			}
			defer file.Close()
			err = ExecuteTemplate(file, filepath.Base(path), data)
			if err != nil {
				return fmt.Errorf("unable to execute template for file %s: %w", path, err)
			}

		}
		return nil
	}

	err = walkTemplateFS(walkDirFunc)
	if err != nil {
		return err
	}
	return nil
}

// removeFilenameSuffix removes a suffix (e.g. `templ`) from a filename.
func removeFilenameSuffix(path, suffix string) (string, error) {
	withoutSuffix, ok := strings.CutSuffix(path, suffix)
	if !ok {
		return "", fmt.Errorf("could not find suffix %s in path %s", suffix, path)
	}
	return withoutSuffix, nil
}
