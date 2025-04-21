package dups

import "testing"

func Test1(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	expected := false
	result := ContainsDuplicate(arr)

	if result != expected {
		t.Errorf("ContainsDuplicate(%v) returned %v, expected %v", arr, result, expected)
	}

}

func Test2(t *testing.T) {
	arr := []int{1, 2, 3, 1}
	expected := true
	result := ContainsDuplicate(arr)

	if result != expected {
		t.Errorf("ContainsDuplicate(%v) returned %v, expected %v", arr, result, expected)
	}

}

func Test3(t *testing.T) {
	arr := []int{3, 2, 6, -1, 2, 1}
	expected := true
	result := ContainsDuplicate(arr)

	if result != expected {
		t.Errorf("ContainsDuplicate(%v) returned %v, expected %v", arr, result, expected)
	}

}
