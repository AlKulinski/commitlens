package infra

import (
	"bufio"
	"os"

	"github.com/alkowskey/commitlens/internal/diff/domain"
)

type BaseDiffer struct {
	result domain.DiffResult
}

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

	err = d.diff(sourceFile, targetFile)

	if err != nil {
		panic(err)
	}

	return d.result, nil
}

func (d *BaseDiffer) diff(sourceFile *os.File, targetFile *os.File) error {
	scannerSource := bufio.NewScanner(sourceFile)
	scannerTarget := bufio.NewScanner(targetFile)

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
			d.result.Added = append(d.result.Added, targetLine)
			d.result.HasDifferences = true
		} else if hasSource && !hasTarget {
			d.result.Removed = append(d.result.Removed, sourceLine)
			d.result.HasDifferences = true
		} else if sourceLine != targetLine {
			d.result.Removed = append(d.result.Removed, sourceLine)
			d.result.Added = append(d.result.Added, targetLine)
			d.result.HasDifferences = true
		}
	}

	if err := scannerSource.Err(); err != nil {
		return err
	}
	if err := scannerTarget.Err(); err != nil {
		return err
	}

	return nil
}
