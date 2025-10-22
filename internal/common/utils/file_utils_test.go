package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetWorkingDirectory(t *testing.T) {
	t.Run("returns current working directory", func(t *testing.T) {
		dir, err := GetWorkingDirectory()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if dir == "" {
			t.Error("expected non-empty directory path")
		}

		// Verify it's an absolute path
		if !filepath.IsAbs(dir) {
			t.Errorf("expected absolute path, got %s", dir)
		}
	})
}

func TestWalkFiles(t *testing.T) {
	t.Run("walks files in directory", func(t *testing.T) {
		tmpDir := t.TempDir()

		// Create test files
		os.WriteFile(filepath.Join(tmpDir, "file1.txt"), []byte("content"), 0644)
		os.WriteFile(filepath.Join(tmpDir, "file2.txt"), []byte("content"), 0644)

		subDir := filepath.Join(tmpDir, "subdir")
		os.Mkdir(subDir, 0755)
		os.WriteFile(filepath.Join(subDir, "file3.txt"), []byte("content"), 0644)

		files, err := WalkFiles(tmpDir)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if len(files) != 3 {
			t.Errorf("expected 3 files, got %d", len(files))
		}
	})

	t.Run("returns empty slice for empty directory", func(t *testing.T) {
		tmpDir := t.TempDir()

		files, err := WalkFiles(tmpDir)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if len(files) != 0 {
			t.Errorf("expected 0 files, got %d", len(files))
		}
	})
}

func TestGetFileStats(t *testing.T) {
	t.Run("returns file stats for existing file", func(t *testing.T) {
		tmpDir := t.TempDir()
		tmpFile := filepath.Join(tmpDir, "test.txt")
		content := []byte("test content for stats")
		err := os.WriteFile(tmpFile, content, 0644)
		if err != nil {
			t.Fatalf("failed to create test file: %v", err)
		}

		stats, err := GetFileStats(tmpFile)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if stats.Size != int64(len(content)) {
			t.Errorf("expected size %d, got %d", len(content), stats.Size)
		}

		if stats.LastModified == 0 {
			t.Error("expected non-zero last modified time")
		}
	})

	t.Run("returns error for non-existent file", func(t *testing.T) {
		_, err := GetFileStats("/non/existent/file.txt")
		if err == nil {
			t.Error("expected error for non-existent file")
		}
	})
}
