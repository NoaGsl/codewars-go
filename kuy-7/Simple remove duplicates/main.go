package main

import "fmt"

func main() {
	fmt.Printf("Duplicates removed: %v\n", Solve([]int{3, 4, 4, 3, 6, 3}))
}

func Solve(arr []int) (result []int) {
	encountered := map[int]bool{}
	for i := len(arr) - 1; i >= 0; i-- {
		if !encountered[arr[i]] {
			encountered[arr[i]] = true
			result = append([]int{arr[i]}, result...)
		}
	}
	return result
}
