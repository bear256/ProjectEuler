package main

import "math"
import "fmt"

// IsPrime ...
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

func test(n int) bool {
	flag := IsPrime(n)
	r2l := n
	if flag {
		ns := fmt.Sprint(n)
		for i := len(ns) - 1; i > 0; i-- {
			base := int(math.Pow10(i))
			l2r := n % base
			r2l = r2l / 10
			// fmt.Println(base, l2r, r2l)
			flag = flag && IsPrime(l2r) && IsPrime(r2l)
			if !flag {
				break
			}
		}
	}
	return flag
}

func main() {
	sum := 0
	i := 0
	for n := 11; i < 11; n += 2 {
		if test(n) {
			fmt.Println(n)
			sum += n
			i++
		}
	}
	fmt.Println(sum)
}
