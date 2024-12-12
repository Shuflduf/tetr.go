package main

import "testing"

func TestGetKickIndex(t *testing.T) {
	tests := []struct {
		prevIndex int
		newIndex  int
		expected  int
	}{
		{0, 1, 0},
		{1, 0, 1},
		{1, 2, 2},
		{2, 1, 3},
		{2, 3, 4},
		{3, 2, 5},
		{3, 0, 6},
		{0, 3, 7},
	}

	for _, tt := range tests {
		result := GetKickIndex(tt.prevIndex, tt.newIndex)
		if result != tt.expected {
			t.Errorf("GetKickIndex(%d, %d) = %d; want %d", tt.prevIndex, tt.newIndex, result, tt.expected)
		}
	}
}
