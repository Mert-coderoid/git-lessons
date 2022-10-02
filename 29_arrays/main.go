package main

import "fmt"

func main() {
	myArr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	cities := []string{"istanbul", "roma", "tahran", "belgrad", "moskova", "ankara", "london", "paris", "madrid", "berlin"}

	// myArr = mySquare(myArr)
	mySquareWithRange(cities)
	fmt.Println(myArr)

}

func mySquare(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
		arr[i] *= arr[i]
	}
	return arr
}

func mySquareWithRange(arr []string) []string {
	for index, city := range arr {
		fmt.Println(index, city)
	}
	return arr
	/*
		output:
			0 istanbul
			1 roma
			2 tahran
			3 belgrad
			4 moskova
			5 ankara
			6 london
			7 paris
			8 madrid
			9 berlin
	*/
}

/* 	city1 := "istanbul"
   	city2 := "roma"
   	city3 := "tahran"
   	city4 := "belgrad"

   	fmt.Println(city1, city2, city3, city4) */

//cities := [4]string{"istanbul", "roma", "tahran", "belgrad"}

/* 	cities := []string{"istanbul", "roma", "tahran", "belgrad"}

   	fmt.Println(cities)
   	fmt.Println(cities[0]) // zero based index
   	fmt.Println(cities[3])

   	fmt.Println(len(cities)) */

/* 	var myArr [5]int
   	fmt.Println(myArr)

   	//myArr[0] = "istanbul"
   	myArr[0] = 100
   	fmt.Println(myArr)
   	myArr[len(myArr)-1] = 200
   	fmt.Println(myArr) */

/* 	var myArr [4]int
   	var myArr2 [5]int */

/* 	if myArr == myArr2 {
	fmt.Println("MESAJ")
} */

/* 	fmt.Println(myArr)
   	fmt.Println(myArr2) */

/* 	cities := [4]string{"istanbul", "roma", "tahran", "belgrad"}

   	for i := 0; i < len(cities); i++ {
   		fmt.Println(i, cities[i])
   	}

   	cities[0] = "ANKARA"
   	fmt.Println()

   	for i := 0; i < len(cities); i++ {
   		fmt.Println(i, cities[i])
   	} */

/* 	myArr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

   	myArr = mySquare(myArr) // First Class Functions

   	fmt.Println(myArr) */

// FOR --- RANGE
// for index, value := range myArr

// cities := [4]string{"istanbul", "roma", "tahran", "belgrad"}

/* 	for index, city := range cities {
	fmt.Println(index, city)
} */

/* 	for _, city := range cities {
	fmt.Println(city)
}
*/

/* func mySquare(arr [10]int) [10]int {

	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
	}

	return arr
} */
