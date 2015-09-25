package main

import (
	"fmt"
	"time"
)

var mondays = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func solution1() {
	count := 0
	for year := 1901; year <= 2000; year++ {
		for month := 1; month <= 12; month++ {
			datestr := fmt.Sprintf("%d-%02d-01", year, month)
			date, _ := time.Parse("2006-01-02", datestr)
			if date.Weekday() == time.Sunday {
				count++
			}
		}
	}
	fmt.Println(count)
}

func solution2() {
	count := 0
	days := 365
	for year := 1901; year <= 2000; year++ {
		for month := 1; month <= 12; month++ {
			if (days + 1) % 7 == 0 {
				count++
			}
			days += mondays[month-1]
			if year % 4 == 0 && month == 2 {
				days++
			}
		}
	}
	fmt.Println(count)
}

func main() {
//	solution1()
	solution2()
}