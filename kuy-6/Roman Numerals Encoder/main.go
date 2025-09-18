package main

import (
	"fmt"
	"strings"
)

func main() {
	number := 29
	fmt.Printf("Roman number for the number: %d => %s\n", number, Solution(number))
}

func Solution(number int) (result string) {
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	for i, value := range values {
		if number < value {
			continue
		}
		result += strings.Repeat(symbols[i], number/value)
		number -= number / value * value
	}

	return result
}
