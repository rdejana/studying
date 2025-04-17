package flags

func sort(input []int) {
	lowerPtr, highPtr := 0, len(input)-1

	for i := 0; i <= highPtr; {
		//	fmt.Println(lowerPtr, highPtr)
		value := input[i]

		if value == 0 {

			swap(input, i, lowerPtr)
			lowerPtr++
			i++

		} else if value == 2 {
			//need to revist the swaped value..

			swap(input, i, highPtr)
			highPtr--

		} else {
			i++
		}

	}

}

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
