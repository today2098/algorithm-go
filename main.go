package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	defer wr.Flush()

	n, m, s := getInt(), getFloat64(), getString()
	out(n, m, s)

	v := getInts(n)
	outArray(v)
}
