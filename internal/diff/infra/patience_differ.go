package infra

import (
	"github.com/alkowskey/commitlens/internal/common/utils"
	"github.com/alkowskey/commitlens/internal/diff/domain"
)

type PatienceDiffer struct {
	result domain.DiffResult
}

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

	d.result = domain.DiffResult{
		Path: targetPath,
	}

	d.patienceDiff(sourceLines, targetLines)
	return d.result, nil
}

func (d *PatienceDiffer) patienceDiff(source, target []string) {
	matches := d.findUniqueMatches(source, target)
	d.buildDiffResult(source, target, matches)
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

func (d *PatienceDiffer) buildDiffResult(source, target []string, matches []match) {
	sourceIdx := 0
	targetIdx := 0

	for _, m := range matches {
		d.processRegion(source, target, sourceIdx, m.sourceIdx, targetIdx, m.targetIdx)

		sourceIdx = m.sourceIdx + 1
		targetIdx = m.targetIdx + 1
	}

	d.processRegion(source, target, sourceIdx, len(source), targetIdx, len(target))
}

func (d *PatienceDiffer) processRegion(source, target []string,
	sourceStart, sourceEnd, targetStart, targetEnd int) {

	sourceLen := sourceEnd - sourceStart
	targetLen := targetEnd - targetStart

	maxLen := max(sourceLen, targetLen)

	for i := 0; i < maxLen; i++ {
		hasSource := sourceStart+i < sourceEnd
		hasTarget := targetStart+i < targetEnd

		if hasSource && hasTarget {
			sourceLine := source[sourceStart+i]
			targetLine := target[targetStart+i]

			if sourceLine != targetLine {
				d.result.Removed = append(d.result.Removed, sourceLine)
				d.result.Added = append(d.result.Added, targetLine)
				d.result.HasDifferences = true
			}
		} else if hasSource {
			d.result.Removed = append(d.result.Removed, source[sourceStart+i])
			d.result.HasDifferences = true
		} else if hasTarget {
			d.result.Added = append(d.result.Added, target[targetStart+i])
			d.result.HasDifferences = true
		}
	}
}
