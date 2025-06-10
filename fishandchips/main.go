package main

import (
	"fmt"
)

func FishAndChips(n int) string {
	if n < 0 {
		return "error: number is not found"
	}
	div2 := n%2 == 0
	div3 := n%3 == 0
	switch {
	case div2 && div3:
		return "fish and chips"
	case div2:
		return "fish"
	case div3:
		return "chips"
	default:
		return "error: no divisible"
	}
}

func main() {
	fmt.Println(FishAndChips(4))
	fmt.Println(FishAndChips(9))
	fmt.Println(FishAndChips(6))
}
