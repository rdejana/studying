package palindrome

import (
	"strings"
	"unicode"
)

func IsPalindrome(s string) bool {
	first := 0
	last := len(s) - 1

	for first < last {
		for first < last && !isLetterOrDigit((rune(s[first]))) {
			first++
		}

		for first < last && !isLetterOrDigit((rune(s[last]))) {
			last--
		}

		//do they match???
		if strings.ToLower((string(s[first]))) != strings.ToLower(string(s[last])) {
			return false
		}

		first++
		last--
	}

	return true
}

// helper
func isLetterOrDigit(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
