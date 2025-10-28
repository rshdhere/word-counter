package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "./words.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("failed to open file: %v", err)
	}

	data, err := io.ReadAll(file)
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
