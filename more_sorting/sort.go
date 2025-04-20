package main

import "fmt"

func merge(a, b []int) []int {

	out := make([]int, len(a)+len(b), len(a)+len(b))

	outPtr := 0
	aPtr := 0
	bPtr := 0

	fmt.Println(len(out))

	for outPtr < len(out) {

		if aPtr < len(a) && bPtr < len(b) {
			fmt.Println(a[aPtr], b[bPtr])
			if a[aPtr] <= b[bPtr] {
				out[outPtr] = a[aPtr]
				aPtr++
			} else {
				out[outPtr] = b[bPtr]
				bPtr++
			}
		} else if aPtr < len(a) {
			out[outPtr] = a[aPtr]
			aPtr++
		} else if bPtr < len(b) {
			out[outPtr] = b[bPtr]
			bPtr++
		} else {
			panic("odd...")
		}

		outPtr++
	}

	return out
}

func main() {

	arr1 := []int{1, 3, 4, 5}
	arr2 := []int{2, 4, 6, 8}

	merged := merge(arr1, arr2)

	fmt.Println(merged)

}
