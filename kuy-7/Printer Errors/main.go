package main

import "fmt"

func main() {
	fmt.Printf("PrinterError : %v\n", PrinterError("aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz"))
}

func PrinterError(s string) string {
	errors := 0
	for _, v := range s {
		if v > 109 {
			errors++
		}
	}
	return fmt.Sprintf("%d/%d", errors, len(s))
}
