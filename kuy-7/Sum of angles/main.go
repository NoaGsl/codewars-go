package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Angle(3): %v\n", Angle(3))
}

func Angle(n int) int {
	return (n - 2) * 180
}
