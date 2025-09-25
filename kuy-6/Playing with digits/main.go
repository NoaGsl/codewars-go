package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("DigPow(89, 1): %v\n", DigPow(89, 1))
}

func DigPow(n, p int) int {
	sum := SumDigPow(n, p)

	if sum%n == 0 {
		return sum / n
	}

	return -1
}

func SumDigPow(a int, p int) (sum int) {
	digits := []int{}
	for a != 0 {
		digits = append(digits, a%10)
		a /= 10
	}
	for i := range digits {
		sum += int(math.Pow(float64(digits[len(digits)-1-i]), float64(p)))
		p++
	}
	return sum
}
