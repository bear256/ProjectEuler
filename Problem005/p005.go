package main

import (
	"fmt"
)

func solution1Search() {
	num := 1
	for ;; num++{
		flag := true
		for i:=2; i <= 20; i++ {
			if num % i != 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(num)
			break
		}
	}
}

func solution2Construct() {
	primes := []int{2,3,5,7,11,13,17,19}
	nums := []int{0,0,0,0,0,0,0,0}
	for i := 2; i <= 20; i++ {
		n := i
		for idx, prime := range primes {
			pn := 0
			for ;n > 1;{
				if n%prime == 0 {
					pn ++
					n = n/prime
				} else {
					break
				}
			}
			if pn > nums[idx] {
				nums[idx] = pn
			}
		}
	}
	num := 1
	for idx, prime := range primes {
		for j := 0; j < nums[idx]; j++ {
			num = num * prime
		}
	}
	fmt.Println(num)
}

func main() {
	solution1Search()
	solution2Construct()
}