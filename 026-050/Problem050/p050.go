package main

import (
	"fmt"
	"math"
	"time"
)

func sqrt(n int) int {
	return int(math.Sqrt(float64(n)))
}

func genPrimes(n int) []int {
	primes := []int{2, 3}
	for p := primes[len(primes)-1] + 2; p <= n; p += 2 {
		flag := true
		for i := 0; i < len(primes); i++ {
			if p%primes[i] == 0 {
				flag = false
				break
			}
		}
		if flag {
			primes = append(primes, p)
		}
	}
	return primes
}

var pos map[int]int

func genPrimes2(n int) ([]int, []bool) {
	pbs := make([]bool, n)
	for i := 0; i < n; i++ {
		pbs[i] = true
	}
	count := n
	pbs[0] = false
	pbs[1] = false
	count -= 2
	for i := 2; i*i < n; i++ {
		if pbs[i] {
			for j := i * i; j < n; j += i {
				if pbs[j] {
					pbs[j] = false
					count--
				}
			}
		}
	}
	primes := make([]int, count)
	pos = make(map[int]int, count)
	idx := 0
	for p, pb := range pbs {
		if pb {
			primes[idx] = p
			pos[p] = idx
			idx++
		}
	}
	return primes, pbs
}

func sum(ps []int) int {
	sum := 0
	for _, p := range ps {
		sum += p
	}
	return sum
}

func order(n int, ns []int) int {
	pos := -1
	start, end := 0, len(ns)-1
	for {
		if start > end {
			// pos = start
			break
		}
		mid := start + (end-start)/2
		// fmt.Println(start, mid, end, pos)
		if n == ns[mid] {
			pos = mid
			break
		}
		if n < ns[mid] {
			end = mid - 1
		}
		if n > ns[mid] {
			start = mid + 1
		}
	}
	return pos
}

func order2(p int) int {
	return pos[p]
}

func calc(ps []int, pbs []bool) {
	prime := 0
	count := 0
	start, end := 0, 0
	for i := range ps {
		// for j := 0; j < order(p, ps); j++ { // A0
		// for j := 0; j < order2(p); j++ { // A1
		// for j := 0; j < i; j++ { // A2
		pbidx := 0 // B
		for j := i; j < len(ps); j++ {
			// pbidx := sum(ps[j:i]) // A
			// pbidx := sum(ps[i : j+1]) // B0
			pbidx += ps[j] // B1
			if pbidx >= len(pbs) {
				break
			}
			if pbs[pbidx] {
				// if count < i-j { // A
				if count < j-i {
					prime = pbidx
					// count = i - j // A
					count = j - i
					// start, end = j, i // A
					start, end = i, j+1
				}
			}
		}
	}
	fmt.Println(prime, count, start, end)
}

func calc2(ps []int, pbs []bool) {
	prime := 0
	count := 0
	start, end := 0, 0
	for i := range ps {
		sum := 0
		// length := 0
		for j := i; j < len(ps); j++ {
			sum += ps[j]
			// length++
			if sum >= len(pbs) {
				break
			}
			if pbs[sum] {
				// if count < length {
				if count < j-i {
					prime = sum
					// count = length
					count = j - i
					start, end = i, j+1
				}
			}
		}
	}
	fmt.Println(prime, count, start, end)
}

func timecount(f func([]int, []bool), ps []int, pbs []bool) {
	t0 := time.Now().UnixNano()
	f(ps, pbs)
	t1 := time.Now().UnixNano()
	fmt.Println((t1 - t0) / 1000)
}

func main() {
	n := 1000000
	t0 := time.Now().UnixNano()
	ps, pbs := genPrimes2(n)
	t1 := time.Now().UnixNano()
	fmt.Println((t1 - t0) / 1000)
	timecount(calc, ps, pbs)
	timecount(calc2, ps, pbs)
}
