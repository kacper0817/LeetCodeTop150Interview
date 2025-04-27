package removeelement

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		val      int
		wantNums []int
		wantLen  int
	}{
		{
			name:     "remove existing element",
			nums:     []int{3, 2, 2, 3},
			val:      3,
			wantNums: []int{2, 2},
			wantLen:  2,
		},
		{
			name:     "no element to remove",
			nums:     []int{1, 2, 3, 4},
			val:      5,
			wantNums: []int{1, 2, 3, 4},
			wantLen:  4,
		},
		{
			name:     "remove all elements",
			nums:     []int{2, 2, 2},
			val:      2,
			wantNums: []int{},
			wantLen:  0,
		},
		{
			name:     "empty array",
			nums:     []int{},
			val:      0,
			wantNums: []int{},
			wantLen:  0,
		},
		{
			name:     "mixed elements",
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			wantNums: []int{0, 1, 3, 0, 4},
			wantLen:  5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLen := removeElement(tt.nums, tt.val)

			if gotLen != tt.wantLen {
				t.Errorf("removeElement() length = %d, want %d", gotLen, tt.wantLen)
			}

			if !reflect.DeepEqual(tt.nums[:gotLen], tt.wantNums) {
				t.Errorf("removeElement() nums = %v, want %v", tt.nums[:gotLen], tt.wantNums)
			}
		})
	}
}
