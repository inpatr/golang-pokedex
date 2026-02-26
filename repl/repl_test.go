package repl_test

import (
	"testing"
	"github.com/inpatr/golang-pokedex/repl"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input: "Lorem Ipsum dolor sit amet",
			expected: []string{"Lorem", "Ipsum", "dolor", "sit", "amet"},
		},
		{
			input: "Kartoffelgratin",
			expected: []string{"Kartoffelgratin"},
		},
	}

	for _, c := range cases {
		actual := repl.CleanInput(c.input)
		
		if len(actual) != len(c.expected) {
			t.Errorf("unexpected length...")
		}
		
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Got: %s,\n Want: %s", word, expectedWord)
			}
		}
	}
}
