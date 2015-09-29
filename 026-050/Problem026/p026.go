package main

import (
	"fmt"
)

func length(d int) int {
	l := 0
	t := 10
	for ; t < d; {
		t *=10
	}
	rems := []int{}
	flag := false
	for i :=0; ; i++ {
		r := t%d
		// fmt.Println(i, r, t/d)
		for idx, rem := range rems {
			if r == rem {
				l = i - idx
				flag = true
				break
			}
		}
		if flag {
			break
		}
		rems = append(rems, r)
		t = r*10
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
	fmt.Println(n)
}