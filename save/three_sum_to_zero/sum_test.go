package three_sum_to_zero

import "testing"

func Test1(t *testing.T) {
	input := []int{-3, 0, 1, 2, -1, 1, -2}
	expected := [][]int{{-3, 1, 2}, {-2, 0, 2}, {-2, 1, 1}, {-1, 0, 1}}

	output := searchTriplets(input)
	result := areEqual(expected, output)
	if result != true {
		t.Errorf("%v is not equal to %v \n", input, output)
	}

}

func Test2(t *testing.T) {
	input := []int{-5, 2, -1, -2, 3}
	expected := [][]int{{-5, 2, 3}, {-2, -1, 3}}

	output := searchTriplets(input)
	result := areEqual(expected, output)
	if result != true {
		t.Errorf("%v is not equal to %v \n", input, output)
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

func areEqual(one [][]int, two [][]int) bool {
	if len(one) != len(two) {
		return false
	}

	for i, v := range one {

		r := areSlicesEqual(v, two[i])
		if r == false {
			return false
		}

	}

	return true
}
