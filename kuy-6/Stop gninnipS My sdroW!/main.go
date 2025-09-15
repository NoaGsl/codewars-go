package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Spinned words %s\n", SpinWords("Hey fellow warriors"))
}

func SpinWords(str string) string {
	result := []string{}
	splittedStrings := strings.Split(str, " ")

	for _, word := range splittedStrings {
		if len(word) >= 5 {
			result = append(result, ReverseString(word))
			continue
		}
		result = append(result, word)
	}

	return strings.Join(result, " ")
} // SpinWords

func ReverseString(str string) (reverse string) {
	for i := len(str) - 1; i >= 0; i-- {
		reverse += string(str[i])
	}
	return
}
