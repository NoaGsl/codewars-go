package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Previous multiple of 3: %d\n", PrevMultOfThree(952406))
}

func PrevMultOfThree(n int) interface{} {
	for newNumber := n; newNumber > 0; newNumber /= 10 {
		if newNumber%3 == 0 {
			return newNumber
		}
	}

	return nil
}
