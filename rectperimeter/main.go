package main

import "fmt"

func RectPerimeter(width int, height int) int {
	if width < 0 || height < 0 {
		return -1
	}
	return 2 * (width + height)
}

func main() {
	fmt.Println(RectPerimeter(4, 5))  // 18
	fmt.Println(RectPerimeter(-3, 5)) // -1
}
