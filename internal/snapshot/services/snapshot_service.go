package services

import (
	"fmt"

	"github.com/alkowskey/commit-suggester/internal/common/utils"
	"github.com/alkowskey/commit-suggester/internal/snapshot/domain"
	"github.com/alkowskey/commit-suggester/internal/snapshot/repository"
	"github.com/google/uuid"
)

type SnapshotService interface {
	TakeSnapshot(string) ([]domain.Snapshot, error)
	Compare(string) ([]domain.Snapshot, error)
	GetSnapshotDirectory(string) string
}

type SnapshotServiceImpl struct {
	snapshotRepository repository.SnapshotRepository
}

func NewSnapshotService(snapshotRepository repository.SnapshotRepository) SnapshotService {
	return &SnapshotServiceImpl{
		snapshotRepository: snapshotRepository,
	}
}

func (s *SnapshotServiceImpl) GetSnapshotDirectory(subdirectory string) string {
	snapshotDir, err := utils.GetWorkingDirectory()

	if err != nil {
		panic(err)
	}

	return fmt.Sprintf(".cache/%s/%s", snapshotDir, subdirectory)
}

func (s *SnapshotServiceImpl) Compare(subdirectory string) ([]domain.Snapshot, error) {
	existingSnapshots, err := s.snapshotRepository.GetSnapshots()
	if err != nil {
		return nil, err
	}

	files, err := utils.WalkFiles(subdirectory)
	if err != nil {
		return nil, err
	}

	fileSnapshots, err := s.buildSnapshotsFromFiles(files)
	if err != nil {
		return nil, err
	}

	snapshotDiffs := compareSnapshots(existingSnapshots, fileSnapshots)
	return snapshotDiffs, nil
}

func (s *SnapshotServiceImpl) TakeSnapshot(subdirectory string) ([]domain.Snapshot, error) {
	err := s.verifyExistingSnapshots()
	if err != nil {
		return nil, err
	}

	files, err := utils.WalkFiles(subdirectory)
	if err != nil {
		return nil, err
	}

	fileSnapshots, err := s.buildSnapshotsFromFiles(files)
	if err != nil {
		return nil, err
	}

	s.snapshotRepository.StoreBatchSnapshots(fileSnapshots)
	return fileSnapshots, nil
}

func (s *SnapshotServiceImpl) buildSnapshotsFromFiles(files []string) ([]domain.Snapshot, error) {
	fileSnapshots := make([]domain.Snapshot, len(files))

	for i, file := range files {
		stats, err := utils.GetFileStats(file)
		if err != nil {
			return nil, err
		}

		hash, err := utils.CalculateContentHash(file)
		if err != nil {
			return nil, err
		}

		fileSnapshots[i] = domain.Snapshot{
			ID:    uuid.New(),
			Path:  file,
			Hash:  hash,
			Size:  stats.Size,
			Mtime: stats.LastModified,
		}
	}

	return fileSnapshots, nil
}

func (s *SnapshotServiceImpl) verifyExistingSnapshots() error {
	existingSnapshots, err := s.snapshotRepository.GetSnapshots()
	if len(existingSnapshots) > 0 || err != nil {
		return domain.ErrSnapshotsAlreadyExist
	}
	return nil
}

func compareSnapshots(existingSnapshots, newSnapshots []domain.Snapshot) []domain.Snapshot {
	var snapshotDiffs []domain.Snapshot

	existingSnapshotMap := createSnapshotMap(existingSnapshots)
	newSnapshotMap := createSnapshotMap(newSnapshots)

	for _, newSnapshot := range newSnapshots {
		snap := existingSnapshotMap[newSnapshot.Hash]
		if snap == nil {
			snapshotDiffs = append(snapshotDiffs, newSnapshot)
		}
	}

	for _, existingSnapshot := range existingSnapshots {
		snap := newSnapshotMap[existingSnapshot.Hash]
		if snap == nil {
			snapshotDiffs = append(snapshotDiffs, existingSnapshot)
		}
	}

	return utils.DedupBy(snapshotDiffs, func(s domain.Snapshot) string { return s.Path })
}

func createSnapshotMap(snapshot []domain.Snapshot) map[uint64]*domain.Snapshot {
	snapshotMap := make(map[uint64]*domain.Snapshot)
	for i := range snapshot {
		s := &snapshot[i]
		snapshotMap[s.Hash] = s
	}
	return snapshotMap
}
