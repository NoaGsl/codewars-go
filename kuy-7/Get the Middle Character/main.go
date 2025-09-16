package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Middle characters: %s\n", GetMiddle("test"))
}

func GetMiddle(s string) string {
	middle := len(s) / 2
	if len(s)%2 == 0 {
		return s[middle-1 : middle+1]
	}
	return string(s[middle])
}
