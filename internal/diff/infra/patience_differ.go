package infra

import (
	"github.com/alkowskey/commitlens/internal/common/utils"
	"github.com/alkowskey/commitlens/internal/diff/domain"
)

type PatienceDiffer struct{}

func NewPatienceDiffer() *PatienceDiffer {
	return &PatienceDiffer{}
}

func (d *PatienceDiffer) Compare(targetPath string, sourcePath string) (domain.DiffResult, error) {
	sourceFile, err := utils.OpenOrEmpty(sourcePath)
	if err != nil {
		panic("error opening source file")
	}
	defer sourceFile.Close()

	targetFile, err := utils.OpenOrEmpty(targetPath)
	if err != nil {
		panic("error opening target file")
	}
	defer targetFile.Close()

	sourceLines, err := utils.ReadLines(sourceFile)
	if err != nil {
		return domain.DiffResult{}, err
	}

	targetLines, err := utils.ReadLines(targetFile)
	if err != nil {
		return domain.DiffResult{}, err
	}

	return d.patienceDiff(sourceLines, targetLines), nil
}

func (d *PatienceDiffer) patienceDiff(source, target []string) domain.DiffResult {
	matches := d.findUniqueMatches(source, target)
	return d.buildDiffResult(source, target, matches)
}

func (d *PatienceDiffer) findUniqueMatches(source, target []string) []match {
	sourceUnique := findUniqueLines(source)
	targetUnique := findUniqueLines(target)

	var matches []match
	for line, sourceIdx := range sourceUnique {
		if targetIdx, exists := targetUnique[line]; exists {
			matches = append(matches, match{sourceIdx: sourceIdx, targetIdx: targetIdx})
		}
	}

	return longestIncreasingSubsequence(matches)
}

type match struct {
	sourceIdx int
	targetIdx int
}

func findUniqueLines(lines []string) map[string]int {
	counts := make(map[string]int)
	lastIdx := make(map[string]int)

	for idx, line := range lines {
		counts[line]++
		lastIdx[line] = idx
	}

	unique := make(map[string]int)
	for line, count := range counts {
		if count == 1 {
			unique[line] = lastIdx[line]
		}
	}

	return unique
}

func longestIncreasingSubsequence(matches []match) []match {
	if len(matches) == 0 {
		return nil
	}

	n := len(matches)
	dp := make([]int, n)
	parent := make([]int, n)

	for i := range dp {
		dp[i] = 1
		parent[i] = -1
	}

	maxLen := 1
	maxIdx := 0

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if matches[j].sourceIdx < matches[i].sourceIdx &&
				matches[j].targetIdx < matches[i].targetIdx &&
				dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				parent[i] = j
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
			maxIdx = i
		}
	}

	result := make([]match, maxLen)
	idx := maxIdx
	for i := maxLen - 1; i >= 0; i-- {
		result[i] = matches[idx]
		idx = parent[idx]
	}

	return result
}

func (d *PatienceDiffer) buildDiffResult(source, target []string, matches []match) domain.DiffResult {
	result := domain.DiffResult{}

	sourceIdx := 0
	targetIdx := 0

	for _, m := range matches {
		d.processRegion(&result, source, target, sourceIdx, m.sourceIdx, targetIdx, m.targetIdx)

		sourceIdx = m.sourceIdx + 1
		targetIdx = m.targetIdx + 1
	}

	d.processRegion(&result, source, target, sourceIdx, len(source), targetIdx, len(target))

	return result
}

func (d *PatienceDiffer) processRegion(result *domain.DiffResult, source, target []string,
	sourceStart, sourceEnd, targetStart, targetEnd int) {

	sourceLen := sourceEnd - sourceStart
	targetLen := targetEnd - targetStart

	maxLen := sourceLen
	if targetLen > maxLen {
		maxLen = targetLen
	}

	for i := 0; i < maxLen; i++ {
		hasSource := sourceStart+i < sourceEnd
		hasTarget := targetStart+i < targetEnd

		if hasSource && hasTarget {
			sourceLine := source[sourceStart+i]
			targetLine := target[targetStart+i]

			if sourceLine != targetLine {
				result.Removed = append(result.Removed, sourceLine)
				result.Added = append(result.Added, targetLine)
				result.HasDifferences = true
			}
		} else if hasSource {
			result.Removed = append(result.Removed, source[sourceStart+i])
			result.HasDifferences = true
		} else if hasTarget {
			result.Added = append(result.Added, target[targetStart+i])
			result.HasDifferences = true
		}
	}
}
