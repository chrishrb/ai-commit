package client

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/briandowns/spinner"
	"github.com/chrishrb/ai-commit/pkg/config"
	"github.com/chrishrb/ai-commit/pkg/git"
)

type Client interface {
  GenerateContent(ctx context.Context, diff string, branchIssue string, streamingFn func(ctx context.Context, chunk []byte) error,) (string, error)
}

func BuildCommitMessage() (string, error) {
  var err error

  // Get diff
  diff, err := git.GetDiff(config.C.IgnoredFiles)
	if err != nil || diff == "" {
		return "", err
	}
  slog.Debug("BuildCommitMessage", "diff", diff)

	// Get issue number from branch, e.g. ISSUE-123
	var issue string
	if config.C.Plugins.AddBranchPrefix {
		issue, err = git.BranchIssue()
		if err != nil {
			return "", err
		}
	}
  slog.Debug("BuildCommitMessage", "issue", issue)

	// Generate commit message
  var c Client
  switch config.C.Client.Provider {
    case "ollama": c = NewOllamaClient(config.C)
    case "copilot": c = NewCopilotClient(config.C)
    default: return "", errors.New("invalid provider, only copilot and ollama are supported")
  }

  s := spinner.New(spinner.CharSets[14], 100 * time.Millisecond)
  s.Start()
  s.Prefix = "💬 "
  content, err := c.GenerateContent(context.Background(), diff, issue, nil)
  s.Stop()
  return content, err
}
