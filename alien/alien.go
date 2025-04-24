package alien

import (
	"errors"
	"fmt"
)

func alien(input []string) string {

	graph := make(map[string][]string)

	for i := 0; i < len(input)-1; i++ { //notice the < and not <=
		one := input[i]
		two := input[i+1]
		a, b, err := handleWords(one, two)
		graph[a] = append(graph[a], b)

		if err != nil {
			return ""
		}
	}

	fmt.Println(graph)

	return ""
}

func handleWords(one, two string) (string, string, error) {
	fmt.Println(one, two)
	lengthOne := len(one)
	lengthTwo := len(two)
	i, j := 0, 0
	for ; i < lengthOne && j < lengthTwo; i, j = i+1, j+1 {
		if one[i] != two[j] {
			fmt.Println("\t", string(one[i]), string(two[j]))
			return string(one[i]), string(two[j]), nil
		}
	}

	//so either all equal or we have an error???
	if lengthOne == lengthTwo {
		fmt.Println("we have a duplicate...")
	}
	//ab,abc
	if lengthTwo > lengthOne {
		fmt.Println("first word is a prefix of the second, no info")
	}

	if lengthOne > lengthTwo {
		fmt.Println("This is an error condition, out of order")
		return "", "", errors.New("one and two are out of order")
	}

	return "", "", nil

}
