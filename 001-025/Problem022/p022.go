package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type ByAlpha []string

func (a ByAlpha) Len() int           { return len(a) }
func (a ByAlpha) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAlpha) Less(i, j int) bool { return a[i] < a[j] }

func worth(name string) int {
	sum := 0
	for _, alpha := range name {
		sum += int(alpha-'A') + 1
	}
	return sum
}

func main() {
	bts, _ := ioutil.ReadFile("p022_names.txt")
	total := 0
	names := strings.Split(string(bts), ",")
	sort.Sort(ByAlpha(names))
	for idx, name := range names {
		name = strings.Replace(name, "\"", "", -1)
		score := (idx + 1) * worth(name)
		total += score
	}
	fmt.Println(total)
}
