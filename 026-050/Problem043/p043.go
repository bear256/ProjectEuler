package main

import "fmt"

// In ...
func In(d int, ns [10]bool) bool {
	return ns[d]
}

func ints2int(number []int) int {
	sum := 0
	for _, e := range number {
		sum = sum*10 + e
	}
	return sum
}

func test(number []int) bool {
	flag := true
	primes := []int{0, 2, 3, 5, 7, 11, 13, 17}
	for i := 1; i < 8; i++ {
		ddd := ints2int(number[i : i+3])
		flag = flag && ddd%primes[i] == 0
		if !flag {
			break
		}
	}
	return flag
}

func generate(n, floor int, in [10]bool, result []int, sum *int) {
	if n == floor {
		flag := test(result)
		if flag {
			num := ints2int(result)
			*sum += num
			fmt.Println(result, num, flag)
		}
	} else {
		for d := 0; d < n; d++ {
			if In(d, in) {
				continue
			}
			in[d] = true
			result[floor] = d
			generate(n, floor+1, in, result, sum)
			in[d] = false
		}
	}
}

func main() {
	// test([10]int{1, 4, 0, 6, 3, 5, 7, 2, 8, 9})
	in := [10]bool{false, false, false, false, false, false, false, false, false, false}
	result := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	sum := 0
	generate(10, 0, in, result, &sum)
	fmt.Println(sum)
}
