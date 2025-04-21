package shortlest

import "math"

func ShortestDistance(words []string, a, b string) int {
	shortestDistance := len(words)
	index1, index2 := -1, -1

	for i, word := range words {
		if word == a {
			index1 = i
		} else if word == b {
			index2 = i
		}

		if index1 != -1 && index2 != -1 {
			distance := index1 - index2
			absDistance := math.Abs(float64(distance))
			shortestDistance = int(math.Min(absDistance, float64(shortestDistance)))
		}

	}

	return shortestDistance

}
