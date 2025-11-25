package scanner

import (
	"fmt"
	"os"
)

// WriteFile writes the given content to the specified file path.
// It returns an error if the file cannot be written.
func WriteFile(filePath, content string) error {
	// Use os.WriteFile to write the content as bytes to the file.
	// 0644 is the file permission mode, making the file readable and writable by the owner, and readable by others.
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		// Wrap the error with a more descriptive message.
		return fmt.Errorf("failed to write file: %w", err)
	}
	// Return nil to indicate success.
	return nil
}