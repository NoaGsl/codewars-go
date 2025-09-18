package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Jaden case srting: %s\n", ToJadenCase("most trees are blue"))
}

func ToJadenCase(str string) string {
	words := strings.Split(str, " ")
	for i, word := range words {
		words[i] = strings.ToUpper(word[:1]) + word[1:]
	}

	return strings.Join(words, " ")
}
