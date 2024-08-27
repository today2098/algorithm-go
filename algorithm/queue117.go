package algorithm

import "errors"

var ErrQueue117Empty = errors.New("Queue117: queue is empty")

// A queue structure.
type Queue117 struct {
	Data []interface{}
}

// Create a new queue.
func NewQueue117() *Queue117 {
	return &Queue117{
		Data: []interface{}{},
	}
}

// Checks if the queue is empty.
func (q *Queue117) Empty() bool {
	return q.Size() == 0
}

// Returns the number of elements.
func (q *Queue117) Size() int {
	return len(q.Data)
}

// Returns the front element.
func (q *Queue117) Front() interface{} {
	if q.Empty() {
		panic(ErrQueue117Empty)
	}
	return q.Data[0]
}

// Inserts an element at the back.
func (q *Queue117) Push(x interface{}) {
	q.Data = append(q.Data, x)
}

// Inserts an elements at the back.
func (q *Queue117) PushRange(v []interface{}) {
	for i := 0; i < len(v); i++ {
		q.Data = append(q.Data, v[i])
	}
}

// Removes and returns the front element.
func (q *Queue117) Pop() interface{} {
	if q.Empty() {
		panic(ErrQueue117Empty)
	}
	res := q.Data[0]
	q.Data = q.Data[1:]
	return res
}
