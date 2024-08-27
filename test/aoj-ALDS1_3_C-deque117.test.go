//go:build ignore

// verification-helper: PROBLEM https://onlinejudge.u-aizu.ac.jp/courses/lesson/1/ALDS1/3/ALDS1_3_C
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

func outInts(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		fmt.Fprintf(wr, "%v ", arr[i])
	}
	if len(arr) > 0 {
		fmt.Fprintf(wr, "%v", arr[len(arr)-1])
	}
	fmt.Fprintf(wr, "\n")
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	defer wr.Flush()

	n := getInt()

	dq := algorithm.NewDeque117()
	for i := 0; i < n; i++ {
		cmd := getString()

		if cmd == "insert" {
			x := getInt()

			dq.PushFront(x)
		} else if cmd == "delete" {
			x := getInt()

			node := dq.Data.Front()
			for node != nil && node.Value.(int) != x {
				node = node.Next()
			}
			if node != nil {
				dq.Data.Remove(node)
			}
		} else if cmd == "deleteFirst" {
			dq.PopFront()
		} else {
			dq.PopBack()
		}
	}

	ans := []int{}
	node := dq.Data.Front()
	for node != nil {
		ans = append(ans, node.Value.(int))
		node = node.Next()
	}

	outInts(ans)
}
