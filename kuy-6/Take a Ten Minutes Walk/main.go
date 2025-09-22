package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Is walk valid? %t\n", IsValidWalk([]rune{'n', 's', 'n', 's', 'n', 's', 'n', 's', 'n', 's'}))
}

func IsValidWalk(walk []rune) bool {
	if len(walk) != 10 {
		return false
	}

	eastWestCount := 0
	northSouthCount := 0
	for _, v := range walk {
		switch v {
		case 'n':
			northSouthCount++
		case 's':
			northSouthCount--
		case 'e':
			eastWestCount++
		case 'w':
			eastWestCount--
		}
	}
	if eastWestCount == 0 && northSouthCount == 0 {
		return true
	}
	return false
}
