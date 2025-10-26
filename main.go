package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("./words.txt")

	wordCount := 0

	for _, value := range data {
		if value == ' ' {
			wordCount++
		}
	}

	wordCount++
	fmt.Println("total words in the text file are", wordCount)
}
