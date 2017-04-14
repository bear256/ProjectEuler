package main

import (
	"fmt"
)

func factorial(n int) int {
	result := 1
	for ; n != 0 && n > 1; n-- {
		result *= n
	}
	return result
}

func factMap() map[string]int {
	fm := make(map[string]int, 10)
	for n := 0; n < 10; n++ {
		ns := fmt.Sprint(n)
		fm[ns] = factorial(n)
	}
	return fm
}

func test(n int, fm map[string]int) bool {
	sum := 0
	for _, ns := range fmt.Sprint(n) {
		sum += fm[string(ns)]
	}
	return n == sum
}

func main() {
	fm := factMap()
	sum := 0
	for n := 3; n < fm["9"]; n++ {
		if test(n, fm) {
			fmt.Println(n)
			sum += n
		}
	}
	fmt.Println(sum)
}
