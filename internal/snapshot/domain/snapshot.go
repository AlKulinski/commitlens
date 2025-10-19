package domain

import (
	"fmt"

	"github.com/alkowskey/commitlens/internal/snapshot/domain/errors"
	"github.com/google/uuid"
)

type Snapshot struct {
	ID        uuid.UUID
	Path      string
	CreatedAt string
	UpdatedAt string
	Hash      uint64
	Size      int64
	Mtime     int64
}

func (s Snapshot) String() string {
	return fmt.Sprintf("Snapshot{Path: %s, Hash: %d, Size: %d}", s.Path, s.Hash, s.Size)
}

var ErrSnapshotsAlreadyExist = errors.ErrSnapshotsAlreadyExist
