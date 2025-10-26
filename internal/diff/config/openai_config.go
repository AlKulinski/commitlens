package config

import (
	"os"

	"github.com/alkowskey/commitlens/internal/diff/prompts"
	"github.com/openai/openai-go/v3"
)

type OpenAIConfig struct {
	ApiKey       string
	MasterPrompt string
	Model        openai.ChatModel
}

func DefaultOpenAIConfig() *OpenAIConfig {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		panic("OPENAI_API_KEY environment variable not set")
	}
	model := openai.ChatModelGPT5Nano
	return &OpenAIConfig{
		ApiKey:       apiKey,
		MasterPrompt: prompts.DiffSummarizerSystemPrompt,
		Model:        model,
	}
}
