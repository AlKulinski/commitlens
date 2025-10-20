package usecases

import "github.com/alkowskey/commitlens/internal/snapshot/services"

type FlushSnapshotsUsecase struct {
	snapshotService services.SnapshotService
}

func NewFlushSnapshotsUsecase(snapshotService services.SnapshotService) *FlushSnapshotsUsecase {
	return &FlushSnapshotsUsecase{
		snapshotService: snapshotService,
	}
}

func (u *FlushSnapshotsUsecase) Execute() error {
	return u.snapshotService.FlushSnapshots()
}
