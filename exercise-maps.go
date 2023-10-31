package main

import (
	"strings"
)

// https://go-tour-jp.appspot.com/moretypes/23
func WordCount(s string) map[string]int {
	result := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		result[word] += 1
	}
	return result
}

func main() {
	// wc.Test(WordCount)
}
