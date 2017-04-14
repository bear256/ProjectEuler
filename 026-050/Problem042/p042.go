package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func test(word string) bool {
	sum := 0
	for _, letter := range word {
		position := int(letter - 'A' + 1)
		sum += position
	}
	n := int(math.Sqrt(float64(sum * 2)))
	return n*(n+1) == sum*2
}

func main() {
	txt, _ := ioutil.ReadFile("p042_words.txt")
	words := strings.Split(string(txt), ",")
	count := 0
	for _, word := range words {
		if test(strings.Trim(word, "\"")) {
			fmt.Println(word)
			count++
		}
	}
	fmt.Println(count)
}
