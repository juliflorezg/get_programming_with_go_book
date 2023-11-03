/*
Save some money to buy a gift for your friend. Write a program that randomly places
nickels ($0.05), dimes ($0.10), and quarters ($0.25) into an empty piggy bank until it contains at
least $20.00. Display the running balance of the piggy bank after each deposit,
formatting it with an appropriate width and precision
*/
package main

import (
	"fmt"
	"math/rand"
)

var coinDenomination = []float64{0.05, 0.10, 0.25}

func piggy() {
	piggyBank := 0.0

	for piggyBank < 20.00 {
		randomCoin := coinDenomination[rand.Intn(3)]
		piggyBank += randomCoin
		fmt.Println("coin to add:", randomCoin)
		fmt.Print("current piggyBank balance:::  ", "$")
		fmt.Printf("%5.2f\n", piggyBank)
	}

	fmt.Printf("$%5.2f", piggyBank)
}
