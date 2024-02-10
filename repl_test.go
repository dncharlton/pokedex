package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string {
				"hello",
				"world",
			},
		},{
			input: "HeLLo woRLd",
			expected: []string {
				"hello",
				"world",
			},
		},
	}

	for _, cs := range cases{
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("The lengths are not equal")
			continue
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]
			if actualWord != expectedWord {
				t.Errorf("words are not the same")
				continue 
			}
		}
	}
}