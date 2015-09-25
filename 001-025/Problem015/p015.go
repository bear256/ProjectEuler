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

func solution1Search(dg int) {
	fmt.Println(count([]int{0,0}, dg))
}

func calcSum(array []int) int {
	for ;len(array) > 1;{
		tmp := make([]int, len(array)-1)
		for i := 0; i < len(array)-1; i++ {
			tmp[i] = array[i] + array[i+1]
		}
		array = tmp
	}
	return array[0]
}

func solution2Construct(dg int) {
	array := []int{2}
	if dg > 1{
		for i := 2; i <= dg; i++ {	
			tmp := array
			array = make([]int, i)
			array[0] = i+1
			array[i-1] = i+1
			for j := 1; j < i-1; j++ {
				array[j] = tmp[j-1]+tmp[j]
			}
		}
	}
	fmt.Println(calcSum(array))
}

func main() {
	// solution1Search(20)
	solution2Construct(20)
}
