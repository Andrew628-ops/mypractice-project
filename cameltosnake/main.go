package main

// camelToSnake converts a camelCase or PascalCase string to snake_case
func camelToSnake(s string) string {
	var result []rune
	for i, r := range s {
		if r >= 'A' && r <= 'Z' {
			if i > 0 {
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
	println(camelToSnake("camelCase"))   // Output: camel_case
	println(camelToSnake("CamelCase"))   // Output: camel_case
	println(camelToSnake("ThisIsATest")) // Output: this_is_a_test
	println(camelToSnake("simpleTest"))  // Output: simple_test
}
