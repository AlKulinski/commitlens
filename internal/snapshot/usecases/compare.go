package usecases

import (
	"github.com/alkowskey/commitlens/internal/diff/domain"
	"github.com/alkowskey/commitlens/internal/diff/utils"
	"github.com/alkowskey/commitlens/internal/snapshot/services"
)

type CompareUsecase struct {
	snapshotService services.SnapshotService
}

func NewCompareUsecase(snapshotService services.SnapshotService) *CompareUsecase {
	return &CompareUsecase{
		snapshotService: snapshotService,
	}
}

func (u *CompareUsecase) Execute(subdirectory string) ([]domain.DiffResult, error) {
	compared, err := u.snapshotService.Compare(subdirectory)
	if err != nil {
		return nil, err
	}
	utils.LogDiff(compared)
	utils.PrintDiff(compared)
	return compared, nil
}
