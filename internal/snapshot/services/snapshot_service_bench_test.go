package services

import (
	"fmt"
	"sync"
	"testing"

	diffDomain "github.com/alkowskey/commitlens/internal/diff/domain"
	"github.com/alkowskey/commitlens/internal/snapshot/domain"
	"github.com/google/uuid"
)

// Mock DiffService for benchmarking
type mockDiffService struct {
	delay bool
}

func (m *mockDiffService) CompareFiles(source, target string) (diffDomain.DiffResult, error) {
	// Simulate some work
	if m.delay {
		// Simulate I/O delay
		sum := 0
		for i := 0; i < 10000; i++ {
			sum += i
		}
	}
	return diffDomain.DiffResult{
		HasDifferences: true,
		Added:          []string{"line1", "line2"},
		Removed:        []string{"line3"},
	}, nil
}

// Old concurrent version (unlimited goroutines)
func (s *SnapshotServiceImpl) processSnapshotDifferencesConcurrentOld(
	snapshotDiffs []domain.Snapshot,
	cachedPathPrefix string,
) []diffDomain.DiffResult {
	results := make(chan diffDomain.DiffResult, len(snapshotDiffs))
	var wg sync.WaitGroup

	for _, snapshotDiff := range snapshotDiffs {
		wg.Add(1)
		go func(sd domain.Snapshot) {
			defer wg.Done()
			cachedPath := fmt.Sprintf("%s/%s", cachedPathPrefix, sd.Path)
			result, err := s.diffService.CompareFiles(sd.Path, cachedPath)
			if err != nil {
				panic(err)
			}
			results <- result
		}(snapshotDiff)
	}

	wg.Wait()
	close(results)
	var diffResults []diffDomain.DiffResult
	for val := range results {
		diffResults = append(diffResults, val)
	}

	return diffResults
}

// Synchronous version for comparison
func (s *SnapshotServiceImpl) processSnapshotDifferencesSynchronous(
	snapshotDiffs []domain.Snapshot,
	cachedPathPrefix string,
) []diffDomain.DiffResult {
	diffResults := make([]diffDomain.DiffResult, 0, len(snapshotDiffs))

	for _, snapshotDiff := range snapshotDiffs {
		cachedPath := fmt.Sprintf("%s/%s", cachedPathPrefix, snapshotDiff.Path)
		result, err := s.diffService.CompareFiles(snapshotDiff.Path, cachedPath)
		if err != nil {
			panic(err)
		}
		diffResults = append(diffResults, result)
	}

	return diffResults
}

// Helper to create test snapshots
func createTestSnapshots(count int) []domain.Snapshot {
	snapshots := make([]domain.Snapshot, count)
	for i := 0; i < count; i++ {
		snapshots[i] = domain.Snapshot{
			ID:   uuid.New(),
			Path: fmt.Sprintf("file_%d.txt", i),
			Hash: uint64(i),
		}
	}
	return snapshots
}

func BenchmarkProcessSnapshotDifferences_Optimized_10(b *testing.B) {
	service := &SnapshotServiceImpl{
		diffService: &mockDiffService{delay: true},
	}
	snapshots := createTestSnapshots(10)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		service.processSnapshotDifferences(snapshots, "/cache")
	}
}

func BenchmarkProcessSnapshotDifferences_ConcurrentOld_10(b *testing.B) {
	service := &SnapshotServiceImpl{
		diffService: &mockDiffService{delay: true},
	}
	snapshots := createTestSnapshots(10)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		service.processSnapshotDifferencesConcurrentOld(snapshots, "/cache")
	}
}

func BenchmarkProcessSnapshotDifferences_Synchronous_10(b *testing.B) {
	service := &SnapshotServiceImpl{
		diffService: &mockDiffService{delay: true},
	}
	snapshots := createTestSnapshots(10)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		service.processSnapshotDifferencesSynchronous(snapshots, "/cache")
	}
}

func BenchmarkProcessSnapshotDifferences_Optimized_100(b *testing.B) {
	service := &SnapshotServiceImpl{
		diffService: &mockDiffService{delay: true},
	}
	snapshots := createTestSnapshots(100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		service.processSnapshotDifferences(snapshots, "/cache")
	}
}

func BenchmarkProcessSnapshotDifferences_ConcurrentOld_100(b *testing.B) {
	service := &SnapshotServiceImpl{
		diffService: &mockDiffService{delay: true},
	}
	snapshots := createTestSnapshots(100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		service.processSnapshotDifferencesConcurrentOld(snapshots, "/cache")
	}
}

func BenchmarkProcessSnapshotDifferences_Synchronous_100(b *testing.B) {
	service := &SnapshotServiceImpl{
		diffService: &mockDiffService{delay: true},
	}
	snapshots := createTestSnapshots(100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		service.processSnapshotDifferencesSynchronous(snapshots, "/cache")
	}
}

func BenchmarkProcessSnapshotDifferences_Optimized_1000(b *testing.B) {
	service := &SnapshotServiceImpl{
		diffService: &mockDiffService{delay: true},
	}
	snapshots := createTestSnapshots(1000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		service.processSnapshotDifferences(snapshots, "/cache")
	}
}

func BenchmarkProcessSnapshotDifferences_ConcurrentOld_1000(b *testing.B) {
	service := &SnapshotServiceImpl{
		diffService: &mockDiffService{delay: true},
	}
	snapshots := createTestSnapshots(1000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		service.processSnapshotDifferencesConcurrentOld(snapshots, "/cache")
	}
}

func BenchmarkProcessSnapshotDifferences_Synchronous_1000(b *testing.B) {
	service := &SnapshotServiceImpl{
		diffService: &mockDiffService{delay: true},
	}
	snapshots := createTestSnapshots(1000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		service.processSnapshotDifferencesSynchronous(snapshots, "/cache")
	}
}
