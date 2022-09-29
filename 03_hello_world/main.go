// package clause
package main

// import statement - 	fmt is a package in the Go standard library that implements formatted I/O with functions analogous to C's printf and scanf
// tr;	import statement - 	fmt Go standart kütüphanesindeki bir pakettir. C'nin printf ve scanf fonksiyonlarına benzer biçimlendirilmiş I/O işlevlerini uygular.
// I/O işlevleri 	- 	Input/Output işlevleri, veri giriş çıkış işlevleri, veri okuma yazma işlevleri, veri işleme işlevleri
import (
	"fmt"
	"time"
)

// main function	- 	Every Go program has a main function that is the entry point of the program and is executed by default when you run the program.
// tr: 	main fonksiyonu	- 	Her Go programı bir main fonksiyonuna sahiptir. Bu fonksiyon programın giriş noktasıdır ve programı çalıştırdığınızda varsayılan olarak çalıştırılır.
func main() {
	fmt.Println("Hello, World!")
	time.Sleep(15 * time.Second)
}
