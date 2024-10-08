package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type client struct {
	Provider          string
	Model             string
	ContextWindowSize int
}

type prompts struct {
	Mission                    string
  OneLineSummaryExample string
  OneLineSummaryExampleWithTicketNumber string
	DiffInstructions           string
	MultiLineCommitGuidelines  string
	GeneralGuidelines          string
}

type plugins struct {
	AddBranchPrefix bool
}

type Config struct {
	Client                 client
	Prompts                prompts
	MultiLineCommitMessage bool
	IgnoredFiles           []string
	Plugins                plugins
}


var C Config = Config{
  Client: client{
    Provider:          "ollama",
    Model:             "llama3.2",
    ContextWindowSize: 12800,
  },
  Prompts: prompts{
    Mission: `You are provided with a git diff output that shows code changes. Your task is to generate a structured and descriptive commit message based on the following guidelines:
1. The commit message should have a short, one-line summary (50 characters or less), starting with:
  a) one of the following keywords:
  - feat: for new features or enhancements
  - fix: for bug fixes
  - refactor: for code restructuring without changing behavior
  - docs: for documentation changes
  - test: for adding or modifying tests
  - chore: for maintenance tasks (e.g., updating dependencies)
`,
    OneLineSummaryExample: "  Example: 'fix: resolve null pointer exception'",
    OneLineSummaryExampleWithTicketNumber: "  Example: 'fix: ISSUE-123 Resolve null pointer exception'",
    MultiLineCommitGuidelines: `
3. After the summary, include a detailed description explaining:
- What has changed and why.
- The issue the changes are addressing (if applicable).
- Any important implications for other parts of the codebase.`,
    GeneralGuidelines: `
4. Only output one commit message and no further explanations.
5. Use an imperative tone (e.g., 'Fix', 'Add', 'Update').
6. Ensure that the message is clear and concise, focusing on the intent of the changes rather than just describing the diff.
7. Don‘t use code to explain the changes.
`,
    DiffInstructions: "Here is the git diff output:\n",
  },
  Plugins: plugins{
    AddBranchPrefix: true,
  },
  MultiLineCommitMessage: true,
  IgnoredFiles:           []string{},
}

func ParseConfig() error {
	err := viper.Unmarshal(&C)
	if err != nil {
		return err
	}
	return nil
}

func (C *Config) BuildPrompt(issue string) string {
	var sb strings.Builder
	sb.WriteString(C.Prompts.Mission)
	if issue == "" {
    sb.WriteString(C.Prompts.OneLineSummaryExample)
	} else {
		sb.WriteString(fmt.Sprintf("  b) The ticket number '%s'.\n", issue))
    sb.WriteString(C.Prompts.OneLineSummaryExampleWithTicketNumber)
  }
	if C.MultiLineCommitMessage {
		sb.WriteString(C.Prompts.MultiLineCommitGuidelines)
	}
	sb.WriteString(C.Prompts.GeneralGuidelines)
	sb.WriteString(C.Prompts.DiffInstructions)
	return sb.String()
}
