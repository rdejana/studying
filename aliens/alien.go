package main

import "fmt"

// https://www.youtube.com/watch?v=zTJx9ZKy8Jg&list=PLMPSjUOg8tP2wSulRK6aai19kdFJelT-P&index=1
func main() {

	words := []string{"wrt", "wrf", "er", "ett", "rftt"}

	//get "vertices" -> each char, store in a map

	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			word1 := words[i]
			word2 := words[j]

			fmt.Println(word1, word2)
			//break
		}

	}
}
