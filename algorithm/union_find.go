package algorithm

import "errors"

var ErrUnionFindOutOfRange = errors.New("UnionFind: index out of range")

type UnionFind struct {
	par []int
	gn  int
}

func NewUnionFind(n int) *UnionFind {
	par := make([]int, n)
	for i := 0; i < n; i++ {
		par[i] = -1
	}
	return &UnionFind{
		par: par,
		gn:  n,
	}
}

func (uf *UnionFind) Vn() int {
	return len(uf.par)
}

func (uf *UnionFind) Gn() int {
	return uf.gn
}

func (uf *UnionFind) Root(x int) int {
	if !(0 <= x && x < uf.Vn()) {
		panic(ErrUnionFindOutOfRange)
	}
	if uf.par[x] < 0 {
		return x
	}
	return uf.Root(uf.par[x])
}

func (uf *UnionFind) Size(x int) int {
	if !(0 <= x && x < uf.Vn()) {
		panic(ErrUnionFindOutOfRange)
	}
	return -uf.par[uf.Root(x)]
}

func (uf *UnionFind) IsSame(x, y int) bool {
	if !(0 <= x && x < uf.Vn() && 0 <= y && y < uf.Vn()) {
		panic(ErrUnionFindOutOfRange)
	}
	return uf.Root(x) == uf.Root(y)
}

func (uf *UnionFind) Unite(x, y int) bool {
	if !(0 <= x && x < uf.Vn() && 0 <= y && y < uf.Vn()) {
		panic(ErrUnionFindOutOfRange)
	}
	x, y = uf.Root(x), uf.Root(y)
	if x == y {
		return false
	}
	if uf.Size(x) < uf.Size(y) {
		x, y = y, x
	}
	uf.par[x] += uf.par[y]
	uf.par[y] = x
	uf.gn--
	return true
}

func (uf *UnionFind) Reset() {
	for i := 0; i < uf.Vn(); i++ {
		uf.par[i] = -1
	}
	uf.gn = uf.Vn()
}
