package scanner

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
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

func Scanner(path string) ([]FileInfo, error) {
	files := []FileInfo{}
	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}

		// Skip unwanted directories
		if d.IsDir() && skipDirs[d.Name()] {
			return fs.SkipDir
		}

		if !d.IsDir() {
			base := filepath.Base(path)
			if shouldSkipFile(base) {
				return nil
			}

			info, _ := d.Info()
			lines := countLines(path)

			file := FileInfo{
				Path:  path,
				Name:  base,
				Size:  info.Size(),
				Lines: lines,
			}

			files = append(files, file)
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

func countLines(path string) int {
	file, err := os.Open(path)
	if err != nil {
		return 0
	}
	defer file.Close()

	buf := make([]byte, 32*1024)
	count := 0
	for {
		c, err := file.Read(buf)
		count += bytes.Count(buf[:c], []byte{'\n'})
		if err != nil {
			break
		}
	}
	return count
}
