package usecases

import (
	"github.com/alkowskey/commit-suggester/internal/snapshot/domain"
	"github.com/alkowskey/commit-suggester/internal/snapshot/services"
)

type CompareUsecase struct {
	snapshotService services.SnapshotService
}

func NewCompareUsecase(snapshotService services.SnapshotService) *CompareUsecase {
	return &CompareUsecase{
		snapshotService: snapshotService,
	}
}

func (u *CompareUsecase) Execute(subdirectory string) ([]domain.Snapshot, error) {
	return u.snapshotService.Compare(subdirectory)
}
