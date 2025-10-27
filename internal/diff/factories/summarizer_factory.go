package factories

import (
	"github.com/alkowskey/commitlens/internal/diff/config"
	"github.com/alkowskey/commitlens/internal/diff/domain"
	"github.com/alkowskey/commitlens/internal/diff/infra"
	"github.com/openai/openai-go/v3"
)

func CreateSummarizer(modelProvider domain.ModelProvider) domain.DiffSumarizer {
	switch modelProvider {
	case domain.OpenAIModelProvider:
		return createOpenAiSummarizer()
	case domain.GroqModelProvider:
		return createGroqSummarizer()
	}

	panic("Illegal model provider")
}

func createOpenAiSummarizer() domain.DiffSumarizer {
	client := openai.NewClient()
	config := config.DefaultOpenAIConfig()
	return infra.NewOpenAIDiffSummarizer(&client, config)
}

func createGroqSummarizer() domain.DiffSumarizer {
	return infra.NewGroqDiffSummarizer()
}
