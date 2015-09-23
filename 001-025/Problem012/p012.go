package main

import (
	"fmt"
	"math"
)

func factorsNum(num int) int {
	count := 2
	if num == 1 {
		count = 1
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num % i == 0 {
			count += 2
		}
	}
//	fmt.Println(num, count)
	return count
}

func main() {
	count := 1
	n := 1
	for ;; {
		l := factorsNum(n)
		if l > 500 {
			fmt.Println(n)
			break
		}
		count ++
		n = n + count
	}
}