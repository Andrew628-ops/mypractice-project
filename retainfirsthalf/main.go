package main

import (
	"fmt"
)

func RetainFirstHalf(str string) string {
	n := len(str)
	if n == 0 || n == 1 {
		return str
	}
	return str[:n/2]
}

func main() {
	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(RetainFirstHalf("A"))
	fmt.Println(RetainFirstHalf(""))
	fmt.Println(RetainFirstHalf("Hello World"))
}
