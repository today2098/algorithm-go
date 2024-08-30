package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	"golang.org/x/exp/constraints"
)

const (
	inf  int     = 1e9
	linf int     = 1e18
	eps  float64 = 1e-10
	mod  int     = 998244353
	mod2 int     = 1e9 + 7
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func getInt() int {
	sc.Scan()
	elem, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return elem
}

func getFloat64() float64 {
	sc.Scan()
	elem, err := strconv.ParseFloat(sc.Text(), 64)
	if err != nil {
		panic(err)
	}
	return elem
}

func getString() string {
	sc.Scan()
	return sc.Text()
}

func getInts(n int) []int {
	v := make([]int, n)
	for i := 0; i < n; i++ {
		v[i] = getInt()
	}
	return v
}

func getStrings(n int) []string {
	vs := make([]string, n)
	for i := 0; i < n; i++ {
		vs[i] = getString()
	}
	return vs
}

func out(x ...any) {
	fmt.Fprintln(wr, x...)
}

func outArray[T any](arr []T) {
	for i := 0; i < len(arr)-1; i++ {
		fmt.Fprintf(wr, "%v ", arr[i])
	}
	if len(arr) > 0 {
		fmt.Fprintf(wr, "%v", arr[len(arr)-1])
	}
	fmt.Fprintf(wr, "\n")
}

func fill[T any](v []T, x T) {
	for i := 0; i < len(v); i++ {
		v[i] = x
	}
}

func min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func chmin[T constraints.Ordered](a *T, b T) bool {
	if *a > b {
		*a = b
		return true
	}
	return false
}

func chmax[T constraints.Ordered](a *T, b T) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	defer wr.Flush()

	out(inf, linf, eps, mod, mod2)

	n, m, s := getInt(), getFloat64(), getString()
	out(n, m, s)

	v, vs := getInts(n), getStrings(n)
	outArray(v)
	outArray(vs)

	fill(v, -1)
	outArray(v)

	out(min("hello", "world"), max("hello", "world"))

	a, b, c := 0, 1, -1
	chmin(&b, a)
	chmax(&c, a)
	out(a, b, c)
}
