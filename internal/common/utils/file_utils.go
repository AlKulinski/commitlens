package utils

import (
	"os"
	"path/filepath"
)

type FileStats struct {
	Size         int64
	LastModified int64
}

func WalkFiles(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func GetFileStats(path string) (FileStats, error) {
	// Get file info
	fileInfo, err := os.Stat(path)
	if err != nil {
		return FileStats{}, err
	}

	// Get file size
	size := fileInfo.Size()

	// Get last modified time
	lastModified := fileInfo.ModTime().Unix()

	return FileStats{
		Size:         size,
		LastModified: lastModified,
	}, nil
}
