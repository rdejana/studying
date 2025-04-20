package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	stack []int
}

func NewStack() *Stack {
	return &Stack{
		stack: []int{},
	}
}

func (s *Stack) Push(r int) {
	s.stack = append(s.stack, r)
}

func (s *Stack) Peek() int {
	index := len(s.stack) - 1

	top := s.stack[index]
	//s.stack = s.stack[:index]

	return top

}

func (s *Stack) Pop() (int, error) {
	index := len(s.stack) - 1
	if index < 0 {
		return 0, errors.New("stack is empty")
	}
	top := s.stack[index]
	s.stack = s.stack[:index]

	return top, nil
}

func (s *Stack) IsEmpty() bool {
	index := len(s.stack) - 1
	return index < 0
}

func nge(arr []int) []int {

	output := make([]int, len(arr))
	stack := NewStack()

	for i := len(output) - 1; i >= 0; i-- {
		value := arr[i]

		for stack.IsEmpty() == false && stack.Peek() <= value {
			stack.Pop()
		}

		if stack.IsEmpty() {
			output[i] = -1
		} else {
			top := stack.Peek()
			if top > value {
				output[i] = top
			}

		}
		stack.Push(value)

	}

	return output
}

func main() {

	/*
		Input: [4, 5, 2, 25]
		 Output: [5, 25, 25, -1]
	*/

	input := []int{4, 5, 2, 25}

	output := nge(input)
	fmt.Println(output)

	input = []int{13, 7, 6, 12}
	output = nge(input)
	fmt.Println(output)

	input = []int{1, 2, 3, 4, 5}
	output = nge(input)
	fmt.Println(output)
	input = []int{13, 7, 6, 12}
	output = nge(input)
	fmt.Println(output)

}
