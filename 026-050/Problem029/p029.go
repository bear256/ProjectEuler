package main

import (
	"fmt"
	"math/big"
)

func pow(a, b int) (bool, int) {
	flag := true
	ab := 1
	i := 0
	for ; i < b; i++ {
		ab *= a
		if ab > 100 {
			flag = false
			break
		}
	}
	return flag, ab
}

func solution1Wrong() {
	count := 99*99
	nums := []int{}
	for a := 2; a <= 100; a++ {
		for b := 2; b <= 100; b++ {
			flag, ab := pow(a, b)
			if flag {
				fmt.Println(a, b, ab, 100/b -1)
				nums = append(nums, ab)
				count -= 100/b-1
			}
		}
	}
	fmt.Println(count)
}

func solution2BinNum() {
	for a := 2; a <= 100; a++ {
		for b := 2; b <= 100; b++ {
		}
	}
}

func main() {
	big.NewInt(2)
}