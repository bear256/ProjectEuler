package main

import (
	"math"
	"strconv"
	"fmt"
)

func isPalindromic(num int) bool {
	flag := true
	numstr := strconv.Itoa(num)
	length := len(numstr)
	for i:=0; i < length/2; i++ {
		if numstr[i] != numstr[length - i - 1] {
			flag = false
			break
		}
	}
	return flag
}

func canFactTwo3digitsNum(num int) bool {
	a := int(math.Sqrt(float64(num)))
	b := 0
	for n := a; n < num; n++ {
		if num%n == 0 {
			a = n
			b = num/n
			break
		}
	}
	return a < 1000 && b < 1000
}

func main() {
	for i:= 1000*1000 - 1; ;i-- {
		if isPalindromic(i) && canFactTwo3digitsNum(i) {
			fmt.Println(i)
			break
		}
	}
}