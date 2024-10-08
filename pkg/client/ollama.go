package client

import (
	"context"

	"github.com/chrishrb/ai-commit/pkg/config"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

type OllamaClient struct {
	config config.Config
}

func NewOllamaClient(c config.Config) *OllamaClient {
  return &OllamaClient{
    config: c,
  }
}

func (c *OllamaClient) GenerateContent(
	ctx context.Context,
	diff string,
	branchIssue string,
	streamingFn func(ctx context.Context, chunk []byte) error,
) (string, error) {
	llm, err := ollama.New(ollama.WithModel(config.C.Client.Model), ollama.WithRunnerNumCtx(config.C.Client.ContextWindowSize))
	if err != nil {
		return "", err
	}

	prompt := config.C.BuildPrompt(branchIssue)
	content := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, prompt),
		llms.TextParts(llms.ChatMessageTypeTool, diff),
	}
	response, err := llm.GenerateContent(ctx, content, llms.WithStreamingFunc(streamingFn))

	if err != nil {
		return "", err
	}

	return response.Choices[0].Content, nil
}
