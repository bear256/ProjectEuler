package main

import (
	"fmt"
	"strconv"
)

func test(a, b, c int) bool {
	flag := true
	check := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	result := fmt.Sprintf("%d%d%d", a, b, c)
	if len(result) != 9 {
		flag = false
	} else {
		for _, p := range result {
			idx, _ := strconv.Atoi(string(p))
			if check[idx] > 0 {
				flag = false
				break
			}
			check[idx]++
		}
		sum := 0
		for _, c := range check[1:] {
			sum += c
		}
		flag = sum == 9
	}
	return flag
}

func main() {
	sum := 0
	m := make(map[int]bool)
	for a := 1; a < 1000; a++ {
		for b := 2; b < 1000000/a; b++ {
			c := a * b
			if _, ok := m[c]; !ok {
				if test(a, b, c) {
					fmt.Println(a, b, c)
					m[c] = true
					sum += c
				}
			}
		}
	}
	fmt.Println(sum)
}
