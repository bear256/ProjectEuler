package main

import (
	"fmt"
)

func quare(n []int) []int {
	var carry int = 0
	for i := 0; i < len(n); i++ {
		n[i] = n[i] * 2 + carry
		carry = 0
		if n[i] >= 10 {
			carry = n[i]/10
			n[i] = n[i]%10
		}
	}
	if carry > 0 {
		n = append(n, carry)
	}
	return n
}

func sum(n []int) int {
	s := 0
	for _, b := range n {
		s += b
	}
	return s
}

func main() {
	n := 1000
	nums := []int{1}
	for i := 0; i < n; i++ {
		nums = quare(nums)
	}
	fmt.Println(nums)
	fmt.Println(sum(nums))
}