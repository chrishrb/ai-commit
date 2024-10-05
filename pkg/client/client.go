package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/chrishrb/ai-commit/pkg/git"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func BuildCommitMessage() (string, error) {
  err := parseConfig()
  if err != nil {
    return "", err
  }

  var sb strings.Builder

  // Add branchIssuerNumber, e.g. ISSUE-123
  var prefix string
  if (Config.AddBranchPrefix) {
    prefix, err := git.BranchIssuerNumber()
    if err != nil {
      return "", err
    }
    sb.WriteString(prefix + " ")
  }

  // Generate commit message
  res, err := llmResponse(prefix)
  if err != nil {
    return "", err
  }
  sb.WriteString(res)
  return sb.String(), err
}

func llmResponse(branchIssuerNumber string) (string, error) {
  llm, err := ollama.New(ollama.WithModel("llama3.2"), ollama.WithRunnerNumCtx(Config.Client.ContextWindowSize))
  if err != nil {
    return "", err
  }

  ctx := context.Background()

  prompt := buildPrompt(Config, branchIssuerNumber)
  diff, err := git.GetDiff(Config.IgnoredFiles)
  if err != nil {
    return "", err
  }
  content := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, prompt),
		llms.TextParts(llms.ChatMessageTypeTool, diff),
	}
  completion, err := llm.GenerateContent(ctx, content, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		fmt.Print(string(chunk))
		return nil
	}))

  if err != nil {
    return "", err
  }
  
  var sb strings.Builder
  for _, comp := range completion.Choices {
    sb.WriteString(comp.Content)
  }
  return sb.String(), nil
}

func buildPrompt(c config, branchIssuerNumber string) string {
  var sb strings.Builder
  sb.WriteString(c.Prompts.Mission + "\n")
  if branchIssuerNumber == "" {
    sb.WriteString(c.Prompts.ConventionalCommitKeywords + "\n")
  }
  sb.WriteString(c.Prompts.DiffInstructions + "\n")
  sb.WriteString(c.Prompts.GeneralGuidelines + "\n")

  if (c.OneLineCommitMessage) {
    sb.WriteString(c.Prompts.OneLineCommitGuidelines + "\n")
  }

  fmt.Println(sb.String())

  return sb.String()
}
