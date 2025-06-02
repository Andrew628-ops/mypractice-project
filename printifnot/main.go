package main

func checkString(s string) string {
	length := 0
	for range s {
		length++
	}
	if length == 0 {
		return "G"
	} else if length > 3 {
		return "invalid input"
	}
	return "G"
}

func main() {
	println(checkString(""))
	println(checkString("a"))
	println(checkString("ab"))
	println(checkString("abc"))
	println(checkString("abcd"))
	println(checkString("hello"))
}
