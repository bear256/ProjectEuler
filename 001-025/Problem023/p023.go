package main

import (
	"fmt"
)

func sumDivisors(n int) int {
	sum := 1
	for i := 2; i*i <= n; i++ {
		if n%i==0 {
			j := n/i
			if i!= j {
				sum += i + n/i
			} else {
				sum += i
			}
			
		}
	}
	return sum
}

func main() {
	abundants := []int{}
	for n := 1; n <= 28123; n++ {
		if sumDivisors(n) > n {
			abundants = append(abundants, n)
		}
	}
	isAbundant := [28123*2+1]bool{false}
	for i := 0; i < len(abundants); i++ {
		for j := 0; j <= i; j++{
			idx := abundants[i] + abundants[j]
			isAbundant[idx] = true
		}
	}
	sum := 0
	for n := 1; n <= 28123; n++ {
		if !isAbundant[n] {
			sum += n
		}
	}
	fmt.Println(sum)
}
