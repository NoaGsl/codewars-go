package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Longest Consecutive strings: %v\n", LongestConsec([]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 2))
}

func LongestConsec(strarr []string, k int) (longestStr string) {
	for i := 0; i < len(strarr)-(k-1); i++ {
		assembledStr := ""
		for j := 0; j < k; j++ {
			assembledStr += strarr[i+j]
		}
		if len(assembledStr) > len(longestStr) {
			longestStr = assembledStr
		}
	}

	return longestStr
}
