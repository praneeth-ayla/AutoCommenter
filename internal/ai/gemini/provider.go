package gemini

import "github.com/praneeth-ayla/AutoCommenter/internal/contextstore"

type GeminiProvider struct{}

func New() *GeminiProvider {
	return &GeminiProvider{}
}

func (g *GeminiProvider) GenerateContext(path string, content string) (contextstore.FileDetails, error) {
	// return a slice because the interface expects []FileDetails
	return contextstore.FileDetails{}, nil
}

func (g *GeminiProvider) GenerateComments(content string, contexts []contextstore.FileDetails) (string, error) {
	return "", nil
}
