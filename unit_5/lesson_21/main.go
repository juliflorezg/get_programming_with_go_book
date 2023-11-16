package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type location struct {
	lat  float64
	long float64
}

func main() {
	var curiosity struct {
		lat  float64
		long float64
	}

	//assigning values to structure fields
	curiosity.lat = -4.5895
	curiosity.long = 137.4417

	fmt.Println(curiosity.lat, curiosity.long)
	fmt.Println(curiosity)

	var spirit location
	spirit.lat = -14.5684
	spirit.long = 175.472636

	var opportunity location
	opportunity.lat = -1.9462
	opportunity.long = 354.4734

	fmt.Println(spirit, opportunity)

	// declare structures with composite literals

	{
		opportunity := location{lat: -1.9462, long: 354.4734}
		insight := location{lat: 4.5, long: 135.9}
		fmt.Println(opportunity, insight)
		// use struct literal without the field names, values must be in order and a value should be provided for every field in the struct
		spirit := location{-14.5684, 175.472636}
		// we can use the + with format verb %v to show field names on screen
		fmt.Printf("%+v\n", spirit)
	}

	// structs are copied, change to one won't affect the other ones
	{
		bradbury := location{-4.5895, 137.4417}
		curiosity := bradbury

		curiosity.long += 0.0106
		fmt.Println(bradbury)
		fmt.Println(curiosity)
	}

	// we an use structs as values of a slice (a slice of structs)
	{
		type location struct {
			name string
			lat  float64
			long float64
		}

		locations := []location{
			{name: "Bradbury Landing", lat: -4.5895, long: 137.4417},
			{name: "Columbia Memorial Station", lat: -14.5684, long: 175.472636},
			{name: "Challenger Memorial Station", lat: -1.9462, long: 354.4734},
		}
		fmt.Println(locations)
	}

	//? encoding to json format
	{
		type location struct {
			//* fields MUST be Exported when working with json format, so it can recognize them
			Lat, Long float64
		}
		curiosity := location{-4.5895, 137.4417}

		// marshal function converts the value to json format
		bytes, err := json.Marshal(curiosity)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(bytes)
		fmt.Println(string(bytes))
	}

	// we can customize location fields for json
	{
		{
			type location struct {
				//* fields MUST be Exported when working with json format, so it can recognize them
				Lat  float64 `json:"latitude"`
				Long float64 `json:"longitude"`
			}
			curiosity := location{-4.5895, 137.4417}

			// marshal function converts the value to json format
			bytes, err := json.Marshal(curiosity)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(bytes)
			fmt.Println(string(bytes))
		}
	}

  landing()

}
