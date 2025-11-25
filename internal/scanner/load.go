package scanner

import (
	"fmt"
	"os"
)

// Load reads the content of multiple files.
func Load(files []Info) []Data {

	filesContent := []Data{}
	for _, file := range files {
		// Read the entire file content.
		fileContent, err := os.ReadFile(file.Path)
		if err != nil {
			// Log the error and skip to the next file if reading fails.
			fmt.Println("Error reading file:", err)
			continue
		}
		filesContent = append(filesContent, Data{
			Path:    file.Path,
			Content: string(fileContent)}, // Convert byte slice to string.
		)
	}
	return filesContent

}

// LoadSingle reads the content of a single file.
func LoadSingle(file Info) Data {

	fileContent, err := os.ReadFile(file.Path)
	if err != nil {
		// Log the error and return an empty Data struct if reading fails.
		fmt.Println("Error reading file:", err)
		return Data{}
	}

	fileData := Data{
		Path:    file.Path,
		Content: string(fileContent), // Convert byte slice to string.
	}

	return fileData
}