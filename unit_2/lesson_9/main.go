package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var blank string

	fmt.Println(blank) // will be the zero value for the string type, which is the empty string

	//how to avoid escaping in strings (use `` instead of "")
	fmt.Println("peace be upon you\nupon you be peace")
	fmt.Println(`strings can span multiple lines with the \n escape sequence`)
	fmt.Println(`
    peace be upon you
    upon you be peace`)

	// runes represent characters and a rune is an alias for int32

	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33

	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang)
	// use of the %c format verb to get the character representation for that value
	fmt.Printf("%c %c %c %c\n", pi, alpha, omega, bang)

	// to store the numeric value of a character

	var grade = 'A'
	fmt.Println(grade) // 65, rune for 'A'
	fmt.Println(rune('*'))
	fmt.Printf("%c %[1]v\n", 128515)
	fmt.Println(rune('é'))

	//indexing a string

	message := "shalom"
	c := message[5]
	fmt.Printf("%c %[1]v\n", c) // m 109

	//strings are inmutable
	// message[5] = "d" //! error, cannot assign to message[5]

	//> Write a program to print each byte (ASCII character) of "shalom", one character per line.

	for i := 0; i < len(message); i++ {
		fmt.Printf("%c %[1]v\n", message[i])
		// fmt.Println(message[i])
	}
	fmt.Println("rune for 'a'", rune('a'))
	fmt.Println("rune for 'b'", rune('b'))
	fmt.Println("rune for 'c'", rune('c'))
	fmt.Println("rune for 'x'", rune('x'))
	fmt.Println("rune for 'y'", rune('y'))
	fmt.Println("rune for 'z'", rune('z'))

	fmt.Println('g' - 'a' + 'A')
	fmt.Printf("%c", 'g'-'a'+'A')

	//> 1 How many runes are in the English alphabet "abcdefghijklmnopqrstuvwxyz"? How many bytes?
	// >2 How many bytes are in the rune '¿'?
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	fmt.Println(len(alphabet), "bytes in the english alphabet")
	fmt.Println(utf8.RuneCountInString(alphabet), "runes in the english alphabet")
	fmt.Println(len("¿"), "bytes in the ¿ character")

	caesar()
	international()

}
