package main

import (
	"fmt"
)

func main() {
	// var name string = "John"
	// fmt.Println("Hello,", name)

	// var name string
	// var lastname string
	// name, lastname = "Mert", "Yılmaz"

	// shorter way to declare variables
	// var name, lastname string = "Mert", "Yılmaz"

	// shortest way to declare a variable
	name, lastname := "Mert", "Yılmaz"

	var age int
	age = 24

	var isCool bool
	isCool = true

	fmt.Println("Hello,", name, lastname)
	fmt.Println("You are", age, "years old")
	fmt.Println("Is he cool?", isCool)
}
