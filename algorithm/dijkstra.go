//go:build go1.18

package algorithm

import (
	"errors"
	"math"

	"golang.org/x/exp/constraints"
)

var ErrDijkstraOutOfRange = errors.New("Dijkstra: index out of range")

type Dijkstra[T constraints.Integer | constraints.Float] struct {
	g   [][]*dijkstraPair[T] // pair of (cost, to)
	d   []T
	pre []int
	inf T
}

type dijkstraPair[T any] struct {
	first  T
	second int
}

func NewDijkstra[T constraints.Integer | constraints.Float](n int, inf T) *Dijkstra[T] {
	d := make([]T, n)
	pre := make([]int, n)
	for i := 0; i < n; i++ {
		d[i] = inf
		pre[i] = -1
	}
	return &Dijkstra[T]{
		g:   make([][]*dijkstraPair[T], n),
		d:   d,
		pre: pre,
		inf: inf,
	}
}

func NewDefaultDijkstra(n int) *Dijkstra[int] {
	return NewDijkstra(n, math.MaxInt)
}

func (d *Dijkstra[T]) Infinity() T {
	return d.inf
}

func (d *Dijkstra[T]) Order() int {
	return len(d.g)
}

func (d *Dijkstra[T]) AddEdge(from, to int, cost T) {
	if !(0 <= from && from < d.Order() && 0 <= to && to < d.Order()) {
		panic(ErrDijkstraOutOfRange)
	}
	d.g[from] = append(d.g[from], &dijkstraPair[T]{
		first:  cost,
		second: to,
	})
}

func (d *Dijkstra[T]) Dijkstra(s int) {
	if !(0 <= s && s < d.Order()) {
		panic(ErrDijkstraOutOfRange)
	}
	for i := 0; i < d.Order(); i++ {
		d.d[i] = d.Infinity()
		d.pre[i] = -1
	}
	d.d[s] = 0
	pq := NewBinaryHeap(func(a, b *dijkstraPair[T]) bool {
		return a.first < b.first
	})
	pq.Push(&dijkstraPair[T]{
		first:  0, // dist
		second: s, // node
	})
	for !pq.Empty() {
		src := pq.Pop()
		if d.d[src.second] < src.first {
			continue
		}
		for _, edge := range d.g[src.second] {
			if d.d[edge.second] > d.d[src.second]+edge.first {
				d.d[edge.second] = d.d[src.second] + edge.first
				d.pre[edge.second] = src.second
				pq.Push(&dijkstraPair[T]{
					first:  d.d[edge.second],
					second: edge.second,
				})
			}
		}
	}
}

func (d *Dijkstra[T]) Distance(t int) T {
	if !(0 <= t && t < d.Order()) {
		panic(ErrDijkstraOutOfRange)
	}
	return d.d[t]
}

func (d *Dijkstra[T]) ShortestPath(t int) []int {
	if !(0 <= t && t < d.Order()) {
		panic(ErrDijkstraOutOfRange)
	}
	var path []int
	if d.Distance(t) == d.Infinity() {
		return path
	}
	for t != -1 {
		path = append(path, t)
		t = d.pre[t]
	}
	for i := 0; i < len(path)/2; i++ {
		path[i], path[len(path)-1-i] = path[len(path)-1-i], path[i]
	}
	return path
}
