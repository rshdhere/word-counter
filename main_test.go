package main

import "testing"

func TestCountWords(t *testing.T) {
	input := "one two three four five"
	wants := 5

	result := wordCount([]byte(input))

	if result != wants {
		t.Fail()
	}
}
