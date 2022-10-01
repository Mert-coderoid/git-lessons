package main

import "fmt"

func main() {
	// (print - println) - printf
	fmt.Print("Hello", "World")
	fmt.Println("")
	fmt.Println("Hello World")
	fmt.Printf("Hello %s", "World")
	fmt.Println("")
	fmt.Printf("Hello %v", "World")
	fmt.Println("")
	fmt.Printf("Hello %T", "World")
	fmt.Println("")
	fmt.Printf("Hello", "World")
}
