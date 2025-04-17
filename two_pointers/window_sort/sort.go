package main

import "fmt"

//come back to this

func find(arr []int) int {
	fmt.Println(arr)
	//how sorted is the orignal array
	lowerPtr := 0
	index := 0
	for index < len(arr) {
		if index+1 < len(arr) {
			v1 := arr[lowerPtr]
			v2 := arr[index+1]
			if v2 < v1 {
				break
			} else {
				lowerPtr++
				index++
			}
		}
	}

	highPtr := len(arr) - 1
	index = len(arr) - 1
	for index >= 0 {
		if index-1 > 0 {
			v1 := arr[highPtr]
			v2 := arr[index-1]
			//we need v1 to be greater than v2

			fmt.Println("--->", v2, v1)

			if v2 > v1 {

				break
			} else {
				highPtr--
				index--
			}
		}
	}

	//so we have a low and high, but doesn't mean we have the min/max
	min := lowerPtr
	index = lowerPtr
	for index < highPtr {
		v1 := arr[min]
		v2 := arr[index+1]
		fmt.Println(v1, v2)
		index++
	}

	//these are indexes, so
	length := 0
	if highPtr > lowerPtr {
		length = highPtr - lowerPtr + 1
	}
	fmt.Println(lowerPtr, highPtr)
	fmt.Println("length:", length)
	return length
}

func main() {
	t1 := []int{1, 2, 5, 3, 7, 10, 9, 12}
	fmt.Println(t1)
	find(t1)
}
