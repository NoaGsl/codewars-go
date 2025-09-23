package main

import "fmt"

func main() {
	fmt.Printf("SumDigPow(1, 100): %v\n", SumDigPow(12157692622039623100, 12157692622039625657))
	// fmt.Printf("SumDigPow(1, 100): %v\n", SumDigPow(1, 100))
}

func SumDigPow(a, b uint64) (list []uint64) {
	for i := a; i <= b; i++ {
		sum := GetSumDigPow(i)
		if sum == i {
			list = append(list, sum)
		}
	}

	return list
}

func GetSumDigPow(a uint64) (sum uint64) {
	digits := []uint64{}
	for a != 0 {
		digits = append(digits, a%10)
		a /= 10
	}
	for i := 1; i <= len(digits); i++ {
		sum += Pow(digits[len(digits)-i], i)
	}

	return sum
}

func Pow(a uint64, power int) (result uint64) {
	result = 1
	for i := 1; i <= power; i++ {
		result *= a
	}
	return result
}
