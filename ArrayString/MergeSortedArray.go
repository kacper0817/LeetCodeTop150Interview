package arraystring

func RunMergeSortedArray() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}

	merge(nums1, 3, nums2, 3)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
	} else if n > 0 {
		copy(nums1, handleMerge(nums1, m, nums2, n))
	}
}

func handleMerge(nums1 []int, m int, nums2 []int, n int) []int {
	temp := make([]int, m+n)
	num1Index := 0
	num2Index := 0
	i := 0

	for num1Index > -1 || num2Index > -1 {
		if num1Index == -1 {
			temp[i] = nums2[num2Index]
			incrementCounter(&num2Index, n)
		} else if num2Index == -1 {
			temp[i] = nums1[num1Index]
			incrementCounter(&num1Index, m)
		} else if nums1[num1Index] < nums2[num2Index] {
			temp[i] = nums1[num1Index]
			incrementCounter(&num1Index, m)
		} else {
			temp[i] = nums2[num2Index]
			incrementCounter(&num2Index, n)
		}
		i++
	}
	return temp
}

func incrementCounter(value *int, maxValue int) {
	if *value+1 < maxValue {
		(*value)++
	} else {
		*value = -1
	}
}
