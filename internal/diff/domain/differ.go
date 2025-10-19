package domain

type DiffRequest struct {
	Options DiffOptions
}

type DiffOptions struct {
	Ignore []string
}

type DiffResult struct {
	HasDifferences bool
	Added          []string
	Removed        []string
}

type Differ interface {
	Compare(sourcePath string, targetPath string) (DiffResult, error)
}
