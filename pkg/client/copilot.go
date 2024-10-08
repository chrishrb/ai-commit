package client

import (
	"context"
	"errors"

	"github.com/chrishrb/ai-commit/pkg/config"
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

	// TODO: implement
	return "", errors.ErrUnsupported
}
