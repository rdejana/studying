package shortlest

import (
	"testing"
)

func TestName(t *testing.T) {
	words1 := []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	word11 := "fox"
	word21 := "dog"
	expectedOutput1 := 5
	actualOutput1 := ShortestDistance(words1, word11, word21)

	if expectedOutput1 != actualOutput1 {
		t.Errorf("Expected: %d, Actual: %d", expectedOutput1, actualOutput1)
	}

}

func Test2(t *testing.T) {
	words := []string{"a", "b", "c", "d", "a", "b"}
	word1 := "a"
	word2 := "b"
	expectedOutput1 := 1
	actualOutput1 := ShortestDistance(words, word1, word2)

	if expectedOutput1 != actualOutput1 {
		t.Errorf("Expected: %d, Actual: %d", expectedOutput1, actualOutput1)
	}
}

func Test3(t *testing.T) {
	// Test case 3
	words := []string{"a", "c", "d", "b", "a"}
	word1 := "a"
	word2 := "b"
	expectedOutput1 := 1
	actualOutput1 := ShortestDistance(words, word1, word2)

	if expectedOutput1 != actualOutput1 {
		t.Errorf("Expected: %d, Actual: %d", expectedOutput1, actualOutput1)
	}
}

func Test4(t *testing.T) {
	// Test case 3
	words := []string{"a", "b", "c", "d", "e"}
	word1 := "a"
	word2 := "e"
	expectedOutput1 := 4
	actualOutput1 := ShortestDistance(words, word1, word2)

	if expectedOutput1 != actualOutput1 {
		t.Errorf("Expected: %d, Actual: %d", expectedOutput1, actualOutput1)
	}
}
