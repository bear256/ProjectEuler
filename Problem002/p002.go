package main

import (
	"fmt"
)

func main() {
	a0, a1 := 0, 1
	a2 := a0 + a1
	sum := 0
	for {
		if a2 > 4000000 {
			fmt.Println(sum)
			break
		}
		if a2%2 == 0 {
			sum += a2
		}
		a0 = a1
		a1 = a2
		a2 = a0 + a1
	}
}
