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
	// Gemini backend requires environment variables for API key.
	key := os.Getenv("GOOGLE_API_KEY")
	if key == "" { // Fallback to GEMINI_API_KEY if GOOGLE_API_KEY is not set.
		key = os.Getenv("GEMINI_API_KEY")
	}
	if key == "" { // Return an error if neither key is found.
		return fmt.Errorf("missing Gemini API key. Set GOOGLE_API_KEY or GEMINI_API_KEY")
	}

	return nil
}