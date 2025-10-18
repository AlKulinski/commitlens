package repository

import "github.com/alkowskey/commit-suggester/internal/snapshot/domain"

type SnapshotEntity struct {
	ID    int
	Path  string
	Hash  string
	Size  int
	Mtime int
}

func entityFromDomain(s domain.Snapshot) SnapshotEntity {
	return SnapshotEntity{
		ID:    s.ID,
		Path:  s.Path,
		Hash:  s.Hash,
		Size:  s.Size,
		Mtime: s.Mtime,
	}
}

func (e SnapshotEntity) toDomain() domain.Snapshot {
	return domain.Snapshot{
		ID:    e.ID,
		Path:  e.Path,
		Hash:  e.Hash,
		Size:  e.Size,
		Mtime: e.Mtime,
	}
}
