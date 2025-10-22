package infra

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alkowskey/commitlens/internal/common/utils"
	"github.com/alkowskey/commitlens/internal/diff/config"
	"github.com/alkowskey/commitlens/internal/diff/domain"
	diffUtils "github.com/alkowskey/commitlens/internal/diff/utils"
)

type groqMessage struct {
	Role    string
	Content string
}

type groqRequest struct {
	Model    string
	Messages []groqMessage
}

type groqChoice struct {
	Message groqMessage
}

type groqResponse struct {
	Choices []groqChoice
}

type GroqDiffSummarizer struct {
}

func NewGroqDiffSummarizer() *GroqDiffSummarizer {
	return &GroqDiffSummarizer{}
}

func (s *GroqDiffSummarizer) Summarize(ctx context.Context, diff []domain.DiffResult) (string, error) {
	return s.requestGroq(ctx, diff)
}

func (s *GroqDiffSummarizer) requestGroq(ctx context.Context, diff []domain.DiffResult) (string, error) {
	req := s.buildHttpGroqRequest(ctx, diff)
	body := utils.ExecuteHTTPRequest(req)

	response, nil := s.parseGroqResponse(body)
	return response, nil
}

func (s *GroqDiffSummarizer) buildHttpGroqRequest(ctx context.Context, diff []domain.DiffResult) *http.Request {
	config := config.GetGroqConfig()

	reqBody := groqRequest{
		Model: config.Model,
		Messages: []groqMessage{
			{
				Role:    "system",
				Content: config.MasterPrompt,
			},
			{
				Role:    "user",
				Content: diffUtils.FormatDiffs(diff),
			},
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		panic("failed to marshal request")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", config.Url, bytes.NewBuffer(jsonData))
	if err != nil {
		panic("failed to create request")
	}

	utils.SetHTTPHeaders(req, map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", config.ApiKey),
	})
	return req
}

func (s *GroqDiffSummarizer) parseGroqResponse(body []byte) (string, error) {
	var groqResp groqResponse
	if err := json.Unmarshal(body, &groqResp); err != nil {
		return "", err
	}

	if len(groqResp.Choices) == 0 {
		return "", fmt.Errorf("no response choices returned")
	}

	return groqResp.Choices[0].Message.Content, nil
}
