package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Len of the shortest word: %v\n", FindShort("bitcoin take over the world maybe who knows perhaps"))
}

func FindShort(s string) (shortest int) {
	words := strings.Split(s, " ")
	shortest = len(words[0])
	for _, word := range words {
		if len(word) < shortest {
			shortest = len(word)
		}
	}
	return shortest
}
