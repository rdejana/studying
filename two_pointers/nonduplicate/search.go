package nonduplicate

import "fmt"

// review
// https://www.designgurus.io/course-play/grokking-the-coding-interview/doc/find-nonduplicate-number-instances-easy
func moveElements(arr []int) int {
	nextNonDuplicate := 1
	for i := 1; i < len(arr); i++ {
		fmt.Println(arr[nextNonDuplicate-1], arr[i])
		if arr[nextNonDuplicate-1] != arr[i] {
			// If they are not equal, update the element at the nextNonDuplicate position with the current element.
			arr[nextNonDuplicate] = arr[i]

			// Increment the nextNonDuplicate pointer.
			nextNonDuplicate++
		}
	}

	return nextNonDuplicate
}
