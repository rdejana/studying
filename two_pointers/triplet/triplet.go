package triplet

import (
	"slices"
)

func searchTriplets(arr []int) [][]int {
	slices.Sort(arr)
	triplets := make([][]int, 0)
	for i := 0; i < len(arr)-2; i++ {
		if i > 0 && arr[i] == arr[i-1] { // skip same element to avoid duplicate triplets
			continue
		}
		searchPair(arr, -arr[i], i+1, &triplets)
	}
	return triplets
}

func searchPair(arr []int, targetSum int, left int, triplets *[][]int) {
	right := len(arr) - 1
	for left < right {
		currentSum := arr[left] + arr[right]
		if currentSum == targetSum {
			*triplets = append(*triplets, []int{-targetSum, arr[left], arr[right]})
			left++
			right--
			//avoid duplicates
			for left < right && arr[left] == arr[left-1] {
				left++ // skip same element to avoid duplicate triplets
			}
			for left < right && arr[right] == arr[right+1] {
				right-- // skip same element to avoid duplicate triplets
			}
		} else if targetSum > currentSum {
			left++ // we need a pair with a bigger sum
		} else {
			right-- // we need a pair with a smaller sum
		}

	}
}
