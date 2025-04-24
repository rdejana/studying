package palindrome

import (
	"strings"
	"unicode"
)

func IsPalindrome(s string) bool {
	//lazy
	s = strings.ToLower(s)
	//use runes
	builder := strings.Builder{}
	for _, s := range s {
		currChar := rune(strings.ToLower(string(s))[0])
		if unicode.IsLetter(currChar) || unicode.IsNumber(currChar) {
			//fmt.Print(string(currChar))
			builder.WriteRune(currChar)
		}

	}
	testString := builder.String()
	start := 0
	end := len(testString) - 1
	for start < end {
		if testString[start] != testString[end] {
			return false
		}
		start++
		end--
	}

	return true
}
