package repository

import (
	"database/sql"

	"github.com/alkowskey/commitlens/internal/snapshot/domain"
)

type SnapshotRepository interface {
	StoreSnapshot(snapshot domain.Snapshot) error
	StoreBatchSnapshots(snapshot []domain.Snapshot) error
	FlushSnapshots() error
	GetSnapshots() ([]domain.Snapshot, error)
}

type SnapshotRepositoryImpl struct {
	db *sql.DB
}

func NewSnapshotRepository(db *sql.DB) SnapshotRepository {
	return &SnapshotRepositoryImpl{
		db: db,
	}
}

func (s *SnapshotRepositoryImpl) GetSnapshots() ([]domain.Snapshot, error) {
	rows, err := s.db.Query("SELECT id, path, hash, size, mtime FROM snapshots")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var snapshots []domain.Snapshot
	for rows.Next() {
		var snapshot SnapshotEntity
		err := rows.Scan(&snapshot.ID, &snapshot.Path, &snapshot.Hash, &snapshot.Size, &snapshot.Mtime)
		if err != nil {
			return nil, err
		}
		snapshots = append(snapshots, snapshot.toDomain())
	}
	return snapshots, nil
}

func (s *SnapshotRepositoryImpl) FlushSnapshots() error {
	_, err := s.db.Exec("DELETE FROM snapshots")
	if err != nil {
		return err
	}
	return nil
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
