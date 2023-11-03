/*
Write a new piggy bank program that uses integers to track the number of cents rather
than dollars. Randomly place nickels (5¢), dimes (10¢), and quarters (25¢) into an empty
piggy bank until it contains at least $20.
Display the running balance of the piggy bank after each deposit in dollars (for example, $1.05).
*/
package main

import (
	"fmt"
	"math/rand"
)

var denominationsInCents = []int{5, 10, 25}

func piggy() {
	piggyBank := 0

	for piggyBank < 2000 {
		randomCoin := denominationsInCents[rand.Intn(3)]
		piggyBank += randomCoin
		fmt.Println("coin to add:", randomCoin)
		fmt.Print("current piggyBank balance:::  ", "$")
		fmt.Printf("%v.%v\n", piggyBank/100, piggyBank%100)
	}

	fmt.Println(piggyBank)
	fmt.Printf("%v.%v\n", piggyBank/100, piggyBank%100)
}
