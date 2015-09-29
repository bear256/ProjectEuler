package main

import (
	"fmt"
)

func sumByLevel(level int) int {
	sum := 0
	max := level*level
	sum += max
	for i := 1; i < 4; i++ {
		t := max - i*(level - 1)
		sum += t
	}
	return sum
}

func main() {
	sum := 1
	for level := 3; level <= 1001; level += 2 {
		sum += sumByLevel(level)
	}
	fmt.Println(sum)
}