package main

import "fmt"

func main() {
	fmt.Printf("FindOutlier([]int{2, 6, 8, -10, 3}): %v\n", FindOutlier([]int{2, 6, 8, -10, 3}))
}

func FindOutlier(integers []int) int {
	odds := []int{}
	evens := []int{}
	for _, v := range integers {
		if v%2 == 0 {
			evens = append(evens, v)
			continue
		}
		odds = append(odds, v)
	}
	if len(odds) == 1 {
		return odds[0]
	}
	return evens[0]
}
