package main

import (
	"fmt"
)

func nthpow(n, pow int) int {
	nt := n
	sum := 0
	for ;nt > 0; {
		u := nt%10
		np := 1
		for i := 0; i < pow; i++ {
			np *= u
		}
		sum += np
		nt /= 10
	}
	return sum
}

func main() {
	sum := 0
	for n := 10; n < 300000; n++ {
		if n == nthpow(n, 5) {
			sum += n
		}
	}
	fmt.Println(sum)
}