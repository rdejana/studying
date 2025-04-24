package two_pointers

func SquareArray(input []int) []int {

	left := 0
	right := len(input) - 1
	output := make([]int, len(input))

	//input is sorted and we know that "negative" will be positive

	for i := len(output) - 1; i >= 0; i-- {
		leftSquare := input[left] * input[left]
		rightSquare := input[right] * input[right]

		if leftSquare > rightSquare {
			output[i] = leftSquare
			left++
		} else {
			output[i] = rightSquare
			right--
		}

	}

	return output
}
