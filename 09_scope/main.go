package main

import (
	"fmt"
)

// packVar := "I am a package variable"		// non-declaration statement outside function body  -> tr : fonksiyon dışında tanımlama ifadesi olamaz
var packVar = "I am a package variable"

func main() {

	if true {
		var blockVar = "I am a block variable"
		fmt.Println(blockVar)
	}

	var funcVar string = "I am a function variable"

	fmt.Println(funcVar)
	fmt.Println(packVar)
	// fmt.Println(blockVar)	// undefined: blockVar
}

func myFunc() {
	fmt.Println(packVar)
	// fmt.Println(blockVar)	// undefined: blockVar
	// fmt.Println(funcVar) 	// undefined: funcVar
}
