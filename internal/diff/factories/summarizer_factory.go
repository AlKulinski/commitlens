package factories

import (
	"github.com/alkowskey/commitlens/internal/diff/config"
	"github.com/alkowskey/commitlens/internal/diff/domain"
	"github.com/alkowskey/commitlens/internal/diff/infra"
	"github.com/openai/openai-go/v3"
)

func CreateSummarizer() domain.DiffSumarizer {
	client := openai.NewClient()
	config := config.DefaultOpenAIConfig()
	return infra.NewOpenAIDiffSummarizer(&client, config)
}
