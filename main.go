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

	for _, value := range data {
		if value == ' ' {
			wordCount++
		}
	}

	wordCount++
	return wordCount
}
