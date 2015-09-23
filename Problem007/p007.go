package main

import (
	"fmt"
	"math"
)

func main() {
	count := 1
	primes := []int{2}
	for i:=3; ;i+=2{
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
			count++
			if count == 10001 {
				fmt.Println(i)
				break
			}
		}
	}
}