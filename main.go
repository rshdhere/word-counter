package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("./words.txt")

	result := wordCount(data)
	fmt.Println(result)
}

func wordCount(data []byte) int {
	words := strings.Fields(string(data))
	return len(words)
}
