package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	var a []int

	readInput(&n, &a)
	process(n, &a)
}

func process(n int, a *[]int) {
	sort.Ints(*a)
	fmt.Print(a)

	var res int = len(*a)
	var l int = 1
	var r int = len(*a) - 1

	for _, v := range *a {
		var curRes int
		l, curRes = binarySearch(a, l, r, v)
		if curRes == -1 {
			break
		}
		res--
	}

	fmt.Print(res)
}

func binarySearch(a *[]int, l, r, x int) (int, int) {
	if l > r {
		return l, -1
	}

	var res int = -1

	for l <= r {
		var mid int = (l + r) / 2
		if (*a)[mid] >= x*2 {
			res = mid
			l = mid + 1
		}
		r = mid - 1
	}

	l = res + 1
	return l, res
}

func readInput(n *int, a *[]int) {
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		var x int
		fmt.Scan(&x)
		*a = append(*a, x)
	}
}
