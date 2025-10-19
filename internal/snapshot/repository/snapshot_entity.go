package repository

import (
	"fmt"
	"strconv"

	"github.com/alkowskey/commitlens/internal/snapshot/domain"
	"github.com/google/uuid"
)

type SnapshotEntity struct {
	ID    uuid.UUID
	Path  string
	Hash  string
	Size  int64
	Mtime int64
}

func entityFromDomain(s domain.Snapshot) SnapshotEntity {
	return SnapshotEntity{
		ID:    s.ID,
		Path:  s.Path,
		Hash:  fmt.Sprintf("%016x", s.Hash),
		Size:  s.Size,
		Mtime: s.Mtime,
	}
}

func (e SnapshotEntity) toDomain() domain.Snapshot {
	hash, err := strconv.ParseUint(e.Hash, 16, 64)
	if err != nil {
		panic(err)
	}
	return domain.Snapshot{
		ID:    e.ID,
		Path:  e.Path,
		Hash:  hash,
		Size:  e.Size,
		Mtime: e.Mtime,
	}
}
