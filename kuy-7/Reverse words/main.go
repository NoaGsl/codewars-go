package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Reversed words: %s\n", ReverseWords("double  spaced  words"))
}

func ReverseWords(str string) string {
	words := strings.Split(str, " ")

	for i, word := range words {
		words[i] = ReverseWord(word)
	}

	return strings.Join(words, " ")
}

func ReverseWord(word string) (reverse string) {
	for i := len(word) - 1; i >= 0; i-- {
		reverse += string(word[i])
	}
	return
}
