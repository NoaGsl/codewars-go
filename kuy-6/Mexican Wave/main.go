package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Mexican wave: %v\n", wave(" x yz"))
}

func wave(str string) (wave []string) {
	for i := 0; i < len(str); i++ {
		if string(str[i]) == " " {
			continue
		}

		s := []byte(str)
		s[i] = strings.ToUpper(string(s[i]))[0]
		wave = append(wave, string(s))
	}

	return wave
}
