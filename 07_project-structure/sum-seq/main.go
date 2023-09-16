package main

import (
	"fmt"
)

func main() {
	var n int
	var a []int

	readInput(&n, &a)
	calcSumSeq(a)
}

func calcSumSeq(a []int) {
	n := len(a)
	if n <= 0 {
		return
	}
	if n == 1 {
		fmt.Print(a)
		return
	}

	fmt.Println(a)
	b := make([]int, 0)
	for i := 0; i < n-1; i++ {
		b = append(b, a[i]+a[i+1])
	}
	calcSumSeq(b)
}

func readInput(n *int, a *[]int) {
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		var x int
		fmt.Scan(&x)
		*a = append(*a, x)
	}
}
