package main

import (
	"fmt"
)

func main() {
	fmt.Printf("String: %s\n", CleanString("abc#d##c"))
}

func CleanString(s string) (result string) {
	for _, v := range s {
		if string(v) != "#" {
			result += string(v)
			continue
		}

		if len(result) > 0 {
			result = result[:len(result)-1]
		}
	}
	return result
}
