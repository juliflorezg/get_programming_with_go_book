package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var spirit = location{coordinate{14, 34, 6.2, 'S'}.decimal(), coordinate{175, 28, 21.5, 'E'}.decimal()}
var opportunity = location{coordinate{1, 56, 46.3, 'S'}.decimal(), coordinate{354, 28, 24.2, 'E'}.decimal()}
var curiosity = location{coordinate{4, 35, 22.2, 'S'}.decimal(), coordinate{137, 26, 30.1, 'E'}.decimal()}
var inSight = location{coordinate{4, 30, 0.0, 'N'}.decimal(), coordinate{135, 54, 0, 'E'}.decimal()}

func getLocations() {

	fmt.Println("location of Spirit rover landing:", spirit)
	fmt.Println("location of Opportunity rover landing:", opportunity)
	fmt.Println("location of Curiosity rover landing:", curiosity)
	fmt.Println("location of InSight rover landing:", inSight)

}

func getDistances() {
	mars := world{radius: 3389.5}
	marsLocations := map[string]location{
		"spirit":      spirit,
		"opportunity": opportunity,
		"curiosity":   curiosity,
		"insight":     inSight,
	}

	// i.e: {"spirit opportunity": 34123.435435}
	locationPairs := make(map[string]float64, (len(marsLocations)*(len(marsLocations)-1))/2) //n(n-1)/2

	for k := range marsLocations {  
		for k1 := range marsLocations {
      // if locations are the same, skip
			if k == k1 {
        continue
			}
      // if location is the same than other one, but in opposite direction, skip (AB == BA)
			if locationPairs[k+" "+k1] != 0 || locationPairs[k1+" "+k] != 0 {
				continue
			}
			//* we could also use the comma, ok idiom, but that leads to two separate conditions:::
			// if _, ok := locationPairs[k+" "+k1]; ok {
			// 	continue
			// }
			// if _, ok := locationPairs[k1+" "+k]; ok {
			// 	continue
			// }
			locationPairs[k+" "+k1] = mars.distance(marsLocations[k], marsLocations[k1])
		}
	}

	fmt.Println(locationPairs)
	// fmt.Println(string(json.MarshalIndent(locationPairs, "", "  ")))

	bytes, err := json.MarshalIndent(locationPairs, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(bytes)
	fmt.Println(string(bytes))

}
