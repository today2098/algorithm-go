package algorithm

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func Accumulate[T any](v []T, init T, op func(acc, x T) T) T {
	for i := 0; i < len(v); i++ {
		init = op(init, v[i])
	}
	return init
}

func DefaultAccumulate[T constraints.Integer | constraints.Float](v []T) T {
	return Accumulate(v, 0, func(acc, x T) T {
		return acc + x
	})
}

func MinElement[T constraints.Ordered](v []T) int {
	if len(v) == 0 {
		return 0
	}
	itr := 0
	mn := v[0]
	for i := 1; i < len(v); i++ {
		if v[i] < mn {
			mn = v[i]
			itr = i
		}
	}
	return itr
}

func MaxElement[T constraints.Ordered](v []T) int {
	if len(v) == 0 {
		return 0
	}
	itr := 0
	mx := v[0]
	for i := 1; i < len(v); i++ {
		if v[i] > mx {
			mx = v[i]
			itr = i
		}
	}
	return itr
}

func LowerBound[T constraints.Ordered](v []T, x T) int {
	itr := sort.Search(len(v), func(i int) bool {
		return v[i] >= x
	})
	return itr
}

func UpperBound[T constraints.Ordered](v []T, x T) int {
	itr := sort.Search(len(v), func(i int) bool {
		return v[i] > x
	})
	return itr
}
