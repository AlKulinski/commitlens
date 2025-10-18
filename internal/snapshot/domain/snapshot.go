package domain

type Snapshot struct {
	ID        int
	Path      string
	CreatedAt string
	UpdatedAt string
	Hash      string
	Size      int
	Mtime     int
}
