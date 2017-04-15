package main

import (
	"fmt"
	"math"
)

func pn(n int) int {
	return n * (3*n - 1) / 2
}

func n(m int) int {
	return (1 + int(math.Sqrt(float64(1+24*m)))) / 6
}

func main() {
	distance1, distance2 := 100, 100
	flag1, flag2 := false, false
	for nb := 2; ; nb++ {
		for na := 1; na < nb; na++ {
			// a,b,c,d
			// method1 if c == b + a && c is Pn then if d = b+c && d is Pn then D =a
			// method2 if c == b + a && c is Pn then if d = a+c && d is Pn then D =b
			b := pn(nb)
			a := pn(na)
			c := a + b
			nc := n(c)
			if c == pn(nc) {
				d1 := b + c
				d2 := a + c
				nd1, nd2 := n(d1), n(d2)
				if d1 == pn(nd1) {
					distance1 = a
					flag1 = true
					break
				}
				if d2 == pn(nd2) {
					distance2 = b
					flag2 = true
					break
				}
			}
		}
		if flag1 || flag2 {
			break
		}
	}
	if flag1 {
		fmt.Println(distance1)
	} else {
		fmt.Println(distance2)
	}
}
