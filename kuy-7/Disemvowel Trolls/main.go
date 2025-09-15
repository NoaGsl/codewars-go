package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Disemvoweled Comment: %s\n", Disemvowel("mEPc4MGFThyyri02BfJQBTDFhUzLQLLiv7EKJcvQiBiZsooDB6owMm5 ho7Xlyqv77bYNK5endOjlxFRc22iFjzBwzA3dXUzvXlp9OsfkUeyF5kbvksg0bS2yA5IG4A0dJBOJDW7UUrOPFxgvURjrsY3if4OWZ77 ZxIPASm1l8s25eLQHzXRlPC"))
}

func Disemvowel(comment string) string {
	disemvoweledComment := ""
	vowels := "aeiou"
	for _, char := range comment {
		if strings.Contains(vowels, strings.ToLower(string(char))) {
			continue
		}
		disemvoweledComment += string(char)
	}
	return disemvoweledComment
}
