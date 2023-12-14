package main

import "fmt"

// It’s boring to see the same line repeated over and over again. Write a pipeline element (a goroutine) that remembers the previous value and only sends the value to the next stage of the pipeline if it’s different from the one that came before. To make things a little simpler, you may assume that the first value is never the empty string.

func source(downStream chan string) {
	for _, v := range []string{"first", "first", "second", "third", "third", "fourth", "fourth"} {
		downStream <- v
	}
	close(downStream)
}

func removeIdentical(upstream, downstream chan string) {
	// isFirstValue bool
	var previous string

	for val := range upstream {
		if previous == "" {
			previous = val
			downstream <- val
		}
		if previous != val {
			previous = val
			downstream <- val
		}
	}
	close(downstream)

}

func finalStage(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}
