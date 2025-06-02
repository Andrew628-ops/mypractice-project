package main

import "fmt"

func countIfNotLetter(s string) int {
	count := 0
	for _, r := range s {
		if !(r >= 'A' && r <= 'Z') && !(r >= 'a' && r <= 'z') {
			count++
		}
	}
	return count
}

func main() {
	text := "Hello 123!"
	result := countIfNotLetter(text)
	fmt.Printf("Characters that are NOT letters: %d\n", result)
}
