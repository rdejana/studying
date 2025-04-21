package dups

func ContainsDuplicate(nums []int) bool {
	//see if we have a duplicate
	myMap := make(map[int]bool)
	for _, v := range nums {
		_, ok := myMap[v]
		if ok { //ok means we've seen this
			return true
		} else {
			myMap[v] = true //set the flag
		}

	}

	return false
}
