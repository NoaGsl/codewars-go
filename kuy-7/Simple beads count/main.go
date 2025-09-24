package main

import "fmt"

func main() {
	fmt.Printf("CountRedBeads(5): %v\n", CountRedBeads(5))
}

func CountRedBeads(n int) int {
	if n == 0 {
		return 0
	}
	return n*2 - 2
}
