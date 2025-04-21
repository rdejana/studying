package pangram

import (
	"testing"
)

func Test1(t *testing.T) {
	sentence := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	expected := true
	//sentence = "az"
	result := checkIfPangram(sentence)
	if expected != result {
		t.Errorf("Expected %t, got %t", expected, result)
	}
}

func Test2(t *testing.T) {

	sentence := "This is not a pangram"
	expected := false
	//sentence = "az"
	result := checkIfPangram(sentence)
	if expected != result {
		t.Errorf("Expected %t, got %t", expected, result)
	}
}
