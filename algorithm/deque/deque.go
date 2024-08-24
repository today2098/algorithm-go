package deque

import (
	"container/list"
	"errors"
)

var ErrDequeEmpty = errors.New("Deque: deque is empty")

type Deque[T any] struct {
	data *list.List
}

func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{data: list.New()}
}

func (dq *Deque[T]) Empty() bool {
	return dq.Size() == 0
}

func (dq *Deque[T]) Size() int {
	return dq.data.Len()
}

func (dq *Deque[T]) Front() T {
	res := dq.data.Front()
	if res == nil {
		panic(ErrDequeEmpty)
	}
	return res.Value.(T)
}

func (dq *Deque[T]) Back() T {
	res := dq.data.Back()
	if res == nil {
		panic(ErrDequeEmpty)
	}
	return res.Value.(T)
}

func (dq *Deque[T]) PushFront(x T) {
	dq.data.PushFront(x)
}

func (dq *Deque[T]) PushFrontRange(v []T) {
	for i := len(v) - 1; i >= 0; i-- {
		dq.data.PushFront(v[i])
	}
}

func (dq *Deque[T]) PushBack(x T) {
	dq.data.PushBack(x)
}

func (dq *Deque[T]) PushBackRange(v []T) {
	for i := 0; i < len(v); i++ {
		dq.data.PushBack(v[i])
	}
}

func (dq *Deque[T]) PopFront() T {
	res := dq.data.Front()
	if res == nil {
		panic(ErrDequeEmpty)
	}
	return dq.data.Remove(res).(T)
}

func (dq *Deque[T]) PopBack() T {
	res := dq.data.Back()
	if res == nil {
		panic(ErrDequeEmpty)
	}
	return dq.data.Remove(res).(T)
}
