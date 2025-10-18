package usecases

import "github.com/alkowskey/commit-suggester/internal/snapshot/repository"

type FlushSnapshotsUsecase struct {
	snapshotRepository repository.SnapshotRepository
}

func NewFlushSnapshotsUsecase(snapshotRepository repository.SnapshotRepository) *FlushSnapshotsUsecase {
	return &FlushSnapshotsUsecase{
		snapshotRepository: snapshotRepository,
	}
}

func (u *FlushSnapshotsUsecase) Execute() error {
	return u.snapshotRepository.FlushSnapshots()
}
