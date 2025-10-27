package main

import "testing"

func TestCountWords(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "At 05 words",
			input: "one two three four five",
			wants: 5,
		},
		{
			name:  "At empty file",
			input: "",
			wants: 0,
		},
		{
			name:  "At single space",
			input: " ",
			wants: 0,
		},
	}

	for _, value := range testCases {
		t.Run(value.name, func(t *testing.T) {
			result := wordCount([]byte(value.input))

			if result != value.wants {
				t.Logf("%s expected: %d received: %d", value.name, value.wants, result)
				t.Fail()
			}
		})
	}
}
