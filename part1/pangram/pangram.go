package pangram

import (
	"strings"
)

/*
A pangram is a sentence where every letter of the English alphabet appears at least once.
*/

func CheckIfPangram(sentence string) bool {
	sentence = strings.ToLower(sentence)
	m := make(map[string]bool)
	for _, letter := range sentence {
		if 97 <= letter && letter <= 122 {
			m[string(letter)] = true
		}
	}

	return len(m) == 26

	//return false
}

func initMap() map[string]bool {
	m := make(map[string]bool)

	return m
}
