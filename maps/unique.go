package main

func firstUniqChar(str string) int {

	charMap := make(map[string]int)

	for _, c := range str {
		key := string(c)

		count, ok := charMap[key]
		if !ok {
			charMap[key] = 1
		} else {
			charMap[key] = count + 1
		}

	}

	for i, c := range str {
		key := string(c)

		if charMap[key] == 1 {
			return i
		}

	}

	return -1
}

func main() {

}
