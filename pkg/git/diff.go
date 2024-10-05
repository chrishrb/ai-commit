package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

var ignoredFiles []string = []string{
	".lock",
	".mod",
	".sum",
	"-lock",
	".svg",
	".png",
	".jpg",
	".jpeg",
	".webp",
	".gif",
}

func getStagedFiles() ([]string, error) {
	cmd := exec.Command("git", "diff", "--relative", "--name-only", "--cached")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	files := strings.Split(strings.TrimSpace(out.String()), "\n")
	return files, nil
}

func GetDiff(additionalIgnoredFiles []string) (string, error) {
	// Get the list of staged files
	files, err := getStagedFiles()
	if err != nil {
		return "", err
	}

	filteredFiles := filter(files, additionalIgnoredFiles)

	// Run git diff for the remaining files
	cmd := exec.Command("git", append([]string{"--no-pager", "diff", "--cached", "--relative"}, filteredFiles...)...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

func filter(files, additionalIgnoredFiles []string) []string {
	var filteredFiles []string
	var ignored []string
	ignoredFiles = append(ignoredFiles, additionalIgnoredFiles...)
	for _, file := range files {
		ignore := false
		for _, pattern := range ignoredFiles {
			if strings.Contains(file, pattern) {
				ignore = true
				ignored = append(ignored, file)
				break
			}
		}
		if !ignore {
			filteredFiles = append(filteredFiles, file)
		}
	}

	if len(ignored) > 0 {
		fmt.Println("Some files are excluded from 'git diff'. No commit messages are generated for this files:")
		for _, file := range ignored {
			fmt.Println(file)
		}
	}

	return filteredFiles
}
