//go:build ignore

// verification-helper: PROBLEM https://onlinejudge.u-aizu.ac.jp/courses/lesson/1/ALDS1/3/ALDS1_3_A

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/today2098/algorithm-go/algorithm/stack"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func out(x ...any) {
	fmt.Fprintln(wr, x...)
}

func main() {
	sc.Split(bufio.ScanLines)
	sc.Buffer([]byte{}, math.MaxInt32)
	defer wr.Flush()

	sc.Scan()
	query := strings.Split(sc.Text(), " ")

	st := stack.NewStack[int]()
	for _, elem := range query {
		num, err := strconv.Atoi(elem)
		if err == nil {
			st.Push(num)
		} else {
			if elem == "+" {
				tmp := st.Pop() + st.Pop()
				st.Push(tmp)
			} else if elem == "-" {
				tmp := -st.Pop() + st.Pop()
				st.Push(tmp)
			} else {
				tmp := st.Pop() * st.Pop()
				st.Push(tmp)
			}
		}
	}

	ans := st.Pop()
	out(ans)
}
