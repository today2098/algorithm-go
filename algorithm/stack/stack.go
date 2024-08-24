package stack

import "errors"

var ErrStackEmpty = errors.New("Stack: stack is empty")

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: []T{}}
}

func (s *Stack[T]) Empty() bool {
	return s.Size() == 0
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func (s *Stack[T]) Top() T {
	if s.Empty() {
		panic(ErrStackEmpty)
	}
	return s.data[len(s.data)-1]
}

func (s *Stack[T]) Push(x T) {
	s.data = append(s.data, x)
}

func (s *Stack[T]) PushRange(v []T) {
	for i := 0; i < len(v); i++ {
		s.data = append(s.data, v[i])
	}
}

func (s *Stack[T]) Pop() T {
	if s.Empty() {
		panic(ErrStackEmpty)
	}
	res := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return res
}
