package window

import (
	"fmt"
	"math"
)

func sort(arr []int) int {
	fmt.Println(arr)
	length := len(arr) - 1
	low := 0
	high := length //len(arr) - 1

	for low < length && arr[low] <= arr[low+1] {
		low++
	}

	if low == length {
		//all sorted return 0
		return 0
	}

	//now we just go in reverse
	for high > 0 && arr[high] >= arr[high-1] {
		high--
	}

	fmt.Println("low:", low, "high", high)
	//we need to keep going because our low and high may be in the subarry we found

	subarrayMax := math.MinInt32
	subarrayMin := math.MaxInt32

	for k := low; k <= high; k++ {
		if arr[k] > subarrayMax {
			subarrayMax = arr[k]
		}
		if arr[k] < subarrayMin {
			subarrayMin = arr[k]
		}
	}

	fmt.Println("submin", subarrayMin, "submax", subarrayMax)
	//we may need to move left
	for low > 0 && arr[low-1] > subarrayMin {
		low--
	}
	//we can also move right...
	for high < length && arr[high+1] < subarrayMax {
		high++
	}

	fmt.Println("low:", low, "high", high)
	return high - low + 1 //+1 is needed to go from 0 based to 1 based for length
}
