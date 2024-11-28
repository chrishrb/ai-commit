package client

import (
	"context"

	"github.com/chrishrb/ai-commit/pkg/config"
	copilot "github.com/stong1994/github-copilot-api"
)

type CopilotClient struct {
	config config.Config
}

func NewCopilotClient(c config.Config) *CopilotClient {
	return &CopilotClient{
		config: c,
	}
}

func (c *CopilotClient) GenerateContent(
	ctx context.Context,
	diff string,
	branchIssue string,
	streamingFn func(ctx context.Context, chunk []byte) error,
) (string, error) {
	client, err := copilot.NewCopilot(copilot.WithCompletionModel(c.config.Client.Model))
	if err != nil {
		return "", err
	}

	prompt := config.C.BuildPrompt(branchIssue)
	response, err := client.CreateCompletion(ctx, &copilot.CompletionRequest{
		Messages: []copilot.Message{
			{
				Role:    "system",
				Content: prompt,
			},
			{
				Role:    "user",
				Content: diff,
			},
		},
		StreamingFunc: streamingFn,
	})

	if err != nil {
		return "", err
	}

	return response.Choices[0].Message.Content, nil
}
