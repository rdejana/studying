package main

import "fmt"

type Solution struct{}

func (s Solution) findMissingNumber(arr []int) int {
	n := len(arr) + 1
	// Find sum of all numbers from 1 to n.
	s1 := 0
	for i := 1; i <= n; i++ {
		s1 += i
	}

	// Subtract all numbers in input from sum.
	for _, num := range arr {
		s1 -= num
	}

	// s1, now, is the missing number
	return s1
}

func (s Solution) findMissingNumberXOR(arr []int) int {
	n := len(arr) + 1
	// find sum of all numbers from 1 to n.
	x1 := 0
	for i := 1; i <= n; i++ {
		x1 = x1 ^ i
	}

	// x2 represents XOR of all values in arr
	x2 := arr[0]
	for i := 1; i < n-1; i++ {
		x2 = x2 ^ arr[i]
	}

	// missing number is the xor of x1 and x2
	return x1 ^ x2
}

func main() {
	arr := []int{1, 5, 2, 6, 4}
	var s Solution
	fmt.Print("Missing number is: ", s.findMissingNumberXOR(arr))
}
