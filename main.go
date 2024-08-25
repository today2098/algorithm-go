package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"

	"golang.org/x/exp/constraints"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

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

func min[T constraints.Ordered](a T, b T) T {
	if a < b {
		return a
	}
	return b
}

func max[T constraints.Ordered](a T, b T) T {
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

func abs[T constraints.Integer | constraints.Float](a T) T {
	if a >= 0 {
		return a
	}
	return -a
}

func minElement[T constraints.Ordered](v []T) int {
	itr := 0
	if len(v) == 0 {
		return itr
	}
	mn := v[0]
	for i := 1; i < len(v); i++ {
		if chmin(&mn, v[i]) {
			itr = i
		}
	}
	return itr
}

func maxElement[T constraints.Ordered](v []T) int {
	itr := 0
	if len(v) == 0 {
		return itr
	}
	mx := v[0]
	for i := 1; i < len(v); i++ {
		if chmax(&mx, v[i]) {
			itr = i
		}
	}
	return itr
}

func lowerBound[T constraints.Ordered](v []T, x T) int {
	itr := sort.Search(len(v), func(i int) bool {
		return v[i] >= x
	})
	return itr
}

func upperBound[T constraints.Ordered](v []T, x T) int {
	itr := sort.Search(len(v), func(i int) bool {
		return v[i] > x
	})
	return itr
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	defer wr.Flush()

	n, m, s := getInt(), getFloat64(), getString()
	out(n, m, s)

	v, vs := getInts(n), getStrings(n)
	outArray(v)
	outArray(vs)

	fill(v, -1)
	outArray(v)

	out(min("hello", "world"), max("hello", "world"))
	out(abs(-10), abs(10))

	a, b, c := 0, 1, -1
	chmin(&b, a)
	chmax(&c, a)
	out(a, b, c)

	w := []int{3, 1, 4, 2, 5}
	out(w, minElement(w), maxElement(w))

	sort.Ints(w)
	out(w, lowerBound(w, 3), upperBound(w, 3))
}
