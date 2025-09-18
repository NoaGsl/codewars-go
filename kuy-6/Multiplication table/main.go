package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Multiplication table: %d\n", MultiplicationTable(10))
}

func MultiplicationTable(size int) (table [][]int) {
	for i := 1; i <= size; i++ {
		row := []int{}
		for j := 1; j <= size; j++ {
			row = append(row, i*j)
		}
		table = append(table, row)
	}

	return table
}
