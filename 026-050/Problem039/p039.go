package main

import (
	"fmt"
)

func test(p int) int {
	n := 0
	for a := 1; a < p; a++ {
		for b := a; b < p-a; b++ {
			c := p - a - b
			if a*a+b*b == c*c {
				// fmt.Println(a, b, c)
				n++
			}
		}
	}
	return n
}

func main() {
	max := 0
	maxP := 0
	for p := 12; p <= 1000; p++ {
		n := test(p)
		if max < n {
			max = n
			maxP = p
		}
	}
	fmt.Println(max, maxP)
}
