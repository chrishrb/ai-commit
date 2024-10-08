package git

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_issueWithBranchName(t *testing.T) {
	tests := []struct {
		branchName string
		expected   string
	}{
		{"feature/ISSUE-123", "ISSUE-123"},
		{"bugfix/BUG-456", "BUG-456"},
		{"release/REL-789", "REL-789"},
		{"no-match-branch", ""},
		{"", ""},
		{"  feature/ISSUE-123  ", "ISSUE-123"},
		{"main", ""},
	}

	for _, test := range tests {
		result, _ := issueWithBranchName(test.branchName)
		if result != test.expected {
			t.Errorf("For branch name '%s', expected '%s' but got '%s'", test.branchName, test.expected, result)
		}
	}
}

func Test_BranchName(t *testing.T) {
  origShellCommandFunc := shellCommandFunc
  defer func() { shellCommandFunc = origShellCommandFunc }()

  shellCommandCalled := false
  shellCommandFunc = func(name string, args ...string) commandExecutor {
      shellCommandCalled = true

      // Careful: relies on implementation details this could
      // make the test fragile.
      assert.Equal(t, "git", name, "command name")
      assert.Len(t, args, 3, "command args")
      assert.Equal(t, "rev-parse", args[0], "1st command arg")
      assert.Equal(t, "--abbrev-ref", args[1], "2st command arg")
      assert.Equal(t, "HEAD", args[2], "3st command arg")

      // Careful: if the stub deviates from how the system under
      // test works this could generate false positives.
      return &mockCommandExecutor{output: "feature/ISSUE-123_new-issue"}
  }

  branch, err := branchName()
  if assert.NoError(t, err) {
    assert.Equal(t, "feature/ISSUE-123_new-issue", branch, "branch name")
  }

  assert.True(t, shellCommandCalled, "shell command called")
}

func Test_BranchIssue(t *testing.T) {
  origShellCommandFunc := shellCommandFunc
  defer func() { shellCommandFunc = origShellCommandFunc }()

  shellCommandCalled := false
  shellCommandFunc = func(name string, args ...string) commandExecutor {
      shellCommandCalled = true

      // Careful: relies on implementation details this could
      // make the test fragile.
      assert.Equal(t, "git", name, "command name")
      assert.Len(t, args, 3, "command args")
      assert.Equal(t, "rev-parse", args[0], "1st command arg")
      assert.Equal(t, "--abbrev-ref", args[1], "2st command arg")
      assert.Equal(t, "HEAD", args[2], "3st command arg")

      // Careful: if the stub deviates from how the system under
      // test works this could generate false positives.
      return &mockCommandExecutor{output: "feature/ISSUE-123_new-issue"}
  }

  issue, err := BranchIssue()
  if assert.NoError(t, err) {
    assert.Equal(t, "ISSUE-123", issue, "issue")
  }

  assert.True(t, shellCommandCalled, "shell command called")
}
