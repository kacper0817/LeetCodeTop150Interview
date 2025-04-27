package removeduplicates

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		wantNums []int
		wantLen  int
	}{
		{
			name:     "basic example",
			nums:     []int{1, 1, 2},
			wantNums: []int{1, 2},
			wantLen:  2,
		},
		{
			name:     "longer array",
			nums:     []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			wantNums: []int{0, 1, 2, 3, 4},
			wantLen:  5,
		},
		{
			name:     "no duplicates",
			nums:     []int{1, 2, 3, 4},
			wantNums: []int{1, 2, 3, 4},
			wantLen:  4,
		},
		{
			name:     "all duplicates",
			nums:     []int{1, 1, 1, 1},
			wantNums: []int{1},
			wantLen:  1,
		},
		{
			name:     "single element",
			nums:     []int{5},
			wantNums: []int{5},
			wantLen:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLen := removeDuplicates(tt.nums)

			if gotLen != tt.wantLen {
				t.Errorf("removeDuplicates() length = %d, want %d", gotLen, tt.wantLen)
			}

			if !reflect.DeepEqual(tt.nums[:gotLen], tt.wantNums) {
				t.Errorf("removeDuplicates() nums = %v, want %v", tt.nums[:gotLen], tt.wantNums)
			}
		})
	}
}
