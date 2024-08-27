//go:build go1.18

package algorithm

import (
	"errors"

	"golang.org/x/exp/constraints"
)

var ErrBinaryHeapEmpty = errors.New("BinaryHeap: binary-heap is empty")

type BinaryHeap[T any] struct {
	comp BinaryHeapCompFunc[T]
	tree []T
}

type BinaryHeapCompFunc[T any] func(a, b T) bool

func NewBinaryHeap[T any](f BinaryHeapCompFunc[T]) *BinaryHeap[T] {
	return &BinaryHeap[T]{
		comp: f,
		tree: make([]T, 1),
	}
}

func NewDefaultBinaryHeap[T constraints.Ordered]() *BinaryHeap[T] {
	return NewBinaryHeap(func(a, b T) bool {
		return a > b
	})
}

func (b *BinaryHeap[T]) shiftUp(i int) {
	p := i >> 1
	for 1 <= p {
		if b.comp(b.tree[p], b.tree[i]) {
			break
		}
		b.tree[p], b.tree[i] = b.tree[i], b.tree[p]
		i = p
		p >>= 1
	}
}

func (b *BinaryHeap[T]) shiftDown(i int) {
	l, r := i<<1, i<<1|1
	for l <= b.Size() {
		if b.Size() < r || b.comp(b.tree[l], b.tree[r]) {
			if b.comp(b.tree[i], b.tree[l]) {
				break
			}
			b.tree[i], b.tree[l] = b.tree[l], b.tree[i]
			i = l
		} else {
			if b.comp(b.tree[i], b.tree[r]) {
				break
			}
			b.tree[i], b.tree[r] = b.tree[r], b.tree[i]
			i = r
		}
		l, r = i<<1, i<<1|1
	}
}

func (b *BinaryHeap[T]) Empty() bool {
	return b.Size() == 0
}

func (b *BinaryHeap[T]) Size() int {
	return len(b.tree) - 1
}

func (b *BinaryHeap[T]) Top() T {
	if b.Empty() {
		panic(ErrBinaryHeapEmpty)
	}
	return b.tree[1]
}

func (b *BinaryHeap[T]) Push(x T) {
	b.tree = append(b.tree, x)
	b.shiftUp(b.Size())
}

func (b *BinaryHeap[T]) Pop() T {
	if b.Empty() {
		panic(ErrBinaryHeapEmpty)
	}
	res := b.tree[1]
	b.tree[1] = b.tree[b.Size()]
	b.tree = b.tree[:len(b.tree)-1]
	if !b.Empty() {
		b.shiftDown(1)
	}
	return res
}
