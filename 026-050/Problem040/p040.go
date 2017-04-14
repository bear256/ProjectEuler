package main

import (
	"fmt"
	"strconv"
)

func method1() {
	prod := 1
	ns := []int{1, 10, 100, 1000, 10000, 100000, 1000000}
	decfrac := ""
	for i := 0; ; i++ {
		decfrac = fmt.Sprintf("%s%d", decfrac, i)
		if len(decfrac) > 1000000 {
			break
		}
	}
	for _, n := range ns {
		dig, _ := strconv.Atoi(string(decfrac[n]))
		prod *= dig
	}
	fmt.Println(prod)
}

func main() {
	prod := 1
	ns := []int{1, 10, 100, 1000, 10000, 100000, 1000000}
	idx := 0
	count := 0
	for n := 1; ; n++ {
		num := fmt.Sprintf("%d", n)
		length := len(num)
		count += length
		if count >= ns[idx] {
			i := length - (count - ns[idx]) - 1
			fmt.Println(n, i, num[i]-'0')
			prod *= int(num[i] - '0')
			idx++
		}
		if count > 1000000 {
			break
		}
	}
	fmt.Println(prod)
}
