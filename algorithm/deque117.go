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

// Create a new deque.
func NewDeque117() *Deque117 {
	return &Deque117{
		Data: list.New(),
	}
}

// Checks if the deque is empty.
func (d *Deque117) Empty() bool {
	return d.Size() == 0
}

// Returns the number of elements.
func (d *Deque117) Size() int {
	return d.Data.Len()
}

// Returns the first element.
func (d *Deque117) Front() interface{} {
	res := d.Data.Front()
	if res == nil {
		panic(ErrDeque117Empty)
	}
	return res.Value
}

// Returns the last element.
func (d *Deque117) Back() interface{} {
	res := d.Data.Back()
	if res == nil {
		panic(ErrDeque117Empty)
	}
	return res.Value
}

// Inserts an element at the front.
func (d *Deque117) PushFront(x interface{}) {
	d.Data.PushFront(x)
}

// Inserts elements at the front.
func (d *Deque117) PushFrontRange(v []interface{}) {
	for i := len(v) - 1; i >= 0; i-- {
		d.Data.PushFront(v[i])
	}
}

// Inserts an element at the back.
func (d *Deque117) PushBack(x interface{}) {
	d.Data.PushBack(x)
}

// Inserts elements at the back.
func (d *Deque117) PushBackRange(v []interface{}) {
	for i := 0; i < len(v); i++ {
		d.Data.PushBack(v[i])
	}
}

// Removes and returns the first element.
func (d *Deque117) PopFront() interface{} {
	res := d.Data.Front()
	if res == nil {
		panic(ErrDeque117Empty)
	}
	return d.Data.Remove(res)
}

// Removes and returns the last element.
func (d *Deque117) PopBack() interface{} {
	res := d.Data.Back()
	if res == nil {
		panic(ErrDeque117Empty)
	}
	return d.Data.Remove(res)
}
