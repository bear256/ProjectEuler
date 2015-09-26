package main

import (
	"fmt"
)

const BASE = 1000

func product(a []int, b int) []int {
	carry := 0
	for i := 0; i < len(a); i++ {
		tmp := a[i]*b + carry
		if tmp > BASE {
			carry = tmp / BASE
			a[i] = tmp % BASE
		} else {
			carry = 0
			a[i] = tmp
		}
	}
	if carry > 0 {
		for ;carry > BASE; {
			a = append(a, carry%BASE)
			carry = carry/BASE
		}
		a = append(a, carry)
	}
	return a
}

func calcSum(ns []int) int {
	sum := 0
	for _, n := range ns {
		sum += n/100+n%100/10+n%10
	}
	return sum
}

func main() {
	prod := []int{1}
	for n := 2; n <= 100; n++ {
		prod = product(prod, n)
	}
	fmt.Println(prod)
	fmt.Println(calcSum(prod))
}
