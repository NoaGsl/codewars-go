package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Fixed string case %s\n", solve("coDe"))
}

func solve(str string) string {
	upperCount := 0
	for _, letter := range str {
		if letter < 'a' {
			upperCount++
		}
	}

	if upperCount == 0 {
		return str
	}
	if upperCount > len(str)/2 {
		return strings.ToUpper(str)
	}
	return strings.ToLower(str)
}
