package triplet

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	input := []int{-5, 2, -1, -2, 3}
	output := searchTriplets(input)

	fmt.Println(output)
}
