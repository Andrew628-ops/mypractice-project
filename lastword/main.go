package main

import (
	"fmt"
	"strings"
)

func LastWord(s string) string {
	fields := strings.Fields(s)
	if len(fields) == 0 {
		return "\n"
	}
	return fields[len(fields)-1] + "\n"
}

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}
