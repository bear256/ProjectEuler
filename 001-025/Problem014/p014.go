package main

import (
	"fmt"
)

func rule(n int) int {
	if n % 2 == 0 {
		n = n/2
	} else {
		n = 3*n+1
	}
	return n
}

func count(n int) int{
	pc := 1
	for {
		if n == 1 {
			break
		}
		n = rule(n)
		pc++
	}
	return pc
}

func main() {
	max, num := 0, 0
	for n := 1; n < 1000000; n++ {
		pc := count(n)
		if max < pc {
			max = pc
			num = n
		}
	}
	fmt.Println(max, num)
}