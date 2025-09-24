package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Diff: %v\n", ArrayDiff([]int{1, 2}, []int{1}))
}

func ArrayDiff(a, b []int) (diff []int) {
	for _, v := range a {
		if !Contains(b, v) {
			diff = append(diff, v)
		}
	}
	return diff
}

func Contains(a []int, val int) bool {
	for _, v := range a {
		if v == val {
			return true
		}
	}
	return false
}
