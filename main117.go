// main117.go
//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
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

func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}

func outInts(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		fmt.Fprintf(wr, "%v ", arr[i])
	}
	if len(arr) > 0 {
		fmt.Fprintf(wr, "%v", arr[len(arr)-1])
	}
	fmt.Fprintf(wr, "\n")
}

func fillInts(v []int, x int) {
	for i := 0; i < len(v); i++ {
		v[i] = x
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minElement(v []int) int {
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

func maxElement(v []int) int {
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

func lowerBound(v []int, x int) int {
	itr := sort.Search(len(v), func(i int) bool {
		return v[i] >= x
	})
	return itr
}

func upperBound(v []int, x int) int {
	itr := sort.Search(len(v), func(i int) bool {
		return v[i] > x
	})
	return itr
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	defer wr.Flush()

	out(inf, linf, eps, mod, mod2)

	n, m, s := getInt(), getFloat64(), getString()
	out(n, m, s)

	v, vs := getInts(3), getStrings(3)
	out(v, vs)

	w := make([]int, 5)
	fillInts(w, -1)
	outInts(w)

	vv := []int{3, 1, 4, 2, 5}
	out(vv, minElement(vv), maxElement(vv))

	sort.Ints(vv)
	out(vv, lowerBound(vv, 3), upperBound(vv, 3))
}
