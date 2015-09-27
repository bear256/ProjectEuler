package main

import (
	"fmt"
)

func isAbundant(n int) bool {
	prod := 1
	for k := 2; k*k <= n; k++ {
		p := 1
		for n%k == 0 {
			p = p*k + 1
			n /= k
		}
		prod *= p
	}
	if n > 1 {
		prod *= 1 + n
	}
	return prod > 2*n+1
}

func main() {
	sum := 0
	abundants := []int{}
	isAbundants := [28123*2 + 1]bool{false}
	fmt.Println(isAbundant(12), isAbundant(18))
	for n := 2; n <= 28123; n++ {
		isAbundants[n] = isAbundant(n)
		if isAbundants[n] {
			abundants = append(abundants, n)
		}
	}
	for i := 0; i < len(abundants); i++ {
		for j := 0; j <= i; j++ {
			idx := abundants[i] + abundants[j]
			isAbundants[idx] = true
		}
	}
	for n := 1; n <= 28123; n++ {
		if !isAbundants[n] {
			sum += n
		}
	}
	fmt.Println(sum)
}
