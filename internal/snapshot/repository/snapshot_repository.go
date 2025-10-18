package repository

import (
	"database/sql"

	"github.com/alkowskey/commit-suggester/internal/snapshot/domain"
)

type SnapshotRepository interface {
	StoreSnapshot(snapshot domain.Snapshot) error
	StoreBatchSnapshots(snapshot []domain.Snapshot) error
}

type SnapshotRepositoryImpl struct {
	db *sql.DB
}

func NewSnapshotRepository(db *sql.DB) SnapshotRepository {
	return &SnapshotRepositoryImpl{
		db: db,
	}
}

func (s *SnapshotRepositoryImpl) StoreSnapshot(snapshot domain.Snapshot) error {
	snapshotEntity := entityFromDomain(snapshot)
	_, err := s.db.Exec("INSERT INTO snapshots (id, path, hash, size, mtime) VALUES (?, ?, ?, ?, ?)", snapshotEntity.ID, snapshotEntity.Path, snapshotEntity.Hash, snapshotEntity.Size, snapshotEntity.Mtime)
	if err != nil {
		return err
	}
	return nil
}

func (s *SnapshotRepositoryImpl) StoreBatchSnapshots(snapshot []domain.Snapshot) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("INSERT INTO snapshots (id, path, hash, size, mtime) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, snapshot := range snapshot {
		snapshotEntity := entityFromDomain(snapshot)
		_, err = stmt.Exec(snapshotEntity.ID, snapshotEntity.Path, snapshotEntity.Hash, snapshotEntity.Size, snapshotEntity.Mtime)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}
