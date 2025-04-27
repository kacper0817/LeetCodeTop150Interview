package rotatearray

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		input    []int
		steps    int
		expected []int
	}{
		{
			input:    []int{1, 2, 3, 4, 5, 6, 7},
			steps:    3,
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			input:    []int{-1, -100, 3, 99},
			steps:    2,
			expected: []int{3, 99, -1, -100},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			rotate(tt.input, tt.steps)

			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Errorf("value %v, expected %v", tt.input, tt.expected)
			}
		})
	}
}
