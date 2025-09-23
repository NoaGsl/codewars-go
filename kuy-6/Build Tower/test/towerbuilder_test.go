package main

import (
	"strings"
	"testing"
)

// --- Code 1 (strings.Repeat) ---
func TowerBuilder1(nFloors int) []string {
	tow := make([]string, nFloors)
	for i := 0; i < nFloors; i++ {
		spaces := strings.Repeat(" ", nFloors-(i+1))
		blocks := strings.Repeat("*", i*2+1)
		tow[i] = spaces + blocks + spaces
	}
	return tow
}

// --- Code 2 (mutating slice) ---
func TowerBuilder2(nFloors int) (tower []string) {
	floor := []byte{}
	for i := 1; i < nFloors*2; i++ {
		floor = append(floor, ' ')
	}

	for i := 0; i < nFloors; i++ {
		floor[nFloors-1-i] = '*'
		floor[nFloors-1+i] = '*'
		tower = append(tower, string(floor))
	}
	return tower
}

// --- Code 3 (preallocated row) ---
func TowerBuilder3(nFloors int) []string {
	tower := make([]string, nFloors)
	width := nFloors * 2

	for i := 0; i < nFloors; i++ {
		row := make([]byte, width)
		for j := range row {
			row[j] = ' '
		}
		start := nFloors - 1 - i
		end := nFloors - 1 + i
		for j := start; j <= end; j++ {
			row[j] = '*'
		}
		tower[i] = string(row)
	}
	return tower
}

// --- Code 5 (optimized reuse + selective copy) ---
func TowerBuilder4(nFloors int) []string {
	tower := make([]string, nFloors)
	width := nFloors * 2
	floor := make([]byte, width)

	// Pre-fill with spaces once
	for i := range floor {
		floor[i] = ' '
	}

	for i := 0; i < nFloors; i++ {
		start := nFloors - 1 - i
		end := nFloors - 1 + i

		// Place stars for this row
		for j := start; j <= end; j++ {
			floor[j] = '*'
		}

		// Copy only this row’s snapshot
		tower[i] = string(append([]byte(nil), floor...))

		// Reset back to spaces
		for j := start; j <= end; j++ {
			floor[j] = ' '
		}
	}

	return tower
}

// --- Code 5 me trying to optimize
func TowerBuilder5(nFloors int) []string {
	tower := make([]string, nFloors)
	floor := make([]byte, nFloors*2)

	for i := 1; i < nFloors*2; i++ {
		floor[i] = ' '
	}

	for i := 0; i < nFloors; i++ {
		floor[nFloors-1-i] = '*'
		floor[nFloors-1+i] = '*'
		tower[i] = string(floor)
	}

	return tower
}

// --- Benchmarks ---
func BenchmarkTowerBuilder1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TowerBuilder1(1000)
	}
}

func BenchmarkTowerBuilder2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TowerBuilder2(1000)
	}
}

func BenchmarkTowerBuilder3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TowerBuilder3(1000)
	}
}

func BenchmarkTowerBuilder4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TowerBuilder4(1000)
	}
}

func BenchmarkTowerBuilder5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TowerBuilder5(1000)
	}
}
