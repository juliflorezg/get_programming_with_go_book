package main

import (
	"fmt"
	"net/url"
)

var urlStr = "https://pk g.go.dev/net/url#Parse"

func parseUrl() {
	result, err := url.Parse(urlStr)

	if err != nil {
		fmt.Printf("%#v\n", err)

		if errs, ok := err.(*url.Error); ok {
			// for k, e := range *errs {
			// }
			// values := reflect.ValueOf(errs)
			// types := values.Type()
			// for i := 0; i < values.NumField(); i++ {
			//   fmt.Println(types.Field(i).Index[0])
			// }
			fmt.Println("Err:", errs.Err)
			fmt.Println("Op:", errs.Op)
			fmt.Println("url:", errs.URL)
		}
	}

	fmt.Println("result", result)
}
