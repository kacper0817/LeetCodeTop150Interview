package removeelement

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name         string
		nums         []int
		val          int
		expectedNums []int
		expectedLen  int
	}{
		{
			name:         "remove existing element",
			nums:         []int{3, 2, 2, 3},
			val:          3,
			expectedNums: []int{2, 2},
			expectedLen:  2,
		},
		{
			name:         "no element to remove",
			nums:         []int{1, 2, 3, 4},
			val:          5,
			expectedNums: []int{1, 2, 3, 4},
			expectedLen:  4,
		},
		{
			name:         "remove all elements",
			nums:         []int{2, 2, 2},
			val:          2,
			expectedNums: []int{},
			expectedLen:  0,
		},
		{
			name:         "empty array",
			nums:         []int{},
			val:          0,
			expectedNums: []int{},
			expectedLen:  0,
		},
		{
			name:         "mixed elements",
			nums:         []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:          2,
			expectedNums: []int{0, 1, 3, 0, 4},
			expectedLen:  5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLen := removeElement(tt.nums, tt.val)

			if gotLen != tt.expectedLen {
				t.Errorf("length = %d, expectedLen %d", gotLen, tt.expectedLen)
			}

			if !reflect.DeepEqual(tt.nums[:gotLen], tt.expectedNums) {
				t.Errorf("value %v, expected %v", tt.nums[:gotLen], tt.expectedNums)
			}
		})
	}
}
