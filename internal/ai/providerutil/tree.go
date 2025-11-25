package providerutil

import (
	"os"
	"path/filepath"
	"strings"
)

// BuildFileTree recursively walks a directory and builds a string
// representation of the file tree.
func BuildFileTree(root string) (string, error) {
	var builder strings.Builder

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // Propagate any error encountered during walking.
		}

		rel, _ := filepath.Rel(root, path) // Get the relative path from the root.
		if rel == "." {
			return nil // Skip the root directory itself.
		}

		depth := strings.Count(rel, string(filepath.Separator)) // Calculate directory depth for indentation.
		prefix := strings.Repeat("  ", depth)                  // Create indentation prefix.

		if info.IsDir() {
			builder.WriteString(prefix + info.Name() + "/\n") // Append directory name with a slash.
		} else {
			builder.WriteString(prefix + info.Name() + "\n") // Append file name.
		}

		return nil
	})

	if err != nil {
		return "", err // Return empty string and the error if walking failed.
	}

	return builder.String(), nil // Return the built file tree string and nil error.
}