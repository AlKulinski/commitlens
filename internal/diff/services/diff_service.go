package services

import (
	"github.com/alkowskey/commitlens/internal/diff/domain"
)

type DiffService interface {
	CompareFiles(file1, file2 string) (domain.DiffResult, error)
}

func NewDiffService(differ domain.Differ) DiffService {
	return &diffService{
		differ: differ,
	}
}

type diffService struct {
	differ domain.Differ
}

func (ds *diffService) CompareFiles(file1, file2 string) (domain.DiffResult, error) {
	return ds.differ.Compare(file1, file2)
}
