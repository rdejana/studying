package reverse_vowls

import "strings"

const VOWELS = "aeiouAEIOU"

func ReverseVowels(str string) string {
	firstPtr := 0
	lastPtr := len(str) - 1

	array := []rune(str)

	for firstPtr < lastPtr {
		for firstPtr < lastPtr && !strings.ContainsRune(VOWELS, rune(array[firstPtr])) {
			firstPtr++
		}

		for firstPtr < lastPtr && !strings.ContainsRune(VOWELS, rune(array[lastPtr])) {
			lastPtr--
		}
		// Swap the vowels found at first and last
		array[firstPtr], array[lastPtr] = array[lastPtr], array[firstPtr]
		firstPtr++
		lastPtr--
	}

	return string(array)
}
