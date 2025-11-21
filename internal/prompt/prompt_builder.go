package prompt

import (
	"fmt"
	"log"

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
