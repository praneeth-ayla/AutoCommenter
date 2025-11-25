package scanner

import (
	"os"
	"path/filepath"
)

// GetProjectRoot attempts to find the project's root directory by searching for a "go.mod" file.
// It starts from the current working directory and traverses up the directory tree.
func GetProjectRoot() string {
	dir, err := os.Getwd()
	if err != nil {
		// If getting the current working directory fails, assume the root is the current directory.
		return "."
	}

	for {
		// Check if go.mod exists in the current directory.
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			// Found go.mod, so this is the project root.
			return dir
		}

		parent := filepath.Dir(dir)

		// If the parent directory is the same as the current directory, we've reached the filesystem root.
		if parent == dir {
			return "." // fallback to current directory if no go.mod is found up to the root.
		}

		// Move up to the parent directory.
		dir = parent
	}
}