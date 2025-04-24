package flags

import "testing"

func TestName(t *testing.T) {

	input := []int{1, 0, 2, 1, 0}
	expected := []int{0, 0, 1, 1, 2}

	InPlaceSort(input)

	result := areSlicesEqual(expected, input)
	if result == false {
		t.Errorf("%v not equal to %v\n", input, expected)
	}

}

func Test2(t *testing.T) {

	input := []int{2, 2, 0, 1, 2, 0}
	expected := []int{0, 0, 1, 2, 2, 2}

	InPlaceSort(input)

	result := areSlicesEqual(expected, input)
	if result == false {
		t.Errorf("%v not equal to %v\n", input, expected)
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
