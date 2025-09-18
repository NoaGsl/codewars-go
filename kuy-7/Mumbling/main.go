package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Mumbling: %s\n", Accum("abcd"))
}

func Accum(s string) string {
	result := []string{}
	for i, v := range s {
		result = append(result, strings.ToUpper(string(v))+strings.Repeat(strings.ToLower(string(v)), i))
	}
	return strings.Join(result, "-")
}
