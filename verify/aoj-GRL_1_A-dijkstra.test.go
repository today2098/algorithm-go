//go:build ignore

// verification-helper: PROBLEM https://onlinejudge.u-aizu.ac.jp/courses/library/5/GRL/1/GRL_1_A

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

	n, m, r := getInt(), getInt(), getInt()

	dijkstra := algorithm.NewDefaultDijkstra(n)
	for i := 0; i < m; i++ {
		s, t, d := getInt(), getInt(), getInt()
		dijkstra.AddEdge(s, t, d)
	}
	dijkstra.Dijkstra(r)

	for i := 0; i < n; i++ {
		ans := dijkstra.Distance(i)
		if ans == dijkstra.Infinity() {
			out("INF")
		} else {
			out(ans)
		}
	}
}
