package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/urfave/cli/v3"
)

type groqMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type groqRequest struct {
	Model    string        `json:"model"`
	Messages []groqMessage `json:"messages"`
}

type groqChoice struct {
	Message groqMessage `json:"message"`
}

type groqResponse struct {
	Choices []groqChoice `json:"choices"`
}

func newGroqCmd() *cli.Command {
	return &cli.Command{
		Name:  "groq",
		Usage: "Groq command",
		Commands: []*cli.Command{
			newGroqChatCmd(),
		},
	}
}

func newGroqChatCmd() *cli.Command {
	return &cli.Command{
		Name:    "chat",
		Usage:   "Chat command",
		Aliases: []string{"c"},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			apiKey := os.Getenv("GROQ_API_KEY")
			if apiKey == "" {
				return fmt.Errorf("GROQ_API_KEY environment variable not set")
			}

			reqBody := groqRequest{
				Model: "openai/gpt-oss-20b",
				Messages: []groqMessage{
					{
						Role:    "user",
						Content: "What is the meaning of life?",
					},
				},
			}

			jsonData, err := json.Marshal(reqBody)
			if err != nil {
				return fmt.Errorf("failed to marshal request: %w", err)
			}

			req, err := http.NewRequestWithContext(ctx, "POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(jsonData))
			if err != nil {
				return fmt.Errorf("failed to create request: %w", err)
			}

			req.Header.Set("Authorization", "Bearer "+apiKey)
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				return fmt.Errorf("failed to send request: %w", err)
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return fmt.Errorf("failed to read response: %w", err)
			}

			if resp.StatusCode != http.StatusOK {
				return fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(body))
			}

			var groqResp groqResponse
			if err := json.Unmarshal(body, &groqResp); err != nil {
				return fmt.Errorf("failed to unmarshal response: %w", err)
			}

			if len(groqResp.Choices) == 0 {
				return fmt.Errorf("no response choices returned")
			}

			println(groqResp.Choices[0].Message.Content)
			return nil
		},
	}
}
