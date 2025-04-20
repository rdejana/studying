package main

import "fmt"

func sum(word string) int32 {

	var sum int32 = 0
	for _, c := range word {
		sum = sum + c
	}

	return sum

}

func main() {

	arr := []string{"act", "god", "cat", "dog", "tac"}

	for _, str := range arr {
		key := sum(str)
		fmt.Println(str, "->", key)
	} //now just how to return
	//can use key to store index for the result array...

}
