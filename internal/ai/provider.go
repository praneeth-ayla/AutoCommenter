package ai

import (
	"github.com/praneeth-ayla/AutoCommenter/internal/ai/gemini"
	"github.com/praneeth-ayla/AutoCommenter/internal/contextstore"
)

type Provider interface {
	GenerateContext(path string, content string) (contextstore.FileDetails, error)
	GenerateComments(content string, contexts []contextstore.FileDetails) (string, error)
}

func NewProvider(name string) Provider {
	switch name {
	case "gemini":
		return gemini.New()
	}

	panic("unknown provider")
}
