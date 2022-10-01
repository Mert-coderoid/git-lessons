package main

/* #### NOTES ####

printf --> formatlı yazdırma içerisinde
	%d --> decimal (sayısal) manasına gelir

Fonksiyon İsimlendirme
	ilk karekter harf
	camel Case -- mySum, myBestFunction
	paket dışı --> ilk harf büyük
*/
import "fmt"

func main() {

	/*
		// Moduler Programming
		// Readable code
		// From Complex To Simple

		fmt.Println(sum(5, 10))

		// func funcName(params) return type { code  }

		/* 	merhaba()

		   	fmt.Println(sum(5, 10))
		   	fmt.Println(sum(3, 5))
		   	fmt.Println(sum(2, 7))

		   	merhaba()
			   merhaba()
	*/

	/*
		var x, y, sum int

		   	x = 5
		   	y = 10
		   	sum = x + y
		   	fmt.Printf("%d ve %d toplamı %d\n", x, y, sum)

		   	x = 7
		   	y = 11
		   	sum = x + y
		   	fmt.Printf("%d ve %d toplamı %d\n", x, y, sum)
	*/

	// fmt.Println(x)

	// return vs print

	/* 	z := sum(5, 10)
	   	fmt.Println(z)

		   sum2(6, 11) */

	text := sayMyName("Elis", 4.2)
	fmt.Println(text)
}

func sum(x, y int) int {
	return x + y
}

func sum2(x, y int) {
	fmt.Println(x + y)
}

func merhaba() {
	fmt.Println("Benim Adım Arin")
}

func sayMyName(name string, age float64) string {
	Converted_age := fmt.Sprintf("%v", age) // int to string - output: "4"
	text := "Benim Adım " + name + " ve " + Converted_age + " yaşındayım"
	return text
}
