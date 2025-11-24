package gemini

import (
	"fmt"
	"os"
)

type GeminiProvider struct{}

func New() *GeminiProvider {
	return &GeminiProvider{}
}

func (g *GeminiProvider) Validate() error {
	// Gemini backend required environment
	key := os.Getenv("GOOGLE_API_KEY")
	if key == "" {
		key = os.Getenv("GEMINI_API_KEY")
	}
	if key == "" {
		return fmt.Errorf("missing Gemini API key. Set GOOGLE_API_KEY or GEMINI_API_KEY")
	}

	return nil
}
