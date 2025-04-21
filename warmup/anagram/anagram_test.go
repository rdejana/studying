package anagram

import (
	"testing"
)

func Test1(t *testing.T) {
	a := "listen"
	b := "silent"
	expected := true
	result := isAnagram(a, b)
	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func Test3(t *testing.T) {
	a := "hello"
	b := "world"
	expected := false
	result := isAnagram(a, b)

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
