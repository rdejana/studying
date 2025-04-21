package palindrome

import (
	"testing"
)

func Test1(t *testing.T) {
	str := "A man, a plan, a canal, Panama!"
	result := IsPalindrome(str)
	expected := true
	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func Test2(t *testing.T) {
	str := "Was it a car or a cat I saw?"
	result := IsPalindrome(str)
	expected := true
	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func Test3(t *testing.T) {
	str := "Was it a car ab or a cat I saw?"
	result := IsPalindrome(str)
	expected := false
	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
