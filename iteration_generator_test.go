package iteration_names

import (
	"strings"
	"testing"
)

var tests = []struct {
	letter  string
	success bool
}{
	{"A", true},
	{"B", true},
	{"C", true},
	{"Asdf", false},
	{"a", false},
	{"123", false},
}

func TestIterationNameGenerator(t *testing.T) {
	for _, test := range tests {
		iterationName, err := iterationNameGenerator(test.letter)

		if (err != nil && test.success) || (err == nil && !test.success) {
			t.Fatal(err)
		}

		if err != nil {
			continue
		}

		splitName := strings.Split(iterationName, " ")
		for _, word := range splitName {
			if word[0:1] != test.letter {
				t.Fatalf("Expected adjective starting with %s. Got %s.", test.letter, word[0:1])
			}
		}
	}
}
