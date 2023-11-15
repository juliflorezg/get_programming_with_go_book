package main

import (
	"fmt"
	"strings"
)

func words() map[string]int {
	var text = "As far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade, and far overhead the multiple transparency of huge leaves filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever he felt able he ran again; the ground continued soft and springy, covered with the same resilient weed which was the first thing his hands had touched in Malacandra. Once or twice a small red creature scuttled across his path, but otherwise there seemed to be no life stirring in the wood; nothing to fear—except the fact of wandering unprovisioned and alone in a forest of unknown vegetation thousands or millions of miles beyond the reach or knowledge of man."

	text = strings.ReplaceAll(text, "-", " ")
	text = strings.ReplaceAll(text, "_`", " ")
	text = strings.ReplaceAll(text, "—", " ")
	text = strings.ReplaceAll(text, ",", " ")
	text = strings.ReplaceAll(text, ".", " ")
	text = strings.ReplaceAll(text, ";", " ")
	text = strings.ReplaceAll(text, ":", " ")
	text = strings.ReplaceAll(text, "  ", " ")
	text = strings.ToLower(text)

	textSlice := strings.Fields(text)
	wordsCount := make(map[string]int)

	fmt.Println(text)
	fmt.Println(textSlice)

	for _, word := range textSlice {
		wordsCount[word]++
	}

	for w, c := range wordsCount {
		if c > 1 {
			fmt.Println(w, ":", c)
		}
	}

	return wordsCount

}
