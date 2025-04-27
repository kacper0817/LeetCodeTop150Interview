package romantointeger

var mapToInt = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	length := len(s)
	arr := make([]int, length)
	for i := range length {
		value := mapToInt[s[i]]
		arr[i] = value
	}

	sum := 0
	for i := 0; i < length; i++ {
		if i+1 < length {
			followingValue := arr[i+1]
			if arr[i] < followingValue {
				sum += followingValue - arr[i]
				i++
				continue
			}
		}
		sum += arr[i]
	}
	return sum
}
