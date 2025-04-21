package h_index

import (
	"fmt"
	"slices"
)

func H_Inddex(nums []int) int {
	//sort first - n log n for quicksort stuff
	/*
		h = 0 to start
		5 -> 1 with >1 citation (5 >= 0)
		4 -> 2 with >2 (4 >=1)
		3 -> 3 with >3 (
		[0 1 3 4 5]
	*/
	slices.Sort(nums) //depends on sorting
	fmt.Println(nums)
	h := 0
	//process by going backwards we know that
	//
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] >= h {
			h++
		} else {
			break
		}
	}

	return h
}
