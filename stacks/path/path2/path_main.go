package main

import (
	"errors"
	"strings"
)

type MyStack struct {
	Stack []string
}

func CreateNewStack() *MyStack {
	return &MyStack{
		Stack: []string{},
	}
}

func (s *MyStack) Push(r string) {
	s.Stack = append(s.Stack, r)
}

func (s *MyStack) Pop() (string, error) {
	index := len(s.Stack) - 1
	if index < 0 {
		return "", errors.New("stack is empty")
	}
	top := s.Stack[index]
	s.Stack = s.Stack[:index]

	return top, nil
}

func (s *MyStack) IsEmpty() bool {
	index := len(s.Stack) - 1
	return index < 0
}

type Solution struct{}

func (s *Solution) simplifyPath(path string) string {
	pathStack := CreateNewStack()
	//split first
	pathSlice := strings.Split(path, "/")

	for _, element := range pathSlice {
		//now . and "" should be considered no ops
		if element == "" || element == "." {
			//just skip and continue..
			continue
		} else if element == ".." {
			//this means try to pop
			if pathStack.IsEmpty() == false {
				pathStack.Pop()

			}
			continue
		}
		pathStack.Push(element)

	}

	//we can cheat a bit here
	output := "/" + strings.Join(pathStack.Stack, "/")

	return output
}

func SimplifyPath(input string) string {

	pathStack := CreateNewStack()
	//split first
	pathSlice := strings.Split(input, "/")

	for _, element := range pathSlice {
		//now . and "" should be considered no ops
		if element == "" || element == "." {
			//just skip and continue..
			continue
		} else if element == ".." {
			//this means try to pop
			if pathStack.IsEmpty() == false {
				pathStack.Pop()

			}
			continue
		}
		pathStack.Push(element)

	}

	//we can cheat a bit here
	output := "/" + strings.Join(pathStack.Stack, "/")

	return output
}

func main() {

}
