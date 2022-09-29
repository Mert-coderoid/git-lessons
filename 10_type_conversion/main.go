package main

import (
	"fmt"
	"strconv"
)

func main() {

	/*

			x := 10
			y := 3.14

			fmt.Printf("%v %T\n", x, x)
			fmt.Printf("%v %T\n", y, y)

			// Type conversion type(value) => int(y) = 10.0 => 10
			fmt.Println(x + int(y))     // 13
			fmt.Println(float64(x) + y) // 13.14

			fmt.Printf("%v %T\n", x, x)
			fmt.Printf("%v %T\n", y, y)

			var a int8 = 10
			var b int16 = 20

			fmt.Println(a + int8(b )) // 30



		var x int16 = 128
		var y int8

		// y = x	// cannot use x (variable of type int8) as type int16 in assignment
		y = int8(x) // -128	// tr: 128 sayısı 8 bitlik bir sayı için 2'lik tabanda 10000000 sayısıdır. 8 bitlik bir sayı için 2'lik tabanda 10000000 sayısı -128 sayısına karşılık gelir.
		// type conversion yaparken büyük olan veri tipi küçük olan veri tipine dönüştürülmesi önerilmez

		fmt.Println(y)

	*/

	x := 10
	y := "10"

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)

	// fmt.Println(x + int(y)) // cannot convert y (type string) to type int

	// convert y to int with strconv.Atoi

	funcAtoi, _ := strconv.Atoi(y)
	fmt.Println(x + funcAtoi)

	num1 := 106

	str1 := string(rune(num1))

	fmt.Printf("%v %T\n", str1, str1)
}

/*
	// how to create public github repo from local repo in terminal
	1 - git init
	2 - git add .
	3 - git commit -m "first commit"
	4 - git remote add origin
	5 - git push -u origin master
*/
