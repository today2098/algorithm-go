//go:build ignore

// verification-helper: PROBLEM https://onlinejudge.u-aizu.ac.jp/courses/library/3/DSL/1/DSL_1_A

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/today2098/algorithm-go/algorithm"
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

func out(x ...any) {
	fmt.Fprintln(wr, x...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	defer wr.Flush()

	n, q := getInt(), getInt()

	uf := algorithm.NewUnionFind(n)
	for i := 0; i < q; i++ {
		com, x, y := getInt(), getInt(), getInt()

		if com == 0 {
			uf.Unite(x, y)
		} else {
			if uf.IsSame(x, y) {
				out(1)
			} else {
				out(0)
			}
		}
	}
}
