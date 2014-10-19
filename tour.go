package main

import (
    "code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
    words := strings.Fields(s)
    counts := make(map[string]int)
	for _, word := range words {
		count, ok := counts[word]
		if ok {
			counts[word] = count + 1
		} else {
			counts[word] = 1
		}
	}
	return counts
}

func main() {
    wc.Test(WordCount)
}