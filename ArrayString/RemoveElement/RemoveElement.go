package removeelement

func removeElement(nums []int, val int) int {
	temp := []int{}
	counter := 0
	for _, number := range nums {
		if number != val {
			temp = append(temp, number)
			counter++
		}
	}

	copy(nums, temp)
	return counter
}
