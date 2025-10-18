package usecases

import (
	"github.com/alkowskey/commit-suggester/internal/snapshot/domain"
	"github.com/alkowskey/commit-suggester/internal/snapshot/repository"
)

type CompareUsecase struct {
	snapshotRepository repository.SnapshotRepository
}

func NewCompareUsecase(snapshotRepository repository.SnapshotRepository) *CompareUsecase {
	return &CompareUsecase{
		snapshotRepository: snapshotRepository,
	}
}

func (u *CompareUsecase) Execute() ([]domain.Snapshot, error) {
	return u.snapshotRepository.GetSnapshots()
}
