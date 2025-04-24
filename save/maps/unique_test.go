package main

import "testing"

func TestName(t *testing.T) {

	input := "apple"
	expected := 0
	output := firstUniqChar(input)

	if output != expected {
		t.Errorf("%d is not equal to %d", expected, output)
	}
}

func Test2(t *testing.T) {

	input := "abcab"
	expected := 2
	output := firstUniqChar(input)

	if output != expected {
		t.Errorf("%d is not equal to %d", expected, output)
	}
}

func Test3(t *testing.T) {

	input := "abab"
	expected := -1
	output := firstUniqChar(input)

	if output != expected {
		t.Errorf("%d is not equal to %d", expected, output)
	}
}
