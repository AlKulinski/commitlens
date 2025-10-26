package infra

import (
	"context"

	"github.com/openai/openai-go/v3"

	"github.com/alkowskey/commitlens/internal/diff/config"
	"github.com/alkowskey/commitlens/internal/diff/domain"
	diffUtils "github.com/alkowskey/commitlens/internal/diff/utils"
)

type OpenAIDiffSummarizer struct {
	OpenAIClient *openai.Client
	Config       *config.OpenAIConfig
}

func NewOpenAIDiffSummarizer(client *openai.Client, config *config.OpenAIConfig) *OpenAIDiffSummarizer {
	return &OpenAIDiffSummarizer{
		OpenAIClient: client,
		Config:       config,
	}
}

func (s *OpenAIDiffSummarizer) Summarize(ctx context.Context, diff []domain.DiffResult) (string, error) {
	chatCompletion, err := s.OpenAIClient.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Model: s.Config.Model,
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(s.Config.MasterPrompt),
			openai.UserMessage(diffUtils.FormatDiffs(diff)),
		},
	})

	if err != nil {
		return "", err
	}

	summary := chatCompletion.Choices[0].Message.Content
	return summary, nil
}
