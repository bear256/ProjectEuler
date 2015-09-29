package main

import (
	"fmt"
)

func formula(a, b, n int) int {
	return n*n + a*n + b
}

func isPrime(n int) bool {
	flag := true
	for i := 2; i*i <= n; i++ {
		if n % i == 0 {
			flag = false
			break
		}
	}
	return flag
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

func main() {
	max := 0
	prod := -1
	for a := -999; a < 1000; a++ {
		for b := -999; b < 1000; b++ {
			count := 0
			for n := 0; ; n++ {
				p := formula(a, b, n)
				if p <= 0 || !isPrime(p) { 
					if max < count {
						max = count
						prod = a * b
					}
					break
				}
				count++
			}
		}
	}
	fmt.Println(prod)
}