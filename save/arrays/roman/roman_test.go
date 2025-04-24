package roman

import "testing"

func Test1(t *testing.T) {
	input := "XLII"
	expected := 42
	output := RomanToInt(input)
	if output != expected {
		t.Error("Expected", expected, "Got", output)
	}

}

func Test2(t *testing.T) {
	input := "CXCIV"
	expected := 194
	output := RomanToInt(input)
	if output != expected {
		t.Error("Expected", expected, "Got", output)
	}

}

func Test3(t *testing.T) {
	input := "MMMCDXLIV"
	expected := 3444
	output := RomanToInt(input)
	if output != expected {
		t.Error("Expected", expected, "Got", output)
	}

}
