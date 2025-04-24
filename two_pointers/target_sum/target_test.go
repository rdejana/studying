package target_sum

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
	arr := []int{1, 2, 3, 4, 6}
	output := Search(arr, 6)
	expected := []int{1, 3}
	if !arrayEquals(output, expected) {
		t.Errorf("Expected: %d, got: %d", expected, output)
	}
}

func Test2(t *testing.T) {
	arr := []int{2, 5, 9, 11}
	output := Search(arr, 11)
	expected := []int{0, 2}
	if !arrayEquals(output, expected) {
		t.Errorf("Expected: %d, got: %d", expected, output)
	}
}
