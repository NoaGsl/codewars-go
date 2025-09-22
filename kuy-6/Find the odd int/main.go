package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Odd int: %v\n", FindOdd([]int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}))
}

func FindOdd(seq []int) int {
	numberCount := map[int]int{}
	for _, number := range seq {
		numberCount[number]++
	}
	for number, count := range numberCount {
		if count%2 == 1 {
			return number
		}
	}
	return 0
}
