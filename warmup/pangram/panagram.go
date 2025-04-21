package pangram

import (
	"strings"
)

func checkIfPangram(sentence string) bool {
	sentence = strings.ToLower(sentence)
	m := make(map[string]bool)
	for _, letter := range sentence {
		//isLetter := unicode.IsLetter(letter)
		if letter >= 97 && letter <= 122 {
			m[string(letter)] = true
		}
	}

	return len(m) == 26
}
