package usecases

import (
	"github.com/alkowskey/commit-suggester/internal/snapshot/services"
)

type StartTrackedDiffUsecase struct {
	SnapshotService services.SnapshotService
}

func NewTrackStartUsecase(snapshotService services.SnapshotService) *StartTrackedDiffUsecase {
	return &StartTrackedDiffUsecase{
		SnapshotService: snapshotService,
	}
}

func (d *StartTrackedDiffUsecase) Execute(subdirectory string) error {
	_, err := d.SnapshotService.TakeSnapshot(subdirectory)

	return err
}
