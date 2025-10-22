package utils

import "testing"

func TestMin(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"a is smaller", 3, 5, 3},
		{"b is smaller", 10, 2, 2},
		{"equal values", 7, 7, 7},
		{"negative numbers", -5, -3, -5},
		{"mixed signs", -2, 5, -2},
		{"zero and positive", 0, 10, 0},
		{"zero and negative", 0, -5, -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Min(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Min(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
