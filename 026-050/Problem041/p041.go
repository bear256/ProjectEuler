package main

import (
	"fmt"
	"math"
)

// In ...
func In(d int, ns [10]bool) bool {
	return ns[d]
}

// IsPrime ..
func IsPrime(n int) bool {
	flag := n > 1
	if n > 2 {
		flag = n%2 != 0
	}
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			flag = false
			break
		}
	}
	return flag
}

func generate(n, floor int, in [10]bool, result [9]int, max *int) {
	if n == floor {
		sum := 0
		for _, e := range result[:n] {
			sum = sum*10 + e
		}
		if IsPrime(sum) {
			if *max < sum {
				*max = sum
			}
		}
		fmt.Println(result[:n], sum)
	} else {
		for d := n; d > 0; d-- {
			if In(d, in) {
				continue
			}
			in[d] = true
			result[floor] = d
			generate(n, floor+1, in, result, max)
			in[d] = false
		}
	}
}

func main() {
	in := [10]bool{true, false, false, false, false, false, false, false, false, false}
	fmt.Println(In(1, in))
	max := 0
	for n := 8; n > 3; n-- {
		generate(n, 0, in, [9]int{}, &max)
	}
	fmt.Println(max)
}
