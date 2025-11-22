package ai

import (
	"github.com/praneeth-ayla/AutoCommenter/internal/ai/gemini"
	"github.com/praneeth-ayla/AutoCommenter/internal/contextstore"
	"github.com/praneeth-ayla/AutoCommenter/internal/scanner"
)

type Provider interface {
	GenerateComments(content string, contexts []contextstore.FileDetails) (string, error)
	GenerateContextBatch(files []scanner.Data) ([]contextstore.FileDetails, error)
}

func NewProvider(name string) Provider {
	switch name {
	case "gemini":
		return gemini.New()
	}

	panic("unknown provider")
}
