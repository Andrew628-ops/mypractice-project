package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println()
		return
	}
	seen := make(map[rune]bool)
	for _, c := range os.Args[1] + os.Args[2] {
		if !seen[c] {
			seen[c] = true
			fmt.Printf("%c", c)
		}
	}
	fmt.Println()
}
