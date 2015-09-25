package main

import (
	"strconv"
	"strings"
	"fmt"
)

const test = `3
7 4
2 4 6
8 5 9 3`

const real = `75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`

var max int = 0

func search(m [][]int, i, j, sum int) {
	if i < len(m) && j < len(m[i]) {
		search(m, i+1, j, sum+m[i][j])
		search(m, i+1, j+1, sum+m[i][j])
	} 
	if max < sum {
		max = sum
	}
}

func main() {
	content := real
	m := [][]int{}
	for _, row := range strings.Split(content, "\n") {
		r := []int{}
		for _, col := range strings.Split(row, " ") {
			i, _ := strconv.Atoi(col)
			r = append(r, i)
		}
		m = append(m, r)
	}
//	fmt.Println(m)
	search(m, 0, 0, 0)
	fmt.Println(max)
}