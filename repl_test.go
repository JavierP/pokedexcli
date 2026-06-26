package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "  Hello  World  ",
			expected: []string{"hello", "world"},
		},
		//add more cases here
	}


	for _, c := range cases {
		actual := cleanInput(c.input)
		//check the length of the actual slice against the expected slice 
		//if they don't match , use t.Errorf to print an error message 
		//and fail the test
		if len(actual) != len(c.expected){
			t.Errorf("Length do not match: %v vs %v", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice 
			// if they don't match, use t.Errorf to print an error  message
			// and fail the test
			if word != expectedWord {
				t.Errorf("Words do not match: %v and %v", word, expectedWord)
			}
		}
	}
}


