package anagram

func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	//really dumb way
	sumA := 0
	sumB := 0

	for i := 0; i < len(a); i++ {
		vA := a[i]
		vB := b[i]

		sumA = sumA + int(vA)
		sumB = sumB + int(vB)

	}

	return sumA == sumB

}
