package git

import (
	"os/exec"
	"regexp"
	"strings"
)

func BranchIssuerNumber() (string, error) {
	branchName, err := branchName()
	if err != nil {
		return "", err
	}
	return issuerNumberWithBranchName(branchName)
}

func branchName() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func issuerNumberWithBranchName(branchName string) (string, error) {
	const branchIssuerNumberRegex = `([.]*\/)([\-\w]*?\-\d+)`
	branchName = strings.TrimSpace(branchName)
	re := regexp.MustCompile(branchIssuerNumberRegex)
	issuerNumber := re.FindStringSubmatch(branchName)
	if len(issuerNumber) < 2 {
		return "", nil
	}
	return issuerNumber[2], nil
}
