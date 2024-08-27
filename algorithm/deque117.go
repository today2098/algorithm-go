package algorithm

import (
	"container/list"
	"errors"
)

var ErrDeque117Empty = errors.New("Deque117: deque is empty")

// A deque data structure.
type Deque117 struct {
	Data *list.List
}

// Create a deque.
func NewDeque117() *Deque117 {
	return &Deque117{Data: list.New()}
}

// Checks if the deque is empty.
func (dq *Deque117) Empty() bool {
	return dq.Size() == 0
}

// Returns the number of elements.
func (dq *Deque117) Size() int {
	return dq.Data.Len()
}

// Returns the first element of deque.
func (dq *Deque117) Front() interface{} {
	res := dq.Data.Front()
	if res == nil {
		panic(ErrDeque117Empty)
	}
	return res.Value
}

// Returns the last element of deque.
func (dq *Deque117) Back() interface{} {
	res := dq.Data.Back()
	if res == nil {
		panic(ErrDeque117Empty)
	}
	return res.Value
}

// Inserts an element to the front of deque.
func (dq *Deque117) PushFront(x interface{}) {
	dq.Data.PushFront(x)
}

// Inserts elements to the front of deque.
func (dq *Deque117) PushFrontRange(v []interface{}) {
	for i := len(v) - 1; i >= 0; i-- {
		dq.Data.PushFront(v[i])
	}
}

// Inserts an element to the back of deque.
func (dq *Deque117) PushBack(x interface{}) {
	dq.Data.PushBack(x)
}

// Inserts elements to the back of deque.
func (dq *Deque117) PushBackRange(v []interface{}) {
	for i := 0; i < len(v); i++ {
		dq.Data.PushBack(v[i])
	}
}

// Removes the first element of deque and returns the element.
func (dq *Deque117) PopFront() interface{} {
	res := dq.Data.Front()
	if res == nil {
		panic(ErrDeque117Empty)
	}
	return dq.Data.Remove(res)
}

// Removes the last element of deque and returns the element.
func (dq *Deque117) PopBack() interface{} {
	res := dq.Data.Back()
	if res == nil {
		panic(ErrDeque117Empty)
	}
	return dq.Data.Remove(res)
}
