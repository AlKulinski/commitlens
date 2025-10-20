package infra

import (
	"bufio"
	"os"

	"github.com/alkowskey/commitlens/internal/diff/domain"
)

type BaseDiffer struct{}

func NewBaseDiffer() *BaseDiffer {
	return &BaseDiffer{}
}

func (d *BaseDiffer) Compare(targetPath string, sourcePath string) (domain.DiffResult, error) {
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

		if !hasSource && !hasTarget {
			break
		}

		var sourceLine, targetLine string

		if hasSource {
			sourceLine = scannerSource.Text()
		}
		if hasTarget {
			targetLine = scannerTarget.Text()
		}

		if !hasSource && hasTarget {
			result.Added = append(result.Added, targetLine)
			result.HasDifferences = true
		} else if hasSource && !hasTarget {
			result.Removed = append(result.Removed, sourceLine)
			result.HasDifferences = true
		} else if sourceLine != targetLine {
			result.Removed = append(result.Removed, sourceLine)
			result.Added = append(result.Added, targetLine)
			result.HasDifferences = true
		}
	}

	if err := scannerSource.Err(); err != nil {
		return domain.DiffResult{}, err
	}
	if err := scannerTarget.Err(); err != nil {
		return domain.DiffResult{}, err
	}

	return result, nil
}
