package main

func main() {
	// sequential - conditional	- iterative
	// for init; condition; post {}
	// for condition {}
	// for {}
	// for key, value := range collection {}

	/*
		for <init statement>; <condition>; <post statement> {
			// body of for
		}
	*/

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	/*	Without shadowing
		i := 0

		for i < 5 {
			fmt.Println(i)
			i++
		}

		for ; i < 5; i++ {
			fmt.Println(i)
		}

		fmt.Println("done: ", i)
	*/

	/*	Break and Continue
		for {
			fmt.Println("loop")
			break
		}

		for n := 0; n <= 20; n++ {
			if n%3 == 0 {
				continue
			}
			fmt.Println(n)
		}

		for i := 0; i < 50; i++ {
			if i == 33 {
				break
			}

			if i%2 == 0 {
				fmt.Println(i, "is even")
			} else {
				fmt.Println(i, "is odd")
			}
		}
	*/
}
