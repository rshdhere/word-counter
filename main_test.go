package main_test

import (
	"testing"

	counter "github.com/rshdhere/counter"
)

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

		{
			name:  "At new lines",
			input: "one two three\nfour five",
			wants: 5,
		},

		{
			name:  "At multiple spaces",
			input: "This is a sentance.    This is another sentance",
			wants: 8,
		},

		{
			name:  "At prefixed multiple spaces",
			input: "   Hello",
			wants: 1,
		},

		{
			name:  "At suffixed multiple spaces",
			input: "Hello   ",
			wants: 1,
		},

		{
			name:  "At tab charachter in code",
			input: "Hello\tWorld\n",
			wants: 2,
		},
	}

	for _, value := range testCases {
		t.Run(value.name, func(t *testing.T) {
			result := counter.WordCount([]byte(value.input))

			if result != value.wants {
				t.Logf("%s expected: %d received: %d", value.name, value.wants, result)
				t.Fail()
			}
		})
	}
}
