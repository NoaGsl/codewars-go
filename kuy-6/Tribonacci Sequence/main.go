package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Solution: %v\n", Tribonacci([3]float64{1, 1, 1}, 30))
}

func Tribonacci(signature [3]float64, n int) (result []float64) {
	if n < 3 {
		return signature[0:n]
	}

	result = signature[:]
	for i := 2; i+1 < n; i++ {
		result = append(result, result[i]+result[i-1]+result[i-2])
	}

	return result
}
