package algorithm

import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Chmin[T constraints.Ordered](a *T, b T) bool {
	if *a > b {
		*a = b
		return true
	}
	return false
}

func Chmax[T constraints.Ordered](a *T, b T) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}
