package main

import (
	"fmt"
)

func plus(a, b []int) []int {
	if len(a) < len(b) {
		a, b = b, a
	}
	c := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		if i < len(b) {
			c[i] = a[i] + b[i]
		} else {
			c[i] = a[i]
		}
	}
	for i := 0; i < len(c); i++ {
		if c[i] >= 10 {
			if i+1 < len(c) {
				c[i+1] += c[i]/10
			} else {
				c = append(c, c[i]/10)
			}
			c[i] = c[i]%10
		}
	}
	return c
}

func main() {
	a, b := []int{1}, []int{1}
	for nth := 3;; nth++{
		c := plus(a, b)
		if len(c) == 1000 {
			fmt.Println(nth)
			break
		}
		a = b
		b = c
	}
}