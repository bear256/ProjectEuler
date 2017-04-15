package main

import (
	"fmt"
)

func main() {
	sum := 0
	max := 10000000000
	for m := 1; m <= 1000; m++ {
		prod := 1
		for n := 1; n <= m; n++ {
			prod *= m
			if prod > max {
				prod %= max
			}
			// Improve the performance
			// if prod == 0 {
			// 	break
			// }
		}
		sum += prod
		if sum > max {
			sum -= max
		}
	}
	fmt.Println(sum)
}
