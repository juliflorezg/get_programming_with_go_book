package main

import (
	"fmt"
	"math/big"
)

func main() {
	const lightSpeed = 299792 // km/s
	const secondsPerDay = 86400

	var distance int64 = 41.3e12
	fmt.Println("Alpha Centauri is ", distance, " km away.")
	days := distance / lightSpeed / secondsPerDay
	fmt.Println("That is", days, "days of travel at light speed.")

	{
		// var distance uint64 = 24e18 //! too big/truncated
	}

	//using big.int
	{
		lightSpeed := big.NewInt(299792)
		secondsPerDay := big.NewInt(86400)
		distance := new(big.Int)
		distance.SetString("24000000000000000000", 10)
		fmt.Println("Andromeda Galaxy is ", distance, " km away.")
		seconds := new(big.Int)
		seconds.Div(distance, lightSpeed)

		days := new(big.Int)
		days.Div(seconds, secondsPerDay)
		fmt.Println("That is", days, "days of travel at light speed.")

	}

	// constants can be untyped
	{
		// no error, untyped constant
		const distance = 24000000000000000000
		fmt.Println("Andromeda Galaxy is", 24000000000000000000/299792/86400, "light days away")
	}

	// we can use unusually big numbers with const
	{
		const distance = 24000000000000000000
		const lightSpeed = 299792
		const secondsPerDay = 86400
		const days = distance / lightSpeed / secondsPerDay
		fmt.Println("Andromeda Galaxy is", days, "light days away.")
		// fmt.Println(distance) //! overflows
	}
}
