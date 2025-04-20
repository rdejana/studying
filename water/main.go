package main

import (
	"fmt"
	"math"
)

func maxWater(arr []int) int {
	left := 1
	right := len(arr) - 2

	lMax := arr[left-1]
	rMax := arr[right+1]

	sum := 0

	for left <= right {
		if rMax <= lMax {
			temp := math.Max(0, float64(rMax-arr[right]))
			sum += int(temp)
			temp = math.Max(float64(rMax), float64(arr[right]))
			rMax = int(temp)
			right -= 1
		} else {
			temp := math.Max(0, float64(lMax-arr[left]))
			sum += int(temp)
			temp = math.Max(float64(lMax), float64(arr[left]))
			lMax = int(temp)
			left += 1
		}
	}

	return sum

}

func main() {
	arr := []int{2, 1, 5, 3, 1, 0, 4}
	max := maxWater(arr)
	fmt.Println("We got:", max)
}
