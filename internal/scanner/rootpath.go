package scanner

import (
	"os"
	"path/filepath"
)

func GetProjectRoot() string {
	dir, err := os.Getwd()
	if err != nil {
		return "."
	}

	for {
		// check if go.mod exists here
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}

		parent := filepath.Dir(dir)

		// reached filesystem root
		if parent == dir {
			return "." // fallback
		}

		dir = parent
	}
}
