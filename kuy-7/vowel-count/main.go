package main

func GetCount(str string) (count int) {
	vowels := []string{"a", "e", "i", "o", "u"}
	for _, char := range str {
		if SliceContains(vowels, string(char)) {
			count++
		}
	}

	return count
}

func SliceContains(strings []string, value string) bool {
	for _, s := range strings {
		if s == value {
			return true
		}
	}

	return false
}
