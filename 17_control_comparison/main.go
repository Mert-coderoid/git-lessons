// Golang program to convert Int data type to Float
package main

import "fmt"

// ToFloat32 converts a int num to a float32 num
func ToFloat32(in int) float32 {
	return float32(in)
}

// ToFloat64 converts a int num to a float64 num
func ToFloat64(in int) float64 {
	return float64(in)
}

func main() {

	var x int64 = 5
	x = int64(ToFloat64(int(x)))

	fmt.Printf("Type: %T, Value: %v", x+5.5, x)

	// r := float64(x)

	// fmt.Printf("r is of type %T and value is %v", r, r)
	fmt.Println("")
	// fmt.Println(y + r + 1.5)
}
