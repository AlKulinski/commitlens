package usecases

import (
	"github.com/alkowskey/commit-suggester/internal/common/utils"
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
	snapshotDir := d.SnapshotService.GetSnapshotDirectory(subdirectory)
	err := utils.CopyDirectory(snapshotDir, subdirectory)
	if err != nil {
		return err
	}
	_, err = d.SnapshotService.TakeSnapshot(subdirectory)

	return err
}
