package algorithm

import "errors"

var ErrBinaryHeap117Empty = errors.New("BinaryHeap117: binary-heap is empty")

// A binary-heap data structure.
type BinaryHeap117 struct {
	cmp  BinaryHeap117CmpFunc
	tree []interface{}
}

type BinaryHeap117CmpFunc func(a, b interface{}) bool

// Create a new binary-heap.
func NewBinaryHeap117(f BinaryHeap117CmpFunc) *BinaryHeap117 {
	return &BinaryHeap117{
		cmp:  f,
		tree: make([]interface{}, 1),
	}
}

// Create a new basic binary-heap.
func NewDefaultBinaryHeap117() *BinaryHeap117 {
	return NewBinaryHeap117(func(a, b interface{}) bool {
		return a.(int) > b.(int)
	})
}

func (b *BinaryHeap117) shiftUp(i int) {
	p := i >> 1
	for 1 <= p {
		if b.cmp(b.tree[p], b.tree[i]) {
			break
		}
		b.tree[p], b.tree[i] = b.tree[i], b.tree[p]
		i = p
		p >>= 1
	}
}

func (b *BinaryHeap117) shiftDown(i int) {
	l, r := i<<1, i<<1|1
	for l <= b.Size() {
		if b.Size() < r || b.cmp(b.tree[l], b.tree[r]) {
			if b.cmp(b.tree[i], b.tree[l]) {
				break
			}
			b.tree[i], b.tree[l] = b.tree[l], b.tree[i]
			i = l
		} else {
			if b.cmp(b.tree[i], b.tree[r]) {
				break
			}
			b.tree[i], b.tree[r] = b.tree[r], b.tree[i]
			i = r
		}
		l, r = i<<1, i<<1|1
	}
}

// Checks if the binary-heap is empty.
func (b *BinaryHeap117) Empty() bool {
	return b.Size() == 0
}

// Returns the number of elements.
func (b *BinaryHeap117) Size() int {
	return len(b.tree) - 1
}

// Returns the largest element.
func (b *BinaryHeap117) Top() interface{} {
	if b.Empty() {
		panic(ErrBinaryHeap117Empty)
	}
	return b.tree[1]
}

// Inserts an element to the binary-heap.
func (b *BinaryHeap117) Push(x interface{}) {
	b.tree = append(b.tree, x)
	b.shiftUp(b.Size())
}

// Removes and returns the largest element.
func (b *BinaryHeap117) Pop() interface{} {
	if b.Empty() {
		panic(ErrBinaryHeap117Empty)
	}
	res := b.tree[1]
	b.tree[1] = b.tree[b.Size()]
	b.tree = b.tree[:len(b.tree)-1]
	if !b.Empty() {
		b.shiftDown(1)
	}
	return res
}
