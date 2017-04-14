package main

import (
	"fmt"
	"strconv"
)

func test(n int) (bool, string, int) {
	flag := true
	check := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	result := ""
	seq := 0
	for i := 1; i < 10; i++ {
		prod := fmt.Sprint(n * i)
		if len(result)+len(prod) > 9 {
			break
		}
		result += prod
		for _, p := range prod {
			idx, _ := strconv.Atoi(string(p))
			if check[idx] > 0 {
				flag = false
				break
			}
			check[idx]++
		}
		if !flag {
			break
		}
		seq = i
	}
	sum := 0
	for _, c := range check[1:] {
		sum += c
	}
	return sum == 9, result, seq
}

func main() {
	max := "0"
	maxN := 1
	maxI := 0
	for n := 1; n < 10000; n++ {
		if flag, prod, i := test(n); flag {
			if max < prod {
				max = prod
				maxN = n
				maxI = i
			}
		}
	}
	fmt.Println(max, maxN, maxI)
}
