package main

import "fmt"

// camelToSnake converts a camelCase or PascalCase string to snake_case
func camelToSnake(s string) string {
	var result []rune
	for i, r := range s {
		if r >= 'A' && r <= 'Z' {
			if i > 0 && s[i-1] != '_' && !(s[i-1] >= 'A' && s[i-1] <= 'Z') {
				result = append(result, '_')
			}
			// Convert uppercase to lowercase
			result = append(result, r+'a'-'A')
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

func main() {
	fmt.Println(camelToSnake("HelloWorld"))       // hello_world
	fmt.Println(camelToSnake("helloWorld"))       // hello_world
	fmt.Println(camelToSnake("camelCase"))        // camel_case
	fmt.Println(camelToSnake("CAMELtoSnackCASE")) // camel_to_snack_case
	fmt.Println(camelToSnake("camelToSnakeCase")) // camel_to_snake_case
	fmt.Println(camelToSnake("hey2"))             // hey2
}
