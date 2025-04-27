package mergesortedarray

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 5, 6},
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1:    []int{4, 0, 0, 0, 0, 0},
			m:        1,
			nums2:    []int{1, 2, 3, 5, 6},
			n:        5,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			nums1:    []int{1, 2, 5, 6, 0},
			m:        4,
			nums2:    []int{0},
			n:        1,
			expected: []int{0, 1, 2, 5, 6},
		},
		{
			nums1:    []int{1, 2, 5, 6},
			m:        4,
			nums2:    []int{0},
			n:        0,
			expected: []int{1, 2, 5, 6},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			merge(tt.nums1, tt.m, tt.nums2, tt.n)
			if !reflect.DeepEqual(tt.nums1, tt.expected) {
				t.Errorf("value %v, expected %v", tt.nums1, tt.expected)
			}
		})
	}
}
