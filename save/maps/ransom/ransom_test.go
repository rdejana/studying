package ransom

import "testing"

func Test1(t *testing.T) {
	note := "hello"
	magazine := "hellworld"
	expected := true
	result := CanWrite(note, magazine)

	if result != expected {
		t.Errorf("Unexpected result %t.  Expected %t", result, expected)
	}
}

func Test2(t *testing.T) {
	note := "notes"
	magazine := "stoned"
	expected := true
	result := CanWrite(note, magazine)

	if result != expected {
		t.Errorf("Unexpected result %t.  Expected %t", result, expected)
	}
}

func Test3(t *testing.T) {
	note := "apple"
	magazine := "pale"
	expected := false
	result := CanWrite(note, magazine)

	if result != expected {
		t.Errorf("Unexpected result %t.  Expected %t", result, expected)
	}
}
