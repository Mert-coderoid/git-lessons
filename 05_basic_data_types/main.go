/*
--- BASIC DATA TYPES ---
NUMERIC TYPES
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

COMPLEX TYPES
complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

OTHER TYPES
byte        alias for uint8
rune        alias for int32

STRING TYPES
string      the set of all strings of 8-bit bytes (not to be confused with UTF-8)

( A string type represents the set of string values. A string value is a (possibly empty) sequence of bytes.
Strings are immutable: once created, it is impossible to change the contents of a string. The predeclared string type is string.)

BOOLEAN TYPES
bool        the set of boolean values, true and false

COMPOSITE TYPES
array       a numbered sequence of elements of a single type
struct      a collection of fields
pointer     a pointer to another type
function    an executable code block
interface   a set of method signatures
slice       a segment of an array
map         an unordered group of elements of one type, indexed by a set of unique keys of another type
channel     a communication mechanism with a type

------------------- TURKCE -------------------

--- BASIT VERİ TİPLERİ ---
NUMERIC TİPLER
uint8       uint8       tüm 8 bitlik tamsayıların kümesi (0 ile 255 arasında)
uint16      uint16      tüm 16 bitlik tamsayıların kümesi (0 ile 65535 arasında)
uint32      uint32      tüm 32 bitlik tamsayıların kümesi (0 ile 4294967295 arasında)
uint64      uint64      tüm 64 bitlik tamsayıların kümesi (0 ile 18446744073709551615 arasında)

int8        int8        tüm 8 bitlik tamsayıların kümesi (-128 ile 127 arasında)
int16       int16       tüm 16 bitlik tamsayıların kümesi (-32768 ile 32767 arasında)
int32       int32       tüm 32 bitlik tamsayıların kümesi (-2147483648 ile 2147483647 arasında)
int64       int64       tüm 64 bitlik tamsayıların kümesi (-9223372036854775808 ile 9223372036854775807 arasında)

float32     float32     tüm IEEE-754 32 bitlik ondalık sayıların kümesi
float64     float64     tüm IEEE-754 64 bitlik ondalık sayıların kümesi

COMPLEX TİPLER
complex64   complex64   tüm 32 bitlik gerçek ve sanal kısımlı karmaşık sayıların kümesi
complex128  complex128  tüm 64 bitlik gerçek ve sanal kısımlı karmaşık sayıların kümesi

DİĞER TİPLER
byte        byte        uint8 türü için takma ad
rune        rune        int32 türü için takma ad

STRING TİPLER
string      string      tüm 8 bitlik bayt dizilerinin kümesi (UTF-8 ile karıştırılmamalıdır)

( String türü, string değerlerin kümesini temsil eder. Bir string değeri, bir bayt dizisidir.
Stringler değiştirilemez: oluşturulduktan sonra, bir stringin içeriğini değiştirmek mümkün değildir. Önceden tanımlanmış string türü string'dir.)

BOOLEAN TİPLER
bool        bool        true ve false değerlerinin kümesi

COMPOSITE TİPLER
array       array       tek bir türün elemanlarının numaralandırılmış dizisi
struct      struct      alanların bir koleksiyonu
pointer     pointer     başka bir türe gösteren bir işaretçi
function    function    çalıştırılabilir kod bloğu
interface   interface   bir dizi yöntem imzası
slice       slice       bir dizi segmenti
map         map         bir türün elemanlarının, başka bir türün benzersiz anahtarları ile indekslenmiş olması
channel     channel     bir türün iletişim mekanizması

*/

package main

import (
	"fmt"
)

func main() {

	name, lastname := "Mert", "Yılmaz"

	var age uint8
	age = 255

	var isCool bool
	isCool = true

	// weight := 70.5 	// deafult float64
	var weight float32 = 70454545454545454545457045454704545454.5
	var weight2 float64 = 70454545454545454545457045454704545454704545454545454545454570454547045454547045454545454545454545704545470454545470454545454545454545457045454704545454704545454545454545454570454547045454547045454545454545454545704545470454545470454545454545454545457045454704545454704545454545454545454570454547045454544442.5

	fmt.Println("Hello,", name, lastname)
	fmt.Println("You are", age, "years old")
	fmt.Println("Is he cool?", isCool)

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", weight)
	fmt.Printf("%T\n", weight2)
}
