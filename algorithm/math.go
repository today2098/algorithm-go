package algorithm

import "golang.org/x/exp/constraints"

func Abs[T constraints.Integer | constraints.Float](a T) T {
	if a >= 0 {
		return a
	}
	return -a
}

func ModPow(n, k, m int) int {
	if k == 0 {
		return 1
	}
	res := ModPow(n, k/2, m)
	res = res * res % m
	if k&1 == 1 {
		res = res * n % m
	}
	return res
}
