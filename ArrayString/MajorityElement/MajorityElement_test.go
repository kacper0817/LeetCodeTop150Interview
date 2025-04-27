package majorityelement

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{3, 2, 3},
			expected: 3,
		},
		{
			nums:     []int{2, 2, 1, 1, 1, 2, 2},
			expected: 2,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {

			majorityElement(tt.nums)
			if !reflect.DeepEqual(tt.nums, tt.expected) {
				t.Errorf("value %v, expected %v", tt.nums, tt.expected)
			}
		})
	}
}
