package ransom

func CanWrite(note string, magazine string) bool {

	letterMap := make(map[string]int)

	for _, c := range magazine {

		key := string(c)
		count, ok := letterMap[key]
		if !ok {
			count = 1
		} else {
			count += 1
		}
		letterMap[key] = count

	}

	//let's loop over
	for _, c := range note {
		key := string(c)
		count, ok := letterMap[key]
		if !ok {
			return false
		}

		if count > 0 {
			letterMap[key] = count - 1
		} else {
			return false
		}

	}

	return true
}
