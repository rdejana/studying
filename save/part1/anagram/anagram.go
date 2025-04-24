package anagram

func IsAnagram(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}

	//use runes, but since we are asci, cna use just the char value for now
	freqMap := make(map[int32]int)

	for _, ch := range a {
		freqMap[ch]++
	}

	for _, ch := range b {
		freqMap[ch]--
	}

	for _, v := range freqMap {
		if v != 0 {
			return false
		}
	}

	return true
}
