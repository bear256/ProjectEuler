package main

import (
	"fmt"
)

func factorial(n int) int {
	prod := 1
	for i := 2; i <= n; i++ {
		prod *= i
	}
	return prod
}

func recur(nth, level int, digits []int) {
	sub := factorial(level)
	rank := 0
	for ;nth - sub > 0;rank++ {
		nth -= sub
	}
	fmt.Print(digits[rank])
	tmp := []int{}
	for _, digit := range digits {
		if digit != digits[rank] {
			tmp = append(tmp, digit)
		}
	}
	if level > 1 {
		recur(nth, level - 1, tmp)
	} else {
		fmt.Println(tmp[0])
	}
}

func main() {
	nth := 1000000
	recur(nth, 9, []int{0,1,2,3,4,5,6,7,8,9})
}