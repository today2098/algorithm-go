//go:build go1.18

package algorithm

import (
	"container/list"
	"errors"
)

var ErrDequeEmpty = errors.New("Deque: deque is empty")

type Deque[T any] struct {
	Data *list.List
}

func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{Data: list.New()}
}

func (dq *Deque[T]) Empty() bool {
	return dq.Size() == 0
}

func (dq *Deque[T]) Size() int {
	return dq.Data.Len()
}

func (dq *Deque[T]) Front() T {
	res := dq.Data.Front()
	if res == nil {
		panic(ErrDequeEmpty)
	}
	return res.Value.(T)
}

func (dq *Deque[T]) Back() T {
	res := dq.Data.Back()
	if res == nil {
		panic(ErrDequeEmpty)
	}
	return res.Value.(T)
}

func (dq *Deque[T]) PushFront(x T) {
	dq.Data.PushFront(x)
}

func (dq *Deque[T]) PushFrontRange(v []T) {
	for i := len(v) - 1; i >= 0; i-- {
		dq.Data.PushFront(v[i])
	}
}

func (dq *Deque[T]) PushBack(x T) {
	dq.Data.PushBack(x)
}

func (dq *Deque[T]) PushBackRange(v []T) {
	for i := 0; i < len(v); i++ {
		dq.Data.PushBack(v[i])
	}
}

func (dq *Deque[T]) PopFront() T {
	res := dq.Data.Front()
	if res == nil {
		panic(ErrDequeEmpty)
	}
	return dq.Data.Remove(res).(T)
}

func (dq *Deque[T]) PopBack() T {
	res := dq.Data.Back()
	if res == nil {
		panic(ErrDequeEmpty)
	}
	return dq.Data.Remove(res).(T)
}
