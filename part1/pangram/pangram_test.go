package pangram

import (
	"testing"
)

func Test0(t *testing.T) {
	s := "abcdefghijklmnopqrstuvwxyz"
	expected := true
	result := CheckIfPangram(s)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

}

func Test1(t *testing.T) {
	s := "TheQuickBrownFoxJumpsOverTheLazyDog"
	expected := true
	result := CheckIfPangram(s)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

}

func Test2(t *testing.T) {
	s := "This is not a pangram"
	expected := false
	result := CheckIfPangram(s)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

}
