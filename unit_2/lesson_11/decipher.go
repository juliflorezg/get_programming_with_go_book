// Write a program to decipher the ciphered text shown in table 11.2. To keep it simple, all
// characters are uppercase English letters for both the text and keyword:
// cipherText := "CSOITEUIWUIZNSROCNKFD"
// keyword := "GOLANG"

// A B C D E F G H I J K  L  M  N  O  P  Q  R  S  T  U  V  W  X  Y  Z
// 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25

// ABCDEFGHIJKLMNOPQRSTUVWYZ
// GOLANGGOLANGGOLANG

package main

import "fmt"

var AlphabetMap = map[string]int{
	"A": 0, "B": 1, "C": 2, "D": 3, "E": 4, "F": 5, "G": 6, "H": 7, "I": 8, "J": 9, "K": 10, "L": 11, "M": 12, "N": 13, "O": 14, "P": 15, "Q": 16, "R": 17, "S": 18, "T": 19, "U": 20, "V": 21, "W": 22, "X": 23, "Y": 24, "Z": 25,
}

func decipher() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	var result string

	keywordPointer := 0

	for _, r := range cipherText {
		if keywordPointer == len(keyword) {
			keywordPointer = 0
		}
		var rest, position int
		var char string

		// fmt.Println(i, string(r), string(keyword[keywordPointer]))
		// fmt.Println(alphabetMap[string(r)], "-", alphabetMap[string(keyword[keywordPointer])])
		//* C - G -> 2 - 6
		rest = AlphabetMap[string(r)] - AlphabetMap[string(keyword[keywordPointer])]

		// fmt.Println("subtraction", rest)
		if rest < 0 {
			// 22
			position = 26 + rest
		} else {
			// fmt.Println(alphabetMap[string(r)], rest)
			position = AlphabetMap[string(r)] - AlphabetMap[string(keyword[keywordPointer])]
		}
		// fmt.Println("position", position)

		for k, v := range AlphabetMap {
			if v == position {
				char = k
				result += char
			}
		}

		keywordPointer++
	}

	fmt.Println(result)

}
