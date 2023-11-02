// Write a guess-the-number program. Make the computer pick random numbers between 1–100 until it guesses your number, which you declare at the top of the program. Display each guess and whether it was too big or too small.

package main

import (
	"fmt"
	"math/rand"
)

func guess() {
	var number int
	fmt.Println("Please give a random number between 1 and 100")

	fmt.Scanln(&number)

	for {
		var guess = rand.Intn(100) + 1
		if guess > number {
			fmt.Println("My guess is greater than your number::", guess)
		} else if guess < number {
			fmt.Println("My guess is lower than your number::", guess)
		} else if guess == number {
			fmt.Println("I guessed the number, its", number)
			break
		}
	}
}
