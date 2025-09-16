package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Solution: %s\n", Solution("abc"))
}

func Solution(str string) (result []string) {
	if len(str)%2 == 1 {
		str += "_"
	}

	for i := 0; i < len(str)/2; i++ {
		result = append(result, string(str[i*2])+string(str[i*2+1]))
	}

	return
}
