package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Value of tje word %v\n", WordsToMarks("attitude"))
}

func WordsToMarks(s string) (result int) {
	for _, v := range s {
		result += int(v - 'a' + 1)
	}
	return result
}
