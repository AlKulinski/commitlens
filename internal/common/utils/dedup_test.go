package utils

import (
	"testing"
)

func TestDedupBy(t *testing.T) {
	t.Run("deduplicates by ID", func(t *testing.T) {
		type Item struct {
			ID   int
			Name string
		}

		items := []Item{
			{ID: 1, Name: "first"},
			{ID: 2, Name: "second"},
			{ID: 1, Name: "duplicate"},
			{ID: 3, Name: "third"},
			{ID: 2, Name: "another duplicate"},
		}

		result := DedupBy(items, func(item Item) int {
			return item.ID
		})

		if len(result) != 3 {
			t.Errorf("expected 3 items, got %d", len(result))
		}

		if result[0].ID != 1 || result[0].Name != "first" {
			t.Errorf("expected first item to be {1, first}, got %v", result[0])
		}
		if result[1].ID != 2 || result[1].Name != "second" {
			t.Errorf("expected second item to be {2, second}, got %v", result[1])
		}
		if result[2].ID != 3 || result[2].Name != "third" {
			t.Errorf("expected third item to be {3, third}, got %v", result[2])
		}
	})

	t.Run("handles empty slice", func(t *testing.T) {
		var items []int
		result := DedupBy(items, func(item int) int {
			return item
		})

		if len(result) != 0 {
			t.Errorf("expected empty slice, got %d items", len(result))
		}
	})

	t.Run("handles slice with no duplicates", func(t *testing.T) {
		items := []string{"a", "b", "c"}
		result := DedupBy(items, func(item string) string {
			return item
		})

		if len(result) != 3 {
			t.Errorf("expected 3 items, got %d", len(result))
		}
	})

	t.Run("handles all duplicates", func(t *testing.T) {
		items := []int{5, 5, 5, 5}
		result := DedupBy(items, func(item int) int {
			return item
		})

		if len(result) != 1 {
			t.Errorf("expected 1 item, got %d", len(result))
		}
		if result[0] != 5 {
			t.Errorf("expected value 5, got %d", result[0])
		}
	})
}
