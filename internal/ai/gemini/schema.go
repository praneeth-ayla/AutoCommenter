package gemini

var GenerateCommentsForFilesSchema = map[string]any{
	"type": "object",
	"properties": map[string]any{
		"files": map[string]any{
			"type": "array",
			"items": map[string]any{
				"type": "object",
				"properties": map[string]any{
					"path": map[string]any{
						"type": "string", // Path of the file to generate comments for.
					},
					"content": map[string]any{
						"type": "string", // Content of the file.
					},
				},
				"required": []string{"path", "content"}, // Both path and content are mandatory for each file.
			},
		},
	},
	"required": []string{"files"}, // The 'files' field is required at the top level.
}

var GenerateContextBatchSchema = map[string]any{
	"type": "object",
	"properties": map[string]any{
		"files": map[string]any{
			"type": "array",
			"items": map[string]any{
				"type": "object",
				"properties": map[string]any{
					"path": map[string]any{
						"type": "string", // Path of the file.
					},
					"file_name": map[string]any{
						"type": "string", // Name of the file.
					},
					"exports": map[string]any{
						"type":  "array",
						"items": map[string]any{"type": "string"}, // List of exported identifiers from the file.
					},
					"imports": map[string]any{
						"type":  "array",
						"items": map[string]any{"type": "string"}, // List of imported packages.
					},
					"summary": map[string]any{
						"type": "string", // Summary of the file's purpose.
					},
				},
				"required": []string{
					"path",
					"file_name",
					"exports",
					"imports",
					"summary",
				}, // All these fields are mandatory for context generation.
			},
		},
	},
	"required": []string{"files"}, // The 'files' field is required.
}