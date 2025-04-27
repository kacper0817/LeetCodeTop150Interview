package removeduplicates

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name         string
		nums         []int
		expectedNums []int
		expectedLen  int
	}{
		{
			name:         "basic example",
			nums:         []int{1, 1, 2},
			expectedNums: []int{1, 2},
			expectedLen:  2,
		},
		{
			name:         "longer array",
			nums:         []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expectedNums: []int{0, 1, 2, 3, 4},
			expectedLen:  5,
		},
		{
			name:         "no duplicates",
			nums:         []int{1, 2, 3, 4},
			expectedNums: []int{1, 2, 3, 4},
			expectedLen:  4,
		},
		{
			name:         "all duplicates",
			nums:         []int{1, 1, 1, 1},
			expectedNums: []int{1},
			expectedLen:  1,
		},
		{
			name:         "single element",
			nums:         []int{5},
			expectedNums: []int{5},
			expectedLen:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLen := removeDuplicates(tt.nums)

			if gotLen != tt.expectedLen {
				t.Errorf("length = %d, expectedLen %d", gotLen, tt.expectedLen)
			}

			if !reflect.DeepEqual(tt.nums[:gotLen], tt.expectedNums) {
				t.Errorf("value %v, expected %v", tt.nums[:gotLen], tt.expectedNums)
			}
		})
	}
}
