package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Highest scoring word: %v\n", High("man i need a taxi up to ubud"))
}

func High(s string) (highest string) {
	highestScore := 0
	words := strings.Split(s, " ")

	for _, word := range words {
		wordScore := 0
		for _, runeLetter := range word {
			wordScore += int(runeLetter - 96)
		}
		if wordScore > highestScore {
			highestScore = wordScore
			highest = word
		}
	}

	return highest
}
