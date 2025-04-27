package rotatearray

import (
	"slices"
)

func rotate(nums []int, k int) {
	splitIndex := k + 1
	splited := slices.Clone(nums[splitIndex:])

	for i := splitIndex - 1; i >= 0; i-- {
		nums[splitIndex+i-1] = nums[i]
	}

	for i := range len(splited) {
		nums[i] = splited[i]
	}
}
