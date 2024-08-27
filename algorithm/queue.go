//go:build go1.18

package algorithm

import "errors"

var ErrQueueEmpty = errors.New("Queue: queue is empty")

type Queue[T any] struct {
	Data []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{Data: []T{}}
}

func (q *Queue[T]) Empty() bool {
	return q.Size() == 0
}

func (q *Queue[T]) Size() int {
	return len(q.Data)
}

func (q *Queue[T]) Front() T {
	if q.Empty() {
		panic(ErrQueueEmpty)
	}
	return q.Data[0]
}

func (q *Queue[T]) Push(x T) {
	q.Data = append(q.Data, x)
}

func (q *Queue[T]) PushRange(v []T) {
	for i := 0; i < len(v); i++ {
		q.Data = append(q.Data, v[i])
	}
}

func (q *Queue[T]) Pop() T {
	if q.Empty() {
		panic(ErrQueueEmpty)
	}
	res := q.Data[0]
	q.Data = q.Data[1:]
	return res
}
