package queue

import "errors"

var ErrQueueEmpty = errors.New("Queue: queue is empty")

type Queue[T any] struct {
	data []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{data: []T{}}
}

func (q *Queue[T]) Empty() bool {
	return q.Size() == 0
}

func (q *Queue[T]) Size() int {
	return len(q.data)
}

func (q *Queue[T]) Front() T {
	if q.Empty() {
		panic(ErrQueueEmpty)
	}
	return q.data[0]
}

func (q *Queue[T]) Push(x T) {
	q.data = append(q.data, x)
}

func (q *Queue[T]) PushRange(v []T) {
	for i := 0; i < len(v); i++ {
		q.data = append(q.data, v[i])
	}
}

func (q *Queue[T]) Pop() T {
	if q.Empty() {
		panic(ErrQueueEmpty)
	}
	res := q.data[0]
	q.data = q.data[1:]
	return res
}
