package main

import (
	"fmt"
	"math"
)

func main() {

	//110
	//001
	num := 6
	//0 1 10 11 100 101 110 111

	bitCount := 0
	n := num
	for n > 0 {
		fmt.Println("n ->", n)
		bitCount++
		n = n >> 1
		fmt.Println("n ->", n)
	}

	allBitsSet := int(math.Pow(2, float64(bitCount))) - 1

	v := num ^ allBitsSet

	fmt.Println(bitCount, allBitsSet)
	fmt.Println(v)
}
