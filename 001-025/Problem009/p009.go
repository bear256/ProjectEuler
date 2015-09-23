package main

import (
	"fmt"
)

func main() {
	for a := 1; a < 1000; a ++ {
		for b := 1; b < 1000; b++ {
			if a*a+b*b == (1000-a-b)*(1000-a-b) {
				fmt.Println(a*b*(1000-a-b))
				break
			}
		}
	}
}