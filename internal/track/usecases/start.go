package usecases

import (
	"github.com/alkowskey/commit-suggester/internal/common/utils"
	"github.com/alkowskey/commit-suggester/internal/snapshot/domain"
	"github.com/alkowskey/commit-suggester/internal/snapshot/repository"
	"github.com/google/uuid"
)

type StartTrackedDiffUsecase struct {
	snapshotRepository repository.SnapshotRepository
}

func NewTrackStartUsecase(snapshotRepository repository.SnapshotRepository) *StartTrackedDiffUsecase {
	return &StartTrackedDiffUsecase{
		snapshotRepository: snapshotRepository,
	}
}

func (d *StartTrackedDiffUsecase) Execute(subdirectory string) error {
	err := d.verifyExistingSnapshots()
	if err != nil {
		return err
	}

	files, err := utils.WalkFiles(subdirectory)
	if err != nil {
		return err
	}

	fileSnapshots := make([]domain.Snapshot, len(files))

	for i, file := range files {
		stats, err := utils.GetFileStats(file)
		if err != nil {
			return err
		}

		hash, err := utils.CalculateContentHash(file)
		if err != nil {
			return err
		}

		fileSnapshots[i] = domain.Snapshot{
			ID:    uuid.New(),
			Path:  file,
			Hash:  hash,
			Size:  stats.Size,
			Mtime: stats.LastModified,
		}
	}

	return d.snapshotRepository.StoreBatchSnapshots(fileSnapshots)
}

func (d *StartTrackedDiffUsecase) verifyExistingSnapshots() error {
	existingSnapshots, err := d.snapshotRepository.GetSnapshots()
	if len(existingSnapshots) > 0 || err != nil {
		return domain.ErrSnapshotsAlreadyExist
	}
	return nil
}
