package prompt

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/alpkeskin/gotoon"
	"github.com/praneeth-ayla/AutoCommenter/internal/contextstore"
)

func BuildFileContextPrompt(path string, content string) string {
	return fmt.Sprintf(TemplateFileContext, path, content)
}

func BuildCommentPrompt(style string, content string, contextData string) (string, error) {
	data := map[string]interface{}{
		"content": content,
		"context": contextData,
	}

	encoded, err := gotoon.Encode(
		data,
		gotoon.WithIndent(0),
		gotoon.WithDelimiter("\t"),
	)
	if err != nil {
		return "", fmt.Errorf("prompt encoding failed: %w", err)
	}

	switch style {
	case "minimalist":
		return fmt.Sprintf(TemplateMinimalist, encoded), nil
	case "explanatory":
		return fmt.Sprintf(TemplateExplanatory, encoded), nil
	case "detailed":
		return fmt.Sprintf(TemplateDetailed, encoded), nil
	case "docstring":
		return fmt.Sprintf(TemplateDocstring, encoded), nil
	case "inline-only":
		return fmt.Sprintf(TemplateInlineOnly, encoded), nil
	default:
		return "", errors.New("unknown style: supported styles are minimalist, explanatory, detailed, docstring, inline-only")
	}
}

func BuildFixesPrompt(original string, aiOutput string) string {
	return fmt.Sprintf(TemplateApplyFixes, original, aiOutput)
}

func BuildReadmePrompt(contexts []contextstore.FileDetails, existingReadme string, fileTree string) (string, error) {
	var sb strings.Builder

	for _, c := range contexts {
		j, err := json.Marshal(c)
		if err != nil {
			return "", fmt.Errorf("context marshal error: %w", err)
		}
		sb.Write(j)
		sb.WriteByte('\n')
	}

	contextStr := sb.String()
	readmeStr := strings.TrimSpace(existingReadme)
	treeStr := strings.TrimSpace(fileTree)

	return fmt.Sprintf(TemplateReadme, contextStr, treeStr, readmeStr), nil
}
