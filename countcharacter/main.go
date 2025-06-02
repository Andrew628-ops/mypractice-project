package main

import (
	"fmt"
)

func countCharactersInOrder(s string) {
	counts := make(map[rune]int)
	order := []rune{}

	for _, r := range s {
		if _, exists := counts[r]; !exists {
			order = append(order, r)
		}
		counts[r]++
	}

	for _, r := range order {
		fmt.Printf("'%c' : %d\n", r, counts[r])
	}
}

func main() {
	text := "hello world"
	countCharactersInOrder(text)
}
