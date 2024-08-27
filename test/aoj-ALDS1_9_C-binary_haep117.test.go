//go:build ignore

// verification-helper: PROBLEM https://onlinejudge.u-aizu.ac.jp/courses/lesson/1/ALDS1/9/ALDS1_9_C

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/today2098/algorithm-go/algorithm"
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

func getString() string {
	sc.Scan()
	return sc.Text()
}

func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	defer wr.Flush()

	pq := algorithm.NewDefaultBinaryHeap117()
	for {
		query := getString()
		if query == "end" {
			break
		}

		if query == "insert" {
			x := getInt()
			pq.Push(x)
		} else {
			out(pq.Pop())
		}
	}
}
