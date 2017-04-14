package main

import (
	"fmt"
)

// IsPalindromic ...
func IsPalindromic(s string) bool {
	length := len(s)
	flag := true
	for i := 0; i < length/2; i++ {
		flag = flag && s[i] == s[length-i-1]
	}
	return flag
}

func test(n int) bool {
	base10 := fmt.Sprint(n)
	base2 := fmt.Sprintf("%b", n)
	return IsPalindromic(base10) && IsPalindromic(base2)
}

func main() {
	sum := 0
	for n := 1; n < 1000000; n++ {
		if test(n) {
			sum += n
		}
	}
	fmt.Println(sum)
}
