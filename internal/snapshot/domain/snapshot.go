package domain

import (
	"github.com/alkowskey/commit-suggester/internal/snapshot/domain/errors"
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

var ErrSnapshotsAlreadyExist = errors.ErrSnapshotsAlreadyExist
