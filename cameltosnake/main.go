package main

import (
	"fmt"
	"unicode"
)

// camelToSnake converts a camelCase or PascalCase string to snake_case
func camelToSnake(s string) string {
	var result []rune
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(r))
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

func main() {
	fmt.Println(camelToSnake("camelCase"))   // Output: camel_case
	fmt.Println(camelToSnake("CamelCase"))   // Output: camel_case
	fmt.Println(camelToSnake("ThisIsATest")) // Output: this_is_a_test
	fmt.Println(camelToSnake("simpleTest"))  // Output: simple_test
}
