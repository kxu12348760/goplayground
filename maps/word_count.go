package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wc_map := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		count, ok := wc_map[word]
		if ok {
			wc_map[word] = count + 1
		} else {
			wc_map[word] = 1
		}
	}
	return wc_map
}

func main() {
	wc.Test(WordCount)
}
