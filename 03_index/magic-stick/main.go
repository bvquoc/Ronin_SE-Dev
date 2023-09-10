package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	var a []int
	var b []int

	readInput(&n, &a, &b)
	solve(&n, &a, &b)
}

func readInput(n *int, a, b *[]int) {
	fmt.Scan(n)

	for i := 0; i < *n; i++ {
		var x int
		fmt.Scan(&x)
		*a = append(*a, x)
	}

	for i := 0; i < *n; i++ {
		var x int
		fmt.Scan(&x)
		*b = append(*b, x)
	}
}

func solve(n *int, a, b *[]int) {
	const YES = "YES"
	const NO = "NO"

	sort.Ints(*a)
	sort.Ints(*b)

	var res string = YES
	for i, v := range *a {
		if v > (*b)[i] {
			res = NO
			break
		}
	}

	fmt.Print(res)
}
