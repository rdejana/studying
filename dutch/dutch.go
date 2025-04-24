package dutch

func dutchFlag(arr []int) {
	low := 0
	high := len(arr) - 1
	for i := 0; i <= high; {
		//value := arr[i]
		if arr[i] == 0 {
			swap(arr, low, i)
			low++
			i++
		} else if arr[i] == 2 {
			swap(arr, high, i)
			high--
			//don't touch i as we don't know what we swapped
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
