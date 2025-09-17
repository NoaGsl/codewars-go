package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Camel case string: %s\n", CamelCase("test case"))
}

func CamelCase(s string) string {
	words := strings.Split(s, " ")
	for i, word := range words {
		if len(word) == 0 {
			continue
		}

		words[i] = strings.ToUpper(word[:1]) + word[1:]
	}

	return strings.Join(words, "")
}
