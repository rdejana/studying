package main

import (
	"fmt"
)

func main() {
	/*
		Given an array arr. Find the majority element in the array. If no majority exists, return -1. A majority element in an array is an element that
		appears strictly more than arr.size() / 2 times in the array.
	*/

	input := []int{3}
	length := len(input)
	majority := float64(length) / float64(2)

	freq := make(map[int]int)

	for _, c := range input {
		count, ok := freq[c]
		if !ok {
			count = 1
		} else {
			count = count + 1
		}
		freq[c] = count

	}

	max := -1
	for k, v := range freq {
		if float64(v) > majority {
			max = k
			break

		}

	}
	fmt.Println("Majority is ", max)

}
