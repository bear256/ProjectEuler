package main

import (
	"fmt"
)

func length(d int) int {
	t := 10
	decs := []int{}
	l := 0
	flag := false
	for i := 0;; i++{
		digit := t/d
		// fmt.Println(digit, decs)
		for idx, dec := range decs {
			if dec == digit {
				l = i - idx
				flag = true
				break
			}
		}
		if flag {
			break
		}
		decs = append(decs, digit)
		t = t%d*10
	}
	return l
}

func main() {
	max, n := 0, 2
	for d := 2; d < 1000; d++ {
		l := length(d)
		if max < l {
			max = l
			n = d
		}
	}
	fmt.Println(max, n)
}