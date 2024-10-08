package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BuildPrompt(t *testing.T) {
	t.Run("no issue number", func(t *testing.T) {
    expected := `You are provided with a git diff output that shows code changes. Your task is to generate a structured and descriptive commit message based on the following guidelines:
1. The commit message should have a short, one-line summary (50 characters or less), starting with:
  a) one of the following keywords:
  - feat: for new features or enhancements
  - fix: for bug fixes
  - refactor: for code restructuring without changing behavior
  - docs: for documentation changes
  - test: for adding or modifying tests
  - chore: for maintenance tasks (e.g., updating dependencies)
  Example: 'fix: resolve null pointer exception'
3. After the summary, include a detailed description explaining:
- What has changed and why.
- The issue the changes are addressing (if applicable).
- Any important implications for other parts of the codebase.
4. Only output one commit message and no further explanations.
5. Use an imperative tone (e.g., 'Fix', 'Add', 'Update').
6. Ensure that the message is clear and concise, focusing on the intent of the changes rather than just describing the diff.
7. Don‘t use code to explain the changes.
Here is the git diff output:
`

    err := ParseConfig()
    if assert.NoError(t, err) {
      assert.Equal(t, expected, C.BuildPrompt(""), "empty issue")
    }
  })

	t.Run("with issue number", func(t *testing.T) {
    expected := `You are provided with a git diff output that shows code changes. Your task is to generate a structured and descriptive commit message based on the following guidelines:
1. The commit message should have a short, one-line summary (50 characters or less), starting with:
  a) one of the following keywords:
  - feat: for new features or enhancements
  - fix: for bug fixes
  - refactor: for code restructuring without changing behavior
  - docs: for documentation changes
  - test: for adding or modifying tests
  - chore: for maintenance tasks (e.g., updating dependencies)
  b) The ticket number 'ISSUE-123'.
  Example: 'fix: ISSUE-123 Resolve null pointer exception'
3. After the summary, include a detailed description explaining:
- What has changed and why.
- The issue the changes are addressing (if applicable).
- Any important implications for other parts of the codebase.
4. Only output one commit message and no further explanations.
5. Use an imperative tone (e.g., 'Fix', 'Add', 'Update').
6. Ensure that the message is clear and concise, focusing on the intent of the changes rather than just describing the diff.
7. Don‘t use code to explain the changes.
Here is the git diff output:
`

    err := ParseConfig()
    if assert.NoError(t, err) {
      assert.Equal(t, expected, C.BuildPrompt("ISSUE-123"), "empty issue")
    }
  })
}
