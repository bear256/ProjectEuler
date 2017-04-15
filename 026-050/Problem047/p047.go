package main

import (
	"fmt"
	"math"
)

var primes []int = []int{2, 3, 5, 7}

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

func newPrime(n int) int {
	for p := primes[n-1] + 2; ; p += 2 {
		if IsPrime(p) {
			return p
		}
	}
}

func isPrimeProd(prod, n int) bool {
	i := 0
	tp := 0
	pn := 0
	for prod != 1 {
		// fmt.Println(i, prod, primes[i], tp, pn)
		if prod%primes[i] == 0 {
			if tp != primes[i] {
				tp = primes[i]
				pn++
			}
			prod = prod / primes[i]
		} else {
			i++
			// if i == len(primes) {
			// 	primes = append(primes, newPrime(len(primes)))
			// }
			if pn > n {
				break
			}
		}
	}
	// fmt.Println(pn)
	return pn == n
}

func consecutive(m, n int) bool {
	flag := true
	for i := 0; i < n; i++ {
		prod := m + i
		flag = flag && isPrimeProd(prod, n)
	}
	return flag
}

func main() {
	n := 4
	for i := 4; ; i++ {
		if i == len(primes) {
			primes = append(primes, newPrime(len(primes)))
		}
		if primes[i]-primes[i-1] > n {
			for m := primes[i-1] + 1; m < primes[i]-n+1; m++ {
				if consecutive(m, n) {
					fmt.Println(m)
					return
				}
			}
		}
	}
}
