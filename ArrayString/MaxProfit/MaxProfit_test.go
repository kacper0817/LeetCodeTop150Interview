package maxprofit

import (
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "basic example",
			nums:     []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			name:     "no profit",
			nums:     []int{4, 0, 0, 0, 0, 0},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxProfit(tt.nums)
			if result != tt.expected {
				t.Errorf("value %v, expected %v", result, tt.expected)
			}
		})
	}
}
