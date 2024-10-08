package git

import (
	"regexp"
	"strings"
)

func BranchIssue() (string, error) {
	branchName, err := branchName()
	if err != nil {
		return "", err
	}
	return issueWithBranchName(branchName)
}

func branchName() (string, error) {
	cmd := shellCommandFunc("git", "rev-parse", "--abbrev-ref", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func issueWithBranchName(branchName string) (string, error) {
	const branchIssuerNumberRegex = `([.]*\/)([\-\w]*?\-\d+)`
	branchName = strings.TrimSpace(branchName)
	re := regexp.MustCompile(branchIssuerNumberRegex)
	issuerNumber := re.FindStringSubmatch(branchName)
	if len(issuerNumber) < 2 {
		return "", nil
	}
	return issuerNumber[2], nil
}
