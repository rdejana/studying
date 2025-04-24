package main

import (
	"errors"
	"fmt"
	"strings"
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

func IntToBinaryWithStack(num int) string {
	stack := NewStack()

	for num > 0 {
		stack.Push(num % 2)
		num /= 2
	}

	builder := strings.Builder{}
	for !stack.IsEmpty() {
		v, _ := stack.Pop()
		builder.WriteString(fmt.Sprintf("%d", v))
	}

	return builder.String()
}

func main() {

	binString := IntToBinaryWithStack(18)
	fmt.Println(binString)

}
