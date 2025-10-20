package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/alkowskey/commitlens/internal/diff/domain"
)

func LogDiff(diffs []domain.DiffResult) error {
	content := FormatDiffsPlain(diffs)
	if err := os.WriteFile("diff.txt", []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write diff to file: %w", err)
	}
	return nil
}

func PrintDiff(diffs []domain.DiffResult) {
	fmt.Print(FormatDiffs(diffs))
}

func FormatDiffs(diffs []domain.DiffResult) string {
	if len(diffs) == 0 {
		return "No files to compare\n"
	}

	var result strings.Builder
	changedFiles := 0

	for _, diff := range diffs {
		if diff.HasDifferences {
			result.WriteString(diff.String())
			result.WriteString("\n")
			changedFiles++
		}
	}

	if changedFiles == 0 {
		return "No differences found in any files\n"
	}

	result.WriteString(fmt.Sprintf("\nSummary: %d file(s) changed\n", changedFiles))
	return result.String()
}

func FormatDiffsPlain(diffs []domain.DiffResult) string {
	if len(diffs) == 0 {
		return "No files to compare\n"
	}

	var result strings.Builder
	changedFiles := 0

	for _, diff := range diffs {
		if diff.HasDifferences {
			result.WriteString(diff.StringPlain())
			result.WriteString("\n")
			changedFiles++
		}
	}

	if changedFiles == 0 {
		return "No differences found in any files\n"
	}

	result.WriteString(fmt.Sprintf("\nSummary: %d file(s) changed\n", changedFiles))
	return result.String()
}
