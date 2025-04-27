package maxprofit

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			nums:     []int{4, 0, 0, 0, 0, 0},
			expected: 0,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := maxProfit(tt.nums)
			if result != tt.expected {
				t.Errorf("value %v, expected %v", result, tt.expected)
			}
		})
	}
}
