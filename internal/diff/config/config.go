package config

import (
	"os"

	"github.com/alkowskey/commitlens/internal/diff/prompts"
)

type groqConfig struct {
	ApiKey       string
	Model        string
	Url          string
	MasterPrompt string
}

func GetGroqConfig() *groqConfig {
	apiKey := os.Getenv("GROQ_API_KEY")
	if apiKey == "" {
		panic("GROQ_API_KEY environment variable not set")
	}

	model := os.Getenv("GROQ_MODEL")

	if model == "" {
		panic("GROQ_MODEL environment variable not set")
	}

	url := os.Getenv("GROQ_API_URL")
	if url == "" {
		panic("GROQ_API_URL environment variable not set")
	}

	return &groqConfig{
		ApiKey:       apiKey,
		Model:        model,
		Url:          url,
		MasterPrompt: prompts.DiffSummarizerSystemPrompt,
	}
}
