package main

import (
	"fmt"
	"strings"
)

func sourceGopher(downstream chan string) {
	for _, v := range []string{"Hello World!", "a bad apple", "goodbye all"} {
		downstream <- v
	}

	downstream <- "" // sentinel value, this indicates the downstream pipelines that the work is done
}

func filterGopher(upstream, downstream chan string) {

	for {
		val := <-upstream
		if val == "" {
			downstream <- ""
			return
		}
		if !strings.Contains(val, "bad") {
			downstream <- val
		}

	}

}

func printGopher (upstream chan string) {
  for {
    val := <- upstream
    if val == "" {
      return
    }

    fmt.Println(val)
  }
}
