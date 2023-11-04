//Write a program that converts strings to Booleans:
//  The strings “true”, “yes”, or “1” are true.
//  The strings “false”, “no”, or “0” are false.
//  Display an error message for any other values.

package main

import "fmt"

func input() bool {
	var str string
	var result bool
	fmt.Scanln(&str)

	switch str {
	case "true", "yes", "1":
		result = true
	case "false", "no", "0":
		result = false
	default:
		fmt.Println("not a valid input")
	}

	return result

}
