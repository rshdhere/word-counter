package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("./words.txt")

	result := wordCount(data)
	fmt.Println(result)
}

func wordCount(data []byte) int {
	wordCount := 0

	wasSpace := true
	for _, x := range data {
		isSpace := (x == ' ' || x == '\n' || x == '\t')
		if wasSpace && !isSpace {
			wordCount++
		}
		wasSpace = isSpace
	}

	return wordCount
}
