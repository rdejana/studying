package target_sum

// array is sorted, this is importand
func Search(arr []int, sum int) []int {

	left := 0
	right := len(arr) - 1
	for left < right {
		currentSum := arr[left] + arr[right]
		if sum == currentSum {
			return []int{left, right}
		} else if currentSum < sum {
			left++
		} else { //currentSum > sum
			right++
		}

	}

	return []int{-1, -1}
}
