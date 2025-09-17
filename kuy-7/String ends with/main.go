package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Strings end the same ?: %t\n", solution("banana", "a"))
}

func solution(str, ending string) bool {
	return strings.HasSuffix(str, ending)
}
