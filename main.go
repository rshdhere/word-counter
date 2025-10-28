package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("./words.txt")

	result := WordCount(data)
	fmt.Println(result)
}

func WordCount(data []byte) int {
	words := bytes.Fields(data)
	return len(words)
}
