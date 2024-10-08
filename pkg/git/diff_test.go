package git

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetStagedFiles(t *testing.T) {
	origShellCommandFunc := shellCommandFunc
	defer func() { shellCommandFunc = origShellCommandFunc }()

	t.Run("returns staged files", func(t *testing.T) {
		shellCommandCalled := false
		shellCommandFunc = func(name string, args ...string) commandExecutor {
			shellCommandCalled = true

			// Careful: relies on implementation details this could
			// make the test fragile.
			assert.Equal(t, "git", name, "command name")
			assert.Len(t, args, 4, "command args")
			assert.Equal(t, "diff", args[0], "1st command arg")
			assert.Equal(t, "--cached", args[1], "2st command arg")
			assert.Equal(t, "--relative", args[2], "3st command arg")
			assert.Equal(t, "--name-only", args[3], "4st command arg")

			// Careful: if the stub deviates from how the system under
			// test works this could generate false positives.
			return &mockCommandExecutor{output: "cmd/root.go\npkg/client/client.go"}
		}

		stagedFiles, err := getStagedFiles()
		if assert.NoError(t, err) {
			assert.Equal(t, []string{"cmd/root.go", "pkg/client/client.go"}, stagedFiles, "staged files")
		}

		assert.True(t, shellCommandCalled, "shell command called")
	})

	t.Run("empty staged files", func(t *testing.T) {
		shellCommandCalled := false
		shellCommandFunc = func(name string, args ...string) commandExecutor {
			shellCommandCalled = true
			return &mockCommandExecutor{output: ""}
		}

		stagedFiles, err := getStagedFiles()
		if assert.NoError(t, err) {
			assert.Equal(t, []string{}, stagedFiles, "staged files")
		}

		assert.True(t, shellCommandCalled, "shell command called")
	})
}
