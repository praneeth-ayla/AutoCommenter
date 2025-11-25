package providerutil

import (
	"os"
	"path/filepath"
	"strings"
)

func BuildFileTree(root string) (string, error) {
	var builder strings.Builder

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		rel, _ := filepath.Rel(root, path)
		if rel == "." {
			return nil
		}

		depth := strings.Count(rel, string(filepath.Separator))
		prefix := strings.Repeat("  ", depth)

		if info.IsDir() {
			builder.WriteString(prefix + info.Name() + "/\n")
		} else {
			builder.WriteString(prefix + info.Name() + "\n")
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return builder.String(), nil
}
