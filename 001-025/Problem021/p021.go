package main

import (
	"fmt"
)

func d(n int) int {
	sum := 1
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	sum := 0
	for a := 2; a < 10000; a++ {
		b := d(a)
		if b != a && d(b) == a {
			fmt.Println(a)
			sum += a
		}
	}
	fmt.Println(sum)
}
