package main

import (
	"fmt"
)

// GCD ...
func GCD(a, b int) int {
	// a = a % b
	// if a == 0 {
	// 	return b
	// } else {
	// 	return GCD(b, a)
	// }
	for b != 0 {
		a = a % b
		a, b = b, a
	}
	return a
}

func main() {
	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	numerator, denominator := 1, 1
	for _, a := range number {
		for _, b := range number {
			for _, c := range number[1:] {
				// 10*a+b : 10*+c == a : c && a:c < 1
				if (10*a+b)*c == a*(10*b+c) && a*1.0/c < 1 {
					fmt.Println(10*a+b, "/", 10*b+c)
					numerator *= a
					denominator *= c
				}
			}
		}
	}
	gcd := GCD(numerator, denominator)
	fmt.Println(denominator / gcd)
}
