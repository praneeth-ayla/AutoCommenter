package prompt

import (
	"fmt"
	"log"
	"strings"

	"github.com/alpkeskin/gotoon"
)

// BuildGenerateCommentsForFilesPrompt constructs the AI prompt to generate comments for given files.
func BuildGenerateCommentsForFilesPrompt(files []string) string {
	data := map[string]interface{}{
		"files": files,
	}

	// Encode the file content into a structured string format for the AI prompt.
	encoded, err := gotoon.Encode(
		data,
		gotoon.WithIndent(0),       // no extra spaces
		gotoon.WithDelimiter("\t"), // tabs tokenize better
	)
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf(GenerateCommentsForFiles, encoded)
}

func BuildFileContextPrompt(path string, content string) string {
	var b strings.Builder

	b.WriteString("Return JSON for this file using the schema fields path, file_name, summary, exports, and imports.\n")
	b.WriteString("Identify exports and imports from the content.\n\n")

	b.WriteString("Path: ")
	b.WriteString(path)
	b.WriteString("\n\n")

	b.WriteString("Content:\n")
	b.WriteString(content)

	return b.String()
}
