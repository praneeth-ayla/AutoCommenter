package scanner

import (
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
)

var skipDirs = map[string]bool{
	"node_modules": true,
	".git":         true,
	"venv":         true,
	".next":        true,
	"build":        true,
}

var skipFilePatterns = []string{
	".env",
	".env.*",
}

func Scanner(path string) ([]string, error) {
	files := []string{}
	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}

		// Skip unwanted directories
		if d.IsDir() && skipDirs[d.Name()] {
			return fs.SkipDir
		}

		// Only print files
		if !d.IsDir() {
			base := filepath.Base(path)
			if shouldSkipFile(base) {
				return nil
			}
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return files, errors.New("something went wrong")
	}

	return files, nil
}

func shouldSkipFile(name string) bool {
	for _, pattern := range skipFilePatterns {
		match, _ := filepath.Match(pattern, name)

		if match {
			return true
		}
	}
	return false
}
