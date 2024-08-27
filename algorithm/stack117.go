package algorithm

import "errors"

var ErrStack117Empty = errors.New("Stack117: stack is empty")

// A stack structure.
type Stack117 struct {
	Data []interface{}
}

// Create a stack.
func NewStack117() *Stack117 {
	return &Stack117{
		Data: []interface{}{},
	}
}

// Checks if the stack is empty.
func (s *Stack117) Empty() bool {
	return s.Size() == 0
}

// Returns the number of elements.
func (s *Stack117) Size() int {
	return len(s.Data)
}

// Returns the top element.
func (s *Stack117) Top() interface{} {
	if s.Empty() {
		panic(ErrStack117Empty)
	}
	return s.Data[len(s.Data)-1]
}

// Inserts an element at the top.
func (s *Stack117) Push(x interface{}) {
	s.Data = append(s.Data, x)
}

// Inserts elements at the top.
func (s *Stack117) PushRange(v []interface{}) {
	for i := 0; i < len(v); i++ {
		s.Data = append(s.Data, v[i])
	}
}

// Removes and returns the top element.
func (s *Stack117) Pop() interface{} {
	if s.Empty() {
		panic(ErrStack117Empty)
	}
	res := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	return res
}
