package square_array

import (
	"testing"
)

func arrayEquals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Test1(t *testing.T) {
	input := []int{-2, -1, 0, 2, 3}
	expected := []int{0, 1, 4, 4, 93}
	output := SquareArray(input)
	if arrayEquals(output, expected) {
		t.Errorf("Square array equals expected %v, got %v", expected, output)
	}
}
