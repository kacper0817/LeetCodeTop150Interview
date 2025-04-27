package removeduplicates2

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
			nums:         []int{1, 1, 1, 2, 2, 3},
			expectedNums: []int{1, 1, 2, 2, 3},
			expectedLen:  5,
		},
		{
			name:         "basic example",
			nums:         []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			expectedNums: []int{0, 0, 1, 1, 2, 3, 3},
			expectedLen:  7,
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
