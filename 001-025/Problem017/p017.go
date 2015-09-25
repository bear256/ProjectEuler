package main

import (
	"fmt"
)

var (
	words = map[int]string{1: "one", 2: "two", 3: "three", 4:"four", 5:"five", 6:"six", 7:"seven", 8:"eight", 9:"nine",
		10:"ten", 11:"eleven", 12:"twelve", 13:"thirteen", 14:"fourteen", 15:"fifteen", 16:"sixteen", 17:"seventeen", 18:"eighteen", 19:"nineteen",
		20:"twenty", 30:"thirty", 40:"forty", 50:"fifty", 60:"sixty", 70:"seventy", 80:"eighty", 90:"ninety", 100:"hundred", 1000:"thousand",
	}
)

func lt100(n int) int {
	sum := 0
	if n <= 20 {
		sum += len(words[n])
	} else if n < 100 {
		units := n%10
		tens := n/10*10
		sum += len(words[tens]) + len(words[units])
	}
	return sum
}

func lt1000(n int) int {
	sum := 0
	tens := n%100
	if n < 100 {
		sum = lt100(n)
	} else {
		if tens > 0 {
			sum += len("and") + lt100(tens)
		}
		hundreds := n/100
		sum += len(words[hundreds]) + len("hundred")
	}
	return sum
}

func main() {
	sum := 0
	// fmt.Println(lt100(42))
	// fmt.Println(lt1000(115))
	// fmt.Println(lt1000(342))
	for n := 1; n <= 1000; n++ {
		 if n < 1000 {
			 sum += lt1000(n)
		 } else {
			 thousands := n/1000
			 sum += len(words[thousands]) + len("thousand")
		 }
	}
	fmt.Println(sum)
}