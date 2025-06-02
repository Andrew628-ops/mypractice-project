package main

func checkString(s string) string {
	if len(s) == 0 || len(s) <= 3 {
		return "G"
	}
	return "invalid input"
}

func main() {
	println(checkString(""))
	println(checkString("a"))
	println(checkString("ab"))
	println(checkString("abc"))
	println(checkString("abcd"))
	println(checkString("hello"))
}
