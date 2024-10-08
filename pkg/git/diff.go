package git

import (
	"fmt"
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

func GetDiff(additionalIgnoredFiles []string) (string, error) {
	// Get the list of staged files
	files, err := getStagedFiles()
	if err != nil {
		return "", err
	}

	filteredFiles := filter(files, additionalIgnoredFiles)
	if len(filteredFiles) == 0 {
		return "", nil
	}

	// Run git diff for the remaining files
	cmd := shellCommandFunc("git", append([]string{"--no-pager", "diff", "--cached", "--relative", "--"}, filteredFiles...)...)
  out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}

func getStagedFiles() ([]string, error) {
	cmd := shellCommandFunc("git", "diff", "--cached", "--relative", "--name-only")
  out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	trimmedOut := strings.TrimSpace(string(out))
	if trimmedOut == "" {
		return []string{}, nil
	}
	files := strings.Split(trimmedOut, "\n")
	return files, nil
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
    fmt.Println("🙈 Some files are excluded from generating commit messages:")
		for _, file := range ignored {
			fmt.Printf("  ・%s\n", file)
		}
	}

	return filteredFiles
}
