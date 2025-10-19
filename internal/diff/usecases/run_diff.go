package usecases

import (
	"github.com/alkowskey/commitlens/internal/diff/domain"
)

type DiffUsecase struct {
	SourcePath string
	TargetPath string
	differ     domain.Differ
}

func NewDiffUsecase(sourcePath string, targetPath string, differ domain.Differ) *DiffUsecase {
	return &DiffUsecase{
		SourcePath: sourcePath,
		TargetPath: targetPath,
		differ:     differ,
	}
}

func (d *DiffUsecase) Execute(request domain.DiffRequest) (domain.DiffResult, error) {
	return d.differ.Compare(d.SourcePath, d.TargetPath)
}
