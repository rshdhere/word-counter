package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	fileName := "./words.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("failed to read file: %v", err)
	}

	result := WordCount(data)
	fmt.Println(result)
}

func WordCount(data []byte) int {
	words := bytes.Fields(data)
	return len(words)
}
