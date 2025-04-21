package duplicates

import (
	"slices"
)

func ContainsDuplicate(nums []int) bool {
	//just use a hashmap
	myMap := make(map[int]bool)
	for _, num := range nums {
		_, ok := myMap[num]
		if ok {
			//we found it
			return true //we have a duplicate
		} else {
			myMap[num] = true
		}
	}

	return false
}

func ContainsDuplicateWithSorting(nums []int) bool {
	slices.Sort(nums)
	//loop stops at the second to last index
	for i := 0; i < len(nums)-1; i++ {
		// if any two elements are the same, return true
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}
