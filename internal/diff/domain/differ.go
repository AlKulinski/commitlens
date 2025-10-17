package domain

type Differ interface {
	Compare(sourcePath string, targetPath string) (DiffResult, error)
}
