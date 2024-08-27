//go:build go1.18

package algorithm

import (
	"errors"
	"math"

	"golang.org/x/exp/constraints"
)

var ErrBellmanFordOutOfRange = errors.New("BellmanFord: index out of range")

type BellmanFord[T constraints.Integer | constraints.Float] struct {
	vn  int
	g   []*bellmanFordEdge[T]
	d   []T
	pre []int
	inf T
}

type bellmanFordEdge[T constraints.Integer | constraints.Float] struct {
	from int
	to   int
	cost T
}

func NewBellmanFord[T constraints.Integer | constraints.Float](n int, inf T) *BellmanFord[T] {
	d := make([]T, n)
	pre := make([]int, n)
	for i := 0; i < n; i++ {
		d[i] = inf
		pre[i] = -1
	}
	return &BellmanFord[T]{
		vn:  n,
		g:   []*bellmanFordEdge[T]{},
		d:   d,
		pre: pre,
		inf: inf,
	}
}

func NewDefaultBellmanFord(n int) *BellmanFord[int] {
	return NewBellmanFord(n, math.MaxInt)
}

func (b *BellmanFord[T]) Infinity() T {
	return b.inf
}

func (b *BellmanFord[T]) Order() int {
	return b.vn
}

func (b *BellmanFord[T]) AddEdge(from, to int, cost T) {
	if !(0 <= from && from < b.Order() && 0 <= to && to < b.Order()) {
		panic(ErrBellmanFordOutOfRange)
	}
	b.g = append(b.g, &bellmanFordEdge[T]{
		from: from,
		to:   to,
		cost: cost,
	})
}

func (b *BellmanFord[T]) FindNegativeCycle() bool {
	nd := make([]T, b.Order())
	for i := 0; i < b.Order(); i++ {
		update := false
		for _, edge := range b.g {
			if nd[edge.to] > nd[edge.from]+edge.cost {
				nd[edge.to] = nd[edge.from] + edge.cost
				update = true
			}
		}
		if !update {
			return false // non negative cycle
		}
	}
	return true // find negative cycle
}

func (b *BellmanFord[T]) BellmanFord(s int) bool {
	if !(0 <= s && s < b.Order()) {
		panic(ErrBellmanFordOutOfRange)
	}
	for i := 0; i < b.Order(); i++ {
		b.d[i] = b.Infinity()
		b.pre[i] = -1
	}
	b.d[s] = 0
	for i := 0; i < b.Order(); i++ {
		update := false
		for _, edge := range b.g {
			if b.d[edge.from] == b.Infinity() {
				continue
			}
			if b.d[edge.to] > b.d[edge.from]+edge.cost {
				b.d[edge.to] = b.d[edge.from] + edge.cost
				b.pre[edge.to] = edge.from
				update = true
			}
		}
		if !update {
			return false // non negative cycle
		}
	}
	for i := 0; i < b.Order(); i++ {
		update := false
		for _, edge := range b.g {
			if b.d[edge.from] == b.Infinity() || b.d[edge.to] == -b.Infinity() {
				continue
			}
			if b.d[edge.from] == -b.Infinity() || b.d[edge.to] > b.d[edge.from]+edge.cost {
				b.d[edge.to] = -b.Infinity()
				update = true
			}
		}
		if !update {
			break
		}
	}
	return true // find negative cycle
}

func (b *BellmanFord[T]) Distance(t int) T {
	if !(0 <= t && t < b.Order()) {
		panic(ErrBellmanFordOutOfRange)
	}
	return b.d[t]
}

func (b *BellmanFord[T]) ShortestPath(t int) []int {
	if !(0 <= t && t < b.Order()) {
		panic(ErrBellmanFordOutOfRange)
	}
	var path []int
	if b.Distance(t) == b.Infinity() || b.Distance(t) == -b.Infinity() {
		return path
	}
	for t != -1 {
		path = append(path, t)
		t = b.pre[t]
	}
	for i := 0; i < len(path)/2; i++ {
		path[i], path[len(path)-1-i] = path[len(path)-1-i], path[i]
	}
	return path
}
