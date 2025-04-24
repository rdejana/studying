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

// add to stack in reverse order
func arrayToStack(arr []int) *Stack {
	stack := NewStack()

	for i := 0; i < len(arr); i++ {
		stack.Push(arr[i])
	}

	return stack
	//add in

}

func main() {

	input := []int{34, 3, 31, 98, 92, 23}
	stack := arrayToStack(input)

	tempStack := NewStack()
	for stack.IsEmpty() == false {
		value, _ := stack.Pop()

		for tempStack.IsEmpty() == false {
			tempTop := tempStack.Peek()
			fmt.Println(tempTop, value)
			if tempTop > value {
				val, _ := tempStack.Pop()
				stack.Push(val)
			} else {
				break
			}
		}
		fmt.Println("pushing value to tempstack")
		tempStack.Push(value)
	}

	fmt.Println(tempStack.stack)

}
