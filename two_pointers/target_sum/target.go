package target_sum

func Search(arr []int, targetSum int) []int {

	left := 0
	right := len(arr) - 1
	for left < right {
		currentSum := arr[left] + arr[right]
		if targetSum == currentSum {
			return []int{left, right}
		} else if currentSum < targetSum {
			left++
		} else { //currentSum > targetSum
			right--
		}

	}

	return []int{-1, -1}
}
