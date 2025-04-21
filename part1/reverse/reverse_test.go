package reverse

import "testing"

func Test1(t *testing.T) {
	input := "hello"
	expected := "holle"

	result := ReverseVowls(input)

	if result != expected {
		t.Errorf("ReverseVowls failed, input: %s, expect: %s, result: %s", input, expected, result)
	}
}

func Test2(t *testing.T) {
	input := "hello"
	expected := "holle"

	result := ReverseVowls(input)

	if result != expected {
		t.Errorf("ReverseVowls failed, input: %s, expect: %s, result: %s", input, expected, result)
	}
}
