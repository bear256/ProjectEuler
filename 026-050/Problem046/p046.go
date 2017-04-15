package main

import (
	"fmt"
	"math"
)

func sqrt(n int) int {
	return int(math.Sqrt(float64(n)))
}

// IsPrime ..
func IsPrime(n int) bool {
	flag := n > 1
	if n > 2 {
		flag = n%2 != 0
	}
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			flag = false
			break
		}
	}
	return flag
}

func test(odd int) bool {
	flag := false
	if !IsPrime(odd) {
		for n := sqrt((odd - 1) / 2); n > 0; n-- {
			p := odd - 2*n*n
			flag = flag || IsPrime(p)
		}
	} else {
		flag = true
	}
	return flag
}

func main() {
	for odd := 9; ; odd += 2 {
		if !test(odd) {
			fmt.Println(odd)
			break
		}
	}
}
