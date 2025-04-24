package two_pointers

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	input := []int{-2, -1, 0, 2, 3}
	expected := []int{0, 1, 4, 4, 9}
	fmt.Println(input)
	output := SquareArray(input)

	r := areSlicesEqual(output, expected)
	if !r {
		t.Errorf("Output does not equal expected %v != %v\n", output, expected)
	}

}

func Test2(t *testing.T) {
	input := []int{-3, -1, 0, 1, 2}
	expected := []int{0, 1, 1, 4, 9}
	fmt.Println(input)
	output := SquareArray(input)
	fmt.Println(output)
	r := areSlicesEqual(output, expected)
	if !r {
		t.Errorf("Output does not equal expected %v != %v\n", output, expected)
	}

}

func areSlicesEqual(one []int, two []int) bool {
	if len(one) != len(two) {
		return false
	}

	for i, _ := range one {
		if one[i] != two[i] {
			return false
		}
	}

	return true
}
