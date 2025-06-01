package main

import "fmt"

func main() {
	name := "otokpa okibe andrew"
	age := 18
	marrital := "single"
	phone := 07025750640
	email := "andrewokibeotokpa@gmail.com"

	const fruit = "apple"

	fmt.Printf("my name is %v i am %d years old i am %v my phone number is %.2d my email address is %v", name, age, marrital, phone, email)
	fmt.Println(fruit)

	step1 := "Breathe in..."
	step2 := "Breathe out..."

	meditation := fmt.Sprintln(step1, step2)

	fmt.Print(meditation)
}
