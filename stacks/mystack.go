package stacks

import "errors"

type Stack[T any] struct {
	Stack []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		Stack: []T{},
	}
}

func (s *Stack[T]) Push(r T) {
	s.Stack = append(s.Stack, r)
}

func (s *Stack[T]) Pop() (T, error) {
	index := len(s.Stack) - 1
	if index < 0 {
		return *new(T), errors.New("stack is empty")
	}
	top := s.Stack[index]
	s.Stack = s.Stack[:index]

	return top, nil
}

func (s *Stack[T]) IsEmpty() bool {
	index := len(s.Stack) - 1
	return index < 0
}
