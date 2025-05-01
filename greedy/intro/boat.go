package intro

import "slices"

func numberOfBoats(people []int, limit int) int {
	slices.Sort(people)
	left := 0
	right := len(people) - 1
	totalBoatsNeeded := 0
	for left < right {
		if people[left]+people[right] == limit {
			left++
		} //if false, we are just taking the bigger one

		right-- //just take
		totalBoatsNeeded++
	}

	return totalBoatsNeeded
}
