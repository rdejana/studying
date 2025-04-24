package square_array

func SquareArray(arr []int) []int {
	//we could have neg numbers...
	left := 0
	right := len(arr) - 1
	result := []int{}
	//slightly lazy and just using append
	//could track the highest index decrement
	for left <= right {
		leftValue := arr[left] * arr[left]
		rightValue := arr[right] * arr[right]
		if leftValue > rightValue {
			result = append([]int{leftValue}, result...)
			left++
		} else if leftValue <= rightValue {
			result = append([]int{rightValue}, result...)
			right--
		}

	}

	return result
}
