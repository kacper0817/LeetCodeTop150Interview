package majorityelement

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "basic example",
			nums: []int{3, 2, 3},
			want: 3,
		},
		{
			name: "1 element nums",
			nums: []int{2, 2, 1, 1, 1, 2, 2},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			majorityElement(tt.nums)
			if !reflect.DeepEqual(tt.want, tt.want) {
				t.Errorf("majorityElement() = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}
