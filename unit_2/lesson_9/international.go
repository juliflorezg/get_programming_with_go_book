package main

import "fmt"

func international() {
	fmt.Println("international")
	message := "Hola Estación Espacial Internacional”"

	// r will be the code point (rune) hence the name r
	for _, r := range message {
		// fmt.Println("original rune:", r)
		// if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' {
		// 	r += 13
		// 	if r > 'z' || r > 'Z' {
		// 		r -= 26
		// 	}
		// }
		if r >= 'a' && r <= 'z' {
			r += 13
			if r > 'z' {
				r -= 26
			}
		} else if r >= 'A' && r <= 'Z' {
			r += 13
			if r > 'Z' {
				r -= 26
			}

		}

		// fmt.Println("ciphered:", r)
		fmt.Printf("%c", r)
	}
}
