package main

import (
	"errors"
	"fmt"
	"strings"
)

type Stack struct {
	stack []rune
}

func NewStack() *Stack {
	return &Stack{
		stack: []rune{},
	}
}

func (s *Stack) Push(r rune) {
	s.stack = append(s.stack, r)
}

func (s *Stack) Pop() (rune, error) {
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

func (s *Stack) Dump() {

	for index := len(s.stack) - 1; index >= 0; index-- {
		fmt.Println(string(s.stack[index]))
	}
}

func ReverseString(str string) string {
	stack := NewStack()
	for _, c := range str {
		stack.Push(c)
	}

	builder := strings.Builder{}
	for stack.IsEmpty() != true {
		c, _ := stack.Pop()
		builder.WriteRune(c)
	}

	return builder.String()
}

func main() {
	s1 := "Hello World!"
	fmt.Println(s1, ReverseString(s1))

}
