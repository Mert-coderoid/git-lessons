/* package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Print("Lütfen Sınav Sonucunuzu Giriniz: ")
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n') // _ blank identifier
	fmt.Println(value)
} */

/* package main

import "fmt"

func main() {

	fmt.Println("Arin")
} */
/*
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("seafood", "foo"))

	fmt.Println(strings.Count("animatrix", "a"))
}
*/

/* package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ToUpper("Gopher"))
} */

/* Main Package

fmt methods

	Println() -> print line
	Print() -> print
	Printf() -> print format
	Sprint() -> string print
	Sprintf() -> string print format
	Sprintln() -> string print line
	Fprint() -> file print
	Fprintf() -> file print format
	Fprintln() -> file print line


string methods

	Contains() -> string içerisinde aranan değer var mı?
	Count() -> string içerisinde aranan değer kaç defa geçiyor?
	HasPrefix() -> string hangi değer ile başlıyor?
	HasSuffix() -> string hangi değer ile bitiyor?
	Index() -> string içerisinde aranan değer hangi index numarasında?
	Join() -> string içerisindeki değerleri birleştirme
	Repeat() -> string içerisindeki değeri tekrar etme
	Replace() -> string içerisindeki değeri değiştirme
	ReplaceAll() -> string içerisindeki değerleri değiştirme
	Split() -> string içerisindeki değerleri ayırma
	ToLower() -> string içerisindeki değerleri küçük harfe çevirme
	ToUpper() -> string içerisindeki değerleri büyük harfe çevirme
	Trim() -> string içerisindeki değerleri silme
	TrimLeft() -> string içerisindeki değerleri soldan silme
	TrimRight() -> string içerisindeki değerleri sağdan silme
	TrimSpace() -> string içerisindeki değerleri boşlukları silme
	Fields() -> string içerisindeki değerleri boşluklarına göre ayırma


os methods

	Args -> komut satırından girilen değerleri almak için kullanılır
	Exit -> programı sonlandırmak için kullanılır
	Getenv -> sistemde tanımlı olan değişkenleri almak için kullanılır
	Stdin -> kullanıcıdan veri almak için kullanılır
	Stdout -> kullanıcıya veri göndermek için kullanılır
	Stderr -> kullanıcıya hata mesajı göndermek için kullanılır
	Opendir -> dizin açmak için kullanılır
	Readdir -> dizin içerisindeki dosyaları okumak için kullanılır
	Stat -> dosya veya dizin bilgilerini okumak için kullanılır
	Remove -> dosya veya dizin silmek için kullanılır
	Rename -> dosya veya dizin ismini değiştirmek için kullanılır
	Chmod -> dosya veya dizin izinlerini değiştirmek için kullanılır
	Chown -> dosya veya dizin sahibini değiştirmek için kullanılır
	Chdir -> dizin değiştirmek için kullanılır
	Getwd -> çalışılan dizini okumak için kullanılır
	Mkdir -> dizin oluşturmak için kullanılır
	MkdirAll -> dizin oluşturmak için kullanılır
	RemoveAll -> dizin silmek için kullanılır
	Truncate -> dosya boyutunu değiştirmek için kullanılır
	Open -> dosya açmak için kullanılır
	Create -> dosya oluşturmak için kullanılır
	NewFile -> dosya oluşturmak için kullanılır
	OpenFile -> dosya açmak için kullanılır


bufio methods

	NewReader -> okuma işlemi için kullanılır
	ReadString -> okuma işlemi için kullanılır
	ReadBytes -> okuma işlemi için kullanılır
	ReadRune -> okuma işlemi için kullanılır
	ReadLine -> okuma işlemi için kullanılır
	ReadSlice -> okuma işlemi için kullanılır
	Read -> okuma işlemi için kullanılır
	NewWriter -> yazma işlemi için kullanılır
	WriteString -> yazma işlemi için kullanılır
	WriteRune -> yazma işlemi için kullanılır
	Write -> yazma işlemi için kullanılır
	Flush -> yazma işlemi için kullanılır
	ReadFrom -> okuma işlemi için kullanılır
	WriteTo -> yazma işlemi için kullanılır
	Reset -> okuma yazma işlemi için kullanılır
	Scan -> okuma yazma işlemi için kullanılır
	Scanln -> okuma yazma işlemi için kullanılır
	Scanf -> okuma yazma işlemi için kullanılır
	ScanBytes -> okuma yazma işlemi için kullanılır
	ScanRunes -> okuma yazma işlemi için kullanılır
	ScanWords -> okuma yazma işlemi için kullanılır
	ScanLines -> okuma yazma işlemi için kullanılır
	ScanComma -> okuma yazma işlemi için kullanılır
	ScanSpace -> okuma yazma işlemi için kullanılır
	ScanText -> okuma yazma işlemi için kullanılır
	ScanUntil -> okuma yazma işlemi için kullanılır
	ScanInt -> okuma yazma işlemi için kullanılır
	ScanUint -> okuma yazma işlemi için kullanılır
	ScanFloat -> okuma yazma işlemi için kullanılır
	ScanBool -> okuma yazma işlemi için kullanılır



*/

package main

func main() {
	// fmt.Println(strings.Contains("seafood", "foo"))
	// fmt.Println(strings.Contains("seafood", "bar"))
	// fmt.Println(strings.Contains("seafood", ""))
	// fmt.Println(strings.Contains("", ""))

	// // create an array
	// list1 := []string{"foo", "bar", "baz"}

	// // create list in list
	// list2 := [][]string{
	// 	{"foo", "bar", "baz"},
	// 	{"foo", "bar", "baz"},
	// 	{"foo", "bar", "baz"},
	// }

	// // create map
	// list3 := map[string]string{
	// 	"foo": "bar",
	// 	"bar": "baz",
	// 	"baz": "foo",
	// }

	// fmt.Println(list1)
	// fmt.Println(list2)
	// fmt.Println(list3)

	// create a array
	list1 := []any{"foo", 2, true, 3.14}

}
