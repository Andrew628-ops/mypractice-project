package main

import (
	"fmt"
)

func HashCode(dec string) string {
	n := len(dec)
	res := []rune{}
	for _, c := range dec {
		h := (int(c) + n) % 127
		if h < 33 {
			h += 33
		}
		res = append(res, rune(h))
	}
	return string(res)
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
