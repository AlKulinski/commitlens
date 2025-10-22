package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCalculateContentHash(t *testing.T) {
	t.Run("calculates hash for existing file", func(t *testing.T) {
		// Create a temporary file
		tmpDir := t.TempDir()
		tmpFile := filepath.Join(tmpDir, "test.txt")
		content := []byte("test content")
		err := os.WriteFile(tmpFile, content, 0644)
		if err != nil {
			t.Fatalf("failed to create test file: %v", err)
		}

		hash, err := CalculateContentHash(tmpFile)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if hash == 0 {
			t.Error("expected non-zero hash")
		}
	})

	t.Run("returns same hash for same content", func(t *testing.T) {
		tmpDir := t.TempDir()
		content := []byte("identical content")

		file1 := filepath.Join(tmpDir, "file1.txt")
		file2 := filepath.Join(tmpDir, "file2.txt")

		os.WriteFile(file1, content, 0644)
		os.WriteFile(file2, content, 0644)

		hash1, _ := CalculateContentHash(file1)
		hash2, _ := CalculateContentHash(file2)

		if hash1 != hash2 {
			t.Errorf("expected same hash for identical content, got %d and %d", hash1, hash2)
		}
	})

	t.Run("returns different hash for different content", func(t *testing.T) {
		tmpDir := t.TempDir()

		file1 := filepath.Join(tmpDir, "file1.txt")
		file2 := filepath.Join(tmpDir, "file2.txt")

		os.WriteFile(file1, []byte("content A"), 0644)
		os.WriteFile(file2, []byte("content B"), 0644)

		hash1, _ := CalculateContentHash(file1)
		hash2, _ := CalculateContentHash(file2)

		if hash1 == hash2 {
			t.Error("expected different hashes for different content")
		}
	})

	t.Run("returns error for non-existent file", func(t *testing.T) {
		_, err := CalculateContentHash("/non/existent/file.txt")
		if err == nil {
			t.Error("expected error for non-existent file")
		}
	})
}
