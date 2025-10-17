package infra

import (
	"bufio"
	"os"

	"github.com/alkowskey/commit-suggester/internal/diff/domain"
)

type BaseDiffer struct{}

func NewBaseDiffer() *BaseDiffer {
	return &BaseDiffer{}
}

func (d *BaseDiffer) Compare(sourcePath string, targetPath string) (domain.DiffResult, error) {
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return domain.DiffResult{}, err
	}
	defer sourceFile.Close()

	targetFile, err := os.Open(targetPath)
	if err != nil {
		return domain.DiffResult{}, err
	}
	defer targetFile.Close()

	if sourceFile.Name() == targetFile.Name() {
		return domain.DiffResult{}, nil
	}

	diffResult, err := d.diff(sourceFile, targetFile)

	if err != nil {
		return domain.DiffResult{}, err
	}

	return diffResult, nil
}

func (d *BaseDiffer) diff(sourceFile *os.File, targetFile *os.File) (domain.DiffResult, error) {
	scannerSource := bufio.NewScanner(sourceFile)
	scannerTarget := bufio.NewScanner(targetFile)

	result := domain.DiffResult{}

	for {
		hasSource := scannerSource.Scan()
		hasTarget := scannerTarget.Scan()

		// Both files ended at the same time
		if !hasSource && !hasTarget {
			break
		}

		var sourceLine, targetLine string

		// Get the current lines (empty string if file ended)
		if hasSource {
			sourceLine = scannerSource.Text()
		}
		if hasTarget {
			targetLine = scannerTarget.Text()
		}

		// Handle different cases
		if !hasSource && hasTarget {
			// Source file ended, but target has more lines (additions)
			result.Added = append(result.Added, targetLine)
			result.HasDifferences = true
		} else if hasSource && !hasTarget {
			// Target file ended, but source has more lines (removals)
			result.Removed = append(result.Removed, sourceLine)
			result.HasDifferences = true
		} else if sourceLine != targetLine {
			// Both files have lines, but they're different
			result.Removed = append(result.Removed, sourceLine)
			result.Added = append(result.Added, targetLine)
			result.HasDifferences = true
		}
	}

	// Check for scanner errors
	if err := scannerSource.Err(); err != nil {
		return domain.DiffResult{}, err
	}
	if err := scannerTarget.Err(); err != nil {
		return domain.DiffResult{}, err
	}

	return result, nil
}
