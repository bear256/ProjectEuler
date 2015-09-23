package main

import (
	"math"
	"fmt"
)

func main() {
	num := 600851475143
	primes := []int{}
	for i:=2; ;i++{
		isPrime := true
		for _, prime := range primes {
			if i%prime == 0{
				isPrime = false
				break
			}
			if prime > int(math.Sqrt(float64(i))) {
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
			prime := i
			for ;num%prime == 0 && num > 1; {
//				fmt.Println(prime)
				num = num/prime
			}
			if num == 1 {
				fmt.Println(prime)
				break
			}
		}
	}
}