package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	/*
		underArray := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		mySlc := underArray[2:5]
		fmt.Println(mySlc)
		mySlc2 := underArray[2:5:7]
		fmt.Println(mySlc2)

		mySlc = append(underArray[:], 10, 11, 12)
		fmt.Println(mySlc)
	*/
	/*
		mySlc := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		myslc2 := []int{10, 11, 12}

		mySlc = append(mySlc, myslc2...)
		fmt.Println(mySlc)
	*/
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	fmt.Printf("b is %s\n", b)
	var f interface{}
	err := json.Unmarshal(b, &f)

	if err != nil {
		fmt.Println(err)
	}

	f = map[string]interface{}{
		"Name": "Wednesday",
		"Age":  6,
		"Parents": []interface{}{
			"Gomez",
			"Morticia",
		},
		"Children": 2,
		"bool":     true,
	}

	m := f.(map[string]interface{})

	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		// if its int
		case int:
			fmt.Println(k, "is int", vv)

		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

	/* 	underArray := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	   	fmt.Println(underArray)

	   	mySlc := underArray[2:5]

	   	mySlc2 := underArray[:6]
	   	fmt.Println(mySlc2)
	   	mySlc3 := underArray[3:]
	   	fmt.Println(mySlc3)
	   	mySlc4 := underArray[:]
	   	fmt.Println(mySlc4)

	   	fmt.Println(mySlc)
	   	mySlc[0] = 100
	   	fmt.Println(mySlc)
	   	fmt.Println(mySlc2)
		   fmt.Println(mySlc4) */

	/* 	mySlc := []int{1, 2, 3} // mySliceUnderArray
	   	fmt.Println(mySlc) */

	/* 	mySlc = append(mySlc, 4, 5)
	   	fmt.Println(mySlc) */

	/* 	mySlc2 := append(mySlc, 4, 5) // mySlice2UnderArray
	   	fmt.Println(mySlc2)

	   	mySlc[0] = 100
	   	fmt.Println(mySlc)
	   	fmt.Println(mySlc2) */

	/* 	mySlc := []int{1, 2, 3}
	   	mySlc2 := []int{4, 5}

	   	mySlc = append(mySlc, mySlc2...)

	   	fmt.Println(mySlc) */

	/* 	mySlc := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	   	fmt.Println(mySlc)
	   	// mySlc = mySlc[2:]  ilk n elemanı silmek [n:]

	   	// mySlc = mySlc[:len(mySlc)-3] son n elamanı silmek [:len(mySlc)-n]
	   	mySlc = mySlc[2:]
	   	mySlc = mySlc[:len(mySlc)-3]

	   	fmt.Println(mySlc) */

	/* 	var myArr [4]int
	   	fmt.Println(myArr)

	   	var mySlc []int
	   	mySlc = make([]int, 4) // Zero değerler Slice Elemanlarına
	   	fmt.Println(mySlc)

	   	var mySlc2 []bool
	   	mySlc2 = make([]bool, 2) // Zero değerler Slice Elemanlarına
	   	fmt.Println(mySlc2)
	*/
	/*
		var mySlc3 []int
		fmt.Printf("%#v", mySlc3)

		fmt.Println()

		mySlc4 := make([]int, 0)
		fmt.Printf("%#v", mySlc4)
	*/
}
