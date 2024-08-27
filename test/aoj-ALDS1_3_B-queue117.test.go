//go:build ignore

// verification-helper: PROBLEM https://onlinejudge.u-aizu.ac.jp/courses/lesson/1/ALDS1/3/ALDS1_3_B

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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	defer wr.Flush()

	n, q := getInt(), getInt()

	type task struct {
		name string
		time int
	}
	que := algorithm.NewQueue117()
	for i := 0; i < n; i++ {
		t := &task{}
		t.name, t.time = getString(), getInt()
		que.Push(t)
	}

	now := 0
	for !que.Empty() {
		t := que.Pop().(*task)
		if t.time <= q {
			now += t.time
			out(t.name, now)
		} else {
			now += q
			t.time -= q
			que.Push(t)
		}
	}
}
