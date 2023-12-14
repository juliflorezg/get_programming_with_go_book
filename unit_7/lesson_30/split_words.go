package main

import "strings"

func splitSentence(ups, dwn chan string) {
	for str := range ups {
		words := strings.Fields(str)
		for _, w := range words {
			dwn <- w
		}
	}

	close(dwn)
}
