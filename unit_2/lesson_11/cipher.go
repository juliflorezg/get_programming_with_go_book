package main

import (
	"fmt"
	"strings"
)

func cipher() {
	cipherText := "we dig you luv the gophers"
	cipherText = strings.ToUpper(strings.ReplaceAll(cipherText, " ", ""))
	keyword := "GOLANG"
	var cipheredText string

	keywordPointer := 0

	for _, r := range cipherText {
		if keywordPointer == len(keyword) {
			keywordPointer = 0
		}
		var sum, position int
		var char string

		// fmt.Println(i, string(r), string(keyword[keywordPointer]))
		// fmt.Println(alphabetMap[string(r)], "-", alphabetMap[string(keyword[keywordPointer])])
		//* C + G -> 2 + 6
		sum = AlphabetMap[string(r)] + AlphabetMap[string(keyword[keywordPointer])]

		// fmt.Println("subtraction", rest)
		if sum > 26 {
			// 22
			position = sum - 26
		} else {
			// fmt.Println(AlphabetMap[string(r)], rest)
			position = AlphabetMap[string(r)] + AlphabetMap[string(keyword[keywordPointer])]
		}
		// fmt.Println("position", position)

		for k, v := range AlphabetMap {
			if v == position {
				char = k
				cipheredText += char
			}
		}

		keywordPointer++
	}

	fmt.Println(cipheredText)

}
