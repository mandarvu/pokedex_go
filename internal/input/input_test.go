package input

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Normal Input",
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			name:     "Multiple words without extra whitespaces",
			input:    "charmender charizard pickachu",
			expected: []string{"charmender", "charizard", "pickachu"},
		},
		{
			name:     "Multiple words with extra spaces",
			input:    "     charmender      charizard       pickachu",
			expected: []string{"charmender", "charizard", "pickachu"},
		},
		{
			name:     "Multiple words with extra spaces and mixed case",
			input:    "  chArmEndEr      cHarIzard       PiCkAcHu",
			expected: []string{"charmender", "charizard", "pickachu"},
		},
		{
			name:     "Single word in uppercase",
			input:    "HELLO",
			expected: []string{"hello"},
		},
	}

	for _, tc := range testCases {
		actual := cleanInput(tc.input)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("TEST: %s FAILED:\ncleanInput(%q) = %q; expected %q", tc.name, tc.input, actual, tc.expected)
		}
	}
}
