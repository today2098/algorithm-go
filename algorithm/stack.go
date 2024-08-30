package algorithm

import "errors"

var ErrStackEmpty = errors.New("Stack: stack is empty")

type Stack[T any] struct {
	Data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{Data: []T{}}
}

func (s *Stack[T]) Empty() bool {
	return s.Size() == 0
}

func (s *Stack[T]) Size() int {
	return len(s.Data)
}

func (s *Stack[T]) Top() T {
	if s.Empty() {
		panic(ErrStackEmpty)
	}
	return s.Data[len(s.Data)-1]
}

func (s *Stack[T]) Push(x T) {
	s.Data = append(s.Data, x)
}

func (s *Stack[T]) PushRange(v []T) {
	for i := 0; i < len(v); i++ {
		s.Data = append(s.Data, v[i])
	}
}

func (s *Stack[T]) Pop() T {
	if s.Empty() {
		panic(ErrStackEmpty)
	}
	res := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	return res
}
