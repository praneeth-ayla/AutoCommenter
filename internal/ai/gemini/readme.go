package gemini

import (
	"context"

	"github.com/praneeth-ayla/AutoCommenter/internal/ai/providerutil"
	"github.com/praneeth-ayla/AutoCommenter/internal/contextstore"
	"github.com/praneeth-ayla/AutoCommenter/internal/prompt"
	"google.golang.org/genai"
)

func (g *GeminiProvider) GenerateReadme(contexts []contextstore.FileDetails, existingReadme string) (string, error) {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		return "", err
	}

	promptText, err := prompt.BuildReadmePrompt(contexts, existingReadme)
	if err != nil {
		return "", err
	}

	config := &genai.GenerateContentConfig{
		SystemInstruction: &genai.Content{
			Parts: []*genai.Part{
				{Text: prompt.SystemInstructionReadme},
			},
		},
		ResponseMIMEType: "text/plain",
	}

	input := []*genai.Content{
		{Parts: []*genai.Part{{Text: promptText}}},
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-pro",
		input,
		config,
	)
	if err != nil {
		return "", err
	}

	out := result.Text()

	out = providerutil.StripCodeFences(out)
	// out = providerutil.NormalizeMarkdownHeaders(out)

	return out, nil
}
