package removeelement

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		nums         []int
		val          int
		expectedNums []int
		expectedLen  int
	}{
		{
			nums:         []int{3, 2, 2, 3},
			val:          3,
			expectedNums: []int{2, 2},
			expectedLen:  2,
		},
		{
			nums:         []int{1, 2, 3, 4},
			val:          5,
			expectedNums: []int{1, 2, 3, 4},
			expectedLen:  4,
		},
		{
			nums:         []int{2, 2, 2},
			val:          2,
			expectedNums: []int{},
			expectedLen:  0,
		},
		{
			nums:         []int{},
			val:          0,
			expectedNums: []int{},
			expectedLen:  0,
		},
		{
			nums:         []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:          2,
			expectedNums: []int{0, 1, 3, 0, 4},
			expectedLen:  5,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
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
