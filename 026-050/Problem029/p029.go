package main

import (
	"fmt"
)

func factor(n int) (int, int) {
	var base, pow, prod int
	for base = 2; base <= n; base++ {
		prod = n
		pow = 0
		flag := true
		for ;prod > 1; {
			if prod % base == 0 {
				pow ++
				prod /= base
			} else {
				flag = false
				break
			}
		}
		if flag {
			break
		}
	}
	// fmt.Println(base, pow)
	return base, pow
}

func main() {
	m := map[int]int{}
	for a := 2; a <= 100; a++ {
		for b:=2; b <=100; b++ {
			base, pow := factor(a)
			pow *= b
			m[base*1000+pow] = 0
		}
	}
	fmt.Println(len(m))
}