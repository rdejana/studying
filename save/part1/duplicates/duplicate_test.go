package duplicates

import "testing"

func Test1(t *testing.T) {
	t1 := []int{3, 2, 6, -1, 2, 1}
	expected := true

	result := ContainsDuplicateWithSorting(t1)
	if result != expected {
		t.Errorf("ContainsDuplicate failed for %v and %v", t1, expected)
	}
}

func Test2(t *testing.T) {
	t1 := []int{3, 2, 6, -1, 1}
	expected := false

	result := ContainsDuplicateWithSorting(t1)
	if result != expected {
		t.Errorf("ContainsDuplicate failed for %v and %v", t1, expected)
	}
}

func Test3(t *testing.T) {
	t1 := []int{1, 2, 3, 1}
	expected := true

	result := ContainsDuplicateWithSorting(t1)
	if result != expected {
		t.Errorf("ContainsDuplicate failed for %v and %v", t1, expected)
	}
}

func Test4(t *testing.T) {
	t1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := false

	result := ContainsDuplicateWithSorting(t1)
	if result != expected {
		t.Errorf("ContainsDuplicate failed for %v and %v", t1, expected)
	}
}
