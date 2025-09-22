package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Alternated Capitalization %s\n", Capitalize("abcdef"))
}

func Capitalize(st string) []string {
	result := [][]byte{[]byte(st), []byte(st)}
	for i, v := range st {
		upper := strings.ToUpper(string(v))
		if i%2 == 0 {
			result[0][i] = upper[0]
			continue
		}

		result[1][i] = upper[0]
	}

	return []string{string(result[0]), string(result[1])}
}
