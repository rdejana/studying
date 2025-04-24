package anagram

import "testing"

func Test0(t *testing.T) {
	a := "listen"
	b := "silen"
	expected := false
	output := IsAnagram(a, b)

	if output != expected {
		t.Errorf("Expected %v, got %v", expected, output)
	}
}

func Test1(t *testing.T) {
	a := "listen"
	b := "silent"
	expected := true
	output := IsAnagram(a, b)

	if output != expected {
		t.Errorf("Expected %v, got %v", expected, output)
	}
}
