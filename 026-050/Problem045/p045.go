package main

import (
	"fmt"
	"math"
)

func tn(n int) int {
	return n * (n + 1) / 2
}

func pn(n int) int {
	return n * (3*n - 1) / 2
}

func hn(n int) int {
	return n * (2*n - 1)
}

func sqrt(n int) int {
	return int(math.Sqrt(float64(n)))
}

func isTn(m int) bool {
	n := (sqrt(1+8*m) - 1) / 2
	return tn(n) == m
}

func isPn(m int) bool {
	n := (sqrt(1+24*m) + 1) / 6
	return pn(n) == m
}

func main() {
	for n := 144; ; n++ {
		m := hn(n)
		if isPn(m) && isTn(m) {
			fmt.Println(m)
			break
		}
	}
}
