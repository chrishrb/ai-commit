package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/chrishrb/ai-commit/pkg/git"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func GetResponse() (*string, error) {
  err := parseConfig()
  if err != nil {
    return nil, err
  }

  llm, err := ollama.New(ollama.WithModel("llama3.2"))
  if err != nil {
    return nil, err
  }

  prompt, err := buildPrompt(Config)
  if err != nil {
    return nil, err
  }

  ctx := context.Background()
  completion, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
  if err != nil {
    return nil, err
  }

  return &completion, nil
}

func buildPrompt(c config) (string, error) {
  var sb strings.Builder
  diff, err := git.GetDiff(Config.IgnoredFiles)
  if err != nil {
    return "", err
  }
  sb.WriteString(diff)
  sb.WriteString("---------------------\n")

  sb.WriteString(c.Prompts.Mission + "\n")
  sb.WriteString(c.Prompts.ConventionalCommitKeywords + "\n")
  sb.WriteString(c.Prompts.GeneralGuidelines + "\n")
  sb.WriteString(c.Prompts.DiffInstructions + "\n")

  if (c.OneLineCommitMessage) {
    sb.WriteString(c.Prompts.OneLineCommitGuidelines + "\n")
  }

  fmt.Println(sb.String())

  return sb.String(), nil
}
