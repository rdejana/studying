package h_index

import "testing"

func Test1(t *testing.T) {
	nums := []int{4, 3, 0, 1, 5}
	expected := 3

	output := H_Inddex(nums)
	if expected != output {
		t.Errorf("expected:%v, actual:%v", expected, output)
	}
}

func Test2(t *testing.T) {

	nums := []int{10, 8, 5, 4, 3, 7, 2, 1}
	expected := 4

	output := H_Inddex(nums)
	if expected != output {
		t.Errorf("expected:%v, actual:%v", expected, output)
	}
}

func Test3(t *testing.T) {

	nums := []int{10, 8, 5, 4, 3, 7, 4, 2, 1}
	expected := 4

	output := H_Inddex(nums)
	if expected != output {
		t.Errorf("expected:%v, actual:%v", expected, output)
	}
}
