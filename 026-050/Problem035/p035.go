package main

import (
	"fmt"
	"math"
	"time"
)

// IsPrime ...
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

// Primes ...
func Primes(max int) []int {
	primes := make([]int, 0)
	primes = append(primes, 2)
	for n := 3; n <= max; n += 2 {
		flag := true
		for _, p := range primes {
			if p > int(math.Sqrt(float64(n))) {
				break
			}
			if n%p == 0 {
				flag = false
				break
			}
		}
		if flag {
			primes = append(primes, n)
		}
	}
	return primes
}

func test(n int) bool {
	ns := fmt.Sprint(n)
	flag := true
	for i := 0; i < len(ns); i++ {
		base := int(math.Pow10(i))
		base1 := int(math.Pow10(len(ns) - i))
		a := n / base
		b := n % base
		// fmt.Println(base, a, b, b*base1+a)
		flag = flag && IsPrime(b*base1+a)
	}
	return flag
}

func rotations(n int) []int {
	ns := fmt.Sprint(n)
	set := make([]int, len(ns))
	for i := 0; i < len(ns); i++ {
		base := int(math.Pow10(i))
		base1 := int(math.Pow10(len(ns) - i))
		a := n / base
		b := n % base
		set[i] = b*base1 + a
	}
	return set
}

// HasPrimes ...
func HasPrimes(max int) map[int]bool {
	primes := make([]int, 0)
	primes = append(primes, 2)
	hasPrime := make(map[int]bool)
	hasPrime[2] = true
	for n := 3; n <= max; n += 2 {
		flag := true
		for _, p := range primes {
			if p > int(math.Sqrt(float64(n))) {
				break
			}
			if n%p == 0 {
				flag = false
				break
			}
		}
		if flag {
			primes = append(primes, n)
			hasPrime[n] = true
		}
	}
	return hasPrime
}

func method1() {
	primes := Primes(1000000)
	// fmt.Println(primes)
	hasPrime := make(map[int]bool)
	for _, p := range primes {
		hasPrime[p] = true
	}
	n := 0
	for _, p := range primes {
		flag := true
		for _, sp := range rotations(p) {
			_, ok := hasPrime[sp]
			flag = flag && ok
		}
		if flag {
			n++
		}
	}
	fmt.Println(n)
}

func method2() {
	primes := Primes(1000000)
	// fmt.Println(primes)
	n := 0
	for _, p := range primes {
		if test(p) {
			// fmt.Println(p)
			n++
		}
	}
	fmt.Println(n)
}

func method3() {
	t0 := time.Now().UnixNano()
	hasPrime := HasPrimes(1000000)
	n := 0
	for p := range hasPrime {
		flag := true
		for _, sp := range rotations(p) {
			_, ok := hasPrime[sp]
			flag = flag && ok
		}
		if flag {
			n++
		}
	}
	fmt.Println(n)
	t1 := time.Now().UnixNano()
	fmt.Println(t1 - t0)
}

func main() {
	t0 := time.Now().UnixNano()
	method1()
	t1 := time.Now().UnixNano()
	fmt.Println(t1 - t0)
	t2 := time.Now().UnixNano()
	method2()
	t3 := time.Now().UnixNano()
	fmt.Println(t3 - t2)
	method3()
}
