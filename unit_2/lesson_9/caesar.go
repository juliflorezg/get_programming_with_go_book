// Decipher the quote from Julius Caesar:
// L fdph, L vdz, L frqtxhuhg

// Your program will need to shift uppercase and lowercase letters by –3. Remember that
// 'a' becomes 'x', 'b' becomes 'y', and 'c' becomes 'z', and likewise for uppercase letters.

package main

import "fmt"

func caesar() {
	fmt.Println("Caesar experiment")
	message := "L fdph, L vdz, L frqtxhuhg"

	for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
			c = c - 3
			if c < 'a' || c < 'A' {
				c = c + 26
			}
		}
		fmt.Printf("%c", c)
	}

}
