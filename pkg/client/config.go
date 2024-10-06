package client

import (
	"github.com/spf13/viper"
)

type client struct {
	Provider          string
	Model             string
	ApiKey            string
	ApiUrl            string
	ContextWindowSize int
}

type prompts struct {
	ConventionalCommitKeywords string
	Mission                    string
	DiffInstructions           string
	MultiLineCommitGuidelines  string
	GeneralGuidelines          string
}

type config struct {
	Client                 client
	Prompts                prompts
	MultiLineCommitMessage bool
	IgnoredFiles           []string
	AddBranchPrefix        bool
}

var (
	Config config = config{
		Client: client{
			Provider:          "ollama",
			Model:             "llama3.2",
			ApiKey:            "",
			ApiUrl:            "",
			ContextWindowSize: 12800,
		},
		Prompts: prompts{
			Mission: `You are provided with a git diff output that shows code changes. Your task is to generate a structured and descriptive commit message based on the following guidelines:
1. The commit message should have a short, one-line summary (50 characters or less)`,
			ConventionalCommitKeywords: ` starting with one of the following keywords:
  - feat: for new features or enhancements
  - fix: for bug fixes
  - refactor: for code restructuring without changing behavior
  - docs: for documentation changes
  - test: for adding or modifying tests
  - chore: for maintenance tasks (e.g., updating dependencies)`,
			MultiLineCommitGuidelines: `
2. After the summary, include a detailed description explaining:
  - What has changed and why.
  - The issue the changes are addressing (if applicable).
  - Any important implications for other parts of the codebase.`,
			GeneralGuidelines: `
3. Only output one commit message and no further explanations.
4. Use an imperative tone (e.g., 'Fix', 'Add', 'Update').
5. Ensure that the message is clear and concise, focusing on the intent of the changes rather than just describing the diff.
6. Don‘t use code to explain the changes.
`,
			DiffInstructions: "Here is the git diff output:",
		},
		MultiLineCommitMessage: true,
		IgnoredFiles:           []string{},
		AddBranchPrefix:        true,
	}
)

func parseConfig() error {
	err := viper.Unmarshal(&Config)
	if err != nil {
		return err
	}
	return nil
}
