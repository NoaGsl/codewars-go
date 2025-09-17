package main

import (
	"fmt"
)

func main() {
	apollo := Dog{"Apollo", "Dobermann", "male", 4}
	fmt.Printf("The dog bark: %s\n", apollo.Bark())
}

type Dog struct {
	Name  string
	Breed string
	Sex   string
	Age   int
}

func (dog *Dog) Bark() string {
	return "Woof!"
}
