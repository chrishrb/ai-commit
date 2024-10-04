package client

import (
	"github.com/spf13/viper"
)

type client struct {
	Provider string
	ApiKey   string
	ApiUrl   string
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
}

var (
	Config config = config{
		Client: client{
			Provider: "ollama",
			ApiKey:   "",
			ApiUrl:   "",
		},
		Prompts: prompts{
			Mission: `You are to act as an author of a commit message in git. Your mission is to create clean and 
comprehensive commit messages as per the conventional commit convention. Don't add any descriptions to the output, only commit message.`,
			ConventionalCommitKeywords: `Do not preface the commit with anything, except for the conventional commit keywords: fix, feat, build, 
chore, ci, docs, style, refactor, perf, test.`,
			DiffInstructions: "Above the line is an output of 'git diff --staged' command, and you are to convert it into a commit message.",
			OneLineCommitGuidelines: `Craft a concise commit message that encapsulates all changes made, with an emphasis on 
the primary updates. If the modifications share a common theme or scope, mention it succinctly; otherwise, leave 
the scope out to maintain focus. The goal is to provide a clear and unified overview of the changes in a one single 
message, without diverging into a list of commit per file change.`,
			GeneralGuidelines: "Use the present tense. Lines must not be longer than 74 characters. Use english for the commit message.",
		},
		OneLineCommitMessage: false,
		IgnoredFiles:         []string{},
	}
)

func parseConfig() error {
	err := viper.Unmarshal(&Config)
	if err != nil {
		return err
	}
	return nil
}
