package client

import (
	"github.com/spf13/viper"
)

type client struct {
	Provider          string
	ApiKey            string
	ApiUrl            string
	ContextWindowSize int
}

type prompts struct {
	ConventionalCommitKeywords string
	Mission                    string
	DiffInstructions           string
	OneLineCommitGuidelines    string
	GeneralGuidelines          string
}

type config struct {
	Client               client
	Prompts              prompts
	OneLineCommitMessage bool
	IgnoredFiles         []string
	AddBranchPrefix      bool
}

var (
	Config config = config{
		Client: client{
			Provider:          "ollama",
			ApiKey:            "",
			ApiUrl:            "",
			ContextWindowSize: 12800,
		},
		Prompts: prompts{
			Mission:                    "You are to act as an author of a commit message in git. Your mission is to create clean and comprehensive commit messages as per the conventional commit convention. Don't add any descriptions to the output, only commit message.",
			ConventionalCommitKeywords: "Preface the commit message headline with the commit keywords: fix, feat, build, chore, ci, docs, style, refactor, perf, test.",
			DiffInstructions:           "An output of 'git diff --staged' command is provided, and you are to convert it into a commit message.",
			OneLineCommitGuidelines:    "Craft a concise commit message that encapsulates all changes made, with an emphasis on the primary updates. If the modifications share a common theme or scope, mention it succinctly; otherwise, leave the scope out to maintain focus. The goal is to provide a clear and unified overview of the changes in a one single message, without diverging into a list of commit per file change.",
			GeneralGuidelines: `Use the present tense. Lines must not be longer than 74 characters. Start with a upper case letter. Only output ONE commit message. Use english for the commit message. Output in the following format:
feat: My commit message

This is a description of my commit. Some things have changed. 
`,
		},
		OneLineCommitMessage: false,
		IgnoredFiles:         []string{},
		AddBranchPrefix:      true,
	}
)

func parseConfig() error {
	err := viper.Unmarshal(&Config)
	if err != nil {
		return err
	}
	return nil
}
