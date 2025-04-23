package stringer

import (
	"errors"
	"strings"
)

type Stack struct {
	Stack []string
}

func NStark() *Stack {
	s := &Stack{
		Stack: []string{},
	}
	return s
}

func NewStack(text string) *Stack {

	s := &Stack{
		Stack: []string{},
	}
	for _, r := range text {
		s.Push(string(r))
	}
	return s
}

func (s *Stack) Push(r string) {
	s.Stack = append(s.Stack, r)
}

func (s *Stack) PopValild() (string, error) {

	index := len(s.Stack) - 1
	if index < 0 {
		return "", errors.New("stack is empty")
	}
	v, err := s.Pop()
	if err != nil {
		return "", err
	}
	if "#" == v {
		return s.PopValild()
	} else {
		return v, nil
	}

}

func (s *Stack) Popper(count int) (string, error) {
	index := len(s.Stack) - 1
	if index < 0 {
		return "", errors.New("stack is empty")
	}
	top := s.Stack[index]
	s.Stack = s.Stack[:index]

	return top, nil
}

func (s *Stack) Pop() (string, error) {
	index := len(s.Stack) - 1
	if index < 0 {
		return "", errors.New("stack is empty")
	}
	top := s.Stack[index]
	s.Stack = s.Stack[:index]

	return top, nil
}

func (s *Stack) IsEmpty() bool {
	index := len(s.Stack) - 1
	return index < 0
}

func myStringer(str string) string {
	//"xyz##"
	stack := NStark() //char would be better
	index := len(str) - 1
	for index >= 0 {
		valid := getNextValidCharIndex(str, index)
		stack.Push(string(str[valid]))
		index = valid - 1
	}

	b := strings.Builder{}
	for stack.IsEmpty() == false {
		s, _ := stack.Pop()
		b.WriteString(s)
		//fmt.Println(s)
	}

	return b.String()
}

func getNextValidCharIndex(str string, index int) int {
	backspaceCount := 0
	for index >= 0 {
		if str[index] == '#' {
			backspaceCount++
		} else if backspaceCount > 0 {
			backspaceCount--
		} else {
			break
		}

		index--
	}

	return index
}
