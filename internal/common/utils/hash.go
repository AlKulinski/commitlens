package utils

import (
	"os"

	"github.com/cespare/xxhash/v2"
)

func CalculateContentHash(path string) (uint64, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}

	hash := xxhash.Sum64(data)
	return hash, nil
}
