package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Highest and Lowest: %s\n", HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4"))
}

func HighAndLow(in string) string {
	stringsNumber := strings.Split(in, " ")
	firstElement, err := strconv.Atoi(stringsNumber[0])
	if err != nil {
		return ""
	}

	lowest := firstElement
	highest := firstElement

	for _, stringNumber := range stringsNumber {
		number, err := strconv.Atoi(stringNumber)
		if err != nil {
			return ""
		}

		if number < lowest {
			lowest = number
			continue
		}
		if number > highest {
			highest = number
		}
	}
	return fmt.Sprintf("%v %v", highest, lowest)
}
