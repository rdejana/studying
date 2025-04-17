package main

import "testing"

func Test1(t *testing.T) {

	arr := []int{1, 2, 5, 3, 7, 10, 9, 12}
	expected := 5
	r := find(arr)

	if r != expected {
		t.Errorf("Got %d, wanted %d\n", r, expected)
	}
}

func Test2(t *testing.T) {
	arr := []int{1, 3, 2, 0, -1, 7, 10}
	expected := 5
	r := find(arr)

	if r != expected {
		t.Errorf("Got %d, wanted %d\n", r, expected)
	}
}

func Test3(t *testing.T) {
	arr := []int{1, 2, 3}
	expected := 0
	r := find(arr)

	if r != expected {
		t.Errorf("Got %d, wanted %d\n", r, expected)
	}
}

func Test4(t *testing.T) {
	arr := []int{3, 2, 1}
	expected := 3
	r := find(arr)

	if r != expected {
		t.Errorf("Got %d, wanted %d\n", r, expected)
	}
}
