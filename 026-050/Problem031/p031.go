package main

import (
	"fmt"
)

var (
	ps = []int{1, 2, 5, 10, 20, 50, 100, 200}
	count = 0
	// ns = [8]int{0}
)

func recur(total, idx int) {
	if total == 0 {
		// fmt.Println(ns)
		count++
		return
	}
	if total < 0 || idx < 0 {
		return
	}
	for n := total/ps[idx]; n >= 0; n-- {
		// ns[idx] = n
		recur(total - n*ps[idx], idx-1)
	}
}

func main() {
	recur(200, 7)
	fmt.Println(count)
}