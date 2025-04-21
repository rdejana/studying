package good_pair

func NumberOfGoodPairs(nums []int) int {
	pairCount := 0
	numberMap := make(map[int]int)

	for _, v := range nums {
		numberMap[v]++
		pairCount = pairCount + numberMap[v] - 1
	}

	return pairCount
}
