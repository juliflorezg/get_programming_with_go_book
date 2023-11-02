/*
Modify listing 4.7 to handle leap years:

	 Generate a random year instead of always using 2018.
	 For February, assign daysInMonth to 29 for leap years and 28 for other years.
	  Hint: you can put an if statement inside of a case block.
	 Use a for loop to generate and display 10 random dates.
*/
package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func generateRandomDates() {

	for i := 0; i < 10; i++ {

		year := rand.Intn(2100) + 1

		var isLeap = year%400 == 0 || (year%4 == 0 && year%100 != 0)

		month := rand.Intn(12) + 1
		daysInMonth := 31

		switch month {
		case 2:
			daysInMonth = 28
			if isLeap {
				daysInMonth = 29
			}
		case 4, 6, 9, 11:
			daysInMonth = 30
		}

		day := rand.Intn(daysInMonth) + 1

		fmt.Println(era, year, month, day)
	}
}
