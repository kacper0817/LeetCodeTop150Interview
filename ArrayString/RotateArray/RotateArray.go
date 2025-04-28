package rotatearray

func rotate(nums []int, k int) {
	if len(nums) <= 1 {
		return
	}

	shiftedNums := make([]int, len(nums))

	for i := range nums {
		newPosition := (i + k) % len(nums)
		shiftedNums[newPosition] = nums[i]
	}
	copy(nums, shiftedNums)

}
