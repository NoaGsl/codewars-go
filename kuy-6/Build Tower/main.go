package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Tower: %v\n", TowerBuilder(6))
}

func TowerBuilder(nFloors int) (tower []string) {
	floor := []byte{}
	for i := 1; i < nFloors*2; i++ {
		floor = append(floor, ' ')
	}

	for i := range nFloors {
		floor[nFloors-1-i] = '*'
		floor[nFloors-1+i] = '*'

		tower = append(tower, string(floor))
	}
	return tower
}
