package main

import (
	"fmt"
	"strings"
	"studying/stacks"
)

func main() {

	input := "/a//b////c/d//././/.."

	stringStack := stacks.NewStack[string]()

	a, err := stringStack.Pop()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("should be some default->|", a, "|")

	split := strings.Split(input, "/")

	fmt.Println(split)
	for _, c := range split {
		if c == "" || c == "." {
			fmt.Println("found a skip char")
			continue
		} else if c == ".." {
			if stringStack.IsEmpty() == false {
				stringStack.Pop()
				continue
			}
		}
		stringStack.Push(c)
	}

	fmt.Println(stringStack.Stack)
	output := "/" + strings.Join(stringStack.Stack, "/")
	fmt.Println(output)

}
