package scanner

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var skipDirs = map[string]bool{
	"node_modules": true,
	".git":         true,
	".next":        true,
	"build":        true,
	"dist":         true,
	"migrations":   true,
	"prisma":       true,
}

var allowedExt = map[string]bool{
	".ts":  true,
	".tsx": true,
	".js":  true,
	".jsx": true,
	".go":  true,
	".py":  true,
}

// Scan returns all allowed files with metadata
func Scan(root string) ([]Info, error) {
	abs, err := filepath.Abs(root)
	if err != nil {
		return nil, err
	}

	var files []Info

	err = filepath.WalkDir(abs, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Walk error:", err)
			// Skip the file/directory that caused the error but continue walking
			return nil
		}

		if d.IsDir() {
			rel, _ := filepath.Rel(abs, path)
			if skipDirs[rel] {
				// Skip this directory and its contents
				return filepath.SkipDir
			}
			return nil
		}

		ext := filepath.Ext(path)
		if !allowedExt[ext] {
			// Skip files with disallowed extensions
			return nil
		}

		info, statErr := d.Info()
		if statErr != nil {
			return nil
		}

		lines := countLines(path)

		files = append(files, Info{
			Path:  filepath.Clean(path),
			Name:  info.Name(),
			Size:  info.Size(),
			Lines: lines,
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Path < files[j].Path
	})

	return files, nil
}

// Filter only returns comment-worthy files
func FilterFilesNeedingComments(files []Info) []Info {
	var result []Info

	for _, f := range files {
		p := f.Path

		if strings.HasSuffix(p, ".d.ts") {
			continue // Skip declaration files
		}
		if strings.Contains(p, "/ui/") ||
			strings.Contains(p, "/types/") ||
			strings.Contains(p, "/__tests__/") ||
			strings.Contains(p, "/.storybook/") {
			continue // Skip files in specific directories
		}
		if strings.HasPrefix(p, "next-env.d.ts") ||
			strings.HasPrefix(p, "next.config") ||
			strings.Contains(p, "seed.") {
			continue // Skip specific files or files containing specific patterns
		}

		result = append(result, f)
	}

	return result
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
		count += bytes.Count(buf[:c], []byte{'\n'}) // Count newline characters
		if err != nil {
			// Break loop on any read error, including EOF
			break
		}
	}
	return count
}