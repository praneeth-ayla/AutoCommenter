package gemini

// AnalyzeFilesForCommentsSchema defines the JSON schema for the AI's response when analyzing files for comments.
var AnalyzeFilesForCommentsSchema = map[string]any{
	"type": "object",
	"properties": map[string]any{
		"files": map[string]any{
			"type":  "array",
			"items": map[string]any{"type": "string"},
		},
	},
	"required": []string{"files"},
}

// GenerateCommentsForFilesSchema defines the JSON schema for the AI's response when generating comments for files.
var GenerateCommentsForFilesSchema = map[string]any{
	"type": "object",
	"properties": map[string]any{
		"files": map[string]any{
			"type": "array",
			"items": map[string]any{
				"type": "object",
				"properties": map[string]any{
					"path": map[string]any{
						"type": "string",
					},
					"content": map[string]any{
						"type": "string",
					},
				},
				"required": []string{"path", "content"},
			},
		},
	},
	"required": []string{"files"},
}

var GenerateContextSchema = map[string]any{
	"type": "object",
	"properties": map[string]any{
		"path": map[string]any{
			"type": "string",
		},
		"file_name": map[string]any{
			"type": "string",
		},
		"exports": map[string]any{
			"type":  "array",
			"items": map[string]any{"type": "string"},
		},
		"imports": map[string]any{
			"type":  "array",
			"items": map[string]any{"type": "string"},
		},
		"summary": map[string]any{
			"type": "string",
		},
	},
	"required": []string{
		"path",
		"file_name",
		"exports",
		"imports",
		"summary",
	},
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
						"type": "string",
					},
					"file_name": map[string]any{
						"type": "string",
					},
					"exports": map[string]any{
						"type":  "array",
						"items": map[string]any{"type": "string"},
					},
					"imports": map[string]any{
						"type":  "array",
						"items": map[string]any{"type": "string"},
					},
					"summary": map[string]any{
						"type": "string",
					},
				},
				"required": []string{
					"path",
					"file_name",
					"exports",
					"imports",
					"summary",
				},
			},
		},
	},
	"required": []string{"files"},
}
