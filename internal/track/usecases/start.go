package usecases

import (
	"github.com/alkowskey/commit-suggester/internal/snapshot/domain"
	"github.com/alkowskey/commit-suggester/internal/snapshot/repository"
)

type StartTrackedDiffUsecase struct {
	snapshotRepository repository.SnapshotRepository
}

func NewTrackStartUsecase(snapshotRepository repository.SnapshotRepository) *StartTrackedDiffUsecase {
	return &StartTrackedDiffUsecase{
		snapshotRepository: snapshotRepository,
	}
}

func (d *StartTrackedDiffUsecase) Execute(sourcePath string, targetPath string) error {
	snapshot := &domain.Snapshot{
		Path:  sourcePath,
		Hash:  "123",
		Size:  32,
		Mtime: 123,
	}

	return d.snapshotRepository.StoreSnapshot(*snapshot)
}
