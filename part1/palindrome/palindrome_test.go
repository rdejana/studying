package palindrome

import (
	"testing"
)

func Test1(t *testing.T) {

	input := "A man, a plan, a canal, Panama!"
	expected := true
	result := IsPalindrome(input)
	if expected != result {
		t.Errorf("Expected: %t, got: %t", expected, result)
	}

}

func Test2(t *testing.T) {

	input := "Was it a car or a cat I saw?"
	expected := true
	result := IsPalindrome(input)
	if expected != result {
		t.Errorf("Expected: %t, got: %t", expected, result)
	}

}
