package reverse

import (
	"strings"
)

//come back to

func ReverseVowls(s string) string {

	vowels := "aeiouAEIOU"
	first, last := 0, len(s)-1
	array := []rune(s)
	for first < last {
		for first < last && !strings.ContainsRune(vowels, array[first]) {
			first++
		}
		// Skip non-vowel characters from the end
		for first < last && !strings.ContainsRune(vowels, array[last]) {
			last--
		}
		// Swap the vowels found at first and last
		array[first], array[last] = array[last], array[first]
		first++
		last--
	}

	return string(array)
}
