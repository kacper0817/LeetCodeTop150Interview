package removeduplicates

func removeDuplicates(nums []int) int {
	length := len(nums)

	index := 0

	uniqueNumsCounter := 0

	if length == 1 {
		return 1
	}

	for index < length-1 {
		currentCount := countRecursive(nums, index, 0)

		nums[uniqueNumsCounter] = nums[index]
		uniqueNumsCounter++

		index += currentCount + 1
	}

	lastValue := nums[length-1]
	if lastValue != nums[length-2] {
		nums[uniqueNumsCounter] = nums[length-1]
		uniqueNumsCounter++
	}

	return uniqueNumsCounter
}

func countRecursive(nums []int, index int, counter int) int {
	currentValue := nums[index]
	if index < len(nums)-1 && currentValue == nums[index+1] {
		index++
		counter = countRecursive(nums, index, counter+1)
	}

	return counter
}
