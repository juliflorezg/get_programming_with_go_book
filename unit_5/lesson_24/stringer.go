package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type coordinate struct {
	d float64
	m float64
	s float64
	h rune
}

type formattedCoordinate struct {
	Decimal float64 `json:"decimal"`
	Dms     string  `json:"dms"`
	D       float64 `json:"degrees"`
	M       float64 `json:"minutes"`
	S       float64 `json:"seconds"`
	H       string  `json:"hemisphere"`
}

// Write a String method on the coordinate type and use it to display coordinates in a more readable format.

// Your program should output: Elysium Planitia is at 4º30'0.0" N, 135º54'0.0" E

func (c coordinate) String() string {
	s := fmt.Sprintf("%v°%v'%v\" %c", c.d, c.m, c.s, c.h)
	return s
}

// decimal returns a decimal coordinate for a dms coordinate
func (c coordinate) decimal() float64 {
	sign := 1.0

	switch c.h {
	case 'S', 's', 'W', 'w':
		sign = -1
	}

	return sign * (c.d + c.m/60 + c.s/3600)
}

type location struct {
	Name string     `json:"name"`
	Lat  coordinate `json:"latitude"`
	Long coordinate `json:"longitude"`
}

func (l location) String() string {
	// we can use l.lat & l.long in Sprintf since they are of type coordinate and that type satisfies the Stringer interface bc of String method
	s := fmt.Sprintf("%v, %v", l.Lat, l.Long)
	return s
}

// this function makes coordinate satisfy the Marshaler interface :
//
//	type Marshaler interface {
//		MarshalJSON() ([]byte, error)
//	}
func (c coordinate) MarshalJSON() ([]byte, error) {
	//formatted coordinate
	fc := formattedCoordinate{
		Decimal: c.decimal(),
		Dms:     c.String(),
		D:       c.d,
		M:       c.m,
		S:       c.s,
		H:       string(c.h),
	}

	return json.Marshal(fc)
	//the code below is what generates a stack overflow error, it calls this function over and over as we can see thanks to the print statement
	// fmt.Println("dentro del marshal")

	// return json.Marshal(c)
}

// * MarshalJSON could be written as the following (defining struct inline)::
//* func (c coordinate) MarshalJSON() ([]byte, error) {

//* 	return json.Marshal(struct {
//* 		D       float64 `json:"degrees"`
//* 		M       float64 `json:"minutes"`
//* 		S       float64 `json:"seconds"`
//* 		H       string  `json:"hemisphere"`
//* 		Decimal float64
//* 		Dms     string
//* 	}{
//* 		D:       c.D,
//* 		M:       c.M,
//* 		S:       c.S,
//* 		H:       string(c.H),
// * 		Dms:     c.String(),
// * 		Decimal: c.decimal(),
// * 	})

// * }

func printLocationMessage() {
	elysium := location{
		Name: "Elysium Planitia",
		Lat:  coordinate{4, 30, 0.0, 'N'},
		Long: coordinate{135, 54, 0.0, 'E'},
	}

	// same as with coordinate, we can use l in Println bc location type satisfies Stringer interface (thanks to String method)
	fmt.Println("Elysium Planitia is at", elysium)

	c := coordinate{135, 54, 0.0, 'E'}

	// bytes, err := json.Marshal(c)
	// if err != nil {
	//   fmt.Println(err)
	//   os.Exit(1)
	// }
	// fmt.Println(string(bytes))

	// bytes, err := json.Marshal(elysium)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Println(string(bytes))

	bytes, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(bytes)) // {"decimal":135.9,"dms":"135°54'0\" E","degrees":135,"minutes":54,"seconds":0,"hemisphere":"E"}

	bytes2, err2 := json.MarshalIndent(elysium, "", " ")
	if err2 != nil {
		fmt.Println(err2)
		os.Exit(1)
	}
	fmt.Println(string(bytes2))
	// 	{
	//  "name": "Elysium Planitia",
	//  "latitude": {
	//   "decimal": 4.5,
	//   "dms": "4°30'0\" N",
	//   "degrees": 4,
	//   "minutes": 30,
	//   "seconds": 0,
	//   "hemisphere": "N"
	//  },
	//  "longitude": {
	//   "decimal": 135.9,
	//   "dms": "135°54'0\" E",
	//   "degrees": 135,
	//   "minutes": 54,
	//   "seconds": 0,
	//   "hemisphere": "E"
	//  }
	// }

}
