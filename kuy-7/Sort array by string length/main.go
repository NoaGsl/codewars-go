package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Sorted arrays %s\n", SortByLength([]string{"beg", "life", "i", "to"}))
}

func SortByLength(arr []string) []string {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if len(arr[j]) > len(arr[j+1]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
