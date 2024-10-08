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

	// Get branchIssuerNumber, e.g. ISSUE-123
	var prefix string
	if Config.Plugins.AddBranchPrefix {
		prefix, err = git.BranchIssuerNumber()
		if err != nil {
			return "", err
		}
	}

	// Generate commit message
	res, err := llmResponse(prefix)
	if err != nil || res == "" {
		return "", err
	}
	return res, err
}

func llmResponse(branchIssuerNumber string) (string, error) {
	llm, err := ollama.New(ollama.WithModel(Config.Client.Model), ollama.WithRunnerNumCtx(Config.Client.ContextWindowSize))
	if err != nil {
		return "", err
	}

	ctx := context.Background()

	prompt := buildPrompt(Config, branchIssuerNumber)
	diff, err := git.GetDiff(Config.IgnoredFiles)
	if err != nil || diff == "" {
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
	sb.WriteString(c.Prompts.Mission)
	if branchIssuerNumber == "" {
    sb.WriteString(c.Prompts.OneLineSummaryExample)
	} else {
		sb.WriteString(fmt.Sprintf("  b) The ticket number `%s`.", branchIssuerNumber))
    sb.WriteString(c.Prompts.OneLineSummaryExampleWithTicketNumber)
  }
	if c.MultiLineCommitMessage {
		sb.WriteString(c.Prompts.MultiLineCommitGuidelines)
	}
	sb.WriteString(c.Prompts.GeneralGuidelines)
	sb.WriteString(c.Prompts.DiffInstructions)

  fmt.Println(sb.String())

	return sb.String()
}
