package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	freq := make(map[string]int)
	for _, word := range strings.Split(s, " ") {
		//fmt.Println(word)
		word = strings.Split(word, ".")[0]
		_, ok := freq[word]
		if ok == false {
			freq[word] = freq[word] + 1
		} else {
			freq[word] = 1
		}
	}
	return freq
}

func main() {
	fmt.Println("Testing wc")
	wc.Test(WordCount)
}
