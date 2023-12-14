package main

import (
	"fmt"
	"strings"
)

func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello hello mom!", "a bad bad apple apple", "goodbye all all"} {
		downstream <- v
	}

	close(downstream)

	// downstream <- "" // sentinel value, this indicates the downstream pipelines that the work is done
}

func filterGopher(upstream, downstream chan string) {

	// for {
	// 	val, ok := <-upstream
	// 	if !ok {
	// 		close(downstream)
	// 		return
	// 	}
	// 	if !strings.Contains(val, "bad") {
	// 		downstream <- val
	// 	}
	// }
	// the previous code can be rewritten like this

	for val := range upstream {
		if !strings.Contains(val, "bad") {
			downstream <- val
		}
	}
	close(downstream)

}

func printGopher(upstream chan string) {
	for val := range upstream {
		fmt.Println(val)
	}
}
