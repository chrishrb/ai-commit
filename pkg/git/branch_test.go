package git

import (
	"testing"
)

func Test_issuerNumberWithBranchName(t *testing.T) {
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
		result, _ := issuerNumberWithBranchName(test.branchName)
		if result != test.expected {
			t.Errorf("For branch name '%s', expected '%s' but got '%s'", test.branchName, test.expected, result)
		}
	}
}
