package main

import (
	"fmt"
)

var sum int = 0

func plus(na, nb []int) []int {
	return []int{na[0]+nb[0],na[1]+nb[1]}
}

func count(root []int, dg int) int {
	if root[0] == dg || root[1] == dg {
		sum ++
		fmt.Println(sum)
		return 1
	} else {
		return count(plus(root, []int{0, 1}), dg) + count(plus(root, []int{1, 0}), dg)
	}
}

func main() {
	fmt.Println(count([]int{0,0}, 4))
}