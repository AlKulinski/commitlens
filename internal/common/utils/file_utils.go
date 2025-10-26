package utils

import (
	"bufio"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type FileStats struct {
	Size         int64
	LastModified int64
}

func GetWorkingDirectory() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return pwd, nil
}

func CopyDirectory(destination string, source string) error {
	println("Copying directory", source, "to", destination)
	err := os.CopyFS(destination, os.DirFS(source))
	return err
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

func OpenOrEmpty(path string) (io.ReadCloser, error) {
	f, err := os.Open(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return io.NopCloser(strings.NewReader("")), nil
		}
		return nil, err
	}
	return f, nil
}

func ReadLines(reader io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
